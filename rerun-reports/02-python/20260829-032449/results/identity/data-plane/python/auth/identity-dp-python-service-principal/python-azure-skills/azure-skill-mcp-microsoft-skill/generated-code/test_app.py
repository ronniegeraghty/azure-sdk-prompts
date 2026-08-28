from __future__ import annotations

import os
import unittest
from unittest.mock import MagicMock, patch

from azure.core.exceptions import ClientAuthenticationError

import app


class AppTests(unittest.TestCase):
    def setUp(self) -> None:
        self.config = app.AzureConfig(
            tenant_id="tenant-id",
            client_id="client-id",
            client_secret="client-secret",
            storage_account_url="https://example.blob.core.windows.net",
        )

    @patch.dict(os.environ, {}, clear=True)
    @patch("app.load_dotenv")
    def test_load_config_reports_all_missing_values(
        self, load_dotenv: MagicMock
    ) -> None:
        with self.assertRaises(app.ConfigurationError) as context:
            app.load_config()

        load_dotenv.assert_called_once_with(override=False)
        self.assertIn("AZURE_TENANT_ID", str(context.exception))
        self.assertIn("AZURE_CLIENT_SECRET", str(context.exception))

    @patch("app.BlobServiceClient")
    @patch("app.ClientSecretCredential")
    def test_lists_containers_with_service_principal(
        self,
        credential_class: MagicMock,
        blob_service_class: MagicMock,
    ) -> None:
        client = blob_service_class.return_value.__enter__.return_value
        client.list_containers.return_value = [{"name": "one"}, {"name": "two"}]

        names = app.list_storage_containers(self.config)

        credential_class.assert_called_once_with(
            tenant_id="tenant-id",
            client_id="client-id",
            client_secret="client-secret",
        )
        blob_service_class.assert_called_once_with(
            account_url="https://example.blob.core.windows.net",
            credential=credential_class.return_value.__enter__.return_value,
        )
        self.assertEqual(names, ["one", "two"])

    @patch("app.list_storage_containers")
    @patch("app.load_config")
    def test_main_returns_authentication_error_exit_code(
        self,
        load_config: MagicMock,
        list_containers: MagicMock,
    ) -> None:
        load_config.return_value = self.config
        list_containers.side_effect = ClientAuthenticationError(
            message="invalid client secret"
        )

        with self.assertLogs(app.logger, level="ERROR"):
            exit_code = app.main()

        self.assertEqual(exit_code, 3)


if __name__ == "__main__":
    unittest.main()
