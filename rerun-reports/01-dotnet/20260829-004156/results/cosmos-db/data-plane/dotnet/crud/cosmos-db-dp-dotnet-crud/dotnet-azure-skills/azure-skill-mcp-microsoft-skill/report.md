# Evaluation Report: cosmos-db-dp-dotnet-crud

**Config:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ✅ PASSED | **Duration:** 166.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `cosmos-db-dp-dotnet-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ✅ PASSED |
| Score | 7/7 |
| Duration | 166.9s |
| Timestamp | 2026-08-28T17:01:39Z |
| Files Generated | 3 |
| Event Count | 4174 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 137.3s |
| Review | 27.9s |
| **Total** | **166.9s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-resource-manager-cosmosdb-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=579624 out=3219 |
| Turn Count | 14 |

## Tool Calls

### ✅ 1. skill — 4401ms

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

### ✅ 2. glob — 7701ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 5492ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-resource-manager-cosmosdb-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. azure-get_azure_bestpractices (via azure) — 6507ms

**Input:**

```json
{"command":"get_best_practices","intent":"Generate a local-only C# console program using Microsoft.Azure.Cosmos for NoSQL database, container, and item CRUD with CosmosException handling and an environment-variable connection string.","parameters":{"language":"C#","operation":"code-generation","resource":"Azure Cosmos DB"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 3245ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a local-only C# console sample for Cosmos DB NoSQL CRUD using Microsoft.Azure.Cosmos.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. powershell — 4993ms

**Input:**

```json
{"command":"dotnet new console --framework net8.0 --use-program-main --name CosmosCrudSample --output . \u0026\u0026 dotnet add package Microsoft.Azure.Cosmos","description":"Scaffold console app and add Cosmos SDK","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Error: Invalid option(s):
--framework net8.0
   'net8.0' is not a valid value for --framework. The possible values are:
      net10.0   - Target net10.0
      net9.0    - Target net9.0

For more information, run:
   dotnet new console -h

For details on the exit code, refer to https://aka.ms/templating-exit-codes#127
<shellId: 0 completed with exit code 127>
```

### ✅ 7. powershell — 7047ms

**Input:**

```json
{"command":"dotnet new console --framework net9.0 --use-program-main --name CosmosCrudSample --output . \u0026\u0026 dotnet add package Microsoft.Azure.Cosmos","description":"Scaffold .NET 9 app and add Cosmos SDK","initial_wait":120,"mode":"sync"}
```

**Output:**

```
The template "Console App" was created successfully.

Processing post-creation actions...
Restoring C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj:
  Determining projects to restore...
  Restored C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj (in 55 ms).
Restore succeeded.


info : X.509 certificate chain validation will use the default trust store selected by .NET for code signing.
info : X.509 certificate chain validation will use the default trust store selected by .NET for timestamping.
info : Adding PackageReference for package 'Microsoft.Azure.Cosmos' into project 'C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj'.
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/registrations2-semver2/microsoft.azure.cosmos/index.json
info : Restoring packages for C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj...
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/index.json
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/vulnerabilitybase.json
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/vulnerabilityupdate.json
info : Package 'Microsoft.Azure.Cosmos' is compatible with all the specified frameworks in project 'C:\Users\
... (truncated)
```

### ✅ 8. view — 6091ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\\Program.cs"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
@@ -1,10 +1,10 @@
 ﻿namespace CosmosCrudSample;
 
 class Program
 {
     static void Main(string[] args)
     {
         Console.WriteLine("Hello, World!");
     }
 }
 

```

### ✅ 9. view — 4529ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\\CosmosCrudSample.csproj"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/CosmosCrudSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/CosmosCrudSample.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/CosmosCrudSample.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/CosmosCrudSample.csproj
@@ -1,15 +1,15 @@
 ﻿<Project Sdk="Microsoft.NET.Sdk">
 
   <PropertyGroup>
     <OutputType>Exe</OutputType>
     <TargetFramework>net9.0</TargetFramework>
     <ImplicitUsings>enable</ImplicitUsings>
     <Nullable>enable</Nullable>
   </PropertyGroup>
 
   <ItemGroup>
     <PackageReference Include="Microsoft.Azure.Cosmos" Version="3.62.1" />
   </ItemGroup>
 
 </Project>
 

```

### ✅ 10. apply_patch — 641ms

**Input:**

```json
"*** Begin Patch\n*** Delete File: Program.cs\n*** Add File: Program.cs\n+using System.Net;\n+using Microsoft.Azure.Cosmos;\n+\n+namespace CosmosCrudSample;\n+\n+internal static class Program\n+{\n+    private const string DatabaseName = \"TestDB\";\n+    private const string ContainerName = \"Items\";\n+    private const string PartitionKeyPath = \"/category\";\n+\n+    public static async Task\u003cint\u003e Main()\n+    {\n+        string? connectionString =\n+            Environment.GetEnvironmentVariable(\"COSMOS_CONNECTION_STRING\");\n+\n+        if (string.IsNullOrWhiteSpace(connectionString))\n+        {\n+            Console.Error.WriteLine(\n+                \"Set the COSMOS_CONNECTION_STRING environment variable before running.\");\n+            return 1;\n+        }\n+\n+        try\n+        {\n+            using CosmosClient client = new(connectionString);\n+\n+            DatabaseResponse databaseResponse =\n+                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);\n+            Database database = databaseResponse.Database;\n+\n+            ContainerResponse containerResponse =\n+                await database.CreateContainerIfNotExistsAsync(\n+                    new ContainerProperties(ContainerName, PartitionKeyPath),\n+                    throughput: 400);\n+            Container container = containerResponse.Container;\n+\n+            Item item = new(\n+                id: Guid.NewGuid().ToString(),\n+                category: \"electronics\",\n+                name: \"Wireless keyboard\",\n+                quantity: 10);\n+            PartitionKey partitionKey = new(item.category);\n+\n+            ItemResponse\u003cItem\u003e createResponse =\n+                await container.CreateItemAsync(item, partitionKey);\n+            Console.WriteLine(\n+                $\"Created item {createResponse.Resource.id} \" +\n+                $\"(request charge: {createResponse.RequestCharge:F2} RUs).\");\n+\n+            ItemResponse\u003cItem\u003e readResponse =\n+                await container.ReadItemAsync\u003cItem\u003e(item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Read item: {readResponse.Resource.name}, \" +\n+                $\"quantity {readResponse.Resource.quantity}.\");\n+\n+            QueryDefinition query = new(\n+                \"SELECT * FROM c WHERE c.category = @category\");\n+            query.WithParameter(\"@category\", \"electronics\");\n+\n+            using FeedIterator\u003cItem\u003e iterator = container.GetItemQueryIterator\u003cItem\u003e(\n+                query,\n+                requestOptions: new QueryRequestOptions\n+                {\n+                    PartitionKey = new PartitionKey(\"electronics\")\n+                });\n+\n+            Console.WriteLine(\"Electronics items:\");\n+            while (iterator.HasMoreResults)\n+            {\n+                FeedResponse\u003cItem\u003e page = await iterator.ReadNextAsync();\n+                foreach (Item result in page)\n+                {\n+                    Console.WriteLine(\n+                        $\"- {result.id}: {result.name}, quantity {result.quantity}\");\n+                }\n+            }\n+\n+            Item updatedItem = item with { quantity = 25 };\n+            ItemResponse\u003cItem\u003e replaceResponse =\n+                await container.ReplaceItemAsync(updatedItem, item.id, partitionKey);\n+            Console.WriteLine(\n+                $\"Updated quantity to {replaceResponse.Resource.quantity}.\");\n+\n+            await container.DeleteItemAsync\u003cItem\u003e(item.id, partitionKey);\n+            Console.WriteLine($\"Deleted item {item.id}.\");\n+\n+            return 0;\n+        }\n+        catch (CosmosException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Cosmos DB request failed: {(int)exception.StatusCode} \" +\n+                $\"({exception.StatusCode}), substatus {exception.SubStatusCode}, \" +\n+                $\"activity ID {exception.ActivityId}.\");\n+            Console.Error.WriteLine(exception.Message);\n+\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            {\n+                Console.Error.WriteLine(\n+                    $\"Retry after {exception.RetryAfter.TotalMilliseconds:F0} ms.\");\n+            }\n+\n+            return 1;\n+        }\n+        catch (ArgumentException exception)\n+        {\n+            Console.Error.WriteLine($\"Invalid Cosmos DB configuration: {exception.Message}\");\n+            return 1;\n+        }\n+    }\n+\n+    private sealed record Item(\n+        string id,\n+        string category,\n+        string name,\n+        int quantity);\n+}\n*** Add File: README.md\n+# Cosmos DB CRUD sample\n+\n+This .NET 9 console application uses the Azure Cosmos DB for NoSQL data-plane\n+SDK to create a database and container, then create, read, query, replace, and\n+delete an item.\n+\n+## Required NuGet package\n+\n+```powershell\n+dotnet add package Microsoft.Azure.Cosmos --version 3.62.1\n+```\n+\n+No Azure resources are provisioned by this project. Run it against an existing\n+account or a local Azure Cosmos DB emulator.\n+\n+## Run\n+\n+Set the connection string without placing credentials in source control:\n+\n+```powershell\n+$env:COSMOS_CONNECTION_STRING = \"\u003cyour-connection-string\u003e\"\n+dotnet run\n+```\n+\n+The account or emulator must already exist. The program idempotently creates\n+the `TestDB` database and the `Items` container with partition key `/category`\n+and 400 RU/s if they do not exist.\n+\n+SDK reference:\n+\u003chttps://learn.microsoft.com/azure/cosmos-db/nosql/quickstart-dotnet\u003e\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs b/dev/null
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
+++ b/dev/null
@@ -1,10 +0,0 @@
-﻿namespace CosmosCrudSample;
-
-class Program
-{
-    static void Main(string[] args)
-    {
-        Console.WriteLine("Hello, World!");
-    }
-}
-



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
@@ -1,0 +1,119 @@
+using System.Net;
+using Microsoft.Azure.Cosmos;
+
+namespace CosmosCrudSample;
+
+internal static class Program
+{
+    private const string DatabaseName = "TestDB";
+    private const string ContainerName = "Items";
+    private const string PartitionKeyPath = "/category";
+
+    public static async Task<int> Main()
+    {
+        string? connectionString =
+            Environment.GetEnvironmentVariable("COSMOS_CONNECTION_STRING");
+
+        if (string.IsNullOrWhiteSpace(connectionString))
+        {
+            Console.Error.WriteLine(
+                "Set the COSMOS_CONNECTION_STRING environment variable before running.");
+            return 1;
+        }
+
+        try
+        {
+            using CosmosClient client = new(connectionString);
+
+            DatabaseResponse databaseResponse =
+                await client.CreateDatabaseIfNotExistsAsync(DatabaseName);
+            Database database = databaseResponse.Database
... (truncated)
```

### ✅ 11. powershell — 1561ms

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Build Cosmos CRUD sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.62.1\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(72,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\.nuget\packages\microsoft.azure.cosmos\3.62.1\buildTransitive\netstandard2.0\Microsoft.Azure.Cosmos.targets(72,9): error : The Newtonsoft.Json package must be explicitly referenced with version >= 10.0.2. Please add a reference to Newtonsoft.Json or set the 'AzureCosmosDisableNewtonsoftJsonCheck' property to 'true' to bypass this check. [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:00.40
<shellId: 2 completed with exit code 1>
```

### ✅ 12. powershell — 2324ms

**Input:**

```json
{"command":"dotnet add package Newtonsoft.Json","description":"Add required JSON package","initial_wait":120,"mode":"sync"}
```

**Output:**

```
info : X.509 certificate chain validation will use the default trust store selected by .NET for code signing.
info : X.509 certificate chain validation will use the default trust store selected by .NET for timestamping.
info : Adding PackageReference for package 'Newtonsoft.Json' into project 'C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj'.
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/registrations2-semver2/newtonsoft.json/index.json
info : Restoring packages for C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj...
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/index.json
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/vulnerabilitybase.json
info :   CACHE https://ms-feed-2.pkgs.visualstudio.com/f5581750-f66a-4ee8-b9cc-8269930bd7c4/_packaging/8dde208a-0fd1-492b-9441-52d063bc5d31/nuget/v3/vulnerabilities/vulnerabilityupdate.json
info : Package 'Newtonsoft.Json' is compatible with all the specified frameworks in project 'C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj'.
info : PackageReference for package 'Newtonsoft.Json' version '13.0.4' added to file 'C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj'.
info : Writing assets file to disk. Path: C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-s
... (truncated)
```

### ✅ 13. apply_patch — 613ms

**Input:**

```json
"*** Begin Patch\n*** Update File: README.md\n@@\n ```powershell\n dotnet add package Microsoft.Azure.Cosmos --version 3.62.1\n+dotnet add package Newtonsoft.Json --version 13.0.4\n ```\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/README.md b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/README.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/README.md
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/README.md
@@ -8,6 +8,7 @@
 
 ```powershell
 dotnet add package Microsoft.Azure.Cosmos --version 3.62.1
+dotnet add package Newtonsoft.Json --version 13.0.4
 ```
 
 No Azure resources are provisioned by this project. Run it against an existing

```

### ✅ 14. powershell — 2154ms

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Build completed Cosmos sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\Program.cs(101,57): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\Program.cs(101,57): error CS1061: 'TimeSpan?' does not contain a definition for 'TotalMilliseconds' and no accessible extension method 'TotalMilliseconds' accepting a first argument of type 'TimeSpan?' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\CosmosCrudSample.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:00.96
<shellId: 4 completed with exit code 1>
```

### ✅ 15. apply_patch — 647ms

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n-            if (exception.StatusCode == HttpStatusCode.TooManyRequests)\n+            if (exception.StatusCode == HttpStatusCode.TooManyRequests \u0026\u0026\n+                exception.RetryAfter is TimeSpan retryAfter)\n             {\n                 Console.Error.WriteLine(\n-                    $\"Retry after {exception.RetryAfter.TotalMilliseconds:F0} ms.\");\n+                    $\"Retry after {retryAfter.TotalMilliseconds:F0} ms.\");\n             }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540/Program.cs
@@ -95,10 +95,11 @@
                 $"activity ID {exception.ActivityId}.");
             Console.Error.WriteLine(exception.Message);
 
-            if (exception.StatusCode == HttpStatusCode.TooManyRequests)
+            if (exception.StatusCode == HttpStatusCode.TooManyRequests &&
+                exception.RetryAfter is TimeSpan retryAfter)
             {
                 Console.Error.WriteLine(
-                    $"Retry after {exception.RetryAfter.TotalMilliseconds:F0} ms.");
+                    $"Retry after {retryAfter.TotalMilliseconds:F0} ms.");
             }
 
             return 1;

```

### ✅ 16. powershell — 1889ms

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Rebuild Cosmos sample","initial_wait":120,"mode":"sync"}
```

**Output:**

```
CosmosCrudSample -> C:\Users\weidxu\AppData\Local\Temp\hyoka-cosmos-db-dp-dotnet-crud-dotnet-azure-skills-azure-skill-mcp-microsoft-skill-2870800540\bin\Debug\net9.0\CosmosCrudSample.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:00.68
<shellId: 5 completed with exit code 0>
```

## Generated Files

- `CosmosCrudSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll scaffold a local .NET console sample using the Cosmos DB data-plane SDK, with configuration kept in an environment variable and no live Azure operations.

The project is scaffolded with `Microsoft.Azure.Cosmos` 3.62.1. I’m now implementing the full create/read/query/replace/delete flow and concise emulator-ready run instructions.

The current Cosmos SDK requires an explicit `Newtonsoft.Json` reference at build time. I’m adding that required package and updating the package list accordingly.

Implemented a complete .NET 9 Cosmos DB CRUD console application.

- `Program.cs`: Creates `TestDB`/`Items`, inserts, reads, parameterized-queries, replaces, and deletes an item with `CosmosException` handling.
- `CosmosCrudSample.csproj`: Uses `Microsoft.Azure.Cosmos` 3.62.1 and `Newtonsoft.Json` 13.0.4.
- `README.md`: Includes package installation, environment-variable configuration, emulator guidance, and the [Microsoft Cosmos DB .NET reference](https://learn.microsoft.com/azure/cosmos-db/nosql/quickstart-dotnet).

Run with:

```powershell
$env:COSMOS_CONNECTION_STRING = "<your-connection-string>"
dotnet run
```

## Reviewer Notes (LLM-as-Judge)

### Summary

7/7 criteria passed. The generated application fully implements the requested Cosmos DB CRUD workflow, parameterized query, package setup, and CosmosException handling.

### Strengths

- Uses the current Microsoft.Azure.Cosmos v3 SDK.
- Keeps the connection string out of source code by using an environment variable.
- Uses the category value consistently as the partition key for CRUD operations.
- Uses a parameterized SQL-like query and processes paginated results.
- Reports useful Cosmos DB failure details, including status, substatus, activity ID, and throttling delay.

### Issues

- Newtonsoft.Json is included but not used and is unnecessary for this sample.

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
hyoka run --prompt-id cosmos-db-dp-dotnet-crud --config dotnet-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
