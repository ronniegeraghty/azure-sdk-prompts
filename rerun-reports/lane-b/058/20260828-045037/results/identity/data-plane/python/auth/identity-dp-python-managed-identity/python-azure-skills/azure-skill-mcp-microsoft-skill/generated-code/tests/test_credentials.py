import unittest
from unittest.mock import patch

from managed_identity_demo.credentials import CredentialMode, create_credential


class CredentialTests(unittest.TestCase):
    @patch("managed_identity_demo.credentials.ManagedIdentityCredential")
    def test_system_assigned_uses_no_identity_selector(self, credential_class):
        create_credential(CredentialMode.SYSTEM_ASSIGNED)

        credential_class.assert_called_once_with()

    @patch("managed_identity_demo.credentials.ManagedIdentityCredential")
    def test_user_assigned_selects_client_id(self, credential_class):
        create_credential(CredentialMode.USER_ASSIGNED, client_id="identity-client-id")

        credential_class.assert_called_once_with(client_id="identity-client-id")

    def test_user_assigned_requires_client_id(self):
        with self.assertRaisesRegex(ValueError, "client ID is required"):
            create_credential(CredentialMode.USER_ASSIGNED)

    @patch("managed_identity_demo.credentials.DefaultAzureCredential")
    def test_local_mode_skips_managed_identity_probe(self, credential_class):
        create_credential(CredentialMode.LOCAL)

        credential_class.assert_called_once_with(
            exclude_managed_identity_credential=True,
            exclude_interactive_browser_credential=True,
        )

    @patch("managed_identity_demo.credentials.DefaultAzureCredential")
    def test_auto_user_passes_managed_identity_client_id(self, credential_class):
        create_credential(CredentialMode.AUTO_USER, client_id="identity-client-id")

        credential_class.assert_called_once_with(
            managed_identity_client_id="identity-client-id",
            exclude_interactive_browser_credential=True,
        )


if __name__ == "__main__":
    unittest.main()
