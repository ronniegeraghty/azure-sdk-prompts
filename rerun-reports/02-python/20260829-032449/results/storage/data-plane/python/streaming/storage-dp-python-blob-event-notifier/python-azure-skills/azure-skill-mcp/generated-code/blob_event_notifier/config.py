"""Passwordless Azure SDK client configuration."""

from __future__ import annotations

import os

from azure.eventgrid import EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient


def storage_account_url() -> str:
    return os.environ["AZURE_STORAGE_ACCOUNT_URL"]


def event_grid_topic_endpoint() -> str:
    return os.environ["AZURE_EVENT_GRID_TOPIC_ENDPOINT"]


def create_blob_service_client() -> BlobServiceClient:
    return BlobServiceClient(
        account_url=storage_account_url(),
        credential=DefaultAzureCredential(),
    )


def create_event_grid_publisher_client(
    endpoint: str | None = None,
) -> EventGridPublisherClient:
    return EventGridPublisherClient(
        endpoint=endpoint or event_grid_topic_endpoint(),
        credential=DefaultAzureCredential(),
    )


def create_async_blob_service_client() -> AsyncBlobServiceClient:
    return AsyncBlobServiceClient(
        account_url=storage_account_url(),
        credential=AsyncDefaultAzureCredential(),
    )


def create_async_event_grid_publisher_client(
    endpoint: str | None = None,
) -> AsyncEventGridPublisherClient:
    return AsyncEventGridPublisherClient(
        endpoint=endpoint or event_grid_topic_endpoint(),
        credential=AsyncDefaultAzureCredential(),
    )
