"""In-memory caches for Key Vault-backed application configuration."""

from __future__ import annotations

from datetime import datetime, timedelta, timezone
from typing import Iterable, Mapping

from .provider import AsyncSecretProvider, SecretProvider, SecretValue


def _is_near_expiry(
    secret: SecretValue,
    warning_window: timedelta,
    now: datetime | None = None,
) -> bool:
    if secret.expires_on is None:
        return False
    current_time = now or datetime.now(timezone.utc)
    expires_on = secret.expires_on
    if expires_on.tzinfo is None:
        expires_on = expires_on.replace(tzinfo=timezone.utc)
    return expires_on <= current_time + warning_window


class SecretCache:
    """Cache secret values and refresh entries that are approaching expiry."""

    def __init__(
        self,
        provider: SecretProvider,
        *,
        warning_window: timedelta = timedelta(days=7),
        defaults: Mapping[str, str | None] | None = None,
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._defaults = dict(defaults or {})
        self._secrets: dict[str, SecretValue] = {}

    def load_required(self, names: Iterable[str]) -> Mapping[str, str | None]:
        for name in names:
            self.refresh(name)
        return self.values()

    def get(self, name: str, default: str | None = None) -> str | None:
        if name not in self._secrets:
            self.refresh(name, default)
        elif _is_near_expiry(self._secrets[name], self._warning_window):
            self.refresh(name, default)
        return self._secrets[name].value

    def refresh(self, name: str, default: str | None = None) -> str | None:
        fallback = self._defaults.get(name, default)
        self._secrets[name] = self._provider.get_secret(name, fallback)
        return self._secrets[name].value

    def near_expiry(self) -> list[SecretValue]:
        return [
            secret
            for secret in self._secrets.values()
            if _is_near_expiry(secret, self._warning_window)
        ]

    def values(self) -> Mapping[str, str | None]:
        return {name: secret.value for name, secret in self._secrets.items()}


class AsyncSecretCache:
    """Async cache equivalent of :class:`SecretCache`."""

    def __init__(
        self,
        provider: AsyncSecretProvider,
        *,
        warning_window: timedelta = timedelta(days=7),
        defaults: Mapping[str, str | None] | None = None,
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._defaults = dict(defaults or {})
        self._secrets: dict[str, SecretValue] = {}

    async def load_required(
        self, names: Iterable[str]
    ) -> Mapping[str, str | None]:
        for name in names:
            await self.refresh(name)
        return self.values()

    async def get(self, name: str, default: str | None = None) -> str | None:
        if name not in self._secrets:
            await self.refresh(name, default)
        elif _is_near_expiry(self._secrets[name], self._warning_window):
            await self.refresh(name, default)
        return self._secrets[name].value

    async def refresh(self, name: str, default: str | None = None) -> str | None:
        fallback = self._defaults.get(name, default)
        self._secrets[name] = await self._provider.get_secret(name, fallback)
        return self._secrets[name].value

    def near_expiry(self) -> list[SecretValue]:
        return [
            secret
            for secret in self._secrets.values()
            if _is_near_expiry(secret, self._warning_window)
        ]

    def values(self) -> Mapping[str, str | None]:
        return {name: secret.value for name, secret in self._secrets.items()}
