from __future__ import annotations

import asyncio
import logging
import os
import time
from typing import List

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")


def _sentinel_keys() -> List[str]:
    raw_keys = os.getenv("DEMO_SENTINEL_KEYS", "Sentinel")
    keys = [key.strip() for key in raw_keys.split(",") if key.strip()]
    if not keys:
        raise ValueError("DEMO_SENTINEL_KEYS must contain at least one key")
    return keys


def run_sync_demo(endpoint: str) -> None:
    print("\n--- Synchronous demo ---")
    config_key = os.getenv("DEMO_CONFIG_KEY", "Demo:Message")
    config_prefix = os.getenv("DEMO_CONFIG_PREFIX", "Demo:")
    flag_name = os.getenv("DEMO_FEATURE_FLAG", "BetaFeature")
    interval = float(os.getenv("DEMO_POLL_INTERVAL_SECONDS", "5"))
    watch_seconds = float(os.getenv("DEMO_WATCH_SECONDS", "15"))

    credential = DefaultAzureCredential()
    configuration = ConfigurationService(endpoint, credential)
    try:
        print("staging:", configuration.get_setting_with_label(config_key, "staging"))
        print("production:", configuration.get_setting_with_label(config_key, "production"))
        print("production settings:", configuration.list_settings(config_prefix, "production"))

        evaluator = FeatureFlagEvaluator(configuration)
        for user_id in ("alice", "bob", "carol", "dave"):
            enabled = evaluator.is_enabled(flag_name, user_id, "production")
            print(f"{flag_name} for {user_id}: {enabled}")

        watcher = ConfigurationWatcher(
            configuration,
            _sentinel_keys(),
            interval,
            label="production",
            on_refresh=lambda keys: print(f"Refreshed after sentinel change: {keys}"),
        )
        print(f"Watching sentinels for {watch_seconds:g} seconds...")
        watcher.start()
        try:
            time.sleep(watch_seconds)
        finally:
            watcher.stop()
    finally:
        configuration.close()
        credential.close()


async def run_async_demo(endpoint: str) -> None:
    print("\n--- Asynchronous demo ---")
    config_key = os.getenv("DEMO_CONFIG_KEY", "Demo:Message")
    config_prefix = os.getenv("DEMO_CONFIG_PREFIX", "Demo:")
    flag_name = os.getenv("DEMO_FEATURE_FLAG", "BetaFeature")
    interval = float(os.getenv("DEMO_POLL_INTERVAL_SECONDS", "5"))
    watch_seconds = float(os.getenv("DEMO_WATCH_SECONDS", "15"))

    credential = AsyncDefaultAzureCredential()
    configuration = AsyncConfigurationService(endpoint, credential)
    try:
        print(
            "staging:",
            await configuration.get_setting_with_label(config_key, "staging"),
        )
        print(
            "production:",
            await configuration.get_setting_with_label(config_key, "production"),
        )
        print(
            "production settings:",
            await configuration.list_settings(config_prefix, "production"),
        )

        evaluator = AsyncFeatureFlagEvaluator(configuration)
        for user_id in ("alice", "bob", "carol", "dave"):
            enabled = await evaluator.is_enabled(flag_name, user_id, "production")
            print(f"{flag_name} for {user_id}: {enabled}")

        watcher = AsyncConfigurationWatcher(
            configuration,
            _sentinel_keys(),
            interval,
            label="production",
            on_refresh=lambda keys: print(f"Refreshed after sentinel change: {keys}"),
        )
        print(f"Watching sentinels for {watch_seconds:g} seconds...")
        watcher.start()
        try:
            await asyncio.sleep(watch_seconds)
        finally:
            await watcher.stop()
    finally:
        await configuration.close()
        await credential.close()


def main() -> None:
    endpoint = os.getenv("AZURE_APP_CONFIG_ENDPOINT")
    if not endpoint:
        raise RuntimeError("Set AZURE_APP_CONFIG_ENDPOINT before running the demo")

    run_sync_demo(endpoint)
    asyncio.run(run_async_demo(endpoint))


if __name__ == "__main__":
    main()
