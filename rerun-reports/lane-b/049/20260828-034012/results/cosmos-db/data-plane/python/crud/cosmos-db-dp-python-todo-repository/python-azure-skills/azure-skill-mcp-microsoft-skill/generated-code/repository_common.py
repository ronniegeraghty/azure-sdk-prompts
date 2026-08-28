from __future__ import annotations

import logging
from collections.abc import Mapping
from typing import Any


class RequestChargeTracker:
    def __init__(self, logger: logging.Logger, operation: str) -> None:
        self._logger = logger
        self._operation = operation
        self.last_charge = 0.0

    def response_hook(
        self, headers: Mapping[str, str], _response: dict[str, Any] | None
    ) -> None:
        self.last_charge = float(headers.get("x-ms-request-charge", 0.0))
        self._logger.info(
            "Cosmos DB operation=%s request_charge=%.2f RU",
            self._operation,
            self.last_charge,
        )
