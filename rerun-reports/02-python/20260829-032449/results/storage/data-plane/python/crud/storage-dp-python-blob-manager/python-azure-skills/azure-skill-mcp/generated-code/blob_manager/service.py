"""High-level synchronous and asynchronous Azure Blob Storage operations."""

from __future__ import annotations

import asyncio
import logging
from dataclasses import dataclass
from pathlib import Path
from typing import Any, AsyncIterator, Generic, TypeVar

from azure.core import MatchConditions
from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ServiceRequestError,
    ServiceResponseError,
)
from azure.storage.blob import BlobLeaseClient, BlobServiceClient
from azure.storage.blob.aio import BlobLeaseClient as AsyncBlobLeaseClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

LOGGER = logging.getLogger(__name__)
T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    succeeded: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class BlobSummary:
    name: str
    size: int
    etag: str | None
    metadata: dict[str, str]
    tags: dict[str, str] | None


def _failure(operation: str, blob_name: str | None, exc: Exception) -> OperationResult[Any]:
    target = f" for blob {blob_name!r}" if blob_name else ""
    if isinstance(exc, ResourceNotFoundError):
        detail = "The blob or container was not found."
    elif isinstance(exc, ClientAuthenticationError):
        detail = "Authentication failed. Check the managed identity and Azure RBAC role."
    elif isinstance(exc, ResourceExistsError):
        detail = "Another writer created the blob before this operation completed."
    elif isinstance(exc, HttpResponseError) and getattr(exc, "error_code", None) in {
        "LeaseAlreadyPresent",
        "LeaseIdMissing",
        "LeaseIdMismatchWithBlobOperation",
        "LeaseNotPresentWithBlobOperation",
    }:
        detail = "The blob lease is held by another client or the lease ID is invalid."
    elif isinstance(exc, HttpResponseError) and exc.status_code in {409, 412}:
        detail = (
            "The blob changed concurrently or is protected by another client's lease."
        )
    elif isinstance(exc, HttpResponseError) and exc.status_code == 403:
        detail = "Permission denied. Check the identity's Blob Data RBAC role."
    elif isinstance(exc, (ServiceRequestError, ServiceResponseError)):
        detail = "Azure Storage could not be reached or returned an invalid response."
    else:
        detail = str(exc) or type(exc).__name__

    message = f"{operation} failed{target}: {detail}"
    LOGGER.warning(message)
    return OperationResult(False, message)


class BlobStorageService:
    def __init__(
        self,
        client: BlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
    ) -> None:
        self._container = client.get_container_client(container_name)
        self._max_concurrency = max_concurrency

    def upload(
        self,
        source: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        timeout: int | None = None,
        lease: str | BlobLeaseClient | None = None,
    ) -> OperationResult[str]:
        path = Path(source)
        if not path.is_file():
            return OperationResult(False, f"Upload failed: file not found: {path}")

        blob = self._container.get_blob_client(blob_name)
        request_options: dict[str, Any] = {}
        if timeout is not None:
            request_options["timeout"] = timeout

        try:
            try:
                properties = blob.get_blob_properties(**request_options)
            except ResourceNotFoundError:
                properties = None

            conditions: dict[str, Any]
            if properties is None:
                conditions = {"overwrite": False}
            else:
                conditions = {
                    "overwrite": True,
                    "etag": properties.etag,
                    "match_condition": MatchConditions.IfNotModified,
                }

            with path.open("rb") as stream:
                response = blob.upload_blob(
                    stream,
                    length=path.stat().st_size,
                    metadata=metadata,
                    tags=tags,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    **conditions,
                    **request_options,
                )
            return OperationResult(
                True,
                f"Uploaded {path} to {blob_name!r}.",
                response.get("etag"),
            )
        except (AzureError, OSError) as exc:
            return _failure("Upload", blob_name, exc)

    def download(
        self,
        blob_name: str,
        destination: str | Path,
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        path = Path(destination)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            path.parent.mkdir(parents=True, exist_ok=True)
            with path.open("wb") as stream:
                self._container.download_blob(
                    blob_name,
                    max_concurrency=self._max_concurrency,
                    **request_options,
                ).readinto(stream)
            return OperationResult(True, f"Downloaded {blob_name!r} to {path}.", path)
        except (AzureError, OSError) as exc:
            path.unlink(missing_ok=True)
            return _failure("Download", blob_name, exc)

    def list_blobs(
        self,
        *,
        include_tags: bool = True,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobSummary]]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            include = ["metadata", "tags"] if include_tags else ["metadata"]
            blobs = [
                BlobSummary(
                    name=item.name,
                    size=item.size or 0,
                    etag=item.etag,
                    metadata=item.metadata or {},
                    tags=getattr(item, "tags", None),
                )
                for item in self._container.list_blobs(
                    include=include, **request_options
                )
            ]
            return OperationResult(True, f"Listed {len(blobs)} blob(s).", blobs)
        except AzureError as exc:
            return _failure("List blobs", None, exc)

    def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: int | None = None,
    ) -> OperationResult[BlobLeaseClient]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            lease = BlobLeaseClient(self._container.get_blob_client(blob_name))
            lease.acquire(lease_duration=duration, **request_options)
            return OperationResult(True, f"Acquired lease for {blob_name!r}.", lease)
        except AzureError as exc:
            return _failure("Acquire lease", blob_name, exc)

    def delete(
        self,
        blob_name: str,
        *,
        timeout: int | None = None,
        lease: str | BlobLeaseClient | None = None,
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            self._container.delete_blob(
                blob_name,
                lease=lease,
                delete_snapshots="include",
                **request_options,
            )
            return OperationResult(True, f"Deleted {blob_name!r}.")
        except AzureError as exc:
            return _failure("Delete", blob_name, exc)


class AsyncBlobStorageService:
    def __init__(
        self,
        client: AsyncBlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
        block_size: int = 8 * 1024 * 1024,
    ) -> None:
        self._container = client.get_container_client(container_name)
        self._max_concurrency = max_concurrency
        self._block_size = block_size

    async def _file_chunks(self, path: Path) -> AsyncIterator[bytes]:
        with path.open("rb") as stream:
            while chunk := await asyncio.to_thread(stream.read, self._block_size):
                yield chunk

    async def upload(
        self,
        source: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        timeout: int | None = None,
        lease: str | AsyncBlobLeaseClient | None = None,
    ) -> OperationResult[str]:
        path = Path(source)
        if not path.is_file():
            return OperationResult(False, f"Upload failed: file not found: {path}")

        blob = self._container.get_blob_client(blob_name)
        request_options: dict[str, Any] = {}
        if timeout is not None:
            request_options["timeout"] = timeout

        try:
            try:
                properties = await blob.get_blob_properties(**request_options)
            except ResourceNotFoundError:
                properties = None

            conditions: dict[str, Any]
            if properties is None:
                conditions = {"overwrite": False}
            else:
                conditions = {
                    "overwrite": True,
                    "etag": properties.etag,
                    "match_condition": MatchConditions.IfNotModified,
                }

            response = await blob.upload_blob(
                self._file_chunks(path),
                length=path.stat().st_size,
                metadata=metadata,
                tags=tags,
                lease=lease,
                max_concurrency=self._max_concurrency,
                **conditions,
                **request_options,
            )
            return OperationResult(
                True,
                f"Uploaded {path} to {blob_name!r}.",
                response.get("etag"),
            )
        except (AzureError, OSError) as exc:
            return _failure("Upload", blob_name, exc)

    async def download(
        self,
        blob_name: str,
        destination: str | Path,
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        path = Path(destination)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            path.parent.mkdir(parents=True, exist_ok=True)
            downloader = await self._container.download_blob(
                blob_name,
                max_concurrency=self._max_concurrency,
                **request_options,
            )
            with path.open("wb") as stream:
                async for chunk in downloader.chunks():
                    await asyncio.to_thread(stream.write, chunk)
            return OperationResult(True, f"Downloaded {blob_name!r} to {path}.", path)
        except (AzureError, OSError) as exc:
            path.unlink(missing_ok=True)
            return _failure("Download", blob_name, exc)

    async def list_blobs(
        self,
        *,
        include_tags: bool = True,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobSummary]]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            include = ["metadata", "tags"] if include_tags else ["metadata"]
            blobs = [
                BlobSummary(
                    name=item.name,
                    size=item.size or 0,
                    etag=item.etag,
                    metadata=item.metadata or {},
                    tags=getattr(item, "tags", None),
                )
                async for item in self._container.list_blobs(
                    include=include, **request_options
                )
            ]
            return OperationResult(True, f"Listed {len(blobs)} blob(s).", blobs)
        except AzureError as exc:
            return _failure("List blobs", None, exc)

    async def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: int | None = None,
    ) -> OperationResult[AsyncBlobLeaseClient]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            lease = AsyncBlobLeaseClient(
                self._container.get_blob_client(blob_name)
            )
            await lease.acquire(lease_duration=duration, **request_options)
            return OperationResult(True, f"Acquired lease for {blob_name!r}.", lease)
        except AzureError as exc:
            return _failure("Acquire lease", blob_name, exc)

    async def delete(
        self,
        blob_name: str,
        *,
        timeout: int | None = None,
        lease: str | AsyncBlobLeaseClient | None = None,
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            await self._container.delete_blob(
                blob_name,
                lease=lease,
                delete_snapshots="include",
                **request_options,
            )
            return OperationResult(True, f"Deleted {blob_name!r}.")
        except AzureError as exc:
            return _failure("Delete", blob_name, exc)
