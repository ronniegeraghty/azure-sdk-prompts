"""Local demonstration of sync and async Blob lifecycle event processing."""

from __future__ import annotations

import asyncio
import logging
from dataclasses import dataclass
from types import SimpleNamespace
from typing import Any

from blob_event_notifier import (
    AsyncBlobEventHandler,
    AsyncEventPublisher,
    AsyncEventReceiver,
    BlobEventHandler,
    CustomEvent,
    EventPublisher,
    EventReceiver,
)

TOPIC_ENDPOINT = "https://example-topic.westus2-1.eventgrid.azure.net/api/events"

EVENT_GRID_CREATED = """{
  "id": "7b233c13-8e1f-4a30-81f8-8f410fd3e1b7",
  "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
  "subject": "/blobServices/default/containers/documents/blobs/invoices/2026-08.pdf",
  "eventType": "Microsoft.Storage.BlobCreated",
  "eventTime": "2026-08-28T03:45:12.123Z",
  "data": {
    "api": "PutBlob",
    "clientRequestId": "c8f50454-b068-4b13-b263-4c0ab526cd41",
    "requestId": "9a5e628f-201e-0065-778b-8ff388000000",
    "eTag": "0x8DC000000000001",
    "contentType": "application/pdf",
    "contentLength": 24576,
    "blobType": "BlockBlob",
    "url": "https://demostorage.blob.core.windows.net/documents/invoices/2026-08.pdf",
    "sequencer": "000000000000000000000000000000010000000000000001"
  },
  "dataVersion": "1",
  "metadataVersion": "1"
}"""

EVENT_GRID_DELETED = """{
  "id": "4fc94e2c-71ff-4695-a62e-9cfbb1f9152a",
  "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
  "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
  "eventType": "Microsoft.Storage.BlobDeleted",
  "eventTime": "2026-08-28T03:46:00.000Z",
  "data": {
    "api": "DeleteBlob",
    "url": "https://demostorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
    "blobType": "BlockBlob"
  },
  "dataVersion": "1",
  "metadataVersion": "1"
}"""

CLOUD_EVENT_CREATED = """{
  "specversion": "1.0",
  "id": "1576d48f-bc8b-4a25-8d05-176f1d862af1",
  "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
  "type": "Microsoft.Storage.BlobCreated",
  "subject": "/blobServices/default/containers/documents/blobs/contracts/vendor-agreement.docx",
  "time": "2026-08-28T04:00:00.000Z",
  "datacontenttype": "application/json",
  "data": {
    "api": "PutBlob",
    "contentType": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
    "contentLength": 32768,
    "blobType": "BlockBlob",
    "url": "https://demostorage.blob.core.windows.net/documents/contracts/vendor-agreement.docx"
  }
}"""

CLOUD_EVENT_DELETED = """{
  "specversion": "1.0",
  "id": "87fb086e-89d1-4070-9f72-ec5e29b90162",
  "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
  "type": "Microsoft.Storage.BlobDeleted",
  "subject": "/blobServices/default/containers/documents/blobs/contracts/draft.docx",
  "time": "2026-08-28T04:01:00.000Z",
  "datacontenttype": "application/json",
  "data": {
    "api": "DeleteBlob",
    "url": "https://demostorage.blob.core.windows.net/documents/contracts/draft.docx"
  }
}"""


@dataclass
class _FakeDownload:
    content: bytes

    def readall(self) -> bytes:
        return self.content


class _FakeBlobClient:
    def __init__(self, name: str) -> None:
        self._name = name
        self._content = f"mock content for {name}".encode()

    def download_blob(self) -> _FakeDownload:
        return _FakeDownload(self._content)

    def get_blob_properties(self) -> Any:
        content_type = "application/pdf" if self._name.endswith(".pdf") else (
            "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        )
        return SimpleNamespace(
            size=len(self._content),
            content_settings=SimpleNamespace(content_type=content_type),
            blob_tier="Hot",
        )


class _FakeBlobServiceClient:
    def get_blob_client(self, container: str, blob: str) -> _FakeBlobClient:
        return _FakeBlobClient(blob)


class _FakePublisherClient:
    def send(self, events: list[Any]) -> None:
        for event in events:
            print(f"Published: type={event.event_type}, subject={event.subject}")

    def close(self) -> None:
        pass


class _AsyncFakeDownload(_FakeDownload):
    async def readall(self) -> bytes:
        return self.content


class _AsyncFakeBlobClient(_FakeBlobClient):
    async def download_blob(self) -> _AsyncFakeDownload:
        return _AsyncFakeDownload(self._content)

    async def get_blob_properties(self) -> Any:
        return super().get_blob_properties()


class _AsyncFakeBlobServiceClient:
    def get_blob_client(self, container: str, blob: str) -> _AsyncFakeBlobClient:
        return _AsyncFakeBlobClient(blob)


class _AsyncFakePublisherClient:
    async def send(self, events: list[Any]) -> None:
        for event in events:
            print(f"Published async: type={event.event_type}, subject={event.subject}")

    async def close(self) -> None:
        pass


def _downstream_event() -> CustomEvent:
    return CustomEvent(
        event_type="Contoso.Documents.DocumentProcessed",
        subject="/documents/invoices/processed",
        data={"documentId": "2026-08", "status": "processed"},
    )


def run_sync_demo() -> None:
    print("=== Sync implementation ===")
    receiver = EventReceiver(BlobEventHandler(_FakeBlobServiceClient()))
    for payload in (
        EVENT_GRID_CREATED,
        EVENT_GRID_DELETED,
        CLOUD_EVENT_CREATED,
        CLOUD_EVENT_DELETED,
    ):
        receiver.receive(payload)

    publisher = EventPublisher(TOPIC_ENDPOINT, client=_FakePublisherClient())
    publisher.publish([_downstream_event()])
    publisher.close()


async def run_async_demo() -> None:
    print("\n=== Async implementation ===")
    receiver = AsyncEventReceiver(AsyncBlobEventHandler(_AsyncFakeBlobServiceClient()))
    for payload in (
        EVENT_GRID_CREATED,
        EVENT_GRID_DELETED,
        CLOUD_EVENT_CREATED,
        CLOUD_EVENT_DELETED,
    ):
        await receiver.receive(payload)

    publisher = AsyncEventPublisher(TOPIC_ENDPOINT, client=_AsyncFakePublisherClient())
    await publisher.publish([_downstream_event()])
    await publisher.close()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(name)s: %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())
