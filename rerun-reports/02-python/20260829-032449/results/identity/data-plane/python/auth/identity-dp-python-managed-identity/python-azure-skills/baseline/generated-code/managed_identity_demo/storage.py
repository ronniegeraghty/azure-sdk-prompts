"""Azure Blob Storage client construction and a simple authenticated operation."""

from __future__ import annotations

from typing import Iterable

from azure.core.credentials import TokenCredential
from azure.storage.blob import BlobServiceClient


def create_blob_service_client(
    account_url: str,
    credential: TokenCredential,
) -> BlobServiceClient:
    if not account_url.startswith("https://") or not account_url.endswith(
        ".blob.core.windows.net"
    ):
        raise ValueError(
            "account_url must look like https://<account>.blob.core.windows.net"
        )
    return BlobServiceClient(account_url=account_url, credential=credential)


def list_container_names(client: BlobServiceClient) -> Iterable[str]:
    """List containers, proving that authentication and RBAC authorization work."""
    for container in client.list_containers():
        yield container["name"]
