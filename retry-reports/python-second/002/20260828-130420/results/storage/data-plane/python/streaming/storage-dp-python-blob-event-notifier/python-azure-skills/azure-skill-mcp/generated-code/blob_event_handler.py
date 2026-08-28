"""Handlers for Azure Blob Storage lifecycle events."""

from __future__ import annotations

import logging
import re
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceModifiedError, ResourceNotFoundError

from event_receiver import ReceivedEvent

logger = logging.getLogger(__name__)

_SUBJECT_PATTERN = re.compile(
    r"^/blobServices/default/containers/(?P<container>[^/]+)/blobs/(?P<blob>.+)$"
)
_EXPECTED_RACE_ERROR_CODES = {
    "BlobArchived",
    "BlobBeingRehydrated",
    "OperationNotAllowedOnArchivedBlob",
}


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str) -> BlobLocation:
    match = _SUBJECT_PATTERN.match(subject)
    if not match:
        raise ValueError(f"Unsupported blob event subject: {subject!r}")
    return BlobLocation(
        container=unquote(match.group("container")),
        name=unquote(match.group("blob")),
    )


def _event_subject(event: ReceivedEvent) -> str:
    subject = event.subject
    if not subject:
        raise ValueError("Blob event does not have a subject")
    return subject


def _tier(properties: Any) -> str:
    value = getattr(properties, "blob_tier", None) or getattr(
        properties, "access_tier", None
    )
    return str(value or "unknown")


def _content_type(properties: Any) -> str:
    settings = getattr(properties, "content_settings", None)
    return str(getattr(settings, "content_type", None) or "unknown")


def _is_expected_tier_race(error: HttpResponseError) -> bool:
    return getattr(error, "error_code", None) in _EXPECTED_RACE_ERROR_CODES


class BlobEventHandler:
    def __init__(self, blob_service: Any) -> None:
        self._blob_service = blob_service

    def handle_created(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_event_subject(event))
        blob = self._blob_service.get_blob_client(
            container=location.container, blob=location.name
        )
        try:
            properties = blob.get_blob_properties()
            content = blob.download_blob().readall()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            logger.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_tier_race(error):
                raise
            logger.warning(
                "Blob %s/%s is unavailable in its current access tier: %s",
                location.container,
                location.name,
                error,
            )
            return

        print(
            "Blob created: "
            f"name={location.name}, size={len(content)}, "
            f"content_type={_content_type(properties)}, access_tier={_tier(properties)}"
        )

    def handle_deleted(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_event_subject(event))
        logger.info("Blob deleted: %s/%s", location.container, location.name)


class AsyncBlobEventHandler:
    def __init__(self, blob_service: Any) -> None:
        self._blob_service = blob_service

    async def handle_created(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_event_subject(event))
        blob = self._blob_service.get_blob_client(
            container=location.container, blob=location.name
        )
        try:
            properties = await blob.get_blob_properties()
            downloader = await blob.download_blob()
            content = await downloader.readall()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            logger.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_tier_race(error):
                raise
            logger.warning(
                "Blob %s/%s is unavailable in its current access tier: %s",
                location.container,
                location.name,
                error,
            )
            return

        print(
            "Blob created: "
            f"name={location.name}, size={len(content)}, "
            f"content_type={_content_type(properties)}, access_tier={_tier(properties)}"
        )

    async def handle_deleted(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_event_subject(event))
        logger.info("Blob deleted: %s/%s", location.container, location.name)
