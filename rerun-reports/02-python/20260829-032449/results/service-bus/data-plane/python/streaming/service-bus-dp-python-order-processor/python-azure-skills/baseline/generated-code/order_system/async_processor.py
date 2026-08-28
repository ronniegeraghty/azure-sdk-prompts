from __future__ import annotations

import logging
from collections.abc import Awaitable, Callable
from typing import Any

from azure.servicebus import (
    NEXT_AVAILABLE_SESSION,
    ServiceBusMessage,
    ServiceBusSubQueue,
)
from azure.servicebus.aio import ServiceBusClient
from azure.servicebus.exceptions import OperationTimeoutError

from .model import Order, OrderStatus
from .processor import DeadLetterRecord, message_body

logger = logging.getLogger(__name__)


async def default_order_handler(order: Order) -> None:
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
    logger.info("Completed order %s", order.order_id)


class AsyncOrderProcessor:
    def __init__(
        self,
        fully_qualified_namespace: str,
        queue_name: str,
        credential: Any,
        handler: Callable[[Order], Awaitable[None]] = default_order_handler,
    ) -> None:
        self._client = ServiceBusClient(
            fully_qualified_namespace=fully_qualified_namespace,
            credential=credential,
        )
        self._queue_name = queue_name
        self._handler = handler

    async def __aenter__(self) -> AsyncOrderProcessor:
        return self

    async def __aexit__(self, *args: object) -> None:
        await self.close()

    async def close(self) -> None:
        await self._client.close()

    async def process_available_orders(
        self,
        max_sessions: int | None = None,
        max_wait_time: float = 5,
    ) -> int:
        processed = 0
        sessions = 0
        while max_sessions is None or sessions < max_sessions:
            try:
                receiver = self._client.get_queue_receiver(
                    queue_name=self._queue_name,
                    session_id=NEXT_AVAILABLE_SESSION,
                    max_wait_time=max_wait_time,
                )
                async with receiver:
                    sessions += 1
                    while messages := await receiver.receive_messages(
                        max_message_count=50,
                        max_wait_time=max_wait_time,
                    ):
                        for message in messages:
                            if await self._process_message(receiver, message):
                                processed += 1
            except OperationTimeoutError:
                break
        return processed

    async def _process_message(self, receiver: Any, message: Any) -> bool:
        try:
            order = Order.from_json(message_body(message))
            await self._handler(order)
        except Exception as exc:
            reason = f"{type(exc).__name__}: order processing failed"
            logger.exception("Dead-lettering invalid order message")
            await receiver.dead_letter_message(
                message,
                reason=reason[:4096],
                error_description=str(exc)[:4096],
            )
            return False
        await receiver.complete_message(message)
        return True

    async def read_dead_letters(
        self,
        reprocess: bool = False,
        max_sessions: int | None = None,
        max_wait_time: float = 5,
    ) -> list[DeadLetterRecord]:
        records: list[DeadLetterRecord] = []
        sender = (
            self._client.get_queue_sender(self._queue_name)
            if reprocess
            else None
        )
        if sender is not None:
            await sender.__aenter__()
        try:
            sessions = 0
            while max_sessions is None or sessions < max_sessions:
                try:
                    receiver = self._client.get_queue_receiver(
                        queue_name=self._queue_name,
                        sub_queue=ServiceBusSubQueue.DEAD_LETTER,
                        session_id=NEXT_AVAILABLE_SESSION,
                        max_wait_time=max_wait_time,
                    )
                    async with receiver:
                        sessions += 1
                        while messages := await receiver.receive_messages(
                            max_message_count=50,
                            max_wait_time=max_wait_time,
                        ):
                            for message in messages:
                                body = message_body(message)
                                records.append(
                                    DeadLetterRecord(
                                        message_id=message.message_id,
                                        reason=message.dead_letter_reason,
                                        description=message.dead_letter_error_description,
                                        body=body.decode("utf-8", errors="replace"),
                                    )
                                )
                                logger.info(
                                    "DLQ message %s: %s",
                                    message.message_id,
                                    message.dead_letter_reason,
                                )
                                if sender is not None:
                                    try:
                                        order = Order.from_json(body)
                                        await sender.send_messages(
                                            ServiceBusMessage(
                                                order.to_json(),
                                                content_type="application/json",
                                                correlation_id=order.order_id,
                                                session_id=order.customer_name,
                                            )
                                        )
                                        await receiver.complete_message(message)
                                    except (KeyError, TypeError, ValueError) as exc:
                                        logger.warning(
                                            "DLQ message %s cannot be reprocessed: %s",
                                            message.message_id,
                                            exc,
                                        )
                except OperationTimeoutError:
                    break
        finally:
            if sender is not None:
                await sender.__aexit__(None, None, None)
        return records
