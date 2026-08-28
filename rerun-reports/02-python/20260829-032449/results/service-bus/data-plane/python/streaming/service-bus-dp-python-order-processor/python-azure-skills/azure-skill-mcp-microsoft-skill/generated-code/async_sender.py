from __future__ import annotations

import logging
from collections import defaultdict
from collections.abc import Iterable

from azure.servicebus import ServiceBusMessageBatch
from azure.servicebus.aio import ServiceBusClient

from messages import create_order_message
from order import Order

logger = logging.getLogger(__name__)


class AsyncOrderSender:
    def __init__(
        self,
        client: ServiceBusClient,
        queue_name: str,
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    async def send_order(self, order: Order) -> None:
        message = create_order_message(order, self._high_priority_threshold)
        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            await sender.send_messages(message)
        logger.info("Sent order %s for %s", order.order_id, order.customer_name)

    async def send_orders(self, orders: Iterable[Order]) -> None:
        orders_by_customer: dict[str, list[Order]] = defaultdict(list)
        for order in orders:
            orders_by_customer[order.customer_name].append(order)

        async with self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            for customer_orders in orders_by_customer.values():
                batch = await sender.create_message_batch()
                for order in customer_orders:
                    message = create_order_message(
                        order,
                        self._high_priority_threshold,
                    )
                    batch = await self._add_or_flush(sender, batch, message)
                if len(batch):
                    await sender.send_messages(batch)

        logger.info("Sent %d customer order group(s)", len(orders_by_customer))

    @staticmethod
    async def _add_or_flush(sender, batch: ServiceBusMessageBatch, message):
        try:
            batch.add_message(message)
            return batch
        except ValueError:
            if not len(batch):
                raise ValueError("An order message exceeds the queue batch size limit")

        await sender.send_messages(batch)
        next_batch = await sender.create_message_batch()
        try:
            next_batch.add_message(message)
        except ValueError as exc:
            raise ValueError(
                "An order message exceeds the queue batch size limit"
            ) from exc
        return next_batch
