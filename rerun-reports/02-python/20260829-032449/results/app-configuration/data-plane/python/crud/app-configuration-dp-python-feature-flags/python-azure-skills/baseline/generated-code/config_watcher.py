"""Sentinel-based configuration watchers."""

from __future__ import annotations

import asyncio
import logging
from collections.abc import Callable, Sequence
from threading import Event, Thread
from typing import Protocol

logger = logging.getLogger(__name__)


class RefreshableConfiguration(Protocol):
    def get_setting(self, key: str, label: str | None = None) -> str | None: ...

    def refresh_all(self) -> None: ...


class AsyncRefreshableConfiguration(Protocol):
    async def get_setting(
        self, key: str, label: str | None = None
    ) -> str | None: ...

    async def refresh_all(self) -> None: ...


class ConfigurationWatcher:
    """Poll sentinel keys and refresh all known configuration when one changes."""

    def __init__(
        self,
        configuration: RefreshableConfiguration,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        *,
        label: str | None = None,
        on_refresh: Callable[[set[str]], None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be positive")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: dict[str, str | None] | None = None
        self._stop_event = Event()
        self._thread: Thread | None = None

    def poll_once(self) -> set[str]:
        current = {
            key: self._configuration.get_setting(key, self._label)
            for key in self._sentinel_keys
        }
        if self._values is None:
            self._values = current
            return set()

        changed = {
            key for key, value in current.items() if value != self._values[key]
        }
        self._values = current
        if changed:
            self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh(changed)
        return changed

    def start(self) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Watcher is already running")
        self._stop_event.clear()
        self.poll_once()
        self._thread = Thread(
            target=self._run,
            name="app-configuration-watcher",
            daemon=True,
        )
        self._thread.start()

    def _run(self) -> None:
        while not self._stop_event.wait(self._polling_interval):
            try:
                self.poll_once()
            except Exception:
                logger.exception("Configuration sentinel poll failed")

    def stop(self, timeout: float | None = None) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join(timeout)
            self._thread = None


class AsyncConfigurationWatcher:
    """Async sentinel watcher with the same behavior as the sync watcher."""

    def __init__(
        self,
        configuration: AsyncRefreshableConfiguration,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        *,
        label: str | None = None,
        on_refresh: Callable[[set[str]], None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be positive")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: dict[str, str | None] | None = None
        self._stop_event = asyncio.Event()
        self._task: asyncio.Task[None] | None = None

    async def poll_once(self) -> set[str]:
        current = {
            key: await self._configuration.get_setting(key, self._label)
            for key in self._sentinel_keys
        }
        if self._values is None:
            self._values = current
            return set()

        changed = {
            key for key, value in current.items() if value != self._values[key]
        }
        self._values = current
        if changed:
            await self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh(changed)
        return changed

    async def start(self) -> None:
        if self._task is not None and not self._task.done():
            raise RuntimeError("Watcher is already running")
        self._stop_event.clear()
        await self.poll_once()
        self._task = asyncio.create_task(
            self._run(), name="app-configuration-watcher"
        )

    async def _run(self) -> None:
        while True:
            try:
                await asyncio.wait_for(
                    self._stop_event.wait(), timeout=self._polling_interval
                )
                return
            except TimeoutError:
                try:
                    await self.poll_once()
                except Exception:
                    logger.exception("Configuration sentinel poll failed")

    async def stop(self) -> None:
        self._stop_event.set()
        if self._task is not None:
            await self._task
            self._task = None
