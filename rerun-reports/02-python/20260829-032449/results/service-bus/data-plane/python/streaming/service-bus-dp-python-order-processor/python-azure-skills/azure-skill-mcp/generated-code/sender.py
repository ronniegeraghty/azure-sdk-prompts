from __future__ import annotations

from collections import defaultdict
from collections.abc import Iterable
from datetime import datetime, timedelta, timezone
from uuid import uuid4

from azure.servicebus import ServiceBusMessage
from azure.servicebus import ServiceBusClient as SyncServiceBusClient
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import MessageSizeExceededError

from order import Order


HIGH_PRIORITY_DELAY = timedelta(seconds=30)


def create_order_message(
    order: Order,
    high_priority_threshold: float,
    message_id: str | None = None,
) -> ServiceBusMessage:
    is_high_priority = order.total_price > high_priority_threshold
    message = ServiceBusMessage(
        order.to_json(),
        content_type="application/json",
        correlation_id=order.order_id,
        message_id=message_id or order.order_id,
        session_id=order.customer_name,
        application_properties={
            "priority": "high" if is_high_priority else "normal",
        },
    )
    if is_high_priority:
        message.scheduled_enqueue_time_utc = (
            datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY
        )
    return message


class OrderSender:
    def __init__(
        self,
        client: SyncServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    def send_order(self, order: Order, *, is_reprocess: bool = False) -> None:
        message_id = (
            f"{order.order_id}:reprocess:{uuid4()}" if is_reprocess else None
        )
        with self._client.get_queue_sender(self._queue_name) as sender:
            sender.send_messages(
                create_order_message(
                    order,
                    self._high_priority_threshold,
                    message_id,
                )
            )

    def send_orders(self, orders: Iterable[Order]) -> None:
        orders_by_customer: dict[str, list[Order]] = defaultdict(list)
        for order in orders:
            orders_by_customer[order.customer_name].append(order)

        with self._client.get_queue_sender(self._queue_name) as sender:
            for customer_orders in orders_by_customer.values():
                batch = sender.create_message_batch()
                for order in customer_orders:
                    message = create_order_message(
                        order,
                        self._high_priority_threshold,
                    )
                    try:
                        batch.add_message(message)
                    except MessageSizeExceededError:
                        if len(batch) == 0:
                            raise
                        sender.send_messages(batch)
                        batch = sender.create_message_batch()
                        batch.add_message(message)

                if len(batch) > 0:
                    sender.send_messages(batch)


class AsyncOrderSender:
    def __init__(
        self,
        client: AsyncServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    async def send_order(
        self,
        order: Order,
        *,
        is_reprocess: bool = False,
    ) -> None:
        message_id = (
            f"{order.order_id}:reprocess:{uuid4()}" if is_reprocess else None
        )
        async with self._client.get_queue_sender(self._queue_name) as sender:
            await sender.send_messages(
                create_order_message(
                    order,
                    self._high_priority_threshold,
                    message_id,
                )
            )

    async def send_orders(self, orders: Iterable[Order]) -> None:
        orders_by_customer: dict[str, list[Order]] = defaultdict(list)
        for order in orders:
            orders_by_customer[order.customer_name].append(order)

        async with self._client.get_queue_sender(self._queue_name) as sender:
            for customer_orders in orders_by_customer.values():
                batch = await sender.create_message_batch()
                for order in customer_orders:
                    message = create_order_message(
                        order,
                        self._high_priority_threshold,
                    )
                    try:
                        batch.add_message(message)
                    except MessageSizeExceededError:
                        if len(batch) == 0:
                            raise
                        await sender.send_messages(batch)
                        batch = await sender.create_message_batch()
                        batch.add_message(message)

                if len(batch) > 0:
                    await sender.send_messages(batch)
