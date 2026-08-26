from __future__ import annotations

import asyncio
import time
from datetime import datetime
from typing import Any, Protocol

from azure.core.exceptions import ResourceNotFoundError


class SyncRotationClient(Protocol):
    def begin_delete_secret(self, name: str) -> Any: ...

    def purge_deleted_secret(self, name: str) -> None: ...

    def get_deleted_secret(self, name: str) -> Any: ...

    def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime,
    ) -> Any: ...


class AsyncRotationClient(Protocol):
    async def delete_secret(self, name: str) -> Any: ...

    async def purge_deleted_secret(self, name: str) -> None: ...

    async def get_deleted_secret(self, name: str) -> Any: ...

    async def set_secret(
        self,
        name: str,
        value: str,
        *,
        expires_on: datetime,
    ) -> Any: ...


class SecretRotator:
    def __init__(
        self,
        client: SyncRotationClient,
        *,
        purge_timeout: float = 120.0,
        poll_interval: float = 2.0,
    ) -> None:
        if purge_timeout <= 0 or poll_interval <= 0:
            raise ValueError("purge_timeout and poll_interval must be positive")
        self._client = client
        self._purge_timeout = purge_timeout
        self._poll_interval = poll_interval

    def rotate(self, name: str, value: str, expires_on: datetime) -> Any:
        deleted = False
        try:
            delete_poller = self._client.begin_delete_secret(name)
            delete_poller.result()
            deleted = True
        except ResourceNotFoundError:
            pass

        if deleted:
            self._purge_after_delete(name)
        else:
            try:
                self._client.purge_deleted_secret(name)
            except ResourceNotFoundError:
                pass
            else:
                self._wait_until_purged(name)

        return self._client.set_secret(name, value, expires_on=expires_on)

    def _purge_after_delete(self, name: str) -> None:
        deadline = time.monotonic() + self._purge_timeout
        while True:
            try:
                self._client.purge_deleted_secret(name)
                break
            except ResourceNotFoundError:
                if time.monotonic() >= deadline:
                    raise TimeoutError(
                        f"Timed out waiting to purge secret {name!r}"
                    )
                time.sleep(self._poll_interval)
        self._wait_until_purged(name)

    def _wait_until_purged(self, name: str) -> None:
        deadline = time.monotonic() + self._purge_timeout
        while True:
            try:
                self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if time.monotonic() >= deadline:
                raise TimeoutError(
                    f"Timed out waiting for secret {name!r} to be purged"
                )
            time.sleep(self._poll_interval)


class AsyncSecretRotator:
    def __init__(
        self,
        client: AsyncRotationClient,
        *,
        purge_timeout: float = 120.0,
        poll_interval: float = 2.0,
    ) -> None:
        if purge_timeout <= 0 or poll_interval <= 0:
            raise ValueError("purge_timeout and poll_interval must be positive")
        self._client = client
        self._purge_timeout = purge_timeout
        self._poll_interval = poll_interval

    async def rotate(
        self,
        name: str,
        value: str,
        expires_on: datetime,
    ) -> Any:
        deleted = False
        try:
            await self._client.delete_secret(name)
            deleted = True
        except ResourceNotFoundError:
            pass

        if deleted:
            await self._purge_after_delete(name)
        else:
            try:
                await self._client.purge_deleted_secret(name)
            except ResourceNotFoundError:
                pass
            else:
                await self._wait_until_purged(name)

        return await self._client.set_secret(
            name,
            value,
            expires_on=expires_on,
        )

    async def _purge_after_delete(self, name: str) -> None:
        loop = asyncio.get_running_loop()
        deadline = loop.time() + self._purge_timeout
        while True:
            try:
                await self._client.purge_deleted_secret(name)
                break
            except ResourceNotFoundError:
                if loop.time() >= deadline:
                    raise TimeoutError(
                        f"Timed out waiting to purge secret {name!r}"
                    )
                await asyncio.sleep(self._poll_interval)
        await self._wait_until_purged(name)

    async def _wait_until_purged(self, name: str) -> None:
        loop = asyncio.get_running_loop()
        deadline = loop.time() + self._purge_timeout
        while True:
            try:
                await self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if loop.time() >= deadline:
                raise TimeoutError(
                    f"Timed out waiting for secret {name!r} to be purged"
                )
            await asyncio.sleep(self._poll_interval)
