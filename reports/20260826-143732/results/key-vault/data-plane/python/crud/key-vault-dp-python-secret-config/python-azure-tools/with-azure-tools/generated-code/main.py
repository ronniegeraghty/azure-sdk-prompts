from __future__ import annotations

import asyncio
import os
from datetime import datetime, timedelta, timezone

from secret_config.factory import AsyncConfiguration, SyncConfiguration

REQUIRED_SECRETS = {
    "database-url": None,
    "api-key": None,
    "feature-flag": "disabled",
}
ROTATION_SECRET_NAME = "api-key"


def print_cache_reads(label: str, values: dict[str, str | None]) -> None:
    for name, value in values.items():
        state = "available" if value is not None else "missing"
        print(f"[{label}] {name}: {state}")


def warn_about_expiry(label: str, names: list[str]) -> None:
    for name in names:
        print(f"[{label}] WARNING: {name} is expired or near expiry")


def run_sync_demo(rotated_value: str) -> None:
    print("Running synchronous Key Vault configuration demo")
    with SyncConfiguration(warning_days=7) as config:
        config.cache.load_required(REQUIRED_SECRETS)
        print_cache_reads(
            "sync",
            {name: config.cache.get(name) for name in REQUIRED_SECRETS},
        )

        config.cache.refresh("feature-flag")
        warn_about_expiry("sync", list(config.cache.expiring()))

        expires_on = datetime.now(timezone.utc) + timedelta(days=90)
        config.rotator.rotate(
            ROTATION_SECRET_NAME,
            rotated_value,
            expires_on,
        )
        config.cache.refresh(ROTATION_SECRET_NAME)
        print(f"[sync] Rotated and refreshed {ROTATION_SECRET_NAME}")


async def run_async_demo(rotated_value: str) -> None:
    print("Running asynchronous Key Vault configuration demo")
    async with AsyncConfiguration(warning_days=7) as config:
        await config.cache.load_required(REQUIRED_SECRETS)
        values = {
            name: await config.cache.get(name)
            for name in REQUIRED_SECRETS
        }
        print_cache_reads("async", values)

        await config.cache.refresh("feature-flag")
        warn_about_expiry("async", list(config.cache.expiring()))

        expires_on = datetime.now(timezone.utc) + timedelta(days=90)
        await config.rotator.rotate(
            ROTATION_SECRET_NAME,
            rotated_value,
            expires_on,
        )
        await config.cache.refresh(ROTATION_SECRET_NAME)
        print(f"[async] Rotated and refreshed {ROTATION_SECRET_NAME}")


def main() -> None:
    try:
        rotated_value = os.environ["DEMO_ROTATED_SECRET_VALUE"]
    except KeyError as error:
        raise RuntimeError(
            "DEMO_ROTATED_SECRET_VALUE must be set for the rotation demo"
        ) from error

    run_sync_demo(rotated_value)
    asyncio.run(run_async_demo(rotated_value))


if __name__ == "__main__":
    main()
