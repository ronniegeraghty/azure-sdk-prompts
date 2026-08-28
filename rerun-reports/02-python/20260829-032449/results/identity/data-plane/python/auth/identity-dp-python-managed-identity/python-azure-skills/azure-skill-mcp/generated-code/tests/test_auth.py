from __future__ import annotations

import unittest
from unittest.mock import patch

from managed_identity_demo.auth import AuthMode, create_credential


class CreateCredentialTests(unittest.TestCase):
    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_system_assigned_has_no_selector(self, credential_type) -> None:
        create_credential(AuthMode.SYSTEM)
        credential_type.assert_called_once_with()

    @patch("managed_identity_demo.auth.ManagedIdentityCredential")
    def test_user_assigned_uses_client_id(self, credential_type) -> None:
        create_credential(
            AuthMode.USER,
            managed_identity_client_id=" 11111111-1111-1111-1111-111111111111 ",
        )
        credential_type.assert_called_once_with(
            client_id="11111111-1111-1111-1111-111111111111"
        )

    def test_user_assigned_requires_client_id(self) -> None:
        with self.assertRaisesRegex(ValueError, "requires"):
            create_credential(AuthMode.USER)

    @patch("managed_identity_demo.auth.DefaultAzureCredential")
    def test_local_default_skips_hosted_credentials(self, credential_type) -> None:
        create_credential(AuthMode.LOCAL_DEFAULT)
        credential_type.assert_called_once_with(
            exclude_environment_credential=True,
            exclude_workload_identity_credential=True,
            exclude_managed_identity_credential=True,
        )


if __name__ == "__main__":
    unittest.main()

