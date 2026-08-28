from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime, timezone
from typing import Generic, List, Mapping, Optional, TypeVar


T = TypeVar("T", covariant=True)


@dataclass
class TodoItem:
    id: str
    title: str
    description: str
    completed: bool
    category: str
    created_at: datetime = field(default_factory=lambda: datetime.now(timezone.utc))
    etag: Optional[str] = field(default=None, repr=False, compare=False)

    def to_cosmos_document(self) -> dict:
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
    def from_cosmos_document(cls, document: Mapping[str, object]) -> "TodoItem":
        created_at = datetime.fromisoformat(str(document["createdAt"]).replace("Z", "+00:00"))
        return cls(
            id=str(document["id"]),
            title=str(document["title"]),
            description=str(document["description"]),
            completed=bool(document["completed"]),
            created_at=created_at,
            category=str(document["category"]),
            etag=str(document["_etag"]) if document.get("_etag") is not None else None,
        )


@dataclass(frozen=True)
class OperationResult(Generic[T]):
    value: T
    request_charge: float


@dataclass(frozen=True)
class QueryPage:
    page_number: int
    items: List[TodoItem]
    request_charge: float
