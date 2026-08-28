"""Azure Service Bus queue and topic/subscription messaging examples."""

import asyncio
import logging
import os
from dataclasses import dataclass

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient


MESSAGE_COUNT = 5
RECEIVE_WAIT_SECONDS = 10


@dataclass(frozen=True)
class Settings:
    fully_qualified_namespace: str
    queue_name: str
    topic_name: str
    subscription_name: str

    @classmethod
    def from_environment(cls) -> "Settings":
        return cls(
            fully_qualified_namespace=os.environ[
                "SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"
            ],
            queue_name=os.environ["SERVICEBUS_QUEUE_NAME"],
            topic_name=os.environ["SERVICEBUS_TOPIC_NAME"],
            subscription_name=os.environ["SERVICEBUS_SUBSCRIPTION_NAME"],
        )


def process_message(source: str, message: object) -> None:
    """Replace this function with application-specific processing."""
    logging.info("Processed message from %s: %s", source, message)


def run_sync_queue_demo(settings: Settings) -> None:
    with DefaultAzureCredential() as credential:
        with ServiceBusClient(
            fully_qualified_namespace=settings.fully_qualified_namespace,
            credential=credential,
        ) as client:
            with client.get_queue_sender(settings.queue_name) as sender:
                batch = sender.create_message_batch()
                for index in range(1, MESSAGE_COUNT + 1):
                    batch.add_message(
                        ServiceBusMessage(
                            f"Synchronous queue message {index}",
                            message_id=f"sync-queue-{index}",
                        )
                    )
                sender.send_messages(batch)
                logging.info("Sent a batch of %d queue messages", MESSAGE_COUNT)

            with client.get_queue_receiver(
                queue_name=settings.queue_name,
                max_wait_time=RECEIVE_WAIT_SECONDS,
            ) as receiver:
                messages = receiver.receive_messages(
                    max_message_count=MESSAGE_COUNT,
                    max_wait_time=RECEIVE_WAIT_SECONDS,
                )
                for message in messages:
                    process_message(settings.queue_name, message)
                    receiver.complete_message(message)
                logging.info("Completed %d queue messages", len(messages))


def run_sync_topic_demo(settings: Settings) -> None:
    with DefaultAzureCredential() as credential:
        with ServiceBusClient(
            fully_qualified_namespace=settings.fully_qualified_namespace,
            credential=credential,
        ) as client:
            with client.get_topic_sender(settings.topic_name) as sender:
                sender.send_messages(
                    ServiceBusMessage(
                        "Synchronous topic message",
                        message_id="sync-topic-1",
                    )
                )
                logging.info("Sent a message to topic %s", settings.topic_name)

            with client.get_subscription_receiver(
                topic_name=settings.topic_name,
                subscription_name=settings.subscription_name,
                max_wait_time=RECEIVE_WAIT_SECONDS,
            ) as receiver:
                messages = receiver.receive_messages(
                    max_message_count=1,
                    max_wait_time=RECEIVE_WAIT_SECONDS,
                )
                for message in messages:
                    process_message(settings.subscription_name, message)
                    receiver.complete_message(message)
                logging.info(
                    "Completed %d message(s) from subscription %s",
                    len(messages),
                    settings.subscription_name,
                )


async def send_async_queue_batch(
    client: AsyncServiceBusClient, settings: Settings
) -> None:
    async with client.get_queue_sender(settings.queue_name) as sender:
        batch = await sender.create_message_batch()
        for index in range(1, MESSAGE_COUNT + 1):
            batch.add_message(
                ServiceBusMessage(
                    f"Asynchronous queue message {index}",
                    message_id=f"async-queue-{index}",
                )
            )
        await sender.send_messages(batch)
        logging.info("Asynchronously sent a batch of %d messages", MESSAGE_COUNT)


async def receive_async_queue(
    client: AsyncServiceBusClient, settings: Settings
) -> None:
    async with client.get_queue_receiver(
        queue_name=settings.queue_name,
        max_wait_time=RECEIVE_WAIT_SECONDS,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=MESSAGE_COUNT,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        for message in messages:
            process_message(settings.queue_name, message)
            await receiver.complete_message(message)
        logging.info("Asynchronously completed %d queue messages", len(messages))


async def send_async_topic_message(
    client: AsyncServiceBusClient, settings: Settings
) -> None:
    async with client.get_topic_sender(settings.topic_name) as sender:
        await sender.send_messages(
            ServiceBusMessage(
                "Asynchronous topic message",
                message_id="async-topic-1",
            )
        )
        logging.info("Asynchronously sent a topic message")


async def receive_async_subscription(
    client: AsyncServiceBusClient, settings: Settings
) -> None:
    async with client.get_subscription_receiver(
        topic_name=settings.topic_name,
        subscription_name=settings.subscription_name,
        max_wait_time=RECEIVE_WAIT_SECONDS,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        for message in messages:
            process_message(settings.subscription_name, message)
            await receiver.complete_message(message)
        logging.info(
            "Asynchronously completed %d subscription message(s)", len(messages)
        )


async def run_async_demo(settings: Settings) -> None:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncServiceBusClient(
            fully_qualified_namespace=settings.fully_qualified_namespace,
            credential=credential,
        ) as client:
            # Independent sends and receives overlap instead of blocking one another.
            await asyncio.gather(
                send_async_queue_batch(client, settings),
                send_async_topic_message(client, settings),
            )
            await asyncio.gather(
                receive_async_queue(client, settings),
                receive_async_subscription(client, settings),
            )


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    settings = Settings.from_environment()

    run_sync_queue_demo(settings)
    run_sync_topic_demo(settings)
    asyncio.run(run_async_demo(settings))


if __name__ == "__main__":
    main()
