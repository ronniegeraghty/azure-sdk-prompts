from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock
from typing import TypeAlias

from azure.appconfiguration import AzureAppConfigurationClient, ConfigurationSetting
from azure.appconfiguration.aio import (
    AzureAppConfigurationClient as AsyncAzureAppConfigurationClient,
)
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError


CacheKey: TypeAlias = tuple[str, str | None]
PrefixQuery: TypeAlias = tuple[str, str | None]


@dataclass(frozen=True)
class _CachedSetting:
    value: str | None
    etag: str


@dataclass(frozen=True)
class _CachedPrefix:
    values: dict[str, str | None]
    page_etags: tuple[str, ...]


def _label_filter(label: str | None) -> str:
    return label if label is not None else "\0"


def _cache_setting(setting: ConfigurationSetting) -> _CachedSetting:
    return _CachedSetting(value=setting.value, etag=str(setting.etag))


class ConfigurationService:
    """Cached, synchronous access to Azure App Configuration."""

    def __init__(self, client: AzureAppConfigurationClient) -> None:
        self._client = client
        self._settings: dict[CacheKey, _CachedSetting] = {}
        self._prefixes: dict[PrefixQuery, _CachedPrefix] = {}
        self._lock = RLock()

    def get_setting(self, key: str, label: str | None = None) -> str | None:
        identity = (key, label)
        with self._lock:
            cached = self._settings.get(identity)
            try:
                if cached is None:
                    setting = self._client.get_configuration_setting(
                        key=key, label=label
                    )
                else:
                    setting = self._client.get_configuration_setting(
                        key=key,
                        label=label,
                        etag=cached.etag,
                        match_condition=MatchConditions.IfModified,
                    )
            except ResourceNotFoundError:
                self._settings.pop(identity, None)
                return None

            if setting is None:
                return cached.value if cached is not None else None

            self._settings[identity] = _cache_setting(setting)
            return setting.value

    def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        query = (key_prefix, label)
        key_filter = f"{key_prefix}*"
        label_filter = _label_filter(label)

        with self._lock:
            cached = self._prefixes.get(query)
            if cached is not None:
                page_etags = tuple(
                    str(page.etag)
                    for page in self._client.check_configuration_settings(
                        key_filter=key_filter, label_filter=label_filter
                    ).by_page()
                )
                if page_etags == cached.page_etags:
                    return dict(cached.values)

            values: dict[str, str | None] = {}
            page_etags_list: list[str] = []
            pages = self._client.list_configuration_settings(
                key_filter=key_filter, label_filter=label_filter
            ).by_page()
            for page in pages:
                page_etags_list.append(str(page.etag))
                for setting in page:
                    values[setting.key] = setting.value

            page_etags = tuple(page_etags_list)
            self._prefixes[query] = _CachedPrefix(values, page_etags)
            return dict(values)

    def refresh_all(self) -> None:
        """Invalidate and reload every key and prefix requested so far."""
        with self._lock:
            setting_queries = tuple(self._settings)
            prefix_queries = tuple(self._prefixes)
            self._settings.clear()
            self._prefixes.clear()

        for key, label in setting_queries:
            self.get_setting(key, label)
        for prefix, label in prefix_queries:
            self.list_settings(prefix, label)


class AsyncConfigurationService:
    """Cached, asynchronous access to Azure App Configuration."""

    def __init__(self, client: AsyncAzureAppConfigurationClient) -> None:
        self._client = client
        self._settings: dict[CacheKey, _CachedSetting] = {}
        self._prefixes: dict[PrefixQuery, _CachedPrefix] = {}
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: str | None = None) -> str | None:
        identity = (key, label)
        async with self._lock:
            cached = self._settings.get(identity)
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
                self._settings.pop(identity, None)
                return None

            if setting is None:
                return cached.value if cached is not None else None

            self._settings[identity] = _cache_setting(setting)
            return setting.value

    async def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        query = (key_prefix, label)
        key_filter = f"{key_prefix}*"
        label_filter = _label_filter(label)

        async with self._lock:
            cached = self._prefixes.get(query)
            if cached is not None:
                page_etags = tuple(
                    [
                        str(page.etag)
                        async for page in self._client.check_configuration_settings(
                            key_filter=key_filter, label_filter=label_filter
                        ).by_page()
                    ]
                )
                if page_etags == cached.page_etags:
                    return dict(cached.values)

            values: dict[str, str | None] = {}
            page_etags_list: list[str] = []
            pages = self._client.list_configuration_settings(
                key_filter=key_filter, label_filter=label_filter
            ).by_page()
            async for page in pages:
                page_etags_list.append(str(page.etag))
                async for setting in page:
                    values[setting.key] = setting.value

            page_etags = tuple(page_etags_list)
            self._prefixes[query] = _CachedPrefix(values, page_etags)
            return dict(values)

    async def refresh_all(self) -> None:
        """Invalidate and reload every key and prefix requested so far."""
        async with self._lock:
            setting_queries = tuple(self._settings)
            prefix_queries = tuple(self._prefixes)
            self._settings.clear()
            self._prefixes.clear()

        for key, label in setting_queries:
            await self.get_setting(key, label)
        for prefix, label in prefix_queries:
            await self.list_settings(prefix, label)
