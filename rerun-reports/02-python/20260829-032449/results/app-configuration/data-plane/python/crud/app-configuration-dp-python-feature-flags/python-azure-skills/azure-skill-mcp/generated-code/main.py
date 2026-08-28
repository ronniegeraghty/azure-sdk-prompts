from __future__ import annotations

import asyncio
import logging
import os
import time

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.core.exceptions import ResourceNotFoundError
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from configuration_service import AsyncConfigurationService, ConfigurationService
from configuration_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


SETTING_KEY = "Demo:Settings:Message"
SETTINGS_PREFIX = "Demo:Settings:"
FLAG_ID = "BetaExperience"
SAMPLE_USERS = ("alice", "bob", "carol", "dave")


def run_sync_demo(
    endpoint: str,
    sentinel_keys: list[str],
    polling_interval: float,
    watch_seconds: float,
) -> None:
    print("\n=== Synchronous implementation ===")
    credential = DefaultAzureCredential()
    client = AzureAppConfigurationClient(base_url=endpoint, credential=credential)
    configuration = ConfigurationService(client)

    try:
        _print_sync_setting(configuration, SETTING_KEY)
        _print_sync_setting(configuration, SETTING_KEY, "production")
        _print_sync_setting(configuration, SETTING_KEY, "staging")

        try:
            print(f"Prefix values: {configuration.list_settings(SETTINGS_PREFIX)}")
        except ResourceNotFoundError:
            print(f"No settings found under {SETTINGS_PREFIX!r}")

        evaluator = FeatureFlagEvaluator(configuration)
        for user_id in SAMPLE_USERS:
            enabled = evaluator.is_enabled(FLAG_ID, user_id, label="production")
            print(f"Flag {FLAG_ID!r} for {user_id}: {enabled}")

        watcher = ConfigurationWatcher(
            configuration,
            sentinel_keys,
            polling_interval,
            on_change=lambda keys: print(f"Sync refresh triggered by: {', '.join(keys)}"),
        )
        print(
            f"Watching {sentinel_keys} for {watch_seconds:g} seconds "
            f"(poll every {polling_interval:g} seconds)..."
        )
        watcher.start()
        try:
            time.sleep(watch_seconds)
        finally:
            watcher.stop()
    finally:
        client.close()
        credential.close()


async def run_async_demo(
    endpoint: str,
    sentinel_keys: list[str],
    polling_interval: float,
    watch_seconds: float,
) -> None:
    print("\n=== Asynchronous implementation ===")
    credential = AsyncDefaultAzureCredential()
    client = AsyncAzureAppConfigurationClient(base_url=endpoint, credential=credential)
    configuration = AsyncConfigurationService(client)

    try:
        await _print_async_setting(configuration, SETTING_KEY)
        await _print_async_setting(configuration, SETTING_KEY, "production")
        await _print_async_setting(configuration, SETTING_KEY, "staging")

        try:
            values = await configuration.list_settings(SETTINGS_PREFIX)
            print(f"Prefix values: {values}")
        except ResourceNotFoundError:
            print(f"No settings found under {SETTINGS_PREFIX!r}")

        evaluator = AsyncFeatureFlagEvaluator(configuration)
        for user_id in SAMPLE_USERS:
            enabled = await evaluator.is_enabled(
                FLAG_ID, user_id, label="production"
            )
            print(f"Flag {FLAG_ID!r} for {user_id}: {enabled}")

        watcher = AsyncConfigurationWatcher(
            configuration,
            sentinel_keys,
            polling_interval,
            on_change=lambda keys: print(f"Async refresh triggered by: {', '.join(keys)}"),
        )
        print(
            f"Watching {sentinel_keys} for {watch_seconds:g} seconds "
            f"(poll every {polling_interval:g} seconds)..."
        )
        watcher_task = asyncio.create_task(watcher.run())
        try:
            await asyncio.sleep(watch_seconds)
        finally:
            watcher.stop()
            await watcher_task
    finally:
        await client.close()
        await credential.close()


def _print_sync_setting(
    configuration: ConfigurationService, key: str, label: str | None = None
) -> None:
    try:
        value = configuration.get_setting(key, label)
    except ResourceNotFoundError:
        value = "<not found>"
    print(f"{key} [{label or 'no label'}]: {value}")


async def _print_async_setting(
    configuration: AsyncConfigurationService, key: str, label: str | None = None
) -> None:
    try:
        value = await configuration.get_setting(key, label)
    except ResourceNotFoundError:
        value = "<not found>"
    print(f"{key} [{label or 'no label'}]: {value}")


def main() -> None:
    logging.basicConfig(
        level=os.getenv("LOG_LEVEL", "INFO"),
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )

    endpoint = os.getenv("AZURE_APPCONFIG_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "Set AZURE_APPCONFIG_ENDPOINT to your App Configuration endpoint"
        )

    sentinel_keys = [
        key.strip()
        for key in os.getenv("APPCONFIG_SENTINEL_KEYS", "Demo:Sentinel").split(",")
        if key.strip()
    ]
    polling_interval = float(os.getenv("APPCONFIG_POLL_INTERVAL", "10"))
    watch_seconds = float(os.getenv("DEMO_WATCH_SECONDS", "30"))
    if watch_seconds < 0:
        raise ValueError("DEMO_WATCH_SECONDS must not be negative")

    run_sync_demo(endpoint, sentinel_keys, polling_interval, watch_seconds)
    asyncio.run(
        run_async_demo(endpoint, sentinel_keys, polling_interval, watch_seconds)
    )


if __name__ == "__main__":
    main()
