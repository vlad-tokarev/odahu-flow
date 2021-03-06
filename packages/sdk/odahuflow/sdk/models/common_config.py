# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models.external_url import ExternalUrl  # noqa: F401,E501
from odahuflow.sdk.models import util


class CommonConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, external_urls: List[ExternalUrl]=None, resource_gpu_name: str=None):  # noqa: E501
        """CommonConfig - a model defined in Swagger

        :param external_urls: The external_urls of this CommonConfig.  # noqa: E501
        :type external_urls: List[ExternalUrl]
        :param resource_gpu_name: The resource_gpu_name of this CommonConfig.  # noqa: E501
        :type resource_gpu_name: str
        """
        self.swagger_types = {
            'external_urls': List[ExternalUrl],
            'resource_gpu_name': str
        }

        self.attribute_map = {
            'external_urls': 'externalUrls',
            'resource_gpu_name': 'resourceGpuName'
        }

        self._external_urls = external_urls
        self._resource_gpu_name = resource_gpu_name

    @classmethod
    def from_dict(cls, dikt) -> 'CommonConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The CommonConfig of this CommonConfig.  # noqa: E501
        :rtype: CommonConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def external_urls(self) -> List[ExternalUrl]:
        """Gets the external_urls of this CommonConfig.

        The collection of external urls, for example: metrics, edge, service catalog and so on  # noqa: E501

        :return: The external_urls of this CommonConfig.
        :rtype: List[ExternalUrl]
        """
        return self._external_urls

    @external_urls.setter
    def external_urls(self, external_urls: List[ExternalUrl]):
        """Sets the external_urls of this CommonConfig.

        The collection of external urls, for example: metrics, edge, service catalog and so on  # noqa: E501

        :param external_urls: The external_urls of this CommonConfig.
        :type external_urls: List[ExternalUrl]
        """

        self._external_urls = external_urls

    @property
    def resource_gpu_name(self) -> str:
        """Gets the resource_gpu_name of this CommonConfig.

        Kubernetes can consume the GPU resource in the <vendor>.com/gpu format. For example, amd.com/gpu or nvidia.com/gpu.  # noqa: E501

        :return: The resource_gpu_name of this CommonConfig.
        :rtype: str
        """
        return self._resource_gpu_name

    @resource_gpu_name.setter
    def resource_gpu_name(self, resource_gpu_name: str):
        """Sets the resource_gpu_name of this CommonConfig.

        Kubernetes can consume the GPU resource in the <vendor>.com/gpu format. For example, amd.com/gpu or nvidia.com/gpu.  # noqa: E501

        :param resource_gpu_name: The resource_gpu_name of this CommonConfig.
        :type resource_gpu_name: str
        """

        self._resource_gpu_name = resource_gpu_name
