import asyncio
import os
from typing import Optional

from azure.eventhub import EventData
from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
from azure.eventhub.aio import PartitionContext


EVENT_HUB_CONNECTION_STRING = os.environ["EVENT_HUB_CONNECTION_STRING"]
EVENT_HUB_NAME = os.environ["EVENT_HUB_NAME"]
EVENT_HUB_CONSUMER_GROUP = os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default")
STORAGE_CONNECTION_STRING = os.environ["AZURE_STORAGE_CONNECTION_STRING"]
BLOB_CONTAINER_NAME = os.environ["BLOB_CONTAINER_NAME"]


async def send_events() -> None:
    producer = EventHubProducerClient.from_connection_string(
        conn_str=EVENT_HUB_CONNECTION_STRING,
        eventhub_name=EVENT_HUB_NAME,
    )

    async with producer:
        batch = await producer.create_batch()
        for index in range(10):
            batch.add(EventData(f"Event {index}"))

        await producer.send_batch(batch)
        print(f"Sent {len(batch)} events")


async def on_event(
    partition_context: PartitionContext,
    event: EventData,
) -> None:
    print(
        f"Partition {partition_context.partition_id}: "
        f"{event.body_as_str(encoding='UTF-8')}"
    )
    await partition_context.update_checkpoint(event)


async def on_error(
    partition_context: Optional[PartitionContext],
    error: Exception,
) -> None:
    if partition_context is None:
        print(f"Consumer error: {error!r}")
    else:
        print(
            f"Partition {partition_context.partition_id} error: {error!r}"
        )


async def receive_events() -> None:
    checkpoint_store = BlobCheckpointStore.from_connection_string(
        STORAGE_CONNECTION_STRING,
        BLOB_CONTAINER_NAME,
    )
    consumer = EventHubConsumerClient.from_connection_string(
        conn_str=EVENT_HUB_CONNECTION_STRING,
        consumer_group=EVENT_HUB_CONSUMER_GROUP,
        eventhub_name=EVENT_HUB_NAME,
        checkpoint_store=checkpoint_store,
    )

    async with consumer:
        await consumer.receive(
            on_event=on_event,
            on_error=on_error,
            starting_position="-1",
        )


async def main() -> None:
    await send_events()
    print("Receiving events. Press Ctrl+C to stop.")
    await receive_events()


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        print("Stopped receiving events")
