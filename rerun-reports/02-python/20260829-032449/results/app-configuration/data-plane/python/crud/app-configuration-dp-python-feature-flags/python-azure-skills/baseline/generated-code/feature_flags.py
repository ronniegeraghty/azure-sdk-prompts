"""Azure App Configuration feature flag evaluation."""

from __future__ import annotations

import hashlib
import json
from typing import Any, Protocol

FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


class SyncSettingReader(Protocol):
    def get_setting(self, key: str, label: str | None = None) -> str | None: ...


class AsyncSettingReader(Protocol):
    async def get_setting(
        self, key: str, label: str | None = None
    ) -> str | None: ...


def _parse_flag(raw_value: str | None, flag_name: str) -> dict[str, Any]:
    if raw_value is None:
        raise ValueError(f"Feature flag {flag_name!r} has no JSON value")
    try:
        payload = json.loads(raw_value)
    except json.JSONDecodeError as exc:
        raise ValueError(f"Feature flag {flag_name!r} contains invalid JSON") from exc
    if not isinstance(payload, dict):
        raise ValueError(f"Feature flag {flag_name!r} must contain a JSON object")
    return payload


def _rollout_bucket(flag_name: str, user_id: str) -> int:
    digest = hashlib.sha256(f"{flag_name}:{user_id}".encode("utf-8")).digest()
    return int.from_bytes(digest[:8], "big") % 10_000


def _percentage(parameters: Any, flag_name: str) -> float:
    if not isinstance(parameters, dict):
        raise ValueError(f"Percentage filter for {flag_name!r} needs parameters")
    raw = parameters.get("Value", parameters.get("Percentage"))
    try:
        value = float(raw)
    except (TypeError, ValueError) as exc:
        raise ValueError(
            f"Percentage filter for {flag_name!r} needs a numeric Value"
        ) from exc
    if not 0 <= value <= 100:
        raise ValueError(
            f"Percentage filter for {flag_name!r} must be between 0 and 100"
        )
    return value


def _evaluate(payload: dict[str, Any], flag_name: str, user_id: str | None) -> bool:
    if payload.get("enabled") is not True:
        return False

    conditions = payload.get("conditions")
    if not conditions:
        return True
    if not isinstance(conditions, dict):
        raise ValueError(f"Feature flag {flag_name!r} has invalid conditions")

    filters = conditions.get("client_filters") or []
    if not isinstance(filters, list):
        raise ValueError(f"Feature flag {flag_name!r} has invalid client_filters")
    if not filters:
        return True

    results: list[bool] = []
    for feature_filter in filters:
        if not isinstance(feature_filter, dict):
            raise ValueError(f"Feature flag {flag_name!r} has an invalid filter")
        if feature_filter.get("name") not in PERCENTAGE_FILTER_NAMES:
            results.append(False)
            continue
        if user_id is None:
            results.append(False)
            continue
        percentage = _percentage(feature_filter.get("parameters"), flag_name)
        results.append(_rollout_bucket(flag_name, user_id) < percentage * 100)

    requirement_type = str(conditions.get("requirement_type", "Any")).lower()
    if requirement_type == "all":
        return all(results)
    if requirement_type == "any":
        return any(results)
    raise ValueError(
        f"Feature flag {flag_name!r} has unknown requirement_type "
        f"{conditions.get('requirement_type')!r}"
    )


class FeatureFlagEvaluator:
    def __init__(self, configuration: SyncSettingReader) -> None:
        self._configuration = configuration

    def is_enabled(
        self,
        flag_name: str,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        raw_value = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _evaluate(_parse_flag(raw_value, flag_name), flag_name, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncSettingReader) -> None:
        self._configuration = configuration

    async def is_enabled(
        self,
        flag_name: str,
        user_id: str | None = None,
        label: str | None = None,
    ) -> bool:
        raw_value = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _evaluate(_parse_flag(raw_value, flag_name), flag_name, user_id)
