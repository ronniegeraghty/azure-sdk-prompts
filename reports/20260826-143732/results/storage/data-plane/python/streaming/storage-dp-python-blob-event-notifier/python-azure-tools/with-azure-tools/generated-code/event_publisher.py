from __future__ import annotations

import logging
from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from typing import Any

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent

from config import async_event_grid_client, sync_event_grid_client

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    data: Mapping[str, Any]
    subject: str | None = None
    data_version: str = "1.0"


class EventPublishingError(RuntimeError):
    pass


def publish_custom_events(
    endpoint: str,
    events: Iterable[CustomEvent],
    subject: str,
    *,
    publisher_client: Any | None = None,
) -> list[EventGridEvent]:
    event_batch = _build_event_batch(events, subject)
    if not event_batch:
        return []

    try:
        if publisher_client is not None:
            publisher_client.send(event_batch)
        else:
            with sync_event_grid_client(endpoint) as client:
                client.send(event_batch)
    except AzureError as exc:
        logger.error(
            "Failed to publish %d event(s) to Event Grid: %s",
            len(event_batch),
            exc,
        )
        raise EventPublishingError("Event Grid publishing failed") from exc

    logger.info("Published %d downstream event(s)", len(event_batch))
    return event_batch


async def publish_custom_events_async(
    endpoint: str,
    events: Iterable[CustomEvent],
    subject: str,
    *,
    publisher_client: Any | None = None,
) -> list[EventGridEvent]:
    event_batch = _build_event_batch(events, subject)
    if not event_batch:
        return []

    try:
        if publisher_client is not None:
            await publisher_client.send(event_batch)
        else:
            async with async_event_grid_client(endpoint) as client:
                await client.send(event_batch)
    except AzureError as exc:
        logger.error(
            "Failed to publish %d event(s) to Event Grid: %s",
            len(event_batch),
            exc,
        )
        raise EventPublishingError("Event Grid publishing failed") from exc

    logger.info("Published %d downstream event(s)", len(event_batch))
    return event_batch


def _build_event_batch(
    events: Iterable[CustomEvent], default_subject: str
) -> list[EventGridEvent]:
    _validate_subject(default_subject)
    batch = []
    for event in events:
        event_subject = event.subject or default_subject
        _validate_subject(event_subject)
        batch.append(
            EventGridEvent(
                subject=event_subject,
                event_type=event.event_type,
                data=dict(event.data),
                data_version=event.data_version,
            )
        )
    return batch


def _validate_subject(subject: str) -> None:
    if not subject.startswith("/") or subject.endswith("/"):
        raise ValueError(
            "Event subject must be an absolute hierarchy without a trailing slash"
        )
