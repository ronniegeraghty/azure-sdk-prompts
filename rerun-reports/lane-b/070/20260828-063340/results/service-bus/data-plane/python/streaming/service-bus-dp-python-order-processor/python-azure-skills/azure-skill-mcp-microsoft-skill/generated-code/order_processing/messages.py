"""Shared Service Bus message helpers."""

from __future__ import annotations

import hashlib
from datetime import datetime
from typing import Any

from azure.servicebus import ServiceBusMessage

from .model import Order


def customer_session_id(customer_name: str) -> str:
    normalized = " ".join(customer_name.casefold().split())
    digest = hashlib.sha256(normalized.encode("utf-8")).hexdigest()
    return f"customer-{digest}"


def order_message(
    order: Order,
    *,
    high_priority: bool,
    scheduled_at: datetime | None = None,
    retried_from_dlq: bool = False,
) -> ServiceBusMessage:
    properties: dict[str, Any] = {
        "priority": "high" if high_priority else "normal",
    }
    if retried_from_dlq:
        properties["retried_from_dlq"] = True

    return ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        message_id=order.order_id,
        correlation_id=order.order_id,
        session_id=customer_session_id(order.customer_name),
        application_properties=properties,
        scheduled_enqueue_time_utc=scheduled_at,
    )


def message_body_text(message: Any) -> str:
    body = message.body
    if isinstance(body, str):
        return body
    if isinstance(body, bytes):
        return body.decode("utf-8")

    chunks: list[bytes] = []
    for chunk in body:
        if isinstance(chunk, bytes):
            chunks.append(chunk)
        elif isinstance(chunk, bytearray):
            chunks.append(bytes(chunk))
        else:
            raise ValueError(f"unsupported message body chunk: {type(chunk).__name__}")
    return b"".join(chunks).decode("utf-8")
