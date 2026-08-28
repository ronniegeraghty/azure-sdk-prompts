from __future__ import annotations

from datetime import timedelta
from typing import Mapping

from secret_provider import AsyncSecretProvider, SecretRecord, SyncSecretProvider


class SyncSecretCache:
    def __init__(
        self,
        provider: SyncSecretProvider,
        *,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._entries: dict[str, SecretRecord] = {}
        self._defaults: dict[str, str | None] = {}

    def bulk_load(
        self,
        required_keys: Mapping[str, str | None],
    ) -> dict[str, str | None]:
        for name, default in required_keys.items():
            self._defaults[name] = default
            self._entries[name] = self._provider.get_secret_record(name, default)
        return {name: self._entries[name].value for name in required_keys}

    def get(self, name: str, default: str | None = None) -> str | None:
        if name not in self._entries:
            self._defaults[name] = default
            self._entries[name] = self._provider.get_secret_record(name, default)
        elif self._entries[name].expires_within(self._warning_window):
            self.refresh(name)
        return self._entries[name].value

    def refresh(self, name: str) -> str | None:
        default = self._defaults.get(name)
        self._entries[name] = self._provider.get_secret_record(name, default)
        return self._entries[name].value

    def refresh_expiring(self) -> list[str]:
        refreshed: list[str] = []
        for name, record in list(self._entries.items()):
            if record.expires_within(self._warning_window):
                self.refresh(name)
                refreshed.append(name)
        return refreshed

    def expiring_secrets(self) -> list[SecretRecord]:
        return [
            record
            for record in self._entries.values()
            if record.expires_within(self._warning_window)
        ]


class AsyncSecretCache:
    def __init__(
        self,
        provider: AsyncSecretProvider,
        *,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._entries: dict[str, SecretRecord] = {}
        self._defaults: dict[str, str | None] = {}

    async def bulk_load(
        self,
        required_keys: Mapping[str, str | None],
    ) -> dict[str, str | None]:
        for name, default in required_keys.items():
            self._defaults[name] = default
            self._entries[name] = await self._provider.get_secret_record(
                name,
                default,
            )
        return {name: self._entries[name].value for name in required_keys}

    async def get(self, name: str, default: str | None = None) -> str | None:
        if name not in self._entries:
            self._defaults[name] = default
            self._entries[name] = await self._provider.get_secret_record(
                name,
                default,
            )
        elif self._entries[name].expires_within(self._warning_window):
            await self.refresh(name)
        return self._entries[name].value

    async def refresh(self, name: str) -> str | None:
        default = self._defaults.get(name)
        self._entries[name] = await self._provider.get_secret_record(
            name,
            default,
        )
        return self._entries[name].value

    async def refresh_expiring(self) -> list[str]:
        refreshed: list[str] = []
        for name, record in list(self._entries.items()):
            if record.expires_within(self._warning_window):
                await self.refresh(name)
                refreshed.append(name)
        return refreshed

    def expiring_secrets(self) -> list[SecretRecord]:
        return [
            record
            for record in self._entries.values()
            if record.expires_within(self._warning_window)
        ]
