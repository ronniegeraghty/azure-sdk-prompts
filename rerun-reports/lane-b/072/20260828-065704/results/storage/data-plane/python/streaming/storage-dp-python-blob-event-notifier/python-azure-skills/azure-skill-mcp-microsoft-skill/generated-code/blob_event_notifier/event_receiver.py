from __future__ import annotations

import json
import logging
from collections.abc import Mapping, Sequence
from typing import Any, Protocol

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from .blob_event_handler import StructuredEvent

BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"

JsonPayload = str | bytes | bytearray | Mapping[str, Any] | Sequence[Mapping[str, Any]]


class EventHandler(Protocol):
    def handle_created(self, event: StructuredEvent) -> None: ...

    def handle_deleted(self, event: StructuredEvent) -> None: ...


class AsyncEventHandler(Protocol):
    async def handle_created(self, event: StructuredEvent) -> None: ...

    async def handle_deleted(self, event: StructuredEvent) -> None: ...


def deserialize_events(payload: JsonPayload) -> list[StructuredEvent]:
    # JSON decoding only unwraps the webhook batch; SDK helpers construct and validate each event.
    decoded: Any = json.loads(payload) if isinstance(payload, (str, bytes, bytearray)) else payload
    items = [decoded] if isinstance(decoded, Mapping) else list(decoded)

    events: list[StructuredEvent] = []
    for item in items:
        if not isinstance(item, Mapping):
            raise ValueError("Each webhook batch item must be a JSON object")
        if item.get("specversion") == "1.0":
            events.append(CloudEvent.from_dict(dict(item)))
        elif "eventType" in item:
            events.append(EventGridEvent.from_dict(dict(item)))
        else:
            raise ValueError("Payload item is neither CloudEvents 1.0 nor Event Grid schema")
    return events


def _event_type(event: StructuredEvent) -> str:
    return event.type if isinstance(event, CloudEvent) else event.event_type


def receive_events(
    payload: JsonPayload,
    handler: EventHandler,
    logger: logging.Logger | None = None,
) -> list[StructuredEvent]:
    log = logger or logging.getLogger(__name__)
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            handler.handle_created(event)
        elif event_type == BLOB_DELETED:
            handler.handle_deleted(event)
        else:
            log.warning("Ignoring unrecognized Event Grid event type: %s", event_type)
    return events


async def receive_events_async(
    payload: JsonPayload,
    handler: AsyncEventHandler,
    logger: logging.Logger | None = None,
) -> list[StructuredEvent]:
    log = logger or logging.getLogger(__name__)
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            await handler.handle_created(event)
        elif event_type == BLOB_DELETED:
            await handler.handle_deleted(event)
        else:
            log.warning("Ignoring unrecognized Event Grid event type: %s", event_type)
    return events
