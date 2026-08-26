from __future__ import annotations

from enum import Enum

from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

AzureCredential = ManagedIdentityCredential | DefaultAzureCredential


class CredentialConfigurationError(ValueError):
    """Raised when the selected identity mode is missing required settings."""


class IdentityMode(str, Enum):
    SYSTEM = "system"
    USER = "user"
    LOCAL = "local"
    AUTO = "auto"


def create_credential(
    mode: IdentityMode,
    managed_identity_client_id: str | None = None,
) -> AzureCredential:
    """Create a credential without performing authentication or network I/O."""
    client_id = managed_identity_client_id.strip() if managed_identity_client_id else None

    if mode is IdentityMode.SYSTEM:
        return ManagedIdentityCredential()

    if mode is IdentityMode.USER:
        if not client_id:
            raise CredentialConfigurationError(
                "User-assigned mode requires MANAGED_IDENTITY_CLIENT_ID or --client-id."
            )
        return ManagedIdentityCredential(client_id=client_id)

    if mode is IdentityMode.LOCAL:
        # Avoid probing the managed identity endpoint during local development.
        return DefaultAzureCredential(exclude_managed_identity_credential=True)

    if mode is IdentityMode.AUTO:
        options = {}
        if client_id:
            options["managed_identity_client_id"] = client_id
        return DefaultAzureCredential(**options)

    raise CredentialConfigurationError(f"Unsupported identity mode: {mode}")
