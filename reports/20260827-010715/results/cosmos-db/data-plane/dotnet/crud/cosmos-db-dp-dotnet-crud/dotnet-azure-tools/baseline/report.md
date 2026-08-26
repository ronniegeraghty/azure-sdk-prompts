# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 103.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 103.6s |
| Timestamp | 2026-08-26T17:10:03Z |
| Files Generated | 2 |
| Event Count | 2941 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 80.4s |
| Review | 22.4s |
| **Total** | **103.6s** |

## Configuration

- **name:** dotnet-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=126411 out=1853 |
| Turn Count | 6 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"pattern":"*.{csproj,sln,cs}"}
```

### 🔧 3. glob

**Input:**

```json
{"pattern":"README*"}
```

### 🔧 4. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: CosmosCrudSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.50.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Net;\n+using Microsoft.Azure.Cosmos;\n+\n+const string databaseName = \"TestDB\";\n+const string containerName = \"Items\";\n+const string category = \"electronics\";\n+const string itemId = \"item-001\";\n+\n+string? connectionString =\n+    Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\");\n+\n+if (string.IsNullOrWhiteSpace(connectionString))\n+{\n+    Console.Error.WriteLine(\n+        \"Set the COSMOS_CONNECTION_STRING environment variable before running the sample.\");\n+    return 1;\n+}\n+\n+try\n+{\n+    using CosmosClient client = new(connectionString);\n+\n+    DatabaseResponse databaseResponse =\n+        await client.CreateDatabaseIfNotExistsAsync(databaseName);\n+    Database database = databaseResponse.Database;\n+\n+    ContainerResponse containerResponse =\n+        await database.CreateContainerIfNotExistsAsync(\n+            new ContainerProperties(containerName, \"/category\"));\n+    Container container = containerResponse.Container;\n+\n+    Item item = new()\n+    {\n+        id = itemId,\n+        category = category,\n+        name = \"Wireless headphones\",\n+        quantity = 10\n+    };\n+\n+    ItemResponse\u003cItem\u003e createResponse = await container.CreateItemAsync(\n+        item,\n+        new PartitionKey(item.category));\n+    Console.WriteLine(\n+        $\"Created {createResponse.Resource.id}; RU charge: {createResponse.RequestCharge:F2}\");\n+\n+    ItemResponse\u003cItem\u003e readResponse = await container.ReadItemAsync\u003cItem\u003e(\n+        itemId,\n+        new PartitionKey(category));\n+    Console.WriteLine(\n+        $\"Read {readResponse.Resource.name}, quantity {readResponse.Resource.quantity}\");\n+\n+    QueryDefinition query = new(\n+        \"SELECT * FROM items i WHERE i.category = @category\");\n+    query.WithParameter(\"@category\", category);\n+\n+    using FeedIterator\u003cItem\u003e queryIterator = container.GetItemQueryIterator\u003cItem\u003e(\n+        query,\n+        requestOptions: new QueryRequestOptions\n+        {\n+            PartitionKey = new PartitionKey(category)\n+        });\n+\n+    while (queryIterator.HasMoreResults)\n+    {\n+        FeedResponse\u003cItem\u003e page = await queryIterator.ReadNextAsync();\n+        foreach (Item result in page)\n+        {\n+            Console.WriteLine(\n+                $\"Query result: {result.id}, {result.name}, quantity {result.quantity}\");\n+        }\n+    }\n+\n+    item.quantity = 25;\n+    ItemResponse\u003cItem\u003e replaceResponse = await container.ReplaceItemAsync(\n+        item,\n+        item.id,\n+        new PartitionKey(item.category));\n+    Console.WriteLine(\n+        $\"Updated {replaceResponse.Resource.id} to quantity {replaceResponse.Resource.quantity}\");\n+\n+    ItemResponse\u003cItem\u003e deleteResponse = await container.DeleteItemAsync\u003cItem\u003e(\n+        item.id,\n+        new PartitionKey(item.category));\n+    Console.WriteLine(\n+        $\"Deleted {item.id}; status: {deleteResponse.StatusCode}\");\n+\n+    return 0;\n+}\n+catch (CosmosException exception)\n+{\n+    Console.Error.WriteLine(\n+        $\"Cosmos DB request failed ({(int)exception.StatusCode} {exception.StatusCode}).\");\n+    Console.Error.WriteLine($\"Message: {exception.Message}\");\n+    Console.Error.WriteLine($\"Activity ID: {exception.ActivityId}\");\n+    Console.Error.WriteLine($\"Request charge: {exception.RequestCharge:F2} RU\");\n+\n+    if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+    {\n+        Console.Error.WriteLine($\"Retry after: {exception.RetryAfter}\");\n+    }\n+\n+    return 1;\n+}\n+catch (ArgumentException exception)\n+{\n+    Console.Error.WriteLine($\"Invalid Cosmos DB configuration: {exception.Message}\");\n+    return 1;\n+}\n+\n+internal sealed class Item\n+{\n+    public required string id { get; init; }\n+\n+    public required string category { get; init; }\n+\n+    public required string name { get; init; }\n+\n+    public int quantity { get; set; }\n+}\n*** End Patch\n"
```

### 🔧 5. powershell

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo","description":"Build Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

### 🔧 6. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: CosmosCrudSample.csproj\n@@\n   \u003cItemGroup\u003e\n     \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.50.0\" /\u003e\n+    \u003cPackageReference Include=\"Newtonsoft.Json\" Version=\"13.0.3\" /\u003e\n   \u003c/ItemGroup\u003e\n*** End Patch\n"
```

### 🔧 7. powershell

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo --no-restore","description":"Rebuild Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `CosmosCrudSample.csproj`
- `Program.cs`

## Copilot Response

I’ll inspect the workspace, then create a local-only sample with the required package reference and Cosmos DB CRUD flow.

The workspace is empty. I’ll add a .NET 8 console project using `Microsoft.Azure.Cosmos`, with the connection string read from an environment variable so no credentials are embedded.

The sample is implemented. I’ll now compile it to catch SDK/API or nullable-type issues without connecting to Azure.

The Cosmos SDK requires an explicit `Newtonsoft.Json` reference in this project configuration. I’ll add that required package and rebuild.

Implemented a complete .NET 8 Cosmos DB CRUD sample:

- `Program.cs` — creates the database/container, inserts, reads, queries, replaces, and deletes an item.
- `CosmosCrudSample.csproj` — references:
  - `Microsoft.Azure.Cosmos` 3.50.0
  - `Newtonsoft.Json` 13.0.3
- Handles `CosmosException`, including status code, activity ID, RU charge, and throttling retry delay.
- Reads credentials securely from `COSMOS_CONNECTION_STRING`.

Run locally with:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-connection-string>"
dotnet run
```

## Reviewer Notes (LLM-as-Judge)

### Summary

All 7 criteria passed. The sample provides the requested Cosmos DB CRUD workflow, parameterized category query, package references, and CosmosException handling.

### Strengths

- Uses the current Microsoft.Azure.Cosmos v3 API throughout.
- Keeps the connection string outside source code in an environment variable.
- Uses parameterized SQL and scopes the query to the category partition.
- Reports useful Cosmos DB diagnostics, including status code, activity ID, request charge, and throttling delay.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
