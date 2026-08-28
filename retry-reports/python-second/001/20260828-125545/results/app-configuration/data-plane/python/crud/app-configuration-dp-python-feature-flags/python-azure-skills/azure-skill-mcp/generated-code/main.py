from __future__ import annotations

import asyncio
import logging
import os
import time

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


ENDPOINT_ENV = "AZURE_APPCONFIG_ENDPOINT"
LABEL_ENV = "AZURE_APPCONFIG_LABEL"
WATCH_SECONDS_ENV = "CONFIG_WATCH_SECONDS"
SAMPLE_USERS = ("alice", "bob", "carol", "dave")


def _endpoint() -> str:
    endpoint = os.getenv(ENDPOINT_ENV)
    if not endpoint:
        raise RuntimeError(
            f"Set {ENDPOINT_ENV} to an Azure App Configuration endpoint, "
            "for example https://your-store.azconfig.io"
        )
    return endpoint


def _watch_seconds() -> float:
    value = float(os.getenv(WATCH_SECONDS_ENV, "15"))
    if value < 0:
        raise ValueError(f"{WATCH_SECONDS_ENV} cannot be negative")
    return value


def run_sync_demo(endpoint: str, label: str | None, watch_seconds: float) -> None:
    print("\n--- Synchronous demo ---")
    credential = DefaultAzureCredential()
    try:
        with ConfigurationService(endpoint, credential) as configuration:
            print("App:Message (no label):", configuration.get_setting("App:Message"))
            print(
                f"App:Message ({label or 'no'} label):",
                configuration.get_setting("App:Message", label),
            )
            print("App settings:", configuration.list_settings("App:", label))

            flags = FeatureFlagEvaluator(configuration)
            for user_id in SAMPLE_USERS:
                enabled = flags.is_enabled("BetaFeature", user_id, label=label)
                print(f"BetaFeature for {user_id}: {enabled}")

            watcher = ConfigurationWatcher(
                configuration,
                sentinel_keys=["App:Sentinel"],
                polling_interval=5,
                label=label,
            )
            watcher.start(
                lambda changed: print(
                    "Configuration refreshed after sentinel changes:", changed
                )
            )
            try:
                print(f"Watching for sync changes for {watch_seconds:g} seconds...")
                time.sleep(watch_seconds)
            finally:
                watcher.stop()
    finally:
        credential.close()


async def run_async_demo(
    endpoint: str, label: str | None, watch_seconds: float
) -> None:
    print("\n--- Asynchronous demo ---")
    credential = AsyncDefaultAzureCredential()
    try:
        async with AsyncConfigurationService(endpoint, credential) as configuration:
            print(
                "App:Message (no label):",
                await configuration.get_setting("App:Message"),
            )
            print(
                f"App:Message ({label or 'no'} label):",
                await configuration.get_setting("App:Message", label),
            )
            print("App settings:", await configuration.list_settings("App:", label))

            flags = AsyncFeatureFlagEvaluator(configuration)
            for user_id in SAMPLE_USERS:
                enabled = await flags.is_enabled("BetaFeature", user_id, label=label)
                print(f"BetaFeature for {user_id}: {enabled}")

            watcher = AsyncConfigurationWatcher(
                configuration,
                sentinel_keys=["App:Sentinel"],
                polling_interval=5,
                label=label,
            )
            watcher_task = asyncio.create_task(
                watcher.run(
                    lambda changed: print(
                        "Configuration refreshed after sentinel changes:", changed
                    )
                )
            )
            try:
                print(f"Watching for async changes for {watch_seconds:g} seconds...")
                await asyncio.sleep(watch_seconds)
            finally:
                watcher.stop()
                await watcher_task
    finally:
        await credential.close()


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    endpoint = _endpoint()
    label = os.getenv(LABEL_ENV, "production")
    watch_seconds = _watch_seconds()

    run_sync_demo(endpoint, label, watch_seconds)
    asyncio.run(run_async_demo(endpoint, label, watch_seconds))


if __name__ == "__main__":
    main()
