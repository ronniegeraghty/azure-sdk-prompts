import {
  earliestEventPosition,
  EventHubConsumerClient,
  EventHubProducerClient,
  type Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { ContainerClient } from "@azure/storage-blob";
import { randomUUID } from "node:crypto";

const EVENT_COUNT = 10;

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

async function main(): Promise<void> {
  const eventHubConnectionString = requiredEnvironmentVariable(
    "EVENTHUB_CONNECTION_STRING",
  );
  const eventHubName = requiredEnvironmentVariable("EVENTHUB_NAME");
  const storageConnectionString = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const storageContainerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );
  const consumerGroup = process.env.EVENTHUB_CONSUMER_GROUP ?? "$Default";
  const runId = randomUUID();

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );

  try {
    const batch = await producer.createBatch();

    for (let index = 1; index <= EVENT_COUNT; index += 1) {
      const wasAdded = batch.tryAdd({
        body: {
          message: `Event ${index}`,
          sentAt: new Date().toISOString(),
        },
        contentType: "application/json",
        properties: {
          eventNumber: index,
          eventType: "typescript-demo",
          runId,
        },
      });

      if (!wasAdded) {
        throw new Error(
          `Event ${index} did not fit in the batch; no events were sent.`,
        );
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${EVENT_COUNT} events for run ${runId}.`);
  } finally {
    await producer.close();
  }

  const containerClient = new ContainerClient(
    storageConnectionString,
    storageContainerName,
  );
  await containerClient.createIfNotExists();

  const checkpointStore = new BlobCheckpointStore(containerClient);
  const consumer = new EventHubConsumerClient(
    consumerGroup,
    eventHubConnectionString,
    eventHubName,
    checkpointStore,
  );

  let subscription: Subscription | undefined;
  let shutdownStarted = false;
  let receivedForRun = 0;
  let resolveFinished: (() => void) | undefined;
  const finished = new Promise<void>((resolve) => {
    resolveFinished = resolve;
  });

  const requestShutdown = (signal: NodeJS.Signals): void => {
    if (!shutdownStarted) {
      shutdownStarted = true;
      console.log(`Received ${signal}; shutting down gracefully.`);
      resolveFinished?.();
    }
  };

  process.once("SIGINT", requestShutdown);
  process.once("SIGTERM", requestShutdown);

  try {
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
            );

            if (event.properties?.runId === runId) {
              receivedForRun += 1;
            }
          }

          await context.updateCheckpoint(events[events.length - 1]);
          console.log(
            `Updated checkpoint for partition ${context.partitionId}.`,
          );

          if (receivedForRun >= EVENT_COUNT && !shutdownStarted) {
            shutdownStarted = true;
            resolveFinished?.();
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
        maxBatchSize: EVENT_COUNT,
        maxWaitTimeInSeconds: 5,
      },
    );

    await finished;
  } finally {
    process.off("SIGINT", requestShutdown);
    process.off("SIGTERM", requestShutdown);
    await subscription?.close();
    await consumer.close();
    console.log("Event Hubs clients closed.");
  }
}

main().catch((error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;
});
