"""Synchronous and asynchronous order senders."""

from __future__ import annotations

from collections import defaultdict
from collections.abc import Iterable
from datetime import datetime, timedelta, timezone

from azure.servicebus import ServiceBusMessageBatch

from .messages import customer_session_id, order_message
from .model import Order

FRAUD_REVIEW_DELAY = timedelta(seconds=30)


class OrderSender:
    def __init__(self, client, queue_name: str, high_priority_threshold: float) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold
        self._blocked_sessions_until: dict[str, datetime] = {}

    def _schedule_for(self, order: Order, now: datetime) -> datetime | None:
        session_id = customer_session_id(order.customer_name)
        blocked_until = self._blocked_sessions_until.get(session_id)
        if order.total_price > self._high_priority_threshold:
            blocked_until = max(blocked_until or now, now + FRAUD_REVIEW_DELAY)
            self._blocked_sessions_until[session_id] = blocked_until
        elif blocked_until is not None and blocked_until <= now:
            self._blocked_sessions_until.pop(session_id, None)
            blocked_until = None
        return blocked_until

    def send_order(self, order: Order) -> None:
        now = datetime.now(timezone.utc)
        scheduled_at = self._schedule_for(order, now)
        message = order_message(
            order,
            high_priority=order.total_price > self._high_priority_threshold,
            scheduled_at=scheduled_at,
        )
        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            if scheduled_at is None:
                sender.send_messages(message)
            else:
                sender.schedule_messages(message, scheduled_at)

    def send_orders(self, orders: Iterable[Order]) -> None:
        grouped: dict[str, list[Order]] = defaultdict(list)
        for order in orders:
            grouped[customer_session_id(order.customer_name)].append(order)

        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:
            for customer_orders in grouped.values():
                self._send_customer_orders(sender, customer_orders)

    def _send_customer_orders(self, sender, orders: list[Order]) -> None:
        now = datetime.now(timezone.utc)
        immediate = sender.create_message_batch()

        for order in orders:
            scheduled_at = self._schedule_for(order, now)
            message = order_message(
                order,
                high_priority=order.total_price > self._high_priority_threshold,
                scheduled_at=scheduled_at,
            )
            if scheduled_at is not None:
                if len(immediate):
                    sender.send_messages(immediate)
                    immediate = sender.create_message_batch()
                sender.schedule_messages(message, scheduled_at)
                continue

            immediate = self._add_to_batch(sender, immediate, message)

        if len(immediate):
            sender.send_messages(immediate)

    @staticmethod
    def _add_to_batch(sender, batch: ServiceBusMessageBatch, message):
        try:
            batch.add_message(message)
            return batch
        except ValueError:
            if not len(batch):
                raise ValueError("an order message exceeds the queue message size limit")
            sender.send_messages(batch)
            next_batch = sender.create_message_batch()
            try:
                next_batch.add_message(message)
            except ValueError as exc:
                raise ValueError(
                    "an order message exceeds the queue message size limit"
                ) from exc
            return next_batch


class AsyncOrderSender:
    def __init__(self, client, queue_name: str, high_priority_threshold: float) -> None:
        self._client = client
        self._queue_name = queue_name
        self._high_priority_threshold = high_priority_threshold
        self._blocked_sessions_until: dict[str, datetime] = {}

    def _schedule_for(self, order: Order, now: datetime) -> datetime | None:
        session_id = customer_session_id(order.customer_name)
        blocked_until = self._blocked_sessions_until.get(session_id)
        if order.total_price > self._high_priority_threshold:
            blocked_until = max(blocked_until or now, now + FRAUD_REVIEW_DELAY)
            self._blocked_sessions_until[session_id] = blocked_until
        elif blocked_until is not None and blocked_until <= now:
            self._blocked_sessions_until.pop(session_id, None)
            blocked_until = None
        return blocked_until

    async def send_order(self, order: Order) -> None:
        now = datetime.now(timezone.utc)
        scheduled_at = self._schedule_for(order, now)
        message = order_message(
            order,
            high_priority=order.total_price > self._high_priority_threshold,
            scheduled_at=scheduled_at,
        )
        async with self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            if scheduled_at is None:
                await sender.send_messages(message)
            else:
                await sender.schedule_messages(message, scheduled_at)

    async def send_orders(self, orders: Iterable[Order]) -> None:
        grouped: dict[str, list[Order]] = defaultdict(list)
        for order in orders:
            grouped[customer_session_id(order.customer_name)].append(order)

        async with self._client.get_queue_sender(
            queue_name=self._queue_name
        ) as sender:
            for customer_orders in grouped.values():
                await self._send_customer_orders(sender, customer_orders)

    async def _send_customer_orders(self, sender, orders: list[Order]) -> None:
        now = datetime.now(timezone.utc)
        immediate = await sender.create_message_batch()

        for order in orders:
            scheduled_at = self._schedule_for(order, now)
            message = order_message(
                order,
                high_priority=order.total_price > self._high_priority_threshold,
                scheduled_at=scheduled_at,
            )
            if scheduled_at is not None:
                if len(immediate):
                    await sender.send_messages(immediate)
                    immediate = await sender.create_message_batch()
                await sender.schedule_messages(message, scheduled_at)
                continue

            immediate = await self._add_to_batch(sender, immediate, message)

        if len(immediate):
            await sender.send_messages(immediate)

    @staticmethod
    async def _add_to_batch(sender, batch: ServiceBusMessageBatch, message):
        try:
            batch.add_message(message)
            return batch
        except ValueError:
            if not len(batch):
                raise ValueError("an order message exceeds the queue message size limit")
            await sender.send_messages(batch)
            next_batch = await sender.create_message_batch()
            try:
                next_batch.add_message(message)
            except ValueError as exc:
                raise ValueError(
                    "an order message exceeds the queue message size limit"
                ) from exc
            return next_batch
