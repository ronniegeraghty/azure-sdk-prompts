"""Credential construction for Azure-hosted and local environments."""

from __future__ import annotations

from typing import Literal, Optional

from azure.core.credentials import TokenCredential
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential

IdentityMode = Literal["system", "user", "default"]


def create_credential(
    mode: IdentityMode,
    *,
    user_assigned_client_id: Optional[str] = None,
) -> TokenCredential:
    """Create the credential appropriate for the selected execution environment."""
    if mode == "system":
        return ManagedIdentityCredential()

    if mode == "user":
        if not user_assigned_client_id:
            raise ValueError(
                "A user-assigned identity requires its client ID. "
                "Set AZURE_CLIENT_ID or pass --client-id."
            )
        return ManagedIdentityCredential(client_id=user_assigned_client_id)

    if mode == "default":
        return DefaultAzureCredential(
            managed_identity_client_id=user_assigned_client_id,
        )

    raise ValueError(f"Unsupported identity mode: {mode}")
