from __future__ import annotations

import asyncio
import logging
import os
import threading
import time

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import (
    AzureAppConfigurationClient as AsyncAzureAppConfigurationClient,
)
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


LABEL = os.getenv("APPCONFIG_LABEL", "production")
FLAG_ID = os.getenv("APPCONFIG_DEMO_FLAG", "beta-dashboard")
SENTINEL_KEY = os.getenv("APPCONFIG_SENTINEL_KEY", "app:sentinel")
POLL_INTERVAL = float(os.getenv("APPCONFIG_POLL_INTERVAL", "5"))
WATCH_SECONDS = float(os.getenv("APPCONFIG_DEMO_WATCH_SECONDS", "6"))
SAMPLE_USERS = ("alice", "bob", "carol", "dave")


def run_sync_demo(endpoint: str) -> None:
    print("\n--- Synchronous demo ---")
    with DefaultAzureCredential() as credential:
        with AzureAppConfigurationClient(endpoint, credential) as client:
            configuration = ConfigurationService(client)
            flags = FeatureFlagEvaluator(configuration)

            print("Unlabeled message:", configuration.get_setting("app:message"))
            print(
                f"{LABEL} message:",
                configuration.get_setting("app:message", LABEL),
            )
            print(
                f"{LABEL} settings:",
                configuration.list_settings("app:", LABEL),
            )
            for user_id in SAMPLE_USERS:
                enabled = flags.is_enabled(FLAG_ID, user_id, LABEL)
                print(f"{FLAG_ID} for {user_id}: {enabled}")

            watcher = ConfigurationWatcher(
                configuration,
                [(SENTINEL_KEY, LABEL)],
                polling_interval=POLL_INTERVAL,
            )
            stop_event = threading.Event()
            thread = threading.Thread(
                target=watcher.run,
                args=(stop_event,),
                name="app-configuration-watcher",
                daemon=True,
            )
            print(f"Watching {SENTINEL_KEY!r} for {WATCH_SECONDS:g} seconds...")
            thread.start()
            time.sleep(WATCH_SECONDS)
            stop_event.set()
            thread.join()


async def run_async_demo(endpoint: str) -> None:
    print("\n--- Asynchronous demo ---")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncAzureAppConfigurationClient(endpoint, credential) as client:
            configuration = AsyncConfigurationService(client)
            flags = AsyncFeatureFlagEvaluator(configuration)

            print(
                "Unlabeled message:",
                await configuration.get_setting("app:message"),
            )
            print(
                f"{LABEL} message:",
                await configuration.get_setting("app:message", LABEL),
            )
            print(
                f"{LABEL} settings:",
                await configuration.list_settings("app:", LABEL),
            )
            for user_id in SAMPLE_USERS:
                enabled = await flags.is_enabled(FLAG_ID, user_id, LABEL)
                print(f"{FLAG_ID} for {user_id}: {enabled}")

            watcher = AsyncConfigurationWatcher(
                configuration,
                [(SENTINEL_KEY, LABEL)],
                polling_interval=POLL_INTERVAL,
            )
            stop_event = asyncio.Event()
            task = asyncio.create_task(watcher.run(stop_event))
            print(f"Watching {SENTINEL_KEY!r} for {WATCH_SECONDS:g} seconds...")
            await asyncio.sleep(WATCH_SECONDS)
            stop_event.set()
            await task


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    endpoint = os.environ["AZURE_APPCONFIGURATION_ENDPOINT"]
    run_sync_demo(endpoint)
    asyncio.run(run_async_demo(endpoint))


if __name__ == "__main__":
    main()
