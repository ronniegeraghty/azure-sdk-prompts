from __future__ import annotations

import asyncio
import time
from datetime import datetime, timezone
from typing import Any

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError


class SecretRotationError(RuntimeError):
    pass


def _utc_expiry(expires_on: datetime) -> datetime:
    if expires_on.tzinfo is None:
        return expires_on.replace(tzinfo=timezone.utc)
    return expires_on.astimezone(timezone.utc)


def rotate_secret(
    client: Any,
    name: str,
    new_value: str,
    expires_on: datetime,
    *,
    deletion_timeout: float = 120.0,
    poll_interval: float = 1.0,
) -> Any:
    try:
        delete_poller = client.begin_delete_secret(name)
        delete_poller.result(timeout=deletion_timeout)
    except ResourceNotFoundError:
        pass

    try:
        client.purge_deleted_secret(name)
    except ResourceNotFoundError:
        pass
    except HttpResponseError as exc:
        raise SecretRotationError(
            "The deleted secret could not be purged. Check purge permissions and "
            "whether purge protection is enabled."
        ) from exc

    deadline = time.monotonic() + deletion_timeout
    while True:
        try:
            client.get_deleted_secret(name)
        except ResourceNotFoundError:
            break
        if time.monotonic() >= deadline:
            raise TimeoutError(f"Timed out waiting for secret {name!r} to be purged")
        time.sleep(poll_interval)

    return client.set_secret(
        name,
        new_value,
        expires_on=_utc_expiry(expires_on),
    )


async def rotate_secret_async(
    client: Any,
    name: str,
    new_value: str,
    expires_on: datetime,
    *,
    deletion_timeout: float = 120.0,
    poll_interval: float = 1.0,
) -> Any:
    try:
        delete_poller = await client.begin_delete_secret(name)
        await delete_poller.result()
    except ResourceNotFoundError:
        pass

    try:
        await client.purge_deleted_secret(name)
    except ResourceNotFoundError:
        pass
    except HttpResponseError as exc:
        raise SecretRotationError(
            "The deleted secret could not be purged. Check purge permissions and "
            "whether purge protection is enabled."
        ) from exc

    deadline = asyncio.get_running_loop().time() + deletion_timeout
    while True:
        try:
            await client.get_deleted_secret(name)
        except ResourceNotFoundError:
            break
        if asyncio.get_running_loop().time() >= deadline:
            raise TimeoutError(f"Timed out waiting for secret {name!r} to be purged")
        await asyncio.sleep(poll_interval)

    return await client.set_secret(
        name,
        new_value,
        expires_on=_utc_expiry(expires_on),
    )
