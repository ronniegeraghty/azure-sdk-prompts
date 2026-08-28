from unittest.mock import patch

import pytest

from managed_identity_demo.auth import (
    IdentityType,
    create_credential,
    create_managed_identity_credential,
)


@patch("managed_identity_demo.auth.ManagedIdentityCredential")
def test_system_assigned_uses_no_client_id(credential_type):
    create_managed_identity_credential(IdentityType.SYSTEM_ASSIGNED)

    credential_type.assert_called_once_with()


@patch("managed_identity_demo.auth.ManagedIdentityCredential")
def test_user_assigned_uses_explicit_client_id(credential_type):
    create_managed_identity_credential(
        IdentityType.USER_ASSIGNED,
        client_id="identity-client-id",
    )

    credential_type.assert_called_once_with(client_id="identity-client-id")


def test_user_assigned_requires_client_id(monkeypatch):
    monkeypatch.delenv("AZURE_CLIENT_ID", raising=False)

    with pytest.raises(ValueError, match="requires its client ID"):
        create_managed_identity_credential(IdentityType.USER_ASSIGNED)


def test_local_fallback_requires_explicit_opt_in(monkeypatch):
    monkeypatch.delenv("AZURE_ALLOW_LOCAL_CREDENTIALS", raising=False)

    with pytest.raises(ValueError, match="fallback is disabled"):
        create_credential(IdentityType.LOCAL)
