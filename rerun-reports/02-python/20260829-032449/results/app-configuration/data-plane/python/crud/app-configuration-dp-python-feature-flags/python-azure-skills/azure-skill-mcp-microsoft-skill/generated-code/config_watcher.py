from __future__ import annotations

import asyncio
import logging
import threading
from collections.abc import Sequence

from azure.core.exceptions import AzureError

from config_service import AsyncConfigurationService, ConfigurationService


logger = logging.getLogger(__name__)
Sentinel = tuple[str, str | None]


def _normalize_sentinels(
    sentinels: Sequence[str | Sentinel],
) -> tuple[Sentinel, ...]:
    return tuple(
        sentinel if isinstance(sentinel, tuple) else (sentinel, None)
        for sentinel in sentinels
    )


class ConfigurationWatcher:
    def __init__(
        self,
        configuration: ConfigurationService,
        sentinels: Sequence[str | Sentinel],
        polling_interval: float = 30.0,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        if not sentinels:
            raise ValueError("at least one sentinel key is required")
        self._configuration = configuration
        self._sentinels = _normalize_sentinels(sentinels)
        self._polling_interval = polling_interval
        self._values: dict[Sentinel, str | None] | None = None

    def poll_once(self) -> bool:
        current = {
            sentinel: self._configuration.get_setting(*sentinel)
            for sentinel in self._sentinels
        }
        changed = self._values is not None and current != self._values
        self._values = current
        if changed:
            logger.info("Sentinel changed; refreshing all cached configuration")
            self._configuration.refresh_all()
        return changed

    def run(self, stop_event: threading.Event) -> None:
        while not stop_event.is_set():
            try:
                self.poll_once()
            except AzureError:
                logger.exception("Unable to poll Azure App Configuration")
            stop_event.wait(self._polling_interval)


class AsyncConfigurationWatcher:
    def __init__(
        self,
        configuration: AsyncConfigurationService,
        sentinels: Sequence[str | Sentinel],
        polling_interval: float = 30.0,
    ) -> None:
        if polling_interval <= 0:
            raise ValueError("polling_interval must be greater than zero")
        if not sentinels:
            raise ValueError("at least one sentinel key is required")
        self._configuration = configuration
        self._sentinels = _normalize_sentinels(sentinels)
        self._polling_interval = polling_interval
        self._values: dict[Sentinel, str | None] | None = None

    async def poll_once(self) -> bool:
        current = {
            sentinel: await self._configuration.get_setting(*sentinel)
            for sentinel in self._sentinels
        }
        changed = self._values is not None and current != self._values
        self._values = current
        if changed:
            logger.info("Sentinel changed; refreshing all cached configuration")
            await self._configuration.refresh_all()
        return changed

    async def run(self, stop_event: asyncio.Event) -> None:
        while not stop_event.is_set():
            try:
                await self.poll_once()
            except AzureError:
                logger.exception("Unable to poll Azure App Configuration")
            try:
                await asyncio.wait_for(
                    stop_event.wait(), timeout=self._polling_interval
                )
            except TimeoutError:
                pass
