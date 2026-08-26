"""Synchronous Azure Service Bus queue and topic/subscription demonstration."""

import os

from azure.identity import DefaultAzureCredential
from azure.servicebus import (
    ServiceBusClient,
    ServiceBusMessage,
    ServiceBusMessageBatch,
    ServiceBusReceiver,
    ServiceBusReceivedMessage,
    ServiceBusSender,
)

BATCH_SIZE = 5
MAX_WAIT_TIME_SECONDS = 10


def required_setting(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Set the required environment variable {name}.")
    return value


def process_message(source: str, message: ServiceBusReceivedMessage) -> None:
    """Replace this function with application-specific processing."""
    print(f"{source}: {str(message)}")


def create_five_message_batch(sender: ServiceBusSender) -> ServiceBusMessageBatch:
    batch = sender.create_message_batch()
    for index in range(1, BATCH_SIZE + 1):
        batch.add_message(
            ServiceBusMessage(
                f"Queue batch message {index}",
                message_id=f"sync-queue-{index}",
            )
        )
    return batch


def receive_up_to_five(
    receiver: ServiceBusReceiver,
) -> list[ServiceBusReceivedMessage]:
    return receiver.receive_messages(
        max_message_count=BATCH_SIZE,
        max_wait_time=MAX_WAIT_TIME_SECONDS,
    )


def main() -> None:
    namespace = required_setting("SERVICEBUS_FULLY_QUALIFIED_NAMESPACE")
    queue_name = required_setting("SERVICEBUS_QUEUE_NAME")
    topic_name = required_setting("SERVICEBUS_TOPIC_NAME")
    subscription_name = required_setting("SERVICEBUS_SUBSCRIPTION_NAME")

    with DefaultAzureCredential() as credential:
        with ServiceBusClient(
            fully_qualified_namespace=namespace,
            credential=credential,
        ) as client:
            with client.get_queue_sender(queue_name=queue_name) as sender:
                batch = create_five_message_batch(sender)
                sender.send_messages(batch)
                print(f"Sent {BATCH_SIZE} messages to queue {queue_name}.")

            with client.get_queue_receiver(queue_name=queue_name) as receiver:
                for message in receive_up_to_five(receiver):
                    process_message("queue", message)
                    receiver.complete_message(message)

            with client.get_topic_sender(topic_name=topic_name) as sender:
                sender.send_messages(
                    ServiceBusMessage(
                        "Hello from the synchronous topic publisher",
                        message_id="sync-topic-1",
                    )
                )
                print(f"Sent a message to topic {topic_name}.")

            with client.get_subscription_receiver(
                topic_name=topic_name,
                subscription_name=subscription_name,
            ) as receiver:
                messages = receiver.receive_messages(
                    max_message_count=1,
                    max_wait_time=MAX_WAIT_TIME_SECONDS,
                )
                for message in messages:
                    process_message("subscription", message)
                    receiver.complete_message(message)


if __name__ == "__main__":
    main()
