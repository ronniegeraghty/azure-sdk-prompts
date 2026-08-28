from __future__ import annotations

from collections import defaultdict
from collections.abc import Iterable
from decimal import Decimal

from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from .message_factory import create_order_message, customer_session_id
from .models import Order


class OrderSender:
    def __init__(
        self,
        client: ServiceBusClient,
        queue_name: str,
        high_priority_threshold: Decimal = Decimal("1000"),
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    def send_order(self, order: Order) -> None:
        with self._client.get_queue_sender(self._queue_name) as sender:
            sender.send_messages(
                create_order_message(order, self._high_priority_threshold)
            )

    def send_orders(self, orders: Iterable[Order]) -> None:
        grouped = _group_orders_by_session(orders)
        with self._client.get_queue_sender(self._queue_name) as sender:
            for session_orders in grouped.values():
                batch = sender.create_message_batch()
                for order in session_orders:
                    message = create_order_message(
                        order, self._high_priority_threshold
                    )
                    try:
                        batch.add_message(message)
                    except ValueError:
                        if len(batch) == 0:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum "
                                "Service Bus message batch size"
                            ) from None
                        sender.send_messages(batch)
                        batch = sender.create_message_batch()
                        try:
                            batch.add_message(message)
                        except ValueError:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum "
                                "Service Bus message batch size"
                            ) from None
                if len(batch):
                    sender.send_messages(batch)

    def send_raw_message(self, message: ServiceBusMessage) -> None:
        with self._client.get_queue_sender(self._queue_name) as sender:
            sender.send_messages(message)


class AsyncOrderSender:
    def __init__(
        self,
        client: AsyncServiceBusClient,
        queue_name: str,
        high_priority_threshold: Decimal = Decimal("1000"),
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    async def send_order(self, order: Order) -> None:
        async with self._client.get_queue_sender(self._queue_name) as sender:
            await sender.send_messages(
                create_order_message(order, self._high_priority_threshold)
            )

    async def send_orders(self, orders: Iterable[Order]) -> None:
        grouped = _group_orders_by_session(orders)
        async with self._client.get_queue_sender(self._queue_name) as sender:
            for session_orders in grouped.values():
                batch = await sender.create_message_batch()
                for order in session_orders:
                    message = create_order_message(
                        order, self._high_priority_threshold
                    )
                    try:
                        batch.add_message(message)
                    except ValueError:
                        if len(batch) == 0:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum "
                                "Service Bus message batch size"
                            ) from None
                        await sender.send_messages(batch)
                        batch = await sender.create_message_batch()
                        try:
                            batch.add_message(message)
                        except ValueError:
                            raise ValueError(
                                f"Order {order.order_id!r} exceeds the maximum "
                                "Service Bus message batch size"
                            ) from None
                if len(batch):
                    await sender.send_messages(batch)

    async def send_raw_message(self, message: ServiceBusMessage) -> None:
        async with self._client.get_queue_sender(self._queue_name) as sender:
            await sender.send_messages(message)


def _group_orders_by_session(orders: Iterable[Order]) -> dict[str, list[Order]]:
    grouped: dict[str, list[Order]] = defaultdict(list)
    for order in orders:
        grouped[customer_session_id(order.customer_name)].append(order)
    return grouped
