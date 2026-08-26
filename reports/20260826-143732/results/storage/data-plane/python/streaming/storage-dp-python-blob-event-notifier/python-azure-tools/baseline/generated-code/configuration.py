from __future__ import annotations

import os
from dataclasses import dataclass

from azure.eventgrid import EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


@dataclass(frozen=True)
class AzureSettings:
    storage_account_url: str
    event_grid_topic_endpoint: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        return cls(
            storage_account_url=os.environ["AZURE_STORAGE_ACCOUNT_URL"],
            event_grid_topic_endpoint=os.environ["EVENT_GRID_TOPIC_ENDPOINT"],
        )


@dataclass
class SyncAzureClients:
    credential: DefaultAzureCredential
    blob_service: BlobServiceClient
    event_grid_publisher: EventGridPublisherClient

    def close(self) -> None:
        self.event_grid_publisher.close()
        self.blob_service.close()
        self.credential.close()

    def __enter__(self) -> "SyncAzureClients":
        return self

    def __exit__(self, *_: object) -> None:
        self.close()


@dataclass
class AsyncAzureClients:
    credential: AsyncDefaultAzureCredential
    blob_service: AsyncBlobServiceClient
    event_grid_publisher: AsyncEventGridPublisherClient

    async def close(self) -> None:
        await self.event_grid_publisher.close()
        await self.blob_service.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncAzureClients":
        return self

    async def __aexit__(self, *_: object) -> None:
        await self.close()


def create_sync_clients(settings: AzureSettings) -> SyncAzureClients:
    credential = DefaultAzureCredential()
    return SyncAzureClients(
        credential=credential,
        blob_service=BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ),
        event_grid_publisher=EventGridPublisherClient(
            endpoint=settings.event_grid_topic_endpoint,
            credential=credential,
        ),
    )


def create_async_clients(settings: AzureSettings) -> AsyncAzureClients:
    credential = AsyncDefaultAzureCredential()
    return AsyncAzureClients(
        credential=credential,
        blob_service=AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ),
        event_grid_publisher=AsyncEventGridPublisherClient(
            endpoint=settings.event_grid_topic_endpoint,
            credential=credential,
        ),
    )
