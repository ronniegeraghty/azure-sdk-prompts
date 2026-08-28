"""Sync and async publishers for custom downstream Event Grid events."""

from __future__ import annotations

import logging
from collections.abc import Mapping, Sequence
from dataclasses import dataclass
from typing import Any

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent

from .config import (
    create_async_event_grid_publisher_client,
    create_event_grid_publisher_client,
)

LOGGER = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    data: Mapping[str, Any]
    subject: str
    data_version: str = "1.0"


def _normalize_subject(subject: str) -> str:
    normalized = "/" + subject.strip("/")
    if normalized == "/" or "//" in normalized:
        raise ValueError("Event subject must contain a valid hierarchy")
    return normalized


def _to_event_grid_events(events: Sequence[CustomEvent]) -> list[EventGridEvent]:
    if not events:
        raise ValueError("At least one custom event is required")
    return [
        EventGridEvent(
            subject=_normalize_subject(event.subject),
            event_type=event.event_type,
            data=dict(event.data),
            data_version=event.data_version,
        )
        for event in events
    ]


def publish_events(
    endpoint: str,
    events: Sequence[CustomEvent],
    *,
    client: Any | None = None,
) -> bool:
    publisher = client or create_event_grid_publisher_client(endpoint)
    owns_client = client is None
    try:
        publisher.send(_to_event_grid_events(events))
        LOGGER.info("Published %d downstream event(s)", len(events))
        return True
    except AzureError:
        LOGGER.exception("Event Grid publishing failed")
        return False
    finally:
        if owns_client:
            publisher.close()


async def publish_events_async(
    endpoint: str,
    events: Sequence[CustomEvent],
    *,
    client: Any | None = None,
) -> bool:
    publisher = client or create_async_event_grid_publisher_client(endpoint)
    owns_client = client is None
    try:
        await publisher.send(_to_event_grid_events(events))
        LOGGER.info("Published %d downstream event(s)", len(events))
        return True
    except AzureError:
        LOGGER.exception("Event Grid publishing failed")
        return False
    finally:
        if owns_client:
            await publisher.close()
