"""In-memory caching for Key Vault-backed configuration."""

from __future__ import annotations

import asyncio
from datetime import datetime, timedelta, timezone
from typing import Iterable, Mapping, Optional

from .provider import AsyncSecretProvider, SecretDetails, SecretProvider


def _utc(value: datetime) -> datetime:
    if value.tzinfo is None:
        return value.replace(tzinfo=timezone.utc)
    return value.astimezone(timezone.utc)


def _is_near_expiry(
    details: SecretDetails, warning_window: timedelta, now: datetime
) -> bool:
    return (
        details.expires_on is not None
        and _utc(details.expires_on) <= now + warning_window
    )


class SecretCache:
    """Cache secret values and refresh entries approaching expiration."""

    def __init__(
        self,
        provider: SecretProvider,
        required_keys: Iterable[str] = (),
        defaults: Mapping[str, Optional[str]] | None = None,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window must not be negative")
        self._provider = provider
        self._required_keys = tuple(dict.fromkeys(required_keys))
        self._defaults = dict(defaults or {})
        self._warning_window = warning_window
        self._entries: dict[str, SecretDetails] = {}

    def load_required(self) -> dict[str, Optional[str]]:
        for name in self._required_keys:
            self.refresh(name)
        return self.values()

    def get(self, name: str, default: Optional[str] = None) -> Optional[str]:
        entry = self._entries.get(name)
        if entry is None or _is_near_expiry(
            entry, self._warning_window, datetime.now(timezone.utc)
        ):
            entry = self.refresh(name, default)
        return entry.value

    def refresh(
        self, name: str, default: Optional[str] = None
    ) -> SecretDetails:
        fallback = self._defaults.get(name, default)
        entry = self._provider.get_secret_details(name, default=fallback)
        self._entries[name] = entry
        return entry

    def refresh_expiring(self) -> tuple[str, ...]:
        now = datetime.now(timezone.utc)
        names = tuple(
            name
            for name, entry in self._entries.items()
            if _is_near_expiry(entry, self._warning_window, now)
        )
        for name in names:
            self.refresh(name)
        return names

    def expiring_secrets(self) -> dict[str, datetime]:
        now = datetime.now(timezone.utc)
        return {
            name: entry.expires_on
            for name, entry in self._entries.items()
            if entry.expires_on is not None
            and _is_near_expiry(entry, self._warning_window, now)
        }

    def values(self) -> dict[str, Optional[str]]:
        return {name: entry.value for name, entry in self._entries.items()}


class AsyncSecretCache:
    """Asynchronous counterpart to :class:`SecretCache`."""

    def __init__(
        self,
        provider: AsyncSecretProvider,
        required_keys: Iterable[str] = (),
        defaults: Mapping[str, Optional[str]] | None = None,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window must not be negative")
        self._provider = provider
        self._required_keys = tuple(dict.fromkeys(required_keys))
        self._defaults = dict(defaults or {})
        self._warning_window = warning_window
        self._entries: dict[str, SecretDetails] = {}

    async def load_required(self) -> dict[str, Optional[str]]:
        await asyncio.gather(*(self.refresh(name) for name in self._required_keys))
        return self.values()

    async def get(self, name: str, default: Optional[str] = None) -> Optional[str]:
        entry = self._entries.get(name)
        if entry is None or _is_near_expiry(
            entry, self._warning_window, datetime.now(timezone.utc)
        ):
            entry = await self.refresh(name, default)
        return entry.value

    async def refresh(
        self, name: str, default: Optional[str] = None
    ) -> SecretDetails:
        fallback = self._defaults.get(name, default)
        entry = await self._provider.get_secret_details(name, default=fallback)
        self._entries[name] = entry
        return entry

    async def refresh_expiring(self) -> tuple[str, ...]:
        now = datetime.now(timezone.utc)
        names = tuple(
            name
            for name, entry in self._entries.items()
            if _is_near_expiry(entry, self._warning_window, now)
        )
        await asyncio.gather(*(self.refresh(name) for name in names))
        return names

    def expiring_secrets(self) -> dict[str, datetime]:
        now = datetime.now(timezone.utc)
        return {
            name: entry.expires_on
            for name, entry in self._entries.items()
            if entry.expires_on is not None
            and _is_near_expiry(entry, self._warning_window, now)
        }

    def values(self) -> dict[str, Optional[str]]:
        return {name: entry.value for name, entry in self._entries.items()}
