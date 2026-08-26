"""Credential construction for Azure-hosted and local environments."""

from enum import Enum
import os

from azure.core.credentials import TokenCredential
from azure.identity import (
    AzureCliCredential,
    ChainedTokenCredential,
    EnvironmentCredential,
    ManagedIdentityCredential,
)


class IdentityMode(str, Enum):
    """Supported authentication modes."""

    SYSTEM_ASSIGNED = "system"
    USER_ASSIGNED = "user"
    LOCAL = "local"


def create_credential(
    mode: IdentityMode,
    *,
    managed_identity_client_id: str | None = None,
) -> TokenCredential:
    """Create a credential suitable for the selected execution environment.

    System-assigned identity is selected by omitting an identity identifier.
    User-assigned identity is selected with its application (client) ID.
    Local mode tries service-principal environment variables, then Azure CLI.
    """
    if mode is IdentityMode.SYSTEM_ASSIGNED:
        if managed_identity_client_id:
            raise ValueError(
                "managed_identity_client_id must be omitted for system-assigned identity"
            )
        return ManagedIdentityCredential()

    if mode is IdentityMode.USER_ASSIGNED:
        if not managed_identity_client_id:
            raise ValueError(
                "AZURE_CLIENT_ID is required for user-assigned managed identity"
            )
        return ManagedIdentityCredential(client_id=managed_identity_client_id)

    if mode is IdentityMode.LOCAL:
        if managed_identity_client_id:
            raise ValueError(
                "managed_identity_client_id is not used by the local credential chain"
            )
        credentials: list[TokenCredential] = []
        service_principal_variables = (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_CLIENT_SECRET",
        )
        if all(os.getenv(name) for name in service_principal_variables):
            credentials.append(EnvironmentCredential())
        credentials.append(AzureCliCredential())
        return ChainedTokenCredential(*credentials)

    raise ValueError(f"Unsupported identity mode: {mode!r}")
