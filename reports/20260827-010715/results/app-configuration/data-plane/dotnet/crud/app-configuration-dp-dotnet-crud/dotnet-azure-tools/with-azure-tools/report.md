# Evaluation Report: app-configuration-dp-dotnet-crud

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 167.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-dotnet-crud` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 6/7 |
| Duration | 167.8s |
| Timestamp | 2026-08-26T17:07:15Z |
| Files Generated | 2 |
| Event Count | 2152 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 137.4s |
| Review | 29.3s |
| **Total** | **167.8s** |

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
| Token Usage | in=348627 out=2748 |
| Turn Count | 8 |

## Prompt Sent

```
Write a C# program that manages
configuration settings in Azure App Configuration:
1. Create a ConfigurationClient using a connection string
2. Set a configuration setting with key "app:Settings:FontSize" and value "24"
3. Set a configuration setting with a label "Production"
4. Get the setting by key and print its value
5. List all settings with the prefix "app:Settings:" using GetConfigurationSettings
6. Create a feature flag setting for "BetaFeature" that is enabled
7. Delete the setting

Show required NuGet packages and proper error handling with RequestFailedException.
```

## Tool Calls

### 🔧 1. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Generate a local-only C# console program using Azure.Data.AppConfiguration with a connection string, CRUD operations, labels, listing, feature flags, and RequestFailedException handling.","learn":true}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a C# console program for Azure App Configuration data-plane CRUD and feature flags.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 3. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official C# Azure.Data.AppConfiguration examples for ConfigurationClient connection string, SetConfigurationSetting, labeled settings, GetConfigurationSetting, GetConfigurationSettings with SettingSelector key prefix, FeatureFlagConfigurationSetting, DeleteConfigurationSetting, and RequestFailedException.","learn":true}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-2102961491","pattern":"**/*"}
```

### 🔧 5. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Locate official Azure App Configuration .NET data-plane client documentation and CRUD/feature flag APIs.","parameters":{"query":"Azure.Data.AppConfiguration ConfigurationClient C# SetConfigurationSetting GetConfigurationSettings SettingSelector KeyFilter FeatureFlagConfigurationSetting DeleteConfigurationSetting"}}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find latest official C# samples for Azure App Configuration setting CRUD, labels, listing by prefix, and feature flags.","parameters":{"language":"csharp","query":"Azure.Data.AppConfiguration ConfigurationClient SetConfigurationSetting FeatureFlagConfigurationSetting GetConfigurationSettings SettingSelector C#"}}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"dotnet --version; dotnet nuget list source","description":"Check .NET SDK and NuGet sources","initial_wait":30,"mode":"sync"}
```

### 🔧 8. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official Azure.Data.AppConfiguration .NET client README identified by search.","parameters":{"url":"https://learn.microsoft.com/dotnet/api/overview/azure/data.appconfiguration-readme?view=azure-dotnet"}}
```

### 🔧 9. rg

**Input:**

```json
{"-C":2,"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787764110841-copilot-tool-output-f084609fcec44c488b3a8ba2d59540d6.txt","pattern":"FeatureFlagConfigurationSetting|SettingSelector|KeyFilter|GetConfigurationSettings|ConfigurationClient\\("}
```

### 🔧 10. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"FeatureFlagConfigurationSetting repo:Azure/azure-sdk-for-net path:sdk/appconfiguration/Azure.Data.AppConfiguration"}
```

### 🔧 11. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"page":1,"perPage":10,"query":"GetConfigurationSettings(new SettingSelector repo:Azure/azure-sdk-for-net path:sdk/appconfiguration/Azure.Data.AppConfiguration"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: AppConfigurationCrud.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Data.AppConfiguration\" Version=\"1.11.1\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure;\n+using Azure.Data.AppConfiguration;\n+\n+const string connectionStringVariable = \"AZURE_APPCONFIG_CONNECTION_STRING\";\n+const string settingKey = \"app:Settings:FontSize\";\n+const string settingPrefix = \"app:Settings:\";\n+const string productionLabel = \"Production\";\n+\n+string? connectionString = Environment.GetEnvironmentVariable(connectionStringVariable);\n+if (string.IsNullOrWhiteSpace(connectionString))\n+{\n+    Console.Error.WriteLine(\n+        $\"Set the {connectionStringVariable} environment variable to an Azure App Configuration connection string.\");\n+    return 1;\n+}\n+\n+try\n+{\n+    var client = new ConfigurationClient(connectionString);\n+\n+    client.SetConfigurationSetting(settingKey, \"24\");\n+    Console.WriteLine($\"Set '{settingKey}' to '24'.\");\n+\n+    var productionSetting = new ConfigurationSetting(settingKey, \"24\", productionLabel);\n+    client.SetConfigurationSetting(productionSetting);\n+    Console.WriteLine($\"Set '{settingKey}' with label '{productionLabel}'.\");\n+\n+    ConfigurationSetting setting = client.GetConfigurationSetting(settingKey).Value;\n+    Console.WriteLine($\"Value for '{settingKey}': {setting.Value}\");\n+\n+    Console.WriteLine($\"Settings with prefix '{settingPrefix}':\");\n+    var selector = new SettingSelector\n+    {\n+        KeyFilter = $\"{settingPrefix}*\"\n+    };\n+\n+    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))\n+    {\n+        string label = matchingSetting.Label ?? \"(no label)\";\n+        Console.WriteLine($\"  {matchingSetting.Key} = {matchingSetting.Value} [{label}]\");\n+    }\n+\n+    var featureFlag = new FeatureFlagConfigurationSetting(\"BetaFeature\", isEnabled: true);\n+    client.SetConfigurationSetting(featureFlag);\n+    Console.WriteLine(\"Enabled feature flag 'BetaFeature'.\");\n+\n+    client.DeleteConfigurationSetting(settingKey);\n+    Console.WriteLine($\"Deleted the unlabeled setting '{settingKey}'.\");\n+\n+    return 0;\n+}\n+catch (RequestFailedException ex)\n+{\n+    Console.Error.WriteLine(\n+        $\"Azure App Configuration request failed. Status: {ex.Status}; \" +\n+        $\"Error code: {ex.ErrorCode ?? \"(none)\"}; Message: {ex.Message}\");\n+    return 1;\n+}\n+catch (ArgumentException ex)\n+{\n+    Console.Error.WriteLine($\"The connection string is invalid: {ex.Message}\");\n+    return 1;\n+}\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Build the C# sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `AppConfigurationCrud.csproj`
- `Program.cs`

## Copilot Response

I’ll inspect the local project and apply current Azure SDK patterns before adding the sample.

The workspace is empty, so I’ll create a minimal console project and keep the connection string in an environment variable.

I found the current `Azure.Data.AppConfiguration` 1.11.1 API; the sample will use its typed feature-flag model and prefix selector.

Implemented the C# console sample in:

- `Program.cs` — CRUD operations, labeled setting, prefix listing, enabled `BetaFeature` flag, deletion, and `RequestFailedException` handling.
- `AppConfigurationCrud.csproj` — references `Azure.Data.AppConfiguration` version `1.11.1`.

Run with:

```powershell
$env:AZURE_APPCONFIG_CONNECTION_STRING = "<connection-string>"
dotnet run
```

The connection string is read from the environment and is never hardcoded. SDK reference: [Azure App Configuration client library for .NET](https://learn.microsoft.com/dotnet/api/overview/azure/data.appconfiguration-readme?view=azure-dotnet).

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 7 criteria passed. The sample covers package setup, client creation, setting creation, prefix filtering, feature flags, deletion, and Azure error handling, but does not retrieve a labeled setting by key and label.

### Strengths

- Uses the required Azure.Data.AppConfiguration package.
- Reads the connection string from an environment variable rather than hardcoding it.
- Uses SettingSelector for prefix filtering and FeatureFlagConfigurationSetting for the enabled feature flag.
- Provides explicit RequestFailedException handling.

### Issues

- The retrieval call omits the Production label, so GetConfigurationSetting(key, label) is not demonstrated.

## Grader Results

- config-values.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (6/7)
      - `Azure.Data.AppConfiguration` NuGet package: Pass
      - `ConfigurationClient` creation with connection string or `DefaultAzureCredential`: Pass
      - `SetConfigurationSetting()` with key, value, and optional label: Pass
      - `GetConfigurationSetting()` by key and label: Fail
      - `GetConfigurationSettings()` with `SettingSelector` for filtering: Pass
      - Feature flag configuration settings with `FeatureFlagConfigurationSetting`: Pass
      - `DeleteConfigurationSetting()` and `RequestFailedException` handling: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8571** | **85.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-dotnet-crud --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
