using System.Text.Json;
using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Messaging.EventHubs.Producer;
using Azure.Storage.Blobs;

const int EventCount = 10;

string eventHubConnectionString = GetRequiredEnvironmentVariable(
    "EVENT_HUB_CONNECTION_STRING");
string eventHubName = GetRequiredEnvironmentVariable("EVENT_HUB_NAME");
string blobStorageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable(
    "BLOB_CONTAINER_NAME");

string runId = Guid.NewGuid().ToString("N");

await using (var producer = new EventHubProducerClient(
    eventHubConnectionString,
    eventHubName))
{
    using EventDataBatch batch = await producer.CreateBatchAsync();

    for (int number = 1; number <= EventCount; number++)
    {
        var body = new
        {
            Number = number,
            Message = $"Event {number}",
            SentAtUtc = DateTimeOffset.UtcNow
        };

        var eventData = new EventData(
            BinaryData.FromString(JsonSerializer.Serialize(body)));

        eventData.Properties["RunId"] = runId;
        eventData.Properties["MessageNumber"] = number;
        eventData.Properties["Source"] = "EventHubsStreamingSample";

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {number} is too large to fit in the batch.");
        }
    }

    await producer.SendAsync(batch);
    Console.WriteLine($"Sent {batch.Count} events.");
}

// The checkpoint container must already exist.
var checkpointStore = new BlobContainerClient(
    blobStorageConnectionString,
    blobContainerName);

var processor = new EventProcessorClient(
    checkpointStore,
    EventHubConsumerClient.DefaultConsumerGroupName,
    eventHubConnectionString,
    eventHubName);

using var cancellationSource = new CancellationTokenSource();
var receivedAllEvents = new TaskCompletionSource(
    TaskCreationOptions.RunContinuationsAsynchronously);
int receivedEventCount = 0;

async Task ProcessEventAsync(ProcessEventArgs args)
{
    if (!args.HasEvent)
    {
        return;
    }

    Console.WriteLine(
        $"Received from partition {args.Partition.PartitionId}: " +
        args.Data.EventBody.ToString());

    // Checkpoint only after the event has been processed successfully.
    await args.UpdateCheckpointAsync(cancellationSource.Token);

    if (args.Data.Properties.TryGetValue("RunId", out object? value) &&
        string.Equals(value?.ToString(), runId, StringComparison.Ordinal) &&
        Interlocked.Increment(ref receivedEventCount) >= EventCount)
    {
        receivedAllEvents.TrySetResult();
    }
}

Task ProcessErrorAsync(ProcessErrorEventArgs args)
{
    Console.Error.WriteLine(
        $"Processor error. Operation: {args.Operation}; " +
        $"Partition: {args.PartitionId ?? "<none>"}; " +
        $"Exception: {args.Exception.Message}");

    return Task.CompletedTask;
}

ConsoleCancelEventHandler cancelHandler = (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    cancellationSource.Cancel();
};

Console.CancelKeyPress += cancelHandler;
processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

try
{
    await processor.StartProcessingAsync(cancellationSource.Token);
    Console.WriteLine("Processing events. Press Ctrl+C to stop.");

    Task cancellationTask = Task.Delay(
        Timeout.InfiniteTimeSpan,
        cancellationSource.Token);

    await Task.WhenAny(receivedAllEvents.Task, cancellationTask);
}
finally
{
    if (processor.IsRunning)
    {
        await processor.StopProcessingAsync();
    }

    processor.ProcessEventAsync -= ProcessEventAsync;
    processor.ProcessErrorAsync -= ProcessErrorAsync;
    Console.CancelKeyPress -= cancelHandler;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException(
            $"Set the {name} environment variable before running the sample.")
        : value;
}
