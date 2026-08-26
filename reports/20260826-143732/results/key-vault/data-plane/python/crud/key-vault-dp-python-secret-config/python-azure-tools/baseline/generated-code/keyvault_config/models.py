from dataclasses import dataclass
from datetime import datetime, timedelta, timezone
from typing import Optional


@dataclass(frozen=True)
class SecretInfo:
    """A secret value and the metadata needed by the configuration cache."""

    name: str
    value: Optional[str]
    version: Optional[str]
    expires_on: Optional[datetime]
    found: bool

    def expires_within(self, warning_window: timedelta) -> bool:
        if self.expires_on is None:
            return False

        expires_on = self.expires_on
        if expires_on.tzinfo is None:
            expires_on = expires_on.replace(tzinfo=timezone.utc)
        return expires_on <= datetime.now(timezone.utc) + warning_window
