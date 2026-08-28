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


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing."""


def _required_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ConfigurationError(f"Required environment variable {name!r} is not set")
    return value


@dataclass(frozen=True)
class Settings:
    storage_account_url: str
    blob_container: str
    key_vault_url: str
    key_name: str
    key_version: str | None = None

    @classmethod
    def from_environment(cls) -> "Settings":
        return cls(
            storage_account_url=_required_environment_variable(
                "AZURE_STORAGE_ACCOUNT_URL"
            ),
            blob_container=_required_environment_variable("AZURE_BLOB_CONTAINER"),
            key_vault_url=_required_environment_variable("AZURE_KEY_VAULT_URL"),
            key_name=_required_environment_variable("AZURE_KEY_NAME"),
            key_version=os.getenv("AZURE_KEY_VERSION") or None,
        )


class SyncAzureConnections:
    """Builds all synchronous Azure clients from one shared credential."""

    def __init__(
        self,
        settings: Settings,
        credential: DefaultAzureCredential | None = None,
    ) -> None:
        self.settings = settings
        self.credential = credential or DefaultAzureCredential()
        self.blob_service_client = BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container_client: ContainerClient = (
            self.blob_service_client.get_container_client(settings.blob_container)
        )
        self.key_client = KeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    def close(self) -> None:
        self.key_client.close()
        self.blob_service_client.close()
        self.credential.close()

    def __enter__(self) -> "SyncAzureConnections":
        return self

    def __exit__(self, *_: object) -> None:
        self.close()


class AsyncAzureConnections:
    """Builds all asynchronous Azure clients from one shared async credential."""

    def __init__(
        self,
        settings: Settings,
        credential: AsyncDefaultAzureCredential | None = None,
    ) -> None:
        self.settings = settings
        self.credential = credential or AsyncDefaultAzureCredential()
        self.blob_service_client = AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container_client: AsyncContainerClient = (
            self.blob_service_client.get_container_client(settings.blob_container)
        )
        self.key_client = AsyncKeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    async def close(self) -> None:
        await self.key_client.close()
        await self.blob_service_client.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncAzureConnections":
        return self

    async def __aexit__(self, *_: object) -> None:
        await self.close()
