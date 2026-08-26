import {
  type ProcessErrorArgs,
  type ServiceBusReceiver,
  type ServiceBusSender,
  ServiceBusClient,
  ServiceBusError,
} from "@azure/service-bus";

const connectionString = requireEnvironmentVariable(
  "AZURE_SERVICE_BUS_CONNECTION_STRING",
);
const queueName = requireEnvironmentVariable("AZURE_SERVICE_BUS_QUEUE_NAME");
const topicName = requireEnvironmentVariable("AZURE_SERVICE_BUS_TOPIC_NAME");
const subscriptionName = requireEnvironmentVariable(
  "AZURE_SERVICE_BUS_SUBSCRIPTION_NAME",
);

const client = new ServiceBusClient(connectionString);

let queueSender: ServiceBusSender | undefined;
let pullReceiver: ServiceBusReceiver | undefined;
let subscribedReceiver: ServiceBusReceiver | undefined;
let messageSubscription:
  | ReturnType<ServiceBusReceiver["subscribe"]>
  | undefined;
let topicSender: ServiceBusSender | undefined;
let topicReceiver: ServiceBusReceiver | undefined;

async function main(): Promise<void> {
  try {
    queueSender = client.createSender(queueName);

    await queueSender.sendMessages({
      body: { kind: "single", text: "Hello from Azure Service Bus" },
      contentType: "application/json",
      messageId: "single-message",
    });
    console.log("Sent one queue message.");

    const batch = await queueSender.createMessageBatch();
    for (let index = 1; index <= 5; index += 1) {
      const added = batch.tryAddMessage({
        body: { kind: "batch", sequence: index },
        contentType: "application/json",
        messageId: `batch-message-${index}`,
      });

      if (!added) {
        throw new Error(`Batch is full; message ${index} could not be added.`);
      }
    }
    await queueSender.sendMessages(batch);
    console.log(`Sent a batch of ${batch.count} queue messages.`);

    pullReceiver = client.createReceiver(queueName, {
      receiveMode: "peekLock",
    });
    const receivedMessages = await pullReceiver.receiveMessages(6, {
      maxWaitTimeInMs: 5_000,
    });

    for (const message of receivedMessages) {
      console.log("Pull receiver processed:", message.body);
      await pullReceiver.completeMessage(message);
    }
    console.log(`Completed ${receivedMessages.length} queue messages.`);

    await pullReceiver.close();
    pullReceiver = undefined;

    subscribedReceiver = client.createReceiver(queueName, {
      receiveMode: "peekLock",
    });
    messageSubscription = subscribedReceiver.subscribe(
      {
        processMessage: async (message) => {
          console.log("Subscriber processed:", message.body);
          await subscribedReceiver!.completeMessage(message);
        },
        processError: async (args) => {
          logServiceBusError("Queue subscription error", args);
        },
      },
      {
        autoCompleteMessages: false,
        maxConcurrentCalls: 1,
      },
    );

    await queueSender.sendMessages({
      body: { kind: "subscription", text: "Hello, subscriber" },
      contentType: "application/json",
      messageId: `subscription-message-${Date.now()}`,
    });
    console.log("Subscriber is listening for 5 seconds.");
    await delay(5_000);

    await messageSubscription.close();
    messageSubscription = undefined;
    await subscribedReceiver.close();
    subscribedReceiver = undefined;

    topicSender = client.createSender(topicName);
    await topicSender.sendMessages({
      body: { event: "order.created", orderId: "order-123" },
      contentType: "application/json",
      subject: "orders/created",
      applicationProperties: { eventType: "order.created" },
    });
    console.log("Sent one topic message.");

    topicReceiver = client.createReceiver(topicName, subscriptionName, {
      receiveMode: "peekLock",
    });
    const topicMessages = await topicReceiver.receiveMessages(1, {
      maxWaitTimeInMs: 5_000,
    });

    for (const message of topicMessages) {
      console.log("Topic subscription processed:", message.body);
      await topicReceiver.completeMessage(message);
    }
  } finally {
    await closeResources();
  }
}

async function closeResources(): Promise<void> {
  await messageSubscription?.close();
  await topicReceiver?.close();
  await subscribedReceiver?.close();
  await pullReceiver?.close();
  await topicSender?.close();
  await queueSender?.close();
  await client.close();
}

function logServiceBusError(context: string, args: ProcessErrorArgs): void {
  const { error } = args;

  if (error instanceof ServiceBusError) {
    console.error(
      `${context}: ${error.code} (retryable: ${error.retryable}) - ${error.message}`,
    );
    return;
  }

  console.error(`${context}:`, error);
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function delay(milliseconds: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

main().catch((error: unknown) => {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
});
