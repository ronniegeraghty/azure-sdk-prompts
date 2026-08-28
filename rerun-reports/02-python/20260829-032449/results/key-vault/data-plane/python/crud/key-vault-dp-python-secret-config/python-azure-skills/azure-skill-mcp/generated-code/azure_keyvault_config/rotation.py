"""Safe secret rotation helpers."""

from __future__ import annotations

import asyncio
import time
from datetime import datetime, timezone

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets import KeyVaultSecret, SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


def _validate_expiry(expires_on: datetime) -> None:
    normalized = (
        expires_on.replace(tzinfo=timezone.utc)
        if expires_on.tzinfo is None
        else expires_on.astimezone(timezone.utc)
    )
    if normalized <= datetime.now(timezone.utc):
        raise ValueError("expires_on must be in the future")


class SecretRotator:
    """Delete, fully remove, and recreate a secret with the same name."""

    def __init__(
        self,
        client: SecretClient,
        purge_timeout: float = 120.0,
        polling_interval: float = 1.0,
    ) -> None:
        if purge_timeout <= 0 or polling_interval <= 0:
            raise ValueError("timeouts and polling intervals must be positive")
        self._client = client
        self._purge_timeout = purge_timeout
        self._polling_interval = polling_interval

    def rotate(
        self, name: str, value: str, expires_on: datetime
    ) -> KeyVaultSecret:
        _validate_expiry(expires_on)
        try:
            poller = self._client.begin_delete_secret(name)
            poller.wait()
            deleted = poller.result()
        except ResourceNotFoundError:
            deleted = None

        if deleted is not None and deleted.recovery_id is not None:
            self._client.purge_deleted_secret(name)
            self._wait_until_purged(name)

        return self._client.set_secret(name, value, expires_on=expires_on)

    def _wait_until_purged(self, name: str) -> None:
        deadline = time.monotonic() + self._purge_timeout
        while True:
            try:
                self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if time.monotonic() >= deadline:
                raise TimeoutError(
                    f"Secret {name!r} was not purged within "
                    f"{self._purge_timeout:.1f} seconds"
                )
            time.sleep(self._polling_interval)


class AsyncSecretRotator:
    """Asynchronously delete, fully remove, and recreate a secret."""

    def __init__(
        self,
        client: AsyncSecretClient,
        purge_timeout: float = 120.0,
        polling_interval: float = 1.0,
    ) -> None:
        if purge_timeout <= 0 or polling_interval <= 0:
            raise ValueError("timeouts and polling intervals must be positive")
        self._client = client
        self._purge_timeout = purge_timeout
        self._polling_interval = polling_interval

    async def rotate(
        self, name: str, value: str, expires_on: datetime
    ) -> KeyVaultSecret:
        _validate_expiry(expires_on)
        try:
            # The aio client runs its deletion poller internally before returning.
            deleted = await self._client.delete_secret(name)
        except ResourceNotFoundError:
            deleted = None

        if deleted is not None and deleted.recovery_id is not None:
            await self._client.purge_deleted_secret(name)
            await self._wait_until_purged(name)

        return await self._client.set_secret(name, value, expires_on=expires_on)

    async def _wait_until_purged(self, name: str) -> None:
        deadline = time.monotonic() + self._purge_timeout
        while True:
            try:
                await self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            if time.monotonic() >= deadline:
                raise TimeoutError(
                    f"Secret {name!r} was not purged within "
                    f"{self._purge_timeout:.1f} seconds"
                )
            await asyncio.sleep(self._polling_interval)
