import {
  ServiceBusClient,
  ServiceBusMessage,
  ServiceBusReceiver,
  ServiceBusSender,
  ServiceBusError,
  type ProcessErrorArgs,
  type ServiceBusReceivedMessage,
} from "@azure/service-bus";

interface Closeable {
  close(): Promise<void>;
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }

  return value;
}

function logMessage(prefix: string, message: ServiceBusReceivedMessage): void {
  console.log(`${prefix} ${String(message.messageId)}:`, message.body);
}

function logSubscriptionError(args: ProcessErrorArgs): void {
  console.error(
    `Subscription error from ${args.errorSource} on ${args.entityPath}:`,
    args.error,
  );

  if (args.error instanceof ServiceBusError) {
    console.error(
      `Service Bus error code=${args.error.code}, retryable=${args.error.retryable}`,
    );
  }
}

async function main(): Promise<void> {
  const connectionString = requireEnvironmentVariable(
    "SERVICEBUS_CONNECTION_STRING",
  );
  const queueName = requireEnvironmentVariable("SERVICEBUS_QUEUE_NAME");
  const topicName = requireEnvironmentVariable("SERVICEBUS_TOPIC_NAME");
  const subscriptionName = requireEnvironmentVariable(
    "SERVICEBUS_SUBSCRIPTION_NAME",
  );

  // The connection string is read from the environment; never embed credentials in code.
  const client = new ServiceBusClient(connectionString);
  const closeables: Closeable[] = [];

  const track = <T extends Closeable>(resource: T): T => {
    closeables.push(resource);
    return resource;
  };

  const closeNow = async (resource: Closeable): Promise<void> => {
    await resource.close();
    const index = closeables.lastIndexOf(resource);
    if (index !== -1) {
      closeables.splice(index, 1);
    }
  };

  try {
    const queueSender: ServiceBusSender = track(
      client.createSender(queueName),
    );

    const singleMessage: ServiceBusMessage = {
      body: { kind: "single", text: "Hello from Azure Service Bus" },
      contentType: "application/json",
      messageId: `single-${crypto.randomUUID()}`,
    };
    await queueSender.sendMessages(singleMessage);
    console.log("Sent one queue message");

    const batch = await queueSender.createMessageBatch();
    for (let index = 1; index <= 5; index += 1) {
      const wasAdded = batch.tryAddMessage({
        body: { kind: "batch", sequence: index },
        contentType: "application/json",
        messageId: `batch-${index}-${crypto.randomUUID()}`,
      });

      if (!wasAdded) {
        throw new Error(`Message ${index} did not fit in the Service Bus batch`);
      }
    }

    await queueSender.sendMessages(batch);
    console.log(`Sent a batch of ${batch.count} queue messages`);

    const queueReceiver: ServiceBusReceiver = track(
      client.createReceiver(queueName, { receiveMode: "peekLock" }),
    );
    const receivedMessages = await queueReceiver.receiveMessages(6, {
      maxWaitTimeInMs: 10_000,
    });

    for (const message of receivedMessages) {
      logMessage("Pulled queue message", message);
      await queueReceiver.completeMessage(message);
    }
    console.log(`Completed ${receivedMessages.length} pulled messages`);
    await closeNow(queueReceiver);

    const subscribedMessageId = `subscribed-${crypto.randomUUID()}`;
    const subscribedReceiver: ServiceBusReceiver = track(
      client.createReceiver(queueName, { receiveMode: "peekLock" }),
    );

    let notifySubscribedMessage: (() => void) | undefined;
    const subscribedMessageProcessed = new Promise<void>((resolve) => {
      notifySubscribedMessage = resolve;
    });

    const streamingSubscription = track(
      subscribedReceiver.subscribe(
        {
          processMessage: async (message) => {
            logMessage("Subscribed queue message", message);
            if (message.messageId === subscribedMessageId) {
              notifySubscribedMessage?.();
            }
            // subscribe() auto-completes when this handler succeeds.
          },
          processError: async (args) => {
            logSubscriptionError(args);
          },
        },
        {
          autoCompleteMessages: true,
          maxConcurrentCalls: 1,
        },
      ),
    );

    await queueSender.sendMessages({
      body: { kind: "subscription", text: "Process me with subscribe()" },
      contentType: "application/json",
      messageId: subscribedMessageId,
    });

    let subscriptionTimeout: NodeJS.Timeout | undefined;
    try {
      await Promise.race([
        subscribedMessageProcessed,
        new Promise<never>((_, reject) => {
          subscriptionTimeout = setTimeout(
            () => reject(new Error("Timed out waiting for subscribed message")),
            15_000,
          );
        }),
      ]);
    } finally {
      if (subscriptionTimeout) {
        clearTimeout(subscriptionTimeout);
      }
    }

    await closeNow(streamingSubscription);
    await closeNow(subscribedReceiver);
    await closeNow(queueSender);

    const topicSender: ServiceBusSender = track(
      client.createSender(topicName),
    );
    await topicSender.sendMessages({
      body: { event: "demo.created", createdAt: new Date().toISOString() },
      contentType: "application/json",
      subject: "demo/created",
      messageId: `topic-${crypto.randomUUID()}`,
    });
    console.log("Sent one topic message");

    const topicReceiver: ServiceBusReceiver = track(
      client.createReceiver(topicName, subscriptionName, {
        receiveMode: "peekLock",
      }),
    );
    const topicMessages = await topicReceiver.receiveMessages(1, {
      maxWaitTimeInMs: 10_000,
    });

    if (topicMessages.length === 0) {
      throw new Error(
        `No message arrived on topic subscription "${subscriptionName}"`,
      );
    }

    for (const message of topicMessages) {
      logMessage("Topic subscription message", message);
      await topicReceiver.completeMessage(message);
    }

    await closeNow(topicReceiver);
    await closeNow(topicSender);
  } finally {
    for (const resource of closeables.reverse()) {
      await resource.close();
    }
    await client.close();
  }
}

main().catch((error: unknown) => {
  console.error("Service Bus demo failed:", error);
  process.exitCode = 1;
});
