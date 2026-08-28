from __future__ import annotations

import unittest
from unittest.mock import patch

from managed_identity_demo.auth import (
    AuthSettings,
    ConfigurationError,
    create_credential,
)


class AuthSettingsTests(unittest.TestCase):
    def test_defaults_to_local_system_identity(self) -> None:
        settings = AuthSettings.from_environment({})

        self.assertEqual(settings.environment, "local")
        self.assertEqual(settings.identity_type, "system")

    def test_user_assigned_identity_requires_client_id_in_azure(self) -> None:
        with self.assertRaisesRegex(ConfigurationError, "AZURE_CLIENT_ID"):
            AuthSettings.from_environment(
                {"APP_ENV": "azure", "MANAGED_IDENTITY_TYPE": "user"}
            )

    @patch("managed_identity_demo.auth.DefaultAzureCredential")
    def test_local_credential_skips_managed_identity_probe(self, credential) -> None:
        create_credential(AuthSettings("local", "system"))

        credential.assert_called_once_with(
            exclude_managed_identity_credential=True
        )

    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_system_assigned_credential_has_no_selector(self, credential) -> None:
        create_credential(AuthSettings("azure", "system"))

        credential.assert_called_once_with()

    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_user_assigned_credential_uses_client_id(self, credential) -> None:
        create_credential(AuthSettings("azure", "user", "identity-client-id"))

        credential.assert_called_once_with(client_id="identity-client-id")


if __name__ == "__main__":
    unittest.main()

