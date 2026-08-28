from __future__ import annotations

import json
import logging
from collections.abc import Callable, Iterable
from dataclasses import dataclass

from azure.servicebus import (
    ServiceBusClient,
    ServiceBusReceivedMessage,
    ServiceBusSubQueue,
)

from messages import create_reprocessed_message
from order import Order, OrderStatus

logger = logging.getLogger(__name__)


@dataclass(frozen=True, slots=True)
class DeadLetterRecord:
    body: str
    reason: str
    description: str
    correlation_id: str | None


class SyncOrderProcessor:
    def __init__(self, client: ServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    def process_customers(
        self,
        customer_names: Iterable[str],
        max_wait_time: float = 5,
    ) -> int:
        processed = 0
        for customer_name in dict.fromkeys(customer_names):
            processed += self.process_customer(customer_name, max_wait_time)
        return processed

    def process_customer(
        self,
        customer_name: str,
        max_wait_time: float = 5,
    ) -> int:
        processed = 0
        receiver = self._client.get_queue_receiver(
            queue_name=self._queue_name,
            session_id=customer_name,
            max_wait_time=max_wait_time,
        )
        with receiver:
            while True:
                messages = receiver.receive_messages(
                    max_message_count=20,
                    max_wait_time=max_wait_time,
                )
                if not messages:
                    break
                for message in messages:
                    if self._process_message(receiver, message):
                        processed += 1
        return processed

    def inspect_dead_letters(
        self,
        customer_name: str,
        replacement_factory: Callable[[DeadLetterRecord], Order | None] | None = None,
        max_wait_time: float = 5,
    ) -> list[DeadLetterRecord]:
        records: list[DeadLetterRecord] = []
        receiver = self._client.get_queue_receiver(
            queue_name=self._queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=customer_name,
            max_wait_time=max_wait_time,
        )
        with receiver, self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            messages = receiver.receive_messages(
                max_message_count=100,
                max_wait_time=max_wait_time,
            )
            for message in messages:
                record = DeadLetterRecord(
                    body=str(message),
                    reason=message.dead_letter_reason or "Unknown",
                    description=message.dead_letter_error_description or "",
                    correlation_id=message.correlation_id,
                )
                records.append(record)
                logger.warning(
                    "DLQ order correlation_id=%s reason=%s description=%s body=%s",
                    record.correlation_id,
                    record.reason,
                    record.description,
                    record.body,
                )

                replacement = replacement_factory(record) if replacement_factory else None
                if replacement is None:
                    receiver.abandon_message(message)
                    continue

                sender.send_messages(
                    create_reprocessed_message(replacement, record.reason)
                )
                receiver.complete_message(message)
                logger.info("Reprocessed dead-lettered order as %s", replacement.order_id)
        return records

    @staticmethod
    def _process_message(receiver, message: ServiceBusReceivedMessage) -> bool:
        try:
            order = Order.from_json(str(message))
            order.status = OrderStatus.PROCESSING
            logger.info(
                "Processing order %s for %s: %d x %s ($%.2f)",
                order.order_id,
                order.customer_name,
                order.quantity,
                order.product,
                order.total_price,
            )
            order.status = OrderStatus.COMPLETED
        except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:
            receiver.dead_letter_message(
                message,
                reason="InvalidOrderPayload",
                error_description=f"{type(exc).__name__}: {exc}",
            )
            logger.error("Dead-lettered invalid order: %s", exc)
            return False
        except Exception as exc:
            receiver.dead_letter_message(
                message,
                reason="OrderProcessingFailed",
                error_description=f"{type(exc).__name__}: {exc}",
            )
            logger.exception("Dead-lettered order after processing failure")
            return False

        receiver.complete_message(message)
        logger.info("Completed order %s", order.order_id)
        return True
