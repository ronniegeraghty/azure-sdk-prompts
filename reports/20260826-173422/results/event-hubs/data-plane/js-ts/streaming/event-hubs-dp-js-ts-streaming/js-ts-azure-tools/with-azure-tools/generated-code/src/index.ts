import {
  earliestEventPosition,
  EventHubConsumerClient,
  EventHubProducerClient,
  type Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { ContainerClient } from "@azure/storage-blob";

const eventCount = 10;

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
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

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;
  let closing = false;

  const close = async (): Promise<void> => {
    if (closing) {
      return;
    }
    closing = true;

    await subscription?.close();
    await consumer?.close();
    await producer.close();
  };

  try {
    const runId = crypto.randomUUID();
    const batch = await producer.createBatch();

    for (let index = 0; index < eventCount; index += 1) {
      const added = batch.tryAdd({
        body: {
          message: `Event ${index + 1}`,
          sentAt: new Date().toISOString(),
        },
        contentType: "application/json",
        properties: {
          eventIndex: index,
          eventType: "sample",
          runId,
        },
      });

      if (!added) {
        throw new Error(`Event ${index + 1} did not fit in the Event Hubs batch`);
      }
    }

    await producer.sendBatch(batch);
    console.log(`Sent ${eventCount} events in one batch (runId: ${runId}).`);

    const containerClient = ContainerClient.fromConnectionString(
      storageConnectionString,
      checkpointContainerName,
    );
    const checkpointStore = new BlobCheckpointStore(containerClient);

    consumer = new EventHubConsumerClient(
      consumerGroup,
      eventHubConnectionString,
      eventHubName,
      checkpointStore,
    );

    let receivedFromThisRun = 0;
    let resolveAllEventsReceived: (() => void) | undefined;
    const allEventsReceived = new Promise<void>((resolve) => {
      resolveAllEventsReceived = resolve;
    });

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
              receivedFromThisRun += 1;
            }
          }

          await context.updateCheckpoint(events[events.length - 1]!);
          console.log(
            `Checkpoint updated for partition ${context.partitionId}.`,
          );

          if (receivedFromThisRun >= eventCount) {
            resolveAllEventsReceived?.();
          }
        },
        processError: async (error, context) => {
          console.error(
            `Error while receiving from partition ${context.partitionId}:`,
            error,
          );
        },
      },
      {
        maxBatchSize: eventCount,
        maxWaitTimeInSeconds: 5,
        startPosition: earliestEventPosition,
      },
    );

    console.log("Listening for events. Press Ctrl+C to stop.");

    const shutdownSignal = new Promise<NodeJS.Signals>((resolve) => {
      process.once("SIGINT", () => resolve("SIGINT"));
      process.once("SIGTERM", () => resolve("SIGTERM"));
    });

    const outcome = await Promise.race([
      allEventsReceived.then(() => "received" as const),
      shutdownSignal,
    ]);

    if (outcome === "received") {
      console.log(`Received all ${eventCount} events from this run.`);
    } else {
      console.log(`Received ${outcome}; shutting down gracefully.`);
    }
  } finally {
    await close();
    console.log("Producer, subscription, and consumer closed.");
  }
}

main().catch((error: unknown) => {
  console.error("Event Hubs sample failed:", error);
  process.exitCode = 1;
});
