from __future__ import annotations

import asyncio
import json
import logging
from contextlib import asynccontextmanager, contextmanager
from dataclasses import dataclass, field
from typing import Any

from blob_event_handler import AsyncBlobEventHandler, BlobEventHandler
from event_publisher import CustomEvent, publish_events, publish_events_async
from event_receiver import receive_events, receive_events_async

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")

DEMO_TOPIC_ENDPOINT = "https://example-topic.eastus-1.eventgrid.azure.net/api/events"
BLOB_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1001.pdf"
)


def sample_payloads() -> tuple[str, str]:
    event_grid_payload = json.dumps(
        [
            {
                "id": "f3a8c2ce-3a1d-4a31-a5d5-111111111111",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
                "subject": BLOB_SUBJECT,
                "eventType": "Microsoft.Storage.BlobCreated",
                "eventTime": "2026-08-29T00:00:00Z",
                "data": {
                    "api": "PutBlob",
                    "clientRequestId": "11111111-1111-1111-1111-111111111111",
                    "requestId": "22222222-2222-2222-2222-222222222222",
                    "eTag": "0x8DC000000000001",
                    "contentType": "application/pdf",
                    "contentLength": 2048,
                    "blobType": "BlockBlob",
                    "url": "https://demostore.blob.core.windows.net/documents/"
                    "invoices/2026/invoice-1001.pdf",
                    "sequencer": "000000000000000000000000000000010000000000000001",
                    "storageDiagnostics": {"batchId": "33333333-3333-3333-3333-333333333333"},
                },
                "dataVersion": "",
                "metadataVersion": "1",
            },
            {
                "id": "f3a8c2ce-3a1d-4a31-a5d5-222222222222",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
                "subject": BLOB_SUBJECT,
                "eventType": "Microsoft.Storage.BlobDeleted",
                "eventTime": "2026-08-29T00:01:00Z",
                "data": {
                    "api": "DeleteBlob",
                    "clientRequestId": "44444444-4444-4444-4444-444444444444",
                    "requestId": "55555555-5555-5555-5555-555555555555",
                    "url": "https://demostore.blob.core.windows.net/documents/"
                    "invoices/2026/invoice-1001.pdf",
                    "sequencer": "000000000000000000000000000000020000000000000001",
                },
                "dataVersion": "",
                "metadataVersion": "1",
            },
        ]
    )

    cloud_events_payload = json.dumps(
        [
            {
                "specversion": "1.0",
                "type": "Microsoft.Storage.BlobCreated",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
                "subject": BLOB_SUBJECT,
                "id": "6c3f5898-8622-4db6-9aaa-111111111111",
                "time": "2026-08-29T00:02:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "PutBlob",
                    "contentType": "application/pdf",
                    "contentLength": 2048,
                    "blobType": "BlockBlob",
                    "url": "https://demostore.blob.core.windows.net/documents/"
                    "invoices/2026/invoice-1001.pdf",
                },
            },
            {
                "specversion": "1.0",
                "type": "Microsoft.Storage.BlobDeleted",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
                "subject": BLOB_SUBJECT,
                "id": "6c3f5898-8622-4db6-9aaa-222222222222",
                "time": "2026-08-29T00:03:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "DeleteBlob",
                    "url": "https://demostore.blob.core.windows.net/documents/"
                    "invoices/2026/invoice-1001.pdf",
                },
            },
        ]
    )
    return event_grid_payload, cloud_events_payload


@dataclass
class _ContentSettings:
    content_type: str = "application/pdf"


@dataclass
class _BlobProperties:
    size: int = 2048
    content_settings: _ContentSettings = field(default_factory=_ContentSettings)
    blob_tier: str = "Hot"


class _Download:
    def readall(self) -> bytes:
        return b"%PDF-demo"


class _AsyncDownload:
    async def readall(self) -> bytes:
        return b"%PDF-demo"


class _BlobClient:
    def get_blob_properties(self) -> _BlobProperties:
        return _BlobProperties()

    def download_blob(self) -> _Download:
        return _Download()


class _AsyncBlobClient:
    async def get_blob_properties(self) -> _BlobProperties:
        return _BlobProperties()

    async def download_blob(self) -> _AsyncDownload:
        return _AsyncDownload()


class _BlobService:
    def get_blob_client(self, container: str, blob: str) -> _BlobClient:
        return _BlobClient()


class _AsyncBlobService:
    def get_blob_client(self, container: str, blob: str) -> _AsyncBlobClient:
        return _AsyncBlobClient()


class _Publisher:
    def send(self, events: list[Any]) -> None:
        print(f"Published {len(events)} downstream event(s) locally")


class _AsyncPublisher:
    async def send(self, events: list[Any]) -> None:
        print(f"Published {len(events)} downstream event(s) locally (async)")


@contextmanager
def _local_publisher(endpoint: str):
    yield _Publisher()


@asynccontextmanager
async def _local_async_publisher(endpoint: str):
    yield _AsyncPublisher()


def downstream_events() -> list[CustomEvent]:
    return [
        CustomEvent(
            event_type="Contoso.Documents.DocumentProcessed",
            subject="/documents/invoices/processed",
            data={"documentId": "invoice-1001", "status": "processed"},
        )
    ]


def run_sync_demo() -> None:
    print("\n--- Synchronous demo ---")
    handler = BlobEventHandler(_BlobService())
    for payload in sample_payloads():
        receive_events(payload, handler.handle_created, handler.handle_deleted)
    publish_events(
        DEMO_TOPIC_ENDPOINT,
        downstream_events(),
        client_context=_local_publisher,
    )


async def run_async_demo() -> None:
    print("\n--- Asynchronous demo ---")
    handler = AsyncBlobEventHandler(_AsyncBlobService())
    for payload in sample_payloads():
        await receive_events_async(
            payload, handler.handle_created, handler.handle_deleted
        )
    await publish_events_async(
        DEMO_TOPIC_ENDPOINT,
        downstream_events(),
        client_context=_local_async_publisher,
    )


if __name__ == "__main__":
    run_sync_demo()
    asyncio.run(run_async_demo())
