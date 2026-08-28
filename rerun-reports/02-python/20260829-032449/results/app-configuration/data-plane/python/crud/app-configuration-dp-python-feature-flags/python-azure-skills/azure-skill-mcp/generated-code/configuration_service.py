from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock
from typing import Optional

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError


@dataclass(frozen=True)
class _CachedSetting:
    value: str
    etag: str


_CacheKey = tuple[str, Optional[str]]


class ConfigurationService:
    """Cached synchronous access to Azure App Configuration."""

    def __init__(self, client: AzureAppConfigurationClient) -> None:
        self._client = client
        self._cache: dict[_CacheKey, _CachedSetting] = {}
        self._single_queries: set[_CacheKey] = set()
        self._prefix_queries: set[_CacheKey] = set()
        self._lock = RLock()

    def get_setting(self, key: str, label: Optional[str] = None) -> str:
        """Get a setting, using its cached ETag to avoid downloading unchanged data."""
        if not key:
            raise ValueError("key must not be empty")

        cache_key = (key, label)
        with self._lock:
            self._single_queries.add(cache_key)
            return self._get_setting_locked(key, label)

    def list_settings(self, prefix: str, label: Optional[str] = None) -> dict[str, str]:
        """List settings under a key prefix, reusing values whose ETags are unchanged."""
        if not prefix:
            raise ValueError("prefix must not be empty")

        query = (prefix, label)
        with self._lock:
            self._prefix_queries.add(query)
            metadata = self._client.list_configuration_settings(
                key_filter=f"{prefix}*",
                label_filter=_label_filter(label),
                fields=["key", "label", "etag"],
            )

            values: dict[str, str] = {}
            seen: set[_CacheKey] = set()
            for setting in metadata:
                if setting.key is None:
                    continue

                cache_key = (setting.key, setting.label)
                seen.add(cache_key)
                cached = self._cache.get(cache_key)
                etag = str(setting.etag)
                if cached is not None and cached.etag == etag:
                    values[setting.key] = cached.value
                else:
                    values[setting.key] = self._get_setting_locked(setting.key, setting.label)

            stale_keys = [
                cache_key
                for cache_key in self._cache
                if cache_key[0].startswith(prefix)
                and cache_key[1] == label
                and cache_key not in seen
            ]
            for cache_key in stale_keys:
                del self._cache[cache_key]

            return values

    def refresh_all(self) -> None:
        """Revalidate all keys and prefixes that have previously been requested."""
        with self._lock:
            single_queries = tuple(self._single_queries)
            prefix_queries = tuple(self._prefix_queries)

        for prefix, label in prefix_queries:
            self.list_settings(prefix, label)
        for key, label in single_queries:
            self.get_setting(key, label)

    def clear_cache(self) -> None:
        with self._lock:
            self._cache.clear()

    def _get_setting_locked(self, key: str, label: Optional[str]) -> str:
        cache_key = (key, label)
        cached = self._cache.get(cache_key)
        kwargs = {}
        if cached is not None:
            kwargs = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfModified,
            }

        try:
            setting = self._client.get_configuration_setting(key=key, label=label, **kwargs)
        except ResourceNotModifiedError:
            if cached is None:
                raise
            return cached.value
        except ResourceNotFoundError:
            self._cache.pop(cache_key, None)
            raise

        if setting is None:
            if cached is not None:
                return cached.value
            raise ResourceNotFoundError(f"Configuration setting {key!r} was not found")
        if setting.value is None:
            raise ValueError(f"Configuration setting {key!r} has no value")

        result = _CachedSetting(value=setting.value, etag=str(setting.etag))
        self._cache[cache_key] = result
        return result.value


class AsyncConfigurationService:
    """Cached asynchronous access to Azure App Configuration."""

    def __init__(self, client: AsyncAzureAppConfigurationClient) -> None:
        self._client = client
        self._cache: dict[_CacheKey, _CachedSetting] = {}
        self._single_queries: set[_CacheKey] = set()
        self._prefix_queries: set[_CacheKey] = set()
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: Optional[str] = None) -> str:
        """Get a setting, using its cached ETag to avoid downloading unchanged data."""
        if not key:
            raise ValueError("key must not be empty")

        cache_key = (key, label)
        async with self._lock:
            self._single_queries.add(cache_key)
            return await self._get_setting_locked(key, label)

    async def list_settings(
        self, prefix: str, label: Optional[str] = None
    ) -> dict[str, str]:
        """List settings under a key prefix, reusing values whose ETags are unchanged."""
        if not prefix:
            raise ValueError("prefix must not be empty")

        query = (prefix, label)
        async with self._lock:
            self._prefix_queries.add(query)
            metadata = self._client.list_configuration_settings(
                key_filter=f"{prefix}*",
                label_filter=_label_filter(label),
                fields=["key", "label", "etag"],
            )

            values: dict[str, str] = {}
            seen: set[_CacheKey] = set()
            async for setting in metadata:
                if setting.key is None:
                    continue

                cache_key = (setting.key, setting.label)
                seen.add(cache_key)
                cached = self._cache.get(cache_key)
                etag = str(setting.etag)
                if cached is not None and cached.etag == etag:
                    values[setting.key] = cached.value
                else:
                    values[setting.key] = await self._get_setting_locked(
                        setting.key, setting.label
                    )

            stale_keys = [
                cache_key
                for cache_key in self._cache
                if cache_key[0].startswith(prefix)
                and cache_key[1] == label
                and cache_key not in seen
            ]
            for cache_key in stale_keys:
                del self._cache[cache_key]

            return values

    async def refresh_all(self) -> None:
        """Revalidate all keys and prefixes that have previously been requested."""
        async with self._lock:
            single_queries = tuple(self._single_queries)
            prefix_queries = tuple(self._prefix_queries)

        for prefix, label in prefix_queries:
            await self.list_settings(prefix, label)
        for key, label in single_queries:
            await self.get_setting(key, label)

    async def clear_cache(self) -> None:
        async with self._lock:
            self._cache.clear()

    async def _get_setting_locked(self, key: str, label: Optional[str]) -> str:
        cache_key = (key, label)
        cached = self._cache.get(cache_key)
        kwargs = {}
        if cached is not None:
            kwargs = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfModified,
            }

        try:
            setting = await self._client.get_configuration_setting(
                key=key, label=label, **kwargs
            )
        except ResourceNotModifiedError:
            if cached is None:
                raise
            return cached.value
        except ResourceNotFoundError:
            self._cache.pop(cache_key, None)
            raise

        if setting is None:
            if cached is not None:
                return cached.value
            raise ResourceNotFoundError(f"Configuration setting {key!r} was not found")
        if setting.value is None:
            raise ValueError(f"Configuration setting {key!r} has no value")

        result = _CachedSetting(value=setting.value, etag=str(setting.etag))
        self._cache[cache_key] = result
        return result.value


def _label_filter(label: Optional[str]) -> str:
    # App Configuration represents the null label with a NUL filter.
    return label if label is not None else "\0"
