from __future__ import annotations

import asyncio
from collections.abc import Awaitable, Callable, Sequence
from threading import Event

from config_service import AsyncConfigurationService, ConfigurationService


class ConfigurationWatcher:
    def __init__(
        self,
        config: ConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        *,
        label: str | None = None,
        on_refresh: Callable[[], None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._config = config
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._stop = Event()

    def run(self, max_polls: int | None = None) -> None:
        previous = {
            key: self._config.get_setting(key, self._label) for key in self._sentinel_keys
        }
        polls = 0
        while max_polls is None or polls < max_polls:
            if self._stop.wait(self._polling_interval):
                break
            current = {
                key: self._config.get_setting(key, self._label) for key in self._sentinel_keys
            }
            polls += 1
            if current != previous:
                self._config.refresh_all()
                if self._on_refresh:
                    self._on_refresh()
                previous = current

    def stop(self) -> None:
        self._stop.set()


class AsyncConfigurationWatcher:
    def __init__(
        self,
        config: AsyncConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        *,
        label: str | None = None,
        on_refresh: Callable[[], Awaitable[None] | None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._config = config
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._stop = asyncio.Event()

    async def run(self, max_polls: int | None = None) -> None:
        previous = {
            key: await self._config.get_setting(key, self._label)
            for key in self._sentinel_keys
        }
        polls = 0
        while max_polls is None or polls < max_polls:
            try:
                await asyncio.wait_for(self._stop.wait(), timeout=self._polling_interval)
                break
            except TimeoutError:
                pass

            current = {
                key: await self._config.get_setting(key, self._label)
                for key in self._sentinel_keys
            }
            polls += 1
            if current != previous:
                await self._config.refresh_all()
                if self._on_refresh:
                    result = self._on_refresh()
                    if result is not None:
                        await result
                previous = current

    def stop(self) -> None:
        self._stop.set()
