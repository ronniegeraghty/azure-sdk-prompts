"""Cached synchronous and asynchronous Azure App Configuration clients."""

from __future__ import annotations

import asyncio
from dataclasses import dataclass
from threading import RLock
from typing import Any

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError

_NULL_LABEL_FILTER = "\0"


@dataclass(frozen=True)
class _CachedSetting:
    value: str | None
    etag: str


def _cache_key(key: str, label: str | None) -> tuple[str, str | None]:
    return key, label


def _label_filter(label: str | None) -> str:
    return _NULL_LABEL_FILTER if label is None else label


def _etag(setting: Any) -> str:
    return str(setting.etag)


class ConfigurationService:
    """Retrieve and cache settings with conditional ETag requests."""

    def __init__(self, client: Any) -> None:
        self._client = client
        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}
        self._key_queries: set[tuple[str, str | None]] = set()
        self._prefix_queries: set[tuple[str, str | None]] = set()
        self._prefix_members: dict[
            tuple[str, str | None], set[tuple[str, str | None]]
        ] = {}
        self._lock = RLock()

    def get_setting(
        self, key: str, label: str | None = None, *, refresh: bool = False
    ) -> str | None:
        """Return a setting value, optionally checking Azure for a newer ETag."""
        with self._lock:
            self._key_queries.add((key, label))
            value, _ = self._fetch_setting(key, label, refresh=refresh)
            return value

    def get_setting_with_label(self, key: str, label: str) -> str | None:
        """Return a setting for a specific environment label."""
        return self.get_setting(key, label)

    def list_settings(
        self, prefix: str, label: str | None = None, *, refresh: bool = False
    ) -> dict[str, str | None]:
        """Return no-label or specifically labeled settings matching a key prefix."""
        with self._lock:
            query = (prefix, label)
            self._prefix_queries.add(query)
            members = self._prefix_members.get(query)
            if members is not None and not refresh:
                return {
                    key: self._cache[(key, item_label)].value
                    for key, item_label in members
                }

            if members is None:
                settings = self._client.list_configuration_settings(
                    key_filter=f"{prefix}*",
                    label_filter=_label_filter(label),
                )
                new_members: set[tuple[str, str | None]] = set()
                for setting in settings:
                    item = _cache_key(setting.key, setting.label)
                    self._cache[item] = _CachedSetting(setting.value, _etag(setting))
                    new_members.add(item)
            else:
                settings = self._client.list_configuration_settings(
                    key_filter=f"{prefix}*",
                    label_filter=_label_filter(label),
                    fields=["key", "label", "etag"],
                )
                new_members = set()
                for setting in settings:
                    item = _cache_key(setting.key, setting.label)
                    new_members.add(item)
                    cached = self._cache.get(item)
                    if cached is None or cached.etag != _etag(setting):
                        self._fetch_setting(setting.key, setting.label, refresh=True)

                for deleted in members - new_members:
                    self._cache.pop(deleted, None)

            self._prefix_members[query] = new_members
            return {
                key: self._cache[(key, item_label)].value
                for key, item_label in new_members
            }

    def check_for_update(self, key: str, label: str | None = None) -> bool:
        """Report whether a conditional poll finds a change or deletion."""
        with self._lock:
            item = _cache_key(key, label)
            if item not in self._cache:
                self.get_setting(key, label)
                return False
            try:
                _, changed = self._fetch_setting(key, label, refresh=True)
                return changed
            except ResourceNotFoundError:
                self._cache.pop(item, None)
                return True

    def refresh_all(self) -> None:
        """Invalidate and reload all keys and prefixes requested so far."""
        with self._lock:
            key_queries = tuple(self._key_queries)
            prefix_queries = tuple(self._prefix_queries)
            self._cache.clear()
            self._prefix_members.clear()
            for key, label in key_queries:
                try:
                    self._fetch_setting(key, label, refresh=False)
                except ResourceNotFoundError:
                    pass
            for prefix, label in prefix_queries:
                self.list_settings(prefix, label)

    def _fetch_setting(
        self, key: str, label: str | None, *, refresh: bool
    ) -> tuple[str | None, bool]:
        item = _cache_key(key, label)
        cached = self._cache.get(item)
        if cached is not None and not refresh:
            return cached.value, False

        kwargs: dict[str, Any] = {"key": key, "label": label}
        if cached is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfModified,
            )
        try:
            setting = self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            return cached.value, False

        current = _CachedSetting(setting.value, _etag(setting))
        self._cache[item] = current
        return current.value, cached is not None and current.etag != cached.etag


class AsyncConfigurationService:
    """Asynchronous counterpart to :class:`ConfigurationService`."""

    def __init__(self, client: Any) -> None:
        self._client = client
        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}
        self._key_queries: set[tuple[str, str | None]] = set()
        self._prefix_queries: set[tuple[str, str | None]] = set()
        self._prefix_members: dict[
            tuple[str, str | None], set[tuple[str, str | None]]
        ] = {}
        self._lock = asyncio.Lock()

    async def get_setting(
        self, key: str, label: str | None = None, *, refresh: bool = False
    ) -> str | None:
        """Return a setting value, optionally checking Azure for a newer ETag."""
        async with self._lock:
            self._key_queries.add((key, label))
            value, _ = await self._fetch_setting(key, label, refresh=refresh)
            return value

    async def get_setting_with_label(self, key: str, label: str) -> str | None:
        """Return a setting for a specific environment label."""
        return await self.get_setting(key, label)

    async def list_settings(
        self, prefix: str, label: str | None = None, *, refresh: bool = False
    ) -> dict[str, str | None]:
        """Return no-label or specifically labeled settings matching a key prefix."""
        async with self._lock:
            return await self._list_settings_locked(prefix, label, refresh=refresh)

    async def check_for_update(
        self, key: str, label: str | None = None
    ) -> bool:
        """Report whether a conditional poll finds a change or deletion."""
        async with self._lock:
            item = _cache_key(key, label)
            if item not in self._cache:
                await self._fetch_setting(key, label, refresh=False)
                return False
            try:
                _, changed = await self._fetch_setting(key, label, refresh=True)
                return changed
            except ResourceNotFoundError:
                self._cache.pop(item, None)
                return True

    async def refresh_all(self) -> None:
        """Invalidate and reload all keys and prefixes requested so far."""
        async with self._lock:
            key_queries = tuple(self._key_queries)
            prefix_queries = tuple(self._prefix_queries)
            self._cache.clear()
            self._prefix_members.clear()
            for key, label in key_queries:
                try:
                    await self._fetch_setting(key, label, refresh=False)
                except ResourceNotFoundError:
                    pass
            for prefix, label in prefix_queries:
                await self._list_settings_locked(prefix, label, refresh=False)

    async def _list_settings_locked(
        self, prefix: str, label: str | None, *, refresh: bool
    ) -> dict[str, str | None]:
        query = (prefix, label)
        self._prefix_queries.add(query)
        members = self._prefix_members.get(query)
        if members is not None and not refresh:
            return {
                key: self._cache[(key, item_label)].value
                for key, item_label in members
            }

        if members is None:
            settings = self._client.list_configuration_settings(
                key_filter=f"{prefix}*",
                label_filter=_label_filter(label),
            )
            new_members: set[tuple[str, str | None]] = set()
            async for setting in settings:
                item = _cache_key(setting.key, setting.label)
                self._cache[item] = _CachedSetting(setting.value, _etag(setting))
                new_members.add(item)
        else:
            settings = self._client.list_configuration_settings(
                key_filter=f"{prefix}*",
                label_filter=_label_filter(label),
                fields=["key", "label", "etag"],
            )
            new_members = set()
            async for setting in settings:
                item = _cache_key(setting.key, setting.label)
                new_members.add(item)
                cached = self._cache.get(item)
                if cached is None or cached.etag != _etag(setting):
                    await self._fetch_setting(setting.key, setting.label, refresh=True)

            for deleted in members - new_members:
                self._cache.pop(deleted, None)

        self._prefix_members[query] = new_members
        return {
            key: self._cache[(key, item_label)].value
            for key, item_label in new_members
        }

    async def _fetch_setting(
        self, key: str, label: str | None, *, refresh: bool
    ) -> tuple[str | None, bool]:
        item = _cache_key(key, label)
        cached = self._cache.get(item)
        if cached is not None and not refresh:
            return cached.value, False

        kwargs: dict[str, Any] = {"key": key, "label": label}
        if cached is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfModified,
            )
        try:
            setting = await self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            return cached.value, False

        current = _CachedSetting(setting.value, _etag(setting))
        self._cache[item] = current
        return current.value, cached is not None and current.etag != cached.etag
