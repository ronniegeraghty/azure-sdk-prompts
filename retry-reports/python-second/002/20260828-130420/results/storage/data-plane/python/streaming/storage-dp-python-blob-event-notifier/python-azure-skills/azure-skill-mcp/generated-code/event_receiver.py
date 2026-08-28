"""Deserialize and route Event Grid webhook payloads."""

from __future__ import annotations

import inspect
import json
import logging
from collections.abc import Awaitable, Callable
from typing import Any, TypeAlias

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

logger = logging.getLogger(__name__)

BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"

ReceivedEvent: TypeAlias = EventGridEvent | CloudEvent
SyncHandler: TypeAlias = Callable[[ReceivedEvent], None]
AsyncHandler: TypeAlias = Callable[[ReceivedEvent], Awaitable[None]]


def _deserialize_one(raw_event: dict[str, Any]) -> ReceivedEvent:
    """Delegate field conversion and validation to Azure SDK model helpers."""
    if raw_event.get("specversion") == "1.0":
        return CloudEvent.from_dict(raw_event)
    return EventGridEvent.from_dict(raw_event)


def deserialize_events(payload: str | bytes) -> list[ReceivedEvent]:
    """Deserialize either a single event or an Event Grid batch."""
    decoded = json.loads(payload)
    envelopes = decoded if isinstance(decoded, list) else [decoded]
    if not all(isinstance(item, dict) for item in envelopes):
        raise ValueError("Event payload must contain a JSON object or an array of objects")
    return [_deserialize_one(item) for item in envelopes]


def event_type(event: ReceivedEvent) -> str:
    return event.type if isinstance(event, CloudEvent) else event.event_type


class EventReceiver:
    def __init__(self, on_created: SyncHandler, on_deleted: SyncHandler) -> None:
        self._on_created = on_created
        self._on_deleted = on_deleted

    def receive(self, payload: str | bytes) -> list[ReceivedEvent]:
        events = deserialize_events(payload)
        for event in events:
            kind = event_type(event)
            if kind == BLOB_CREATED:
                self._on_created(event)
            elif kind == BLOB_DELETED:
                self._on_deleted(event)
            else:
                logger.warning("Ignoring unrecognized Event Grid event type %s", kind)
        return events


class AsyncEventReceiver:
    def __init__(self, on_created: AsyncHandler, on_deleted: AsyncHandler) -> None:
        self._on_created = on_created
        self._on_deleted = on_deleted

    async def receive(self, payload: str | bytes) -> list[ReceivedEvent]:
        events = deserialize_events(payload)
        for event in events:
            kind = event_type(event)
            if kind == BLOB_CREATED:
                result = self._on_created(event)
            elif kind == BLOB_DELETED:
                result = self._on_deleted(event)
            else:
                logger.warning("Ignoring unrecognized Event Grid event type %s", kind)
                continue

            if not inspect.isawaitable(result):
                raise TypeError("Async event handlers must return an awaitable")
            await result
        return events
