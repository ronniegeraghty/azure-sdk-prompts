"""Run sync and async Azure App Configuration demonstrations."""

from __future__ import annotations

import asyncio
import logging
import os

from azure.appconfiguration import AzureAppConfigurationClient
from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from appconfig_demo import (
    AsyncConfigurationService,
    AsyncConfigurationWatcher,
    AsyncFeatureFlagEvaluator,
    ConfigurationService,
    ConfigurationWatcher,
    FeatureFlagEvaluator,
    SentinelKey,
)

logging.basicConfig(level=os.getenv("LOG_LEVEL", "INFO"))
logger = logging.getLogger(__name__)


def _endpoint() -> str:
    endpoint = os.getenv("AZURE_APPCONFIGURATION_ENDPOINT")
    if not endpoint:
        raise RuntimeError("AZURE_APPCONFIGURATION_ENDPOINT must be set")
    return endpoint


def _demo_options() -> tuple[str, str, str, str, list[str], float, int]:
    key = os.getenv("DEMO_CONFIG_KEY", "demo:message")
    label = os.getenv("DEMO_CONFIG_LABEL", "production")
    prefix = os.getenv("DEMO_CONFIG_PREFIX", "demo:")
    flag = os.getenv("DEMO_FEATURE_FLAG", "gradual-rollout")
    users = [
        user.strip()
        for user in os.getenv("DEMO_USER_IDS", "alice,bob,charlie,diana").split(",")
        if user.strip()
    ]
    interval = float(os.getenv("DEMO_POLL_INTERVAL", "5"))
    polls = int(os.getenv("DEMO_MAX_POLLS", "3"))
    return key, label, prefix, flag, users, interval, polls


def run_sync_demo(endpoint: str) -> None:
    """Run the synchronous client, evaluator, and watcher."""
    key, label, prefix, flag, users, interval, polls = _demo_options()
    sentinel = SentinelKey(os.getenv("DEMO_SENTINEL_KEY", "demo:sentinel"), label)

    print("\n=== Synchronous demo ===")
    with DefaultAzureCredential() as credential:
        with AzureAppConfigurationClient(
            base_url=endpoint,
            credential=credential,
        ) as client:
            configuration = ConfigurationService(client)
            flags = FeatureFlagEvaluator(configuration)

            print(f"{key} (unlabeled): {configuration.get_setting(key)}")
            print(f"{key} ({label}): {configuration.get_setting(key, label)}")
            print(
                f"{prefix}* ({label}): "
                f"{configuration.list_settings(prefix, label)}"
            )
            for user_id in users:
                enabled = flags.is_enabled(flag, user_id=user_id, label=label)
                print(f"feature {flag!r} for {user_id}: {enabled}")

            def refreshed(changed: set[SentinelKey]) -> None:
                print(f"sync refresh triggered by: {sorted(s.key for s in changed)}")

            print(
                f"Watching {sentinel.key!r} for {polls} polls; "
                "change its value to trigger a refresh."
            )
            ConfigurationWatcher(
                configuration,
                [sentinel],
                interval,
                on_refresh=refreshed,
            ).run(max_polls=polls)


async def run_async_demo(endpoint: str) -> None:
    """Run the asynchronous client, evaluator, and watcher."""
    key, label, prefix, flag, users, interval, polls = _demo_options()
    sentinel = SentinelKey(os.getenv("DEMO_SENTINEL_KEY", "demo:sentinel"), label)

    print("\n=== Asynchronous demo ===")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncAzureAppConfigurationClient(
            base_url=endpoint,
            credential=credential,
        ) as client:
            configuration = AsyncConfigurationService(client)
            flags = AsyncFeatureFlagEvaluator(configuration)

            print(f"{key} (unlabeled): {await configuration.get_setting(key)}")
            print(f"{key} ({label}): {await configuration.get_setting(key, label)}")
            print(
                f"{prefix}* ({label}): "
                f"{await configuration.list_settings(prefix, label)}"
            )
            for user_id in users:
                enabled = await flags.is_enabled(flag, user_id=user_id, label=label)
                print(f"feature {flag!r} for {user_id}: {enabled}")

            async def refreshed(changed: set[SentinelKey]) -> None:
                print(f"async refresh triggered by: {sorted(s.key for s in changed)}")

            print(
                f"Watching {sentinel.key!r} for {polls} polls; "
                "change its value to trigger a refresh."
            )
            await AsyncConfigurationWatcher(
                configuration,
                [sentinel],
                interval,
                on_refresh=refreshed,
            ).run(max_polls=polls)


def main() -> None:
    endpoint = _endpoint()
    run_sync_demo(endpoint)
    asyncio.run(run_async_demo(endpoint))


if __name__ == "__main__":
    try:
        main()
    except (ValueError, RuntimeError):
        logger.exception("Demo configuration is invalid")
        raise
