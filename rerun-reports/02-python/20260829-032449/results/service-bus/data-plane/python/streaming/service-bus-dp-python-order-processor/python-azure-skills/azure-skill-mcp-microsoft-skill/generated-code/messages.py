from __future__ import annotations

from datetime import datetime, timedelta, timezone
from uuid import uuid4

from azure.servicebus import ServiceBusMessage

from order import Order


def create_order_message(
    order: Order,
    high_priority_threshold: float,
) -> ServiceBusMessage:
    is_high_priority = order.total_price > high_priority_threshold
    scheduled_time = (
        datetime.now(timezone.utc) + timedelta(seconds=30)
        if is_high_priority
        else None
    )

    return ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        message_id=str(uuid4()),
        session_id=order.customer_name,
        scheduled_enqueue_time_utc=scheduled_time,
        application_properties={
            "order_id": order.order_id,
            "priority": "high" if is_high_priority else "normal",
        },
    )


def create_reprocessed_message(order: Order, dead_letter_reason: str) -> ServiceBusMessage:
    return ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        message_id=str(uuid4()),
        session_id=order.customer_name,
        application_properties={
            "order_id": order.order_id,
            "priority": "normal",
            "reprocessed_from_dead_letter": True,
            "original_dead_letter_reason": dead_letter_reason,
        },
    )
