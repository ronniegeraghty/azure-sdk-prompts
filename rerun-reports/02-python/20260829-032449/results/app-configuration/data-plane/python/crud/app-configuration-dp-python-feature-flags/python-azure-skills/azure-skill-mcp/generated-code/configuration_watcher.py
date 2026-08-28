from __future__ import annotations

import asyncio
import logging
from collections.abc import Callable, Sequence
from threading import Event, Thread
from typing import Optional

from azure.core.exceptions import AzureError

from configuration_service import AsyncConfigurationService, ConfigurationService


logger = logging.getLogger(__name__)
ChangeCallback = Callable[[list[str]], None]


class ConfigurationWatcher:
    """Poll sentinel keys and refresh cached selectors when their values change."""

    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        *,
        label: Optional[str] = None,
        on_change: Optional[ChangeCallback] = None,
    ) -> None:
        _validate_watcher_options(sentinel_keys, polling_interval)
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_change = on_change
        self._values: dict[str, str] = {}
        self._stop_event = Event()
        self._thread: Optional[Thread] = None

    def start(self) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("watcher is already running")
        self._stop_event.clear()
        self._thread = Thread(
            target=self.run,
            name="app-configuration-watcher",
            daemon=True,
        )
        self._thread.start()

    def run(self) -> None:
        while not self._stop_event.is_set():
            try:
                self._poll()
            except AzureError:
                logger.exception("Failed to poll Azure App Configuration sentinels")
            if self._stop_event.wait(self._polling_interval):
                return

    def stop(self) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join()
            self._thread = None

    def _poll(self) -> None:
        current = {
            key: self._configuration.get_setting(key, self._label)
            for key in self._sentinel_keys
        }
        changed = [
            key
            for key, value in current.items()
            if key in self._values and self._values[key] != value
        ]
        self._values = current

        if changed:
            self._configuration.refresh_all()
            if self._on_change is not None:
                self._on_change(changed)


class AsyncConfigurationWatcher:
    """Asynchronous sentinel watcher for Azure App Configuration."""

    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Sequence[str],
        polling_interval: float,
        *,
        label: Optional[str] = None,
        on_change: Optional[ChangeCallback] = None,
    ) -> None:
        _validate_watcher_options(sentinel_keys, polling_interval)
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        self._polling_interval = polling_interval
        self._label = label
        self._on_change = on_change
        self._values: dict[str, str] = {}
        self._stop_event = asyncio.Event()

    async def run(self) -> None:
        while True:
            try:
                await self._poll()
            except AzureError:
                logger.exception("Failed to poll Azure App Configuration sentinels")

            try:
                await asyncio.wait_for(
                    self._stop_event.wait(), timeout=self._polling_interval
                )
                return
            except TimeoutError:
                continue

    def stop(self) -> None:
        self._stop_event.set()

    async def _poll(self) -> None:
        current = {
            key: await self._configuration.get_setting(key, self._label)
            for key in self._sentinel_keys
        }
        changed = [
            key
            for key, value in current.items()
            if key in self._values and self._values[key] != value
        ]
        self._values = current

        if changed:
            await self._configuration.refresh_all()
            if self._on_change is not None:
                self._on_change(changed)


def _validate_watcher_options(
    sentinel_keys: Sequence[str], polling_interval: float
) -> None:
    if not sentinel_keys or any(not key for key in sentinel_keys):
        raise ValueError("sentinel_keys must contain at least one non-empty key")
    if polling_interval <= 0:
        raise ValueError("polling_interval must be greater than zero")
