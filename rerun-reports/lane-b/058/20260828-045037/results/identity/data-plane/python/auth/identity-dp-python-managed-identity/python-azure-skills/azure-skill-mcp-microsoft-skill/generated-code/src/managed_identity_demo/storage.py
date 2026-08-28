"""Blob Storage client usage with Microsoft Entra token credentials."""

from typing import List

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.storage.blob import BlobServiceClient

from .credentials import AzureCredential


class AzureOperationError(RuntimeError):
    """An actionable Azure authentication or service failure."""


def build_blob_service_client(
    account_url: str, credential: AzureCredential
) -> BlobServiceClient:
    """Construct a Blob client without making a network request."""
    if not account_url.startswith("https://") or not account_url.rstrip("/").endswith(
        ".blob.core.windows.net"
    ):
        raise ValueError(
            "Account URL must look like "
            "'https://<account>.blob.core.windows.net'."
        )
    return BlobServiceClient(
        account_url=account_url.rstrip("/"),
        credential=credential,
    )


def list_container_names(
    account_url: str, credential: AzureCredential
) -> List[str]:
    """Authenticate and list containers visible to the selected identity."""
    try:
        with build_blob_service_client(account_url, credential) as client:
            return [container.name for container in client.list_containers()]
    except ClientAuthenticationError as exc:
        raise AzureOperationError(
            "Authentication failed. Confirm the identity is enabled and, for a "
            "user-assigned identity, that its client ID is correct."
        ) from exc
    except ServiceRequestError as exc:
        raise AzureOperationError(
            "Azure could not be reached. Check DNS, proxy, firewall, and the "
            "managed identity endpoint availability."
        ) from exc
    except HttpResponseError as exc:
        if exc.status_code == 403:
            detail = (
                "Authentication succeeded but access was denied. Assign a data-plane "
                "role such as Storage Blob Data Reader and allow time for propagation."
            )
        else:
            detail = f"Blob Storage returned HTTP {exc.status_code or 'unknown'}."
        raise AzureOperationError(detail) from exc
