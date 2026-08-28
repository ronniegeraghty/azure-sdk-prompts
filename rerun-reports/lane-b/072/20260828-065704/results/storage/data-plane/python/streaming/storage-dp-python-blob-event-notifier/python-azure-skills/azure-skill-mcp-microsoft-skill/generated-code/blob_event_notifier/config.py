from __future__ import annotations

import os
from contextlib import asynccontextmanager, contextmanager
from dataclasses import dataclass
from typing import AsyncIterator, Iterator

from azure.eventgrid import EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


@dataclass(frozen=True)
class AzureSettings:
    storage_account_url: str
    eventgrid_topic_endpoint: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        return cls(
            storage_account_url=os.environ["AZURE_STORAGE_ACCOUNT_URL"],
            eventgrid_topic_endpoint=os.environ["EVENTGRID_TOPIC_ENDPOINT"],
        )


@contextmanager
def azure_clients(
    settings: AzureSettings,
) -> Iterator[tuple[BlobServiceClient, EventGridPublisherClient]]:
    with DefaultAzureCredential() as credential:
        with BlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ) as blob_service:
            with EventGridPublisherClient(
                endpoint=settings.eventgrid_topic_endpoint,
                credential=credential,
            ) as eventgrid_publisher:
                yield blob_service, eventgrid_publisher


@asynccontextmanager
async def async_azure_clients(
    settings: AzureSettings,
) -> AsyncIterator[tuple[AsyncBlobServiceClient, AsyncEventGridPublisherClient]]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncBlobServiceClient(
            account_url=settings.storage_account_url,
            credential=credential,
        ) as blob_service:
            async with AsyncEventGridPublisherClient(
                endpoint=settings.eventgrid_topic_endpoint,
                credential=credential,
            ) as eventgrid_publisher:
                yield blob_service, eventgrid_publisher
