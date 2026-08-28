"""Asynchronous Azure Service Bus queue and topic/subscription example."""

from __future__ import annotations

import argparse
import asyncio
import os
from dataclasses import dataclass

from azure.servicebus import ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient

BATCH_SIZE = 5


@dataclass(frozen=True)
class ServiceBusSettings:
    connection_string: str
    queue_name: str
    topic_name: str
    subscription_name: str

    @classmethod
    def from_environment(cls) -> "ServiceBusSettings":
        names = (
            "SERVICEBUS_CONNECTION_STR",
            "SERVICEBUS_QUEUE_NAME",
            "SERVICEBUS_TOPIC_NAME",
            "SERVICEBUS_SUBSCRIPTION_NAME",
        )
        missing = [name for name in names if not os.environ.get(name)]
        if missing:
            raise RuntimeError(
                "Missing required environment variables: " + ", ".join(missing)
            )

        return cls(
            connection_string=os.environ["SERVICEBUS_CONNECTION_STR"],
            queue_name=os.environ["SERVICEBUS_QUEUE_NAME"],
            topic_name=os.environ["SERVICEBUS_TOPIC_NAME"],
            subscription_name=os.environ["SERVICEBUS_SUBSCRIPTION_NAME"],
        )


async def send_queue_batch(
    client: ServiceBusClient, queue_name: str
) -> None:
    sender = client.get_queue_sender(queue_name=queue_name)
    async with sender:
        batch = await sender.create_message_batch()
        for index in range(1, BATCH_SIZE + 1):
            batch.add_message(
                ServiceBusMessage(
                    f"Async queue message {index}",
                    message_id=f"async-queue-{index}",
                    application_properties={"sequence": index},
                )
            )
        await sender.send_messages(batch)
        print(f"Sent an async queue batch containing {BATCH_SIZE} messages.")


async def receive_and_complete_queue(
    client: ServiceBusClient, queue_name: str
) -> None:
    receiver = client.get_queue_receiver(
        queue_name=queue_name,
        max_wait_time=10,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=BATCH_SIZE,
            max_wait_time=10,
        )
        for message in messages:
            print(f"Processed async queue message: {message}")
            await receiver.complete_message(message)
        print(f"Completed {len(messages)} async queue messages.")


async def send_topic_message(
    client: ServiceBusClient, topic_name: str
) -> None:
    sender = client.get_topic_sender(topic_name=topic_name)
    async with sender:
        await sender.send_messages(
            ServiceBusMessage(
                "Async topic message",
                message_id="async-topic-1",
                subject="async-demo",
            )
        )
        print("Sent one async topic message.")


async def receive_and_complete_subscription(
    client: ServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    receiver = client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
        max_wait_time=10,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=10,
        )
        for message in messages:
            print(f"Processed async subscription message: {message}")
            await receiver.complete_message(message)
        print(f"Completed {len(messages)} async subscription messages.")


async def run(settings: ServiceBusSettings) -> None:
    client = ServiceBusClient.from_connection_string(
        conn_str=settings.connection_string
    )
    async with client:
        await send_queue_batch(client, settings.queue_name)
        await receive_and_complete_queue(client, settings.queue_name)
        await send_topic_message(client, settings.topic_name)
        await receive_and_complete_subscription(
            client,
            settings.topic_name,
            settings.subscription_name,
        )


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--execute",
        action="store_true",
        help="Connect to Azure Service Bus; without this flag, perform a dry run.",
    )
    args = parser.parse_args()

    if not args.execute:
        print(
            "Dry run: would asynchronously send a 5-message queue batch, receive "
            "and complete the messages, then send to a topic and receive from a "
            "subscription."
        )
        return

    asyncio.run(run(ServiceBusSettings.from_environment()))


if __name__ == "__main__":
    main()
