from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Any, Mapping


@dataclass(slots=True)
class TodoItem:
    id: str
    title: str
    description: str
    completed: bool
    category: str
    created_at: datetime = field(default_factory=lambda: datetime.now(timezone.utc))
    etag: str | None = field(default=None, repr=False, compare=False)

    def to_document(self) -> dict[str, Any]:
        created_at = self.created_at
        if created_at.tzinfo is None:
            created_at = created_at.replace(tzinfo=timezone.utc)

        return {
            "id": self.id,
            "title": self.title,
            "description": self.description,
            "completed": self.completed,
            "createdAt": created_at.astimezone(timezone.utc).isoformat(),
            "category": self.category,
        }

    @classmethod
    def from_document(cls, document: Mapping[str, Any]) -> TodoItem:
        created_at = str(document["createdAt"])
        if created_at.endswith("Z"):
            created_at = f"{created_at[:-1]}+00:00"

        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document["description"]),
            completed=bool(document["completed"]),
            created_at=datetime.fromisoformat(created_at),
            category=str(document["category"]),
            etag=str(document["_etag"]) if document.get("_etag") else None,
        )
