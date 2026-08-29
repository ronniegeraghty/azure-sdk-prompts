import {
  EventHubConsumerClient,
  EventHubProducerClient,
  earliestEventPosition,
  type Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { BlobServiceClient } from "@azure/storage-blob";
import { randomUUID } from "node:crypto";

const eventHubsConnectionString = requireEnvironmentVariable(
  "EVENT_HUB_CONNECTION_STRING",
);
const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
const storageConnectionString = requireEnvironmentVariable(
  "AZURE_STORAGE_CONNECTION_STRING",
);
const checkpointContainerName = requireEnvironmentVariable(
  "CHECKPOINT_CONTAINER_NAME",
);
const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

async function main(): Promise<void> {
  const producer = new EventHubProducerClient(
    eventHubsConnectionString,
    eventHubName,
  );
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;
  let resolveStop: (() => void) | undefined;
  const stopRequested = new Promise<void>((resolve) => {
    resolveStop = resolve;
  });
  const requestStop = (signal: NodeJS.Signals): void => {
    console.log(`Received ${signal}; shutting down.`);
    resolveStop?.();
  };

  process.once("SIGINT", requestStop);
  process.once("SIGTERM", requestStop);

  try {
    const runId = randomUUID();
    const batch = await producer.createBatch();

    for (let index = 1; index <= 10; index += 1) {
      const added = batch.tryAdd({
        body: {
          message: `Event ${index}`,
          createdAt: new Date().toISOString(),
        },
        properties: {
          runId,
          eventNumber: index,
          source: "typescript-event-hubs-demo",
        },
      });

      if (!added) {
        throw new Error(`Event ${index} is too large for the Event Hubs batch.`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events for run ${runId}.`);

    const blobServiceClient = BlobServiceClient.fromConnectionString(
      storageConnectionString,
    );
    const containerClient =
      blobServiceClient.getContainerClient(checkpointContainerName);
    const checkpointStore = new BlobCheckpointStore(containerClient);

    consumer = new EventHubConsumerClient(
      consumerGroup,
      eventHubsConnectionString,
      eventHubName,
      checkpointStore,
    );

    const receivedEventNumbers = new Set<number>();
    subscription = consumer.subscribe(
      {
        processEvents: async (events, context) => {
          for (const event of events) {
            console.log(
              `Partition ${context.partitionId} received:`,
              event.body,
            );

            if (
              event.properties?.runId === runId &&
              typeof event.properties.eventNumber === "number"
            ) {
              receivedEventNumbers.add(event.properties.eventNumber);
            }

            await context.updateCheckpoint(event);
          }

          if (receivedEventNumbers.size === 10) {
            console.log("Received and checkpointed all 10 demo events.");
            resolveStop?.();
          }
        },
        processError: async (error, context) => {
          console.error(
            `Error on partition ${context.partitionId ?? "unknown"}:`,
            error,
          );
        },
      },
      {
        startPosition: earliestEventPosition,
      },
    );

    console.log("Listening for events. Press Ctrl+C to stop.");
    await stopRequested;
  } finally {
    process.removeListener("SIGINT", requestStop);
    process.removeListener("SIGTERM", requestStop);
    await subscription?.close();
    await consumer?.close();
    await producer.close();
  }
}

main().catch((error: unknown) => {
  console.error("Event Hubs example failed:", error);
  process.exitCode = 1;
});
