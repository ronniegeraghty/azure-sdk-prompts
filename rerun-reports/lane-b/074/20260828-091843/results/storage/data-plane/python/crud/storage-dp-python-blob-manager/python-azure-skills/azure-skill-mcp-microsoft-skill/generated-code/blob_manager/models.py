"""Result types returned by blob management operations."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Generic, TypeVar

T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    """A non-throwing result with a caller-friendly status message."""

    success: bool
    message: str
    value: T | None = None


@dataclass(frozen=True, slots=True)
class BlobInfo:
    """Selected properties for a listed blob."""

    name: str
    size: int
    last_modified: datetime | None
    metadata: dict[str, str]
    tags: dict[str, str]
