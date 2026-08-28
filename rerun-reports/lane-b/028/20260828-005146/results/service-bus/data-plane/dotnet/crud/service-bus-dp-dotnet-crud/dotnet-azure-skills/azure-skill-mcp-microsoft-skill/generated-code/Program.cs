using Azure.Messaging.ServiceBus;

string connectionString = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_CONNECTION_STRING");
string queueName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_QUEUE_NAME");
string topicName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_TOPIC_NAME");
string subscriptionName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_SUBSCRIPTION_NAME");

await using var client = new ServiceBusClient(connectionString);

await SendAndReceiveQueueMessagesAsync(client, queueName);
await ProcessQueueMessageAsync(client, queueName);
await SendAndReceiveTopicMessageAsync(client, topicName, subscriptionName);

static async Task SendAndReceiveQueueMessagesAsync(
    ServiceBusClient client,
    string queueName)
{
    await using ServiceBusSender sender = client.CreateSender(queueName);

    await sender.SendMessageAsync(new ServiceBusMessage("Single queue message"));
    Console.WriteLine("Sent one message to the queue.");

    using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();

    for (int i = 1; i <= 5; i++)
    {
        var message = new ServiceBusMessage($"Batch message {i}")
        {
            MessageId = Guid.NewGuid().ToString()
        };

        if (!batch.TryAddMessage(message))
        {
            throw new InvalidOperationException(
                $"Batch message {i} is too large to fit in the Service Bus batch.");
        }
    }

    await sender.SendMessagesAsync(batch);
    Console.WriteLine("Sent a batch of 5 messages to the queue.");

    await using ServiceBusReceiver receiver = client.CreateReceiver(queueName);
    IReadOnlyList<ServiceBusReceivedMessage> messages =
        await receiver.ReceiveMessagesAsync(
            maxMessages: 6,
            maxWaitTime: TimeSpan.FromSeconds(10));

    Console.WriteLine($"Received {messages.Count} queue message(s).");

    foreach (ServiceBusReceivedMessage message in messages)
    {
        Console.WriteLine($"Queue message: {message.Body}");
        await receiver.CompleteMessageAsync(message);
        Console.WriteLine($"Completed message {message.MessageId}.");
    }
}

static async Task ProcessQueueMessageAsync(
    ServiceBusClient client,
    string queueName)
{
    await using ServiceBusSender sender = client.CreateSender(queueName);
    await sender.SendMessageAsync(new ServiceBusMessage("Message for the processor"));

    await using ServiceBusProcessor processor = client.CreateProcessor(
        queueName,
        new ServiceBusProcessorOptions
        {
            AutoCompleteMessages = false,
            MaxConcurrentCalls = 1
        });

    var processed = new TaskCompletionSource(
        TaskCreationOptions.RunContinuationsAsynchronously);

    processor.ProcessMessageAsync += async args =>
    {
        try
        {
            Console.WriteLine($"Processor message: {args.Message.Body}");
            await args.CompleteMessageAsync(args.Message);
            processed.TrySetResult();
        }
        catch (Exception exception)
        {
            processed.TrySetException(exception);
            throw;
        }
    };

    processor.ProcessErrorAsync += args =>
    {
        Console.Error.WriteLine(
            $"Processor error ({args.ErrorSource}, {args.EntityPath}): " +
            args.Exception.Message);
        processed.TrySetException(args.Exception);
        return Task.CompletedTask;
    };

    await processor.StartProcessingAsync();
    Console.WriteLine("Processor started.");

    try
    {
        await processed.Task.WaitAsync(TimeSpan.FromSeconds(30));
    }
    finally
    {
        await processor.StopProcessingAsync();
        Console.WriteLine("Processor stopped.");
    }
}

static async Task SendAndReceiveTopicMessageAsync(
    ServiceBusClient client,
    string topicName,
    string subscriptionName)
{
    await using ServiceBusSender topicSender = client.CreateSender(topicName);
    await topicSender.SendMessageAsync(
        new ServiceBusMessage("Message published to the topic"));
    Console.WriteLine("Sent one message to the topic.");

    await using ServiceBusReceiver subscriptionReceiver =
        client.CreateReceiver(topicName, subscriptionName);

    IReadOnlyList<ServiceBusReceivedMessage> messages =
        await subscriptionReceiver.ReceiveMessagesAsync(
            maxMessages: 1,
            maxWaitTime: TimeSpan.FromSeconds(10));

    foreach (ServiceBusReceivedMessage message in messages)
    {
        Console.WriteLine($"Subscription message: {message.Body}");
        await subscriptionReceiver.CompleteMessageAsync(message);
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.")
        : value;
}
