using Azure.Messaging.ServiceBus;

string connectionString = GetRequiredSetting("AZURE_SERVICE_BUS_CONNECTION_STRING");
string queueName = GetRequiredSetting("AZURE_SERVICE_BUS_QUEUE_NAME");
string topicName = GetRequiredSetting("AZURE_SERVICE_BUS_TOPIC_NAME");
string subscriptionName = GetRequiredSetting("AZURE_SERVICE_BUS_SUBSCRIPTION_NAME");

await using var client = new ServiceBusClient(connectionString);

// Send one queue message, followed by a batch of five messages.
await using (ServiceBusSender queueSender = client.CreateSender(queueName))
{
    await queueSender.SendMessageAsync(
        new ServiceBusMessage("Single queue message")
        {
            ContentType = "text/plain",
            MessageId = Guid.NewGuid().ToString()
        });

    using ServiceBusMessageBatch batch = await queueSender.CreateMessageBatchAsync();

    for (int i = 1; i <= 5; i++)
    {
        var message = new ServiceBusMessage($"Batch message {i}")
        {
            MessageId = Guid.NewGuid().ToString(),
            ApplicationProperties = { ["Sequence"] = i }
        };

        if (!batch.TryAddMessage(message))
        {
            throw new InvalidOperationException(
                $"Message {i} is too large to fit in an empty Service Bus batch.");
        }
    }

    await queueSender.SendMessagesAsync(batch);
    Console.WriteLine("Sent one message and a batch of five messages to the queue.");
}

// Receive queue messages explicitly and settle each one after processing.
await using (ServiceBusReceiver queueReceiver = client.CreateReceiver(
    queueName,
    new ServiceBusReceiverOptions
    {
        ReceiveMode = ServiceBusReceiveMode.PeekLock
    }))
{
    IReadOnlyList<ServiceBusReceivedMessage> messages =
        await queueReceiver.ReceiveMessagesAsync(
            maxMessages: 6,
            maxWaitTime: TimeSpan.FromSeconds(10));

    foreach (ServiceBusReceivedMessage message in messages)
    {
        Console.WriteLine($"Received queue message: {message.Body}");

        // Complete only after processing succeeds.
        await queueReceiver.CompleteMessageAsync(message);
    }
}

// Continuously process queue messages with callback handlers.
var processedByHandler = new TaskCompletionSource(
    TaskCreationOptions.RunContinuationsAsynchronously);

await using (ServiceBusProcessor processor = client.CreateProcessor(
    queueName,
    new ServiceBusProcessorOptions
    {
        AutoCompleteMessages = false,
        MaxConcurrentCalls = 1
    }))
{
    processor.ProcessMessageAsync += async args =>
    {
        Console.WriteLine($"Processor received: {args.Message.Body}");
        await args.CompleteMessageAsync(args.Message, args.CancellationToken);
        processedByHandler.TrySetResult();
    };

    processor.ProcessErrorAsync += args =>
    {
        Console.Error.WriteLine(
            $"Processor error ({args.ErrorSource}, {args.EntityPath}): {args.Exception}");
        return Task.CompletedTask;
    };

    await processor.StartProcessingAsync();

    try
    {
        await using ServiceBusSender processorDemoSender = client.CreateSender(queueName);
        await processorDemoSender.SendMessageAsync(
            new ServiceBusMessage("Message for the continuous processor"));

        await processedByHandler.Task.WaitAsync(TimeSpan.FromSeconds(30));
    }
    finally
    {
        await processor.StopProcessingAsync();
    }
}

// Send to a topic, then receive and complete the copy from a subscription.
await using (ServiceBusSender topicSender = client.CreateSender(topicName))
{
    await topicSender.SendMessageAsync(
        new ServiceBusMessage("Topic message")
        {
            Subject = "ServiceBusDemo"
        });
}

await using (ServiceBusReceiver subscriptionReceiver =
    client.CreateReceiver(topicName, subscriptionName))
{
    ServiceBusReceivedMessage? topicMessage =
        await subscriptionReceiver.ReceiveMessageAsync(TimeSpan.FromSeconds(10));

    if (topicMessage is null)
    {
        Console.WriteLine("No topic message was received within the timeout.");
    }
    else
    {
        Console.WriteLine($"Received subscription message: {topicMessage.Body}");
        await subscriptionReceiver.CompleteMessageAsync(topicMessage);
    }
}

static string GetRequiredSetting(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException(
        $"Set the required environment variable '{name}' before running the sample.");
