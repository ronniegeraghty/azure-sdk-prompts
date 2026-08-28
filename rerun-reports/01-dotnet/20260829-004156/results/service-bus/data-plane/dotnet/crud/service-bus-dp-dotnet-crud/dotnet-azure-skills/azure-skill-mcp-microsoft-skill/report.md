# Evaluation Report: service-bus-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 90.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 90.6s |
| Timestamp | 2026-08-28T18:34:49Z |
| Files Generated | 3 |
| Event Count | 4047 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 64.0s |
| Review | 25.0s |
| **Total** | **90.6s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-servicebus-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=273783 out=3007 |
| Turn Count | 7 |

## Tool Calls

### ✅ 1. skill — 723ms

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

### ✅ 2. glob — 1481ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 3. view — 1485ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-servicebus-dotnet\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 4. view — 1482ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-servicebus-dotnet\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 5. view — 1023ms

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

### ✅ 6. azure-get_azure_bestpractices (via azure) — 1364ms

**Input:**

```json
{"intent":"Generate a local-only C# console sample using Azure.Messaging.ServiceBus with connection-string client creation, queue single and batch send, pull receive and completion, processor handlers, topic send, subscription receive, and async disposal. Do not provision or contact Azure.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 717ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure SDK .NET code-generation best practices to a local-only Service Bus console sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. apply_patch — 675ms

**Input:**

```json
"*** Begin Patch\n*** Add File: ServiceBusMessagingDemo.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Messaging.ServiceBus\" Version=\"7.20.1\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure.Messaging.ServiceBus;\n+\n+const string ConnectionStringVariable = \"AZURE_SERVICEBUS_CONNECTION_STRING\";\n+const string QueueNameVariable = \"AZURE_SERVICEBUS_QUEUE_NAME\";\n+const string TopicNameVariable = \"AZURE_SERVICEBUS_TOPIC_NAME\";\n+const string SubscriptionNameVariable = \"AZURE_SERVICEBUS_SUBSCRIPTION_NAME\";\n+\n+string connectionString = GetRequiredEnvironmentVariable(ConnectionStringVariable);\n+string queueName = GetRequiredEnvironmentVariable(QueueNameVariable);\n+string topicName = GetRequiredEnvironmentVariable(TopicNameVariable);\n+string subscriptionName = GetRequiredEnvironmentVariable(SubscriptionNameVariable);\n+\n+try\n+{\n+    await using ServiceBusClient client = new(connectionString);\n+    await using ServiceBusSender queueSender = client.CreateSender(queueName);\n+\n+    await queueSender.SendMessageAsync(\n+        new ServiceBusMessage(\"Single queue message\")\n+        {\n+            MessageId = Guid.NewGuid().ToString()\n+        });\n+    Console.WriteLine(\"Sent one message to the queue.\");\n+\n+    using (ServiceBusMessageBatch batch = await queueSender.CreateMessageBatchAsync())\n+    {\n+        for (int i = 1; i \u003c= 5; i++)\n+        {\n+            ServiceBusMessage message = new($\"Batch message {i}\")\n+            {\n+                MessageId = Guid.NewGuid().ToString(),\n+                ApplicationProperties =\n+                {\n+                    [\"BatchSequence\"] = i\n+                }\n+            };\n+\n+            if (!batch.TryAddMessage(message))\n+            {\n+                throw new InvalidOperationException(\n+                    $\"Batch message {i} is too large for an empty or partially filled batch.\");\n+            }\n+        }\n+\n+        await queueSender.SendMessagesAsync(batch);\n+    }\n+    Console.WriteLine(\"Sent a batch of five messages to the queue.\");\n+\n+    await using ServiceBusReceiver queueReceiver = client.CreateReceiver(\n+        queueName,\n+        new ServiceBusReceiverOptions\n+        {\n+            ReceiveMode = ServiceBusReceiveMode.PeekLock\n+        });\n+\n+    IReadOnlyList\u003cServiceBusReceivedMessage\u003e receivedMessages =\n+        await queueReceiver.ReceiveMessagesAsync(\n+            maxMessages: 6,\n+            maxWaitTime: TimeSpan.FromSeconds(10));\n+\n+    foreach (ServiceBusReceivedMessage message in receivedMessages)\n+    {\n+        Console.WriteLine($\"Pull receiver processed: {message.Body}\");\n+        await queueReceiver.CompleteMessageAsync(message);\n+    }\n+\n+    await using ServiceBusProcessor processor = client.CreateProcessor(\n+        queueName,\n+        new ServiceBusProcessorOptions\n+        {\n+            AutoCompleteMessages = false,\n+            MaxConcurrentCalls = 2\n+        });\n+\n+    processor.ProcessMessageAsync += async args =\u003e\n+    {\n+        Console.WriteLine($\"Processor handled: {args.Message.Body}\");\n+        await args.CompleteMessageAsync(args.Message);\n+    };\n+\n+    processor.ProcessErrorAsync += args =\u003e\n+    {\n+        Console.Error.WriteLine(\n+            $\"Processor error ({args.ErrorSource}) on {args.EntityPath}: {args.Exception.Message}\");\n+        return Task.CompletedTask;\n+    };\n+\n+    bool processorStarted = false;\n+    try\n+    {\n+        await processor.StartProcessingAsync();\n+        processorStarted = true;\n+\n+        await queueSender.SendMessageAsync(\n+            new ServiceBusMessage(\"Message for the continuous processor\")\n+            {\n+                MessageId = Guid.NewGuid().ToString()\n+            });\n+\n+        Console.WriteLine(\"Processor is running for 10 seconds.\");\n+        await Task.Delay(TimeSpan.FromSeconds(10));\n+    }\n+    finally\n+    {\n+        if (processorStarted)\n+        {\n+            await processor.StopProcessingAsync();\n+        }\n+    }\n+\n+    await using ServiceBusSender topicSender = client.CreateSender(topicName);\n+    await topicSender.SendMessageAsync(\n+        new ServiceBusMessage(\"Message published to the topic\")\n+        {\n+            MessageId = Guid.NewGuid().ToString()\n+        });\n+    Console.WriteLine(\"Sent one message to the topic.\");\n+\n+    await using ServiceBusReceiver subscriptionReceiver =\n+        client.CreateReceiver(topicName, subscriptionName);\n+\n+    ServiceBusReceivedMessage? subscriptionMessage =\n+        await subscriptionReceiver.ReceiveMessageAsync(TimeSpan.FromSeconds(10));\n+\n+    if (subscriptionMessage is null)\n+    {\n+        Console.WriteLine(\"No subscription message arrived within 10 seconds.\");\n+    }\n+    else\n+    {\n+        Console.WriteLine($\"Subscription received: {subscriptionMessage.Body}\");\n+        await subscriptionReceiver.CompleteMessageAsync(subscriptionMessage);\n+    }\n+}\n+catch (ServiceBusException exception)\n+{\n+    Console.Error.WriteLine(\n+        $\"Service Bus operation failed ({exception.Reason}): {exception.Message}\");\n+    Environment.ExitCode = 1;\n+}\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+    if (string.IsNullOrWhiteSpace(value))\n+    {\n+        throw new InvalidOperationException(\n+            $\"Set the required environment variable '{name}' before running the sample.\");\n+    }\n+\n+    return value;\n+}\n*** Add File: README.md\n+# Azure Service Bus messaging demo\n+\n+This .NET 8 console application demonstrates:\n+\n+- sending one queue message and a five-message `ServiceBusMessageBatch`;\n+- receiving queue messages with `ReceiveMessagesAsync` and settling them with\n+  `CompleteMessageAsync`;\n+- continuous processing with `ServiceBusProcessor` message and error handlers;\n+- publishing to a topic and receiving from a subscription; and\n+- asynchronous disposal of Service Bus clients, senders, receivers, and the\n+  processor with `await using`.\n+\n+## NuGet package\n+\n+The sample requires `Azure.Messaging.ServiceBus` 7.20.1:\n+\n+```powershell\n+dotnet add package Azure.Messaging.ServiceBus --version 7.20.1\n+```\n+\n+`Azure.Identity` is not required because this example specifically uses a\n+connection string. For production applications, prefer Microsoft Entra ID and\n+managed identity rather than connection strings.\n+\n+## Configuration\n+\n+Create the queue, topic, and subscription before running the sample. Then set\n+the configuration through environment variables; do not put credentials in\n+source code.\n+\n+```powershell\n+$env:AZURE_SERVICEBUS_CONNECTION_STRING = \"Endpoint=sb://\u003cnamespace\u003e.servicebus.windows.net/;SharedAccessKeyName=\u003cpolicy\u003e;SharedAccessKey=\u003ckey\u003e\"\n+$env:AZURE_SERVICEBUS_QUEUE_NAME = \"\u003cqueue-name\u003e\"\n+$env:AZURE_SERVICEBUS_TOPIC_NAME = \"\u003ctopic-name\u003e\"\n+$env:AZURE_SERVICEBUS_SUBSCRIPTION_NAME = \"\u003csubscription-name\u003e\"\n+\n+dotnet run\n+```\n+\n+The connection string policy needs data-plane permission to send and receive.\n+This application does not create or modify Service Bus entities.\n+\n+## Reference\n+\n+- [Azure Service Bus client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/messaging.servicebus-readme)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/ServiceBusMessagingDemo.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/ServiceBusMessagingDemo.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/ServiceBusMessagingDemo.csproj
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



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758/Program.cs
@@ -1,0 +1,153 @@
+using Azure.Messaging.ServiceBus;
+
+const string ConnectionStringVariable = "AZURE_SERVICEBUS_CONNECTION_STRING";
+const string QueueNameVariable = "AZURE_SERVICEBUS_QUEUE_NAME";
+const string TopicNameVariable = "AZURE_SERVICEBUS_TOPIC_NAME";
+const string SubscriptionNameVariable = "AZURE_SERVICEBUS_SUBSCRIPTION_NAME";
+
+string connectionString = GetRequiredEnvironmentVariable(ConnectionStringVariable);
+string queueName = GetRequiredEnvironmentVariable(QueueNameVariable);
+string topicName = GetRequiredEnvir
... (truncated)
```

### ✅ 9. powershell — 2692ms

**Input:**

```json
{"command":"dotnet build .\\ServiceBusMessagingDemo.csproj --nologo","description":"Build the Service Bus demo","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Determining projects to restore...
  Restored C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758\ServiceBusMessagingDemo.csproj (in 225 ms).
  ServiceBusMessagingDemo -> C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-4206348758\bin\Debug\net8.0\ServiceBusMessagingDemo.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.36
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `Program.cs`
- `README.md`
- `ServiceBusMessagingDemo.csproj`

## Copilot Response

I’ll scaffold the local console sample from the SDK guidance, then build it without connecting to Azure.

The sample will use environment-based configuration, fail fast on missing settings, and run the processor for a bounded demo interval so it does not hang indefinitely.

Created a complete .NET 8 Azure Service Bus console sample:

- `Program.cs` — single/batch queue sends, pull receiver, message completion, processor handlers, topic/subscription messaging, and `await using` disposal.
- `ServiceBusMessagingDemo.csproj` — references `Azure.Messaging.ServiceBus` 7.20.1.
- `README.md` — setup, required environment variables, NuGet command, permissions, and reference documentation.

Credentials are read from `AZURE_SERVICEBUS_CONNECTION_STRING`; no secrets are embedded. Queue, topic, and subscription entities must already exist.

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The sample comprehensively demonstrates the requested Service Bus operations, but it does not show abandoning or dead-lettering messages as required by check_6.

### Strengths

- Includes the required Azure.Messaging.ServiceBus dependency and connection-string client creation.
- Correctly implements single and batched queue sends, pull receiving, explicit completion, processor handlers, and topic/subscription messaging.
- Uses await using for asynchronous disposal and keeps credentials in environment variables.

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
hyoka run --prompt-id service-bus-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
