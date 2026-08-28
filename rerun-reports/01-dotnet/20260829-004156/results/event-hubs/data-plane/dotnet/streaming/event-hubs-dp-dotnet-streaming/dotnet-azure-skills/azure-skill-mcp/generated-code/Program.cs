using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Producer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Storage.Blobs;

string eventHubsConnectionString = GetRequiredSetting("EVENT_HUB_CONNECTION_STRING");
string eventHubName = GetRequiredSetting("EVENT_HUB_NAME");
string blobStorageConnectionString = GetRequiredSetting("BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredSetting("BLOB_CONTAINER_NAME");
string consumerGroup = Environment.GetEnvironmentVariable("EVENT_HUB_CONSUMER_GROUP")
    ?? EventHubConsumerClient.DefaultConsumerGroupName;

await SendEventsAsync(eventHubsConnectionString, eventHubName);

BlobContainerClient checkpointStore = new(
    blobStorageConnectionString,
    blobContainerName);

EventProcessorClient processor = new(
    checkpointStore,
    consumerGroup,
    eventHubsConnectionString,
    eventHubName);

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

TaskCompletionSource shutdown = new(TaskCreationOptions.RunContinuationsAsynchronously);
ConsoleCancelEventHandler cancelHandler = (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    shutdown.TrySetResult();
};

Console.CancelKeyPress += cancelHandler;

try
{
    await processor.StartProcessingAsync();
    Console.WriteLine("Processing events. Press Ctrl+C to stop.");
    await shutdown.Task;
}
finally
{
    Console.WriteLine("Stopping event processor...");
    await processor.StopProcessingAsync();
    Console.CancelKeyPress -= cancelHandler;
    processor.ProcessEventAsync -= ProcessEventAsync;
    processor.ProcessErrorAsync -= ProcessErrorAsync;
}

static async Task SendEventsAsync(string connectionString, string eventHubName)
{
    await using EventHubProducerClient producer = new(connectionString, eventHubName);
    using EventDataBatch batch = await producer.CreateBatchAsync();

    for (int eventNumber = 1; eventNumber <= 10; eventNumber++)
    {
        EventData eventData = new($"Event body {eventNumber}");
        eventData.Properties["EventNumber"] = eventNumber;
        eventData.Properties["Source"] = "EventHubsStreamingDemo";
        eventData.Properties["CreatedUtc"] = DateTimeOffset.UtcNow;

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {eventNumber} does not fit in the batch. " +
                "Send the current batch and create another before retrying.");
        }
    }

    await producer.SendAsync(batch);
    Console.WriteLine($"Sent {batch.Count} events.");
}

static async Task ProcessEventAsync(ProcessEventArgs eventArgs)
{
    if (!eventArgs.HasEvent)
    {
        return;
    }

    string body = eventArgs.Data.EventBody.ToString();
    Console.WriteLine(
        $"Received from partition {eventArgs.Partition.PartitionId}: {body}");

    foreach ((string key, object value) in eventArgs.Data.Properties)
    {
        Console.WriteLine($"  {key}: {value}");
    }

    // Checkpoint only after processing succeeds so a failure can be retried.
    await eventArgs.UpdateCheckpointAsync();
}

static Task ProcessErrorAsync(ProcessErrorEventArgs eventArgs)
{
    Console.Error.WriteLine(
        $"Processor error on partition {eventArgs.PartitionId ?? "<none>"} " +
        $"during {eventArgs.Operation}: {eventArgs.Exception}");

    return Task.CompletedTask;
}

static string GetRequiredSetting(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    return !string.IsNullOrWhiteSpace(value)
        ? value
        : throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.");
}
