"""Event Grid native-schema and CloudEvents 1.0 receivers."""

from __future__ import annotations

import logging
from typing import Any

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

LOGGER = logging.getLogger(__name__)

BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"


def deserialize_event(payload: str | bytes) -> EventGridEvent | CloudEvent[Any]:
    """Deserialize one structured event with Azure SDK-provided helpers."""
    try:
        return CloudEvent.from_json(payload)
    except (KeyError, TypeError, ValueError):
        event = EventGridEvent.from_json(payload)
        if not event.event_type:
            raise ValueError(
                "Payload is neither a CloudEvents 1.0 event nor an Event Grid event"
            )
        return event


def _event_type(event: EventGridEvent | CloudEvent[Any]) -> str:
    if isinstance(event, CloudEvent):
        return event.type
    return event.event_type


class EventReceiver:
    def __init__(self, handler: Any) -> None:
        self._handler = handler

    def receive(self, payload: str | bytes) -> EventGridEvent | CloudEvent[Any]:
        event = deserialize_event(payload)
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            self._handler.handle_blob_created(event)
        elif event_type == BLOB_DELETED:
            self._handler.handle_blob_deleted(event)
        else:
            LOGGER.warning("Ignoring unsupported Event Grid event type: %s", event_type)
        return event


class AsyncEventReceiver:
    def __init__(self, handler: Any) -> None:
        self._handler = handler

    async def receive(self, payload: str | bytes) -> EventGridEvent | CloudEvent[Any]:
        event = deserialize_event(payload)
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            await self._handler.handle_blob_created(event)
        elif event_type == BLOB_DELETED:
            await self._handler.handle_blob_deleted(event)
        else:
            LOGGER.warning("Ignoring unsupported Event Grid event type: %s", event_type)
        return event
