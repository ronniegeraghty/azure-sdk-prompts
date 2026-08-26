using Azure.Messaging.ServiceBus;

string connectionString = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_CONNECTION_STRING");
string queueName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_QUEUE_NAME");
string topicName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_TOPIC_NAME");
string subscriptionName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_SUBSCRIPTION_NAME");

await using var client = new ServiceBusClient(connectionString);

await SendToQueueAsync(client, queueName);
await ReceiveFromQueueAsync(client, queueName);
await ProcessQueueContinuouslyAsync(client, queueName);
await SendToTopicAndReceiveFromSubscriptionAsync(
    client,
    topicName,
    subscriptionName);

static async Task SendToQueueAsync(ServiceBusClient client, string queueName)
{
    await using ServiceBusSender sender = client.CreateSender(queueName);

    await sender.SendMessageAsync(
        new ServiceBusMessage("Single queue message")
        {
            MessageId = Guid.NewGuid().ToString()
        });

    using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();

    for (int index = 1; index <= 5; index++)
    {
        var message = new ServiceBusMessage($"Batch message {index}")
        {
            MessageId = Guid.NewGuid().ToString(),
            ApplicationProperties =
            {
                ["BatchIndex"] = index
            }
        };

        if (!batch.TryAddMessage(message))
        {
            throw new InvalidOperationException(
                $"Batch message {index} is too large to fit in an empty Service Bus batch.");
        }
    }

    await sender.SendMessagesAsync(batch);
    Console.WriteLine("Sent one queue message and a batch of five queue messages.");
}

static async Task ReceiveFromQueueAsync(ServiceBusClient client, string queueName)
{
    await using ServiceBusReceiver receiver = client.CreateReceiver(
        queueName,
        new ServiceBusReceiverOptions
        {
            ReceiveMode = ServiceBusReceiveMode.PeekLock
        });

    IReadOnlyList<ServiceBusReceivedMessage> messages =
        await receiver.ReceiveMessagesAsync(
            maxMessages: 6,
            maxWaitTime: TimeSpan.FromSeconds(10));

    foreach (ServiceBusReceivedMessage message in messages)
    {
        Console.WriteLine($"Received from queue: {message.Body}");

        // Complete only after successful processing to remove the message.
        await receiver.CompleteMessageAsync(message);
    }
}

static async Task ProcessQueueContinuouslyAsync(
    ServiceBusClient client,
    string queueName)
{
    await using ServiceBusProcessor processor = client.CreateProcessor(
        queueName,
        new ServiceBusProcessorOptions
        {
            AutoCompleteMessages = false,
            MaxConcurrentCalls = 2
        });

    var processedMessage = new TaskCompletionSource(
        TaskCreationOptions.RunContinuationsAsynchronously);

    processor.ProcessMessageAsync += async args =>
    {
        Console.WriteLine($"Processor received: {args.Message.Body}");
        await args.CompleteMessageAsync(args.Message);
        processedMessage.TrySetResult();
    };

    processor.ProcessErrorAsync += args =>
    {
        Console.Error.WriteLine(
            $"Processor error on {args.EntityPath} ({args.ErrorSource}): " +
            args.Exception.Message);
        return Task.CompletedTask;
    };

    await processor.StartProcessingAsync();

    try
    {
        await using ServiceBusSender sender = client.CreateSender(queueName);
        await sender.SendMessageAsync(
            new ServiceBusMessage("Message for the continuous processor")
            {
                MessageId = Guid.NewGuid().ToString()
            });

        await processedMessage.Task.WaitAsync(TimeSpan.FromSeconds(30));
    }
    finally
    {
        await processor.StopProcessingAsync();
    }
}

static async Task SendToTopicAndReceiveFromSubscriptionAsync(
    ServiceBusClient client,
    string topicName,
    string subscriptionName)
{
    await using ServiceBusSender topicSender = client.CreateSender(topicName);
    await using ServiceBusReceiver subscriptionReceiver =
        client.CreateReceiver(topicName, subscriptionName);

    await topicSender.SendMessageAsync(
        new ServiceBusMessage("Topic message")
        {
            MessageId = Guid.NewGuid().ToString()
        });

    ServiceBusReceivedMessage? message =
        await subscriptionReceiver.ReceiveMessageAsync(
            maxWaitTime: TimeSpan.FromSeconds(10));

    if (message is null)
    {
        throw new TimeoutException(
            $"No message arrived on subscription '{subscriptionName}'.");
    }

    Console.WriteLine($"Received from topic subscription: {message.Body}");
    await subscriptionReceiver.CompleteMessageAsync(message);
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
