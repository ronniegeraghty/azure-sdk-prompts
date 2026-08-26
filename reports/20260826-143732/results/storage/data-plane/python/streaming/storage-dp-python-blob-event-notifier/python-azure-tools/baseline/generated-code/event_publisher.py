from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any, Iterable

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent, EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    subject: str
    data: Any
    data_version: str = "1.0"

    def to_event_grid_event(self) -> EventGridEvent:
        if not self.subject.startswith("/"):
            raise ValueError("Custom event subjects must start with '/'")
        return EventGridEvent(
            subject=self.subject,
            event_type=self.event_type,
            data=self.data,
            data_version=self.data_version,
        )


class EventPublisher:
    def __init__(
        self,
        endpoint: str,
        credential: Any | None = None,
        *,
        client: EventGridPublisherClient | Any | None = None,
    ) -> None:
        if client is None and credential is None:
            raise ValueError("A credential is required when no publisher client is supplied")
        self._client = client or EventGridPublisherClient(endpoint, credential)

    def publish(self, events: Iterable[CustomEvent]) -> bool:
        sdk_events = [event.to_event_grid_event() for event in events]
        if not sdk_events:
            return True
        try:
            self._client.send(sdk_events)
        except AzureError:
            logger.exception("Failed to publish %d Event Grid event(s)", len(sdk_events))
            return False
        logger.info("Published %d Event Grid event(s)", len(sdk_events))
        return True


class AsyncEventPublisher:
    def __init__(
        self,
        endpoint: str,
        credential: Any | None = None,
        *,
        client: AsyncEventGridPublisherClient | Any | None = None,
    ) -> None:
        if client is None and credential is None:
            raise ValueError("A credential is required when no publisher client is supplied")
        self._client = client or AsyncEventGridPublisherClient(endpoint, credential)

    async def publish(self, events: Iterable[CustomEvent]) -> bool:
        sdk_events = [event.to_event_grid_event() for event in events]
        if not sdk_events:
            return True
        try:
            await self._client.send(sdk_events)
        except AzureError:
            logger.exception("Failed to publish %d Event Grid event(s)", len(sdk_events))
            return False
        logger.info("Published %d Event Grid event(s)", len(sdk_events))
        return True
