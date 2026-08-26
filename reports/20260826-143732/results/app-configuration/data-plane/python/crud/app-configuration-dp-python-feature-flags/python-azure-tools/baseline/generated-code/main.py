"""Run synchronous and asynchronous Azure App Configuration demos."""

from __future__ import annotations

import asyncio
import os
import time

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from appconfig_manager import (
    AsyncConfigurationService,
    AsyncConfigurationWatcher,
    AsyncFeatureFlagEvaluator,
    ConfigurationService,
    ConfigurationWatcher,
    FeatureFlagEvaluator,
)

SAMPLE_USERS = ("alice", "bob", "charlie", "diana")


def _endpoint() -> str:
    try:
        return os.environ["AZURE_APPCONFIG_ENDPOINT"]
    except KeyError as exc:
        raise RuntimeError(
            "Set AZURE_APPCONFIG_ENDPOINT to an App Configuration endpoint"
        ) from exc


def run_sync_demo(endpoint: str, watch_seconds: float) -> None:
    print("\n--- Synchronous demo ---")
    with DefaultAzureCredential() as credential:
        with AzureAppConfigurationClient(endpoint, credential) as client:
            configuration = ConfigurationService(client)
            flags = FeatureFlagEvaluator(configuration)

            print(
                "production API URL:",
                configuration.get_setting_with_label("Demo:ApiUrl", "production"),
            )
            print(
                "staging settings:",
                configuration.list_settings("Demo:", label="staging"),
            )
            for user_id in SAMPLE_USERS:
                enabled = flags.is_enabled(
                    "NewCheckout", user_id=user_id, label="production"
                )
                print(f"NewCheckout for {user_id}: {enabled}")

            watcher = ConfigurationWatcher(
                configuration,
                sentinel_keys=["Demo:Sentinel"],
                polling_interval=5,
                label="production",
                on_refresh=lambda: print("Sync configuration cache refreshed"),
            )
            watcher.start()
            print(f"Watching for sentinel changes for {watch_seconds:g} seconds...")
            time.sleep(watch_seconds)
            watcher.stop()


async def run_async_demo(endpoint: str, watch_seconds: float) -> None:
    print("\n--- Asynchronous demo ---")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncClient(endpoint, credential) as client:
            configuration = AsyncConfigurationService(client)
            flags = AsyncFeatureFlagEvaluator(configuration)

            print(
                "production API URL:",
                await configuration.get_setting_with_label(
                    "Demo:ApiUrl", "production"
                ),
            )
            print(
                "staging settings:",
                await configuration.list_settings("Demo:", label="staging"),
            )
            for user_id in SAMPLE_USERS:
                enabled = await flags.is_enabled(
                    "NewCheckout", user_id=user_id, label="production"
                )
                print(f"NewCheckout for {user_id}: {enabled}")

            watcher = AsyncConfigurationWatcher(
                configuration,
                sentinel_keys=["Demo:Sentinel"],
                polling_interval=5,
                label="production",
                on_refresh=lambda: print("Async configuration cache refreshed"),
            )
            await watcher.start()
            print(f"Watching for sentinel changes for {watch_seconds:g} seconds...")
            await asyncio.sleep(watch_seconds)
            await watcher.stop()


def main() -> None:
    endpoint = _endpoint()
    watch_seconds = float(os.getenv("DEMO_WATCH_SECONDS", "15"))
    if watch_seconds < 0:
        raise ValueError("DEMO_WATCH_SECONDS cannot be negative")
    run_sync_demo(endpoint, watch_seconds)
    asyncio.run(run_async_demo(endpoint, watch_seconds))


if __name__ == "__main__":
    main()

