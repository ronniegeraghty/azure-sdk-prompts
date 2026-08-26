from __future__ import annotations

import json
import logging
from enum import Enum
from typing import Any

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from blob_event_handler import AsyncBlobEventHandler, BlobEvent, BlobEventHandler

logger = logging.getLogger(__name__)

BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"


class EventSchema(str, Enum):
    AUTO = "auto"
    EVENT_GRID = "event-grid"
    CLOUD_EVENTS = "cloud-events"


def deserialize_events(
    payload: str | bytes,
    schema: EventSchema = EventSchema.AUTO,
) -> list[BlobEvent]:
    envelope = json.loads(payload)
    raw_events = envelope if isinstance(envelope, list) else [envelope]
    if not all(isinstance(item, dict) for item in raw_events):
        raise ValueError("Event Grid payload must contain JSON event objects")

    events: list[BlobEvent] = []
    for raw_event in raw_events:
        selected_schema = schema
        if selected_schema is EventSchema.AUTO:
            selected_schema = (
                EventSchema.CLOUD_EVENTS
                if "specversion" in raw_event
                else EventSchema.EVENT_GRID
            )

        # The SDK helper performs schema validation, field mapping, and time conversion.
        event_json = json.dumps(raw_event)
        if selected_schema is EventSchema.CLOUD_EVENTS:
            events.append(CloudEvent.from_json(event_json))
        else:
            events.append(EventGridEvent.from_json(event_json))
    return events


def event_type(event: BlobEvent) -> str:
    if isinstance(event, CloudEvent):
        return event.type
    return event.event_type


class EventReceiver:
    def __init__(self, handler: BlobEventHandler) -> None:
        self._handler = handler

    def receive(
        self,
        payload: str | bytes,
        schema: EventSchema = EventSchema.AUTO,
    ) -> list[BlobEvent]:
        events = deserialize_events(payload, schema)
        for event in events:
            kind = event_type(event)
            if kind == BLOB_CREATED:
                self._handler.handle_created(event)
            elif kind == BLOB_DELETED:
                self._handler.handle_deleted(event)
            else:
                logger.warning("Ignoring unrecognized event type: %s", kind)
        return events


class AsyncEventReceiver:
    def __init__(self, handler: AsyncBlobEventHandler) -> None:
        self._handler = handler

    async def receive(
        self,
        payload: str | bytes,
        schema: EventSchema = EventSchema.AUTO,
    ) -> list[BlobEvent]:
        events = deserialize_events(payload, schema)
        for event in events:
            kind = event_type(event)
            if kind == BLOB_CREATED:
                await self._handler.handle_created(event)
            elif kind == BLOB_DELETED:
                await self._handler.handle_deleted(event)
            else:
                logger.warning("Ignoring unrecognized event type: %s", kind)
        return events
