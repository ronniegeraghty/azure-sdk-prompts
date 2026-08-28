from __future__ import annotations

import hashlib
import json
from typing import Any, Dict, Optional

from azure.core.exceptions import ResourceNotFoundError

from config_service import AsyncConfigurationService, ConfigurationService


FEATURE_FLAG_PREFIX = ".appconfig.featureflag/"


def _percentage_from_flag(flag: Dict[str, Any]) -> Optional[float]:
    conditions = flag.get("conditions") or {}
    filters = conditions.get("client_filters") or []
    for client_filter in filters:
        name = str(client_filter.get("name", "")).rsplit(".", 1)[-1].lower()
        if name != "percentage":
            continue

        parameters = client_filter.get("parameters") or {}
        raw_percentage = parameters.get("Value", parameters.get("value"))
        try:
            percentage = float(raw_percentage)
        except (TypeError, ValueError) as exc:
            raise ValueError("Percentage feature filter has an invalid Value") from exc
        if not 0 <= percentage <= 100:
            raise ValueError("Percentage feature filter Value must be between 0 and 100")
        return percentage
    return None


def _is_enabled(payload: str, flag_name: str, user_id: Optional[str]) -> bool:
    try:
        flag = json.loads(payload)
    except json.JSONDecodeError as exc:
        raise ValueError(f"Feature flag {flag_name!r} contains invalid JSON") from exc
    if not isinstance(flag, dict):
        raise ValueError(f"Feature flag {flag_name!r} must contain a JSON object")
    if not flag.get("enabled", False):
        return False

    percentage = _percentage_from_flag(flag)
    if percentage is None:
        return True
    if user_id is None:
        return False

    digest = hashlib.sha256(f"{flag_name}:{user_id}".encode("utf-8")).digest()
    bucket = int.from_bytes(digest[:8], "big") % 10_000
    return bucket < round(percentage * 100)


class FeatureFlagEvaluator:
    def __init__(self, configuration: ConfigurationService) -> None:
        self._configuration = configuration

    def is_enabled(
        self,
        flag_name: str,
        user_id: Optional[str] = None,
        label: Optional[str] = None,
    ) -> bool:
        try:
            payload = self._configuration.get_setting(
                f"{FEATURE_FLAG_PREFIX}{flag_name}",
                label,
            )
        except ResourceNotFoundError:
            return False
        if payload is None:
            return False
        return _is_enabled(payload, flag_name, user_id)


class AsyncFeatureFlagEvaluator:
    def __init__(self, configuration: AsyncConfigurationService) -> None:
        self._configuration = configuration

    async def is_enabled(
        self,
        flag_name: str,
        user_id: Optional[str] = None,
        label: Optional[str] = None,
    ) -> bool:
        try:
            payload = await self._configuration.get_setting(
                f"{FEATURE_FLAG_PREFIX}{flag_name}",
                label,
            )
        except ResourceNotFoundError:
            return False
        if payload is None:
            return False
        return _is_enabled(payload, flag_name, user_id)
