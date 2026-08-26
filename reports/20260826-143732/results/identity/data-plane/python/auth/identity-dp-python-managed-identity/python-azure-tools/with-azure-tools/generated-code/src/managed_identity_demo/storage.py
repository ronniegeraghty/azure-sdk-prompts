from __future__ import annotations

from urllib.parse import urlsplit

from azure.core.credentials import TokenCredential
from azure.storage.blob import BlobServiceClient


class StorageConfigurationError(ValueError):
    """Raised when the Blob Storage endpoint is unsafe or malformed."""


def validate_account_url(account_url: str) -> str:
    value = account_url.strip()
    parsed = urlsplit(value)

    if parsed.scheme != "https" or not parsed.hostname:
        raise StorageConfigurationError(
            "The Blob Storage account URL must be an absolute HTTPS URL."
        )
    if parsed.username or parsed.password or parsed.query or parsed.fragment:
        raise StorageConfigurationError(
            "The account URL must not contain credentials, a query string, or a fragment."
        )

    return value.rstrip("/")


def list_container_names(
    account_url: str,
    credential: TokenCredential,
) -> list[str]:
    """List Blob Storage containers using Microsoft Entra token authentication."""
    validated_url = validate_account_url(account_url)
    with BlobServiceClient(
        account_url=validated_url,
        credential=credential,
        retry_total=3,
        retry_backoff_factor=0.8,
        connection_timeout=10,
        read_timeout=30,
    ) as client:
        return [container["name"] for container in client.list_containers()]
