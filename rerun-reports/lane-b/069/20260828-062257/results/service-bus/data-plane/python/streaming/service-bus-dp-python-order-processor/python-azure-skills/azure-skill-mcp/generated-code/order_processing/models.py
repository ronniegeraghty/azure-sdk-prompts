from __future__ import annotations

import json
from dataclasses import asdict, dataclass
from decimal import Decimal
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
    total_price: Decimal
    status: OrderStatus = OrderStatus.PENDING

    def __post_init__(self) -> None:
        if not self.order_id.strip():
            raise ValueError("order_id must not be empty")
        if not self.customer_name.strip():
            raise ValueError("customer_name must not be empty")
        if self.quantity <= 0:
            raise ValueError("quantity must be greater than zero")
        if self.total_price < 0:
            raise ValueError("total_price must not be negative")

    def to_dict(self) -> dict[str, Any]:
        data = asdict(self)
        data["total_price"] = str(self.total_price)
        data["status"] = self.status.value
        return data

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), separators=(",", ":"))

    @classmethod
    def from_dict(cls, data: dict[str, Any]) -> Order:
        return cls(
            order_id=str(data["order_id"]),
            customer_name=str(data["customer_name"]),
            product=str(data["product"]),
            quantity=int(data["quantity"]),
            total_price=Decimal(str(data["total_price"])),
            status=OrderStatus(data["status"]),
        )

    @classmethod
    def from_json(cls, payload: str | bytes) -> Order:
        data = json.loads(payload)
        if not isinstance(data, dict):
            raise ValueError("order JSON must contain an object")
        return cls.from_dict(data)
