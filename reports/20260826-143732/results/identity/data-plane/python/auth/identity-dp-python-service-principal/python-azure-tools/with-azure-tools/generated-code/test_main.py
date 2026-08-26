import os
import unittest
from unittest.mock import MagicMock, patch

from azure.core.exceptions import ClientAuthenticationError

import main


class SettingsTests(unittest.TestCase):
    @patch.dict(os.environ, {}, clear=True)
    def test_missing_environment_variables_are_reported(self) -> None:
        with self.assertRaisesRegex(
            main.ConfigurationError, "AZURE_TENANT_ID"
        ):
            main.Settings.from_environment()

    @patch.dict(
        os.environ,
        {
            "AZURE_TENANT_ID": "tenant",
            "AZURE_CLIENT_ID": "client",
            "AZURE_CLIENT_SECRET": "secret",
            "AZURE_STORAGE_ACCOUNT_URL": "http://account.example",
        },
        clear=True,
    )
    def test_storage_url_must_use_https(self) -> None:
        with self.assertRaisesRegex(main.ConfigurationError, "HTTPS"):
            main.Settings.from_environment()


class AzureClientTests(unittest.TestCase):
    def setUp(self) -> None:
        self.settings = main.Settings(
            tenant_id="tenant",
            client_id="client",
            client_secret="secret",
            storage_account_url="https://account.blob.core.windows.net",
        )

    @patch("main.BlobServiceClient")
    @patch("main.ClientSecretCredential")
    def test_credential_is_used_by_blob_client(
        self,
        credential_type: MagicMock,
        blob_client_type: MagicMock,
    ) -> None:
        credential = credential_type.return_value.__enter__.return_value
        blob_client = blob_client_type.return_value.__enter__.return_value
        blob_client.list_containers.return_value = iter(
            [{"name": "one"}, {"name": "two"}]
        )

        names = main.list_container_names(self.settings)

        credential_type.assert_called_once_with(
            tenant_id="tenant",
            client_id="client",
            client_secret="secret",
        )
        credential.get_token.assert_called_once_with(main.STORAGE_SCOPE)
        blob_client_type.assert_called_once_with(
            account_url=self.settings.storage_account_url,
            credential=credential,
        )
        self.assertEqual(names, ["one", "two"])

    @patch("main.load_dotenv")
    @patch("main.Settings.from_environment")
    @patch("main.list_container_names")
    def test_authentication_failure_returns_distinct_exit_code(
        self,
        list_names: MagicMock,
        from_environment: MagicMock,
        _load_dotenv: MagicMock,
    ) -> None:
        from_environment.return_value = self.settings
        list_names.side_effect = ClientAuthenticationError("invalid secret")

        with self.assertLogs(main.logger, level="ERROR") as logs:
            exit_code = main.run()

        self.assertEqual(exit_code, 3)
        self.assertIn("authentication failed", logs.output[0].lower())
        self.assertNotIn("invalid secret", logs.output[0])


if __name__ == "__main__":
    unittest.main()
