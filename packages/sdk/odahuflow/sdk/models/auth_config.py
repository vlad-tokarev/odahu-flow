# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models import util


class AuthConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, api_token: str=None, api_url: str=None, client_id: str=None, client_secret: str=None, oauth_oidc_token_endpoint: str=None):  # noqa: E501
        """AuthConfig - a model defined in Swagger

        :param api_token: The api_token of this AuthConfig.  # noqa: E501
        :type api_token: str
        :param api_url: The api_url of this AuthConfig.  # noqa: E501
        :type api_url: str
        :param client_id: The client_id of this AuthConfig.  # noqa: E501
        :type client_id: str
        :param client_secret: The client_secret of this AuthConfig.  # noqa: E501
        :type client_secret: str
        :param oauth_oidc_token_endpoint: The oauth_oidc_token_endpoint of this AuthConfig.  # noqa: E501
        :type oauth_oidc_token_endpoint: str
        """
        self.swagger_types = {
            'api_token': str,
            'api_url': str,
            'client_id': str,
            'client_secret': str,
            'oauth_oidc_token_endpoint': str
        }

        self.attribute_map = {
            'api_token': 'apiToken',
            'api_url': 'apiUrl',
            'client_id': 'clientId',
            'client_secret': 'clientSecret',
            'oauth_oidc_token_endpoint': 'oauthOidcTokenEndpoint'
        }

        self._api_token = api_token
        self._api_url = api_url
        self._client_id = client_id
        self._client_secret = client_secret
        self._oauth_oidc_token_endpoint = oauth_oidc_token_endpoint

    @classmethod
    def from_dict(cls, dikt) -> 'AuthConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The AuthConfig of this AuthConfig.  # noqa: E501
        :rtype: AuthConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def api_token(self) -> str:
        """Gets the api_token of this AuthConfig.

        It is a mock for the future. Currently, it is always empty.  # noqa: E501

        :return: The api_token of this AuthConfig.
        :rtype: str
        """
        return self._api_token

    @api_token.setter
    def api_token(self, api_token: str):
        """Sets the api_token of this AuthConfig.

        It is a mock for the future. Currently, it is always empty.  # noqa: E501

        :param api_token: The api_token of this AuthConfig.
        :type api_token: str
        """

        self._api_token = api_token

    @property
    def api_url(self) -> str:
        """Gets the api_url of this AuthConfig.

        ODAHU API URL  # noqa: E501

        :return: The api_url of this AuthConfig.
        :rtype: str
        """
        return self._api_url

    @api_url.setter
    def api_url(self, api_url: str):
        """Sets the api_url of this AuthConfig.

        ODAHU API URL  # noqa: E501

        :param api_url: The api_url of this AuthConfig.
        :type api_url: str
        """

        self._api_url = api_url

    @property
    def client_id(self) -> str:
        """Gets the client_id of this AuthConfig.

        OpenID client_id credential for service account  # noqa: E501

        :return: The client_id of this AuthConfig.
        :rtype: str
        """
        return self._client_id

    @client_id.setter
    def client_id(self, client_id: str):
        """Sets the client_id of this AuthConfig.

        OpenID client_id credential for service account  # noqa: E501

        :param client_id: The client_id of this AuthConfig.
        :type client_id: str
        """

        self._client_id = client_id

    @property
    def client_secret(self) -> str:
        """Gets the client_secret of this AuthConfig.

        OpenID client_secret credential for service account  # noqa: E501

        :return: The client_secret of this AuthConfig.
        :rtype: str
        """
        return self._client_secret

    @client_secret.setter
    def client_secret(self, client_secret: str):
        """Sets the client_secret of this AuthConfig.

        OpenID client_secret credential for service account  # noqa: E501

        :param client_secret: The client_secret of this AuthConfig.
        :type client_secret: str
        """

        self._client_secret = client_secret

    @property
    def oauth_oidc_token_endpoint(self) -> str:
        """Gets the oauth_oidc_token_endpoint of this AuthConfig.

        OpenID token url  # noqa: E501

        :return: The oauth_oidc_token_endpoint of this AuthConfig.
        :rtype: str
        """
        return self._oauth_oidc_token_endpoint

    @oauth_oidc_token_endpoint.setter
    def oauth_oidc_token_endpoint(self, oauth_oidc_token_endpoint: str):
        """Sets the oauth_oidc_token_endpoint of this AuthConfig.

        OpenID token url  # noqa: E501

        :param oauth_oidc_token_endpoint: The oauth_oidc_token_endpoint of this AuthConfig.
        :type oauth_oidc_token_endpoint: str
        """

        self._oauth_oidc_token_endpoint = oauth_oidc_token_endpoint
