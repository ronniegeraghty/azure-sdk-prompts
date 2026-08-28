import {
  EventDataBatch,
  EventHubConsumerClient,
  EventHubProducerClient,
  ReceivedEventData,
  Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

async function addEvents(batch: EventDataBatch): Promise<void> {
  for (let index = 1; index <= 10; index += 1) {
    const added = batch.tryAdd({
      body: {
        id: index,
        message: `Event ${index}`,
        createdAt: new Date().toISOString(),
      },
      properties: {
        eventType: "sample.created",
        sequenceNumber: index,
        source: "typescript-sample",
      },
    });

    if (!added) {
      throw new Error(`Event ${index} did not fit in the Event Hubs batch.`);
    }
  }
}

function waitForShutdownSignal(): Promise<NodeJS.Signals> {
  return new Promise((resolve) => {
    const signals: NodeJS.Signals[] = ["SIGINT", "SIGTERM"];
    for (const signal of signals) {
      process.once(signal, () => resolve(signal));
    }
  });
}

async function main(): Promise<void> {
  const eventHubConnectionString = requireEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "STORAGE_CONNECTION_STRING",
  );
  const blobContainerName = requireEnvironmentVariable("BLOB_CONTAINER_NAME");
  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

  const producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName,
  );
  const checkpointStore = new BlobCheckpointStore(
    storageConnectionString,
    blobContainerName,
  );
  const consumer = new EventHubConsumerClient(
    consumerGroup,
    eventHubConnectionString,
    eventHubName,
    checkpointStore,
  );

  let subscription: Subscription | undefined;

  try {
    const batch = await producer.createBatch();
    await addEvents(batch);
    await producer.sendBatch(batch);
    console.log(`Sent ${batch.count} events.`);

    subscription = consumer.subscribe({
      processEvents: async (events, context) => {
        for (const event of events) {
          printEvent(event, context.partitionId);
        }

        const lastEvent = events.at(-1);
        if (lastEvent) {
          await context.updateCheckpoint(lastEvent);
          console.log(
            `Updated checkpoint for partition ${context.partitionId} at sequence ${lastEvent.sequenceNumber}.`,
          );
        }
      },
      processError: async (error, context) => {
        console.error(
          `Error while processing partition ${context.partitionId}:`,
          error,
        );
      },
    });

    console.log("Receiving events. Press Ctrl+C to stop.");
    const signal = await waitForShutdownSignal();
    console.log(`Received ${signal}; shutting down.`);
  } finally {
    await subscription?.close();
    await consumer.close();
    await producer.close();
  }
}

function printEvent(event: ReceivedEventData, partitionId: string): void {
  console.log(
    `Partition ${partitionId}, sequence ${event.sequenceNumber}:`,
    event.body,
  );
}

main().catch((error: unknown) => {
  console.error("Event Hubs sample failed:", error);
  process.exitCode = 1;
});
