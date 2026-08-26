"""Synchronous and asynchronous Azure Service Bus order senders."""

from __future__ import annotations

from collections import defaultdict
from datetime import datetime, timedelta, timezone
from typing import Iterable

from azure.servicebus import ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus import ServiceBusClient

from order_model import Order

DEFAULT_HIGH_PRIORITY_THRESHOLD = 1_000.0
FRAUD_REVIEW_DELAY = timedelta(seconds=30)


def _order_message(order: Order, high_priority_threshold: float) -> ServiceBusMessage:
    high_priority = order.total_price > high_priority_threshold
    scheduled_time = (
        datetime.now(timezone.utc) + FRAUD_REVIEW_DELAY if high_priority else None
    )
    return ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        message_id=order.order_id,
        session_id=order.customer_name,
        subject="high-priority-order" if high_priority else "order",
        application_properties={"priority": "high" if high_priority else "normal"},
        scheduled_enqueue_time_utc=scheduled_time,
    )


def _orders_by_customer(orders: Iterable[Order]) -> dict[str, list[Order]]:
    grouped: dict[str, list[Order]] = defaultdict(list)
    for order in orders:
        grouped[order.customer_name].append(order)
    return grouped


class OrderSender:
    def __init__(
        self,
        client: ServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = DEFAULT_HIGH_PRIORITY_THRESHOLD,
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    def send_order(self, order: Order) -> None:
        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            sender.send_messages(
                _order_message(order, self._high_priority_threshold)
            )

    def send_orders(self, orders: Iterable[Order]) -> None:
        grouped_orders = _orders_by_customer(orders)
        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            for customer_orders in grouped_orders.values():
                batch = sender.create_message_batch()
                batch_count = 0
                for order in customer_orders:
                    message = _order_message(order, self._high_priority_threshold)
                    try:
                        batch.add_message(message)
                        batch_count += 1
                    except ValueError:
                        if batch_count == 0:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum message size"
                            ) from None
                        sender.send_messages(batch)
                        batch = sender.create_message_batch()
                        try:
                            batch.add_message(message)
                        except ValueError:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum message size"
                            ) from None
                        batch_count = 1
                if batch_count:
                    sender.send_messages(batch)

    def send_invalid_message_for_demo(self, customer_name: str) -> None:
        message = ServiceBusMessage(
            '{"order_id": "invalid"',
            content_type="application/json",
            correlation_id="invalid-order",
            message_id=f"invalid-{customer_name}",
            session_id=customer_name,
        )
        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            sender.send_messages(message)


class AsyncOrderSender:
    def __init__(
        self,
        client: AsyncServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = DEFAULT_HIGH_PRIORITY_THRESHOLD,
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    async def send_order(self, order: Order) -> None:
        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            await sender.send_messages(
                _order_message(order, self._high_priority_threshold)
            )

    async def send_orders(self, orders: Iterable[Order]) -> None:
        grouped_orders = _orders_by_customer(orders)
        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            for customer_orders in grouped_orders.values():
                batch = await sender.create_message_batch()
                batch_count = 0
                for order in customer_orders:
                    message = _order_message(order, self._high_priority_threshold)
                    try:
                        batch.add_message(message)
                        batch_count += 1
                    except ValueError:
                        if batch_count == 0:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum message size"
                            ) from None
                        await sender.send_messages(batch)
                        batch = await sender.create_message_batch()
                        try:
                            batch.add_message(message)
                        except ValueError:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum message size"
                            ) from None
                        batch_count = 1
                if batch_count:
                    await sender.send_messages(batch)

    async def send_invalid_message_for_demo(self, customer_name: str) -> None:
        message = ServiceBusMessage(
            '{"order_id": "invalid"',
            content_type="application/json",
            correlation_id="invalid-order",
            message_id=f"invalid-{customer_name}",
            session_id=customer_name,
        )
        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            await sender.send_messages(message)
