"""Azure SDK client construction using token credentials."""

from __future__ import annotations

from dataclasses import dataclass

from azure.core.credentials import TokenCredential
from azure.keyvault.secrets import SecretClient
from azure.storage.blob import BlobServiceClient


@dataclass(frozen=True)
class AzureClients:
    secret_client: SecretClient | None
    blob_service_client: BlobServiceClient | None


def create_clients(
    credential: TokenCredential,
    *,
    key_vault_url: str | None = None,
    storage_account_url: str | None = None,
) -> AzureClients:
    """Create configured clients without making network requests."""
    secret_client = (
        SecretClient(vault_url=key_vault_url, credential=credential)
        if key_vault_url
        else None
    )
    blob_service_client = (
        BlobServiceClient(account_url=storage_account_url, credential=credential)
        if storage_account_url
        else None
    )
    return AzureClients(
        secret_client=secret_client,
        blob_service_client=blob_service_client,
    )

