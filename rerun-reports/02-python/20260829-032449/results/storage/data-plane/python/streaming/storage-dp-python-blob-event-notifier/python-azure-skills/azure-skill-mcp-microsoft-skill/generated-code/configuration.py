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
    event_grid_topic_endpoint: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        return cls(
            storage_account_url=os.environ["AZURE_STORAGE_ACCOUNT_URL"],
            event_grid_topic_endpoint=os.environ["EVENTGRID_TOPIC_ENDPOINT"],
        )


@contextmanager
def open_blob_service_client(account_url: str) -> Iterator[BlobServiceClient]:
    with DefaultAzureCredential() as credential:
        with BlobServiceClient(account_url, credential=credential) as client:
            yield client


@asynccontextmanager
async def open_async_blob_service_client(
    account_url: str,
) -> AsyncIterator[AsyncBlobServiceClient]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncBlobServiceClient(
            account_url, credential=credential
        ) as client:
            yield client


@contextmanager
def open_event_grid_publisher(
    endpoint: str,
) -> Iterator[EventGridPublisherClient]:
    with DefaultAzureCredential() as credential:
        with EventGridPublisherClient(endpoint, credential) as client:
            yield client


@asynccontextmanager
async def open_async_event_grid_publisher(
    endpoint: str,
) -> AsyncIterator[AsyncEventGridPublisherClient]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncEventGridPublisherClient(endpoint, credential) as client:
            yield client
