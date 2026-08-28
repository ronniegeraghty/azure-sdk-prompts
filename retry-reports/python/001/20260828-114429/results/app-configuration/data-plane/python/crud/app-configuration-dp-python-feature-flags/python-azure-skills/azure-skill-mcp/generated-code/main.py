from __future__ import annotations

import asyncio
import os

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator

ENDPOINT_ENV = "AZURE_APPCONFIG_ENDPOINT"
SAMPLE_USERS = ("alice", "bob", "carol", "dave")


def run_sync_demo(endpoint: str, polling_interval: float) -> None:
    print("=== Synchronous demo ===")
    credential = DefaultAzureCredential()
    client = AzureAppConfigurationClient(base_url=endpoint, credential=credential)
    try:
        config = ConfigurationService(client)
        print("Demo:Message (production):", config.get_setting_with_label("Demo:Message", "production"))
        print("Demo:Message (staging):", config.get_setting_with_label("Demo:Message", "staging"))
        print("Demo settings:", config.list_settings("Demo:", "production"))

        evaluator = FeatureFlagEvaluator(config, label="production")
        for user_id in SAMPLE_USERS:
            enabled = evaluator.is_enabled("BetaFeature", user_id=user_id)
            print(f"BetaFeature for {user_id}: {enabled}")

        watcher = ConfigurationWatcher(
            config,
            sentinel_keys=["Sentinel"],
            polling_interval=polling_interval,
            label="production",
            on_refresh=lambda: print("Sync sentinel changed; cache refreshed."),
        )
        print("Polling sync sentinel once...")
        watcher.run(max_polls=1)
    finally:
        client.close()
        credential.close()


async def run_async_demo(endpoint: str, polling_interval: float) -> None:
    print("\n=== Asynchronous demo ===")
    credential = AsyncDefaultAzureCredential()
    client = AsyncAzureAppConfigurationClient(base_url=endpoint, credential=credential)
    try:
        config = AsyncConfigurationService(client)
        print(
            "Demo:Message (production):",
            await config.get_setting_with_label("Demo:Message", "production"),
        )
        print(
            "Demo:Message (staging):",
            await config.get_setting_with_label("Demo:Message", "staging"),
        )
        print("Demo settings:", await config.list_settings("Demo:", "production"))

        evaluator = AsyncFeatureFlagEvaluator(config, label="production")
        for user_id in SAMPLE_USERS:
            enabled = await evaluator.is_enabled("BetaFeature", user_id=user_id)
            print(f"BetaFeature for {user_id}: {enabled}")

        watcher = AsyncConfigurationWatcher(
            config,
            sentinel_keys=["Sentinel"],
            polling_interval=polling_interval,
            label="production",
            on_refresh=lambda: print("Async sentinel changed; cache refreshed."),
        )
        print("Polling async sentinel once...")
        await watcher.run(max_polls=1)
    finally:
        await client.close()
        await credential.close()


def main() -> None:
    endpoint = os.environ.get(ENDPOINT_ENV)
    if not endpoint:
        raise RuntimeError(f"Set {ENDPOINT_ENV} to your Azure App Configuration endpoint")
    polling_interval = float(os.environ.get("APPCONFIG_POLL_INTERVAL", "5"))
    run_sync_demo(endpoint, polling_interval)
    asyncio.run(run_async_demo(endpoint, polling_interval))


if __name__ == "__main__":
    main()
