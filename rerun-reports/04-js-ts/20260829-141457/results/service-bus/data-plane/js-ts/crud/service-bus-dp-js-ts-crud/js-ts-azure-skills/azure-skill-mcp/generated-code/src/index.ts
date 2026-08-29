import {
  ServiceBusClient,
  type ProcessErrorArgs,
  type ServiceBusReceivedMessage,
} from "@azure/service-bus";

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
const topicSender = client.createSender(topicName);
const subscriptionReceiver = client.createReceiver(
  topicName,
  subscriptionName,
  { receiveMode: "peekLock" },
);

async function main(): Promise<void> {
  await queueSender.sendMessages({
    body: "Single queue message",
    subject: "single-message",
  });
  console.log("Sent one message to the queue.");

  const batch = await queueSender.createMessageBatch();
  for (let messageNumber = 1; messageNumber <= 5; messageNumber += 1) {
    const added = batch.tryAddMessage({
      body: `Batch queue message ${messageNumber}`,
      subject: "batch-message",
      applicationProperties: { messageNumber },
    });

    if (!added) {
      throw new Error(`Message ${messageNumber} did not fit in the batch.`);
    }
  }

  await queueSender.sendMessages(batch);
  console.log(`Sent a batch of ${batch.count} messages to the queue.`);

  const receivedMessages = await queueReceiver.receiveMessages(6, {
    maxWaitTimeInMs: 10_000,
  });

  for (const message of receivedMessages) {
    await processMessage(message, "Queue receiveMessages()");
    await queueReceiver.completeMessage(message);
  }
  console.log(`Received and completed ${receivedMessages.length} queue messages.`);

  let resolveSubscribedMessage!: () => void;
  let rejectSubscribedMessage!: (error: Error) => void;
  const subscribedMessageProcessed = new Promise<void>((resolve, reject) => {
    resolveSubscribedMessage = resolve;
    rejectSubscribedMessage = reject;
  });

  queueReceiver.subscribe(
    {
      processMessage: async (message) => {
        await processMessage(message, "Queue subscribe()");
        await queueReceiver.completeMessage(message);
        resolveSubscribedMessage();
      },
      processError: async (args: ProcessErrorArgs) => {
        logProcessError(args);
        rejectSubscribedMessage(args.error);
      },
    },
    {
      autoCompleteMessages: false,
      maxConcurrentCalls: 1,
    },
  );

  await queueSender.sendMessages({
    body: "Message for the subscribe() handler",
    subject: "subscribed-message",
  });
  await withTimeout(
    subscribedMessageProcessed,
    15_000,
    "Timed out waiting for the queue subscribe() handler.",
  );

  await topicSender.sendMessages({
    body: "Message sent through a topic",
    subject: "topic-message",
  });
  console.log("Sent one message to the topic.");

  const topicMessages = await subscriptionReceiver.receiveMessages(1, {
    maxWaitTimeInMs: 10_000,
  });

  if (topicMessages.length === 0) {
    throw new Error("No topic message arrived at the subscription.");
  }

  for (const message of topicMessages) {
    await processMessage(message, "Topic subscription");
    await subscriptionReceiver.completeMessage(message);
  }
  console.log("Received and completed the topic message from the subscription.");
}

async function processMessage(
  message: ServiceBusReceivedMessage,
  source: string,
): Promise<void> {
  console.log(`${source}:`, {
    messageId: message.messageId,
    subject: message.subject,
    body: message.body,
  });
}

function logProcessError(args: ProcessErrorArgs): void {
  console.error("Service Bus subscription error:", {
    namespace: args.fullyQualifiedNamespace,
    entityPath: args.entityPath,
    errorSource: args.errorSource,
    error: args.error,
  });
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

async function withTimeout(
  operation: Promise<void>,
  timeoutInMs: number,
  message: string,
): Promise<void> {
  let timeout: NodeJS.Timeout | undefined;
  const timeoutPromise = new Promise<never>((_, reject) => {
    timeout = setTimeout(() => reject(new Error(message)), timeoutInMs);
  });

  try {
    await Promise.race([operation, timeoutPromise]);
  } finally {
    if (timeout) {
      clearTimeout(timeout);
    }
  }
}

try {
  await main();
} catch (error) {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
} finally {
  const closeResults = await Promise.allSettled([
    queueReceiver.close(),
    subscriptionReceiver.close(),
    queueSender.close(),
    topicSender.close(),
  ]);

  for (const result of closeResults) {
    if (result.status === "rejected") {
      console.error("Failed to close a Service Bus resource:", result.reason);
      process.exitCode = 1;
    }
  }

  try {
    await client.close();
  } catch (error) {
    console.error("Failed to close ServiceBusClient:", error);
    process.exitCode = 1;
  }
}
