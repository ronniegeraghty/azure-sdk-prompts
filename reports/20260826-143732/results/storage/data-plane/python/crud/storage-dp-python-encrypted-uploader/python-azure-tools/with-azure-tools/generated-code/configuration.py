from __future__ import annotations

import os
from contextlib import AsyncExitStack, ExitStack, asynccontextmanager, contextmanager
from dataclasses import dataclass
from typing import AsyncIterator, Iterator

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


class ConfigurationError(ValueError):
    """Raised when required application configuration is missing."""


@dataclass(frozen=True)
class Settings:
    storage_account_url: str
    storage_container: str
    key_vault_url: str
    key_name: str
    blob_name: str

    @classmethod
    def from_environment(cls) -> "Settings":
        required = {
            "AZURE_STORAGE_ACCOUNT_URL": os.getenv("AZURE_STORAGE_ACCOUNT_URL"),
            "AZURE_STORAGE_CONTAINER": os.getenv("AZURE_STORAGE_CONTAINER"),
            "AZURE_KEY_VAULT_URL": os.getenv("AZURE_KEY_VAULT_URL"),
            "AZURE_KEY_VAULT_KEY_NAME": os.getenv("AZURE_KEY_VAULT_KEY_NAME"),
        }
        missing = [name for name, value in required.items() if not value]
        if missing:
            raise ConfigurationError(
                "Missing required environment variables: " + ", ".join(sorted(missing))
            )

        return cls(
            storage_account_url=required["AZURE_STORAGE_ACCOUNT_URL"] or "",
            storage_container=required["AZURE_STORAGE_CONTAINER"] or "",
            key_vault_url=required["AZURE_KEY_VAULT_URL"] or "",
            key_name=required["AZURE_KEY_VAULT_KEY_NAME"] or "",
            blob_name=os.getenv("AZURE_BLOB_NAME", "encrypted-demo.bin"),
        )


@dataclass(frozen=True)
class SyncAzureClients:
    credential: DefaultAzureCredential
    blob_service: BlobServiceClient
    key_client: KeyClient


@dataclass(frozen=True)
class AsyncAzureClients:
    credential: AsyncDefaultAzureCredential
    blob_service: AsyncBlobServiceClient
    key_client: AsyncKeyClient


@contextmanager
def create_sync_clients(settings: Settings) -> Iterator[SyncAzureClients]:
    with ExitStack() as stack:
        credential = stack.enter_context(DefaultAzureCredential())
        blob_service = stack.enter_context(
            BlobServiceClient(
                account_url=settings.storage_account_url,
                credential=credential,
            )
        )
        key_client = stack.enter_context(
            KeyClient(vault_url=settings.key_vault_url, credential=credential)
        )
        yield SyncAzureClients(credential, blob_service, key_client)


@asynccontextmanager
async def create_async_clients(
    settings: Settings,
) -> AsyncIterator[AsyncAzureClients]:
    async with AsyncExitStack() as stack:
        credential = await stack.enter_async_context(AsyncDefaultAzureCredential())
        blob_service = await stack.enter_async_context(
            AsyncBlobServiceClient(
                account_url=settings.storage_account_url,
                credential=credential,
            )
        )
        key_client = await stack.enter_async_context(
            AsyncKeyClient(vault_url=settings.key_vault_url, credential=credential)
        )
        yield AsyncAzureClients(credential, blob_service, key_client)
