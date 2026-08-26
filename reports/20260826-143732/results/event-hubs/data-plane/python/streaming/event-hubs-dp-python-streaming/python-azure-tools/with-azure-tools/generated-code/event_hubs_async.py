"""Send and receive Azure Event Hubs events with async checkpointing."""

import asyncio
import logging
import os
import sys

from azure.core.exceptions import AzureError
from azure.eventhub import EventData
from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
from azure.identity.aio import DefaultAzureCredential

LOGGER = logging.getLogger("event_hubs_async")


def required_setting(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise ValueError(f"Set the required environment variable {name}.")
    return value


async def send_events(
    credential: DefaultAzureCredential,
    namespace: str,
    event_hub_name: str,
) -> None:
    async with EventHubProducerClient(
        fully_qualified_namespace=namespace,
        eventhub_name=event_hub_name,
        credential=credential,
    ) as producer:
        batch = await producer.create_batch()
        for event_number in range(1, 6):
            batch.add(EventData(f"Async event {event_number}"))

        await producer.send_batch(batch)
        LOGGER.info("Sent %d events.", len(batch))


async def receive_events(
    credential: DefaultAzureCredential,
    namespace: str,
    event_hub_name: str,
    consumer_group: str,
    storage_account_url: str,
    checkpoint_container: str,
) -> None:
    async def on_event(partition_context, event) -> None:
        print(
            f"Partition {partition_context.partition_id}: "
            f"{event.body_as_str(encoding='UTF-8')}"
        )
        await partition_context.update_checkpoint(event)

    async def on_error(partition_context, error: Exception) -> None:
        partition_id = (
            partition_context.partition_id if partition_context else "all partitions"
        )
        LOGGER.error("Receive error for %s: %s", partition_id, error)

    async with BlobCheckpointStore(
        blob_account_url=storage_account_url,
        container_name=checkpoint_container,
        credential=credential,
    ) as checkpoint_store:
        async with EventHubConsumerClient(
            fully_qualified_namespace=namespace,
            eventhub_name=event_hub_name,
            consumer_group=consumer_group,
            credential=credential,
            checkpoint_store=checkpoint_store,
        ) as consumer:
            LOGGER.info("Receiving events. Press Ctrl+C to stop.")
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
        await send_events(credential, namespace, event_hub_name)
        await receive_events(
            credential,
            namespace,
            event_hub_name,
            consumer_group,
            storage_account_url,
            checkpoint_container,
        )


if __name__ == "__main__":
    logging.basicConfig(
        level=os.getenv("LOG_LEVEL", "INFO").upper(),
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        LOGGER.info("Receiver stopped.")
    except (AzureError, ValueError) as error:
        LOGGER.error("%s", error)
        sys.exit(1)
