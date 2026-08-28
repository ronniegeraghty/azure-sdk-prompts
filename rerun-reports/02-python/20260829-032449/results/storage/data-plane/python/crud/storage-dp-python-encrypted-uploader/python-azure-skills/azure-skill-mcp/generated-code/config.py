from __future__ import annotations

import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


@dataclass(frozen=True)
class Settings:
    storage_account_url: str
    storage_container_name: str
    key_vault_url: str
    key_name: str

    @classmethod
    def from_env(cls) -> "Settings":
        return cls(
            storage_account_url=_required_env("AZURE_STORAGE_ACCOUNT_URL"),
            storage_container_name=_required_env("AZURE_STORAGE_CONTAINER_NAME"),
            key_vault_url=_required_env("AZURE_KEY_VAULT_URL"),
            key_name=_required_env("AZURE_KEY_NAME"),
        )


@dataclass
class SyncConnections:
    credential: DefaultAzureCredential
    blob_service_client: BlobServiceClient
    key_client: KeyClient

    def close(self) -> None:
        self.key_client.close()
        self.blob_service_client.close()
        self.credential.close()

    def __enter__(self) -> "SyncConnections":
        return self

    def __exit__(self, exc_type: object, exc: object, traceback: object) -> None:
        self.close()


@dataclass
class AsyncConnections:
    credential: AsyncDefaultAzureCredential
    blob_service_client: AsyncBlobServiceClient
    key_client: AsyncKeyClient

    async def close(self) -> None:
        await self.key_client.close()
        await self.blob_service_client.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncConnections":
        return self

    async def __aexit__(
        self, exc_type: object, exc: object, traceback: object
    ) -> None:
        await self.close()


def build_sync_connections(settings: Settings) -> SyncConnections:
    credential = DefaultAzureCredential()
    return SyncConnections(
        credential=credential,
        blob_service_client=BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ),
        key_client=KeyClient(
            vault_url=settings.key_vault_url,
            credential=credential,
        ),
    )


def build_async_connections(settings: Settings) -> AsyncConnections:
    credential = AsyncDefaultAzureCredential()
    return AsyncConnections(
        credential=credential,
        blob_service_client=AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ),
        key_client=AsyncKeyClient(
            vault_url=settings.key_vault_url,
            credential=credential,
        ),
    )


def _required_env(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set")
    return value
