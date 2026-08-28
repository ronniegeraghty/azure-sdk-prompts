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


@dataclass(frozen=True)
class AzureSettings:
    storage_account_url: str
    event_grid_topic_endpoint: str

    @classmethod
    def from_environment(cls) -> "AzureSettings":
        return cls(
            storage_account_url=_required_environment_url(
                "AZURE_STORAGE_ACCOUNT_URL", ".blob.core.windows.net"
            ),
            event_grid_topic_endpoint=_required_environment_url(
                "AZURE_EVENTGRID_TOPIC_ENDPOINT", ".eventgrid.azure.net"
            ),
        )


def create_blob_service_client(
    account_url: str, credential: DefaultAzureCredential | None = None
) -> BlobServiceClient:
    return BlobServiceClient(
        account_url=account_url,
        credential=credential or DefaultAzureCredential(),
    )


def create_event_grid_publisher_client(
    endpoint: str, credential: DefaultAzureCredential | None = None
) -> EventGridPublisherClient:
    return EventGridPublisherClient(
        endpoint=endpoint,
        credential=credential or DefaultAzureCredential(),
    )


def create_async_blob_service_client(
    account_url: str, credential: AsyncDefaultAzureCredential | None = None
) -> AsyncBlobServiceClient:
    return AsyncBlobServiceClient(
        account_url=account_url,
        credential=credential or AsyncDefaultAzureCredential(),
    )


def create_async_event_grid_publisher_client(
    endpoint: str, credential: AsyncDefaultAzureCredential | None = None
) -> AsyncEventGridPublisherClient:
    return AsyncEventGridPublisherClient(
        endpoint=endpoint,
        credential=credential or AsyncDefaultAzureCredential(),
    )


def _required_environment_url(name: str, expected_host_suffix: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set")
    if not value.startswith("https://") or expected_host_suffix not in value:
        raise ValueError(f"{name} must be an HTTPS Azure endpoint")
    return value.rstrip("/")
