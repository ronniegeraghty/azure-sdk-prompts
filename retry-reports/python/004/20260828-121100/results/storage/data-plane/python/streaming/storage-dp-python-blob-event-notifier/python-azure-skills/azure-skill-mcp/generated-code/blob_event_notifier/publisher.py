"""Sync and async publishers for downstream Event Grid notifications."""

from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any, Mapping, Sequence

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent, EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

LOGGER = logging.getLogger(__name__)


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    subject: str
    data: Mapping[str, Any]
    data_version: str = "1.0"

    def to_event_grid_event(self) -> EventGridEvent:
        subject = "/" + self.subject.strip("/")
        return EventGridEvent(
            subject=subject,
            event_type=self.event_type,
            data=dict(self.data),
            data_version=self.data_version,
        )


class EventPublishError(RuntimeError):
    """Raised when Event Grid rejects a downstream notification."""


class EventPublisher:
    def __init__(
        self,
        topic_endpoint: str,
        client: EventGridPublisherClient | None = None,
    ) -> None:
        self._credential: DefaultAzureCredential | None = None
        if client is None:
            self._credential = DefaultAzureCredential()
            client = EventGridPublisherClient(topic_endpoint, self._credential)
        self._client = client

    def publish(self, events: Sequence[CustomEvent]) -> None:
        sdk_events = [event.to_event_grid_event() for event in events]
        if not sdk_events:
            return
        try:
            self._client.send(sdk_events)
        except AzureError as error:
            LOGGER.error("Failed to publish %d Event Grid event(s): %s", len(events), error)
            raise EventPublishError("Event Grid publishing failed") from error

    def close(self) -> None:
        self._client.close()
        if self._credential is not None:
            self._credential.close()


class AsyncEventPublisher:
    def __init__(
        self,
        topic_endpoint: str,
        client: AsyncEventGridPublisherClient | None = None,
    ) -> None:
        self._credential: AsyncDefaultAzureCredential | None = None
        if client is None:
            self._credential = AsyncDefaultAzureCredential()
            client = AsyncEventGridPublisherClient(topic_endpoint, self._credential)
        self._client = client

    async def publish(self, events: Sequence[CustomEvent]) -> None:
        sdk_events = [event.to_event_grid_event() for event in events]
        if not sdk_events:
            return
        try:
            await self._client.send(sdk_events)
        except AzureError as error:
            LOGGER.error("Failed to publish %d Event Grid event(s): %s", len(events), error)
            raise EventPublishError("Event Grid publishing failed") from error

    async def close(self) -> None:
        await self._client.close()
        if self._credential is not None:
            await self._credential.close()
