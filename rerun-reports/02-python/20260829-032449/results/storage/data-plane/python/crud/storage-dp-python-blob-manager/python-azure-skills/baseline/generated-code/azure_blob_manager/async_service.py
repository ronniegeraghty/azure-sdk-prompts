"""Asynchronous Azure Blob Storage operations."""

from __future__ import annotations

import asyncio
import os
from pathlib import Path
from typing import Any, AsyncIterator, BinaryIO, Mapping

from azure.core import MatchConditions
from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.storage.blob.aio import BlobLeaseClient, BlobServiceClient

from .models import BlobInfo, LeaseHandle, OperationResult, UploadInfo
from .service import _error_message, _request_options


async def _file_chunks(
    stream: BinaryIO, chunk_size: int
) -> AsyncIterator[bytes]:
    while chunk := await asyncio.to_thread(stream.read, chunk_size):
        yield chunk


class AsyncBlobStorageService:
    """Async, memory-efficient, optimistic-concurrency-safe blob operations."""

    def __init__(
        self,
        client: BlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
        chunk_size: int = 8 * 1024 * 1024,
    ) -> None:
        self._container = client.get_container_client(container_name)
        self._max_concurrency = max_concurrency
        self._chunk_size = chunk_size

    async def upload(
        self,
        source_path: str | Path,
        blob_name: str,
        *,
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        lease_id: str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[UploadInfo]:
        source = Path(source_path)
        blob = self._container.get_blob_client(blob_name)
        options: dict[str, Any] = _request_options(timeout)
        if lease_id:
            options["lease"] = lease_id

        try:
            try:
                current = await blob.get_blob_properties(**_request_options(timeout))
                options["etag"] = current.etag
                options["match_condition"] = MatchConditions.IfNotModified
                overwrite = True
            except ResourceNotFoundError:
                overwrite = False

            with source.open("rb") as data:
                response = await blob.upload_blob(
                    _file_chunks(data, self._chunk_size),
                    length=source.stat().st_size,
                    overwrite=overwrite,
                    metadata=dict(metadata) if metadata else None,
                    tags=dict(tags) if tags else None,
                    max_concurrency=self._max_concurrency,
                    **options,
                )
            info = UploadInfo(
                name=blob_name,
                etag=str(response["etag"]),
                last_modified=response.get("last_modified"),
            )
            return OperationResult(True, f"Uploaded {blob_name!r}", info)
        except (OSError, HttpResponseError) as exc:
            return OperationResult(False, _error_message("Upload", blob_name, exc))

    async def download(
        self,
        blob_name: str,
        destination_path: str | Path,
        *,
        timeout: float | None = None,
    ) -> OperationResult[Path]:
        destination = Path(destination_path)
        temporary = destination.with_name(f"{destination.name}.part")
        try:
            await asyncio.to_thread(destination.parent.mkdir, parents=True, exist_ok=True)
            stream = await self._container.download_blob(
                blob_name,
                max_concurrency=self._max_concurrency,
                **_request_options(timeout),
            )
            with temporary.open("wb") as output:
                async for chunk in stream.chunks():
                    await asyncio.to_thread(output.write, chunk)
            await asyncio.to_thread(os.replace, temporary, destination)
            return OperationResult(
                True, f"Downloaded {blob_name!r} to {destination}", destination
            )
        except (OSError, HttpResponseError) as exc:
            await asyncio.to_thread(temporary.unlink, missing_ok=True)
            return OperationResult(False, _error_message("Download", blob_name, exc))

    async def list_blobs(
        self, *, timeout: float | None = None
    ) -> OperationResult[list[BlobInfo]]:
        try:
            blobs = [
                BlobInfo(
                    name=item.name,
                    size=item.size or 0,
                    etag=item.etag,
                    last_modified=item.last_modified,
                    metadata=item.metadata or {},
                    tags=item.tags or {},
                )
                async for item in self._container.list_blobs(
                    include=["metadata", "tags"], **_request_options(timeout)
                )
            ]
            return OperationResult(True, f"Listed {len(blobs)} blob(s)", blobs)
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("List blobs", None, exc))

    async def delete(
        self,
        blob_name: str,
        *,
        lease_id: str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[None]:
        options: dict[str, Any] = _request_options(timeout)
        if lease_id:
            options["lease"] = lease_id
        try:
            await self._container.delete_blob(blob_name, **options)
            return OperationResult(True, f"Deleted {blob_name!r}")
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("Delete", blob_name, exc))

    async def acquire_lease(
        self,
        blob_name: str,
        *,
        lease_duration: int = 60,
        timeout: float | None = None,
    ) -> OperationResult[LeaseHandle]:
        try:
            lease = await self._container.get_blob_client(blob_name).acquire_lease(
                lease_duration=lease_duration, **_request_options(timeout)
            )
            handle = LeaseHandle(blob_name, lease.id)
            return OperationResult(True, f"Acquired lease for {blob_name!r}", handle)
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("Acquire lease", blob_name, exc))

    async def release_lease(
        self, handle: LeaseHandle, *, timeout: float | None = None
    ) -> OperationResult[None]:
        try:
            blob = self._container.get_blob_client(handle.blob_name)
            await BlobLeaseClient(blob, lease_id=handle.lease_id).release(
                **_request_options(timeout)
            )
            return OperationResult(True, f"Released lease for {handle.blob_name!r}")
        except HttpResponseError as exc:
            return OperationResult(
                False, _error_message("Release lease", handle.blob_name, exc)
            )
