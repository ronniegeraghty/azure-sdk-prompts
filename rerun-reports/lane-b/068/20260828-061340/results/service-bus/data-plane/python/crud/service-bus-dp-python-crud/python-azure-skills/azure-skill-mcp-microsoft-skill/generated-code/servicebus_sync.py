"""Synchronous Azure Service Bus queue and topic/subscription examples."""

import os

from azure.identity import DefaultAzureCredential
from azure.servicebus import (
    ServiceBusClient,
    ServiceBusMessage,
    ServiceBusMessageBatch,
)


def process_message(source: str, message: ServiceBusMessage) -> None:
    """Replace this function with application-specific processing."""
    print(f"{source}: {message}")


def send_queue_batch(client: ServiceBusClient, queue_name: str) -> None:
    with client.get_queue_sender(queue_name=queue_name) as sender:
        batch: ServiceBusMessageBatch = sender.create_message_batch()
        for index in range(1, 6):
            batch.add_message(
                ServiceBusMessage(
                    f"Synchronous queue message {index}",
                    message_id=f"sync-queue-{index}",
                )
            )
        sender.send_messages(batch)
        print(f"Sent {len(batch)} messages to queue {queue_name!r}.")


def receive_queue_messages(client: ServiceBusClient, queue_name: str) -> None:
    with client.get_queue_receiver(
        queue_name=queue_name,
        max_wait_time=5,
    ) as receiver:
        messages = receiver.receive_messages(
            max_message_count=5,
            max_wait_time=5,
        )
        for message in messages:
            process_message("queue", message)
            receiver.complete_message(message)


def topic_subscription_roundtrip(
    client: ServiceBusClient,
    topic_name: str,
    subscription_name: str,
) -> None:
    with client.get_topic_sender(topic_name=topic_name) as sender:
        sender.send_messages(
            ServiceBusMessage(
                "Synchronous topic message",
                message_id="sync-topic-1",
            )
        )

    with client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
        max_wait_time=5,
    ) as receiver:
        messages = receiver.receive_messages(
            max_message_count=1,
            max_wait_time=5,
        )
        for message in messages:
            process_message("topic subscription", message)
            receiver.complete_message(message)


def main() -> None:
    namespace = os.environ["SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.environ["SERVICEBUS_QUEUE_NAME"]
    topic_name = os.environ["SERVICEBUS_TOPIC_NAME"]
    subscription_name = os.environ["SERVICEBUS_SUBSCRIPTION_NAME"]

    with DefaultAzureCredential() as credential:
        with ServiceBusClient(
            fully_qualified_namespace=namespace,
            credential=credential,
        ) as client:
            send_queue_batch(client, queue_name)
            receive_queue_messages(client, queue_name)
            topic_subscription_roundtrip(client, topic_name, subscription_name)


if __name__ == "__main__":
    main()
