from __future__ import annotations

import asyncio
import json
import unittest
from collections.abc import AsyncIterator, Iterator
from typing import Any

from azure.appconfiguration import ConfigurationSetting
from azure.core import MatchConditions
from azure.core.exceptions import ResourceNotFoundError

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher
from feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator


def setting(
    key: str, value: str, etag: str, label: str | None = None
) -> ConfigurationSetting:
    return ConfigurationSetting(key=key, value=value, etag=etag, label=label)


class FakeCredential:
    pass


class FakeClient:
    def __init__(self, values: list[ConfigurationSetting]) -> None:
        self.values = {(item.key, item.label): item for item in values}
        self.get_calls: list[dict[str, Any]] = []
        self.closed = False

    def get_configuration_setting(self, **kwargs: Any) -> ConfigurationSetting | None:
        self.get_calls.append(kwargs)
        item = self.values.get((kwargs["key"], kwargs.get("label")))
        if item is None:
            raise ResourceNotFoundError("missing")
        if (
            kwargs.get("match_condition") == MatchConditions.IfModified
            and kwargs.get("etag") == item.etag
        ):
            return None
        return item

    def list_configuration_settings(
        self, **kwargs: Any
    ) -> Iterator[ConfigurationSetting]:
        prefix = kwargs["key_filter"][:-1]
        label = kwargs.get("label_filter")
        for (key, item_label), item in self.values.items():
            if key.startswith(prefix) and item_label == label:
                yield ConfigurationSetting(
                    key=item.key,
                    label=item.label,
                    etag=item.etag,
                )

    def close(self) -> None:
        self.closed = True


class FakeAsyncPage:
    def __init__(self, values: list[ConfigurationSetting]) -> None:
        self._values = values

    async def __aiter__(self) -> AsyncIterator[ConfigurationSetting]:
        for value in self._values:
            yield value


class FakeAsyncClient(FakeClient):
    async def get_configuration_setting(
        self, **kwargs: Any
    ) -> ConfigurationSetting | None:
        return super().get_configuration_setting(**kwargs)

    def list_configuration_settings(self, **kwargs: Any) -> FakeAsyncPage:
        return FakeAsyncPage(list(super().list_configuration_settings(**kwargs)))

    async def close(self) -> None:
        self.closed = True


class ConfigurationServiceTests(unittest.TestCase):
    def setUp(self) -> None:
        self.client = FakeClient(
            [
                setting("App:Message", "hello", "1", "production"),
                setting("App:Color", "blue", "2", "production"),
            ]
        )
        self.service = ConfigurationService(
            "https://example.azconfig.io",
            FakeCredential(),  # type: ignore[arg-type]
            client=self.client,  # type: ignore[arg-type]
        )

    def test_get_uses_etag_after_first_read(self) -> None:
        self.assertEqual(
            self.service.get_setting("App:Message", "production"), "hello"
        )
        self.assertEqual(
            self.service.get_setting("App:Message", "production"), "hello"
        )
        self.assertEqual(
            self.client.get_calls[-1]["match_condition"],
            MatchConditions.IfModified,
        )
        self.assertEqual(self.client.get_calls[-1]["etag"], "1")

    def test_prefix_read_only_fetches_changed_values(self) -> None:
        expected = {"App:Message": "hello", "App:Color": "blue"}
        self.assertEqual(self.service.list_settings("App:", "production"), expected)
        initial_get_count = len(self.client.get_calls)
        self.assertEqual(self.service.list_settings("App:", "production"), expected)
        self.assertEqual(len(self.client.get_calls), initial_get_count)

        self.client.values[("App:Color", "production")] = setting(
            "App:Color", "green", "3", "production"
        )
        expected["App:Color"] = "green"
        self.assertEqual(self.service.list_settings("App:", "production"), expected)
        self.assertEqual(len(self.client.get_calls), initial_get_count + 1)

    def test_watcher_refreshes_when_sentinel_changes(self) -> None:
        self.client.values[("App:Sentinel", "production")] = setting(
            "App:Sentinel", "1", "s1", "production"
        )
        watcher = ConfigurationWatcher(
            self.service, ["App:Sentinel"], 1, label="production"
        )
        self.assertEqual(watcher.poll_once(), [])
        self.client.values[("App:Sentinel", "production")] = setting(
            "App:Sentinel", "2", "s2", "production"
        )
        self.assertEqual(watcher.poll_once(), ["App:Sentinel"])

    def test_feature_flag_percentage_is_deterministic(self) -> None:
        payload = json.dumps(
            {
                "id": "BetaFeature",
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
        self.client.values[(".appconfig.featureflag/BetaFeature", None)] = setting(
            ".appconfig.featureflag/BetaFeature", payload, "f1"
        )
        evaluator = FeatureFlagEvaluator(self.service)
        first = evaluator.is_enabled("BetaFeature", "alice")
        self.assertEqual(first, evaluator.is_enabled("BetaFeature", "alice"))


class AsyncProjectTests(unittest.IsolatedAsyncioTestCase):
    async def asyncSetUp(self) -> None:
        payload = json.dumps(
            {
                "id": "BetaFeature",
                "enabled": True,
                "conditions": {
                    "client_filters": [
                        {
                            "name": "Microsoft.Percentage",
                            "parameters": {"Value": 100},
                        }
                    ]
                },
            }
        )
        self.client = FakeAsyncClient(
            [
                setting("App:Message", "hello", "1", "production"),
                setting("App:Sentinel", "1", "s1", "production"),
                setting(".appconfig.featureflag/BetaFeature", payload, "f1"),
            ]
        )
        self.service = AsyncConfigurationService(
            "https://example.azconfig.io",
            FakeCredential(),  # type: ignore[arg-type]
            client=self.client,  # type: ignore[arg-type]
        )

    async def test_async_service_flag_and_watcher(self) -> None:
        self.assertEqual(
            await self.service.get_setting("App:Message", "production"), "hello"
        )
        self.assertEqual(
            await self.service.get_setting("App:Message", "production"), "hello"
        )

        evaluator = AsyncFeatureFlagEvaluator(self.service)
        self.assertTrue(await evaluator.is_enabled("BetaFeature", "alice"))

        watcher = AsyncConfigurationWatcher(
            self.service, ["App:Sentinel"], 1, label="production"
        )
        self.assertEqual(await watcher.poll_once(), [])
        self.client.values[("App:Sentinel", "production")] = setting(
            "App:Sentinel", "2", "s2", "production"
        )
        self.assertEqual(await watcher.poll_once(), ["App:Sentinel"])


if __name__ == "__main__":
    unittest.main()
