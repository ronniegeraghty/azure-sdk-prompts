from __future__ import annotations

from dataclasses import dataclass
from threading import RLock
from typing import Any

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotModifiedError


SettingId = tuple[str, str | None]
PrefixId = tuple[str, str | None]


@dataclass(frozen=True)
class CachedSetting:
    value: str | None
    etag: str | None


class ConfigurationService:
    """Cached access to an Azure App Configuration synchronous client."""

    def __init__(self, client: Any) -> None:
        self._client = client
        self._settings: dict[SettingId, CachedSetting] = {}
        self._prefixes: dict[PrefixId, dict[str, str | None]] = {}
        self._lock = RLock()

    def get_setting(
        self, key: str, label: str | None = None, *, force_refresh: bool = False
    ) -> str | None:
        setting_id = (key, label)
        with self._lock:
            cached = self._settings.get(setting_id)

        kwargs: dict[str, Any] = {"key": key, "label": label}
        if cached is not None and cached.etag is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfModified,
            )
        elif cached is not None and not force_refresh:
            return cached.value

        try:
            setting = self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            return cached.value if cached is not None else None

        updated = CachedSetting(setting.value, _etag_text(setting.etag))
        with self._lock:
            self._settings[setting_id] = updated
        return updated.value

    def get_setting_with_label(self, key: str, label: str) -> str | None:
        return self.get_setting(key, label)

    def list_settings(
        self,
        key_prefix: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> dict[str, str | None]:
        prefix_id = (key_prefix, label)
        with self._lock:
            cached = self._prefixes.get(prefix_id)
            if cached is not None and not force_refresh:
                return dict(cached)

        settings = self._client.list_configuration_settings(
            key_filter=f"{key_prefix}*", label_filter=label
        )
        values = {setting.key: setting.value for setting in settings}
        with self._lock:
            self._prefixes[prefix_id] = values
            for setting in settings:
                self._settings[(setting.key, setting.label)] = CachedSetting(
                    setting.value, _etag_text(setting.etag)
                )
        return dict(values)

    def refresh_all(self) -> None:
        with self._lock:
            setting_ids = list(self._settings)
            prefix_ids = list(self._prefixes)

        for key, label in setting_ids:
            self.get_setting(key, label, force_refresh=True)
        for key_prefix, label in prefix_ids:
            self.list_settings(key_prefix, label, force_refresh=True)


class AsyncConfigurationService:
    """Cached access to an Azure App Configuration asynchronous client."""

    def __init__(self, client: Any) -> None:
        self._client = client
        self._settings: dict[SettingId, CachedSetting] = {}
        self._prefixes: dict[PrefixId, dict[str, str | None]] = {}

    async def get_setting(
        self, key: str, label: str | None = None, *, force_refresh: bool = False
    ) -> str | None:
        setting_id = (key, label)
        cached = self._settings.get(setting_id)
        kwargs: dict[str, Any] = {"key": key, "label": label}
        if cached is not None and cached.etag is not None:
            kwargs.update(
                etag=cached.etag,
                match_condition=MatchConditions.IfModified,
            )
        elif cached is not None and not force_refresh:
            return cached.value

        try:
            setting = await self._client.get_configuration_setting(**kwargs)
        except ResourceNotModifiedError:
            return cached.value if cached is not None else None

        updated = CachedSetting(setting.value, _etag_text(setting.etag))
        self._settings[setting_id] = updated
        return updated.value

    async def get_setting_with_label(self, key: str, label: str) -> str | None:
        return await self.get_setting(key, label)

    async def list_settings(
        self,
        key_prefix: str,
        label: str | None = None,
        *,
        force_refresh: bool = False,
    ) -> dict[str, str | None]:
        prefix_id = (key_prefix, label)
        cached = self._prefixes.get(prefix_id)
        if cached is not None and not force_refresh:
            return dict(cached)

        settings = self._client.list_configuration_settings(
            key_filter=f"{key_prefix}*", label_filter=label
        )
        received = [setting async for setting in settings]
        values = {setting.key: setting.value for setting in received}
        self._prefixes[prefix_id] = values
        for setting in received:
            self._settings[(setting.key, setting.label)] = CachedSetting(
                setting.value, _etag_text(setting.etag)
            )
        return dict(values)

    async def refresh_all(self) -> None:
        setting_ids = list(self._settings)
        prefix_ids = list(self._prefixes)
        for key, label in setting_ids:
            await self.get_setting(key, label, force_refresh=True)
        for key_prefix, label in prefix_ids:
            await self.list_settings(key_prefix, label, force_refresh=True)


def _etag_text(etag: Any) -> str | None:
    return str(etag) if etag is not None else None
