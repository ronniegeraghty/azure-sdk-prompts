import asyncio
import unittest

from order_processing.processor import AsyncOrderProcessor, OrderProcessor


class FakeMessage:
    body = [b"{not-json"]
    message_id = "bad-1"


class FakeSyncReceiver:
    def __init__(self) -> None:
        self.messages = [FakeMessage()]
        self.dead_lettered = []

    def __enter__(self):
        return self

    def __exit__(self, *_args) -> None:
        return None

    def receive_messages(self, **_kwargs):
        messages, self.messages = self.messages, []
        return messages

    def complete_message(self, message) -> None:
        raise AssertionError(f"invalid message {message.message_id} was completed")

    def dead_letter_message(self, message, **kwargs) -> None:
        self.dead_lettered.append((message, kwargs))


class FakeSyncClient:
    def __init__(self, receiver: FakeSyncReceiver) -> None:
        self.receiver = receiver

    def get_queue_receiver(self, **_kwargs):
        return self.receiver


class FakeAsyncReceiver(FakeSyncReceiver):
    async def __aenter__(self):
        return self

    async def __aexit__(self, *_args) -> None:
        return None

    async def receive_messages(self, **_kwargs):
        messages, self.messages = self.messages, []
        return messages

    async def complete_message(self, message) -> None:
        raise AssertionError(f"invalid message {message.message_id} was completed")

    async def dead_letter_message(self, message, **kwargs) -> None:
        self.dead_lettered.append((message, kwargs))


class FakeAsyncClient:
    def __init__(self, receiver: FakeAsyncReceiver) -> None:
        self.receiver = receiver

    def get_queue_receiver(self, **_kwargs):
        return self.receiver


class ProcessorTests(unittest.TestCase):
    def test_sync_dead_letters_invalid_json(self) -> None:
        receiver = FakeSyncReceiver()
        processor = OrderProcessor(FakeSyncClient(receiver), "orders")

        completed = processor.process(max_messages=1)

        self.assertEqual(completed, [])
        self.assertEqual(len(receiver.dead_lettered), 1)
        self.assertEqual(
            receiver.dead_lettered[0][1]["reason"],
            "OrderDeserializationFailed",
        )

    def test_async_dead_letters_invalid_json(self) -> None:
        async def run() -> None:
            receiver = FakeAsyncReceiver()
            processor = AsyncOrderProcessor(FakeAsyncClient(receiver), "orders")

            completed = await processor.process(max_messages=1)

            self.assertEqual(completed, [])
            self.assertEqual(len(receiver.dead_lettered), 1)
            self.assertEqual(
                receiver.dead_lettered[0][1]["reason"],
                "OrderDeserializationFailed",
            )

        asyncio.run(run())


if __name__ == "__main__":
    unittest.main()
