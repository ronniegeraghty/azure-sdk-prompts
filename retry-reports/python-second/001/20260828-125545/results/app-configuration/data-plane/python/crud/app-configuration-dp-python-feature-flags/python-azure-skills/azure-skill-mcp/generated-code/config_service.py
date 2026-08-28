from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock
from typing import Any

from azure.appconfiguration import AzureAppConfigurationClient, ConfigurationSetting
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ResourceNotFoundError


@dataclass(frozen=True)
class _CachedSetting:
    value: str | None
    etag: str | None


class ConfigurationService:
    """Cached synchronous access to Azure App Configuration."""

    def __init__(
        self,
        endpoint: str,
        credential: TokenCredential,
        *,
        client: AzureAppConfigurationClient | None = None,
    ) -> None:
        self._client = client or AzureAppConfigurationClient(endpoint, credential)
        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}
        self._prefix_queries: set[tuple[str, str | None]] = set()
        self._lock = RLock()

    def get_setting(self, key: str, label: str | None = None) -> str | None:
        cache_key = (key, label)
        with self._lock:
            cached = self._cache.get(cache_key)
            try:
                if cached is None:
                    setting = self._client.get_configuration_setting(key=key, label=label)
                else:
                    setting = self._client.get_configuration_setting(
                        key=key,
                        label=label,
                        etag=cached.etag,
                        match_condition=MatchConditions.IfModified,
                    )
            except ResourceNotFoundError:
                self._cache.pop(cache_key, None)
                raise

            if setting is None:
                return cached.value if cached is not None else None

            self._cache_setting(setting)
            return setting.value

    def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return settings keyed by their full App Configuration key."""
        with self._lock:
            self._prefix_queries.add((key_prefix, label))
            seen: set[tuple[str, str | None]] = set()
            result: dict[str, str | None] = {}
            settings = self._client.list_configuration_settings(
                key_filter=f"{key_prefix}*",
                label_filter=label,
                fields=["key", "label", "etag"],
            )

            for metadata in settings:
                if metadata.key is None:
                    continue
                cache_key = (metadata.key, metadata.label)
                seen.add(cache_key)
                cached = self._cache.get(cache_key)
                if cached is not None and cached.etag == metadata.etag:
                    result[metadata.key] = cached.value
                    continue

                setting = self._client.get_configuration_setting(
                    key=metadata.key,
                    label=metadata.label,
                )
                self._cache_setting(setting)
                result[metadata.key] = setting.value

            self._remove_deleted_prefix_entries(key_prefix, label, seen)
            return result

    def refresh_all(self) -> None:
        """Discard and rebuild all configuration previously read by this service."""
        with self._lock:
            cached_keys = set(self._cache)
            prefix_queries = set(self._prefix_queries)
            self._cache.clear()

            for key, label in cached_keys:
                try:
                    setting = self._client.get_configuration_setting(key=key, label=label)
                except ResourceNotFoundError:
                    continue
                self._cache_setting(setting)

            for prefix, label in prefix_queries:
                self.list_settings(prefix, label)

    def close(self) -> None:
        self._client.close()

    def __enter__(self) -> ConfigurationService:
        return self

    def __exit__(self, *args: Any) -> None:
        self.close()

    def _cache_setting(self, setting: ConfigurationSetting) -> None:
        if setting.key is None:
            raise ValueError("App Configuration returned a setting without a key")
        self._cache[(setting.key, setting.label)] = _CachedSetting(
            value=setting.value,
            etag=setting.etag,
        )

    def _remove_deleted_prefix_entries(
        self,
        key_prefix: str,
        label: str | None,
        seen: set[tuple[str, str | None]],
    ) -> None:
        stale = [
            cache_key
            for cache_key in self._cache
            if cache_key[0].startswith(key_prefix)
            and cache_key[1] == label
            and cache_key not in seen
        ]
        for cache_key in stale:
            del self._cache[cache_key]


class AsyncConfigurationService:
    """Cached asynchronous access to Azure App Configuration."""

    def __init__(
        self,
        endpoint: str,
        credential: AsyncTokenCredential,
        *,
        client: AsyncAzureAppConfigurationClient | None = None,
    ) -> None:
        self._client = client or AsyncAzureAppConfigurationClient(endpoint, credential)
        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}
        self._prefix_queries: set[tuple[str, str | None]] = set()
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: str | None = None) -> str | None:
        async with self._lock:
            return await self._get_setting_unlocked(key, label)

    async def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return settings keyed by their full App Configuration key."""
        async with self._lock:
            return await self._list_settings_unlocked(key_prefix, label)

    async def refresh_all(self) -> None:
        """Discard and rebuild all configuration previously read by this service."""
        async with self._lock:
            cached_keys = set(self._cache)
            prefix_queries = set(self._prefix_queries)
            self._cache.clear()

            for key, label in cached_keys:
                try:
                    setting = await self._client.get_configuration_setting(
                        key=key, label=label
                    )
                except ResourceNotFoundError:
                    continue
                self._cache_setting(setting)

            for prefix, label in prefix_queries:
                await self._list_settings_unlocked(prefix, label)

    async def close(self) -> None:
        await self._client.close()

    async def __aenter__(self) -> AsyncConfigurationService:
        return self

    async def __aexit__(self, *args: Any) -> None:
        await self.close()

    async def _get_setting_unlocked(
        self, key: str, label: str | None
    ) -> str | None:
        cache_key = (key, label)
        cached = self._cache.get(cache_key)
        try:
            if cached is None:
                setting = await self._client.get_configuration_setting(
                    key=key, label=label
                )
            else:
                setting = await self._client.get_configuration_setting(
                    key=key,
                    label=label,
                    etag=cached.etag,
                    match_condition=MatchConditions.IfModified,
                )
        except ResourceNotFoundError:
            self._cache.pop(cache_key, None)
            raise

        if setting is None:
            return cached.value if cached is not None else None

        self._cache_setting(setting)
        return setting.value

    async def _list_settings_unlocked(
        self, key_prefix: str, label: str | None
    ) -> dict[str, str | None]:
        self._prefix_queries.add((key_prefix, label))
        seen: set[tuple[str, str | None]] = set()
        result: dict[str, str | None] = {}
        settings = self._client.list_configuration_settings(
            key_filter=f"{key_prefix}*",
            label_filter=label,
            fields=["key", "label", "etag"],
        )

        async for metadata in settings:
            if metadata.key is None:
                continue
            cache_key = (metadata.key, metadata.label)
            seen.add(cache_key)
            cached = self._cache.get(cache_key)
            if cached is not None and cached.etag == metadata.etag:
                result[metadata.key] = cached.value
                continue

            setting = await self._client.get_configuration_setting(
                key=metadata.key,
                label=metadata.label,
            )
            self._cache_setting(setting)
            result[metadata.key] = setting.value

        self._remove_deleted_prefix_entries(key_prefix, label, seen)
        return result

    def _cache_setting(self, setting: ConfigurationSetting) -> None:
        if setting.key is None:
            raise ValueError("App Configuration returned a setting without a key")
        self._cache[(setting.key, setting.label)] = _CachedSetting(
            value=setting.value,
            etag=setting.etag,
        )

    def _remove_deleted_prefix_entries(
        self,
        key_prefix: str,
        label: str | None,
        seen: set[tuple[str, str | None]],
    ) -> None:
        stale = [
            cache_key
            for cache_key in self._cache
            if cache_key[0].startswith(key_prefix)
            and cache_key[1] == label
            and cache_key not in seen
        ]
        for cache_key in stale:
            del self._cache[cache_key]
