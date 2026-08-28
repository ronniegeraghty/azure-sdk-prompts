from __future__ import annotations

import hashlib
import json
from typing import Any

from config_service import AsyncConfigurationService, ConfigurationService


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


class FeatureFlagError(ValueError):
    pass


def _flag_key(flag_id: str) -> str:
    return flag_id if flag_id.startswith(FEATURE_FLAG_PREFIX) else FEATURE_FLAG_PREFIX + flag_id


def _parse_flag(raw_value: str | None, flag_id: str) -> dict[str, Any]:
    if raw_value is None:
        raise FeatureFlagError(f"Feature flag {flag_id!r} has no JSON value")
    try:
        flag = json.loads(raw_value)
    except json.JSONDecodeError as exc:
        raise FeatureFlagError(f"Feature flag {flag_id!r} contains invalid JSON") from exc
    if not isinstance(flag, dict):
        raise FeatureFlagError(f"Feature flag {flag_id!r} must contain a JSON object")
    return flag


def _percentage(flag: dict[str, Any], flag_id: str) -> float | None:
    conditions = flag.get("conditions", {})
    if not isinstance(conditions, dict):
        raise FeatureFlagError(f"Feature flag {flag_id!r} has invalid conditions")
    filters = conditions.get("client_filters", [])
    if not isinstance(filters, list):
        raise FeatureFlagError(f"Feature flag {flag_id!r} has invalid client_filters")

    for client_filter in filters:
        if not isinstance(client_filter, dict):
            continue
        if client_filter.get("name") not in PERCENTAGE_FILTER_NAMES:
            continue
        parameters = client_filter.get("parameters", {})
        if not isinstance(parameters, dict):
            raise FeatureFlagError(f"Feature flag {flag_id!r} has invalid percentage parameters")
        value = parameters.get("Value", parameters.get("value"))
        try:
            percentage = float(value)
        except (TypeError, ValueError) as exc:
            raise FeatureFlagError(
                f"Feature flag {flag_id!r} has an invalid rollout percentage"
            ) from exc
        if not 0 <= percentage <= 100:
            raise FeatureFlagError(
                f"Feature flag {flag_id!r} rollout percentage must be between 0 and 100"
            )
        return percentage
    return None


def _is_in_rollout(flag_id: str, user_id: str, percentage: float) -> bool:
    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], byteorder="big") / 2**64
    return bucket < percentage / 100


def _evaluate(flag: dict[str, Any], flag_id: str, user_id: str | None) -> bool:
    if flag.get("enabled") is not True:
        return False
    percentage = _percentage(flag, flag_id)
    if percentage is None:
        return True
    if user_id is None:
        raise ValueError(f"user_id is required for percentage flag {flag_id!r}")
    return _is_in_rollout(flag_id, user_id, percentage)


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self,
        flag_id: str,
        user_id: str | None = None,
        *,
        label: str | None = None,
    ) -> bool:
        raw_value = self._configuration.get_setting(_flag_key(flag_id), label)
        return _evaluate(_parse_flag(raw_value, flag_id), flag_id, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self,
        flag_id: str,
        user_id: str | None = None,
        *,
        label: str | None = None,
    ) -> bool:
        raw_value = await self._configuration.get_setting(_flag_key(flag_id), label)
        return _evaluate(_parse_flag(raw_value, flag_id), flag_id, user_id)
