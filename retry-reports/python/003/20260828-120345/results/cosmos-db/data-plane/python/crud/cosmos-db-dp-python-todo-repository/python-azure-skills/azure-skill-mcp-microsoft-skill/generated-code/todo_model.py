from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Any
from uuid import uuid4


@dataclass(slots=True)
class TodoItem:
    title: str
    description: str
    category: str
    completed: bool = False
    id: str = field(default_factory=lambda: str(uuid4()))
    created_at: datetime = field(default_factory=lambda: datetime.now(timezone.utc))
    etag: str | None = field(default=None, repr=False, compare=False)

    def to_document(self) -> dict[str, Any]:
        return {
            "id": self.id,
            "title": self.title,
            "description": self.description,
            "completed": self.completed,
            "created_at": self.created_at.isoformat(),
            "category": self.category,
        }

    @classmethod
    def from_document(cls, document: dict[str, Any]) -> TodoItem:
        created_at = datetime.fromisoformat(document["created_at"])
        if created_at.tzinfo is None:
            created_at = created_at.replace(tzinfo=timezone.utc)

        return cls(
            id=document["id"],
            title=document["title"],
            description=document["description"],
            completed=document["completed"],
            created_at=created_at,
            category=document["category"],
            etag=document.get("_etag"),
        )
