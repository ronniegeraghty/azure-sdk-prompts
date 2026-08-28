import asyncio
import os
from typing import Optional

from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
from azure.eventhub import EventData
from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore


EVENT_HUB_CONNECTION_STRING = os.environ["EVENT_HUB_CONNECTION_STRING"]
EVENT_HUB_NAME = os.environ["EVENT_HUB_NAME"]
STORAGE_CONNECTION_STRING = os.environ["AZURE_STORAGE_CONNECTION_STRING"]
BLOB_CONTAINER_NAME = os.environ["BLOB_CHECKPOINT_CONTAINER"]
CONSUMER_GROUP = os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default")


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


async def on_event(partition_context, event) -> None:
    print(
        f"Partition {partition_context.partition_id}: "
        f"{event.body_as_str(encoding='UTF-8')}"
    )
    await partition_context.update_checkpoint(event)


async def on_error(partition_context, error: Exception) -> None:
    partition_id: Optional[str] = (
        partition_context.partition_id if partition_context else None
    )
    if partition_id is None:
        print(f"Consumer error: {error!r}")
    else:
        print(f"Consumer error on partition {partition_id}: {error!r}")


async def receive_events() -> None:
    checkpoint_store = BlobCheckpointStore.from_connection_string(
        STORAGE_CONNECTION_STRING,
        BLOB_CONTAINER_NAME,
    )
    consumer = EventHubConsumerClient.from_connection_string(
        conn_str=EVENT_HUB_CONNECTION_STRING,
        consumer_group=CONSUMER_GROUP,
        eventhub_name=EVENT_HUB_NAME,
        checkpoint_store=checkpoint_store,
    )

    async with consumer:
        print("Receiving events. Press Ctrl+C to stop.")
        await consumer.receive(
            on_event=on_event,
            on_error=on_error,
            starting_position="-1",
        )


async def main() -> None:
    await send_events()
    await receive_events()


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        print("Receiver stopped.")
