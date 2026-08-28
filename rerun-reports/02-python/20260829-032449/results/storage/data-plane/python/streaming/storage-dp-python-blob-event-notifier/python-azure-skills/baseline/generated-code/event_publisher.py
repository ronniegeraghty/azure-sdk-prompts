"""Publish custom downstream notifications to an Event Grid topic."""

from __future__ import annotations

import logging
from collections.abc import Iterable
from dataclasses import dataclass
from datetime import datetime, timezone
from typing import Any
from uuid import uuid4

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent, EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

LOGGER = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    data: Any
    subject: str
    data_version: str = "1.0"

    def to_event_grid_event(self) -> EventGridEvent:
        return EventGridEvent(
            id=str(uuid4()),
            event_time=datetime.now(timezone.utc),
            subject=self.subject,
            event_type=self.event_type,
            data=self.data,
            data_version=self.data_version,
        )


def publish_custom_events(
    topic_endpoint: str,
    custom_events: Iterable[CustomEvent],
    *,
    client: Any | None = None,
) -> bool:
    owned_credential = None
    owned_client = None
    if client is None:
        owned_credential = DefaultAzureCredential()
        owned_client = EventGridPublisherClient(topic_endpoint, owned_credential)
        client = owned_client

    events = [event.to_event_grid_event() for event in custom_events]
    try:
        client.send(events)
        return True
    except AzureError:
        LOGGER.exception("Failed to publish %d custom Event Grid event(s)", len(events))
        return False
    finally:
        if owned_client is not None:
            owned_client.close()
        if owned_credential is not None:
            owned_credential.close()


async def publish_custom_events_async(
    topic_endpoint: str,
    custom_events: Iterable[CustomEvent],
    *,
    client: Any | None = None,
) -> bool:
    owned_credential = None
    owned_client = None
    if client is None:
        owned_credential = AsyncDefaultAzureCredential()
        owned_client = AsyncEventGridPublisherClient(topic_endpoint, owned_credential)
        client = owned_client

    events = [event.to_event_grid_event() for event in custom_events]
    try:
        await client.send(events)
        return True
    except AzureError:
        LOGGER.exception("Failed to publish %d custom Event Grid event(s)", len(events))
        return False
    finally:
        if owned_client is not None:
            await owned_client.close()
        if owned_credential is not None:
            await owned_credential.close()
