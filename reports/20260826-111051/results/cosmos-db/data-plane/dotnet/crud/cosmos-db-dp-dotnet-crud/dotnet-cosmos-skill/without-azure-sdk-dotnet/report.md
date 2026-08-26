# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-cosmos-skill/without-azure-sdk-dotnet | **Result:** ✅ PASSED | **Duration:** 293.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-cosmos-skill/without-azure-sdk-dotnet |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 293.4s |
| Timestamp | 2026-08-26T03:10:51Z |
| Files Generated | 2 |
| Event Count | 4606 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 241.4s |
| Review | 50.9s |
| **Total** | **293.4s** |

## Configuration

- **name:** dotnet-cosmos-skill/without-azure-sdk-dotnet
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=715362 out=4317 |
| Turn Count | 17 |

## Prompt Sent

```
Write a C# program that performs
CRUD operations on items in an Azure Cosmos DB NoSQL container:
1. Create a CosmosClient using a connection string
2. Create a database named "TestDB" and a container named "Items" with partition key "/category"
3. Insert a JSON item with properties: id, category, name, and quantity
4. Read the item back by id and partition key
5. Query items where category equals "electronics" using SQL-like syntax
6. Replace the item with updated quantity
7. Delete the item

Show required NuGet packages and proper error handling with CosmosException.
```

## Tool Calls

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local C# console program using Microsoft.Azure.Cosmos that demonstrates CRUD operations against an Azure Cosmos DB for NoSQL container with CosmosException handling. Do not provision or connect to live Azure resources.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a local C# Cosmos DB CRUD sample without provisioning resources.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705","pattern":"**/*"}
```

### 🔧 4. powershell

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK version","initial_wait":30,"mode":"sync"}
```

### 🔧 5. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Microsoft documentation for .NET Azure Cosmos DB for NoSQL point CRUD operations, parameterized SQL queries, and CosmosException handling.","learn":true}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve latest official C# Cosmos DB NoSQL CRUD SDK patterns.","parameters":{"language":"csharp","query":"Microsoft.Azure.Cosmos C# CosmosClient CreateDatabaseIfNotExistsAsync CreateContainerIfNotExistsAsync CreateItemAsync ReadItemAsync QueryDefinition ReplaceItemAsync DeleteItemAsync CosmosException"}}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Cosmos DB .NET SDK CRUD documentation.","parameters":{"query":"Azure Cosmos DB for NoSQL .NET SDK CRUD items parameterized query CosmosException"}}
```

### 🔧 9. rg

**Input:**

```json
{"head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787713951241-copilot-tool-output-35d8aef913a1493eba32bc58a98869ff.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\]+"}
```

### 🔧 10. rg

**Input:**

```json
{"head_limit":10,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787713951985-copilot-tool-output-a9c3c91a5ec94ac4a1179bcaa65855c1.txt","pattern":"https://learn\\.microsoft\\.com[^\"\\\\]+"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"$p='C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787713951241-copilot-tool-output-35d8aef913a1493eba32bc58a98869ff.txt'; (Get-Content -Raw $p | ConvertFrom-Json).results | Select-Object -First 5 title,contentUrl | Format-Table -AutoSize","description":"Extract Microsoft documentation result URLs","initial_wait":30,"mode":"sync"}
```

### 🔧 12. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787713951241-copilot-tool-output-35d8aef913a1493eba32bc58a98869ff.txt","view_range":[1,5]}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the official .NET Cosmos DB quickstart details used as the implementation reference.","parameters":{"url":"https://learn.microsoft.com/azure/cosmos-db/quickstart-dotnet#code-examples"}}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"dotnet new console --framework net8.0 --use-program-main --force","description":"Create .NET console project","initial_wait":120,"mode":"sync"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"dotnet new console --framework net9.0 --use-program-main --force \u0026\u0026 dotnet add package Microsoft.Azure.Cosmos --version 3.*","description":"Create project and add Cosmos SDK","initial_wait":180,"mode":"sync"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\Program.cs"}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705.csproj"}
```

### 🔧 18. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Delete File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\Program.cs\n*** Add File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\Program.cs\n+using System.Net;\n+using Microsoft.Azure.Cosmos;\n+\n+namespace CosmosCrudSample;\n+\n+internal static class Program\n+{\n+    private const string DatabaseName = \"TestDB\";\n+    private const string ContainerName = \"Items\";\n+    private const string PartitionKeyPath = \"/category\";\n+\n+    public static async Task\u003cint\u003e Main()\n+    {\n+        string? connectionString =\n+            Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\");\n+\n+        if (string.IsNullOrWhiteSpace(connectionString))\n+        {\n+            Console.Error.WriteLine(\n+                \"Set the COSMOS_CONNECTION_STRING environment variable before running the program.\");\n+            return 1;\n+        }\n+\n+        try\n+        {\n+            using CosmosClient client = new(\n+                connectionString,\n+                new CosmosClientOptions\n+                {\n+                    ApplicationName = \"CosmosCrudSample\"\n+                });\n+\n+            DatabaseResponse databaseResponse =\n+                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);\n+            Database database = databaseResponse.Database;\n+\n+            ContainerResponse containerResponse =\n+                await database.CreateContainerIfNotExistsAsync(\n+                    id: ContainerName,\n+                    partitionKeyPath: PartitionKeyPath);\n+            Container container = containerResponse.Container;\n+\n+            Console.WriteLine(\n+                $\"Using database '{DatabaseName}' and container '{ContainerName}'.\");\n+\n+            CosmosItem item = new(\n+                id: Guid.NewGuid().ToString(),\n+                category: \"electronics\",\n+                name: \"Wireless keyboard\",\n+                quantity: 10);\n+            PartitionKey partitionKey = new(item.category);\n+\n+            ItemResponse\u003cCosmosItem\u003e createResponse =\n+                await container.CreateItemAsync(item, partitionKey);\n+            Console.WriteLine(\n+                $\"Created item '{createResponse.Resource.id}' \" +\n+                $\"({createResponse.RequestCharge:F2} RU).\");\n+\n+            ItemResponse\u003cCosmosItem\u003e readResponse =\n+                await container.ReadItemAsync\u003cCosmosItem\u003e(item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Read item: {readResponse.Resource.name}, \" +\n+                $\"quantity {readResponse.Resource.quantity}.\");\n+\n+            QueryDefinition query = new(\n+                \"SELECT * FROM c WHERE c.category = @category\");\n+            query.WithParameter(\"@category\", \"electronics\");\n+\n+            using FeedIterator\u003cCosmosItem\u003e iterator =\n+                container.GetItemQueryIterator\u003cCosmosItem\u003e(query);\n+\n+            Console.WriteLine(\"Electronics items:\");\n+            while (iterator.HasMoreResults)\n+            {\n+                FeedResponse\u003cCosmosItem\u003e page = await iterator.ReadNextAsync();\n+                foreach (CosmosItem result in page)\n+                {\n+                    Console.WriteLine(\n+                        $\"- {result.id}: {result.name}, quantity {result.quantity}\");\n+                }\n+            }\n+\n+            CosmosItem updatedItem = item with { quantity = 25 };\n+            ItemResponse\u003cCosmosItem\u003e replaceResponse =\n+                await container.ReplaceItemAsync(\n+                    updatedItem,\n+                    updatedItem.id,\n+                    partitionKey);\n+            Console.WriteLine(\n+                $\"Updated quantity to {replaceResponse.Resource.quantity}.\");\n+\n+            await container.DeleteItemAsync\u003cCosmosItem\u003e(item.id, partitionKey);\n+            Console.WriteLine($\"Deleted item '{item.id}'.\");\n+\n+            return 0;\n+        }\n+        catch (CosmosException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Cosmos DB request failed with status {(int)exception.StatusCode} \" +\n+                $\"({exception.StatusCode}).\");\n+            Console.Error.WriteLine($\"Message: {exception.Message}\");\n+            Console.Error.WriteLine($\"Activity ID: {exception.ActivityId}\");\n+            Console.Error.WriteLine($\"Request charge: {exception.RequestCharge:F2} RU\");\n+\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            {\n+                Console.Error.WriteLine(\n+                    $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms\");\n+            }\n+\n+            return 1;\n+        }\n+    }\n+}\n+\n+internal sealed record CosmosItem(\n+    string id,\n+    string category,\n+    string name,\n+    int quantity);\n*** End Patch\n"
```

### 🔧 19. powershell

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Build Cosmos CRUD sample","initial_wait":180,"mode":"sync"}
```

### 🔧 20. powershell

**Input:**

```json
{"command":"dotnet add package Newtonsoft.Json \u0026\u0026 dotnet build --no-restore","description":"Add JSON dependency and rebuild","initial_wait":180,"mode":"sync"}
```

### 🔧 21. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\Program.cs\n@@\n-            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests \u0026\u0026\n+                exception.RetryAfter is TimeSpan retryAfter)\n             {\n                 Console.Error.WriteLine(\n-                    $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms\");\n+                    $\"Retry after: {retryAfter.TotalMilliseconds:F0} ms\");\n             }\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705.csproj\n@@\n-    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.*\" /\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.62.1\" /\u003e\n*** End Patch\n"
```

### 🔧 22. powershell

**Input:**

```json
{"command":"dotnet restore --nologo \u0026\u0026 dotnet build --no-restore --nologo; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; Remove-Item Env:COSMOS_CONNECTION_STRING -ErrorAction SilentlyContinue; dotnet run --no-build; if ($LASTEXITCODE -ne 1) { Write-Error \"Expected missing-configuration exit code 1, got $LASTEXITCODE\"; exit 1 }; exit 0","description":"Build and validate configuration handling","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `Program.cs`
- `hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705.csproj`

## Copilot Response

I’ll inspect the workspace, then add the smallest runnable console project with configuration kept outside source code.

The installed SDK supports .NET 9/10, so I’m targeting .NET 9 to match the current Microsoft quickstart and adding the v3 Cosmos SDK.

The Cosmos SDK requires an explicit Newtonsoft.Json reference in this target framework; I’m adding that required dependency before rebuilding.

Implemented the .NET 9 console application in `Program.cs`. It creates `TestDB`/`Items`, performs all requested CRUD operations, runs a parameterized category query, and handles `CosmosException` diagnostics including throttling.

**Required NuGet packages:**
```powershell
dotnet add package Microsoft.Azure.Cosmos --version 3.62.1
dotnet add package Newtonsoft.Json --version 13.0.4
```

Run using a connection string without hardcoding credentials:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-connection-string>"
dotnet run
```

Project file: `hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-without-azure-sdk-dotnet-1953959705.csproj`

Reference: [Azure Cosmos DB for NoSQL .NET quickstart](https://learn.microsoft.com/azure/cosmos-db/quickstart-dotnet#code-examples)

## Reviewer Notes (LLM-as-Judge)

### Summary

All 7 criteria passed. The generated application covers database and container provisioning, the complete item CRUD lifecycle, a parameterized category query, required package setup, and Cosmos DB-specific error handling.

### Strengths

- Uses an environment variable instead of hardcoding the connection string.
- Uses the correct /category partition key consistently for item operations.
- Uses a parameterized query and processes all feed pages.
- Reports useful CosmosException diagnostics, including throttling retry guidance.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-cosmos-skill/without-azure-sdk-dotnet
```

---

[← Back to Summary](../../../../../../summary.md)
