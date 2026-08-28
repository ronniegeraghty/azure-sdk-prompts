from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceModifiedError, ResourceNotFoundError

from event_receiver import ReceivedEvent

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class BlobLocation:
    container: str
    name: str


def parse_blob_subject(subject: str | None) -> BlobLocation:
    if not subject:
        raise ValueError("Blob event is missing its subject")

    container_marker = "/containers/"
    blob_marker = "/blobs/"
    if container_marker not in subject or blob_marker not in subject:
        raise ValueError(f"Unexpected blob event subject: {subject}")

    container_and_blob = subject.split(container_marker, 1)[1]
    container, separator, blob_name = container_and_blob.partition(blob_marker)
    if not separator or not container or not blob_name:
        raise ValueError(f"Unexpected blob event subject: {subject}")
    return BlobLocation(unquote(container), unquote(blob_name))


def _subject(event: ReceivedEvent) -> str | None:
    return event.subject


def _is_expected_race(error: HttpResponseError) -> bool:
    return error.status_code in {404, 409, 412} or getattr(error, "error_code", None) in {
        "BlobArchived",
        "BlobNotFound",
        "ConditionNotMet",
    }


class BlobEventHandler:
    def __init__(self, blob_service_client: Any) -> None:
        self._client = blob_service_client

    def handle_created(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_subject(event))
        blob_client = self._client.get_blob_client(location.container, location.name)
        try:
            properties = blob_client.get_blob_properties()
            downloader = blob_client.download_blob()
            downloader.readall()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            logger.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_race(error):
                raise
            logger.warning(
                "Blob %s/%s is no longer readable at its original tier: %s",
                location.container,
                location.name,
                error,
            )
            return

        print(
            "Blob created: "
            f"name={location.name}, size={properties.size}, "
            f"content_type={properties.content_settings.content_type}, "
            f"access_tier={properties.blob_tier}"
        )

    def handle_deleted(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_subject(event))
        logger.info("Blob deleted: %s/%s", location.container, location.name)


class AsyncBlobEventHandler:
    def __init__(self, blob_service_client: Any) -> None:
        self._client = blob_service_client

    async def handle_created(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_subject(event))
        blob_client = self._client.get_blob_client(location.container, location.name)
        try:
            properties = await blob_client.get_blob_properties()
            downloader = await blob_client.download_blob()
            await downloader.readall()
        except (ResourceNotFoundError, ResourceModifiedError) as error:
            logger.warning(
                "Blob %s/%s changed or disappeared before it could be read: %s",
                location.container,
                location.name,
                error,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_race(error):
                raise
            logger.warning(
                "Blob %s/%s is no longer readable at its original tier: %s",
                location.container,
                location.name,
                error,
            )
            return

        print(
            "Blob created (async): "
            f"name={location.name}, size={properties.size}, "
            f"content_type={properties.content_settings.content_type}, "
            f"access_tier={properties.blob_tier}"
        )

    async def handle_deleted(self, event: ReceivedEvent) -> None:
        location = parse_blob_subject(_subject(event))
        logger.info("Blob deleted (async): %s/%s", location.container, location.name)
