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

function waitForShutdownSignal(): Promise<"SIGINT" | "SIGTERM"> {
  return new Promise((resolve) => {
    process.once("SIGINT", () => resolve("SIGINT"));
    process.once("SIGTERM", () => resolve("SIGTERM"));
  });
}

async function closeResources(
  subscription: Subscription | undefined,
  consumer: EventHubConsumerClient | undefined,
  producer: EventHubProducerClient | undefined,
): Promise<void> {
  let closeError: unknown;

  for (const close of [
    () => subscription?.close(),
    () => consumer?.close(),
    () => producer?.close(),
  ]) {
    try {
      await close();
    } catch (error) {
      closeError ??= error;
    }
  }

  if (closeError) {
    throw closeError;
  }
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const checkpointContainerName =
    process.env.CHECKPOINT_CONTAINER_NAME ?? "event-hub-checkpoints";
  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

  let producer: EventHubProducerClient | undefined;
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;

  try {
    producer = new EventHubProducerClient(
      eventHubConnectionString,
      eventHubName,
    );

    const batch = await producer.createBatch();
    for (let index = 1; index <= 10; index += 1) {
      const added = batch.tryAdd({
        body: {
          eventNumber: index,
          message: `Hello from event ${index}`,
          sentAt: new Date().toISOString(),
        },
        contentType: "application/json",
        properties: {
          eventType: "sample",
          source: "typescript-demo",
          eventNumber: index,
        },
      });

      if (!added) {
        throw new Error(`Event ${index} did not fit in the Event Hubs batch.`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events.`);

    const containerClient = ContainerClient.fromConnectionString(
      storageConnectionString,
      checkpointContainerName,
    );
    await containerClient.createIfNotExists();

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
          if (events.length === 0) {
            return;
          }

          for (const event of events) {
            console.log(
              `Partition ${context.partitionId}, sequence ${event.sequenceNumber}:`,
              event.body,
              event.properties,
            );
          }

          const lastEvent = events.at(-1);
          if (lastEvent) {
            await context.updateCheckpoint(lastEvent);
            console.log(
              `Updated partition ${context.partitionId} checkpoint to sequence ${lastEvent.sequenceNumber}.`,
            );
          }
        },
        processError: async (error, context) => {
          console.error(
            `Event Hubs error on partition ${context.partitionId}:`,
            error,
          );
        },
      },
      {
        startPosition: earliestEventPosition,
        maxBatchSize: 100,
        maxWaitTimeInSeconds: 10,
      },
    );

    console.log("Receiving events. Press Ctrl+C to stop.");
    const signal = await waitForShutdownSignal();
    console.log(`Received ${signal}; shutting down gracefully.`);
  } finally {
    await closeResources(subscription, consumer, producer);
  }
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;
});
