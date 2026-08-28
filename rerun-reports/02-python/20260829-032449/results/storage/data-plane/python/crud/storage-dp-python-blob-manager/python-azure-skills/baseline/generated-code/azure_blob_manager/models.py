"""Value objects returned by blob storage operations."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Generic, Mapping, TypeVar

T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    success: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class UploadInfo:
    name: str
    etag: str
    last_modified: datetime | None


@dataclass(frozen=True, slots=True)
class BlobInfo:
    name: str
    size: int
    etag: str | None
    last_modified: datetime | None
    metadata: Mapping[str, str]
    tags: Mapping[str, str]


@dataclass(frozen=True, slots=True)
class LeaseHandle:
    blob_name: str
    lease_id: str
