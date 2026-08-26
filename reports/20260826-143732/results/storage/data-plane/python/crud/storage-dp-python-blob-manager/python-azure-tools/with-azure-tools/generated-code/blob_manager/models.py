"""Return models shared by the sync and async services."""

from __future__ import annotations

from dataclasses import dataclass
from typing import Generic, TypeVar

T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    success: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class BlobInfo:
    name: str
    size: int
    etag: str | None
    metadata: dict[str, str]
    tags: dict[str, str]


@dataclass(frozen=True, slots=True)
class LeaseHandle:
    container_name: str
    blob_name: str
    lease_id: str
