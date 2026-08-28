# Evaluation Report: resource-manager-mp-java-rg-crud

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 575.7s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `resource-manager-mp-java-rg-crud` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/19 |
| Duration | 575.7s |
| Timestamp | 2026-08-27T18:15:10Z |
| Files Generated | 3 |
| Event Count | 6089 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 328.3s |
| Review | 245.7s |
| **Total** | **575.7s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=524151 out=5471 |
| Turn Count | 11 |

## Tool Calls

### ✅ 1. skill — 3816ms

**Input:**

```json
{"skill":"azure-identity-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Identity library for Java

Authentication library for Azure SDK clients using Microsoft Entra ID.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.15.0</version>
</dependency>
```

## Key Concepts

| Credential | Use Case |
|------------|----------|
| `DefaultAzureCredential` | **Recommended** - Works in dev and production |
| `ManagedIdentityCredential` | Azure-hosted apps (App Service, Functions, VMs) |
| `EnvironmentCredential` | CI/CD pipelines with env vars |
| `ClientSecretCredential` | Service principals with secret |
| `ClientCertificateCredential` | Service principals with certificate |
| `AzureCliCredential` | Local dev using `az login` |
| `InteractiveBrowserCredential` | Interactive login flow |
| `DeviceCodeCredential` | Headless device authentication |

## DefaultAzureCredential (Recommended)

The `DefaultAzureCredential` tries multiple authentication methods in order. See [DefaultAzureCredential overview](https://aka.ms/azsdk/java/identity/credential-chains#defaultazurecredential-overview) for the current credential chain order and defaults.

```java
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

// Simple usage
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();

// Use with any Azure client
BlobServiceClient blobClient = new BlobServiceClientBuilder()
    .endpoint("https://<storage-account>.blob.core.windows.net")
    .credential(credential)
    .buildClient();

KeyClient keyClient = new KeyClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();
```

### Configure DefaultAzureCredential

```java
DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
    .managedIdentityClientId("<user-assigned-identity-client-id>")  // For user-assigned MI
    .tenantId("<tenant-id
... (truncated)
```

### ✅ 2. view — 190246ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-identity-java/references/examples.md
@@ -1,495 +1,495 @@
 # Azure Identity library for Java - Examples
 
 Comprehensive code examples for the Azure Identity library for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [DefaultAzureCredential](#defaultazurecredential)
 - [ChainedTokenCredential](#chainedtokencredential)
 - [ClientSecretCredential](#clientsecretcredential)
 - [ClientCertificateCredential](#clientcertificatecredential)
 - [ManagedIdentityCredential](#managedidentitycredential)
 - [EnvironmentCredential](#environmentcredential)
 - [InteractiveBrowserCredential](#interactivebrowsercredential)
 - [DeviceCodeCredential](#devicecodecredential)
 - [AzureCliCredential](#azureclicredential)
 - [Using Credentials with Azure SDK Clients](#using-credentials-with-azure-sdk-clients)
 
 ## Maven Dependency
 
 ```xml
 <!-- Using Azure SDK BOM (recommended) -->
 <dependencyManagement>
     <dependencies>
         <dependency>
             <groupId>com.azure</groupId>
             <artifactId>azure-sdk-bom</artifactId>
             <version>1.2.29</version>
             <type>pom</type>
             <scope>import</scope>
         </dependency>
     </dependencies>
 </dependencyManagement>
 
 <dependencies>
     <dependency>
         <groupId>com.azure</groupId>
         <artifactId>azure-identity</artifactId>
     </dependency>
 </dependencies>
 
 <!-- Or direct depend
... (truncated)
```

### ✅ 3. glob — 188830ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ❌ 4. azure-get_azure_bestpractices (via azure) — 190250ms

**Input:**

```json
{"intent":"Generate a local-only Java Maven application using DefaultAzureCredential and the modern azure-resourcemanager management-plane SDK to create, list, get, tag, and delete a resource group, with proper exception handling. Do not execute Azure operations.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 5. azure-documentation (via azure) — 190251ms

**Input:**

```json
{"intent":"Find current official Java examples and Maven dependency guidance for com.azure.resourcemanager AzureResourceManager authentication and resource group create, list, get, update tags, and delete operations.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 6. web_fetch — 6697ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/maven-metadata.xml"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azure?view=azure-java-stable - status code 404
```

### ✅ 7. web_fetch — 6694ms

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 8. web_fetch — 6280ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azure?view=azure-java-stable"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure.resourcemanager</groupId>
  <artifactId>azure-resourcemanager</artifactId>
  <versioning>
    <latest>2.63.0</latest>
    <release>2.63.0</release>
    <versions>
      <version>1.0.0-hybrid</version>
      <version>2.0.0-beta.1</version>
      <version>2.0.0-beta.2</version>
      <version>2.0.0-beta.3</version>
      <version>2.0.0-beta.4</version>
      <version>2.0.0-beta.5</version>
      <version>2.0.0</version>
      <version>2.1.0</version>
      <version>2.2.0</version>
      <version>2.3.0</version>
      <version>2.4.0</version>
      <version>2.5.0</version>
      <version>2.6.0</version>
      <version>2.7.0</version>
      <version>2.8.0</version>
      <version>2.9.0</version>
      <version>2.10.0</version>
      <version>2.11.0</version>
      <version>2.12.0</version>
      <version>2.13.0</version>
      <version>2.14.0</version>
      <version>2.15.0</version>
      <version>2.16.0</version>
      <version>2.17.0</version>
      <version>2.18.0</version>
      <version>2.19.0</version>
      <version>2.20.0</version>
      <version>2.21.0</version>
      <version>2.22.0</version>
      <version>2.23.0</version>
      <version>2.24.0</version>
      <version>2.25.0</version>
      <version>2.26.0</version>
      <version>2.27.0</version>
      <version>2.28.0</version>
      <version>2.29.0</version>
      <version>2.30.0</version>
      <version>2.31.0</version>
      <version>2.32.0</version>
      <version>2.33.0</version>
      <version>2.34.0</version>
      <version>2.35.0</version>
      <version>2.36.0</version>
      <version>2.37.0</version>
      <version>2.38.0</version>
      <version>2.39.0</version>
      <version>2.40.0</version>
      <version>2.41.0</version>
      <versi
... (truncated)
```

### ✅ 9. powershell — 7144ms

**Input:**

```json
{"command":"$java = Get-Command java -ErrorAction SilentlyContinue; $mvn = Get-Command mvn -ErrorAction SilentlyContinue; [pscustomobject]@{ Java = if ($java) { \u0026 java -version 2\u003e\u00261 | Select-Object -First 1 } else { 'missing' }; Maven = if ($mvn) { \u0026 mvn -version 2\u003e\u00261 | Select-Object -First 1 } else { 'missing' } } | Format-List","description":"Check installed Java and Maven versions"}
```

**Output:**

```
Java  : openjdk version "25.0.1" 2025-10-21 LTS
Maven : Apache Maven 3.9.11 (3e54c93a704957b63ee3494413a2b544fd3d825b)
<shellId: 0 completed with exit code 0>
```

### ❌ 10. web_fetch — 5258ms

**Input:**

```json
{"max_length":4000,"start_index":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azure?view=azure-java-stable&preserve-view=true - status code 404
```

### ✅ 11. web_fetch — 5255ms

**Input:**

```json
{"max_length":12000,"url":"https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/2.63.0/azure-resourcemanager-2.63.0.pom"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/resourcemanager/azure-resourcemanager/2.63.0/azure-resourcemanager-2.63.0.pom:
<!--
 Copyright (c) Microsoft Corporation. All rights reserved.
 Licensed under the MIT License. See License.txt in the project root for
 license information.
-->
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <parent>
    <groupId>com.azure</groupId>
    <artifactId>azure-client-sdk-parent</artifactId>
    <version>1.7.0</version> <!-- {x-version-update;com.azure:azure-client-sdk-parent;current} -->
    <relativePath>../../parents/azure-client-sdk-parent</relativePath>
  </parent>

  <groupId>com.azure.resourcemanager</groupId>
  <artifactId>azure-resourcemanager</artifactId>
  <version>2.63.0</version> <!-- {x-version-update;com.azure.resourcemanager:azure-resourcemanager;current} -->
  <packaging>jar</packaging>

  <name>Microsoft Azure SDK for Management</name>
  <description>This package contains Microsoft Azure SDK for Management. For documentation on how to use this package, please see https://aka.ms/azsdk/java/mgmt</description>
  <url>https://github.com/Azure/azure-sdk-for-java</url>

  <licenses>
    <license>
      <name>The MIT License (MIT)</name>
      <url>http://opensource.org/licenses/MIT</url>
      <distribution>repo</distribution>
    </license>
  </licenses>

  <scm>
    <url>https://github.com/Azure/azure-sdk-for-java</url>
    <connection>scm:git:git@github.com:Azure/azure-sdk-for-java.git</connection>
    <tag>HEAD</tag>
  </scm>

  <properties>
    <!-- RMJacoco -->
    <jacoco.min.linecoverage>0.10</jacoco.min.linecoverage>
    <jacoco.min.branchcoverage>0.10</jacoco.min.branchcoverage>
    <!-- Configures the Java 9+ run to perform the 
... (truncated)
```

### ✅ 12. web_fetch — 4867ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azure?view=azure-java-stable\u0026preserve-view=true"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
5.4</version>
      <version>1.16.0-beta.1</version>
      <version>1.16.0</version>
      <version>1.16.1</version>
      <version>1.16.2</version>
      <version>1.16.3</version>
      <version>1.17.0-beta.1</version>
      <version>1.17.0</version>
      <version>1.18.0</version>
      <version>1.18.1</version>
      <version>1.18.2</version>
      <version>1.18.3</version>
      <version>1.18.4</version>
      <version>1.18.5</version>
      <version>1.19.0-beta.1</version>
      <version>1.19.0-beta.2</version>
    </versions>
    <lastUpdated>20260824185316</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 13. apply_patch — 724ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eresource-group-manager\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure-resourcemanager.version\u003e2.63.0\u003c/azure-resourcemanager.version\u003e\n+        \u003cazure-identity.version\u003e1.18.5\u003c/azure-identity.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure-resourcemanager.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure-identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azure.ResourceGroupManager\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azure/ResourceGroupManager.java\n+package com.example.azure;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.core.management.AzureEnvironment;\n+import com.azure.core.management.AzureProfile;\n+import com.azure.core.management.exception.ManagementException;\n+import com.azure.identity.CredentialUnavailableException;\n+import com.azure.identity.DefaultAzureCredential;\n+import com.azure.identity.DefaultAzureCredentialBuilder;\n+import com.azure.resourcemanager.AzureResourceManager;\n+import com.azure.resourcemanager.resources.models.ResourceGroup;\n+\n+import java.util.UUID;\n+\n+public final class ResourceGroupManager {\n+    private static final String LOCATION = \"eastus\";\n+    private static final String TAG_NAME = \"managed-by\";\n+    private static final String TAG_VALUE = \"azure-sdk-for-java\";\n+\n+    private ResourceGroupManager() {\n+    }\n+\n+    public static void main(String[] args) {\n+        String subscriptionId = requireEnvironmentVariable(\"AZURE_SUBSCRIPTION_ID\");\n+        String resourceGroupName = getResourceGroupName();\n+\n+        ResourceGroup createdGroup = null;\n+        boolean deleted = false;\n+\n+        try {\n+            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();\n+            AzureProfile profile = new AzureProfile(\n+                null,\n+                null,\n+                subscriptionId,\n+                AzureEnvironment.AZURE);\n+\n+            AzureResourceManager azure = AzureResourceManager\n+                .authenticate(credential, profile)\n+                .withSubscription(subscriptionId);\n+\n+            System.out.printf(\"Creating resource group '%s' in %s...%n\",\n+                resourceGroupName, LOCATION);\n+            createdGroup = azure.resourceGroups()\n+                .define(resourceGroupName)\n+                .withRegion(LOCATION)\n+                .create();\n+\n+            System.out.println(\"\\nResource groups in the subscription:\");\n+            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {\n+                System.out.printf(\"- %s (%s)%n\",\n+                    resourceGroup.name(), resourceGroup.regionName());\n+            }\n+\n+            ResourceGroup fetchedGroup = azure.resourceGroups()\n+                .getByName(resourceGroupName);\n+            if (fetchedGroup == null) {\n+                throw new IllegalStateException(\n+                    \"The created resource group could not be retrieved: \" + resourceGroupName);\n+            }\n+\n+            System.out.println(\"\\nCreated resource group details:\");\n+            printResourceGroup(fetchedGroup);\n+\n+            ResourceGroup taggedGroup = fetchedGroup.update()\n+                .withTag(TAG_NAME, TAG_VALUE)\n+                .apply();\n+            System.out.printf(\"%nAdded tag %s=%s. Current tags: %s%n\",\n+                TAG_NAME, TAG_VALUE, taggedGroup.tags());\n+\n+            System.out.printf(\"%nDeleting resource group '%s'...%n\", resourceGroupName);\n+            azure.resourceGroups().deleteByName(resourceGroupName);\n+            deleted = true;\n+            System.out.println(\"Resource group deleted.\");\n+        } catch (CredentialUnavailableException exception) {\n+            System.err.println(\"No credential was available to DefaultAzureCredential: \"\n+                + exception.getMessage());\n+        } catch (ClientAuthenticationException exception) {\n+            System.err.println(\"Azure authentication failed: \" + exception.getMessage());\n+        } catch (ManagementException exception) {\n+            System.err.printf(\"Azure Resource Manager request failed (status %d): %s%n\",\n+                exception.getResponse().getStatusCode(), exception.getMessage());\n+        } finally {\n+            if (createdGroup != null \u0026\u0026 !deleted) {\n+                deleteAfterFailure(createdGroup);\n+            }\n+        }\n+    }\n+\n+    private static void printResourceGroup(ResourceGroup resourceGroup) {\n+        System.out.println(\"Name: \" + resourceGroup.name());\n+        System.out.println(\"ID: \" + resourceGroup.id());\n+        System.out.println(\"Region: \" + resourceGroup.regionName());\n+        System.out.println(\"Provisioning state: \" + resourceGroup.provisioningState());\n+        System.out.println(\"Tags: \" + resourceGroup.tags());\n+    }\n+\n+    private static void deleteAfterFailure(ResourceGroup resourceGroup) {\n+        try {\n+            System.err.printf(\"Cleaning up resource group '%s' after failure...%n\",\n+                resourceGroup.name());\n+            resourceGroup.manager().resourceGroups().deleteByName(resourceGroup.name());\n+        } catch (ClientAuthenticationException | ManagementException cleanupException) {\n+            System.err.printf(\"Cleanup failed for resource group '%s': %s%n\",\n+                resourceGroup.name(), cleanupException.getMessage());\n+        }\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(\n+                \"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static String getResourceGroupName() {\n+        String configuredName = System.getenv(\"AZURE_RESOURCE_GROUP_NAME\");\n+        if (configuredName != null \u0026\u0026 !configuredName.isBlank()) {\n+            return configuredName;\n+        }\n+        return \"java-sdk-rg-\" + UUID.randomUUID().toString().substring(0, 8);\n+    }\n+}\n*** Add File: README.md\n+# Azure Resource Group Manager\n+\n+This Maven application uses `DefaultAzureCredential` and the modern\n+`azure-resourcemanager` management-plane SDK to:\n+\n+1. Create a resource group in `eastus`.\n+2. List all resource groups in the subscription.\n+3. Retrieve and display the created resource group.\n+4. Add a `managed-by=azure-sdk-for-java` tag.\n+5. Delete the resource group.\n+\n+The required Maven dependencies are declared in `pom.xml`:\n+\n+```xml\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure.resourcemanager\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-resourcemanager\u003c/artifactId\u003e\n+    \u003cversion\u003e2.63.0\u003c/version\u003e\n+\u003c/dependency\u003e\n+\u003cdependency\u003e\n+    \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+    \u003cversion\u003e1.18.5\u003c/version\u003e\n+\u003c/dependency\u003e\n+```\n+\n+## Run\n+\n+Use Java 17 or later. Configure `DefaultAzureCredential` for your environment\n+(for example, Azure CLI credentials for local development or managed identity\n+in Azure), then set the subscription:\n+\n+```powershell\n+$env:AZURE_SUBSCRIPTION_ID = \"\u003csubscription-id\u003e\"\n+$env:AZURE_RESOURCE_GROUP_NAME = \"\u003coptional-resource-group-name\u003e\"\n+mvn compile exec:java\n+```\n+\n+If `AZURE_RESOURCE_GROUP_NAME` is omitted, the application generates a unique\n+name. The signed-in identity needs permission to read, create, update, and\n+delete resource groups in the subscription. The application deletes the\n+created resource group during normal execution and attempts cleanup if a later\n+operation fails.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/pom.xml
@@ -1,0 +1,49 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example.azure</groupId>
+    <artifactId>resource-group-manager</artifactId>
+    <version>1.0.0</version>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure-resourcemanager.version>2.63.0</azure-resourcemanager.version>
+        <azure-identity.version>1.18.5</azure-identity.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure.resourcemanager</groupId>
+            <artifactId>azure-resourcemanager</artifactId>
+            <version>${azure-resourcemanager.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure-identity.version}</version>
+        </dependency>
+    </dependencies>
+
+    <build>
+        <plugins>
+            <plugin>
+                <groupId>org.apache.maven.plugins</groupId>
+                <artifactId>maven-compiler-plugin</artifactId>
+                <version>3.14.1</version>
+            </plugin>

... (truncated)
```

### ✅ 14. powershell — 4494ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven Java application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example.azure:resource-group-manager >--------------
[INFO] Building resource-group-manager 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ resource-group-manager ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ resource-group-manager ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java:[5,33] cannot find symbol
  symbol:   class AzureProfile
  location: package com.azure.core.management
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java:[32,13] cannot find symbol
  symbol:   class AzureProfile
  location: class com.example.azure.Re
... (truncated)
```

### ✅ 15. apply_patch — 685ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azure/ResourceGroupManager.java\n@@\n import com.azure.core.exception.ClientAuthenticationException;\n import com.azure.core.management.AzureEnvironment;\n-import com.azure.core.management.AzureProfile;\n import com.azure.core.management.exception.ManagementException;\n+import com.azure.core.management.profile.AzureProfile;\n@@\n         String resourceGroupName = getResourceGroupName();\n \n+        AzureResourceManager azure = null;\n         ResourceGroup createdGroup = null;\n         boolean deleted = false;\n@@\n-            AzureResourceManager azure = AzureResourceManager\n+            azure = AzureResourceManager\n                 .authenticate(credential, profile)\n                 .withSubscription(subscriptionId);\n@@\n-            if (createdGroup != null \u0026\u0026 !deleted) {\n-                deleteAfterFailure(createdGroup);\n+            if (azure != null \u0026\u0026 createdGroup != null \u0026\u0026 !deleted) {\n+                deleteAfterFailure(azure, resourceGroupName);\n@@\n-    private static void deleteAfterFailure(ResourceGroup resourceGroup) {\n+    private static void deleteAfterFailure(\n+        AzureResourceManager azure,\n+        String resourceGroupName\n+    ) {\n         try {\n             System.err.printf(\"Cleaning up resource group '%s' after failure...%n\",\n-                resourceGroup.name());\n-            resourceGroup.manager().resourceGroups().deleteByName(resourceGroup.name());\n+                resourceGroupName);\n+            azure.resourceGroups().deleteByName(resourceGroupName);\n         } catch (ClientAuthenticationException | ManagementException cleanupException) {\n             System.err.printf(\"Cleanup failed for resource group '%s': %s%n\",\n-                resourceGroup.name(), cleanupException.getMessage());\n+                resourceGroupName, cleanupException.getMessage());\n         }\n     }\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
@@ -2,8 +2,8 @@
 
 import com.azure.core.exception.ClientAuthenticationException;
 import com.azure.core.management.AzureEnvironment;
-import com.azure.core.management.AzureProfile;
 import com.azure.core.management.exception.ManagementException;
+import com.azure.core.management.profile.AzureProfile;
 import com.azure.identity.CredentialUnavailableException;
 import com.azure.identity.DefaultAzureCredential;
 import com.azure.identity.DefaultAzureCredentialBuilder;
@@ -24,6 +24,7 @@
         String subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
         String resourceGroupName = getResourceGroupName();
 
+        AzureResourceManager azure = null;
         ResourceGroup createdGroup = null;
         boolean deleted = false;
 
@@ -35,7 +36,7 @@
                 subscriptionId,
                 AzureEnvironment.AZURE);
 
-            AzureResourceManager azure = AzureResourceManager
+            azure = AzureResourceManager
                 .authenticate(credential, profile)
                 .withSubscription(subscriptionId);
 
@@ -81,8 +82,8 @@
             System.err.printf("Azure Resource Manager request failed (status %d): %s
... (truncated)
```

### ✅ 16. powershell — 4804ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Recompile after SDK API fixes","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example.azure:resource-group-manager >--------------
[INFO] Building resource-group-manager 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ resource-group-manager ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ resource-group-manager ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java:[33,36] no suitable constructor found for AzureProfile(<nulltype>,<nulltype>,java.lang.String,com.azure.core.management.AzureEnvironment)
    constructor com.azure.core.management.profile.AzureProfile.AzureProfile(com.azure.core.models.AzureCloud) is not applicable
      (actual and formal argument lists differ in length)
    constructor com.azure.core.management.profile.AzureProfile.AzureProfile(j
... (truncated)
```

### ✅ 17. apply_patch — 721ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azure/ResourceGroupManager.java\n@@\n-            AzureProfile profile = new AzureProfile(\n-                null,\n-                null,\n-                subscriptionId,\n-                AzureEnvironment.AZURE);\n+            AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727/src/main/java/com/example/azure/ResourceGroupManager.java
@@ -30,11 +30,7 @@
 
         try {
             DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
-            AzureProfile profile = new AzureProfile(
-                null,
-                null,
-                subscriptionId,
-                AzureEnvironment.AZURE);
+            AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);
 
             azure = AzureResourceManager
                 .authenticate(credential, profile)

```

### ✅ 18. powershell — 4332ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile corrected resource manager sample","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] --------------< com.example.azure:resource-group-manager >--------------
[INFO] Building resource-group-manager 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ resource-group-manager ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-resource-manager-mp-java-rg-crud-java-azure-skills-azure-skill-mcp-microsoft-skill-4017631727\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ resource-group-manager ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 1 source file with javac [debug release 17] to target\classes
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
[INFO] Total time:  1.893 s
[INFO] Finished at: 2026-08-28T02:20:28+08:00
[INFO] ------------------------------------------------------------------------
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azure\ResourceGroupManager.java`

## Copilot Response

I’ll scaffold a local Maven sample and keep all Azure operations unexecuted; the application will require an explicit subscription ID and resource-group name at runtime.

The project is empty. I’m pinning compatible SDK versions from Maven metadata, then I’ll add the sample and compile it locally without invoking Azure.

The sample is in place. I’m compiling it now to catch any SDK signature or exception-hierarchy mismatches; this only resolves dependencies and compiles bytecode.

Implemented a complete Maven Java application using the modern Azure management-plane SDK.

- `pom.xml` — `azure-resourcemanager` 2.63.0 and `azure-identity` 1.18.5
- `src\main\java\com\example\azure\ResourceGroupManager.java` — creates, lists, retrieves, tags, and deletes the resource group using `DefaultAzureCredential`
- `README.md` — configuration and execution instructions
- Includes authentication, ARM request, cleanup, and missing-resource handling

Run with:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "<optional-resource-group-name>"
mvn compile exec:java
```

If no resource-group name is provided, the application generates a unique one. It attempts to delete the created group if a later operation fails.

## Reviewer Notes (LLM-as-Judge)

### Summary

The sole criterion is satisfied.

### Strengths

- No AutoCloseable Azure SDK client is left unclosed.

## Grader Results

- resource-group-crud.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (7/7)
      - Maven dependency for `azure-resourcemanager` and `azure-identity`: Pass
      - `AzureResourceManager.authenticate()` with credential and profile: Pass
      - `resourceGroups().define().withRegion().create()`: Pass
      - `resourceGroups().list()` iteration: Pass
      - `resourceGroups().getByName()` for details: Pass
      - Tag management via `update().withTag().apply()`: Pass
      - `resourceGroups().deleteByName()` for cleanup: Pass
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
  - Client Builder Pattern (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Fail
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Client Builder Pattern` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 10.0% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.0000** | **76.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id resource-manager-mp-java-rg-crud --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
