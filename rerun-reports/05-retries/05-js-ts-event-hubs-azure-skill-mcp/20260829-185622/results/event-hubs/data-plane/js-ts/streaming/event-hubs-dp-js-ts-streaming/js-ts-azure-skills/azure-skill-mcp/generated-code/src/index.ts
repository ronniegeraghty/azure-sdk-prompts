import {
  earliestEventPosition,
  EventHubConsumerClient,
  EventHubProducerClient,
  type Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { ContainerClient } from "@azure/storage-blob";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

function waitForShutdownSignal(): Promise<NodeJS.Signals> {
  return new Promise((resolve) => {
    process.once("SIGINT", () => resolve("SIGINT"));
    process.once("SIGTERM", () => resolve("SIGTERM"));
  });
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING",
  );
  const blobContainerName = requireEnvironmentVariable("BLOB_CONTAINER_NAME");
  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;

  try {
    const batch = await producer.createBatch();

    for (let eventNumber = 1; eventNumber <= 10; eventNumber += 1) {
      const wasAdded = batch.tryAdd({
        body: {
          message: `Event ${eventNumber}`,
          sentAt: new Date().toISOString(),
        },
        properties: {
          eventNumber,
          source: "typescript-event-hubs-sample",
        },
      });

      if (!wasAdded) {
        throw new Error(`Event ${eventNumber} did not fit in the batch.`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events.`);

    const containerClient = new ContainerClient(
      storageConnectionString,
      blobContainerName,
    );
    const checkpointStore = new BlobCheckpointStore(containerClient);

    consumer = new EventHubConsumerClient(
      consumerGroup,
      eventHubConnectionString,
      eventHubName,
      checkpointStore,
    );

    subscription = consumer.subscribe(
      {
        processEvents: async (events, context) => {
          for (const event of events) {
            console.log(
              `Partition ${context.partitionId}, sequence ${event.sequenceNumber}:`,
              event.body,
              "properties:",
              event.properties,
            );
          }

          const lastEvent = events.at(-1);
          if (lastEvent) {
            await context.updateCheckpoint(lastEvent);
            console.log(
              `Checkpoint updated for partition ${context.partitionId} at sequence ${lastEvent.sequenceNumber}.`,
            );
          }
        },
        processError: async (error, context) => {
          console.error(
            `Error processing partition ${context.partitionId}:`,
            error,
          );
        },
      },
      {
        startPosition: earliestEventPosition,
        maxBatchSize: 10,
        maxWaitTimeInSeconds: 5,
      },
    );

    console.log("Receiving events. Press Ctrl+C to stop.");
    const signal = await waitForShutdownSignal();
    console.log(`Received ${signal}; shutting down.`);
  } finally {
    await subscription?.close();
    await consumer?.close();
    await producer.close();
  }
}

main().catch((error: unknown) => {
  console.error("Event Hubs sample failed:", error);
  process.exitCode = 1;
});
