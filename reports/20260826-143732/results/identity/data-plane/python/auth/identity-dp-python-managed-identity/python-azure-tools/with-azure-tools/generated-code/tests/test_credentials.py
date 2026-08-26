from unittest.mock import patch

import pytest

from managed_identity_demo.credentials import (
    CredentialConfigurationError,
    IdentityMode,
    create_credential,
)


@patch("managed_identity_demo.credentials.ManagedIdentityCredential")
def test_system_assigned_uses_default_managed_identity(mock_credential):
    create_credential(IdentityMode.SYSTEM)

    mock_credential.assert_called_once_with()


@patch("managed_identity_demo.credentials.ManagedIdentityCredential")
def test_user_assigned_uses_client_id(mock_credential):
    create_credential(IdentityMode.USER, " identity-client-id ")

    mock_credential.assert_called_once_with(client_id="identity-client-id")


def test_user_assigned_requires_client_id():
    with pytest.raises(CredentialConfigurationError):
        create_credential(IdentityMode.USER)


@patch("managed_identity_demo.credentials.DefaultAzureCredential")
def test_local_mode_skips_managed_identity_probe(mock_credential):
    create_credential(IdentityMode.LOCAL)

    mock_credential.assert_called_once_with(exclude_managed_identity_credential=True)


@patch("managed_identity_demo.credentials.DefaultAzureCredential")
def test_auto_mode_targets_user_assigned_identity_when_configured(mock_credential):
    create_credential(IdentityMode.AUTO, "identity-client-id")

    mock_credential.assert_called_once_with(
        managed_identity_client_id="identity-client-id"
    )
