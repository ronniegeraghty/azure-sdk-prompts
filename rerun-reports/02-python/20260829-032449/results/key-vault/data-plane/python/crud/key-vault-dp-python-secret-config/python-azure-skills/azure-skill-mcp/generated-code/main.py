"""Demonstrate synchronous and asynchronous Key Vault configuration."""

from __future__ import annotations

import asyncio
import os
from datetime import datetime, timedelta, timezone

from azure_keyvault_config import (
    AsyncSecretCache,
    AsyncSecretRotator,
    SecretCache,
    SecretRotator,
    create_async_provider,
    create_sync_provider,
)

REQUIRED_KEYS = ("database-url", "api-key", "feature-flags")
WARNING_WINDOW = timedelta(days=7)


def _print_cached_values(values: dict[str, str | None]) -> None:
    for name, value in values.items():
        status = "loaded" if value is not None else "missing"
        print(f"{name}: {status}")


def _print_expiry_warnings(expiring: dict[str, datetime]) -> None:
    for name, expires_on in expiring.items():
        print(f"WARNING: {name} expires at {expires_on.isoformat()}")


def run_sync_demo(rotation_value: str) -> None:
    print("Synchronous implementation")
    with create_sync_provider() as provider:
        cache = SecretCache(
            provider,
            required_keys=REQUIRED_KEYS,
            warning_window=WARNING_WINDOW,
        )
        cache.load_required()
        _print_cached_values({name: cache.get(name) for name in REQUIRED_KEYS})

        cache.refresh("api-key")
        cache.refresh_expiring()
        _print_expiry_warnings(cache.expiring_secrets())

        rotator = SecretRotator(provider.client)
        rotator.rotate(
            "api-key",
            rotation_value,
            datetime.now(timezone.utc) + timedelta(days=90),
        )
        cache.refresh("api-key")


async def run_async_demo(rotation_value: str) -> None:
    print("Asynchronous implementation")
    async with create_async_provider() as provider:
        cache = AsyncSecretCache(
            provider,
            required_keys=REQUIRED_KEYS,
            warning_window=WARNING_WINDOW,
        )
        await cache.load_required()
        values = {
            name: await cache.get(name)
            for name in REQUIRED_KEYS
        }
        _print_cached_values(values)

        await cache.refresh("api-key")
        await cache.refresh_expiring()
        _print_expiry_warnings(cache.expiring_secrets())

        rotator = AsyncSecretRotator(provider.client)
        await rotator.rotate(
            "api-key",
            rotation_value,
            datetime.now(timezone.utc) + timedelta(days=90),
        )
        await cache.refresh("api-key")


def main() -> None:
    rotation_value = os.environ.get("ROTATED_SECRET_VALUE")
    if not rotation_value:
        raise RuntimeError(
            "ROTATED_SECRET_VALUE must contain the demo's replacement secret value"
        )

    run_sync_demo(rotation_value)
    asyncio.run(run_async_demo(rotation_value))


if __name__ == "__main__":
    main()
