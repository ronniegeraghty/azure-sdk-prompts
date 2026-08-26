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
            storage_account_url=_required_https_url("AZURE_STORAGE_ACCOUNT_URL"),
            event_grid_topic_endpoint=_required_https_url(
                "EVENTGRID_TOPIC_ENDPOINT"
            ),
        )


def _required_https_url(variable_name: str) -> str:
    value = os.getenv(variable_name)
    if not value:
        raise ValueError(f"Required environment variable {variable_name} is not set")
    if not value.startswith("https://"):
        raise ValueError(f"{variable_name} must use HTTPS")
    return value


@contextmanager
def sync_blob_client(account_url: str) -> Iterator[BlobServiceClient]:
    with DefaultAzureCredential() as credential:
        with BlobServiceClient(account_url, credential=credential) as client:
            yield client


@asynccontextmanager
async def async_blob_client(
    account_url: str,
) -> AsyncIterator[AsyncBlobServiceClient]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncBlobServiceClient(
            account_url, credential=credential
        ) as client:
            yield client


@contextmanager
def sync_event_grid_client(
    endpoint: str,
) -> Iterator[EventGridPublisherClient]:
    with DefaultAzureCredential() as credential:
        with EventGridPublisherClient(endpoint, credential) as client:
            yield client


@asynccontextmanager
async def async_event_grid_client(
    endpoint: str,
) -> AsyncIterator[AsyncEventGridPublisherClient]:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncEventGridPublisherClient(endpoint, credential) as client:
            yield client
