"""Demonstrate synchronous and asynchronous Key Vault configuration."""

from __future__ import annotations

import asyncio
import os
from datetime import datetime, timedelta, timezone

from keyvault_config.cache import AsyncSecretCache, SecretCache
from keyvault_config.config import (
    create_async_key_vault_resources,
    create_key_vault_resources,
)
from keyvault_config.provider import AsyncSecretProvider, SecretProvider
from keyvault_config.rotation import rotate_secret, rotate_secret_async

REQUIRED_KEYS = ("database-url", "api-key", "feature-flags")
DEFAULTS = {"feature-flags": "{}"}
ROTATION_SECRET_ENV = "ROTATION_SECRET_NAME"
ROTATION_VALUE_ENV = "ROTATION_SECRET_VALUE"


def print_expiry_warnings(cache: SecretCache | AsyncSecretCache) -> None:
    for secret in cache.near_expiry():
        print(
            f"WARNING: {secret.name} expires at "
            f"{secret.expires_on.isoformat()}"
        )


def run_sync_demo() -> None:
    print("=== Synchronous Key Vault configuration ===")
    resources = create_key_vault_resources()
    try:
        provider = SecretProvider(resources.client)
        cache = SecretCache(provider, defaults=DEFAULTS)
        cache.load_required(REQUIRED_KEYS)

        for name in REQUIRED_KEYS:
            print(f"{name}: {cache.get(name)!r}")

        cache.refresh("api-key")
        print("Refreshed api-key")
        print_expiry_warnings(cache)

        rotation_name = os.getenv(ROTATION_SECRET_ENV)
        rotation_value = os.getenv(ROTATION_VALUE_ENV)
        if rotation_name and rotation_value:
            expires_on = datetime.now(timezone.utc) + timedelta(days=90)
            rotate_secret(
                resources.client, rotation_name, rotation_value, expires_on
            )
            cache.refresh(rotation_name)
            print(f"Rotated {rotation_name}")
        else:
            print(
                "Rotation skipped; set ROTATION_SECRET_NAME and "
                "ROTATION_SECRET_VALUE to enable it"
            )
    finally:
        resources.close()


async def run_async_demo() -> None:
    print("\n=== Asynchronous Key Vault configuration ===")
    resources = create_async_key_vault_resources()
    try:
        provider = AsyncSecretProvider(resources.client)
        cache = AsyncSecretCache(provider, defaults=DEFAULTS)
        await cache.load_required(REQUIRED_KEYS)

        for name in REQUIRED_KEYS:
            print(f"{name}: {await cache.get(name)!r}")

        await cache.refresh("api-key")
        print("Refreshed api-key")
        print_expiry_warnings(cache)

        rotation_name = os.getenv(ROTATION_SECRET_ENV)
        rotation_value = os.getenv(ROTATION_VALUE_ENV)
        if rotation_name and rotation_value:
            expires_on = datetime.now(timezone.utc) + timedelta(days=90)
            await rotate_secret_async(
                resources.client, rotation_name, rotation_value, expires_on
            )
            await cache.refresh(rotation_name)
            print(f"Rotated {rotation_name}")
        else:
            print(
                "Rotation skipped; set ROTATION_SECRET_NAME and "
                "ROTATION_SECRET_VALUE to enable it"
            )
    finally:
        await resources.close()


def main() -> None:
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
