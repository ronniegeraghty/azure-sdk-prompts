import {
  type ProcessErrorArgs,
  ServiceBusClient,
  type ServiceBusReceiver,
  type ServiceBusSender,
} from "@azure/service-bus";

type ServiceBusSubscription = ReturnType<ServiceBusReceiver["subscribe"]>;

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function logSubscriptionError(args: ProcessErrorArgs): void {
  console.error("Service Bus subscription error", {
    source: args.errorSource,
    entityPath: args.entityPath,
    namespace: args.fullyQualifiedNamespace,
    error: args.error,
  });
}

async function run(): Promise<void> {
  const connectionString = requiredEnvironmentVariable(
    "SERVICEBUS_CONNECTION_STRING",
  );
  const queueName = requiredEnvironmentVariable("SERVICEBUS_QUEUE_NAME");
  const topicName = requiredEnvironmentVariable("SERVICEBUS_TOPIC_NAME");
  const topicSubscriptionName = requiredEnvironmentVariable(
    "SERVICEBUS_SUBSCRIPTION_NAME",
  );

  const client = new ServiceBusClient(connectionString);
  let queueSender: ServiceBusSender | undefined;
  let pullReceiver: ServiceBusReceiver | undefined;
  let subscriptionReceiver: ServiceBusReceiver | undefined;
  let subscription: ServiceBusSubscription | undefined;
  let topicSender: ServiceBusSender | undefined;
  let topicReceiver: ServiceBusReceiver | undefined;

  try {
    queueSender = client.createSender(queueName);

    await queueSender.sendMessages({
      body: { kind: "single", text: "Hello from a single message" },
      contentType: "application/json",
      messageId: `single-${Date.now()}`,
    });
    console.log("Sent one queue message");

    const batch = await queueSender.createMessageBatch();
    for (let number = 1; number <= 5; number += 1) {
      const wasAdded = batch.tryAddMessage({
        body: { kind: "batch", number },
        contentType: "application/json",
        messageId: `batch-${Date.now()}-${number}`,
      });

      if (!wasAdded) {
        throw new Error(`Batch message ${number} did not fit in the batch`);
      }
    }
    await queueSender.sendMessages(batch);
    console.log(`Sent a batch of ${batch.count} queue messages`);

    pullReceiver = client.createReceiver(queueName, {
      receiveMode: "peekLock",
    });
    const pulledMessages = await pullReceiver.receiveMessages(3, {
      maxWaitTimeInMs: 10_000,
    });

    for (const message of pulledMessages) {
      console.log("Pulled queue message:", message.body);
      await pullReceiver.completeMessage(message);
    }

    subscriptionReceiver = client.createReceiver(queueName, {
      receiveMode: "peekLock",
    });

    let subscribedMessageCount = 0;
    let resolveSubscription: (() => void) | undefined;
    let rejectSubscription: ((error: unknown) => void) | undefined;
    const subscriptionFinished = new Promise<void>((resolve, reject) => {
      resolveSubscription = resolve;
      rejectSubscription = reject;
    });

    subscription = subscriptionReceiver.subscribe(
      {
        processMessage: async (message) => {
          console.log("Subscribed queue message:", message.body);
          await subscriptionReceiver?.completeMessage(message);
          subscribedMessageCount += 1;

          if (subscribedMessageCount >= 3) {
            resolveSubscription?.();
          }
        },
        processError: async (args) => {
          logSubscriptionError(args);
          rejectSubscription?.(args.error);
        },
      },
      {
        autoCompleteMessages: false,
        maxConcurrentCalls: 1,
      },
    );

    let timeoutId: NodeJS.Timeout | undefined;
    const timeout = new Promise<never>((_, reject) => {
      timeoutId = setTimeout(
        () => reject(new Error("Timed out waiting for subscribed messages")),
        15_000,
      );
    });
    try {
      await Promise.race([subscriptionFinished, timeout]);
    } finally {
      clearTimeout(timeoutId);
    }
    await subscription.close();
    subscription = undefined;

    topicSender = client.createSender(topicName);
    await topicSender.sendMessages({
      body: { event: "order.created", orderId: "order-123" },
      contentType: "application/json",
      applicationProperties: { eventType: "order.created" },
      messageId: `topic-${Date.now()}`,
    });
    console.log("Sent one topic message");

    topicReceiver = client.createReceiver(topicName, topicSubscriptionName, {
      receiveMode: "peekLock",
    });
    const topicMessages = await topicReceiver.receiveMessages(1, {
      maxWaitTimeInMs: 10_000,
    });

    for (const message of topicMessages) {
      console.log("Received topic subscription message:", message.body);
      await topicReceiver.completeMessage(message);
    }
  } finally {
    try {
      await subscription?.close();
    } finally {
      try {
        await Promise.all([
          topicReceiver?.close(),
          topicSender?.close(),
          subscriptionReceiver?.close(),
          pullReceiver?.close(),
          queueSender?.close(),
        ]);
      } finally {
        await client.close();
      }
    }
  }
}

run().catch((error: unknown) => {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
});
