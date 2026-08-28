from __future__ import annotations

import asyncio
from datetime import datetime, timedelta, timezone
from uuid import uuid4

from keyvault_config.cache import AsyncSecretCache, SecretCache
from keyvault_config.factory import create_async_provider, create_sync_provider
from keyvault_config.rotation import rotate_secret, rotate_secret_async

REQUIRED_CONFIG = {
    "database-url": "sqlite:///local.db",
    "external-api-key": None,
    "feature-flag": "disabled",
}
ROTATING_SECRET = "demo-rotating-secret"
WARNING_WINDOW = timedelta(days=7)


def print_cache_status(label: str, values: dict[str, str | None]) -> None:
    print(label)
    for name, value in values.items():
        print(f"  {name}: {'available' if value is not None else 'missing'}")


def warn_about_expiry(cache: SecretCache | AsyncSecretCache) -> None:
    for secret in cache.expiring_secrets():
        print(f"WARNING: {secret.name} expires at {secret.expires_on}")


def run_sync_demo() -> None:
    print("Running synchronous demo")
    with create_sync_provider() as provider:
        cache = SecretCache(provider, WARNING_WINDOW)
        loaded = dict(cache.load_required(REQUIRED_CONFIG))
        print_cache_status("Startup configuration:", loaded)

        cached = {name: cache.get(name) for name in REQUIRED_CONFIG}
        print_cache_status("Read from cache:", cached)

        cache.refresh("feature-flag")
        cache.refresh_expiring()
        warn_about_expiry(cache)

        expires_on = datetime.now(timezone.utc) + timedelta(days=90)
        rotate_secret(
            provider.client,
            ROTATING_SECRET,
            f"sync-{uuid4()}",
            expires_on,
        )
        cache.refresh(ROTATING_SECRET)
        print(f"Rotated {ROTATING_SECRET} with the synchronous client")


async def run_async_demo() -> None:
    print("Running asynchronous demo")
    async with create_async_provider() as provider:
        cache = AsyncSecretCache(provider, WARNING_WINDOW)
        loaded = dict(await cache.load_required(REQUIRED_CONFIG))
        print_cache_status("Startup configuration:", loaded)

        cached = {name: await cache.get(name) for name in REQUIRED_CONFIG}
        print_cache_status("Read from cache:", cached)

        await cache.refresh("feature-flag")
        await cache.refresh_expiring()
        warn_about_expiry(cache)

        expires_on = datetime.now(timezone.utc) + timedelta(days=90)
        await rotate_secret_async(
            provider.client,
            ROTATING_SECRET,
            f"async-{uuid4()}",
            expires_on,
        )
        await cache.refresh(ROTATING_SECRET)
        print(f"Rotated {ROTATING_SECRET} with the asynchronous client")


def main() -> None:
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
