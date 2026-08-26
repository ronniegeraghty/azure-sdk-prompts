from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Any, Mapping


@dataclass(slots=True)
class ToDoItem:
    id: str
    title: str
    description: str
    completed: bool
    created_at: str
    category: str
    etag: str | None = field(default=None, repr=False, compare=False)

    @classmethod
    def new(
        cls,
        *,
        id: str,
        title: str,
        description: str,
        category: str,
        completed: bool = False,
    ) -> "ToDoItem":
        return cls(
            id=id,
            title=title,
            description=description,
            completed=completed,
            created_at=datetime.now(timezone.utc).isoformat(),
            category=category,
        )

    @classmethod
    def from_document(cls, document: Mapping[str, Any]) -> "ToDoItem":
        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document.get("description", "")),
            completed=bool(document.get("completed", False)),
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
class QueryPage:
    page_number: int
    items: list[ToDoItem]
    request_charge: float
