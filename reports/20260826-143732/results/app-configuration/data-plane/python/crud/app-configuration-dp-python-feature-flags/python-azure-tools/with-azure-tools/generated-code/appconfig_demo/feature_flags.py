"""Feature flag evaluation for Azure App Configuration JSON payloads."""

from __future__ import annotations

import hashlib
import json
from collections.abc import Mapping
from typing import Any

from .configuration_service import AsyncConfigurationService, ConfigurationService

FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = frozenset(
    {
        "microsoft.percentage",
        "percentage",
    }
)


class FeatureFlagFormatError(ValueError):
    """Raised when an App Configuration feature flag payload is malformed."""


def _parse_flag(payload: str | None, flag_name: str) -> Mapping[str, Any]:
    if payload is None:
        raise FeatureFlagFormatError(f"feature flag {flag_name!r} has no JSON payload")
    try:
        parsed = json.loads(payload)
    except json.JSONDecodeError as error:
        raise FeatureFlagFormatError(
            f"feature flag {flag_name!r} contains invalid JSON"
        ) from error
    if not isinstance(parsed, dict):
        raise FeatureFlagFormatError(
            f"feature flag {flag_name!r} must contain a JSON object"
        )
    return parsed


def _percentage(parameters: object, flag_name: str) -> float:
    if not isinstance(parameters, dict):
        raise FeatureFlagFormatError(
            f"percentage filter for {flag_name!r} must have parameters"
        )
    raw_value = parameters.get("Value", parameters.get("value"))
    if isinstance(raw_value, bool):
        raise FeatureFlagFormatError(
            f"percentage for {flag_name!r} must be a number from 0 to 100"
        )
    try:
        value = float(raw_value)
    except (TypeError, ValueError) as error:
        raise FeatureFlagFormatError(
            f"percentage for {flag_name!r} must be a number from 0 to 100"
        ) from error
    if not 0 <= value <= 100:
        raise FeatureFlagFormatError(
            f"percentage for {flag_name!r} must be between 0 and 100"
        )
    return value


def _is_in_rollout(flag_id: str, user_id: str, percentage: float) -> bool:
    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], byteorder="big") % 10_000
    return bucket < round(percentage * 100)


def _evaluate(payload: str | None, requested_name: str, user_id: str | None) -> bool:
    flag = _parse_flag(payload, requested_name)
    if flag.get("enabled") is not True:
        return False

    flag_id = flag.get("id", requested_name)
    if not isinstance(flag_id, str):
        raise FeatureFlagFormatError(f"feature flag {requested_name!r} has an invalid id")

    conditions = flag.get("conditions", {})
    if not isinstance(conditions, dict):
        raise FeatureFlagFormatError(
            f"feature flag {requested_name!r} has invalid conditions"
        )
    filters = conditions.get("client_filters", [])
    if not isinstance(filters, list):
        raise FeatureFlagFormatError(
            f"feature flag {requested_name!r} has invalid client_filters"
        )
    if not filters:
        return True

    for client_filter in filters:
        if not isinstance(client_filter, dict):
            raise FeatureFlagFormatError(
                f"feature flag {requested_name!r} contains an invalid filter"
            )
        name = client_filter.get("name")
        if isinstance(name, str) and name.casefold() in PERCENTAGE_FILTER_NAMES:
            if user_id is None:
                return False
            rollout = _percentage(client_filter.get("parameters"), requested_name)
            if _is_in_rollout(flag_id, user_id, rollout):
                return True

    # Client filters have OR semantics; unsupported filters fail closed.
    return False


class FeatureFlagEvaluator:
    """Evaluate feature flags through the sync configuration service."""

    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self,
        flag_name: str,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        if not flag_name:
            raise ValueError("flag_name must not be empty")
        payload = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}",
            label,
        )
        return _evaluate(payload, flag_name, user_id)


class AsyncFeatureFlagEvaluator:
    """Evaluate feature flags through the async configuration service."""

    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self,
        flag_name: str,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        if not flag_name:
            raise ValueError("flag_name must not be empty")
        payload = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}",
            label,
        )
        return _evaluate(payload, flag_name, user_id)
