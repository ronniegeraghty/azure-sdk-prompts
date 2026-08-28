from __future__ import annotations

import asyncio
import inspect
import logging
from collections.abc import Awaitable, Callable, Sequence
from threading import Event, Thread
from typing import TypeAlias

from azure.core.exceptions import ResourceNotFoundError

from config_service import AsyncConfigurationService, ConfigurationService


logger = logging.getLogger(__name__)
_MISSING = object()
SentinelValue: TypeAlias = str | None | object
SyncCallback: TypeAlias = Callable[[list[str]], None]
AsyncCallback: TypeAlias = Callable[[list[str]], None | Awaitable[None]]


class ConfigurationWatcher:
    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        *,
        label: str | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._values: dict[str, SentinelValue] | None = None
        self._stop_event = Event()
        self._thread: Thread | None = None

    def poll_once(self, callback: SyncCallback | None = None) -> list[str]:
        current = {key: self._read_sentinel(key) for key in self._sentinel_keys}
        if self._values is None:
            self._values = current
            return []

        changed = [key for key in self._sentinel_keys if current[key] != self._values[key]]
        self._values = current
        if changed:
            self._configuration.refresh_all()
            if callback is not None:
                callback(changed)
        return changed

    def start(self, callback: SyncCallback | None = None) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Configuration watcher is already running")
        self._stop_event.clear()
        self._thread = Thread(
            target=self._run,
            args=(callback,),
            name="app-configuration-watcher",
            daemon=True,
        )
        self._thread.start()

    def stop(self) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join()
            self._thread = None

    def _run(self, callback: SyncCallback | None) -> None:
        while not self._stop_event.is_set():
            try:
                self.poll_once(callback)
            except Exception:
                logger.exception("Configuration watcher polling failed")
            self._stop_event.wait(self._polling_interval)

    def _read_sentinel(self, key: str) -> SentinelValue:
        try:
            return self._configuration.get_setting(key, self._label)
        except ResourceNotFoundError:
            return _MISSING


class AsyncConfigurationWatcher:
    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float = 30.0,
        *,
        label: str | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._values: dict[str, SentinelValue] | None = None
        self._stop_event = asyncio.Event()

    async def poll_once(self, callback: AsyncCallback | None = None) -> list[str]:
        current = {
            key: await self._read_sentinel(key) for key in self._sentinel_keys
        }
        if self._values is None:
            self._values = current
            return []

        changed = [key for key in self._sentinel_keys if current[key] != self._values[key]]
        self._values = current
        if changed:
            await self._configuration.refresh_all()
            if callback is not None:
                callback_result = callback(changed)
                if inspect.isawaitable(callback_result):
                    await callback_result
        return changed

    async def run(self, callback: AsyncCallback | None = None) -> None:
        self._stop_event.clear()
        while not self._stop_event.is_set():
            try:
                await self.poll_once(callback)
            except asyncio.CancelledError:
                raise
            except Exception:
                logger.exception("Async configuration watcher polling failed")
            try:
                await asyncio.wait_for(
                    self._stop_event.wait(), timeout=self._polling_interval
                )
            except TimeoutError:
                pass

    def stop(self) -> None:
        self._stop_event.set()

    async def _read_sentinel(self, key: str) -> SentinelValue:
        try:
            return await self._configuration.get_setting(key, self._label)
        except ResourceNotFoundError:
            return _MISSING
