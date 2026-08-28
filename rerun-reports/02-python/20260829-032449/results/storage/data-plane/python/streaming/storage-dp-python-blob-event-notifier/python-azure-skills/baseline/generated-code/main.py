"""Offline demonstration of synchronous and asynchronous event processing."""

from __future__ import annotations

import asyncio
import logging
from dataclasses import dataclass
from typing import Any

from event_publisher import (
    CustomEvent,
    publish_custom_events,
    publish_custom_events_async,
)
from event_receiver import receive_events, receive_events_async

MOCK_TOPIC_ENDPOINT = "https://example-topic.eastus-1.eventgrid.azure.net/api/events"

EVENT_GRID_PAYLOAD = [
    {
        "id": "4d8e1d41-52c5-4c24-a654-00b4df095a30",
        "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
        "subject": "/blobServices/default/containers/documents/blobs/invoices/2026-08.pdf",
        "eventType": "Microsoft.Storage.BlobCreated",
        "eventTime": "2026-08-29T00:00:00Z",
        "data": {
            "api": "PutBlob",
            "clientRequestId": "28c5b5af-5474-4a40-92ca-577456c1c2b8",
            "requestId": "00000000-0000-0000-0000-000000000000",
            "eTag": "0x8DE000000000000",
            "contentType": "application/pdf",
            "contentLength": 24576,
            "blobType": "BlockBlob",
            "accessTier": "Hot",
            "url": "https://demostorage.blob.core.windows.net/documents/invoices/2026-08.pdf",
            "sequencer": "000000000000000000000000000000010000000000000000",
            "storageDiagnostics": {"batchId": "6cdb9ea9-a006-006b-0085-ef6ed6000000"},
        },
        "dataVersion": "",
        "metadataVersion": "1",
    },
    {
        "id": "e12e103e-0fb8-4e2c-a144-40f0108f814e",
        "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
        "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
        "eventType": "Microsoft.Storage.BlobDeleted",
        "eventTime": "2026-08-29T00:01:00Z",
        "data": {
            "api": "DeleteBlob",
            "clientRequestId": "3d231a72-e83d-48bc-b06f-3d19ec592fc2",
            "requestId": "00000000-0000-0000-0000-000000000001",
            "contentType": "application/pdf",
            "blobType": "BlockBlob",
            "url": "https://demostorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
            "sequencer": "000000000000000000000000000000020000000000000000",
        },
        "dataVersion": "",
        "metadataVersion": "1",
    },
]

CLOUD_EVENTS_PAYLOAD = [
    {
        "specversion": "1.0",
        "id": "5b394789-53e5-4f88-a64a-c9ba5326b77e",
        "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
        "subject": "/blobServices/default/containers/documents/blobs/reports/summary.txt",
        "type": "Microsoft.Storage.BlobCreated",
        "time": "2026-08-29T00:02:00Z",
        "data": {
            "api": "PutBlob",
            "contentType": "text/plain",
            "contentLength": 1024,
            "blobType": "BlockBlob",
            "accessTier": "Cool",
            "url": "https://demostorage.blob.core.windows.net/documents/reports/summary.txt",
        },
    },
    {
        "specversion": "1.0",
        "id": "a9bc6301-ef98-4285-bd1c-daef9d000bbc",
        "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
        "subject": "/blobServices/default/containers/documents/blobs/reports/draft.txt",
        "type": "Microsoft.Storage.BlobDeleted",
        "time": "2026-08-29T00:03:00Z",
        "data": {
            "api": "DeleteBlob",
            "contentType": "text/plain",
            "blobType": "BlockBlob",
            "url": "https://demostorage.blob.core.windows.net/documents/reports/draft.txt",
        },
    },
]


@dataclass
class FakeContentSettings:
    content_type: str


@dataclass
class FakeBlobProperties:
    size: int
    content_settings: FakeContentSettings
    blob_tier: str


class FakeDownloader:
    def readall(self) -> bytes:
        return b"offline demo content"


class AsyncFakeDownloader:
    async def readall(self) -> bytes:
        return b"offline demo content"


class FakeBlobClient:
    def __init__(self, name: str) -> None:
        self.name = name

    def get_blob_properties(self) -> FakeBlobProperties:
        content_type = "application/pdf" if self.name.endswith(".pdf") else "text/plain"
        return FakeBlobProperties(24576, FakeContentSettings(content_type), "Hot")

    def download_blob(self) -> FakeDownloader:
        return FakeDownloader()


class AsyncFakeBlobClient(FakeBlobClient):
    async def get_blob_properties(self) -> FakeBlobProperties:
        return super().get_blob_properties()

    async def download_blob(self) -> AsyncFakeDownloader:
        return AsyncFakeDownloader()


class FakeBlobService:
    def get_blob_client(self, container: str, blob: str) -> FakeBlobClient:
        return FakeBlobClient(blob)


class AsyncFakeBlobService:
    def get_blob_client(self, container: str, blob: str) -> AsyncFakeBlobClient:
        return AsyncFakeBlobClient(blob)


class FakePublisher:
    def send(self, events: list[Any]) -> None:
        print(f"Published {len(events)} downstream event(s) synchronously")


class AsyncFakePublisher:
    async def send(self, events: list[Any]) -> None:
        print(f"Published {len(events)} downstream event(s) asynchronously")


def downstream_event() -> CustomEvent:
    return CustomEvent(
        event_type="Contoso.Documents.Processed",
        subject="/documents/invoices/processed",
        data={"document": "2026-08.pdf", "status": "processed"},
    )


def run_sync_demo() -> None:
    print("=== Synchronous demo ===")
    blob_service = FakeBlobService()
    receive_events(EVENT_GRID_PAYLOAD, blob_service)
    receive_events(CLOUD_EVENTS_PAYLOAD, blob_service)
    publish_custom_events(
        MOCK_TOPIC_ENDPOINT,
        [downstream_event()],
        client=FakePublisher(),
    )


async def run_async_demo() -> None:
    print("\n=== Asynchronous demo ===")
    blob_service = AsyncFakeBlobService()
    await receive_events_async(EVENT_GRID_PAYLOAD, blob_service)
    await receive_events_async(CLOUD_EVENTS_PAYLOAD, blob_service)
    await publish_custom_events_async(
        MOCK_TOPIC_ENDPOINT,
        [downstream_event()],
        client=AsyncFakePublisher(),
    )


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
