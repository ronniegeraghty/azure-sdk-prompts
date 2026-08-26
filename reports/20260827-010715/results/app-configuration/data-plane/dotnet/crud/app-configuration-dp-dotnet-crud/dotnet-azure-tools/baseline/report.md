# Evaluation Report: app-configuration-dp-dotnet-crud

**Config:** dotnet-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 122.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-dotnet-crud` |
| Config | dotnet-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 6/7 |
| Duration | 122.0s |
| Timestamp | 2026-08-26T17:11:46Z |
| Files Generated | 3 |
| Event Count | 2101 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 99.5s |
| Review | 21.9s |
| **Total** | **122.0s** |

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
| Token Usage | in=110349 out=1842 |
| Turn Count | 5 |

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

### 🔧 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-tools-baseline-1327657672","pattern":"**/*"}
```

### 🔧 2. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-tools-baseline-1327657672","pattern":"**/*.csproj"}
```

### 🔧 3. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-tools-baseline-1327657672","pattern":"**/*.cs"}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure.Data.AppConfiguration latest .NET FeatureFlagConfigurationSetting constructor SetConfigurationSetting ConfigurationClient 2026"}
```

### 🔧 5. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure.Data.AppConfiguration .NET ConfigurationClient GetConfigurationSettings SettingSelector KeyFilter DeleteConfigurationSetting example"}
```

### 🔧 6. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: AppConfigurationSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Data.AppConfiguration\" Version=\"1.11.1\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure;\n+using Azure.Data.AppConfiguration;\n+\n+const string connectionStringVariable = \"AZURE_APP_CONFIGURATION_CONNECTION_STRING\";\n+const string key = \"app:Settings:FontSize\";\n+const string value = \"24\";\n+const string productionLabel = \"Production\";\n+\n+string? connectionString = Environment.GetEnvironmentVariable(connectionStringVariable);\n+if (string.IsNullOrWhiteSpace(connectionString))\n+{\n+    Console.Error.WriteLine(\n+        $\"Set the {connectionStringVariable} environment variable to an Azure App Configuration connection string.\");\n+    return 1;\n+}\n+\n+try\n+{\n+    var client = new ConfigurationClient(connectionString);\n+\n+    client.SetConfigurationSetting(key, value);\n+    Console.WriteLine($\"Set {key} = {value}\");\n+\n+    client.SetConfigurationSetting(key, value, productionLabel);\n+    Console.WriteLine($\"Set {key} = {value} with label {productionLabel}\");\n+\n+    ConfigurationSetting setting = client.GetConfigurationSetting(key).Value;\n+    Console.WriteLine($\"Retrieved {setting.Key} = {setting.Value}\");\n+\n+    Console.WriteLine(\"Settings with prefix \\\"app:Settings:\\\":\");\n+    var selector = new SettingSelector\n+    {\n+        KeyFilter = \"app:Settings:*\"\n+    };\n+\n+    foreach (ConfigurationSetting matchingSetting in client.GetConfigurationSettings(selector))\n+    {\n+        string label = matchingSetting.Label is null ? \"(no label)\" : matchingSetting.Label;\n+        Console.WriteLine($\"  {matchingSetting.Key} = {matchingSetting.Value}, label = {label}\");\n+    }\n+\n+    var featureFlag = new FeatureFlagConfigurationSetting(\"BetaFeature\", isEnabled: true);\n+    client.SetConfigurationSetting(featureFlag);\n+    Console.WriteLine(\"Enabled feature flag BetaFeature\");\n+\n+    client.DeleteConfigurationSetting(key);\n+    Console.WriteLine($\"Deleted the unlabeled setting {key}\");\n+\n+    return 0;\n+}\n+catch (RequestFailedException ex)\n+{\n+    Console.Error.WriteLine(\n+        $\"Azure App Configuration request failed. Status: {ex.Status}, ErrorCode: {ex.ErrorCode ?? \"unknown\"}\");\n+    Console.Error.WriteLine(ex.Message);\n+    return 2;\n+}\n*** Add File: README.md\n+# Azure App Configuration CRUD sample\n+\n+Required NuGet package:\n+\n+```powershell\n+dotnet add package Azure.Data.AppConfiguration\n+```\n+\n+The project already references this package. Set a connection string, then run the sample:\n+\n+```powershell\n+$env:AZURE_APP_CONFIGURATION_CONNECTION_STRING = \"\u003cyour-connection-string\u003e\"\n+dotnet run\n+```\n+\n+The connection string needs permission to read and write configuration settings. The sample creates an\n+unlabeled setting, a `Production`-labeled setting, and an enabled `BetaFeature` feature flag. It deletes the\n+unlabeled setting at the end; the labeled setting and feature flag remain in the store.\n*** End Patch\n"
```

### 🔧 7. powershell

**Input:**

```json
{"command":"dotnet build --nologo --verbosity quiet","description":"Build the App Configuration sample","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `AppConfigurationSample.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll create a local-only .NET console sample with the Azure SDK package reference, placeholder connection-string input, and explicit `RequestFailedException` handling.

The sample will use only `Azure.Data.AppConfiguration` because connection-string authentication does not require `Azure.Identity`; the connection string is read from an environment variable rather than embedded.

Implementation is complete; I’m compiling it now to catch any SDK API or type mismatches without connecting to Azure.

Created a complete .NET 8 console sample:

- `Program.cs` implements all requested Azure App Configuration operations.
- `AppConfigurationSample.csproj` references `Azure.Data.AppConfiguration` 1.11.1.
- `README.md` documents package installation and execution.
- Connection strings are read securely from `AZURE_APP_CONFIGURATION_CONNECTION_STRING`.
- Azure failures are handled with `RequestFailedException`, including status and error code.
- The project builds successfully with no warnings or errors.

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 7 criteria passed. The sample covers the required package, client construction, setting creation, prefix listing, enabled feature flag, deletion, and Azure error handling, but does not get a setting by both key and label.

### Strengths

- Uses an environment variable rather than embedding the connection string.
- Includes both unlabeled and Production-labeled setting creation.
- Uses SettingSelector for prefix filtering and handles RequestFailedException with useful status details.

### Issues

- The Production-labeled setting is created but never retrieved with GetConfigurationSetting(key, productionLabel).

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
hyoka run --prompt-id app-configuration-dp-dotnet-crud --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
