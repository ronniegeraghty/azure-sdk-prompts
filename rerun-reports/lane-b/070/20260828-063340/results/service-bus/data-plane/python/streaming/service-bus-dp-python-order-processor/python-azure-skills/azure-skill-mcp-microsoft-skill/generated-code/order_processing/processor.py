"""Synchronous and asynchronous session-aware order processors."""

from __future__ import annotations

import logging
from collections.abc import Callable
from typing import Any

from azure.servicebus import NEXT_AVAILABLE_SESSION, ServiceBusSubQueue
from azure.servicebus.exceptions import OperationTimeoutError

from .messages import message_body_text, order_message
from .model import Order, OrderStatus

LOGGER = logging.getLogger(__name__)
DeadLetterTransformer = Callable[[str, Any], Order]


def _processed(order: Order) -> Order:
    order.status = OrderStatus.PROCESSING
    LOGGER.info(
        "Processing order %s for %s: %d x %s",
        order.order_id,
        order.customer_name,
        order.quantity,
        order.product,
    )
    order.status = OrderStatus.COMPLETED
    LOGGER.info("Completed order %s", order.order_id)
    return order


class OrderProcessor:
    def __init__(self, client, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    def process(self, max_messages: int = 100, max_wait_time: float = 5) -> list[Order]:
        completed: list[Order] = []
        handled = 0
        while handled < max_messages:
            try:
                with self._client.get_queue_receiver(
                    queue_name=self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    max_wait_time=max_wait_time,
                ) as receiver:
                    while handled < max_messages:
                        messages = receiver.receive_messages(
                            max_message_count=max_messages - handled,
                            max_wait_time=max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            handled += 1
                            try:
                                completed.append(
                                    _processed(Order.from_json(message_body_text(message)))
                                )
                                receiver.complete_message(message)
                            except (TypeError, ValueError) as exc:
                                LOGGER.error(
                                    "Dead-lettering message %s: %s",
                                    message.message_id,
                                    exc,
                                )
                                receiver.dead_letter_message(
                                    message,
                                    reason="OrderDeserializationFailed",
                                    error_description=str(exc)[:4096],
                                )
            except OperationTimeoutError:
                break
        return completed

    def inspect_dead_letters(
        self, max_messages: int = 100, max_wait_time: float = 5
    ) -> list[dict[str, Any]]:
        return self._handle_dead_letters(
            max_messages=max_messages,
            max_wait_time=max_wait_time,
            transformer=None,
        )

    def reprocess_dead_letters(
        self,
        transformer: DeadLetterTransformer,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> list[dict[str, Any]]:
        return self._handle_dead_letters(
            max_messages=max_messages,
            max_wait_time=max_wait_time,
            transformer=transformer,
        )

    def _handle_dead_letters(
        self,
        *,
        max_messages: int,
        max_wait_time: float,
        transformer: DeadLetterTransformer | None,
    ) -> list[dict[str, Any]]:
        inspected: list[dict[str, Any]] = []
        handled = 0
        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            while handled < max_messages:
                try:
                    with self._client.get_queue_receiver(
                        queue_name=self._queue_name,
                        sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                        session_id=NEXT_AVAILABLE_SESSION,
                        max_wait_time=max_wait_time,
                    ) as receiver:
                        messages = receiver.receive_messages(
                            max_message_count=max_messages - handled,
                            max_wait_time=max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            handled += 1
                            body = message_body_text(message)
                            details = {
                                "message_id": message.message_id,
                                "reason": message.dead_letter_reason,
                                "description": message.dead_letter_error_description,
                                "body": body,
                            }
                            inspected.append(details)
                            LOGGER.warning("DLQ message: %s", details)
                            if transformer is None:
                                receiver.abandon_message(message)
                                return inspected

                            try:
                                repaired = transformer(body, message)
                                sender.send_messages(
                                    order_message(
                                        repaired,
                                        high_priority=False,
                                        retried_from_dlq=True,
                                    )
                                )
                                receiver.complete_message(message)
                            except (TypeError, ValueError) as exc:
                                LOGGER.error(
                                    "DLQ message %s could not be repaired: %s",
                                    message.message_id,
                                    exc,
                                )
                                receiver.abandon_message(message)
                except OperationTimeoutError:
                    break
        return inspected


class AsyncOrderProcessor:
    def __init__(self, client, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    async def process(
        self, max_messages: int = 100, max_wait_time: float = 5
    ) -> list[Order]:
        completed: list[Order] = []
        handled = 0
        while handled < max_messages:
            try:
                async with self._client.get_queue_receiver(
                    queue_name=self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    max_wait_time=max_wait_time,
                ) as receiver:
                    while handled < max_messages:
                        messages = await receiver.receive_messages(
                            max_message_count=max_messages - handled,
                            max_wait_time=max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            handled += 1
                            try:
                                completed.append(
                                    _processed(Order.from_json(message_body_text(message)))
                                )
                                await receiver.complete_message(message)
                            except (TypeError, ValueError) as exc:
                                LOGGER.error(
                                    "Dead-lettering message %s: %s",
                                    message.message_id,
                                    exc,
                                )
                                await receiver.dead_letter_message(
                                    message,
                                    reason="OrderDeserializationFailed",
                                    error_description=str(exc)[:4096],
                                )
            except OperationTimeoutError:
                break
        return completed

    async def inspect_dead_letters(
        self, max_messages: int = 100, max_wait_time: float = 5
    ) -> list[dict[str, Any]]:
        return await self._handle_dead_letters(
            max_messages=max_messages,
            max_wait_time=max_wait_time,
            transformer=None,
        )

    async def reprocess_dead_letters(
        self,
        transformer: DeadLetterTransformer,
        max_messages: int = 100,
        max_wait_time: float = 5,
    ) -> list[dict[str, Any]]:
        return await self._handle_dead_letters(
            max_messages=max_messages,
            max_wait_time=max_wait_time,
            transformer=transformer,
        )

    async def _handle_dead_letters(
        self,
        *,
        max_messages: int,
        max_wait_time: float,
        transformer: DeadLetterTransformer | None,
    ) -> list[dict[str, Any]]:
        inspected: list[dict[str, Any]] = []
        handled = 0
        async with self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            while handled < max_messages:
                try:
                    async with self._client.get_queue_receiver(
                        queue_name=self._queue_name,
                        sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                        session_id=NEXT_AVAILABLE_SESSION,
                        max_wait_time=max_wait_time,
                    ) as receiver:
                        messages = await receiver.receive_messages(
                            max_message_count=max_messages - handled,
                            max_wait_time=max_wait_time,
                        )
                        if not messages:
                            break
                        for message in messages:
                            handled += 1
                            body = message_body_text(message)
                            details = {
                                "message_id": message.message_id,
                                "reason": message.dead_letter_reason,
                                "description": message.dead_letter_error_description,
                                "body": body,
                            }
                            inspected.append(details)
                            LOGGER.warning("DLQ message: %s", details)
                            if transformer is None:
                                await receiver.abandon_message(message)
                                return inspected

                            try:
                                repaired = transformer(body, message)
                                await sender.send_messages(
                                    order_message(
                                        repaired,
                                        high_priority=False,
                                        retried_from_dlq=True,
                                    )
                                )
                                await receiver.complete_message(message)
                            except (TypeError, ValueError) as exc:
                                LOGGER.error(
                                    "DLQ message %s could not be repaired: %s",
                                    message.message_id,
                                    exc,
                                )
                                await receiver.abandon_message(message)
                except OperationTimeoutError:
                    break
        return inspected
