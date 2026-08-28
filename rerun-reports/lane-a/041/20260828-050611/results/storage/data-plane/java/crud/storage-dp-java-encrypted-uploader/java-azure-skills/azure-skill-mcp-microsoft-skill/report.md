# Evaluation Report: storage-dp-java-encrypted-uploader

**Config:** java-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 581.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-encrypted-uploader` |
| Config | java-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 31/32 |
| Duration | 581.1s |
| Timestamp | 2026-08-27T21:06:11Z |
| Files Generated | 11 |
| Event Count | 15806 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 337.9s |
| Review | 242.0s |
| **Total** | **581.1s** |

## Configuration

- **name:** java-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-storage-blob-java, azure-security-keyvault-keys-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=454390 out=10586 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. skill — 6024ms

**Input:**

```json
{"skill":"azure-security-keyvault-keys-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Key Vault Keys (Java)

Manage cryptographic keys and perform cryptographic operations in Azure Key Vault and Managed HSM.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-keys</artifactId>
    <version>4.9.0</version>
</dependency>
```

## Client Creation

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

// Key management client
KeyClient keyClient = new KeyClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildClient();

// Async client
KeyAsyncClient keyAsyncClient = new KeyClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(credential)
    .buildAsyncClient();

// Cryptography client (for encrypt/decrypt/sign/verify)
CryptographyClient cryptoClient = new CryptographyClientBuilder()
    .keyIdentifier("https://<vault-name>.vault.azure.net/keys/<key-name>/<key-version>")
    .credential(credential)
    .buildClient();
```

## Key T
... (truncated)
```

### ✅ 2. skill — 4763ms

**Input:**

```json
{"skill":"azure-storage-blob-java"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Storage Blob SDK for Java

Build blob storage applications using the Azure Storage Blob SDK for Java.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-storage-blob</artifactId>
    <version>12.33.0</version>
</dependency>
```

## Client Creation

### BlobServiceClient

```java
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

// With SAS token
BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .endpoint("<storage-account-url>")
    .sasToken("<sas-token>")
    .buildClient();

// With connection string
BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .connectionString("<connection-string>")
    .buildClient();
```

### With DefaultAzureCredential

```java
import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureIdentityEnvVars;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

// Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
TokenCredential credential = new DefaultAzureCredentialBuilder()
    .requireEnvVars(AzureIdentityEnvVars.AZURE_TOKEN_CREDENTIALS)
    .build();
// Or use a specific credential directly in production:
// See https://learn.microsoft.com/java/api/overview/azure/identity-readme?view=azure-java-stable#credential-classes
// TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

BlobServiceClient serviceClient = new BlobServiceClientBuilder()
    .endpoint("<storage-account-url>")
    .credential(credential)
    .buildClient();
```

### BlobContainerClient

```java
import com.azure.storage.blob.BlobContainerClient;

// From service client
BlobContainerClient containerClient = serviceClient.getBlobContainerClient("mycontainer");

// Direct construction
BlobContainerClient containerCli
... (truncated)
```

### ❌ 3. azure-get_azure_bestpractices (via azure) — 187412ms

**Input:**

```json
{"intent":"Generate a local Java 17 Maven sample using Azure Blob Storage and Azure Key Vault Keys with managed identity, sync and async clients, envelope encryption, metadata, and robust service error handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 4. glob — 187407ms

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-keys-java","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java"],"pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-java\skills\azure-storage-blob-java\references\examples.md
C:\Users\weidxu\.hyoka\cache\repos\microsoft\skills\default\.github\plugins\azure-sdk-java\skills\azure-security-keyvault-keys-java\references\examples.md
```

### ✅ 5. glob — 186177ms

**Input:**

```json
{"paths":".","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. view — 3762ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-keys-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-keys-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-keys-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-keys-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-security-keyvault-keys-java/references/examples.md
@@ -1,515 +1,515 @@
 # Azure Key Vault Keys SDK for Java - Examples
 
 Comprehensive code examples for the Azure Key Vault Keys SDK for Java.
 
 ## Table of Contents
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Creating Keys](#creating-keys)
 - [Getting and Listing Keys](#getting-and-listing-keys)
 - [Updating Key Properties](#updating-key-properties)
 - [Deleting and Recovering Keys](#deleting-and-recovering-keys)
 - [Key Rotation](#key-rotation)
 - [Cryptographic Operations](#cryptographic-operations)
 - [Async Client Patterns](#async-client-patterns)
 - [Error Handling](#error-handling)
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-security-keyvault-keys</artifactId>
     <version>4.9.0</version>
 </dependency>
 
 <!-- Required for authentication -->
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-identity</artifactId>
     <version>1.14.0</version>
 </dependency>
 ```
 
 ## Client Creation
 
 ### Sync KeyClient
 
 ```java
 import com.azure.identity.DefaultAzureCredentialBuilder;
 import com.azure.security.keyvault.keys.KeyClient;
 import com.azure.security.keyvault.keys.KeyClientBuilder;
 
 KeyClient keyClient = new KeyClientBuilder()
     .vaultUrl("<your-key-vault-url>")
     .
... (truncated)
```

### ✅ 7. view — 3437ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java\\references\\examples.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-java/skills/azure-storage-blob-java/references/examples.md
@@ -1,411 +1,411 @@
 # Azure Storage Blob Java SDK - Examples
 
 Comprehensive code examples for the Azure Storage Blob SDK for Java.
 
 ## Table of Contents
 
 - [Maven Dependency](#maven-dependency)
 - [Client Creation](#client-creation)
 - [Container Operations](#container-operations)
 - [Upload Blobs](#upload-blobs)
 - [Download Blobs](#download-blobs)
 - [List Blobs](#list-blobs)
 - [SAS Token Generation](#sas-token-generation)
 - [Error Handling](#error-handling)
 
 ---
 
 ## Maven Dependency
 
 ```xml
 <dependency>
     <groupId>com.azure</groupId>
     <artifactId>azure-storage-blob</artifactId>
     <version>12.33.0</version>
 </dependency>
 ```
 
 Or use the BOM:
 
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
         <artifactId>azure-storage-blob</artifactId>
     </dependency>
 </dependencies>
 ```
 
 ---
 
 ## Client Creation
 
 ### Using Shared Key Credential
 
 ```java
 import com.azure.storage.blob.BlobServiceClient;
 import com.azure.storage.blob.BlobSe
... (truncated)
```

### ✅ 8. apply_patch — 882ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eencrypted-blob-uploader\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eEncrypted Azure Blob Uploader\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cazure.identity.version\u003e1.14.0\u003c/azure.identity.version\u003e\n+        \u003cazure.keyvault.keys.version\u003e4.9.0\u003c/azure.keyvault.keys.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.33.0\u003c/azure.storage.blob.version\u003e\n+        \u003cslf4j.version\u003e2.0.16\u003c/slf4j.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-keys\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.keyvault.keys.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e${slf4j.version}\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.13.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.encryptedblob.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/encryptedblob/EnvelopeEncryptionException.java\n+package com.example.encryptedblob;\n+\n+public final class EnvelopeEncryptionException extends RuntimeException {\n+    public EnvelopeEncryptionException(String message) {\n+        super(message);\n+    }\n+\n+    public EnvelopeEncryptionException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyEnvelope.java\n+package com.example.encryptedblob;\n+\n+import java.util.Arrays;\n+\n+record ProtectedDataKey(String keyId, byte[] wrappedKey) {\n+    ProtectedDataKey {\n+        if (keyId == null || keyId.isBlank()) {\n+            throw new IllegalArgumentException(\"keyId must not be blank\");\n+        }\n+        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+\n+    @Override\n+    public byte[] wrappedKey() {\n+        return Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+}\n+\n+final class DataKeyEnvelope implements AutoCloseable {\n+    private final byte[] plaintextKey;\n+    private final ProtectedDataKey protectedKey;\n+\n+    DataKeyEnvelope(byte[] plaintextKey, ProtectedDataKey protectedKey) {\n+        this.plaintextKey = plaintextKey;\n+        this.protectedKey = protectedKey;\n+    }\n+\n+    byte[] plaintextKey() {\n+        return plaintextKey;\n+    }\n+\n+    ProtectedDataKey protectedKey() {\n+        return protectedKey;\n+    }\n+\n+    @Override\n+    public void close() {\n+        Arrays.fill(plaintextKey, (byte) 0);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagement.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import com.azure.security.keyvault.keys.models.KeyVaultKey;\n+\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+\n+public final class KeyManagement {\n+    private static final int DATA_KEY_BYTES = 32;\n+    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;\n+\n+    private final KeyClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public KeyManagement(KeyClient keyClient, TokenCredential credential, String keyName) {\n+        this.keyClient = keyClient;\n+        this.credential = credential;\n+        this.keyName = keyName;\n+        this.secureRandom = new SecureRandom();\n+    }\n+\n+    DataKeyEnvelope generateAndWrapKey() {\n+        byte[] dataKey = new byte[DATA_KEY_BYTES];\n+        secureRandom.nextBytes(dataKey);\n+\n+        try {\n+            KeyVaultKey vaultKey = keyClient.getKey(keyName);\n+            String versionedKeyId = vaultKey.getId();\n+            CryptographyClient cryptographyClient = cryptoClient(versionedKeyId);\n+            byte[] wrappedKey = cryptographyClient.wrapKey(WRAP_ALGORITHM, dataKey).getEncryptedKey();\n+            return new DataKeyEnvelope(dataKey, new ProtectedDataKey(versionedKeyId, wrappedKey));\n+        } catch (HttpResponseException exception) {\n+            Arrays.fill(dataKey, (byte) 0);\n+            throw keyVaultFailure(\"wrap a data encryption key\", exception);\n+        } catch (RuntimeException exception) {\n+            Arrays.fill(dataKey, (byte) 0);\n+            throw exception;\n+        }\n+    }\n+\n+    byte[] unwrapKey(ProtectedDataKey protectedKey) {\n+        try {\n+            return cryptoClient(protectedKey.keyId())\n+                .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())\n+                .getKey();\n+        } catch (HttpResponseException exception) {\n+            throw keyVaultFailure(\"unwrap the data encryption key\", exception);\n+        }\n+    }\n+\n+    private CryptographyClient cryptoClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+            .keyIdentifier(keyId)\n+            .credential(credential)\n+            .buildClient();\n+    }\n+\n+    private static EnvelopeEncryptionException keyVaultFailure(String operation, HttpResponseException exception) {\n+        int status = exception.getResponse() == null ? -1 : exception.getResponse().getStatusCode();\n+        return new EnvelopeEncryptionException(\n+            \"Key Vault could not \" + operation + \" (HTTP \" + status\n+                + \"). Verify that the stored key version is enabled and permits wrapKey/unwrapKey.\",\n+            exception);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AsyncKeyManagement.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import reactor.core.publisher.Mono;\n+\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+\n+public final class AsyncKeyManagement {\n+    private static final int DATA_KEY_BYTES = 32;\n+    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;\n+\n+    private final KeyAsyncClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public AsyncKeyManagement(KeyAsyncClient keyClient, TokenCredential credential, String keyName) {\n+        this.keyClient = keyClient;\n+        this.credential = credential;\n+        this.keyName = keyName;\n+        this.secureRandom = new SecureRandom();\n+    }\n+\n+    Mono\u003cDataKeyEnvelope\u003e generateAndWrapKey() {\n+        return Mono.defer(() -\u003e {\n+            byte[] dataKey = new byte[DATA_KEY_BYTES];\n+            secureRandom.nextBytes(dataKey);\n+\n+            return keyClient.getKey(keyName)\n+                .flatMap(key -\u003e {\n+                    String versionedKeyId = key.getId();\n+                    return cryptoClient(versionedKeyId).wrapKey(WRAP_ALGORITHM, dataKey)\n+                        .map(result -\u003e new DataKeyEnvelope(\n+                            dataKey,\n+                            new ProtectedDataKey(versionedKeyId, result.getEncryptedKey())));\n+                })\n+                .doOnError(ignored -\u003e Arrays.fill(dataKey, (byte) 0))\n+                .doOnCancel(() -\u003e Arrays.fill(dataKey, (byte) 0))\n+                .onErrorMap(HttpResponseException.class, exception -\u003e\n+                    keyVaultFailure(\"wrap a data encryption key\", exception));\n+        });\n+    }\n+\n+    Mono\u003cbyte[]\u003e unwrapKey(ProtectedDataKey protectedKey) {\n+        return cryptoClient(protectedKey.keyId())\n+            .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())\n+            .map(result -\u003e result.getKey())\n+            .onErrorMap(HttpResponseException.class, exception -\u003e\n+                keyVaultFailure(\"unwrap the data encryption key\", exception));\n+    }\n+\n+    private CryptographyAsyncClient cryptoClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+            .keyIdentifier(keyId)\n+            .credential(credential)\n+            .buildAsyncClient();\n+    }\n+\n+    private static EnvelopeEncryptionException keyVaultFailure(String operation, HttpResponseException exception) {\n+        int status = exception.getResponse() == null ? -1 : exception.getResponse().getStatusCode();\n+        return new EnvelopeEncryptionException(\n+            \"Key Vault could not \" + operation + \" (HTTP \" + status\n+                + \"). Verify that the stored key version is enabled and permits wrapKey/unwrapKey.\",\n+            exception);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n+\n+import javax.crypto.AEADBadTagException;\n+import javax.crypto.Cipher;\n+import javax.crypto.spec.GCMParameterSpec;\n+import javax.crypto.spec.SecretKeySpec;\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.security.GeneralSecurityException;\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+import java.util.Base64;\n+import java.util.Map;\n+\n+public final class EncryptedBlobClient {\n+    private static final String CONTENT_ALGORITHM = \"AES/GCM/NoPadding\";\n+    private static final int GCM_TAG_BITS = 128;\n+    private static final int IV_BYTES = 12;\n+\n+    private final BlobContainerClient containerClient;\n+    private final KeyManagement keyManagement;\n+    private final SecureRandom secureRandom;\n+\n+    public EncryptedBlobClient(BlobContainerClient containerClient, KeyManagement keyManagement) {\n+        this.containerClient = containerClient;\n+        this.keyManagement = keyManagement;\n+        this.secureRandom = new SecureRandom();\n+    }\n+\n+    public UploadResult upload(Path source, String blobName) {\n+        try {\n+            return upload(Files.readAllBytes(source), blobName);\n+        } catch (IOException exception) {\n+            throw new EnvelopeEncryptionException(\"Could not read source file: \" + source, exception);\n+        }\n+    }\n+\n+    public UploadResult upload(byte[] plaintext, String blobName) {\n+        try (DataKeyEnvelope envelope = keyManagement.generateAndWrapKey()) {\n+            byte[] iv = new byte[IV_BYTES];\n+            secureRandom.nextBytes(iv);\n+            byte[] ciphertext = encrypt(plaintext, envelope.plaintextKey(), iv);\n+            ProtectedDataKey protectedKey = envelope.protectedKey();\n+\n+            Map\u003cString, String\u003e metadata = BlobEncryptionMetadata.create(protectedKey, iv);\n+            containerClient.getBlobClient(blobName).uploadWithResponse(\n+                new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))\n+                    .setMetadata(metadata),\n+                null,\n+                null);\n+\n+            return new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey());\n+        } catch (BlobStorageException exception) {\n+            throw blobFailure(\"upload\", blobName, exception);\n+        }\n+    }\n+\n+    public byte[] download(String blobName) {\n+        try {\n+            var blobClient = containerClient.getBlobClient(blobName);\n+            Map\u003cString, String\u003e metadata = blobClient.getProperties().getMetadata();\n+            BlobEncryptionMetadata encryptionMetadata = BlobEncryptionMetadata.parse(metadata);\n+            byte[] ciphertext = blobClient.downloadContent().toBytes();\n+            byte[] dataKey = keyManagement.unwrapKey(encryptionMetadata.protectedKey());\n+\n+            try {\n+                return decrypt(ciphertext, dataKey, encryptionMetadata.iv());\n+            } finally {\n+                Arrays.fill(dataKey, (byte) 0);\n+            }\n+        } catch (BlobStorageException exception) {\n+            throw blobFailure(\"download\", blobName, exception);\n+        }\n+    }\n+\n+    public void download(String blobName, Path destination) {\n+        byte[] plaintext = download(blobName);\n+        try {\n+            Files.write(destination, plaintext);\n+        } catch (IOException exception) {\n+            throw new EnvelopeEncryptionException(\"Could not write decrypted file: \" + destination, exception);\n+        } finally {\n+            Arrays.fill(plaintext, (byte) 0);\n+        }\n+    }\n+\n+    private static byte[] encrypt(byte[] plaintext, byte[] dataKey, byte[] iv) {\n+        try {\n+            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);\n+            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(dataKey, \"AES\"), new GCMParameterSpec(GCM_TAG_BITS, iv));\n+            return cipher.doFinal(plaintext);\n+        } catch (GeneralSecurityException exception) {\n+            throw new EnvelopeEncryptionException(\"Local AES-GCM encryption failed.\", exception);\n+        }\n+    }\n+\n+    static byte[] decrypt(byte[] ciphertext, byte[] dataKey, byte[] iv) {\n+        try {\n+            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);\n+            cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(dataKey, \"AES\"), new GCMParameterSpec(GCM_TAG_BITS, iv));\n+            return cipher.doFinal(ciphertext);\n+        } catch (AEADBadTagException exception) {\n+            throw new EnvelopeEncryptionException(\n+                \"Ciphertext authentication failed; the blob data or encryption metadata was modified.\", exception);\n+        } catch (GeneralSecurityException exception) {\n+            throw new EnvelopeEncryptionException(\"Local AES-GCM decryption failed.\", exception);\n+        }\n+    }\n+\n+    private static EnvelopeEncryptionException blobFailure(\n+        String operation,\n+        String blobName,\n+        BlobStorageException exception\n+    ) {\n+        String detail = exception.getStatusCode() == 404\n+            ? \"The blob or container does not exist.\"\n+            : \"Storage returned \" + exception.getErrorCode() + \".\";\n+        return new EnvelopeEncryptionException(\n+            \"Could not \" + operation + \" blob '\" + blobName + \"' (HTTP \"\n+                + exception.getStatusCode() + \"). \" + detail,\n+            exception);\n+    }\n+\n+    public record UploadResult(String keyId, byte[] wrappedKey) {\n+        public UploadResult {\n+            wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);\n+        }\n+\n+        @Override\n+        public byte[] wrappedKey() {\n+            return Arrays.copyOf(wrappedKey, wrappedKey.length);\n+        }\n+\n+        public String wrappedKeyBase64() {\n+            return Base64.getEncoder().encodeToString(wrappedKey);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/BlobEncryptionMetadata.java\n+package com.example.encryptedblob;\n+\n+import java.util.Base64;\n+import java.util.HashMap;\n+import java.util.Map;\n+\n+record BlobEncryptionMetadata(ProtectedDataKey protectedKey, byte[] iv) {\n+    private static final String VERSION = \"1\";\n+    private static final String CONTENT_ALGORITHM = \"AES256GCM\";\n+    private static final String WRAP_ALGORITHM = \"RSAOAEP256\";\n+\n+    static Map\u003cString, String\u003e create(ProtectedDataKey protectedKey, byte[] iv) {\n+        Map\u003cString, String\u003e metadata = new HashMap\u003c\u003e();\n+        metadata.put(\"encryptionversion\", VERSION);\n+        metadata.put(\"contentalgorithm\", CONTENT_ALGORITHM);\n+        metadata.put(\"wrapalgorithm\", WRAP_ALGORITHM);\n+        metadata.put(\"keyid\", protectedKey.keyId());\n+        metadata.put(\"wrappedkey\", Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));\n+        metadata.put(\"iv\", Base64.getEncoder().encodeToString(iv));\n+        return metadata;\n+    }\n+\n+    static BlobEncryptionMetadata parse(Map\u003cString, String\u003e metadata) {\n+        requireValue(metadata, \"encryptionversion\", VERSION);\n+        requireValue(metadata, \"contentalgorithm\", CONTENT_ALGORITHM);\n+        requireValue(metadata, \"wrapalgorithm\", WRAP_ALGORITHM);\n+\n+        String keyId = require(metadata, \"keyid\");\n+        try {\n+            byte[] wrappedKey = Base64.getDecoder().decode(require(metadata, \"wrappedkey\"));\n+            byte[] iv = Base64.getDecoder().decode(require(metadata, \"iv\"));\n+            if (iv.length != 12) {\n+                throw new EnvelopeEncryptionException(\"Invalid AES-GCM IV length in blob metadata.\");\n+            }\n+            return new BlobEncryptionMetadata(new ProtectedDataKey(keyId, wrappedKey), iv);\n+        } catch (IllegalArgumentException exception) {\n+            throw new EnvelopeEncryptionException(\"Blob encryption metadata contains invalid Base64.\", exception);\n+        }\n+    }\n+\n+    private static String require(Map\u003cString, String\u003e metadata, String name) {\n+        String value = metadata.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new EnvelopeEncryptionException(\"Blob is missing required encryption metadata: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static void requireValue(Map\u003cString, String\u003e metadata, String name, String expected) {\n+        String actual = require(metadata, name);\n+        if (!expected.equals(actual)) {\n+            throw new EnvelopeEncryptionException(\n+                \"Unsupported blob encryption metadata \" + name + \": \" + actual);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n+import reactor.core.publisher.Mono;\n+import reactor.core.scheduler.Schedulers;\n+\n+import javax.crypto.Cipher;\n+import javax.crypto.spec.GCMParameterSpec;\n+import javax.crypto.spec.SecretKeySpec;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.security.GeneralSecurityException;\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+import java.util.Map;\n+\n+public final class AsyncEncryptedBlobClient {\n+    private static final String CONTENT_ALGORITHM = \"AES/GCM/NoPadding\";\n+    private static final int GCM_TAG_BITS = 128;\n+    private static final int IV_BYTES = 12;\n+\n+    private final BlobContainerAsyncClient containerClient;\n+    private final AsyncKeyManagement keyManagement;\n+    private final SecureRandom secureRandom;\n+\n+    public AsyncEncryptedBlobClient(\n+        BlobContainerAsyncClient containerClient,\n+        AsyncKeyManagement keyManagement\n+    ) {\n+        this.containerClient = containerClient;\n+        this.keyManagement = keyManagement;\n+        this.secureRandom = new SecureRandom();\n+    }\n+\n+    public Mono\u003cEncryptedBlobClient.UploadResult\u003e upload(Path source, String blobName) {\n+        return Mono.fromCallable(() -\u003e Files.readAllBytes(source))\n+            .subscribeOn(Schedulers.boundedElastic())\n+            .onErrorMap(exception -\u003e\n+                new EnvelopeEncryptionException(\"Could not read source file: \" + source, exception))\n+            .flatMap(plaintext -\u003e upload(plaintext, blobName));\n+    }\n+\n+    public Mono\u003cEncryptedBlobClient.UploadResult\u003e upload(byte[] plaintext, String blobName) {\n+        return Mono.usingWhen(\n+                keyManagement.generateAndWrapKey(),\n+                envelope -\u003e Mono.defer(() -\u003e {\n+                    byte[] iv = new byte[IV_BYTES];\n+                    secureRandom.nextBytes(iv);\n+                    byte[] ciphertext = encrypt(plaintext, envelope.plaintextKey(), iv);\n+                    ProtectedDataKey protectedKey = envelope.protectedKey();\n+                    Map\u003cString, String\u003e metadata = BlobEncryptionMetadata.create(protectedKey, iv);\n+\n+                    return containerClient.getBlobAsyncClient(blobName)\n+                        .uploadWithResponse(\n+                            new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))\n+                                .setMetadata(metadata))\n+                        .map(ignored -\u003e new EncryptedBlobClient.UploadResult(\n+                            protectedKey.keyId(),\n+                            protectedKey.wrappedKey()));\n+                }),\n+                envelope -\u003e Mono.fromRunnable(envelope::close),\n+                (envelope, ignored) -\u003e Mono.fromRunnable(envelope::close),\n+                envelope -\u003e Mono.fromRunnable(envelope::close))\n+            .onErrorMap(BlobStorageException.class, exception -\u003e\n+                blobFailure(\"upload\", blobName, exception));\n+    }\n+\n+    public Mono\u003cbyte[]\u003e download(String blobName) {\n+        var blobClient = containerClient.getBlobAsyncClient(blobName);\n+        return blobClient.getProperties()\n+            .flatMap(properties -\u003e {\n+                BlobEncryptionMetadata metadata = BlobEncryptionMetadata.parse(properties.getMetadata());\n+                return blobClient.downloadContent()\n+                    .flatMap(content -\u003e keyManagement.unwrapKey(metadata.protectedKey())\n+                        .flatMap(dataKey -\u003e Mono.using(\n+                            () -\u003e dataKey,\n+                            key -\u003e Mono.fromCallable(() -\u003e\n+                                EncryptedBlobClient.decrypt(content.toBytes(), key, metadata.iv())),\n+                            key -\u003e Arrays.fill(key, (byte) 0))));\n+            })\n+            .onErrorMap(BlobStorageException.class, exception -\u003e\n+                blobFailure(\"download\", blobName, exception));\n+    }\n+\n+    public Mono\u003cVoid\u003e download(String blobName, Path destination) {\n+        return download(blobName)\n+            .flatMap(plaintext -\u003e Mono.fromRunnable(() -\u003e {\n+                try {\n+                    Files.write(destination, plaintext);\n+                } catch (Exception exception) {\n+                    throw new EnvelopeEncryptionException(\n+                        \"Could not write decrypted file: \" + destination, exception);\n+                } finally {\n+                    Arrays.fill(plaintext, (byte) 0);\n+                }\n+            }).subscribeOn(Schedulers.boundedElastic()));\n+    }\n+\n+    private static byte[] encrypt(byte[] plaintext, byte[] dataKey, byte[] iv) {\n+        try {\n+            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);\n+            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(dataKey, \"AES\"), new GCMParameterSpec(GCM_TAG_BITS, iv));\n+            return cipher.doFinal(plaintext);\n+        } catch (GeneralSecurityException exception) {\n+            throw new EnvelopeEncryptionException(\"Local AES-GCM encryption failed.\", exception);\n+        }\n+    }\n+\n+    private static EnvelopeEncryptionException blobFailure(\n+        String operation,\n+        String blobName,\n+        BlobStorageException exception\n+    ) {\n+        String detail = exception.getStatusCode() == 404\n+            ? \"The blob or container does not exist.\"\n+            : \"Storage returned \" + exception.getErrorCode() + \".\";\n+        return new EnvelopeEncryptionException(\n+            \"Could not \" + operation + \" blob '\" + blobName + \"' (HTTP \"\n+                + exception.getStatusCode() + \"). \" + detail,\n+            exception);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AzureConfiguration.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.KeyClientBuilder;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+public final class AzureConfiguration {\n+    private final TokenCredential credential;\n+    private final BlobContainerClient blobContainerClient;\n+    private final BlobContainerAsyncClient blobContainerAsyncClient;\n+    private final KeyClient keyClient;\n+    private final KeyAsyncClient keyAsyncClient;\n+    private final String keyName;\n+\n+    private AzureConfiguration(\n+        TokenCredential credential,\n+        BlobContainerClient blobContainerClient,\n+        BlobContainerAsyncClient blobContainerAsyncClient,\n+        KeyClient keyClient,\n+        KeyAsyncClient keyAsyncClient,\n+        String keyName\n+    ) {\n+        this.credential = credential;\n+        this.blobContainerClient = blobContainerClient;\n+        this.blobContainerAsyncClient = blobContainerAsyncClient;\n+        this.keyClient = keyClient;\n+        this.keyAsyncClient = keyAsyncClient;\n+        this.keyName = keyName;\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        String blobEndpoint = requiredEnvironment(\"AZURE_STORAGE_BLOB_ENDPOINT\");\n+        String containerName = requiredEnvironment(\"AZURE_STORAGE_CONTAINER\");\n+        String vaultEndpoint = requiredEnvironment(\"AZURE_KEY_VAULT_ENDPOINT\");\n+        String keyName = requiredEnvironment(\"AZURE_KEY_VAULT_KEY_NAME\");\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        String clientId = System.getenv(\"AZURE_CLIENT_ID\");\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            credentialBuilder.clientId(clientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()\n+            .endpoint(blobEndpoint)\n+            .credential(credential);\n+        BlobServiceClient blobServiceClient = blobBuilder.buildClient();\n+        BlobServiceAsyncClient blobServiceAsyncClient = blobBuilder.buildAsyncClient();\n+\n+        KeyClientBuilder keyBuilder = new KeyClientBuilder()\n+            .vaultUrl(vaultEndpoint)\n+            .credential(credential);\n+\n+        return new AzureConfiguration(\n+            credential,\n+            blobServiceClient.getBlobContainerClient(containerName),\n+            blobServiceAsyncClient.getBlobContainerAsyncClient(containerName),\n+            keyBuilder.buildClient(),\n+            keyBuilder.buildAsyncClient(),\n+            keyName);\n+    }\n+\n+    public TokenCredential credential() {\n+        return credential;\n+    }\n+\n+    public BlobContainerClient blobContainerClient() {\n+        return blobContainerClient;\n+    }\n+\n+    public BlobContainerAsyncClient blobContainerAsyncClient() {\n+        return blobContainerAsyncClient;\n+    }\n+\n+    public KeyClient keyClient() {\n+        return keyClient;\n+    }\n+\n+    public KeyAsyncClient keyAsyncClient() {\n+        return keyAsyncClient;\n+    }\n+\n+    public String keyName() {\n+        return keyName;\n+    }\n+\n+    private static String requiredEnvironment(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/Main.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+\n+public final class Main {\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws Exception {\n+        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();\n+        Path workDirectory = Files.createTempDirectory(\"encrypted-blob-demo-\");\n+\n+        runSyncDemo(configuration, workDirectory);\n+        runAsyncDemo(configuration, workDirectory);\n+    }\n+\n+    private static void runSyncDemo(AzureConfiguration configuration, Path workDirectory) throws Exception {\n+        Path source = workDirectory.resolve(\"sync-source.txt\");\n+        Path destination = workDirectory.resolve(\"sync-downloaded.txt\");\n+        Files.writeString(source, \"Hello from the synchronous encrypted uploader.\", StandardCharsets.UTF_8);\n+\n+        KeyManagement keyManagement = new KeyManagement(\n+            configuration.keyClient(),\n+            configuration.credential(),\n+            configuration.keyName());\n+        EncryptedBlobClient client = new EncryptedBlobClient(\n+            configuration.blobContainerClient(),\n+            keyManagement);\n+\n+        EncryptedBlobClient.UploadResult result = client.upload(source, \"demo/sync-encrypted.bin\");\n+        client.download(\"demo/sync-encrypted.bin\", destination);\n+\n+        System.out.println(\"Sync vault key ID: \" + result.keyId());\n+        System.out.println(\"Sync wrapped DEK (base64): \" + result.wrappedKeyBase64());\n+        System.out.println(\"Sync decrypted output: \"\n+            + Files.readString(destination, StandardCharsets.UTF_8));\n+    }\n+\n+    private static void runAsyncDemo(AzureConfiguration configuration, Path workDirectory) throws Exception {\n+        Path source = workDirectory.resolve(\"async-source.txt\");\n+        Path destination = workDirectory.resolve(\"async-downloaded.txt\");\n+        Files.writeString(source, \"Hello from the asynchronous encrypted uploader.\", StandardCharsets.UTF_8);\n+\n+        AsyncKeyManagement keyManagement = new AsyncKeyManagement(\n+            configuration.keyAsyncClient(),\n+            configuration.credential(),\n+            configuration.keyName());\n+        AsyncEncryptedBlobClient client = new AsyncEncryptedBlobClient(\n+            configuration.blobContainerAsyncClient(),\n+            keyManagement);\n+\n+        EncryptedBlobClient.UploadResult result = client\n+            .upload(source, \"demo/async-encrypted.bin\")\n+            .flatMap(upload -\u003e client.download(\"demo/async-encrypted.bin\", destination).thenReturn(upload))\n+            .block();\n+\n+        if (result == null) {\n+            throw new IllegalStateException(\"The asynchronous round trip completed without an upload result.\");\n+        }\n+\n+        System.out.println(\"Async vault key ID: \" + result.keyId());\n+        System.out.println(\"Async wrapped DEK (base64): \" + result.wrappedKeyBase64());\n+        System.out.println(\"Async decrypted output: \"\n+            + Files.readString(destination, StandardCharsets.UTF_8));\n+    }\n+}\n*** Add File: README.md\n+# Encrypted Azure Blob Uploader\n+\n+Java 17 sample that encrypts files locally with a fresh AES-256-GCM data encryption key (DEK), wraps that DEK with a versioned RSA key in Azure Key Vault, and stores only ciphertext and the wrapped DEK in Azure Blob Storage.\n+\n+## Prerequisites\n+\n+- Java 17 and Maven 3.9+\n+- An existing blob container\n+- An existing RSA or RSA-HSM Key Vault key enabled for `wrapKey` and `unwrapKey`\n+- A managed identity with `Storage Blob Data Contributor` on the container and permission to get, wrap, and unwrap the Key Vault key\n+\n+Set these environment variables:\n+\n+```text\n+AZURE_STORAGE_BLOB_ENDPOINT=https://\u003caccount\u003e.blob.core.windows.net\n+AZURE_STORAGE_CONTAINER=\u003cexisting-container\u003e\n+AZURE_KEY_VAULT_ENDPOINT=https://\u003cvault\u003e.vault.azure.net\n+AZURE_KEY_VAULT_KEY_NAME=\u003cexisting-rsa-key\u003e\n+AZURE_CLIENT_ID=\u003coptional-user-assigned-managed-identity-client-id\u003e\n+```\n+\n+Run the demo from an Azure-hosted environment with that managed identity:\n+\n+```text\n+mvn compile exec:java\n+```\n+\n+The demo performs synchronous and asynchronous round trips. Blob metadata contains the versioned Key Vault key ID, RSA-OAEP-256-wrapped DEK, AES-GCM IV, and algorithm identifiers. The plaintext DEK exists only in process memory and is overwritten after each operation.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/pom.xml
@@ -1,0 +1,63 @@
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
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <maven.compiler.release>17</maven.compiler.release>
+        <azure.identity.version>1.14.0</azure.identity.version>
+        <azure.keyvault.keys.version>4.9.0</azure.keyvault.keys.version>
+        <azure.storage.blob.version>12.33.0</azure.storage.blob.version>
+        <slf4j.version>2.0.16</slf4j.version>
+    </properties>
+
+    <dependencies>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-identity</artifactId>
+            <version>${azure.identity.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-security-keyvault-keys</artifactId>
+            <version>${azure.keyvault.keys.version}</version>
+        </dependency>
+        <dependency>
+            <groupId>com.azure</groupId>
+            <artifactId>azure-storage-blob</artif
... (truncated)
```

### ✅ 9. powershell — 4606ms

**Input:**

```json
{"command":"mvn --no-transfer-progress compile","description":"Compile the Maven project","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:encrypted-blob-uploader >-----------------
[INFO] Building Encrypted Azure Blob Uploader 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460\src\main\resources
[INFO] 
[INFO] --- compiler:3.13.0:compile (default-compile) @ encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java:[90,21] incompatible types: inference variable R has incompatible bounds
    equality constraints: java.lang.Void
    lower bounds: java.lang.Object
[INFO] 1 error
[INFO] -------------------------------------------------------------
[INFO] ------------------------------------------------------------------------
[INFO] BUILD FAILURE
[INFO] ----------------------------------
... (truncated)
```

### ✅ 10. view — 631ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460\\src\\main\\java\\com\\example\\encryptedblob\\AsyncEncryptedBlobClient.java","view_range":[80,107]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
@@ -80,28 +80,28 @@
                             key -> Mono.fromCallable(() ->
                                 EncryptedBlobClient.decrypt(content.toBytes(), key, metadata.iv())),
                             key -> Arrays.fill(key, (byte) 0))));
             })
             .onErrorMap(BlobStorageException.class, exception ->
                 blobFailure("download", blobName, exception));
     }
 
     public Mono<Void> download(String blobName, Path destination) {
         return download(blobName)
             .flatMap(plaintext -> Mono.fromRunnable(() -> {
                 try {
                     Files.write(destination, plaintext);
                 } catch (Exception exception) {
                     throw new EnvelopeEncryptionException(
                         "Could not write decrypted file: " + destination, exception);
                 } finally {
                     Arrays.fill(plaintext, (byte) 0);
                 }
             }).subscribeOn(Schedulers.boundedElastic()));
     }
 
     private static byte[] encrypt(byte[] plaintext, byte[] dataKey, byte[] iv) {
         try {

... (truncated)
```

### ✅ 11. apply_patch — 683ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java\n@@\n-            .flatMap(plaintext -\u003e Mono.fromRunnable(() -\u003e {\n+            .flatMap(plaintext -\u003e Mono.\u003cVoid\u003efromRunnable(() -\u003e {\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460/src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java
@@ -87,7 +87,7 @@
 
     public Mono<Void> download(String blobName, Path destination) {
         return download(blobName)
-            .flatMap(plaintext -> Mono.fromRunnable(() -> {
+            .flatMap(plaintext -> Mono.<Void>fromRunnable(() -> {
                 try {
                     Files.write(destination, plaintext);
                 } catch (Exception exception) {

```

### ✅ 12. powershell — 5120ms

**Input:**

```json
{"command":"mvn --no-transfer-progress package -DskipTests","description":"Build the Maven package","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] ----------------< com.example:encrypted-blob-uploader >-----------------
[INFO] Building Encrypted Azure Blob Uploader 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460\src\main\resources
[INFO] 
[INFO] --- compiler:3.13.0:compile (default-compile) @ encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 9 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-skill-2611529460\src\test\resources
[INFO] 
[INFO] --- compiler:3.13.0:testCompile (default-testCompile) @ encrypted-blob-uploader ---
[INFO] No sources to compile
[INFO] 
[INFO] --- surefire:3.2.5:test (default-test) @ encrypted-blob-uploader ---
[INFO] Tests are skipped.
[INFO] 
[INFO] --- jar:3.4.1:jar (default-jar) @ encrypted-blob-uploader ---
[INFO] Building jar: C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-azure-skill-mcp-microsoft-
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\encryptedblob\AsyncEncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\AsyncKeyManagement.java`
- `src\main\java\com\example\encryptedblob\AzureConfiguration.java`
- `src\main\java\com\example\encryptedblob\BlobEncryptionMetadata.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\EnvelopeEncryptionException.java`
- `src\main\java\com\example\encryptedblob\KeyEnvelope.java`
- `src\main\java\com\example\encryptedblob\KeyManagement.java`
- `src\main\java\com\example\encryptedblob\Main.java`

## Copilot Response

I’ll implement versioned RSA-OAEP key wrapping with local AES-256-GCM, atomic blob metadata, and explicit zeroization of in-memory DEKs.

Created a complete Java 17 Maven project implementing:

- AES-256-GCM client-side file encryption with per-upload DEKs.
- RSA-OAEP-256 DEK wrapping/unwrapping through versioned Azure Key Vault Keys.
- Synchronous and Reactor-based asynchronous implementations.
- Atomic storage of ciphertext and cryptographic metadata in Blob Storage.
- Shared `ManagedIdentityCredential` configuration for all Azure clients.
- Service-specific error handling for missing blobs, disabled keys, permissions, and tampered ciphertext.
- A `Main` demo running both sync and async round trips and printing the key ID, wrapped DEK, and decrypted output.
- Environment and RBAC setup documentation in `README.md`.

The packaged artifact is `target\encrypted-blob-uploader-1.0.0.jar`. No Azure resources were created or modified.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client is left unclosed; DataKeyEnvelope, the project's own AutoCloseable resource, is correctly cleaned up in both sync and async flows.

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-java-encrypted-uploader --config java-azure-skills/azure-skill-mcp-microsoft-skill
```

---

[← Back to Summary](../../../../../../summary.md)
