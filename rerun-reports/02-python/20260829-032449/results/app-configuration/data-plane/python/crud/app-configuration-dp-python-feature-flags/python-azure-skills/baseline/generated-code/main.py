"""Demonstrate sync and async Azure App Configuration access."""

from __future__ import annotations

import asyncio
import logging
import os
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

CONFIG_PREFIX = "Demo:"
SETTING_KEY = "Demo:Message"
FEATURE_FLAG = "PercentageRollout"
SENTINEL_KEYS = ["Demo:Sentinel"]
LABELS = ("production", "staging")
SAMPLE_USERS = ("alice", "bob", "carol", "dave")


def _watch_seconds() -> float:
    return float(os.getenv("DEMO_WATCH_SECONDS", "10"))


def run_sync_demo(endpoint: str) -> None:
    print("\n--- Synchronous demo ---")
    credential = DefaultAzureCredential()
    client = AzureAppConfigurationClient(endpoint, credential)
    try:
        configuration = ConfigurationService(client)
        flags = FeatureFlagEvaluator(configuration)

        for label in LABELS:
            value = configuration.get_setting(SETTING_KEY, label)
            print(f"{SETTING_KEY} [{label}]: {value}")
        print(
            f"{CONFIG_PREFIX}* [production]: "
            f"{configuration.list_settings(CONFIG_PREFIX, 'production')}"
        )

        for user_id in SAMPLE_USERS:
            enabled = flags.is_enabled(
                FEATURE_FLAG, user_id=user_id, label="production"
            )
            print(f"{FEATURE_FLAG} for {user_id}: {enabled}")

        watcher = ConfigurationWatcher(
            configuration,
            SENTINEL_KEYS,
            polling_interval=2,
            label="production",
            on_refresh=lambda keys: print(
                f"Sync configuration refreshed; changed sentinels: {sorted(keys)}"
            ),
        )
        print(f"Watching sentinels for {_watch_seconds():g} seconds...")
        watcher.start()
        try:
            time.sleep(_watch_seconds())
        finally:
            watcher.stop()
    finally:
        client.close()
        credential.close()


async def run_async_demo(endpoint: str) -> None:
    print("\n--- Asynchronous demo ---")
    credential = AsyncDefaultAzureCredential()
    client = AsyncAzureAppConfigurationClient(endpoint, credential)
    try:
        configuration = AsyncConfigurationService(client)
        flags = AsyncFeatureFlagEvaluator(configuration)

        for label in LABELS:
            value = await configuration.get_setting(SETTING_KEY, label)
            print(f"{SETTING_KEY} [{label}]: {value}")
        print(
            f"{CONFIG_PREFIX}* [production]: "
            f"{await configuration.list_settings(CONFIG_PREFIX, 'production')}"
        )

        for user_id in SAMPLE_USERS:
            enabled = await flags.is_enabled(
                FEATURE_FLAG, user_id=user_id, label="production"
            )
            print(f"{FEATURE_FLAG} for {user_id}: {enabled}")

        watcher = AsyncConfigurationWatcher(
            configuration,
            SENTINEL_KEYS,
            polling_interval=2,
            label="production",
            on_refresh=lambda keys: print(
                f"Async configuration refreshed; changed sentinels: {sorted(keys)}"
            ),
        )
        print(f"Watching sentinels for {_watch_seconds():g} seconds...")
        await watcher.start()
        try:
            await asyncio.sleep(_watch_seconds())
        finally:
            await watcher.stop()
    finally:
        await client.close()
        await credential.close()


async def main() -> None:
    endpoint = os.environ.get("AZURE_APPCONFIGURATION_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "Set AZURE_APPCONFIGURATION_ENDPOINT to the App Configuration endpoint"
        )
    run_sync_demo(endpoint)
    await run_async_demo(endpoint)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(main())
