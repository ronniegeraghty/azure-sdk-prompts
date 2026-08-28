from __future__ import annotations

import asyncio
import logging
from collections.abc import Awaitable, Callable, Sequence
from threading import Event, Thread

from azure.core.exceptions import ResourceNotFoundError

from config_service import AsyncConfigurationService, ConfigurationService


logger = logging.getLogger(__name__)
_MISSING = object()
SentinelValue = str | None | object


class ConfigurationWatcher:
    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        label: str | None = None,
        on_refresh: Callable[[], None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("Polling interval must be greater than zero")

        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: dict[str, SentinelValue] = {}
        self._stop_event = Event()
        self._thread: Thread | None = None

    def _read_sentinel(self, key: str) -> SentinelValue:
        try:
            return self._configuration.get_setting(key, self._label)
        except ResourceNotFoundError:
            return _MISSING

    def _poll(self) -> bool:
        changed = False
        for key in self._sentinel_keys:
            value = self._read_sentinel(key)
            previous = self._values.get(key, _MISSING)
            if key in self._values and value != previous:
                changed = True
            self._values[key] = value

        if changed:
            self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh()
        return changed

    def _run(self) -> None:
        try:
            self._poll()
            while not self._stop_event.wait(self._polling_interval):
                self._poll()
        except Exception:
            logger.exception("Configuration watcher stopped after a polling failure")

    def start(self) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Configuration watcher is already running")
        self._stop_event.clear()
        self._thread = Thread(
            target=self._run, name="configuration-watcher", daemon=True
        )
        self._thread.start()

    def stop(self) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join()
            self._thread = None

    def __enter__(self) -> ConfigurationWatcher:
        self.start()
        return self

    def __exit__(self, *args: object) -> None:
        self.stop()


class AsyncConfigurationWatcher:
    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        label: str | None = None,
        on_refresh: Callable[[], Awaitable[None] | None] | None = None,
    ) -> None:
        if not sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        if polling_interval <= 0:
            raise ValueError("Polling interval must be greater than zero")

        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: dict[str, SentinelValue] = {}
        self._stop_event = asyncio.Event()
        self._task: asyncio.Task[None] | None = None

    async def _read_sentinel(self, key: str) -> SentinelValue:
        try:
            return await self._configuration.get_setting(key, self._label)
        except ResourceNotFoundError:
            return _MISSING

    async def _poll(self) -> bool:
        changed = False
        for key in self._sentinel_keys:
            value = await self._read_sentinel(key)
            previous = self._values.get(key, _MISSING)
            if key in self._values and value != previous:
                changed = True
            self._values[key] = value

        if changed:
            await self._configuration.refresh_all()
            if self._on_refresh is not None:
                result = self._on_refresh()
                if result is not None:
                    await result
        return changed

    async def _run(self) -> None:
        try:
            await self._poll()
            while True:
                try:
                    await asyncio.wait_for(
                        self._stop_event.wait(), timeout=self._polling_interval
                    )
                    return
                except TimeoutError:
                    await self._poll()
        except Exception:
            logger.exception("Async configuration watcher stopped after a polling failure")

    def start(self) -> None:
        if self._task is not None and not self._task.done():
            raise RuntimeError("Async configuration watcher is already running")
        self._stop_event.clear()
        self._task = asyncio.create_task(
            self._run(), name="async-configuration-watcher"
        )

    async def stop(self) -> None:
        self._stop_event.set()
        if self._task is not None:
            await self._task
            self._task = None

    async def __aenter__(self) -> AsyncConfigurationWatcher:
        self.start()
        return self

    async def __aexit__(self, *args: object) -> None:
        await self.stop()

