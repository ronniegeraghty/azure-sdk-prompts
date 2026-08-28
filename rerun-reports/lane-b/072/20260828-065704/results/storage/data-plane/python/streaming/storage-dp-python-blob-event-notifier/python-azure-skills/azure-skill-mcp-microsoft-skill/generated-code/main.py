from __future__ import annotations

import asyncio
import json
import logging
from types import SimpleNamespace
from typing import Any, Sequence

from azure.eventgrid import EventGridEvent

from blob_event_notifier.blob_event_handler import AsyncBlobEventHandler, BlobEventHandler
from blob_event_notifier.event_publisher import (
    AsyncEventPublisher,
    CustomEvent,
    EventPublisher,
)
from blob_event_notifier.event_receiver import receive_events, receive_events_async

CREATED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1001.pdf"
)
DELETED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/2025/invoice-0999.pdf"
)

EVENTGRID_PAYLOAD = json.dumps(
    [
        {
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
            "subject": CREATED_SUBJECT,
            "eventType": "Microsoft.Storage.BlobCreated",
            "id": "11111111-1111-1111-1111-111111111111",
            "data": {
                "api": "PutBlob",
                "clientRequestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
                "requestId": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
                "eTag": "0x8DC000000000000",
                "contentType": "application/pdf",
                "contentLength": 27,
                "blobType": "BlockBlob",
                "url": "https://demostorage.blob.core.windows.net/documents/"
                "invoices/2026/invoice-1001.pdf",
                "sequencer": "000000000000000000000000000000000000000000000001",
                "storageDiagnostics": {"batchId": "cccccccc-cccc-cccc-cccc-cccccccccccc"},
            },
            "dataVersion": "",
            "metadataVersion": "1",
            "eventTime": "2026-08-28T00:00:00Z",
        },
        {
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
            "subject": DELETED_SUBJECT,
            "eventType": "Microsoft.Storage.BlobDeleted",
            "id": "22222222-2222-2222-2222-222222222222",
            "data": {
                "api": "DeleteBlob",
                "url": "https://demostorage.blob.core.windows.net/documents/"
                "invoices/2025/invoice-0999.pdf",
                "sequencer": "000000000000000000000000000000000000000000000002",
            },
            "dataVersion": "",
            "metadataVersion": "1",
            "eventTime": "2026-08-28T00:01:00Z",
        },
    ]
)

CLOUDEVENT_PAYLOAD = json.dumps(
    [
        {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobCreated",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
            "subject": CREATED_SUBJECT,
            "id": "33333333-3333-3333-3333-333333333333",
            "time": "2026-08-28T00:02:00Z",
            "datacontenttype": "application/json",
            "data": {
                "api": "PutBlob",
                "contentType": "application/pdf",
                "contentLength": 27,
                "blobType": "BlockBlob",
                "url": "https://demostorage.blob.core.windows.net/documents/"
                "invoices/2026/invoice-1001.pdf",
            },
        },
        {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobDeleted",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
            "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
            "subject": DELETED_SUBJECT,
            "id": "44444444-4444-4444-4444-444444444444",
            "time": "2026-08-28T00:03:00Z",
            "datacontenttype": "application/json",
            "data": {
                "api": "DeleteBlob",
                "url": "https://demostorage.blob.core.windows.net/documents/"
                "invoices/2025/invoice-0999.pdf",
            },
        },
    ]
)


class DemoDownloader:
    def __init__(self, content: bytes) -> None:
        self._content = content
        self.properties = SimpleNamespace(
            size=len(content),
            content_settings=SimpleNamespace(content_type="application/pdf"),
            blob_tier="Hot",
        )

    def readall(self) -> bytes:
        return self._content


class DemoBlobClient:
    def download_blob(self) -> DemoDownloader:
        return DemoDownloader(b"%PDF-1.7 demo invoice data")


class DemoBlobService:
    def get_blob_client(self, container: str, blob: str) -> DemoBlobClient:
        print(f"Downloading mock blob from {container}/{blob}")
        return DemoBlobClient()


class DemoPublisherClient:
    def send(self, events: Sequence[EventGridEvent]) -> None:
        for event in events:
            print(f"Published mock event: type={event.event_type}, subject={event.subject}")


class AsyncDemoDownloader(DemoDownloader):
    async def readall(self) -> bytes:
        return self._content


class AsyncDemoBlobClient:
    async def download_blob(self) -> AsyncDemoDownloader:
        return AsyncDemoDownloader(b"%PDF-1.7 demo invoice data")


class AsyncDemoBlobService:
    def get_blob_client(self, container: str, blob: str) -> AsyncDemoBlobClient:
        print(f"Downloading mock blob from {container}/{blob}")
        return AsyncDemoBlobClient()


class AsyncDemoPublisherClient:
    async def send(self, events: Sequence[EventGridEvent]) -> None:
        for event in events:
            print(f"Published mock event: type={event.event_type}, subject={event.subject}")


DOWNSTREAM_EVENT = CustomEvent(
    event_type="Contoso.Documents.DocumentProcessed",
    data={"documentId": "invoice-1001", "status": "processed"},
)
DOWNSTREAM_SUBJECT = "/documents/invoices/processed"


def run_sync_demo() -> None:
    print("\n=== Sync demo ===")
    handler = BlobEventHandler(DemoBlobService())
    receive_events(EVENTGRID_PAYLOAD, handler)
    receive_events(CLOUDEVENT_PAYLOAD, handler)
    EventPublisher(DemoPublisherClient()).publish(
        [DOWNSTREAM_EVENT],
        default_subject=DOWNSTREAM_SUBJECT,
    )


async def run_async_demo() -> None:
    print("\n=== Async demo ===")
    handler = AsyncBlobEventHandler(AsyncDemoBlobService())
    await receive_events_async(EVENTGRID_PAYLOAD, handler)
    await receive_events_async(CLOUDEVENT_PAYLOAD, handler)
    await AsyncEventPublisher(AsyncDemoPublisherClient()).publish(
        [DOWNSTREAM_EVENT],
        default_subject=DOWNSTREAM_SUBJECT,
    )


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
