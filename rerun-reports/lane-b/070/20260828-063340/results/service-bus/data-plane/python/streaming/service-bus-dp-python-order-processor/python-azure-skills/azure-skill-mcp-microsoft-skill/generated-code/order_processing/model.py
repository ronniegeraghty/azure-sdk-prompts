"""Order domain model."""

from __future__ import annotations

import json
import math
from dataclasses import asdict, dataclass
from enum import Enum
from typing import Any


class OrderStatus(str, Enum):
    PENDING = "pending"
    PROCESSING = "processing"
    COMPLETED = "completed"
    FAILED = "failed"


@dataclass(slots=True)
class Order:
    order_id: str
    customer_name: str
    product: str
    quantity: int
    total_price: float
    status: OrderStatus = OrderStatus.PENDING

    def __post_init__(self) -> None:
        if not isinstance(self.order_id, str) or not self.order_id.strip():
            raise ValueError("order_id must be a non-empty string")
        if not isinstance(self.customer_name, str) or not self.customer_name.strip():
            raise ValueError("customer_name must be a non-empty string")
        if not isinstance(self.product, str) or not self.product.strip():
            raise ValueError("product must be a non-empty string")
        if isinstance(self.quantity, bool) or not isinstance(self.quantity, int):
            raise ValueError("quantity must be an integer")
        if self.quantity <= 0:
            raise ValueError("quantity must be greater than zero")
        if isinstance(self.total_price, bool) or not isinstance(
            self.total_price, (int, float)
        ):
            raise ValueError("total_price must be numeric")
        self.total_price = float(self.total_price)
        if not math.isfinite(self.total_price) or self.total_price < 0:
            raise ValueError("total_price must be a finite, non-negative number")
        if isinstance(self.status, str):
            try:
                self.status = OrderStatus(self.status)
            except ValueError as exc:
                raise ValueError(f"invalid order status: {self.status}") from exc
        elif not isinstance(self.status, OrderStatus):
            raise ValueError("status must be a valid OrderStatus")

    def to_dict(self) -> dict[str, Any]:
        data = asdict(self)
        data["status"] = self.status.value
        return data

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), separators=(",", ":"), ensure_ascii=True)

    @classmethod
    def from_json(cls, payload: str | bytes) -> Order:
        try:
            data = json.loads(payload)
        except (json.JSONDecodeError, UnicodeDecodeError) as exc:
            raise ValueError("order payload is not valid JSON") from exc
        if not isinstance(data, dict):
            raise ValueError("order payload must be a JSON object")

        required_fields = {
            "order_id",
            "customer_name",
            "product",
            "quantity",
            "total_price",
            "status",
        }
        missing = required_fields - data.keys()
        extra = data.keys() - required_fields
        if missing:
            raise ValueError(f"missing order fields: {', '.join(sorted(missing))}")
        if extra:
            raise ValueError(f"unexpected order fields: {', '.join(sorted(extra))}")

        try:
            return cls(**data)
        except (TypeError, ValueError) as exc:
            raise ValueError(f"invalid order payload: {exc}") from exc
