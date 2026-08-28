import asyncio
import os

from azure.identity.aio import DefaultAzureCredential
from azure.servicebus import ServiceBusMessage, ServiceBusReceivedMessage
from azure.servicebus.aio import ServiceBusClient, ServiceBusReceiver


MESSAGE_COUNT = 5
RECEIVE_WAIT_SECONDS = 5


def required_setting(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Set the {name} environment variable.")
    return value


async def process_message(source: str, message: ServiceBusReceivedMessage) -> None:
    await asyncio.sleep(0)
    print(f"{source}: {str(message)}")


async def send_queue_batch(
    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT
) -> None:
    sender = client.get_queue_sender(queue_name=queue_name)
    async with sender:
        batch = await sender.create_message_batch()
        for index in range(1, message_count + 1):
            try:
                batch.add_message(
                    ServiceBusMessage(
                        f"Async queue message {index}",
                        message_id=f"async-queue-{index}",
                        content_type="text/plain",
                    )
                )
            except ValueError as error:
                raise RuntimeError(
                    "The five-message demonstration batch exceeds the entity size limit."
                ) from error

        await sender.send_messages(batch)
        print(f"Sent {message_count} queue messages in one async batch.")


async def process_and_complete(
    receiver: ServiceBusReceiver, source: str, message: ServiceBusReceivedMessage
) -> None:
    await process_message(source, message)
    await receiver.complete_message(message)


async def receive_queue_messages(
    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT
) -> None:
    receiver = client.get_queue_receiver(
        queue_name=queue_name,
        prefetch_count=message_count,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=message_count,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        await asyncio.gather(
            *(
                process_and_complete(receiver, "queue", message)
                for message in messages
            )
        )
        print(f"Completed {len(messages)} queue messages concurrently.")


async def send_topic_message(client: ServiceBusClient, topic_name: str) -> None:
    sender = client.get_topic_sender(topic_name=topic_name)
    async with sender:
        await sender.send_messages(
            ServiceBusMessage(
                "Async topic message",
                message_id="async-topic-1",
                content_type="text/plain",
            )
        )
        print("Sent one async topic message.")


async def receive_subscription_message(
    client: ServiceBusClient, topic_name: str, subscription_name: str
) -> None:
    receiver = client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
    )
    async with receiver:
        messages = await receiver.receive_messages(
            max_message_count=1,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        await asyncio.gather(
            *(
                process_and_complete(receiver, "subscription", message)
                for message in messages
            )
        )
        print(f"Completed {len(messages)} subscription messages.")


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
            await send_queue_batch(client, queue_name)
            await receive_queue_messages(client, queue_name)
            await send_topic_message(client, topic_name)
            await receive_subscription_message(
                client,
                topic_name,
                subscription_name,
            )


if __name__ == "__main__":
    asyncio.run(main())
