from __future__ import annotations

from collections.abc import Mapping
from typing import Any


REQUEST_CHARGE_HEADER = "x-ms-request-charge"


class ConcurrencyConflictError(RuntimeError):
    """Raised when an update uses an ETag that is no longer current."""


class RequestChargeTracker:
    def __init__(self) -> None:
        self._charges: list[float] = []

    def __call__(self, headers: Mapping[str, str], _response: Any) -> None:
        charge = headers.get(REQUEST_CHARGE_HEADER)
        if charge is not None:
            self._charges.append(float(charge))

    @property
    def total(self) -> float:
        return sum(self._charges)

    def take(self) -> float:
        total = self.total
        self.clear()
        return total

    def clear(self) -> None:
        self._charges.clear()
