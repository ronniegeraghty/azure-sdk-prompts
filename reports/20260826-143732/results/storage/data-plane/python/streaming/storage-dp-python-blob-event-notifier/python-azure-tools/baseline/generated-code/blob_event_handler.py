from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

logger = logging.getLogger(__name__)

BlobEvent = EventGridEvent | CloudEvent[Any]


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    container_marker = "/containers/"
    blob_marker = "/blobs/"
    if container_marker not in subject or blob_marker not in subject:
        raise ValueError(f"Invalid blob event subject: {subject!r}")

    _, container_and_blob = subject.split(container_marker, 1)
    container, separator, blob_name = container_and_blob.partition(blob_marker)
    if not separator or not container or not blob_name:
        raise ValueError(f"Invalid blob event subject: {subject!r}")
    return BlobLocation(unquote(container), unquote(blob_name))


def _subject(event: BlobEvent) -> str:
    if not event.subject:
        raise ValueError("Blob event is missing a subject")
    return event.subject


def _tier_name(properties: Any) -> str:
    tier = getattr(properties, "blob_tier", None)
    return getattr(tier, "value", tier) or "Unknown"


def _content_type(properties: Any) -> str:
    content_settings = getattr(properties, "content_settings", None)
    return getattr(content_settings, "content_type", None) or "application/octet-stream"


def _is_tier_race(error: HttpResponseError) -> bool:
    return error.error_code in {
        "BlobArchived",
        "BlobBeingRehydrated",
        "BlobOperationNotSupported",
    }


class BlobEventHandler:
    def __init__(self, blob_service: Any) -> None:
        self._blob_service = blob_service

    def handle_created(self, event: BlobEvent) -> None:
        location = parse_blob_subject(_subject(event))
        blob = self._blob_service.get_blob_client(location.container, location.name)
        try:
            downloader = blob.download_blob()
            content = downloader.readall()
            properties = downloader.properties
        except ResourceNotFoundError:
            logger.warning(
                "Blob %s/%s no longer exists; skipping created event",
                location.container,
                location.name,
            )
            return
        except HttpResponseError as error:
            if _is_tier_race(error):
                logger.warning(
                    "Blob %s/%s changed access tier before it could be read: %s",
                    location.container,
                    location.name,
                    error.error_code,
                )
                return
            raise

        size = getattr(properties, "size", None)
        if size is None:
            size = len(content)
        print(
            f"Blob created: name={location.name}, size={size}, "
            f"content_type={_content_type(properties)}, tier={_tier_name(properties)}"
        )

    def handle_deleted(self, event: BlobEvent) -> None:
        location = parse_blob_subject(_subject(event))
        logger.info("Blob deleted: %s/%s", location.container, location.name)


class AsyncBlobEventHandler:
    def __init__(self, blob_service: Any) -> None:
        self._blob_service = blob_service

    async def handle_created(self, event: BlobEvent) -> None:
        location = parse_blob_subject(_subject(event))
        blob = self._blob_service.get_blob_client(location.container, location.name)
        try:
            downloader = await blob.download_blob()
            content = await downloader.readall()
            properties = downloader.properties
        except ResourceNotFoundError:
            logger.warning(
                "Blob %s/%s no longer exists; skipping created event",
                location.container,
                location.name,
            )
            return
        except HttpResponseError as error:
            if _is_tier_race(error):
                logger.warning(
                    "Blob %s/%s changed access tier before it could be read: %s",
                    location.container,
                    location.name,
                    error.error_code,
                )
                return
            raise

        size = getattr(properties, "size", None)
        if size is None:
            size = len(content)
        print(
            f"Blob created: name={location.name}, size={size}, "
            f"content_type={_content_type(properties)}, tier={_tier_name(properties)}"
        )

    async def handle_deleted(self, event: BlobEvent) -> None:
        location = parse_blob_subject(_subject(event))
        logger.info("Blob deleted: %s/%s", location.container, location.name)
