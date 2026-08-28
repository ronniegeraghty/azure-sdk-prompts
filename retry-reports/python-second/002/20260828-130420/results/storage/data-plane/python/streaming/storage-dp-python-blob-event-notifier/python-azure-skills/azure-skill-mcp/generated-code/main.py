"""Offline demonstration of synchronous and asynchronous event processing."""

from __future__ import annotations

import asyncio
import json
import logging
from dataclasses import dataclass
from typing import Any

from blob_event_handler import AsyncBlobEventHandler, BlobEventHandler
from event_publisher import AsyncEventPublisher, CustomEvent, EventPublisher
from event_receiver import AsyncEventReceiver, EventReceiver

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")

CREATED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1042.pdf"
)
DELETED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2025/invoice-0087.pdf"
)

EVENT_GRID_PAYLOAD = json.dumps(
    [
        {
            "id": "0f47f202-b4b8-4a87-a72d-42c4a59324af",
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "subject": CREATED_SUBJECT,
            "eventType": "Microsoft.Storage.BlobCreated",
            "eventTime": "2026-08-28T04:55:02.123456Z",
            "data": {
                "api": "PutBlob",
                "clientRequestId": "f87e01ad-22e4-4a8f-938f-3d8e05467f36",
                "requestId": "6f19d1c1-901e-0024-2a85-bfea60000000",
                "eTag": "0x8DC000000000001",
                "contentType": "application/pdf",
                "contentLength": 28,
                "blobType": "BlockBlob",
                "url": "https://demostore.blob.core.windows.net/documents/"
                "invoices/2026/invoice-1042.pdf",
                "sequencer": "000000000000000000000000000000010000000000000001",
                "storageDiagnostics": {"batchId": "0a9a3f63-18c4-4f70-8d32-3bf986ae973f"},
            },
            "dataVersion": "",
            "metadataVersion": "1",
        },
        {
            "id": "635e8504-4dc0-45b7-9818-90b325690b96",
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "subject": DELETED_SUBJECT,
            "eventType": "Microsoft.Storage.BlobDeleted",
            "eventTime": "2026-08-28T04:56:10.123456Z",
            "data": {
                "api": "DeleteBlob",
                "requestId": "4b2d6ce9-201e-0017-5a85-bf6d4e000000",
                "blobType": "BlockBlob",
                "url": "https://demostore.blob.core.windows.net/documents/"
                "invoices/2025/invoice-0087.pdf",
                "sequencer": "000000000000000000000000000000020000000000000001",
            },
            "dataVersion": "",
            "metadataVersion": "1",
        },
    ]
)

CLOUD_EVENTS_PAYLOAD = json.dumps(
    [
        {
            "specversion": "1.0",
            "id": "74414088-4b23-4bb8-8069-66c2f9f93fc7",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "subject": CREATED_SUBJECT,
            "type": "Microsoft.Storage.BlobCreated",
            "time": "2026-08-28T04:57:12.123456Z",
            "datacontenttype": "application/json",
            "data": {
                "api": "PutBlob",
                "contentType": "application/pdf",
                "contentLength": 28,
                "blobType": "BlockBlob",
                "url": "https://demostore.blob.core.windows.net/documents/"
                "invoices/2026/invoice-1042.pdf",
            },
        },
        {
            "specversion": "1.0",
            "id": "e5867778-40a7-47f4-8942-f5a366d06bc3",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "subject": DELETED_SUBJECT,
            "type": "Microsoft.Storage.BlobDeleted",
            "time": "2026-08-28T04:58:16.123456Z",
            "datacontenttype": "application/json",
            "data": {
                "api": "DeleteBlob",
                "blobType": "BlockBlob",
                "url": "https://demostore.blob.core.windows.net/documents/"
                "invoices/2025/invoice-0087.pdf",
            },
        },
    ]
)


@dataclass
class _ContentSettings:
    content_type: str


@dataclass
class _BlobProperties:
    content_settings: _ContentSettings
    blob_tier: str


class _Download:
    def readall(self) -> bytes:
        return b"%PDF-1.7 mock invoice data"


class _BlobClient:
    def get_blob_properties(self) -> _BlobProperties:
        return _BlobProperties(_ContentSettings("application/pdf"), "Hot")

    def download_blob(self) -> _Download:
        return _Download()


class _BlobService:
    def get_blob_client(self, container: str, blob: str) -> _BlobClient:
        del container, blob
        return _BlobClient()


class _PublisherClient:
    def send(self, events: list[Any]) -> None:
        print(f"Mock published {len(events)} downstream event(s)")


class _AsyncDownload:
    async def readall(self) -> bytes:
        return b"%PDF-1.7 mock invoice data"


class _AsyncBlobClient:
    async def get_blob_properties(self) -> _BlobProperties:
        return _BlobProperties(_ContentSettings("application/pdf"), "Hot")

    async def download_blob(self) -> _AsyncDownload:
        return _AsyncDownload()


class _AsyncBlobService:
    def get_blob_client(self, container: str, blob: str) -> _AsyncBlobClient:
        del container, blob
        return _AsyncBlobClient()


class _AsyncPublisherClient:
    async def send(self, events: list[Any]) -> None:
        print(f"Mock published {len(events)} downstream event(s)")


def downstream_event() -> CustomEvent:
    return CustomEvent(
        event_type="Contoso.Documents.Processed",
        subject="/documents/invoices/processed",
        data={"document": "invoice-1042.pdf", "status": "processed"},
    )


def run_sync_demo() -> None:
    print("\n--- synchronous demo ---")
    handler = BlobEventHandler(_BlobService())
    receiver = EventReceiver(handler.handle_created, handler.handle_deleted)
    receiver.receive(EVENT_GRID_PAYLOAD)
    receiver.receive(CLOUD_EVENTS_PAYLOAD)
    EventPublisher(_PublisherClient()).publish([downstream_event()])


async def run_async_demo() -> None:
    print("\n--- asynchronous demo ---")
    handler = AsyncBlobEventHandler(_AsyncBlobService())
    receiver = AsyncEventReceiver(handler.handle_created, handler.handle_deleted)
    await receiver.receive(EVENT_GRID_PAYLOAD)
    await receiver.receive(CLOUD_EVENTS_PAYLOAD)
    await AsyncEventPublisher(_AsyncPublisherClient()).publish([downstream_event()])


if __name__ == "__main__":
    run_sync_demo()
    asyncio.run(run_async_demo())
