"""Session-aware synchronous and asynchronous order processors."""

from __future__ import annotations

import logging
from dataclasses import replace
from typing import Any
from uuid import uuid4

from azure.servicebus import (
    NEXT_AVAILABLE_SESSION,
    ServiceBusClient,
    ServiceBusMessage,
    ServiceBusSubQueue,
)
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import OperationTimeoutError

from order_model import Order, OrderDeserializationError, OrderStatus

LOGGER = logging.getLogger(__name__)
DEAD_LETTER_REASON = "OrderDeserializationFailed"


def _message_body_as_text(message: Any) -> str:
    body = message.body
    if isinstance(body, str):
        return body
    if isinstance(body, (bytes, bytearray)):
        return bytes(body).decode("utf-8")
    return b"".join(body).decode("utf-8")


def _process_message_body(message: Any) -> Order:
    order = Order.from_json(_message_body_as_text(message))
    processing_order = replace(order, status=OrderStatus.PROCESSING)
    LOGGER.info(
        "Processing order %s for customer %s",
        processing_order.order_id,
        processing_order.customer_name,
    )
    completed_order = replace(processing_order, status=OrderStatus.COMPLETED)
    LOGGER.info(
        "Completed order %s: %s x%d, total $%.2f",
        completed_order.order_id,
        completed_order.product,
        completed_order.quantity,
        completed_order.total_price,
    )
    return completed_order


def _retry_message(message: Any) -> ServiceBusMessage:
    properties = dict(message.application_properties or {})
    properties["reprocessed_from_dlq"] = True
    return ServiceBusMessage(
        _message_body_as_text(message),
        content_type=message.content_type or "application/json",
        correlation_id=message.correlation_id,
        message_id=f"{message.message_id}-retry-{uuid4()}",
        session_id=message.session_id,
        subject=message.subject,
        application_properties=properties,
    )


class OrderProcessor:
    def __init__(self, client: ServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    def process_available(
        self, max_wait_time: float = 5, max_messages_per_receive: int = 50
    ) -> int:
        processed = 0
        while True:
            receiver = self._client.get_queue_receiver(
                queue_name=self._queue_name,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=max_wait_time,
            )
            try:
                with receiver:
                    while True:
                        messages = receiver.receive_messages(
                            max_message_count=max_messages_per_receive,
                            max_wait_time=1,
                        )
                        if not messages:
                            break
                        for message in messages:
                            try:
                                _process_message_body(message)
                            except OrderDeserializationError as exc:
                                receiver.dead_letter_message(
                                    message,
                                    reason=DEAD_LETTER_REASON,
                                    error_description=str(exc)[:4096],
                                )
                                LOGGER.error(
                                    "Dead-lettered message %s: %s",
                                    message.message_id,
                                    exc,
                                )
                            else:
                                receiver.complete_message(message)
                                processed += 1
            except OperationTimeoutError:
                break
        return processed

    def process_dead_letters(
        self, reprocess: bool = False, max_wait_time: float = 5
    ) -> int:
        sender_context = (
            self._client.get_queue_sender(queue_name=self._queue_name)
            if reprocess
            else None
        )
        if sender_context is None:
            return self._drain_dead_letters(None, max_wait_time)
        with sender_context as sender:
            return self._drain_dead_letters(sender, max_wait_time)

    def _drain_dead_letters(self, sender: Any, max_wait_time: float) -> int:
        handled = 0
        while True:
            receiver = self._client.get_queue_receiver(
                queue_name=self._queue_name,
                sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=max_wait_time,
            )
            try:
                with receiver:
                    while True:
                        messages = receiver.receive_messages(
                            max_message_count=50, max_wait_time=1
                        )
                        if not messages:
                            break
                        for message in messages:
                            LOGGER.warning(
                                "DLQ message %s: reason=%s description=%s body=%s",
                                message.message_id,
                                message.dead_letter_reason,
                                message.dead_letter_error_description,
                                _message_body_as_text(message),
                            )
                            if sender is not None:
                                sender.send_messages(_retry_message(message))
                            receiver.complete_message(message)
                            handled += 1
            except OperationTimeoutError:
                break
        return handled


class AsyncOrderProcessor:
    def __init__(self, client: AsyncServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    async def process_available(
        self, max_wait_time: float = 5, max_messages_per_receive: int = 50
    ) -> int:
        processed = 0
        while True:
            receiver = self._client.get_queue_receiver(
                queue_name=self._queue_name,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=max_wait_time,
            )
            try:
                async with receiver:
                    while True:
                        messages = await receiver.receive_messages(
                            max_message_count=max_messages_per_receive,
                            max_wait_time=1,
                        )
                        if not messages:
                            break
                        for message in messages:
                            try:
                                _process_message_body(message)
                            except OrderDeserializationError as exc:
                                await receiver.dead_letter_message(
                                    message,
                                    reason=DEAD_LETTER_REASON,
                                    error_description=str(exc)[:4096],
                                )
                                LOGGER.error(
                                    "Dead-lettered message %s: %s",
                                    message.message_id,
                                    exc,
                                )
                            else:
                                await receiver.complete_message(message)
                                processed += 1
            except OperationTimeoutError:
                break
        return processed

    async def process_dead_letters(
        self, reprocess: bool = False, max_wait_time: float = 5
    ) -> int:
        if not reprocess:
            return await self._drain_dead_letters(None, max_wait_time)
        async with self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            return await self._drain_dead_letters(sender, max_wait_time)

    async def _drain_dead_letters(self, sender: Any, max_wait_time: float) -> int:
        handled = 0
        while True:
            receiver = self._client.get_queue_receiver(
                queue_name=self._queue_name,
                sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                session_id=NEXT_AVAILABLE_SESSION,
                max_wait_time=max_wait_time,
            )
            try:
                async with receiver:
                    while True:
                        messages = await receiver.receive_messages(
                            max_message_count=50, max_wait_time=1
                        )
                        if not messages:
                            break
                        for message in messages:
                            LOGGER.warning(
                                "DLQ message %s: reason=%s description=%s body=%s",
                                message.message_id,
                                message.dead_letter_reason,
                                message.dead_letter_error_description,
                                _message_body_as_text(message),
                            )
                            if sender is not None:
                                await sender.send_messages(_retry_message(message))
                            await receiver.complete_message(message)
                            handled += 1
            except OperationTimeoutError:
                break
        return handled
