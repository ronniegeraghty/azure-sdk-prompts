from __future__ import annotations

from dataclasses import asdict, dataclass, field
from datetime import datetime, timezone
from typing import Generic, TypeVar
from uuid import uuid4


@dataclass(slots=True)
class TodoItem:
    id: str
    title: str
    description: str
    completed: bool
    created_at: datetime
    category: str
    etag: str | None = field(default=None, repr=False, compare=False)

    @classmethod
    def new(cls, title: str, description: str, category: str) -> TodoItem:
        return cls(
            id=str(uuid4()),
            title=title,
            description=description,
            completed=False,
            created_at=datetime.now(timezone.utc),
            category=category,
        )

    @classmethod
    def from_document(cls, document: dict[str, object]) -> TodoItem:
        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document["description"]),
            completed=bool(document["completed"]),
            created_at=datetime.fromisoformat(str(document["created_at"])),
            category=str(document["category"]),
            etag=str(document["_etag"]) if document.get("_etag") else None,
        )

    def to_document(self) -> dict[str, object]:
        document = asdict(self)
        document.pop("etag")
        document["created_at"] = self.created_at.isoformat()
        return document


T = TypeVar("T")


@dataclass(frozen=True, slots=True)
class OperationResult(Generic[T]):
    value: T
    request_charge: float


@dataclass(frozen=True, slots=True)
class QueryPage:
    number: int
    items: list[TodoItem]
    request_charge: float


class TodoConflictError(RuntimeError):
    """Raised when an update would overwrite a newer version of an item."""
