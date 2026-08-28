from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock
from typing import Any

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    ConfigurationSettingFields,
)
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError


SettingId = tuple[str, str | None]


@dataclass
class _CachedSetting:
    value: str | None
    etag: str


def _cache_key(key: str, label: str | None) -> SettingId:
    return key, label


def _list_label_filter(label: str | None) -> str:
    # App Configuration represents the null label with the NUL filter.
    return label if label is not None else "\0"


class ConfigurationService:
    """Synchronous Azure App Configuration reader with an ETag-aware cache."""

    def __init__(self, endpoint: str, credential: Any) -> None:
        self._client = AzureAppConfigurationClient(endpoint, credential)
        self._cache: dict[SettingId, _CachedSetting] = {}
        self._lock = RLock()

    def get_setting(self, key: str, label: str | None = None) -> str | None:
        with self._lock:
            return self._get_setting_locked(key, label)

    def _get_setting_locked(self, key: str, label: str | None) -> str | None:
        setting_id = _cache_key(key, label)
        cached = self._cache.get(setting_id)

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
            self._cache.pop(setting_id, None)
            raise

        if setting is None:
            if cached is None:
                raise RuntimeError("App Configuration returned no setting")
            return cached.value
        self._cache[setting_id] = _CachedSetting(setting.value, str(setting.etag))
        return setting.value

    def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return matching values while downloading only new or changed payloads."""
        with self._lock:
            metadata = self._client.list_configuration_settings(
                key_filter=f"{key_prefix}*",
                label_filter=_list_label_filter(label),
                fields=[
                    ConfigurationSettingFields.KEY,
                    ConfigurationSettingFields.LABEL,
                    ConfigurationSettingFields.ETAG,
                ],
            )

            result: dict[str, str | None] = {}
            current_ids: set[SettingId] = set()
            for item in metadata:
                setting_id = _cache_key(item.key, item.label)
                current_ids.add(setting_id)
                cached = self._cache.get(setting_id)
                if cached is not None and cached.etag == str(item.etag):
                    result[item.key] = cached.value
                else:
                    result[item.key] = self._get_setting_locked(item.key, item.label)

            stale_ids = {
                setting_id
                for setting_id in self._cache
                if setting_id[1] == label
                and setting_id[0].startswith(key_prefix)
                and setting_id not in current_ids
            }
            for setting_id in stale_ids:
                del self._cache[setting_id]

            return result

    def refresh_all(self) -> None:
        """Force a complete refresh of every setting currently held in the cache."""
        with self._lock:
            setting_ids = list(self._cache)
            for key, label in setting_ids:
                try:
                    setting = self._client.get_configuration_setting(key=key, label=label)
                except ResourceNotFoundError:
                    self._cache.pop((key, label), None)
                    continue
                self._cache[(key, label)] = _CachedSetting(
                    setting.value, str(setting.etag)
                )

    def close(self) -> None:
        self._client.close()

    def __enter__(self) -> ConfigurationService:
        return self

    def __exit__(self, *args: object) -> None:
        self.close()


class AsyncConfigurationService:
    """Asynchronous Azure App Configuration reader with an ETag-aware cache."""

    def __init__(self, endpoint: str, credential: Any) -> None:
        self._client = AsyncAzureAppConfigurationClient(endpoint, credential)
        self._cache: dict[SettingId, _CachedSetting] = {}
        self._lock = asyncio.Lock()

    async def get_setting(self, key: str, label: str | None = None) -> str | None:
        async with self._lock:
            return await self._get_setting_locked(key, label)

    async def _get_setting_locked(
        self, key: str, label: str | None
    ) -> str | None:
        setting_id = _cache_key(key, label)
        cached = self._cache.get(setting_id)

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
            self._cache.pop(setting_id, None)
            raise

        if setting is None:
            if cached is None:
                raise RuntimeError("App Configuration returned no setting")
            return cached.value
        self._cache[setting_id] = _CachedSetting(setting.value, str(setting.etag))
        return setting.value

    async def list_settings(
        self, key_prefix: str, label: str | None = None
    ) -> dict[str, str | None]:
        """Return matching values while downloading only new or changed payloads."""
        async with self._lock:
            metadata = self._client.list_configuration_settings(
                key_filter=f"{key_prefix}*",
                label_filter=_list_label_filter(label),
                fields=[
                    ConfigurationSettingFields.KEY,
                    ConfigurationSettingFields.LABEL,
                    ConfigurationSettingFields.ETAG,
                ],
            )

            result: dict[str, str | None] = {}
            current_ids: set[SettingId] = set()
            async for item in metadata:
                setting_id = _cache_key(item.key, item.label)
                current_ids.add(setting_id)
                cached = self._cache.get(setting_id)
                if cached is not None and cached.etag == str(item.etag):
                    result[item.key] = cached.value
                else:
                    result[item.key] = await self._get_setting_locked(
                        item.key, item.label
                    )

            stale_ids = {
                setting_id
                for setting_id in self._cache
                if setting_id[1] == label
                and setting_id[0].startswith(key_prefix)
                and setting_id not in current_ids
            }
            for setting_id in stale_ids:
                del self._cache[setting_id]

            return result

    async def refresh_all(self) -> None:
        """Force a complete refresh of every setting currently held in the cache."""
        async with self._lock:
            setting_ids = list(self._cache)
            for key, label in setting_ids:
                try:
                    setting = await self._client.get_configuration_setting(
                        key=key, label=label
                    )
                except ResourceNotFoundError:
                    self._cache.pop((key, label), None)
                    continue
                self._cache[(key, label)] = _CachedSetting(
                    setting.value, str(setting.etag)
                )

    async def close(self) -> None:
        await self._client.close()

    async def __aenter__(self) -> AsyncConfigurationService:
        return self

    async def __aexit__(self, *args: object) -> None:
        await self.close()
