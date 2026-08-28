import asyncio
import json
import unittest

from feature_flags import (
    AsyncFeatureFlagEvaluator,
    FeatureFlagError,
    FeatureFlagEvaluator,
)


def flag_payload(enabled=True, percentage=None):
    filters = []
    if percentage is not None:
        filters.append(
            {
                "name": "Microsoft.Percentage",
                "parameters": {"Value": percentage},
            }
        )
    return json.dumps(
        {
            "id": "test-flag",
            "enabled": enabled,
            "conditions": {"client_filters": filters},
        }
    )


class FakeConfiguration:
    def __init__(self, payload):
        self.payload = payload

    def get_setting(self, key, label=None):
        return self.payload


class AsyncFakeConfiguration:
    def __init__(self, payload):
        self.payload = payload

    async def get_setting(self, key, label=None):
        return self.payload


class FeatureFlagEvaluatorTests(unittest.TestCase):
    def test_disabled_flag_is_disabled(self):
        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(False)))
        self.assertFalse(evaluator.is_enabled("test-flag", "alice"))

    def test_enabled_flag_without_filters_is_enabled(self):
        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload()))
        self.assertTrue(evaluator.is_enabled("test-flag", "alice"))

    def test_percentage_rollout_is_deterministic(self):
        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=30)))
        first = evaluator.is_enabled("test-flag", "alice")
        self.assertEqual(first, evaluator.is_enabled("test-flag", "alice"))

    def test_rollout_boundaries(self):
        zero = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=0)))
        full = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=100)))
        self.assertFalse(zero.is_enabled("test-flag", "alice"))
        self.assertTrue(full.is_enabled("test-flag", "alice"))

    def test_invalid_percentage_raises(self):
        evaluator = FeatureFlagEvaluator(FakeConfiguration(flag_payload(percentage=101)))
        with self.assertRaises(FeatureFlagError):
            evaluator.is_enabled("test-flag", "alice")

    def test_async_evaluator(self):
        evaluator = AsyncFeatureFlagEvaluator(
            AsyncFakeConfiguration(flag_payload(percentage=100))
        )
        self.assertTrue(asyncio.run(evaluator.is_enabled("test-flag", "alice")))


if __name__ == "__main__":
    unittest.main()
