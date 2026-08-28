"""Event Grid webhook payload deserialization and routing."""

from __future__ import annotations

import json
import logging
from collections.abc import Callable
from typing import Any, TypeAlias

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from .blob_handler import (
    handle_blob_created,
    handle_blob_created_async,
    handle_blob_deleted,
    handle_blob_deleted_async,
)

LOGGER = logging.getLogger(__name__)
BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"
GridEvent: TypeAlias = EventGridEvent | CloudEvent


def deserialize_events(payload: str | bytes) -> list[GridEvent]:
    try:
        envelope = json.loads(payload)
    except (TypeError, json.JSONDecodeError) as exc:
        raise ValueError("Webhook payload is not valid JSON") from exc

    raw_events = envelope if isinstance(envelope, list) else [envelope]
    events: list[GridEvent] = []
    for raw_event in raw_events:
        if not isinstance(raw_event, dict):
            raise ValueError("Each webhook event must be a JSON object")

        # SDK helpers own the schema-to-model mapping; this code only identifies
        # which Event Grid-supported envelope was delivered.
        if "specversion" in raw_event:
            events.append(CloudEvent.from_dict(raw_event))
        else:
            events.append(EventGridEvent.from_dict(raw_event))
    return events


def _event_type(event: GridEvent) -> str:
    return event.type if isinstance(event, CloudEvent) else event.event_type


def receive_events(
    payload: str | bytes,
    blob_service_client: Any,
    *,
    on_created: Callable[[GridEvent, Any], None] = handle_blob_created,
    on_deleted: Callable[[GridEvent], None] = handle_blob_deleted,
) -> list[GridEvent]:
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            on_created(event, blob_service_client)
        elif event_type == BLOB_DELETED:
            on_deleted(event)
        else:
            LOGGER.warning("Ignoring unsupported event type: %s", event_type)
    return events


async def receive_events_async(
    payload: str | bytes,
    blob_service_client: Any,
) -> list[GridEvent]:
    events = deserialize_events(payload)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            await handle_blob_created_async(event, blob_service_client)
        elif event_type == BLOB_DELETED:
            await handle_blob_deleted_async(event)
        else:
            LOGGER.warning("Ignoring unsupported event type: %s", event_type)
    return events
