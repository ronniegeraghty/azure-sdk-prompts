# Evaluation Report: key-vault-dp-java-secret-config

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 715.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-secret-config` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 21/22 |
| Duration | 715.2s |
| Timestamp | 2026-08-27T17:38:09Z |
| Files Generated | 11 |
| Event Count | 13640 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 505.3s |
| Review | 208.8s |
| **Total** | **715.2s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=263790 out=10177 |
| Turn Count | 6 |

## Tool Calls

### ❌ 1. azure-get_azure_bestpractices (via azure) — 187518ms

**Input:**

```json
{"intent":"Get current best practices for a Java 17 application using Azure Key Vault Secrets with managed identity, synchronous and asynchronous clients, caching, expiry inspection, and safe secret rotation. Code generation only; do not modify live Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 2. azure-documentation (via azure) — 187521ms

**Input:**

```json
{"intent":"Find official documentation for Azure Key Vault Secrets Java SDK: SecretClient and SecretAsyncClient getSecret with versions, ResourceNotFound handling, expiry properties, beginDeleteSecret polling completion, purge/recovery behavior, and DefaultAzureCredential managed identity.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. glob — 187512ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. web_search (via github-mcp-server) — 185766ms

**Input:**

```json
{"query":"site:learn.microsoft.com Java Azure Key Vault SecretClient beginDeleteSecret purgeDeletedSecret getDeletedSecret setSecret KeyVaultSecret expiry ResourceNotFoundException managed identity Maven current SDK"}
```

**Output:**

```
{"type":"output_text","text":{"value":"As of August 2026, the latest available versions on Maven Central for the requested Azure libraries are:\n\n- **com.azure:azure-identity:** `1.19.0-beta.2`\n- **com.azure:azure-security-keyvault-secrets:** `4.11.2`\n\nThese versions were confirmed from the Maven Central repository and reflect updates up to at least mid-2026【3:2†source】【3:3†source】. If you need a direct Maven dependency snippet or links for these versions, let me know!","annotations":[{"text":"","start_index":320,"end_index":332,"url_citation":{"title":"Maven Central: com.azure:azure-identity","url":"https://central.sonatype.com/artifact/com.azure/azure-identity"}},{"text":"","start_index":332,"end_index":344,"url_citation":{"title":"Maven Central: com.azure:azure-security-keyvault-secrets","url":"https://central.sonatype.com/artifact/com.azure/azure-security-keyvault-secrets"}}]},"bing_searches":[{"text":"Maven Central latest versions August 2026 com.azure azure-security-keyvault-secrets azure-identity","url":"https://www.bing.com/search?q=Maven+Central+latest+versions+August+2026+com.azure+azure-security-keyvault-secrets+azure-identity"}],"annotations":null}
```

### ✅ 5. web_search (via github-mcp-server) — 185415ms

**Input:**

```json
{"query":"Maven Central latest com.azure azure-security-keyvault-secrets azure-identity versions August 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here’s a reference summary answering your query using the most current Microsoft documentation. This covers code samples, Maven dependencies, methods for secret management, managed identity, error handling, and expiry settings—focused on Java and the current Azure Key Vault SDK.\n\n---\n\n### 1. Maven Dependency (Current SDK)\n\nYou should declare these dependencies (note the version may update, check the [official docs](https://learn.microsoft.com/en-us/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)):\n\n**With BOM (recommended):**\n```xml\n<dependencyManagement>\n  <dependencies>\n    <dependency>\n      <groupId>com.azure</groupId>\n      <artifactId>azure-sdk-bom</artifactId>\n      <version>1.2.18</version> <!-- check the latest BOM version -->\n      <type>pom</type>\n      <scope>import</scope>\n    </dependency>\n  </dependencies>\n</dependencyManagement>\n<dependencies>\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-security-keyvault-secrets</artifactId>\n  </dependency>\n</dependencies>\n```\n**Direct version (if you don’t use BOM):**\n```xml\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-security-keyvault-secrets</artifactId>\n  <version>4.11.0</version> <!-- This is one of the recent GA versions -->\n</dependency>\n```\n【3:1†source】\n\n---\n\n### 2. Setting up SecretClient with Managed Identity\n\n```java\nSecretClient secretClient = new SecretClientBuilder()\n    .vaultUrl(\"<your-key-vault-url>\")\n    .credential(new DefaultAzureCredentialBuilder().build()) // supports Managed Identity\n    .buildClient();\n```\nDefaultAzureCredential works for local dev (env vars, Azure CLI) and for managed identity in production【3:0†source】【3:1†source】.\n\n---\n\n### 3. Common Operations\n\n#### Set a Secret (with expiry)\n```java\nKeyVaultSecret secret = new KeyVaultSecret(\"secretName\", \"secretValue\");\nOffsetDateTime ex
... (truncated)
```

### ❌ 6. azure-documentation (via azure) — 185768ms

**Input:**

```json
{"command":"search","intent":"Official Java API examples for SecretClient and SecretAsyncClient: versioned getSecret, expiry properties, asynchronous delete poller completion, purgeDeletedSecret, and ManagedIdentityCredentialBuilder. Return exact method signatures and recommended Maven artifacts.","learn":false}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 7. web_fetch — 3948ms

**Input:**

```json
{"max_length":10000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/com.azure.security.keyvault.secrets.secretasyncclient?view=azure-java-stable:
---
layout: Reference
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: SecretAsyncClient Class | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/com.azure.security.keyvault.secrets.secretasyncclient?view=azure-java-stable
config_moniker_range: azure-java-stable
uid: com.azure.security.keyvault.secrets.SecretAsyncClient
inheritedClassMethods:
- System.Collections.Generic.Dictionary`2[System.String,System.Object]
desc: >
  <p>The SecretAsyncClient provides asynchronous methods to manage <a class=no-loc href='com.azure.security.keyvault.secrets.models.keyvaultsecret'>KeyVaultSecret</a> in the Azure Key Vault. The client supports creating, retrieving, updating, deleting, purging, backing up, restoring, and listing the <a class=no-loc href='com.azure.security.keyvault.secrets.models.keyvaultsecret'>KeyVaultSecret</a>. The client also supports listing <a class=no-loc href='com.azure.security.keyvault.secrets.models.deletedsecret'>DeletedSecret</a> for a soft-delete enabled key vault.</p>

  <h2 id="getting-started">Getting Started</h2>

  <p>In order to interact with the Azure Key Vault service, you will need to create an instance of the <a class=no-loc href='com.azure.security.keyvault.secrets.secretasyncclient'>SecretAsyncClient</a> class, a vault url and a credential object.</p>

  <p>The examples shown in this document use a credential object named DefaultAzureCredential for authentication, which is appropriate for most scenarios, including local development and production environments. Additionally, we recommend using a <a href="/azure/active-directory/managed-identities-azure-resources/"> managed identity</a> for authentication in production environments. You can find more information on different ways of authenticating and their corresponding credential types in the
... (truncated)
```

### ✅ 8. web_fetch — 3585ms

**Input:**

```json
{"max_length":10000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-sdk-bom</artifactId>
  <versioning>
    <latest>1.3.8</latest>
    <release>1.3.8</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.2.6</version>
      <version>1.2.7</version>
      <version>1.2.8</version>
      <version>1.2.9</version>
      <version>1.2.10</version>
      <version>1.2.11</version>
      <version>1.2.12</version>
      <version>1.2.13</version>
      <version>1.2.14</version>
      <version>1.2.15</version>
      <version>1.2.16</version>
      <version>1.2.17</version>
      <version>1.2.18</version>
      <version>1.2.19</version>
      <version>1.2.20</version>
      <version>1.2.21</version>
      <version>1.2.22</version>
      <version>1.2.23</version>
      <version>1.2.24</version>
      <version>1.2.25</version>
      <version>1.2.26</version>
      <version>1.2.27</version>
      <version>1.2.28</version>
      <version>1.2.29</version>
      <version>1.2.30</version>
      <version>1.2.31</version>
      <version>1.2.32</version>
      <version>1.2.33</version>
      <version>1.2.34</version>
      <version>1.2.35</version>
      <version>1.2.36</version>
      <version>1.2.37</version>
      <version>1.2.38</version>
      <version>1.3.0</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version
... (truncated)
```

### ✅ 9. web_fetch — 3246ms

**Input:**

```json
{"max_length":10000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <versioning>
    <latest>1.19.0-beta.2</latest>
    <release>1.19.0-beta.2</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.0-preview.1</version>
      <version>1.0.0-preview.2</version>
      <version>1.0.0-preview.3</version>
      <version>1.0.0-preview.4</version>
      <version>1.0.0-preview.5</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.0.7</version>
      <version>1.0.8</version>
      <version>1.0.9</version>
      <version>1.1.0-beta.1</version>
      <version>1.1.0-beta.2</version>
      <version>1.1.0-beta.3</version>
      <version>1.1.0-beta.4</version>
      <version>1.1.0-beta.5</version>
      <version>1.1.0-beta.6</version>
      <version>1.1.0-beta.7</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.2.0-beta.1</version>
      <version>1.2.0-beta.2</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.3.0-beta.1</version>
      <version>1.3.0-beta.2</version>
      <version>1.3.0</version>
      <version>1.3.1</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version>1.3.4</version>
      <version>1.3.5</version>
      <version>1.3.6</version>
      <version>1.3.7</version>
      <version>1.4.0-beta.1</version>
      <version>1.4.0</version>
      <version>1.4.1</version>
      <version>
... (truncated)
```

### ✅ 10. web_fetch — 2897ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.security.keyvault.secrets.secretasyncclient?view=azure-java-stable"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-secrets/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-security-keyvault-secrets</artifactId>
  <versioning>
    <latest>4.11.2</latest>
    <release>4.11.2</release>
    <versions>
      <version>4.0.0</version>
      <version>4.0.1</version>
      <version>4.1.0</version>
      <version>4.1.1</version>
      <version>4.1.2</version>
      <version>4.1.3</version>
      <version>4.1.4</version>
      <version>4.1.5</version>
      <version>4.2.0-beta.1</version>
      <version>4.2.0-beta.2</version>
      <version>4.2.0-beta.3</version>
      <version>4.2.0-beta.4</version>
      <version>4.2.0</version>
      <version>4.2.1</version>
      <version>4.2.2</version>
      <version>4.2.3</version>
      <version>4.2.4</version>
      <version>4.2.5</version>
      <version>4.2.6</version>
      <version>4.2.7</version>
      <version>4.2.8</version>
      <version>4.3.0-beta.1</version>
      <version>4.3.0-beta.2</version>
      <version>4.3.0-beta.3</version>
      <version>4.3.0-beta.4</version>
      <version>4.3.0-beta.5</version>
      <version>4.3.0-beta.6</version>
      <version>4.3.0</version>
      <version>4.3.1</version>
      <version>4.3.2</version>
      <version>4.3.3</version>
      <version>4.3.4</version>
      <version>4.3.5</version>
      <version>4.3.6</version>
      <version>4.3.7</version>
      <version>4.3.8</version>
      <version>4.4.0-beta.1</version>
      <version>4.4.0-beta.2</version>
      <version>4.4.0-beta.3</version>
      <version>4.4.0-beta.4</version>
      <version>4.4.0-beta.5</version>
      <version>4.4.0</version>
      <version>4.4.1</version>
      <version>4.4.2</version>
      <version>4.4.3</version>
      <version>4.4.4</version>
      <version>4.4.5</version>
      <version>4.4.6</ve
... (truncated)
```

### ✅ 11. apply_patch — 708ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-config-provider\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure Key Vault configuration provider\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.keyvaultconfig.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/keyvaultconfig/ConfigSecret.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public record ConfigSecret(\n+        String name,\n+        String value,\n+        String version,\n+        OffsetDateTime expiresOn,\n+        boolean defaultValue) {\n+\n+    public ConfigSecret {\n+        Objects.requireNonNull(name, \"name\");\n+        Objects.requireNonNull(value, \"value\");\n+    }\n+\n+    public boolean expiresWithin(Duration warningWindow, Clock clock) {\n+        Objects.requireNonNull(warningWindow, \"warningWindow\");\n+        Objects.requireNonNull(clock, \"clock\");\n+        if (warningWindow.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return expiresOn != null\n+                \u0026\u0026 !expiresOn.isAfter(OffsetDateTime.now(clock).plus(warningWindow));\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/KeyVaultSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+import java.util.Objects;\n+\n+public final class KeyVaultSecretProvider {\n+    private final SecretClient client;\n+\n+    public KeyVaultSecretProvider(SecretClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public ConfigSecret getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    public ConfigSecret getSecret(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+\n+        try {\n+            KeyVaultSecret secret = version == null || version.isBlank()\n+                    ? client.getSecret(name)\n+                    : client.getSecret(name, version);\n+            return new ConfigSecret(\n+                    secret.getName(),\n+                    secret.getValue(),\n+                    secret.getProperties().getVersion(),\n+                    secret.getProperties().getExpiresOn(),\n+                    false);\n+        } catch (ResourceNotFoundException exception) {\n+            return new ConfigSecret(name, defaultValue, version, null, true);\n+        }\n+    }\n+\n+    private static void requireText(String value, String field) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(field + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncKeyVaultSecretProvider.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public final class AsyncKeyVaultSecretProvider {\n+    private final SecretAsyncClient client;\n+\n+    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public Mono\u003cConfigSecret\u003e getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    public Mono\u003cConfigSecret\u003e getSecret(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+\n+        Mono\u003cKeyVaultSecret\u003e request = version == null || version.isBlank()\n+                ? client.getSecret(name)\n+                : client.getSecret(name, version);\n+\n+        return request\n+                .map(secret -\u003e new ConfigSecret(\n+                        secret.getName(),\n+                        secret.getValue(),\n+                        secret.getProperties().getVersion(),\n+                        secret.getProperties().getExpiresOn(),\n+                        false))\n+                .onErrorResume(\n+                        ResourceNotFoundException.class,\n+                        exception -\u003e Mono.just(new ConfigSecret(name, defaultValue, version, null, true)));\n+    }\n+\n+    private static void requireText(String value, String field) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(field + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretCache.java\n+package com.example.keyvaultconfig;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class SecretCache {\n+    private final KeyVaultSecretProvider provider;\n+    private final Map\u003cString, String\u003e requiredDefaults;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final ConcurrentHashMap\u003cString, ConfigSecret\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public SecretCache(\n+            KeyVaultSecretProvider provider,\n+            Map\u003cString, String\u003e requiredDefaults,\n+            Duration warningWindow) {\n+        this(provider, requiredDefaults, warningWindow, Clock.systemUTC());\n+    }\n+\n+    SecretCache(\n+            KeyVaultSecretProvider provider,\n+            Map\u003cString, String\u003e requiredDefaults,\n+            Duration warningWindow,\n+            Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.requiredDefaults = Map.copyOf(Objects.requireNonNull(requiredDefaults, \"requiredDefaults\"));\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public void loadRequired() {\n+        requiredDefaults.forEach((name, defaultValue) -\u003e\n+                cache.put(name, provider.getSecret(name, defaultValue)));\n+    }\n+\n+    public ConfigSecret get(String name) {\n+        String defaultValue = defaultFor(name);\n+        return cache.compute(name, (key, current) -\u003e current == null || isNearExpiry(current)\n+                ? provider.getSecret(key, defaultValue)\n+                : current);\n+    }\n+\n+    public ConfigSecret refresh(String name) {\n+        ConfigSecret refreshed = provider.getSecret(name, defaultFor(name));\n+        cache.put(name, refreshed);\n+        return refreshed;\n+    }\n+\n+    public List\u003cConfigSecret\u003e expiringSecrets() {\n+        return cache.values().stream().filter(this::isNearExpiry).toList();\n+    }\n+\n+    public List\u003cConfigSecret\u003e refreshExpiringSecrets() {\n+        List\u003cConfigSecret\u003e refreshed = new ArrayList\u003c\u003e();\n+        for (ConfigSecret secret : expiringSecrets()) {\n+            refreshed.add(refresh(secret.name()));\n+        }\n+        return List.copyOf(refreshed);\n+    }\n+\n+    private String defaultFor(String name) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        return requiredDefaults.getOrDefault(name, \"\");\n+    }\n+\n+    private boolean isNearExpiry(ConfigSecret secret) {\n+        return secret.expiresWithin(warningWindow, clock);\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretCache.java\n+package com.example.keyvaultconfig;\n+\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class AsyncSecretCache {\n+    private final AsyncKeyVaultSecretProvider provider;\n+    private final Map\u003cString, String\u003e requiredDefaults;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final ConcurrentHashMap\u003cString, ConfigSecret\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncSecretCache(\n+            AsyncKeyVaultSecretProvider provider,\n+            Map\u003cString, String\u003e requiredDefaults,\n+            Duration warningWindow) {\n+        this(provider, requiredDefaults, warningWindow, Clock.systemUTC());\n+    }\n+\n+    AsyncSecretCache(\n+            AsyncKeyVaultSecretProvider provider,\n+            Map\u003cString, String\u003e requiredDefaults,\n+            Duration warningWindow,\n+            Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.requiredDefaults = Map.copyOf(Objects.requireNonNull(requiredDefaults, \"requiredDefaults\"));\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public Mono\u003cVoid\u003e loadRequired() {\n+        return Flux.fromIterable(requiredDefaults.entrySet())\n+                .flatMap(entry -\u003e fetchAndCache(entry.getKey(), entry.getValue()))\n+                .then();\n+    }\n+\n+    public Mono\u003cConfigSecret\u003e get(String name) {\n+        String defaultValue = defaultFor(name);\n+        ConfigSecret current = cache.get(name);\n+        return current == null || isNearExpiry(current)\n+                ? fetchAndCache(name, defaultValue)\n+                : Mono.just(current);\n+    }\n+\n+    public Mono\u003cConfigSecret\u003e refresh(String name) {\n+        return fetchAndCache(name, defaultFor(name));\n+    }\n+\n+    public List\u003cConfigSecret\u003e expiringSecrets() {\n+        return cache.values().stream().filter(this::isNearExpiry).toList();\n+    }\n+\n+    public Flux\u003cConfigSecret\u003e refreshExpiringSecrets() {\n+        return Flux.fromIterable(expiringSecrets())\n+                .flatMap(secret -\u003e refresh(secret.name()));\n+    }\n+\n+    private Mono\u003cConfigSecret\u003e fetchAndCache(String name, String defaultValue) {\n+        return provider.getSecret(name, defaultValue)\n+                .doOnNext(secret -\u003e cache.put(name, secret));\n+    }\n+\n+    private String defaultFor(String name) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        return requiredDefaults.getOrDefault(name, \"\");\n+    }\n+\n+    private boolean isNearExpiry(ConfigSecret secret) {\n+        return secret.expiresWithin(warningWindow, clock);\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/KeyVaultClientFactory.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+\n+import java.net.URI;\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class KeyVaultClientFactory {\n+    public static final String VAULT_URL_ENV = \"KEY_VAULT_URL\";\n+    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = \"AZURE_CLIENT_ID\";\n+\n+    private KeyVaultClientFactory() {\n+    }\n+\n+    public static Clients fromEnvironment() {\n+        return fromEnvironment(System.getenv());\n+    }\n+\n+    static Clients fromEnvironment(Map\u003cString, String\u003e environment) {\n+        Objects.requireNonNull(environment, \"environment\");\n+        String vaultUrl = requireHttpsVaultUrl(environment.get(VAULT_URL_ENV));\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        String managedIdentityClientId = environment.get(MANAGED_IDENTITY_CLIENT_ID_ENV);\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        SecretClientBuilder clientBuilder = new SecretClientBuilder()\n+                .vaultUrl(vaultUrl)\n+                .credential(credential);\n+        return new Clients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());\n+    }\n+\n+    private static String requireHttpsVaultUrl(String value) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" must be set\");\n+        }\n+        URI uri;\n+        try {\n+            uri = URI.create(value);\n+        } catch (IllegalArgumentException exception) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" must be a valid URI\", exception);\n+        }\n+        if (!\"https\".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" must be an HTTPS URL\");\n+        }\n+        return value;\n+    }\n+\n+    public record Clients(SecretClient syncClient, SecretAsyncClient asyncClient) {\n+        public Clients {\n+            Objects.requireNonNull(syncClient, \"syncClient\");\n+            Objects.requireNonNull(asyncClient, \"asyncClient\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/SecretRotationHelper.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.util.polling.SyncPoller;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.DeletedSecret;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public final class SecretRotationHelper {\n+    private final SecretClient client;\n+    private final Duration timeout;\n+    private final Duration pollInterval;\n+\n+    public SecretRotationHelper(SecretClient client, Duration timeout, Duration pollInterval) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.timeout = requirePositive(timeout, \"timeout\");\n+        this.pollInterval = requirePositive(pollInterval, \"pollInterval\");\n+    }\n+\n+    public KeyVaultSecret rotate(\n+            String name,\n+            String newValue,\n+            OffsetDateTime newExpiresOn) {\n+        requireRotationArguments(name, newValue, newExpiresOn);\n+        long deadlineNanos = deadlineNanos();\n+\n+        SyncPoller\u003cDeletedSecret, Void\u003e deletePoller = client.beginDeleteSecret(name);\n+        deletePoller.waitForCompletion(remaining(deadlineNanos));\n+\n+        // Soft-deleted names cannot be recreated until the deleted record is purged.\n+        client.purgeDeletedSecret(name);\n+        waitUntilPurged(name, deadlineNanos);\n+\n+        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);\n+        replacement.getProperties().setExpiresOn(newExpiresOn);\n+        return createAfterPurge(replacement, deadlineNanos);\n+    }\n+\n+    private void waitUntilPurged(String name, long deadlineNanos) {\n+        while (true) {\n+            try {\n+                client.getDeletedSecret(name);\n+            } catch (ResourceNotFoundException exception) {\n+                return;\n+            }\n+            sleepOrThrow(deadlineNanos);\n+        }\n+    }\n+\n+    private KeyVaultSecret createAfterPurge(KeyVaultSecret replacement, long deadlineNanos) {\n+        while (true) {\n+            try {\n+                return client.setSecret(replacement);\n+            } catch (HttpResponseException exception) {\n+                if (exception.getResponse() == null\n+                        || exception.getResponse().getStatusCode() != 409) {\n+                    throw exception;\n+                }\n+                sleepOrThrow(deadlineNanos);\n+            }\n+        }\n+    }\n+\n+    private void sleepOrThrow(long deadlineNanos) {\n+        Duration remaining = remaining(deadlineNanos);\n+        Duration sleep = remaining.compareTo(pollInterval) \u003c 0 ? remaining : pollInterval;\n+        try {\n+            Thread.sleep(sleep.toMillis());\n+        } catch (InterruptedException exception) {\n+            Thread.currentThread().interrupt();\n+            throw new IllegalStateException(\"Interrupted while waiting for Key Vault rotation\", exception);\n+        }\n+    }\n+\n+    private long deadlineNanos() {\n+        return System.nanoTime() + timeout.toNanos();\n+    }\n+\n+    private static Duration remaining(long deadlineNanos) {\n+        long nanos = deadlineNanos - System.nanoTime();\n+        if (nanos \u003c= 0) {\n+            throw new IllegalStateException(\"Timed out waiting for Key Vault secret rotation\");\n+        }\n+        return Duration.ofNanos(nanos);\n+    }\n+\n+    static void requireRotationArguments(String name, String newValue, OffsetDateTime newExpiresOn) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        Objects.requireNonNull(newValue, \"newValue\");\n+        Objects.requireNonNull(newExpiresOn, \"newExpiresOn\");\n+    }\n+\n+    static Duration requirePositive(Duration duration, String field) {\n+        Objects.requireNonNull(duration, field);\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(field + \" must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/AsyncSecretRotationHelper.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+import reactor.util.retry.Retry;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public final class AsyncSecretRotationHelper {\n+    private final SecretAsyncClient client;\n+    private final Duration timeout;\n+    private final Duration pollInterval;\n+\n+    public AsyncSecretRotationHelper(\n+            SecretAsyncClient client,\n+            Duration timeout,\n+            Duration pollInterval) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.timeout = SecretRotationHelper.requirePositive(timeout, \"timeout\");\n+        this.pollInterval = SecretRotationHelper.requirePositive(pollInterval, \"pollInterval\");\n+    }\n+\n+    public Mono\u003cKeyVaultSecret\u003e rotate(\n+            String name,\n+            String newValue,\n+            OffsetDateTime newExpiresOn) {\n+        SecretRotationHelper.requireRotationArguments(name, newValue, newExpiresOn);\n+\n+        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue);\n+        replacement.getProperties().setExpiresOn(newExpiresOn);\n+\n+        return client.beginDeleteSecret(name)\n+                .last()\n+                .then(client.purgeDeletedSecret(name))\n+                .then(waitUntilPurged(name))\n+                .then(createAfterPurge(replacement))\n+                .timeout(timeout);\n+    }\n+\n+    private Mono\u003cVoid\u003e waitUntilPurged(String name) {\n+        return Flux.interval(Duration.ZERO, pollInterval)\n+                .concatMap(ignored -\u003e client.getDeletedSecret(name)\n+                        .map(secret -\u003e false)\n+                        .onErrorResume(\n+                                ResourceNotFoundException.class,\n+                                exception -\u003e Mono.just(true)))\n+                .filter(Boolean::booleanValue)\n+                .next()\n+                .then();\n+    }\n+\n+    private Mono\u003cKeyVaultSecret\u003e createAfterPurge(KeyVaultSecret replacement) {\n+        long retries = Math.max(1, timeout.dividedBy(pollInterval));\n+        return client.setSecret(replacement)\n+                .retryWhen(Retry.fixedDelay(retries, pollInterval)\n+                        .filter(this::isNameConflict));\n+    }\n+\n+    private boolean isNameConflict(Throwable throwable) {\n+        return throwable instanceof HttpResponseException exception\n+                \u0026\u0026 exception.getResponse() != null\n+                \u0026\u0026 exception.getResponse().getStatusCode() == 409;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvaultconfig/Main.java\n+package com.example.keyvaultconfig;\n+\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.UUID;\n+\n+public final class Main {\n+    private static final Duration WARNING_WINDOW = Duration.ofDays(7);\n+    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(5);\n+    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);\n+    private static final String ROTATION_SECRET = \"demo-rotating-secret\";\n+    private static final Map\u003cString, String\u003e REQUIRED_CONFIG = Map.of(\n+            \"database-connection-string\", \"not-configured\",\n+            \"external-api-key\", \"not-configured\",\n+            \"feature-flags\", \"{}\");\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        KeyVaultClientFactory.Clients clients = KeyVaultClientFactory.fromEnvironment();\n+        runSyncDemo(clients);\n+        runAsyncDemo(clients).block();\n+    }\n+\n+    private static void runSyncDemo(KeyVaultClientFactory.Clients clients) {\n+        System.out.println(\"=== Synchronous provider ===\");\n+        KeyVaultSecretProvider provider = new KeyVaultSecretProvider(clients.syncClient());\n+        SecretCache cache = new SecretCache(provider, REQUIRED_CONFIG, WARNING_WINDOW);\n+\n+        cache.loadRequired();\n+        REQUIRED_CONFIG.keySet().forEach(name -\u003e printRead(cache.get(name)));\n+\n+        printRead(cache.refresh(\"external-api-key\"));\n+        warnAbout(cache.expiringSecrets());\n+        cache.refreshExpiringSecrets();\n+\n+        ConfigSecret versioned = provider.getSecret(\n+                \"external-api-key\",\n+                System.getenv(\"DEMO_SECRET_VERSION\"),\n+                \"not-configured\");\n+        System.out.printf(\"Versioned read: %s (version=%s, default=%s)%n\",\n+                versioned.name(), versioned.version(), versioned.defaultValue());\n+\n+        SecretRotationHelper rotation = new SecretRotationHelper(\n+                clients.syncClient(), ROTATION_TIMEOUT, ROTATION_POLL_INTERVAL);\n+        KeyVaultSecret rotated = rotation.rotate(\n+                ROTATION_SECRET,\n+                \"sync-\" + UUID.randomUUID(),\n+                OffsetDateTime.now(ZoneOffset.UTC).plusDays(30));\n+        System.out.printf(\"Rotated %s to version %s%n\",\n+                rotated.getName(), rotated.getProperties().getVersion());\n+    }\n+\n+    private static Mono\u003cVoid\u003e runAsyncDemo(KeyVaultClientFactory.Clients clients) {\n+        System.out.println(\"=== Asynchronous provider ===\");\n+        AsyncKeyVaultSecretProvider provider = new AsyncKeyVaultSecretProvider(clients.asyncClient());\n+        AsyncSecretCache cache = new AsyncSecretCache(provider, REQUIRED_CONFIG, WARNING_WINDOW);\n+        AsyncSecretRotationHelper rotation = new AsyncSecretRotationHelper(\n+                clients.asyncClient(), ROTATION_TIMEOUT, ROTATION_POLL_INTERVAL);\n+\n+        return cache.loadRequired()\n+                .thenMany(Flux.fromIterable(REQUIRED_CONFIG.keySet())\n+                        .concatMap(cache::get)\n+                        .doOnNext(Main::printRead))\n+                .then(cache.refresh(\"external-api-key\").doOnNext(Main::printRead))\n+                .then(Mono.fromRunnable(() -\u003e warnAbout(cache.expiringSecrets())))\n+                .thenMany(cache.refreshExpiringSecrets())\n+                .then(provider.getSecret(\n+                        \"external-api-key\",\n+                        System.getenv(\"DEMO_SECRET_VERSION\"),\n+                        \"not-configured\"))\n+                .doOnNext(secret -\u003e System.out.printf(\n+                        \"Versioned read: %s (version=%s, default=%s)%n\",\n+                        secret.name(), secret.version(), secret.defaultValue()))\n+                .then(rotation.rotate(\n+                        ROTATION_SECRET,\n+                        \"async-\" + UUID.randomUUID(),\n+                        OffsetDateTime.now(ZoneOffset.UTC).plusDays(30)))\n+                .doOnNext(secret -\u003e System.out.printf(\n+                        \"Rotated %s to version %s%n\",\n+                        secret.getName(), secret.getProperties().getVersion()))\n+                .then();\n+    }\n+\n+    private static void printRead(ConfigSecret secret) {\n+        System.out.printf(\"Cache read: %s (version=%s, default=%s)%n\",\n+                secret.name(), secret.version(), secret.defaultValue());\n+    }\n+\n+    private static void warnAbout(List\u003cConfigSecret\u003e secrets) {\n+        secrets.forEach(secret -\u003e System.out.printf(\n+                \"WARNING: %s expires at %s%n\", secret.name(), secret.expiresOn()));\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+A small Java 17 example with synchronous and Reactor-based asynchronous providers, in-memory caches, expiry-aware refresh, and safe secret rotation.\n+\n+## Configuration\n+\n+Set `KEY_VAULT_URL` to an HTTPS vault URL. The application authenticates only with Azure managed identity. For a user-assigned identity, also set `AZURE_CLIENT_ID`; otherwise the system-assigned identity is used.\n+\n+The identity needs secret read/list permissions for the provider. The rotation demo additionally needs delete and purge permissions.\n+\n+```powershell\n+$env:KEY_VAULT_URL = \"https://your-vault.vault.azure.net\"\n+mvn compile exec:java\n+```\n+\n+`DEMO_SECRET_VERSION` is optional and demonstrates fetching an exact secret version. Secret values are deliberately not printed.\n+\n+## Rotation behavior\n+\n+Key Vault soft-delete retains a deleted secret name, so the helper waits for deletion, purges the deleted record, waits until that record is no longer visible, and only then creates the replacement. Rotation therefore requires purge permission and a vault without purge protection. When purge protection is enabled, delete-and-recreate rotation is intentionally rejected by Key Vault; use normal secret versioning instead.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-2691076187/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-2691076187/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-key-vault-dp-java-secret-config-java-azure-skills-azure-skill-mcp-2691076187/pom.xml
@@ -1,0 +1,65 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>key-vault-config-provider</artifactId>
+    <version>1.0.0</version>
+    <name>Azure Key Vault configuration provider</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.sdk.bom.version>1.3.8</azure.sdk.bom.version>
+    </properties>
+
+    <dependencyManagement>
+        <dependencies>
+            <dependency>
+                <groupId>com.azure</groupId>
+                <artifactId>azure-sdk-bom</artifactId>
+                <version>${azure.sdk.bom.version}</version>
+                <type>pom</type>
+                <scope>import</scope>
+            </dependency>
+        </dependencies>
+    </dependencyManagement>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-security-keyvault-secrets</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>org.slf4j</groupId>
+            <a
... (truncated)
```

### ✅ 12. powershell — 4435ms

**Input:**

```json
{"command":"mvn -q -DskipTests compile","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 0 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\keyvaultconfig\AsyncKeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretCache.java`
- `src\main\java\com\example\keyvaultconfig\AsyncSecretRotationHelper.java`
- `src\main\java\com\example\keyvaultconfig\ConfigSecret.java`
- `src\main\java\com\example\keyvaultconfig\KeyVaultClientFactory.java`
- `src\main\java\com\example\keyvaultconfig\KeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvaultconfig\Main.java`
- `src\main\java\com\example\keyvaultconfig\SecretCache.java`
- `src\main\java\com\example\keyvaultconfig\SecretRotationHelper.java`

## Copilot Response

I’ll ground the implementation in current Azure Java SDK patterns, then scaffold and verify the Maven project locally without contacting Azure.

The workspace is empty. I’m confirming the exact Java SDK polling and secret-model APIs before writing the provider, cache, rotation helpers, and demo.

The design is set: missing secrets are represented explicitly with defaults, cache reads refresh near-expiry entries, and rotation waits for delete, purges the soft-deleted record, waits for purge propagation, then retries creation only on transient name conflicts.

The complete project is in place. I’m compiling against the selected Azure SDK BOM now and will adjust any API mismatches exposed by the actual artifacts.

Created a complete Java 17 Maven project with:

- Sync and Reactor-based async Key Vault secret providers
- Missing-secret defaults and version-specific retrieval
- Expiry metadata and warning-window checks
- Thread-safe in-memory caches with startup bulk loading, refresh, and expiry-triggered re-fetch
- Managed identity authentication using `KEY_VAULT_URL` and optional `AZURE_CLIENT_ID`
- Safe sync/async rotation that waits for deletion, purges the soft-deleted secret, waits for propagation, then creates the replacement
- A `Main` demo running the full synchronous workflow followed by the asynchronous workflow
- Azure SDK BOM-based `pom.xml` and usage documentation

Rotation requires purge permission and a vault without purge protection; this limitation is documented in `README.md`.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Client lifecycle handling is appropriate for the Azure Key Vault SecretClient and SecretAsyncClient APIs.

## Grader Results

- secret-config.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (10/10)
      - Secret versioning: retrieves specific version via `getSecret(name, version)`: Pass
      - Secret expiry: accesses `properties().getExpiresOn()` on `SecretProperties`: Pass
      - Configurable warning window for near-expiry detection: Pass
      - In-memory caching (e.g., `ConcurrentHashMap`) with bulk-load and single-key refresh: Pass
      - Secret rotation uses `beginDeleteSecret()` as a long-running operation: Pass
      - Sync uses `SyncPoller` to wait for delete completion: Pass
      - Async uses `PollerFlux` to wait for delete completion: Pass
      - Creates new secret only after delete completes (not concurrently): Pass
      - Returns a default value when secret is not found (does not crash): Pass
      - NOT using fire-and-forget `deleteSecret()` without waiting for completion: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Pass
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
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Fail
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-java-secret-config --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
