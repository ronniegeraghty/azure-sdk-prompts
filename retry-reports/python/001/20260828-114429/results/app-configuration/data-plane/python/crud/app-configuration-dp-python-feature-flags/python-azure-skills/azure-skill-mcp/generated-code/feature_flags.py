from __future__ import annotations

import hashlib
import json
from collections.abc import Mapping
from typing import Any

from config_service import AsyncConfigurationService, ConfigurationService

FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Percentage", "Microsoft.Percentage"}


class FeatureFlagError(ValueError):
    pass


def _percentage_bucket(flag_id: str, user_id: str) -> float:
    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    return int.from_bytes(digest[:8], "big") / 2**64 * 100


def _parse_percentage(parameters: Mapping[str, Any]) -> float:
    raw_value = parameters.get("Value", parameters.get("value"))
    try:
        percentage = float(raw_value)
    except (TypeError, ValueError) as exc:
        raise FeatureFlagError("Percentage filter requires a numeric Value") from exc
    if not 0 <= percentage <= 100:
        raise FeatureFlagError("Percentage filter Value must be between 0 and 100")
    return percentage


def _evaluate_payload(payload: str, flag_name: str, user_id: str | None) -> bool:
    try:
        flag = json.loads(payload)
    except json.JSONDecodeError as exc:
        raise FeatureFlagError(f"Feature flag {flag_name!r} contains invalid JSON") from exc
    if not isinstance(flag, dict):
        raise FeatureFlagError(f"Feature flag {flag_name!r} must contain a JSON object")

    if flag.get("enabled") is not True:
        return False

    conditions = flag.get("conditions")
    if not isinstance(conditions, dict):
        return True

    filters = conditions.get("client_filters", [])
    if not isinstance(filters, list):
        raise FeatureFlagError(f"Feature flag {flag_name!r} client_filters must be a list")
    if not filters:
        return conditions.get("requirement_type", "Any").lower() != "all"

    results: list[bool] = []
    for filter_config in filters:
        if not isinstance(filter_config, dict):
            raise FeatureFlagError(f"Feature flag {flag_name!r} has an invalid filter")
        filter_name = filter_config.get("name")
        if filter_name not in PERCENTAGE_FILTER_NAMES:
            results.append(False)
            continue
        if user_id is None:
            results.append(False)
            continue
        parameters = filter_config.get("parameters", {})
        if not isinstance(parameters, dict):
            raise FeatureFlagError(f"Feature flag {flag_name!r} filter parameters must be an object")
        percentage = _parse_percentage(parameters)
        results.append(_percentage_bucket(flag_name, user_id) < percentage)

    requirement_type = str(conditions.get("requirement_type", "Any")).lower()
    if requirement_type == "all":
        return all(results)
    if requirement_type != "any":
        raise FeatureFlagError(f"Feature flag {flag_name!r} has invalid requirement_type")
    return any(results)


class FeatureFlagEvaluator:
    def __init__(self, config: ConfigurationService, label: str | None = None) -> None:
        self._config = config
        self._label = label

    def is_enabled(self, flag_name: str, user_id: str | None = None) -> bool:
        payload = self._config.get_setting(FEATURE_FLAG_PREFIX + flag_name, self._label)
        return False if payload is None else _evaluate_payload(payload, flag_name, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, config: AsyncConfigurationService, label: str | None = None) -> None:
        self._config = config
        self._label = label

    async def is_enabled(self, flag_name: str, user_id: str | None = None) -> bool:
        payload = await self._config.get_setting(FEATURE_FLAG_PREFIX + flag_name, self._label)
        return False if payload is None else _evaluate_payload(payload, flag_name, user_id)
