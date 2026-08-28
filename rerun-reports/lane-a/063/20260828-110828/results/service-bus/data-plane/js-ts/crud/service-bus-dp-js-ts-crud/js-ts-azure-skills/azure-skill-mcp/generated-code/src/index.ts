import {
  ServiceBusClient,
  type ProcessErrorArgs,
  type ServiceBusReceivedMessage,
} from "@azure/service-bus";

const connectionString = requireEnvironmentVariable(
  "SERVICE_BUS_CONNECTION_STRING",
);
const queueName = requireEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
const topicName = requireEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
const subscriptionName = requireEnvironmentVariable(
  "SERVICE_BUS_SUBSCRIPTION_NAME",
);

const client = new ServiceBusClient(connectionString);
const queueSender = client.createSender(queueName);
const queueReceiver = client.createReceiver(queueName, {
  receiveMode: "peekLock",
});
const topicSender = client.createSender(topicName);
const topicSubscriptionReceiver = client.createReceiver(
  topicName,
  subscriptionName,
  { receiveMode: "peekLock" },
);

let closeQueueSubscription: (() => Promise<void>) | undefined;

try {
  await queueSender.sendMessages({
    body: "Single queue message",
    contentType: "text/plain",
  });
  console.log("Sent one message to the queue.");

  const batch = await queueSender.createMessageBatch();
  for (let sequence = 1; sequence <= 5; sequence += 1) {
    const added = batch.tryAddMessage({
      body: { sequence, text: `Batch message ${sequence}` },
      contentType: "application/json",
    });

    if (!added) {
      throw new Error(`Batch message ${sequence} exceeded the batch size limit.`);
    }
  }
  await queueSender.sendMessages(batch);
  console.log("Sent a batch of 5 messages to the queue.");

  const receivedMessages = await queueReceiver.receiveMessages(6, {
    maxWaitTimeInMs: 10_000,
  });

  for (const message of receivedMessages) {
    await processMessage(message);
    await queueReceiver.completeMessage(message);
  }
  console.log(`Received and completed ${receivedMessages.length} queue messages.`);

  const subscribedMessageProcessed = new Promise<void>((resolve) => {
    const subscription = queueReceiver.subscribe({
      processMessage: async (message: ServiceBusReceivedMessage) => {
        await processMessage(message);
        await queueReceiver.completeMessage(message);
        resolve();
      },
      processError: async (args: ProcessErrorArgs) => {
        console.error(
          `Queue subscription error (${args.errorSource}) in ${args.entityPath}:`,
          args.error,
        );
      },
    });

    closeQueueSubscription = () => subscription.close();
  });

  await queueSender.sendMessages({
    body: "Message handled by subscribe()",
    contentType: "text/plain",
  });
  await waitWithTimeout(
    subscribedMessageProcessed,
    15_000,
    "Timed out waiting for the subscribe() handler.",
  );
  console.log("Processed and completed one message using subscribe().");

  await topicSender.sendMessages({
    body: "Topic message",
    subject: "topic-demo",
    contentType: "text/plain",
  });

  const topicMessages = await topicSubscriptionReceiver.receiveMessages(1, {
    maxWaitTimeInMs: 10_000,
  });
  if (topicMessages.length === 0) {
    throw new Error("No message was received from the topic subscription.");
  }

  for (const message of topicMessages) {
    await processMessage(message);
    await topicSubscriptionReceiver.completeMessage(message);
  }
  console.log("Sent to a topic and received from its subscription.");
} finally {
  try {
    if (closeQueueSubscription) {
      await closeQueueSubscription();
    }
  } finally {
    try {
      await Promise.all([
        queueReceiver.close(),
        topicSubscriptionReceiver.close(),
        queueSender.close(),
        topicSender.close(),
      ]);
    } finally {
      await client.close();
    }
  }
}

async function processMessage(
  message: ServiceBusReceivedMessage,
): Promise<void> {
  console.log(`Processing message ${message.messageId}:`, message.body);
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function waitWithTimeout(
  operation: Promise<void>,
  timeoutInMs: number,
  timeoutMessage: string,
): Promise<void> {
  return new Promise((resolve, reject) => {
    const timeout = setTimeout(() => reject(new Error(timeoutMessage)), timeoutInMs);

    operation.then(
      () => {
        clearTimeout(timeout);
        resolve();
      },
      (error: unknown) => {
        clearTimeout(timeout);
        reject(error);
      },
    );
  });
}
