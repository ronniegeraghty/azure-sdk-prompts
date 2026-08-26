using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Producer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Storage.Blobs;

const string consumerGroup = EventHubConsumerClient.DefaultConsumerGroupName;

string eventHubConnectionString = GetRequiredEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING");
string eventHubName = GetRequiredEnvironmentVariable("EVENT_HUB_NAME");
string blobStorageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable("BLOB_CONTAINER_NAME");

await SendEventsAsync(eventHubConnectionString, eventHubName);

var checkpointStore = new BlobContainerClient(
    blobStorageConnectionString,
    blobContainerName);

var processor = new EventProcessorClient(
    checkpointStore,
    consumerGroup,
    eventHubConnectionString,
    eventHubName);

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

using var cancellationSource = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

Console.WriteLine("Starting the event processor. Press Ctrl+C to stop.");

try
{
    await processor.StartProcessingAsync(cancellationSource.Token);
    await Task.Delay(Timeout.InfiniteTimeSpan, cancellationSource.Token);
}
catch (OperationCanceledException) when (cancellationSource.IsCancellationRequested)
{
    Console.WriteLine("Stopping the event processor.");
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

static async Task SendEventsAsync(
    string connectionString,
    string eventHubName)
{
    await using var producer = new EventHubProducerClient(
        connectionString,
        eventHubName);
    using EventDataBatch batch = await producer.CreateBatchAsync();

    for (int index = 1; index <= 10; index++)
    {
        var eventData = new EventData($"Event body {index}");
        eventData.Properties["EventNumber"] = index;
        eventData.Properties["CreatedUtc"] = DateTimeOffset.UtcNow;
        eventData.Properties["Source"] = "CSharpEventHubsSample";

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {index} is too large to fit in the batch.");
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

    // Checkpoint only after the event has been processed successfully.
    await eventArgs.UpdateCheckpointAsync(eventArgs.CancellationToken);
}

static Task ProcessErrorAsync(ProcessErrorEventArgs eventArgs)
{
    Console.Error.WriteLine(
        $"Processing error on partition {eventArgs.PartitionId ?? "(none)"} " +
        $"during {eventArgs.Operation}: {eventArgs.Exception.Message}");
    return Task.CompletedTask;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Set the {name} environment variable before running the sample.");
    }

    return value;
}
