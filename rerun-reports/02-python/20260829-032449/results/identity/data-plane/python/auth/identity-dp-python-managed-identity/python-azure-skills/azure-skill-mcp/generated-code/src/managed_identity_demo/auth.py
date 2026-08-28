"""Credential factories for hosted and local execution."""

from __future__ import annotations

from enum import Enum

from azure.core.credentials import TokenCredential
from azure.identity import (
    AzureCliCredential,
    DefaultAzureCredential,
    ManagedIdentityCredential,
)


class AuthMode(str, Enum):
    SYSTEM = "system"
    USER = "user"
    LOCAL_DEFAULT = "local-default"
    LOCAL_CLI = "local-cli"


def create_credential(
    mode: AuthMode,
    *,
    managed_identity_client_id: str | None = None,
) -> TokenCredential:
    """Create a credential appropriate for the selected runtime."""
    if mode is AuthMode.SYSTEM:
        return ManagedIdentityCredential()

    if mode is AuthMode.USER:
        if not managed_identity_client_id or not managed_identity_client_id.strip():
            raise ValueError(
                "User-assigned identity requires "
                "AZURE_MANAGED_IDENTITY_CLIENT_ID or --client-id."
            )
        return ManagedIdentityCredential(client_id=managed_identity_client_id.strip())

    if mode is AuthMode.LOCAL_DEFAULT:
        # Local development should not wait for an Azure-hosted identity endpoint.
        return DefaultAzureCredential(
            exclude_environment_credential=True,
            exclude_workload_identity_credential=True,
            exclude_managed_identity_credential=True,
        )

    if mode is AuthMode.LOCAL_CLI:
        return AzureCliCredential()

    raise ValueError(f"Unsupported authentication mode: {mode}")

