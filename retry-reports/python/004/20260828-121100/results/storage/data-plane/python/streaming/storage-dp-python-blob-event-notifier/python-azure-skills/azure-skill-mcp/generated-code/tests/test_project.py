from __future__ import annotations

import asyncio
import unittest
from types import SimpleNamespace
from unittest.mock import AsyncMock, Mock

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from blob_event_notifier.blob_handler import BlobEventHandler, parse_blob_subject
from blob_event_notifier.publisher import AsyncEventPublisher, CustomEvent, EventPublisher
from blob_event_notifier.receiver import AsyncEventReceiver, EventReceiver, deserialize_event
from main import (
    CLOUD_EVENT_CREATED,
    CLOUD_EVENT_DELETED,
    EVENT_GRID_CREATED,
    EVENT_GRID_DELETED,
)


class ReceiverTests(unittest.TestCase):
    def test_deserializes_both_supported_schemas(self) -> None:
        self.assertIsInstance(deserialize_event(EVENT_GRID_CREATED), EventGridEvent)
        self.assertIsInstance(deserialize_event(CLOUD_EVENT_CREATED), CloudEvent)

    def test_routes_sync_events(self) -> None:
        handler = Mock()
        receiver = EventReceiver(handler)

        receiver.receive(EVENT_GRID_CREATED)
        receiver.receive(EVENT_GRID_DELETED)

        handler.handle_blob_created.assert_called_once()
        handler.handle_blob_deleted.assert_called_once()

    def test_routes_async_events(self) -> None:
        handler = SimpleNamespace(
            handle_blob_created=AsyncMock(),
            handle_blob_deleted=AsyncMock(),
        )

        async def run() -> None:
            receiver = AsyncEventReceiver(handler)
            await receiver.receive(CLOUD_EVENT_CREATED)
            await receiver.receive(CLOUD_EVENT_DELETED)

        asyncio.run(run())
        handler.handle_blob_created.assert_awaited_once()
        handler.handle_blob_deleted.assert_awaited_once()


class BlobHandlerTests(unittest.TestCase):
    def test_parses_encoded_blob_name_and_preserves_hierarchy(self) -> None:
        location = parse_blob_subject(
            "/blobServices/default/containers/documents/blobs/folder/invoice%202026.pdf"
        )
        self.assertEqual(location.container, "documents")
        self.assertEqual(location.name, "folder/invoice 2026.pdf")

    def test_handles_disappearing_blob(self) -> None:
        from azure.core.exceptions import ResourceNotFoundError

        blob_client = Mock()
        blob_client.download_blob.side_effect = ResourceNotFoundError("gone")
        service = Mock()
        service.get_blob_client.return_value = blob_client

        BlobEventHandler(service).handle_blob_created(deserialize_event(EVENT_GRID_CREATED))


class PublisherTests(unittest.TestCase):
    def test_publishes_subject_hierarchy(self) -> None:
        client = Mock()
        publisher = EventPublisher("https://example.invalid", client=client)
        publisher.publish(
            [
                CustomEvent(
                    event_type="Document.Processed",
                    subject="documents/invoices/processed",
                    data={"id": "42"},
                )
            ]
        )

        sent_event = client.send.call_args.args[0][0]
        self.assertEqual(sent_event.subject, "/documents/invoices/processed")

    def test_async_publisher(self) -> None:
        client = SimpleNamespace(send=AsyncMock(), close=AsyncMock())

        async def run() -> None:
            publisher = AsyncEventPublisher("https://example.invalid", client=client)
            await publisher.publish(
                [
                    CustomEvent(
                        event_type="Document.Processed",
                        subject="/documents/invoices/processed",
                        data={"id": "42"},
                    )
                ]
            )

        asyncio.run(run())
        client.send.assert_awaited_once()


if __name__ == "__main__":
    unittest.main()
