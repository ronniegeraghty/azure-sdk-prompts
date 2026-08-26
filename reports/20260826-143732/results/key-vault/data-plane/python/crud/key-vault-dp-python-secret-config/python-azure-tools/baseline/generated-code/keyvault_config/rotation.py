import asyncio
import time
from datetime import datetime

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


class SecretRotationError(RuntimeError):
    pass


class SecretRotator:
    """Delete, purge, and safely recreate a soft-deleted secret."""

    def __init__(
        self,
        client: SecretClient,
        timeout_seconds: float = 120.0,
        poll_interval_seconds: float = 2.0,
    ) -> None:
        self._client = client
        self._timeout_seconds = timeout_seconds
        self._poll_interval_seconds = poll_interval_seconds

    def rotate(self, name: str, value: str, expires_on: datetime):
        deadline = time.monotonic() + self._timeout_seconds
        delete_poller = self._client.begin_delete_secret(name)
        delete_poller.wait()

        try:
            self._client.purge_deleted_secret(name)
        except HttpResponseError as error:
            raise SecretRotationError(
                "The deleted secret could not be purged. Ensure purge permission "
                "is granted and purge protection is disabled."
            ) from error

        self._wait_until_purged(name, deadline)
        return self._create_after_purge(name, value, expires_on, deadline)

    def _wait_until_purged(self, name: str, deadline: float) -> None:
        while time.monotonic() < deadline:
            try:
                self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            time.sleep(self._poll_interval_seconds)
        raise SecretRotationError(f"Timed out waiting for {name!r} to be purged")

    def _create_after_purge(
        self, name: str, value: str, expires_on: datetime, deadline: float
    ):
        while True:
            try:
                return self._client.set_secret(
                    name, value, expires_on=expires_on
                )
            except HttpResponseError as error:
                if error.status_code != 409 or time.monotonic() >= deadline:
                    raise
                time.sleep(self._poll_interval_seconds)


class AsyncSecretRotator:
    """Asynchronous delete, purge, and recreate workflow."""

    def __init__(
        self,
        client: AsyncSecretClient,
        timeout_seconds: float = 120.0,
        poll_interval_seconds: float = 2.0,
    ) -> None:
        self._client = client
        self._timeout_seconds = timeout_seconds
        self._poll_interval_seconds = poll_interval_seconds

    async def rotate(self, name: str, value: str, expires_on: datetime):
        loop = asyncio.get_running_loop()
        deadline = loop.time() + self._timeout_seconds
        delete_poller = await self._client.begin_delete_secret(name)
        await delete_poller.wait()

        try:
            await self._client.purge_deleted_secret(name)
        except HttpResponseError as error:
            raise SecretRotationError(
                "The deleted secret could not be purged. Ensure purge permission "
                "is granted and purge protection is disabled."
            ) from error

        await self._wait_until_purged(name, deadline)
        return await self._create_after_purge(
            name, value, expires_on, deadline
        )

    async def _wait_until_purged(self, name: str, deadline: float) -> None:
        loop = asyncio.get_running_loop()
        while loop.time() < deadline:
            try:
                await self._client.get_deleted_secret(name)
            except ResourceNotFoundError:
                return
            await asyncio.sleep(self._poll_interval_seconds)
        raise SecretRotationError(f"Timed out waiting for {name!r} to be purged")

    async def _create_after_purge(
        self, name: str, value: str, expires_on: datetime, deadline: float
    ):
        loop = asyncio.get_running_loop()
        while True:
            try:
                return await self._client.set_secret(
                    name, value, expires_on=expires_on
                )
            except HttpResponseError as error:
                if error.status_code != 409 or loop.time() >= deadline:
                    raise
                await asyncio.sleep(self._poll_interval_seconds)
