import asyncio
import json
import unittest

from feature_flags import (
    AsyncFeatureFlagEvaluator,
    FeatureFlagError,
    FeatureFlagEvaluator,
    _evaluate_payload,
)


class _SyncConfig:
    def __init__(self, payload):
        self.payload = payload

    def get_setting(self, key, label=None):
        return self.payload


class _AsyncConfig:
    def __init__(self, payload):
        self.payload = payload

    async def get_setting(self, key, label=None):
        return self.payload


class FeatureFlagTests(unittest.TestCase):
    def test_simple_enabled_and_disabled_flags(self):
        self.assertTrue(_evaluate_payload('{"enabled": true}', "Flag", None))
        self.assertFalse(_evaluate_payload('{"enabled": false}', "Flag", None))

    def test_percentage_rollout_is_deterministic(self):
        payload = json.dumps(
            {
                "enabled": True,
                "conditions": {
                    "client_filters": [
                        {"name": "Microsoft.Percentage", "parameters": {"Value": 30}}
                    ]
                },
            }
        )
        first = [_evaluate_payload(payload, "Beta", "alice") for _ in range(10)]
        self.assertEqual(len(set(first)), 1)

    def test_percentage_boundaries(self):
        zero = '{"enabled":true,"conditions":{"client_filters":[{"name":"Percentage","parameters":{"Value":0}}]}}'
        full = '{"enabled":true,"conditions":{"client_filters":[{"name":"Percentage","parameters":{"Value":100}}]}}'
        self.assertFalse(_evaluate_payload(zero, "Beta", "alice"))
        self.assertTrue(_evaluate_payload(full, "Beta", "alice"))

    def test_invalid_percentage_raises(self):
        payload = '{"enabled":true,"conditions":{"client_filters":[{"name":"Percentage","parameters":{"Value":101}}]}}'
        with self.assertRaises(FeatureFlagError):
            _evaluate_payload(payload, "Beta", "alice")

    def test_sync_evaluator(self):
        evaluator = FeatureFlagEvaluator(_SyncConfig('{"enabled": true}'))
        self.assertTrue(evaluator.is_enabled("Flag"))

    def test_async_evaluator(self):
        evaluator = AsyncFeatureFlagEvaluator(_AsyncConfig('{"enabled": true}'))
        self.assertTrue(asyncio.run(evaluator.is_enabled("Flag")))


if __name__ == "__main__":
    unittest.main()
