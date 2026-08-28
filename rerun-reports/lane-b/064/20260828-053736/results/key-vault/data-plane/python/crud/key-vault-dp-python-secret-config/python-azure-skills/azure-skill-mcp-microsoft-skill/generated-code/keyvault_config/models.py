from __future__ import annotations

from dataclasses import dataclass
from datetime import datetime, timezone


@dataclass(frozen=True, slots=True)
class SecretInfo:
    name: str
    value: str | None
    version: str | None
    expires_on: datetime | None

    def expires_within(self, deadline: datetime) -> bool:
        if self.expires_on is None:
            return False
        expires_on = self.expires_on
        if expires_on.tzinfo is None:
            expires_on = expires_on.replace(tzinfo=timezone.utc)
        return expires_on <= deadline
