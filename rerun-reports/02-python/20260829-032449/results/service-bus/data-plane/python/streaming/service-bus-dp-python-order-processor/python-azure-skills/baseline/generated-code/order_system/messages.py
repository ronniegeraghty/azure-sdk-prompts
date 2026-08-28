from __future__ import annotations

from datetime import datetime, timedelta, timezone

from azure.servicebus import ServiceBusMessage

from .model import Order


HIGH_PRIORITY_DELAY = timedelta(seconds=30)


def create_order_message(
    order: Order, high_priority_threshold: float
) -> ServiceBusMessage:
    is_high_priority = order.total_price > high_priority_threshold
    return ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        session_id=order.customer_name,
        application_properties={
            "priority": "high" if is_high_priority else "normal",
        },
        scheduled_enqueue_time_utc=(
            datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY
            if is_high_priority
            else None
        ),
    )

