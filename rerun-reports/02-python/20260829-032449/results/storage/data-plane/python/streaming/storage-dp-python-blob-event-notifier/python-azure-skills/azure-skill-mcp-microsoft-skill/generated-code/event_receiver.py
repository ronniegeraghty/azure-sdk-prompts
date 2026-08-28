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


def deserialize_events(payload: str | bytes) -> list[ReceivedEvent]:
    """Deserialize an Event Grid webhook batch using Azure SDK model helpers."""
    envelope: Any = json.loads(payload)
    raw_events = envelope if isinstance(envelope, list) else [envelope]
    if not all(isinstance(item, dict) for item in raw_events):
        raise ValueError("Event Grid payload must contain JSON event objects")

    events: list[ReceivedEvent] = []
    for raw_event in raw_events:
        serialized_event = json.dumps(raw_event)
        if "specversion" in raw_event:
            events.append(CloudEvent.from_json(serialized_event))
        elif "eventType" in raw_event:
            events.append(EventGridEvent.from_json(serialized_event))
        else:
            raise ValueError("Event does not match Event Grid or CloudEvents 1.0 schema")
    return events


def _event_type(event: ReceivedEvent) -> str:
    return event.type if isinstance(event, CloudEvent) else event.event_type


def receive_events(
    payload: str | bytes,
    on_blob_created: SyncHandler,
    on_blob_deleted: SyncHandler,
) -> list[ReceivedEvent]:
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            on_blob_created(event)
        elif event_type == BLOB_DELETED:
            on_blob_deleted(event)
        else:
            logger.warning("Ignoring unrecognized event type %s", event_type)
    return events


async def receive_events_async(
    payload: str | bytes,
    on_blob_created: AsyncHandler,
    on_blob_deleted: AsyncHandler,
) -> list[ReceivedEvent]:
    if not inspect.iscoroutinefunction(on_blob_created):
        raise TypeError("on_blob_created must be an async callable")
    if not inspect.iscoroutinefunction(on_blob_deleted):
        raise TypeError("on_blob_deleted must be an async callable")

    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            await on_blob_created(event)
        elif event_type == BLOB_DELETED:
            await on_blob_deleted(event)
        else:
            logger.warning("Ignoring unrecognized event type %s", event_type)
    return events
