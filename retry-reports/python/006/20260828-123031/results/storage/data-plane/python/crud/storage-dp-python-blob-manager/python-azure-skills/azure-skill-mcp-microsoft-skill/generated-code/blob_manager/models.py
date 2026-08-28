"""Public result models for blob operations."""

from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime
from typing import Generic, TypeVar

T = TypeVar("T")


@dataclass(frozen=True)
class OperationResult(Generic[T]):
    success: bool
    message: str
    value: T | None = None

    @classmethod
    def ok(cls, message: str, value: T | None = None) -> "OperationResult[T]":
        return cls(success=True, message=message, value=value)

    @classmethod
    def fail(cls, message: str) -> "OperationResult[T]":
        return cls(success=False, message=message)


@dataclass(frozen=True)
class BlobInfo:
    name: str
    size: int
    last_modified: datetime | None
    metadata: dict[str, str]
    tags: dict[str, str]
