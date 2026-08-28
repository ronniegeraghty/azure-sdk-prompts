# Evaluation Report: app-configuration-dp-dotnet-crud

**Config:** dotnet-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 81.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-dotnet-crud` |
| Config | dotnet-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 6/7 |
| Duration | 81.2s |
| Timestamp | 2026-08-28T16:49:00Z |
| Files Generated | 2 |
| Event Count | 1760 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 51.8s |
| Review | 28.0s |
| **Total** | **81.2s** |

## Configuration

- **name:** dotnet-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=81702 out=1503 |
| Turn Count | 4 |

## Tool Calls

### ✅ 1. glob — 7060ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 6730ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586","pattern":"**/*.csproj"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 6400ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586","pattern":"**/*.cs"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. apply_patch — 3660ms

**Input:**

```json
"*** Begin Patch\n*** Add File: AppConfigurationSample.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.Data.AppConfiguration\" Version=\"1.6.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure;\n+using Azure.Data.AppConfiguration;\n+\n+const string key = \"app:Settings:FontSize\";\n+const string productionLabel = \"Production\";\n+const string featureFlagName = \"BetaFeature\";\n+\n+string? connectionString =\n+    Environment.GetEnvironmentVariable(\"AZURE_APP_CONFIGURATION_CONNECTION_STRING\");\n+\n+if (string.IsNullOrWhiteSpace(connectionString))\n+{\n+    Console.Error.WriteLine(\n+        \"Set AZURE_APP_CONFIGURATION_CONNECTION_STRING to an Azure App Configuration connection string.\");\n+    return 1;\n+}\n+\n+try\n+{\n+    var client = new ConfigurationClient(connectionString);\n+\n+    // Create an unlabeled setting and a separate Production-labeled setting.\n+    client.SetConfigurationSetting(new ConfigurationSetting(key, \"24\"));\n+    client.SetConfigurationSetting(new ConfigurationSetting(key, \"24\", productionLabel));\n+\n+    Response\u003cConfigurationSetting\u003e response = client.GetConfigurationSetting(key);\n+    Console.WriteLine($\"{response.Value.Key} = {response.Value.Value}\");\n+\n+    Console.WriteLine(\"Settings with prefix app:Settings:\");\n+    var selector = new SettingSelector\n+    {\n+        KeyFilter = \"app:Settings:*\"\n+    };\n+\n+    foreach (ConfigurationSetting setting in client.GetConfigurationSettings(selector))\n+    {\n+        string label = setting.Label ?? \"(no label)\";\n+        Console.WriteLine($\"{setting.Key} = {setting.Value} [Label: {label}]\");\n+    }\n+\n+    var featureFlag = new FeatureFlagConfigurationSetting(featureFlagName, isEnabled: true);\n+    client.SetConfigurationSetting(featureFlag);\n+    Console.WriteLine($\"Feature flag {featureFlagName} created and enabled.\");\n+\n+    client.DeleteConfigurationSetting(key);\n+    Console.WriteLine($\"Deleted unlabeled setting {key}.\");\n+\n+    return 0;\n+}\n+catch (RequestFailedException ex)\n+{\n+    Console.Error.WriteLine(\n+        $\"Azure App Configuration request failed. Status: {ex.Status}, \" +\n+        $\"ErrorCode: {ex.ErrorCode ?? \"(none)\"}, Message: {ex.Message}\");\n+    return 1;\n+}\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/AppConfigurationSample.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/AppConfigurationSample.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/AppConfigurationSample.csproj
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
+    <PackageReference Include="Azure.Data.AppConfiguration" Version="1.6.0" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-dotnet-crud-dotnet-azure-skills-baseline-2687697586/Program.cs
@@ -1,0 +1,57 @@
+using Azure;
+using Azure.Data.AppConfiguration;
+
+const string key = "app:Settings:FontSize";
+const string productionLabel = "Production";
+const string featureFlagName = "BetaFeature";
+
+string? connectionString =
+    Environment.GetEnvironmentVariable("AZURE_APP_CONFIGURATION_CONNECTION_STRING");
+
+if (string.IsNullOrWhiteSpace(connectionString))
+{
+    Console.Error.WriteLine(
+        "Set AZURE_APP_CONFIGURATION_CONNECTION_STRING to an Azure App Configuration connection string.");
+    return 1;
+}
+
+try
+{
+    var client = new ConfigurationClient(connectionString);
+
+    // Create an unlabeled setting
... (truncated)
```

### ✅ 5. powershell — 6442ms

**Input:**

```json
{"command":"dotnet build .\\AppConfigurationSample.csproj --nologo --verbosity quiet","description":"Build the Azure App Configuration sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:03.08
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `AppConfigurationSample.csproj`
- `Program.cs`

## Copilot Response

I’ll inspect the workspace, then create a local-ready sample with package references and Azure SDK error handling.

The sample is in place; I’m compiling it now to catch any Azure SDK API or package-version issues.

Created a complete C# sample:

- `Program.cs` implements setting creation, Production label, retrieval, prefix listing, enabled `BetaFeature` flag, deletion, and `RequestFailedException` handling.
- `AppConfigurationSample.csproj` targets .NET 8 and includes:

```xml
<PackageReference Include="Azure.Data.AppConfiguration" Version="1.6.0" />
```

Provide the connection string through an environment variable before running:

```powershell
$env:AZURE_APP_CONFIGURATION_CONNECTION_STRING = "<connection-string>"
dotnet run
```

The program deletes the unlabeled `app:Settings:FontSize` setting; the separately labeled `Production` setting remains.

## Reviewer Notes (LLM-as-Judge)

### Summary

6 of 7 criteria passed. The sample covers the required package, client creation, setting operations, filtered listing, feature flag creation, deletion, and Azure request error handling, but it does not retrieve a setting by both key and label.

### Strengths

- Uses an environment variable rather than embedding the connection string.
- Correctly filters settings by the app:Settings: prefix with SettingSelector.
- Provides explicit RequestFailedException handling and enables the BetaFeature feature flag.

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
hyoka run --prompt-id app-configuration-dp-dotnet-crud --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
