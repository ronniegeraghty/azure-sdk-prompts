"""Credential factories for Azure-hosted and local execution."""

from __future__ import annotations

import os
from enum import Enum
from typing import Optional

from azure.core.credentials import TokenCredential
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential


class IdentityType(str, Enum):
    SYSTEM_ASSIGNED = "system"
    USER_ASSIGNED = "user"
    LOCAL = "local"


def create_managed_identity_credential(
    identity_type: IdentityType,
    *,
    client_id: Optional[str] = None,
) -> ManagedIdentityCredential:
    """Create a credential for a system- or user-assigned managed identity."""
    if identity_type is IdentityType.SYSTEM_ASSIGNED:
        if client_id:
            raise ValueError("client_id must not be supplied for a system-assigned identity")
        return ManagedIdentityCredential()

    if identity_type is IdentityType.USER_ASSIGNED:
        resolved_client_id = client_id or os.getenv("AZURE_CLIENT_ID")
        if not resolved_client_id:
            raise ValueError(
                "A user-assigned identity requires its client ID. Pass --client-id "
                "or set AZURE_CLIENT_ID."
            )
        return ManagedIdentityCredential(client_id=resolved_client_id)

    raise ValueError("ManagedIdentityCredential supports only 'system' or 'user'")


def create_credential(
    identity_type: IdentityType,
    *,
    client_id: Optional[str] = None,
) -> TokenCredential:
    """Create an explicit managed identity credential or an opt-in local fallback."""
    if identity_type is not IdentityType.LOCAL:
        return create_managed_identity_credential(identity_type, client_id=client_id)

    if os.getenv("AZURE_ALLOW_LOCAL_CREDENTIALS") != "1":
        raise ValueError(
            "Local credential fallback is disabled. Set AZURE_ALLOW_LOCAL_CREDENTIALS=1 "
            "after signing in with Azure CLI, Azure Developer CLI, or a supported IDE."
        )

    # Exclude environment credentials to avoid accidentally using a client secret
    # intended for another application.
    return DefaultAzureCredential(
        exclude_environment_credential=True,
        exclude_managed_identity_credential=True,
        exclude_shared_token_cache_credential=True,
        exclude_interactive_browser_credential=True,
    )
