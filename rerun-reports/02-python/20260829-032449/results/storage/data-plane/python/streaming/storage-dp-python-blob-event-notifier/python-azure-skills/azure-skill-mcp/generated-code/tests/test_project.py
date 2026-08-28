from __future__ import annotations

import unittest

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from blob_event_notifier.blob_handler import parse_blob_subject
from blob_event_notifier.event_publisher import CustomEvent, publish_events
from blob_event_notifier.event_receiver import deserialize_events
from main import MockPublisherClient, sample_payloads


class ProjectTests(unittest.TestCase):
    def test_deserializes_both_supported_schemas(self) -> None:
        native_payload, cloud_payload = sample_payloads()

        native = deserialize_events(native_payload)
        cloud = deserialize_events(cloud_payload)

        self.assertIsInstance(native[0], EventGridEvent)
        self.assertIsInstance(cloud[0], CloudEvent)

    def test_parses_nested_and_encoded_blob_name(self) -> None:
        location = parse_blob_subject(
            "/blobServices/default/containers/docs/blobs/2026/invoice%2001.pdf"
        )

        self.assertEqual("docs", location.container)
        self.assertEqual("2026/invoice 01.pdf", location.name)

    def test_publisher_normalizes_subject(self) -> None:
        client = MockPublisherClient()

        published = publish_events(
            "https://example.invalid",
            [CustomEvent("Contoso.Test", {"ok": True}, "documents/processed")],
            client=client,
        )

        self.assertTrue(published)


if __name__ == "__main__":
    unittest.main()
