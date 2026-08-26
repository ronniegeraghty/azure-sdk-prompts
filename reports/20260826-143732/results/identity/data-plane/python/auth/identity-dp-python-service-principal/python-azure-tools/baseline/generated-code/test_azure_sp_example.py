import os
import unittest
from unittest.mock import MagicMock, patch

from azure.mgmt.resource import ResourceManagementClient

from azure_sp_example import (
    AzureSettings,
    ConfigurationError,
    create_credential,
    create_resource_client,
    list_resource_groups,
)


class AzureServicePrincipalExampleTests(unittest.TestCase):
    def test_missing_environment_is_rejected(self) -> None:
        with patch.dict(os.environ, {}, clear=True):
            with self.assertRaisesRegex(
                ConfigurationError, "AZURE_TENANT_ID"
            ):
                AzureSettings.from_environment()

    def test_credential_and_client_are_created_without_network_access(self) -> None:
        settings = AzureSettings("tenant", "client", "secret", "subscription")

        credential = create_credential(settings)
        client = create_resource_client(settings, credential)

        self.assertIsInstance(client, ResourceManagementClient)
        client.close()
        credential.close()

    def test_resource_groups_are_read_through_sdk_client(self) -> None:
        client = MagicMock()
        group = MagicMock()
        group.name = "example-rg"
        client.resource_groups.list.return_value = [group]

        with patch("builtins.print") as output:
            list_resource_groups(client)

        client.resource_groups.list.assert_called_once_with()
        output.assert_any_call("- example-rg")


if __name__ == "__main__":
    unittest.main()
