# Evaluation Report: resource-manager-mp-dotnet-rg-crud

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 183.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-dotnet-rg-crud` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 6/6 |
| Duration | 183.0s |
| Timestamp | 2026-08-26T18:08:58Z |
| Files Generated | 2 |
| Event Count | 3915 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 155.1s |
| Review | 26.7s |
| **Total** | **183.0s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** dotnet-azure-tools/with-azure-tools

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=660964 out=3771 |
| Turn Count | 14 |

## Prompt Sent

```
Write a C# program that manages
Azure Resource Groups using the management plane SDK:
1. Authenticate using DefaultAzureCredential
2. Create a new resource group in "eastus" region
3. List all resource groups in the subscription
4. Get details of the created resource group
5. Add a tag to the resource group
6. Delete the resource group

Show required NuGet packages and proper error handling.
Use the Azure.ResourceManager SDK (not the older Microsoft.Azure.Management packages).
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get","intent":"Generate a local-only C# console sample using Azure.ResourceManager and DefaultAzureCredential to create, list, inspect, tag, and delete a resource group with proper error handling."}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720","pattern":"**/*"}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure SDK code-generation best practices to a C# Azure Resource Manager resource-group CRUD sample.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

### 🔧 7. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720"}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official current .NET Azure.ResourceManager examples for creating, listing, getting, tagging, and deleting resource groups.","learn":true}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"dotnet --version; Get-ChildItem -Force | Select-Object Name,Length,Mode","description":"Check .NET SDK and workspace contents","initial_wait":30,"mode":"sync"}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official current Azure.ResourceManager .NET resource group CRUD and tagging documentation.","parameters":{"query":"Azure.ResourceManager .NET ResourceGroupCollection CreateOrUpdateAsync GetAllAsync ResourceGroupResource SetTagsAsync DeleteAsync"}}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find official C# Azure.ResourceManager resource group CRUD and tagging samples.","parameters":{"language":"csharp","query":"Azure.ResourceManager C# create list get update tags delete resource group ArmClient DefaultAzureCredential"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the full official Azure Resource Manager .NET resource-management examples before implementing the sample.","parameters":{"url":"https://learn.microsoft.com/dotnet/azure/sdk/resource-management"}}
```

### 🔧 13. powershell

**Input:**

```json
{"command":"dotnet new console --framework net10.0 --use-program-main --force; dotnet add package Azure.Identity; dotnet add package Azure.ResourceManager; dotnet add package Azure.ResourceManager.Resources","description":"Create console project and add Azure SDK packages","initial_wait":180,"mode":"sync"}
```

### 🔧 14. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720\\Program.cs"}
```

### 🔧 15. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720.csproj"}
```

### 🔧 16. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720\\Program.cs\n@@\n-﻿namespace hyoka_resource_manager_mp_dotnet_rg_crud_dotnet_azure_tools_with_azure_tools_121524720;\n-\n-class Program\n-{\n-    static void Main(string[] args)\n-    {\n-        Console.WriteLine(\"Hello, World!\");\n-    }\n-}\n+﻿using Azure;\n+using Azure.Core;\n+using Azure.Identity;\n+using Azure.ResourceManager;\n+using Azure.ResourceManager.Resources;\n+\n+namespace hyoka_resource_manager_mp_dotnet_rg_crud_dotnet_azure_tools_with_azure_tools_121524720;\n+\n+internal static class Program\n+{\n+    private static async Task\u003cint\u003e Main(string[] args)\n+    {\n+        string resourceGroupName = args.Length \u003e 0\n+            ? args[0]\n+            : $\"resource-manager-demo-{Guid.NewGuid():N}\"[..37];\n+\n+        using CancellationTokenSource cancellationTokenSource = new();\n+        Console.CancelKeyPress += (_, eventArgs) =\u003e\n+        {\n+            eventArgs.Cancel = true;\n+            cancellationTokenSource.Cancel();\n+        };\n+\n+        ResourceGroupResource? resourceGroup = null;\n+        bool deleted = false;\n+        int exitCode = 0;\n+\n+        try\n+        {\n+            TokenCredential credential = new DefaultAzureCredential();\n+            ArmClient armClient = new(credential);\n+            SubscriptionResource subscription =\n+                await armClient.GetDefaultSubscriptionAsync(cancellationTokenSource.Token);\n+            ResourceGroupCollection resourceGroups = subscription.GetResourceGroups();\n+\n+            Console.WriteLine(\n+                $\"Creating resource group '{resourceGroupName}' in '{AzureLocation.EastUS}'...\");\n+            ResourceGroupData resourceGroupData = new(AzureLocation.EastUS);\n+            ArmOperation\u003cResourceGroupResource\u003e createOperation =\n+                await resourceGroups.CreateOrUpdateAsync(\n+                    WaitUntil.Completed,\n+                    resourceGroupName,\n+                    resourceGroupData,\n+                    cancellationTokenSource.Token);\n+            resourceGroup = createOperation.Value;\n+            Console.WriteLine($\"Created: {resourceGroup.Id}\");\n+\n+            Console.WriteLine(\"\\nResource groups in the subscription:\");\n+            await foreach (ResourceGroupResource item in\n+                resourceGroups.GetAllAsync().WithCancellation(cancellationTokenSource.Token))\n+            {\n+                Console.WriteLine($\"- {item.Data.Name} ({item.Data.Location})\");\n+            }\n+\n+            Console.WriteLine($\"\\nGetting details for '{resourceGroupName}'...\");\n+            resourceGroup = (await resourceGroups.GetAsync(\n+                resourceGroupName,\n+                cancellationTokenSource.Token)).Value;\n+            PrintDetails(resourceGroup.Data);\n+\n+            const string tagName = \"managed-by\";\n+            const string tagValue = \"azure-resource-manager-sdk\";\n+            Console.WriteLine($\"\\nAdding tag '{tagName}={tagValue}'...\");\n+            resourceGroup = (await resourceGroup.AddTagAsync(\n+                tagName,\n+                tagValue,\n+                cancellationTokenSource.Token)).Value;\n+            PrintDetails(resourceGroup.Data);\n+\n+            Console.WriteLine($\"\\nDeleting resource group '{resourceGroupName}'...\");\n+            await resourceGroup.DeleteAsync(\n+                WaitUntil.Completed,\n+                cancellationTokenSource.Token);\n+            deleted = true;\n+            Console.WriteLine(\"Resource group deleted.\");\n+        }\n+        catch (CredentialUnavailableException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"No credential was available for DefaultAzureCredential: {exception.Message}\");\n+            exitCode = 1;\n+        }\n+        catch (AuthenticationFailedException exception)\n+        {\n+            Console.Error.WriteLine($\"Azure authentication failed: {exception.Message}\");\n+            exitCode = 1;\n+        }\n+        catch (RequestFailedException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Azure request failed (HTTP {exception.Status}, code '{exception.ErrorCode}'): \" +\n+                exception.Message);\n+            exitCode = 1;\n+        }\n+        catch (OperationCanceledException) when (cancellationTokenSource.IsCancellationRequested)\n+        {\n+            Console.Error.WriteLine(\"Operation canceled.\");\n+            exitCode = 2;\n+        }\n+        finally\n+        {\n+            if (resourceGroup is not null \u0026\u0026 !deleted)\n+            {\n+                try\n+                {\n+                    Console.WriteLine(\n+                        $\"\\nCleaning up resource group '{resourceGroupName}' after the failure...\");\n+                    await resourceGroup.DeleteAsync(WaitUntil.Completed, CancellationToken.None);\n+                    Console.WriteLine(\"Cleanup completed.\");\n+                }\n+                catch (RequestFailedException cleanupException)\n+                {\n+                    Console.Error.WriteLine(\n+                        $\"Cleanup failed (HTTP {cleanupException.Status}, \" +\n+                        $\"code '{cleanupException.ErrorCode}'): {cleanupException.Message}\");\n+                    exitCode = 1;\n+                }\n+            }\n+        }\n+\n+        return exitCode;\n+    }\n+\n+    private static void PrintDetails(ResourceGroupData data)\n+    {\n+        Console.WriteLine($\"Name: {data.Name}\");\n+        Console.WriteLine($\"Location: {data.Location}\");\n+        Console.WriteLine($\"Resource ID: {data.Id}\");\n+        Console.WriteLine(\n+            $\"Tags: {(data.Tags.Count == 0 ? \"(none)\" : string.Join(\", \", data.Tags))}\");\n+    }\n+}\n*** End Patch\n"
```

### 🔧 17. powershell

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Build the Azure resource group sample","initial_wait":120,"mode":"sync"}
```

### 🔧 18. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720\\Program.cs","view_range":[65,112]}
```

### 🔧 19. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720\\Program.cs\n@@\n             await resourceGroup.DeleteAsync(\n                 WaitUntil.Completed,\n-                cancellationTokenSource.Token);\n+                cancellationToken: cancellationTokenSource.Token);\n@@\n-                    await resourceGroup.DeleteAsync(WaitUntil.Completed, CancellationToken.None);\n+                    await resourceGroup.DeleteAsync(\n+                        WaitUntil.Completed,\n+                        cancellationToken: CancellationToken.None);\n*** End Patch\n"
```

### 🔧 20. powershell

**Input:**

```json
{"command":"dotnet build --no-restore","description":"Rebuild the Azure resource group sample","initial_wait":120,"mode":"sync"}
```

## Generated Files

- `Program.cs`
- `hyoka-resource-manager-mp-dotnet-rg-crud-dotnet-azure-tools-with-azure-tools-121524720.csproj`

## Copilot Response

I’ll build this as a runnable console sample with cleanup in `finally`, so a failure after creation does not leave the demo resource group behind.

Implemented a runnable C# console project using the modern `Azure.ResourceManager` management-plane SDK. `Program.cs` creates an `eastus` resource group, lists resource groups, retrieves details, adds a tag, and deletes it. Cleanup runs from `finally` if an intermediate operation fails.

**Required packages:**

```powershell
dotnet add package Azure.Identity
dotnet add package Azure.ResourceManager
dotnet add package Azure.ResourceManager.Resources
```

Run with an automatically generated resource-group name:

```powershell
dotnet run
```

Or provide a name:

```powershell
dotnet run -- my-resource-group
```

`DefaultAzureCredential` uses available developer credentials or managed identity and selects the default subscription. The identity needs resource-group management permissions, such as **Contributor**, at subscription scope.

Reference: [Resource management using the Azure SDK for .NET](https://learn.microsoft.com/dotnet/azure/sdk/resource-management)

## Reviewer Notes (LLM-as-Judge)

### Summary

6/6 criteria passed. The generated project uses the modern Azure Resource Manager SDK and implements the complete resource-group lifecycle with asynchronous completion and structured error handling.

### Strengths

- Includes all required Azure SDK package references.
- Uses DefaultAzureCredential with ArmClient and the default subscription.
- Implements create, list, get, tag, and completed deletion operations.
- Handles authentication, Azure request, cancellation, and cleanup failures.

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
hyoka run --prompt-id resource-manager-mp-dotnet-rg-crud --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
