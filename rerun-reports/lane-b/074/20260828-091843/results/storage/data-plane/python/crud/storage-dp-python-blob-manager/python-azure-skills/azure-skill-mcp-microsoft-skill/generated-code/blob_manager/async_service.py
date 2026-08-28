"""Asynchronous Azure Blob Storage management service."""

from __future__ import annotations

import logging
import os
from pathlib import Path
from types import TracebackType
from typing import Any, Self

from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.storage.blob.aio import BlobLeaseClient, BlobServiceClient

from .config import StorageSettings
from .errors import storage_error_message
from .models import BlobInfo, OperationResult

LOGGER = logging.getLogger(__name__)


class AsyncBlobStorageService:
    """Async streamed blob operations with lease-protected updates."""

    def __init__(self, settings: StorageSettings, container_name: str) -> None:
        self.settings = settings
        self.container_name = container_name
        self._credential: Any = None
        self._client: BlobServiceClient | None = None

    async def __aenter__(self) -> Self:
        self.settings.configure_logging()
        self._credential = self.settings.create_async_credential()
        await self._credential.__aenter__()
        try:
            self._client = self.settings.create_async_client(self._credential)
            await self._client.__aenter__()
        except BaseException:
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

    @property
    def client(self) -> BlobServiceClient:
        if self._client is None:
            raise RuntimeError("Use AsyncBlobStorageService as an async context manager.")
        return self._client

    @staticmethod
    def _timeout_kwargs(timeout: int | None) -> dict[str, int]:
        return {} if timeout is None else {"timeout": timeout}

    async def upload(
        self,
        local_path: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        """Upload a file in bounded-size SDK blocks without reading it all at once."""

        path = Path(local_path)
        if not path.is_file():
            return OperationResult(False, f"Upload failed: file not found: {path}.")

        blob_client = self.client.get_blob_client(self.container_name, blob_name)
        owned_lease: BlobLeaseClient | None = None
        try:
            exists = True
            try:
                await blob_client.get_blob_properties(
                    **self._timeout_kwargs(timeout)
                )
            except ResourceNotFoundError:
                exists = False

            active_lease: BlobLeaseClient | str | None = lease_id
            if exists and lease_id is None:
                owned_lease = await blob_client.acquire_lease(
                    lease_duration=-1, **self._timeout_kwargs(timeout)
                )
                active_lease = owned_lease

            with path.open("rb") as data:
                await blob_client.upload_blob(
                    data,
                    length=os.path.getsize(path),
                    overwrite=exists,
                    metadata=metadata,
                    tags=tags,
                    lease=active_lease,
                    max_concurrency=self.settings.max_concurrency,
                    **self._timeout_kwargs(timeout),
                )
            return OperationResult(
                True,
                f"Uploaded {path} to {self.container_name}/{blob_name}.",
            )
        except AzureError as error:
            return OperationResult(False, storage_error_message("Upload", error))
        finally:
            if owned_lease is not None:
                try:
                    await owned_lease.release(**self._timeout_kwargs(timeout))
                except AzureError as error:
                    LOGGER.warning(
                        "Could not release the upload lease for %s: %s",
                        blob_name,
                        storage_error_message("Lease release", error),
                    )

    async def download(
        self,
        blob_name: str,
        destination: str | Path,
        *,
        timeout: int | None = None,
    ) -> OperationResult[Path]:
        """Download directly into a file without buffering the full blob."""

        destination_path = Path(destination)
        temporary_path = destination_path.with_name(destination_path.name + ".part")
        try:
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            blob_client = self.client.get_blob_client(
                self.container_name, blob_name
            )
            stream = await blob_client.download_blob(
                max_concurrency=self.settings.max_concurrency,
                **self._timeout_kwargs(timeout),
            )
            with temporary_path.open("wb") as output:
                await stream.readinto(output)
            temporary_path.replace(destination_path)
            return OperationResult(
                True,
                f"Downloaded {self.container_name}/{blob_name} to {destination_path}.",
                destination_path,
            )
        except (AzureError, OSError) as error:
            temporary_path.unlink(missing_ok=True)
            if isinstance(error, AzureError):
                message = storage_error_message("Download", error)
            else:
                message = f"Download failed: could not write {destination_path}: {error}."
            return OperationResult(False, message)

    async def list_blobs(
        self,
        *,
        prefix: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobInfo]]:
        """List blobs and include metadata and blob index tags."""

        try:
            container = self.client.get_container_client(self.container_name)
            blobs: list[BlobInfo] = []
            async for blob in container.list_blobs(
                name_starts_with=prefix,
                include=["metadata", "tags"],
                **self._timeout_kwargs(timeout),
            ):
                blobs.append(
                    BlobInfo(
                        name=blob.name,
                        size=blob.size or 0,
                        last_modified=blob.last_modified,
                        metadata=dict(blob.metadata or {}),
                        tags=dict(blob.tags or {}),
                    )
                )
            return OperationResult(
                True,
                f"Listed {len(blobs)} blob(s) in {self.container_name}.",
                blobs,
            )
        except AzureError as error:
            return OperationResult(False, storage_error_message("List", error))

    async def delete(
        self,
        blob_name: str,
        *,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        try:
            blob_client = self.client.get_blob_client(
                self.container_name, blob_name
            )
            await blob_client.delete_blob(
                delete_snapshots="include",
                lease=lease_id,
                **self._timeout_kwargs(timeout),
            )
            return OperationResult(
                True, f"Deleted {self.container_name}/{blob_name}."
            )
        except AzureError as error:
            return OperationResult(False, storage_error_message("Delete", error))

    async def acquire_lease(
        self, blob_name: str, *, timeout: int | None = None
    ) -> OperationResult[str]:
        try:
            blob_client = self.client.get_blob_client(
                self.container_name, blob_name
            )
            lease = await blob_client.acquire_lease(
                lease_duration=-1, **self._timeout_kwargs(timeout)
            )
            return OperationResult(
                True, f"Acquired a lease for {blob_name}.", lease.id
            )
        except AzureError as error:
            return OperationResult(
                False, storage_error_message("Lease acquisition", error)
            )

    async def release_lease(
        self,
        blob_name: str,
        lease_id: str,
        *,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        try:
            blob_client = self.client.get_blob_client(
                self.container_name, blob_name
            )
            await BlobLeaseClient(blob_client, lease_id=lease_id).release(
                **self._timeout_kwargs(timeout)
            )
            return OperationResult(True, f"Released the lease for {blob_name}.")
        except AzureError as error:
            return OperationResult(
                False, storage_error_message("Lease release", error)
            )
