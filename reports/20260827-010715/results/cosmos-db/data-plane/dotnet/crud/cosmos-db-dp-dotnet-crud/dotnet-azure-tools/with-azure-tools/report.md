# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 164.5s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 164.5s |
| Timestamp | 2026-08-26T17:15:30Z |
| Files Generated | 2 |
| Event Count | 4232 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 136.9s |
| Review | 26.4s |
| **Total** | **164.5s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=513030 out=3831 |
| Turn Count | 12 |

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
{"intent":"Generate a local-only C# console application using Microsoft.Azure.Cosmos that demonstrates database/container creation and item CRUD with CosmosException handling. Do not provision or connect to real Azure resources; use an environment-variable placeholder or Cosmos DB emulator.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation best practices for a local C# Microsoft.Azure.Cosmos CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268","pattern":"**/*"}
```

### 🔧 4. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos .NET SDK CRUD operations, CreateDatabaseIfNotExistsAsync, CreateContainerIfNotExistsAsync, parameterized QueryDefinition, ReplaceItemAsync, DeleteItemAsync, and CosmosException handling.","learn":true}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate official Cosmos DB for NoSQL .NET CRUD guidance and API references.","parameters":{"query":"Azure Cosmos DB for NoSQL .NET SDK v3 CRUD CreateDatabaseIfNotExistsAsync CreateContainerIfNotExistsAsync ReadItemAsync ReplaceItemAsync DeleteItemAsync CosmosException"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# Microsoft.Azure.Cosmos CRUD and parameterized query samples.","parameters":{"language":"csharp","query":"Microsoft.Azure.Cosmos C# CRUD QueryDefinition WithParameter CosmosException"}}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"dotnet --version; dotnet new console --framework net8.0 --use-program-main --no-restore","description":"Create .NET 8 console project","initial_wait":30,"mode":"sync"}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"dotnet new console --framework net10.0 --use-program-main --no-restore; if ($LASTEXITCODE -eq 0) { dotnet add package Microsoft.Azure.Cosmos --version 3.58.0 --no-restore }","description":"Create project and add Cosmos package","initial_wait":120,"mode":"sync"}
```

### 🔧 10. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268\\Program.cs"}
```

### 🔧 11. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268.csproj"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Delete File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268\\Program.cs\n*** Add File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268\\Program.cs\n+using Microsoft.Azure.Cosmos;\n+using System.Net;\n+\n+namespace hyoka_cosmos_db_dp_dotnet_crud_dotnet_azure_tools_with_azure_tools_3348390268;\n+\n+internal static class Program\n+{\n+    private const string DatabaseName = \"TestDB\";\n+    private const string ContainerName = \"Items\";\n+    private const string Category = \"electronics\";\n+\n+    private static async Task\u003cint\u003e Main()\n+    {\n+        try\n+        {\n+            string connectionString =\n+                Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\")\n+                ?? throw new InvalidOperationException(\n+                    \"Set COSMOS_CONNECTION_STRING to a Cosmos DB connection string.\");\n+\n+            using CosmosClient client = new(\n+                connectionString,\n+                new CosmosClientOptions\n+                {\n+                    ApplicationName = \"CosmosCrudSample\"\n+                });\n+\n+            DatabaseResponse databaseResponse =\n+                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);\n+            Database database = databaseResponse.Database;\n+\n+            ContainerResponse containerResponse =\n+                await database.CreateContainerIfNotExistsAsync(\n+                    ContainerName,\n+                    partitionKeyPath: \"/category\");\n+            Container container = containerResponse.Container;\n+\n+            Console.WriteLine(\n+                $\"Using database '{DatabaseName}' and container '{ContainerName}'.\");\n+\n+            Item item = new()\n+            {\n+                id = Guid.NewGuid().ToString(),\n+                category = Category,\n+                name = \"Wireless headphones\",\n+                quantity = 5\n+            };\n+            PartitionKey partitionKey = new(item.category);\n+\n+            ItemResponse\u003cItem\u003e createResponse =\n+                await container.CreateItemAsync(item, partitionKey);\n+            Console.WriteLine(\n+                $\"Created item {createResponse.Resource.id} \" +\n+                $\"({createResponse.RequestCharge:F2} RU).\");\n+\n+            ItemResponse\u003cItem\u003e readResponse =\n+                await container.ReadItemAsync\u003cItem\u003e(item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Read item: {readResponse.Resource.name}, \" +\n+                $\"quantity {readResponse.Resource.quantity}.\");\n+\n+            QueryDefinition query = new(\n+                \"SELECT * FROM c WHERE c.category = @category\");\n+            query.WithParameter(\"@category\", Category);\n+\n+            using FeedIterator\u003cItem\u003e results =\n+                container.GetItemQueryIterator\u003cItem\u003e(\n+                    query,\n+                    requestOptions: new QueryRequestOptions\n+                    {\n+                        PartitionKey = partitionKey\n+                    });\n+\n+            while (results.HasMoreResults)\n+            {\n+                FeedResponse\u003cItem\u003e page = await results.ReadNextAsync();\n+\n+                foreach (Item result in page)\n+                {\n+                    Console.WriteLine(\n+                        $\"Query result: {result.id} - {result.name} \" +\n+                        $\"(quantity {result.quantity}).\");\n+                }\n+            }\n+\n+            item.quantity = 10;\n+            ItemResponse\u003cItem\u003e replaceResponse =\n+                await container.ReplaceItemAsync(item, item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Updated item quantity to {replaceResponse.Resource.quantity} \" +\n+                $\"({replaceResponse.RequestCharge:F2} RU).\");\n+\n+            ItemResponse\u003cItem\u003e deleteResponse =\n+                await container.DeleteItemAsync\u003cItem\u003e(item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Deleted item {item.id} \" +\n+                $\"({deleteResponse.RequestCharge:F2} RU).\");\n+\n+            return 0;\n+        }\n+        catch (CosmosException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Cosmos DB request failed with HTTP {(int)exception.StatusCode} \" +\n+                $\"({exception.StatusCode}), substatus {exception.SubStatusCode}.\");\n+            Console.Error.WriteLine($\"Activity ID: {exception.ActivityId}\");\n+\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            {\n+                Console.Error.WriteLine(\n+                    $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms.\");\n+            }\n+\n+            Console.Error.WriteLine(exception.Message);\n+            return 1;\n+        }\n+        catch (InvalidOperationException exception)\n+        {\n+            Console.Error.WriteLine($\"Configuration error: {exception.Message}\");\n+            return 2;\n+        }\n+        catch (ArgumentException exception)\n+        {\n+            Console.Error.WriteLine($\"Invalid Cosmos DB configuration: {exception.Message}\");\n+            return 2;\n+        }\n+    }\n+}\n+\n+internal sealed class Item\n+{\n+    public required string id { get; init; }\n+\n+    public required string category { get; init; }\n+\n+    public required string name { get; init; }\n+\n+    public int quantity { get; set; }\n+}\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"dotnet build --nologo","description":"Build Cosmos CRUD sample","initial_wait":180,"mode":"sync"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"dotnet add package Newtonsoft.Json --version 13.0.4; if ($LASTEXITCODE -eq 0) { dotnet build --nologo --no-restore }","description":"Add JSON dependency and rebuild","initial_wait":180,"mode":"sync"}
```

### 🔧 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268\\Program.cs\n@@\n-            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests\n+                \u0026\u0026 exception.RetryAfter is TimeSpan retryAfter)\n             {\n                 Console.Error.WriteLine(\n-                    $\"Retry after: {exception.RetryAfter.TotalMilliseconds:F0} ms.\");\n+                    $\"Retry after: {retryAfter.TotalMilliseconds:F0} ms.\");\n             }\n*** End Patch\n"
```

### 🔧 16. powershell

**Input:**

```json
{"command":"dotnet build --nologo --no-restore; if ($LASTEXITCODE -eq 0) { Remove-Item Env:COSMOS_CONNECTION_STRING -ErrorAction SilentlyContinue; dotnet run --no-build; if ($LASTEXITCODE -eq 2) { exit 0 } else { exit $LASTEXITCODE } }","description":"Build and check configuration handling","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `Program.cs`
- `hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3348390268.csproj`

## Copilot Response

I’ll build this as a local .NET console sample with placeholder configuration, Cosmos SDK CRUD flow, and targeted exception handling.

The workspace is empty, so I’m creating a self-contained console project and grounding the API usage in the current official Cosmos DB .NET guidance.

The project target is now aligned with the installed SDK; I’m wiring the seven operations in order, using a partition-scoped parameterized query and preserving Cosmos diagnostics on failures.

The Cosmos SDK’s build target requires an explicit Newtonsoft.Json reference, so I’m adding that required package rather than bypassing the compatibility check.

Implemented the Cosmos DB CRUD console app in `Program.cs`.

**Required NuGet packages:**
```xml
<PackageReference Include="Microsoft.Azure.Cosmos" Version="3.58.0" />
<PackageReference Include="Newtonsoft.Json" Version="13.0.4" />
```

The program creates `TestDB/Items`, inserts, reads, queries with parameterized SQL, replaces, and deletes an item. It reads the connection string securely from `COSMOS_CONNECTION_STRING` and reports `CosmosException` status, substatus, activity ID, and throttling retry delay.

```powershell
$env:COSMOS_CONNECTION_STRING = "<Cosmos DB or emulator connection string>"
dotnet run
```

References: [Cosmos DB .NET quickstart](https://learn.microsoft.com/azure/cosmos-db/quickstart-dotnet) and [querying items with the .NET SDK](https://learn.microsoft.com/azure/cosmos-db/how-to-dotnet-query-items).

## Reviewer Notes (LLM-as-Judge)

### Summary

7/7 criteria passed. The implementation fully covers the requested Cosmos DB CRUD workflow, parameterized category query, package declaration, and CosmosException handling.

### Strengths

- Uses the current Microsoft.Azure.Cosmos SDK with correctly configured client, database, container, and partition key operations.
- Uses a parameterized SQL query and iterates through all result pages.
- Provides detailed CosmosException diagnostics, including explicit handling information for HTTP 429 throttling.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
