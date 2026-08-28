import os

from azure.identity import DefaultAzureCredential
from azure.servicebus import (
    ServiceBusClient,
    ServiceBusMessage,
    ServiceBusReceivedMessage,
)


MESSAGE_COUNT = 5
RECEIVE_WAIT_SECONDS = 5


def required_setting(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Set the {name} environment variable.")
    return value


def process_message(source: str, message: ServiceBusReceivedMessage) -> None:
    print(f"{source}: {str(message)}")


def send_queue_batch(
    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT
) -> None:
    with client.get_queue_sender(queue_name=queue_name) as sender:
        batch = sender.create_message_batch()
        for index in range(1, message_count + 1):
            try:
                batch.add_message(
                    ServiceBusMessage(
                        f"Queue message {index}",
                        message_id=f"sync-queue-{index}",
                        content_type="text/plain",
                    )
                )
            except ValueError as error:
                raise RuntimeError(
                    "The five-message demonstration batch exceeds the entity size limit."
                ) from error

        sender.send_messages(batch)
        print(f"Sent {message_count} queue messages in one batch.")


def receive_queue_messages(
    client: ServiceBusClient, queue_name: str, message_count: int = MESSAGE_COUNT
) -> None:
    with client.get_queue_receiver(queue_name=queue_name) as receiver:
        messages = receiver.receive_messages(
            max_message_count=message_count,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        for message in messages:
            process_message("queue", message)
            receiver.complete_message(message)

        print(f"Completed {len(messages)} queue messages.")


def send_topic_message(client: ServiceBusClient, topic_name: str) -> None:
    with client.get_topic_sender(topic_name=topic_name) as sender:
        sender.send_messages(
            ServiceBusMessage(
                "Topic message",
                message_id="sync-topic-1",
                content_type="text/plain",
            )
        )
        print("Sent one topic message.")


def receive_subscription_message(
    client: ServiceBusClient, topic_name: str, subscription_name: str
) -> None:
    with client.get_subscription_receiver(
        topic_name=topic_name,
        subscription_name=subscription_name,
    ) as receiver:
        messages = receiver.receive_messages(
            max_message_count=1,
            max_wait_time=RECEIVE_WAIT_SECONDS,
        )
        for message in messages:
            process_message("subscription", message)
            receiver.complete_message(message)

        print(f"Completed {len(messages)} subscription messages.")


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
            send_queue_batch(client, queue_name)
            receive_queue_messages(client, queue_name)
            send_topic_message(client, topic_name)
            receive_subscription_message(client, topic_name, subscription_name)


if __name__ == "__main__":
    main()
