from __future__ import annotations

import asyncio
import time
from datetime import datetime

from azure.core.exceptions import ResourceExistsError, ResourceNotFoundError
from azure.keyvault.secrets import KeyVaultSecret, SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient


def _wait_until_purged(
    client: SecretClient,
    name: str,
    timeout: float,
    poll_interval: float,
) -> None:
    deadline = time.monotonic() + timeout
    while True:
        try:
            client.get_deleted_secret(name)
        except ResourceNotFoundError:
            return
        if time.monotonic() >= deadline:
            raise TimeoutError(f"Timed out waiting for secret {name!r} to be purged")
        time.sleep(poll_interval)


def _set_after_purge(
    client: SecretClient,
    name: str,
    value: str,
    expires_on: datetime,
    timeout: float,
    poll_interval: float,
) -> KeyVaultSecret:
    deadline = time.monotonic() + timeout
    while True:
        try:
            return client.set_secret(name, value, expires_on=expires_on)
        except ResourceExistsError:
            if time.monotonic() >= deadline:
                raise
            time.sleep(poll_interval)


def rotate_secret(
    client: SecretClient,
    name: str,
    value: str,
    expires_on: datetime,
    *,
    timeout: float = 300,
    poll_interval: float = 2,
) -> KeyVaultSecret:
    deleted_secret_exists = False
    try:
        delete_poller = client.begin_delete_secret(name)
        delete_poller.result(timeout=timeout)
        deleted_secret_exists = True
    except ResourceNotFoundError:
        try:
            client.get_deleted_secret(name)
            deleted_secret_exists = True
        except ResourceNotFoundError:
            pass

    if deleted_secret_exists:
        client.purge_deleted_secret(name)
        _wait_until_purged(client, name, timeout, poll_interval)

    return _set_after_purge(
        client,
        name,
        value,
        expires_on,
        timeout,
        poll_interval,
    )


async def _wait_until_purged_async(
    client: AsyncSecretClient,
    name: str,
    timeout: float,
    poll_interval: float,
) -> None:
    deadline = time.monotonic() + timeout
    while True:
        try:
            await client.get_deleted_secret(name)
        except ResourceNotFoundError:
            return
        if time.monotonic() >= deadline:
            raise TimeoutError(f"Timed out waiting for secret {name!r} to be purged")
        await asyncio.sleep(poll_interval)


async def _set_after_purge_async(
    client: AsyncSecretClient,
    name: str,
    value: str,
    expires_on: datetime,
    timeout: float,
    poll_interval: float,
) -> KeyVaultSecret:
    deadline = time.monotonic() + timeout
    while True:
        try:
            return await client.set_secret(name, value, expires_on=expires_on)
        except ResourceExistsError:
            if time.monotonic() >= deadline:
                raise
            await asyncio.sleep(poll_interval)


async def rotate_secret_async(
    client: AsyncSecretClient,
    name: str,
    value: str,
    expires_on: datetime,
    *,
    timeout: float = 300,
    poll_interval: float = 2,
) -> KeyVaultSecret:
    deleted_secret_exists = False
    try:
        delete_poller = await client.begin_delete_secret(name)
        await asyncio.wait_for(delete_poller.result(), timeout=timeout)
        deleted_secret_exists = True
    except ResourceNotFoundError:
        try:
            await client.get_deleted_secret(name)
            deleted_secret_exists = True
        except ResourceNotFoundError:
            pass

    if deleted_secret_exists:
        await client.purge_deleted_secret(name)
        await _wait_until_purged_async(client, name, timeout, poll_interval)

    return await _set_after_purge_async(
        client,
        name,
        value,
        expires_on,
        timeout,
        poll_interval,
    )
