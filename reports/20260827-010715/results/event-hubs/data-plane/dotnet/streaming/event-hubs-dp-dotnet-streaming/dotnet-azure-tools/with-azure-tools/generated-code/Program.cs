using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Messaging.EventHubs.Producer;
using Azure.Storage.Blobs;

string eventHubsConnectionString = GetRequiredEnvironmentVariable(
    "EVENTHUB_CONNECTION_STRING");
string eventHubName = GetRequiredEnvironmentVariable("EVENTHUB_NAME");
string blobStorageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable("BLOB_CONTAINER_NAME");

using var shutdown = new CancellationTokenSource();
Console.CancelKeyPress += (_, eventArgs) =>
{
    eventArgs.Cancel = true;
    shutdown.Cancel();
};

await SendEventsAsync(
    eventHubsConnectionString,
    eventHubName,
    shutdown.Token);

var checkpointStore = new BlobContainerClient(
    blobStorageConnectionString,
    blobContainerName);
await checkpointStore.CreateIfNotExistsAsync(cancellationToken: shutdown.Token);

var processor = new EventProcessorClient(
    checkpointStore,
    EventHubConsumerClient.DefaultConsumerGroupName,
    eventHubsConnectionString,
    eventHubName);

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

try
{
    await processor.StartProcessingAsync(shutdown.Token);
    Console.WriteLine("Receiving events. Press Ctrl+C to stop.");

    try
    {
        await Task.Delay(Timeout.InfiniteTimeSpan, shutdown.Token);
    }
    catch (OperationCanceledException) when (shutdown.IsCancellationRequested)
    {
        // Expected when Ctrl+C requests a graceful shutdown.
    }
}
finally
{
    if (processor.IsRunning)
    {
        await processor.StopProcessingAsync(CancellationToken.None);
    }

    processor.ProcessEventAsync -= ProcessEventAsync;
    processor.ProcessErrorAsync -= ProcessErrorAsync;
}

static async Task SendEventsAsync(
    string connectionString,
    string eventHubName,
    CancellationToken cancellationToken)
{
    await using var producer = new EventHubProducerClient(
        connectionString,
        eventHubName);
    using EventDataBatch batch = await producer.CreateBatchAsync(cancellationToken);

    string runId = Guid.NewGuid().ToString("N");

    for (int eventNumber = 1; eventNumber <= 10; eventNumber++)
    {
        var eventData = new EventData($"Event {eventNumber} from run {runId}")
        {
            ContentType = "text/plain"
        };
        eventData.Properties["EventNumber"] = eventNumber;
        eventData.Properties["RunId"] = runId;
        eventData.Properties["Source"] = "EventHubsSample";

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {eventNumber} does not fit in the batch. " +
                "Send the current batch and create another batch for larger workloads.");
        }
    }

    await producer.SendAsync(batch, cancellationToken);
    Console.WriteLine($"Sent {batch.Count} events.");
}

static async Task ProcessEventAsync(ProcessEventArgs args)
{
    if (!args.HasEvent)
    {
        return;
    }

    Console.WriteLine(
        $"Partition {args.Partition.PartitionId}: {args.Data.EventBody}");

    foreach ((string key, object value) in args.Data.Properties)
    {
        Console.WriteLine($"  {key}: {value}");
    }

    // Checkpoint only after the event has been processed successfully.
    await args.UpdateCheckpointAsync(args.CancellationToken);
}

static Task ProcessErrorAsync(ProcessErrorEventArgs args)
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
