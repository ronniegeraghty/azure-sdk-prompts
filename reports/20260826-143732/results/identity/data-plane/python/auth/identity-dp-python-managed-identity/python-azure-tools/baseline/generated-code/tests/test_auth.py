from unittest.mock import patch

import pytest

from managed_identity_demo.auth import IdentityMode, create_credential


@patch("managed_identity_demo.auth.ManagedIdentityCredential")
def test_system_assigned_omits_client_id(managed_identity_credential):
    create_credential(IdentityMode.SYSTEM_ASSIGNED)

    managed_identity_credential.assert_called_once_with()


@patch("managed_identity_demo.auth.ManagedIdentityCredential")
def test_user_assigned_passes_client_id(managed_identity_credential):
    create_credential(
        IdentityMode.USER_ASSIGNED,
        managed_identity_client_id="identity-client-id",
    )

    managed_identity_credential.assert_called_once_with(
        client_id="identity-client-id"
    )


def test_user_assigned_requires_client_id():
    with pytest.raises(ValueError, match="AZURE_CLIENT_ID is required"):
        create_credential(IdentityMode.USER_ASSIGNED)


def test_system_assigned_rejects_client_id():
    with pytest.raises(ValueError, match="must be omitted"):
        create_credential(
            IdentityMode.SYSTEM_ASSIGNED,
            managed_identity_client_id="unexpected",
        )


@patch("managed_identity_demo.auth.ChainedTokenCredential")
@patch("managed_identity_demo.auth.AzureCliCredential")
@patch("managed_identity_demo.auth.EnvironmentCredential")
@patch.dict(
    "os.environ",
    {
        "AZURE_TENANT_ID": "tenant",
        "AZURE_CLIENT_ID": "client",
        "AZURE_CLIENT_SECRET": "secret",
    },
    clear=True,
)
def test_local_uses_environment_then_cli(
    environment_credential,
    azure_cli_credential,
    chained_credential,
):
    create_credential(IdentityMode.LOCAL)

    chained_credential.assert_called_once_with(
        environment_credential.return_value,
        azure_cli_credential.return_value,
    )


@patch("managed_identity_demo.auth.ChainedTokenCredential")
@patch("managed_identity_demo.auth.AzureCliCredential")
@patch("managed_identity_demo.auth.EnvironmentCredential")
@patch.dict("os.environ", {"AZURE_TENANT_ID": "partial"}, clear=True)
def test_local_skips_incomplete_environment_credential(
    environment_credential,
    azure_cli_credential,
    chained_credential,
):
    create_credential(IdentityMode.LOCAL)

    environment_credential.assert_not_called()
    chained_credential.assert_called_once_with(azure_cli_credential.return_value)
