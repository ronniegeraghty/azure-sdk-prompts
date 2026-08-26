"""Cached sync and async access to Azure App Configuration."""

from __future__ import annotations

import asyncio
import logging
from dataclasses import dataclass
from threading import RLock
from typing import TypeAlias

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError

logger = logging.getLogger(__name__)

SettingIdentity: TypeAlias = tuple[str, str | None]
PrefixQuery: TypeAlias = tuple[str, str | None]


@dataclass(frozen=True)
class _CachedSetting:
    value: str | None
    etag: str


def _identity(key: str, label: str | None) -> SettingIdentity:
    return key, label


def _label_filter(label: str | None) -> str:
    # App Configuration uses the NUL label filter to select only unlabeled values.
    return label if label is not None else "\0"


class ConfigurationService:
    """Retrieve and cache settings using conditional ETag requests."""

    def __init__(self, client: AzureAppConfigurationClient) -> None:
        self._client = client
        self._cache: dict[SettingIdentity, _CachedSetting] = {}
        self._exact_queries: set[SettingIdentity] = set()
        self._prefix_members: dict[PrefixQuery, set[SettingIdentity]] = {}
        self._lock = RLock()

    def get_setting(
        self,
        key: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> str | None:
        """Get one value, optionally selecting an environment label."""
        if not key:
            raise ValueError("key must not be empty")

        identity = _identity(key, label)
        with self._lock:
            self._exact_queries.add(identity)
            return self._get_setting_locked(identity, force_refresh=force_refresh)

    def _get_setting_locked(
        self,
        identity: SettingIdentity,
        *,
        force_refresh: bool,
    ) -> str | None:
        key, label = identity
        cached = self._cache.get(identity)
        request_options: dict[str, object] = {}
        if cached is not None and not force_refresh:
            # If-None-Match returns no payload when the cached ETag is current.
            request_options = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfModified,
            }

        try:
            setting = self._client.get_configuration_setting(
                key=key,
                label=label,
                **request_options,
            )
        except ResourceNotModifiedError:
            if cached is None:
                raise RuntimeError("received a not-modified response without a cached value")
            return cached.value
        except ResourceNotFoundError:
            self._cache.pop(identity, None)
            raise

        # Some transports represent HTTP 304 as None instead of raising.
        if setting is None:
            if cached is None:
                raise RuntimeError("App Configuration returned no setting")
            return cached.value

        self._cache[identity] = _CachedSetting(setting.value, str(setting.etag))
        return setting.value

    def list_settings(
        self,
        key_prefix: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> dict[str, str | None]:
        """List values under a key prefix while only downloading changed values."""
        if not key_prefix:
            raise ValueError("key_prefix must not be empty")

        query = (key_prefix, label)
        with self._lock:
            known_members = self._prefix_members.get(query)
            if known_members is None or force_refresh:
                settings = self._client.list_configuration_settings(
                    key_filter=f"{key_prefix}*",
                    label_filter=_label_filter(label),
                )
                current_members: set[SettingIdentity] = set()
                result: dict[str, str | None] = {}
                for setting in settings:
                    identity = _identity(setting.key, setting.label)
                    current_members.add(identity)
                    self._cache[identity] = _CachedSetting(
                        setting.value,
                        str(setting.etag),
                    )
                    result[setting.key] = setting.value
            else:
                # Fetch only identity and ETag metadata, then download changed values.
                settings = self._client.list_configuration_settings(
                    key_filter=f"{key_prefix}*",
                    label_filter=_label_filter(label),
                    fields=["key", "label", "etag"],
                )
                current_members = set()
                result = {}
                for setting in settings:
                    identity = _identity(setting.key, setting.label)
                    current_members.add(identity)
                    cached = self._cache.get(identity)
                    if cached is None or cached.etag != str(setting.etag):
                        self._get_setting_locked(identity, force_refresh=True)
                    result[setting.key] = self._cache[identity].value

            self._prefix_members[query] = current_members
            self._remove_unreferenced(known_members or set(), current_members)
            return result

    def _remove_unreferenced(
        self,
        previous_members: set[SettingIdentity],
        current_members: set[SettingIdentity],
    ) -> None:
        for identity in previous_members - current_members:
            referenced_by_prefix = any(
                identity in members for members in self._prefix_members.values()
            )
            if identity not in self._exact_queries and not referenced_by_prefix:
                self._cache.pop(identity, None)

    def refresh_all(self) -> None:
        """Fully reload every exact key and prefix previously requested."""
        with self._lock:
            exact_queries = tuple(self._exact_queries)
            prefix_queries = tuple(self._prefix_members)

        for identity in exact_queries:
            try:
                self.get_setting(*identity, force_refresh=True)
            except ResourceNotFoundError:
                logger.info("Cached setting was deleted: key=%s label=%s", *identity)

        for prefix, label in prefix_queries:
            self.list_settings(prefix, label, force_refresh=True)

    def clear_cache(self) -> None:
        """Clear values and remembered queries."""
        with self._lock:
            self._cache.clear()
            self._exact_queries.clear()
            self._prefix_members.clear()


class AsyncConfigurationService:
    """Async counterpart to :class:`ConfigurationService`."""

    def __init__(self, client: AsyncAzureAppConfigurationClient) -> None:
        self._client = client
        self._cache: dict[SettingIdentity, _CachedSetting] = {}
        self._exact_queries: set[SettingIdentity] = set()
        self._prefix_members: dict[PrefixQuery, set[SettingIdentity]] = {}
        self._lock = asyncio.Lock()

    async def get_setting(
        self,
        key: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> str | None:
        """Get one value, optionally selecting an environment label."""
        if not key:
            raise ValueError("key must not be empty")

        identity = _identity(key, label)
        async with self._lock:
            self._exact_queries.add(identity)
            return await self._get_setting_locked(identity, force_refresh=force_refresh)

    async def _get_setting_locked(
        self,
        identity: SettingIdentity,
        *,
        force_refresh: bool,
    ) -> str | None:
        key, label = identity
        cached = self._cache.get(identity)
        request_options: dict[str, object] = {}
        if cached is not None and not force_refresh:
            request_options = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfModified,
            }

        try:
            setting = await self._client.get_configuration_setting(
                key=key,
                label=label,
                **request_options,
            )
        except ResourceNotModifiedError:
            if cached is None:
                raise RuntimeError("received a not-modified response without a cached value")
            return cached.value
        except ResourceNotFoundError:
            self._cache.pop(identity, None)
            raise

        if setting is None:
            if cached is None:
                raise RuntimeError("App Configuration returned no setting")
            return cached.value

        self._cache[identity] = _CachedSetting(setting.value, str(setting.etag))
        return setting.value

    async def list_settings(
        self,
        key_prefix: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> dict[str, str | None]:
        """List values under a key prefix while only downloading changed values."""
        if not key_prefix:
            raise ValueError("key_prefix must not be empty")

        query = (key_prefix, label)
        async with self._lock:
            known_members = self._prefix_members.get(query)
            if known_members is None or force_refresh:
                settings = self._client.list_configuration_settings(
                    key_filter=f"{key_prefix}*",
                    label_filter=_label_filter(label),
                )
                current_members: set[SettingIdentity] = set()
                result: dict[str, str | None] = {}
                async for setting in settings:
                    identity = _identity(setting.key, setting.label)
                    current_members.add(identity)
                    self._cache[identity] = _CachedSetting(
                        setting.value,
                        str(setting.etag),
                    )
                    result[setting.key] = setting.value
            else:
                settings = self._client.list_configuration_settings(
                    key_filter=f"{key_prefix}*",
                    label_filter=_label_filter(label),
                    fields=["key", "label", "etag"],
                )
                current_members = set()
                result = {}
                async for setting in settings:
                    identity = _identity(setting.key, setting.label)
                    current_members.add(identity)
                    cached = self._cache.get(identity)
                    if cached is None or cached.etag != str(setting.etag):
                        await self._get_setting_locked(identity, force_refresh=True)
                    result[setting.key] = self._cache[identity].value

            self._prefix_members[query] = current_members
            self._remove_unreferenced(known_members or set(), current_members)
            return result

    def _remove_unreferenced(
        self,
        previous_members: set[SettingIdentity],
        current_members: set[SettingIdentity],
    ) -> None:
        for identity in previous_members - current_members:
            referenced_by_prefix = any(
                identity in members for members in self._prefix_members.values()
            )
            if identity not in self._exact_queries and not referenced_by_prefix:
                self._cache.pop(identity, None)

    async def refresh_all(self) -> None:
        """Fully reload every exact key and prefix previously requested."""
        async with self._lock:
            exact_queries = tuple(self._exact_queries)
            prefix_queries = tuple(self._prefix_members)

        for identity in exact_queries:
            try:
                await self.get_setting(*identity, force_refresh=True)
            except ResourceNotFoundError:
                logger.info("Cached setting was deleted: key=%s label=%s", *identity)

        for prefix, label in prefix_queries:
            await self.list_settings(prefix, label, force_refresh=True)

    async def clear_cache(self) -> None:
        """Clear values and remembered queries."""
        async with self._lock:
            self._cache.clear()
            self._exact_queries.clear()
            self._prefix_members.clear()
