# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models.model_training_spec import ModelTrainingSpec  # noqa: F401,E501
from odahuflow.sdk.models.model_training_status import ModelTrainingStatus  # noqa: F401,E501
from odahuflow.sdk.models import util


class ModelTraining(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, id: str=None, spec: ModelTrainingSpec=None, status: ModelTrainingStatus=None):  # noqa: E501
        """ModelTraining - a model defined in Swagger

        :param id: The id of this ModelTraining.  # noqa: E501
        :type id: str
        :param spec: The spec of this ModelTraining.  # noqa: E501
        :type spec: ModelTrainingSpec
        :param status: The status of this ModelTraining.  # noqa: E501
        :type status: ModelTrainingStatus
        """
        self.swagger_types = {
            'id': str,
            'spec': ModelTrainingSpec,
            'status': ModelTrainingStatus
        }

        self.attribute_map = {
            'id': 'id',
            'spec': 'spec',
            'status': 'status'
        }

        self._id = id
        self._spec = spec
        self._status = status

    @classmethod
    def from_dict(cls, dikt) -> 'ModelTraining':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The ModelTraining of this ModelTraining.  # noqa: E501
        :rtype: ModelTraining
        """
        return util.deserialize_model(dikt, cls)

    @property
    def id(self) -> str:
        """Gets the id of this ModelTraining.

        Model training ID  # noqa: E501

        :return: The id of this ModelTraining.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id: str):
        """Sets the id of this ModelTraining.

        Model training ID  # noqa: E501

        :param id: The id of this ModelTraining.
        :type id: str
        """

        self._id = id

    @property
    def spec(self) -> ModelTrainingSpec:
        """Gets the spec of this ModelTraining.

        Model training specification  # noqa: E501

        :return: The spec of this ModelTraining.
        :rtype: ModelTrainingSpec
        """
        return self._spec

    @spec.setter
    def spec(self, spec: ModelTrainingSpec):
        """Sets the spec of this ModelTraining.

        Model training specification  # noqa: E501

        :param spec: The spec of this ModelTraining.
        :type spec: ModelTrainingSpec
        """

        self._spec = spec

    @property
    def status(self) -> ModelTrainingStatus:
        """Gets the status of this ModelTraining.

        Model training status  # noqa: E501

        :return: The status of this ModelTraining.
        :rtype: ModelTrainingStatus
        """
        return self._status

    @status.setter
    def status(self, status: ModelTrainingStatus):
        """Sets the status of this ModelTraining.

        Model training status  # noqa: E501

        :param status: The status of this ModelTraining.
        :type status: ModelTrainingStatus
        """

        self._status = status
