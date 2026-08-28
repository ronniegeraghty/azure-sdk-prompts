from __future__ import annotations

import hashlib
import json
from typing import Any

from config_service import AsyncConfigurationService, ConfigurationService


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self, flag_id: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_id}", label
        )
        return _evaluate(flag_id, payload, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self, flag_id: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_id}", label
        )
        return _evaluate(flag_id, payload, user_id)


def _evaluate(flag_id: str, payload: str | None, user_id: str | None) -> bool:
    if payload is None:
        return False
    try:
        flag = json.loads(payload)
    except (TypeError, json.JSONDecodeError) as exc:
        raise ValueError(f"Feature flag {flag_id!r} contains invalid JSON") from exc
    if not isinstance(flag, dict):
        raise ValueError(f"Feature flag {flag_id!r} must contain a JSON object")
    if not flag.get("enabled", False):
        return False

    filters = flag.get("conditions", {}).get("client_filters", [])
    percentage = _percentage_from(filters)
    if percentage is None:
        return True
    if user_id is None:
        return False
    return _bucket(flag_id, user_id) < percentage


def _percentage_from(filters: Any) -> float | None:
    if not isinstance(filters, list):
        return None
    for item in filters:
        if not isinstance(item, dict) or item.get("name") not in PERCENTAGE_FILTER_NAMES:
            continue
        parameters = item.get("parameters", {})
        raw_value = parameters.get("Value", parameters.get("value"))
        try:
            percentage = float(raw_value)
        except (TypeError, ValueError) as exc:
            raise ValueError("Percentage feature filter requires a numeric Value") from exc
        if not 0 <= percentage <= 100:
            raise ValueError("Percentage feature filter Value must be between 0 and 100")
        return percentage
    return None


def _bucket(flag_id: str, user_id: str) -> float:
    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    return int.from_bytes(digest[:8], "big") / 2**64 * 100

