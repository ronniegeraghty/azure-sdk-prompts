"""Asynchronous Azure Service Bus examples for higher-throughput processing."""

import asyncio
import os

from azure.identity.aio import DefaultAzureCredential
from azure.servicebus import ServiceBusMessage, ServiceBusMessageBatch
from azure.servicebus.aio import ServiceBusClient, ServiceBusReceiver


async def process_message(source: str, message: ServiceBusMessage) -> None:
    """Replace this function with application-specific asynchronous work."""
    await asyncio.sleep(0)
    print(f"{source}: {message}")


async def process_and_complete(
    receiver: ServiceBusReceiver,
    source: str,
    message: ServiceBusMessage,
) -> None:
    await process_message(source, message)
    await receiver.complete_message(message)


async def send_queue_batch(client: ServiceBusClient, queue_name: str) -> None:
    async with client.get_queue_sender(queue_name=queue_name) as sender:
        batch: ServiceBusMessageBatch = await sender.create_message_batch()
        for index in range(1, 6):
            batch.add_message(
                ServiceBusMessage(
                    f"Asynchronous queue message {index}",
                    message_id=f"async-queue-{index}",
                )
            )
        await sender.send_messages(batch)
        print(f"Sent {len(batch)} messages to queue {queue_name!r}.")


async def receive_queue_messages(
    client: ServiceBusClient,
    queue_name: str,
) -> None:
    async with client.get_queue_receiver(
        queue_name=queue_name,
        max_wait_time=5,
        prefetch_count=10,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=5,
            max_wait_time=5,
        )
        await asyncio.gather(
            *(
                process_and_complete(receiver, "queue", message)
                for message in messages
            )
        )


async def topic_subscription_roundtrip(
    client: ServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    async with client.get_topic_sender(topic_name=topic_name) as sender:
        await sender.send_messages(
            ServiceBusMessage(
                "Asynchronous topic message",
                message_id="async-topic-1",
            )
        )

    async with client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
        max_wait_time=5,
    ) as receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=5,
        )
        await asyncio.gather(
            *(
                process_and_complete(
                    receiver,
                    "topic subscription",
                    message,
                )
                for message in messages
            )
        )


async def main() -> None:
    namespace = os.environ["SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.environ["SERVICEBUS_QUEUE_NAME"]
    topic_name = os.environ["SERVICEBUS_TOPIC_NAME"]
    subscription_name = os.environ["SERVICEBUS_SUBSCRIPTION_NAME"]

    async with DefaultAzureCredential() as credential:
        async with ServiceBusClient(
            fully_qualified_namespace=namespace,
            credential=credential,
        ) as client:
            await send_queue_batch(client, queue_name)
            await receive_queue_messages(client, queue_name)
            await topic_subscription_roundtrip(
                client,
                topic_name,
                subscription_name,
            )


if __name__ == "__main__":
    asyncio.run(main())
