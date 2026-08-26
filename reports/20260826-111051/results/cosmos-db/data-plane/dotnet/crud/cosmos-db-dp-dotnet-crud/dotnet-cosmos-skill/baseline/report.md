# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-cosmos-skill/baseline | **Result:** ✅ PASSED | **Duration:** 148.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-cosmos-skill/baseline |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 148.3s |
| Timestamp | 2026-08-26T03:15:45Z |
| Files Generated | 3 |
| Event Count | 3750 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 112.7s |
| Review | 34.3s |
| **Total** | **148.3s** |

## Configuration

- **name:** dotnet-cosmos-skill/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=270321 out=3217 |
| Turn Count | 7 |

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
{"intent":"Generate a local-only C# console program using Microsoft.Azure.Cosmos for NoSQL item CRUD with connection-string authentication, partition key usage, parameterized queries, and CosmosException handling. Do not provision live resources.","learn":true}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-cosmos-skill-baseline-2251974249","pattern":"**/*"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a local-only C# Cosmos DB NoSQL CRUD console sample using the emulator and Microsoft.Azure.Cosmos.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"dotnet --version; dotnet package search Microsoft.Azure.Cosmos --exact-match --take 1 --format json","description":"Check .NET and latest Cosmos package","initial_wait":120,"mode":"sync"}
```

### 🔧 6. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: CosmosCrudSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.62.1\" /\u003e\n+    \u003cPackageReference Include=\"Newtonsoft.Json\" Version=\"13.0.4\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Net;\n+using Microsoft.Azure.Cosmos;\n+using Newtonsoft.Json.Linq;\n+\n+internal static class Program\n+{\n+    private const string DatabaseName = \"TestDB\";\n+    private const string ContainerName = \"Items\";\n+    private const string PartitionKeyPath = \"/category\";\n+\n+    public static async Task\u003cint\u003e Main()\n+    {\n+        try\n+        {\n+            string connectionString =\n+                Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\")\n+                ?? throw new InvalidOperationException(\n+                    \"Set COSMOS_CONNECTION_STRING to the Azure Cosmos DB Emulator connection string.\");\n+\n+            using CosmosClient client = new(\n+                connectionString,\n+                new CosmosClientOptions\n+                {\n+                    ApplicationName = \"CosmosCrudSample\"\n+                });\n+\n+            DatabaseResponse databaseResponse =\n+                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);\n+            Database database = databaseResponse.Database;\n+\n+            ContainerResponse containerResponse =\n+                await database.CreateContainerIfNotExistsAsync(\n+                    new ContainerProperties(ContainerName, PartitionKeyPath));\n+            Container container = containerResponse.Container;\n+\n+            const string itemId = \"item-001\";\n+            const string category = \"electronics\";\n+            PartitionKey partitionKey = new(category);\n+\n+            JObject item = JObject.FromObject(new\n+            {\n+                id = itemId,\n+                category,\n+                name = \"Wireless Mouse\",\n+                quantity = 10\n+            });\n+\n+            ItemResponse\u003cJObject\u003e createResponse =\n+                await container.CreateItemAsync(item, partitionKey);\n+            Console.WriteLine(\n+                $\"Created {createResponse.Resource[\"id\"]} \" +\n+                $\"(RU: {createResponse.RequestCharge:F2})\");\n+\n+            ItemResponse\u003cJObject\u003e readResponse =\n+                await container.ReadItemAsync\u003cJObject\u003e(itemId, partitionKey);\n+            Console.WriteLine($\"Read: {readResponse.Resource}\");\n+\n+            QueryDefinition queryDefinition = new(\n+                \"SELECT * FROM c WHERE c.category = @category\");\n+            queryDefinition.WithParameter(\"@category\", category);\n+\n+            using FeedIterator\u003cJObject\u003e query =\n+                container.GetItemQueryIterator\u003cJObject\u003e(\n+                    queryDefinition,\n+                    requestOptions: new QueryRequestOptions\n+                    {\n+                        PartitionKey = partitionKey\n+                    });\n+\n+            while (query.HasMoreResults)\n+            {\n+                FeedResponse\u003cJObject\u003e page = await query.ReadNextAsync();\n+                foreach (JObject queryItem in page)\n+                {\n+                    Console.WriteLine($\"Query result: {queryItem}\");\n+                }\n+            }\n+\n+            item[\"quantity\"] = 25;\n+            ItemResponse\u003cJObject\u003e replaceResponse =\n+                await container.ReplaceItemAsync(item, itemId, partitionKey);\n+            Console.WriteLine(\n+                $\"Updated quantity to {replaceResponse.Resource[\"quantity\"]} \" +\n+                $\"(RU: {replaceResponse.RequestCharge:F2})\");\n+\n+            ItemResponse\u003cJObject\u003e deleteResponse =\n+                await container.DeleteItemAsync\u003cJObject\u003e(itemId, partitionKey);\n+            Console.WriteLine(\n+                $\"Deleted {itemId} (status: {deleteResponse.StatusCode}, \" +\n+                $\"RU: {deleteResponse.RequestCharge:F2})\");\n+\n+            return 0;\n+        }\n+        catch (CosmosException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Cosmos DB request failed: {exception.StatusCode} \" +\n+                $\"({(int)exception.StatusCode}), substatus {exception.SubStatusCode}.\");\n+            Console.Error.WriteLine($\"Activity ID: {exception.ActivityId}\");\n+            Console.Error.WriteLine($\"Request charge: {exception.RequestCharge:F2} RU\");\n+            Console.Error.WriteLine(exception.Message);\n+\n+            if (exception.StatusCode == HttpStatusCode.Conflict)\n+            {\n+                Console.Error.WriteLine(\n+                    \"The item already exists. Delete item-001 or use a different id.\");\n+            }\n+\n+            return 1;\n+        }\n+        catch (InvalidOperationException exception)\n+        {\n+            Console.Error.WriteLine($\"Configuration error: {exception.Message}\");\n+            return 2;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB NoSQL CRUD sample\n+\n+This .NET 8 console application creates `TestDB` and its `Items` container,\n+then creates, reads, queries, replaces, and deletes one item. The container\n+uses `/category` as its partition key.\n+\n+## Required NuGet packages\n+\n+```powershell\n+dotnet add package Microsoft.Azure.Cosmos --version 3.62.1\n+dotnet add package Newtonsoft.Json --version 13.0.4\n+```\n+\n+`Microsoft.Azure.Cosmos` is the Azure Cosmos DB NoSQL SDK. `Newtonsoft.Json`\n+provides `JObject`, which lets the sample work directly with JSON documents.\n+\n+## Run locally\n+\n+Start the Azure Cosmos DB Emulator, copy its connection string, and set it in\n+an environment variable. Do not commit the connection string.\n+\n+```powershell\n+$env:COSMOS_CONNECTION_STRING = \"\u003cyour-Cosmos-DB-Emulator-connection-string\u003e\"\n+dotnet run\n+```\n+\n+The query uses a parameter rather than interpolating user input:\n+\n+```sql\n+SELECT * FROM c WHERE c.category = @category\n+```\n+\n+References:\n+\n+- [Azure Cosmos DB .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/sdk-dotnet-v3)\n+- [Work with items using the .NET SDK](https://learn.microsoft.com/azure/cosmos-db/nosql/how-to-dotnet-create-item)\n+- [Azure Cosmos DB Emulator](https://learn.microsoft.com/azure/cosmos-db/emulator)\n*** End Patch\n"
```

### 🔧 7. powershell

**Input:**

```json
{"command":"dotnet build --nologo --verbosity minimal","description":"Build Cosmos CRUD sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `CosmosCrudSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll inspect the workspace and ground the implementation in current Azure .NET guidance before creating the runnable sample.

The workspace is empty, so I’ll create a minimal .NET console project with the Cosmos SDK package, emulator-friendly configuration, and the full CRUD sequence.

The sample is implemented; I’m compiling it now to catch SDK/API or package compatibility issues.

Implemented a complete .NET 8 Cosmos DB NoSQL CRUD sample:

- `Program.cs` — creates the client, database/container, then inserts, reads, queries, replaces, and deletes a JSON item.
- `CosmosCrudSample.csproj` — references `Microsoft.Azure.Cosmos 3.62.1` and `Newtonsoft.Json 13.0.4`.
- `README.md` — includes emulator setup, NuGet commands, and documentation links.

Run locally with:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-Cosmos-DB-Emulator-connection-string>"
dotnet run
```

The implementation uses a parameterized SQL query, correct partition-key routing, client disposal, and detailed `CosmosException` handling.

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 7 criteria passed. The sample completely implements the requested Cosmos DB NoSQL CRUD workflow, query, dependencies, and error handling.

### Strengths

- Uses parameterized SQL-like querying and correct partition-key routing.
- Includes all required CRUD APIs and database/container creation operations.
- Provides explicit NuGet dependencies and detailed CosmosException diagnostics.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-cosmos-skill/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
