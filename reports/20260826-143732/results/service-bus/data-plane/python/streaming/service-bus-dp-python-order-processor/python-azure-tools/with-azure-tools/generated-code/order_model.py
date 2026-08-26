"""Order domain model and JSON serialization."""

from __future__ import annotations

import json
from dataclasses import dataclass
from enum import Enum
from typing import Any


class OrderStatus(str, Enum):
    PENDING = "pending"
    PROCESSING = "processing"
    COMPLETED = "completed"
    FAILED = "failed"


class OrderDeserializationError(ValueError):
    """Raised when a message body cannot be converted to an Order."""


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
        if isinstance(self.quantity, bool) or self.quantity <= 0:
            raise ValueError("quantity must be a positive integer")
        if isinstance(self.total_price, bool) or self.total_price < 0:
            raise ValueError("total_price must be non-negative")
        if not isinstance(self.status, OrderStatus):
            self.status = OrderStatus(self.status)

    def to_dict(self) -> dict[str, Any]:
        return {
            "order_id": self.order_id,
            "customer_name": self.customer_name,
            "product": self.product,
            "quantity": self.quantity,
            "total_price": self.total_price,
            "status": self.status.value,
        }

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), separators=(",", ":"))

    @classmethod
    def from_dict(cls, data: dict[str, Any]) -> "Order":
        try:
            return cls(
                order_id=str(data["order_id"]),
                customer_name=str(data["customer_name"]),
                product=str(data["product"]),
                quantity=int(data["quantity"]),
                total_price=float(data["total_price"]),
                status=OrderStatus(data["status"]),
            )
        except (KeyError, TypeError, ValueError) as exc:
            raise OrderDeserializationError(f"Invalid order data: {exc}") from exc

    @classmethod
    def from_json(cls, payload: str) -> "Order":
        try:
            data = json.loads(payload)
        except json.JSONDecodeError as exc:
            raise OrderDeserializationError(f"Invalid order JSON: {exc}") from exc

        if not isinstance(data, dict):
            raise OrderDeserializationError("Order JSON must contain an object")
        return cls.from_dict(data)
