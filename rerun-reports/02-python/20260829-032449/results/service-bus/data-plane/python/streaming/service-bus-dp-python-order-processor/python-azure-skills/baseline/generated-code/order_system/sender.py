from __future__ import annotations

from collections.abc import Iterable
from typing import Any

from azure.servicebus import ServiceBusClient
from azure.servicebus.exceptions import MessageSizeExceededError

from .messages import create_order_message
from .model import Order


class OrderSender:
    def __init__(
        self,
        fully_qualified_namespace: str,
        queue_name: str,
        credential: Any,
        high_priority_threshold: float = 1_000.0,
    ) -> None:
        self._client = ServiceBusClient(
            fully_qualified_namespace=fully_qualified_namespace,
            credential=credential,
        )
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold

    def __enter__(self) -> OrderSender:
        return self

    def __exit__(self, *args: object) -> None:
        self.close()

    def close(self) -> None:
        self._client.close()

    def send_order(self, order: Order) -> None:
        with self._client.get_queue_sender(self._queue_name) as sender:
            sender.send_messages(
                create_order_message(order, self._high_priority_threshold)
            )

    def send_orders(self, orders: Iterable[Order]) -> None:
        with self._client.get_queue_sender(self._queue_name) as sender:
            batch = sender.create_message_batch()
            batch_count = 0

            for order in orders:
                message = create_order_message(
                    order, self._high_priority_threshold
                )
                try:
                    batch.add_message(message)
                    batch_count += 1
                except MessageSizeExceededError:
                    if batch_count == 0:
                        raise
                    sender.send_messages(batch)
                    batch = sender.create_message_batch()
                    batch.add_message(message)
                    batch_count = 1

            if batch_count:
                sender.send_messages(batch)

