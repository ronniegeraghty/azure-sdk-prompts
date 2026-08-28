# Evaluation Report: app-configuration-dp-java-feature-flags

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 843.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-java-feature-flags` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 18/21 |
| Duration | 843.7s |
| Timestamp | 2026-08-27T14:11:56Z |
| Files Generated | 11 |
| Event Count | 13446 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 562.2s |
| Review | 280.2s |
| **Total** | **843.7s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-appconfiguration-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=346876 out=9651 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 4118ms

**Input:**

```json
{"skill":"azure-appconfiguration-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure App Configuration SDK for Java

Client library for Azure App Configuration, a managed service for centralizing application configurations.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-data-appconfiguration</artifactId>
    <version>1.8.0</version>
</dependency>
```

Or use Azure SDK BOM:

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-sdk-bom</artifactId>
            <version>{bom_version}</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-data-appconfiguration</artifactId>
    </dependency>
</dependencies>
```

## Prerequisites

- Azure App Configuration store
- Connection string or Entra ID credentials

## Environment Variables

```bash
AZURE_APPCONFIG_CONNECTION_STRING=Endpoint=https://<store>.azconfig.io;Id=<id>;Secret=<secret>  # Alternative to Entra ID auth
AZURE_APPCONFIG_ENDPOINT=https://<store>.azconfig.io  # Required for all auth methods
AZURE_TOKEN_CREDENTIALS=prod  # Required only if DefaultAzureCredential is used in production
```

## Client Creation

### With Connection String

```java
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;

ConfigurationClient configClient = new ConfigurationClientBuilder()
    .connectionString(System.getenv("AZURE_APPCONFIG_CONNECTION_STRING"))
    .buildClient();
```

### Async Client

```java
import com.azure.data.appconfiguration.ConfigurationAsyncClient;

ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()
    .connectionString(connectionString)
    .buildAsyncClient();
```

### With Entra ID (Recommended)

```java
import com.azure.core.credential.TokenCredential;

... (truncated)
```

### ✅ 2. view — 186372ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md"}
```

**Output:**

```
File too large to read at once (26.1 KB). Consider using the grep tool to search within the file, or view with view_range to read specific sections.
```

### ✅ 3. glob — 185897ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java","pattern":"**/acceptance-criteria.md"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. glob — 185464ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 5. azure-get_azure_bestpractices (via azure) — 186373ms

**Input:**

```json
{"intent":"Create a local-only Java 17 Maven demo using Azure App Configuration sync and async SDK clients, managed identity, conditional ETag reads, feature flags, and sentinel polling. No Azure resources will be deployed or modified.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 6. rg — 183311ms

**Input:**

```json
{"-n":true,"head_limit":100,"output_mode":"content","paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md","pattern":"^## |Conditional|ifChanged|Feature|Async|ManagedIdentity|ETag"}
```

**Output:**

```
[grep content: 80 matches across 1 file(s) under C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references]

examples.md (80 match(es)):
  5: ## Table of Contents
  11: - [Feature Flags](#feature-flags)
  16: - [Async Operations](#async-operations)
  20: ## Maven Dependency
  60: ## Client Creation
  88: ### Async Client
  91: import com.azure.data.appconfiguration.ConfigurationAsyncClient;
  93: ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()
  96:     .buildAsyncClient();
  99: ## Configuration Settings CRUD
  119: System.out.println("  ETag: " + setting.getETag());
  162: System.out.println("ETag: " + setting.getETag());
  166: ### Conditional Get (If Changed)
  174: // First, get the setting with its ETag
  179:     setting,      // Setting with ETag
  181:     true,         // ifChanged - only fetch if modified
  206: ### Conditional Update (If Unchanged)
  211: // Get current setting with ETag
  217: // Update only if ETag matches (optimistic concurrency)
  219:     current,     // Setting with current ETag
  220:     true,        // ifUnchanged - only update if ETag matches
  244: ### Conditional Delete
  248:     setting,     // Setting with ETag
  258: ## List and Filter Settings
  349: ## Feature Flags
  351: ### Create Feature Flag
  354: import com.azure.data.appconfiguration.models.FeatureFlagConfigurationSetting;
  355: import com.azure.data.appconfiguration.models.FeatureFlagFilter;
  359: FeatureFlagConfigurationSetting simpleFlag = new FeatureFlagConfigurationSetting("dark-mode", true)
  364: // Feature flag with percentage filter
  365: FeatureFlagFilter percentageFilter = new FeatureFlagFilter("Microsoft.Percentage")
  368: FeatureFlagConfigurationSetting betaFlag = new FeatureFlagConfigurationSetting("beta-feature", true)
  372: FeatureFlagConfigurationSetting created = (FeatureFlagConfigurationSetting)
  375: System.out.println("Feature flag cr
... (truncated)
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 183314ms

**Input:**

```json
{"intent":"Get Java Azure App Configuration code-generation best practices for sync/async clients, managed identity, conditional ETag requests, feature flags, and sentinel polling. Local demo only; do not deploy or mutate Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. azure-documentation (via azure) — 183319ms

**Input:**

```json
{"intent":"Find official Azure App Configuration Java SDK documentation for conditional ETag reads, feature flag JSON schema, labels, async clients, managed identity authentication, and sentinel-based refresh.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 9. view — 2879ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md","view_range":[20,198]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
@@ -841,25 +841,25 @@
 ## Environment Variables
 
 ```bash
 AZURE_APPCONFIG_CONNECTION_STRING=Endpoint=https://<store>.azconfig.io;Id=<id>;Secret=<secret>
 AZURE_APPCONFIG_ENDPOINT=https://<store>.azconfig.io
 
 # For DefaultAzureCredential
 AZURE_CLIENT_ID=<service-principal-client-id>
 AZURE_CLIENT_SECRET=<service-principal-secret>
 AZURE_TENANT_ID=<tenant-id>
 ```
 
 ## Best Practices
 
 1. **Use labels** — Separate configurations by environment (Dev, Staging, Production)
 2. **Use snapshots** — Create immutable snapshots for releases
 3. **Feature flags** — Use for gradual rollouts and A/B testing
 4. **Secret references** — Store sensitive values in Key Vault, not App Configuration
 5. **Conditional requests** — Use ETags for optimistic concurrency
 6. **Read-only protection** — Lock critical production settings
 7. **Use Entra ID** — Preferred over connection strings for security
 8. **Async client** — Use for high-throughput scenarios
 

```

### ✅ 10. view — 2528ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md","view_range":[258,451]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
@@ -630,35 +630,35 @@
 ## Async Operations
 
 ### Async List with Reactive Streams
 
 ```java
 ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()
     .credential(new DefaultAzureCredentialBuilder().build())
     .endpoint(endpoint)
     .buildAsyncClient();
 
 asyncClient.listConfigurationSettings(new SettingSelector().setLabelFilter("Production"))
     .subscribe(
         setting -> System.out.println(setting.getKey() + " = " + setting.getValue()),
         error -> System.err.println("Error: " + error.getMessage()),
         () -> System.out.println("Completed listing")
     );
 
 // Keep application running for async operations
 Thread.sleep(5000);
 ```
 
 ### Async Get and Set
 
 ```java
 asyncClient.getConfigurationSetting("app/timeout", "Production")
     .flatMap(setting -> {
         setting.setValue("45000");
         return asyncClient.setConfigurationSetting(setting);
     })
     .subscribe(
         updated -> System.out.println("Updated: " + updated.getValue()),
         error -> System.err.println("Error: " + error.getMessage())
     );
 ```
 

```

### ✅ 11. view — 2191ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md","view_range":[630,664]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
@@ -258,194 +258,194 @@
 ## List and Filter Settings
 
 ### List All Settings
 
 ```java
 import com.azure.data.appconfiguration.models.SettingSelector;
 import com.azure.core.http.rest.PagedIterable;
 
 PagedIterable<ConfigurationSetting> settings = client.listConfigurationSettings(
     new SettingSelector()
 );
 
 System.out.println("=== All Settings ===");
 for (ConfigurationSetting s : settings) {
     System.out.printf("%s [%s] = %s%n", s.getKey(), s.getLabel(), s.getValue());
 }
 ```
 
 ### List by Key Pattern
 
 ```java
 SettingSelector selector = new SettingSelector()
     .setKeyFilter("app/*");  // Wildcard pattern
 
 PagedIterable<ConfigurationSetting> settings = client.listConfigurationSettings(selector);
 
 for (ConfigurationSetting s : settings) {
     System.out.println(s.getKey() + " = " + s.getValue());
 }
 ```
 
 ### List by Label
 
 ```java
 SettingSelector selector = new SettingSelector()
     .setKeyFilter("*")
     .setLabelFilter("Production");
 
 PagedIterable<ConfigurationSetting> settings = client.listConfigurationSettings(selector);
 ```
 
 ### List by Multiple Keys
 
 ```java
 SettingSelector selector = new SettingSelector()
     .setKeyFilter("app/database/*,app/cache/*");  // Comma-separated patterns
 
 PagedIterable<ConfigurationSetting> settings = client.listConfigurationSettings(selector);
 
... (truncated)
```

### ✅ 12. view — 1839ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-appconfiguration-java\\references\\examples.md","view_range":[841,865]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-appconfiguration-java/references/examples.md
@@ -20,179 +20,179 @@
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-data-appconfiguration</artifactId>
     <version>1.8.0</version>
 </dependency>
 
 <!-- For DefaultAzureCredential -->
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.14.2</version>
 </dependency>
 ```
 
 Or use Azure SDK BOM:
 
 ```xml
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>1.2.28</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-data-appconfiguration</artifactId>
     </dependency>
 </dependencies>
 ```
 
 ## Client Creation
 
 ### With Connection String
 
 ```java
 import com.azure.data.appconfiguration.ConfigurationClient;
 import com.azure.data.appconfiguration.ConfigurationClientBuilder;
 
 String connectionString = System.getenv("AZURE_APPCONFIG_CONNECTION_STRING");
 
 ConfigurationClient client = new ConfigurationClientBuilder()
     .connectionString(connectionString)
     .buildClient();
 ```
 
 ### Wi
... (truncated)
```

### ✅ 13. github-mcp-server-search_code (via github-mcp-server) — 2221ms

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"getConfigurationSettingWithResponse(ConfigurationSetting setting repo:Azure/azure-sdk-for-java path:sdk/appconfiguration/azure-data-appconfiguration/src/main/java"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/appconfiguration/azure-data-appconfiguration/src/main/java/com/azure/data/appconfiguration/ConfigurationClient.java","text_matches":[{"fragment":" * <p>The {@link ConfigurationClient#setReadOnly(ConfigurationSetting, boolean)}\n * method can be used to conditionally set a configuration setting to read-only in the Azure App Configuration.</p>\n *","matches":[{"indices":[62,69],"text":"Setting"},{"indices":[140,147],"text":"setting"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/appconfiguration/azure-data-appconfiguration/src/main/java/com/azure/data/appconfiguration/ConfigurationClient.java?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"},{"fragment":"    @ServiceMethod(returns = ReturnType.SINGLE)\n    public Response<ConfigurationSetting> getConfigurationSettingWithResponse(ConfigurationSetting setting,\n        OffsetDateTime acceptDateTime, boolean ifChanged, Context context) {","matches":[{"indices":[90,146],"text":"getConfigurationSettingWithResponse(ConfigurationSetting"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/appconfiguration/azure-data-appconfiguration/src/main/java/com/azure/data/appconfiguration/ConfigurationClient.java?ref=adb5c6c0d3ccc0c2ca17af296e814a24b404de43","property":"content"}]},{"path":"sdk/appconfiguration/azure-data-appconfiguration/src/main/java/com/azure/data/appconfiguration/ConfigurationAsyncClient.java","text_matches":[{"fragment":" * <p>The {@link com.azure.data.appconfiguration.ConfigurationAsyncClient#setReadOnly(ConfigurationSetting, boolean)}\n * method can be used to conditionally clear read-only of the setting in the Azure App Configuration.</p>\n *","matches":[{"indices":[99,106],"text":"Setting"},{"indices":[180,187],"text":"setting"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928948/contents/sdk/appconfiguratio
... (truncated)
```

### ✅ 14. apply_patch — 904ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-app-configuration-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.appconfiguration.version\u003e1.8.0\u003c/azure.appconfiguration.version\u003e\n+        \u003cazure.identity.version\u003e1.14.2\u003c/azure.identity.version\u003e\n+        \u003cjackson.version\u003e2.17.2\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.11.0\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-data-appconfiguration\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.appconfiguration.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cuseModulePath\u003efalse\u003c/useModulePath\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.4.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.appconfig.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/appconfig/ConfigurationReader.java\n+package com.example.appconfig;\n+\n+import java.util.Map;\n+\n+public interface ConfigurationReader {\n+    String getSetting(String key, String label);\n+\n+    Map\u003cString, String\u003e listSettings(String keyPrefix, String label);\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncConfigurationReader.java\n+package com.example.appconfig;\n+\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Map;\n+\n+public interface AsyncConfigurationReader {\n+    Mono\u003cString\u003e getSetting(String key, String label);\n+\n+    Mono\u003cMap\u003cString, String\u003e\u003e listSettings(String keyPrefix, String label);\n+}\n*** Add File: src/main/java/com/example/appconfig/ConfigurationService.java\n+package com.example.appconfig;\n+\n+import com.azure.core.http.rest.Response;\n+import com.azure.core.util.Context;\n+import com.azure.data.appconfiguration.ConfigurationClient;\n+import com.azure.data.appconfiguration.models.ConfigurationSetting;\n+import com.azure.data.appconfiguration.models.SettingSelector;\n+\n+import java.util.LinkedHashMap;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class ConfigurationService implements ConfigurationReader {\n+    private static final String NULL_LABEL_FILTER = \"\\0\";\n+\n+    private final ConfigurationClient client;\n+    private final Map\u003cSettingKey, ConfigurationSetting\u003e settingCache = new ConcurrentHashMap\u003c\u003e();\n+    private final Map\u003cPrefixKey, Map\u003cString, String\u003e\u003e prefixCache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public ConfigurationService(ConfigurationClient client) {\n+        this.client = client;\n+    }\n+\n+    public String getSetting(String key) {\n+        return getSetting(key, null);\n+    }\n+\n+    @Override\n+    public String getSetting(String key, String label) {\n+        SettingKey cacheKey = new SettingKey(key, label);\n+        ConfigurationSetting cached = settingCache.get(cacheKey);\n+        if (cached == null) {\n+            ConfigurationSetting loaded = client.getConfigurationSetting(key, label);\n+            settingCache.put(cacheKey, loaded);\n+            return loaded.getValue();\n+        }\n+\n+        Response\u003cConfigurationSetting\u003e response = client.getConfigurationSettingWithResponse(\n+            cached, null, true, Context.NONE);\n+        if (response.getStatusCode() == 304) {\n+            return cached.getValue();\n+        }\n+\n+        ConfigurationSetting updated = response.getValue();\n+        settingCache.put(cacheKey, updated);\n+        return updated.getValue();\n+    }\n+\n+    public Map\u003cString, String\u003e listSettings(String keyPrefix) {\n+        return listSettings(keyPrefix, null);\n+    }\n+\n+    @Override\n+    public Map\u003cString, String\u003e listSettings(String keyPrefix, String label) {\n+        PrefixKey cacheKey = new PrefixKey(keyPrefix, label);\n+        return prefixCache.computeIfAbsent(cacheKey, this::loadPrefix);\n+    }\n+\n+    public synchronized void refreshAll() {\n+        List\u003cSettingKey\u003e settings = List.copyOf(settingCache.keySet());\n+        List\u003cPrefixKey\u003e prefixes = List.copyOf(prefixCache.keySet());\n+        settingCache.clear();\n+        prefixCache.clear();\n+\n+        settings.forEach(key -\u003e getSetting(key.key(), key.label()));\n+        prefixes.forEach(key -\u003e listSettings(key.prefix(), key.label()));\n+    }\n+\n+    private Map\u003cString, String\u003e loadPrefix(PrefixKey key) {\n+        SettingSelector selector = new SettingSelector()\n+            .setKeyFilter(escapeFilter(key.prefix()) + \"*\")\n+            .setLabelFilter(key.label() == null ? NULL_LABEL_FILTER : key.label());\n+        Map\u003cString, String\u003e values = new LinkedHashMap\u003c\u003e();\n+        client.listConfigurationSettings(selector).forEach(setting -\u003e {\n+            values.put(setting.getKey(), setting.getValue());\n+            settingCache.put(new SettingKey(setting.getKey(), setting.getLabel()), setting);\n+        });\n+        return Map.copyOf(values);\n+    }\n+\n+    private static String escapeFilter(String value) {\n+        return value.replace(\"\\\\\", \"\\\\\\\\\").replace(\",\", \"\\\\,\").replace(\"*\", \"\\\\*\");\n+    }\n+\n+    private record SettingKey(String key, String label) {\n+    }\n+\n+    private record PrefixKey(String prefix, String label) {\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncConfigurationService.java\n+package com.example.appconfig;\n+\n+import com.azure.data.appconfiguration.ConfigurationAsyncClient;\n+import com.azure.data.appconfiguration.models.ConfigurationSetting;\n+import com.azure.data.appconfiguration.models.SettingSelector;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.List;\n+import java.util.Map;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class AsyncConfigurationService implements AsyncConfigurationReader {\n+    private static final String NULL_LABEL_FILTER = \"\\0\";\n+\n+    private final ConfigurationAsyncClient client;\n+    private final Map\u003cSettingKey, ConfigurationSetting\u003e settingCache = new ConcurrentHashMap\u003c\u003e();\n+    private final Map\u003cPrefixKey, Map\u003cString, String\u003e\u003e prefixCache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncConfigurationService(ConfigurationAsyncClient client) {\n+        this.client = client;\n+    }\n+\n+    public Mono\u003cString\u003e getSetting(String key) {\n+        return getSetting(key, null);\n+    }\n+\n+    @Override\n+    public Mono\u003cString\u003e getSetting(String key, String label) {\n+        return Mono.defer(() -\u003e {\n+            SettingKey cacheKey = new SettingKey(key, label);\n+            ConfigurationSetting cached = settingCache.get(cacheKey);\n+            if (cached == null) {\n+                return client.getConfigurationSetting(key, label)\n+                    .doOnNext(setting -\u003e settingCache.put(cacheKey, setting))\n+                    .map(ConfigurationSetting::getValue);\n+            }\n+\n+            return client.getConfigurationSettingWithResponse(cached, null, true)\n+                .map(response -\u003e {\n+                    if (response.getStatusCode() == 304) {\n+                        return cached.getValue();\n+                    }\n+                    ConfigurationSetting updated = response.getValue();\n+                    settingCache.put(cacheKey, updated);\n+                    return updated.getValue();\n+                });\n+        });\n+    }\n+\n+    public Mono\u003cMap\u003cString, String\u003e\u003e listSettings(String keyPrefix) {\n+        return listSettings(keyPrefix, null);\n+    }\n+\n+    @Override\n+    public Mono\u003cMap\u003cString, String\u003e\u003e listSettings(String keyPrefix, String label) {\n+        return Mono.defer(() -\u003e {\n+            PrefixKey cacheKey = new PrefixKey(keyPrefix, label);\n+            Map\u003cString, String\u003e cached = prefixCache.get(cacheKey);\n+            return cached == null ? loadPrefix(cacheKey) : Mono.just(cached);\n+        });\n+    }\n+\n+    public Mono\u003cVoid\u003e refreshAll() {\n+        return Mono.defer(() -\u003e {\n+            List\u003cSettingKey\u003e settings = List.copyOf(settingCache.keySet());\n+            List\u003cPrefixKey\u003e prefixes = List.copyOf(prefixCache.keySet());\n+            settingCache.clear();\n+            prefixCache.clear();\n+\n+            Mono\u003cVoid\u003e reloadSettings = Flux.fromIterable(settings)\n+                .concatMap(key -\u003e getSetting(key.key(), key.label()))\n+                .then();\n+            Mono\u003cVoid\u003e reloadPrefixes = Flux.fromIterable(prefixes)\n+                .concatMap(key -\u003e listSettings(key.prefix(), key.label()))\n+                .then();\n+            return reloadSettings.then(reloadPrefixes);\n+        });\n+    }\n+\n+    private Mono\u003cMap\u003cString, String\u003e\u003e loadPrefix(PrefixKey key) {\n+        SettingSelector selector = new SettingSelector()\n+            .setKeyFilter(escapeFilter(key.prefix()) + \"*\")\n+            .setLabelFilter(key.label() == null ? NULL_LABEL_FILTER : key.label());\n+        return client.listConfigurationSettings(selector)\n+            .doOnNext(setting -\u003e settingCache.put(\n+                new SettingKey(setting.getKey(), setting.getLabel()), setting))\n+            .collectMap(ConfigurationSetting::getKey, ConfigurationSetting::getValue)\n+            .map(Map::copyOf)\n+            .doOnNext(values -\u003e prefixCache.put(key, values));\n+    }\n+\n+    private static String escapeFilter(String value) {\n+        return value.replace(\"\\\\\", \"\\\\\\\\\").replace(\",\", \"\\\\,\").replace(\"*\", \"\\\\*\");\n+    }\n+\n+    private record SettingKey(String key, String label) {\n+    }\n+\n+    private record PrefixKey(String prefix, String label) {\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/FeatureFlagEvaluator.java\n+package com.example.appconfig;\n+\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.JsonNode;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.nio.ByteBuffer;\n+import java.nio.charset.StandardCharsets;\n+import java.security.MessageDigest;\n+import java.security.NoSuchAlgorithmException;\n+\n+public final class FeatureFlagEvaluator {\n+    static final String FEATURE_FLAG_PREFIX = \".appconfig.featureflag/\";\n+    private static final String PERCENTAGE_FILTER = \"Microsoft.Percentage\";\n+    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();\n+\n+    private final ConfigurationReader configuration;\n+\n+    public FeatureFlagEvaluator(ConfigurationReader configuration) {\n+        this.configuration = configuration;\n+    }\n+\n+    public boolean isEnabled(String flagId, String userId) {\n+        return isEnabled(flagId, userId, null);\n+    }\n+\n+    public boolean isEnabled(String flagId, String userId, String label) {\n+        String payload = configuration.getSetting(FEATURE_FLAG_PREFIX + flagId, label);\n+        return evaluate(flagId, userId, payload);\n+    }\n+\n+    static boolean evaluate(String flagId, String userId, String payload) {\n+        JsonNode flag;\n+        try {\n+            flag = OBJECT_MAPPER.readTree(payload);\n+        } catch (JsonProcessingException e) {\n+            throw new IllegalArgumentException(\"Invalid JSON for feature flag '\" + flagId + \"'\", e);\n+        }\n+\n+        if (!flag.path(\"enabled\").asBoolean(false)) {\n+            return false;\n+        }\n+\n+        for (JsonNode filter : flag.path(\"conditions\").path(\"client_filters\")) {\n+            if (PERCENTAGE_FILTER.equals(filter.path(\"name\").asText())) {\n+                double percentage = filter.path(\"parameters\").path(\"Value\").asDouble(0);\n+                return percentage \u003e= 100 || percentage \u003e 0 \u0026\u0026 rolloutBucket(flagId, userId) \u003c percentage;\n+            }\n+        }\n+        return true;\n+    }\n+\n+    static double rolloutBucket(String flagId, String userId) {\n+        if (userId == null || userId.isBlank()) {\n+            throw new IllegalArgumentException(\"userId is required for percentage rollout\");\n+        }\n+        try {\n+            byte[] digest = MessageDigest.getInstance(\"SHA-256\")\n+                .digest((flagId + \":\" + userId).getBytes(StandardCharsets.UTF_8));\n+            long unsignedPrefix = Integer.toUnsignedLong(ByteBuffer.wrap(digest).getInt());\n+            return (unsignedPrefix % 10_000) / 100.0;\n+        } catch (NoSuchAlgorithmException e) {\n+            throw new IllegalStateException(\"SHA-256 is unavailable\", e);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncFeatureFlagEvaluator.java\n+package com.example.appconfig;\n+\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncFeatureFlagEvaluator {\n+    private final AsyncConfigurationReader configuration;\n+\n+    public AsyncFeatureFlagEvaluator(AsyncConfigurationReader configuration) {\n+        this.configuration = configuration;\n+    }\n+\n+    public Mono\u003cBoolean\u003e isEnabled(String flagId, String userId) {\n+        return isEnabled(flagId, userId, null);\n+    }\n+\n+    public Mono\u003cBoolean\u003e isEnabled(String flagId, String userId, String label) {\n+        return configuration.getSetting(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + flagId, label)\n+            .map(payload -\u003e FeatureFlagEvaluator.evaluate(flagId, userId, payload));\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/ConfigurationWatcher.java\n+package com.example.appconfig;\n+\n+import reactor.core.Disposable;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.LinkedHashMap;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.Executors;\n+import java.util.concurrent.ScheduledExecutorService;\n+import java.util.concurrent.TimeUnit;\n+import java.util.function.Consumer;\n+\n+public final class ConfigurationWatcher implements AutoCloseable {\n+    private final Runnable startAction;\n+    private final Runnable closeAction;\n+    private boolean started;\n+\n+    private ConfigurationWatcher(Runnable startAction, Runnable closeAction) {\n+        this.startAction = startAction;\n+        this.closeAction = closeAction;\n+    }\n+\n+    public static ConfigurationWatcher forSync(\n+        ConfigurationService service,\n+        List\u003cString\u003e sentinelKeys,\n+        String label,\n+        Duration pollingInterval,\n+        Runnable onRefresh,\n+        Consumer\u003cThrowable\u003e errorHandler\n+    ) {\n+        validate(sentinelKeys, pollingInterval);\n+        ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(runnable -\u003e {\n+            Thread thread = new Thread(runnable, \"app-configuration-watcher\");\n+            thread.setDaemon(true);\n+            return thread;\n+        });\n+        Map\u003cString, String\u003e previousValues = new LinkedHashMap\u003c\u003e();\n+\n+        Runnable poll = () -\u003e {\n+            try {\n+                boolean changed = false;\n+                for (String key : sentinelKeys) {\n+                    String current = service.getSetting(key, label);\n+                    String previous = previousValues.put(key, current);\n+                    changed |= previous != null \u0026\u0026 !Objects.equals(previous, current);\n+                }\n+                if (changed) {\n+                    service.refreshAll();\n+                    onRefresh.run();\n+                }\n+            } catch (RuntimeException error) {\n+                errorHandler.accept(error);\n+            }\n+        };\n+\n+        return new ConfigurationWatcher(\n+            () -\u003e scheduler.scheduleWithFixedDelay(\n+                poll, 0, pollingInterval.toMillis(), TimeUnit.MILLISECONDS),\n+            scheduler::shutdownNow);\n+    }\n+\n+    public static ConfigurationWatcher forAsync(\n+        AsyncConfigurationService service,\n+        List\u003cString\u003e sentinelKeys,\n+        String label,\n+        Duration pollingInterval,\n+        Runnable onRefresh,\n+        Consumer\u003cThrowable\u003e errorHandler\n+    ) {\n+        validate(sentinelKeys, pollingInterval);\n+        Map\u003cString, String\u003e previousValues = new LinkedHashMap\u003c\u003e();\n+        Disposable[] subscription = new Disposable[1];\n+\n+        Mono\u003cVoid\u003e poll = Flux.fromIterable(sentinelKeys)\n+            .concatMap(key -\u003e service.getSetting(key, label).map(value -\u003e Map.entry(key, value)))\n+            .collectMap(Map.Entry::getKey, Map.Entry::getValue, LinkedHashMap::new)\n+            .flatMap(currentValues -\u003e {\n+                boolean changed = !previousValues.isEmpty() \u0026\u0026 sentinelKeys.stream()\n+                    .anyMatch(key -\u003e !Objects.equals(previousValues.get(key), currentValues.get(key)));\n+                previousValues.clear();\n+                previousValues.putAll(currentValues);\n+                return changed ? service.refreshAll().doOnSuccess(ignored -\u003e onRefresh.run()) : Mono.empty();\n+            });\n+\n+        return new ConfigurationWatcher(\n+            () -\u003e subscription[0] = Flux.interval(Duration.ZERO, pollingInterval)\n+                .concatMap(ignored -\u003e poll)\n+                .subscribe(ignored -\u003e {\n+                }, errorHandler),\n+            () -\u003e {\n+                if (subscription[0] != null) {\n+                    subscription[0].dispose();\n+                }\n+            });\n+    }\n+\n+    public synchronized void start() {\n+        if (started) {\n+            throw new IllegalStateException(\"Watcher has already been started\");\n+        }\n+        started = true;\n+        startAction.run();\n+    }\n+\n+    @Override\n+    public synchronized void close() {\n+        if (started) {\n+            closeAction.run();\n+            started = false;\n+        }\n+    }\n+\n+    private static void validate(List\u003cString\u003e sentinelKeys, Duration pollingInterval) {\n+        if (sentinelKeys.isEmpty()) {\n+            throw new IllegalArgumentException(\"At least one sentinel key is required\");\n+        }\n+        if (pollingInterval.isZero() || pollingInterval.isNegative()) {\n+            throw new IllegalArgumentException(\"Polling interval must be positive\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/Main.java\n+package com.example.appconfig;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.data.appconfiguration.ConfigurationAsyncClient;\n+import com.azure.data.appconfiguration.ConfigurationClient;\n+import com.azure.data.appconfiguration.ConfigurationClientBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class Main {\n+    private static final String ENVIRONMENT_LABEL = \"production\";\n+    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(5);\n+    private static final List\u003cString\u003e SENTINELS = List.of(\"demo:sentinel\");\n+    private static final List\u003cString\u003e SAMPLE_USERS = List.of(\"alice\", \"bob\", \"carol\", \"dave\");\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String endpoint = requireEnvironmentVariable(\"AZURE_APPCONFIG_ENDPOINT\");\n+        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();\n+\n+        runSyncDemo(new ConfigurationClientBuilder()\n+            .endpoint(endpoint)\n+            .credential(credential)\n+            .buildClient());\n+\n+        runAsyncDemo(new ConfigurationClientBuilder()\n+            .endpoint(endpoint)\n+            .credential(credential)\n+            .buildAsyncClient());\n+    }\n+\n+    private static void runSyncDemo(ConfigurationClient client) throws InterruptedException {\n+        System.out.println(\"=== Synchronous demo ===\");\n+        ConfigurationService configuration = new ConfigurationService(client);\n+        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(configuration);\n+\n+        printConfiguration(\n+            configuration.getSetting(\"demo:message\"),\n+            configuration.getSetting(\"demo:message\", ENVIRONMENT_LABEL),\n+            configuration.listSettings(\"demo:\", ENVIRONMENT_LABEL));\n+        SAMPLE_USERS.forEach(user -\u003e System.out.printf(\n+            \"beta-dashboard for %-5s: %s%n\",\n+            user,\n+            flags.isEnabled(\"beta-dashboard\", user, ENVIRONMENT_LABEL)));\n+\n+        try (ConfigurationWatcher watcher = ConfigurationWatcher.forSync(\n+            configuration,\n+            SENTINELS,\n+            ENVIRONMENT_LABEL,\n+            POLLING_INTERVAL,\n+            () -\u003e System.out.println(\"Sync cache refreshed after sentinel change.\"),\n+            error -\u003e System.err.println(\"Sync watcher failed: \" + error.getMessage()))) {\n+            watcher.start();\n+            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());\n+        }\n+    }\n+\n+    private static void runAsyncDemo(ConfigurationAsyncClient client) throws InterruptedException {\n+        System.out.println(\"\\n=== Asynchronous demo ===\");\n+        AsyncConfigurationService configuration = new AsyncConfigurationService(client);\n+        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(configuration);\n+\n+        String unlabeled = configuration.getSetting(\"demo:message\").block();\n+        String labeled = configuration.getSetting(\"demo:message\", ENVIRONMENT_LABEL).block();\n+        Map\u003cString, String\u003e prefixed = configuration.listSettings(\"demo:\", ENVIRONMENT_LABEL).block();\n+        printConfiguration(unlabeled, labeled, prefixed);\n+\n+        flagsForUsers(flags).block();\n+\n+        try (ConfigurationWatcher watcher = ConfigurationWatcher.forAsync(\n+            configuration,\n+            SENTINELS,\n+            ENVIRONMENT_LABEL,\n+            POLLING_INTERVAL,\n+            () -\u003e System.out.println(\"Async cache refreshed after sentinel change.\"),\n+            error -\u003e System.err.println(\"Async watcher failed: \" + error.getMessage()))) {\n+            watcher.start();\n+            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());\n+        }\n+    }\n+\n+    private static reactor.core.publisher.Mono\u003cVoid\u003e flagsForUsers(AsyncFeatureFlagEvaluator flags) {\n+        return reactor.core.publisher.Flux.fromIterable(SAMPLE_USERS)\n+            .concatMap(user -\u003e flags.isEnabled(\"beta-dashboard\", user, ENVIRONMENT_LABEL)\n+                .doOnNext(enabled -\u003e System.out.printf(\n+                    \"beta-dashboard for %-5s: %s%n\", user, enabled)))\n+            .then();\n+    }\n+\n+    private static void printConfiguration(\n+        String unlabeled,\n+        String labeled,\n+        Map\u003cString, String\u003e prefixed\n+    ) {\n+        System.out.println(\"Unlabeled message: \" + unlabeled);\n+        System.out.println(\"Production message: \" + labeled);\n+        System.out.println(\"Production demo settings: \" + prefixed);\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(name + \" must contain the App Configuration endpoint\");\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/test/java/com/example/appconfig/FeatureFlagEvaluatorTest.java\n+package com.example.appconfig;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertFalse;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class FeatureFlagEvaluatorTest {\n+    @Test\n+    void evaluatesSimpleEnabledAndDisabledFlags() {\n+        assertTrue(FeatureFlagEvaluator.evaluate(\"flag\", \"user\", \"\"\"\n+            {\"id\":\"flag\",\"enabled\":true,\"conditions\":{\"client_filters\":[]}}\n+            \"\"\"));\n+        assertFalse(FeatureFlagEvaluator.evaluate(\"flag\", \"user\", \"\"\"\n+            {\"id\":\"flag\",\"enabled\":false,\"conditions\":{\"client_filters\":[]}}\n+            \"\"\"));\n+    }\n+\n+    @Test\n+    void percentageRolloutIsDeterministic() {\n+        String payload = \"\"\"\n+            {\n+              \"id\": \"flag\",\n+              \"enabled\": true,\n+              \"conditions\": {\n+                \"client_filters\": [\n+                  {\"name\":\"Microsoft.Percentage\",\"parameters\":{\"Value\":30}}\n+                ]\n+              }\n+            }\n+            \"\"\";\n+\n+        boolean first = FeatureFlagEvaluator.evaluate(\"flag\", \"alice\", payload);\n+        assertEquals(first, FeatureFlagEvaluator.evaluate(\"flag\", \"alice\", payload));\n+        assertEquals(\n+            FeatureFlagEvaluator.rolloutBucket(\"flag\", \"alice\") \u003c 30,\n+            first);\n+    }\n+\n+    @Test\n+    void rejectsMissingUserForPercentageRollout() {\n+        String payload = \"\"\"\n+            {\n+              \"enabled\": true,\n+              \"conditions\": {\n+                \"client_filters\": [\n+                  {\"name\":\"Microsoft.Percentage\",\"parameters\":{\"Value\":50}}\n+                ]\n+              }\n+            }\n+            \"\"\";\n+        assertThrows(\n+            IllegalArgumentException.class,\n+            () -\u003e FeatureFlagEvaluator.evaluate(\"flag\", null, payload));\n+    }\n+\n+    @Test\n+    void evaluatorUsesAzureFeatureFlagKeyAndLabel() {\n+        ConfigurationReader reader = new ConfigurationReader() {\n+            @Override\n+            public String getSetting(String key, String label) {\n+                assertEquals(\".appconfig.featureflag/beta\", key);\n+                assertEquals(\"staging\", label);\n+                return \"\"\"\n+                    {\"id\":\"beta\",\"enabled\":true,\"conditions\":{\"client_filters\":[]}}\n+                    \"\"\";\n+            }\n+\n+            @Override\n+            public Map\u003cString, String\u003e listSettings(String keyPrefix, String label) {\n+                return Map.of();\n+            }\n+        };\n+\n+        assertTrue(new FeatureFlagEvaluator(reader).isEnabled(\"beta\", \"user-1\", \"staging\"));\n+    }\n+}\n*** Add File: README.md\n+# Azure App Configuration Java demo\n+\n+This Java 17 project demonstrates synchronous and asynchronous configuration reads, labels,\n+prefix queries, ETag-aware caching, feature flags with deterministic percentage rollout, and\n+sentinel-driven cache refresh.\n+\n+## App Configuration data\n+\n+Create these settings in an existing App Configuration store:\n+\n+| Key | Label | Example value |\n+| --- | --- | --- |\n+| `demo:message` | *(no label)* | `Hello from the default configuration` |\n+| `demo:message` | `production` | `Hello from production` |\n+| `demo:sentinel` | `production` | `1` |\n+| `.appconfig.featureflag/beta-dashboard` | `production` | See below |\n+\n+Example feature flag payload:\n+\n+```json\n+{\n+  \"id\": \"beta-dashboard\",\n+  \"description\": \"Gradual dashboard rollout\",\n+  \"enabled\": true,\n+  \"conditions\": {\n+    \"client_filters\": [\n+      {\n+        \"name\": \"Microsoft.Percentage\",\n+        \"parameters\": {\n+          \"Value\": 30\n+        }\n+      }\n+    ]\n+  }\n+}\n+```\n+\n+Assign the managed identity the **App Configuration Data Reader** role, then set the endpoint:\n+\n+```powershell\n+$env:AZURE_APPCONFIG_ENDPOINT = \"https://\u003cstore-name\u003e.azconfig.io\"\n+mvn compile exec:java\n+```\n+\n+The demo runs the synchronous flow first, watches for ten seconds, and then repeats with the\n+asynchronous client. Change `demo:sentinel` while either watcher is running to trigger a full\n+refresh of all entries that the application has cached.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548/pom.xml
@@ -1,0 +1,70 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>azure-app-configuration-demo</artifactId>
+    <version>1.0.0</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.appconfiguration.version>1.8.0</azure.appconfiguration.version>
+        <azure.identity.version>1.14.2</azure.identity.version>
+        <jackson.version>2.17.2</jackson.version>
+        <junit.version>5.11.0</junit.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-data-appconfiguration</artifactId>
+            <version>${azure.appconfiguration.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.fasterxml.jackson.core</groupId>
+            <artifactId>jackson-databind</artifactId>
+            <version>${
... (truncated)
```

### ✅ 15. powershell — 14217ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Compile project and run Maven tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example:azure-app-configuration-demo >--------------
[INFO] Building azure-app-configuration-demo 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-app-configuration-demo ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548\src\main\resources
[INFO] 
[INFO] --- compiler:3.13.0:compile (default-compile) @ azure-app-configuration-demo ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 8 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-app-configuration-demo ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-app-configuration-dp-java-feature-flags-java-azure-skills-azure-skill-mcp-microsoft-skill-519534548\src\test\resources
[INFO] 
[INFO] --- compiler:3.13.0:testCompile (default-testCompile) @ azure-app-configuration-demo ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.0:test (default-test) @ azure-app-configuration-demo ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvide
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\appconfig\AsyncConfigurationReader.java`
- `src\main\java\com\example\appconfig\AsyncConfigurationService.java`
- `src\main\java\com\example\appconfig\AsyncFeatureFlagEvaluator.java`
- `src\main\java\com\example\appconfig\ConfigurationReader.java`
- `src\main\java\com\example\appconfig\ConfigurationService.java`
- `src\main\java\com\example\appconfig\ConfigurationWatcher.java`
- `src\main\java\com\example\appconfig\FeatureFlagEvaluator.java`
- `src\main\java\com\example\appconfig\Main.java`
- `src\test\java\com\example\appconfig\FeatureFlagEvaluatorTest.java`

## Copilot Response

I’m grounding the implementation in the App Configuration SDK patterns, then I’ll scaffold and validate the local Maven project.

The design will use ETag-based conditional reads for normal polling and sentinel-triggered cache rehydration for coordinated refreshes; feature rollout decisions use a stable SHA-256 bucket.

Created a complete Java 17 Maven project with:

- Synchronous and Reactor-based asynchronous configuration services
- Label-aware single-key and prefix queries
- ETag conditional reads and cached prefix results
- JSON feature-flag evaluation with deterministic SHA-256 percentage rollout
- Sync/async sentinel watchers that refresh all cached configuration
- Managed-identity demo flows in `Main`
- Azure SDK dependencies, unit tests, and setup instructions in `README.md`

Run with:

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
mvn compile exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Correctly scopes both ConfigurationWatcher instances with try-with-resources; no unclosed AutoCloseable Azure SDK clients are present.

## Grader Results

- feature-flags.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Retrieves settings with a specific label parameter using `SettingSelector`: Pass
      - Lists settings filtered by key prefix using `setKeyFilter()`: Pass
      - Implements conditional reads with `matchConditions` / `setIfNoneMatch()` using the setting's ETag: Pass
      - Handles 304 Not Modified (setting unchanged since last read): Pass
      - Uses `.appconfig.featureflag/` prefix for feature flag keys: Pass
      - Parses the JSON payload in feature flag setting values: Pass
      - Implements deterministic percentage rollout (consistent hash, not `Math.random()`): Pass
      - Implements sentinel key watching with configurable polling interval: Pass
      - Detects sentinel value change via ETag or value comparison and triggers full refresh: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Pass
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Fail
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Fail
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-java-feature-flags --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
