from __future__ import annotations

import re
from datetime import datetime, timedelta, timezone
from decimal import Decimal

from azure.servicebus import ServiceBusMessage

from .models import Order

HIGH_PRIORITY_DELAY = timedelta(seconds=30)


def customer_session_id(customer_name: str) -> str:
    normalized = re.sub(r"[^A-Za-z0-9._-]+", "-", customer_name.strip().lower())
    return normalized.strip("-") or "customer"


def create_order_message(
    order: Order,
    high_priority_threshold: Decimal,
) -> ServiceBusMessage:
    is_high_priority = order.total_price > high_priority_threshold
    message = ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        message_id=order.order_id,
        session_id=customer_session_id(order.customer_name),
        application_properties={
            "order_id": order.order_id,
            "priority": "high" if is_high_priority else "normal",
        },
    )
    if is_high_priority:
        message.scheduled_enqueue_time_utc = (
            datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY
        )
    return message
