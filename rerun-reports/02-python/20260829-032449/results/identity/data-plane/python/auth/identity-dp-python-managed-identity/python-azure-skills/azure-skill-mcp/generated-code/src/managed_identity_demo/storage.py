"""Azure Blob Storage client construction and sample operation."""

from __future__ import annotations

from collections.abc import Iterator
from urllib.parse import urlparse

from azure.core.credentials import TokenCredential
from azure.storage.blob import BlobServiceClient


def create_blob_service_client(
    account_url: str,
    credential: TokenCredential,
) -> BlobServiceClient:
    """Create a token-authenticated BlobServiceClient with bounded retries."""
    parsed = urlparse(account_url)
    if parsed.scheme != "https" or not parsed.netloc:
        raise ValueError(
            "AZURE_STORAGE_ACCOUNT_URL must be an HTTPS URL, for example "
            "https://myaccount.blob.core.windows.net."
        )

    return BlobServiceClient(
        account_url=account_url.rstrip("/"),
        credential=credential,
        retry_total=4,
        retry_connect=4,
        retry_read=4,
        retry_status=4,
    )


def list_container_names(client: BlobServiceClient) -> Iterator[str]:
    """Yield container names visible to the authenticated principal."""
    for container in client.list_containers(name_starts_with=None):
        yield container["name"]

