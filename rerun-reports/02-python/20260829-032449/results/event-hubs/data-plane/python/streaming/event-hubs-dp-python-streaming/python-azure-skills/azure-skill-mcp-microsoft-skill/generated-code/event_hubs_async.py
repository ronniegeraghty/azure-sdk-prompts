#!/usr/bin/env python3
"""Send a batch to Azure Event Hubs, then receive and checkpoint events.

Required environment variables:
    EVENT_HUB_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net
    EVENT_HUB_NAME=<event-hub-name>
    STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net

Optional environment variables:
    EVENT_HUB_CONSUMER_GROUP=$Default
    CHECKPOINT_CONTAINER=event-hub-checkpoints

The Blob Storage container must already exist. DefaultAzureCredential uses
developer credentials locally and managed identity when hosted in Azure.
"""

import asyncio
import json
import logging
import os
from dataclasses import dataclass
from datetime import datetime, timezone

from azure.eventhub import EventData
from azure.eventhub.aio import EventHubConsumerClient, EventHubProducerClient
from azure.eventhub.extensions.checkpointstoreblobaio import BlobCheckpointStore
from azure.identity.aio import DefaultAzureCredential


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s %(levelname)s %(name)s: %(message)s",
)
LOGGER = logging.getLogger("event-hubs-demo")


@dataclass(frozen=True)
class Settings:
    fully_qualified_namespace: str
    event_hub_name: str
    storage_account_url: str
    consumer_group: str
    checkpoint_container: str

    @classmethod
    def from_environment(cls) -> "Settings":
        required = {
            "EVENT_HUB_FULLY_QUALIFIED_NAMESPACE": os.getenv(
                "EVENT_HUB_FULLY_QUALIFIED_NAMESPACE"
            ),
            "EVENT_HUB_NAME": os.getenv("EVENT_HUB_NAME"),
            "STORAGE_ACCOUNT_URL": os.getenv("STORAGE_ACCOUNT_URL"),
        }
        missing = [name for name, value in required.items() if not value]
        if missing:
            names = ", ".join(missing)
            raise RuntimeError(f"Missing required environment variables: {names}")

        return cls(
            fully_qualified_namespace=required[
                "EVENT_HUB_FULLY_QUALIFIED_NAMESPACE"
            ],
            event_hub_name=required["EVENT_HUB_NAME"],
            storage_account_url=required["STORAGE_ACCOUNT_URL"],
            consumer_group=os.getenv("EVENT_HUB_CONSUMER_GROUP", "$Default"),
            checkpoint_container=os.getenv(
                "CHECKPOINT_CONTAINER", "event-hub-checkpoints"
            ),
        )


async def send_batch(
    settings: Settings,
    credential: DefaultAzureCredential,
) -> None:
    async with EventHubProducerClient(
        fully_qualified_namespace=settings.fully_qualified_namespace,
        eventhub_name=settings.event_hub_name,
        credential=credential,
    ) as producer:
        batch = await producer.create_batch()
        for event_number in range(1, 6):
            payload = {
                "event_number": event_number,
                "message": f"Async Event Hubs sample event {event_number}",
                "sent_at": datetime.now(timezone.utc).isoformat(),
            }
            batch.add(EventData(json.dumps(payload)))

        await producer.send_batch(batch)
        LOGGER.info("Sent %d events in one batch", len(batch))


async def receive_events(
    settings: Settings,
    credential: DefaultAzureCredential,
) -> None:
    checkpoint_store = BlobCheckpointStore(
        blob_account_url=settings.storage_account_url,
        container_name=settings.checkpoint_container,
        credential=credential,
    )

    async def on_event(partition_context, event) -> None:
        print(
            f"Partition {partition_context.partition_id}: "
            f"{event.body_as_str(encoding='UTF-8')}"
        )
        await partition_context.update_checkpoint(event)
        LOGGER.info(
            "Updated checkpoint for partition %s at sequence number %s",
            partition_context.partition_id,
            event.sequence_number,
        )

    async def on_error(partition_context, error: Exception) -> None:
        if partition_context is None:
            LOGGER.error("Consumer error: %s", error)
        else:
            LOGGER.error(
                "Consumer error on partition %s: %s",
                partition_context.partition_id,
                error,
            )

    async with checkpoint_store:
        async with EventHubConsumerClient(
            fully_qualified_namespace=settings.fully_qualified_namespace,
            eventhub_name=settings.event_hub_name,
            consumer_group=settings.consumer_group,
            credential=credential,
            checkpoint_store=checkpoint_store,
        ) as consumer:
            LOGGER.info("Receiving events; press Ctrl+C to stop")
            await consumer.receive(
                on_event=on_event,
                on_error=on_error,
                starting_position="-1",
            )


async def main() -> None:
    settings = Settings.from_environment()
    async with DefaultAzureCredential() as credential:
        await send_batch(settings, credential)
        await receive_events(settings, credential)


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        LOGGER.info("Receiver stopped")
