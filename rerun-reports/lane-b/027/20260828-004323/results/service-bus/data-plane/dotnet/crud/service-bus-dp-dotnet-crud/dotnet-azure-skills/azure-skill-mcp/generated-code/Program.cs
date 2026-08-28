using Azure.Messaging.ServiceBus;

namespace ServiceBusMessagingDemo;

internal static class Program
{
    private static async Task Main()
    {
        string connectionString = GetRequiredEnvironmentVariable("SERVICE_BUS_CONNECTION_STRING");
        string queueName = GetRequiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
        string topicName = GetRequiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
        string subscriptionName = GetRequiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");

        await using ServiceBusClient client = new(connectionString);

        await RunQueueDemoAsync(client, queueName);
        await RunProcessorDemoAsync(client, queueName);
        await RunTopicSubscriptionDemoAsync(client, topicName, subscriptionName);
    }

    private static async Task RunQueueDemoAsync(ServiceBusClient client, string queueName)
    {
        await using ServiceBusSender sender = client.CreateSender(queueName);

        await sender.SendMessageAsync(new ServiceBusMessage("Single queue message"));
        Console.WriteLine("Sent one queue message.");

        using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();
        for (int messageNumber = 1; messageNumber <= 5; messageNumber++)
        {
            var message = new ServiceBusMessage($"Batch message {messageNumber}")
            {
                MessageId = Guid.NewGuid().ToString()
            };

            if (!batch.TryAddMessage(message))
            {
                throw new InvalidOperationException(
                    $"Message {messageNumber} is too large to fit in an empty Service Bus batch.");
            }
        }

        await sender.SendMessagesAsync(batch);
        Console.WriteLine("Sent a batch of five queue messages.");

        await using ServiceBusReceiver receiver = client.CreateReceiver(
            queueName,
            new ServiceBusReceiverOptions
            {
                ReceiveMode = ServiceBusReceiveMode.PeekLock
            });

        IReadOnlyList<ServiceBusReceivedMessage> messages =
            await receiver.ReceiveMessagesAsync(
                maxMessages: 6,
                maxWaitTime: TimeSpan.FromSeconds(30));

        foreach (ServiceBusReceivedMessage message in messages)
        {
            Console.WriteLine($"Received queue message: {message.Body}");
            await receiver.CompleteMessageAsync(message);
        }

        Console.WriteLine($"Completed {messages.Count} queue message(s).");
    }

    private static async Task RunProcessorDemoAsync(ServiceBusClient client, string queueName)
    {
        var messageProcessed = new TaskCompletionSource(
            TaskCreationOptions.RunContinuationsAsynchronously);

        await using ServiceBusProcessor processor = client.CreateProcessor(
            queueName,
            new ServiceBusProcessorOptions
            {
                AutoCompleteMessages = false,
                MaxConcurrentCalls = 1
            });

        processor.ProcessMessageAsync += async args =>
        {
            Console.WriteLine($"Processor received: {args.Message.Body}");
            await args.CompleteMessageAsync(args.Message);
            messageProcessed.TrySetResult();
        };

        processor.ProcessErrorAsync += args =>
        {
            Console.Error.WriteLine(
                $"Processor error ({args.ErrorSource}, {args.EntityPath}): {args.Exception}");
            messageProcessed.TrySetException(args.Exception);
            return Task.CompletedTask;
        };

        await processor.StartProcessingAsync();
        try
        {
            await using ServiceBusSender sender = client.CreateSender(queueName);
            await sender.SendMessageAsync(new ServiceBusMessage("Message for the processor"));

            await messageProcessed.Task.WaitAsync(TimeSpan.FromSeconds(30));
        }
        finally
        {
            await processor.StopProcessingAsync();
        }
    }

    private static async Task RunTopicSubscriptionDemoAsync(
        ServiceBusClient client,
        string topicName,
        string subscriptionName)
    {
        await using ServiceBusSender topicSender = client.CreateSender(topicName);
        await topicSender.SendMessageAsync(new ServiceBusMessage("Topic message"));
        Console.WriteLine("Sent one topic message.");

        await using ServiceBusReceiver subscriptionReceiver =
            client.CreateReceiver(topicName, subscriptionName);

        IReadOnlyList<ServiceBusReceivedMessage> messages =
            await subscriptionReceiver.ReceiveMessagesAsync(
                maxMessages: 1,
                maxWaitTime: TimeSpan.FromSeconds(30));

        if (messages.Count == 0)
        {
            throw new TimeoutException(
                $"No message arrived on subscription '{subscriptionName}' within 30 seconds.");
        }

        ServiceBusReceivedMessage message = messages[0];
        Console.WriteLine($"Received subscription message: {message.Body}");
        await subscriptionReceiver.CompleteMessageAsync(message);
    }

    private static string GetRequiredEnvironmentVariable(string name)
    {
        string? value = Environment.GetEnvironmentVariable(name);
        return !string.IsNullOrWhiteSpace(value)
            ? value
            : throw new InvalidOperationException(
                $"Set the {name} environment variable before running the sample.");
    }
}
