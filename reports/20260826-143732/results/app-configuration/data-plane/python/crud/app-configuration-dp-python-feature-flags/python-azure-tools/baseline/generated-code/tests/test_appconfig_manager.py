from __future__ import annotations

import asyncio
import json
import unittest
from types import SimpleNamespace
from typing import Any

from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError

from appconfig_manager.configuration import (
    AsyncConfigurationService,
    ConfigurationService,
)
from appconfig_manager.feature_flags import (
    AsyncFeatureFlagEvaluator,
    FeatureFlagEvaluator,
)


def _setting(
    key: str, value: str | None, etag: str, label: str | None = None
) -> SimpleNamespace:
    return SimpleNamespace(key=key, value=value, etag=etag, label=label)


class FakeClient:
    def __init__(self, settings: list[SimpleNamespace]) -> None:
        self.settings = {(item.key, item.label): item for item in settings}
        self.get_calls = 0
        self.list_calls: list[list[str] | None] = []

    def get_configuration_setting(self, **kwargs: Any) -> SimpleNamespace:
        self.get_calls += 1
        item = self.settings.get((kwargs["key"], kwargs["label"]))
        if item is None:
            raise ResourceNotFoundError("missing")
        if kwargs.get("etag") == item.etag:
            raise ResourceNotModifiedError("unchanged")
        return item

    def list_configuration_settings(self, **kwargs: Any) -> list[SimpleNamespace]:
        fields = kwargs.get("fields")
        self.list_calls.append(fields)
        prefix = kwargs["key_filter"][:-1]
        label = None if kwargs["label_filter"] == "\0" else kwargs["label_filter"]
        matches = [
            item
            for (key, item_label), item in self.settings.items()
            if key.startswith(prefix) and item_label == label
        ]
        if fields:
            return [
                _setting(item.key, None, item.etag, item.label) for item in matches
            ]
        return matches


class AsyncItems:
    def __init__(self, items: list[SimpleNamespace]) -> None:
        self._items = items

    def __aiter__(self) -> AsyncItems:
        self._iterator = iter(self._items)
        return self

    async def __anext__(self) -> SimpleNamespace:
        try:
            return next(self._iterator)
        except StopIteration as exc:
            raise StopAsyncIteration from exc


class FakeAsyncClient(FakeClient):
    async def get_configuration_setting(self, **kwargs: Any) -> SimpleNamespace:
        return super().get_configuration_setting(**kwargs)

    def list_configuration_settings(self, **kwargs: Any) -> AsyncItems:
        return AsyncItems(super().list_configuration_settings(**kwargs))


def _percentage_flag(value: float) -> str:
    return json.dumps(
        {
            "id": "rollout",
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


class ConfigurationServiceTests(unittest.TestCase):
    def test_get_uses_cache_and_conditional_refresh(self) -> None:
        client = FakeClient([_setting("Api:Url", "one", "1")])
        service = ConfigurationService(client)

        self.assertEqual(service.get_setting("Api:Url"), "one")
        self.assertEqual(service.get_setting("Api:Url"), "one")
        self.assertEqual(client.get_calls, 1)
        self.assertFalse(service.check_for_update("Api:Url"))
        self.assertEqual(client.get_calls, 2)

        client.settings[("Api:Url", None)] = _setting("Api:Url", "two", "2")
        self.assertTrue(service.check_for_update("Api:Url"))
        self.assertEqual(service.get_setting("Api:Url"), "two")

    def test_prefix_refresh_fetches_only_changed_values(self) -> None:
        client = FakeClient(
            [_setting("App:A", "a", "1"), _setting("App:B", "b", "1")]
        )
        service = ConfigurationService(client)

        self.assertEqual(service.list_settings("App:"), {"App:A": "a", "App:B": "b"})
        client.settings[("App:B", None)] = _setting("App:B", "new", "2")
        self.assertEqual(
            service.list_settings("App:", refresh=True),
            {"App:A": "a", "App:B": "new"},
        )
        self.assertEqual(client.list_calls, [None, ["key", "label", "etag"]])
        self.assertEqual(client.get_calls, 1)

    def test_percentage_rollout_is_deterministic(self) -> None:
        client = FakeClient(
            [_setting(".appconfig.featureflag/Test", _percentage_flag(30), "1")]
        )
        evaluator = FeatureFlagEvaluator(ConfigurationService(client))

        first = evaluator.is_enabled("Test", user_id="alice")
        self.assertEqual(first, evaluator.is_enabled("Test", user_id="alice"))
        outcomes = {
            evaluator.is_enabled("Test", user_id=f"user-{number}")
            for number in range(100)
        }
        self.assertEqual(outcomes, {False, True})


class AsyncConfigurationServiceTests(unittest.IsolatedAsyncioTestCase):
    async def test_async_cache_refresh_and_rollout(self) -> None:
        client = FakeAsyncClient(
            [
                _setting("Api:Url", "one", "1", "production"),
                _setting(
                    ".appconfig.featureflag/Test",
                    _percentage_flag(100),
                    "1",
                    "production",
                ),
            ]
        )
        service = AsyncConfigurationService(client)
        evaluator = AsyncFeatureFlagEvaluator(service)

        self.assertEqual(
            await service.get_setting_with_label("Api:Url", "production"), "one"
        )
        self.assertFalse(await service.check_for_update("Api:Url", "production"))
        self.assertTrue(
            await evaluator.is_enabled(
                "Test", user_id="alice", label="production"
            )
        )

        client.settings[("Api:Url", "production")] = _setting(
            "Api:Url", "two", "2", "production"
        )
        self.assertTrue(await service.check_for_update("Api:Url", "production"))
        self.assertEqual(
            await service.get_setting("Api:Url", "production"), "two"
        )


if __name__ == "__main__":
    unittest.main()
