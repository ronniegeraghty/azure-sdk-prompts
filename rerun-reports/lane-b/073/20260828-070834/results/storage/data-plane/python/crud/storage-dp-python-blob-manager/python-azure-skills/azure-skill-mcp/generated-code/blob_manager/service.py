"""Synchronous and asynchronous Azure Blob Storage management services."""

from __future__ import annotations

import asyncio
import os
from dataclasses import dataclass
from pathlib import Path
from typing import AsyncIterator, Generic, TypeVar

from azure.core import MatchConditions
from azure.core.exceptions import (
    AzureError,
    HttpResponseError,
    ResourceExistsError,
    ResourceModifiedError,
    ResourceNotFoundError,
)
from azure.storage.blob import BlobLeaseClient, BlobServiceClient
from azure.storage.blob.aio import BlobLeaseClient as AsyncBlobLeaseClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

T = TypeVar("T")


@dataclass(frozen=True)
class OperationResult(Generic[T]):
    """A storage operation outcome that callers can handle without exceptions."""

    succeeded: bool
    message: str
    value: T | None = None


def _error_result(operation: str, error: Exception) -> OperationResult[None]:
    if isinstance(error, ResourceNotFoundError):
        detail = "the container or blob was not found"
    elif isinstance(error, ResourceModifiedError):
        detail = "the blob changed concurrently; retry with the latest version"
    elif isinstance(error, ResourceExistsError):
        detail = "the blob already exists or is leased by another client"
    elif isinstance(error, HttpResponseError) and error.status_code == 403:
        detail = "permission was denied; verify the managed identity RBAC role"
    elif isinstance(error, HttpResponseError) and error.status_code == 409:
        detail = (
            "the request conflicted with the blob state, possibly because another "
            "client holds its lease"
        )
    elif isinstance(error, HttpResponseError) and error.status_code == 412:
        detail = "the blob changed concurrently or the supplied lease is invalid"
    elif isinstance(error, OSError):
        detail = f"local file error: {error}"
    else:
        detail = str(error) or error.__class__.__name__
    return OperationResult(False, f"{operation} failed: {detail}.")


class BlobStorageManager:
    """Reusable synchronous blob operations."""

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
        source: str | os.PathLike[str],
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease: BlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        """Stream a file and conditionally create or replace a block blob."""
        source_path = Path(source)
        if not source_path.is_file():
            return OperationResult(False, f"Upload failed: file '{source_path}' not found.")

        blob = self._container.get_blob_client(blob_name)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            try:
                etag = blob.get_blob_properties(**request_options).etag
            except ResourceNotFoundError:
                etag = None

            with source_path.open("rb") as stream:
                if etag is None:
                    blob.upload_blob(
                        stream,
                        length=source_path.stat().st_size,
                        overwrite=False,
                        metadata=metadata,
                        tags=tags,
                        lease=lease,
                        max_concurrency=self._max_concurrency,
                        **request_options,
                    )
                else:
                    blob.upload_blob(
                        stream,
                        length=source_path.stat().st_size,
                        overwrite=True,
                        metadata=metadata,
                        tags=tags,
                        lease=lease,
                        etag=etag,
                        match_condition=MatchConditions.IfNotModified,
                        max_concurrency=self._max_concurrency,
                        **request_options,
                    )
            return OperationResult(True, f"Uploaded '{source_path}' to '{blob_name}'.")
        except (AzureError, OSError) as error:
            return _error_result("Upload", error)

    def download(
        self,
        blob_name: str,
        destination: str | os.PathLike[str],
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        """Download directly to a file without buffering the full blob."""
        destination_path = Path(destination)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            with destination_path.open("wb") as stream:
                downloader = self._container.download_blob(
                    blob_name,
                    max_concurrency=self._max_concurrency,
                    **request_options,
                )
                downloader.readinto(stream)
            return OperationResult(
                True, f"Downloaded '{blob_name}' to '{destination_path}'.", destination_path
            )
        except (AzureError, OSError) as error:
            try:
                destination_path.unlink(missing_ok=True)
            except OSError:
                pass
            return _error_result("Download", error)

    def list_blobs(
        self, *, name_starts_with: str | None = None, timeout: int | None = None
    ) -> OperationResult[list[str]]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            names = [
                blob.name
                for blob in self._container.list_blobs(
                    name_starts_with=name_starts_with, **request_options
                )
            ]
            return OperationResult(True, f"Listed {len(names)} blob(s).", names)
        except AzureError as error:
            return _error_result("List", error)

    def delete(
        self,
        blob_name: str,
        *,
        lease: BlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            self._container.delete_blob(
                blob_name, lease=lease, delete_snapshots="include", **request_options
            )
            return OperationResult(True, f"Deleted '{blob_name}'.")
        except AzureError as error:
            return _error_result("Delete", error)

    def acquire_lease(
        self, blob_name: str, *, duration: int = 60, timeout: int | None = None
    ) -> OperationResult[BlobLeaseClient]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            lease = BlobLeaseClient(self._container.get_blob_client(blob_name))
            lease.acquire(lease_duration=duration, **request_options)
            return OperationResult(True, f"Acquired a lease on '{blob_name}'.", lease)
        except AzureError as error:
            return _error_result("Acquire lease", error)

    def release_lease(
        self, lease: BlobLeaseClient, *, timeout: int | None = None
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            lease.release(**request_options)
            return OperationResult(True, "Released blob lease.")
        except AzureError as error:
            return _error_result("Release lease", error)


async def _file_chunks(path: Path, chunk_size: int) -> AsyncIterator[bytes]:
    stream = await asyncio.to_thread(path.open, "rb")
    try:
        while chunk := await asyncio.to_thread(stream.read, chunk_size):
            yield chunk
    finally:
        await asyncio.to_thread(stream.close)


class AsyncBlobStorageManager:
    """Reusable asynchronous blob operations."""

    def __init__(
        self,
        client: AsyncBlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
        upload_chunk_size: int = 4 * 1024 * 1024,
    ) -> None:
        self._container = client.get_container_client(container_name)
        self._max_concurrency = max_concurrency
        self._upload_chunk_size = upload_chunk_size

    async def upload(
        self,
        source: str | os.PathLike[str],
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        source_path = Path(source)
        if not source_path.is_file():
            return OperationResult(False, f"Upload failed: file '{source_path}' not found.")

        blob = self._container.get_blob_client(blob_name)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            try:
                etag = (await blob.get_blob_properties(**request_options)).etag
            except ResourceNotFoundError:
                etag = None

            upload_options = {
                "length": source_path.stat().st_size,
                "metadata": metadata,
                "tags": tags,
                "lease": lease,
                "max_concurrency": self._max_concurrency,
                **request_options,
            }
            chunks = _file_chunks(source_path, self._upload_chunk_size)
            try:
                if etag is None:
                    await blob.upload_blob(chunks, overwrite=False, **upload_options)
                else:
                    await blob.upload_blob(
                        chunks,
                        overwrite=True,
                        etag=etag,
                        match_condition=MatchConditions.IfNotModified,
                        **upload_options,
                    )
            finally:
                await chunks.aclose()
            return OperationResult(True, f"Uploaded '{source_path}' to '{blob_name}'.")
        except (AzureError, OSError) as error:
            return _error_result("Upload", error)

    async def download(
        self,
        blob_name: str,
        destination: str | os.PathLike[str],
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            downloader = await self._container.download_blob(
                blob_name,
                max_concurrency=self._max_concurrency,
                **request_options,
            )
            stream = await asyncio.to_thread(destination_path.open, "wb")
            try:
                async for chunk in downloader.chunks():
                    await asyncio.to_thread(stream.write, chunk)
            finally:
                await asyncio.to_thread(stream.close)
            return OperationResult(
                True, f"Downloaded '{blob_name}' to '{destination_path}'.", destination_path
            )
        except (AzureError, OSError) as error:
            try:
                destination_path.unlink(missing_ok=True)
            except OSError:
                pass
            return _error_result("Download", error)

    async def list_blobs(
        self, *, name_starts_with: str | None = None, timeout: int | None = None
    ) -> OperationResult[list[str]]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            names = [
                blob.name
                async for blob in self._container.list_blobs(
                    name_starts_with=name_starts_with, **request_options
                )
            ]
            return OperationResult(True, f"Listed {len(names)} blob(s).", names)
        except AzureError as error:
            return _error_result("List", error)

    async def delete(
        self,
        blob_name: str,
        *,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            await self._container.delete_blob(
                blob_name, lease=lease, delete_snapshots="include", **request_options
            )
            return OperationResult(True, f"Deleted '{blob_name}'.")
        except AzureError as error:
            return _error_result("Delete", error)

    async def acquire_lease(
        self, blob_name: str, *, duration: int = 60, timeout: int | None = None
    ) -> OperationResult[AsyncBlobLeaseClient]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            lease = AsyncBlobLeaseClient(self._container.get_blob_client(blob_name))
            await lease.acquire(lease_duration=duration, **request_options)
            return OperationResult(True, f"Acquired a lease on '{blob_name}'.", lease)
        except AzureError as error:
            return _error_result("Acquire lease", error)

    async def release_lease(
        self, lease: AsyncBlobLeaseClient, *, timeout: int | None = None
    ) -> OperationResult[None]:
        request_options = {"timeout": timeout} if timeout is not None else {}
        try:
            await lease.release(**request_options)
            return OperationResult(True, "Released blob lease.")
        except AzureError as error:
            return _error_result("Release lease", error)
