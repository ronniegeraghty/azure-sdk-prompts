"""Synchronous Azure Blob Storage management service."""

from __future__ import annotations

import logging
import os
from pathlib import Path
from types import TracebackType
from typing import Any, Self

from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.storage.blob import BlobLeaseClient, BlobServiceClient

from .config import StorageSettings
from .errors import storage_error_message
from .models import BlobInfo, OperationResult

LOGGER = logging.getLogger(__name__)


class BlobStorageService:
    """Streamed blob operations with lease-protected updates."""

    def __init__(self, settings: StorageSettings, container_name: str) -> None:
        self.settings = settings
        self.container_name = container_name
        self._credential: Any = None
        self._client: BlobServiceClient | None = None

    def __enter__(self) -> Self:
        self.settings.configure_logging()
        self._credential = self.settings.create_credential()
        self._credential.__enter__()
        try:
            self._client = self.settings.create_client(self._credential)
            self._client.__enter__()
        except BaseException:
            self._credential.__exit__(None, None, None)
            self._credential = None
            raise
        return self

    def __exit__(
        self,
        exc_type: type[BaseException] | None,
        exc_value: BaseException | None,
        traceback: TracebackType | None,
    ) -> None:
        if self._client is not None:
            self._client.__exit__(exc_type, exc_value, traceback)
        if self._credential is not None:
            self._credential.__exit__(exc_type, exc_value, traceback)
        self._client = None
        self._credential = None

    @property
    def client(self) -> BlobServiceClient:
        if self._client is None:
            raise RuntimeError("Use BlobStorageService as a context manager.")
        return self._client

    @staticmethod
    def _timeout_kwargs(timeout: int | None) -> dict[str, int]:
        return {} if timeout is None else {"timeout": timeout}

    def upload(
        self,
        local_path: str | Path,
        blob_name: str,
        *,
        metadata: dict[str, str] | None = None,
        tags: dict[str, str] | None = None,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        """Upload from a file stream, leasing an existing blob before replacing it."""

        path = Path(local_path)
        if not path.is_file():
            return OperationResult(False, f"Upload failed: file not found: {path}.")

        blob_client = self.client.get_blob_client(self.container_name, blob_name)
        owned_lease: BlobLeaseClient | None = None
        try:
            exists = True
            try:
                blob_client.get_blob_properties(**self._timeout_kwargs(timeout))
            except ResourceNotFoundError:
                exists = False

            active_lease: BlobLeaseClient | str | None = lease_id
            if exists and lease_id is None:
                owned_lease = blob_client.acquire_lease(
                    lease_duration=-1, **self._timeout_kwargs(timeout)
                )
                active_lease = owned_lease

            with path.open("rb") as data:
                blob_client.upload_blob(
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
                    owned_lease.release(**self._timeout_kwargs(timeout))
                except AzureError as error:
                    LOGGER.warning(
                        "Could not release the upload lease for %s: %s",
                        blob_name,
                        storage_error_message("Lease release", error),
                    )

    def download(
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
            stream = blob_client.download_blob(
                max_concurrency=self.settings.max_concurrency,
                **self._timeout_kwargs(timeout),
            )
            with temporary_path.open("wb") as output:
                stream.readinto(output)
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

    def list_blobs(
        self,
        *,
        prefix: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[list[BlobInfo]]:
        """List blobs and include metadata and blob index tags."""

        try:
            container = self.client.get_container_client(self.container_name)
            blobs = [
                BlobInfo(
                    name=blob.name,
                    size=blob.size or 0,
                    last_modified=blob.last_modified,
                    metadata=dict(blob.metadata or {}),
                    tags=dict(blob.tags or {}),
                )
                for blob in container.list_blobs(
                    name_starts_with=prefix,
                    include=["metadata", "tags"],
                    **self._timeout_kwargs(timeout),
                )
            ]
            return OperationResult(
                True,
                f"Listed {len(blobs)} blob(s) in {self.container_name}.",
                blobs,
            )
        except AzureError as error:
            return OperationResult(False, storage_error_message("List", error))

    def delete(
        self,
        blob_name: str,
        *,
        lease_id: str | None = None,
        timeout: int | None = None,
    ) -> OperationResult[None]:
        try:
            self.client.get_blob_client(self.container_name, blob_name).delete_blob(
                delete_snapshots="include",
                lease=lease_id,
                **self._timeout_kwargs(timeout),
            )
            return OperationResult(
                True, f"Deleted {self.container_name}/{blob_name}."
            )
        except AzureError as error:
            return OperationResult(False, storage_error_message("Delete", error))

    def acquire_lease(
        self, blob_name: str, *, timeout: int | None = None
    ) -> OperationResult[str]:
        try:
            lease = self.client.get_blob_client(
                self.container_name, blob_name
            ).acquire_lease(lease_duration=-1, **self._timeout_kwargs(timeout))
            return OperationResult(
                True, f"Acquired a lease for {blob_name}.", lease.id
            )
        except AzureError as error:
            return OperationResult(
                False, storage_error_message("Lease acquisition", error)
            )

    def release_lease(
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
            BlobLeaseClient(blob_client, lease_id=lease_id).release(
                **self._timeout_kwargs(timeout)
            )
            return OperationResult(True, f"Released the lease for {blob_name}.")
        except AzureError as error:
            return OperationResult(
                False, storage_error_message("Lease release", error)
            )
