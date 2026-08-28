import {
  earliestEventPosition,
  EventData,
  EventHubConsumerClient,
  EventHubProducerClient,
  Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { BlobServiceClient } from "@azure/storage-blob";
import { randomUUID } from "node:crypto";

const eventHubConnectionString = requiredEnv("EVENTHUB_CONNECTION_STRING");
const eventHubName = requiredEnv("EVENTHUB_NAME");
const storageConnectionString = requiredEnv("AZURE_STORAGE_CONNECTION_STRING");
const consumerGroup = process.env.EVENTHUB_CONSUMER_GROUP ?? "$Default";
const containerName =
  process.env.BLOB_CONTAINER_NAME ?? "event-hub-checkpoints";

let producer: EventHubProducerClient | undefined;
let consumer: EventHubConsumerClient | undefined;
let subscription: Subscription | undefined;
let closing: Promise<void> | undefined;

function requiredEnv(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function close(): Promise<void> {
  if (!closing) {
    closing = (async () => {
      console.log("Shutting down...");

      await subscription?.close();
      await consumer?.close();
      await producer?.close();

      console.log("All Event Hubs clients are closed.");
    })();
  }

  return closing;
}

function registerShutdownSignals(): void {
  for (const signal of ["SIGINT", "SIGTERM"] as const) {
    process.once(signal, () => {
      void close().catch((error: unknown) => {
        console.error("Graceful shutdown failed:", error);
        process.exitCode = 1;
      });
    });
  }
}

async function main(): Promise<void> {
  registerShutdownSignals();
  const runId = randomUUID();

  producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );

  const batch = await producer.createBatch();
  for (let sequence = 1; sequence <= 10; sequence += 1) {
    const event: EventData = {
      body: {
        message: `Event ${sequence}`,
        createdAt: new Date().toISOString(),
      },
      contentType: "application/json",
      properties: {
        runId,
        sequence,
        source: "typescript-demo",
      },
    };

    if (!batch.tryAdd(event)) {
      throw new Error(`Event ${sequence} did not fit in the Event Hubs batch.`);
    }
  }

  await producer.sendBatch(batch);
  console.log(`Sent ${batch.count} events for run ${runId}.`);

  const blobServiceClient =
    BlobServiceClient.fromConnectionString(storageConnectionString);
  const containerClient = blobServiceClient.getContainerClient(containerName);
  await containerClient.createIfNotExists();

  const checkpointStore = new BlobCheckpointStore(containerClient);
  consumer = new EventHubConsumerClient(
    consumerGroup,
    eventHubConnectionString,
    eventHubName,
    checkpointStore,
  );

  let receivedForThisRun = 0;
  let resolveAllEventsReceived!: () => void;
  const allEventsReceived = new Promise<void>((resolve) => {
    resolveAllEventsReceived = resolve;
  });

  subscription = consumer.subscribe(
    {
      processEvents: async (events, context) => {
        for (const event of events) {
          console.log(
            `Partition ${context.partitionId}:`,
            JSON.stringify(event.body),
            event.properties,
          );

          if (event.properties?.runId === runId) {
            receivedForThisRun += 1;
          }
        }

        if (events.length > 0) {
          await context.updateCheckpoint(events[events.length - 1]);
          console.log(
            `Updated checkpoint for partition ${context.partitionId}.`,
          );
        }

        if (receivedForThisRun >= 10) {
          resolveAllEventsReceived();
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
      maxWaitTimeInSeconds: 5,
    },
  );

  console.log("Listening for events. Press Ctrl+C to stop.");
  await allEventsReceived;
  console.log("Received all 10 events sent by this run.");
}

void main()
  .catch((error: unknown) => {
    console.error("Application failed:", error);
    process.exitCode = 1;
  })
  .finally(close)
  .catch((error: unknown) => {
    console.error("Shutdown failed:", error);
    process.exitCode = 1;
  });
