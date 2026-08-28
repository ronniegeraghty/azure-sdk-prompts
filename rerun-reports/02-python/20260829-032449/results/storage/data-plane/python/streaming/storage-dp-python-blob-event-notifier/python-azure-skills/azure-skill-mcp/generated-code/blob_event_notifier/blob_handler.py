"""Handlers for Blob Storage lifecycle events."""

from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError

LOGGER = logging.getLogger(__name__)
SUBJECT_MARKER = "/blobServices/default/containers/"
ARCHIVE_ERROR_CODES = {"BlobArchived", "BlobBeingRehydrated", "BlobNotFound"}


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    if SUBJECT_MARKER not in subject:
        raise ValueError(f"Invalid Blob Storage event subject: {subject!r}")

    _, relative_path = subject.split(SUBJECT_MARKER, maxsplit=1)
    try:
        container, encoded_name = relative_path.split("/blobs/", maxsplit=1)
    except ValueError as exc:
        raise ValueError(f"Invalid Blob Storage event subject: {subject!r}") from exc

    if not container or not encoded_name:
        raise ValueError(f"Invalid Blob Storage event subject: {subject!r}")
    return BlobLocation(unquote(container), unquote(encoded_name))


def _access_tier(properties: Any) -> str:
    tier = getattr(properties, "blob_tier", None)
    return str(tier or "unknown")


def _is_lifecycle_race(exc: HttpResponseError) -> bool:
    return getattr(exc, "error_code", None) in ARCHIVE_ERROR_CODES


def handle_blob_created(event: Any, blob_service_client: Any) -> None:
    location = parse_blob_subject(event.subject)
    blob_client = blob_service_client.get_blob_client(
        container=location.container,
        blob=location.name,
    )

    try:
        downloader = blob_client.download_blob()
        content = downloader.readall()
        properties = downloader.properties
    except ResourceNotFoundError:
        LOGGER.warning(
            "Blob %s/%s no longer exists; skipping created event",
            location.container,
            location.name,
        )
        return
    except HttpResponseError as exc:
        if _is_lifecycle_race(exc):
            LOGGER.warning(
                "Blob %s/%s is unavailable after a lifecycle change (%s)",
                location.container,
                location.name,
                exc.error_code,
            )
            return
        raise

    content_settings = getattr(properties, "content_settings", None)
    content_type = getattr(content_settings, "content_type", None) or "unknown"
    size = getattr(properties, "size", None)
    LOGGER.info(
        "Blob created: name=%s size=%s content_type=%s access_tier=%s",
        location.name,
        size if size is not None else len(content),
        content_type,
        _access_tier(properties),
    )


async def handle_blob_created_async(event: Any, blob_service_client: Any) -> None:
    location = parse_blob_subject(event.subject)
    blob_client = blob_service_client.get_blob_client(
        container=location.container,
        blob=location.name,
    )

    try:
        downloader = await blob_client.download_blob()
        content = await downloader.readall()
        properties = downloader.properties
    except ResourceNotFoundError:
        LOGGER.warning(
            "Blob %s/%s no longer exists; skipping created event",
            location.container,
            location.name,
        )
        return
    except HttpResponseError as exc:
        if _is_lifecycle_race(exc):
            LOGGER.warning(
                "Blob %s/%s is unavailable after a lifecycle change (%s)",
                location.container,
                location.name,
                exc.error_code,
            )
            return
        raise

    content_settings = getattr(properties, "content_settings", None)
    content_type = getattr(content_settings, "content_type", None) or "unknown"
    size = getattr(properties, "size", None)
    LOGGER.info(
        "Blob created: name=%s size=%s content_type=%s access_tier=%s",
        location.name,
        size if size is not None else len(content),
        content_type,
        _access_tier(properties),
    )


def handle_blob_deleted(event: Any) -> None:
    location = parse_blob_subject(event.subject)
    LOGGER.info(
        "Blob deleted: container=%s name=%s",
        location.container,
        location.name,
    )


async def handle_blob_deleted_async(event: Any) -> None:
    handle_blob_deleted(event)
