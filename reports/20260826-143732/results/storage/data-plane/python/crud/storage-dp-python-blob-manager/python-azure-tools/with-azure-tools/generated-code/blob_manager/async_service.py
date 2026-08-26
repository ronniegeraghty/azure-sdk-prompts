"""Asynchronous Azure Blob Storage management service."""

from __future__ import annotations

from pathlib import Path
from types import TracebackType
from typing import Mapping

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError
from azure.storage.blob.aio import BlobLeaseClient

from .config import BlobStorageSettings
from .errors import STORAGE_EXCEPTIONS, storage_failure
from .models import BlobInfo, LeaseHandle, OperationResult


class AsyncBlobStorageManager:
    """Memory-efficient asynchronous blob operations with safe conditional writes."""

    def __init__(self, settings: BlobStorageSettings) -> None:
        self._settings = settings
        self._credential, self._client = settings.create_async_client()

    async def __aenter__(self) -> "AsyncBlobStorageManager":
        await self._credential.__aenter__()
        await self._client.__aenter__()
        return self

    async def __aexit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> None:
        await self._client.close()
        await self._credential.close()

    async def upload(
        self,
        container_name: str,
        blob_name: str,
        source_path: str | Path,
        *,
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[str]:
        blob_client = self._client.get_blob_client(container_name, blob_name)
        try:
            try:
                properties = await blob_client.get_blob_properties(timeout=timeout)
                write_conditions = {
                    "overwrite": True,
                    "etag": properties.etag,
                    "match_condition": MatchConditions.IfNotModified,
                }
            except ResourceNotFoundError:
                write_conditions = {"overwrite": False}

            path = Path(source_path)
            with path.open("rb") as data:
                response = await blob_client.upload_blob(
                    data,
                    length=path.stat().st_size,
                    metadata=dict(metadata) if metadata else None,
                    tags=dict(tags) if tags else None,
                    lease=lease_id,
                    timeout=timeout,
                    max_concurrency=self._settings.max_concurrency,
                    **write_conditions,
                )
            return OperationResult(
                True,
                f"Uploaded {path} to {container_name}/{blob_name}.",
                response["etag"],
            )
        except STORAGE_EXCEPTIONS as exc:
            return storage_failure(f"upload {container_name}/{blob_name}", exc)
        except OSError as exc:
            return OperationResult(False, f"Could not read {source_path}: {exc}.")

    async def download(
        self,
        container_name: str,
        blob_name: str,
        destination_path: str | Path,
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        destination = Path(destination_path)
        blob_client = self._client.get_blob_client(container_name, blob_name)
        try:
            destination.parent.mkdir(parents=True, exist_ok=True)
            downloader = await blob_client.download_blob(
                timeout=timeout,
                max_concurrency=self._settings.max_concurrency,
            )
            with destination.open("wb") as output:
                async for chunk in downloader.chunks():
                    output.write(chunk)
            return OperationResult(
                True,
                f"Downloaded {container_name}/{blob_name} to {destination}.",
                destination,
            )
        except STORAGE_EXCEPTIONS as exc:
            destination.unlink(missing_ok=True)
            return storage_failure(f"download {container_name}/{blob_name}", exc)
        except OSError as exc:
            destination.unlink(missing_ok=True)
            return OperationResult(False, f"Could not write {destination}: {exc}.")

    async def list_blobs(
        self,
        container_name: str,
        *,
        prefix: str | None = None,
        include_tags: bool = True,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobInfo]]:
        container_client = self._client.get_container_client(container_name)
        try:
            include = ["metadata", "tags"] if include_tags else ["metadata"]
            blobs = []
            async for item in container_client.list_blobs(
                name_starts_with=prefix,
                include=include,
                timeout=timeout,
            ):
                blobs.append(
                    BlobInfo(
                        name=item.name,
                        size=item.size or 0,
                        etag=item.etag,
                        metadata=dict(item.metadata or {}),
                        tags=dict(item.tags or {}),
                    )
                )
            return OperationResult(
                True, f"Listed {len(blobs)} blob(s) in {container_name}.", blobs
            )
        except STORAGE_EXCEPTIONS as exc:
            return storage_failure(f"list blobs in {container_name}", exc)

    async def delete(
        self,
        container_name: str,
        blob_name: str,
        *,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        blob_client = self._client.get_blob_client(container_name, blob_name)
        try:
            await blob_client.delete_blob(
                delete_snapshots="include", lease=lease_id, timeout=timeout
            )
            return OperationResult(True, f"Deleted {container_name}/{blob_name}.")
        except STORAGE_EXCEPTIONS as exc:
            return storage_failure(f"delete {container_name}/{blob_name}", exc)

    async def acquire_lease(
        self,
        container_name: str,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: int | None = None,
    ) -> OperationResult[LeaseHandle]:
        blob_client = self._client.get_blob_client(container_name, blob_name)
        lease = BlobLeaseClient(blob_client)
        try:
            await lease.acquire(lease_duration=duration, timeout=timeout)
            handle = LeaseHandle(container_name, blob_name, lease.id)
            return OperationResult(True, f"Acquired lease for {blob_name}.", handle)
        except STORAGE_EXCEPTIONS as exc:
            return storage_failure(f"acquire lease for {blob_name}", exc)

    async def release_lease(
        self,
        handle: LeaseHandle,
        *,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        blob_client = self._client.get_blob_client(
            handle.container_name, handle.blob_name
        )
        lease = BlobLeaseClient(blob_client, lease_id=handle.lease_id)
        try:
            await lease.release(timeout=timeout)
            return OperationResult(True, f"Released lease for {handle.blob_name}.")
        except STORAGE_EXCEPTIONS as exc:
            return storage_failure(f"release lease for {handle.blob_name}", exc)
