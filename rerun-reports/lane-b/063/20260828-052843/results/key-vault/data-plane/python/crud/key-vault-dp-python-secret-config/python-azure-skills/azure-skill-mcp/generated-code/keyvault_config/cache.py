from __future__ import annotations

from datetime import datetime, timedelta, timezone
from typing import Mapping

from .provider import AsyncSecretProvider, SecretProvider, SecretSnapshot


def _is_near_expiry(
    snapshot: SecretSnapshot,
    warning_window: timedelta,
    now: datetime | None = None,
) -> bool:
    if not snapshot.found or snapshot.expires_on is None:
        return False
    current_time = now or datetime.now(timezone.utc)
    return snapshot.expires_on <= current_time + warning_window


class SecretCache:
    def __init__(
        self,
        provider: SecretProvider,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._cache: dict[str, SecretSnapshot] = {}
        self._defaults: dict[str, str | None] = {}

    def load_required(
        self,
        required: Mapping[str, str | None],
    ) -> Mapping[str, str | None]:
        self._defaults.update(required)
        for name, default in required.items():
            self.refresh(name, default)
        self.refresh_expiring()
        return {name: self._cache[name].value for name in required}

    def get(self, name: str, default: str | None = None) -> str | None:
        if name not in self._cache:
            self._defaults.setdefault(name, default)
            return self.refresh(name, self._defaults[name]).value
        if _is_near_expiry(self._cache[name], self._warning_window):
            self.refresh(name)
        return self._cache[name].value

    def refresh(
        self,
        name: str,
        default: str | None = None,
    ) -> SecretSnapshot:
        if name in self._defaults and default is None:
            default = self._defaults[name]
        else:
            self._defaults[name] = default
        snapshot = self._provider.get_secret_with_metadata(name, default)
        self._cache[name] = snapshot
        return snapshot

    def refresh_expiring(self) -> list[SecretSnapshot]:
        refreshed: list[SecretSnapshot] = []
        for name, snapshot in list(self._cache.items()):
            if _is_near_expiry(snapshot, self._warning_window):
                refreshed.append(self.refresh(name))
        return refreshed

    def expiring_secrets(self) -> list[SecretSnapshot]:
        return [
            snapshot
            for snapshot in self._cache.values()
            if _is_near_expiry(snapshot, self._warning_window)
        ]


class AsyncSecretCache:
    def __init__(
        self,
        provider: AsyncSecretProvider,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._cache: dict[str, SecretSnapshot] = {}
        self._defaults: dict[str, str | None] = {}

    async def load_required(
        self,
        required: Mapping[str, str | None],
    ) -> Mapping[str, str | None]:
        self._defaults.update(required)
        for name, default in required.items():
            await self.refresh(name, default)
        await self.refresh_expiring()
        return {name: self._cache[name].value for name in required}

    async def get(
        self,
        name: str,
        default: str | None = None,
    ) -> str | None:
        if name not in self._cache:
            self._defaults.setdefault(name, default)
            return (await self.refresh(name, self._defaults[name])).value
        if _is_near_expiry(self._cache[name], self._warning_window):
            await self.refresh(name)
        return self._cache[name].value

    async def refresh(
        self,
        name: str,
        default: str | None = None,
    ) -> SecretSnapshot:
        if name in self._defaults and default is None:
            default = self._defaults[name]
        else:
            self._defaults[name] = default
        snapshot = await self._provider.get_secret_with_metadata(name, default)
        self._cache[name] = snapshot
        return snapshot

    async def refresh_expiring(self) -> list[SecretSnapshot]:
        refreshed: list[SecretSnapshot] = []
        for name, snapshot in list(self._cache.items()):
            if _is_near_expiry(snapshot, self._warning_window):
                refreshed.append(await self.refresh(name))
        return refreshed

    def expiring_secrets(self) -> list[SecretSnapshot]:
        return [
            snapshot
            for snapshot in self._cache.values()
            if _is_near_expiry(snapshot, self._warning_window)
        ]
