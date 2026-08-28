"""Graceful, streaming Azure Blob Storage operations."""

from __future__ import annotations

import asyncio
import math
import os
from dataclasses import dataclass
from pathlib import Path
from typing import Generic, TypeVar

from azure.core import MatchConditions
from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceModifiedError,
    ResourceNotFoundError,
    ServiceRequestError,
)
from azure.storage.blob import BlobLeaseClient, BlobServiceClient
from azure.storage.blob.aio import (
    BlobLeaseClient as AsyncBlobLeaseClient,
    BlobServiceClient as AsyncBlobServiceClient,
)

T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class BlobOperationResult(Generic[T]):
    """A storage operation result that callers can inspect without exceptions."""

    succeeded: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class BlobSummary:
    name: str
    size: int
    etag: str | None


def _request_options(timeout: float | None) -> dict[str, int]:
    if timeout is None:
        return {}
    if timeout <= 0:
        raise ValueError("timeout must be greater than zero")
    seconds = max(1, math.ceil(timeout))
    return {
        "timeout": seconds,
        "connection_timeout": seconds,
        "read_timeout": seconds,
    }


def _failure(
    operation: str, exc: AzureError | OSError | ValueError
) -> BlobOperationResult[object]:
    if isinstance(exc, ResourceNotFoundError):
        detail = "blob or container was not found"
    elif isinstance(exc, ClientAuthenticationError):
        detail = "authentication failed; check the managed identity configuration"
    elif isinstance(exc, ResourceModifiedError):
        detail = "the blob changed concurrently; retry with the latest version"
    elif isinstance(exc, ResourceExistsError):
        detail = "another writer created the blob concurrently"
    elif isinstance(exc, ServiceRequestError):
        detail = "the service could not be reached before the request timed out"
    elif isinstance(exc, HttpResponseError) and exc.status_code in {401, 403}:
        detail = "permission denied; check the identity's Blob Storage data role"
    elif isinstance(exc, HttpResponseError) and (
        "lease" in str(getattr(exc, "error_code", "")).lower()
        or "lease" in str(exc).lower()
    ):
        detail = "the blob lease is held by another client or the lease ID is invalid"
    elif isinstance(exc, OSError):
        detail = f"local file error: {exc}"
    elif isinstance(exc, ValueError):
        detail = f"invalid argument: {exc}"
    else:
        detail = str(exc) or exc.__class__.__name__
    return BlobOperationResult(False, f"{operation} failed: {detail}")


def _timeout_failure(operation: str) -> BlobOperationResult[object]:
    return BlobOperationResult(False, f"{operation} failed: operation timed out")


class BlobStorageService:
    """Synchronous blob operations using streaming transfers and ETag protection."""

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
        source_path: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease: BlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[str]:
        """Stream a file and reject writes based on stale blob state."""
        operation = f"upload {blob_name!r}"
        try:
            options = _request_options(timeout)
            source = Path(source_path)
            size = source.stat().st_size
            blob = self._container.get_blob_client(blob_name)

            try:
                properties = blob.get_blob_properties(**options)
            except ResourceNotFoundError:
                properties = None

            conditions: dict[str, object]
            if properties is None:
                conditions = {"overwrite": False}
            else:
                conditions = {
                    "overwrite": True,
                    "etag": properties.etag,
                    "match_condition": MatchConditions.IfNotModified,
                }

            with source.open("rb") as stream:
                blob.upload_blob(
                    stream,
                    length=size,
                    metadata=metadata,
                    tags=tags,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    **conditions,
                    **options,
                )
            return BlobOperationResult(True, f"uploaded {blob_name!r}", blob.url)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    def download(
        self,
        blob_name: str,
        destination_path: str | Path,
        *,
        timeout: float | None = None,
    ) -> BlobOperationResult[Path]:
        """Stream a blob to disk without buffering the complete blob in memory."""
        operation = f"download {blob_name!r}"
        destination = Path(destination_path)
        temporary = destination.with_name(f"{destination.name}.part")
        try:
            options = _request_options(timeout)
            destination.parent.mkdir(parents=True, exist_ok=True)
            blob = self._container.get_blob_client(blob_name)
            stream = blob.download_blob(
                max_concurrency=self._max_concurrency,
                **options,
            )
            with temporary.open("wb") as output:
                stream.readinto(output)
            os.replace(temporary, destination)
            return BlobOperationResult(
                True, f"downloaded {blob_name!r} to {destination}", destination
            )
        except (AzureError, OSError, ValueError) as exc:
            temporary.unlink(missing_ok=True)
            return _failure(operation, exc)

    def list_blobs(
        self,
        *,
        name_starts_with: str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[list[BlobSummary]]:
        operation = "list blobs"
        try:
            blobs = [
                BlobSummary(blob.name, blob.size or 0, blob.etag)
                for blob in self._container.list_blobs(
                    name_starts_with=name_starts_with,
                    **_request_options(timeout),
                )
            ]
            return BlobOperationResult(True, f"listed {len(blobs)} blob(s)", blobs)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    def delete(
        self,
        blob_name: str,
        *,
        lease: BlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[None]:
        operation = f"delete {blob_name!r}"
        try:
            self._container.delete_blob(
                blob_name,
                delete_snapshots="include",
                lease=lease,
                **_request_options(timeout),
            )
            return BlobOperationResult(True, f"deleted {blob_name!r}")
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: float | None = None,
    ) -> BlobOperationResult[BlobLeaseClient]:
        operation = f"acquire lease for {blob_name!r}"
        try:
            lease = self._container.get_blob_client(blob_name).acquire_lease(
                lease_duration=duration,
                **_request_options(timeout),
            )
            return BlobOperationResult(True, f"acquired lease for {blob_name!r}", lease)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)


class AsyncBlobStorageService:
    """Asynchronous blob operations using streaming transfers and ETag protection."""

    def __init__(
        self,
        client: AsyncBlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
    ) -> None:
        self._container = client.get_container_client(container_name)
        self._max_concurrency = max_concurrency

    async def upload(
        self,
        source_path: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[str]:
        operation = f"upload {blob_name!r}"
        try:
            async with asyncio.timeout(timeout):
                options = _request_options(timeout)
                source = Path(source_path)
                size = await asyncio.to_thread(lambda: source.stat().st_size)
                blob = self._container.get_blob_client(blob_name)

                try:
                    properties = await blob.get_blob_properties(**options)
                except ResourceNotFoundError:
                    properties = None

                conditions: dict[str, object]
                if properties is None:
                    conditions = {"overwrite": False}
                else:
                    conditions = {
                        "overwrite": True,
                        "etag": properties.etag,
                        "match_condition": MatchConditions.IfNotModified,
                    }

                stream = await asyncio.to_thread(source.open, "rb")
                try:
                    await blob.upload_blob(
                        stream,
                        length=size,
                        metadata=metadata,
                        tags=tags,
                        lease=lease,
                        max_concurrency=self._max_concurrency,
                        **conditions,
                        **options,
                    )
                finally:
                    await asyncio.to_thread(stream.close)
            return BlobOperationResult(True, f"uploaded {blob_name!r}", blob.url)
        except TimeoutError:
            return _timeout_failure(operation)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    async def download(
        self,
        blob_name: str,
        destination_path: str | Path,
        *,
        timeout: float | None = None,
    ) -> BlobOperationResult[Path]:
        operation = f"download {blob_name!r}"
        destination = Path(destination_path)
        temporary = destination.with_name(f"{destination.name}.part")
        try:
            async with asyncio.timeout(timeout):
                options = _request_options(timeout)
                await asyncio.to_thread(
                    destination.parent.mkdir, parents=True, exist_ok=True
                )
                stream = await self._container.download_blob(
                    blob_name,
                    max_concurrency=self._max_concurrency,
                    **options,
                )
                output = await asyncio.to_thread(temporary.open, "wb")
                try:
                    async for chunk in stream.chunks():
                        await asyncio.to_thread(output.write, chunk)
                finally:
                    await asyncio.to_thread(output.close)
                await asyncio.to_thread(os.replace, temporary, destination)
            return BlobOperationResult(
                True, f"downloaded {blob_name!r} to {destination}", destination
            )
        except TimeoutError:
            temporary.unlink(missing_ok=True)
            return _timeout_failure(operation)
        except (AzureError, OSError, ValueError) as exc:
            temporary.unlink(missing_ok=True)
            return _failure(operation, exc)

    async def list_blobs(
        self,
        *,
        name_starts_with: str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[list[BlobSummary]]:
        operation = "list blobs"
        try:
            async with asyncio.timeout(timeout):
                blobs = [
                    BlobSummary(blob.name, blob.size or 0, blob.etag)
                    async for blob in self._container.list_blobs(
                        name_starts_with=name_starts_with,
                        **_request_options(timeout),
                    )
                ]
            return BlobOperationResult(True, f"listed {len(blobs)} blob(s)", blobs)
        except TimeoutError:
            return _timeout_failure(operation)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    async def delete(
        self,
        blob_name: str,
        *,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> BlobOperationResult[None]:
        operation = f"delete {blob_name!r}"
        try:
            async with asyncio.timeout(timeout):
                await self._container.delete_blob(
                    blob_name,
                    delete_snapshots="include",
                    lease=lease,
                    **_request_options(timeout),
                )
            return BlobOperationResult(True, f"deleted {blob_name!r}")
        except TimeoutError:
            return _timeout_failure(operation)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)

    async def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: float | None = None,
    ) -> BlobOperationResult[AsyncBlobLeaseClient]:
        operation = f"acquire lease for {blob_name!r}"
        try:
            async with asyncio.timeout(timeout):
                lease = await self._container.get_blob_client(blob_name).acquire_lease(
                    lease_duration=duration,
                    **_request_options(timeout),
                )
            return BlobOperationResult(True, f"acquired lease for {blob_name!r}", lease)
        except TimeoutError:
            return _timeout_failure(operation)
        except (AzureError, OSError, ValueError) as exc:
            return _failure(operation, exc)
