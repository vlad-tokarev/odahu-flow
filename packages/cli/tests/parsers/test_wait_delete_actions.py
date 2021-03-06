#
#    Copyright 2020 EPAM Systems
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
import http
import json
import pathlib
import typing

import pytest
from click.testing import CliRunner
from odahuflow.sdk.clients.api import WrongHttpStatusCode
from pytest_mock import MockFixture

from .data import ENTITY_ID, EntityTestData, generate_entities_for_test, DEPLOYMENT

ENTITY_TEST_DATA: typing.Dict[str, EntityTestData] = {
    "deployment": DEPLOYMENT,
}


def pytest_generate_tests(metafunc):
    generate_entities_for_test(metafunc, list(ENTITY_TEST_DATA.keys()))


@pytest.fixture
def entity_test_data(request) -> EntityTestData:
    return ENTITY_TEST_DATA[request.param]


def test_delete_by_file(tmp_path: pathlib.Path, mocker: MockFixture, cli_runner: CliRunner,
                        entity_test_data: EntityTestData):
    message = "was deleted"
    entity_file = tmp_path / "entity.yaml"
    entity_file.write_text(
        json.dumps(
            {**entity_test_data.entity.to_dict(), **{'kind': entity_test_data.kind}}))
    delete_client_mock = mocker.patch.object(entity_test_data.entity_client.__class__,
                                             'delete',
                                             return_value=message)
    get_client_mock = mocker.patch.object(entity_test_data.entity_client.__class__,
                                          'get',
                                          side_effect=WrongHttpStatusCode(status_code=http.HTTPStatus.NOT_FOUND))

    result = cli_runner.invoke(entity_test_data.click_group,
                               ['delete', '-f', entity_file],
                               obj=entity_test_data.entity_client)

    assert result.exit_code == 0
    assert message in result.stdout
    delete_client_mock.assert_called_once_with(entity_test_data.entity.id)
    get_client_mock.assert_called_once()


def test_delete_by_id(mocker: MockFixture, cli_runner: CliRunner,
                      entity_test_data: EntityTestData):
    message = "was deleted"
    delete_client_mock = mocker.patch.object(entity_test_data.entity_client.__class__,
                                             'delete',
                                             return_value=message)
    get_client_mock = mocker.patch.object(entity_test_data.entity_client.__class__,
                                          'get',
                                          side_effect=WrongHttpStatusCode(status_code=http.HTTPStatus.NOT_FOUND))

    result = cli_runner.invoke(entity_test_data.click_group,
                               ['delete', '--id', ENTITY_ID],
                               obj=entity_test_data.entity_client)

    assert result.exit_code == 0
    assert message in result.stdout
    delete_client_mock.assert_called_once_with(entity_test_data.entity.id)
    get_client_mock.assert_called_once()
