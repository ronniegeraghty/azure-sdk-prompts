from __future__ import annotations

from collections.abc import Mapping
from datetime import datetime, timedelta, timezone

from .providers import AsyncSecretProvider, SecretProvider, SecretResult


class SecretCache:
    def __init__(
        self,
        provider: SecretProvider,
        *,
        expiry_warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if expiry_warning_window < timedelta(0):
            raise ValueError("expiry_warning_window cannot be negative")
        self._provider = provider
        self._warning_window = expiry_warning_window
        self._entries: dict[str, SecretResult] = {}
        self._defaults: dict[str, str | None] = {}

    def load_required(self, secrets: Mapping[str, str | None]) -> None:
        self._defaults.update(secrets)
        for name, default in secrets.items():
            self._entries[name] = self._provider.get(name, default)

    def get(self, name: str, default: str | None = None) -> str | None:
        effective_default = self._defaults.get(name, default)
        entry = self._entries.get(name)
        if entry is None or entry.expires_within(self._warning_window):
            entry = self._provider.get(name, effective_default)
            self._entries[name] = entry
        return entry.value

    def refresh(self, name: str) -> SecretResult:
        entry = self._provider.get(name, self._defaults.get(name))
        self._entries[name] = entry
        return entry

    def expiring(
        self,
        *,
        now: datetime | None = None,
    ) -> dict[str, SecretResult]:
        current_time = now or datetime.now(timezone.utc)
        return {
            name: entry
            for name, entry in self._entries.items()
            if entry.expires_within(self._warning_window, now=current_time)
        }


class AsyncSecretCache:
    def __init__(
        self,
        provider: AsyncSecretProvider,
        *,
        expiry_warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if expiry_warning_window < timedelta(0):
            raise ValueError("expiry_warning_window cannot be negative")
        self._provider = provider
        self._warning_window = expiry_warning_window
        self._entries: dict[str, SecretResult] = {}
        self._defaults: dict[str, str | None] = {}

    async def load_required(self, secrets: Mapping[str, str | None]) -> None:
        self._defaults.update(secrets)
        for name, default in secrets.items():
            self._entries[name] = await self._provider.get(name, default)

    async def get(self, name: str, default: str | None = None) -> str | None:
        effective_default = self._defaults.get(name, default)
        entry = self._entries.get(name)
        if entry is None or entry.expires_within(self._warning_window):
            entry = await self._provider.get(name, effective_default)
            self._entries[name] = entry
        return entry.value

    async def refresh(self, name: str) -> SecretResult:
        entry = await self._provider.get(name, self._defaults.get(name))
        self._entries[name] = entry
        return entry

    def expiring(
        self,
        *,
        now: datetime | None = None,
    ) -> dict[str, SecretResult]:
        current_time = now or datetime.now(timezone.utc)
        return {
            name: entry
            for name, entry in self._entries.items()
            if entry.expires_within(self._warning_window, now=current_time)
        }
