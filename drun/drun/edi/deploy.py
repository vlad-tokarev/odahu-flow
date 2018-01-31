#
#    Copyright 2017 EPAM Systems
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
#
"""
Deploy logic for DRun
"""

import os
import logging

import drun.utils
from drun.utils import Colors, ExternalFileReader, normalize_name_to_dns_1123
import drun.model.io
import drun.const.env
import drun.const.headers
import drun.external.grafana
import drun.external.edi
import drun.containers.docker
import drun.containers.k8s

import docker
import docker.errors

LOGGER = logging.getLogger('deploy')
VALID_SERVING_WORKERS = drun.containers.docker.VALID_SERVING_WORKERS


def build_model(args):
    """
    Build model

    :param args: command arguments
    :type args: :py:class:`argparse.Namespace`
    :return: :py:class:`docker.model.Image` docker image
    """
    client = drun.containers.docker.build_docker_client(args)

    with ExternalFileReader(args.model_file) as external_reader:
        if not os.path.exists(external_reader.path):
            raise Exception('Cannot find model file: %s' % external_reader.path)

        with drun.model.io.ModelContainer(external_reader.path, do_not_load_model=True) as container:
            model_id = container.get('model.id', None)
            if args.model_id:
                model_id = args.model_id

        if not model_id:
            raise Exception('Cannot get model id (not setted in container and not setted in arguments)')

        image_labels = drun.containers.docker.generate_docker_labels_for_image(external_reader.path, model_id, args)

        base_docker_image = args.base_docker_image
        if not base_docker_image:
            base_docker_image = 'drun/base-python-image:latest'

        image = drun.containers.docker.build_docker_image(
            client,
            base_docker_image,
            model_id,
            external_reader.path,
            image_labels,
            args.python_package,
            args.docker_image_tag,
            args.serving
        )

        LOGGER.info('Built image: %s with python package: %s' % (image, args.python_package))

        print('Successfully created docker image %s for model %s' % (image.short_id, model_id))

        if args.push_to_registry:
            uri = args.push_to_registry  # type: str
            tag_start_position = uri.rfind(':')
            slash_latest_position = uri.rfind('/')

            if 0 < tag_start_position < slash_latest_position and slash_latest_position > 0:
                repository = uri
                tag = 'latest'

            print('Tagging image %s for model %s as %s' % (image.short_id, model_id, args.push_to_registry))
            image.tag(args.push_to_registry)
            print('Successfully tagged image %s for model %s as %s' % (image.short_id, model_id, args.push_to_registry))
            client.images.push(args.push_to_registry)
            print('Successfully pushed image %s for model %s to %s' % (image.short_id, model_id, args.push_to_registry))

        return image


def inspect_kubernetes(args):
    """
    Inspect kubernetes

    :param args: command arguments with .namespace
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    edi_client = drun.external.edi.build_client(args)
    model_deployments = edi_client.inspect()

    print('%sModel deployments:%s' % (Colors.BOLD, Colors.ENDC))

    for deployment in model_deployments:
        if deployment.status == 'ok':
            line_color = Colors.OKGREEN
        elif deployment.status == 'warning':
            line_color = Colors.WARNING
        else:
            line_color = Colors.FAIL

        arguments = (
            line_color, Colors.ENDC,
            Colors.UNDERLINE, deployment.model, Colors.ENDC,
            deployment.image, deployment.version,
            line_color, deployment.ready_replicas, deployment.scale,
            Colors.ENDC
        )
        print('%s*%s %s%s%s %s (version: %s) - %s%s / %d pods ready%s' % arguments)

    if len(model_deployments) == 0:
        print('%s-- cannot find any model deployments --%s' % (Colors.WARNING, Colors.ENDC))


def undeploy_kubernetes(args):
    """
    Undeploy model to kubernetes

    :param args: command arguments with .model_id, .namespace
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    edi_client = drun.external.edi.build_client(args)
    edi_client.undeploy(args.model_id, args.grace_period)


def scale_kubernetes(args):
    """
    Scale model instances

    :param args: command arguments with .model_id, .namespace and .scale
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    edi_client = drun.external.edi.build_client(args)
    edi_client.scale(args.model_id, args.scale)


def deploy_kubernetes(args):
    """
    Deploy kubernetes model

    :param args: command arguments with .model_id, .namespace and .scale
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    edi_client = drun.external.edi.build_client(args)
    edi_client.deploy(args.image, args.scale, args.image_for_k8s)


def deploy_model(args):
    """
    Deploy model to docker host

    :param args: command arguments with .model_id, .model_file, .docker_network
    :type args: :py:class:`argparse.Namespace`
    :return: :py:class:`docker.model.Container` new instance
    """
    client = drun.containers.docker.build_docker_client(args)
    network_id = drun.containers.docker.find_network(client, args)
    grafana_client = drun.external.grafana.build_grafana_client(args)

    if args.model_id and args.docker_image:
        print('Use only --model-id or --docker-image')
        exit(1)
    elif not args.model_id and not args.docker_image:
        print('Use with --model-id or --docker-image')
        exit(1)

    current_containers = drun.containers.docker.get_stack_containers_and_images(client, network_id)

    if args.model_id:
        for image in current_containers['model_images']:
            model_name = image.labels.get('com.epam.drun.model.id', None)
            if model_name == args.model_id:
                image = image
                model_id = model_name
                break
        else:
            raise Exception('Cannot found image for model_id = %s' % (args.model_id,))
    elif args.docker_image:
        try:
            image = client.images.get(args.docker_image)
        except docker.errors.ImageNotFound:
            print('Cannot find %s locally. Pulling' % args.docker_image)
            image = client.images.pull(args.docker_image)

        model_id = image.labels.get('com.epam.drun.model.id', None)
        if not model_id:
            raise Exception('Cannot detect model_id in image')
    else:
        raise Exception('Provide model-id or docker-image')

    # Detect current existing containers with models, stop and remove them
    LOGGER.info('Founding containers with model_id=%s' % model_id)

    for container in current_containers['models']:
        model_name = container.labels.get('com.epam.drun.model.id', None)
        if model_name == model_id:
            LOGGER.info('Stopping container #%s' % container.short_id)
            container.stop()
            LOGGER.info('Removing container #%s' % container.short_id)
            container.remove()

    container_labels = drun.containers.docker.generate_docker_labels_for_container(image)

    ports = {}
    if args.expose_model_port:
        exposing_port = args.expose_model_port
        ports['%d/tcp' % os.getenv(*drun.const.env.LEGION_PORT)] = exposing_port

    LOGGER.info('Starting container with image #%s for model %s' % (image.short_id, model_id))
    container = client.containers.run(image,
                                      network=network_id,
                                      stdout=True,
                                      stderr=True,
                                      detach=True,
                                      ports=ports,
                                      labels=container_labels)

    LOGGER.info('Creating Grafana dashboard for model %s' % (model_id,))
    grafana_client.create_dashboard_for_model_by_labels(container_labels)

    print('Successfully created docker container %s for model %s' % (container.short_id, model_id))
    return container


def undeploy_model(args):
    """
    Undeploy model from Docker Host

    :param args: command arguments
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    client = drun.containers.docker.build_docker_client(args)
    network_id = drun.containers.docker.find_network(client, args)
    grafana_client = drun.external.grafana.build_grafana_client(args)

    current_containers = drun.containers.docker.get_stack_containers_and_images(client, network_id)

    for container in current_containers['models']:
        model_name = container.labels.get('com.epam.drun.model.id', None)
        if model_name == args.model_id:
            target_container = container
            break
    else:
        raise Exception('Cannot found container for model_id = %s' % (args.model_id,))

    LOGGER.info('Stopping container #%s' % target_container.short_id)
    target_container.stop()
    LOGGER.info('Removing container #%s' % target_container.short_id)
    target_container.remove()
    LOGGER.info('Removing Grafana dashboard for model %s' % (args.model_id,))
    grafana_client.remove_dashboard_for_model(args.model_id)

    print('Successfully undeployed model %s' % (args.model_id,))


def inspect(args):
    """
    Print information about current containers / images state

    :param args: command arguments
    :type args: :py:class:`argparse.Namespace`
    :return: None
    """
    client = drun.containers.docker.build_docker_client(args)
    network_id = drun.containers.docker.find_network(client, args)
    containers = drun.containers.docker.get_stack_containers_and_images(client, network_id)

    all_required_containers_is_ok = True

    print('%sServices:%s' % (Colors.BOLD, Colors.ENDC))
    for container in containers['services']:
        is_running = container.status == 'running'
        container_name = container.labels.get('com.epam.drun.container_description', container.image.tags[0])
        container_required = container.labels.get('com.epam.drun.container_required', 'true').lower()
        container_required = container_required in ('1', 'yes', 'true')
        container_status = container.status

        if is_running:
            line_color = Colors.OKGREEN
        elif container_required:
            line_color = Colors.FAIL
            all_required_containers_is_ok = False
        else:
            line_color = Colors.WARNING

        if container.status == 'exited':
            exit_code = container.attrs['State']['ExitCode']
            container_status = 'exited with code %d' % (exit_code,)
        elif container.status == 'running':
            ports = list(container.attrs['NetworkSettings']['Ports'].values())
            ports = [item['HostPort'] for sublist in ports if sublist for item in sublist if item]

            if ports:
                container_status = 'running on ports: %s' % (', '.join(ports),)
        print('%s*%s %s #%s - %s%s%s' % (line_color, Colors.ENDC,
                                         container_name, container.short_id,
                                         line_color, container_status, Colors.ENDC))

    if not containers['services']:
        all_required_containers_is_ok = False
        print('%s-- looks like DRun stack hasn\'t been deployed --%s' % (Colors.FAIL, Colors.ENDC))

    print('%sModel instances:%s' % (Colors.BOLD, Colors.ENDC))
    for container in containers['models']:
        is_running = container.status == 'running'
        line_color = Colors.OKGREEN if is_running else Colors.FAIL
        container_status = '%s #%s' % (container.status, container.short_id)

        model_name = container.labels.get('com.epam.drun.model.id', 'Undefined model ' + ','.join(container.image.tags))
        model_image_id = container.image.short_id
        model_version = container.labels.get('com.epam.drun.model.version', '?')

        print('%s*%s %s%s%s #%s (version: %s) - %s%s%s' % (line_color, Colors.ENDC,
                                                           Colors.UNDERLINE, model_name, Colors.ENDC,
                                                           model_image_id, model_version,
                                                           line_color, container_status, Colors.ENDC))

    if not containers['models']:
        print('%s-- cannot find any model instances --%s' % (Colors.WARNING, Colors.ENDC))

    print('%sModel images:%s' % (Colors.BOLD, Colors.ENDC))
    for image in containers['model_images']:
        model_name = image.labels.get('com.epam.drun.model.id', 'Undefined model')
        model_image_id = image.short_id
        model_version = image.labels.get('com.epam.drun.model.version', '?')
        print('* %s%s%s #%s (version: %s)' % (Colors.UNDERLINE, model_name, Colors.ENDC, model_image_id, model_version))

    if not containers['model_images']:
        print('%s-- cannot find any model images --%s' % (Colors.WARNING, Colors.ENDC))

    if not all_required_containers_is_ok:
        return 2
