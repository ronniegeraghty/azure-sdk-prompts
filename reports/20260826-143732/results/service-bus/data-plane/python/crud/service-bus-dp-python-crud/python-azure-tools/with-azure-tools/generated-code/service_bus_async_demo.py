"""Asynchronous Azure Service Bus queue and pub/sub demonstration."""

import asyncio
import os

from azure.identity.aio import DefaultAzureCredential
from azure.servicebus import (
    ServiceBusMessage,
    ServiceBusMessageBatch,
    ServiceBusReceivedMessage,
)
from azure.servicebus.aio import (
    ServiceBusClient,
    ServiceBusReceiver,
    ServiceBusSender,
)

BATCH_SIZE = 5
MAX_WAIT_TIME_SECONDS = 10


def required_setting(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Set the required environment variable {name}.")
    return value


async def process_message(
    source: str, message: ServiceBusReceivedMessage
) -> None:
    """Replace this coroutine with application-specific asynchronous work."""
    await asyncio.sleep(0)
    print(f"{source}: {str(message)}")


async def create_five_message_batch(
    sender: ServiceBusSender,
) -> ServiceBusMessageBatch:
    batch = await sender.create_message_batch()
    for index in range(1, BATCH_SIZE + 1):
        batch.add_message(
            ServiceBusMessage(
                f"Async queue batch message {index}",
                message_id=f"async-queue-{index}",
            )
        )
    return batch


async def queue_round_trip(
    client: ServiceBusClient, queue_name: str
) -> None:
    sender = client.get_queue_sender(queue_name=queue_name)
    async with sender:
        batch = await create_five_message_batch(sender)
        await sender.send_messages(batch)
        print(f"Sent {BATCH_SIZE} messages to queue {queue_name}.")

    receiver = client.get_queue_receiver(
        queue_name=queue_name,
        prefetch_count=BATCH_SIZE,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=BATCH_SIZE,
            max_wait_time=MAX_WAIT_TIME_SECONDS,
        )
        await process_and_complete(receiver, messages, "queue")


async def process_and_complete(
    receiver: ServiceBusReceiver,
    messages: list[ServiceBusReceivedMessage],
    source: str,
) -> None:
    for message in messages:
        await process_message(source, message)
        await receiver.complete_message(message)


async def topic_round_trip(
    client: ServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    sender = client.get_topic_sender(topic_name=topic_name)
    async with sender:
        await sender.send_messages(
            ServiceBusMessage(
                "Hello from the asynchronous topic publisher",
                message_id="async-topic-1",
            )
        )
        print(f"Sent a message to topic {topic_name}.")

    receiver = client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=MAX_WAIT_TIME_SECONDS,
        )
        await process_and_complete(receiver, messages, "subscription")


async def main() -> None:
    namespace = required_setting("SERVICEBUS_FULLY_QUALIFIED_NAMESPACE")
    queue_name = required_setting("SERVICEBUS_QUEUE_NAME")
    topic_name = required_setting("SERVICEBUS_TOPIC_NAME")
    subscription_name = required_setting("SERVICEBUS_SUBSCRIPTION_NAME")

    async with DefaultAzureCredential() as credential:
        async with ServiceBusClient(
            fully_qualified_namespace=namespace,
            credential=credential,
        ) as client:
            # Independent queue and pub/sub flows overlap for higher throughput.
            await asyncio.gather(
                queue_round_trip(client, queue_name),
                topic_round_trip(client, topic_name, subscription_name),
            )


if __name__ == "__main__":
    asyncio.run(main())
