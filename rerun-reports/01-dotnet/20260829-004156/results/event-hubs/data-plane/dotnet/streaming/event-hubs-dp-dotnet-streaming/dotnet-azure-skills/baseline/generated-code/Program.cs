using Azure.Messaging.EventHubs;
using Azure.Messaging.EventHubs.Consumer;
using Azure.Messaging.EventHubs.Producer;
using Azure.Messaging.EventHubs.Processor;
using Azure.Storage.Blobs;
using System.Text;

string eventHubsConnectionString = GetRequiredEnvironmentVariable(
    "EVENT_HUBS_CONNECTION_STRING");
string eventHubName = GetRequiredEnvironmentVariable("EVENT_HUB_NAME");
string blobStorageConnectionString = GetRequiredEnvironmentVariable(
    "BLOB_STORAGE_CONNECTION_STRING");
string blobContainerName = GetRequiredEnvironmentVariable(
    "BLOB_CONTAINER_NAME");

await using (var producer = new EventHubProducerClient(
    eventHubsConnectionString,
    eventHubName))
{
    using EventDataBatch batch = await producer.CreateBatchAsync();

    for (int eventNumber = 1; eventNumber <= 10; eventNumber++)
    {
        var eventData = new EventData(
            Encoding.UTF8.GetBytes($"Event body {eventNumber}"));

        eventData.Properties["EventNumber"] = eventNumber;
        eventData.Properties["Source"] = "EventHubsSample";
        eventData.Properties["CreatedUtc"] = DateTimeOffset.UtcNow;

        if (!batch.TryAdd(eventData))
        {
            throw new InvalidOperationException(
                $"Event {eventNumber} is too large for the current batch.");
        }
    }

    await producer.SendAsync(batch);
    Console.WriteLine($"Sent {batch.Count} events.");
}

var checkpointStore = new BlobContainerClient(
    blobStorageConnectionString,
    blobContainerName);

var processor = new EventProcessorClient(
    checkpointStore,
    EventHubConsumerClient.DefaultConsumerGroupName,
    eventHubsConnectionString,
    eventHubName);

processor.ProcessEventAsync += ProcessEventAsync;
processor.ProcessErrorAsync += ProcessErrorAsync;

using var cancellationSource = new CancellationTokenSource();
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
    Console.WriteLine("Stopping event processing.");
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

static async Task ProcessEventAsync(ProcessEventArgs eventArgs)
{
    if (eventArgs.CancellationToken.IsCancellationRequested)
    {
        return;
    }

    if (eventArgs.Data is null)
    {
        return;
    }

    string body = Encoding.UTF8.GetString(eventArgs.Data.EventBody.ToArray());
    Console.WriteLine(
        $"Partition {eventArgs.Partition.PartitionId}: {body}");

    foreach ((string key, object value) in eventArgs.Data.Properties)
    {
        Console.WriteLine($"  {key}: {value}");
    }

    // The checkpoint advances only after the event has been handled successfully.
    await eventArgs.UpdateCheckpointAsync(eventArgs.CancellationToken);
}

static Task ProcessErrorAsync(ProcessErrorEventArgs eventArgs)
{
    Console.Error.WriteLine(
        $"Error in operation '{eventArgs.Operation}' on partition " +
        $"'{eventArgs.PartitionId ?? "N/A"}': {eventArgs.Exception}");

    return Task.CompletedTask;
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException(
            $"Set the required environment variable '{name}'.")
        : value;
}
