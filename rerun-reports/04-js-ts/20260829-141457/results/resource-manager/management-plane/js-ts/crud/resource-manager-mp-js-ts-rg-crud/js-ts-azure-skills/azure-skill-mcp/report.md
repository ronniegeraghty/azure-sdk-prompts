# Evaluation Report: resource-manager-mp-js-ts-rg-crud

**Config:** js-ts-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 331.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-js-ts-rg-crud` |
| Config | js-ts-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 15/18 |
| Duration | 331.8s |
| Timestamp | 2026-08-29T08:53:19Z |
| Files Generated | 5 |
| Event Count | 3849 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 181.4s |
| Review | 149.2s |
| **Total** | **331.8s** |

## Configuration

- **name:** js-ts-azure-skills/azure-skill-mcp
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
| Token Usage | in=576856 out=3612 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. skill — 1789ms

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

### ✅ 2. azure-get_azure_bestpractices (via azure) — 5522ms

**Input:**

```json
{"intent":"Generate a TypeScript program using @azure/identity DefaultAzureCredential and @azure/arm-resources ResourceManagementClient to create, list, get, tag, and delete a resource group with correct async/await and safe cleanup patterns.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 5889ms

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

### ✅ 4. glob — 5158ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 15857ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure TypeScript SDK code-generation best practices for a management-plane resource-group CRUD example.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. powershell — 16221ms

**Input:**

```json
{"command":"npm view @azure/identity version; npm view @azure/arm-resources version; npm view typescript version; npm view tsx version; npm view @types/node version","description":"Check current npm package versions","initial_wait":120,"mode":"sync"}
```

**Output:**

```
4.13.2
8.0.0
7.0.2
4.23.12
26.2.0
<shellId: 0 completed with exit code 0>
```

### ✅ 7. azure-documentation (via azure) — 3182ms

**Input:**

```json
{"intent":"Find official TypeScript Azure SDK documentation and samples for @azure/arm-resources ResourceManagementClient resourceGroups.createOrUpdate, list iteration, get, update, and beginDeleteAndWait, plus DefaultAzureCredential.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ❌ 8. web_fetch — 2822ms

**Input:**

```json
{"max_length":8000,"url":"https://www.npmjs.com/package/@azure/arm-resources"}
```

**Error:**

```
Error: Failed to fetch https://www.npmjs.com/package/@azure/arm-resources - status code 403
```

### ❌ 9. web_fetch — 2818ms

**Input:**

```json
{"max_length":6000,"url":"https://www.npmjs.com/package/@azure/identity"}
```

**Error:**

```
Error: Failed to fetch https://www.npmjs.com/package/@azure/identity - status code 403
```

### ✅ 10. azure-documentation (via azure) — 2576ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve the latest official TypeScript sample for resource group lifecycle operations using @azure/arm-resources and DefaultAzureCredential.","parameters":{"language":"typescript","query":"TypeScript @azure/arm-resources ResourceManagementClient resourceGroups createOrUpdate list get update beginDeleteAndWait DefaultAzureCredential"}}
```

**Output:**

```
{"results":[{"title":"ResourcesOperations interface-Property Details","content":"# ResourcesOperations interface\r\n\r\n- Package:\r\n    - [@azure/arm-resources](https://learn.microsoft.com/en-us/javascript/api/@azure/arm-resources/)\r\n\r\nInterface representing a Resources operations.\r\n\r\n## Property Details\r\n\r\n### beginCreateOrUpdateAndWait\r\n\r\n Warning \r\n\r\nThis API is now deprecated.\r\n\r\nuse createOrUpdate instead\r\n\r\n\u0060\u0060\u0060typescript\r\nbeginCreateOrUpdateAndWait: (resourceGroupName: string, resourceProviderNamespace: string, parentResourcePath: string, resourceType: string, resourceName: string, apiVersion: string, parameters: GenericResource, options?: ResourcesCreateOrUpdateOptionalParams) =\u003E Promise\u003CGenericResource\u003E\r\n\u0060\u0060\u0060\r\n\r\n#### Property Value\r\n\r\n(resourceGroupName: string, resourceProviderNamespace: string, parentResourcePath: string, resourceType: string, resourceName: string, apiVersion: string, parameters: GenericResource, options?: ResourcesCreateOrUpdateOptionalParams) =\u0026gt; Promise\u0026lt;GenericResource\u0026gt;","contentUrl":"https://learn.microsoft.com/javascript/api/@azure/arm-resources/resourcesoperations?view=azure-node-latest"},{"title":"Resources interface-Method Details","content":"# Resources interface\r\n\r\n- Package:\r\n    - [@azure/arm-resources-profile-2020-09-01-hybrid](https://learn.microsoft.com/en-us/javascript/api/@azure/arm-resources-profile-2020-09-01-hybrid/)\r\n\r\nInterface representing a Resources.\r\n\r\n## Method Details\r\n\r\n### beginDeleteAndWait(string, string, string, string, string, string, ResourcesDeleteOptionalParams)\r\n\r\nDeletes a resource.\r\n\r\n\u0060\u0060\u0060typescript\r\nfunction beginDeleteAndWait(resourceGroupName: string, resourceProviderNamespace: string, parentResourcePath: string, resourceType: string, resourceName: string, apiVersion: string, options?: ResourcesDeleteOptionalParams): Promise\u003Cvoid\u003E\r\n\u006
... (truncated)
```

### ✅ 11. azure-documentation (via azure) — 2197ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official API references for the TypeScript ResourceManagementClient and DefaultAzureCredential.","parameters":{"query":"site:learn.microsoft.com javascript api @azure/arm-resources ResourceGroups createOrUpdate beginDeleteAndWait DefaultAzureCredential TypeScript"}}
```

**Output:**

```
{"results":[{"description":"description: Imports DefaultAzureCredential from @azure/identity and ResourceManagementClient from @azure/arm-resources, then uses these to authenticate and list all Azure resource groups accessible with the given subscription ID.\npackage: @azure/identity, @azure/arm-resources\nlanguage: typescript\n","codeSnippet":"import { DefaultAzureCredential } from \u0022@azure/identity\u0022;\r\nimport { ResourceManagementClient } from \u0022@azure/arm-resources\u0022;\r\n\r\nconst subscriptionId = process.env.AZURE_SUBSCRIPTION_ID!;\r\nif (!subscriptionId) {\r\n  throw new Error(\u0022AZURE_SUBSCRIPTION_ID environment variable is not set.\u0022);\r\n}\r\n\r\nconsole.log(\u0060Using Subscription ID: ${subscriptionId}\u0060);\r\n\r\nasync function main() {\r\n\r\n    const credential = new DefaultAzureCredential();\r\n    const client = new ResourceManagementClient(credential, subscriptionId);\r\n\r\n    let i=0;\r\n\r\n    for await (const item of client.resourceGroups.list()) {\r\n        console.log(\u0060${\u002B\u002Bi}: ${item.name}\u0060);\r\n    }\r\n    console.log(\u0060Found ${i} resource group(s).\u0060);\r\n}\r\n\r\nmain().catch((err) =\u003E {\r\n  console.error(err);\r\n});","link":"https://learn.microsoft.com/azure/developer/javascript/sdk/use-azure-sdk#create-an-sdk-client-and-call-methods","language":"typescript"},{"description":"description: Creates and authenticates an ArmClient instance using DefaultAzureCredential, retrieves the default subscription, and creates or updates a resource group with a specified name and location.\npackage: Azure.Identity, Azure.ResourceManager\nlanguage: c\n","codeSnippet":"ArmClient armClient = new ArmClient(new DefaultAzureCredential());\r\nSubscriptionResource subscription = await armClient.GetDefaultSubscriptionAsync();\r\nResourceGroupCollection rgCollection = subscription.GetResourceGroups();\r\n// With the collection, we can create a new resource group with an specific name\r\nstring rgName 
... (truncated)
```

### ✅ 12. rg — 855ms

**Input:**

```json
{"-n":true,"head_limit":60,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787993675282-copilot-tool-output-5124a59176894bc086d8e8f05f2ae786.txt","pattern":"ResourceGroupsOperations|beginDeleteAndWait|createOrUpdate\\(|update\\(|list\\(|resource group"}
```

**Output:**

```
Output too large to read at once (23.8 KB). Saved to: C:\Users\weidxu\AppData\Local\Temp\copilot-tool-output-1787993680547-1.txt
Consider using tools like rg (for searching), head/tail (for viewing start/end), view with view_range (for specific sections), or jq (for JSON) to examine portions of the output.

Preview (first 500 chars):
C:\Users\weidxu\AppData\Local\Temp\1787993675282-copilot-tool-output-5124a59176894bc086d8e8f05f2ae786.txt:1:{"results":[{"title":"ResourcesOperations interface-Property Details","content":"# ResourcesOperations interface\r\n\r\n- Package:\r\n    - [@azure/arm-resources](https://learn.microsoft.com/en-us/javascript/api/@azure/arm-resources/)\r\n\r\nInterface representing a Resources operations.\r\n\r\n## Property Details\r\n\r\n### beginCreateOrUpdateAndWait\r\n\r\n Warning \r\n\r\nThis API is no
```

### ✅ 13. apply_patch — 923ms

**Input:**

```json
"*** Begin Patch\n*** Add File: package.json\n+{\n+  \"name\": \"azure-resource-group-management-example\",\n+  \"version\": \"1.0.0\",\n+  \"private\": true,\n+  \"type\": \"module\",\n+  \"scripts\": {\n+    \"build\": \"tsc\",\n+    \"start\": \"tsx src/index.ts\"\n+  },\n+  \"dependencies\": {\n+    \"@azure/arm-resources\": \"8.0.0\",\n+    \"@azure/identity\": \"4.13.2\"\n+  },\n+  \"devDependencies\": {\n+    \"@types/node\": \"26.2.0\",\n+    \"tsx\": \"4.23.12\",\n+    \"typescript\": \"7.0.2\"\n+  }\n+}\n*** Add File: tsconfig.json\n+{\n+  \"compilerOptions\": {\n+    \"target\": \"ES2022\",\n+    \"module\": \"NodeNext\",\n+    \"moduleResolution\": \"NodeNext\",\n+    \"outDir\": \"dist\",\n+    \"strict\": true,\n+    \"noUncheckedIndexedAccess\": true,\n+    \"exactOptionalPropertyTypes\": true,\n+    \"skipLibCheck\": true\n+  },\n+  \"include\": [\"src/**/*.ts\"]\n+}\n*** Add File: src/index.ts\n+import { ResourceManagementClient } from \"@azure/arm-resources\";\n+import { DefaultAzureCredential } from \"@azure/identity\";\n+\n+const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;\n+if (!subscriptionId) {\n+  throw new Error(\"Set the AZURE_SUBSCRIPTION_ID environment variable.\");\n+}\n+\n+const resourceGroupName =\n+  process.env.AZURE_RESOURCE_GROUP_NAME ?? `sdk-rg-example-${Date.now()}`;\n+const location = \"eastus\";\n+\n+async function main(): Promise\u003cvoid\u003e {\n+  const credential = new DefaultAzureCredential();\n+  const client = new ResourceManagementClient(credential, subscriptionId);\n+  let resourceGroupCreated = false;\n+\n+  try {\n+    console.log(`Creating resource group \"${resourceGroupName}\" in \"${location}\"...`);\n+    const created = await client.resourceGroups.createOrUpdate(resourceGroupName, {\n+      location,\n+    });\n+    resourceGroupCreated = true;\n+    console.log(`Created: ${created.id}`);\n+\n+    console.log(\"\\nResource groups in the subscription:\");\n+    for await (const resourceGroup of client.resourceGroups.list()) {\n+      console.log(`- ${resourceGroup.name} (${resourceGroup.location})`);\n+    }\n+\n+    const details = await client.resourceGroups.get(resourceGroupName);\n+    console.log(\"\\nCreated resource group details:\", {\n+      id: details.id,\n+      name: details.name,\n+      location: details.location,\n+      tags: details.tags,\n+    });\n+\n+    const updated = await client.resourceGroups.update(resourceGroupName, {\n+      tags: {\n+        ...details.tags,\n+        environment: \"demo\",\n+      },\n+    });\n+    console.log(\"\\nUpdated tags:\", updated.tags);\n+  } finally {\n+    if (resourceGroupCreated) {\n+      console.log(`\\nDeleting resource group \"${resourceGroupName}\"...`);\n+      await client.resourceGroups.beginDeleteAndWait(resourceGroupName);\n+      console.log(\"Resource group deleted.\");\n+    }\n+  }\n+}\n+\n+main().catch((error: unknown) =\u003e {\n+  console.error(\"Resource group operation failed:\", error);\n+  process.exitCode = 1;\n+});\n*** Add File: README.md\n+# Azure Resource Group Management with TypeScript\n+\n+This example uses the Azure management-plane SDK to create, list, read, tag,\n+and delete a resource group. The `finally` block deletes the resource group\n+even when a later operation fails.\n+\n+## Packages\n+\n+Runtime packages:\n+\n+```powershell\n+npm install @azure/identity @azure/arm-resources\n+```\n+\n+TypeScript development packages:\n+\n+```powershell\n+npm install --save-dev typescript tsx @types/node\n+```\n+\n+## Configuration and execution\n+\n+Set the subscription ID and, optionally, a resource group name:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_RESOURCE_GROUP_NAME = \"sdk-rg-example\"\n+npm install\n+npm run build\n+npm start\n+```\n+\n+`DefaultAzureCredential` uses the available credential in its standard\n+credential chain. In Azure-hosted production environments, use a managed\n+identity with only the resource-group permissions the application requires.\n+\n+\u003e Running the program performs live create, update, and delete operations in\n+\u003e the configured Azure subscription.\n+\n+## References\n+\n+- [Azure SDK for JavaScript: create a client and call methods](https://learn.microsoft.com/azure/developer/javascript/sdk/use-azure-sdk#create-an-sdk-client-and-call-methods)\n+- [`@azure/arm-resources` overview](https://learn.microsoft.com/javascript/api/overview/azure/arm-resources-readme?view=azure-node-latest)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/package.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/package.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/package.json
@@ -1,0 +1,20 @@
+{
+  "name": "azure-resource-group-management-example",
+  "version": "1.0.0",
+  "private": true,
+  "type": "module",
+  "scripts": {
+    "build": "tsc",
+    "start": "tsx src/index.ts"
+  },
+  "dependencies": {
+    "@azure/arm-resources": "8.0.0",
+    "@azure/identity": "4.13.2"
+  },
+  "devDependencies": {
+    "@types/node": "26.2.0",
+    "tsx": "4.23.12",
+    "typescript": "7.0.2"
+  }
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json
@@ -1,0 +1,14 @@
+{
+  "compilerOptions": {
+    "target": "ES2022",
+    "module": "NodeNext",
+    "moduleResolution": "NodeNext",
+    "outDir": "dist",
+    "strict": true,
+    "noUncheckedIndexedAccess": true,
+    "exactOptionalPropertyTypes": true,
+    "skipLibCheck": true
+  },
+  "include": ["src/**/*.ts"]
+}
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/
... (truncated)
```

### ✅ 14. powershell — 50623ms

**Input:**

```json
{"command":"npm install --no-audit --no-fund \u0026\u0026 npm run build","description":"Install dependencies and compile TypeScript","initial_wait":180,"mode":"sync"}
```

**Output:**

```
added 53 packages in 43s

> azure-resource-group-management-example@1.0.0 build
> tsc

tsconfig.json(6,5): error TS5011: The common source directory of 'tsconfig.json' is './src'. The 'rootDir' setting must be explicitly set to this or another path to adjust your output's file layout.
  Visit https://aka.ms/ts6 for migration information.
<shellId: 1 completed with exit code 2>
```

### ✅ 15. apply_patch — 697ms

**Input:**

```json
"*** Begin Patch\n*** Update File: tsconfig.json\n@@\n     \"module\": \"NodeNext\",\n     \"moduleResolution\": \"NodeNext\",\n+    \"rootDir\": \"src\",\n     \"outDir\": \"dist\",\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/tsconfig.json
@@ -3,6 +3,7 @@
     "target": "ES2022",
     "module": "NodeNext",
     "moduleResolution": "NodeNext",
+    "rootDir": "src",
     "outDir": "dist",
     "strict": true,
     "noUncheckedIndexedAccess": true,

```

### ✅ 16. powershell — 1858ms

**Input:**

```json
{"command":"npm run build","description":"Compile TypeScript program","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-resource-group-management-example@1.0.0 build
> tsc

src/index.ts(15,59): error TS2769: No overload matches this call.
  The last overload gave the following error.
    Argument of type 'string | undefined' is not assignable to parameter of type 'string'.
      Type 'undefined' is not assignable to type 'string'.
<shellId: 2 completed with exit code 2>
```

### ✅ 17. apply_patch — 722ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/index.ts\n@@\n-const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;\n-if (!subscriptionId) {\n-  throw new Error(\"Set the AZURE_SUBSCRIPTION_ID environment variable.\");\n+function requireEnvironmentVariable(name: string): string {\n+  const value = process.env[name];\n+  if (!value) {\n+    throw new Error(`Set the ${name} environment variable.`);\n+  }\n+  return value;\n }\n \n+const subscriptionId = requireEnvironmentVariable(\"AZURE_SUBSCRIPTION_ID\");\n const resourceGroupName =\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/src/index.ts b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/src/index.ts
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/src/index.ts
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-js-ts-rg-crud-js-ts-azure-skills-azure-skill-mcp-1787000947/src/index.ts
@@ -1,11 +1,15 @@
 import { ResourceManagementClient } from "@azure/arm-resources";
 import { DefaultAzureCredential } from "@azure/identity";
 
-const subscriptionId = process.env.AZURE_SUBSCRIPTION_ID;
-if (!subscriptionId) {
-  throw new Error("Set the AZURE_SUBSCRIPTION_ID environment variable.");
+function requireEnvironmentVariable(name: string): string {
+  const value = process.env[name];
+  if (!value) {
+    throw new Error(`Set the ${name} environment variable.`);
+  }
+  return value;
 }
 
+const subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
 const resourceGroupName =
   process.env.AZURE_RESOURCE_GROUP_NAME ?? `sdk-rg-example-${Date.now()}`;
 const location = "eastus";

```

### ✅ 18. powershell — 1758ms

**Input:**

```json
{"command":"npm run build","description":"Compile corrected TypeScript program","initial_wait":120,"mode":"sync"}
```

**Output:**

```
> azure-resource-group-management-example@1.0.0 build
> tsc
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `package-lock.json`
- `package.json`
- `src\index.ts`
- `tsconfig.json`

## Copilot Response

I’ll produce a local TypeScript example with environment-based configuration and no live Azure operations.

The example will preserve existing tags during update and guarantee cleanup with `finally` after creation.

Created a complete TypeScript example:

- `src/index.ts` — resource-group create, async iteration, get, tag update, and `beginDeleteAndWait`
- `package.json` — required Azure SDK and TypeScript packages
- `tsconfig.json` — strict Node.js TypeScript configuration
- `README.md` — setup, authentication, execution, and Microsoft references

The program reads `AZURE_SUBSCRIPTION_ID` from the environment, uses `DefaultAzureCredential`, preserves existing tags during updates, and deletes the created group in a `finally` block.

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "sdk-rg-example" # optional
npm install
npm run build
npm start
```

Running `npm start` performs live create, update, and delete operations in the configured Azure subscription.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Provides all required Azure SDK dependencies and appropriate TypeScript tooling.
- Includes a valid tsconfig.json with strict type checking and NodeNext module configuration.
- Pins dependency versions and supplies a matching package-lock.json.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (8/8)
      - `@azure/arm-resources` and `@azure/identity` npm packages: Pass
      - `DefaultAzureCredential` for authentication: Pass
      - `ResourceManagementClient` constructor with credential and subscriptionId: Pass
      - `resourceGroups.createOrUpdate()` with resource group name and parameters: Pass
      - `resourceGroups.list()` with async iteration (`for await...of`): Pass
      - `resourceGroups.get()` for fetching details: Pass
      - `resourceGroups.beginDeleteAndWait()` for long-running delete: Pass
      - Tag updates via `resourceGroups.update()` with tags parameter: Pass
- js-ts.yaml (criteria file):
  - Correct @azure/ Scoped Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct @azure/ Scoped Packages**: All Azure SDK dependencies use @azure/ scoped packages (e.g., @azure/storage-blob, @azure/keyvault-secrets, @azure/identity). No deprecated unscoped packages (azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-*).: Pass
  - @azure/identity for Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**@azure/identity for Authentication**: Includes @azure/identity as a dependency. For local development, DefaultAzureCredential is acceptable. For production code, uses a specific Entra token credential (e.g., ManagedIdentityCredential, ClientSecretCredential, WorkloadIdentityCredential) or configures DefaultAzureCredential with the AZURE_TOKEN_CREDENTIALS environment variable set to "prod" or a specific token credential. No hardcoded connection strings, account keys, SAS tokens, or secrets in code.: Pass
  - Client Constructor with Endpoint and Credential (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Constructor with Endpoint and Credential**: SDK clients constructed with endpoint URL and credential object, e.g. new BlobServiceClient(url, credential) or new SecretClient(vaultUrl, credential). No legacy constructor patterns or connection string-based construction when identity-based auth is available.: Pass
  - Async/Await Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async/Await Pattern**: Uses async/await for all asynchronous operations. No raw .then()/.catch() chains when await is cleaner. No callback-based patterns. Top-level async function or IIFE for entry point.: Pass
  - Pagination with for-await-of (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination with for-await-of**: List/query operations use for-await-of iteration over the PagedAsyncIterableIterator returned by SDK methods. Does not collect all pages into an array in memory (e.g., no Array.from() or manual page accumulation). Uses byPage() when per-page control is needed.: Pass
  - LRO Pattern (beginXxx + pollUntilDone) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (beginXxx + pollUntilDone)**: Long-running operations use the begin* method prefix (e.g., beginDeleteSecret, beginCreateOrUpdate) and await poller.pollUntilDone(). No manual setTimeout/setInterval polling loops.: Fail
  - RestError Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**RestError Exception Handling**: Catches RestError (from @azure/core-rest-pipeline or the service package) with statusCode inspection for error-specific handling. Not just generic Error or try/catch without type checking.: Fail
  - No Deprecated Packages (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated Packages**: No deprecated or legacy Azure packages in package.json: azure-storage, azure-arm-*, ms-rest-azure, @azure/ms-rest-js, @azure/ms-rest-nodeauth, @azure/ms-rest-azure-js. All dependencies use the current @azure/* track.: Pass
  - Logging via @azure/logger (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Logging via @azure/logger**: Uses @azure/logger with setLogLevel() or the AZURE_LOG_LEVEL environment variable for SDK diagnostic logging. No custom logging wrappers around SDK internals.: Fail
  - package.json with Correct Dependencies (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**package.json with Correct Dependencies**: Includes a valid package.json with all required @azure/* dependencies listed. TypeScript projects include tsconfig.json and appropriate @types/* packages or use a recent TypeScript version with built-in types.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Correct @azure/ Scoped Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `@azure/identity for Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Client Constructor with Endpoint and Credential` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Async/Await Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Pagination with for-await-of` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `LRO Pattern (beginXxx + pollUntilDone)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `RestError Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated Packages` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| `Logging via @azure/logger` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `package.json with Correct Dependencies` | prompt_review | 100% | 1.00 | 1.0000 | 12.5% | ✅ |
| **Final** | | | **Σ 11.00** | **Σ 8.0000** | **72.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-js-ts-rg-crud --config js-ts-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
