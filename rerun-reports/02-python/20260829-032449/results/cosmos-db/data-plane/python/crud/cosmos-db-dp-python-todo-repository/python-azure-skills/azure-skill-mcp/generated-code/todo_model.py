from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Any, Generic, Mapping, TypeVar
from uuid import uuid4


T = TypeVar("T")


@dataclass
class ToDoItem:
    id: str
    title: str
    description: str
    completed: bool
    created_at: datetime
    category: str
    etag: str | None = field(default=None, repr=False, compare=False)

    def __post_init__(self) -> None:
        if not self.id:
            raise ValueError("id must not be empty")
        if not self.title:
            raise ValueError("title must not be empty")
        if not self.category:
            raise ValueError("category must not be empty")
        if self.created_at.tzinfo is None:
            raise ValueError("created_at must include timezone information")

    @classmethod
    def new(cls, title: str, description: str, category: str) -> "ToDoItem":
        return cls(
            id=str(uuid4()),
            title=title,
            description=description,
            completed=False,
            created_at=datetime.now(timezone.utc),
            category=category,
        )

    @classmethod
    def from_document(cls, document: Mapping[str, Any]) -> "ToDoItem":
        created_at = datetime.fromisoformat(
            str(document["created_at"]).replace("Z", "+00:00")
        )
        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document["description"]),
            completed=bool(document["completed"]),
            created_at=created_at,
            category=str(document["category"]),
            etag=document.get("_etag"),
        )

    def to_document(self) -> dict[str, Any]:
        return {
            "id": self.id,
            "title": self.title,
            "description": self.description,
            "completed": self.completed,
            "created_at": self.created_at.astimezone(timezone.utc).isoformat(),
            "category": self.category,
        }


@dataclass(frozen=True)
class OperationResult(Generic[T]):
    value: T
    request_charge: float


@dataclass(frozen=True)
class QueryPage:
    number: int
    items: tuple[ToDoItem, ...]
    request_charge: float
