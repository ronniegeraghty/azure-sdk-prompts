import pytest

from managed_identity_demo.auth import IdentityMode
from managed_identity_demo.config import Settings


def test_user_assigned_settings_require_client_id():
    settings = Settings(
        identity_mode=IdentityMode.USER_ASSIGNED,
        storage_account_url="https://storage.example",
        key_vault_url="https://vault.example",
    )

    with pytest.raises(ValueError, match="AZURE_CLIENT_ID is required"):
        settings.validate()


def test_endpoints_must_use_https():
    settings = Settings(
        identity_mode=IdentityMode.LOCAL,
        storage_account_url="http://storage.example",
        key_vault_url="https://vault.example",
    )

    with pytest.raises(ValueError, match="absolute HTTPS URL"):
        settings.validate()

