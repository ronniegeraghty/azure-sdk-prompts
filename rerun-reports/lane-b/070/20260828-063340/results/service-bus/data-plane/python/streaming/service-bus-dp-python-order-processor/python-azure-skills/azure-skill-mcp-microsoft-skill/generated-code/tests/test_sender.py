import asyncio
import unittest

from order_processing.model import Order
from order_processing.sender import AsyncOrderSender, OrderSender


class FakeBatch:
    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.messages = []

    def add_message(self, message) -> None:
        if len(self.messages) >= self.capacity:
            raise ValueError("batch full")
        self.messages.append(message)

    def __len__(self) -> int:
        return len(self.messages)


class FakeSyncSender:
    def __init__(self, capacity: int = 1) -> None:
        self.capacity = capacity
        self.sent = []
        self.scheduled = []

    def __enter__(self):
        return self

    def __exit__(self, *_args) -> None:
        return None

    def create_message_batch(self) -> FakeBatch:
        return FakeBatch(self.capacity)

    def send_messages(self, messages) -> None:
        self.sent.append(messages)

    def schedule_messages(self, message, when) -> None:
        self.scheduled.append((message, when))


class FakeSyncClient:
    def __init__(self, sender: FakeSyncSender) -> None:
        self.sender = sender

    def get_queue_sender(self, **_kwargs):
        return self.sender


class FakeAsyncSender(FakeSyncSender):
    async def __aenter__(self):
        return self

    async def __aexit__(self, *_args) -> None:
        return None

    async def create_message_batch(self) -> FakeBatch:
        return FakeBatch(self.capacity)

    async def send_messages(self, messages) -> None:
        self.sent.append(messages)

    async def schedule_messages(self, message, when) -> None:
        self.scheduled.append((message, when))


class FakeAsyncClient:
    def __init__(self, sender: FakeAsyncSender) -> None:
        self.sender = sender

    def get_queue_sender(self, **_kwargs):
        return self.sender


def orders() -> list[Order]:
    return [
        Order("o-1", "Ada", "Keyboard", 1, 100.0),
        Order("o-2", "Ada", "Mouse", 1, 50.0),
    ]


class SyncSenderTests(unittest.TestCase):
    def test_splits_full_batches(self) -> None:
        fake = FakeSyncSender(capacity=1)
        sender = OrderSender(FakeSyncClient(fake), "orders", 1_000.0)

        sender.send_orders(orders())

        self.assertEqual(len(fake.sent), 2)
        self.assertTrue(all(len(batch) == 1 for batch in fake.sent))

    def test_schedules_high_priority_and_following_customer_order(self) -> None:
        fake = FakeSyncSender(capacity=10)
        sender = OrderSender(FakeSyncClient(fake), "orders", 75.0)

        sender.send_orders(orders())

        self.assertEqual(len(fake.scheduled), 2)
        self.assertEqual(fake.scheduled[0][1], fake.scheduled[1][1])
        first_message = fake.scheduled[0][0]
        self.assertEqual(first_message.correlation_id, "o-1")
        self.assertEqual(first_message.application_properties["priority"], "high")
        self.assertIsNotNone(first_message.scheduled_enqueue_time_utc)


class AsyncSenderTests(unittest.TestCase):
    def test_splits_full_batches(self) -> None:
        async def run() -> None:
            fake = FakeAsyncSender(capacity=1)
            sender = AsyncOrderSender(FakeAsyncClient(fake), "orders", 1_000.0)

            await sender.send_orders(orders())

            self.assertEqual(len(fake.sent), 2)
            self.assertTrue(all(len(batch) == 1 for batch in fake.sent))

        asyncio.run(run())


if __name__ == "__main__":
    unittest.main()
