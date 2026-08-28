"""Asynchronous Azure Blob Storage management service."""

from __future__ import annotations

import asyncio
import os
from collections.abc import AsyncIterator, Mapping
from pathlib import Path
from types import TracebackType
from uuid import uuid4

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError
from azure.identity.aio import DefaultAzureCredential
from azure.storage.blob.aio import BlobLeaseClient, BlobServiceClient

from .config import StorageSettings, create_async_client
from .errors import HANDLED_AZURE_ERRORS, describe_storage_error, timeout_options
from .models import BlobInfo, OperationResult


async def _file_chunks(path: Path, chunk_size: int) -> AsyncIterator[bytes]:
    stream = await asyncio.to_thread(path.open, "rb")
    try:
        while chunk := await asyncio.to_thread(stream.read, chunk_size):
            yield chunk
    finally:
        await asyncio.to_thread(stream.close)


class AsyncBlobStorageManager:
    """Context-managed asynchronous wrapper around common blob operations."""

    def __init__(self, settings: StorageSettings) -> None:
        self.settings = settings
        self._credential: DefaultAzureCredential | None = None
        self._client: BlobServiceClient | None = None

    async def __aenter__(self) -> "AsyncBlobStorageManager":
        self._credential = DefaultAzureCredential()
        await self._credential.__aenter__()
        try:
            self._client = create_async_client(self.settings, self._credential)
            await self._client.__aenter__()
        except Exception:
            await self._credential.__aexit__(None, None, None)
            self._credential = None
            raise
        return self

    async def __aexit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> None:
        if self._client is not None:
            await self._client.__aexit__(exc_type, exc_value, traceback)
        if self._credential is not None:
            await self._credential.__aexit__(exc_type, exc_value, traceback)
        self._client = None
        self._credential = None

    def _service_client(self) -> BlobServiceClient:
        if self._client is None:
            raise RuntimeError("Use AsyncBlobStorageManager as a context manager.")
        return self._client

    async def upload(
        self,
        container: str,
        blob_name: str,
        source: str | Path,
        *,
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        lease: BlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        source_path = Path(source)
        try:
            file_size = (await asyncio.to_thread(source_path.stat)).st_size
            blob_client = self._service_client().get_blob_client(
                container=container, blob=blob_name
            )
            request_options = timeout_options(timeout)
            try:
                properties = await blob_client.get_blob_properties(**request_options)
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

            await blob_client.upload_blob(
                _file_chunks(source_path, self.settings.max_block_size),
                length=file_size,
                metadata=dict(metadata) if metadata else None,
                tags=dict(tags) if tags else None,
                lease=lease,
                max_concurrency=self.settings.max_concurrency,
                **conditions,
                **request_options,
            )
            return OperationResult.ok(
                f"Uploaded {source_path} to {container}/{blob_name}."
            )
        except OSError as error:
            return OperationResult.fail(f"Could not read {source_path}: {error}.")
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(describe_storage_error("upload blob", error))

    async def download(
        self,
        container: str,
        blob_name: str,
        destination: str | Path,
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        temporary_path = destination_path.with_name(
            f".{destination_path.name}.{uuid4().hex}.part"
        )
        stream = None
        try:
            await asyncio.to_thread(
                destination_path.parent.mkdir, parents=True, exist_ok=True
            )
            blob_client = self._service_client().get_blob_client(
                container=container, blob=blob_name
            )
            downloader = await blob_client.download_blob(
                max_concurrency=self.settings.max_concurrency,
                **timeout_options(timeout),
            )
            stream = await asyncio.to_thread(temporary_path.open, "wb")
            async for chunk in downloader.chunks():
                await asyncio.to_thread(stream.write, chunk)
            await asyncio.to_thread(stream.close)
            stream = None
            await asyncio.to_thread(os.replace, temporary_path, destination_path)
            return OperationResult.ok(
                f"Downloaded {container}/{blob_name} to {destination_path}.",
                destination_path,
            )
        except OSError as error:
            return OperationResult.fail(
                f"Could not write download to {destination_path}: {error}."
            )
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(describe_storage_error("download blob", error))
        finally:
            if stream is not None:
                await asyncio.to_thread(stream.close)
            await asyncio.to_thread(temporary_path.unlink, missing_ok=True)

    async def list_blobs(
        self,
        container: str,
        *,
        prefix: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobInfo]]:
        try:
            container_client = self._service_client().get_container_client(container)
            blobs = []
            async for item in container_client.list_blobs(
                name_starts_with=prefix,
                include=["metadata", "tags"],
                **timeout_options(timeout),
            ):
                blobs.append(
                    BlobInfo(
                        name=item.name,
                        size=item.size or 0,
                        last_modified=item.last_modified,
                        metadata=dict(item.metadata or {}),
                        tags=dict(item.tags or {}),
                    )
                )
            return OperationResult.ok(
                f"Listed {len(blobs)} blob(s) in {container}.", blobs
            )
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(describe_storage_error("list blobs", error))

    async def delete(
        self,
        container: str,
        blob_name: str,
        *,
        lease: BlobLeaseClient | str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        try:
            blob_client = self._service_client().get_blob_client(
                container=container, blob=blob_name
            )
            await blob_client.delete_blob(
                delete_snapshots="include",
                lease=lease,
                **timeout_options(timeout),
            )
            return OperationResult.ok(f"Deleted {container}/{blob_name}.")
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(describe_storage_error("delete blob", error))

    async def acquire_lease(
        self,
        container: str,
        blob_name: str,
        *,
        duration: int = 30,
        timeout: int | None = None,
    ) -> OperationResult[BlobLeaseClient]:
        try:
            blob_client = self._service_client().get_blob_client(
                container=container, blob=blob_name
            )
            lease = await blob_client.acquire_lease(
                lease_duration=duration, **timeout_options(timeout)
            )
            return OperationResult.ok(
                f"Acquired a lease on {container}/{blob_name}.", lease
            )
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(
                describe_storage_error("acquire blob lease", error)
            )

    async def release_lease(
        self, lease: BlobLeaseClient, *, timeout: int | None = None
    ) -> OperationResult[None]:
        try:
            await lease.release(**timeout_options(timeout))
            return OperationResult.ok("Released the blob lease.")
        except HANDLED_AZURE_ERRORS as error:
            return OperationResult.fail(
                describe_storage_error("release blob lease", error)
            )
