from __future__ import annotations

import logging
from collections.abc import Callable, Iterable
from dataclasses import dataclass
from typing import Any

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
    ServiceResponseError,
)
from azure.eventgrid import EventGridEvent

from configuration import (
    open_async_event_grid_publisher,
    open_event_grid_publisher,
)

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    data: dict[str, Any]
    subject: str
    data_version: str = "1.0"

    def to_event_grid_event(self) -> EventGridEvent:
        if not self.subject.startswith("/"):
            raise ValueError("Custom event subject must be an absolute hierarchy")
        return EventGridEvent(
            subject=self.subject,
            event_type=self.event_type,
            data=self.data,
            data_version=self.data_version,
        )


def publish_events(
    endpoint: str,
    events: Iterable[CustomEvent],
    *,
    client_context: Callable[[str], Any] = open_event_grid_publisher,
) -> bool:
    sdk_events = [event.to_event_grid_event() for event in events]
    if not sdk_events:
        logger.warning("No custom events supplied for publishing")
        return True

    try:
        with client_context(endpoint) as client:
            client.send(sdk_events)
    except (
        ClientAuthenticationError,
        HttpResponseError,
        ServiceRequestError,
        ServiceResponseError,
    ) as error:
        logger.error("Event Grid publishing failed: %s", error)
        return False

    logger.info("Published %d custom event(s)", len(sdk_events))
    return True


async def publish_events_async(
    endpoint: str,
    events: Iterable[CustomEvent],
    *,
    client_context: Callable[[str], Any] = open_async_event_grid_publisher,
) -> bool:
    sdk_events = [event.to_event_grid_event() for event in events]
    if not sdk_events:
        logger.warning("No custom events supplied for publishing")
        return True

    try:
        async with client_context(endpoint) as client:
            await client.send(sdk_events)
    except (
        ClientAuthenticationError,
        HttpResponseError,
        ServiceRequestError,
        ServiceResponseError,
    ) as error:
        logger.error("Async Event Grid publishing failed: %s", error)
        return False

    logger.info("Published %d custom event(s) asynchronously", len(sdk_events))
    return True
