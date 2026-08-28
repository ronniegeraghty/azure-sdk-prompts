from __future__ import annotations

import asyncio
import json
import unittest
from dataclasses import dataclass

from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotModifiedError

from config_service import AsyncConfigurationService, ConfigurationService
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator
from watcher import AsyncConfigurationWatcher, ConfigurationWatcher


@dataclass
class FakeSetting:
    key: str
    value: str
    label: str | None = None
    etag: str = "1"


class FakeClient:
    def __init__(self, settings: list[FakeSetting]) -> None:
        self.settings = {(item.key, item.label): item for item in settings}
        self.get_calls: list[dict[str, object]] = []
        self.list_calls = 0

    def get_configuration_setting(self, **kwargs: object) -> FakeSetting:
        self.get_calls.append(kwargs)
        setting = self.settings[(str(kwargs["key"]), kwargs.get("label"))]
        if (
            kwargs.get("match_condition") == MatchConditions.IfModified
            and str(kwargs.get("etag")) == setting.etag
        ):
            raise ResourceNotModifiedError()
        return setting

    def list_configuration_settings(
        self, *, key_filter: str, label_filter: str | None
    ) -> list[FakeSetting]:
        self.list_calls += 1
        prefix = key_filter[:-1]
        return [
            item
            for item in self.settings.values()
            if item.key.startswith(prefix) and item.label == label_filter
        ]


class AsyncItems:
    def __init__(self, items: list[FakeSetting]) -> None:
        self._items = items

    def __aiter__(self):
        self._iterator = iter(self._items)
        return self

    async def __anext__(self) -> FakeSetting:
        try:
            return next(self._iterator)
        except StopIteration as exc:
            raise StopAsyncIteration from exc


class AsyncFakeClient(FakeClient):
    async def get_configuration_setting(self, **kwargs: object) -> FakeSetting:
        return super().get_configuration_setting(**kwargs)

    def list_configuration_settings(
        self, *, key_filter: str, label_filter: str | None
    ) -> AsyncItems:
        return AsyncItems(
            super().list_configuration_settings(
                key_filter=key_filter, label_filter=label_filter
            )
        )


def flag_payload(percentage: int) -> str:
    return json.dumps(
        {
            "id": "rollout",
            "enabled": True,
            "conditions": {
                "client_filters": [
                    {
                        "name": "Microsoft.Percentage",
                        "parameters": {"Value": percentage},
                    }
                ]
            },
        }
    )


class SyncTests(unittest.TestCase):
    def setUp(self) -> None:
        self.client = FakeClient(
            [
                FakeSetting("app:a", "A"),
                FakeSetting("app:b", "B"),
                FakeSetting("sentinel", "v1"),
                FakeSetting(".appconfig.featureflag/rollout", flag_payload(30)),
            ]
        )
        self.configuration = ConfigurationService(self.client)

    def test_conditional_reads_and_prefix_cache(self) -> None:
        self.assertEqual(self.configuration.get_setting("app:a"), "A")
        self.assertEqual(self.configuration.get_setting("app:a"), "A")
        self.assertEqual(
            self.client.get_calls[-1]["match_condition"], MatchConditions.IfModified
        )
        self.assertEqual(len(self.configuration.list_settings("app:")), 2)
        self.configuration.list_settings("app:")
        self.assertEqual(self.client.list_calls, 1)

    def test_percentage_is_deterministic(self) -> None:
        evaluator = FeatureFlagEvaluator(self.configuration)
        first = evaluator.is_enabled("rollout", "alice")
        self.assertEqual(first, evaluator.is_enabled("rollout", "alice"))

    def test_sentinel_change_refreshes_cache(self) -> None:
        self.configuration.get_setting("app:a")
        watcher = ConfigurationWatcher(self.configuration, ["sentinel"], 1)
        watcher._values = {("sentinel", None): "v1"}
        self.client.settings[("sentinel", None)] = FakeSetting(
            "sentinel", "v2", etag="2"
        )
        self.assertTrue(watcher._sentinel_changed())


class AsyncTests(unittest.IsolatedAsyncioTestCase):
    async def asyncSetUp(self) -> None:
        self.client = AsyncFakeClient(
            [
                FakeSetting("app:a", "A"),
                FakeSetting("app:b", "B"),
                FakeSetting("sentinel", "v1"),
                FakeSetting(".appconfig.featureflag/rollout", flag_payload(30)),
            ]
        )
        self.configuration = AsyncConfigurationService(self.client)

    async def test_async_service_and_evaluator(self) -> None:
        self.assertEqual(await self.configuration.get_setting("app:a"), "A")
        self.assertEqual(len(await self.configuration.list_settings("app:")), 2)
        evaluator = AsyncFeatureFlagEvaluator(self.configuration)
        first = await evaluator.is_enabled("rollout", "alice")
        self.assertEqual(first, await evaluator.is_enabled("rollout", "alice"))

    async def test_async_watcher_detects_change(self) -> None:
        await self.configuration.get_setting("sentinel")
        watcher = AsyncConfigurationWatcher(self.configuration, ["sentinel"], 1)
        watcher._values = {("sentinel", None): "v1"}
        self.client.settings[("sentinel", None)] = FakeSetting(
            "sentinel", "v2", etag="2"
        )
        self.assertTrue(await watcher._sentinel_changed())


if __name__ == "__main__":
    unittest.main()
