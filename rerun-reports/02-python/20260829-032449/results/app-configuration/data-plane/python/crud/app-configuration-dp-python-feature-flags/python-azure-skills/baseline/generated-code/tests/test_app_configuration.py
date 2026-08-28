from __future__ import annotations

import asyncio
import json
import unittest
from types import SimpleNamespace
from unittest.mock import Mock

from azure.core.exceptions import ResourceNotModifiedError

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


def setting(key: str, value: str, etag: str) -> SimpleNamespace:
    return SimpleNamespace(key=key, value=value, etag=etag)


def percentage_flag(value: float) -> str:
    return json.dumps(
        {
            "id": "Rollout",
            "enabled": True,
            "conditions": {
                "client_filters": [
                    {
                        "name": "Microsoft.Percentage",
                        "parameters": {"Value": value},
                    }
                ]
            },
        }
    )


class AsyncItems:
    def __init__(self, items: list[SimpleNamespace]) -> None:
        self._items = items

    def __aiter__(self):
        async def iterate():
            for item in self._items:
                yield item

        return iterate()


class ConfigurationServiceTests(unittest.TestCase):
    def test_conditional_get_reuses_cached_value(self) -> None:
        client = Mock()
        client.get_configuration_setting.side_effect = [
            setting("Demo:Message", "hello", "etag-1"),
            None,
        ]
        service = ConfigurationService(client)

        self.assertEqual(service.get_setting("Demo:Message"), "hello")
        self.assertEqual(service.get_setting("Demo:Message"), "hello")
        self.assertEqual(client.get_configuration_setting.call_count, 2)
        self.assertEqual(
            client.get_configuration_setting.call_args.kwargs["etag"], "etag-1"
        )

    def test_prefix_is_cached_until_refresh(self) -> None:
        client = Mock()
        client.list_configuration_settings.side_effect = [
            [setting("Demo:A", "one", "1")],
            [setting("Demo:A", "two", "2")],
        ]
        service = ConfigurationService(client)

        self.assertEqual(service.list_settings("Demo:"), {"Demo:A": "one"})
        self.assertEqual(service.list_settings("Demo:"), {"Demo:A": "one"})
        service.refresh_all()
        self.assertEqual(service.list_settings("Demo:"), {"Demo:A": "two"})
        self.assertEqual(client.list_configuration_settings.call_count, 2)


class FeatureFlagTests(unittest.TestCase):
    def test_percentage_is_deterministic_and_respects_bounds(self) -> None:
        reader = Mock()
        evaluator = FeatureFlagEvaluator(reader)

        reader.get_setting.return_value = percentage_flag(30)
        first = evaluator.is_enabled("Rollout", "same-user")
        self.assertEqual(first, evaluator.is_enabled("Rollout", "same-user"))

        reader.get_setting.return_value = percentage_flag(0)
        self.assertFalse(evaluator.is_enabled("Rollout", "alice"))
        reader.get_setting.return_value = percentage_flag(100)
        self.assertTrue(evaluator.is_enabled("Rollout", "alice"))

    def test_disabled_flag_is_always_disabled(self) -> None:
        reader = Mock()
        reader.get_setting.return_value = json.dumps(
            {"id": "Rollout", "enabled": False}
        )
        self.assertFalse(FeatureFlagEvaluator(reader).is_enabled("Rollout", "alice"))


class WatcherTests(unittest.TestCase):
    def test_changed_sentinel_refreshes_all_configuration(self) -> None:
        service = Mock()
        service.get_setting.side_effect = ["v1", "v2"]
        watcher = ConfigurationWatcher(service, ["Sentinel"], polling_interval=1)

        self.assertEqual(watcher.poll_once(), set())
        self.assertEqual(watcher.poll_once(), {"Sentinel"})
        service.refresh_all.assert_called_once_with()


class AsyncTests(unittest.IsolatedAsyncioTestCase):
    async def test_async_service_evaluator_and_watcher(self) -> None:
        client = Mock()
        client.get_configuration_setting = unittest.mock.AsyncMock(
            side_effect=[
                setting(
                    ".appconfig.featureflag/Rollout",
                    percentage_flag(100),
                    "1",
                ),
                setting("Sentinel", "v1", "2"),
                setting("Sentinel", "v2", "3"),
                setting(
                    ".appconfig.featureflag/Rollout",
                    percentage_flag(100),
                    "1",
                ),
                setting("Sentinel", "v2", "3"),
            ]
        )
        client.list_configuration_settings.return_value = AsyncItems([])
        service = AsyncConfigurationService(client)
        evaluator = AsyncFeatureFlagEvaluator(service)

        self.assertTrue(await evaluator.is_enabled("Rollout", "alice"))
        watcher = AsyncConfigurationWatcher(service, ["Sentinel"], polling_interval=1)
        self.assertEqual(await watcher.poll_once(), set())
        self.assertEqual(await watcher.poll_once(), {"Sentinel"})


if __name__ == "__main__":
    unittest.main()
