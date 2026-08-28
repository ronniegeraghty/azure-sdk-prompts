# Evaluation Report: storage-dp-java-encrypted-uploader

**Config:** java-azure-skills/azure-skill-mcp | **Result:** ✅ PASSED | **Duration:** 774.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-encrypted-uploader` |
| Config | java-azure-skills/azure-skill-mcp |
| Result | ✅ PASSED |
| Score | 32/32 |
| Duration | 774.1s |
| Timestamp | 2026-08-27T20:53:16Z |
| Files Generated | 15 |
| Event Count | 16307 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 543.2s |
| Review | 229.6s |
| **Total** | **774.1s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=369524 out=11750 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 3515ms

**Input:**

```json
{"skill":"azure-storage"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Services

## Services

| Service | Use When | MCP Tools | CLI |
|---------|----------|-----------|-----|
| Blob Storage | Objects, files, backups, static content | `azure__storage` | `az storage blob` |
| File Shares | SMB file shares, lift-and-shift | - | `az storage file` |
| Queue Storage | Async messaging, task queues | - | `az storage queue` |
| Table Storage | NoSQL key-value (consider Cosmos DB) | - | `az storage table` |
| Data Lake | Big data analytics, hierarchical namespace | - | `az storage fs` |

## MCP Server (Preferred)

When Azure MCP is enabled:

- `azure__storage` with command `storage_account_list` - List storage accounts
- `azure__storage` with command `storage_container_list` - List containers in account
- `azure__storage` with command `storage_blob_list` - List blobs in container
- `azure__storage` with command `storage_blob_get` - Download blob content
- `azure__storage` with command `storage_blob_put` - Upload blob content

**If Azure MCP is not enabled:** Run `/azure:setup` or enable via `/mcp`.

## CLI Fallback

```bash
# List storage accounts
az storage account list --output table

# List containers
az storage container list --account-name ACCOUNT --output table

# List blobs
az storage blob list --account-name ACCOUNT --container-name CONTAINER --output table

# Download blob
az storage blob download --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH

# Upload blob
az storage blob upload --account-name ACCOUNT --container-name CONTAINER --name BLOB --file LOCAL_PATH
```

## Storage Account Tiers

| Tier | Use Case | Performance |
|------|----------|-------------|
| Standard | General purpose, backup | Milliseconds |
| Premium | Databases, high IOPS | Sub-millisecond |

## Blob Access Tiers

| Tier | Access Frequency | Cost |
|------|-----------------|------|
| Hot | Frequent | Higher storage, lower access |
| Cool | Infrequent (30+ days) | Lower storage, higher ac
... (truncated)
```

### ❌ 2. azure-get_azure_bestpractices (via azure) — 191217ms

**Input:**

```json
{"intent":"Generate a local-only Java 17 Maven sample using Azure Blob Storage and Azure Key Vault Keys with managed identity, shared TokenCredential, sync and async clients, and envelope encryption. Do not deploy or modify Azure resources.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 3. view — 191213ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\sdk\\azure-storage-blob-java.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/sdk/azure-storage-blob-java.md
@@ -1,37 +1,37 @@
 # Blob Storage — Java SDK Quick Reference
 
 > Condensed from **azure-storage-blob-java**. Full patterns (SAS tokens,
 > streaming, lease management, parallel uploads, proxy config)
 > in the **azure-storage-blob-java** plugin skill if installed.
 
 ## Install
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-storage-blob</artifactId>
     <version>12.33.0</version>
 </dependency>
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
 </dependency>
 ```
 
 ## Quick Start
 ```java
 import com.azure.storage.blob.BlobServiceClientBuilder;
 import com.azure.identity.DefaultAzureCredentialBuilder;
 var serviceClient = new BlobServiceClientBuilder()
     .endpoint("<storage-account-url>")
     .credential(new DefaultAzureCredentialBuilder().build())
     .buildClient();
 ```
 
 ## Best Practices
 - Use DefaultAzureCredential for **local development only** — in production, use ManagedIdentityCredential. See [auth-best-practices.md](../auth-best-practices.md)
 - Use `BinaryData.fromString()` for string uploads
 - Use `createIfNotExists()` for idempotent container creation
 - Use `BlobParallelUploadOptions` for large file uploads with headers/metadata
 - Use `BlobInputStream`/`BlobOutputStream` for streaming 
... (truncated)
```

### ✅ 4. view — 189647ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-storage\\references\\auth-best-practices.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-storage/references/auth-best-practices.md
@@ -1,129 +1,129 @@
 # Azure Authentication Best Practices
 
 > Source: [Microsoft — Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview) and [Azure Identity client libraries](https://learn.microsoft.com/dotnet/azure/sdk/authentication/).
 
 ## Golden Rule
 
 Use **managed identities** and **Azure RBAC** in production. Reserve `DefaultAzureCredential` for **local development only**.
 
 ## Authentication by Environment
 
 | Environment | Recommended Credential | Why |
 |---|---|---|
 | **Production (Azure-hosted)** | `ManagedIdentityCredential` (system- or user-assigned) | No secrets to manage; auto-rotated by Azure |
 | **Production (on-premises)** | `ClientCertificateCredential` or `WorkloadIdentityCredential` | Deterministic; no fallback chain overhead |
 | **CI/CD pipelines** | `AzurePipelinesCredential` / `WorkloadIdentityCredential` | Scoped to pipeline identity |
 | **Local development** | `DefaultAzureCredential` | Chains CLI, PowerShell, and VS Code credentials for convenience |
 
 ## Why Not `DefaultAzureCredential` in Production?
 
 1. **Unpredictable fallback chain** — walks through multiple credential types, adding latency and making failures harder to diagnose.
 2. **Broad surface area** — checks environment variables, CLI tokens, and oth
... (truncated)
```

### ✅ 5. glob — 188313ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 187372ms

**Input:**

```json
{"intent":"Get Java Azure SDK code-generation best practices for Blob Storage, Key Vault Keys, managed identity, and Reactor async APIs.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-documentation (via azure) — 187375ms

**Input:**

```json
{"intent":"Find current official Java SDK examples and API guidance for CryptographyClient wrapKey/unwrapKey and CryptographyAsyncClient, BlobClient/BlobAsyncClient upload and metadata, and ManagedIdentityCredentialBuilder.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 8. web_search (via github-mcp-server) — 187366ms

**Input:**

```json
{"query":"site:learn.microsoft.com java azure security keyvault keys cryptographyclient wrapKey unwrapKey async blobclient metadata upload BinaryData Maven latest SDK"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's an overview and guidance for your needs using the latest Azure SDK for Java, specifically regarding Azure Key Vault for keys/cryptography async key wrapping/unwrapping, and BlobClient metadata upload with BinaryData and Maven usage.\n\n### 1. **Key Vault Keys & CryptographyClient Async (Java)**\n- **Package**: `com.azure.security.keyvault.keys.cryptography`\n- **Maven dependency** (managed by BOM):\n    ```xml\n    <dependencyManagement>\n      <dependencies>\n        <dependency>\n          <groupId>com.azure</groupId>\n          <artifactId>azure-sdk-bom</artifactId>\n          <version>1.2.21</version> <!-- Use latest version from BOM list -->\n          <type>pom</type>\n          <scope>import</scope>\n        </dependency>\n      </dependencies>\n    </dependencyManagement>\n    <dependencies>\n      <dependency>\n        <groupId>com.azure</groupId>\n        <artifactId>azure-security-keyvault-keys</artifactId>\n      </dependency>\n    </dependencies>\n    ```\n  Or add directly:\n    ```xml\n    <dependency>\n      <groupId>com.azure</groupId>\n      <artifactId>azure-security-keyvault-keys</artifactId>\n      <version>4.11.0</version> <!-- Check for latest on Maven Central -->\n    </dependency>\n    ```\n- **Async usage**: Use the `CryptographyAsyncClient`.\n  ```java\n  CryptographyAsyncClient cryptoAsyncClient = new CryptographyClientBuilder()\n      .keyIdentifier(\"<your-key-id>\")\n      .credential(new DefaultAzureCredentialBuilder().build())\n      .buildAsyncClient();\n  ```\n- **Async wrap/unwrap example**:\n  ```java\n  // Wrap a key\n  cryptoAsyncClient.wrapKey(KeyWrapAlgorithm.RSA_OAEP, keyToWrap)\n      .subscribe(wrapResult -> {\n          byte[] wrappedKey = wrapResult.getEncryptedKey();\n          // store or use wrappedKey\n      });\n\n  // Unwrap a key\n  cryptoAsyncClient.unwrapKey(KeyWrapAlgorithm.RSA_OAEP, wrappedKey)\n      .subscribe(unwrapResult -> {\n          byte[] unwrappedKey = un
... (truncated)
```

### ✅ 9. web_fetch — 7798ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-keys/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-security-keyvault-keys</artifactId>
  <versioning>
    <latest>4.12.0-beta.1</latest>
    <release>4.12.0-beta.1</release>
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
      <version>4.2.0-beta.5</version>
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
      <version>4.3.0-beta.7</version>
      <version>4.3.0-beta.8</version>
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
      <version>4.4.0-beta.6</version>
      <version>4.4.0-beta.7</version>
      <versi
... (truncated)
```

### ✅ 10. web_fetch — 7289ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-keys/maven-metadata.xml"}
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

### ✅ 11. web_fetch — 6818ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
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

### ✅ 12. web_fetch — 6314ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-storage-blob</artifactId>
  <versioning>
    <latest>12.36.0-beta.1</latest>
    <release>12.36.0-beta.1</release>
    <versions>
      <version>12.7.0</version>
      <version>12.8.0-beta.1</version>
      <version>12.8.0</version>
      <version>12.9.0-beta.1</version>
      <version>12.9.0-beta.2</version>
      <version>12.9.0</version>
      <version>12.10.0-beta.1</version>
      <version>12.10.0</version>
      <version>12.10.1</version>
      <version>12.10.2</version>
      <version>12.11.0-beta.1</version>
      <version>12.11.0-beta.2</version>
      <version>12.11.0-beta.3</version>
      <version>12.11.0</version>
      <version>12.11.1</version>
      <version>12.12.0-beta.1</version>
      <version>12.12.0</version>
      <version>12.13.0</version>
      <version>12.14.0-beta.1</version>
      <version>12.14.0</version>
      <version>12.14.1</version>
      <version>12.14.2</version>
      <version>12.14.3</version>
      <version>12.14.4</version>
      <version>12.15.0-beta.1</version>
      <version>12.15.0-beta.2</version>
      <version>12.15.0-beta.3</version>
      <version>12.15.0</version>
      <version>12.16.0</version>
      <version>12.16.1</version>
      <version>12.17.0-beta.1</version>
      <version>12.17.0</version>
      <version>12.17.1</version>
      <version>12.18.0-beta.1</version>
      <version>12.18.0</version>
      <version>12.19.0</version>
      <version>12.19.1</version>
      <version>12.20.0-beta.1</version>
      <version>12.20.0</version>
      <version>12.20.1</version>
      <version>12.20.2</version>
      <version>12.20.3</version>
      <version>12.21.0-beta.1</version>
      <version>12.21.0</version>
      <version>12.21.1</version>
     
... (truncated)
```

### ✅ 13. apply_patch — 1240ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eencrypted-blob-uploader\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eEncrypted Azure Blob Uploader\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-keys\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.encryptedblob.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/encryptedblob/AzureConfiguration.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.KeyClientBuilder;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+import java.util.Map;\n+\n+public final class AzureConfiguration {\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final BlobContainerClient blobContainerClient;\n+    private final BlobContainerAsyncClient blobContainerAsyncClient;\n+    private final KeyClient keyClient;\n+    private final KeyAsyncClient keyAsyncClient;\n+\n+    private AzureConfiguration(\n+            TokenCredential credential,\n+            String keyName,\n+            BlobContainerClient blobContainerClient,\n+            BlobContainerAsyncClient blobContainerAsyncClient,\n+            KeyClient keyClient,\n+            KeyAsyncClient keyAsyncClient) {\n+        this.credential = credential;\n+        this.keyName = keyName;\n+        this.blobContainerClient = blobContainerClient;\n+        this.blobContainerAsyncClient = blobContainerAsyncClient;\n+        this.keyClient = keyClient;\n+        this.keyAsyncClient = keyAsyncClient;\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        return fromEnvironment(System.getenv());\n+    }\n+\n+    static AzureConfiguration fromEnvironment(Map\u003cString, String\u003e environment) {\n+        String storageEndpoint = required(environment, \"AZURE_STORAGE_BLOB_ENDPOINT\");\n+        String containerName = required(environment, \"AZURE_STORAGE_CONTAINER_NAME\");\n+        String vaultEndpoint = required(environment, \"AZURE_KEY_VAULT_ENDPOINT\");\n+        String keyName = required(environment, \"AZURE_KEY_VAULT_KEY_NAME\");\n+        String managedIdentityClientId = environment.get(\"AZURE_CLIENT_ID\");\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()\n+                .endpoint(storageEndpoint)\n+                .credential(credential);\n+        KeyClientBuilder keyBuilder = new KeyClientBuilder()\n+                .vaultUrl(vaultEndpoint)\n+                .credential(credential);\n+\n+        return new AzureConfiguration(\n+                credential,\n+                keyName,\n+                blobBuilder.buildClient().getBlobContainerClient(containerName),\n+                blobBuilder.buildAsyncClient().getBlobContainerAsyncClient(containerName),\n+                keyBuilder.buildClient(),\n+                keyBuilder.buildAsyncClient());\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e environment, String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    public TokenCredential credential() {\n+        return credential;\n+    }\n+\n+    public String keyName() {\n+        return keyName;\n+    }\n+\n+    public BlobContainerClient blobContainerClient() {\n+        return blobContainerClient;\n+    }\n+\n+    public BlobContainerAsyncClient blobContainerAsyncClient() {\n+        return blobContainerAsyncClient;\n+    }\n+\n+    public KeyClient keyClient() {\n+        return keyClient;\n+    }\n+\n+    public KeyAsyncClient keyAsyncClient() {\n+        return keyAsyncClient;\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/DataKeyMaterial.java\n+package com.example.encryptedblob;\n+\n+import javax.crypto.spec.SecretKeySpec;\n+import java.util.Arrays;\n+\n+final class DataKeyMaterial implements AutoCloseable {\n+    private final byte[] keyBytes;\n+\n+    DataKeyMaterial(byte[] keyBytes) {\n+        this.keyBytes = keyBytes;\n+    }\n+\n+    SecretKeySpec asAesKey() {\n+        return new SecretKeySpec(keyBytes, \"AES\");\n+    }\n+\n+    byte[] bytesForWrapping() {\n+        return keyBytes;\n+    }\n+\n+    @Override\n+    public void close() {\n+        Arrays.fill(keyBytes, (byte) 0);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/ProtectedDataKey.java\n+package com.example.encryptedblob;\n+\n+import java.util.Arrays;\n+\n+public record ProtectedDataKey(String keyId, String wrapAlgorithm, byte[] wrappedKey) {\n+    public ProtectedDataKey {\n+        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+\n+    @Override\n+    public byte[] wrappedKey() {\n+        return Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/GeneratedDataKey.java\n+package com.example.encryptedblob;\n+\n+final class GeneratedDataKey implements AutoCloseable {\n+    private final DataKeyMaterial plaintextKey;\n+    private final ProtectedDataKey protectedKey;\n+\n+    GeneratedDataKey(DataKeyMaterial plaintextKey, ProtectedDataKey protectedKey) {\n+        this.plaintextKey = plaintextKey;\n+        this.protectedKey = protectedKey;\n+    }\n+\n+    DataKeyMaterial plaintextKey() {\n+        return plaintextKey;\n+    }\n+\n+    ProtectedDataKey protectedKey() {\n+        return protectedKey;\n+    }\n+\n+    @Override\n+    public void close() {\n+        plaintextKey.close();\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagementClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import com.azure.security.keyvault.keys.models.KeyVaultKey;\n+\n+import java.security.SecureRandom;\n+\n+public final class KeyManagementClient {\n+    public static final String WRAP_ALGORITHM = \"RSA-OAEP-256\";\n+    private static final int DATA_KEY_BYTES = 32;\n+\n+    private final KeyClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public KeyManagementClient(KeyClient keyClient, TokenCredential credential, String keyName) {\n+        this(keyClient, credential, keyName, new SecureRandom());\n+    }\n+\n+    KeyManagementClient(\n+            KeyClient keyClient,\n+            TokenCredential credential,\n+            String keyName,\n+            SecureRandom secureRandom) {\n+        this.keyClient = keyClient;\n+        this.credential = credential;\n+        this.keyName = keyName;\n+        this.secureRandom = secureRandom;\n+    }\n+\n+    GeneratedDataKey generateAndProtectDataKey() {\n+        byte[] rawKey = new byte[DATA_KEY_BYTES];\n+        secureRandom.nextBytes(rawKey);\n+        DataKeyMaterial plaintextKey = new DataKeyMaterial(rawKey);\n+\n+        try {\n+            KeyVaultKey vaultKey = keyClient.getKey(keyName);\n+            String keyId = vaultKey.getId();\n+            CryptographyClient cryptographyClient = cryptographyClient(keyId);\n+            byte[] wrappedKey = cryptographyClient\n+                    .wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, plaintextKey.bytesForWrapping())\n+                    .getEncryptedKey();\n+            return new GeneratedDataKey(\n+                    plaintextKey,\n+                    new ProtectedDataKey(keyId, WRAP_ALGORITHM, wrappedKey));\n+        } catch (HttpResponseException exception) {\n+            plaintextKey.close();\n+            throw new KeyManagementException(\"Key Vault could not protect the data key\", exception);\n+        } catch (RuntimeException exception) {\n+            plaintextKey.close();\n+            throw exception;\n+        }\n+    }\n+\n+    DataKeyMaterial recoverDataKey(ProtectedDataKey protectedKey) {\n+        validateAlgorithm(protectedKey.wrapAlgorithm());\n+        try {\n+            byte[] rawKey = cryptographyClient(protectedKey.keyId())\n+                    .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, protectedKey.wrappedKey())\n+                    .getKey();\n+            return new DataKeyMaterial(rawKey);\n+        } catch (HttpResponseException exception) {\n+            throw new KeyManagementException(\n+                    \"Key Vault could not recover the data key; the key may be disabled or inaccessible\",\n+                    exception);\n+        }\n+    }\n+\n+    private CryptographyClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+                .keyIdentifier(keyId)\n+                .credential(credential)\n+                .buildClient();\n+    }\n+\n+    static void validateAlgorithm(String algorithm) {\n+        if (!WRAP_ALGORITHM.equals(algorithm)) {\n+            throw new IllegalArgumentException(\"Unsupported key wrap algorithm: \" + algorithm);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AsyncKeyManagementClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import reactor.core.publisher.Mono;\n+\n+import java.security.SecureRandom;\n+\n+public final class AsyncKeyManagementClient {\n+    private static final int DATA_KEY_BYTES = 32;\n+\n+    private final KeyAsyncClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public AsyncKeyManagementClient(\n+            KeyAsyncClient keyClient,\n+            TokenCredential credential,\n+            String keyName) {\n+        this(keyClient, credential, keyName, new SecureRandom());\n+    }\n+\n+    AsyncKeyManagementClient(\n+            KeyAsyncClient keyClient,\n+            TokenCredential credential,\n+            String keyName,\n+            SecureRandom secureRandom) {\n+        this.keyClient = keyClient;\n+        this.credential = credential;\n+        this.keyName = keyName;\n+        this.secureRandom = secureRandom;\n+    }\n+\n+    Mono\u003cGeneratedDataKey\u003e generateAndProtectDataKey() {\n+        return Mono.defer(() -\u003e {\n+            byte[] rawKey = new byte[DATA_KEY_BYTES];\n+            secureRandom.nextBytes(rawKey);\n+            DataKeyMaterial plaintextKey = new DataKeyMaterial(rawKey);\n+\n+            return keyClient.getKey(keyName)\n+                    .flatMap(vaultKey -\u003e {\n+                        String keyId = vaultKey.getId();\n+                        return cryptographyClient(keyId)\n+                                .wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, plaintextKey.bytesForWrapping())\n+                                .map(result -\u003e new GeneratedDataKey(\n+                                        plaintextKey,\n+                                        new ProtectedDataKey(\n+                                                keyId,\n+                                                KeyManagementClient.WRAP_ALGORITHM,\n+                                                result.getEncryptedKey())));\n+                    })\n+                    .doOnError(ignored -\u003e plaintextKey.close())\n+                    .onErrorMap(\n+                            HttpResponseException.class,\n+                            exception -\u003e new KeyManagementException(\n+                                    \"Key Vault could not protect the data key\",\n+                                    exception));\n+        });\n+    }\n+\n+    Mono\u003cDataKeyMaterial\u003e recoverDataKey(ProtectedDataKey protectedKey) {\n+        return Mono.defer(() -\u003e {\n+            KeyManagementClient.validateAlgorithm(protectedKey.wrapAlgorithm());\n+            return cryptographyClient(protectedKey.keyId())\n+                    .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, protectedKey.wrappedKey())\n+                    .map(result -\u003e new DataKeyMaterial(result.getKey()))\n+                    .onErrorMap(\n+                            HttpResponseException.class,\n+                            exception -\u003e new KeyManagementException(\n+                                    \"Key Vault could not recover the data key; \"\n+                                            + \"the key may be disabled or inaccessible\",\n+                                    exception));\n+        });\n+    }\n+\n+    private CryptographyAsyncClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+                .keyIdentifier(keyId)\n+                .credential(credential)\n+                .buildAsyncClient();\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagementException.java\n+package com.example.encryptedblob;\n+\n+public final class KeyManagementException extends RuntimeException {\n+    public KeyManagementException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobException.java\n+package com.example.encryptedblob;\n+\n+public final class EncryptedBlobException extends RuntimeException {\n+    public EncryptedBlobException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+\n+    public EncryptedBlobException(String message) {\n+        super(message);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EnvelopeEncryption.java\n+package com.example.encryptedblob;\n+\n+import javax.crypto.AEADBadTagException;\n+import javax.crypto.Cipher;\n+import javax.crypto.spec.GCMParameterSpec;\n+import java.nio.charset.StandardCharsets;\n+import java.security.GeneralSecurityException;\n+import java.security.SecureRandom;\n+import java.util.Base64;\n+import java.util.HashMap;\n+import java.util.Map;\n+\n+final class EnvelopeEncryption {\n+    static final String VERSION = \"1\";\n+    static final String CONTENT_ALGORITHM = \"A256GCM\";\n+    static final String META_VERSION = \"ce_version\";\n+    static final String META_CONTENT_ALGORITHM = \"ce_algorithm\";\n+    static final String META_IV = \"ce_iv\";\n+    static final String META_KEY_ID = \"ce_key_id\";\n+    static final String META_WRAP_ALGORITHM = \"ce_wrap_algorithm\";\n+    static final String META_WRAPPED_KEY = \"ce_wrapped_key\";\n+\n+    private static final int IV_BYTES = 12;\n+    private static final int TAG_BITS = 128;\n+\n+    private final SecureRandom secureRandom;\n+\n+    EnvelopeEncryption() {\n+        this(new SecureRandom());\n+    }\n+\n+    EnvelopeEncryption(SecureRandom secureRandom) {\n+        this.secureRandom = secureRandom;\n+    }\n+\n+    EncryptedPayload encrypt(byte[] plaintext, DataKeyMaterial dataKey, String blobName) {\n+        byte[] iv = new byte[IV_BYTES];\n+        secureRandom.nextBytes(iv);\n+        try {\n+            Cipher cipher = Cipher.getInstance(\"AES/GCM/NoPadding\");\n+            cipher.init(Cipher.ENCRYPT_MODE, dataKey.asAesKey(), new GCMParameterSpec(TAG_BITS, iv));\n+            cipher.updateAAD(aad(blobName));\n+            return new EncryptedPayload(cipher.doFinal(plaintext), iv);\n+        } catch (GeneralSecurityException exception) {\n+            throw new EncryptedBlobException(\"Local encryption failed\", exception);\n+        }\n+    }\n+\n+    byte[] decrypt(\n+            byte[] ciphertext,\n+            byte[] iv,\n+            DataKeyMaterial dataKey,\n+            String blobName) {\n+        try {\n+            Cipher cipher = Cipher.getInstance(\"AES/GCM/NoPadding\");\n+            cipher.init(Cipher.DECRYPT_MODE, dataKey.asAesKey(), new GCMParameterSpec(TAG_BITS, iv));\n+            cipher.updateAAD(aad(blobName));\n+            return cipher.doFinal(ciphertext);\n+        } catch (AEADBadTagException exception) {\n+            throw new EncryptedBlobException(\n+                    \"Ciphertext or encryption metadata failed authentication\",\n+                    exception);\n+        } catch (GeneralSecurityException exception) {\n+            throw new EncryptedBlobException(\"Local decryption failed\", exception);\n+        }\n+    }\n+\n+    Map\u003cString, String\u003e metadata(EncryptedPayload payload, ProtectedDataKey protectedKey) {\n+        Map\u003cString, String\u003e metadata = new HashMap\u003c\u003e();\n+        metadata.put(META_VERSION, VERSION);\n+        metadata.put(META_CONTENT_ALGORITHM, CONTENT_ALGORITHM);\n+        metadata.put(META_IV, Base64.getEncoder().encodeToString(payload.iv()));\n+        metadata.put(META_KEY_ID, protectedKey.keyId());\n+        metadata.put(META_WRAP_ALGORITHM, protectedKey.wrapAlgorithm());\n+        metadata.put(META_WRAPPED_KEY, Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));\n+        return metadata;\n+    }\n+\n+    EnvelopeMetadata parseMetadata(Map\u003cString, String\u003e metadata) {\n+        String version = required(metadata, META_VERSION);\n+        if (!VERSION.equals(version)) {\n+            throw new EncryptedBlobException(\"Unsupported envelope metadata version: \" + version);\n+        }\n+        String contentAlgorithm = required(metadata, META_CONTENT_ALGORITHM);\n+        if (!CONTENT_ALGORITHM.equals(contentAlgorithm)) {\n+            throw new EncryptedBlobException(\n+                    \"Unsupported content encryption algorithm: \" + contentAlgorithm);\n+        }\n+\n+        try {\n+            byte[] iv = Base64.getDecoder().decode(required(metadata, META_IV));\n+            if (iv.length != IV_BYTES) {\n+                throw new EncryptedBlobException(\"Invalid AES-GCM IV length: \" + iv.length);\n+            }\n+            ProtectedDataKey protectedKey = new ProtectedDataKey(\n+                    required(metadata, META_KEY_ID),\n+                    required(metadata, META_WRAP_ALGORITHM),\n+                    Base64.getDecoder().decode(required(metadata, META_WRAPPED_KEY)));\n+            return new EnvelopeMetadata(iv, protectedKey);\n+        } catch (IllegalArgumentException exception) {\n+            throw new EncryptedBlobException(\"Encryption metadata contains invalid Base64\", exception);\n+        }\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e metadata, String name) {\n+        String value = metadata.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new EncryptedBlobException(\"Blob is missing encryption metadata: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static byte[] aad(String blobName) {\n+        return (\"encrypted-blob:v1:\" + blobName).getBytes(StandardCharsets.UTF_8);\n+    }\n+\n+    record EncryptedPayload(byte[] ciphertext, byte[] iv) {\n+    }\n+\n+    record EnvelopeMetadata(byte[] iv, ProtectedDataKey protectedKey) {\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/UploadResult.java\n+package com.example.encryptedblob;\n+\n+public record UploadResult(String keyId, String wrappedDataKeyBase64) {\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.core.util.BinaryData;\n+import com.azure.core.util.Context;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.models.BlobRequestConditions;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.models.DownloadRetryOptions;\n+import com.azure.storage.blob.options.BlockBlobSimpleUploadOptions;\n+import com.azure.storage.blob.specialized.BlockBlobClient;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.time.Duration;\n+import java.util.Base64;\n+\n+public final class EncryptedBlobClient {\n+    private final BlobContainerClient containerClient;\n+    private final KeyManagementClient keyManagementClient;\n+    private final EnvelopeEncryption encryption;\n+\n+    public EncryptedBlobClient(\n+            BlobContainerClient containerClient,\n+            KeyManagementClient keyManagementClient) {\n+        this.containerClient = containerClient;\n+        this.keyManagementClient = keyManagementClient;\n+        this.encryption = new EnvelopeEncryption();\n+    }\n+\n+    public UploadResult upload(Path source, String blobName) {\n+        try {\n+            return upload(Files.readAllBytes(source), blobName);\n+        } catch (IOException exception) {\n+            throw new EncryptedBlobException(\"Could not read source file: \" + source, exception);\n+        }\n+    }\n+\n+    public UploadResult upload(byte[] plaintext, String blobName) {\n+        try (GeneratedDataKey generatedKey = keyManagementClient.generateAndProtectDataKey()) {\n+            EnvelopeEncryption.EncryptedPayload payload =\n+                    encryption.encrypt(plaintext, generatedKey.plaintextKey(), blobName);\n+            ProtectedDataKey protectedKey = generatedKey.protectedKey();\n+            BlockBlobSimpleUploadOptions options = new BlockBlobSimpleUploadOptions(\n+                    BinaryData.fromBytes(payload.ciphertext()))\n+                    .setMetadata(encryption.metadata(payload, protectedKey));\n+\n+            containerClient.getBlobClient(blobName)\n+                    .getBlockBlobClient()\n+                    .uploadWithResponse(options, Duration.ofMinutes(2), Context.NONE);\n+\n+            return new UploadResult(\n+                    protectedKey.keyId(),\n+                    Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));\n+        } catch (BlobStorageException exception) {\n+            throw new EncryptedBlobException(\n+                    \"Blob Storage upload failed for '\" + blobName + \"': \"\n+                            + exception.getErrorCode(),\n+                    exception);\n+        }\n+    }\n+\n+    public byte[] download(String blobName) {\n+        BlockBlobClient blobClient = containerClient.getBlobClient(blobName).getBlockBlobClient();\n+        try {\n+            var properties = blobClient.getProperties();\n+            var metadata = encryption.parseMetadata(properties.getMetadata());\n+            BlobRequestConditions conditions = new BlobRequestConditions()\n+                    .setIfMatch(properties.getETag());\n+            byte[] ciphertext = blobClient.downloadContentWithResponse(\n+                            new DownloadRetryOptions(),\n+                            conditions,\n+                            Duration.ofMinutes(2),\n+                            Context.NONE)\n+                    .getValue()\n+                    .toBytes();\n+\n+            try (DataKeyMaterial dataKey =\n+                         keyManagementClient.recoverDataKey(metadata.protectedKey())) {\n+                return encryption.decrypt(\n+                        ciphertext,\n+                        metadata.iv(),\n+                        dataKey,\n+                        blobName);\n+            }\n+        } catch (BlobStorageException exception) {\n+            String detail = exception.getStatusCode() == 404\n+                    ? \"blob does not exist\"\n+                    : String.valueOf(exception.getErrorCode());\n+            throw new EncryptedBlobException(\n+                    \"Blob Storage download failed for '\" + blobName + \"': \" + detail,\n+                    exception);\n+        } catch (HttpResponseException exception) {\n+            throw new EncryptedBlobException(\n+                    \"Blob Storage download failed for '\" + blobName + \"'\",\n+                    exception);\n+        }\n+    }\n+\n+    public void download(String blobName, Path destination) {\n+        try {\n+            Files.write(destination, download(blobName));\n+        } catch (IOException exception) {\n+            throw new EncryptedBlobException(\n+                    \"Could not write decrypted file: \" + destination,\n+                    exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.models.BlobRequestConditions;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.models.DownloadRetryOptions;\n+import com.azure.storage.blob.options.BlockBlobSimpleUploadOptions;\n+import reactor.core.publisher.Mono;\n+import reactor.core.scheduler.Schedulers;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.util.Base64;\n+\n+public final class EncryptedBlobAsyncClient {\n+    private final BlobContainerAsyncClient containerClient;\n+    private final AsyncKeyManagementClient keyManagementClient;\n+    private final EnvelopeEncryption encryption;\n+\n+    public EncryptedBlobAsyncClient(\n+            BlobContainerAsyncClient containerClient,\n+            AsyncKeyManagementClient keyManagementClient) {\n+        this.containerClient = containerClient;\n+        this.keyManagementClient = keyManagementClient;\n+        this.encryption = new EnvelopeEncryption();\n+    }\n+\n+    public Mono\u003cUploadResult\u003e upload(Path source, String blobName) {\n+        return Mono.fromCallable(() -\u003e Files.readAllBytes(source))\n+                .subscribeOn(Schedulers.boundedElastic())\n+                .onErrorMap(\n+                        IOException.class,\n+                        exception -\u003e new EncryptedBlobException(\n+                                \"Could not read source file: \" + source,\n+                                exception))\n+                .flatMap(bytes -\u003e upload(bytes, blobName));\n+    }\n+\n+    public Mono\u003cUploadResult\u003e upload(byte[] plaintext, String blobName) {\n+        return keyManagementClient.generateAndProtectDataKey()\n+                .flatMap(generatedKey -\u003e Mono.using(\n+                        () -\u003e generatedKey,\n+                        key -\u003e {\n+                            EnvelopeEncryption.EncryptedPayload payload =\n+                                    encryption.encrypt(plaintext, key.plaintextKey(), blobName);\n+                            ProtectedDataKey protectedKey = key.protectedKey();\n+                            BlockBlobSimpleUploadOptions options =\n+                                    new BlockBlobSimpleUploadOptions(\n+                                            BinaryData.fromBytes(payload.ciphertext()))\n+                                            .setMetadata(encryption.metadata(payload, protectedKey));\n+\n+                            return containerClient.getBlobAsyncClient(blobName)\n+                                    .getBlockBlobAsyncClient()\n+                                    .uploadWithResponse(options)\n+                                    .map(ignored -\u003e new UploadResult(\n+                                            protectedKey.keyId(),\n+                                            Base64.getEncoder().encodeToString(\n+                                                    protectedKey.wrappedKey())));\n+                        },\n+                        GeneratedDataKey::close))\n+                .onErrorMap(\n+                        BlobStorageException.class,\n+                        exception -\u003e new EncryptedBlobException(\n+                                \"Blob Storage upload failed for '\" + blobName + \"': \"\n+                                        + exception.getErrorCode(),\n+                                exception));\n+    }\n+\n+    public Mono\u003cbyte[]\u003e download(String blobName) {\n+        var blobClient = containerClient.getBlobAsyncClient(blobName).getBlockBlobAsyncClient();\n+        return blobClient.getProperties()\n+                .flatMap(properties -\u003e {\n+                    var metadata = encryption.parseMetadata(properties.getMetadata());\n+                    BlobRequestConditions conditions = new BlobRequestConditions()\n+                            .setIfMatch(properties.getETag());\n+                    return blobClient.downloadContentWithResponse(\n+                                    new DownloadRetryOptions(),\n+                                    conditions)\n+                            .map(response -\u003e response.getValue().toBytes())\n+                            .flatMap(ciphertext -\u003e keyManagementClient\n+                                    .recoverDataKey(metadata.protectedKey())\n+                                    .flatMap(dataKey -\u003e Mono.using(\n+                                            () -\u003e dataKey,\n+                                            key -\u003e Mono.fromCallable(() -\u003e encryption.decrypt(\n+                                                    ciphertext,\n+                                                    metadata.iv(),\n+                                                    key,\n+                                                    blobName)),\n+                                            DataKeyMaterial::close)));\n+                })\n+                .onErrorMap(\n+                        BlobStorageException.class,\n+                        exception -\u003e {\n+                            String detail = exception.getStatusCode() == 404\n+                                    ? \"blob does not exist\"\n+                                    : String.valueOf(exception.getErrorCode());\n+                            return new EncryptedBlobException(\n+                                    \"Blob Storage download failed for '\" + blobName + \"': \" + detail,\n+                                    exception);\n+                        });\n+    }\n+\n+    public Mono\u003cVoid\u003e download(String blobName, Path destination) {\n+        return download(blobName)\n+                .flatMap(bytes -\u003e Mono.fromCallable(() -\u003e {\n+                            Files.write(destination, bytes);\n+                            return destination;\n+                        })\n+                        .subscribeOn(Schedulers.boundedElastic())\n+                        .onErrorMap(\n+                                IOException.class,\n+                                exception -\u003e new EncryptedBlobException(\n+                                        \"Could not write decrypted file: \" + destination,\n+                                        exception)))\n+                .then();\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/Main.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+\n+public final class Main {\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws Exception {\n+        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();\n+        Path source = args.length \u003e 0\n+                ? Path.of(args[0])\n+                : Files.createTempFile(\"encrypted-blob-demo-\", \".txt\");\n+        if (args.length == 0) {\n+            Files.writeString(\n+                    source,\n+                    \"Client-side envelope encryption with Azure Key Vault.\",\n+                    StandardCharsets.UTF_8);\n+        }\n+\n+        KeyManagementClient keyManager = new KeyManagementClient(\n+                configuration.keyClient(),\n+                configuration.credential(),\n+                configuration.keyName());\n+        EncryptedBlobClient syncClient = new EncryptedBlobClient(\n+                configuration.blobContainerClient(),\n+                keyManager);\n+\n+        UploadResult syncResult = syncClient.upload(source, \"sync-encrypted-demo.bin\");\n+        byte[] syncPlaintext = syncClient.download(\"sync-encrypted-demo.bin\");\n+        printResult(\"sync\", syncResult, syncPlaintext);\n+\n+        AsyncKeyManagementClient asyncKeyManager = new AsyncKeyManagementClient(\n+                configuration.keyAsyncClient(),\n+                configuration.credential(),\n+                configuration.keyName());\n+        EncryptedBlobAsyncClient asyncClient = new EncryptedBlobAsyncClient(\n+                configuration.blobContainerAsyncClient(),\n+                asyncKeyManager);\n+\n+        asyncClient.upload(source, \"async-encrypted-demo.bin\")\n+                .flatMap(result -\u003e asyncClient.download(\"async-encrypted-demo.bin\")\n+                        .doOnNext(plaintext -\u003e printResult(\"async\", result, plaintext)))\n+                .block();\n+    }\n+\n+    private static void printResult(String implementation, UploadResult result, byte[] plaintext) {\n+        System.out.println(\"[\" + implementation + \"] vault key ID: \" + result.keyId());\n+        System.out.println(\"[\" + implementation + \"] wrapped DEK (base64): \"\n+                + result.wrappedDataKeyBase64());\n+        System.out.println(\"[\" + implementation + \"] decrypted output: \"\n+                + new String(plaintext, StandardCharsets.UTF_8));\n+    }\n+}\n*** Add File: README.md\n+# Encrypted Azure Blob Uploader\n+\n+Java 17 sample using envelope encryption:\n+\n+- Each upload generates a random AES-256 data encryption key (DEK).\n+- File contents are encrypted locally with AES-GCM.\n+- Azure Key Vault Keys wraps the DEK with `RSA-OAEP-256`.\n+- Blob metadata stores the wrapped DEK, Key Vault key ID, IV, and algorithm identifiers.\n+- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in memory and are zeroed after use.\n+\n+The managed identity needs Blob data access to the target container and Key Vault key permissions for `get`, `wrapKey`, and `unwrapKey`. The configured Key Vault key must be an RSA key that supports those operations.\n+\n+## Configuration\n+\n+```text\n+AZURE_STORAGE_BLOB_ENDPOINT=https://\u003caccount\u003e.blob.core.windows.net\n+AZURE_STORAGE_CONTAINER_NAME=\u003ccontainer\u003e\n+AZURE_KEY_VAULT_ENDPOINT=https://\u003cvault\u003e.vault.azure.net\n+AZURE_KEY_VAULT_KEY_NAME=\u003crsa-key-name\u003e\n+AZURE_CLIENT_ID=\u003coptional-user-assigned-managed-identity-client-id\u003e\n+```\n+\n+The container and RSA key must already exist. This project does not create or modify Azure resources.\n+\n+## Build and run\n+\n+```text\n+mvn clean package\n+mvn exec:java -Dexec.args=\"C:\\path\\to\\input.txt\"\n+```\n+\n+Without an argument, `Main` creates a temporary UTF-8 demo file. It performs a sync round trip followed by an async round trip.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/pom.xml
@@ -1,0 +1,69 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>encrypted-blob-uploader</artifactId>
+    <version>1.0.0</version>
+    <name>Encrypted Azure Blob Uploader</name>
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
+            <artifactId>azure-security-keyvault-keys</artifactId>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifa
... (truncated)
```

### ✅ 14. powershell — 6792ms

**Input:**

```json
{"command":"mvn -q -DskipTests package","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
<shellId: 0 completed with exit code 0>
```

### ✅ 15. apply_patch — 644ms

**Input:**

```json
"*** Begin Patch\n*** Update File: README.md\n@@\n-- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in memory and are zeroed after use.\n+- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in process memory; the application's raw-key buffers are zeroed after use.\n*** Add File: .gitignore\n+target/\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/README.md b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/README.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/README.md
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/README.md
@@ -6,7 +6,7 @@
 - File contents are encrypted locally with AES-GCM.
 - Azure Key Vault Keys wraps the DEK with `RSA-OAEP-256`.
 - Blob metadata stores the wrapped DEK, Key Vault key ID, IV, and algorithm identifiers.
-- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in memory and are zeroed after use.
+- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in process memory; the application's raw-key buffers are zeroed after use.
 
 The managed identity needs Blob data access to the target container and Key Vault key permissions for `get`, `wrapKey`, and `unwrapKey`. The configured Key Vault key must be an RSA key that supports those operations.
 


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/.gitignore b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/.gitignore
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-1129941592/.gitignore
@@ -1,0 +1,2 @@
+target/
+


```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\encryptedblob\AsyncKeyManagementClient.java`
- `src\main\java\com\example\encryptedblob\AzureConfiguration.java`
- `src\main\java\com\example\encryptedblob\DataKeyMaterial.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobAsyncClient.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobException.java`
- `src\main\java\com\example\encryptedblob\EnvelopeEncryption.java`
- `src\main\java\com\example\encryptedblob\GeneratedDataKey.java`
- `src\main\java\com\example\encryptedblob\KeyManagementClient.java`
- `src\main\java\com\example\encryptedblob\KeyManagementException.java`
- `src\main\java\com\example\encryptedblob\Main.java`
- `src\main\java\com\example\encryptedblob\ProtectedDataKey.java`
- `src\main\java\com\example\encryptedblob\UploadResult.java`

## Copilot Response

I’m grounding the implementation in the current Java SDK patterns, then I’ll scaffold and compile the local project.

The design is set: AES-256-GCM encrypts locally, RSA-OAEP-256 wraps each ephemeral DEK in Key Vault, and the blob write stores ciphertext plus versioned envelope metadata atomically.

Created the Java 17 Maven project with:

- Sync and async Key Vault **Keys** envelope-encryption clients using RSA-OAEP-256.
- Sync and async Blob Storage upload/download clients using local AES-256-GCM.
- Versioned blob metadata containing the wrapped DEK, key version ID, IV, and algorithms.
- One shared `ManagedIdentityCredential` for all Azure clients.
- A `Main` demo running sync and async round trips.
- Complete `pom.xml`, `.gitignore`, and setup instructions in `README.md`.

Configure these variables before running:

```text
AZURE_STORAGE_BLOB_ENDPOINT
AZURE_STORAGE_CONTAINER_NAME
AZURE_KEY_VAULT_ENDPOINT
AZURE_KEY_VAULT_KEY_NAME
AZURE_CLIENT_ID                 # optional user-assigned identity
```

Build and run:

```powershell
mvn clean package
mvn exec:java -Dexec.args="C:\path\to\input.txt"
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No closeable Azure SDK clients are leaked; sensitive data-key resources are scoped with try-with-resources or Reactor Mono.using.

## Grader Results

- encrypted-uploader.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (20/20)
      - Uses `com.azure:azure-security-keyvault-keys` (Keys, NOT Secrets) — critical distinction: Pass
      - Uses `javax.crypto` or `java.security` for local AES-GCM encryption: Pass
      - Uses `KeyClient` / `CryptographyClient` builder for Key Vault Keys (NOT `SecretClient`): Pass
      - Uses `CryptographyClient` for `wrapKey()` and `unwrapKey()` operations: Pass
      - Specifies RSA key wrap algorithm (`KeyWrapAlgorithm.RSA_OAEP` or `RSA_OAEP_256`): Pass
      - Key material never leaves Key Vault (wrap/unwrap is server-side): Pass
      - Generates a random AES-256 DEK locally (32 bytes): Pass
      - Encrypts data with AES-GCM locally using the DEK: Pass
      - Wraps the DEK via Key Vault `wrapKey()`: Pass
      - Stores wrapped DEK as blob metadata: Pass
      - Stores IV (initialization vector) in blob metadata: Pass
      - Stores vault key identifier in blob metadata: Pass
      - For decryption: retrieves wrapped DEK from metadata, unwraps via Key Vault, decrypts locally: Pass
      - Uses AES-GCM (not AES-CBC, AES-ECB, or other modes): Pass
      - Generates random IV for each encryption (typically 12 bytes for GCM): Pass
      - Handles Key Vault errors (key disabled, key not found): Pass
      - Uses `BlobAsyncClient` and `CryptographyAsyncClient` for async: Pass
      - NOT using `SecretClient` instead of `KeyClient`/`CryptographyClient`: Pass
      - NOT encrypting data directly with the vault key (should be envelope encryption): Pass
      - NOT storing raw DEK in plaintext: Pass
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
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 7.7% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 13.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-encrypted-uploader --config java-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
