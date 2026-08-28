"""Handlers for Azure Blob Storage lifecycle events."""

from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceModifiedError, ResourceNotFoundError

LOGGER = logging.getLogger(__name__)

_ARCHIVE_ERROR_CODES = {
    "BlobArchived",
    "BlobBeingRehydrated",
    "OperationNotAllowedOnArchivedBlob",
}


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    marker = "/containers/"
    blob_marker = "/blobs/"
    if marker not in subject:
        raise ValueError(f"Blob event subject has no container segment: {subject!r}")

    _, container_and_blob = subject.split(marker, 1)
    if blob_marker not in container_and_blob:
        raise ValueError(f"Blob event subject has no blob segment: {subject!r}")

    container, name = container_and_blob.split(blob_marker, 1)
    if not container or not name:
        raise ValueError(f"Blob event subject is incomplete: {subject!r}")
    return BlobLocation(unquote(container), unquote(name))


class BlobEventHandler:
    def __init__(self, blob_service_client: Any) -> None:
        self._blob_service_client = blob_service_client

    def handle_blob_created(self, event: Any) -> None:
        location = parse_blob_subject(event.subject)
        blob_client = self._blob_service_client.get_blob_client(
            container=location.container,
            blob=location.name,
        )

        try:
            content = blob_client.download_blob().readall()
            properties = blob_client.get_blob_properties()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            LOGGER.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if getattr(error, "error_code", None) in _ARCHIVE_ERROR_CODES:
                LOGGER.warning(
                    "Blob %s/%s is archived or rehydrating and cannot be downloaded",
                    location.container,
                    location.name,
                )
                return
            raise

        _print_summary(location, properties, len(content))

    def handle_blob_deleted(self, event: Any) -> None:
        location = parse_blob_subject(event.subject)
        LOGGER.info("Blob deleted: %s/%s", location.container, location.name)


class AsyncBlobEventHandler:
    def __init__(self, blob_service_client: Any) -> None:
        self._blob_service_client = blob_service_client

    async def handle_blob_created(self, event: Any) -> None:
        location = parse_blob_subject(event.subject)
        blob_client = self._blob_service_client.get_blob_client(
            container=location.container,
            blob=location.name,
        )

        try:
            downloader = await blob_client.download_blob()
            content = await downloader.readall()
            properties = await blob_client.get_blob_properties()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            LOGGER.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if getattr(error, "error_code", None) in _ARCHIVE_ERROR_CODES:
                LOGGER.warning(
                    "Blob %s/%s is archived or rehydrating and cannot be downloaded",
                    location.container,
                    location.name,
                )
                return
            raise

        _print_summary(location, properties, len(content))

    async def handle_blob_deleted(self, event: Any) -> None:
        location = parse_blob_subject(event.subject)
        LOGGER.info("Blob deleted: %s/%s", location.container, location.name)


def _print_summary(location: BlobLocation, properties: Any, downloaded_size: int) -> None:
    content_settings = getattr(properties, "content_settings", None)
    content_type = getattr(content_settings, "content_type", None) or "unknown"
    access_tier = getattr(properties, "blob_tier", None) or "unknown"
    if hasattr(access_tier, "value"):
        access_tier = access_tier.value
    size = getattr(properties, "size", None)
    if size is None:
        size = downloaded_size

    print(
        "Blob summary: "
        f"name={location.name}, size={size}, "
        f"content_type={content_type}, access_tier={access_tier}"
    )
