"""Run an offline demonstration of sync and async event processing."""

from __future__ import annotations

import asyncio
import json
import logging
from dataclasses import dataclass
from typing import Any

from blob_event_notifier.event_publisher import (
    CustomEvent,
    publish_events,
    publish_events_async,
)
from blob_event_notifier.event_receiver import receive_events, receive_events_async

DEMO_ENDPOINT = "https://example-topic.westus2-1.eventgrid.azure.net/api/events"
CREATED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2026/august-001.pdf"
)
DELETED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf"
)


@dataclass
class MockContentSettings:
    content_type: str


@dataclass
class MockBlobProperties:
    size: int
    content_settings: MockContentSettings
    blob_tier: str


class MockDownloader:
    def __init__(self, content: bytes) -> None:
        self._content = content
        self.properties = MockBlobProperties(
            size=len(content),
            content_settings=MockContentSettings("application/pdf"),
            blob_tier="Hot",
        )

    def readall(self) -> bytes:
        return self._content


class AsyncMockDownloader(MockDownloader):
    async def readall(self) -> bytes:
        return self._content


class MockBlobClient:
    def download_blob(self) -> MockDownloader:
        return MockDownloader(b"%PDF-1.7 demo invoice")


class AsyncMockBlobClient:
    async def download_blob(self) -> AsyncMockDownloader:
        return AsyncMockDownloader(b"%PDF-1.7 demo invoice")


class MockBlobServiceClient:
    def get_blob_client(self, **_: Any) -> MockBlobClient:
        return MockBlobClient()


class AsyncMockBlobServiceClient:
    def get_blob_client(self, **_: Any) -> AsyncMockBlobClient:
        return AsyncMockBlobClient()


class MockPublisherClient:
    def send(self, events: list[Any]) -> None:
        logging.info("Mock publisher accepted %d event(s)", len(events))


class AsyncMockPublisherClient:
    async def send(self, events: list[Any]) -> None:
        logging.info("Async mock publisher accepted %d event(s)", len(events))


def sample_payloads() -> tuple[str, str]:
    native_events = [
        {
            "id": "6f159fb5-006e-001b-66f6-75d8ed06f101",
            "topic": (
                "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore"
            ),
            "subject": CREATED_SUBJECT,
            "eventType": "Microsoft.Storage.BlobCreated",
            "eventTime": "2026-08-28T21:45:00Z",
            "data": {
                "api": "PutBlob",
                "clientRequestId": "8b1b282a-9c65-4d5e-9c21-87bdba80a601",
                "requestId": "6f159fb5-006e-001b-66f6-75d8ed000000",
                "eTag": "0x8DC000000000001",
                "contentType": "application/pdf",
                "contentLength": 32145,
                "blobType": "BlockBlob",
                "url": (
                    "https://demostore.blob.core.windows.net/"
                    "documents/invoices/2026/august-001.pdf"
                ),
                "sequencer": "000000000000000000000000000001",
                "storageDiagnostics": {"batchId": "demo-batch-001"},
            },
            "dataVersion": "",
            "metadataVersion": "1",
        }
    ]
    cloud_events = [
        {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobDeleted",
            "source": (
                "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore"
            ),
            "subject": DELETED_SUBJECT,
            "id": "9aeb0fdf-c01e-0131-6b6e-75a90f06ffff",
            "time": "2026-08-28T21:46:00Z",
            "datacontenttype": "application/json",
            "data": {
                "api": "DeleteBlob",
                "requestId": "9aeb0fdf-c01e-0131-6b6e-75a90f000000",
                "url": (
                    "https://demostore.blob.core.windows.net/"
                    "documents/archive/old-invoice.pdf"
                ),
                "sequencer": "000000000000000000000000000002",
            },
        }
    ]
    return json.dumps(native_events), json.dumps(cloud_events)


def downstream_event() -> CustomEvent:
    return CustomEvent(
        event_type="Contoso.Documents.DocumentProcessed",
        subject="/documents/invoices/processed",
        data={"documentName": "august-001.pdf", "status": "processed"},
    )


def run_sync_demo() -> None:
    logging.info("Starting synchronous demo")
    native_payload, cloud_payload = sample_payloads()
    blob_client = MockBlobServiceClient()
    receive_events(native_payload, blob_client)
    receive_events(cloud_payload, blob_client)
    publish_events(
        DEMO_ENDPOINT,
        [downstream_event()],
        client=MockPublisherClient(),
    )


async def run_async_demo() -> None:
    logging.info("Starting asynchronous demo")
    native_payload, cloud_payload = sample_payloads()
    blob_client = AsyncMockBlobServiceClient()
    await receive_events_async(native_payload, blob_client)
    await receive_events_async(cloud_payload, blob_client)
    await publish_events_async(
        DEMO_ENDPOINT,
        [downstream_event()],
        client=AsyncMockPublisherClient(),
    )


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())
