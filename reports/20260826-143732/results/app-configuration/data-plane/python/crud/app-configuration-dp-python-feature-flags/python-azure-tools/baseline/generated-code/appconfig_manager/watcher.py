"""Sentinel-based configuration watchers."""

from __future__ import annotations

import asyncio
import inspect
from collections.abc import Awaitable, Callable, Sequence
from threading import Event, Thread
from typing import Any

from .configuration import AsyncConfigurationService, ConfigurationService


class ConfigurationWatcher:
    """Poll sentinel keys in a background thread and refresh on changes."""

    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        on_refresh: Callable[[], None] | None = None,
        label: str | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._on_refresh = on_refresh
        self._label = label
        self._stop_event = Event()
        self._thread: Thread | None = None

    def start(self) -> None:
        """Prime sentinels and start polling in a daemon thread."""
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Configuration watcher is already running")
        for key in self._sentinel_keys:
            self._configuration.get_setting(key, self._label)
        self._stop_event.clear()
        self._thread = Thread(target=self.run, name="appconfig-watcher", daemon=True)
        self._thread.start()

    def run(self) -> None:
        """Poll until stopped."""
        while not self._stop_event.wait(self._polling_interval):
            changed = any(
                self._configuration.check_for_update(key, self._label)
                for key in self._sentinel_keys
            )
            if changed:
                self._configuration.refresh_all()
                if self._on_refresh is not None:
                    self._on_refresh()

    def stop(self, timeout: float | None = None) -> None:
        """Stop polling and wait for the background thread."""
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join(timeout)
            if self._thread.is_alive():
                raise TimeoutError("Configuration watcher did not stop in time")
            self._thread = None

    def __enter__(self) -> ConfigurationWatcher:
        self.start()
        return self

    def __exit__(self, *args: Any) -> None:
        self.stop()


class AsyncConfigurationWatcher:
    """Poll sentinel keys in an asyncio task and refresh on changes."""

    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        on_refresh: Callable[[], None | Awaitable[None]] | None = None,
        label: str | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._on_refresh = on_refresh
        self._label = label
        self._stop_event = asyncio.Event()
        self._task: asyncio.Task[None] | None = None

    async def start(self) -> None:
        """Prime sentinels and start polling in an asyncio task."""
        if self._task is not None and not self._task.done():
            raise RuntimeError("Configuration watcher is already running")
        for key in self._sentinel_keys:
            await self._configuration.get_setting(key, self._label)
        self._stop_event.clear()
        self._task = asyncio.create_task(self.run(), name="appconfig-watcher")

    async def run(self) -> None:
        """Poll until stopped."""
        while True:
            try:
                await asyncio.wait_for(
                    self._stop_event.wait(), timeout=self._polling_interval
                )
                return
            except TimeoutError:
                pass

            changed = False
            for key in self._sentinel_keys:
                if await self._configuration.check_for_update(key, self._label):
                    changed = True
            if changed:
                await self._configuration.refresh_all()
                if self._on_refresh is not None:
                    result = self._on_refresh()
                    if inspect.isawaitable(result):
                        await result

    async def stop(self) -> None:
        """Stop polling and wait for the asyncio task."""
        self._stop_event.set()
        if self._task is not None:
            await self._task
            self._task = None

    async def __aenter__(self) -> AsyncConfigurationWatcher:
        await self.start()
        return self

    async def __aexit__(self, *args: Any) -> None:
        await self.stop()

