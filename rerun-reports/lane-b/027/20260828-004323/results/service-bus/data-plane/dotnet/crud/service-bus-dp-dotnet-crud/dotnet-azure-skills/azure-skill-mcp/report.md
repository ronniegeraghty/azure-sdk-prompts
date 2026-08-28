# Evaluation Report: service-bus-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 501.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 501.4s |
| Timestamp | 2026-08-27T16:43:23Z |
| Files Generated | 2 |
| Event Count | 3631 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 476.2s |
| Review | 23.9s |
| **Total** | **501.4s** |

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
| Token Usage | in=374639 out=3675 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 3760ms

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

### ✅ 2. glob — 186162ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 185017ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 186164ms

**Input:**

```json
{"intent":"Generate a local-only C# console sample using Azure.Messaging.ServiceBus that sends single and batched queue messages, receives and completes messages, runs a processor, and demonstrates topic/subscription messaging with async disposal.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 5. view — 188986ms

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

### ✅ 6. view — 187612ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098 b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098
@@ -1,0 +1,0 @@


```

### ❌ 7. azure-documentation (via azure) — 188986ms

**Input:**

```json
{"intent":"Find current official .NET Azure.Messaging.ServiceBus examples for ServiceBusClient connection strings, queue sender and ServiceBusMessageBatch, ReceiveMessagesAsync and CompleteMessageAsync, ServiceBusProcessor handlers, topic sender and subscription receiver, and async disposal.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. view — 2062ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\\SKILL.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 9. glob — 2771ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-cost\cost-query\examples.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\azure-skills\default\.github\plugins\azure-skills\skills\azure-cost\cost-forecast\examples.md
```

### ✅ 10. glob — 2382ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills","pattern":"**/examples.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 11. powershell — 1298ms

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK version"}
```

**Output:**

```
10.0.111
<shellId: 0 completed with exit code 0>
```

### ✅ 12. powershell — 7579ms

**Input:**

```json
{"command":"dotnet new console --framework net10.0 --use-program-main --name ServiceBusMessagingDemo --output . --no-restore \u0026\u0026 dotnet add package Azure.Messaging.ServiceBus","description":"Scaffold console app and add Service Bus SDK","initial_wait":120,"mode":"sync"}
```

**Output:**

```
The template "Console App" was created successfully.

info : X.509 certificate chain validation will use the default trust store selected by .NET for code signing.
info : X.509 certificate chain validation will use the default trust store selected by .NET for timestamping.
info : Adding PackageReference for package 'Azure.Messaging.ServiceBus' into project 'C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\ServiceBusMessagingDemo.csproj'.
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/registrations2-semver2/azure.messaging.servicebus/index.json
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/registrations2-semver2/azure.messaging.servicebus/index.json 684ms
info : Restoring packages for C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\ServiceBusMessagingDemo.csproj...
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/flat2/azure.messaging.servicebus/index.json
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/flat2/azure.messaging.servicebus/index.json 463ms
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/flat2/azure.messaging.servicebus/7.20.2/azure.messaging.servicebus.7.20.2.nupkg
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/flat2/azure.messaging.servicebus/7.20.2/azure.messaging.servicebus.7.20.2.nupkg 806ms
info :   GET https://ms-feed-25.pkgs.vis
... (truncated)
```

### ✅ 13. view — 1660ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\\Program.cs"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/ServiceBusMessagingDemo.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/ServiceBusMessagingDemo.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/ServiceBusMessagingDemo.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/ServiceBusMessagingDemo.csproj
@@ -1,15 +1,15 @@
 ﻿<Project Sdk="Microsoft.NET.Sdk">
 
   <PropertyGroup>
     <OutputType>Exe</OutputType>
     <TargetFramework>net10.0</TargetFramework>
     <ImplicitUsings>enable</ImplicitUsings>
     <Nullable>enable</Nullable>
   </PropertyGroup>
 
   <ItemGroup>
     <PackageReference Include="Azure.Messaging.ServiceBus" Version="7.20.2" />
   </ItemGroup>
 
 </Project>
 

```

### ✅ 14. view — 1235ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\\ServiceBusMessagingDemo.csproj"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
@@ -1,10 +1,10 @@
 ﻿namespace ServiceBusMessagingDemo;
 
 class Program
 {
     static void Main(string[] args)
     {
         Console.WriteLine("Hello, World!");
     }
 }
 

```

### ✅ 15. apply_patch — 659ms

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\\Program.cs\n@@\n-﻿namespace ServiceBusMessagingDemo;\n-\n-class Program\n-{\n-    static void Main(string[] args)\n-    {\n-        Console.WriteLine(\"Hello, World!\");\n-    }\n-}\n+using Azure.Messaging.ServiceBus;\n+\n+namespace ServiceBusMessagingDemo;\n+\n+internal static class Program\n+{\n+    private static async Task Main()\n+    {\n+        string connectionString = GetRequiredEnvironmentVariable(\"SERVICE_BUS_CONNECTION_STRING\");\n+        string queueName = GetRequiredEnvironmentVariable(\"SERVICE_BUS_QUEUE_NAME\");\n+        string topicName = GetRequiredEnvironmentVariable(\"SERVICE_BUS_TOPIC_NAME\");\n+        string subscriptionName = GetRequiredEnvironmentVariable(\"SERVICE_BUS_SUBSCRIPTION_NAME\");\n+\n+        await using ServiceBusClient client = new(connectionString);\n+\n+        await RunQueueDemoAsync(client, queueName);\n+        await RunProcessorDemoAsync(client, queueName);\n+        await RunTopicSubscriptionDemoAsync(client, topicName, subscriptionName);\n+    }\n+\n+    private static async Task RunQueueDemoAsync(ServiceBusClient client, string queueName)\n+    {\n+        await using ServiceBusSender sender = client.CreateSender(queueName);\n+\n+        await sender.SendMessageAsync(new ServiceBusMessage(\"Single queue message\"));\n+        Console.WriteLine(\"Sent one queue message.\");\n+\n+        using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();\n+        for (int messageNumber = 1; messageNumber \u003c= 5; messageNumber++)\n+        {\n+            var message = new ServiceBusMessage($\"Batch message {messageNumber}\")\n+            {\n+                MessageId = Guid.NewGuid().ToString()\n+            };\n+\n+            if (!batch.TryAddMessage(message))\n+            {\n+                throw new InvalidOperationException(\n+                    $\"Message {messageNumber} is too large to fit in an empty Service Bus batch.\");\n+            }\n+        }\n+\n+        await sender.SendMessagesAsync(batch);\n+        Console.WriteLine(\"Sent a batch of five queue messages.\");\n+\n+        await using ServiceBusReceiver receiver = client.CreateReceiver(\n+            queueName,\n+            new ServiceBusReceiverOptions\n+            {\n+                ReceiveMode = ServiceBusReceiveMode.PeekLock\n+            });\n+\n+        IReadOnlyList\u003cServiceBusReceivedMessage\u003e messages =\n+            await receiver.ReceiveMessagesAsync(\n+                maxMessages: 6,\n+                maxWaitTime: TimeSpan.FromSeconds(30));\n+\n+        foreach (ServiceBusReceivedMessage message in messages)\n+        {\n+            Console.WriteLine($\"Received queue message: {message.Body}\");\n+            await receiver.CompleteMessageAsync(message);\n+        }\n+\n+        Console.WriteLine($\"Completed {messages.Count} queue message(s).\");\n+    }\n+\n+    private static async Task RunProcessorDemoAsync(ServiceBusClient client, string queueName)\n+    {\n+        var messageProcessed = new TaskCompletionSource(\n+            TaskCreationOptions.RunContinuationsAsynchronously);\n+\n+        await using ServiceBusProcessor processor = client.CreateProcessor(\n+            queueName,\n+            new ServiceBusProcessorOptions\n+            {\n+                AutoCompleteMessages = false,\n+                MaxConcurrentCalls = 1\n+            });\n+\n+        processor.ProcessMessageAsync += async args =\u003e\n+        {\n+            Console.WriteLine($\"Processor received: {args.Message.Body}\");\n+            await args.CompleteMessageAsync(args.Message);\n+            messageProcessed.TrySetResult();\n+        };\n+\n+        processor.ProcessErrorAsync += args =\u003e\n+        {\n+            Console.Error.WriteLine(\n+                $\"Processor error ({args.ErrorSource}, {args.EntityPath}): {args.Exception}\");\n+            messageProcessed.TrySetException(args.Exception);\n+            return Task.CompletedTask;\n+        };\n+\n+        await processor.StartProcessingAsync();\n+        try\n+        {\n+            await using ServiceBusSender sender = client.CreateSender(queueName);\n+            await sender.SendMessageAsync(new ServiceBusMessage(\"Message for the processor\"));\n+\n+            await messageProcessed.Task.WaitAsync(TimeSpan.FromSeconds(30));\n+        }\n+        finally\n+        {\n+            await processor.StopProcessingAsync();\n+        }\n+    }\n+\n+    private static async Task RunTopicSubscriptionDemoAsync(\n+        ServiceBusClient client,\n+        string topicName,\n+        string subscriptionName)\n+    {\n+        await using ServiceBusSender topicSender = client.CreateSender(topicName);\n+        await topicSender.SendMessageAsync(new ServiceBusMessage(\"Topic message\"));\n+        Console.WriteLine(\"Sent one topic message.\");\n+\n+        await using ServiceBusReceiver subscriptionReceiver =\n+            client.CreateReceiver(topicName, subscriptionName);\n+\n+        IReadOnlyList\u003cServiceBusReceivedMessage\u003e messages =\n+            await subscriptionReceiver.ReceiveMessagesAsync(\n+                maxMessages: 1,\n+                maxWaitTime: TimeSpan.FromSeconds(30));\n+\n+        if (messages.Count == 0)\n+        {\n+            throw new TimeoutException(\n+                $\"No message arrived on subscription '{subscriptionName}' within 30 seconds.\");\n+        }\n+\n+        ServiceBusReceivedMessage message = messages[0];\n+        Console.WriteLine($\"Received subscription message: {message.Body}\");\n+        await subscriptionReceiver.CompleteMessageAsync(message);\n+    }\n+\n+    private static string GetRequiredEnvironmentVariable(string name)\n+    {\n+        string? value = Environment.GetEnvironmentVariable(name);\n+        return !string.IsNullOrWhiteSpace(value)\n+            ? value\n+            : throw new InvalidOperationException(\n+                $\"Set the {name} environment variable before running the sample.\");\n+    }\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098/Program.cs
@@ -1,10 +1,146 @@
-﻿namespace ServiceBusMessagingDemo;
+using Azure.Messaging.ServiceBus;
+
+namespace ServiceBusMessagingDemo;
 
-class Program
+internal static class Program
 {
-    static void Main(string[] args)
+    private static async Task Main()
+    {
+        string connectionString = GetRequiredEnvironmentVariable("SERVICE_BUS_CONNECTION_STRING");
+        string queueName = GetRequiredEnvironmentVariable("SERVICE_BUS_QUEUE_NAME");
+        string topicName = GetRequiredEnvironmentVariable("SERVICE_BUS_TOPIC_NAME");
+        string subscriptionName = GetRequiredEnvironmentVariable("SERVICE_BUS_SUBSCRIPTION_NAME");
+
+        await using ServiceBusClient client = new(connectionString);
+
+        await RunQueueDemoAsync(client, queueName);
+        await RunProcessorDemoAsync(client, queueName);
+        await RunTopicSubscriptionDemoAsync(client, topicName, subscriptionName);
+    }
+
+    private static async Task RunQueueDemoAsync(ServiceBusClient client, string queueName)
+    {
+        await using ServiceBusSender sender = client.CreateSender(queueName);
+
+        await sender.SendMessageAsync(new ServiceBusMessage("Single queue message"));
+        Console.WriteLine("Sent one queue message.");
+
+        using ServiceBusMessageBatch batch = await sender.CreateMessageBatchAsync();
+        for (int messageNumber = 1; messageNumber <= 5; messageNumber++)
+        {
+            var message = new 
... (truncated)
```

### ✅ 16. powershell — 2324ms

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Build Service Bus sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
ServiceBusMessagingDemo -> C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3796983098\bin\Debug\net10.0\ServiceBusMessagingDemo.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:00.88
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `Program.cs`
- `ServiceBusMessagingDemo.csproj`

## Copilot Response

I’m checking the SDK guidance and workspace shape before creating the sample.

The workspace is empty, so I’m scaffolding a console app and adding the official Service Bus SDK before implementing the complete flow.

The sample now covers queue send/receive, explicit settlement, processor handlers, and topic/subscription flow; I’m compiling it to catch API or disposal issues.

Created a complete C# console sample in `Program.cs` using `Azure.Messaging.ServiceBus` **7.20.2**. It demonstrates:

- Single and five-message batch queue sends
- `ReceiveMessagesAsync` with PeekLock and `CompleteMessageAsync`
- Continuous `ServiceBusProcessor` handlers
- Topic send and subscription receive
- Proper `await using` disposal for clients, senders, receivers, and processor

Configure these environment variables before `dotnet run`:

```powershell
$env:SERVICE_BUS_CONNECTION_STRING = "<connection-string>"
$env:SERVICE_BUS_QUEUE_NAME = "<queue>"
$env:SERVICE_BUS_TOPIC_NAME = "<topic>"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription>"
dotnet run
```

The queue, topic, and subscription must already exist. The required NuGet package is declared in `ServiceBusMessagingDemo.csproj`:

```xml
<PackageReference Include="Azure.Messaging.ServiceBus" Version="7.20.2" />
```

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The sample comprehensively demonstrates the requested Service Bus workflows but omits two explicitly evaluated message-settlement operations.

### Strengths

- Uses the required Azure.Messaging.ServiceBus package.
- Demonstrates single, batch, receiver, processor, and topic/subscription messaging.
- Uses await using for asynchronous disposal of Service Bus clients and entities.

### Issues

- Missing demonstrations of AbandonMessageAsync and DeadLetterMessageAsync.

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
hyoka run --prompt-id service-bus-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
