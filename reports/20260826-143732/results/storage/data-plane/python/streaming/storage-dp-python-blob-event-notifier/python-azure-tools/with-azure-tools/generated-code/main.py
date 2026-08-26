from __future__ import annotations

import asyncio
import json
import logging
from functools import partial
from types import SimpleNamespace
from typing import Any

from blob_event_handler import (
    handle_blob_created,
    handle_blob_created_async,
    handle_blob_deleted,
    handle_blob_deleted_async,
)
from event_publisher import (
    CustomEvent,
    publish_custom_events,
    publish_custom_events_async,
)
from event_receiver import receive_events, receive_events_async

DEMO_ENDPOINT = "https://example-topic.eastus-1.eventgrid.azure.net/api/events"
CREATED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/invoices/invoice-1001.pdf"
)
DELETED_SUBJECT = (
    "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf"
)


class DemoDownloader:
    def chunks(self) -> list[bytes]:
        return [b"demo invoice content"]


class AsyncDemoDownloader:
    async def chunks(self) -> Any:
        yield b"demo invoice content"


class DemoBlobClient:
    def get_blob_properties(self) -> SimpleNamespace:
        return SimpleNamespace(
            content_settings=SimpleNamespace(content_type="application/pdf"),
            blob_tier="Hot",
        )

    def download_blob(self) -> DemoDownloader:
        return DemoDownloader()


class AsyncDemoBlobClient:
    async def get_blob_properties(self) -> SimpleNamespace:
        return SimpleNamespace(
            content_settings=SimpleNamespace(content_type="application/pdf"),
            blob_tier="Hot",
        )

    async def download_blob(self) -> AsyncDemoDownloader:
        return AsyncDemoDownloader()


class DemoBlobServiceClient:
    def get_blob_client(self, *, container: str, blob: str) -> DemoBlobClient:
        logging.getLogger(__name__).debug(
            "Demo blob lookup: container=%s blob=%s", container, blob
        )
        return DemoBlobClient()


class AsyncDemoBlobServiceClient:
    def get_blob_client(
        self, *, container: str, blob: str
    ) -> AsyncDemoBlobClient:
        logging.getLogger(__name__).debug(
            "Async demo blob lookup: container=%s blob=%s", container, blob
        )
        return AsyncDemoBlobClient()


class DemoPublisherClient:
    def send(self, events: list[Any]) -> None:
        logging.getLogger(__name__).info(
            "Demo publisher accepted %d event(s)", len(events)
        )


class AsyncDemoPublisherClient:
    async def send(self, events: list[Any]) -> None:
        logging.getLogger(__name__).info(
            "Async demo publisher accepted %d event(s)", len(events)
        )


def event_grid_payload() -> str:
    return json.dumps(
        [
            {
                "id": "4fcbfb95-35a7-4c72-bd3d-bf6fd18e6e1a",
                "topic": (
                    "/subscriptions/00000000-0000-0000-0000-000000000000/"
                    "resourceGroups/demo/providers/Microsoft.Storage/"
                    "storageAccounts/demostorage"
                ),
                "subject": CREATED_SUBJECT,
                "eventType": "Microsoft.Storage.BlobCreated",
                "eventTime": "2026-08-26T08:00:00Z",
                "data": {
                    "api": "PutBlob",
                    "contentType": "application/pdf",
                    "contentLength": 20,
                    "url": (
                        "https://demostorage.blob.core.windows.net/"
                        "documents/invoices/invoice-1001.pdf"
                    ),
                    "sequencer": "0000000000000000000000000000001",
                },
                "dataVersion": "3",
                "metadataVersion": "1",
            },
            {
                "id": "c16e389b-e75b-469f-bd3d-d9babe441a97",
                "topic": (
                    "/subscriptions/00000000-0000-0000-0000-000000000000/"
                    "resourceGroups/demo/providers/Microsoft.Storage/"
                    "storageAccounts/demostorage"
                ),
                "subject": DELETED_SUBJECT,
                "eventType": "Microsoft.Storage.BlobDeleted",
                "eventTime": "2026-08-26T08:01:00Z",
                "data": {
                    "api": "DeleteBlob",
                    "url": (
                        "https://demostorage.blob.core.windows.net/"
                        "documents/archive/old-invoice.pdf"
                    ),
                    "sequencer": "0000000000000000000000000000002",
                },
                "dataVersion": "3",
                "metadataVersion": "1",
            },
        ]
    )


def cloud_events_payload() -> str:
    return json.dumps(
        [
            {
                "specversion": "1.0",
                "id": "ec724c7d-4dc8-4c14-8f3d-983d430ffa0c",
                "source": (
                    "/subscriptions/00000000-0000-0000-0000-000000000000/"
                    "resourceGroups/demo/providers/Microsoft.Storage/"
                    "storageAccounts/demostorage"
                ),
                "type": "Microsoft.Storage.BlobCreated",
                "subject": CREATED_SUBJECT,
                "time": "2026-08-26T08:02:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "PutBlob",
                    "contentType": "application/pdf",
                    "contentLength": 20,
                    "url": (
                        "https://demostorage.blob.core.windows.net/"
                        "documents/invoices/invoice-1001.pdf"
                    ),
                    "sequencer": "0000000000000000000000000000003",
                },
            },
            {
                "specversion": "1.0",
                "id": "f47fef06-bf57-4381-b17f-c63fd7c2b25f",
                "source": (
                    "/subscriptions/00000000-0000-0000-0000-000000000000/"
                    "resourceGroups/demo/providers/Microsoft.Storage/"
                    "storageAccounts/demostorage"
                ),
                "type": "Microsoft.Storage.BlobDeleted",
                "subject": DELETED_SUBJECT,
                "time": "2026-08-26T08:03:00Z",
                "datacontenttype": "application/json",
                "data": {
                    "api": "DeleteBlob",
                    "url": (
                        "https://demostorage.blob.core.windows.net/"
                        "documents/archive/old-invoice.pdf"
                    ),
                    "sequencer": "0000000000000000000000000000004",
                },
            },
        ]
    )


def run_sync_demo() -> None:
    blob_client = DemoBlobServiceClient()
    created_handler = partial(
        handle_blob_created, blob_service_client=blob_client
    )
    receive_events(event_grid_payload(), created_handler, handle_blob_deleted)
    receive_events(cloud_events_payload(), created_handler, handle_blob_deleted)
    publish_custom_events(
        DEMO_ENDPOINT,
        [
            CustomEvent(
                event_type="Contoso.Documents.DocumentProcessed",
                data={"documentId": "invoice-1001", "status": "processed"},
            )
        ],
        "/documents/invoices/processed",
        publisher_client=DemoPublisherClient(),
    )


async def run_async_demo() -> None:
    blob_client = AsyncDemoBlobServiceClient()
    created_handler = partial(
        handle_blob_created_async, blob_service_client=blob_client
    )
    await receive_events_async(
        event_grid_payload(), created_handler, handle_blob_deleted_async
    )
    await receive_events_async(
        cloud_events_payload(), created_handler, handle_blob_deleted_async
    )
    await publish_custom_events_async(
        DEMO_ENDPOINT,
        [
            CustomEvent(
                event_type="Contoso.Documents.DocumentProcessed",
                data={"documentId": "invoice-1001", "status": "processed"},
            )
        ],
        "/documents/invoices/processed",
        publisher_client=AsyncDemoPublisherClient(),
    )


async def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    logging.info("Running synchronous demo")
    run_sync_demo()
    logging.info("Running asynchronous demo")
    await run_async_demo()


if __name__ == "__main__":
    asyncio.run(main())
