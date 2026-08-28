from __future__ import annotations

import asyncio
import os
import secrets
from datetime import datetime, timedelta, timezone

from configuration import create_async_provider, create_sync_provider
from secret_cache import AsyncSecretCache, SyncSecretCache
from secret_rotation import rotate_secret, rotate_secret_async

WARNING_WINDOW = timedelta(days=7)
REQUIRED_KEYS = {
    "database-connection": None,
    "service-api-key": None,
    "feature-flag": "disabled",
}
ROTATION_SECRET_ENV = "ROTATION_SECRET_NAME"


def _describe(value: str | None) -> str:
    return "<missing>" if value is None else f"<loaded: {len(value)} characters>"


def run_sync_demo() -> None:
    print("Sync implementation")
    with create_sync_provider() as provider:
        cache = SyncSecretCache(provider, warning_window=WARNING_WINDOW)
        cache.bulk_load(REQUIRED_KEYS)
        for key in REQUIRED_KEYS:
            print(f"  {key}: {_describe(cache.get(key))}")

        refreshed_key = next(iter(REQUIRED_KEYS))
        cache.refresh(refreshed_key)
        print(f"  Refreshed {refreshed_key}")

        for record in cache.expiring_secrets():
            print(f"  WARNING: {record.name} expires on {record.expires_on}")

        rotation_name = os.getenv(ROTATION_SECRET_ENV, "demo-rotating-secret")
        rotate_secret(
            provider.client,
            rotation_name,
            secrets.token_urlsafe(32),
            datetime.now(timezone.utc) + timedelta(days=90),
        )
        print(f"  Rotated {rotation_name}")


async def run_async_demo() -> None:
    print("Async implementation")
    async with create_async_provider() as provider:
        cache = AsyncSecretCache(provider, warning_window=WARNING_WINDOW)
        await cache.bulk_load(REQUIRED_KEYS)
        for key in REQUIRED_KEYS:
            print(f"  {key}: {_describe(await cache.get(key))}")

        refreshed_key = next(iter(REQUIRED_KEYS))
        await cache.refresh(refreshed_key)
        print(f"  Refreshed {refreshed_key}")

        for record in cache.expiring_secrets():
            print(f"  WARNING: {record.name} expires on {record.expires_on}")

        rotation_name = os.getenv(ROTATION_SECRET_ENV, "demo-rotating-secret")
        await rotate_secret_async(
            provider.client,
            rotation_name,
            secrets.token_urlsafe(32),
            datetime.now(timezone.utc) + timedelta(days=90),
        )
        print(f"  Rotated {rotation_name}")


def main() -> None:
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
