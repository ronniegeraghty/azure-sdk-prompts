using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Messaging.EventHubs.Producer;
using Azure.Storage.Blobs;

const int eventCount = 10;

string eventHubsConnectionString = GetRequiredEnvironmentVariable(
    "EVENTHUB_CONNECTION_STRING");
string eventHubName = GetRequiredEnvironmentVariable("EVENTHUB_NAME");
string storageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable(
    "BLOB_CONTAINER_NAME");

using var cancellationSource = new CancellationTokenSource();
Console.CancelKeyPress += (_, args) =>
{
    args.Cancel = true;
    cancellationSource.Cancel();
};

var blobContainerClient = new BlobContainerClient(
    storageConnectionString,
    blobContainerName);
await blobContainerClient.CreateIfNotExistsAsync(
    cancellationToken: cancellationSource.Token);

string runId = Guid.NewGuid().ToString("N");

await using (var producer = new EventHubProducerClient(
    eventHubsConnectionString,
    eventHubName))
{
    using EventDataBatch batch = await producer.CreateBatchAsync(
        cancellationSource.Token);

    for (int i = 1; i <= eventCount; i++)
    {
        var eventData = new EventData($"Event body {i}")
        {
            ContentType = "text/plain"
        };
        eventData.Properties["EventNumber"] = i;
        eventData.Properties["SampleRunId"] = runId;
        eventData.Properties["SentAtUtc"] = DateTimeOffset.UtcNow.ToString("O");

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {i} is too large to fit in the Event Hubs batch.");
        }
    }

    await producer.SendAsync(batch, cancellationSource.Token);
    Console.WriteLine($"Sent {batch.Count} events.");
}

// A stored checkpoint takes precedence. Without one, EventProcessorClient
// starts at the beginning of each partition, including the events sent above.
var processor = new EventProcessorClient(
    blobContainerClient,
    EventHubConsumerClient.DefaultConsumerGroupName,
    eventHubsConnectionString,
    eventHubName);

int receivedSampleEvents = 0;
var allSampleEventsReceived = new TaskCompletionSource(
    TaskCreationOptions.RunContinuationsAsynchronously);

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

try
{
    await processor.StartProcessingAsync(cancellationSource.Token);
    Console.WriteLine("Processing events. Press Ctrl+C to stop.");

    await Task.WhenAny(
        allSampleEventsReceived.Task,
        Task.Delay(Timeout.InfiniteTimeSpan, cancellationSource.Token));
}
catch (OperationCanceledException) when (cancellationSource.IsCancellationRequested)
{
    Console.WriteLine("Cancellation requested.");
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

async Task ProcessEventAsync(ProcessEventArgs args)
{
    if (args.CancellationToken.IsCancellationRequested)
    {
        return;
    }

    string body = args.Data.EventBody.ToString();
    Console.WriteLine(
        $"Partition {args.Partition.PartitionId}, " +
        $"sequence {args.Data.SequenceNumber}: {body}");

    // Checkpoint only after successful processing. For higher throughput,
    // checkpoint after a count or time interval instead of every event.
    await args.UpdateCheckpointAsync(args.CancellationToken);

    if (args.Data.Properties.TryGetValue("SampleRunId", out object? value) &&
        string.Equals(value?.ToString(), runId, StringComparison.Ordinal) &&
        Interlocked.Increment(ref receivedSampleEvents) == eventCount)
    {
        allSampleEventsReceived.TrySetResult();
    }
}

Task ProcessErrorAsync(ProcessErrorEventArgs args)
{
    Console.Error.WriteLine(
        $"Event processor error. Operation: {args.Operation}; " +
        $"Partition: {args.PartitionId ?? "<none>"}; " +
        $"Exception: {args.Exception}");
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
