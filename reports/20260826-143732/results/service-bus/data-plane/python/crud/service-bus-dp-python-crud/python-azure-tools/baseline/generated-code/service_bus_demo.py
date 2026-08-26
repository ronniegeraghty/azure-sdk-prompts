"""Azure Service Bus queue and topic messaging examples."""

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
        variable_names = (
            "AZURE_SERVICEBUS_CONNECTION_STRING",
            "AZURE_SERVICEBUS_QUEUE_NAME",
            "AZURE_SERVICEBUS_TOPIC_NAME",
            "AZURE_SERVICEBUS_SUBSCRIPTION_NAME",
        )
        missing = [name for name in variable_names if not os.environ.get(name)]
        if missing:
            raise RuntimeError(
                "Set the following environment variables: " + ", ".join(missing)
            )

        return cls(
            connection_string=os.environ["AZURE_SERVICEBUS_CONNECTION_STRING"],
            queue_name=os.environ["AZURE_SERVICEBUS_QUEUE_NAME"],
            topic_name=os.environ["AZURE_SERVICEBUS_TOPIC_NAME"],
            subscription_name=os.environ["AZURE_SERVICEBUS_SUBSCRIPTION_NAME"],
        )


def send_queue_batch(client: ServiceBusClient, queue_name: str) -> None:
    with client.get_queue_sender(queue_name=queue_name) as sender:
        batch = sender.create_message_batch()
        for message_number in range(1, 6):
            batch.add_message(
                ServiceBusMessage(
                    f"Sync queue message {message_number}",
                    application_properties={"sequence": message_number},
                )
            )
        sender.send_messages(batch)
    print("Sent a batch of 5 messages to the queue.")


def receive_queue_messages(client: ServiceBusClient, queue_name: str) -> None:
    with client.get_queue_receiver(
        queue_name=queue_name,
        max_wait_time=10,
    ) as receiver:
        messages = receiver.receive_messages(max_message_count=5, max_wait_time=10)
        for message in messages:
            print(f"Processing queue message: {message}")
            receiver.complete_message(message)
    print(f"Completed {len(messages)} queue message(s).")


def send_topic_message(client: ServiceBusClient, topic_name: str) -> None:
    with client.get_topic_sender(topic_name=topic_name) as sender:
        sender.send_messages(ServiceBusMessage("Sync topic message"))
    print("Sent a message to the topic.")


def receive_subscription_message(
    client: ServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    with client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
        max_wait_time=10,
    ) as receiver:
        messages = receiver.receive_messages(max_message_count=1, max_wait_time=10)
        for message in messages:
            print(f"Processing subscription message: {message}")
            receiver.complete_message(message)
    print(f"Completed {len(messages)} subscription message(s).")


def run_sync(settings: Settings) -> None:
    with ServiceBusClient.from_connection_string(
        conn_str=settings.connection_string
    ) as client:
        send_queue_batch(client, settings.queue_name)
        receive_queue_messages(client, settings.queue_name)
        send_topic_message(client, settings.topic_name)
        receive_subscription_message(
            client,
            settings.topic_name,
            settings.subscription_name,
        )


async def send_queue_batch_async(
    client: AsyncServiceBusClient,
    queue_name: str,
) -> None:
    async with client.get_queue_sender(queue_name=queue_name) as sender:
        batch = await sender.create_message_batch()
        for message_number in range(1, 6):
            batch.add_message(
                ServiceBusMessage(
                    f"Async queue message {message_number}",
                    application_properties={"sequence": message_number},
                )
            )
        await sender.send_messages(batch)
    print("Asynchronously sent a batch of 5 messages to the queue.")


async def send_topic_message_async(
    client: AsyncServiceBusClient,
    topic_name: str,
) -> None:
    async with client.get_topic_sender(topic_name=topic_name) as sender:
        await sender.send_messages(ServiceBusMessage("Async topic message"))
    print("Asynchronously sent a message to the topic.")


async def receive_queue_messages_async(
    client: AsyncServiceBusClient,
    queue_name: str,
) -> None:
    async with client.get_queue_receiver(
        queue_name=queue_name,
        max_wait_time=10,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=5,
            max_wait_time=10,
        )
        for message in messages:
            print(f"Async processing queue message: {message}")
            await receiver.complete_message(message)
    print(f"Asynchronously completed {len(messages)} queue message(s).")


async def receive_subscription_message_async(
    client: AsyncServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    async with client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
        max_wait_time=10,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=10,
        )
        for message in messages:
            print(f"Async processing subscription message: {message}")
            await receiver.complete_message(message)
    print(f"Asynchronously completed {len(messages)} subscription message(s).")


async def run_async(settings: Settings) -> None:
    async with AsyncServiceBusClient.from_connection_string(
        conn_str=settings.connection_string
    ) as client:
        # Independent queue and topic operations run concurrently for higher throughput.
        await asyncio.gather(
            send_queue_batch_async(client, settings.queue_name),
            send_topic_message_async(client, settings.topic_name),
        )
        await asyncio.gather(
            receive_queue_messages_async(client, settings.queue_name),
            receive_subscription_message_async(
                client,
                settings.topic_name,
                settings.subscription_name,
            ),
        )


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--mode",
        choices=("sync", "async", "both"),
        default="both",
        help="Messaging pattern to run (default: both).",
    )
    return parser.parse_args()


def main() -> None:
    args = parse_args()
    settings = Settings.from_environment()

    if args.mode in ("sync", "both"):
        run_sync(settings)
    if args.mode in ("async", "both"):
        asyncio.run(run_async(settings))


if __name__ == "__main__":
    main()
