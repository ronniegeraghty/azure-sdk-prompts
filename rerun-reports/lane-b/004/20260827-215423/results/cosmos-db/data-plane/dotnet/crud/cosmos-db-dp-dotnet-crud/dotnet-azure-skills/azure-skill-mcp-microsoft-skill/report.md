# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 374.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 374.9s |
| Timestamp | 2026-08-27T13:54:23Z |
| Files Generated | 2 |
| Event Count | 3640 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 332.5s |
| Review | 41.1s |
| **Total** | **374.9s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=400505 out=3748 |
| Turn Count | 10 |

## Tool Calls

### ❌ 1. skill — 2069ms

**Input:**

```json
{"skill":"azure-cosmos-db-dotnet"}
```

**Error:**

```
Skill not found: azure-cosmos-db-dotnet
```

### ✅ 2. glob — 188489ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 3. azure-get_azure_bestpractices (via azure) — 188493ms

**Input:**

```json
{"intent":"Get current C# Azure Cosmos DB NoSQL SDK best practices for a local console program demonstrating CRUD, parameterized queries, CosmosException handling, and connection strings without provisioning Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 4. azure-documentation (via azure) — 188495ms

**Input:**

```json
{"intent":"Find official current Microsoft documentation for .NET Azure Cosmos DB for NoSQL CRUD with Microsoft.Azure.Cosmos, CreateDatabaseIfNotExistsAsync, CreateContainerIfNotExistsAsync, CreateItemAsync, ReadItemAsync, parameterized QueryDefinition, ReplaceItemAsync, DeleteItemAsync, and CosmosException.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. web_fetch — 6535ms

**Input:**

```json
{"max_length":20000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 6. web_fetch — 6528ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/quickstart-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/cosmos-db/quickstart-dotnet (redirected from https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/quickstart-dotnet):
---
layout: Conceptual
title: Quickstart - Azure SDK for .NET - Azure Cosmos DB | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/cosmos-db/quickstart-dotnet
uhfHeaderId: MSDocsHeader-NoSQL
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
breadcrumb_path: ../breadcrumb/azure-nosql/toc.json
feedback_system: Standard
feedback_product_url: https://feedback.azure.com/d365community/forum/3002b3be-0d25-ec11-b6e6-000d3a4f0858
feedback_help_link_url: /answers/tags/187/azure-cosmos-db/
feedback_help_link_type: get-help-at-qna
adobe-target: true
ms.service: azure-cosmos-db
ms.subservice: nosql
ms.reviewer: mjbrown
ms.author: sidandrews
author: seesharprun
description: Deploy a .NET web application that uses the Azure SDK for .NET to interact with Azure Cosmos DB for NoSQL data in this quickstart.
ms.devlang: csharp
ms.topic: quickstart-sdk
ms.date: 2026-03-25T00:00:00.0000000Z
ms.custom: devx-track-csharp, devx-track-dotnet, devx-track-extended-azdevcli
locale: en-us
document_id: 03f9c18e-79b0-4fbb-361d-ef08497e9e38
document_version_independent_id: 03f9c18e-79b0-4fbb-361d-ef08497e9e38
updated_at: 2026-06-16T17:14:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/nosql-docs-pr/blob/live/azure/cosmos-db/quickstart-dotnet.md
gitcommit: https://github.com/MicrosoftDocs/nosql-docs-pr/blob/8b73e146062fd96672b45b48656e2402664ba3f9/azure/cosmos-db/quickstart-dotnet.md
git_commit_id: 8b73e146062fd96672b45b48656e2402664ba3f9
site_name: Docs
depot_name: Learn.azure-nosql
page_type: conceptual
toc_rel: toc.json
word_count: 991
asset_id: cosmos-db/quickstart-dotnet
moniker_range_name: 
monikers: []
item_type: Content
source_path: azure/cosmos-db/quickstart-dotnet.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/dev
... (truncated)
```

### ✅ 7. powershell — 5377ms

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
10.0.111
<shellId: 0 completed with exit code 0>
```

### ✅ 8. apply_patch — 3665ms

**Input:**

```json
"*** Begin Patch\n*** Add File: CosmosCrudSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.*\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Net;\n+using Microsoft.Azure.Cosmos;\n+\n+const string databaseName = \"TestDB\";\n+const string containerName = \"Items\";\n+const string partitionKeyPath = \"/category\";\n+const string category = \"electronics\";\n+\n+try\n+{\n+    string connectionString =\n+        Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\")\n+        ?? throw new InvalidOperationException(\n+            \"Set COSMOS_CONNECTION_STRING to an Azure Cosmos DB connection string.\");\n+\n+    using CosmosClient client = new(\n+        connectionString,\n+        new CosmosClientOptions\n+        {\n+            ApplicationName = \"CosmosCrudSample\"\n+        });\n+\n+    DatabaseResponse databaseResponse =\n+        await client.CreateDatabaseIfNotExistsAsync(databaseName);\n+    Database database = databaseResponse.Database;\n+\n+    ContainerResponse containerResponse =\n+        await database.CreateContainerIfNotExistsAsync(\n+            id: containerName,\n+            partitionKeyPath: partitionKeyPath);\n+    Container container = containerResponse.Container;\n+\n+    Item item = new(\n+        id: Guid.NewGuid().ToString(),\n+        category: category,\n+        name: \"Bluetooth Speaker\",\n+        quantity: 10);\n+    PartitionKey partitionKey = new(item.category);\n+\n+    ItemResponse\u003cItem\u003e createResponse =\n+        await container.CreateItemAsync(item, partitionKey);\n+    Console.WriteLine(\n+        $\"Created '{createResponse.Resource.name}' \" +\n+        $\"(RU: {createResponse.RequestCharge:F2}).\");\n+\n+    ItemResponse\u003cItem\u003e readResponse =\n+        await container.ReadItemAsync\u003cItem\u003e(item.id, partitionKey);\n+    Console.WriteLine(\n+        $\"Read '{readResponse.Resource.name}' with quantity \" +\n+        $\"{readResponse.Resource.quantity} \" +\n+        $\"(RU: {readResponse.RequestCharge:F2}).\");\n+\n+    QueryDefinition query = new(\n+        \"SELECT * FROM c WHERE c.category = @category\");\n+    query.WithParameter(\"@category\", category);\n+\n+    using FeedIterator\u003cItem\u003e iterator =\n+        container.GetItemQueryIterator\u003cItem\u003e(\n+            query,\n+            requestOptions: new QueryRequestOptions\n+            {\n+                PartitionKey = partitionKey\n+            });\n+\n+    while (iterator.HasMoreResults)\n+    {\n+        FeedResponse\u003cItem\u003e page = await iterator.ReadNextAsync();\n+        foreach (Item result in page)\n+        {\n+            Console.WriteLine(\n+                $\"Query result: {result.id}, {result.name}, \" +\n+                $\"quantity {result.quantity}\");\n+        }\n+    }\n+\n+    Item updatedItem = item with { quantity = 25 };\n+    ItemResponse\u003cItem\u003e replaceResponse =\n+        await container.ReplaceItemAsync(\n+            updatedItem,\n+            updatedItem.id,\n+            partitionKey);\n+    Console.WriteLine(\n+        $\"Updated quantity to {replaceResponse.Resource.quantity} \" +\n+        $\"(RU: {replaceResponse.RequestCharge:F2}).\");\n+\n+    ItemResponse\u003cItem\u003e deleteResponse =\n+        await container.DeleteItemAsync\u003cItem\u003e(item.id, partitionKey);\n+    Console.WriteLine(\n+        $\"Deleted item '{item.id}' \" +\n+        $\"(RU: {deleteResponse.RequestCharge:F2}).\");\n+}\n+catch (CosmosException exception)\n+{\n+    Console.Error.WriteLine(\n+        $\"Cosmos DB request failed: {(int)exception.StatusCode} \" +\n+        $\"{exception.StatusCode}\");\n+    Console.Error.WriteLine($\"Message: {exception.Message}\");\n+    Console.Error.WriteLine($\"Activity ID: {exception.ActivityId}\");\n+    Console.Error.WriteLine($\"Substatus code: {exception.SubStatusCode}\");\n+    Console.Error.WriteLine($\"Request charge: {exception.RequestCharge:F2} RU\");\n+\n+    if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+    {\n+        Console.Error.WriteLine(\n+            $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms\");\n+    }\n+\n+    Environment.ExitCode = 1;\n+}\n+catch (InvalidOperationException exception)\n+{\n+    Console.Error.WriteLine($\"Configuration error: {exception.Message}\");\n+    Environment.ExitCode = 2;\n+}\n+catch (Exception exception)\n+{\n+    Console.Error.WriteLine(\n+        $\"Unexpected error ({exception.GetType().Name}): {exception.Message}\");\n+    Environment.ExitCode = 3;\n+}\n+\n+public sealed record Item(\n+    string id,\n+    string category,\n+    string name,\n+    int quantity);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj
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
+    <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.*" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs
@@ -1,0 +1,127 @@
+using System.Net;
+using Microsoft.Azure.Cosmos;
+
+const string databaseName = "TestDB";
+const string containerName = "Items";
+const string partitionKeyPath = "/category";
+const string category = "electronics";
+
+try
+{
+    string connectionString =
+        Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
+        ?? throw new InvalidOperationException(
+            "Set COSMOS_CONNECTION_STRING to an Azure Cosmos DB connection string.");
+
+    using CosmosClient client = new(
+        connectionString,
+        new CosmosClientOp
... (truncated)
```

### ✅ 9. powershell — 4606ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo","description":"Build Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Determining projects to restore...
  Restored C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj (in 735 ms).
C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.62.1\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(72,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.62.1\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(72,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:01.51
<shellId: 1 completed with exit code 1>
```

### ✅ 10. apply_patch — 743ms

**Input:**

```json
"*** Begin Patch\n*** Update File: CosmosCrudSample.csproj\n@@\n   \u003cItemGroup\u003e\n-    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.*\" /\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.62.1\" /\u003e\n+    \u003cPackageReference Include=\"Newtonsoft.Json\" Version=\"13.0.3\" /\u003e\n   \u003c/ItemGroup\u003e\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/CosmosCrudSample.csproj
@@ -8,7 +8,8 @@
   </PropertyGroup>
 
   <ItemGroup>
-    <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.*" />
+    <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.62.1" />
+    <PackageReference Include="Newtonsoft.Json" Version="13.0.3" />
   </ItemGroup>
 
 </Project>

```

### ✅ 11. powershell — 3424ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo","description":"Rebuild Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Determining projects to restore...
  Restored C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj (in 331 ms).
C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\Program.cs(105,50): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\Program.cs(105,50): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:01.81
<shellId: 2 completed with exit code 1>
```

### ✅ 12. apply_patch — 803ms

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n-    if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+    if (exception.StatusCode == HttpStatusCode.TooManyRequests \u0026\u0026\n+        exception.RetryAfter is TimeSpan retryAfter)\n     {\n         Console.Error.WriteLine(\n-            $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms\");\n+            $\"Retry after: {retryAfter.TotalMilliseconds:F0} ms\");\n     }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660/Program.cs
@@ -99,10 +99,11 @@
     Console.Error.WriteLine($"Substatus code: {exception.SubStatusCode}");
     Console.Error.WriteLine($"Request charge: {exception.RequestCharge:F2} RU");
 
-    if (exception.StatusCode == HttpStatusCode.TooManyRequests)
+    if (exception.StatusCode == HttpStatusCode.TooManyRequests &&
+        exception.RetryAfter is TimeSpan retryAfter)
     {
         Console.Error.WriteLine(
-            $"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms");
+            $"Retry after: {retryAfter.TotalMilliseconds:F0} ms");
     }
 
     Environment.ExitCode = 1;

```

### ✅ 13. powershell — 2534ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo","description":"Confirm sample compiles","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Determining projects to restore...
  All projects are up-to-date for restore.
  CosmosCrudSample -> C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-1591801660\bin\Debug\net8.0\CosmosCrudSample.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.15
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `CosmosCrudSample.csproj`
- `Program.cs`

## Copilot Response

I’ll build this as a local console app using placeholders only, with the Cosmos DB data-plane SDK and no live Azure actions.

The project will use the emulator-compatible connection-string path, a parameterized query, point operations with the partition key, and structured Cosmos diagnostics on failure.

The current Cosmos SDK adds an explicit Newtonsoft.Json requirement, so I’m adding that required package and pinning the resolved Cosmos SDK version for reproducible builds.

Implemented a complete .NET 8 Cosmos DB CRUD console app:

- `Program.cs` — database/container creation, insert, point read, parameterized query, replace, delete, and `CosmosException` diagnostics.
- `CosmosCrudSample.csproj` — required packages:
  - `Microsoft.Azure.Cosmos` 3.62.1
  - `Newtonsoft.Json` 13.0.3

Run against a local Cosmos DB emulator:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-emulator-connection-string>"
dotnet run
```

The connection string is read from the environment and is never embedded in source code.

## Reviewer Notes (LLM-as-Judge)

### Summary

All 7 criteria passed. The solution provides a complete asynchronous Cosmos DB CRUD workflow with parameterized querying, partition-key usage, configuration handling, and detailed CosmosException diagnostics.

### Strengths

- Uses an environment variable instead of embedding the connection string.
- Uses the partition key consistently for create, read, query, replace, and delete operations.
- Uses a parameterized query and handles paginated query results.
- Provides detailed CosmosException diagnostics, including throttling retry information.

### Issues

- Newtonsoft.Json is referenced but unused and is not required by this sample.

## Grader Results

- crud-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - `Microsoft.Azure.Cosmos` NuGet package: Pass
      - `CosmosClient` creation and configuration: Pass
      - `Database.CreateDatabaseIfNotExistsAsync()`: Pass
      - `Container.CreateContainerIfNotExistsAsync()` with partition key: Pass
      - `Container.CreateItemAsync<T>()`, `ReadItemAsync<T>()`, `ReplaceItemAsync<T>()`, `DeleteItemAsync<T>()`: Pass
      - `Container.GetItemQueryIterator<T>()` with `QueryDefinition`: Pass
      - `CosmosException` handling with status codes: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
