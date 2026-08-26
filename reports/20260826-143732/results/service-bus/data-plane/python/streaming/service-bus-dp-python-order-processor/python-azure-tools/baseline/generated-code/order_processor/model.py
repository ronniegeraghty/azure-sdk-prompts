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


@dataclass
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
        if self.quantity <= 0:
            raise ValueError("quantity must be greater than zero")
        if self.total_price < 0:
            raise ValueError("total_price must not be negative")
        if not isinstance(self.status, OrderStatus):
            self.status = OrderStatus(self.status)

    def to_json(self) -> str:
        data = asdict(self)
        data["status"] = self.status.value
        return json.dumps(data, separators=(",", ":"))

    @classmethod
    def from_json(cls, payload: str | bytes) -> Order:
        data: Any = json.loads(payload)
        if not isinstance(data, dict):
            raise ValueError("order JSON must contain an object")

        required = {
            "order_id",
            "customer_name",
            "product",
            "quantity",
            "total_price",
            "status",
        }
        missing = required.difference(data)
        if missing:
            raise ValueError(f"order JSON is missing fields: {', '.join(sorted(missing))}")

        return cls(
            order_id=str(data["order_id"]),
            customer_name=str(data["customer_name"]),
            product=str(data["product"]),
            quantity=int(data["quantity"]),
            total_price=float(data["total_price"]),
            status=OrderStatus(data["status"]),
        )
