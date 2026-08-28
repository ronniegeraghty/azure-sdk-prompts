# Evaluation Report: service-bus-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 479.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 479.4s |
| Timestamp | 2026-08-27T16:51:46Z |
| Files Generated | 3 |
| Event Count | 3870 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 439.9s |
| Review | 38.4s |
| **Total** | **479.4s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-servicebus-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=228790 out=2655 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. skill — 5024ms

**Input:**

```json
{"skill":"azure-servicebus-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure.Messaging.ServiceBus (.NET)

Enterprise messaging SDK for reliable message delivery with queues, topics, subscriptions, and sessions.

## Installation

```bash
dotnet add package Azure.Messaging.ServiceBus
dotnet add package Azure.Identity
```

**Current Version**: v7.20.1 (stable)

## Environment Variables

```bash
AZURE_SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net  # Required: Service Bus fully qualified namespace
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
AZURE_SERVICEBUS_CONNECTION_STRING=Endpoint=sb://...  # Alternative to Entra ID auth
```

## Authentication

### Microsoft Entra Token Credential

```csharp
using Azure.Identity;
using Azure.Messaging.ServiceBus;

string fullyQualifiedNamespace = "<namespace>.servicebus.windows.net";
// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
var credential = new DefaultAzureCredential(
    DefaultAzureCredential.DefaultEnvironmentVariableName
);
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet#credential-classes
// var credential = new ManagedIdentityCredential();
await using ServiceBusClient client = new(fullyQualifiedNamespace, credential);
```

### Connection String

```csharp
string connectionString = "<connection_string>";
await using ServiceBusClient client = new(connectionString);
```

### ASP.NET Core Dependency Injection

```csharp
services.AddAzureClients(builder =>
{
    builder.AddServiceBusClientWithNamespace("<namespace>.servicebus.windows.net");
    builder.UseCredential(new DefaultAzureCredential());
});
```

## Client Hierarchy

```
ServiceBusClient
├── CreateSender(queueOrTopicName)      → ServiceBusSender
├── CreateReceiver(queueName)           → ServiceBusReceiver
├── CreateReceiver(t
... (truncated)
```

### ❌ 2. view — 188399ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-servicebus-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 188397ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-servicebus-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 188401ms

**Input:**

```json
{"intent":"Generate a local-only C# console sample demonstrating Azure Service Bus queue sends, batching, receives, settlement, processor handlers, topics, subscriptions, and async disposal.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. glob — 188394ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. view — 183339ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-servicebus-dotnet"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-servicebus-dotnet b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-servicebus-dotnet
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-dotnet/skills/azure-servicebus-dotnet
@@ -1,0 +1,1 @@
 SKILL.md

```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 183341ms

**Input:**

```json
{"intent":"Discover the command for Azure code-generation best practices for a C# Azure Service Bus console sample.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. apply_patch — 2636ms

**Input:**

```json
"*** Begin Patch\n*** Add File: ServiceBusMessagingDemo.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Messaging.ServiceBus\" Version=\"7.20.1\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure.Messaging.ServiceBus;\n+\n+string connectionString = GetRequiredEnvironmentVariable(\"AZURE_SERVICEBUS_CONNECTION_STRING\");\n+string queueName = GetRequiredEnvironmentVariable(\"AZURE_SERVICEBUS_QUEUE_NAME\");\n+string topicName = GetRequiredEnvironmentVariable(\"AZURE_SERVICEBUS_TOPIC_NAME\");\n+string subscriptionName = GetRequiredEnvironmentVariable(\"AZURE_SERVICEBUS_SUBSCRIPTION_NAME\");\n+\n+await using var client = new ServiceBusClient(connectionString);\n+\n+await SendAndReceiveQueueMessagesAsync(client, queueName);\n+await ProcessQueueMessageAsync(client, queueName);\n+await SendAndReceiveTopicMessageAsync(client, topicName, subscriptionName);\n+\n+static async Task SendAndReceiveQueueMessagesAsync(\n+    ServiceBusClient client,\n+    string queueName)\n+{\n+    await using ServiceBusSender sender = client.CreateSender(queueName);\n+\n+    await sender.SendMessageAsync(new ServiceBusMessage(\"Single queue message\"));\n+    Console.WriteLine(\"Sent one message to the queue.\");\n+\n+    using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();\n+\n+    for (int i = 1; i \u003c= 5; i++)\n+    {\n+        var message = new ServiceBusMessage($\"Batch message {i}\")\n+        {\n+            MessageId = Guid.NewGuid().ToString()\n+        };\n+\n+        if (!batch.TryAddMessage(message))\n+        {\n+            throw new InvalidOperationException(\n+                $\"Batch message {i} is too large to fit in the Service Bus batch.\");\n+        }\n+    }\n+\n+    await sender.SendMessagesAsync(batch);\n+    Console.WriteLine(\"Sent a batch of 5 messages to the queue.\");\n+\n+    await using ServiceBusReceiver receiver = client.CreateReceiver(queueName);\n+    IReadOnlyList\u003cServiceBusReceivedMessage\u003e messages =\n+        await receiver.ReceiveMessagesAsync(\n+            maxMessages: 6,\n+            maxWaitTime: TimeSpan.FromSeconds(10));\n+\n+    Console.WriteLine($\"Received {messages.Count} queue message(s).\");\n+\n+    foreach (ServiceBusReceivedMessage message in messages)\n+    {\n+        Console.WriteLine($\"Queue message: {message.Body}\");\n+        await receiver.CompleteMessageAsync(message);\n+        Console.WriteLine($\"Completed message {message.MessageId}.\");\n+    }\n+}\n+\n+static async Task ProcessQueueMessageAsync(\n+    ServiceBusClient client,\n+    string queueName)\n+{\n+    await using ServiceBusSender sender = client.CreateSender(queueName);\n+    await sender.SendMessageAsync(new ServiceBusMessage(\"Message for the processor\"));\n+\n+    await using ServiceBusProcessor processor = client.CreateProcessor(\n+        queueName,\n+        new ServiceBusProcessorOptions\n+        {\n+            AutoCompleteMessages = false,\n+            MaxConcurrentCalls = 1\n+        });\n+\n+    var processed = new TaskCompletionSource(\n+        TaskCreationOptions.RunContinuationsAsynchronously);\n+\n+    processor.ProcessMessageAsync += async args =\u003e\n+    {\n+        try\n+        {\n+            Console.WriteLine($\"Processor message: {args.Message.Body}\");\n+            await args.CompleteMessageAsync(args.Message);\n+            processed.TrySetResult();\n+        }\n+        catch (Exception exception)\n+        {\n+            processed.TrySetException(exception);\n+            throw;\n+        }\n+    };\n+\n+    processor.ProcessErrorAsync += args =\u003e\n+    {\n+        Console.Error.WriteLine(\n+            $\"Processor error ({args.ErrorSource}, {args.EntityPath}): \" +\n+            args.Exception.Message);\n+        processed.TrySetException(args.Exception);\n+        return Task.CompletedTask;\n+    };\n+\n+    await processor.StartProcessingAsync();\n+    Console.WriteLine(\"Processor started.\");\n+\n+    try\n+    {\n+        await processed.Task.WaitAsync(TimeSpan.FromSeconds(30));\n+    }\n+    finally\n+    {\n+        await processor.StopProcessingAsync();\n+        Console.WriteLine(\"Processor stopped.\");\n+    }\n+}\n+\n+static async Task SendAndReceiveTopicMessageAsync(\n+    ServiceBusClient client,\n+    string topicName,\n+    string subscriptionName)\n+{\n+    await using ServiceBusSender topicSender = client.CreateSender(topicName);\n+    await topicSender.SendMessageAsync(\n+        new ServiceBusMessage(\"Message published to the topic\"));\n+    Console.WriteLine(\"Sent one message to the topic.\");\n+\n+    await using ServiceBusReceiver subscriptionReceiver =\n+        client.CreateReceiver(topicName, subscriptionName);\n+\n+    IReadOnlyList\u003cServiceBusReceivedMessage\u003e messages =\n+        await subscriptionReceiver.ReceiveMessagesAsync(\n+            maxMessages: 1,\n+            maxWaitTime: TimeSpan.FromSeconds(10));\n+\n+    foreach (ServiceBusReceivedMessage message in messages)\n+    {\n+        Console.WriteLine($\"Subscription message: {message.Body}\");\n+        await subscriptionReceiver.CompleteMessageAsync(message);\n+    }\n+}\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+\n+    return string.IsNullOrWhiteSpace(value)\n+        ? throw new InvalidOperationException(\n+            $\"Set the required environment variable '{name}'.\")\n+        : value;\n+}\n*** Add File: README.md\n+# Azure Service Bus messaging demo\n+\n+This .NET 8 console program demonstrates:\n+\n+- Sending one queue message\n+- Sending five queue messages with `ServiceBusMessageBatch`\n+- Receiving queue messages with `ReceiveMessagesAsync`\n+- Settling messages with `CompleteMessageAsync`\n+- Continuous processing with `ServiceBusProcessor` handlers\n+- Publishing to a topic and receiving from a subscription\n+- Asynchronous disposal of clients, senders, receivers, and processors\n+\n+## Required NuGet package\n+\n+```powershell\n+dotnet add package Azure.Messaging.ServiceBus --version 7.20.1\n+```\n+\n+`Azure.Identity` is not required because this sample explicitly uses a connection\n+string. For production applications, prefer Microsoft Entra authentication and add\n+`Azure.Identity`.\n+\n+## Configuration\n+\n+Create the queue, topic, and subscription before running the sample. Then set:\n+\n+```powershell\n+$env:AZURE_SERVICEBUS_CONNECTION_STRING = \"Endpoint=sb://\u003cnamespace\u003e.servicebus.windows.net/;SharedAccessKeyName=\u003ckey-name\u003e;SharedAccessKey=\u003ckey\u003e\"\n+$env:AZURE_SERVICEBUS_QUEUE_NAME = \"\u003cqueue-name\u003e\"\n+$env:AZURE_SERVICEBUS_TOPIC_NAME = \"\u003ctopic-name\u003e\"\n+$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = \"\u003csubscription-name\u003e\"\n+```\n+\n+The connection string policy needs send and listen permissions for the configured\n+entities.\n+\n+## Run\n+\n+```powershell\n+dotnet run\n+```\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/ServiceBusMessagingDemo.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/ServiceBusMessagingDemo.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/ServiceBusMessagingDemo.csproj
@@ -1,0 +1,15 @@
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
+    <PackageReference Include="Azure.Messaging.ServiceBus" Version="7.20.1" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2183224978/Program.cs
@@ -1,0 +1,148 @@
+using Azure.Messaging.ServiceBus;
+
+string connectionString = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_CONNECTION_STRING");
+string queueName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_QUEUE_NAME");
+string topicName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_TOPIC_NAME");
+string subscriptionName = GetRequiredEnvironmentVariable("AZURE_SERVICEBUS_SUBSCRIPTION_NAME");
+
+await using var client = new ServiceBusClient(connectionString);
+
+await SendAndReceiveQueueMessagesAsync(client, queu
... (truncated)
```

### ✅ 9. powershell — 5485ms

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Build the Service Bus demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.55
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `Program.cs`
- `README.md`
- `ServiceBusMessagingDemo.csproj`

## Copilot Response

I’ll build this as a local console sample using environment-variable configuration, with queue, processor, and topic/subscription flows kept explicit.

The skill’s reference files aren’t at the advertised root, so I’m locating their actual paths before creating the sample.

I found no separate reference files in the installed skill—the SDK guidance is contained in `SKILL.md`. I’m creating a self-contained .NET console project with the package pinned and all entity names supplied through environment variables.

The sample is implemented; I’m compiling it now to catch any SDK signature or disposal issues.

Implemented a complete .NET 8 Azure Service Bus sample:

- `Program.cs` — queue sending, 5-message batching, receiving, completion, processor handlers, and topic/subscription messaging
- `ServiceBusMessagingDemo.csproj` — references `Azure.Messaging.ServiceBus` 7.20.1
- `README.md` — required NuGet command, environment variables, permissions, and run instructions

All `ServiceBusClient`, sender, receiver, and processor instances use proper `await using` disposal. Configuration is loaded from environment variables; no credentials are embedded.

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The sample covers queue and topic messaging, batching, receiving, completion, processor handlers, package setup, and asynchronous disposal, but it omits two required settlement operations.

### Strengths

- Uses the modern Azure.Messaging.ServiceBus SDK with an explicit package reference.
- Demonstrates single-message and five-message batch sending with proper batch-size handling.
- Uses ReceiveMessagesAsync, manual completion, processor handlers, and topic/subscription messaging.
- Uses await using for disposable Service Bus clients, senders, receivers, and processors.

### Issues

- No demonstration of AbandonMessageAsync.
- No demonstration of DeadLetterMessageAsync.

## Grader Results

- send-receive-messages.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `Azure.Messaging.ServiceBus` NuGet package: Pass
      - `ServiceBusClient` creation with connection string or `DefaultAzureCredential`: Pass
      - `ServiceBusSender` and `ServiceBusMessage` for sending: Pass
      - `ServiceBusMessageBatch` and `TryAddMessage()`: Pass
      - `ServiceBusReceiver` and `ReceiveMessagesAsync()`: Pass
      - `CompleteMessageAsync()`, `AbandonMessageAsync()`, `DeadLetterMessageAsync()`: Fail
      - `ServiceBusProcessor` with `ProcessMessageAsync` and `ProcessErrorAsync`: Pass
      - Topic/subscription with `CreateSender(topicName)` and `CreateReceiver(topicName, subscriptionName)`: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
