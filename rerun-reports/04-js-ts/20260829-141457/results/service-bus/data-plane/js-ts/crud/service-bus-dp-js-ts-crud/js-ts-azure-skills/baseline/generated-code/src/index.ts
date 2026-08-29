import { ServiceBusClient, ServiceBusReceiver } from "@azure/service-bus";

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

async function receiveAndCompleteQueueMessages(
  receiver: ServiceBusReceiver,
  expectedCount: number,
): Promise<void> {
  const messages = await receiver.receiveMessages(expectedCount, {
    maxWaitTimeInMs: 10_000,
  });

  for (const message of messages) {
    console.log("Queue message:", message.body);
    await receiver.completeMessage(message);
  }

  console.log(`Received and completed ${messages.length} queue message(s).`);
}

async function demonstrateSubscription(
  client: ServiceBusClient,
  queueName: string,
): Promise<void> {
  const receiver = client.createReceiver(queueName, {
    receiveMode: "peekLock",
  });
  const sender = client.createSender(queueName);

  try {
    const processed = new Promise<void>((resolve, reject) => {
      let settled = false;
      const timeout = setTimeout(() => {
        if (!settled) {
          settled = true;
          reject(new Error("Timed out waiting for the subscribed queue message."));
        }
      }, 15_000);

      receiver.subscribe(
        {
          processMessage: async (message) => {
            console.log("Subscribed queue message:", message.body);
            await receiver.completeMessage(message);
            if (!settled) {
              settled = true;
              clearTimeout(timeout);
              resolve();
            }
          },
          processError: async (args) => {
            console.error("Subscription error:", args.error);
            if (!settled) {
              settled = true;
              clearTimeout(timeout);
              reject(args.error);
            }
          },
        },
        { autoCompleteMessages: false },
      );
    });

    await sender.sendMessages({
      body: "Message handled by subscribe()",
      contentType: "text/plain",
    });
    await processed;
  } finally {
    await receiver.close();
    await sender.close();
  }
}

async function demonstrateTopicAndSubscription(
  client: ServiceBusClient,
  topicName: string,
  subscriptionName: string,
): Promise<void> {
  const sender = client.createSender(topicName);
  const receiver = client.createReceiver(topicName, subscriptionName, {
    receiveMode: "peekLock",
  });

  try {
    await sender.sendMessages({
      body: "Hello from the topic",
      subject: "topic-demo",
    });

    const messages = await receiver.receiveMessages(1, {
      maxWaitTimeInMs: 10_000,
    });

    for (const message of messages) {
      console.log("Topic subscription message:", message.body);
      await receiver.completeMessage(message);
    }
  } finally {
    await receiver.close();
    await sender.close();
  }
}

async function main(): Promise<void> {
  const connectionString = requiredEnvironmentVariable(
    "SERVICE_BUS_CONNECTION_STRING",
  );
  const queueName = requiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
  const topicName = requiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
  const subscriptionName = requiredEnvironmentVariable(
    "SERVICE_BUS_SUBSCRIPTION_NAME",
  );

  const client = new ServiceBusClient(connectionString);
  const queueSender = client.createSender(queueName);
  const queueReceiver = client.createReceiver(queueName, {
    receiveMode: "peekLock",
  });

  try {
    try {
      await queueSender.sendMessages({
        body: "Single queue message",
        contentType: "text/plain",
      });

      const batch = await queueSender.createMessageBatch();
      for (let index = 1; index <= 5; index += 1) {
        const added = batch.tryAddMessage({
          body: `Batch queue message ${index}`,
          messageId: `batch-${index}`,
        });

        if (!added) {
          throw new Error(
            `Batch queue message ${index} did not fit in the batch.`,
          );
        }
      }
      await queueSender.sendMessages(batch);

      await receiveAndCompleteQueueMessages(queueReceiver, 6);
    } finally {
      await queueReceiver.close();
      await queueSender.close();
    }

    await demonstrateSubscription(client, queueName);
    await demonstrateTopicAndSubscription(
      client,
      topicName,
      subscriptionName,
    );
  } finally {
    await client.close();
  }
}

main().catch((error: unknown) => {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
});
