from __future__ import annotations

import json
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
        if not self.order_id.strip():
            raise ValueError("order_id must not be empty")
        if not self.customer_name.strip():
            raise ValueError("customer_name must not be empty")
        if not self.product.strip():
            raise ValueError("product must not be empty")
        if isinstance(self.quantity, bool) or not isinstance(self.quantity, int):
            raise TypeError("quantity must be an integer")
        if self.quantity <= 0:
            raise ValueError("quantity must be greater than zero")
        if isinstance(self.total_price, bool) or not isinstance(
            self.total_price, (int, float)
        ):
            raise TypeError("total_price must be numeric")
        if self.total_price < 0:
            raise ValueError("total_price must not be negative")
        if not isinstance(self.status, OrderStatus):
            self.status = OrderStatus(self.status)
        self.total_price = float(self.total_price)

    def to_json(self) -> str:
        payload = asdict(self)
        payload["status"] = self.status.value
        return json.dumps(payload, separators=(",", ":"))

    @classmethod
    def from_json(cls, value: str | bytes) -> Order:
        payload: Any = json.loads(value)
        if not isinstance(payload, dict):
            raise ValueError("order JSON must contain an object")

        required_fields = {
            "order_id",
            "customer_name",
            "product",
            "quantity",
            "total_price",
            "status",
        }
        missing_fields = required_fields.difference(payload)
        if missing_fields:
            missing = ", ".join(sorted(missing_fields))
            raise ValueError(f"order JSON is missing fields: {missing}")

        return cls(
            order_id=payload["order_id"],
            customer_name=payload["customer_name"],
            product=payload["product"],
            quantity=payload["quantity"],
            total_price=payload["total_price"],
            status=OrderStatus(payload["status"]),
        )
