from __future__ import annotations

import asyncio
import logging
import os
import time

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from config_service import AsyncConfigurationService, ConfigurationService
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator
from watcher import AsyncConfigurationWatcher, ConfigurationWatcher


ENDPOINT_VARIABLE = "AZURE_APPCONFIGURATION_ENDPOINT"
USERS = ("alice", "bob", "charlie", "dana")


def run_sync_demo(endpoint: str) -> None:
    print("=== Synchronous demo ===")
    credential = DefaultAzureCredential()
    try:
        with AzureAppConfigurationClient(endpoint, credential) as client:
            configuration = ConfigurationService(client)
            flags = FeatureFlagEvaluator(configuration)

            print("Production message:", configuration.get_setting_with_label(
                "demo:message", "production"
            ))
            print("Staging message:", configuration.get_setting_with_label(
                "demo:message", "staging"
            ))
            print("Demo settings:", configuration.list_settings("demo:"))
            for user_id in USERS:
                enabled = flags.is_enabled("percentage-rollout", user_id)
                print(f"percentage-rollout for {user_id}: {enabled}")

            watcher = ConfigurationWatcher(
                configuration, ["demo:sentinel"], polling_interval=5
            )
            watcher.start(lambda: print("Sync configuration cache refreshed."))
            print("Watching the sync sentinel for 10 seconds...")
            time.sleep(10)
            watcher.stop()
    finally:
        credential.close()


async def run_async_demo(endpoint: str) -> None:
    print("\n=== Asynchronous demo ===")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncClient(endpoint, credential) as client:
            configuration = AsyncConfigurationService(client)
            flags = AsyncFeatureFlagEvaluator(configuration)

            print("Production message:", await configuration.get_setting_with_label(
                "demo:message", "production"
            ))
            print("Staging message:", await configuration.get_setting_with_label(
                "demo:message", "staging"
            ))
            print("Demo settings:", await configuration.list_settings("demo:"))
            for user_id in USERS:
                enabled = await flags.is_enabled("percentage-rollout", user_id)
                print(f"percentage-rollout for {user_id}: {enabled}")

            watcher = AsyncConfigurationWatcher(
                configuration, ["demo:sentinel"], polling_interval=5
            )

            async def refreshed() -> None:
                print("Async configuration cache refreshed.")

            await watcher.start(refreshed)
            print("Watching the async sentinel for 10 seconds...")
            await asyncio.sleep(10)
            await watcher.stop()


async def main() -> None:
    endpoint = os.environ.get(ENDPOINT_VARIABLE)
    if not endpoint:
        raise RuntimeError(
            f"Set {ENDPOINT_VARIABLE} to your App Configuration endpoint."
        )
    run_sync_demo(endpoint)
    await run_async_demo(endpoint)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(main())

