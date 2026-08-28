from unittest import TestCase
from unittest.mock import patch

from managed_identity_demo.auth import create_credential


class CreateCredentialTests(TestCase):
    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_system_assigned_has_no_identity_selector(self, credential_type):
        credential = create_credential("system")

        self.assertIs(credential, credential_type.return_value)
        credential_type.assert_called_once_with()

    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_user_assigned_uses_client_id(self, credential_type):
        credential = create_credential(
            "user",
            user_assigned_client_id="00000000-0000-0000-0000-000000000001",
        )

        self.assertIs(credential, credential_type.return_value)
        credential_type.assert_called_once_with(
            client_id="00000000-0000-0000-0000-000000000001"
        )

    def test_user_assigned_requires_client_id(self):
        with self.assertRaisesRegex(ValueError, "requires its client ID"):
            create_credential("user")

    @patch("managed_identity_demo.auth.DefaultAzureCredential")
    def test_default_mode_supports_user_assigned_identity_in_azure(
        self, credential_type
    ):
        credential = create_credential(
            "default",
            user_assigned_client_id="00000000-0000-0000-0000-000000000001",
        )

        self.assertIs(credential, credential_type.return_value)
        credential_type.assert_called_once_with(
            managed_identity_client_id="00000000-0000-0000-0000-000000000001"
        )
