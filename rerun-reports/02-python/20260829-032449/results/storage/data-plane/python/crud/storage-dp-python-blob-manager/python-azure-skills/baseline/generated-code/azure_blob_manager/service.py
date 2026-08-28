"""Synchronous Azure Blob Storage operations."""

from __future__ import annotations

import os
from pathlib import Path
from typing import Any, Mapping

from azure.core import MatchConditions
from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceModifiedError,
    ResourceNotFoundError,
)
from azure.storage.blob import BlobLeaseClient, BlobServiceClient

from .models import BlobInfo, LeaseHandle, OperationResult, UploadInfo


def _request_options(timeout: float | None) -> dict[str, float]:
    return {"timeout": timeout} if timeout is not None else {}


def _error_message(operation: str, blob_name: str | None, exc: Exception) -> str:
    target = f" for blob {blob_name!r}" if blob_name else ""
    if isinstance(exc, ResourceNotFoundError):
        detail = "the container or blob was not found"
    elif isinstance(exc, ClientAuthenticationError):
        detail = "authentication failed; check the managed identity and role assignment"
    elif isinstance(exc, ResourceExistsError):
        detail = "the blob changed or was created by another writer"
    elif isinstance(exc, ResourceModifiedError):
        detail = "the blob was modified by another writer"
    elif isinstance(exc, HttpResponseError):
        code = getattr(exc, "error_code", None)
        if code in {"LeaseAlreadyPresent", "LeaseIsBreakingAndCannotBeAcquired"}:
            detail = "a lease is already held by another client"
        elif code in {"LeaseIdMissing", "LeaseIdMismatchWithBlobOperation"}:
            detail = "a valid lease ID is required"
        elif exc.status_code == 403:
            detail = "permission denied; check the managed identity role assignment"
        else:
            detail = f"Azure Storage returned {code or exc.status_code or 'an error'}"
    else:
        detail = str(exc)
    return f"{operation}{target} failed: {detail}"


class BlobStorageService:
    """Memory-efficient, optimistic-concurrency-safe blob operations."""

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
        metadata: Mapping[str, str] | None = None,
        tags: Mapping[str, str] | None = None,
        lease_id: str | None = None,
        timeout: float | None = None,
    ) -> OperationResult[UploadInfo]:
        """Stream a file and update only if the observed blob version is unchanged."""
        source = Path(source_path)
        blob = self._container.get_blob_client(blob_name)
        options: dict[str, Any] = _request_options(timeout)
        if lease_id:
            options["lease"] = lease_id

        try:
            try:
                current = blob.get_blob_properties(**_request_options(timeout))
                options["etag"] = current.etag
                options["match_condition"] = MatchConditions.IfNotModified
                overwrite = True
            except ResourceNotFoundError:
                overwrite = False

            with source.open("rb") as data:
                response = blob.upload_blob(
                    data,
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

    def download(
        self,
        blob_name: str,
        destination_path: str | Path,
        *,
        timeout: float | None = None,
    ) -> OperationResult[Path]:
        """Download a blob incrementally and atomically replace the destination."""
        destination = Path(destination_path)
        temporary = destination.with_name(f"{destination.name}.part")
        try:
            destination.parent.mkdir(parents=True, exist_ok=True)
            stream = self._container.download_blob(
                blob_name,
                max_concurrency=self._max_concurrency,
                **_request_options(timeout),
            )
            with temporary.open("wb") as output:
                for chunk in stream.chunks():
                    output.write(chunk)
            os.replace(temporary, destination)
            return OperationResult(
                True, f"Downloaded {blob_name!r} to {destination}", destination
            )
        except (OSError, HttpResponseError) as exc:
            temporary.unlink(missing_ok=True)
            return OperationResult(False, _error_message("Download", blob_name, exc))

    def list_blobs(
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
                for item in self._container.list_blobs(
                    include=["metadata", "tags"], **_request_options(timeout)
                )
            ]
            return OperationResult(True, f"Listed {len(blobs)} blob(s)", blobs)
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("List blobs", None, exc))

    def delete(
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
            self._container.delete_blob(blob_name, **options)
            return OperationResult(True, f"Deleted {blob_name!r}")
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("Delete", blob_name, exc))

    def acquire_lease(
        self,
        blob_name: str,
        *,
        lease_duration: int = 60,
        timeout: float | None = None,
    ) -> OperationResult[LeaseHandle]:
        try:
            lease = self._container.get_blob_client(blob_name).acquire_lease(
                lease_duration=lease_duration, **_request_options(timeout)
            )
            handle = LeaseHandle(blob_name, lease.id)
            return OperationResult(True, f"Acquired lease for {blob_name!r}", handle)
        except HttpResponseError as exc:
            return OperationResult(False, _error_message("Acquire lease", blob_name, exc))

    def release_lease(
        self, handle: LeaseHandle, *, timeout: float | None = None
    ) -> OperationResult[None]:
        try:
            blob = self._container.get_blob_client(handle.blob_name)
            BlobLeaseClient(blob, lease_id=handle.lease_id).release(
                **_request_options(timeout)
            )
            return OperationResult(True, f"Released lease for {handle.blob_name!r}")
        except HttpResponseError as exc:
            return OperationResult(
                False, _error_message("Release lease", handle.blob_name, exc)
            )
