# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models import util


class ModelDeploymentIstioConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, namespace: str=None, service_name: str=None):  # noqa: E501
        """ModelDeploymentIstioConfig - a model defined in Swagger

        :param namespace: The namespace of this ModelDeploymentIstioConfig.  # noqa: E501
        :type namespace: str
        :param service_name: The service_name of this ModelDeploymentIstioConfig.  # noqa: E501
        :type service_name: str
        """
        self.swagger_types = {
            'namespace': str,
            'service_name': str
        }

        self.attribute_map = {
            'namespace': 'namespace',
            'service_name': 'serviceName'
        }

        self._namespace = namespace
        self._service_name = service_name

    @classmethod
    def from_dict(cls, dikt) -> 'ModelDeploymentIstioConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The ModelDeploymentIstioConfig of this ModelDeploymentIstioConfig.  # noqa: E501
        :rtype: ModelDeploymentIstioConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def namespace(self) -> str:
        """Gets the namespace of this ModelDeploymentIstioConfig.

        Istio ingress gateway namespace  # noqa: E501

        :return: The namespace of this ModelDeploymentIstioConfig.
        :rtype: str
        """
        return self._namespace

    @namespace.setter
    def namespace(self, namespace: str):
        """Sets the namespace of this ModelDeploymentIstioConfig.

        Istio ingress gateway namespace  # noqa: E501

        :param namespace: The namespace of this ModelDeploymentIstioConfig.
        :type namespace: str
        """

        self._namespace = namespace

    @property
    def service_name(self) -> str:
        """Gets the service_name of this ModelDeploymentIstioConfig.

        Istio ingress gateway service name  # noqa: E501

        :return: The service_name of this ModelDeploymentIstioConfig.
        :rtype: str
        """
        return self._service_name

    @service_name.setter
    def service_name(self, service_name: str):
        """Sets the service_name of this ModelDeploymentIstioConfig.

        Istio ingress gateway service name  # noqa: E501

        :param service_name: The service_name of this ModelDeploymentIstioConfig.
        :type service_name: str
        """

        self._service_name = service_name
