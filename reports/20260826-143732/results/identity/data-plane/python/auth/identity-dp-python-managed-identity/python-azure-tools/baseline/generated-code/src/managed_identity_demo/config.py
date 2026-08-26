"""Environment-based application configuration."""

from dataclasses import dataclass
import os
from urllib.parse import urlparse

from .auth import IdentityMode


@dataclass(frozen=True)
class Settings:
    identity_mode: IdentityMode
    storage_account_url: str
    key_vault_url: str
    managed_identity_client_id: str | None = None

    @classmethod
    def from_environment(cls) -> "Settings":
        raw_mode = os.getenv("AZURE_IDENTITY_MODE", IdentityMode.LOCAL.value)
        try:
            mode = IdentityMode(raw_mode)
        except ValueError as error:
            choices = ", ".join(item.value for item in IdentityMode)
            raise ValueError(
                f"AZURE_IDENTITY_MODE must be one of: {choices}"
            ) from error

        settings = cls(
            identity_mode=mode,
            storage_account_url=os.getenv(
                "AZURE_STORAGE_ACCOUNT_URL",
                "https://example.blob.core.windows.net",
            ),
            key_vault_url=os.getenv(
                "AZURE_KEY_VAULT_URL",
                "https://example.vault.azure.net",
            ),
            managed_identity_client_id=os.getenv("AZURE_CLIENT_ID"),
        )
        settings.validate()
        return settings

    def validate(self) -> None:
        _validate_https_url("AZURE_STORAGE_ACCOUNT_URL", self.storage_account_url)
        _validate_https_url("AZURE_KEY_VAULT_URL", self.key_vault_url)

        if (
            self.identity_mode is IdentityMode.USER_ASSIGNED
            and not self.managed_identity_client_id
        ):
            raise ValueError(
                "AZURE_CLIENT_ID is required when AZURE_IDENTITY_MODE=user"
            )
        if (
            self.identity_mode is IdentityMode.SYSTEM_ASSIGNED
            and self.managed_identity_client_id
        ):
            raise ValueError(
                "Unset AZURE_CLIENT_ID when AZURE_IDENTITY_MODE=system"
            )


def _validate_https_url(variable_name: str, value: str) -> None:
    parsed = urlparse(value)
    if parsed.scheme != "https" or not parsed.netloc:
        raise ValueError(f"{variable_name} must be an absolute HTTPS URL")

