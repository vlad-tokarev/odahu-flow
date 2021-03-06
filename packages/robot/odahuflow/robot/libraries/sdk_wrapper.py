from odahuflow.sdk.clients.api_aggregated import parse_resources_file_with_one_item
from odahuflow.sdk.clients.configuration import ConfigurationClient
from odahuflow.sdk.clients.connection import ConnectionClient
from odahuflow.sdk.clients.deployment import ModelDeploymentClient
from odahuflow.sdk.clients.model import ModelClient
from odahuflow.sdk.clients.packaging import ModelPackagingClient
from odahuflow.sdk.clients.packaging_integration import PackagingIntegrationClient
from odahuflow.sdk.clients.route import ModelRouteClient
from odahuflow.sdk.clients.toolchain_integration import ToolchainIntegrationClient
from odahuflow.sdk.clients.training import ModelTrainingClient

from odahuflow.sdk.clients.api import RemoteAPIClient

remote_api = RemoteAPIClient()


class Configuration:

    @staticmethod
    def config_get():
        return ConfigurationClient().get()

    @staticmethod
    def config_put(payload_file):
        return ConfigurationClient().put(payload_file)


class Connection:

    @staticmethod
    def connection_get():
        return ConnectionClient().get_all()

    @staticmethod
    def connection_get_id(conn_id: str):
        return ConnectionClient().get(conn_id)

    @staticmethod
    def connection_get_id_decrypted(conn_id: str):
        return ConnectionClient().get_decrypted(conn_id)

    @staticmethod
    def connection_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ConnectionClient().edit(api_object)

    @staticmethod
    def connection_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ConnectionClient().create(api_object)

    @staticmethod
    def connection_delete(conn_id: str):
        return ConnectionClient().delete(conn_id)


class ModelDeployment:

    @staticmethod
    def deployment_get():
        return ModelDeploymentClient().get_all()

    @staticmethod
    def deployment_get_id(dep_id: str):
        return ModelDeploymentClient().get(dep_id)

    @staticmethod
    def deployment_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelDeploymentClient().edit(api_object)

    @staticmethod
    def deployment_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelDeploymentClient().create(api_object)

    @staticmethod
    def deployment_delete(dep_id: str):
        return ModelDeploymentClient().delete(dep_id)


class ModelPackaging:

    @staticmethod
    def packaging_get():
        return ModelPackagingClient().get_all()

    @staticmethod
    def packaging_get_id(pack_id: str):
        return ModelPackagingClient().get(pack_id)

    @staticmethod
    def packaging_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelPackagingClient().edit(api_object)

    @staticmethod
    def packaging_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelPackagingClient().create(api_object)

    @staticmethod
    def packaging_delete(pack_id: str):
        return ModelPackagingClient().delete(pack_id)

    @staticmethod
    def packaging_get_log(pack_id):
        for msg in ModelPackagingClient().log(pack_id, follow=False):
            print(msg)


class ModelTraining:

    @staticmethod
    def training_get():
        return ModelTrainingClient().get_all()

    @staticmethod
    def training_get_id(train_id: str):
        return ModelTrainingClient().get(train_id)

    @staticmethod
    def training_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelTrainingClient().edit(api_object)

    @staticmethod
    def training_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelTrainingClient().create(api_object)

    @staticmethod
    def training_delete(train_id: str):
        return ModelTrainingClient().delete(train_id)

    @staticmethod
    def training_get_log(train_id):
        for msg in ModelTrainingClient().log(train_id, follow=False):
            print(msg)


class ModelRoute:

    @staticmethod
    def route_get():
        return ModelRouteClient().get_all()

    @staticmethod
    def route_get_id(route_id: str):
        return ModelRouteClient().get(route_id)

    @staticmethod
    def route_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelRouteClient().edit(api_object)

    @staticmethod
    def route_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ModelRouteClient().create(api_object)

    @staticmethod
    def route_delete(route_id: str):
        return ModelRouteClient().delete(route_id)


class Packager:

    @staticmethod
    def packager_get():
        return PackagingIntegrationClient().get_all()

    @staticmethod
    def packager_get_id(pi_id: str):
        return PackagingIntegrationClient().get(pi_id)

    @staticmethod
    def packager_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return PackagingIntegrationClient().edit(api_object)

    @staticmethod
    def packager_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return PackagingIntegrationClient().create(api_object)

    @staticmethod
    def packager_delete(pi_id: str):
        return PackagingIntegrationClient().delete(pi_id)


class Toolchain:

    @staticmethod
    def toolchain_get():
        return ToolchainIntegrationClient().get_all()

    @staticmethod
    def toolchain_get_id(ti_id: str):
        return ToolchainIntegrationClient().get(ti_id)

    @staticmethod
    def toolchain_put(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ToolchainIntegrationClient().edit(api_object)

    @staticmethod
    def toolchain_post(payload_file):
        api_object = parse_resources_file_with_one_item(payload_file).resource
        return ToolchainIntegrationClient().create(api_object)

    @staticmethod
    def toolchain_delete(ti_id: str):
        return ToolchainIntegrationClient().delete(ti_id)


class Model:

    @staticmethod
    def model_get():
        return ModelClient().info()

    @staticmethod
    def model_post(**parameters):
        return ModelClient().invoke(**parameters)
