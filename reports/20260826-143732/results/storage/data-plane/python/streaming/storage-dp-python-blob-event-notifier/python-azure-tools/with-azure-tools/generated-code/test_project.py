from __future__ import annotations

import asyncio
import unittest

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from blob_event_handler import parse_blob_subject
from event_publisher import CustomEvent, publish_custom_events
from event_receiver import deserialize_events, receive_events, receive_events_async
from main import cloud_events_payload, event_grid_payload


class RecordingPublisher:
    def __init__(self) -> None:
        self.events: list[EventGridEvent] = []

    def send(self, events: list[EventGridEvent]) -> None:
        self.events.extend(events)


class ProjectTests(unittest.TestCase):
    def test_deserializes_both_supported_schemas(self) -> None:
        native = deserialize_events(event_grid_payload())
        cloud = deserialize_events(cloud_events_payload())

        self.assertTrue(all(isinstance(event, EventGridEvent) for event in native))
        self.assertTrue(all(isinstance(event, CloudEvent) for event in cloud))

    def test_routes_sync_events(self) -> None:
        created: list[object] = []
        deleted: list[object] = []

        receive_events(event_grid_payload(), created.append, deleted.append)

        self.assertEqual(len(created), 1)
        self.assertEqual(len(deleted), 1)

    def test_routes_async_events(self) -> None:
        created: list[object] = []
        deleted: list[object] = []

        async def on_created(event: object) -> None:
            created.append(event)

        async def on_deleted(event: object) -> None:
            deleted.append(event)

        asyncio.run(
            receive_events_async(
                cloud_events_payload(), on_created, on_deleted
            )
        )

        self.assertEqual(len(created), 1)
        self.assertEqual(len(deleted), 1)

    def test_parses_encoded_blob_name(self) -> None:
        location = parse_blob_subject(
            "/blobServices/default/containers/docs/blobs/a%20folder/file.pdf"
        )

        self.assertEqual(location.container, "docs")
        self.assertEqual(location.name, "a folder/file.pdf")

    def test_publisher_applies_subject_hierarchy(self) -> None:
        publisher = RecordingPublisher()

        events = publish_custom_events(
            "https://example.eventgrid.azure.net/api/events",
            [CustomEvent("Contoso.DocumentProcessed", {"id": "1"})],
            "/documents/invoices/processed",
            publisher_client=publisher,
        )

        self.assertEqual(events[0].subject, "/documents/invoices/processed")
        self.assertEqual(publisher.events, events)


if __name__ == "__main__":
    unittest.main()
