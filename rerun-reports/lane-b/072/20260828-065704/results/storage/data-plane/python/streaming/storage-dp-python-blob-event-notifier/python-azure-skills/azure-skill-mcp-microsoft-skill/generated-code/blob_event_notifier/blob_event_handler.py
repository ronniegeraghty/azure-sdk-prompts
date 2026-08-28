from __future__ import annotations

import logging
from typing import Any, Protocol
from urllib.parse import unquote

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.core.messaging import CloudEvent
from azure.eventgrid import EventGridEvent

StructuredEvent = EventGridEvent | CloudEvent[Any]

_CONTAINER_MARKER = "/containers/"
_BLOB_MARKER = "/blobs/"
_EXPECTED_RACE_ERROR_CODES = {
    "BlobArchived",
    "BlobBeingRehydrated",
    "BlobNotFound",
    "OperationNotAllowedOnArchivedBlob",
}


class BlobService(Protocol):
    def get_blob_client(self, container: str, blob: str) -> Any: ...


class AsyncBlobService(Protocol):
    def get_blob_client(self, container: str, blob: str) -> Any: ...


def parse_blob_subject(subject: str) -> tuple[str, str]:
    container_start = subject.find(_CONTAINER_MARKER)
    blob_start = subject.find(_BLOB_MARKER, container_start + len(_CONTAINER_MARKER))
    if container_start < 0 or blob_start < 0:
        raise ValueError(f"Invalid Blob Storage event subject: {subject!r}")

    container = subject[container_start + len(_CONTAINER_MARKER) : blob_start]
    blob_name = subject[blob_start + len(_BLOB_MARKER) :]
    if not container or not blob_name:
        raise ValueError(f"Invalid Blob Storage event subject: {subject!r}")
    return unquote(container), unquote(blob_name)


def _subject(event: StructuredEvent) -> str:
    subject = event.subject
    if not subject:
        raise ValueError("Blob lifecycle event has no subject")
    return subject


def _property(properties: Any, name: str, default: Any = None) -> Any:
    if isinstance(properties, dict):
        return properties.get(name, default)
    return getattr(properties, name, default)


def _content_type(properties: Any) -> str:
    settings = _property(properties, "content_settings")
    value = _property(settings, "content_type") if settings else None
    return value or "application/octet-stream"


def _access_tier(properties: Any) -> str:
    value = _property(properties, "blob_tier") or _property(properties, "access_tier")
    return str(value or "unknown")


def _is_expected_race(error: HttpResponseError) -> bool:
    return error.status_code in {404, 409} or error.error_code in _EXPECTED_RACE_ERROR_CODES


class BlobEventHandler:
    def __init__(self, blob_service: BlobService, logger: logging.Logger | None = None) -> None:
        self._blob_service = blob_service
        self._logger = logger or logging.getLogger(__name__)

    def handle_created(self, event: StructuredEvent) -> None:
        container, blob_name = parse_blob_subject(_subject(event))
        blob_client = self._blob_service.get_blob_client(container=container, blob=blob_name)
        try:
            downloader = blob_client.download_blob()
            content = downloader.readall()
            properties = downloader.properties
        except ResourceNotFoundError:
            self._logger.warning(
                "Blob %s/%s no longer exists; skipping created event",
                container,
                blob_name,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_race(error):
                raise
            self._logger.warning(
                "Blob %s/%s is unavailable after a lifecycle change (%s)",
                container,
                blob_name,
                error.error_code or error.status_code,
            )
            return

        size = _property(properties, "size", len(content))
        print(
            f"Blob created: name={blob_name}, size={size}, "
            f"content_type={_content_type(properties)}, access_tier={_access_tier(properties)}"
        )

    def handle_deleted(self, event: StructuredEvent) -> None:
        container, blob_name = parse_blob_subject(_subject(event))
        self._logger.info("Blob deleted: %s/%s", container, blob_name)


class AsyncBlobEventHandler:
    def __init__(
        self,
        blob_service: AsyncBlobService,
        logger: logging.Logger | None = None,
    ) -> None:
        self._blob_service = blob_service
        self._logger = logger or logging.getLogger(__name__)

    async def handle_created(self, event: StructuredEvent) -> None:
        container, blob_name = parse_blob_subject(_subject(event))
        blob_client = self._blob_service.get_blob_client(container=container, blob=blob_name)
        try:
            downloader = await blob_client.download_blob()
            content = await downloader.readall()
            properties = downloader.properties
        except ResourceNotFoundError:
            self._logger.warning(
                "Blob %s/%s no longer exists; skipping created event",
                container,
                blob_name,
            )
            return
        except HttpResponseError as error:
            if not _is_expected_race(error):
                raise
            self._logger.warning(
                "Blob %s/%s is unavailable after a lifecycle change (%s)",
                container,
                blob_name,
                error.error_code or error.status_code,
            )
            return

        size = _property(properties, "size", len(content))
        print(
            f"Blob created: name={blob_name}, size={size}, "
            f"content_type={_content_type(properties)}, access_tier={_access_tier(properties)}"
        )

    async def handle_deleted(self, event: StructuredEvent) -> None:
        container, blob_name = parse_blob_subject(_subject(event))
        self._logger.info("Blob deleted: %s/%s", container, blob_name)
