"""Example of passing a token credential to an Azure SDK client."""

from __future__ import annotations

from typing import List

from azure.core.credentials import TokenCredential
from azure.storage.blob import BlobServiceClient


def list_blob_containers(
    account_url: str,
    credential: TokenCredential,
) -> List[str]:
    """List container names using Microsoft Entra authentication."""
    with BlobServiceClient(
        account_url=account_url,
        credential=credential,
    ) as client:
        return [container["name"] for container in client.list_containers()]
