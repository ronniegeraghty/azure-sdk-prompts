# Evaluation Report: event-hubs-dp-dotnet-streaming

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 316.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `event-hubs-dp-dotnet-streaming` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 316.5s |
| Timestamp | 2026-08-27T14:28:54Z |
| Files Generated | 2 |
| Event Count | 3458 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 282.8s |
| Review | 32.5s |
| **Total** | **316.5s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=297233 out=3354 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 3638ms

**Input:**

```json
{"skill":"azure-messaging"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Messaging SDK Troubleshooting

## Quick Reference

| Property | Value |
|----------|-------|
| **Services** | Azure Event Hubs, Azure Service Bus |
| **MCP Tools** | `mcp_azure_mcp_eventhubs`, `mcp_azure_mcp_servicebus` |
| **Best For** | Diagnosing SDK connection, auth, and message processing issues |

## When to Use This Skill

- SDK connection failures, auth errors, or AMQP link errors
- Idle timeout, connection inactivity, or slow reconnection after disconnect
- AMQP link detach or detach-forced errors
- Message lock lost, message lock expired, lock renewal failures, or batch lock timeouts
- Session lock lost, session lock expired, or session receiver errors
- Event processor or message handler stops processing
- Duplicate events or checkpoint offset resets
- SDK configuration questions (retry, prefetch, batch size, receive batch behavior)

## MCP Tools

| Tool | Command | Use |
|------|---------|-----|
| `mcp_azure_mcp_eventhubs` | Namespace/hub ops | List namespaces, hubs, consumer groups |
| `mcp_azure_mcp_servicebus` | Queue/topic ops | List namespaces, queues, topics, subscriptions |
| `mcp_azure_mcp_monitor` | `logs_query` | Query diagnostic logs with KQL |
| `mcp_azure_mcp_resourcehealth` | `get` | Check service health status |
| `mcp_azure_mcp_documentation` | Doc search | Search Microsoft Learn for troubleshooting docs |

## Diagnosis Workflow

1. **Identify the SDK and version** — Check the prompt for SDK and version clues; if not stated, proceed with diagnosis and ask later if needed
2. **Check resource health** — Use `mcp_azure_mcp_resourcehealth` to verify the namespace is healthy
3. **Review the error message** — Match against language-specific troubleshooting guide
4. **Look up documentation** — Use `mcp_azure_mcp_documentation` to search Microsoft Learn for the error or topic
5. **Check configuration** — Verify connection string, entity name, consumer group
6. **Recommend fix** — Apply remediatio
... (truncated)
```

### ❌ 2. view — 7409ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 7406ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. glob — 7404ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. view — 186304ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
@@ -1,0 +1,1 @@
 SKILL.md

```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 186307ms

**Input:**

```json
{"intent":"Generate a local-only C# console sample using Azure Event Hubs producer batching and EventProcessorClient with Blob checkpointing, connection strings from environment variables, async disposal, handlers, and checkpoint updates.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-documentation (via azure) — 186309ms

**Input:**

```json
{"intent":"Find current official .NET guidance and API examples for EventHubProducerClient CreateBatchAsync, EventProcessorClient with BlobContainerClient, ProcessEventAsync/ProcessErrorAsync handlers, UpdateCheckpointAsync, and disposal.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. apply_patch — 798ms

**Input:**

```json
"*** Begin Patch\n*** Add File: EventHubsSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Messaging.EventHubs\" Version=\"5.12.2\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Messaging.EventHubs.Processor\" Version=\"5.12.2\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Storage.Blobs\" Version=\"12.25.1\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Collections.Concurrent;\n+using Azure.Messaging.EventHubs;\n+using Azure.Messaging.EventHubs.Consumer;\n+using Azure.Messaging.EventHubs.Processor;\n+using Azure.Messaging.EventHubs.Producer;\n+using Azure.Storage.Blobs;\n+\n+const int eventCount = 10;\n+const int checkpointFrequency = 5;\n+\n+string eventHubsConnectionString = GetRequiredEnvironmentVariable(\n+    \"EVENT_HUBS_CONNECTION_STRING\");\n+string blobStorageConnectionString = GetRequiredEnvironmentVariable(\n+    \"BLOB_STORAGE_CONNECTION_STRING\");\n+string blobContainerName = GetRequiredEnvironmentVariable(\n+    \"BLOB_CONTAINER_NAME\");\n+string consumerGroup = Environment.GetEnvironmentVariable(\n+    \"EVENT_HUB_CONSUMER_GROUP\") ?? EventHubConsumerClient.DefaultConsumerGroupName;\n+\n+await SendEventsAsync(eventHubsConnectionString);\n+\n+BlobContainerClient checkpointStore = new(\n+    blobStorageConnectionString,\n+    blobContainerName);\n+\n+EventProcessorClient processor = new(\n+    checkpointStore,\n+    consumerGroup,\n+    eventHubsConnectionString,\n+    new EventProcessorClientOptions\n+    {\n+        DefaultStartingPosition = EventPosition.Earliest\n+    });\n+\n+ConcurrentDictionary\u003cstring, int\u003e eventsSinceCheckpoint = new();\n+\n+processor.ProcessEventAsync += ProcessEventAsync;\n+processor.ProcessErrorAsync += ProcessErrorAsync;\n+\n+using CancellationTokenSource cancellationSource = new();\n+Console.CancelKeyPress += (_, eventArgs) =\u003e\n+{\n+    eventArgs.Cancel = true;\n+    cancellationSource.Cancel();\n+};\n+\n+try\n+{\n+    await processor.StartProcessingAsync(cancellationSource.Token);\n+    Console.WriteLine(\"Processing events. Press Ctrl+C to stop.\");\n+\n+    await Task.Delay(Timeout.InfiniteTimeSpan, cancellationSource.Token);\n+}\n+catch (OperationCanceledException) when (cancellationSource.IsCancellationRequested)\n+{\n+    Console.WriteLine(\"Stopping event processor.\");\n+}\n+finally\n+{\n+    if (processor.IsRunning)\n+    {\n+        await processor.StopProcessingAsync();\n+    }\n+\n+    processor.ProcessEventAsync -= ProcessEventAsync;\n+    processor.ProcessErrorAsync -= ProcessErrorAsync;\n+}\n+\n+async Task SendEventsAsync(string connectionString)\n+{\n+    await using EventHubProducerClient producer = new(connectionString);\n+    using EventDataBatch batch = await producer.CreateBatchAsync();\n+\n+    for (int index = 1; index \u003c= eventCount; index++)\n+    {\n+        EventData eventData = new(BinaryData.FromString($\"Event body {index}\"));\n+        eventData.Properties[\"eventNumber\"] = index;\n+        eventData.Properties[\"source\"] = \"EventHubsSample\";\n+        eventData.Properties[\"createdUtc\"] = DateTimeOffset.UtcNow.ToString(\"O\");\n+\n+        if (!batch.TryAdd(eventData))\n+        {\n+            throw new InvalidOperationException(\n+                $\"Event {index} is too large to fit in the Event Hubs batch.\");\n+        }\n+    }\n+\n+    await producer.SendAsync(batch);\n+    Console.WriteLine($\"Sent {batch.Count} events.\");\n+}\n+\n+async Task ProcessEventAsync(ProcessEventArgs eventArgs)\n+{\n+    if (!eventArgs.HasEvent)\n+    {\n+        return;\n+    }\n+\n+    string body = eventArgs.Data.EventBody.ToString();\n+    Console.WriteLine(\n+        $\"Partition {eventArgs.Partition.PartitionId}: {body}\");\n+\n+    int processedCount = eventsSinceCheckpoint.AddOrUpdate(\n+        eventArgs.Partition.PartitionId,\n+        1,\n+        (_, count) =\u003e count + 1);\n+\n+    if (processedCount \u003e= checkpointFrequency)\n+    {\n+        await eventArgs.UpdateCheckpointAsync(eventArgs.CancellationToken);\n+        eventsSinceCheckpoint[eventArgs.Partition.PartitionId] = 0;\n+        Console.WriteLine(\n+            $\"Checkpoint updated for partition {eventArgs.Partition.PartitionId}.\");\n+    }\n+}\n+\n+Task ProcessErrorAsync(ProcessErrorEventArgs eventArgs)\n+{\n+    Console.Error.WriteLine(\n+        $\"Error in partition {eventArgs.PartitionId ?? \"(processor)\"} \" +\n+        $\"during {eventArgs.Operation}: {eventArgs.Exception}\");\n+\n+    return Task.CompletedTask;\n+}\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+\n+    if (string.IsNullOrWhiteSpace(value))\n+    {\n+        throw new InvalidOperationException(\n+            $\"Set the required environment variable '{name}'.\");\n+    }\n+\n+    return value;\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/EventHubsSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/EventHubsSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/EventHubsSample.csproj
@@ -1,0 +1,17 @@
+<Project Sdk="Microsoft.NET.Sdk">
+
+  <PropertyGroup>
+    <OutputType>Exe</OutputType>
+    <TargetFramework>net8.0</TargetFramework>
+    <ImplicitUsings>enable</ImplicitUsings>
+    <Nullable>enable</Nullable>
+  </PropertyGroup>
+
+  <ItemGroup>
+    <PackageReference Include="Azure.Messaging.EventHubs" Version="5.12.2" />
+    <PackageReference Include="Azure.Messaging.EventHubs.Processor" Version="5.12.2" />
+    <PackageReference Include="Azure.Storage.Blobs" Version="12.25.1" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs
@@ -1,0 +1,138 @@
+using System.Collections.Concurrent;
+using Azure.Messaging.EventHubs;
+using Azure.Messaging.EventHubs.Consumer;
+using Azure.Messaging.EventHubs.Processor;
+using Azure.Messaging.EventHubs.Producer;
+using Azure.Storage.Blobs;
+
+const int eventCount = 10;
+const int checkpointFrequency = 5;
+
+string eventHubsConnectionString = GetRequiredEnvironmentVariable(
+    "EVENT_HUBS_CONNECTION_STRING");
+string blobStorageConnectionString = GetRequired
... (truncated)
```

### ✅ 9. powershell — 3353ms

**Input:**

```json
{"command":"dotnet build .\\EventHubsSample.csproj --nologo","description":"Build the Event Hubs sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Determining projects to restore...
  Restored C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\EventHubsSample.csproj (in 315 ms).
C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\Program.cs(32,9): error CS0117: 'EventProcessorClientOptions' does not contain a definition for 'DefaultStartingPosition' [C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\EventHubsSample.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\Program.cs(32,9): error CS0117: 'EventProcessorClientOptions' does not contain a definition for 'DefaultStartingPosition' [C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\EventHubsSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:01.75
<shellId: 0 completed with exit code 1>
```

### ✅ 10. apply_patch — 880ms

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n EventProcessorClient processor = new(\n     checkpointStore,\n     consumerGroup,\n-    eventHubsConnectionString,\n-    new EventProcessorClientOptions\n-    {\n-        DefaultStartingPosition = EventPosition.Earliest\n-    });\n+    eventHubsConnectionString);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012/Program.cs
@@ -26,11 +26,7 @@
 EventProcessorClient processor = new(
     checkpointStore,
     consumerGroup,
-    eventHubsConnectionString,
-    new EventProcessorClientOptions
-    {
-        DefaultStartingPosition = EventPosition.Earliest
-    });
+    eventHubsConnectionString);
 
 ConcurrentDictionary<string, int> eventsSinceCheckpoint = new();
 

```

### ✅ 11. powershell — 2702ms

**Input:**

```json
{"command":"dotnet build .\\EventHubsSample.csproj --nologo --no-restore","description":"Rebuild the Event Hubs sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
EventHubsSample -> C:\Users\weidxu\AppData\Local\Temp\hyoka-event-hubs-dp-dotnet-streaming-dotnet-azure-skills-azure-skill-mcp-3715211012\bin\Debug\net8.0\EventHubsSample.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:00.98
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `EventHubsSample.csproj`
- `Program.cs`

## Copilot Response

I’ll create a minimal console project with environment-based configuration, bounded batching, partition-safe checkpoint cadence, and guaranteed processor shutdown.

The installed SDK confirms `EventProcessorClient` has no configurable default starting-position option; it already begins at the earliest available event when no checkpoint exists, so I’m removing that invalid option.

Created a complete .NET 8 sample:

- `Program.cs` — sends 10 events with custom properties, processes received events, handles errors, and checkpoints every five events per partition.
- `EventHubsSample.csproj` — references:
  - `Azure.Messaging.EventHubs`
  - `Azure.Messaging.EventHubs.Processor`
  - `Azure.Storage.Blobs`

Configuration uses these environment variables:

```powershell
$env:EVENT_HUBS_CONNECTION_STRING = "<connection string including EntityPath>"
$env:BLOB_STORAGE_CONNECTION_STRING = "<storage connection string>"
$env:BLOB_CONTAINER_NAME = "<existing checkpoint container>"
$env:EVENT_HUB_CONSUMER_GROUP = '$Default' # optional

dotnet run
```

The producer uses `await using`, the event batch uses `using`, and the processor is stopped and handlers are unregistered in `finally`. The checkpoint Blob container must already exist.

## Reviewer Notes (LLM-as-Judge)

### Summary

7/7 criteria passed. The sample covers event batching, publishing, processor-based consumption, error handling, checkpointing, required packages, and disposal or shutdown patterns.

### Strengths

- Adds exactly 10 events with multiple custom properties and checks TryAdd() failures.
- Uses per-partition checkpoint counters and passes the processing cancellation token to UpdateCheckpointAsync().
- Uses await using for the producer, using for the batch, and a finally block to stop processing and unregister handlers.

## Grader Results

- send-receive-events.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `Azure.Messaging.EventHubs` and `Azure.Messaging.EventHubs.Processor` NuGet packages: Pass
      - `EventHubProducerClient` and `EventHubConsumerClient`: Pass
      - `CreateBatchAsync()` and `EventDataBatch.TryAdd()`: Pass
      - `SendAsync()` for publishing events: Pass
      - `EventProcessorClient` with `BlobContainerClient` for checkpointing: Pass
      - Event handler delegates and `ProcessEventArgs`: Pass
      - `UpdateCheckpointAsync()` for reliable processing: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id event-hubs-dp-dotnet-streaming --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
