from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Any, Mapping
from uuid import uuid4


@dataclass(slots=True)
class TodoItem:
    id: str
    title: str
    description: str
    completed: bool
    created_at: str
    category: str
    etag: str | None = field(default=None, repr=False, compare=False)

    @classmethod
    def new(cls, title: str, description: str, category: str) -> TodoItem:
        return cls(
            id=str(uuid4()),
            title=title,
            description=description,
            completed=False,
            created_at=datetime.now(timezone.utc).isoformat(),
            category=category,
        )

    @classmethod
    def from_document(cls, document: Mapping[str, Any]) -> TodoItem:
        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document["description"]),
            completed=bool(document["completed"]),
            created_at=str(document["created_at"]),
            category=str(document["category"]),
            etag=document.get("_etag"),
        )

    def to_document(self) -> dict[str, Any]:
        return {
            "id": self.id,
            "title": self.title,
            "description": self.description,
            "completed": self.completed,
            "created_at": self.created_at,
            "category": self.category,
        }


@dataclass(frozen=True, slots=True)
class TodoPage:
    number: int
    items: tuple[TodoItem, ...]
    request_charge: float
