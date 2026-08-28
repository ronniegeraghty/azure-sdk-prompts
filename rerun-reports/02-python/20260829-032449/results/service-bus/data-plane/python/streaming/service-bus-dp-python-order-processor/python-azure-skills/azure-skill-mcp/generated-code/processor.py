from __future__ import annotations

import json
import logging
from collections.abc import Callable
from dataclasses import dataclass

from azure.servicebus import (
    NEXT_AVAILABLE_SESSION,
    ServiceBusMessage,
    ServiceBusReceivedMessage,
    ServiceBusSubQueue,
)
from azure.servicebus import ServiceBusClient as SyncServiceBusClient
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import OperationTimeoutError

from order import Order, OrderStatus
from sender import AsyncOrderSender, OrderSender


LOGGER = logging.getLogger(__name__)
RecoveryFunction = Callable[[ServiceBusReceivedMessage, Exception], Order | None]


@dataclass(frozen=True, slots=True)
class ProcessingResult:
    completed: int = 0
    dead_lettered: int = 0

    @property
    def settled(self) -> int:
        return self.completed + self.dead_lettered


@dataclass(frozen=True, slots=True)
class DeadLetterRecord:
    body: bytes
    correlation_id: str | None
    reason: str | None
    error_description: str | None


def _message_body(message: ServiceBusReceivedMessage) -> bytes:
    return b"".join(message.body)


def _deserialize_order(message: ServiceBusReceivedMessage) -> Order:
    return Order.from_json(_message_body(message).decode("utf-8"))


def _process_order(order: Order) -> Order:
    processing_order = order.with_status(OrderStatus.PROCESSING)
    LOGGER.info(
        "Processing order %s for customer %s",
        processing_order.order_id,
        processing_order.customer_name,
    )
    completed_order = processing_order.with_status(OrderStatus.COMPLETED)
    LOGGER.info("Completed order: %s", completed_order.to_json())
    return completed_order


def _dead_letter_description(exc: Exception) -> str:
    return str(exc)[:4_096]


class OrderProcessor:
    def __init__(self, client: SyncServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    def process_orders(
        self,
        expected_messages: int,
        max_wait_time: float = 35,
    ) -> ProcessingResult:
        completed = 0
        dead_lettered = 0

        while completed + dead_lettered < expected_messages:
            try:
                with self._client.get_queue_receiver(
                    queue_name=self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    max_wait_time=max_wait_time,
                ) as receiver:
                    while completed + dead_lettered < expected_messages:
                        messages = receiver.receive_messages(
                            max_message_count=100,
                            max_wait_time=2,
                        )
                        if not messages:
                            break

                        for message in messages:
                            try:
                                order = _deserialize_order(message)
                            except (
                                json.JSONDecodeError,
                                TypeError,
                                UnicodeDecodeError,
                                ValueError,
                            ) as exc:
                                LOGGER.error(
                                    "Dead-lettering invalid order %s: %s",
                                    message.correlation_id,
                                    exc,
                                )
                                receiver.dead_letter_message(
                                    message,
                                    reason="OrderDeserializationFailed",
                                    error_description=_dead_letter_description(exc),
                                )
                                dead_lettered += 1
                                continue

                            try:
                                _process_order(order)
                            except (ArithmeticError, RuntimeError, ValueError) as exc:
                                LOGGER.error(
                                    "Dead-lettering order %s after processing failure: %s",
                                    order.order_id,
                                    exc,
                                )
                                receiver.dead_letter_message(
                                    message,
                                    reason="OrderProcessingFailed",
                                    error_description=_dead_letter_description(exc),
                                )
                                dead_lettered += 1
                            else:
                                receiver.complete_message(message)
                                completed += 1
            except OperationTimeoutError:
                LOGGER.warning("No session became available before the timeout")
                break

        return ProcessingResult(completed, dead_lettered)

    def inspect_dead_letters(
        self,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> list[DeadLetterRecord]:
        with self._client.get_queue_receiver(
            queue_name=self._queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=max_wait_time,
        ) as receiver:
            messages = receiver.receive_messages(
                max_message_count=max_messages,
                max_wait_time=max_wait_time,
            )
            records = [
                DeadLetterRecord(
                    body=_message_body(message),
                    correlation_id=message.correlation_id,
                    reason=message.dead_letter_reason,
                    error_description=message.dead_letter_error_description,
                )
                for message in messages
            ]
            for message, record in zip(messages, records, strict=True):
                LOGGER.info(
                    "Dead letter correlation_id=%s reason=%s description=%s body=%r",
                    record.correlation_id,
                    record.reason,
                    record.error_description,
                    record.body,
                )
                receiver.abandon_message(message)
            return records

    def reprocess_dead_letters(
        self,
        recovery: RecoveryFunction | None = None,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> int:
        resent = 0
        order_sender = OrderSender(self._client, self._queue_name)
        with self._client.get_queue_receiver(
            queue_name=self._queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=max_wait_time,
        ) as receiver:
            messages = receiver.receive_messages(
                max_message_count=max_messages,
                max_wait_time=max_wait_time,
            )
            for message in messages:
                try:
                    order = _deserialize_order(message)
                except (
                    json.JSONDecodeError,
                    TypeError,
                    UnicodeDecodeError,
                    ValueError,
                ) as exc:
                    order = recovery(message, exc) if recovery else None
                    if order is None:
                        LOGGER.warning(
                            "Leaving unrecoverable message %s in the dead-letter queue",
                            message.correlation_id,
                        )
                        receiver.abandon_message(message)
                        continue

                order_sender.send_order(
                    order.with_status(OrderStatus.PENDING),
                    is_reprocess=True,
                )
                receiver.complete_message(message)
                resent += 1
                LOGGER.info("Resubmitted dead-lettered order %s", order.order_id)
        return resent


class AsyncOrderProcessor:
    def __init__(self, client: AsyncServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    async def process_orders(
        self,
        expected_messages: int,
        max_wait_time: float = 35,
    ) -> ProcessingResult:
        completed = 0
        dead_lettered = 0

        while completed + dead_lettered < expected_messages:
            try:
                async with self._client.get_queue_receiver(
                    queue_name=self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    max_wait_time=max_wait_time,
                ) as receiver:
                    while completed + dead_lettered < expected_messages:
                        messages = await receiver.receive_messages(
                            max_message_count=100,
                            max_wait_time=2,
                        )
                        if not messages:
                            break

                        for message in messages:
                            try:
                                order = _deserialize_order(message)
                            except (
                                json.JSONDecodeError,
                                TypeError,
                                UnicodeDecodeError,
                                ValueError,
                            ) as exc:
                                LOGGER.error(
                                    "Dead-lettering invalid order %s: %s",
                                    message.correlation_id,
                                    exc,
                                )
                                await receiver.dead_letter_message(
                                    message,
                                    reason="OrderDeserializationFailed",
                                    error_description=_dead_letter_description(exc),
                                )
                                dead_lettered += 1
                                continue

                            try:
                                _process_order(order)
                            except (ArithmeticError, RuntimeError, ValueError) as exc:
                                LOGGER.error(
                                    "Dead-lettering order %s after processing failure: %s",
                                    order.order_id,
                                    exc,
                                )
                                await receiver.dead_letter_message(
                                    message,
                                    reason="OrderProcessingFailed",
                                    error_description=_dead_letter_description(exc),
                                )
                                dead_lettered += 1
                            else:
                                await receiver.complete_message(message)
                                completed += 1
            except OperationTimeoutError:
                LOGGER.warning("No session became available before the timeout")
                break

        return ProcessingResult(completed, dead_lettered)

    async def inspect_dead_letters(
        self,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> list[DeadLetterRecord]:
        async with self._client.get_queue_receiver(
            queue_name=self._queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=max_wait_time,
        ) as receiver:
            messages = await receiver.receive_messages(
                max_message_count=max_messages,
                max_wait_time=max_wait_time,
            )
            records = [
                DeadLetterRecord(
                    body=_message_body(message),
                    correlation_id=message.correlation_id,
                    reason=message.dead_letter_reason,
                    error_description=message.dead_letter_error_description,
                )
                for message in messages
            ]
            for message, record in zip(messages, records, strict=True):
                LOGGER.info(
                    "Dead letter correlation_id=%s reason=%s description=%s body=%r",
                    record.correlation_id,
                    record.reason,
                    record.error_description,
                    record.body,
                )
                await receiver.abandon_message(message)
            return records

    async def reprocess_dead_letters(
        self,
        recovery: RecoveryFunction | None = None,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> int:
        resent = 0
        order_sender = AsyncOrderSender(self._client, self._queue_name)
        async with self._client.get_queue_receiver(
            queue_name=self._queue_name,
            sub_queue=ServiceBusSubQueue.DEAD_LETTER,
            session_id=NEXT_AVAILABLE_SESSION,
            max_wait_time=max_wait_time,
        ) as receiver:
            messages = await receiver.receive_messages(
                max_message_count=max_messages,
                max_wait_time=max_wait_time,
            )
            for message in messages:
                try:
                    order = _deserialize_order(message)
                except (
                    json.JSONDecodeError,
                    TypeError,
                    UnicodeDecodeError,
                    ValueError,
                ) as exc:
                    order = recovery(message, exc) if recovery else None
                    if order is None:
                        LOGGER.warning(
                            "Leaving unrecoverable message %s in the dead-letter queue",
                            message.correlation_id,
                        )
                        await receiver.abandon_message(message)
                        continue

                await order_sender.send_order(
                    order.with_status(OrderStatus.PENDING),
                    is_reprocess=True,
                )
                await receiver.complete_message(message)
                resent += 1
                LOGGER.info("Resubmitted dead-lettered order %s", order.order_id)
        return resent
