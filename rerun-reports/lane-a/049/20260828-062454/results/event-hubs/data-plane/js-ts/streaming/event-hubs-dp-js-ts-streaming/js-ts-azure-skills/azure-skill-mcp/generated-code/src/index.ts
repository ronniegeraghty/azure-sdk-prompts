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
  const consumerGroup =
    process.env.EVENT_HUB_CONSUMER_GROUP ?? EventHubConsumerClient.defaultConsumerGroupName;
  const storageConnectionString = requireEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const blobContainerName = requireEnvironmentVariable("BLOB_CONTAINER_NAME");

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  let consumer: EventHubConsumerClient | undefined;
  let subscription: Subscription | undefined;

  try {
    const batch = await producer.createBatch();

    for (let sequence = 1; sequence <= 10; sequence += 1) {
      const wasAdded = batch.tryAdd({
        body: {
          message: `Event ${sequence}`,
          createdAt: new Date().toISOString(),
        },
        properties: {
          sequence,
          source: "typescript-event-hubs-sample",
          category: sequence % 2 === 0 ? "even" : "odd",
        },
      });

      if (!wasAdded) {
        throw new Error(`Event ${sequence} did not fit in the Event Hubs batch.`);
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
            );
            await context.updateCheckpoint(event);
          }
        },
        processError: async (error, context) => {
          const partition = context.partitionId ?? "unassigned";
          console.error(`Error processing partition ${partition}:`, error);
        },
      },
      {
        startPosition: earliestEventPosition,
      },
    );

    console.log("Listening for events. Press Ctrl+C to stop.");
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
