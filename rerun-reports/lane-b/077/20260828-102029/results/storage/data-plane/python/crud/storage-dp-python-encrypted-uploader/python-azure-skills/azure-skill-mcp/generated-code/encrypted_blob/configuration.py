"""Environment-based Azure client configuration."""

from __future__ import annotations

import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing."""


def _required_environment_value(name: str) -> str:
    value = os.getenv(name)
    if value is None or not value.strip():
        raise ConfigurationError(f"Required environment variable {name} is not set.")
    return value.strip()


@dataclass(frozen=True)
class AzureSettings:
    """Endpoints and resource names used by the demo."""

    storage_account_url: str
    container_name: str
    key_vault_url: str
    key_name: str
    key_version: str | None
    blob_name: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        key_version = os.getenv("AZURE_KEY_VERSION")
        return cls(
            storage_account_url=_required_environment_value(
                "AZURE_STORAGE_ACCOUNT_URL"
            ).rstrip("/"),
            container_name=_required_environment_value(
                "AZURE_STORAGE_CONTAINER_NAME"
            ),
            key_vault_url=_required_environment_value("AZURE_KEY_VAULT_URL").rstrip(
                "/"
            ),
            key_name=_required_environment_value("AZURE_KEY_NAME"),
            key_version=key_version.strip() if key_version and key_version.strip() else None,
            blob_name=os.getenv("AZURE_BLOB_NAME", "encrypted-demo.bin").strip()
            or "encrypted-demo.bin",
        )


class SyncAzureClients:
    """Synchronous Azure clients sharing one credential instance."""

    def __init__(self, settings: AzureSettings) -> None:
        self.settings = settings
        self.credential = DefaultAzureCredential()
        self.blob_service_client = BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container_client = self.blob_service_client.get_container_client(
            settings.container_name
        )
        self.key_client = KeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    def close(self) -> None:
        self.key_client.close()
        self.blob_service_client.close()
        self.credential.close()

    def __enter__(self) -> "SyncAzureClients":
        return self

    def __exit__(self, exc_type: object, exc_value: object, traceback: object) -> None:
        self.close()


class AsyncAzureClients:
    """Asynchronous Azure clients sharing one async credential instance."""

    def __init__(self, settings: AzureSettings) -> None:
        self.settings = settings
        self.credential = AsyncDefaultAzureCredential()
        self.blob_service_client = AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=self.credential,
        )
        self.container_client = self.blob_service_client.get_container_client(
            settings.container_name
        )
        self.key_client = AsyncKeyClient(
            vault_url=settings.key_vault_url,
            credential=self.credential,
        )

    async def close(self) -> None:
        await self.key_client.close()
        await self.blob_service_client.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncAzureClients":
        return self

    async def __aexit__(
        self, exc_type: object, exc_value: object, traceback: object
    ) -> None:
        await self.close()
