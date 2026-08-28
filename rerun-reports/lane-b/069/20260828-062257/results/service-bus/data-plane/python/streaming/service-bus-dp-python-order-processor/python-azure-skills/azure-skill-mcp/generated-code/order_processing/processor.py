from __future__ import annotations

import logging
from collections.abc import Callable
from decimal import Decimal

from azure.servicebus import (
    NEXT_AVAILABLE_SESSION,
    ServiceBusClient,
    ServiceBusMessage,
    ServiceBusReceiveMode,
    ServiceBusReceivedMessage,
    ServiceBusSubQueue,
)
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient
from azure.servicebus.exceptions import OperationTimeoutError

from .message_factory import create_order_message
from .models import Order, OrderStatus

LOGGER = logging.getLogger(__name__)
OrderHandler = Callable[[Order], None]
DeadLetterRepair = Callable[[ServiceBusReceivedMessage], Order]


def _message_body(message: ServiceBusReceivedMessage) -> bytes:
    return b"".join(bytes(section) for section in message.body)


def _mark_completed(order: Order) -> Order:
    order.status = OrderStatus.PROCESSING
    LOGGER.info(
        "Processing order %s for %s: %s x %s",
        order.order_id,
        order.customer_name,
        order.quantity,
        order.product,
    )
    order.status = OrderStatus.COMPLETED
    return order


class OrderProcessor:
    def __init__(self, client: ServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    def process_orders(
        self,
        max_messages: int,
        handler: OrderHandler | None = None,
        wait_time: float = 5,
    ) -> list[Order]:
        processed: list[Order] = []
        handler = handler or (lambda order: None)

        while len(processed) < max_messages:
            try:
                receiver = self._client.get_queue_receiver(
                    self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,
                    max_wait_time=wait_time,
                )
                with receiver:
                    messages = receiver.receive_messages(
                        max_message_count=max_messages - len(processed),
                        max_wait_time=wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        try:
                            order = Order.from_json(_message_body(message))
                            handler(order)
                            _mark_completed(order)
                        except Exception as exc:
                            reason = f"{type(exc).__name__}: {exc}"[:4096]
                            LOGGER.exception(
                                "Order message %s failed; dead-lettering it",
                                message.message_id,
                            )
                            receiver.dead_letter_message(
                                message,
                                reason="OrderProcessingFailed",
                                error_description=reason,
                            )
                        else:
                            receiver.complete_message(message)
                            processed.append(order)
                            LOGGER.info("Completed order %s", order.order_id)
            except OperationTimeoutError:
                break
        return processed

    def inspect_dead_letters(
        self,
        max_messages: int = 20,
        reprocess: bool = False,
        repair: DeadLetterRepair | None = None,
        wait_time: float = 5,
    ) -> list[ServiceBusReceivedMessage]:
        inspected: list[ServiceBusReceivedMessage] = []
        while len(inspected) < max_messages:
            try:
                receiver = self._client.get_queue_receiver(
                    self._queue_name,
                    sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                    session_id=NEXT_AVAILABLE_SESSION,
                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,
                    max_wait_time=wait_time,
                )
                with receiver:
                    messages = receiver.receive_messages(
                        max_message_count=max_messages - len(inspected),
                        max_wait_time=wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        inspected.append(message)
                        LOGGER.info(
                            "Dead letter %s: reason=%s description=%s",
                            message.message_id,
                            message.dead_letter_reason,
                            message.dead_letter_error_description,
                        )
                        if reprocess:
                            self._reprocess_message(message, repair)
                            receiver.complete_message(message)
                        else:
                            receiver.abandon_message(message)
                    if not reprocess:
                        # Abandoned messages remain in this session for later inspection.
                        return inspected
            except OperationTimeoutError:
                break
        return inspected

    def _reprocess_message(
        self,
        message: ServiceBusReceivedMessage,
        repair: DeadLetterRepair | None,
    ) -> None:
        order = repair(message) if repair else Order.from_json(_message_body(message))
        with self._client.get_queue_sender(self._queue_name) as sender:
            sender.send_messages(
                create_order_message(order, high_priority_threshold=Decimal("Infinity"))
            )


class AsyncOrderProcessor:
    def __init__(self, client: AsyncServiceBusClient, queue_name: str) -> None:
        self._client = client
        self._queue_name = queue_name

    async def process_orders(
        self,
        max_messages: int,
        handler: OrderHandler | None = None,
        wait_time: float = 5,
    ) -> list[Order]:
        processed: list[Order] = []
        handler = handler or (lambda order: None)

        while len(processed) < max_messages:
            try:
                receiver = self._client.get_queue_receiver(
                    self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,
                    max_wait_time=wait_time,
                )
                async with receiver:
                    messages = await receiver.receive_messages(
                        max_message_count=max_messages - len(processed),
                        max_wait_time=wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        try:
                            order = Order.from_json(_message_body(message))
                            handler(order)
                            _mark_completed(order)
                        except Exception as exc:
                            reason = f"{type(exc).__name__}: {exc}"[:4096]
                            LOGGER.exception(
                                "Order message %s failed; dead-lettering it",
                                message.message_id,
                            )
                            await receiver.dead_letter_message(
                                message,
                                reason="OrderProcessingFailed",
                                error_description=reason,
                            )
                        else:
                            await receiver.complete_message(message)
                            processed.append(order)
                            LOGGER.info("Completed order %s", order.order_id)
            except OperationTimeoutError:
                break
        return processed

    async def inspect_dead_letters(
        self,
        max_messages: int = 20,
        reprocess: bool = False,
        repair: DeadLetterRepair | None = None,
        wait_time: float = 5,
    ) -> list[ServiceBusReceivedMessage]:
        inspected: list[ServiceBusReceivedMessage] = []
        while len(inspected) < max_messages:
            try:
                receiver = self._client.get_queue_receiver(
                    self._queue_name,
                    sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                    session_id=NEXT_AVAILABLE_SESSION,
                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,
                    max_wait_time=wait_time,
                )
                async with receiver:
                    messages = await receiver.receive_messages(
                        max_message_count=max_messages - len(inspected),
                        max_wait_time=wait_time,
                    )
                    if not messages:
                        break
                    for message in messages:
                        inspected.append(message)
                        LOGGER.info(
                            "Dead letter %s: reason=%s description=%s",
                            message.message_id,
                            message.dead_letter_reason,
                            message.dead_letter_error_description,
                        )
                        if reprocess:
                            await self._reprocess_message(message, repair)
                            await receiver.complete_message(message)
                        else:
                            await receiver.abandon_message(message)
                    if not reprocess:
                        # Abandoned messages remain in this session for later inspection.
                        return inspected
            except OperationTimeoutError:
                break
        return inspected

    async def _reprocess_message(
        self,
        message: ServiceBusReceivedMessage,
        repair: DeadLetterRepair | None,
    ) -> None:
        order = repair(message) if repair else Order.from_json(_message_body(message))
        async with self._client.get_queue_sender(self._queue_name) as sender:
            await sender.send_messages(
                create_order_message(order, high_priority_threshold=Decimal("Infinity"))
            )
