using Azure.Messaging.ServiceBus;

const string connectionStringVariable = "AZURE_SERVICE_BUS_CONNECTION_STRING";
const string queueNameVariable = "AZURE_SERVICE_BUS_QUEUE_NAME";
const string topicNameVariable = "AZURE_SERVICE_BUS_TOPIC_NAME";
const string subscriptionNameVariable = "AZURE_SERVICE_BUS_SUBSCRIPTION_NAME";

string connectionString = GetRequiredEnvironmentVariable(connectionStringVariable);
string queueName = GetRequiredEnvironmentVariable(queueNameVariable);
string topicName = GetRequiredEnvironmentVariable(topicNameVariable);
string subscriptionName = GetRequiredEnvironmentVariable(subscriptionNameVariable);

await using var client = new ServiceBusClient(connectionString);

// Send one queue message and then a batch of five queue messages.
await using (ServiceBusSender queueSender = client.CreateSender(queueName))
{
    await queueSender.SendMessageAsync(
        new ServiceBusMessage("Single queue message")
        {
            ContentType = "text/plain",
            Subject = "single"
        });

    using ServiceBusMessageBatch batch = await queueSender.CreateMessageBatchAsync();

    for (int messageNumber = 1; messageNumber <= 5; messageNumber++)
    {
        var message = new ServiceBusMessage($"Batch queue message {messageNumber}")
        {
            ContentType = "text/plain",
            Subject = "batch",
            MessageId = Guid.NewGuid().ToString()
        };

        if (!batch.TryAddMessage(message))
        {
            throw new InvalidOperationException(
                $"Message {messageNumber} is too large to fit in an empty Service Bus batch.");
        }
    }

    await queueSender.SendMessagesAsync(batch);
    Console.WriteLine("Sent one message and a batch of five messages to the queue.");
}

// Receive queue messages in PeekLock mode and settle each one after processing.
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
        Console.WriteLine($"Queue receiver processed: {message.Body}");
        await queueReceiver.CompleteMessageAsync(message);
    }
}

// Continuously process queue messages with success and error handlers.
var processorOptions = new ServiceBusProcessorOptions
{
    AutoCompleteMessages = false,
    MaxConcurrentCalls = 1,
    ReceiveMode = ServiceBusReceiveMode.PeekLock
};

await using (ServiceBusProcessor processor =
             client.CreateProcessor(queueName, processorOptions))
{
    processor.ProcessMessageAsync += ProcessQueueMessageAsync;
    processor.ProcessErrorAsync += ProcessErrorAsync;

    await processor.StartProcessingAsync();
    Console.WriteLine("Queue processor is running. Press Enter to stop.");
    Console.ReadLine();
    await processor.StopProcessingAsync();
}

// Send to a topic.
await using (ServiceBusSender topicSender = client.CreateSender(topicName))
{
    await topicSender.SendMessageAsync(
        new ServiceBusMessage("Topic message")
        {
            ContentType = "text/plain",
            Subject = "topic-demo"
        });

    Console.WriteLine("Sent one message to the topic.");
}

// Receive the topic message from a subscription and complete it.
await using (ServiceBusReceiver subscriptionReceiver =
             client.CreateReceiver(topicName, subscriptionName))
{
    ServiceBusReceivedMessage? message =
        await subscriptionReceiver.ReceiveMessageAsync(TimeSpan.FromSeconds(10));

    if (message is null)
    {
        Console.WriteLine("No topic message was available in the subscription.");
    }
    else
    {
        Console.WriteLine($"Subscription receiver processed: {message.Body}");
        await subscriptionReceiver.CompleteMessageAsync(message);
    }
}

async Task ProcessQueueMessageAsync(ProcessMessageEventArgs args)
{
    Console.WriteLine($"Processor handled: {args.Message.Body}");
    await args.CompleteMessageAsync(args.Message);
}

Task ProcessErrorAsync(ProcessErrorEventArgs args)
{
    Console.Error.WriteLine(
        $"Processor error ({args.ErrorSource}) for {args.EntityPath}: {args.Exception}");
    return Task.CompletedTask;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException(
            $"Set the required environment variable '{name}' before running the sample.");
}
