using System.Collections.Concurrent;
using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Messaging.EventHubs.Producer;
using Azure.Storage.Blobs;

const int eventCount = 10;
const int checkpointFrequency = 5;

string eventHubsConnectionString = GetRequiredEnvironmentVariable(
    "EVENT_HUBS_CONNECTION_STRING");
string blobStorageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable(
    "BLOB_CONTAINER_NAME");
string consumerGroup = Environment.GetEnvironmentVariable(
    "EVENT_HUB_CONSUMER_GROUP") ?? EventHubConsumerClient.DefaultConsumerGroupName;

await SendEventsAsync(eventHubsConnectionString);

BlobContainerClient checkpointStore = new(
    blobStorageConnectionString,
    blobContainerName);

EventProcessorClient processor = new(
    checkpointStore,
    consumerGroup,
    eventHubsConnectionString);

ConcurrentDictionary<string, int> eventsSinceCheckpoint = new();

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

using CancellationTokenSource cancellationSource = new();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

try
{
    await processor.StartProcessingAsync(cancellationSource.Token);
    Console.WriteLine("Processing events. Press Ctrl+C to stop.");

    await Task.Delay(Timeout.InfiniteTimeSpan, cancellationSource.Token);
}
catch (OperationCanceledException) when (cancellationSource.IsCancellationRequested)
{
    Console.WriteLine("Stopping event processor.");
}
finally
{
    if (processor.IsRunning)
    {
        await processor.StopProcessingAsync();
    }

    processor.ProcessEventAsync -= ProcessEventAsync;
    processor.ProcessErrorAsync -= ProcessErrorAsync;
}

async Task SendEventsAsync(string connectionString)
{
    await using EventHubProducerClient producer = new(connectionString);
    using EventDataBatch batch = await producer.CreateBatchAsync();

    for (int index = 1; index <= eventCount; index++)
    {
        EventData eventData = new(BinaryData.FromString($"Event body {index}"));
        eventData.Properties["eventNumber"] = index;
        eventData.Properties["source"] = "EventHubsSample";
        eventData.Properties["createdUtc"] = DateTimeOffset.UtcNow.ToString("O");

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {index} is too large to fit in the Event Hubs batch.");
        }
    }

    await producer.SendAsync(batch);
    Console.WriteLine($"Sent {batch.Count} events.");
}

async Task ProcessEventAsync(ProcessEventArgs eventArgs)
{
    if (!eventArgs.HasEvent)
    {
        return;
    }

    string body = eventArgs.Data.EventBody.ToString();
    Console.WriteLine(
        $"Partition {eventArgs.Partition.PartitionId}: {body}");

    int processedCount = eventsSinceCheckpoint.AddOrUpdate(
        eventArgs.Partition.PartitionId,
        1,
        (_, count) => count + 1);

    if (processedCount >= checkpointFrequency)
    {
        await eventArgs.UpdateCheckpointAsync(eventArgs.CancellationToken);
        eventsSinceCheckpoint[eventArgs.Partition.PartitionId] = 0;
        Console.WriteLine(
            $"Checkpoint updated for partition {eventArgs.Partition.PartitionId}.");
    }
}

Task ProcessErrorAsync(ProcessErrorEventArgs eventArgs)
{
    Console.Error.WriteLine(
        $"Error in partition {eventArgs.PartitionId ?? "(processor)"} " +
        $"during {eventArgs.Operation}: {eventArgs.Exception}");

    return Task.CompletedTask;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.");
    }

    return value;
}
