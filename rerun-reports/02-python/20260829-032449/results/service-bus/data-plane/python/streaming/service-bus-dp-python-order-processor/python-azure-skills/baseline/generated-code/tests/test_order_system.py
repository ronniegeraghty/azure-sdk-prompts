from __future__ import annotations

import asyncio
import json
import unittest
from datetime import datetime, timezone
from unittest.mock import AsyncMock, Mock

from azure.servicebus import ServiceBusMessage

from order_system.async_processor import AsyncOrderProcessor
from order_system.messages import HIGH_PRIORITY_DELAY, create_order_message
from order_system.model import Order, OrderStatus
from order_system.processor import OrderProcessor


class OrderModelTests(unittest.TestCase):
    def test_json_round_trip(self) -> None:
        order = Order(
            "order-1",
            "Ada Lovelace",
            "Keyboard",
            2,
            259.98,
            OrderStatus.PROCESSING,
        )

        self.assertEqual(Order.from_json(order.to_json()), order)
        self.assertEqual(json.loads(order.to_json())["status"], "processing")

    def test_high_priority_message_is_correlated_and_scheduled(self) -> None:
        before = datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY
        message = create_order_message(
            Order("order-2", "Grace Hopper", "Workstation", 1, 2_000),
            high_priority_threshold=1_000,
        )

        self.assertEqual(message.correlation_id, "order-2")
        self.assertEqual(message.session_id, "Grace Hopper")
        self.assertEqual(message.application_properties["priority"], "high")
        self.assertGreaterEqual(message.scheduled_enqueue_time_utc, before)

    def test_normal_priority_message_is_not_scheduled(self) -> None:
        message = create_order_message(
            Order("order-3", "Grace Hopper", "Mouse", 1, 50),
            high_priority_threshold=1_000,
        )

        self.assertEqual(message.application_properties["priority"], "normal")
        self.assertIsNone(message.scheduled_enqueue_time_utc)


class ProcessorTests(unittest.TestCase):
    @staticmethod
    def _processor(handler: Mock) -> OrderProcessor:
        processor = object.__new__(OrderProcessor)
        processor._handler = handler
        return processor

    def test_processing_failure_is_dead_lettered(self) -> None:
        receiver = Mock()
        message = ServiceBusMessage("{invalid")
        processor = self._processor(Mock())

        processed = processor._process_message(receiver, message)

        self.assertFalse(processed)
        receiver.dead_letter_message.assert_called_once()
        receiver.complete_message.assert_not_called()

    def test_successful_message_is_completed(self) -> None:
        receiver = Mock()
        order = Order("order-4", "Ada Lovelace", "Mouse", 1, 50)
        message = ServiceBusMessage(order.to_json())
        handler = Mock()
        processor = self._processor(handler)

        processed = processor._process_message(receiver, message)

        self.assertTrue(processed)
        handler.assert_called_once_with(order)
        receiver.complete_message.assert_called_once_with(message)


class AsyncProcessorTests(unittest.TestCase):
    def test_processing_failure_is_dead_lettered(self) -> None:
        async def run() -> None:
            processor = object.__new__(AsyncOrderProcessor)
            processor._handler = AsyncMock()
            receiver = AsyncMock()
            message = ServiceBusMessage("{invalid")

            processed = await processor._process_message(receiver, message)

            self.assertFalse(processed)
            receiver.dead_letter_message.assert_awaited_once()
            receiver.complete_message.assert_not_awaited()

        asyncio.run(run())


if __name__ == "__main__":
    unittest.main()
