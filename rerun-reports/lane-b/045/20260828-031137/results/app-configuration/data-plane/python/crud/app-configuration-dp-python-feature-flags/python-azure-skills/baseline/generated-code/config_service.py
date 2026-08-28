from __future__ import annotations

from dataclasses import dataclass
from threading import RLock
from typing import Dict, Optional, Tuple

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core import MatchConditions
from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError


_NULL_LABEL = "\0"


@dataclass(frozen=True)
class _CachedSetting:
    value: Optional[str]
    etag: Optional[str]


class ConfigurationService:
    """Synchronous, ETag-aware access to Azure App Configuration."""

    def __init__(
        self,
        endpoint: str,
        credential: TokenCredential,
        *,
        client: Optional[AzureAppConfigurationClient] = None,
    ) -> None:
        self._client = client or AzureAppConfigurationClient(endpoint, credential)
        self._settings: Dict[Tuple[str, Optional[str]], _CachedSetting] = {}
        self._prefixes: Dict[Tuple[str, Optional[str]], Dict[str, Optional[str]]] = {}
        self._lock = RLock()

    def get_setting(
        self,
        key: str,
        label: Optional[str] = None,
        *,
        force_refresh: bool = False,
    ) -> Optional[str]:
        cache_key = (key, label)
        with self._lock:
            cached = self._settings.get(cache_key)

        request_options = {}
        if cached is not None and cached.etag is not None and not force_refresh:
            request_options = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfNotModified,
            }

        try:
            setting = self._client.get_configuration_setting(
                key=key,
                label=label,
                **request_options,
            )
        except ResourceNotModifiedError:
            return cached.value if cached is not None else None

        if setting is None:
            if cached is None:
                raise RuntimeError("Azure App Configuration returned no setting without a cache entry")
            return cached.value

        updated = _CachedSetting(setting.value, setting.etag)
        with self._lock:
            self._settings[cache_key] = updated
        return updated.value

    def get_setting_with_label(self, key: str, label: str) -> Optional[str]:
        return self.get_setting(key, label)

    def list_settings(
        self,
        key_prefix: str,
        label: Optional[str] = None,
        *,
        force_refresh: bool = False,
    ) -> Dict[str, Optional[str]]:
        query_key = (key_prefix, label)
        with self._lock:
            cached = self._prefixes.get(query_key)
            if cached is not None and not force_refresh:
                return dict(cached)

        values: Dict[str, Optional[str]] = {}
        for setting in self._client.list_configuration_settings(
            key_filter=f"{key_prefix}*",
            label_filter=label if label is not None else _NULL_LABEL,
        ):
            values[setting.key] = setting.value
            with self._lock:
                self._settings[(setting.key, setting.label)] = _CachedSetting(
                    setting.value,
                    setting.etag,
                )

        with self._lock:
            self._prefixes[query_key] = values
        return dict(values)

    def refresh_all(self) -> None:
        """Reload every setting and prefix query that has been cached."""
        with self._lock:
            setting_keys = list(self._settings)
            prefix_queries = list(self._prefixes)

        for key, label in setting_keys:
            try:
                self.get_setting(key, label, force_refresh=True)
            except ResourceNotFoundError:
                with self._lock:
                    self._settings.pop((key, label), None)
        for prefix, label in prefix_queries:
            self.list_settings(prefix, label, force_refresh=True)

    def close(self) -> None:
        self._client.close()


class AsyncConfigurationService:
    """Asynchronous, ETag-aware access to Azure App Configuration."""

    def __init__(
        self,
        endpoint: str,
        credential: AsyncTokenCredential,
        *,
        client: Optional[AsyncAzureAppConfigurationClient] = None,
    ) -> None:
        self._client = client or AsyncAzureAppConfigurationClient(endpoint, credential)
        self._settings: Dict[Tuple[str, Optional[str]], _CachedSetting] = {}
        self._prefixes: Dict[Tuple[str, Optional[str]], Dict[str, Optional[str]]] = {}

    async def get_setting(
        self,
        key: str,
        label: Optional[str] = None,
        *,
        force_refresh: bool = False,
    ) -> Optional[str]:
        cache_key = (key, label)
        cached = self._settings.get(cache_key)

        request_options = {}
        if cached is not None and cached.etag is not None and not force_refresh:
            request_options = {
                "etag": cached.etag,
                "match_condition": MatchConditions.IfNotModified,
            }

        try:
            setting = await self._client.get_configuration_setting(
                key=key,
                label=label,
                **request_options,
            )
        except ResourceNotModifiedError:
            return cached.value if cached is not None else None

        if setting is None:
            if cached is None:
                raise RuntimeError("Azure App Configuration returned no setting without a cache entry")
            return cached.value

        updated = _CachedSetting(setting.value, setting.etag)
        self._settings[cache_key] = updated
        return updated.value

    async def get_setting_with_label(self, key: str, label: str) -> Optional[str]:
        return await self.get_setting(key, label)

    async def list_settings(
        self,
        key_prefix: str,
        label: Optional[str] = None,
        *,
        force_refresh: bool = False,
    ) -> Dict[str, Optional[str]]:
        query_key = (key_prefix, label)
        cached = self._prefixes.get(query_key)
        if cached is not None and not force_refresh:
            return dict(cached)

        values: Dict[str, Optional[str]] = {}
        settings = self._client.list_configuration_settings(
            key_filter=f"{key_prefix}*",
            label_filter=label if label is not None else _NULL_LABEL,
        )
        async for setting in settings:
            values[setting.key] = setting.value
            self._settings[(setting.key, setting.label)] = _CachedSetting(
                setting.value,
                setting.etag,
            )

        self._prefixes[query_key] = values
        return dict(values)

    async def refresh_all(self) -> None:
        """Reload every setting and prefix query that has been cached."""
        setting_keys = list(self._settings)
        prefix_queries = list(self._prefixes)

        for key, label in setting_keys:
            try:
                await self.get_setting(key, label, force_refresh=True)
            except ResourceNotFoundError:
                self._settings.pop((key, label), None)
        for prefix, label in prefix_queries:
            await self.list_settings(prefix, label, force_refresh=True)

    async def close(self) -> None:
        await self._client.close()
