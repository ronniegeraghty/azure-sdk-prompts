"""Event Grid webhook payload deserialization and routing."""

from __future__ import annotations

import logging
from collections.abc import Mapping, Sequence
from typing import Any, Literal, TypeAlias

from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

from blob_event_handler import (
    handle_blob_created,
    handle_blob_created_async,
    handle_blob_deleted,
    handle_blob_deleted_async,
)

LOGGER = logging.getLogger(__name__)

BLOB_CREATED = "Microsoft.Storage.BlobCreated"
BLOB_DELETED = "Microsoft.Storage.BlobDeleted"
Schema = Literal["auto", "eventgrid", "cloudevent"]
JsonDocument: TypeAlias = str | bytes | Mapping[str, Any]
JsonPayload: TypeAlias = JsonDocument | Sequence[JsonDocument]


def _documents(payload: JsonPayload) -> list[JsonDocument]:
    if isinstance(payload, (str, bytes, Mapping)):
        return [payload]
    return list(payload)


def _detect_schema(document: JsonDocument) -> Literal["eventgrid", "cloudevent"]:
    if not isinstance(document, Mapping):
        raise ValueError("schema='auto' requires a decoded JSON object")
    if document.get("specversion") == "1.0":
        return "cloudevent"
    if "eventType" in document:
        return "eventgrid"
    raise ValueError("Payload is neither Event Grid schema nor CloudEvents 1.0")


def deserialize_events(payload: JsonPayload, schema: Schema = "auto") -> list[Any]:
    events: list[Any] = []
    for document in _documents(payload):
        if schema == "auto" and not isinstance(document, Mapping):
            try:
                events.append(CloudEvent.from_json(document))
            except ValueError:
                events.append(EventGridEvent.from_json(document))
            continue

        selected_schema = _detect_schema(document) if schema == "auto" else schema
        model = CloudEvent if selected_schema == "cloudevent" else EventGridEvent
        if isinstance(document, Mapping):
            events.append(model.from_dict(dict(document)))
        else:
            events.append(model.from_json(document))
    return events


def _event_type(event: Any) -> str:
    return event.type if isinstance(event, CloudEvent) else event.event_type


def receive_events(
    payload: JsonPayload,
    blob_service: Any,
    schema: Schema = "auto",
) -> list[Any]:
    events = deserialize_events(payload, schema)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            handle_blob_created(event, blob_service)
        elif event_type == BLOB_DELETED:
            handle_blob_deleted(event)
        else:
            LOGGER.warning("Ignoring unrecognized Event Grid event type: %s", event_type)
    return events


async def receive_events_async(
    payload: JsonPayload,
    blob_service: Any,
    schema: Schema = "auto",
) -> list[Any]:
    events = deserialize_events(payload, schema)
    for event in events:
        event_type = _event_type(event)
        if event_type == BLOB_CREATED:
            await handle_blob_created_async(event, blob_service)
        elif event_type == BLOB_DELETED:
            await handle_blob_deleted_async(event)
        else:
            LOGGER.warning("Ignoring unrecognized Event Grid event type: %s", event_type)
    return events
