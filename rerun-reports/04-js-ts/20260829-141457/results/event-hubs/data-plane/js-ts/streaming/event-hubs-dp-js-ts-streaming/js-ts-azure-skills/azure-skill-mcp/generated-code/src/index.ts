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
    process.once("SIGINT", resolve);
    process.once("SIGTERM", resolve);
  });
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const blobContainerName = requireEnvironmentVariable("BLOB_CONTAINER_NAME");
  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  const containerClient = new ContainerClient(
    storageConnectionString,
    blobContainerName,
  );
  const checkpointStore = new BlobCheckpointStore(containerClient);
  const consumer = new EventHubConsumerClient(
    consumerGroup,
    eventHubConnectionString,
    eventHubName,
    checkpointStore,
  );

  let subscription: Subscription | undefined;

  try {
    const batch = await producer.createBatch();

    for (let sequence = 1; sequence <= 10; sequence += 1) {
      const added = batch.tryAdd({
        body: {
          sequence,
          message: `Event ${sequence}`,
          sentAt: new Date().toISOString(),
        },
        properties: {
          source: "typescript-sample",
          sequence,
          category: sequence % 2 === 0 ? "even" : "odd",
        },
      });

      if (!added) {
        throw new Error(`Event ${sequence} did not fit in the batch.`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events.`);

    subscription = consumer.subscribe(
      {
        processEvents: async (events, context) => {
          if (events.length === 0) {
            return;
          }

          for (const event of events) {
            console.log(
              `Partition ${context.partitionId} received:`,
              event.body,
              "properties:",
              event.properties,
            );
          }

          await context.updateCheckpoint(events[events.length - 1]);
          console.log(
            `Checkpoint updated for partition ${context.partitionId}.`,
          );
        },
        processError: async (error, context) => {
          const scope = context.partitionId
            ? `partition ${context.partitionId}`
            : "the consumer";
          console.error(`Error processing ${scope}:`, error);
        },
      },
      {
        startPosition: earliestEventPosition,
        maxBatchSize: 10,
        maxWaitTimeInSeconds: 10,
      },
    );

    console.log("Receiving events. Press Ctrl+C to stop.");
    const signal = await waitForShutdownSignal();
    console.log(`Received ${signal}; shutting down.`);
  } finally {
    await subscription?.close();
    await Promise.all([consumer.close(), producer.close()]);
  }
}

main().catch((error: unknown) => {
  console.error("Fatal error:", error);
  process.exitCode = 1;
});
