# Evaluation Report: resource-manager-mp-dotnet-rg-crud

**Config:** dotnet-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 111.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-dotnet-rg-crud` |
| Config | dotnet-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 6/6 |
| Duration | 111.0s |
| Timestamp | 2026-08-28T18:27:22Z |
| Files Generated | 3 |
| Event Count | 4462 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 96.7s |
| Review | 13.2s |
| **Total** | **111.0s** |

## Configuration

- **name:** dotnet-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-resource-lookup |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=473130 out=4293 |
| Turn Count | 10 |

## Tool Calls

### ✅ 1. skill — 633ms

**Input:**

```json
{"skill":"azure-resource-lookup"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Resource Lookup

List, find, and discover Azure resources of any type across subscriptions and resource groups. Use Azure Resource Graph (ARG) for fast, cross-cutting queries when dedicated MCP tools don't cover the resource type.

## When to Use This Skill

Use this skill when the user wants to:
- **List resources** of any type (VMs, web apps, storage accounts, container apps, databases, etc.)
- **Show resources** in a specific subscription or resource group
- Query resources **across multiple subscriptions** or resource types
- Find **orphaned resources** (unattached disks, unused NICs, idle IPs)
- Discover resources **missing required tags** or configurations
- Get a **resource inventory** spanning multiple types
- Find resources in a **specific state** (unhealthy, failed provisioning, stopped)
- Answer "**what resources do I have?**" or "**show me my Azure resources**"
- **List web apps, websites, or App Services**

> ⚠️ **Warning:** App Service / Web Apps have no dedicated MCP `list` command. Prompts like "list websites", "list web apps", or "list app services" **must** route through this skill to use Azure Resource Graph.

> 💡 **Tip:** For single-resource-type queries, first check if a dedicated MCP tool can handle it (see routing table below). If none exists, use Azure Resource Graph.

## Quick Reference

| Property | Value |
|----------|-------|
| **Query Language** | KQL (Kusto Query Language subset) |
| **CLI Command** | `az graph query -q "<KQL>" -o table` |
| **Extension** | `az extension add --name resource-graph` |
| **MCP Tool** | `extension_cli_generate` with intent for `az graph query` |
| **Best For** | Cross-subscription queries, orphaned resources, tag audits |

## MCP Tools

| Tool | Purpose | When to Use |
|------|---------|-------------|
| `extension_cli_generate` | Generate `az graph query` commands | Primary tool — generate ARG queries from user intent |
| `mcp_azure_mcp_subscription_list` | Li
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 4636ms

**Input:**

```json
{"command":"get","intent":"Generate a local C# console application using Azure.ResourceManager and DefaultAzureCredential to create, list, retrieve, tag, and delete an Azure resource group with proper error handling. Do not deploy or contact Azure.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 5281ms

**Input:**

```json
{"command":"search","intent":"Find official current .NET Azure.ResourceManager examples and API signatures for ArmClient, subscriptions, ResourceGroupCollection CreateOrUpdateAsync, ResourceGroupResource GetAsync, SetTagsAsync or UpdateAsync, and DeleteAsync, plus DefaultAzureCredential.","parameters":{"query":"Azure.ResourceManager .NET create list get update tags delete resource groups ArmClient DefaultAzureCredential"}}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 4314ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. glob — 3997ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-resource-lookup","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. view — 4961ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-resource-lookup\\references\\azure-resource-graph.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-resource-lookup/references/azure-resource-graph.md
@@ -1,185 +1,185 @@
 # Azure Resource Graph Query Patterns
 
 Azure Resource Graph (ARG) queries use a KQL subset against indexed Azure resource metadata. Results are near real-time across all subscriptions.
 
 ## Command Format
 
 ```bash
 az graph query -q "<KQL>" --query "data[].{col1:field1, col2:field2}" -o table
 ```
 
 | Flag | Purpose |
 |------|---------|
 | `-q` | KQL query string |
 | `--query` | JMESPath to shape output columns |
 | `--first N` | Limit to N results |
 | `--subscriptions` | Scope to specific subscription IDs |
 | `-o table` | Table output (also: json, tsv) |
 
 ## Key Tables
 
 | Table | Contents |
 |-------|----------|
 | `Resources` | All ARM resources — name, type, location, properties, tags, sku |
 | `ResourceContainers` | Subscriptions, resource groups, management groups |
 | `HealthResources` | Resource health availability status |
 | `ServiceHealthResources` | Azure service health events/incidents |
 | `AuthorizationResources` | Role assignments and definitions |
 | `AdvisorResources` | Azure Advisor recommendations |
 
 ## KQL Essentials
 
 - `=~` case-insensitive equals (use for `type` field — types are lowercase)
 - `properties.fieldName` navigates the properties JSON bag
 - `mv-expand` flattens arrays (subnets, IP configs)
 - `isempty()` / `is
... (truncated)
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 2383ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure SDK code-generation best practices for a local C# Azure.ResourceManager resource-group CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 3345ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find current official Azure.ResourceManager .NET resource group CRUD API documentation.","parameters":{"query":"Azure.ResourceManager .NET ArmClient ResourceGroupCollection CreateOrUpdateAsync GetAsync SetTagsAsync DeleteAsync DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"title":"Microsoft Azure Resource Manager client library for .NET","content":"# Microsoft Azure Resource Manager client library for .NET\n## Examples\n### Create a resource group\n### List all resource groups\n### Update a resource group\n\u0060\u0060\u0060C\n// Note: Resource group named \u0027myRgName\u0027 should exist for this example to work.\nArmClient client = new ArmClient(new DefaultAzureCredential());\nSubscriptionResource subscription = await client.GetDefaultSubscriptionAsync();\nResourceGroupCollection resourceGroups = subscription.GetResourceGroups();\nstring resourceGroupName = \u0022myRgName\u0022;\nResourceGroupResource resourceGroup = await resourceGroups.GetAsync(resourceGroupName);\nresourceGroup = await resourceGroup.AddTagAsync(\u0022key\u0022, \u0022value\u0022);\n\u0060\u0060\u0060\n### Delete a resource group\n\u0060\u0060\u0060C\nArmClient client = new ArmClient(new DefaultAzureCredential());\nSubscriptionResource subscription = await client.GetDefaultSubscriptionAsync();\nResourceGroupCollection resourceGroups = subscription.GetResourceGroups();\nstring resourceGroupName = \u0022myRgName\u0022;\nResourceGroupResource resourceGroup = await resourceGroups.GetAsync(resourceGroupName);\nawait resourceGroup.DeleteAsync(WaitUntil.Completed);\n\u0060\u0060\u0060\n### Get GenericResource List\n### Create GenericResource\n### Update GenericResource\n### Update GenericResource Tags\n### Get GenericResource\n### Check whether GenericResource exists\n\u0060\u0060\u0060C\nArmClient client = new ArmClient(new DefaultAzureCredential());\nResourceIdentifier id = new ResourceIdentifier(\u0022/subscriptions/{subscription_id}/resourceGroups/{resourcegroup_name}/providers/Microsoft.Network/virtualNetworks/{vnet_name}\u0022);\n\nbool existResult = await client.GetGenericResources().ExistsAsync(id);\nConsole.WriteLine($\u0022Resource exists: {existResult}\u0022);\n\u0060\u0060\u0060\n### Delete GenericResource\n\u0060\u0060\u0060C\nArmClient client
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 3032ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official C# Azure.ResourceManager resource group CRUD code samples.","parameters":{"language":"csharp","query":"Azure.ResourceManager C# ArmClient create list get update tags delete resource group DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"description":"description: Creates and authenticates an ArmClient instance using DefaultAzureCredential, retrieves the default subscription, and creates or updates a resource group with a specified name and location.\npackage: Azure.Identity, Azure.ResourceManager\nlanguage: c\n","codeSnippet":"ArmClient armClient = new ArmClient(new DefaultAzureCredential());\r\nSubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\r\nResourceGroupCollection rgCollection = subscription.GetResourceGroups();\r\n// With the collection, we can create a new resource group with an specific name\r\nstring rgName = \u0022sample-rg\u0022;\r\nAzureLocation location = AzureLocation.WestUS2;\r\nArmOperation\u003CResourceGroupResource\u003E resourceGroupLro = await rgCollection.CreateOrUpdateAsync(WaitUntil.Completed, rgName, new ResourceGroupData(location));\r\nResourceGroupResource resourceGroup = lro.Value;","link":"https://learn.microsoft.com/dotnet/api/overview/azure/resourcemanager.loadtesting-readme?view=azure-dotnet#examples","language":"c"},{"description":"description: Deletes a specified availability set from a resource group using the Azure Resource Manager client.\npackage: Azure.ResourceManager\nlanguage: c\n","codeSnippet":"// First, initialize the ArmClient and get the default subscription\r\nArmClient armClient = new ArmClient(new DefaultAzureCredential());\r\n// Now we get a ResourceGroupResource collection for that subscription\r\nSubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\r\nResourceGroupCollection rgCollection = subscription.GetResourceGroups();\r\n\r\n// With the collection, we can create a new resource group with an specific name\r\nstring rgName = \u0022myRgName\u0022;\r\nResourceGroupResource resourceGroup = await rgCollection.GetAsync(rgName);\r\nAvailabilitySetCollection availabilitySetCollection = resourceGroup.GetAvailabilitySets();\r\nstring availabilitySetName = \u0022myAvailabilityS
... (truncated)
```

### ✅ 10. powershell — 2718ms

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK version","initial_wait":30,"mode":"sync"}
```

**Output:**

```
10.0.111
<shellId: 0 completed with exit code 0>
```

### ✅ 11. rg — 6178ms

**Input:**

```json
{"-C":3,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787941680437-copilot-tool-output-b66f86a47c02415684d85bb8dcc93460.txt","pattern":"Update a resource group|ResourceGroupPatch|SetTags|Delete a resource group|List all resource groups"}
```

**Output:**

```
Output too large to read at once (48.6 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787941688398-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787941680437-copilot-tool-output-b66f86a47c02415684d85bb8dcc93460.txt:{"results":[{"title":"Microsoft Azure Resource Manager client library for .NET","content":"# Microsoft Azure Resource Manager client library for .NET\n## Examples\n### Create a resource group\n### List all resource groups\n### Update a resource group\n\u0060\u0060\u0060C\n// Note: Resource group named \u0027myRgName\u0027 should exist for this example to work.\nArmClient client = new ArmClie
```

### ✅ 12. powershell — 6491ms

**Input:**

```json
{"command":"dotnet new console --framework net10.0 --use-program-main --name AzureResourceGroupManager --output . --no-restore; dotnet add .\\AzureResourceGroupManager.csproj package Azure.Identity; dotnet add .\\AzureResourceGroupManager.csproj package Azure.ResourceManager","description":"Create project and add Azure SDK packages","initial_wait":180,"mode":"sync"}
```

**Output:**

```
The template "Console App" was created successfully.

info : X.509 certificate chain validation will use the default trust store selected by .NET for code signing.
info : X.509 certificate chain validation will use the default trust store selected by .NET for timestamping.
info : Adding PackageReference for package 'Azure.Identity' into project '.\AzureResourceGroupManager.csproj'.
info :   CACHE https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/registrations2-semver2/azure.identity/index.json
info : Restoring packages for C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\AzureResourceGroupManager.csproj...
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/index.json
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/index.json 553ms
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/vulnerabilitybase.json
info :   GET https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/vulnerabilityupdate.json
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/vulnerabilitybase.json 302ms
info :   OK https://ms-feed-25.pkgs.visualstudio.com/6f084628-a36d-42cb-934d-057357e379dc/_packaging/52ee4efa-5537-4eae-acb1-a914a4cca686/nuget/v3/vulnerabilities/vulnerabilityupdate.json 535ms
info : Package 'Azure.Identity' is compatible with all the specified frameworks in project '.\AzureResourceGrou
... (truncated)
```

### ✅ 13. view — 2073ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\AzureResourceGroupManager.csproj"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/AzureResourceGroupManager.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/AzureResourceGroupManager.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/AzureResourceGroupManager.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/AzureResourceGroupManager.csproj
@@ -1,16 +1,16 @@
 ﻿<Project Sdk="Microsoft.NET.Sdk">
 
   <PropertyGroup>
     <OutputType>Exe</OutputType>
     <TargetFramework>net10.0</TargetFramework>
     <ImplicitUsings>enable</ImplicitUsings>
     <Nullable>enable</Nullable>
   </PropertyGroup>
 
   <ItemGroup>
     <PackageReference Include="Azure.Identity" Version="1.21.0" />
     <PackageReference Include="Azure.ResourceManager" Version="1.14.0" />
   </ItemGroup>
 
 </Project>
 

```

### ✅ 14. view — 1765ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\Program.cs"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
@@ -1,10 +1,10 @@
 ﻿namespace AzureResourceGroupManager;
 
 class Program
 {
     static void Main(string[] args)
     {
         Console.WriteLine("Hello, World!");
     }
 }
 

```

### ✅ 15. azure-documentation (via azure) — 2386ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Fetch the official Azure SDK resource-management examples page for complete current CRUD examples.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/resource-management"}}
```

**Output:**

```
# Resource management using the Azure SDK for .NET

The Azure SDK for .NET management plane libraries will help you create, provision, and manage Azure resources from within .NET applications. All Azure services have corresponding management libraries.

With the management libraries (namespaces beginning with `Azure.ResourceManager`, for example, `Azure.ResourceManager.Compute`), you can write configuration and deployment programs to perform the same tasks that you can through the Azure portal, Azure CLI, or other resource management tools.

Those packages follow the [new Azure SDK guidelines](https://azure.github.io/azure-sdk/general_introduction.html), which provide [core capabilities](https://azure.github.io/azure-sdk/general_azurecore.html) that are shared amongst all Azure SDKs, including:

- The intuitive Azure Identity library.
- An HTTP pipeline with custom policies.
- Error handling.
- Distributed tracing.

Note

You might notice that some packages are still prerelease version. Phased releases of additional Azure services' management plane libraries are in process. If you're looking for a stable version package for a particular Azure resource and currently only a prerelease version is available, please raise an issue in [Azure SDK for .NET GitHub repo](https://github.com/Azure/azure-sdk-for-net/issues/new?assignees=&amp;labels=&amp;template=02_feature_request.yml&amp;title=%5BFEATURE+REQ%5D).

## Get started

### Prerequisites

- An [Azure subscription](https://azure.microsoft.com/pricing/purchase-options/azure-account?cid=msft_learn).
- A [TokenCredential](/en-us/dotnet/api/azure.core.tokencredential?view=azure-dotnet&amp;preserve-view=false) implementation, such as an [Azure Identity library credential type](/en-us/dotnet/api/overview/azure/Identity-readme#credential-classes).

### Install the package

Install the Azure Identity and Azure resource management NuGet packages for .NET. For example:

# [PowerShell](#tab/PowerShell)
... (truncated)
```

### ✅ 16. apply_patch — 656ms

**Input:**

```json
"*** Begin Patch\n*** Delete File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\Program.cs\n*** Add File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\Program.cs\n+using Azure;\n+using Azure.Core;\n+using Azure.Identity;\n+using Azure.ResourceManager;\n+using Azure.ResourceManager.Resources;\n+\n+namespace AzureResourceGroupManager;\n+\n+internal static class Program\n+{\n+    private static async Task\u003cint\u003e Main()\n+    {\n+        ResourceGroupResource? createdResourceGroup = null;\n+\n+        try\n+        {\n+            DefaultAzureCredential credential = new();\n+            ArmClient armClient = new(credential);\n+            SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\n+            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();\n+\n+            string resourceGroupName =\n+                $\"sdk-rg-{DateTimeOffset.UtcNow:yyyyMMddHHmmss}-{Guid.NewGuid():N}\"[..39];\n+\n+            Console.WriteLine(\n+                $\"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...\");\n+\n+            ResourceGroupData resourceGroupData = new(AzureLocation.EastUS);\n+            ArmOperation\u003cResourceGroupResource\u003e createOperation =\n+                await resourceGroups.CreateOrUpdateAsync(\n+                    WaitUntil.Completed,\n+                    resourceGroupName,\n+                    resourceGroupData);\n+\n+            createdResourceGroup = createOperation.Value;\n+            Console.WriteLine($\"Created: {createdResourceGroup.Id}\");\n+\n+            Console.WriteLine(\"\\nResource groups in the subscription:\");\n+            await foreach (ResourceGroupResource resourceGroup in resourceGroups.GetAllAsync())\n+            {\n+                Console.WriteLine(\n+                    $\"- {resourceGroup.Data.Name} ({resourceGroup.Data.Location})\");\n+            }\n+\n+            ResourceGroupResource resourceGroupDetails =\n+                await resourceGroups.GetAsync(resourceGroupName);\n+\n+            Console.WriteLine(\"\\nCreated resource group details:\");\n+            Console.WriteLine($\"Name: {resourceGroupDetails.Data.Name}\");\n+            Console.WriteLine($\"Location: {resourceGroupDetails.Data.Location}\");\n+            Console.WriteLine($\"Resource ID: {resourceGroupDetails.Id}\");\n+            Console.WriteLine(\n+                $\"Provisioning state: {resourceGroupDetails.Data.Properties.ProvisioningState}\");\n+\n+            const string tagName = \"ManagedBy\";\n+            const string tagValue = \"Azure.ResourceManager\";\n+            createdResourceGroup =\n+                await resourceGroupDetails.AddTagAsync(tagName, tagValue);\n+\n+            Console.WriteLine($\"\\nAdded tag: {tagName}={createdResourceGroup.Data.Tags[tagName]}\");\n+\n+            Console.WriteLine($\"\\nDeleting resource group '{resourceGroupName}'...\");\n+            await createdResourceGroup.DeleteAsync(WaitUntil.Completed);\n+            createdResourceGroup = null;\n+            Console.WriteLine(\"Resource group deleted.\");\n+\n+            return 0;\n+        }\n+        catch (AuthenticationFailedException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Authentication failed. Configure a credential supported by \" +\n+                $\"DefaultAzureCredential. Details: {exception.Message}\");\n+            return 1;\n+        }\n+        catch (RequestFailedException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Azure request failed (HTTP {exception.Status}, \" +\n+                $\"error code '{exception.ErrorCode ?? \"unknown\"}'): {exception.Message}\");\n+            return 1;\n+        }\n+        catch (OperationCanceledException)\n+        {\n+            Console.Error.WriteLine(\"The operation was canceled.\");\n+            return 1;\n+        }\n+        finally\n+        {\n+            if (createdResourceGroup is not null)\n+            {\n+                try\n+                {\n+                    Console.Error.WriteLine(\n+                        $\"\\nCleaning up resource group '{createdResourceGroup.Data.Name}'...\");\n+                    await createdResourceGroup.DeleteAsync(WaitUntil.Completed);\n+                    Console.Error.WriteLine(\"Cleanup completed.\");\n+                }\n+                catch (RequestFailedException exception)\n+                {\n+                    Console.Error.WriteLine(\n+                        $\"Cleanup failed (HTTP {exception.Status}, \" +\n+                        $\"error code '{exception.ErrorCode ?? \"unknown\"}'): {exception.Message}\");\n+                }\n+            }\n+        }\n+    }\n+}\n*** Add File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\README.md\n+# Azure Resource Group Manager\n+\n+This .NET console application uses the current Azure management-plane SDK to:\n+\n+1. Authenticate with `DefaultAzureCredential`.\n+2. Create a uniquely named resource group in `eastus`.\n+3. List every resource group in the default subscription.\n+4. Retrieve and display the new resource group's details.\n+5. Add a `ManagedBy=Azure.ResourceManager` tag.\n+6. Delete the resource group.\n+\n+If a later operation fails after creation, the `finally` block attempts to delete the\n+resource group so the sample does not leave it behind.\n+\n+## Required NuGet packages\n+\n+```powershell\n+dotnet add package Azure.Identity\n+dotnet add package Azure.ResourceManager\n+```\n+\n+The project currently pins:\n+\n+```xml\n+\u003cPackageReference Include=\"Azure.Identity\" Version=\"1.21.0\" /\u003e\n+\u003cPackageReference Include=\"Azure.ResourceManager\" Version=\"1.14.0\" /\u003e\n+```\n+\n+Resource-group types are included in `Azure.ResourceManager`; the older\n+`Microsoft.Azure.Management.*` packages are not used.\n+\n+## Run\n+\n+Authenticate using any credential supported by `DefaultAzureCredential`, such as\n+Azure CLI for local development or managed identity when hosted in Azure. The identity\n+needs permission to create, read, tag, and delete resource groups in the subscription;\n+the built-in `Resource Group Contributor` role is sufficient when scoped appropriately.\n+\n+```powershell\n+dotnet run\n+```\n+\n+The program uses the default subscription visible to the selected credential and creates\n+then deletes a real resource group. Do not interrupt it after creation; if cleanup fails,\n+delete the printed resource group manually.\n+\n+## References\n+\n+- [Resource management using the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/resource-management)\n+- [Azure Resource Manager client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/resource-manager)\n+- [DefaultAzureCredential class](https://learn.microsoft.com/dotnet/api/azure.identity.defaultazurecredential)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs b/dev/null
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
+++ b/dev/null
@@ -1,10 +0,0 @@
-﻿namespace AzureResourceGroupManager;
-
-class Program
-{
-    static void Main(string[] args)
-    {
-        Console.WriteLine("Hello, World!");
-    }
-}
-



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
@@ -1,0 +1,109 @@
+using Azure;
+using Azure.Core;
+using Azure.Identity;
+using Azure.ResourceManager;
+using Azure.ResourceManager.Resources;
+
+namespace AzureResourceGroupManager;
+
+internal static class Program
+{
+    private static async Task<int> Main()
+    {
+        ResourceGroupResource? createdResourceGroup = null;
+
+        try
+        {
+            DefaultAzureCredential credential = new();
+            ArmClient armClient = new(credential);
+            SubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();
+            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();
+
+            string resourceGroupName =
+                $"sdk-rg-{DateTimeOffset.UtcNow:yyyyMMddHHmmss}-{Guid.NewGuid():N}"[..39];
+
+            Console.WriteLine(
+                $"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...");
+
+            ResourceGroupData resourceGroupData = new(AzureLocation.EastUS)
... (truncated)
```

### ✅ 17. powershell — 1844ms

**Input:**

```json
{"command":"dotnet build .\\AzureResourceGroupManager.csproj --no-restore --nologo","description":"Build the resource group manager","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\Program.cs(53,66): error CS1061: 'ResourceGroupData' does not contain a definition for 'Properties' and no accessible extension method 'Properties' accepting a first argument of type 'ResourceGroupData' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\AzureResourceGroupManager.csproj]

Build FAILED.

C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\Program.cs(53,66): error CS1061: 'ResourceGroupData' does not contain a definition for 'Properties' and no accessible extension method 'Properties' accepting a first argument of type 'ResourceGroupData' could be found (are you missing a using directive or an assembly reference?) [C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\AzureResourceGroupManager.csproj]
    0 Warning(s)
    1 Error(s)

Time Elapsed 00:00:00.59
<shellId: 2 completed with exit code 1>
```

### ✅ 18. apply_patch — 658ms

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\\Program.cs\n@@\n             Console.WriteLine($\"Name: {resourceGroupDetails.Data.Name}\");\n             Console.WriteLine($\"Location: {resourceGroupDetails.Data.Location}\");\n             Console.WriteLine($\"Resource ID: {resourceGroupDetails.Id}\");\n-            Console.WriteLine(\n-                $\"Provisioning state: {resourceGroupDetails.Data.Properties.ProvisioningState}\");\n+            Console.WriteLine($\"Tag count: {resourceGroupDetails.Data.Tags.Count}\");\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307/Program.cs
@@ -49,8 +49,7 @@
             Console.WriteLine($"Name: {resourceGroupDetails.Data.Name}");
             Console.WriteLine($"Location: {resourceGroupDetails.Data.Location}");
             Console.WriteLine($"Resource ID: {resourceGroupDetails.Id}");
-            Console.WriteLine(
-                $"Provisioning state: {resourceGroupDetails.Data.Properties.ProvisioningState}");
+            Console.WriteLine($"Tag count: {resourceGroupDetails.Data.Tags.Count}");
 
             const string tagName = "ManagedBy";
             const string tagValue = "Azure.ResourceManager";

```

### ✅ 19. powershell — 1937ms

**Input:**

```json
{"command":"dotnet build .\\AzureResourceGroupManager.csproj --no-restore --nologo","description":"Rebuild the resource group manager","initial_wait":180,"mode":"sync"}
```

**Output:**

```
AzureResourceGroupManager -> C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-skills-azure-skill-mcp-691285307\bin\Debug\net10.0\AzureResourceGroupManager.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:00.68
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `AzureResourceGroupManager.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll ground the sample in the current Azure.ResourceManager APIs, then scaffold and run it locally without contacting Azure.

The workspace is empty and .NET 10 is available. I’m creating a minimal console project and pinning current Azure SDK packages before implementing the lifecycle.

The SDK exposes resource-group tagging directly through `AddTagAsync`. I’m using a generated name to avoid overwriting or deleting an existing group, with best-effort cleanup if an intermediate operation fails.

`ResourceGroupData` in the pinned SDK does not expose a provisioning-state property. I’m limiting the detail output to supported fields—name, location, resource ID, and tags—then rebuilding.

Implemented a complete C# console application:

- `Program.cs` — authenticates with `DefaultAzureCredential`, creates an `eastus` resource group, lists groups, retrieves details, adds a tag, and deletes it.
- `AzureResourceGroupManager.csproj` — uses `Azure.Identity` **1.21.0** and `Azure.ResourceManager` **1.14.0**.
- `README.md` — includes setup, permissions, execution instructions, and cleanup behavior.

The program uses targeted handling for authentication failures, Azure HTTP errors, cancellation, and best-effort cleanup. Run with `dotnet run`; it will create and delete a real Azure resource group in the credential’s default subscription.

Reference: [Resource management using the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/resource-management)

## Reviewer Notes (LLM-as-Judge)

### Summary

6/6 criteria passed. The implementation fulfills the requested resource-group lifecycle using the modern Azure.ResourceManager SDK and asynchronous API equivalents.

### Strengths

- Uses DefaultAzureCredential with ArmClient and the default subscription.
- Implements create, list, retrieve, tag, and completed deletion operations.
- Includes required package references and targeted error handling with cleanup.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (6/6)
      - `Azure.ResourceManager` NuGet package: Pass
      - `ArmClient` creation with `DefaultAzureCredential`: Pass
      - `GetDefaultSubscription()` and `GetResourceGroups()` collection: Pass
      - `CreateOrUpdate()`, `Get()`, `GetAll()` operations: Pass
      - Tag management with `SetTags()` or `AddTag()`: Pass
      - `Delete()` with `WaitForCompletion()`: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-dotnet-rg-crud --config dotnet-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
