from __future__ import annotations

import json
from dataclasses import asdict, dataclass, replace
from enum import StrEnum
from typing import Any


class OrderStatus(StrEnum):
    PENDING = "pending"
    PROCESSING = "processing"
    COMPLETED = "completed"
    FAILED = "failed"


@dataclass(frozen=True, slots=True)
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

    def to_json(self) -> str:
        return json.dumps(asdict(self), separators=(",", ":"))

    @classmethod
    def from_json(cls, value: str | bytes) -> Order:
        data: Any = json.loads(value)
        if not isinstance(data, dict):
            raise ValueError("order JSON must contain an object")

        required_fields = {
            "order_id",
            "customer_name",
            "product",
            "quantity",
            "total_price",
        }
        missing_fields = required_fields - data.keys()
        if missing_fields:
            missing = ", ".join(sorted(missing_fields))
            raise ValueError(f"order JSON is missing required fields: {missing}")

        try:
            status = OrderStatus(data.get("status", OrderStatus.PENDING))
            return cls(
                order_id=str(data["order_id"]),
                customer_name=str(data["customer_name"]),
                product=str(data["product"]),
                quantity=int(data["quantity"]),
                total_price=float(data["total_price"]),
                status=status,
            )
        except (TypeError, ValueError) as exc:
            raise ValueError(f"invalid order data: {exc}") from exc

    def with_status(self, status: OrderStatus) -> Order:
        return replace(self, status=status)
