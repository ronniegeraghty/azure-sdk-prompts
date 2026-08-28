"""Passwordless Azure client configuration."""

from __future__ import annotations

import os
from dataclasses import dataclass

from azure.eventgrid import EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def _required_setting(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set")
    return value


@dataclass
class AzureClients:
    credential: DefaultAzureCredential
    blob_service: BlobServiceClient
    event_grid_publisher: EventGridPublisherClient

    def close(self) -> None:
        self.event_grid_publisher.close()
        self.blob_service.close()
        self.credential.close()


@dataclass
class AsyncAzureClients:
    credential: AsyncDefaultAzureCredential
    blob_service: AsyncBlobServiceClient
    event_grid_publisher: AsyncEventGridPublisherClient

    async def close(self) -> None:
        await self.event_grid_publisher.close()
        await self.blob_service.close()
        await self.credential.close()


def create_azure_clients() -> AzureClients:
    """Create synchronous clients authenticated without keys or SAS tokens."""
    credential = DefaultAzureCredential()
    return AzureClients(
        credential=credential,
        blob_service=BlobServiceClient(
            account_url=_required_setting("AZURE_STORAGE_ACCOUNT_URL"),
            credential=credential,
        ),
        event_grid_publisher=EventGridPublisherClient(
            endpoint=_required_setting("AZURE_EVENTGRID_TOPIC_ENDPOINT"),
            credential=credential,
        ),
    )


def create_async_azure_clients() -> AsyncAzureClients:
    """Create asynchronous clients authenticated without keys or SAS tokens."""
    credential = AsyncDefaultAzureCredential()
    return AsyncAzureClients(
        credential=credential,
        blob_service=AsyncBlobServiceClient(
            account_url=_required_setting("AZURE_STORAGE_ACCOUNT_URL"),
            credential=credential,
        ),
        event_grid_publisher=AsyncEventGridPublisherClient(
            endpoint=_required_setting("AZURE_EVENTGRID_TOPIC_ENDPOINT"),
            credential=credential,
        ),
    )
