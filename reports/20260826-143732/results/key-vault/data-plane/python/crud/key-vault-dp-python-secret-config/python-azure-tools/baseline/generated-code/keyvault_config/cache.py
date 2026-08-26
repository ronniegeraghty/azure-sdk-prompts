import asyncio
from datetime import timedelta
from typing import Dict, Mapping, Optional, Sequence

from .models import SecretInfo
from .provider import AsyncKeyVaultSecretProvider, KeyVaultSecretProvider


class SecretCache:
    """In-memory, expiry-aware cache for synchronous configuration access."""

    def __init__(
        self,
        provider: KeyVaultSecretProvider,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        self._provider = provider
        self._warning_window = warning_window
        self._entries: Dict[str, SecretInfo] = {}
        self._defaults: Dict[str, Optional[str]] = {}

    def load_required(
        self, keys: Sequence[str] | Mapping[str, Optional[str]]
    ) -> Dict[str, Optional[str]]:
        defaults = (
            dict(keys)
            if isinstance(keys, Mapping)
            else {key: None for key in keys}
        )
        self._defaults.update(defaults)
        for name, default in defaults.items():
            self._entries[name] = self._provider.get_secret_info(name, default)
        return {name: entry.value for name, entry in self._entries.items()}

    def get(self, name: str, default: Optional[str] = None) -> Optional[str]:
        if name not in self._entries:
            self._defaults[name] = default
            return self.refresh(name, default)

        if self._entries[name].expires_within(self._warning_window):
            return self.refresh(name)
        return self._entries[name].value

    def refresh(
        self, name: str, default: Optional[str] = None
    ) -> Optional[str]:
        fallback = self._defaults.get(name, default)
        self._defaults[name] = fallback
        self._entries[name] = self._provider.get_secret_info(name, fallback)
        return self._entries[name].value

    def expiring_secrets(self) -> Dict[str, SecretInfo]:
        return {
            name: entry
            for name, entry in self._entries.items()
            if entry.expires_within(self._warning_window)
        }

    def refresh_expiring(self) -> Dict[str, Optional[str]]:
        return {
            name: self.refresh(name)
            for name in tuple(self.expiring_secrets())
        }


class AsyncSecretCache:
    """In-memory, expiry-aware cache for asynchronous configuration access."""

    def __init__(
        self,
        provider: AsyncKeyVaultSecretProvider,
        warning_window: timedelta = timedelta(days=7),
    ) -> None:
        self._provider = provider
        self._warning_window = warning_window
        self._entries: Dict[str, SecretInfo] = {}
        self._defaults: Dict[str, Optional[str]] = {}

    async def load_required(
        self, keys: Sequence[str] | Mapping[str, Optional[str]]
    ) -> Dict[str, Optional[str]]:
        defaults = (
            dict(keys)
            if isinstance(keys, Mapping)
            else {key: None for key in keys}
        )
        self._defaults.update(defaults)
        entries = await asyncio.gather(
            *(
                self._provider.get_secret_info(name, default)
                for name, default in defaults.items()
            )
        )
        self._entries.update(zip(defaults, entries))
        return {name: entry.value for name, entry in self._entries.items()}

    async def get(
        self, name: str, default: Optional[str] = None
    ) -> Optional[str]:
        if name not in self._entries:
            self._defaults[name] = default
            return await self.refresh(name, default)

        if self._entries[name].expires_within(self._warning_window):
            return await self.refresh(name)
        return self._entries[name].value

    async def refresh(
        self, name: str, default: Optional[str] = None
    ) -> Optional[str]:
        fallback = self._defaults.get(name, default)
        self._defaults[name] = fallback
        self._entries[name] = await self._provider.get_secret_info(
            name, fallback
        )
        return self._entries[name].value

    def expiring_secrets(self) -> Dict[str, SecretInfo]:
        return {
            name: entry
            for name, entry in self._entries.items()
            if entry.expires_within(self._warning_window)
        }

    async def refresh_expiring(self) -> Dict[str, Optional[str]]:
        names = tuple(self.expiring_secrets())
        values = await asyncio.gather(*(self.refresh(name) for name in names))
        return dict(zip(names, values))
