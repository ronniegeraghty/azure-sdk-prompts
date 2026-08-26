from __future__ import annotations

import asyncio
import json
import logging
from dataclasses import dataclass
from types import SimpleNamespace
from typing import Any

from blob_event_handler import AsyncBlobEventHandler, BlobEventHandler
from event_publisher import AsyncEventPublisher, CustomEvent, EventPublisher
from event_receiver import AsyncEventReceiver, EventReceiver, EventSchema

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")

SUBJECT_PREFIX = "/blobServices/default/containers/documents/blobs/"


def event_grid_payload() -> str:
    return json.dumps(
        [
            {
                "id": "8f5ef45a-cd91-4f20-b4f4-76f90c01f844",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": f"{SUBJECT_PREFIX}invoices/2026-08.pdf",
                "eventType": "Microsoft.Storage.BlobCreated",
                "eventTime": "2026-08-26T08:30:00Z",
                "data": {
                    "api": "PutBlob",
                    "contentType": "application/pdf",
                    "contentLength": 2048,
                    "url": "https://demostorage.blob.core.windows.net/"
                    "documents/invoices/2026-08.pdf",
                    "sequencer": "0000000000000000000000000000001",
                },
                "dataVersion": "",
                "metadataVersion": "1",
            },
            {
                "id": "26ed4c1d-b7cb-4829-ab52-ad950719fc51",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": f"{SUBJECT_PREFIX}archive/old-invoice.pdf",
                "eventType": "Microsoft.Storage.BlobDeleted",
                "eventTime": "2026-08-26T08:31:00Z",
                "data": {
                    "api": "DeleteBlob",
                    "url": "https://demostorage.blob.core.windows.net/"
                    "documents/archive/old-invoice.pdf",
                    "sequencer": "0000000000000000000000000000002",
                },
                "dataVersion": "",
                "metadataVersion": "1",
            },
        ]
    )


def cloud_events_payload() -> str:
    return json.dumps(
        [
            {
                "specversion": "1.0",
                "id": "8a94fdb0-97d0-454b-8978-a14f5cbd2571",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": f"{SUBJECT_PREFIX}reports/quarterly.csv",
                "type": "Microsoft.Storage.BlobCreated",
                "time": "2026-08-26T08:32:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "PutBlob",
                    "contentType": "text/csv",
                    "contentLength": 128,
                    "url": "https://demostorage.blob.core.windows.net/"
                    "documents/reports/quarterly.csv",
                    "sequencer": "0000000000000000000000000000003",
                },
            },
            {
                "specversion": "1.0",
                "id": "0e490515-7349-4e70-99f7-680908ba35d0",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/"
                "resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": f"{SUBJECT_PREFIX}reports/draft.csv",
                "type": "Microsoft.Storage.BlobDeleted",
                "time": "2026-08-26T08:33:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "DeleteBlob",
                    "url": "https://demostorage.blob.core.windows.net/"
                    "documents/reports/draft.csv",
                    "sequencer": "0000000000000000000000000000004",
                },
            },
        ]
    )


@dataclass
class DemoDownloader:
    content: bytes
    properties: Any

    def readall(self) -> bytes:
        return self.content


class DemoBlobClient:
    def __init__(self, name: str) -> None:
        content_type = "text/csv" if name.endswith(".csv") else "application/pdf"
        self._download = DemoDownloader(
            content=b"demo blob content",
            properties=SimpleNamespace(
                size=len(b"demo blob content"),
                content_settings=SimpleNamespace(content_type=content_type),
                blob_tier="Hot",
            ),
        )

    def download_blob(self) -> DemoDownloader:
        return self._download


class DemoBlobService:
    def get_blob_client(self, container: str, blob: str) -> DemoBlobClient:
        logging.getLogger(__name__).debug("Reading %s/%s", container, blob)
        return DemoBlobClient(blob)


class DemoPublisherClient:
    def send(self, events: list[Any]) -> None:
        print(f"Published downstream event(s): {[event.subject for event in events]}")


class AsyncDemoDownloader(DemoDownloader):
    async def readall(self) -> bytes:
        return self.content


class AsyncDemoBlobClient(DemoBlobClient):
    async def download_blob(self) -> AsyncDemoDownloader:
        return AsyncDemoDownloader(self._download.content, self._download.properties)


class AsyncDemoBlobService:
    def get_blob_client(self, container: str, blob: str) -> AsyncDemoBlobClient:
        logging.getLogger(__name__).debug("Reading %s/%s", container, blob)
        return AsyncDemoBlobClient(blob)


class AsyncDemoPublisherClient:
    async def send(self, events: list[Any]) -> None:
        print(f"Published downstream event(s): {[event.subject for event in events]}")


def downstream_event() -> CustomEvent:
    return CustomEvent(
        event_type="Contoso.Documents.DocumentProcessed",
        subject="/documents/invoices/processed",
        data={"document": "invoices/2026-08.pdf", "status": "processed"},
    )


def run_sync_demo() -> None:
    print("=== Sync implementation ===")
    receiver = EventReceiver(BlobEventHandler(DemoBlobService()))
    receiver.receive(event_grid_payload(), EventSchema.EVENT_GRID)
    receiver.receive(cloud_events_payload(), EventSchema.CLOUD_EVENTS)
    publisher = EventPublisher(
        "https://example.invalid/api/events",
        client=DemoPublisherClient(),
    )
    publisher.publish([downstream_event()])


async def run_async_demo() -> None:
    print("=== Async implementation ===")
    receiver = AsyncEventReceiver(AsyncBlobEventHandler(AsyncDemoBlobService()))
    await receiver.receive(event_grid_payload(), EventSchema.EVENT_GRID)
    await receiver.receive(cloud_events_payload(), EventSchema.CLOUD_EVENTS)
    publisher = AsyncEventPublisher(
        "https://example.invalid/api/events",
        client=AsyncDemoPublisherClient(),
    )
    await publisher.publish([downstream_event()])


if __name__ == "__main__":
    run_sync_demo()
    asyncio.run(run_async_demo())
