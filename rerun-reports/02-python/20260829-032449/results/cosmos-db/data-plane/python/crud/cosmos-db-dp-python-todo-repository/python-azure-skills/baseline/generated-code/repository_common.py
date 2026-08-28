from __future__ import annotations

from typing import Mapping


class TodoConflictError(RuntimeError):
    """Raised when an update is based on a stale version of an item."""


class MissingEtagError(ValueError):
    """Raised when an item that was not read from Cosmos DB is updated."""


class RequestChargeTracker:
    def __init__(self) -> None:
        self._unreported_charge = 0.0

    def response_hook(self, headers: Mapping[str, str], _body: object) -> None:
        for name, value in headers.items():
            if name.lower() == "x-ms-request-charge":
                self._unreported_charge += float(value)
                return

    def take_charge(self) -> float:
        charge = self._unreported_charge
        self._unreported_charge = 0.0
        return charge

