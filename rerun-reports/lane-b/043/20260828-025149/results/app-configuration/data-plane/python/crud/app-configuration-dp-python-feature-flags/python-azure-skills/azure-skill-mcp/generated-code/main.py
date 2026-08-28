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


logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
logger = logging.getLogger(__name__)

SETTING_KEY = "Demo:Message"
SETTING_PREFIX = "Demo:"
FEATURE_FLAG_NAME = "BetaExperience"
SENTINEL_KEYS = ("Demo:Sentinel",)
SAMPLE_USERS = ("alice", "bob", "charlie", "diana")


def _required_endpoint() -> str:
    endpoint = os.getenv("AZURE_APPCONFIG_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "Set AZURE_APPCONFIG_ENDPOINT to the URL of an App Configuration store"
        )
    return endpoint


def run_sync_demo(
    endpoint: str, label: str, polling_interval: float, watch_duration: float
) -> None:
    logger.info("Starting synchronous demo")
    credential = DefaultAzureCredential()
    try:
        with ConfigurationService(endpoint, credential) as configuration:
            logger.info(
                "%s [%s] = %r",
                SETTING_KEY,
                label,
                configuration.get_setting(SETTING_KEY, label),
            )
            logger.info(
                "Settings under %s [%s]: %s",
                SETTING_PREFIX,
                label,
                configuration.list_settings(SETTING_PREFIX, label),
            )

            evaluator = FeatureFlagEvaluator(configuration)
            for user_id in SAMPLE_USERS:
                enabled = evaluator.is_enabled(FEATURE_FLAG_NAME, user_id, label)
                logger.info(
                    "Feature %s for user %s [%s]: %s",
                    FEATURE_FLAG_NAME,
                    user_id,
                    label,
                    enabled,
                )

            watcher = ConfigurationWatcher(
                configuration,
                SENTINEL_KEYS,
                polling_interval,
                label,
                on_refresh=lambda: logger.info(
                    "Sentinel changed; synchronous cache fully refreshed"
                ),
            )
            with watcher:
                logger.info(
                    "Watching synchronous sentinels for %.1f seconds", watch_duration
                )
                time.sleep(watch_duration)
    finally:
        credential.close()


async def run_async_demo(
    endpoint: str, label: str, polling_interval: float, watch_duration: float
) -> None:
    logger.info("Starting asynchronous demo")
    credential = AsyncDefaultAzureCredential()
    try:
        async with AsyncConfigurationService(endpoint, credential) as configuration:
            logger.info(
                "%s [%s] = %r",
                SETTING_KEY,
                label,
                await configuration.get_setting(SETTING_KEY, label),
            )
            logger.info(
                "Settings under %s [%s]: %s",
                SETTING_PREFIX,
                label,
                await configuration.list_settings(SETTING_PREFIX, label),
            )

            evaluator = AsyncFeatureFlagEvaluator(configuration)
            for user_id in SAMPLE_USERS:
                enabled = await evaluator.is_enabled(
                    FEATURE_FLAG_NAME, user_id, label
                )
                logger.info(
                    "Feature %s for user %s [%s]: %s",
                    FEATURE_FLAG_NAME,
                    user_id,
                    label,
                    enabled,
                )

            async def on_refresh() -> None:
                logger.info("Sentinel changed; asynchronous cache fully refreshed")

            watcher = AsyncConfigurationWatcher(
                configuration,
                SENTINEL_KEYS,
                polling_interval,
                label,
                on_refresh=on_refresh,
            )
            async with watcher:
                logger.info(
                    "Watching asynchronous sentinels for %.1f seconds", watch_duration
                )
                await asyncio.sleep(watch_duration)
    finally:
        await credential.close()


def main() -> None:
    endpoint = _required_endpoint()
    label = os.getenv("CONFIG_LABEL", "production")
    polling_interval = float(os.getenv("CONFIG_POLL_INTERVAL_SECONDS", "5"))
    watch_duration = float(os.getenv("CONFIG_WATCH_DURATION_SECONDS", "15"))

    run_sync_demo(endpoint, label, polling_interval, watch_duration)
    asyncio.run(run_async_demo(endpoint, label, polling_interval, watch_duration))


if __name__ == "__main__":
    main()
