"""Handlers for Azure Blob Storage lifecycle events."""

from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError

LOGGER = logging.getLogger(__name__)


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    marker = "/containers/"
    blob_marker = "/blobs/"
    if marker not in subject:
        raise ValueError(f"Event subject has no container segment: {subject!r}")

    remainder = subject.split(marker, 1)[1]
    if blob_marker not in f"/{remainder}":
        raise ValueError(f"Event subject has no blob segment: {subject!r}")

    container, name = remainder.split(blob_marker, 1)
    if not container or not name:
        raise ValueError(f"Event subject has an empty container or blob name: {subject!r}")
    return BlobLocation(unquote(container), unquote(name))


def _print_summary(location: BlobLocation, properties: Any) -> None:
    content_settings = getattr(properties, "content_settings", None)
    content_type = getattr(content_settings, "content_type", None) or "unknown"
    tier = getattr(properties, "blob_tier", None) or "unknown"
    size = getattr(properties, "size", None)
    print(
        f"Blob created: name={location.name}, size={size}, "
        f"content_type={content_type}, access_tier={tier}"
    )


def _log_unavailable(location: BlobLocation, error: HttpResponseError) -> None:
    LOGGER.warning(
        "Blob %s/%s is no longer readable after its event (%s). "
        "It may have been deleted, replaced, or moved to an offline tier.",
        location.container,
        location.name,
        error,
    )


def _handle_read_error(location: BlobLocation, error: HttpResponseError) -> None:
    transient_blob_states = {
        "BlobArchived",
        "BlobBeingRehydrated",
        "BlobNotFound",
        "ConditionNotMet",
    }
    if (
        isinstance(error, ResourceNotFoundError)
        or getattr(error, "error_code", None) in transient_blob_states
    ):
        _log_unavailable(location, error)
        return
    raise error


def handle_blob_created(event: Any, blob_service: Any) -> None:
    location = parse_blob_subject(event.subject)
    blob_client = blob_service.get_blob_client(location.container, location.name)
    try:
        properties = blob_client.get_blob_properties()
        blob_client.download_blob().readall()
    except HttpResponseError as error:
        _handle_read_error(location, error)
        return
    _print_summary(location, properties)


async def handle_blob_created_async(event: Any, blob_service: Any) -> None:
    location = parse_blob_subject(event.subject)
    blob_client = blob_service.get_blob_client(location.container, location.name)
    try:
        properties = await blob_client.get_blob_properties()
        stream = await blob_client.download_blob()
        await stream.readall()
    except HttpResponseError as error:
        _handle_read_error(location, error)
        return
    _print_summary(location, properties)


def handle_blob_deleted(event: Any) -> None:
    location = parse_blob_subject(event.subject)
    LOGGER.info("Blob deleted: %s/%s", location.container, location.name)


async def handle_blob_deleted_async(event: Any) -> None:
    handle_blob_deleted(event)
