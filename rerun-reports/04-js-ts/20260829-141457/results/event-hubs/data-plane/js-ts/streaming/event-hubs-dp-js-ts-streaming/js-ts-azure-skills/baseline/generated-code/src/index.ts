import {
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

async function waitForShutdownSignal(): Promise<void> {
  await new Promise<void>((resolve) => {
    const shutdown = (signal: NodeJS.Signals): void => {
      console.log(`Received ${signal}; shutting down...`);
      resolve();
    };

    process.once("SIGINT", shutdown);
    process.once("SIGTERM", shutdown);
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
    const sentAt = new Date();
    const batch = await producer.createBatch();

    for (let index = 1; index <= 10; index += 1) {
      const wasAdded = batch.tryAdd({
        body: {
          id: index,
          message: `Event ${index}`,
          createdAt: new Date().toISOString(),
        },
        properties: {
          source: "typescript-sample",
          eventNumber: index,
          category: index % 2 === 0 ? "even" : "odd",
        },
      });

      if (!wasAdded) {
        throw new Error(`Event ${index} did not fit in the Event Hubs batch`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events`);

    subscription = consumer.subscribe(
      {
        processEvents: async (events, context) => {
          for (const event of events) {
            console.log(
              `Partition ${context.partitionId}, sequence ${event.sequenceNumber}:`,
              event.body,
            );
            await context.updateCheckpoint(event);
          }
        },
        processError: async (error, context) => {
          console.error(
            `Error processing partition ${context.partitionId ?? "unknown"}:`,
            error,
          );
        },
      },
      {
        // On a new checkpoint store, include events sent immediately before subscribing.
        startPosition: { enqueuedOn: sentAt },
      },
    );

    console.log("Listening for events. Press Ctrl+C to stop.");
    await waitForShutdownSignal();
  } finally {
    await subscription?.close();
    await consumer.close();
    await producer.close();
    console.log("Event Hubs clients closed");
  }
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;
});
