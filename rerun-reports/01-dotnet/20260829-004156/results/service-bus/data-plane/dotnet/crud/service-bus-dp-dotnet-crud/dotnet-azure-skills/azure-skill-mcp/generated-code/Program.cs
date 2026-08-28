using Azure.Messaging.ServiceBus;

internal static class Program
{
    private const string ConnectionStringVariable = "AZURE_SERVICE_BUS_CONNECTION_STRING";
    private const string QueueNameVariable = "AZURE_SERVICE_BUS_QUEUE_NAME";
    private const string TopicNameVariable = "AZURE_SERVICE_BUS_TOPIC_NAME";
    private const string SubscriptionNameVariable = "AZURE_SERVICE_BUS_SUBSCRIPTION_NAME";

    public static async Task<int> Main()
    {
        try
        {
            string connectionString = GetRequiredSetting(ConnectionStringVariable);
            string queueName = GetRequiredSetting(QueueNameVariable);
            string topicName = GetRequiredSetting(TopicNameVariable);
            string subscriptionName = GetRequiredSetting(SubscriptionNameVariable);

            await RunQueueDemoAsync(connectionString, queueName);
            await RunTopicDemoAsync(connectionString, topicName, subscriptionName);
            return 0;
        }
        catch (ServiceBusException exception)
        {
            Console.Error.WriteLine(
                $"Service Bus operation failed ({exception.Reason}): {exception.Message}");
            return 2;
        }
        catch (InvalidOperationException exception)
        {
            Console.Error.WriteLine(exception.Message);
            return 1;
        }
    }

    private static async Task RunQueueDemoAsync(string connectionString, string queueName)
    {
        await using var client = new ServiceBusClient(connectionString);
        await using ServiceBusSender sender = client.CreateSender(queueName);

        var singleMessage = new ServiceBusMessage("This is a single queue message.")
        {
            ContentType = "text/plain",
            MessageId = Guid.NewGuid().ToString()
        };

        await sender.SendMessageAsync(singleMessage);
        Console.WriteLine($"Sent one message to queue '{queueName}'.");

        using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();
        for (int index = 1; index <= 5; index++)
        {
            var message = new ServiceBusMessage($"Batch message {index}")
            {
                ContentType = "text/plain",
                MessageId = Guid.NewGuid().ToString()
            };

            if (!batch.TryAddMessage(message))
            {
                throw new InvalidOperationException(
                    $"Batch message {index} is too large for the current Service Bus batch.");
            }
        }

        await sender.SendMessagesAsync(batch);
        Console.WriteLine($"Sent a batch of {batch.Count} messages.");

        await using ServiceBusReceiver receiver = client.CreateReceiver(
            queueName,
            new ServiceBusReceiverOptions
            {
                ReceiveMode = ServiceBusReceiveMode.PeekLock
            });

        IReadOnlyList<ServiceBusReceivedMessage> receivedMessages =
            await receiver.ReceiveMessagesAsync(
                maxMessages: 6,
                maxWaitTime: TimeSpan.FromSeconds(10));

        foreach (ServiceBusReceivedMessage message in receivedMessages)
        {
            Console.WriteLine(
                $"Received queue message {message.MessageId}: {message.Body}");

            await receiver.CompleteMessageAsync(message);
            Console.WriteLine($"Completed queue message {message.MessageId}.");
        }

        await RunProcessorDemoAsync(client, sender, queueName);
    }

    private static async Task RunProcessorDemoAsync(
        ServiceBusClient client,
        ServiceBusSender sender,
        string queueName)
    {
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
            Console.WriteLine(
                $"Processor received {args.Message.MessageId}: {args.Message.Body}");
            await args.CompleteMessageAsync(args.Message);
            processed.TrySetResult();
        };

        processor.ProcessErrorAsync += args =>
        {
            Console.Error.WriteLine(
                $"Processor error from {args.ErrorSource} " +
                $"in namespace '{args.FullyQualifiedNamespace}': {args.Exception}");
            return Task.CompletedTask;
        };

        await processor.StartProcessingAsync();
        try
        {
            await sender.SendMessageAsync(
                new ServiceBusMessage("Process this queue message continuously.")
                {
                    ContentType = "text/plain",
                    MessageId = Guid.NewGuid().ToString()
                });

            await processed.Task.WaitAsync(TimeSpan.FromSeconds(30));
        }
        finally
        {
            await processor.StopProcessingAsync();
        }
    }

    private static async Task RunTopicDemoAsync(
        string connectionString,
        string topicName,
        string subscriptionName)
    {
        await using var client = new ServiceBusClient(connectionString);
        await using ServiceBusSender topicSender = client.CreateSender(topicName);
        await using ServiceBusReceiver subscriptionReceiver =
            client.CreateReceiver(topicName, subscriptionName);

        var topicMessage = new ServiceBusMessage("Hello from the topic.")
        {
            ContentType = "text/plain",
            MessageId = Guid.NewGuid().ToString(),
            Subject = "TopicDemo"
        };

        await topicSender.SendMessageAsync(topicMessage);
        Console.WriteLine($"Sent one message to topic '{topicName}'.");

        IReadOnlyList<ServiceBusReceivedMessage> subscriptionMessages =
            await subscriptionReceiver.ReceiveMessagesAsync(
                maxMessages: 1,
                maxWaitTime: TimeSpan.FromSeconds(10));

        foreach (ServiceBusReceivedMessage message in subscriptionMessages)
        {
            Console.WriteLine(
                $"Received from subscription '{subscriptionName}': {message.Body}");
            await subscriptionReceiver.CompleteMessageAsync(message);
        }
    }

    private static string GetRequiredSetting(string variableName)
    {
        string? value = Environment.GetEnvironmentVariable(variableName);
        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException(
                $"Set the required environment variable {variableName}.");
        }

        return value;
    }
}
