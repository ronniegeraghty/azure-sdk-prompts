"""High-level synchronous and asynchronous Azure Blob Storage operations."""

from __future__ import annotations

import os
import tempfile
from dataclasses import dataclass
from pathlib import Path
from typing import Generic, Mapping, TypeVar

from azure.core.exceptions import (
    AzureError,
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
)
from azure.storage.blob import BlobLeaseClient, BlobServiceClient
from azure.storage.blob.aio import BlobLeaseClient as AsyncBlobLeaseClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    success: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class BlobSummary:
    name: str
    size: int
    metadata: Mapping[str, str]
    tags: Mapping[str, str]


def _success(message: str, value: T | None = None) -> OperationResult[T]:
    return OperationResult(success=True, message=message, value=value)


def _failure(message: str) -> OperationResult[T]:
    return OperationResult(success=False, message=message)


def _request_options(timeout: float | None) -> dict[str, float]:
    if timeout is None:
        return {}
    if timeout <= 0:
        raise ValueError("timeout must be greater than zero.")
    return {"timeout": timeout}


def _storage_error_message(action: str, error: AzureError) -> str:
    if isinstance(error, ResourceNotFoundError):
        return f"{action} failed: the container or blob was not found."
    if isinstance(error, ResourceExistsError):
        return f"{action} failed: the blob already exists and overwrite was not requested."
    if isinstance(error, ClientAuthenticationError):
        return f"{action} failed: Azure could not authenticate the configured identity."

    if isinstance(error, HttpResponseError):
        if error.status_code == 403:
            return f"{action} failed: the identity does not have permission for this operation."

        error_code = getattr(error, "error_code", None)
        lease_messages = {
            "LeaseAlreadyPresent": "another client already holds a lease",
            "LeaseIdMissing": "the blob has an active lease but no lease ID was supplied",
            "LeaseIdMismatchWithBlobOperation": "the supplied lease does not own the blob",
            "LeaseLost": "the lease expired or was lost before the operation completed",
        }
        if error_code in lease_messages:
            return f"{action} failed: {lease_messages[error_code]}."
        if error_code in {"ConditionNotMet", "TargetConditionNotMet"}:
            return f"{action} failed: the blob changed during the operation."

    detail = str(error).strip().splitlines()[0] if str(error).strip() else type(error).__name__
    return f"{action} failed: {detail}"


def _blob_summary(blob: object) -> BlobSummary:
    return BlobSummary(
        name=str(getattr(blob, "name", "")),
        size=int(getattr(blob, "size", 0) or 0),
        metadata=dict(getattr(blob, "metadata", None) or {}),
        tags=dict(getattr(blob, "tags", None) or {}),
    )


class BlobStorageService:
    def __init__(self, client: BlobServiceClient, max_concurrency: int = 4) -> None:
        self._client = client
        self._max_concurrency = max_concurrency

    def upload_file(
        self,
        container: str,
        blob_name: str,
        source: str | Path,
        *,
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        overwrite: bool = False,
        lease: BlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[None]:
        if overwrite and lease is None:
            return _failure(
                "Upload refused: overwriting requires a lease so concurrent writers "
                "cannot silently replace each other's changes."
            )

        source_path = Path(source)
        if not source_path.is_file():
            return _failure(f"Upload failed: source file does not exist: {source_path}")

        try:
            options = _request_options(timeout)
            blob_client = self._client.get_blob_client(container, blob_name)
            with source_path.open("rb") as stream:
                blob_client.upload_blob(
                    stream,
                    length=source_path.stat().st_size,
                    metadata=dict(metadata) if metadata else None,
                    tags=dict(tags) if tags else None,
                    overwrite=overwrite,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    **options,
                )
            return _success(
                f"Uploaded {source_path} to {container}/{blob_name} using a streaming transfer."
            )
        except (OSError, ValueError) as error:
            return _failure(f"Upload failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Upload", error))

    def download_file(
        self,
        container: str,
        blob_name: str,
        destination: str | Path,
        *,
        overwrite_local: bool = False,
        timeout: float | None = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        if destination_path.exists() and not overwrite_local:
            return _failure(f"Download refused: destination already exists: {destination_path}")

        temporary_path: Path | None = None
        try:
            options = _request_options(timeout)
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            with tempfile.NamedTemporaryFile(
                mode="w+b",
                prefix=f".{destination_path.name}.",
                suffix=".part",
                dir=destination_path.parent,
                delete=False,
            ) as stream:
                temporary_path = Path(stream.name)
                downloader = self._client.get_blob_client(container, blob_name).download_blob(
                    max_concurrency=self._max_concurrency,
                    **options,
                )
                downloader.readinto(stream)
            os.replace(temporary_path, destination_path)
            return _success(
                f"Downloaded {container}/{blob_name} to {destination_path}.",
                destination_path,
            )
        except (OSError, ValueError) as error:
            return _failure(f"Download failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Download", error))
        finally:
            if temporary_path is not None and temporary_path.exists():
                temporary_path.unlink(missing_ok=True)

    def list_blobs(
        self,
        container: str,
        *,
        timeout: float | None = None,
    ) -> OperationResult[list[BlobSummary]]:
        try:
            options = _request_options(timeout)
            blobs = [
                _blob_summary(blob)
                for blob in self._client.get_container_client(container).list_blobs(
                    include=["metadata", "tags"],
                    **options,
                )
            ]
            return _success(f"Found {len(blobs)} blob(s) in {container}.", blobs)
        except ValueError as error:
            return _failure(f"List failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("List", error))

    def delete_blob(
        self,
        container: str,
        blob_name: str,
        *,
        lease: BlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[None]:
        try:
            options = _request_options(timeout)
            self._client.get_blob_client(container, blob_name).delete_blob(
                lease=lease,
                **options,
            )
            return _success(f"Deleted {container}/{blob_name}.")
        except ValueError as error:
            return _failure(f"Delete failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Delete", error))

    def acquire_lease(
        self,
        container: str,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: float | None = None,
    ) -> OperationResult[BlobLeaseClient]:
        try:
            options = _request_options(timeout)
            lease = self._client.get_blob_client(container, blob_name).acquire_lease(
                lease_duration=duration,
                **options,
            )
            return _success(f"Acquired a lease for {container}/{blob_name}.", lease)
        except ValueError as error:
            return _failure(f"Acquire lease failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Acquire lease", error))


class AsyncBlobStorageService:
    def __init__(self, client: AsyncBlobServiceClient, max_concurrency: int = 4) -> None:
        self._client = client
        self._max_concurrency = max_concurrency

    async def upload_file(
        self,
        container: str,
        blob_name: str,
        source: str | Path,
        *,
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        overwrite: bool = False,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[None]:
        if overwrite and lease is None:
            return _failure(
                "Upload refused: overwriting requires a lease so concurrent writers "
                "cannot silently replace each other's changes."
            )

        source_path = Path(source)
        if not source_path.is_file():
            return _failure(f"Upload failed: source file does not exist: {source_path}")

        try:
            options = _request_options(timeout)
            blob_client = self._client.get_blob_client(container, blob_name)
            with source_path.open("rb") as stream:
                await blob_client.upload_blob(
                    stream,
                    length=source_path.stat().st_size,
                    metadata=dict(metadata) if metadata else None,
                    tags=dict(tags) if tags else None,
                    overwrite=overwrite,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    **options,
                )
            return _success(
                f"Uploaded {source_path} to {container}/{blob_name} using a streaming transfer."
            )
        except (OSError, ValueError) as error:
            return _failure(f"Upload failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Upload", error))

    async def download_file(
        self,
        container: str,
        blob_name: str,
        destination: str | Path,
        *,
        overwrite_local: bool = False,
        timeout: float | None = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        if destination_path.exists() and not overwrite_local:
            return _failure(f"Download refused: destination already exists: {destination_path}")

        temporary_path: Path | None = None
        try:
            options = _request_options(timeout)
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            with tempfile.NamedTemporaryFile(
                mode="w+b",
                prefix=f".{destination_path.name}.",
                suffix=".part",
                dir=destination_path.parent,
                delete=False,
            ) as stream:
                temporary_path = Path(stream.name)
                downloader = await self._client.get_blob_client(
                    container, blob_name
                ).download_blob(
                    max_concurrency=self._max_concurrency,
                    **options,
                )
                await downloader.readinto(stream)
            os.replace(temporary_path, destination_path)
            return _success(
                f"Downloaded {container}/{blob_name} to {destination_path}.",
                destination_path,
            )
        except (OSError, ValueError) as error:
            return _failure(f"Download failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Download", error))
        finally:
            if temporary_path is not None and temporary_path.exists():
                temporary_path.unlink(missing_ok=True)

    async def list_blobs(
        self,
        container: str,
        *,
        timeout: float | None = None,
    ) -> OperationResult[list[BlobSummary]]:
        try:
            options = _request_options(timeout)
            blobs = [
                _blob_summary(blob)
                async for blob in self._client.get_container_client(container).list_blobs(
                    include=["metadata", "tags"],
                    **options,
                )
            ]
            return _success(f"Found {len(blobs)} blob(s) in {container}.", blobs)
        except ValueError as error:
            return _failure(f"List failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("List", error))

    async def delete_blob(
        self,
        container: str,
        blob_name: str,
        *,
        lease: AsyncBlobLeaseClient | str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[None]:
        try:
            options = _request_options(timeout)
            await self._client.get_blob_client(container, blob_name).delete_blob(
                lease=lease,
                **options,
            )
            return _success(f"Deleted {container}/{blob_name}.")
        except ValueError as error:
            return _failure(f"Delete failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Delete", error))

    async def acquire_lease(
        self,
        container: str,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: float | None = None,
    ) -> OperationResult[AsyncBlobLeaseClient]:
        try:
            options = _request_options(timeout)
            lease = await self._client.get_blob_client(container, blob_name).acquire_lease(
                lease_duration=duration,
                **options,
            )
            return _success(f"Acquired a lease for {container}/{blob_name}.", lease)
        except ValueError as error:
            return _failure(f"Acquire lease failed: {error}")
        except AzureError as error:
            return _failure(_storage_error_message("Acquire lease", error))
