from __future__ import annotations

import logging
from dataclasses import dataclass
from typing import Any

from azure.servicebus import (
    NEXT_AVAILABLE_SESSION,
    ServiceBusClient,
    ServiceBusReceivedMessage,
    ServiceBusSubQueue,
)
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import ServiceBusError

from .model import Order, OrderStatus
from .sender import AsyncOrderSender, OrderSender

LOGGER = logging.getLogger(__name__)


@dataclass(frozen=True)
class DeadLetterRecord:
    order_id: str | None
    reason: str | None
    description: str | None
    body: str


def _message_body(message: ServiceBusReceivedMessage) -> str:
    return b"".join(message.body).decode("utf-8")


def _property_is_true(properties: dict[str | bytes, Any] | None, name: str) -> bool:
    if not properties:
        return False
    value = properties.get(name, properties.get(name.encode()))
    if isinstance(value, bytes):
        value = value.decode("utf-8")
    return value is True or value == "true"


def _process_order(message: ServiceBusReceivedMessage) -> Order:
    order = Order.from_json(_message_body(message))
    order.status = OrderStatus.PROCESSING
    if _property_is_true(message.application_properties, "simulate_failure"):
        raise RuntimeError("simulated processing failure")
    order.status = OrderStatus.COMPLETED
    return order


def _dead_letter_record(message: ServiceBusReceivedMessage) -> DeadLetterRecord:
    correlation_id = message.correlation_id
    return DeadLetterRecord(
        order_id=str(correlation_id) if correlation_id is not None else None,
        reason=message.dead_letter_reason,
        description=message.dead_letter_error_description,
        body=_message_body(message),
    )


class OrderProcessor:
    def __init__(
        self,
        client: ServiceBusClient,
        queue_name: str,
        max_wait_time: float = 5,
    ) -> None:
        self.client = client
        self.queue_name = queue_name
        self.max_wait_time = max_wait_time

    def process(self, max_messages: int | None = None) -> int:
        processed = 0
        while max_messages is None or processed < max_messages:
            receiver = self.client.get_queue_receiver(
                queue_name=self.queue_name,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=self.max_wait_time,
            )
            try:
                with receiver:
                    while max_messages is None or processed < max_messages:
                        messages = receiver.receive_messages(
                            max_message_count=(
                                min(20, max_messages - processed)
                                if max_messages is not None
                                else 20
                            ),
                            max_wait_time=self.max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            try:
                                order = _process_order(message)
                            except (TypeError, ValueError, UnicodeError, RuntimeError) as exc:
                                LOGGER.error(
                                    "Order %s failed: %s", message.correlation_id, exc
                                )
                                receiver.dead_letter_message(
                                    message,
                                    reason="OrderProcessingError",
                                    error_description=str(exc)[:4096],
                                )
                            else:
                                receiver.complete_message(message)
                                LOGGER.info(
                                    "Processed order %s for %s: %s",
                                    order.order_id,
                                    order.customer_name,
                                    order.status.value,
                                )
                            processed += 1
            except ServiceBusError as exc:
                LOGGER.debug("No available order session: %s", exc)
                break
        return processed

    def inspect_dead_letters(self, max_messages: int = 20) -> list[DeadLetterRecord]:
        receiver = self.client.get_queue_receiver(
            queue_name=self.queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=self.max_wait_time,
        )
        try:
            with receiver:
                return [
                    _dead_letter_record(message)
                    for message in receiver.peek_messages(
                        max_message_count=max_messages,
                        from_sequence_number=1,
                    )
                ]
        except ServiceBusError as exc:
            LOGGER.debug("No available dead-letter session: %s", exc)
            return []

    def reprocess_dead_letters(
        self, sender: OrderSender, max_messages: int = 20
    ) -> int:
        reprocessed = 0
        examined = 0
        while examined < max_messages:
            receiver = self.client.get_queue_receiver(
                queue_name=self.queue_name,
                sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=self.max_wait_time,
            )
            try:
                with receiver:
                    messages = receiver.receive_messages(
                        max_message_count=min(20, max_messages - examined),
                        max_wait_time=self.max_wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        examined += 1
                        try:
                            order = Order.from_json(_message_body(message))
                            sender.send_order(order)
                        except (TypeError, ValueError, UnicodeError) as exc:
                            LOGGER.error(
                                "Dead-lettered message %s cannot be reprocessed: %s",
                                message.correlation_id,
                                exc,
                            )
                            receiver.abandon_message(message)
                        else:
                            receiver.complete_message(message)
                            reprocessed += 1
                            LOGGER.info("Requeued order %s", order.order_id)
            except ServiceBusError as exc:
                LOGGER.debug("No available dead-letter session: %s", exc)
                break
        return reprocessed


class AsyncOrderProcessor:
    def __init__(
        self,
        client: AsyncServiceBusClient,
        queue_name: str,
        max_wait_time: float = 5,
    ) -> None:
        self.client = client
        self.queue_name = queue_name
        self.max_wait_time = max_wait_time

    async def process(self, max_messages: int | None = None) -> int:
        processed = 0
        while max_messages is None or processed < max_messages:
            receiver = self.client.get_queue_receiver(
                queue_name=self.queue_name,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=self.max_wait_time,
            )
            try:
                async with receiver:
                    while max_messages is None or processed < max_messages:
                        messages = await receiver.receive_messages(
                            max_message_count=(
                                min(20, max_messages - processed)
                                if max_messages is not None
                                else 20
                            ),
                            max_wait_time=self.max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            try:
                                order = _process_order(message)
                            except (TypeError, ValueError, UnicodeError, RuntimeError) as exc:
                                LOGGER.error(
                                    "Order %s failed: %s", message.correlation_id, exc
                                )
                                await receiver.dead_letter_message(
                                    message,
                                    reason="OrderProcessingError",
                                    error_description=str(exc)[:4096],
                                )
                            else:
                                await receiver.complete_message(message)
                                LOGGER.info(
                                    "Processed order %s for %s: %s",
                                    order.order_id,
                                    order.customer_name,
                                    order.status.value,
                                )
                            processed += 1
            except ServiceBusError as exc:
                LOGGER.debug("No available order session: %s", exc)
                break
        return processed

    async def inspect_dead_letters(
        self, max_messages: int = 20
    ) -> list[DeadLetterRecord]:
        receiver = self.client.get_queue_receiver(
            queue_name=self.queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=self.max_wait_time,
        )
        try:
            async with receiver:
                messages = await receiver.peek_messages(
                    max_message_count=max_messages,
                    from_sequence_number=1,
                )
                return [_dead_letter_record(message) for message in messages]
        except ServiceBusError as exc:
            LOGGER.debug("No available dead-letter session: %s", exc)
            return []

    async def reprocess_dead_letters(
        self, sender: AsyncOrderSender, max_messages: int = 20
    ) -> int:
        reprocessed = 0
        examined = 0
        while examined < max_messages:
            receiver = self.client.get_queue_receiver(
                queue_name=self.queue_name,
                sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=self.max_wait_time,
            )
            try:
                async with receiver:
                    messages = await receiver.receive_messages(
                        max_message_count=min(20, max_messages - examined),
                        max_wait_time=self.max_wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        examined += 1
                        try:
                            order = Order.from_json(_message_body(message))
                            await sender.send_order(order)
                        except (TypeError, ValueError, UnicodeError) as exc:
                            LOGGER.error(
                                "Dead-lettered message %s cannot be reprocessed: %s",
                                message.correlation_id,
                                exc,
                            )
                            await receiver.abandon_message(message)
                        else:
                            await receiver.complete_message(message)
                            reprocessed += 1
                            LOGGER.info("Requeued order %s", order.order_id)
            except ServiceBusError as exc:
                LOGGER.debug("No available dead-letter session: %s", exc)
                break
        return reprocessed
