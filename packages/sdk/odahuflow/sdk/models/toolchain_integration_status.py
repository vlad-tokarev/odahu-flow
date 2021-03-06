# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models import util


class ToolchainIntegrationStatus(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, created_at: str=None, updated_at: str=None):  # noqa: E501
        """ToolchainIntegrationStatus - a model defined in Swagger

        :param created_at: The created_at of this ToolchainIntegrationStatus.  # noqa: E501
        :type created_at: str
        :param updated_at: The updated_at of this ToolchainIntegrationStatus.  # noqa: E501
        :type updated_at: str
        """
        self.swagger_types = {
            'created_at': str,
            'updated_at': str
        }

        self.attribute_map = {
            'created_at': 'createdAt',
            'updated_at': 'updatedAt'
        }

        self._created_at = created_at
        self._updated_at = updated_at

    @classmethod
    def from_dict(cls, dikt) -> 'ToolchainIntegrationStatus':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The ToolchainIntegrationStatus of this ToolchainIntegrationStatus.  # noqa: E501
        :rtype: ToolchainIntegrationStatus
        """
        return util.deserialize_model(dikt, cls)

    @property
    def created_at(self) -> str:
        """Gets the created_at of this ToolchainIntegrationStatus.


        :return: The created_at of this ToolchainIntegrationStatus.
        :rtype: str
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at: str):
        """Sets the created_at of this ToolchainIntegrationStatus.


        :param created_at: The created_at of this ToolchainIntegrationStatus.
        :type created_at: str
        """

        self._created_at = created_at

    @property
    def updated_at(self) -> str:
        """Gets the updated_at of this ToolchainIntegrationStatus.


        :return: The updated_at of this ToolchainIntegrationStatus.
        :rtype: str
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at: str):
        """Sets the updated_at of this ToolchainIntegrationStatus.


        :param updated_at: The updated_at of this ToolchainIntegrationStatus.
        :type updated_at: str
        """

        self._updated_at = updated_at
