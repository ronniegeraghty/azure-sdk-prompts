from __future__ import annotations

import asyncio
import time
from datetime import datetime

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


class SecretRotationTimeoutError(TimeoutError):
    pass


class SecretRotator:
    def __init__(
        self,
        client: SecretClient,
        *,
        timeout: float = 120.0,
        poll_interval: float = 1.0,
    ) -> None:
        self._client = client
        self._timeout = timeout
        self._poll_interval = poll_interval

    def rotate(
        self, name: str, value: str, *, expires_on: datetime
    ) -> object:
        delete_poller = self._client.begin_delete_secret(name)
        delete_poller.result(timeout=self._timeout)

        self._client.purge_deleted_secret(name)
        self._wait_until_purged(name)
        return self._client.set_secret(
            name, value, expires_on=expires_on
        )

    def _wait_until_purged(self, name: str) -> None:
        deadline = time.monotonic() + self._timeout
        while True:
            try:
                self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if time.monotonic() >= deadline:
                raise SecretRotationTimeoutError(
                    f"Timed out waiting for secret {name!r} to be purged"
                )
            time.sleep(self._poll_interval)


class AsyncSecretRotator:
    def __init__(
        self,
        client: AsyncSecretClient,
        *,
        timeout: float = 120.0,
        poll_interval: float = 1.0,
    ) -> None:
        self._client = client
        self._timeout = timeout
        self._poll_interval = poll_interval

    async def rotate(
        self, name: str, value: str, *, expires_on: datetime
    ) -> object:
        delete_poller = await self._client.begin_delete_secret(name)
        await asyncio.wait_for(delete_poller.result(), timeout=self._timeout)

        await self._client.purge_deleted_secret(name)
        await self._wait_until_purged(name)
        return await self._client.set_secret(
            name, value, expires_on=expires_on
        )

    async def _wait_until_purged(self, name: str) -> None:
        deadline = asyncio.get_running_loop().time() + self._timeout
        while True:
            try:
                await self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if asyncio.get_running_loop().time() >= deadline:
                raise SecretRotationTimeoutError(
                    f"Timed out waiting for secret {name!r} to be purged"
                )
            await asyncio.sleep(self._poll_interval)
