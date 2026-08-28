# Evaluation Report: cosmos-db-dp-dotnet-pagination

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 520.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-pagination` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 520.1s |
| Timestamp | 2026-08-27T14:20:12Z |
| Files Generated | 3 |
| Event Count | 4527 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 494.8s |
| Review | 24.5s |
| **Total** | **520.1s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-resource-manager-cosmosdb-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=351031 out=3889 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 4465ms

**Input:**

```json
{"skill":"azure-resource-manager-cosmosdb-dotnet"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure.ResourceManager.CosmosDB (.NET)

Management plane SDK for provisioning and managing Azure Cosmos DB resources via Azure Resource Manager.

> **⚠️ Management vs Data Plane**
> - **This SDK (Azure.ResourceManager.CosmosDB)**: Create accounts, databases, containers, configure throughput, manage RBAC
> - **Data Plane SDK (Microsoft.Azure.Cosmos)**: CRUD operations on documents, queries, stored procedures execution

## Installation

```bash
dotnet add package Azure.ResourceManager.CosmosDB
dotnet add package Azure.Identity
```

**Current Versions**: Stable v1.4.0, Preview v1.4.0-beta.13

## Environment Variables

```bash
AZURE_SUBSCRIPTION_ID=<your-subscription-id> # Required: Azure subscription ID
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
AZURE_TENANT_ID=<tenant-id> # For service principal auth (optional)
AZURE_CLIENT_ID=<client-id> # For service principal auth (optional)
AZURE_CLIENT_SECRET=<client-secret> # For service principal auth (optional)
```

## Authentication

```csharp
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
var credential = new DefaultAzureCredential(
    DefaultAzureCredential.DefaultEnvironmentVariableName
);
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/dotnet/api/overview/azure/identity-readme?view=azure-dotnet#credential-classes
// var credential = new ManagedIdentityCredential();
var armClient = new ArmClient(credential);

// Get subscription
var subscriptionId = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
var subscription = armClient.GetSubscriptionResource(
    new ResourceIdentifier($"/subscriptions/{subscriptionId}"));
```

## Resource Hierarchy

```
ArmClient
└── SubscriptionResource
    └── ResourceGro
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 189438ms

**Input:**

```json
{"intent":"Generate a local-only C# console sample using Microsoft.Azure.Cosmos SDK v3 FeedIterator pagination, continuation tokens, MaxItemCount, and RU tracking. Do not deploy or modify Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 3. azure-documentation (via azure) — 189443ms

**Input:**

```json
{"intent":"Find official Microsoft documentation for Microsoft.Azure.Cosmos SDK v3 query pagination with QueryRequestOptions.MaxItemCount, GetItemQueryIterator continuation tokens, FeedIterator, FeedResponse.RequestCharge, and LINQ ToFeedIterator.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 189428ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 188737ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-resource-manager-cosmosdb-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 185816ms

**Input:**

```json
{"intent":"Get Azure SDK code generation best practices for a Microsoft.Azure.Cosmos v3 console query sample. Local-only code generation; no resource operations.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-documentation (via azure) — 185811ms

**Input:**

```json
{"intent":"Official Microsoft.Azure.Cosmos .NET v3 FeedIterator and continuation token query pagination documentation.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. web_search (via github-mcp-server) — 185806ms

**Input:**

```json
{"query":"site:learn.microsoft.com Microsoft.Azure.Cosmos .NET v3 FeedIterator QueryRequestOptions MaxItemCount continuationToken RequestCharge LINQ ToFeedIterator"}
```

**Output:**

```
No relevant web results found.
```

### ❌ 9. web_fetch — 3475ms

**Input:**

```json
{"max_length":20000,"url":"https://api.nuget.org/v3-flatcontainer/microsoft.azure.cosmos/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 10. apply_patch — 3023ms

**Input:**

```json
"*** Begin Patch\n*** Add File: CosmosPaginationSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.47.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Microsoft.Azure.Cosmos;\n+\n+const string QueryText =\n+    \"SELECT * FROM c WHERE c.category = 'electronics'\";\n+const int PageSize = 50;\n+\n+string connectionString = GetRequiredEnvironmentVariable(\n+    \"COSMOS_CONNECTION_STRING\");\n+string databaseId = GetRequiredEnvironmentVariable(\"COSMOS_DATABASE_ID\");\n+string containerId = GetRequiredEnvironmentVariable(\"COSMOS_CONTAINER_ID\");\n+string tokenFile = GetOption(args, \"--token-file\")\n+    ?? \"continuation-token.txt\";\n+\n+string? continuationToken = await ReadContinuationTokenAsync(tokenFile);\n+\n+using CosmosClient client = new(connectionString);\n+Container container = client.GetContainer(databaseId, containerId);\n+\n+QueryRequestOptions requestOptions = new()\n+{\n+    // This is a maximum, not a guarantee: Cosmos DB can return fewer items.\n+    MaxItemCount = PageSize\n+};\n+\n+FeedIterator\u003cdynamic\u003e iterator = container.GetItemQueryIterator\u003cdynamic\u003e(\n+    new QueryDefinition(QueryText),\n+    continuationToken,\n+    requestOptions);\n+\n+double totalRequestCharge = 0;\n+int pageNumber = 0;\n+int totalItemCount = 0;\n+\n+if (continuationToken is not null)\n+{\n+    Console.WriteLine($\"Resuming from token stored in '{tokenFile}'.\");\n+}\n+\n+while (iterator.HasMoreResults)\n+{\n+    FeedResponse\u003cdynamic\u003e page = await iterator.ReadNextAsync();\n+\n+    pageNumber++;\n+    totalRequestCharge += page.RequestCharge;\n+    totalItemCount += page.Count;\n+\n+    Console.WriteLine(\n+        $\"Page {pageNumber}: {page.Count} items, \" +\n+        $\"{page.RequestCharge:F2} RU\");\n+\n+    foreach (dynamic item in page)\n+    {\n+        Console.WriteLine(item);\n+    }\n+\n+    continuationToken = page.ContinuationToken;\n+    Console.WriteLine(\n+        $\"Continuation token: {continuationToken ?? \"\u003cend\u003e\"}\");\n+\n+    // Persist only after the entire page has been processed. If the program\n+    // stops before this point, restarting safely reprocesses that page.\n+    await SaveContinuationTokenAsync(tokenFile, continuationToken);\n+}\n+\n+Console.WriteLine($\"Total items processed: {totalItemCount}\");\n+Console.WriteLine($\"Total request charge: {totalRequestCharge:F2} RU\");\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    return Environment.GetEnvironmentVariable(name)\n+        ?? throw new InvalidOperationException(\n+            $\"Set the {name} environment variable before running.\");\n+}\n+\n+static string? GetOption(string[] arguments, string optionName)\n+{\n+    for (int index = 0; index \u003c arguments.Length; index++)\n+    {\n+        if (arguments[index] != optionName)\n+        {\n+            continue;\n+        }\n+\n+        if (index + 1 \u003e= arguments.Length ||\n+            string.IsNullOrWhiteSpace(arguments[index + 1]))\n+        {\n+            throw new ArgumentException(\n+                $\"Option {optionName} requires a value.\");\n+        }\n+\n+        return arguments[index + 1];\n+    }\n+\n+    return null;\n+}\n+\n+static async Task\u003cstring?\u003e ReadContinuationTokenAsync(string path)\n+{\n+    if (!File.Exists(path))\n+    {\n+        return null;\n+    }\n+\n+    string token = await File.ReadAllTextAsync(path);\n+    return string.IsNullOrWhiteSpace(token) ? null : token;\n+}\n+\n+static async Task SaveContinuationTokenAsync(\n+    string path,\n+    string? continuationToken)\n+{\n+    if (continuationToken is null)\n+    {\n+        if (File.Exists(path))\n+        {\n+            File.Delete(path);\n+        }\n+\n+        return;\n+    }\n+\n+    string? directory = Path.GetDirectoryName(Path.GetFullPath(path));\n+    if (directory is not null)\n+    {\n+        Directory.CreateDirectory(directory);\n+    }\n+\n+    await File.WriteAllTextAsync(path, continuationToken);\n+}\n*** Add File: README.md\n+# Cosmos DB query pagination with .NET\n+\n+This sample uses `Microsoft.Azure.Cosmos` SDK v3 to execute:\n+\n+```sql\n+SELECT * FROM c WHERE c.category = 'electronics'\n+```\n+\n+It reads results with `FeedIterator`, requests at most 50 items per page,\n+prints and persists each continuation token, and accumulates\n+`FeedResponse.RequestCharge` across all pages.\n+\n+## Run\n+\n+Set the connection details for a Cosmos DB account or local Cosmos DB\n+emulator. Do not store credentials in source control.\n+\n+```powershell\n+$env:COSMOS_CONNECTION_STRING = \"\u003cconnection-string\u003e\"\n+$env:COSMOS_DATABASE_ID = \"\u003cdatabase-id\u003e\"\n+$env:COSMOS_CONTAINER_ID = \"\u003ccontainer-id\u003e\"\n+dotnet run\n+```\n+\n+By default, the current continuation token is stored in\n+`continuation-token.txt`. To choose another location:\n+\n+```powershell\n+dotnet run -- --token-file .\\state\\electronics.token\n+```\n+\n+Stop the program after any completed page and run the same command again to\n+resume from the saved token. The token file is removed after the final page.\n+Because a token is saved only after processing its entire page, a failure\n+during page processing can cause that page to be processed again. Consumers\n+should therefore make side effects idempotent.\n+\n+## `MaxItemCount`\n+\n+`QueryRequestOptions.MaxItemCount = 50` asks Cosmos DB to return no more than\n+50 items in each response. It is not an exact page size: a response can\n+contain fewer items because of response-size limits, available request units,\n+query execution limits, or because the query has reached its end.\n+\n+Continuation tokens are opaque SDK/service state. Store and pass them back\n+unchanged, and avoid exposing them unnecessarily.\n+\n+## `FeedIterator` and LINQ queries\n+\n+`GetItemQueryIterator\u003cT\u003e` accepts Cosmos SQL directly and returns a\n+`FeedIterator\u003cT\u003e`. It is the clearest option when the SQL text is known and\n+the application needs explicit control over continuation tokens, request\n+options, page boundaries, and per-page diagnostics such as RU charge.\n+\n+The LINQ API starts with `GetItemLinqQueryable\u003cT\u003e()` and translates supported\n+C# expressions into Cosmos SQL. LINQ is useful for typed, composable query\n+construction, but unsupported expressions can fail translation and the\n+generated SQL may be less obvious. LINQ queries are not automatically\n+page-by-page: call `ToFeedIterator()` on the resulting query and then use the\n+same `HasMoreResults`/`ReadNextAsync` loop. For resumption, supply the saved\n+token and `QueryRequestOptions` to `GetItemLinqQueryable\u003cT\u003e` before calling\n+`ToFeedIterator()`.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj
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



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/Program.cs
@@ -1,0 +1,130 @@
+using Microsoft.Azure.Cosmos;
+
+const string QueryText =
+    "SELECT * FROM c WHERE c.category = 'electronics'";
+const int PageSize = 50;
+
+string connectionString = GetRequiredEnvironmentVariable(
+    "COSMOS_CONNECTION_STRING");
+string databaseId = GetRequiredEnvironmentVariable("COSMOS_DATABASE_ID");
+string containerId = GetRequiredEnvironmentVariable("COSMOS_CONTAINER_ID");
+string tokenFile = GetOption(args, "--token-file")
+    ?? "continuation-token.txt";
+
+string? continuationToke
... (truncated)
```

### ✅ 11. powershell — 2923ms

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Build Cosmos pagination sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.47.0\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(73,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892\CosmosPaginationSample.csproj]

Build FAILED.

C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.47.0\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(73,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892\CosmosPaginationSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:01.38
<shellId: 0 completed with exit code 1>
```

### ✅ 12. apply_patch — 783ms

**Input:**

```json
"*** Begin Patch\n*** Update File: CosmosPaginationSample.csproj\n@@\n   \u003cItemGroup\u003e\n     \u003cPackageReference Include=\"Microsoft.Azure.Cosmos\" Version=\"3.47.0\" /\u003e\n+    \u003cPackageReference Include=\"Newtonsoft.Json\" Version=\"13.0.3\" /\u003e\n   \u003c/ItemGroup\u003e\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-pagination-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2892834892/CosmosPaginationSample.csproj
@@ -9,6 +9,7 @@
 
   <ItemGroup>
     <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.47.0" />
+    <PackageReference Include="Newtonsoft.Json" Version="13.0.3" />
   </ItemGroup>
 
 </Project>

```

### ✅ 13. powershell — 3022ms

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Rebuild Cosmos pagination sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.77
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `CosmosPaginationSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll scaffold a local-only .NET sample with resumable pagination, RU accounting, and concise usage notes.

The available Cosmos skill is management-plane focused, so I’m grounding the data-plane implementation in the v3 SDK guidance before writing it.

The Azure guidance endpoints timed out; I’ll retry the official SDK documentation lookup directly, then create the standalone sample.

I’ve settled on an environment-configured console app that saves the token only after a page is fully processed, so restarts do not skip unprocessed items.

The sample is implemented; I’m compiling it now to catch SDK API or type mismatches.

The SDK’s build target requires an explicit Newtonsoft.Json reference, so I’m adding the matching serializer dependency and rebuilding.

Implemented the Cosmos DB pagination sample:

- `Program.cs` — SQL query, `FeedIterator`, `MaxItemCount = 50`, continuation-token persistence/resume, and total RU tracking.
- `CosmosPaginationSample.csproj` — .NET 8 with `Microsoft.Azure.Cosmos` v3.
- `README.md` — configuration, execution instructions, pagination semantics, and `FeedIterator` vs. LINQ explanation.

Run with:

```powershell
$env:COSMOS_CONNECTION_STRING = "<connection-string>"
$env:COSMOS_DATABASE_ID = "<database-id>"
$env:COSMOS_CONTAINER_ID = "<container-id>"
dotnet run
```

Use `dotnet run -- --token-file .\state\electronics.token` to specify the saved-token location.

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The implementation correctly covers SQL query pagination, continuation-token resumption, page-size configuration, RU accounting, and the LINQ alternative, but omits cross-partition query considerations.

### Strengths

- Uses the required Cosmos SDK v3 iterator APIs and query definition correctly.
- Persists continuation tokens only after processing a page and documents possible at-least-once reprocessing.
- Accumulates per-page request charges and clearly explains MaxItemCount semantics.
- Provides a useful comparison between direct SQL FeedIterator queries and LINQ queries converted with ToFeedIterator().

### Issues

- Cross-partition behavior is not explained; the sample should note that omitting PartitionKey may fan the query out across physical partitions and increase latency/RU consumption, while a known partition key can be supplied through QueryRequestOptions.PartitionKey.

## Grader Results

- pagination-query-items.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `Container.GetItemQueryIterator<T>()` with `QueryDefinition`: Pass
      - `QueryRequestOptions.MaxItemCount` for page size control: Pass
      - `FeedIterator<T>.HasMoreResults` and `ReadNextAsync()` loop pattern: Pass
      - `FeedResponse<T>.ContinuationToken` for resumable pagination: Pass
      - Passing continuation token to resume a query: Pass
      - `FeedResponse<T>.RequestCharge` for RU tracking: Pass
      - Cross-partition query considerations: Fail
      - LINQ alternative via `GetItemLinqQueryable<T>()`: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id cosmos-db-dp-dotnet-pagination --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
