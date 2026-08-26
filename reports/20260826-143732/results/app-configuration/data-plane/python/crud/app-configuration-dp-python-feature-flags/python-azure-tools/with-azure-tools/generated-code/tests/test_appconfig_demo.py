from __future__ import annotations

import json
import unittest
from types import SimpleNamespace

from azure.core import MatchConditions

from appconfig_demo.configuration_service import (
    AsyncConfigurationService,
    ConfigurationService,
)
from appconfig_demo.feature_flags import (
    AsyncFeatureFlagEvaluator,
    FeatureFlagEvaluator,
    FeatureFlagFormatError,
)
from appconfig_demo.watcher import (
    AsyncConfigurationWatcher,
    ConfigurationWatcher,
    SentinelKey,
)


def setting(key: str, value: str | None, etag: str, label: str | None = None):
    return SimpleNamespace(key=key, value=value, etag=etag, label=label)


class FakeSyncClient:
    def __init__(self) -> None:
        self.settings = {
            ("app:one", None): setting("app:one", "1", "e1"),
            ("app:two", None): setting("app:two", "2", "e2"),
        }
        self.get_calls: list[dict[str, object]] = []
        self.list_calls: list[dict[str, object]] = []

    def get_configuration_setting(self, **kwargs):
        self.get_calls.append(kwargs)
        current = self.settings[(kwargs["key"], kwargs.get("label"))]
        if (
            kwargs.get("match_condition") is MatchConditions.IfModified
            and kwargs.get("etag") == current.etag
        ):
            return None
        return current

    def list_configuration_settings(self, **kwargs):
        self.list_calls.append(kwargs)
        prefix = str(kwargs["key_filter"])[:-1]
        label_filter = kwargs["label_filter"]
        selected = [
            item
            for item in self.settings.values()
            if item.key.startswith(prefix)
            and (item.label if item.label is not None else "\0") == label_filter
        ]
        if "fields" in kwargs:
            return [
                setting(item.key, None, item.etag, item.label) for item in selected
            ]
        return selected


class SyncTests(unittest.TestCase):
    def test_single_setting_uses_conditional_request_after_first_read(self) -> None:
        client = FakeSyncClient()
        service = ConfigurationService(client)

        self.assertEqual("1", service.get_setting("app:one"))
        self.assertEqual("1", service.get_setting("app:one"))

        self.assertEqual(2, len(client.get_calls))
        self.assertEqual("e1", client.get_calls[1]["etag"])
        self.assertIs(
            MatchConditions.IfModified,
            client.get_calls[1]["match_condition"],
        )

    def test_prefix_listing_only_downloads_changed_values(self) -> None:
        client = FakeSyncClient()
        service = ConfigurationService(client)

        self.assertEqual(
            {"app:one": "1", "app:two": "2"},
            service.list_settings("app:"),
        )
        client.settings[("app:two", None)] = setting("app:two", "updated", "e3")
        self.assertEqual(
            {"app:one": "1", "app:two": "updated"},
            service.list_settings("app:"),
        )

        self.assertEqual(["key", "label", "etag"], client.list_calls[1]["fields"])
        self.assertEqual(["app:two"], [call["key"] for call in client.get_calls])

    def test_percentage_rollout_is_deterministic(self) -> None:
        client = FakeSyncClient()
        payload = json.dumps(
            {
                "id": "gradual",
                "enabled": True,
                "conditions": {
                    "client_filters": [
                        {
                            "name": "Microsoft.Percentage",
                            "parameters": {"Value": 30},
                        }
                    ]
                },
            }
        )
        client.settings[(".appconfig.featureflag/gradual", None)] = setting(
            ".appconfig.featureflag/gradual",
            payload,
            "flag1",
        )
        evaluator = FeatureFlagEvaluator(ConfigurationService(client))

        first = evaluator.is_enabled("gradual", user_id="alice")
        self.assertEqual(first, evaluator.is_enabled("gradual", user_id="alice"))
        outcomes = {
            evaluator.is_enabled("gradual", user_id=f"user-{index}")
            for index in range(100)
        }
        self.assertEqual({False, True}, outcomes)

    def test_invalid_percentage_fails_explicitly(self) -> None:
        client = FakeSyncClient()
        payload = json.dumps(
            {
                "id": "bad",
                "enabled": True,
                "conditions": {
                    "client_filters": [
                        {
                            "name": "Microsoft.Percentage",
                            "parameters": {"Value": 101},
                        }
                    ]
                },
            }
        )
        client.settings[(".appconfig.featureflag/bad", None)] = setting(
            ".appconfig.featureflag/bad",
            payload,
            "flag1",
        )
        evaluator = FeatureFlagEvaluator(ConfigurationService(client))

        with self.assertRaises(FeatureFlagFormatError):
            evaluator.is_enabled("bad", user_id="alice")

    def test_sentinel_change_triggers_full_refresh(self) -> None:
        values = iter(["v1", "v2"])

        class FakeService:
            refreshes = 0

            def get_setting(self, key, label=None):
                return next(values)

            def refresh_all(self):
                self.refreshes += 1

        service = FakeService()
        watcher = ConfigurationWatcher(service, [SentinelKey("sentinel")], 0.01)

        self.assertEqual(set(), watcher.poll_once())
        self.assertEqual({SentinelKey("sentinel")}, watcher.poll_once())
        self.assertEqual(1, service.refreshes)


class AsyncPage:
    def __init__(self, items) -> None:
        self._items = iter(items)

    def __aiter__(self):
        return self

    async def __anext__(self):
        try:
            return next(self._items)
        except StopIteration as error:
            raise StopAsyncIteration from error


class FakeAsyncClient(FakeSyncClient):
    async def get_configuration_setting(self, **kwargs):
        return super().get_configuration_setting(**kwargs)

    def list_configuration_settings(self, **kwargs):
        return AsyncPage(super().list_configuration_settings(**kwargs))


class AsyncTests(unittest.IsolatedAsyncioTestCase):
    async def test_async_service_and_evaluator(self) -> None:
        client = FakeAsyncClient()
        payload = json.dumps(
            {
                "id": "always",
                "enabled": True,
                "conditions": {"client_filters": []},
            }
        )
        client.settings[(".appconfig.featureflag/always", "production")] = setting(
            ".appconfig.featureflag/always",
            payload,
            "f1",
            "production",
        )
        service = AsyncConfigurationService(client)
        evaluator = AsyncFeatureFlagEvaluator(service)

        self.assertTrue(await evaluator.is_enabled("always", label="production"))
        self.assertTrue(await evaluator.is_enabled("always", label="production"))
        self.assertIs(
            MatchConditions.IfModified,
            client.get_calls[1]["match_condition"],
        )

    async def test_async_watcher_refreshes(self) -> None:
        class FakeService:
            def __init__(self) -> None:
                self.values = iter(["v1", "v2"])
                self.refreshes = 0

            async def get_setting(self, key, label=None):
                return next(self.values)

            async def refresh_all(self):
                self.refreshes += 1

        service = FakeService()
        watcher = AsyncConfigurationWatcher(
            service,
            [SentinelKey("sentinel")],
            0.01,
        )

        self.assertEqual(set(), await watcher.poll_once())
        self.assertEqual({SentinelKey("sentinel")}, await watcher.poll_once())
        self.assertEqual(1, service.refreshes)
