"""Azure Service Bus queue, batch, async, topic, and subscription examples."""

from __future__ import annotations

import argparse
import asyncio
import os
from dataclasses import dataclass

from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient


@dataclass(frozen=True)
class Settings:
    connection_string: str
    queue_name: str
    topic_name: str
    subscription_name: str

    @classmethod
    def from_environment(cls) -> "Settings":
        names = (
            "SERVICE_BUS_CONNECTION_STRING",
            "SERVICE_BUS_QUEUE_NAME",
            "SERVICE_BUS_TOPIC_NAME",
            "SERVICE_BUS_SUBSCRIPTION_NAME",
        )
        missing = [name for name in names if not os.getenv(name)]
        if missing:
            raise RuntimeError(
                "Set the following environment variables before running: "
                + ", ".join(missing)
            )

        return cls(
            connection_string=os.environ["SERVICE_BUS_CONNECTION_STRING"],
            queue_name=os.environ["SERVICE_BUS_QUEUE_NAME"],
            topic_name=os.environ["SERVICE_BUS_TOPIC_NAME"],
            subscription_name=os.environ["SERVICE_BUS_SUBSCRIPTION_NAME"],
        )


def add_five_messages_to_batch(sender: object) -> object:
    """Create one ServiceBusMessageBatch containing exactly five messages."""
    batch = sender.create_message_batch()
    for index in range(1, 6):
        message = ServiceBusMessage(
            f"Synchronous queue message {index}",
            application_properties={"sequence": index, "example": "sync-queue"},
        )
        if not batch.try_add_message(message):
            raise RuntimeError(
                "The configured Service Bus message size limit cannot fit all 5 messages."
            )
    return batch


def run_sync_queue(settings: Settings) -> None:
    """Send a five-message batch, receive it, and complete each message."""
    with ServiceBusClient.from_connection_string(
        settings.connection_string
    ) as client:
        with client.get_queue_sender(settings.queue_name) as sender:
            batch = add_five_messages_to_batch(sender)
            sender.send_messages(batch)
            print(f"Sent {len(batch)} messages to queue {settings.queue_name!r}.")

        with client.get_queue_receiver(
            settings.queue_name, max_wait_time=10
        ) as receiver:
            messages = receiver.receive_messages(max_message_count=5, max_wait_time=10)
            for message in messages:
                print(f"Processed queue message: {message}")
                receiver.complete_message(message)
            print(f"Completed {len(messages)} queue messages.")


def run_sync_topic(settings: Settings) -> None:
    """Send to a topic, then receive and complete from its subscription."""
    with ServiceBusClient.from_connection_string(
        settings.connection_string
    ) as client:
        with client.get_topic_sender(settings.topic_name) as sender:
            sender.send_messages(
                ServiceBusMessage(
                    "Topic message",
                    subject="service-bus-demo",
                    application_properties={"example": "topic"},
                )
            )
            print(f"Sent a message to topic {settings.topic_name!r}.")

        with client.get_subscription_receiver(
            topic_name=settings.topic_name,
            subscription_name=settings.subscription_name,
            max_wait_time=10,
        ) as receiver:
            messages = receiver.receive_messages(max_message_count=1, max_wait_time=10)
            for message in messages:
                print(f"Processed subscription message: {message}")
                receiver.complete_message(message)
            print(f"Completed {len(messages)} subscription messages.")


async def run_async_queue(settings: Settings) -> None:
    """Use azure.servicebus.aio to send and process messages asynchronously."""
    client = AsyncServiceBusClient.from_connection_string(
        settings.connection_string
    )
    async with client:
        sender = client.get_queue_sender(settings.queue_name)
        async with sender:
            batch = await sender.create_message_batch()
            for index in range(1, 6):
                message = ServiceBusMessage(
                    f"Asynchronous queue message {index}",
                    application_properties={
                        "sequence": index,
                        "example": "async-queue",
                    },
                )
                if not batch.try_add_message(message):
                    raise RuntimeError(
                        "The configured Service Bus message size limit cannot fit "
                        "all 5 messages."
                    )
            await sender.send_messages(batch)
            print(f"Sent {len(batch)} messages asynchronously.")

        receiver = client.get_queue_receiver(
            settings.queue_name, max_wait_time=10
        )
        async with receiver:
            messages = await receiver.receive_messages(
                max_message_count=5, max_wait_time=10
            )

            async def process_and_complete(message: ServiceBusMessage) -> None:
                print(f"Processed asynchronously: {message}")
                await receiver.complete_message(message)

            await asyncio.gather(
                *(process_and_complete(message) for message in messages)
            )
            print(f"Completed {len(messages)} messages asynchronously.")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "mode",
        choices=("sync-queue", "async-queue", "topic", "all"),
        nargs="?",
        default="all",
        help="Example to run (default: all).",
    )
    return parser.parse_args()


def main() -> None:
    mode = parse_args().mode
    settings = Settings.from_environment()

    if mode in ("sync-queue", "all"):
        run_sync_queue(settings)
    if mode in ("async-queue", "all"):
        asyncio.run(run_async_queue(settings))
    if mode in ("topic", "all"):
        run_sync_topic(settings)


if __name__ == "__main__":
    main()
