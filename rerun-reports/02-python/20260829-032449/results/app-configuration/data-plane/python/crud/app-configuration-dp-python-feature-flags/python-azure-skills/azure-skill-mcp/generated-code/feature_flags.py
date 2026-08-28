from __future__ import annotations

import hashlib
import json
from dataclasses import dataclass
from typing import Any, Optional

from azure.core.exceptions import ResourceNotFoundError

from configuration_service import AsyncConfigurationService, ConfigurationService


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


@dataclass(frozen=True)
class _FeatureFlag:
    flag_id: str
    enabled: bool
    rollout_percentage: Optional[float]


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self, flag_id: str, user_id: Optional[str] = None, label: Optional[str] = None
    ) -> bool:
        if not flag_id:
            raise ValueError("flag_id must not be empty")

        try:
            payload = self._configuration.get_setting(
                f"{FEATURE_FLAG_PREFIX}{flag_id}", label
            )
        except ResourceNotFoundError:
            return False
        return _evaluate(_parse_feature_flag(payload, flag_id), user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self, flag_id: str, user_id: Optional[str] = None, label: Optional[str] = None
    ) -> bool:
        if not flag_id:
            raise ValueError("flag_id must not be empty")

        try:
            payload = await self._configuration.get_setting(
                f"{FEATURE_FLAG_PREFIX}{flag_id}", label
            )
        except ResourceNotFoundError:
            return False
        return _evaluate(_parse_feature_flag(payload, flag_id), user_id)


def _parse_feature_flag(payload: str, requested_flag_id: str) -> _FeatureFlag:
    try:
        document: Any = json.loads(payload)
    except json.JSONDecodeError as error:
        raise ValueError(
            f"Feature flag {requested_flag_id!r} contains invalid JSON"
        ) from error

    if not isinstance(document, dict):
        raise ValueError(f"Feature flag {requested_flag_id!r} must be a JSON object")

    stored_id = document.get("id", requested_flag_id)
    if not isinstance(stored_id, str):
        raise ValueError(f"Feature flag {requested_flag_id!r} has an invalid id")

    enabled = document.get("enabled", False)
    if not isinstance(enabled, bool):
        raise ValueError(f"Feature flag {requested_flag_id!r} has an invalid enabled value")

    percentage: Optional[float] = None
    conditions = document.get("conditions", {})
    if not isinstance(conditions, dict):
        raise ValueError(f"Feature flag {requested_flag_id!r} has invalid conditions")

    filters = conditions.get("client_filters", [])
    if not isinstance(filters, list):
        raise ValueError(f"Feature flag {requested_flag_id!r} has invalid client filters")

    for client_filter in filters:
        if not isinstance(client_filter, dict):
            continue
        if client_filter.get("name") not in PERCENTAGE_FILTER_NAMES:
            continue

        parameters = client_filter.get("parameters", {})
        if not isinstance(parameters, dict):
            raise ValueError(
                f"Feature flag {requested_flag_id!r} has invalid percentage parameters"
            )
        raw_value = parameters.get("Value", parameters.get("value"))
        try:
            percentage = float(raw_value)
        except (TypeError, ValueError) as error:
            raise ValueError(
                f"Feature flag {requested_flag_id!r} has an invalid percentage"
            ) from error
        if not 0.0 <= percentage <= 100.0:
            raise ValueError(
                f"Feature flag {requested_flag_id!r} percentage must be 0 through 100"
            )
        break

    return _FeatureFlag(
        flag_id=stored_id,
        enabled=enabled,
        rollout_percentage=percentage,
    )


def _evaluate(flag: _FeatureFlag, user_id: Optional[str]) -> bool:
    if not flag.enabled:
        return False
    if flag.rollout_percentage is None:
        return True
    if user_id is None:
        return False

    digest = hashlib.sha256(f"{flag.flag_id}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], byteorder="big") / 2**64 * 100.0
    return bucket < flag.rollout_percentage
