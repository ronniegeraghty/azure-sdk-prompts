import {
  ProcessErrorArgs,
  ServiceBusClient,
  ServiceBusMessage,
} from "@azure/service-bus";

const requiredEnvironmentVariable = (name: string): string => {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Set the ${name} environment variable before running this example.`);
  }
  return value;
};

const delay = (milliseconds: number): Promise<void> =>
  new Promise((resolve) => setTimeout(resolve, milliseconds));

async function main(): Promise<void> {
  const connectionString = requiredEnvironmentVariable(
    "AZURE_SERVICE_BUS_CONNECTION_STRING",
  );
  const queueName = requiredEnvironmentVariable("AZURE_SERVICE_BUS_QUEUE_NAME");
  const topicName = requiredEnvironmentVariable("AZURE_SERVICE_BUS_TOPIC_NAME");
  const subscriptionName = requiredEnvironmentVariable(
    "AZURE_SERVICE_BUS_SUBSCRIPTION_NAME",
  );

  // The queue, topic, and subscription must already exist.
  const client = new ServiceBusClient(connectionString);

  try {
    const queueSender = client.createSender(queueName);
    try {
      await queueSender.sendMessages({
        body: "Single queue message",
        contentType: "text/plain",
      });

      const batch = await queueSender.createMessageBatch();
      for (let sequence = 1; sequence <= 5; sequence += 1) {
        const message: ServiceBusMessage = {
          body: { sequence, text: `Batch message ${sequence}` },
          contentType: "application/json",
          messageId: `batch-${sequence}`,
        };

        if (!batch.tryAddMessage(message)) {
          throw new Error(`Batch message ${sequence} is too large for the batch.`);
        }
      }
      await queueSender.sendMessages(batch);
    } finally {
      await queueSender.close();
    }

    const queueReceiver = client.createReceiver(queueName);
    try {
      const messages = await queueReceiver.receiveMessages(6, {
        maxWaitTimeInMs: 10_000,
      });

      for (const message of messages) {
        console.log("Received from queue:", message.body);
        await queueReceiver.completeMessage(message);
      }
    } finally {
      await queueReceiver.close();
    }

    const subscriptionReceiver = client.createReceiver(queueName);
    const subscription = subscriptionReceiver.subscribe({
      processMessage: async (message) => {
        console.log("Subscription handler received:", message.body);
        await subscriptionReceiver.completeMessage(message);
      },
      processError: async (args: ProcessErrorArgs) => {
        console.error(
          `Subscription error from ${args.errorSource}:`,
          args.error,
        );
      },
    });

    try {
      // Keep the handler active briefly for this finite demo.
      await delay(10_000);
    } finally {
      await subscription.close();
      await subscriptionReceiver.close();
    }

    const topicSender = client.createSender(topicName);
    try {
      await topicSender.sendMessages({
        body: "Message sent to a topic",
        subject: "topic-demo",
      });
    } finally {
      await topicSender.close();
    }

    const topicReceiver = client.createReceiver(topicName, subscriptionName);
    try {
      const [topicMessage] = await topicReceiver.receiveMessages(1, {
        maxWaitTimeInMs: 10_000,
      });

      if (topicMessage) {
        console.log("Received from topic subscription:", topicMessage.body);
        await topicReceiver.completeMessage(topicMessage);
      } else {
        console.log("No topic message arrived before the receive timeout.");
      }
    } finally {
      await topicReceiver.close();
    }
  } finally {
    await client.close();
  }
}

main().catch((error: unknown) => {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
});
