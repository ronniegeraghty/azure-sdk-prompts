from __future__ import annotations

import json
import logging
from collections.abc import Awaitable, Callable, Mapping
from typing import Any, TypeAlias

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent, SystemEventNames

BlobEvent: TypeAlias = EventGridEvent | CloudEvent[Any]
SyncCreatedHandler: TypeAlias = Callable[[BlobEvent], None]
SyncDeletedHandler: TypeAlias = Callable[[BlobEvent], None]
AsyncCreatedHandler: TypeAlias = Callable[[BlobEvent], Awaitable[None]]
AsyncDeletedHandler: TypeAlias = Callable[[BlobEvent], Awaitable[None]]

logger = logging.getLogger(__name__)


def deserialize_events(payload: str | bytes) -> list[BlobEvent]:
    try:
        envelope = json.loads(payload)
    except (json.JSONDecodeError, UnicodeDecodeError) as exc:
        raise ValueError("Event Grid payload is not valid JSON") from exc

    raw_events = envelope if isinstance(envelope, list) else [envelope]
    if not raw_events:
        return []

    events: list[BlobEvent] = []
    for raw_event in raw_events:
        if not isinstance(raw_event, Mapping):
            raise ValueError("Each Event Grid payload item must be a JSON object")

        serialized_event = json.dumps(raw_event)
        if "specversion" in raw_event:
            events.append(CloudEvent.from_json(serialized_event))
        elif "eventType" in raw_event:
            events.append(EventGridEvent.from_json(serialized_event))
        else:
            raise ValueError(
                "Event does not match CloudEvents 1.0 or Event Grid schema"
            )
    return events


def receive_events(
    payload: str | bytes,
    on_blob_created: SyncCreatedHandler,
    on_blob_deleted: SyncDeletedHandler,
) -> list[BlobEvent]:
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == SystemEventNames.StorageBlobCreatedEventName:
            on_blob_created(event)
        elif event_type == SystemEventNames.StorageBlobDeletedEventName:
            on_blob_deleted(event)
        else:
            logger.warning(
                "Ignoring unrecognized Event Grid event type %s (id=%s)",
                event_type,
                event.id,
            )
    return events


async def receive_events_async(
    payload: str | bytes,
    on_blob_created: AsyncCreatedHandler,
    on_blob_deleted: AsyncDeletedHandler,
) -> list[BlobEvent]:
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == SystemEventNames.StorageBlobCreatedEventName:
            await on_blob_created(event)
        elif event_type == SystemEventNames.StorageBlobDeletedEventName:
            await on_blob_deleted(event)
        else:
            logger.warning(
                "Ignoring unrecognized Event Grid event type %s (id=%s)",
                event_type,
                event.id,
            )
    return events


def _event_type(event: BlobEvent) -> str:
    if isinstance(event, EventGridEvent):
        return event.event_type
    return event.type
