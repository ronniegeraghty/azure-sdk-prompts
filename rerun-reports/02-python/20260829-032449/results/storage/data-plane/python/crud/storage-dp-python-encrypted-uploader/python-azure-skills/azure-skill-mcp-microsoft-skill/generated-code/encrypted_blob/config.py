"""Environment-based Azure client configuration."""

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


def _required_environment_variable(name: str) -> str:
    value = os.getenv(name, "").strip()
    if not value:
        raise ValueError(f"Required environment variable {name} is not set")
    return value


@dataclass(frozen=True)
class Settings:
    storage_account_url: str
    storage_container_name: str
    key_vault_url: str
    key_name: str
    input_file: str
    sync_blob_name: str
    async_blob_name: str
    sync_output_file: str
    async_output_file: str

    @classmethod
    def from_environment(cls) -> "Settings":
        return cls(
            storage_account_url=_required_environment_variable(
                "AZURE_STORAGE_ACCOUNT_URL"
            ),
            storage_container_name=_required_environment_variable(
                "AZURE_STORAGE_CONTAINER_NAME"
            ),
            key_vault_url=_required_environment_variable("AZURE_KEYVAULT_URL"),
            key_name=_required_environment_variable("AZURE_KEY_NAME"),
            input_file=os.getenv("DEMO_INPUT_FILE", "sample.txt"),
            sync_blob_name=os.getenv(
                "DEMO_SYNC_BLOB_NAME", "encrypted-sync/sample.txt"
            ),
            async_blob_name=os.getenv(
                "DEMO_ASYNC_BLOB_NAME", "encrypted-async/sample.txt"
            ),
            sync_output_file=os.getenv(
                "DEMO_SYNC_OUTPUT_FILE", "downloaded-sync.txt"
            ),
            async_output_file=os.getenv(
                "DEMO_ASYNC_OUTPUT_FILE", "downloaded-async.txt"
            ),
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
    """Build sync Azure clients that all use one credential instance."""
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
    """Build async Azure clients that all use one credential instance."""
    async with AsyncExitStack() as stack:
        credential = await stack.enter_async_context(AsyncDefaultAzureCredential())
        blob_service = await stack.enter_async_context(
            AsyncBlobServiceClient(
                account_url=settings.storage_account_url,
                credential=credential,
            )
        )
        key_client = await stack.enter_async_context(
            AsyncKeyClient(
                vault_url=settings.key_vault_url,
                credential=credential,
            )
        )
        yield AsyncAzureClients(credential, blob_service, key_client)
