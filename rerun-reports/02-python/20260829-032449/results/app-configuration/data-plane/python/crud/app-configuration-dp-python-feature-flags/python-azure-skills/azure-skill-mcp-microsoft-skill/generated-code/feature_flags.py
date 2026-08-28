from __future__ import annotations

import hashlib
import json
from typing import Any

from config_service import AsyncConfigurationService, ConfigurationService


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"
PERCENTAGE_FILTER_NAMES = {"Microsoft.Percentage", "Percentage"}


class FeatureFlagError(ValueError):
    pass


def _parse_flag(payload: str | None, flag_id: str) -> dict[str, Any] | None:
    if payload is None:
        return None
    try:
        flag = json.loads(payload)
    except json.JSONDecodeError as error:
        raise FeatureFlagError(f"Feature flag {flag_id!r} contains invalid JSON") from error
    if not isinstance(flag, dict):
        raise FeatureFlagError(f"Feature flag {flag_id!r} must contain a JSON object")
    return flag


def _rollout_percentage(flag: dict[str, Any], flag_id: str) -> float | None:
    conditions = flag.get("conditions", {})
    filters = conditions.get("client_filters", []) if isinstance(conditions, dict) else []
    if not isinstance(filters, list):
        raise FeatureFlagError(f"Feature flag {flag_id!r} has invalid client filters")

    for client_filter in filters:
        if not isinstance(client_filter, dict):
            raise FeatureFlagError(f"Feature flag {flag_id!r} has an invalid filter")
        if client_filter.get("name") not in PERCENTAGE_FILTER_NAMES:
            continue
        parameters = client_filter.get("parameters", {})
        try:
            percentage = float(parameters["Value"])
        except (KeyError, TypeError, ValueError) as error:
            raise FeatureFlagError(
                f"Feature flag {flag_id!r} has an invalid rollout percentage"
            ) from error
        if not 0 <= percentage <= 100:
            raise FeatureFlagError(
                f"Feature flag {flag_id!r} rollout percentage must be between 0 and 100"
            )
        return percentage
    return None


def _is_enabled_for_user(
    flag: dict[str, Any] | None, flag_id: str, user_id: str | None
) -> bool:
    if flag is None or flag.get("enabled") is not True:
        return False

    percentage = _rollout_percentage(flag, flag_id)
    if percentage is None:
        return True
    if user_id is None:
        return False

    digest = hashlib.sha256(f"{flag_id}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], "big") * 100 / 2**64
    return bucket < percentage


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self, flag_id: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_id}", label
        )
        return _is_enabled_for_user(_parse_flag(payload, flag_id), flag_id, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self, flag_id: str, user_id: str | None = None, label: str | None = None
    ) -> bool:
        payload = await self._configuration.get_setting(
            f"{FEATURE_FLAG_PREFIX}{flag_id}", label
        )
        return _is_enabled_for_user(_parse_flag(payload, flag_id), flag_id, user_id)
