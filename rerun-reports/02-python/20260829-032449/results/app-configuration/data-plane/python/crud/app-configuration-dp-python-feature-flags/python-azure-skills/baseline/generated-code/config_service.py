"""Cached sync and async access to Azure App Configuration."""

from __future__ import annotations

from dataclasses import dataclass
from threading import RLock
from typing import Any, Protocol

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError

NULL_LABEL = "\0"


class SyncConfigurationClient(Protocol):
    def get_configuration_setting(self, **kwargs: Any) -> Any: ...

    def list_configuration_settings(self, **kwargs: Any) -> Any: ...


class AsyncConfigurationClient(Protocol):
    async def get_configuration_setting(self, **kwargs: Any) -> Any: ...

    def list_configuration_settings(self, **kwargs: Any) -> Any: ...


@dataclass(frozen=True)
class _CacheEntry:
    value: str | None
    etag: Any


def _label_filter(label: str | None) -> str:
    return NULL_LABEL if label is None else label


class ConfigurationService:
    """Retrieve and cache App Configuration values with conditional requests."""

    def __init__(self, client: SyncConfigurationClient) -> None:
        self._client = client
        self._settings: dict[tuple[str, str | None], _CacheEntry] = {}
        self._direct_requests: set[tuple[str, str | None]] = set()
        self._prefixes: dict[tuple[str, str | None], dict[str, str | None]] = {}
        self._lock = RLock()

    def get_setting(self, key: str, label: str | None = None) -> str | None:
        """Return a setting, using its ETag to avoid downloading unchanged data."""
        cache_key = (key, label)
        with self._lock:
            self._direct_requests.add(cache_key)
            return self._get_setting(key, label, conditional=True)

    def _get_setting(
        self, key: str, label: str | None, *, conditional: bool
    ) -> str | None:
        cache_key = (key, label)
        cached = self._settings.get(cache_key)
        kwargs: dict[str, Any] = {"key": key, "label": label}
        if conditional and cached is not None and cached.etag is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfNotModified,
            )

        try:
            setting = self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            if cached is None:
                raise
            return cached.value
        except ResourceNotFoundError:
            self._settings.pop(cache_key, None)
            raise

        if setting is None:
            if cached is None:
                raise RuntimeError(
                    "App Configuration returned no setting without a cached value"
                )
            return cached.value
        entry = _CacheEntry(setting.value, setting.etag)
        self._settings[cache_key] = entry
        return entry.value

    def list_settings(
        self, prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return settings matching a prefix, cached until a coordinated refresh."""
        request = (prefix, label)
        with self._lock:
            cached = self._prefixes.get(request)
            if cached is not None:
                return dict(cached)
            return self._load_prefix(prefix, label)

    def _load_prefix(
        self, prefix: str, label: str | None
    ) -> dict[str, str | None]:
        values: dict[str, str | None] = {}
        entries: dict[tuple[str, str | None], _CacheEntry] = {}
        settings = self._client.list_configuration_settings(
            key_filter=f"{prefix}*",
            label_filter=_label_filter(label),
        )
        for setting in settings:
            values[setting.key] = setting.value
            entries[(setting.key, label)] = _CacheEntry(setting.value, setting.etag)

        self._settings.update(entries)
        self._prefixes[(prefix, label)] = values
        return dict(values)

    def refresh_all(self) -> None:
        """Force a full refresh of every directly read key and cached prefix."""
        with self._lock:
            direct = tuple(self._direct_requests)
            prefixes = tuple(self._prefixes)
            for key, label in direct:
                self._get_setting(key, label, conditional=False)
            for prefix, label in prefixes:
                self._load_prefix(prefix, label)

    def clear_cache(self) -> None:
        """Discard cached values and remembered requests."""
        with self._lock:
            self._settings.clear()
            self._direct_requests.clear()
            self._prefixes.clear()


class AsyncConfigurationService:
    """Async counterpart to :class:`ConfigurationService`."""

    def __init__(self, client: AsyncConfigurationClient) -> None:
        import asyncio

        self._client = client
        self._settings: dict[tuple[str, str | None], _CacheEntry] = {}
        self._direct_requests: set[tuple[str, str | None]] = set()
        self._prefixes: dict[tuple[str, str | None], dict[str, str | None]] = {}
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: str | None = None) -> str | None:
        """Return a setting, using its ETag to avoid downloading unchanged data."""
        cache_key = (key, label)
        async with self._lock:
            self._direct_requests.add(cache_key)
            return await self._get_setting(key, label, conditional=True)

    async def _get_setting(
        self, key: str, label: str | None, *, conditional: bool
    ) -> str | None:
        cache_key = (key, label)
        cached = self._settings.get(cache_key)
        kwargs: dict[str, Any] = {"key": key, "label": label}
        if conditional and cached is not None and cached.etag is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfNotModified,
            )

        try:
            setting = await self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            if cached is None:
                raise
            return cached.value
        except ResourceNotFoundError:
            self._settings.pop(cache_key, None)
            raise

        if setting is None:
            if cached is None:
                raise RuntimeError(
                    "App Configuration returned no setting without a cached value"
                )
            return cached.value
        entry = _CacheEntry(setting.value, setting.etag)
        self._settings[cache_key] = entry
        return entry.value

    async def list_settings(
        self, prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return settings matching a prefix, cached until a coordinated refresh."""
        request = (prefix, label)
        async with self._lock:
            cached = self._prefixes.get(request)
            if cached is not None:
                return dict(cached)
            return await self._load_prefix(prefix, label)

    async def _load_prefix(
        self, prefix: str, label: str | None
    ) -> dict[str, str | None]:
        values: dict[str, str | None] = {}
        entries: dict[tuple[str, str | None], _CacheEntry] = {}
        settings = self._client.list_configuration_settings(
            key_filter=f"{prefix}*",
            label_filter=_label_filter(label),
        )
        async for setting in settings:
            values[setting.key] = setting.value
            entries[(setting.key, label)] = _CacheEntry(setting.value, setting.etag)

        self._settings.update(entries)
        self._prefixes[(prefix, label)] = values
        return dict(values)

    async def refresh_all(self) -> None:
        """Force a full refresh of every directly read key and cached prefix."""
        async with self._lock:
            direct = tuple(self._direct_requests)
            prefixes = tuple(self._prefixes)
            for key, label in direct:
                await self._get_setting(key, label, conditional=False)
            for prefix, label in prefixes:
                await self._load_prefix(prefix, label)

    async def clear_cache(self) -> None:
        """Discard cached values and remembered requests."""
        async with self._lock:
            self._settings.clear()
            self._direct_requests.clear()
            self._prefixes.clear()
