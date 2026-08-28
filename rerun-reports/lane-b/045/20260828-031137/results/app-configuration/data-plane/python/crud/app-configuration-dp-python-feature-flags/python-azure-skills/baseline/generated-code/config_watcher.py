from __future__ import annotations

import asyncio
import logging
from threading import Event, Thread
from typing import Callable, Dict, Iterable, List, Optional

from azure.core.exceptions import AzureError, ResourceNotFoundError

from config_service import AsyncConfigurationService, ConfigurationService


logger = logging.getLogger(__name__)
ChangeCallback = Callable[[List[str]], None]


class ConfigurationWatcher:
    def __init__(
        self,
        configuration: ConfigurationService,
        sentinel_keys: Iterable[str],
        polling_interval: float,
        *,
        label: Optional[str] = None,
        on_refresh: Optional[ChangeCallback] = None,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        if not self._sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: Dict[str, Optional[str]] = {}
        self._stop_event = Event()
        self._thread: Optional[Thread] = None

    def poll_once(self) -> List[str]:
        changed: List[str] = []
        for key in self._sentinel_keys:
            try:
                value = self._configuration.get_setting(key, self._label)
            except ResourceNotFoundError:
                value = None
            if key in self._values and self._values[key] != value:
                changed.append(key)
            self._values[key] = value

        if changed:
            self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh(changed)
        return changed

    def start(self) -> None:
        if self._thread is not None and self._thread.is_alive():
            raise RuntimeError("Configuration watcher is already running")
        self._stop_event.clear()
        self._thread = Thread(target=self._run, name="config-watcher", daemon=True)
        self._thread.start()

    def stop(self) -> None:
        self._stop_event.set()
        if self._thread is not None:
            self._thread.join()
            self._thread = None

    def _run(self) -> None:
        while not self._stop_event.is_set():
            try:
                self.poll_once()
            except AzureError:
                logger.exception("Azure App Configuration sentinel poll failed")
            self._stop_event.wait(self._polling_interval)


class AsyncConfigurationWatcher:
    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinel_keys: Iterable[str],
        polling_interval: float,
        *,
        label: Optional[str] = None,
        on_refresh: Optional[ChangeCallback] = None,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        self._configuration = configuration
        self._sentinel_keys = tuple(sentinel_keys)
        if not self._sentinel_keys:
            raise ValueError("At least one sentinel key is required")
        self._polling_interval = polling_interval
        self._label = label
        self._on_refresh = on_refresh
        self._values: Dict[str, Optional[str]] = {}
        self._stop_event = asyncio.Event()
        self._task: Optional[asyncio.Task[None]] = None

    async def poll_once(self) -> List[str]:
        changed: List[str] = []
        for key in self._sentinel_keys:
            try:
                value = await self._configuration.get_setting(key, self._label)
            except ResourceNotFoundError:
                value = None
            if key in self._values and self._values[key] != value:
                changed.append(key)
            self._values[key] = value

        if changed:
            await self._configuration.refresh_all()
            if self._on_refresh is not None:
                self._on_refresh(changed)
        return changed

    def start(self) -> None:
        if self._task is not None and not self._task.done():
            raise RuntimeError("Configuration watcher is already running")
        self._stop_event.clear()
        self._task = asyncio.create_task(self._run(), name="config-watcher")

    async def stop(self) -> None:
        self._stop_event.set()
        if self._task is not None:
            await self._task
            self._task = None

    async def _run(self) -> None:
        while not self._stop_event.is_set():
            try:
                await self.poll_once()
            except AzureError:
                logger.exception("Azure App Configuration sentinel poll failed")
            try:
                await asyncio.wait_for(
                    self._stop_event.wait(),
                    timeout=self._polling_interval,
                )
            except asyncio.TimeoutError:
                pass
