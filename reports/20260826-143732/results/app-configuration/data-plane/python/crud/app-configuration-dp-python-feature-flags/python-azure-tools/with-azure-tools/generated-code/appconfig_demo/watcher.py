"""Sentinel-based configuration watchers."""

from __future__ import annotations

import asyncio
import logging
from collections.abc import Awaitable, Callable, Sequence
from dataclasses import dataclass
from threading import Event

from azure.core.exceptions import ResourceNotFoundError

from .configuration_service import AsyncConfigurationService, ConfigurationService

logger = logging.getLogger(__name__)

_UNSET = object()
_MISSING = object()


@dataclass(frozen=True)
class SentinelKey:
    """A sentinel key and its optional App Configuration label."""

    key: str
    label: str | None = None

    def __post_init__(self) -> None:
        if not self.key:
            raise ValueError("sentinel key must not be empty")


class ConfigurationWatcher:
    """Poll sentinel keys and fully refresh the sync cache after a change."""

    def __init__(
        self,
        configuration: ConfigurationService,
        sentinels: Sequence[SentinelKey],
        polling_interval: float,
        on_refresh: Callable[[set[SentinelKey]], None] | None = None,
    ) -> None:
        if not sentinels:
            raise ValueError("at least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinels = tuple(sentinels)
        self._polling_interval = polling_interval
        self._on_refresh = on_refresh
        self._last_values: dict[SentinelKey, object] = {
            sentinel: _UNSET for sentinel in sentinels
        }

    def poll_once(self) -> set[SentinelKey]:
        """Check all sentinels once and return those whose values changed."""
        changed: set[SentinelKey] = set()
        for sentinel in self._sentinels:
            try:
                value: object = self._configuration.get_setting(
                    sentinel.key,
                    sentinel.label,
                )
            except ResourceNotFoundError:
                value = _MISSING

            previous = self._last_values[sentinel]
            if previous is not _UNSET and previous != value:
                changed.add(sentinel)
            self._last_values[sentinel] = value

        if changed:
            logger.info("Sentinel change detected; refreshing all cached configuration")
            self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh(changed)
        return changed

    def run(
        self,
        stop_event: Event | None = None,
        *,
        max_polls: int | None = None,
    ) -> None:
        """Poll until stopped, or until max_polls is reached for finite demos."""
        if max_polls is not None and max_polls <= 0:
            raise ValueError("max_polls must be greater than zero")
        stop_event = stop_event or Event()
        polls = 0
        while not stop_event.is_set():
            self.poll_once()
            polls += 1
            if max_polls is not None and polls >= max_polls:
                return
            stop_event.wait(self._polling_interval)


AsyncRefreshCallback = Callable[[set[SentinelKey]], Awaitable[None] | None]


class AsyncConfigurationWatcher:
    """Poll sentinel keys and fully refresh the async cache after a change."""

    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinels: Sequence[SentinelKey],
        polling_interval: float,
        on_refresh: AsyncRefreshCallback | None = None,
    ) -> None:
        if not sentinels:
            raise ValueError("at least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinels = tuple(sentinels)
        self._polling_interval = polling_interval
        self._on_refresh = on_refresh
        self._last_values: dict[SentinelKey, object] = {
            sentinel: _UNSET for sentinel in sentinels
        }

    async def poll_once(self) -> set[SentinelKey]:
        """Check all sentinels once and return those whose values changed."""
        changed: set[SentinelKey] = set()
        for sentinel in self._sentinels:
            try:
                value: object = await self._configuration.get_setting(
                    sentinel.key,
                    sentinel.label,
                )
            except ResourceNotFoundError:
                value = _MISSING

            previous = self._last_values[sentinel]
            if previous is not _UNSET and previous != value:
                changed.add(sentinel)
            self._last_values[sentinel] = value

        if changed:
            logger.info("Sentinel change detected; refreshing all cached configuration")
            await self._configuration.refresh_all()
            if self._on_refresh is not None:
                callback_result = self._on_refresh(changed)
                if callback_result is not None:
                    await callback_result
        return changed

    async def run(
        self,
        stop_event: asyncio.Event | None = None,
        *,
        max_polls: int | None = None,
    ) -> None:
        """Poll until stopped, or until max_polls is reached for finite demos."""
        if max_polls is not None and max_polls <= 0:
            raise ValueError("max_polls must be greater than zero")
        stop_event = stop_event or asyncio.Event()
        polls = 0
        while not stop_event.is_set():
            await self.poll_once()
            polls += 1
            if max_polls is not None and polls >= max_polls:
                return
            try:
                await asyncio.wait_for(
                    stop_event.wait(),
                    timeout=self._polling_interval,
                )
            except TimeoutError:
                pass
