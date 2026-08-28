"""Publish custom downstream events to an Event Grid custom topic."""

from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    subject: str
    data: Any
    data_version: str = "1.0"

    def to_event_grid_event(self) -> EventGridEvent:
        if not self.subject.startswith("/"):
            raise ValueError(
                "Custom event subjects must start with '/' to support hierarchy filters"
            )
        return EventGridEvent(
            event_type=self.event_type,
            subject=self.subject,
            data=self.data,
            data_version=self.data_version,
        )


class EventPublisher:
    def __init__(self, publisher_client: Any) -> None:
        self._client = publisher_client

    def publish(self, events: list[CustomEvent]) -> bool:
        if not events:
            return True
        sdk_events = [event.to_event_grid_event() for event in events]
        try:
            self._client.send(sdk_events)
        except AzureError:
            logger.exception("Failed to publish %d downstream event(s)", len(events))
            return False
        logger.info("Published %d downstream event(s)", len(events))
        return True


class AsyncEventPublisher:
    def __init__(self, publisher_client: Any) -> None:
        self._client = publisher_client

    async def publish(self, events: list[CustomEvent]) -> bool:
        if not events:
            return True
        sdk_events = [event.to_event_grid_event() for event in events]
        try:
            await self._client.send(sdk_events)
        except AzureError:
            logger.exception("Failed to publish %d downstream event(s)", len(events))
            return False
        logger.info("Published %d downstream event(s)", len(events))
        return True
