from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any, Protocol, Sequence

from azure.core.exceptions import AzureError
from azure.eventgrid import EventGridEvent
from azure.eventgrid import EventGridPublisherClient
from azure.eventgrid.aio import EventGridPublisherClient as AsyncEventGridPublisherClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential


class PublisherClient(Protocol):
    def send(self, events: Sequence[EventGridEvent]) -> None: ...


class AsyncPublisherClient(Protocol):
    async def send(self, events: Sequence[EventGridEvent]) -> None: ...


@dataclass(frozen=True)
class CustomEvent:
    event_type: str
    data: dict[str, Any]
    subject: str | None = None
    data_version: str = "1.0"


def _to_eventgrid_events(
    events: Sequence[CustomEvent],
    default_subject: str | None,
) -> list[EventGridEvent]:
    sdk_events: list[EventGridEvent] = []
    for event in events:
        subject = event.subject or default_subject
        if not subject or not subject.startswith("/"):
            raise ValueError("Each custom event subject must be an absolute hierarchy beginning with '/'")
        sdk_events.append(
            EventGridEvent(
                subject=subject,
                event_type=event.event_type,
                data=event.data,
                data_version=event.data_version,
            )
        )
    return sdk_events


class EventPublisher:
    def __init__(self, client: PublisherClient, logger: logging.Logger | None = None) -> None:
        self._client = client
        self._logger = logger or logging.getLogger(__name__)

    def publish(
        self,
        events: Sequence[CustomEvent],
        default_subject: str | None = None,
    ) -> bool:
        if not events:
            return True
        try:
            self._client.send(_to_eventgrid_events(events, default_subject))
        except AzureError:
            self._logger.exception("Event Grid publishing failed for %d event(s)", len(events))
            return False
        self._logger.info("Published %d downstream event(s)", len(events))
        return True


class AsyncEventPublisher:
    def __init__(
        self,
        client: AsyncPublisherClient,
        logger: logging.Logger | None = None,
    ) -> None:
        self._client = client
        self._logger = logger or logging.getLogger(__name__)

    async def publish(
        self,
        events: Sequence[CustomEvent],
        default_subject: str | None = None,
    ) -> bool:
        if not events:
            return True
        try:
            await self._client.send(_to_eventgrid_events(events, default_subject))
        except AzureError:
            self._logger.exception("Event Grid publishing failed for %d event(s)", len(events))
            return False
        self._logger.info("Published %d downstream event(s)", len(events))
        return True


def publish_custom_events(
    topic_endpoint: str,
    events: Sequence[CustomEvent],
    default_subject: str | None = None,
    logger: logging.Logger | None = None,
) -> bool:
    with DefaultAzureCredential() as credential:
        with EventGridPublisherClient(topic_endpoint, credential) as client:
            return EventPublisher(client, logger).publish(events, default_subject)


async def publish_custom_events_async(
    topic_endpoint: str,
    events: Sequence[CustomEvent],
    default_subject: str | None = None,
    logger: logging.Logger | None = None,
) -> bool:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncEventGridPublisherClient(topic_endpoint, credential) as client:
            return await AsyncEventPublisher(client, logger).publish(events, default_subject)
