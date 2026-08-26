from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import (
    HttpResponseError,
    ResourceModifiedError,
    ResourceNotFoundError,
)
from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

logger = logging.getLogger(__name__)

BlobEvent = EventGridEvent | CloudEvent[Any]


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    marker = "/containers/"
    blob_marker = "/blobs/"
    if marker not in subject:
        raise ValueError(f"Blob event subject has no container segment: {subject}")

    _, remainder = subject.split(marker, 1)
    if blob_marker not in f"/{remainder}":
        raise ValueError(f"Blob event subject has no blob segment: {subject}")

    container, blob_name = remainder.split(blob_marker, 1)
    if not container or not blob_name:
        raise ValueError(f"Blob event subject is incomplete: {subject}")
    return BlobLocation(container=unquote(container), name=unquote(blob_name))


def handle_blob_created(event: BlobEvent, blob_service_client: Any) -> None:
    location = parse_blob_subject(event.subject or "")
    blob_client = blob_service_client.get_blob_client(
        container=location.container,
        blob=location.name,
    )

    try:
        properties = blob_client.get_blob_properties()
        downloader = blob_client.download_blob()
        downloaded_size = sum(len(chunk) for chunk in downloader.chunks())
    except (ResourceNotFoundError, ResourceModifiedError) as exc:
        logger.warning(
            "Blob %s/%s changed or disappeared before it could be read: %s",
            location.container,
            location.name,
            exc,
        )
        return
    except HttpResponseError as exc:
        if _is_tier_transition_error(exc):
            logger.warning(
                "Blob %s/%s is unavailable while its access tier changes: %s",
                location.container,
                location.name,
                exc,
            )
            return
        raise

    content_type = properties.content_settings.content_type or "unknown"
    access_tier = properties.blob_tier or "unknown"
    logger.info(
        "Blob created: name=%s size=%s content_type=%s access_tier=%s",
        location.name,
        downloaded_size,
        content_type,
        access_tier,
    )


async def handle_blob_created_async(
    event: BlobEvent, blob_service_client: Any
) -> None:
    location = parse_blob_subject(event.subject or "")
    blob_client = blob_service_client.get_blob_client(
        container=location.container,
        blob=location.name,
    )

    try:
        properties = await blob_client.get_blob_properties()
        downloader = await blob_client.download_blob()
        downloaded_size = 0
        async for chunk in downloader.chunks():
            downloaded_size += len(chunk)
    except (ResourceNotFoundError, ResourceModifiedError) as exc:
        logger.warning(
            "Blob %s/%s changed or disappeared before it could be read: %s",
            location.container,
            location.name,
            exc,
        )
        return
    except HttpResponseError as exc:
        if _is_tier_transition_error(exc):
            logger.warning(
                "Blob %s/%s is unavailable while its access tier changes: %s",
                location.container,
                location.name,
                exc,
            )
            return
        raise

    content_type = properties.content_settings.content_type or "unknown"
    access_tier = properties.blob_tier or "unknown"
    logger.info(
        "Blob created: name=%s size=%s content_type=%s access_tier=%s",
        location.name,
        downloaded_size,
        content_type,
        access_tier,
    )


def handle_blob_deleted(event: BlobEvent) -> None:
    location = parse_blob_subject(event.subject or "")
    logger.info(
        "Blob deleted: container=%s name=%s",
        location.container,
        location.name,
    )


async def handle_blob_deleted_async(event: BlobEvent) -> None:
    handle_blob_deleted(event)


def _is_tier_transition_error(exc: HttpResponseError) -> bool:
    error_code = getattr(exc, "error_code", None)
    return exc.status_code == 409 and error_code in {
        "BlobArchived",
        "BlobBeingRehydrated",
        "BlobOperationNotSupported",
    }
