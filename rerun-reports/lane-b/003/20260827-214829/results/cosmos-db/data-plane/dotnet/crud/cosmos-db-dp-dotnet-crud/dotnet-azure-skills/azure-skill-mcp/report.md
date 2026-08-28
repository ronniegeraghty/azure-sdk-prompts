# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 349.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 349.4s |
| Timestamp | 2026-08-27T13:48:29Z |
| Files Generated | 3 |
| Event Count | 5258 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 320.6s |
| Review | 27.7s |
| **Total** | **349.4s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=375402 out=3817 |
| Turn Count | 10 |

## Tool Calls

### ❌ 1. skill — 2494ms

**Input:**

```json
{"skill":"azure-cosmos"}
```

**Error:**

```
Skill not found: azure-cosmos
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 187900ms

**Input:**

```json
{"intent":"Generate a local-only C# console application using Microsoft.Azure.Cosmos that demonstrates CRUD operations, parameterized queries, partition keys, and CosmosException handling without creating live Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 187904ms

**Input:**

```json
{"intent":"Find current official Microsoft documentation for .NET Azure Cosmos DB for NoSQL SDK CRUD operations: CosmosClient connection string, CreateDatabaseIfNotExistsAsync, CreateContainerIfNotExistsAsync, CreateItemAsync, ReadItemAsync, parameterized QueryDefinition, ReplaceItemAsync, DeleteItemAsync, and CosmosException handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 187895ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. web_fetch — 3238ms

**Input:**

```json
{"max_length":20000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 6. apply_patch — 3214ms

**Input:**

```json
"*** Begin Patch\n*** Add File: CosmosCrudSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.47.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using System.Net;\n+using System.Text.Json.Serialization;\n+using Microsoft.Azure.Cosmos;\n+\n+const string databaseName = \"TestDB\";\n+const string containerName = \"Items\";\n+const string partitionKeyPath = \"/category\";\n+\n+try\n+{\n+    string connectionString = Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\")\n+        ?? throw new InvalidOperationException(\n+            \"Set COSMOS_CONNECTION_STRING to the Azure Cosmos DB Emulator connection string.\");\n+\n+    EnsureLocalEmulatorConnection(connectionString);\n+\n+    using CosmosClient client = new(\n+        connectionString,\n+        new CosmosClientOptions { ApplicationName = \"CosmosCrudSample\" });\n+\n+    DatabaseResponse databaseResponse =\n+        await client.CreateDatabaseIfNotExistsAsync(databaseName);\n+    Database database = databaseResponse.Database;\n+\n+    ContainerResponse containerResponse =\n+        await database.CreateContainerIfNotExistsAsync(\n+            new ContainerProperties(containerName, partitionKeyPath));\n+    Container container = containerResponse.Container;\n+\n+    Item item = new()\n+    {\n+        Id = Guid.NewGuid().ToString(),\n+        Category = \"electronics\",\n+        Name = \"Wireless keyboard\",\n+        Quantity = 10\n+    };\n+\n+    ItemResponse\u003cItem\u003e createResponse = await container.CreateItemAsync(\n+        item,\n+        new PartitionKey(item.Category));\n+    Console.WriteLine(\n+        $\"Created item {createResponse.Resource.Id}; request charge: {createResponse.RequestCharge} RUs.\");\n+\n+    ItemResponse\u003cItem\u003e readResponse = await container.ReadItemAsync\u003cItem\u003e(\n+        item.Id,\n+        new PartitionKey(item.Category));\n+    Console.WriteLine(\n+        $\"Read item: {readResponse.Resource.Name}, quantity {readResponse.Resource.Quantity}.\");\n+\n+    QueryDefinition query = new(\n+        \"SELECT * FROM items i WHERE i.category = @category\");\n+    query.WithParameter(\"@category\", \"electronics\");\n+\n+    using FeedIterator\u003cItem\u003e results = container.GetItemQueryIterator\u003cItem\u003e(\n+        query,\n+        requestOptions: new QueryRequestOptions\n+        {\n+            PartitionKey = new PartitionKey(\"electronics\")\n+        });\n+\n+    Console.WriteLine(\"Electronics:\");\n+    while (results.HasMoreResults)\n+    {\n+        FeedResponse\u003cItem\u003e page = await results.ReadNextAsync();\n+        foreach (Item result in page)\n+        {\n+            Console.WriteLine($\"- {result.Id}: {result.Name} ({result.Quantity})\");\n+        }\n+    }\n+\n+    Item updatedItem = readResponse.Resource;\n+    updatedItem.Quantity = 25;\n+\n+    ItemResponse\u003cItem\u003e replaceResponse = await container.ReplaceItemAsync(\n+        updatedItem,\n+        updatedItem.Id,\n+        new PartitionKey(updatedItem.Category));\n+    Console.WriteLine(\n+        $\"Updated item {replaceResponse.Resource.Id} to quantity {replaceResponse.Resource.Quantity}.\");\n+\n+    await container.DeleteItemAsync\u003cItem\u003e(\n+        updatedItem.Id,\n+        new PartitionKey(updatedItem.Category));\n+    Console.WriteLine($\"Deleted item {updatedItem.Id}.\");\n+}\n+catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.Conflict)\n+{\n+    ReportCosmosError(\"An item with the same id and partition key already exists.\", exception);\n+    Environment.ExitCode = 1;\n+}\n+catch (CosmosException exception) when (exception.StatusCode == HttpStatusCode.NotFound)\n+{\n+    ReportCosmosError(\"The requested database, container, or item was not found.\", exception);\n+    Environment.ExitCode = 1;\n+}\n+catch (CosmosException exception) when (\n+    exception.StatusCode == HttpStatusCode.TooManyRequests)\n+{\n+    ReportCosmosError(\n+        $\"Request rate was too high. Retry after {exception.RetryAfter.TotalMilliseconds:N0} ms.\",\n+        exception);\n+    Environment.ExitCode = 1;\n+}\n+catch (CosmosException exception)\n+{\n+    ReportCosmosError(\"Azure Cosmos DB operation failed.\", exception);\n+    Environment.ExitCode = 1;\n+}\n+catch (InvalidOperationException exception)\n+{\n+    Console.Error.WriteLine($\"Configuration error: {exception.Message}\");\n+    Environment.ExitCode = 1;\n+}\n+catch (UriFormatException exception)\n+{\n+    Console.Error.WriteLine($\"Invalid Cosmos DB endpoint: {exception.Message}\");\n+    Environment.ExitCode = 1;\n+}\n+\n+static void EnsureLocalEmulatorConnection(string connectionString)\n+{\n+    string? endpoint = connectionString\n+        .Split(';', StringSplitOptions.RemoveEmptyEntries)\n+        .Select(part =\u003e part.Split('=', 2))\n+        .Where(parts =\u003e parts.Length == 2)\n+        .Where(parts =\u003e parts[0].Trim().Equals(\n+            \"AccountEndpoint\",\n+            StringComparison.OrdinalIgnoreCase))\n+        .Select(parts =\u003e parts[1].Trim())\n+        .SingleOrDefault();\n+\n+    if (endpoint is null)\n+    {\n+        throw new InvalidOperationException(\n+            \"COSMOS_CONNECTION_STRING must contain AccountEndpoint.\");\n+    }\n+\n+    Uri endpointUri = new(endpoint, UriKind.Absolute);\n+    if (!endpointUri.IsLoopback)\n+    {\n+        throw new InvalidOperationException(\n+            \"This sample is restricted to a local Azure Cosmos DB Emulator endpoint.\");\n+    }\n+}\n+\n+static void ReportCosmosError(string message, CosmosException exception)\n+{\n+    Console.Error.WriteLine(message);\n+    Console.Error.WriteLine(\n+        $\"Status: {(int)exception.StatusCode} ({exception.StatusCode}); \" +\n+        $\"substatus: {exception.SubStatusCode}; \" +\n+        $\"activity id: {exception.ActivityId}; \" +\n+        $\"request charge: {exception.RequestCharge} RUs.\");\n+}\n+\n+internal sealed class Item\n+{\n+    [JsonPropertyName(\"id\")]\n+    public required string Id { get; init; }\n+\n+    [JsonPropertyName(\"category\")]\n+    public required string Category { get; init; }\n+\n+    [JsonPropertyName(\"name\")]\n+    public required string Name { get; init; }\n+\n+    [JsonPropertyName(\"quantity\")]\n+    public int Quantity { get; set; }\n+}\n*** Add File: README.md\n+# Azure Cosmos DB CRUD sample\n+\n+This .NET 8 console application uses the local Azure Cosmos DB Emulator. It\n+refuses non-loopback endpoints so it cannot modify a live Azure account.\n+\n+## Required NuGet package\n+\n+```powershell\n+dotnet add package Microsoft.Azure.Cosmos --version 3.47.0\n+```\n+\n+The package is already declared in `CosmosCrudSample.csproj`.\n+\n+## Run\n+\n+Start the Azure Cosmos DB Emulator, copy its connection string, and run:\n+\n+```powershell\n+$env:COSMOS_CONNECTION_STRING = \"\u003clocal-emulator-connection-string\u003e\"\n+dotnet run\n+```\n+\n+The program creates `TestDB` and the `Items` container in the emulator, creates\n+an item, reads and queries it, updates its quantity, and deletes it.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj
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
+    <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.47.0" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
@@ -1,0 +1,170 @@
+using System.Net;
+using System.Text.Json.Serialization;
+using Microsoft.Azure.Cosmos;
+
+const string databaseName = "TestDB";
+const string containerName = "Items";
+const string partitionKeyPath = "/category";
+
+try
+{
+    string connectionString = Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING")
+        ?? throw new InvalidOperationException(
+            "Set COSMOS_CONNECTION_STRING to the Azure Cosmos DB Emulator connection string.");
+
+    EnsureLocalEmulatorConnection(connectionString);
+
+    using CosmosClient client = new(
+        connectionString,
+        new CosmosClientOptions { ApplicationName = "CosmosCrud
... (truncated)
```

### ✅ 7. powershell — 5935ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo --verbosity quiet","description":"Build Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.47.0\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(73,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.47.0\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(73,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:01.75
<shellId: 0 completed with exit code 1>
```

### ✅ 8. apply_patch — 2527ms

**Input:**

```json
"*** Begin Patch\n*** Update File: CosmosCrudSample.csproj\n@@\n   \u003cItemGroup\u003e\n     \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.47.0\" /\u003e\n+    \u003cPackageReference Include=\"Newtonsoft.Json\" Version=\"13.0.3\" /\u003e\n   \u003c/ItemGroup\u003e\n*** Update File: Program.cs\n@@\n using System.Net;\n-using System.Text.Json.Serialization;\n using Microsoft.Azure.Cosmos;\n+using Newtonsoft.Json;\n@@\n-    [JsonPropertyName(\"id\")]\n+    [JsonProperty(\"id\")]\n     public required string Id { get; init; }\n \n-    [JsonPropertyName(\"category\")]\n+    [JsonProperty(\"category\")]\n     public required string Category { get; init; }\n \n-    [JsonPropertyName(\"name\")]\n+    [JsonProperty(\"name\")]\n     public required string Name { get; init; }\n \n-    [JsonPropertyName(\"quantity\")]\n+    [JsonProperty(\"quantity\")]\n     public int Quantity { get; set; }\n }\n*** Update File: README.md\n@@\n ```powershell\n dotnet add package Microsoft.Azure.Cosmos --version 3.47.0\n+dotnet add package Newtonsoft.Json --version 13.0.3\n ```\n \n-The package is already declared in `CosmosCrudSample.csproj`.\n+Both packages are already declared in `CosmosCrudSample.csproj`.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/CosmosCrudSample.csproj
@@ -9,6 +9,7 @@
 
   <ItemGroup>
     <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.47.0" />
+    <PackageReference Include="Newtonsoft.Json" Version="13.0.3" />
   </ItemGroup>
 
 </Project>


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
@@ -1,6 +1,6 @@
 using System.Net;
-using System.Text.Json.Serialization;
 using Microsoft.Azure.Cosmos;
+using Newtonsoft.Json;
 
 const string databaseName = "TestDB";
 const string containerName = "Items";
@@ -155,16 +155,16 @@
 
 internal sealed class Item
 {
-    [JsonPropertyName("id")]
+    [JsonProperty("id")]
     public required string Id { get; init; }
 
-    [JsonPropertyName("category")]
+    [JsonProperty("category")]
     public required string Category { get; init; }
 
-    [JsonPropertyName("name")]
+    [JsonProperty("name")]
     public required string Name { get; init; }
 
-    [JsonPropertyName("quantit
... (truncated)
```

### ✅ 9. powershell — 3894ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo --verbosity quiet","description":"Rebuild Cosmos DB CRUD sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\Program.cs(100,72): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\Program.cs(100,72): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:02.28
<shellId: 1 completed with exit code 1>
```

### ✅ 10. apply_patch — 912ms

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n     ReportCosmosError(\n-        $\"Request rate was too high. Retry after {exception.RetryAfter.TotalMilliseconds:N0} ms.\",\n+        $\"Request rate was too high. Retry after \" +\n+        $\"{exception.RetryAfter?.TotalMilliseconds ?? 0:N0} ms.\",\n         exception);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-3905696715/Program.cs
@@ -97,7 +97,8 @@
     exception.StatusCode == HttpStatusCode.TooManyRequests)
 {
     ReportCosmosError(
-        $"Request rate was too high. Retry after {exception.RetryAfter.TotalMilliseconds:N0} ms.",
+        $"Request rate was too high. Retry after " +
+        $"{exception.RetryAfter?.TotalMilliseconds ?? 0:N0} ms.",
         exception);
     Environment.ExitCode = 1;
 }

```

### ✅ 11. powershell — 3123ms

**Input:**

```json
{"command":"dotnet build .\\CosmosCrudSample.csproj --nologo --verbosity quiet","description":"Verify Cosmos DB sample builds","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.57
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `CosmosCrudSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

No Cosmos DB-specific skill is installed, so I’m grounding the implementation in Azure SDK guidance and inspecting the workspace before creating the local project.

The Azure guidance endpoints timed out. I’ll keep the sample emulator-only, use the current stable Cosmos SDK package from NuGet, and implement the complete CRUD flow with parameterized queries and status-aware exception reporting.

The implementation is complete. I’m compiling it now to catch SDK API, serialization, and nullable-reference issues before handing it off.

The build exposed an SDK requirement: this Cosmos SDK version requires an explicit Newtonsoft.Json reference, and its default serializer needs Newtonsoft property attributes to preserve the required lowercase `id` and partition-key fields. I’m correcting both together.

Implemented a complete .NET 8 Azure Cosmos DB CRUD sample:

- `Program.cs` — creates the client, database/container, inserts, reads, queries, replaces, and deletes an item.
- Uses parameterized SQL queries and partition keys correctly.
- Handles `CosmosException` cases including `409`, `404`, `429`, and general failures.
- Restricted to the local Cosmos DB Emulator to prevent changes to live Azure resources.
- `CosmosCrudSample.csproj` includes:
  - `Microsoft.Azure.Cosmos` 3.47.0
  - `Newtonsoft.Json` 13.0.3
- `README.md` contains setup and execution instructions.

## Reviewer Notes (LLM-as-Judge)

### Summary

7/7 criteria passed. The generated .NET 8 sample implements the complete Cosmos DB NoSQL CRUD workflow, includes the required dependencies, and provides status-aware CosmosException handling.

### Strengths

- Complete create, read, query, replace, and delete workflow
- Correct database, container, and partition-key configuration
- Parameterized query with a partition-scoped iterator
- Detailed CosmosException handling for common HTTP status codes
- NuGet dependencies and run instructions are documented

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
