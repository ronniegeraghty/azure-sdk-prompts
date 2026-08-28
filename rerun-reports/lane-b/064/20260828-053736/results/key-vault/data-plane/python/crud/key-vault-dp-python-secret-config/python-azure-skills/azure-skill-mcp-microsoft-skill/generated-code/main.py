from __future__ import annotations

import asyncio
import os
from contextlib import asynccontextmanager, contextmanager
from datetime import datetime, timedelta, timezone
from typing import AsyncIterator, Iterator

from keyvault_config.async_cache import AsyncSecretCache
from keyvault_config.async_provider import AsyncSecretProvider
from keyvault_config.cache import SecretCache
from keyvault_config.demo_backend import (
    AsyncInMemorySecretClient,
    InMemorySecretClient,
)
from keyvault_config.factory import (
    open_async_secret_client,
    open_secret_client,
)
from keyvault_config.provider import SecretProvider
from keyvault_config.rotation import AsyncSecretRotator, SecretRotator

REQUIRED_CONFIG = {
    "database-url": None,
    "api-key": None,
    "feature-flag": "disabled",
}


def _seed_local_client() -> InMemorySecretClient:
    now = datetime.now(timezone.utc)
    client = InMemorySecretClient()
    client.set_secret(
        "database-url", "postgresql://localhost/app",
        expires_on=now + timedelta(days=30),
    )
    client.set_secret(
        "api-key", "local-demo-key",
        expires_on=now + timedelta(days=5),
    )
    return client


@contextmanager
def _sync_client() -> Iterator[object]:
    if os.getenv("DEMO_MODE", "local") == "azure":
        with open_secret_client() as client:
            yield client
    else:
        yield _seed_local_client()


@asynccontextmanager
async def _async_client() -> AsyncIterator[object]:
    if os.getenv("DEMO_MODE", "local") == "azure":
        async with open_async_secret_client() as client:
            yield client
    else:
        yield AsyncInMemorySecretClient(_seed_local_client())


def _rotation_value() -> str:
    if os.getenv("DEMO_MODE", "local") == "azure":
        value = os.getenv("DEMO_ROTATION_VALUE")
        if not value:
            raise RuntimeError(
                "DEMO_ROTATION_VALUE is required in Azure demo mode"
            )
        return value
    return "rotated-local-demo-key"


def run_sync_demo() -> None:
    print("Sync implementation")
    with _sync_client() as client:
        provider = SecretProvider(client)
        cache = SecretCache(provider)
        cache.load_required(REQUIRED_CONFIG)

        for key in REQUIRED_CONFIG:
            print(f"  {key}: configured={cache.get(key) is not None}")

        cache.refresh("database-url")
        expiring = cache.expiring_keys()
        if expiring:
            print(f"  Warning: near expiry: {', '.join(expiring)}")

        SecretRotator(client, poll_interval=0).rotate(
            "api-key",
            _rotation_value(),
            expires_on=datetime.now(timezone.utc) + timedelta(days=90),
        )
        cache.refresh("api-key")
        print("  api-key rotated and cache refreshed")


async def run_async_demo() -> None:
    print("Async implementation")
    async with _async_client() as client:
        provider = AsyncSecretProvider(client)
        cache = AsyncSecretCache(provider)
        await cache.load_required(REQUIRED_CONFIG)

        for key in REQUIRED_CONFIG:
            print(
                f"  {key}: configured={await cache.get(key) is not None}"
            )

        await cache.refresh("database-url")
        expiring = cache.expiring_keys()
        if expiring:
            print(f"  Warning: near expiry: {', '.join(expiring)}")

        await AsyncSecretRotator(client, poll_interval=0).rotate(
            "api-key",
            _rotation_value(),
            expires_on=datetime.now(timezone.utc) + timedelta(days=90),
        )
        await cache.refresh("api-key")
        print("  api-key rotated and cache refreshed")


def main() -> None:
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
