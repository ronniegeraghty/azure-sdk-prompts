"""Feature flag evaluation, including deterministic percentage rollout."""

from __future__ import annotations

import hashlib
import json
from typing import Any

from .configuration import AsyncConfigurationService, ConfigurationService

FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
_PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


def _rollout_percentage(payload: dict[str, Any]) -> float | None:
    conditions = payload.get("conditions") or {}
    for item in conditions.get("client_filters") or []:
        if item.get("name") not in _PERCENTAGE_FILTER_NAMES:
            continue
        parameters = item.get("parameters") or {}
        raw_value = next(
            (value for key, value in parameters.items() if key.lower() == "value"),
            None,
        )
        if raw_value is None:
            raise ValueError(
                "Percentage feature filter is missing its Value parameter"
            )
        try:
            percentage = float(raw_value)
        except (TypeError, ValueError) as exc:
            raise ValueError(
                "Percentage feature filter Value must be numeric"
            ) from exc
        if not 0 <= percentage <= 100:
            raise ValueError(
                "Percentage feature filter Value must be between 0 and 100"
            )
        return percentage
    return None


def _parse_flag(flag_name: str, raw_value: str | None, user_id: str | None) -> bool:
    if raw_value is None:
        raise ValueError(f"Feature flag {flag_name!r} has no JSON value")
    try:
        payload = json.loads(raw_value)
    except json.JSONDecodeError as exc:
        raise ValueError(f"Feature flag {flag_name!r} contains invalid JSON") from exc
    if not isinstance(payload, dict):
        raise ValueError(f"Feature flag {flag_name!r} must contain a JSON object")
    if payload.get("enabled") is not True:
        return False

    percentage = _rollout_percentage(payload)
    if percentage is None:
        return True
    if user_id is None:
        raise ValueError("A user_id is required for percentage-based feature flags")

    flag_id = str(payload.get("id") or flag_name)
    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], "big") * 100 / 2**64
    return bucket < percentage


class FeatureFlagEvaluator:
    """Evaluate feature flag JSON obtained through a configuration service."""

    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self,
        flag_name: str,
        *,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        raw_value = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _parse_flag(flag_name, raw_value, user_id)


class AsyncFeatureFlagEvaluator:
    """Asynchronously evaluate feature flag JSON."""

    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self,
        flag_name: str,
        *,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        raw_value = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _parse_flag(flag_name, raw_value, user_id)
