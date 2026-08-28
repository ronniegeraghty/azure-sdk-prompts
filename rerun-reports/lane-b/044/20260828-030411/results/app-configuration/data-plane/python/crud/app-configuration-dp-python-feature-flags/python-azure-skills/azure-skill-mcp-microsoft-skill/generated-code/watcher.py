from __future__ import annotations

import asyncio
import logging
from collections.abc import Awaitable, Callable, Sequence
from threading import Event, Thread

from config_service import AsyncConfigurationService, ConfigurationService


LOGGER = logging.getLogger(__name__)
Sentinel = tuple[str, str | None]


class ConfigurationWatcher:
    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Sequence[str | Sentinel],
        polling_interval: float = 30.0,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinels = [_normalize_sentinel(item) for item in sentinel_keys]
        if not self._sentinels:
            raise ValueError("At least one sentinel key is required")
        self._polling_interval = polling_interval
        self._stop_event = Event()
        self._thread: Thread | None = None
        self._values: dict[Sentinel, str | None] = {}

    def start(self, on_refresh: Callable[[], None] | None = None) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Configuration watcher is already running")
        self._stop_event.clear()
        self._values = {
            sentinel: self._configuration.get_setting(*sentinel)
            for sentinel in self._sentinels
        }
        self._thread = Thread(
            target=self._run, args=(on_refresh,), daemon=True, name="config-watcher"
        )
        self._thread.start()

    def stop(self) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join()
        self._thread = None

    def _run(self, on_refresh: Callable[[], None] | None) -> None:
        while not self._stop_event.wait(self._polling_interval):
            try:
                if self._sentinel_changed():
                    self._configuration.refresh_all()
                    if on_refresh is not None:
                        on_refresh()
            except Exception:
                LOGGER.exception("Configuration watcher poll failed")

    def _sentinel_changed(self) -> bool:
        changed = False
        for sentinel in self._sentinels:
            current = self._configuration.get_setting(*sentinel, force_refresh=True)
            if current != self._values[sentinel]:
                self._values[sentinel] = current
                changed = True
        return changed

    def __enter__(self) -> ConfigurationWatcher:
        self.start()
        return self

    def __exit__(self, *_: object) -> None:
        self.stop()


class AsyncConfigurationWatcher:
    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Sequence[str | Sentinel],
        polling_interval: float = 30.0,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinels = [_normalize_sentinel(item) for item in sentinel_keys]
        if not self._sentinels:
            raise ValueError("At least one sentinel key is required")
        self._polling_interval = polling_interval
        self._task: asyncio.Task[None] | None = None
        self._values: dict[Sentinel, str | None] = {}

    async def start(
        self, on_refresh: Callable[[], Awaitable[None] | None] | None = None
    ) -> None:
        if self._task is not None and not self._task.done():
            raise RuntimeError("Configuration watcher is already running")
        self._values = {
            sentinel: await self._configuration.get_setting(*sentinel)
            for sentinel in self._sentinels
        }
        self._task = asyncio.create_task(self._run(on_refresh))

    async def stop(self) -> None:
        if self._task is None:
            return
        self._task.cancel()
        try:
            await self._task
        except asyncio.CancelledError:
            pass
        self._task = None

    async def _run(
        self, on_refresh: Callable[[], Awaitable[None] | None] | None
    ) -> None:
        while True:
            await asyncio.sleep(self._polling_interval)
            try:
                if await self._sentinel_changed():
                    await self._configuration.refresh_all()
                    if on_refresh is not None:
                        result = on_refresh()
                        if result is not None:
                            await result
            except asyncio.CancelledError:
                raise
            except Exception:
                LOGGER.exception("Async configuration watcher poll failed")

    async def _sentinel_changed(self) -> bool:
        changed = False
        for sentinel in self._sentinels:
            current = await self._configuration.get_setting(
                *sentinel, force_refresh=True
            )
            if current != self._values[sentinel]:
                self._values[sentinel] = current
                changed = True
        return changed

    async def __aenter__(self) -> AsyncConfigurationWatcher:
        await self.start()
        return self

    async def __aexit__(self, *_: object) -> None:
        await self.stop()


def _normalize_sentinel(item: str | Sentinel) -> Sentinel:
    return (item, None) if isinstance(item, str) else item

