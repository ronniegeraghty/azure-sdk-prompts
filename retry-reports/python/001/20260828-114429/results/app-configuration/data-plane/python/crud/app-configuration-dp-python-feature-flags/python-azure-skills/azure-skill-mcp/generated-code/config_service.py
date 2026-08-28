from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError


@dataclass(frozen=True)
class _CachedSetting:
    value: str | None
    etag: str | None


@dataclass(frozen=True)
class _CachedPrefix:
    values: dict[str, str | None]
    page_etags: tuple[str, ...]


def _prefix_filter(prefix: str) -> str:
    escaped = prefix.replace("\\", "\\\\").replace(",", "\\,").replace("*", "\\*")
    return f"{escaped}*"


class ConfigurationService:
    def __init__(self, client: AzureAppConfigurationClient) -> None:
        self._client = client
        self._settings: dict[tuple[str, str | None], _CachedSetting] = {}
        self._prefixes: dict[tuple[str, str | None], _CachedPrefix] = {}
        self._tracked_settings: set[tuple[str, str | None]] = set()
        self._tracked_prefixes: set[tuple[str, str | None]] = set()
        self._lock = RLock()

    def get_setting(self, key: str, label: str | None = None) -> str | None:
        identity = (key, label)
        with self._lock:
            self._tracked_settings.add(identity)
            cached = self._settings.get(identity)

            try:
                if cached and cached.etag:
                    setting = self._client.get_configuration_setting(
                        key=key,
                        label=label,
                        etag=cached.etag,
                        match_condition=MatchConditions.IfModified,
                    )
                else:
                    setting = self._client.get_configuration_setting(key=key, label=label)
            except ResourceNotModifiedError:
                return cached.value if cached else None
            except ResourceNotFoundError:
                self._settings[identity] = _CachedSetting(None, None)
                return None

            if setting is None:
                self._settings[identity] = _CachedSetting(None, None)
                return None

            value = setting.value
            self._settings[identity] = _CachedSetting(
                value=value,
                etag=str(setting.etag) if setting.etag is not None else None,
            )
            return value

    def get_setting_with_label(self, key: str, label: str) -> str | None:
        return self.get_setting(key, label)

    def list_settings(self, key_prefix: str, label: str | None = None) -> dict[str, str | None]:
        identity = (key_prefix, label)
        with self._lock:
            self._tracked_prefixes.add(identity)
            cached = self._prefixes.get(identity)
            if cached:
                current_etags = self._check_prefix(key_prefix, label)
                if current_etags == cached.page_etags:
                    return dict(cached.values)

            values, page_etags = self._download_prefix(key_prefix, label)
            self._prefixes[identity] = _CachedPrefix(values, page_etags)
            return dict(values)

    def refresh_all(self) -> None:
        with self._lock:
            settings = tuple(self._tracked_settings)
            prefixes = tuple(self._tracked_prefixes)
            self._settings.clear()
            self._prefixes.clear()

        for key, label in settings:
            self.get_setting(key, label)
        for prefix, label in prefixes:
            self.list_settings(prefix, label)

    def _check_prefix(self, key_prefix: str, label: str | None) -> tuple[str, ...]:
        pager = self._client.check_configuration_settings(
            key_filter=_prefix_filter(key_prefix),
            label_filter=label,
        )
        return tuple(str(page.etag) for page in pager.by_page())

    def _download_prefix(
        self, key_prefix: str, label: str | None
    ) -> tuple[dict[str, str | None], tuple[str, ...]]:
        values: dict[str, str | None] = {}
        page_etags: list[str] = []
        pager = self._client.list_configuration_settings(
            key_filter=_prefix_filter(key_prefix),
            label_filter=label,
        )
        for page in pager.by_page():
            page_etags.append(str(page.etag))
            for setting in page:
                values[setting.key] = setting.value
        return values, tuple(page_etags)


class AsyncConfigurationService:
    def __init__(self, client: AsyncAzureAppConfigurationClient) -> None:
        self._client = client
        self._settings: dict[tuple[str, str | None], _CachedSetting] = {}
        self._prefixes: dict[tuple[str, str | None], _CachedPrefix] = {}
        self._tracked_settings: set[tuple[str, str | None]] = set()
        self._tracked_prefixes: set[tuple[str, str | None]] = set()
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: str | None = None) -> str | None:
        identity = (key, label)
        async with self._lock:
            self._tracked_settings.add(identity)
            cached = self._settings.get(identity)

            try:
                if cached and cached.etag:
                    setting = await self._client.get_configuration_setting(
                        key=key,
                        label=label,
                        etag=cached.etag,
                        match_condition=MatchConditions.IfModified,
                    )
                else:
                    setting = await self._client.get_configuration_setting(key=key, label=label)
            except ResourceNotModifiedError:
                return cached.value if cached else None
            except ResourceNotFoundError:
                self._settings[identity] = _CachedSetting(None, None)
                return None

            if setting is None:
                self._settings[identity] = _CachedSetting(None, None)
                return None

            value = setting.value
            self._settings[identity] = _CachedSetting(
                value=value,
                etag=str(setting.etag) if setting.etag is not None else None,
            )
            return value

    async def get_setting_with_label(self, key: str, label: str) -> str | None:
        return await self.get_setting(key, label)

    async def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        identity = (key_prefix, label)
        async with self._lock:
            self._tracked_prefixes.add(identity)
            cached = self._prefixes.get(identity)
            if cached:
                current_etags = await self._check_prefix(key_prefix, label)
                if current_etags == cached.page_etags:
                    return dict(cached.values)

            values, page_etags = await self._download_prefix(key_prefix, label)
            self._prefixes[identity] = _CachedPrefix(values, page_etags)
            return dict(values)

    async def refresh_all(self) -> None:
        async with self._lock:
            settings = tuple(self._tracked_settings)
            prefixes = tuple(self._tracked_prefixes)
            self._settings.clear()
            self._prefixes.clear()

        for key, label in settings:
            await self.get_setting(key, label)
        for prefix, label in prefixes:
            await self.list_settings(prefix, label)

    async def _check_prefix(self, key_prefix: str, label: str | None) -> tuple[str, ...]:
        pager = self._client.check_configuration_settings(
            key_filter=_prefix_filter(key_prefix),
            label_filter=label,
        )
        etags: list[str] = []
        async for page in pager.by_page():
            etags.append(str(page.etag))
        return tuple(etags)

    async def _download_prefix(
        self, key_prefix: str, label: str | None
    ) -> tuple[dict[str, str | None], tuple[str, ...]]:
        values: dict[str, str | None] = {}
        page_etags: list[str] = []
        pager = self._client.list_configuration_settings(
            key_filter=_prefix_filter(key_prefix),
            label_filter=label,
        )
        async for page in pager.by_page():
            page_etags.append(str(page.etag))
            async for setting in page:
                values[setting.key] = setting.value
        return values, tuple(page_etags)
