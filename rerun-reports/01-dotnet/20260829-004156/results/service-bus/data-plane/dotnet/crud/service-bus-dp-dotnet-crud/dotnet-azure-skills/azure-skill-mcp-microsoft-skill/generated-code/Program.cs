using Azure.Messaging.ServiceBus;

const string ConnectionStringVariable = "AZURE_SERVICEBUS_CONNECTION_STRING";
const string QueueNameVariable = "AZURE_SERVICEBUS_QUEUE_NAME";
const string TopicNameVariable = "AZURE_SERVICEBUS_TOPIC_NAME";
const string SubscriptionNameVariable = "AZURE_SERVICEBUS_SUBSCRIPTION_NAME";

string connectionString = GetRequiredEnvironmentVariable(ConnectionStringVariable);
string queueName = GetRequiredEnvironmentVariable(QueueNameVariable);
string topicName = GetRequiredEnvironmentVariable(TopicNameVariable);
string subscriptionName = GetRequiredEnvironmentVariable(SubscriptionNameVariable);

try
{
    await using ServiceBusClient client = new(connectionString);
    await using ServiceBusSender queueSender = client.CreateSender(queueName);

    await queueSender.SendMessageAsync(
        new ServiceBusMessage("Single queue message")
        {
            MessageId = Guid.NewGuid().ToString()
        });
    Console.WriteLine("Sent one message to the queue.");

    using (ServiceBusMessageBatch batch = await queueSender.CreateMessageBatchAsync())
    {
        for (int i = 1; i <= 5; i++)
        {
            ServiceBusMessage message = new($"Batch message {i}")
            {
                MessageId = Guid.NewGuid().ToString(),
                ApplicationProperties =
                {
                    ["BatchSequence"] = i
                }
            };

            if (!batch.TryAddMessage(message))
            {
                throw new InvalidOperationException(
                    $"Batch message {i} is too large for an empty or partially filled batch.");
            }
        }

        await queueSender.SendMessagesAsync(batch);
    }
    Console.WriteLine("Sent a batch of five messages to the queue.");

    await using ServiceBusReceiver queueReceiver = client.CreateReceiver(
        queueName,
        new ServiceBusReceiverOptions
        {
            ReceiveMode = ServiceBusReceiveMode.PeekLock
        });

    IReadOnlyList<ServiceBusReceivedMessage> receivedMessages =
        await queueReceiver.ReceiveMessagesAsync(
            maxMessages: 6,
            maxWaitTime: TimeSpan.FromSeconds(10));

    foreach (ServiceBusReceivedMessage message in receivedMessages)
    {
        Console.WriteLine($"Pull receiver processed: {message.Body}");
        await queueReceiver.CompleteMessageAsync(message);
    }

    await using ServiceBusProcessor processor = client.CreateProcessor(
        queueName,
        new ServiceBusProcessorOptions
        {
            AutoCompleteMessages = false,
            MaxConcurrentCalls = 2
        });

    processor.ProcessMessageAsync += async args =>
    {
        Console.WriteLine($"Processor handled: {args.Message.Body}");
        await args.CompleteMessageAsync(args.Message);
    };

    processor.ProcessErrorAsync += args =>
    {
        Console.Error.WriteLine(
            $"Processor error ({args.ErrorSource}) on {args.EntityPath}: {args.Exception.Message}");
        return Task.CompletedTask;
    };

    bool processorStarted = false;
    try
    {
        await processor.StartProcessingAsync();
        processorStarted = true;

        await queueSender.SendMessageAsync(
            new ServiceBusMessage("Message for the continuous processor")
            {
                MessageId = Guid.NewGuid().ToString()
            });

        Console.WriteLine("Processor is running for 10 seconds.");
        await Task.Delay(TimeSpan.FromSeconds(10));
    }
    finally
    {
        if (processorStarted)
        {
            await processor.StopProcessingAsync();
        }
    }

    await using ServiceBusSender topicSender = client.CreateSender(topicName);
    await topicSender.SendMessageAsync(
        new ServiceBusMessage("Message published to the topic")
        {
            MessageId = Guid.NewGuid().ToString()
        });
    Console.WriteLine("Sent one message to the topic.");

    await using ServiceBusReceiver subscriptionReceiver =
        client.CreateReceiver(topicName, subscriptionName);

    ServiceBusReceivedMessage? subscriptionMessage =
        await subscriptionReceiver.ReceiveMessageAsync(TimeSpan.FromSeconds(10));

    if (subscriptionMessage is null)
    {
        Console.WriteLine("No subscription message arrived within 10 seconds.");
    }
    else
    {
        Console.WriteLine($"Subscription received: {subscriptionMessage.Body}");
        await subscriptionReceiver.CompleteMessageAsync(subscriptionMessage);
    }
}
catch (ServiceBusException exception)
{
    Console.Error.WriteLine(
        $"Service Bus operation failed ({exception.Reason}): {exception.Message}");
    Environment.ExitCode = 1;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Set the required environment variable '{name}' before running the sample.");
    }

    return value;
}
