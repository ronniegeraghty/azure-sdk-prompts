import asyncio
import os
from datetime import datetime, timedelta, timezone
from typing import Optional

from keyvault_config.cache import AsyncSecretCache, SecretCache
from keyvault_config.factory import (
    create_async_configuration,
    create_sync_configuration,
)
from keyvault_config.rotation import AsyncSecretRotator, SecretRotator

DEFAULT_CONFIG_KEYS = ("database-url", "api-key", "feature-flags")


def _required_keys():
    configured = os.environ.get("REQUIRED_CONFIG_KEYS")
    return (
        tuple(key.strip() for key in configured.split(",") if key.strip())
        if configured
        else DEFAULT_CONFIG_KEYS
    )


def _rotation_settings():
    name = os.environ.get("ROTATION_SECRET_NAME", "api-key")
    value = os.environ.get("ROTATED_SECRET_VALUE")
    if value is None:
        raise RuntimeError("ROTATED_SECRET_VALUE must contain the new value")
    expires_on = datetime.now(timezone.utc) + timedelta(days=90)
    return name, value, expires_on


def _display(name: str, value: Optional[str]) -> None:
    state = "<missing>" if value is None else "<loaded>"
    print(f"{name}: {state}")


def run_sync_demo() -> None:
    print("Sync Key Vault configuration demo")
    rotation_name, rotation_value, expires_on = _rotation_settings()
    with create_sync_configuration() as configuration:
        cache = SecretCache(configuration.provider)
        loaded = cache.load_required(_required_keys())
        for name, value in loaded.items():
            _display(name, value)

        for name in _required_keys():
            _display(f"cached {name}", cache.get(name))

        cache.refresh(_required_keys()[0])
        cache.refresh_expiring()
        for name, info in cache.expiring_secrets().items():
            print(f"WARNING: {name} expires at {info.expires_on}")

        SecretRotator(configuration.client).rotate(
            rotation_name, rotation_value, expires_on
        )
        cache.refresh(rotation_name)
        print(f"Rotated {rotation_name}")


async def run_async_demo() -> None:
    print("Async Key Vault configuration demo")
    rotation_name, rotation_value, expires_on = _rotation_settings()
    async with create_async_configuration() as configuration:
        cache = AsyncSecretCache(configuration.provider)
        loaded = await cache.load_required(_required_keys())
        for name, value in loaded.items():
            _display(name, value)

        for name in _required_keys():
            _display(f"cached {name}", await cache.get(name))

        await cache.refresh(_required_keys()[0])
        await cache.refresh_expiring()
        for name, info in cache.expiring_secrets().items():
            print(f"WARNING: {name} expires at {info.expires_on}")

        await AsyncSecretRotator(configuration.client).rotate(
            rotation_name, rotation_value, expires_on
        )
        await cache.refresh(rotation_name)
        print(f"Rotated {rotation_name}")


def main() -> None:
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
