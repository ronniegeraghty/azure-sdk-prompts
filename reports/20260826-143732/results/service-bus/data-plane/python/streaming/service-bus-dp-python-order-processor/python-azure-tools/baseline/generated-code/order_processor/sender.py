from __future__ import annotations

from collections import defaultdict
from collections.abc import Iterable, Mapping
from datetime import datetime, timedelta, timezone
from typing import Any

from azure.servicebus import ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import MessageSizeExceededError
from azure.servicebus import ServiceBusClient

from .model import Order


class _MessageFactory:
    def __init__(self, high_priority_threshold: float) -> None:
        self.high_priority_threshold = high_priority_threshold

    def create(
        self,
        order: Order,
        application_properties: Mapping[str, Any] | None = None,
    ) -> ServiceBusMessage:
        properties: dict[str, Any] = dict(application_properties or {})
        is_high_priority = order.total_price > self.high_priority_threshold
        properties["priority"] = "high" if is_high_priority else "normal"

        scheduled_time = (
            datetime.now(timezone.utc) + timedelta(seconds=30)
            if is_high_priority
            else None
        )
        return ServiceBusMessage(
            order.to_json(),
            content_type="application/json",
            correlation_id=order.order_id,
            session_id=order.customer_name,
            application_properties=properties,
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
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self.client = client
        self.queue_name = queue_name
        self._messages = _MessageFactory(high_priority_threshold)

    def send_order(
        self,
        order: Order,
        application_properties: Mapping[str, Any] | None = None,
    ) -> None:
        with self.client.get_queue_sender(queue_name=self.queue_name) as sender:
            sender.send_messages(self._messages.create(order, application_properties))

    def send_orders(self, orders: Iterable[Order]) -> None:
        with self.client.get_queue_sender(queue_name=self.queue_name) as sender:
            for customer_orders in _orders_by_customer(orders).values():
                batch = sender.create_message_batch()
                for order in customer_orders:
                    message = self._messages.create(order)
                    try:
                        batch.add_message(message)
                    except MessageSizeExceededError:
                        if len(batch) == 0:
                            raise
                        sender.send_messages(batch)
                        batch = sender.create_message_batch()
                        batch.add_message(message)
                if len(batch):
                    sender.send_messages(batch)


class AsyncOrderSender:
    def __init__(
        self,
        client: AsyncServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self.client = client
        self.queue_name = queue_name
        self._messages = _MessageFactory(high_priority_threshold)

    async def send_order(
        self,
        order: Order,
        application_properties: Mapping[str, Any] | None = None,
    ) -> None:
        async with self.client.get_queue_sender(queue_name=self.queue_name) as sender:
            await sender.send_messages(
                self._messages.create(order, application_properties)
            )

    async def send_orders(self, orders: Iterable[Order]) -> None:
        async with self.client.get_queue_sender(queue_name=self.queue_name) as sender:
            for customer_orders in _orders_by_customer(orders).values():
                batch = await sender.create_message_batch()
                for order in customer_orders:
                    message = self._messages.create(order)
                    try:
                        batch.add_message(message)
                    except MessageSizeExceededError:
                        if len(batch) == 0:
                            raise
                        await sender.send_messages(batch)
                        batch = await sender.create_message_batch()
                        batch.add_message(message)
                if len(batch):
                    await sender.send_messages(batch)
