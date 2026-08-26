"""Reusable synchronous and asynchronous Azure Blob Storage services."""

from __future__ import annotations

import os
import uuid
from dataclasses import dataclass
from pathlib import Path
from typing import Any, Generic, Optional, TypeVar

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ServiceRequestError,
)
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient

T = TypeVar("T")


@dataclass(frozen=True)
class OperationResult(Generic[T]):
    success: bool
    message: str
    value: Optional[T] = None


def _storage_error(action: str, exc: Exception) -> str:
    if isinstance(exc, ResourceNotFoundError):
        detail = "the container or blob was not found"
    elif isinstance(exc, ResourceExistsError):
        error_code = getattr(exc, "error_code", "")
        if error_code and "Lease" in error_code:
            detail = "the blob is already leased by another client"
        else:
            detail = "the blob already exists or is currently leased"
    elif isinstance(exc, ClientAuthenticationError):
        detail = "authentication failed; verify the managed identity and RBAC role"
    elif isinstance(exc, ServiceRequestError):
        detail = "the storage endpoint could not be reached"
    elif isinstance(exc, HttpResponseError):
        status = getattr(exc, "status_code", None)
        error_code = getattr(exc, "error_code", None)
        if status == 403:
            detail = "permission was denied; verify the identity's data-plane RBAC role"
        elif error_code and "Lease" in error_code:
            detail = f"the lease condition failed ({error_code})"
        else:
            suffix = f" ({error_code})" if error_code else ""
            detail = f"Azure Storage returned HTTP {status or 'error'}{suffix}"
    elif isinstance(exc, OSError):
        detail = f"local file error: {exc}"
    else:
        detail = str(exc)
    return f"Could not {action}: {detail}."


def _timeout_kwargs(timeout: Optional[int]) -> dict[str, int]:
    if timeout is None:
        return {}
    if timeout <= 0:
        raise ValueError("timeout must be greater than zero")
    return {"timeout": timeout}


def _temporary_download_path(destination: Path) -> Path:
    return destination.with_name(f".{destination.name}.{uuid.uuid4().hex}.part")


class BlobStorageService:
    def __init__(
        self,
        client: BlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
    ) -> None:
        self._client = client
        self._container_name = container_name
        self._max_concurrency = max_concurrency

    def upload_file(
        self,
        source: str | os.PathLike[str],
        blob_name: str,
        *,
        metadata: Optional[dict[str, str]] = None,
        tags: Optional[dict[str, str]] = None,
        overwrite: bool = False,
        lease: Any = None,
        timeout: Optional[int] = None,
    ) -> OperationResult[Any]:
        if overwrite and lease is None:
            return OperationResult(
                False,
                "Could not upload blob: overwrites require an active lease to "
                "prevent concurrent writers.",
            )

        try:
            timeout_kwargs = _timeout_kwargs(timeout)
            blob_client = self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            )
            with Path(source).open("rb") as stream:
                response = blob_client.upload_blob(
                    stream,
                    overwrite=overwrite,
                    metadata=metadata,
                    tags=tags,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    validate_content=False,
                    **timeout_kwargs,
                )
            return OperationResult(True, f"Uploaded '{blob_name}'.", response)
        except (HttpResponseError, ServiceRequestError, OSError, ValueError) as exc:
            return OperationResult(False, _storage_error("upload blob", exc))

    def download_file(
        self,
        blob_name: str,
        destination: str | os.PathLike[str],
        *,
        timeout: Optional[int] = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        temporary_path = _temporary_download_path(destination_path)
        try:
            timeout_kwargs = _timeout_kwargs(timeout)
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            downloader = self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).download_blob(
                max_concurrency=self._max_concurrency,
                **timeout_kwargs,
            )
            with temporary_path.open("wb") as stream:
                downloader.readinto(stream)
            temporary_path.replace(destination_path)
            return OperationResult(
                True, f"Downloaded '{blob_name}' to '{destination_path}'.", destination_path
            )
        except (HttpResponseError, ServiceRequestError, OSError, ValueError) as exc:
            temporary_path.unlink(missing_ok=True)
            return OperationResult(False, _storage_error("download blob", exc))

    def list_blobs(
        self, *, timeout: Optional[int] = None
    ) -> OperationResult[list[dict[str, Any]]]:
        try:
            timeout_kwargs = _timeout_kwargs(timeout)
            blobs = [
                {
                    "name": blob.name,
                    "size": blob.size,
                    "metadata": blob.metadata,
                    "tags": blob.tags,
                }
                for blob in self._client.get_container_client(
                    self._container_name
                ).list_blobs(include=["metadata", "tags"], **timeout_kwargs)
            ]
            return OperationResult(
                True, f"Found {len(blobs)} blob(s) in '{self._container_name}'.", blobs
            )
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("list blobs", exc))

    def delete_blob(
        self,
        blob_name: str,
        *,
        lease: Any = None,
        timeout: Optional[int] = None,
    ) -> OperationResult[None]:
        try:
            self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).delete_blob(lease=lease, **_timeout_kwargs(timeout))
            return OperationResult(True, f"Deleted '{blob_name}'.")
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("delete blob", exc))

    def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: Optional[int] = None,
    ) -> OperationResult[Any]:
        try:
            lease = self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).acquire_lease(
                lease_duration=duration,
                **_timeout_kwargs(timeout),
            )
            return OperationResult(True, f"Acquired a lease on '{blob_name}'.", lease)
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("acquire blob lease", exc))

    def release_lease(
        self, lease: Any, *, timeout: Optional[int] = None
    ) -> OperationResult[None]:
        try:
            lease.release(**_timeout_kwargs(timeout))
            return OperationResult(True, "Released the blob lease.")
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("release blob lease", exc))


class AsyncBlobStorageService:
    def __init__(
        self,
        client: AsyncBlobServiceClient,
        container_name: str,
        *,
        max_concurrency: int = 4,
    ) -> None:
        self._client = client
        self._container_name = container_name
        self._max_concurrency = max_concurrency

    async def upload_file(
        self,
        source: str | os.PathLike[str],
        blob_name: str,
        *,
        metadata: Optional[dict[str, str]] = None,
        tags: Optional[dict[str, str]] = None,
        overwrite: bool = False,
        lease: Any = None,
        timeout: Optional[int] = None,
    ) -> OperationResult[Any]:
        if overwrite and lease is None:
            return OperationResult(
                False,
                "Could not upload blob: overwrites require an active lease to "
                "prevent concurrent writers.",
            )

        try:
            blob_client = self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            )
            with Path(source).open("rb") as stream:
                response = await blob_client.upload_blob(
                    stream,
                    overwrite=overwrite,
                    metadata=metadata,
                    tags=tags,
                    lease=lease,
                    max_concurrency=self._max_concurrency,
                    validate_content=False,
                    **_timeout_kwargs(timeout),
                )
            return OperationResult(True, f"Uploaded '{blob_name}'.", response)
        except (HttpResponseError, ServiceRequestError, OSError, ValueError) as exc:
            return OperationResult(False, _storage_error("upload blob", exc))

    async def download_file(
        self,
        blob_name: str,
        destination: str | os.PathLike[str],
        *,
        timeout: Optional[int] = None,
    ) -> OperationResult[Path]:
        destination_path = Path(destination)
        temporary_path = _temporary_download_path(destination_path)
        try:
            destination_path.parent.mkdir(parents=True, exist_ok=True)
            downloader = await self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).download_blob(
                max_concurrency=self._max_concurrency,
                **_timeout_kwargs(timeout),
            )
            with temporary_path.open("wb") as stream:
                async for chunk in downloader.chunks():
                    stream.write(chunk)
            temporary_path.replace(destination_path)
            return OperationResult(
                True, f"Downloaded '{blob_name}' to '{destination_path}'.", destination_path
            )
        except (HttpResponseError, ServiceRequestError, OSError, ValueError) as exc:
            temporary_path.unlink(missing_ok=True)
            return OperationResult(False, _storage_error("download blob", exc))

    async def list_blobs(
        self, *, timeout: Optional[int] = None
    ) -> OperationResult[list[dict[str, Any]]]:
        try:
            blobs = []
            iterator = self._client.get_container_client(
                self._container_name
            ).list_blobs(
                include=["metadata", "tags"],
                **_timeout_kwargs(timeout),
            )
            async for blob in iterator:
                blobs.append(
                    {
                        "name": blob.name,
                        "size": blob.size,
                        "metadata": blob.metadata,
                        "tags": blob.tags,
                    }
                )
            return OperationResult(
                True, f"Found {len(blobs)} blob(s) in '{self._container_name}'.", blobs
            )
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("list blobs", exc))

    async def delete_blob(
        self,
        blob_name: str,
        *,
        lease: Any = None,
        timeout: Optional[int] = None,
    ) -> OperationResult[None]:
        try:
            await self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).delete_blob(lease=lease, **_timeout_kwargs(timeout))
            return OperationResult(True, f"Deleted '{blob_name}'.")
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("delete blob", exc))

    async def acquire_lease(
        self,
        blob_name: str,
        *,
        duration: int = 60,
        timeout: Optional[int] = None,
    ) -> OperationResult[Any]:
        try:
            lease = await self._client.get_blob_client(
                container=self._container_name, blob=blob_name
            ).acquire_lease(
                lease_duration=duration,
                **_timeout_kwargs(timeout),
            )
            return OperationResult(True, f"Acquired a lease on '{blob_name}'.", lease)
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("acquire blob lease", exc))

    async def release_lease(
        self, lease: Any, *, timeout: Optional[int] = None
    ) -> OperationResult[None]:
        try:
            await lease.release(**_timeout_kwargs(timeout))
            return OperationResult(True, "Released the blob lease.")
        except (HttpResponseError, ServiceRequestError, ValueError) as exc:
            return OperationResult(False, _storage_error("release blob lease", exc))
