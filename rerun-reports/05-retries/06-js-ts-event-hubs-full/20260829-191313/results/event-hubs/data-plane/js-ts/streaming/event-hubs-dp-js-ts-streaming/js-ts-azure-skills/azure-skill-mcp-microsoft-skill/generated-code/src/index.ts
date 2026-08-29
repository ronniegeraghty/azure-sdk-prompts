import {
  EventHubConsumerClient,
  EventHubProducerClient,
  earliestEventPosition,
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

function waitForShutdownSignal(): {
  promise: Promise<void>;
  dispose: () => void;
} {
  let resolveShutdown: () => void;
  const promise = new Promise<void>((resolve) => {
    resolveShutdown = resolve;
  });

  const handleSignal = (signal: NodeJS.Signals): void => {
    console.log(`\nReceived ${signal}; shutting down...`);
    resolveShutdown();
  };

  process.once("SIGINT", handleSignal);
  process.once("SIGTERM", handleSignal);

  return {
    promise,
    dispose: () => {
      process.off("SIGINT", handleSignal);
      process.off("SIGTERM", handleSignal);
    },
  };
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
    "EVENTHUB_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENTHUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const consumerGroup = process.env.EVENTHUB_CONSUMER_GROUP ?? "$Default";
  const containerName =
    process.env.BLOB_CONTAINER_NAME ?? "event-hub-checkpoints";

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;
  const shutdown = waitForShutdownSignal();

  try {
    const batch = await producer.createBatch();

    for (let index = 1; index <= 10; index += 1) {
      const wasAdded = batch.tryAdd({
        body: {
          message: `Event ${index}`,
          createdAt: new Date().toISOString(),
        },
        contentType: "application/json",
        properties: {
          eventType: "demo",
          eventNumber: index,
          source: "typescript-sample",
        },
      });

      if (!wasAdded) {
        throw new Error(`Event ${index} did not fit in the Event Hubs batch.`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events.`);

    const containerClient = new ContainerClient(
      storageConnectionString,
      containerName,
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
              `Partition ${context.partitionId}:`,
              JSON.stringify(event.body),
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
            `Error while processing partition ${context.partitionId}:`,
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
    await shutdown.promise;
  } finally {
    shutdown.dispose();
    await subscription?.close();
    await consumer?.close();
    await producer.close();
    console.log("Event Hubs clients closed.");
  }
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;
});
