import {
  earliestEventPosition,
  EventHubConsumerClient,
  EventHubProducerClient,
  type Subscription,
} from "@azure/event-hubs";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { ContainerClient } from "@azure/storage-blob";

let producer: EventHubProducerClient | undefined;
let consumer: EventHubConsumerClient | undefined;
let subscription: Subscription | undefined;
let shuttingDown = false;

async function main(): Promise<void> {
  const eventHubsConnectionString = requireEnvironmentVariable(
    "EVENT_HUBS_CONNECTION_STRING",
  );
  const eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
  const storageConnectionString = requireEnvironmentVariable(
    "AZURE_STORAGE_CONNECTION_STRING",
  );
  const checkpointContainerName = requireEnvironmentVariable(
    "CHECKPOINT_CONTAINER_NAME",
  );
  const consumerGroup = process.env.EVENT_HUB_CONSUMER_GROUP ?? "$Default";

  producer = new EventHubProducerClient(
    eventHubsConnectionString,
    eventHubName,
  );

  const containerClient = new ContainerClient(
    storageConnectionString,
    checkpointContainerName,
  );
  const checkpointStore = new BlobCheckpointStore(containerClient);
  consumer = new EventHubConsumerClient(
    consumerGroup,
    eventHubsConnectionString,
    eventHubName,
    checkpointStore,
  );

  // The checkpoint container must exist before BlobCheckpointStore can use it.
  await containerClient.createIfNotExists();

  const batch = await producer.createBatch();

  for (let index = 1; index <= 10; index += 1) {
    const added = batch.tryAdd({
      body: {
        id: index,
        message: `Event ${index}`,
        createdAt: new Date().toISOString(),
      },
      properties: {
        source: "typescript-sample",
        sequence: index,
        category: index % 2 === 0 ? "even" : "odd",
      },
    });

    if (!added) {
      throw new Error(`Event ${index} is too large to fit in the batch.`);
    }
  }

  await producer.sendBatch(batch);
  console.log(`Sent ${batch.count} events.`);

  subscription = consumer.subscribe(
    {
      processEvents: async (events, context) => {
        const lastEvent = events.at(-1);
        if (!lastEvent) {
          return;
        }

        for (const event of events) {
          console.log(
            `Received from partition ${context.partitionId}:`,
            event.body,
          );
        }

        // Checkpoint the last event only after every event in this delivery succeeds.
        await context.updateCheckpoint(lastEvent);
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
    },
  );

  console.log("Listening for events. Press Ctrl+C to stop.");
}

async function shutdown(signal: string): Promise<void> {
  if (shuttingDown) {
    return;
  }

  shuttingDown = true;
  console.log(`\nReceived ${signal}; shutting down...`);

  await subscription?.close();
  await Promise.all([consumer?.close(), producer?.close()]);
  console.log("Event Hubs clients closed.");
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

for (const signal of ["SIGINT", "SIGTERM"] as const) {
  process.once(signal, () => {
    void shutdown(signal).catch((error: unknown) => {
      console.error("Shutdown failed:", error);
      process.exitCode = 1;
    });
  });
}

main().catch(async (error: unknown) => {
  console.error("Application failed:", error);
  process.exitCode = 1;

  try {
    await shutdown("application error");
  } catch (shutdownError: unknown) {
    console.error("Cleanup failed:", shutdownError);
  }
});
