"""Azure connection configuration for the encrypted blob demo."""

from __future__ import annotations

import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.storage.blob import BlobServiceClient, ContainerClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
from azure.storage.blob.aio import ContainerClient as AsyncContainerClient


def _required_environment(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set")
    return value


@dataclass(frozen=True)
class AzureSettings:
    storage_account_url: str
    storage_container_name: str
    key_vault_url: str
    key_name: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        return cls(
            storage_account_url=_required_environment(
                "AZURE_STORAGE_ACCOUNT_URL"
            ).rstrip("/"),
            storage_container_name=_required_environment(
                "AZURE_STORAGE_CONTAINER_NAME"
            ),
            key_vault_url=_required_environment("AZURE_KEY_VAULT_URL").rstrip("/"),
            key_name=_required_environment("AZURE_KEY_NAME"),
        )


class SyncAzureConnections:
    """Owns one credential shared by all synchronous Azure clients."""

    def __init__(self, settings: AzureSettings) -> None:
        self.settings = settings
        self.credential = DefaultAzureCredential()
        self.blob_service = BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container: ContainerClient = self.blob_service.get_container_client(
            settings.storage_container_name
        )
        self.key_client = KeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    def close(self) -> None:
        self.key_client.close()
        self.blob_service.close()
        self.credential.close()

    def __enter__(self) -> "SyncAzureConnections":
        return self

    def __exit__(self, *_: object) -> None:
        self.close()


class AsyncAzureConnections:
    """Owns one async credential shared by all asynchronous Azure clients."""

    def __init__(self, settings: AzureSettings) -> None:
        self.settings = settings
        self.credential = AsyncDefaultAzureCredential()
        self.blob_service = AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container: AsyncContainerClient = (
            self.blob_service.get_container_client(settings.storage_container_name)
        )
        self.key_client = AsyncKeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    async def close(self) -> None:
        await self.key_client.close()
        await self.blob_service.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncAzureConnections":
        return self

    async def __aexit__(self, *_: object) -> None:
        await self.close()
