"""Safe synchronous and asynchronous secret rotation helpers."""

from __future__ import annotations

import asyncio
import time
from datetime import datetime
from typing import Any

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError


def _purge_error(name: str, error: HttpResponseError) -> RuntimeError:
    return RuntimeError(
        f"Secret {name!r} was deleted but could not be purged. "
        "Key Vault must allow purge and the identity needs purge permission "
        "before the name can be reused."
    )


def rotate_secret(
    client: Any,
    name: str,
    value: str,
    expires_on: datetime,
    *,
    purge_timeout: float = 120.0,
    poll_interval: float = 2.0,
) -> Any:
    """Delete, purge, and recreate a secret after deletion fully completes."""
    if expires_on.tzinfo is None:
        raise ValueError("expires_on must be timezone-aware")

    deletion_poller = client.begin_delete_secret(name)
    deletion_poller.result()

    try:
        client.purge_deleted_secret(name)
    except HttpResponseError as error:
        raise _purge_error(name, error) from error

    deadline = time.monotonic() + purge_timeout
    while True:
        try:
            client.get_deleted_secret(name)
        except ResourceNotFoundError:
            break
        if time.monotonic() >= deadline:
            raise TimeoutError(
                f"Timed out waiting for deleted secret {name!r} to be purged"
            )
        time.sleep(poll_interval)

    return client.set_secret(name, value, expires_on=expires_on)


async def rotate_secret_async(
    client: Any,
    name: str,
    value: str,
    expires_on: datetime,
    *,
    purge_timeout: float = 120.0,
    poll_interval: float = 2.0,
) -> Any:
    """Async delete, purge, and recreate rotation."""
    if expires_on.tzinfo is None:
        raise ValueError("expires_on must be timezone-aware")

    deletion_poller = await client.begin_delete_secret(name)
    await deletion_poller.result()

    try:
        await client.purge_deleted_secret(name)
    except HttpResponseError as error:
        raise _purge_error(name, error) from error

    deadline = asyncio.get_running_loop().time() + purge_timeout
    while True:
        try:
            await client.get_deleted_secret(name)
        except ResourceNotFoundError:
            break
        if asyncio.get_running_loop().time() >= deadline:
            raise TimeoutError(
                f"Timed out waiting for deleted secret {name!r} to be purged"
            )
        await asyncio.sleep(poll_interval)

    return await client.set_secret(name, value, expires_on=expires_on)
