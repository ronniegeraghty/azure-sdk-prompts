from __future__ import annotations

import hashlib
import json
from typing import Any, Protocol


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
_PERCENTAGE_FILTER_NAMES = {"percentage", "random"}


class ConfigurationReader(Protocol):
    def get_setting(self, key: str, label: str | None = None) -> str | None: ...


class AsyncConfigurationReader(Protocol):
    async def get_setting(
        self, key: str, label: str | None = None
    ) -> str | None: ...


def _percentage_bucket(flag_name: str, user_id: str) -> float:
    digest = hashlib.sha256(f"{flag_name}:{user_id}".encode("utf-8")).digest()
    value = int.from_bytes(digest[:8], byteorder="big", signed=False)
    return value / 2**64 * 100.0


def _read_percentage(parameters: dict[str, Any]) -> float:
    normalized = {str(key).lower(): value for key, value in parameters.items()}
    raw_value = normalized.get("value", normalized.get("percentage"))
    if raw_value is None:
        raise ValueError("Percentage filter requires a 'Value' or 'Percentage' parameter")

    try:
        percentage = float(raw_value)
    except (TypeError, ValueError) as error:
        raise ValueError("Percentage filter value must be numeric") from error

    if not 0.0 <= percentage <= 100.0:
        raise ValueError("Percentage filter value must be between 0 and 100")
    return percentage


def _evaluate_payload(payload: str | None, flag_name: str, user_id: str | None) -> bool:
    if payload is None:
        return False

    try:
        flag = json.loads(payload)
    except json.JSONDecodeError as error:
        raise ValueError(f"Feature flag '{flag_name}' contains invalid JSON") from error

    if not isinstance(flag, dict):
        raise ValueError(f"Feature flag '{flag_name}' must contain a JSON object")
    if not flag.get("enabled", False):
        return False

    conditions = flag.get("conditions") or {}
    filters = conditions.get("client_filters") or []
    if not filters:
        return True
    if not isinstance(filters, list):
        raise ValueError(f"Feature flag '{flag_name}' has invalid client filters")

    evaluations: list[bool] = []
    for feature_filter in filters:
        if not isinstance(feature_filter, dict):
            raise ValueError(f"Feature flag '{flag_name}' has an invalid filter")

        filter_name = str(feature_filter.get("name", ""))
        short_name = filter_name.rsplit(".", maxsplit=1)[-1].lower()
        if short_name not in _PERCENTAGE_FILTER_NAMES:
            raise ValueError(
                f"Feature flag '{flag_name}' uses unsupported filter '{filter_name}'"
            )
        if user_id is None:
            raise ValueError(
                f"Feature flag '{flag_name}' requires a user ID for percentage rollout"
            )

        parameters = feature_filter.get("parameters") or {}
        if not isinstance(parameters, dict):
            raise ValueError(f"Feature flag '{flag_name}' has invalid filter parameters")
        percentage = _read_percentage(parameters)
        evaluations.append(_percentage_bucket(flag_name, user_id) < percentage)

    requirement_type = str(conditions.get("requirement_type", "Any")).lower()
    if requirement_type == "all":
        return all(evaluations)
    if requirement_type == "any":
        return any(evaluations)
    raise ValueError(
        f"Feature flag '{flag_name}' has unsupported requirement type "
        f"'{conditions.get('requirement_type')}'"
    )


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationReader) -> None:
        self._configuration = configuration

    def is_enabled(
        self, flag_name: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _evaluate_payload(payload, flag_name, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationReader) -> None:
        self._configuration = configuration

    async def is_enabled(
        self, flag_name: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_name}", label
        )
        return _evaluate_payload(payload, flag_name, user_id)

