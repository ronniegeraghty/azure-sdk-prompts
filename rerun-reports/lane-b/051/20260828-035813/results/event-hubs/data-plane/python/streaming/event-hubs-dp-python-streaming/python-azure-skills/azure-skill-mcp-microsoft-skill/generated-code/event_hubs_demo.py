"""Send and receive Azure Event Hubs events with Blob Storage checkpoints."""

import asyncio
import os

from azure.eventhub import EventData
from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
from azure.identity.aio import DefaultAzureCredential


def required_setting(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Set the {name} environment variable.")
    return value


async def send_events(
    namespace: str,
    event_hub_name: str,
    credential: DefaultAzureCredential,
) -> None:
    async with EventHubProducerClient(
        fully_qualified_namespace=namespace,
        eventhub_name=event_hub_name,
        credential=credential,
    ) as producer:
        batch = await producer.create_batch()
        for event_number in range(1, 6):
            batch.add(EventData(f"Sample event {event_number}"))

        await producer.send_batch(batch)
        print(f"Sent {len(batch)} events.")


async def receive_events(
    namespace: str,
    event_hub_name: str,
    consumer_group: str,
    storage_account_url: str,
    checkpoint_container: str,
    credential: DefaultAzureCredential,
) -> None:
    checkpoint_store = BlobCheckpointStore(
        blob_account_url=storage_account_url,
        container_name=checkpoint_container,
        credential=credential,
    )

    async def on_event(partition_context, event) -> None:
        print(
            f"Partition {partition_context.partition_id}: "
            f"{event.body_as_str(encoding='UTF-8')}"
        )
        await partition_context.update_checkpoint(event)

    async def on_error(partition_context, error: Exception) -> None:
        if partition_context is None:
            print(f"Consumer error: {error}")
        else:
            print(
                f"Consumer error on partition "
                f"{partition_context.partition_id}: {error}"
            )

    async with EventHubConsumerClient(
        fully_qualified_namespace=namespace,
        eventhub_name=event_hub_name,
        consumer_group=consumer_group,
        credential=credential,
        checkpoint_store=checkpoint_store,
    ) as consumer:
        print("Receiving events. Press Ctrl+C to stop.")
        await consumer.receive(
            on_event=on_event,
            on_error=on_error,
            starting_position="-1",
        )


async def main() -> None:
    namespace = required_setting("EVENT_HUB_FULLY_QUALIFIED_NAMESPACE")
    event_hub_name = required_setting("EVENT_HUB_NAME")
    storage_account_url = required_setting("STORAGE_ACCOUNT_URL")
    checkpoint_container = required_setting("CHECKPOINT_CONTAINER")
    consumer_group = os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default")

    async with DefaultAzureCredential() as credential:
        await send_events(namespace, event_hub_name, credential)
        await receive_events(
            namespace,
            event_hub_name,
            consumer_group,
            storage_account_url,
            checkpoint_container,
            credential,
        )


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        print("Stopped.")
