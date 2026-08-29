# Evaluation Report: storage-dp-java-encrypted-uploader

**Config:** java-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 519.2s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-encrypted-uploader` |
| Config | java-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 31/32 |
| Duration | 519.2s |
| Timestamp | 2026-08-29T04:35:06Z |
| Files Generated | 14 |
| Event Count | 17771 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 265.6s |
| Review | 252.6s |
| **Total** | **519.2s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-skills/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=457247 out=14117 |
| Turn Count | 13 |

## Tool Calls

### ✅ 1. glob — 4792ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. powershell — 5676ms

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace root"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 3. web_search (via github-mcp-server) — 71737ms

**Input:**

```json
{"query":"Maven Central latest stable versions August 2026 com.azure azure-storage-blob azure-security-keyvault-keys azure-identity"}
```

**Output:**

```
{"type":"output_text","text":{"value":"To upload a blob using the Azure SDK for Java and set custom metadata, you can use either the synchronous `BlobClient` or the asynchronous `BlobAsyncClient`. The upload API supports passing `BlobUploadOptions` (for synchronous) or `BlobUploadOptions` with `uploadWithResponse` (for asynchronous) to configure metadata, HTTP headers, and more.\n\n### Synchronous Example (`BlobClient` + `BinaryData` + Metadata)\n\n```java\nimport com.azure.core.util.BinaryData;\nimport com.azure.storage.blob.*;\nimport com.azure.storage.blob.models.*;\nimport java.util.HashMap;\nimport java.util.Map;\n\nBlobClient blobClient = new BlobServiceClientBuilder()\n    .connectionString(\"<your-connection-string>\")\n    .buildClient()\n    .getBlobContainerClient(\"<your-container>\")\n    .getBlobClient(\"<your-blob-name>\");\n\nMap<String, String> metadata = new HashMap<>();\nmetadata.put(\"key1\", \"value1\");\nmetadata.put(\"key2\", \"value2\");\n\nBinaryData data = BinaryData.fromBytes(\"Hello, Azure!\".getBytes());\n\nBlobUploadOptions options = new BlobUploadOptions(data)\n    .setMetadata(metadata);\n\nblobClient.uploadWithResponse(options, null, null);\n```\n- You can also set HTTP headers via `.setHeaders(new BlobHttpHeaders().setContentType(\"text/plain\"))` inside options.\n\n---\n\n### Asynchronous Example (`BlobAsyncClient` + `BlobUploadOptions`)\n\n```java\nimport com.azure.core.util.BinaryData;\nimport com.azure.storage.blob.*;\nimport com.azure.storage.blob.options.*;\nimport java.util.HashMap;\nimport java.util.Map;\n\nBlobAsyncClient blobAsyncClient = new BlobServiceClientBuilder()\n    .connectionString(\"<your-connection-string>\")\n    .buildAsyncClient()\n    .getBlobContainerAsyncClient(\"<your-container>\")\n    .getBlobAsyncClient(\"<your-blob-name>\");\n\nMap<String, String> metadata = new HashMap<>();\nmetadata.put(\"key1\", \"value1\");\nmetadata.put(\"key2\", \"value2\");\n\nBinaryData data = BinaryData.fromBytes(\"Hello, As
... (truncated)
```

### ✅ 4. web_search (via github-mcp-server) — 70200ms

**Input:**

```json
{"query":"Azure Java SDK CryptographyClient wrapKey unwrapKey async API KeyWrapAlgorithm RSA_OAEP_256 examples"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here is a guide and code example for using the Azure Java SDK's `CryptographyAsyncClient` to perform asynchronous key wrapping and unwrapping with the `RSA_OAEP_256` algorithm. The required methods are `wrapKey` and `unwrapKey`.\n\n### Maven Dependency\nFirst, ensure you have the following Maven dependency:\n```xml\n<dependency>\n  <groupId>com.azure</groupId>\n  <artifactId>azure-security-keyvault-keys</artifactId>\n  <version>4.11.0</version>\n</dependency>\n```\n\n### Client Setup\nCreate the async client:\n```java\nimport com.azure.identity.DefaultAzureCredentialBuilder;\nimport com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;\nimport com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n\nString keyId = \"<your-key-id>\"; // Full URL to your Key Vault key\n\nCryptographyAsyncClient client = new CryptographyClientBuilder()\n    .keyIdentifier(keyId)\n    .credential(new DefaultAzureCredentialBuilder().build())\n    .buildAsyncClient();\n```\n\n### Wrap and Unwrap Key with RSA_OAEP_256 (Async)\n```java\nimport com.azure.security.keyvault.keys.cryptography.KeyWrapAlgorithm;\nimport reactor.core.publisher.Mono;\n\nbyte[] keyToWrap = ...; // The symmetric key to be wrapped (for example, a 256-bit AES key)\n\n// Async wrap operation\nMono<byte[]> wrappedKeyMono = client.wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, keyToWrap)\n    .map(result -> result.getEncryptedKey());\n\n// Async unwrap operation (after obtaining wrappedKey)\nMono<byte[]> unwrappedKeyMono = wrappedKeyMono.flatMap(wrappedKey -> \n    client.unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, wrappedKey)\n        .map(result -> result.getKey())\n);\n\n// Example: subscribe to see the result\nunwrappedKeyMono.subscribe(unwrappedKey -> {\n    // Use the unwrapped key\n    System.out.println(\"Unwrapped key length: \" + unwrappedKey.length);\n});\n```\n\n- `wrapKey`: Encrypts a symmetric key using `RSA_OAEP_256` and returns an encrypted by
... (truncated)
```

### ✅ 5. web_search (via github-mcp-server) — 68746ms

**Input:**

```json
{"query":"Azure Java BlobClient upload BinaryData metadata setMetadata BlobAsyncClient uploadWithResponse BlobUploadOptions API"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here are the latest stable Maven Central versions for August 2026 for the requested Azure Java SDK packages:\n\n1. **com.azure:azure-storage-blob**  \n   - Latest stable version: **12.26.0** \n\n2. **com.azure:azure-security-keyvault-keys**  \n   - Latest stable version: **4.11.0**【5:0†source】\n\n3. **com.azure:azure-identity**  \n   - Latest stable version: **1.18.4**【7:8†source】\n\nThese versions reflect the most current stable releases found as of August 2026. Let me know if you would like Maven XML dependency snippets or additional package information!","annotations":[{"text":"","start_index":276,"end_index":288,"url_citation":{"title":"Maven Repository: com.azure » azure-security-keyvault-keys","url":"https://mvnrepository.com/artifact/com.azure/azure-security-keyvault-keys"}},{"text":"","start_index":362,"end_index":374,"url_citation":{"title":"Central Repository: com/azure/azure-identity/1.18.4","url":"https://repo1.maven.org/maven2/com/azure/azure-identity/1.18.4/"}}]},"bing_searches":[{"text":"Maven Central latest stable version August 2026 com.azure:azure-storage-blob","url":"https://www.bing.com/search?q=Maven+Central+latest+stable+version+August+2026+com.azure%3Aazure-storage-blob"},{"text":"Maven Central latest stable version August 2026 com.azure azure-security-keyvault-keys","url":"https://www.bing.com/search?q=Maven+Central+latest+stable+version+August+2026+com.azure+azure-security-keyvault-keys"},{"text":"Maven Central latest stable version August 2026 com.azure azure-identity","url":"https://www.bing.com/search?q=Maven+Central+latest+stable+version+August+2026+com.azure+azure-identity"}],"annotations":null}
```

### ✅ 6. web_fetch — 4868ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
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

### ✅ 7. web_fetch — 4567ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-keys/maven-metadata.xml"}
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

### ✅ 8. web_fetch — 4263ms

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
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

### ✅ 9. web_fetch — 4793ms

**Input:**

```json
{"max_length":5000,"start_index":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-keys/maven-metadata.xml:
on>
      <version>4.9.4</version>
      <version>4.10.0-beta.1</version>
      <version>4.10.0</version>
      <version>4.10.1</version>
      <version>4.10.2</version>
      <version>4.10.3</version>
      <version>4.10.4</version>
      <version>4.10.5</version>
      <version>4.10.6</version>
      <version>4.10.7</version>
      <version>4.11.0</version>
      <version>4.11.1</version>
      <version>4.11.2</version>
      <version>4.12.0-beta.1</version>
    </versions>
    <lastUpdated>20260818232407</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 10. web_fetch — 4485ms

**Input:**

```json
{"max_length":3000,"start_index":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-security-keyvault-keys/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<error>No more content available.</error>
```

### ✅ 11. web_fetch — 4175ms

**Input:**

```json
{"max_length":3000,"start_index":5000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Content type text/xml cannot be simplified to markdown. Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
1.0</version>
      <version>12.31.1</version>
      <version>12.31.2</version>
      <version>12.31.3</version>
      <version>12.32.0-beta.1</version>
      <version>12.32.0</version>
      <version>12.33.0-beta.1</version>
      <version>12.33.0</version>
      <version>12.33.1</version>
      <version>12.33.2</version>
      <version>12.33.3</version>
      <version>12.33.4</version>
      <version>12.34.0-beta.1</version>
      <version>12.34.0</version>
      <version>12.35.0-beta.1</version>
      <version>12.35.0</version>
      <version>12.35.1</version>
      <version>12.36.0-beta.1</version>
    </versions>
    <lastUpdated>20260818232411</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 12. apply_patch — 3676ms

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-encrypted-blob-uploader\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure client-side encrypted blob uploader\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.4\u003c/azure.identity.version\u003e\n+        \u003cazure.keyvault.keys.version\u003e4.11.2\u003c/azure.keyvault.keys.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.35.1\u003c/azure.storage.blob.version\u003e\n+        \u003cjunit.version\u003e5.14.0\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-keys\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.keyvault.keys.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.6.2\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.encryptedblob.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/encryptedblob/EnvelopeEncryptionException.java\n+package com.example.encryptedblob;\n+\n+public final class EnvelopeEncryptionException extends RuntimeException {\n+    public EnvelopeEncryptionException(String message) {\n+        super(message);\n+    }\n+\n+    public EnvelopeEncryptionException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/WrappedDataKey.java\n+package com.example.encryptedblob;\n+\n+import java.util.Arrays;\n+\n+public record WrappedDataKey(byte[] plaintextKey, byte[] wrappedKey, String keyId) {\n+    public WrappedDataKey {\n+        plaintextKey = Arrays.copyOf(plaintextKey, plaintextKey.length);\n+        wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+\n+    @Override\n+    public byte[] plaintextKey() {\n+        return Arrays.copyOf(plaintextKey, plaintextKey.length);\n+    }\n+\n+    @Override\n+    public byte[] wrappedKey() {\n+        return Arrays.copyOf(wrappedKey, wrappedKey.length);\n+    }\n+\n+    public void destroy() {\n+        Arrays.fill(plaintextKey, (byte) 0);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/UploadReceipt.java\n+package com.example.encryptedblob;\n+\n+import java.util.Arrays;\n+import java.util.Base64;\n+\n+public record UploadReceipt(String keyId, byte[] wrappedDataKey) {\n+    public UploadReceipt {\n+        wrappedDataKey = Arrays.copyOf(wrappedDataKey, wrappedDataKey.length);\n+    }\n+\n+    @Override\n+    public byte[] wrappedDataKey() {\n+        return Arrays.copyOf(wrappedDataKey, wrappedDataKey.length);\n+    }\n+\n+    public String wrappedDataKeyBase64() {\n+        return Base64.getEncoder().encodeToString(wrappedDataKey);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/CipherSupport.java\n+package com.example.encryptedblob;\n+\n+import javax.crypto.AEADBadTagException;\n+import javax.crypto.Cipher;\n+import javax.crypto.spec.GCMParameterSpec;\n+import javax.crypto.spec.SecretKeySpec;\n+import java.security.GeneralSecurityException;\n+import java.security.SecureRandom;\n+\n+final class CipherSupport {\n+    static final String CONTENT_ALGORITHM = \"AES/GCM/NoPadding\";\n+    static final int DATA_KEY_BYTES = 32;\n+    static final int IV_BYTES = 12;\n+    private static final int GCM_TAG_BITS = 128;\n+    private static final SecureRandom RANDOM = new SecureRandom();\n+\n+    private CipherSupport() {\n+    }\n+\n+    static byte[] generateDataKey() {\n+        byte[] key = new byte[DATA_KEY_BYTES];\n+        RANDOM.nextBytes(key);\n+        return key;\n+    }\n+\n+    static EncryptedData encrypt(byte[] plaintext, byte[] key, byte[] authenticatedData) {\n+        byte[] iv = new byte[IV_BYTES];\n+        RANDOM.nextBytes(iv);\n+        try {\n+            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);\n+            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key, \"AES\"),\n+                    new GCMParameterSpec(GCM_TAG_BITS, iv));\n+            cipher.updateAAD(authenticatedData);\n+            return new EncryptedData(iv, cipher.doFinal(plaintext));\n+        } catch (GeneralSecurityException e) {\n+            throw new EnvelopeEncryptionException(\"Local encryption failed\", e);\n+        }\n+    }\n+\n+    static byte[] decrypt(byte[] ciphertext, byte[] key, byte[] iv, byte[] authenticatedData) {\n+        try {\n+            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);\n+            cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(key, \"AES\"),\n+                    new GCMParameterSpec(GCM_TAG_BITS, iv));\n+            cipher.updateAAD(authenticatedData);\n+            return cipher.doFinal(ciphertext);\n+        } catch (AEADBadTagException e) {\n+            throw new EnvelopeEncryptionException(\n+                    \"Ciphertext or encryption metadata failed authentication\", e);\n+        } catch (GeneralSecurityException e) {\n+            throw new EnvelopeEncryptionException(\"Local decryption failed\", e);\n+        }\n+    }\n+\n+    record EncryptedData(byte[] iv, byte[] ciphertext) {\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EnvelopeMetadata.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.util.Base64;\n+import java.util.LinkedHashMap;\n+import java.util.Map;\n+\n+final class EnvelopeMetadata {\n+    static final String WRAP_ALGORITHM = \"RSA-OAEP-256\";\n+    private static final String FORMAT_VERSION = \"1\";\n+    private static final String VERSION = \"enc-version\";\n+    private static final String CONTENT_ALG = \"enc-content-alg\";\n+    private static final String WRAP_ALG = \"enc-wrap-alg\";\n+    private static final String KEY_ID = \"enc-key-id\";\n+    private static final String WRAPPED_KEY = \"enc-wrapped-key\";\n+    private static final String IV = \"enc-iv\";\n+\n+    private EnvelopeMetadata() {\n+    }\n+\n+    static Map\u003cString, String\u003e create(String keyId, byte[] wrappedKey, byte[] iv) {\n+        Map\u003cString, String\u003e metadata = new LinkedHashMap\u003c\u003e();\n+        metadata.put(VERSION, FORMAT_VERSION);\n+        metadata.put(CONTENT_ALG, CipherSupport.CONTENT_ALGORITHM);\n+        metadata.put(WRAP_ALG, WRAP_ALGORITHM);\n+        metadata.put(KEY_ID, keyId);\n+        metadata.put(WRAPPED_KEY, encode(wrappedKey));\n+        metadata.put(IV, encode(iv));\n+        return metadata;\n+    }\n+\n+    static Parsed parse(Map\u003cString, String\u003e metadata) {\n+        String version = required(metadata, VERSION);\n+        String contentAlgorithm = required(metadata, CONTENT_ALG);\n+        String wrapAlgorithm = required(metadata, WRAP_ALG);\n+        if (!FORMAT_VERSION.equals(version)\n+                || !CipherSupport.CONTENT_ALGORITHM.equals(contentAlgorithm)\n+                || !WRAP_ALGORITHM.equals(wrapAlgorithm)) {\n+            throw new EnvelopeEncryptionException(\"Blob uses an unsupported encryption format\");\n+        }\n+\n+        byte[] wrappedKey = decode(required(metadata, WRAPPED_KEY), WRAPPED_KEY);\n+        byte[] iv = decode(required(metadata, IV), IV);\n+        if (wrappedKey.length == 0 || iv.length != CipherSupport.IV_BYTES) {\n+            throw new EnvelopeEncryptionException(\"Blob encryption metadata is malformed\");\n+        }\n+        return new Parsed(required(metadata, KEY_ID), wrappedKey, iv);\n+    }\n+\n+    static byte[] authenticatedData(String blobName, String keyId) {\n+        return (\"azure-envelope-v1\\n\" + keyId + \"\\n\" + blobName)\n+                .getBytes(StandardCharsets.UTF_8);\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e metadata, String name) {\n+        String value = metadata.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new EnvelopeEncryptionException(\n+                    \"Blob is missing required encryption metadata: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static String encode(byte[] value) {\n+        return Base64.getEncoder().encodeToString(value);\n+    }\n+\n+    private static byte[] decode(String value, String name) {\n+        try {\n+            return Base64.getDecoder().decode(value);\n+        } catch (IllegalArgumentException e) {\n+            throw new EnvelopeEncryptionException(\n+                    \"Blob encryption metadata is not valid base64: \" + name, e);\n+        }\n+    }\n+\n+    record Parsed(String keyId, byte[] wrappedKey, byte[] iv) {\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagementService.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class KeyManagementService {\n+    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;\n+    private final TokenCredential credential;\n+    private final String currentKeyId;\n+    private final CryptographyClient currentKeyClient;\n+\n+    public KeyManagementService(\n+            KeyClient keyClient, TokenCredential credential, String keyName) {\n+        this.credential = Objects.requireNonNull(credential, \"credential\");\n+        try {\n+            this.currentKeyId = keyClient.getKey(keyName).getId();\n+            this.currentKeyClient = cryptographyClient(currentKeyId);\n+        } catch (HttpResponseException e) {\n+            throw keyVaultFailure(\"resolve wrapping key '\" + keyName + \"'\", e);\n+        }\n+    }\n+\n+    public WrappedDataKey generateAndWrapDataKey() {\n+        byte[] plaintextKey = CipherSupport.generateDataKey();\n+        try {\n+            byte[] wrappedKey = currentKeyClient.wrapKey(WRAP_ALGORITHM, plaintextKey)\n+                    .getEncryptedKey();\n+            return new WrappedDataKey(plaintextKey, wrappedKey, currentKeyId);\n+        } catch (HttpResponseException e) {\n+            throw keyVaultFailure(\"wrap the data encryption key\", e);\n+        } finally {\n+            Arrays.fill(plaintextKey, (byte) 0);\n+        }\n+    }\n+\n+    public byte[] unwrapDataKey(String keyId, byte[] wrappedKey) {\n+        try {\n+            return clientFor(keyId).unwrapKey(WRAP_ALGORITHM, wrappedKey).getKey();\n+        } catch (HttpResponseException e) {\n+            throw keyVaultFailure(\n+                    \"unwrap the data encryption key with Key Vault key \" + keyId, e);\n+        }\n+    }\n+\n+    private CryptographyClient clientFor(String keyId) {\n+        return currentKeyId.equals(keyId) ? currentKeyClient : cryptographyClient(keyId);\n+    }\n+\n+    private CryptographyClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+                .keyIdentifier(keyId)\n+                .credential(credential)\n+                .buildClient();\n+    }\n+\n+    private static EnvelopeEncryptionException keyVaultFailure(\n+            String operation, HttpResponseException cause) {\n+        return new EnvelopeEncryptionException(\n+                \"Azure Key Vault could not \" + operation\n+                        + \" (HTTP \" + cause.getResponse().getStatusCode() + \")\",\n+                cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagementAsyncService.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class KeyManagementAsyncService {\n+    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;\n+    private final TokenCredential credential;\n+    private final String currentKeyId;\n+    private final CryptographyAsyncClient currentKeyClient;\n+\n+    private KeyManagementAsyncService(TokenCredential credential, String keyId) {\n+        this.credential = credential;\n+        this.currentKeyId = keyId;\n+        this.currentKeyClient = cryptographyClient(keyId);\n+    }\n+\n+    public static Mono\u003cKeyManagementAsyncService\u003e create(\n+            KeyAsyncClient keyClient, TokenCredential credential, String keyName) {\n+        Objects.requireNonNull(credential, \"credential\");\n+        return keyClient.getKey(keyName)\n+                .map(key -\u003e new KeyManagementAsyncService(credential, key.getId()))\n+                .onErrorMap(HttpResponseException.class,\n+                        e -\u003e keyVaultFailure(\"resolve wrapping key '\" + keyName + \"'\", e));\n+    }\n+\n+    public Mono\u003cWrappedDataKey\u003e generateAndWrapDataKey() {\n+        return Mono.defer(() -\u003e {\n+            byte[] plaintextKey = CipherSupport.generateDataKey();\n+            return currentKeyClient.wrapKey(WRAP_ALGORITHM, plaintextKey)\n+                    .map(result -\u003e new WrappedDataKey(\n+                            plaintextKey, result.getEncryptedKey(), currentKeyId))\n+                    .doOnError(ignored -\u003e Arrays.fill(plaintextKey, (byte) 0));\n+        }).onErrorMap(HttpResponseException.class,\n+                e -\u003e keyVaultFailure(\"wrap the data encryption key\", e));\n+    }\n+\n+    public Mono\u003cbyte[]\u003e unwrapDataKey(String keyId, byte[] wrappedKey) {\n+        return clientFor(keyId).unwrapKey(WRAP_ALGORITHM, wrappedKey)\n+                .map(result -\u003e result.getKey())\n+                .onErrorMap(HttpResponseException.class,\n+                        e -\u003e keyVaultFailure(\n+                                \"unwrap the data encryption key with Key Vault key \" + keyId, e));\n+    }\n+\n+    private CryptographyAsyncClient clientFor(String keyId) {\n+        return currentKeyId.equals(keyId) ? currentKeyClient : cryptographyClient(keyId);\n+    }\n+\n+    private CryptographyAsyncClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+                .keyIdentifier(keyId)\n+                .credential(credential)\n+                .buildAsyncClient();\n+    }\n+\n+    private static EnvelopeEncryptionException keyVaultFailure(\n+            String operation, HttpResponseException cause) {\n+        return new EnvelopeEncryptionException(\n+                \"Azure Key Vault could not \" + operation\n+                        + \" (HTTP \" + cause.getResponse().getStatusCode() + \")\",\n+                cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobUploadFromFileOptions;\n+import com.azure.storage.blob.options.BlobUploadOptions;\n+import com.azure.core.util.BinaryData;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.util.Arrays;\n+import java.util.Map;\n+\n+public final class EncryptedBlobClient {\n+    private final BlobContainerClient containerClient;\n+    private final KeyManagementService keyManagement;\n+\n+    public EncryptedBlobClient(\n+            BlobContainerClient containerClient, KeyManagementService keyManagement) {\n+        this.containerClient = containerClient;\n+        this.keyManagement = keyManagement;\n+    }\n+\n+    public UploadReceipt upload(String blobName, byte[] plaintext) {\n+        WrappedDataKey dataKey = keyManagement.generateAndWrapDataKey();\n+        byte[] plaintextKey = dataKey.plaintextKey();\n+        try {\n+            byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());\n+            CipherSupport.EncryptedData encrypted =\n+                    CipherSupport.encrypt(plaintext, plaintextKey, aad);\n+            Map\u003cString, String\u003e metadata = EnvelopeMetadata.create(\n+                    dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());\n+            blob(blobName).uploadWithResponse(\n+                    new BlobUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))\n+                            .setMetadata(metadata),\n+                    null,\n+                    null);\n+            return new UploadReceipt(dataKey.keyId(), dataKey.wrappedKey());\n+        } catch (BlobStorageException e) {\n+            throw blobFailure(\"upload\", blobName, e);\n+        } finally {\n+            Arrays.fill(plaintextKey, (byte) 0);\n+            dataKey.destroy();\n+        }\n+    }\n+\n+    public UploadReceipt uploadFile(String blobName, Path source) {\n+        try {\n+            return upload(blobName, Files.readAllBytes(source));\n+        } catch (IOException e) {\n+            throw new EnvelopeEncryptionException(\"Could not read input file \" + source, e);\n+        }\n+    }\n+\n+    public byte[] download(String blobName) {\n+        BlobClient blob = blob(blobName);\n+        try {\n+            Map\u003cString, String\u003e metadata = blob.getProperties().getMetadata();\n+            EnvelopeMetadata.Parsed envelope = EnvelopeMetadata.parse(metadata);\n+            byte[] plaintextKey =\n+                    keyManagement.unwrapDataKey(envelope.keyId(), envelope.wrappedKey());\n+            try {\n+                byte[] aad = EnvelopeMetadata.authenticatedData(blobName, envelope.keyId());\n+                return CipherSupport.decrypt(\n+                        blob.downloadContent().toBytes(), plaintextKey, envelope.iv(), aad);\n+            } finally {\n+                Arrays.fill(plaintextKey, (byte) 0);\n+            }\n+        } catch (BlobStorageException e) {\n+            throw blobFailure(\"download\", blobName, e);\n+        }\n+    }\n+\n+    public void downloadFile(String blobName, Path destination) {\n+        try {\n+            Files.write(destination, download(blobName));\n+        } catch (IOException e) {\n+            throw new EnvelopeEncryptionException(\n+                    \"Could not write decrypted file \" + destination, e);\n+        }\n+    }\n+\n+    private BlobClient blob(String blobName) {\n+        return containerClient.getBlobClient(blobName);\n+    }\n+\n+    private static EnvelopeEncryptionException blobFailure(\n+            String operation, String blobName, BlobStorageException cause) {\n+        return new EnvelopeEncryptionException(\n+                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName\n+                        + \"' (HTTP \" + cause.getStatusCode()\n+                        + \", error \" + cause.getErrorCode() + \")\",\n+                cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobUploadOptions;\n+import reactor.core.publisher.Mono;\n+import reactor.core.scheduler.Schedulers;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.util.Arrays;\n+import java.util.Map;\n+\n+public final class EncryptedBlobAsyncClient {\n+    private final BlobContainerAsyncClient containerClient;\n+    private final KeyManagementAsyncService keyManagement;\n+\n+    public EncryptedBlobAsyncClient(\n+            BlobContainerAsyncClient containerClient,\n+            KeyManagementAsyncService keyManagement) {\n+        this.containerClient = containerClient;\n+        this.keyManagement = keyManagement;\n+    }\n+\n+    public Mono\u003cUploadReceipt\u003e upload(String blobName, byte[] plaintext) {\n+        return keyManagement.generateAndWrapDataKey()\n+                .flatMap(dataKey -\u003e uploadEncrypted(blobName, plaintext, dataKey)\n+                        .doFinally(ignored -\u003e dataKey.destroy()));\n+    }\n+\n+    public Mono\u003cUploadReceipt\u003e uploadFile(String blobName, Path source) {\n+        return Mono.fromCallable(() -\u003e Files.readAllBytes(source))\n+                .subscribeOn(Schedulers.boundedElastic())\n+                .onErrorMap(IOException.class,\n+                        e -\u003e new EnvelopeEncryptionException(\n+                                \"Could not read input file \" + source, e))\n+                .flatMap(bytes -\u003e upload(blobName, bytes));\n+    }\n+\n+    public Mono\u003cbyte[]\u003e download(String blobName) {\n+        BlobAsyncClient blob = blob(blobName);\n+        return blob.getProperties()\n+                .map(properties -\u003e EnvelopeMetadata.parse(properties.getMetadata()))\n+                .flatMap(envelope -\u003e keyManagement\n+                        .unwrapDataKey(envelope.keyId(), envelope.wrappedKey())\n+                        .flatMap(key -\u003e blob.downloadContent()\n+                                .map(content -\u003e decrypt(blobName, content, envelope, key))\n+                                .doFinally(ignored -\u003e Arrays.fill(key, (byte) 0))))\n+                .onErrorMap(BlobStorageException.class,\n+                        e -\u003e blobFailure(\"download\", blobName, e));\n+    }\n+\n+    public Mono\u003cVoid\u003e downloadFile(String blobName, Path destination) {\n+        return download(blobName)\n+                .flatMap(bytes -\u003e Mono.fromRunnable(() -\u003e write(destination, bytes))\n+                        .subscribeOn(Schedulers.boundedElastic()));\n+    }\n+\n+    private Mono\u003cUploadReceipt\u003e uploadEncrypted(\n+            String blobName, byte[] plaintext, WrappedDataKey dataKey) {\n+        byte[] plaintextKey = dataKey.plaintextKey();\n+        try {\n+            byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());\n+            CipherSupport.EncryptedData encrypted =\n+                    CipherSupport.encrypt(plaintext, plaintextKey, aad);\n+            Map\u003cString, String\u003e metadata = EnvelopeMetadata.create(\n+                    dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());\n+            return blob(blobName)\n+                    .uploadWithResponse(\n+                            new BlobUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))\n+                                    .setMetadata(metadata))\n+                    .map(ignored -\u003e new UploadReceipt(dataKey.keyId(), dataKey.wrappedKey()))\n+                    .onErrorMap(BlobStorageException.class,\n+                            e -\u003e blobFailure(\"upload\", blobName, e))\n+                    .doFinally(ignored -\u003e Arrays.fill(plaintextKey, (byte) 0));\n+        } catch (RuntimeException e) {\n+            Arrays.fill(plaintextKey, (byte) 0);\n+            return Mono.error(e);\n+        }\n+    }\n+\n+    private byte[] decrypt(\n+            String blobName,\n+            BinaryData content,\n+            EnvelopeMetadata.Parsed envelope,\n+            byte[] plaintextKey) {\n+        byte[] aad = EnvelopeMetadata.authenticatedData(blobName, envelope.keyId());\n+        return CipherSupport.decrypt(content.toBytes(), plaintextKey, envelope.iv(), aad);\n+    }\n+\n+    private BlobAsyncClient blob(String blobName) {\n+        return containerClient.getBlobAsyncClient(blobName);\n+    }\n+\n+    private static void write(Path destination, byte[] bytes) {\n+        try {\n+            Files.write(destination, bytes);\n+        } catch (IOException e) {\n+            throw new EnvelopeEncryptionException(\n+                    \"Could not write decrypted file \" + destination, e);\n+        }\n+    }\n+\n+    private static EnvelopeEncryptionException blobFailure(\n+            String operation, String blobName, BlobStorageException cause) {\n+        return new EnvelopeEncryptionException(\n+                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName\n+                        + \"' (HTTP \" + cause.getStatusCode()\n+                        + \", error \" + cause.getErrorCode() + \")\",\n+                cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AzureConfiguration.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.KeyClientBuilder;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.BlobServiceAsyncClient;\n+import com.azure.storage.blob.BlobServiceClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Map;\n+\n+public final class AzureConfiguration {\n+    private final TokenCredential credential;\n+    private final BlobContainerClient blobContainerClient;\n+    private final BlobContainerAsyncClient blobContainerAsyncClient;\n+    private final KeyClient keyClient;\n+    private final KeyAsyncClient keyAsyncClient;\n+    private final String keyName;\n+\n+    private AzureConfiguration(\n+            TokenCredential credential,\n+            BlobContainerClient blobContainerClient,\n+            BlobContainerAsyncClient blobContainerAsyncClient,\n+            KeyClient keyClient,\n+            KeyAsyncClient keyAsyncClient,\n+            String keyName) {\n+        this.credential = credential;\n+        this.blobContainerClient = blobContainerClient;\n+        this.blobContainerAsyncClient = blobContainerAsyncClient;\n+        this.keyClient = keyClient;\n+        this.keyAsyncClient = keyAsyncClient;\n+        this.keyName = keyName;\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        Map\u003cString, String\u003e environment = System.getenv();\n+        String blobEndpoint = required(environment, \"AZURE_STORAGE_BLOB_ENDPOINT\");\n+        String containerName = required(environment, \"AZURE_STORAGE_CONTAINER\");\n+        String vaultEndpoint = required(environment, \"AZURE_KEY_VAULT_ENDPOINT\");\n+        String keyName = required(environment, \"AZURE_KEY_VAULT_KEY_NAME\");\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder =\n+                new ManagedIdentityCredentialBuilder();\n+        String clientId = environment.get(\"AZURE_CLIENT_ID\");\n+        if (clientId != null \u0026\u0026 !clientId.isBlank()) {\n+            credentialBuilder.clientId(clientId);\n+        }\n+        TokenCredential credential = credentialBuilder.build();\n+\n+        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()\n+                .endpoint(blobEndpoint)\n+                .credential(credential);\n+        BlobServiceClient syncBlobService = blobBuilder.buildClient();\n+        BlobServiceAsyncClient asyncBlobService = blobBuilder.buildAsyncClient();\n+\n+        KeyClientBuilder keyBuilder = new KeyClientBuilder()\n+                .vaultUrl(vaultEndpoint)\n+                .credential(credential);\n+\n+        return new AzureConfiguration(\n+                credential,\n+                syncBlobService.getBlobContainerClient(containerName),\n+                asyncBlobService.getBlobContainerAsyncClient(containerName),\n+                keyBuilder.buildClient(),\n+                keyBuilder.buildAsyncClient(),\n+                keyName);\n+    }\n+\n+    public EncryptedBlobClient encryptedBlobClient() {\n+        KeyManagementService keyManagement =\n+                new KeyManagementService(keyClient, credential, keyName);\n+        return new EncryptedBlobClient(blobContainerClient, keyManagement);\n+    }\n+\n+    public Mono\u003cEncryptedBlobAsyncClient\u003e encryptedBlobAsyncClient() {\n+        return KeyManagementAsyncService.create(keyAsyncClient, credential, keyName)\n+                .map(keyManagement -\u003e\n+                        new EncryptedBlobAsyncClient(blobContainerAsyncClient, keyManagement));\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e environment, String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/Main.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+\n+public final class Main {\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();\n+        byte[] plaintext = \"Client-side encryption with an Azure Key Vault KEK.\"\n+                .getBytes(StandardCharsets.UTF_8);\n+\n+        EncryptedBlobClient syncClient = configuration.encryptedBlobClient();\n+        UploadReceipt syncReceipt = syncClient.upload(\"sync-encrypted-demo.bin\", plaintext);\n+        byte[] syncDecrypted = syncClient.download(\"sync-encrypted-demo.bin\");\n+        printResult(\"sync\", syncReceipt, syncDecrypted);\n+\n+        configuration.encryptedBlobAsyncClient()\n+                .flatMap(asyncClient -\u003e\n+                        asyncClient.upload(\"async-encrypted-demo.bin\", plaintext)\n+                                .flatMap(receipt -\u003e asyncClient\n+                                        .download(\"async-encrypted-demo.bin\")\n+                                        .doOnNext(decrypted -\u003e\n+                                                printResult(\"async\", receipt, decrypted))))\n+                .block();\n+    }\n+\n+    private static void printResult(\n+            String implementation, UploadReceipt receipt, byte[] decrypted) {\n+        System.out.println(\"[\" + implementation + \"] Vault key ID: \" + receipt.keyId());\n+        System.out.println(\"[\" + implementation + \"] Wrapped DEK (base64): \"\n+                + receipt.wrappedDataKeyBase64());\n+        System.out.println(\"[\" + implementation + \"] Decrypted output: \"\n+                + new String(decrypted, StandardCharsets.UTF_8));\n+    }\n+}\n*** Add File: src/test/java/com/example/encryptedblob/CipherSupportTest.java\n+package com.example.encryptedblob;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.util.Arrays;\n+\n+import static org.junit.jupiter.api.Assertions.assertArrayEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class CipherSupportTest {\n+    @Test\n+    void encryptsAndDecryptsWithAuthenticatedMetadata() {\n+        byte[] key = CipherSupport.generateDataKey();\n+        byte[] plaintext = \"test data\".getBytes(StandardCharsets.UTF_8);\n+        byte[] aad = EnvelopeMetadata.authenticatedData(\n+                \"example.bin\", \"https://example.vault.azure.net/keys/test/version\");\n+\n+        try {\n+            CipherSupport.EncryptedData encrypted =\n+                    CipherSupport.encrypt(plaintext, key, aad);\n+\n+            assertArrayEquals(plaintext, CipherSupport.decrypt(\n+                    encrypted.ciphertext(), key, encrypted.iv(), aad));\n+        } finally {\n+            Arrays.fill(key, (byte) 0);\n+        }\n+    }\n+\n+    @Test\n+    void rejectsChangedAuthenticatedMetadata() {\n+        byte[] key = CipherSupport.generateDataKey();\n+        byte[] plaintext = \"test data\".getBytes(StandardCharsets.UTF_8);\n+        byte[] aad = EnvelopeMetadata.authenticatedData(\"example.bin\", \"key-id\");\n+\n+        try {\n+            CipherSupport.EncryptedData encrypted =\n+                    CipherSupport.encrypt(plaintext, key, aad);\n+\n+            assertThrows(EnvelopeEncryptionException.class, () -\u003e CipherSupport.decrypt(\n+                    encrypted.ciphertext(),\n+                    key,\n+                    encrypted.iv(),\n+                    EnvelopeMetadata.authenticatedData(\"other.bin\", \"key-id\")));\n+        } finally {\n+            Arrays.fill(key, (byte) 0);\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure client-side encrypted blob uploader\n+\n+Java 17 example using envelope encryption:\n+\n+- A fresh 256-bit AES data encryption key (DEK) is generated for each upload.\n+- Blob content is encrypted locally with AES-GCM.\n+- Azure Key Vault Keys wraps the DEK with `RSA-OAEP-256`; Key Vault key material never\n+  leaves the vault.\n+- The wrapped DEK, IV, algorithms, format version, and versioned Key Vault key ID are stored\n+  as blob metadata. The plaintext DEK is kept only in memory and cleared after use.\n+\n+The sample reads these environment variables:\n+\n+| Variable | Description |\n+|---|---|\n+| `AZURE_STORAGE_BLOB_ENDPOINT` | Blob service endpoint, such as `https://account.blob.core.windows.net` |\n+| `AZURE_STORAGE_CONTAINER` | Existing container name |\n+| `AZURE_KEY_VAULT_ENDPOINT` | Vault endpoint, such as `https://vault.vault.azure.net` |\n+| `AZURE_KEY_VAULT_KEY_NAME` | Existing RSA key name |\n+| `AZURE_CLIENT_ID` | Optional client ID for a user-assigned managed identity |\n+\n+The managed identity needs blob data read/write permissions and Key Vault `keys/get`,\n+`keys/wrapKey`, and `keys/unwrapKey` permissions. The container and RSA key must already exist;\n+the sample does not create Azure resources.\n+\n+Run the sync round trip followed by the async round trip:\n+\n+```text\n+mvn compile exec:java\n+```\n+\n+This compact example buffers each file in memory and is intended for small files. For large\n+files, use a framed/chunked authenticated-encryption format rather than loading the entire\n+plaintext or ciphertext into one byte array.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/pom.xml b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/pom.xml
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/pom.xml
@@ -1,0 +1,68 @@
+<?xml version="1.0" encoding="UTF-8"?>
+<project xmlns="http://maven.apache.org/POM/4.0.0"
+         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
+         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
+    <modelVersion>4.0.0</modelVersion>
+
+    <groupId>com.example</groupId>
+    <artifactId>azure-encrypted-blob-uploader</artifactId>
+    <version>1.0.0</version>
+    <name>Azure client-side encrypted blob uploader</name>
+
+    <properties>
+        <maven.compiler.release>17</maven.compiler.release>
+        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
+        <azure.identity.version>1.18.4</azure.identity.version>
+        <azure.keyvault.keys.version>4.11.2</azure.keyvault.keys.version>
+        <azure.storage.blob.version>12.35.1</azure.storage.blob.version>
+        <junit.version>5.14.0</junit.version>
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
+            <artifactId>azure-storage-blob</artifactId>
+            <version>${azure.storage.blob.ver
... (truncated)
```

### ✅ 13. powershell — 6977ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -------------< com.example:azure-encrypted-blob-uploader >--------------
[INFO] Building Azure client-side encrypted blob uploader 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ azure-encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 11 source files with javac [debug release 17] to target\classes
[INFO] -------------------------------------------------------------
[ERROR] COMPILATION ERROR : 
[INFO] -------------------------------------------------------------
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java:[7,38] cannot find symbol
  symbol:   class BlobUploadOptions
  location: package com.azure.storage.blob.options
[ERROR] /C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java:[7,38] cannot find symbol
  symbol:   class BlobUploadOptions
  location: package com.azure.storag
... (truncated)
```

### ✅ 14. powershell — 5850ms

**Input:**

```json
{"command":"$jar = Get-ChildItem \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-storage-blob\\12.35.1\\*.jar\" | Select-Object -First 1 -ExpandProperty FullName; \u0026 jar tf $jar | Select-String 'Upload.*Options|BlobAsyncClient.class|BlobClient.class'","description":"Inspect Azure Blob SDK classes"}
```

**Output:**

```
com/azure/storage/blob/specialized/AppendBlobAsyncClient.class
com/azure/storage/blob/specialized/PageBlobAsyncClient.class
com/azure/storage/blob/specialized/BlockBlobAsyncClient.class
com/azure/storage/blob/specialized/AppendBlobClient.class
com/azure/storage/blob/specialized/PageBlobClient.class
com/azure/storage/blob/specialized/BlockBlobClient.class
com/azure/storage/blob/BlobAsyncClient.class
com/azure/storage/blob/BlobClient.class
com/azure/storage/blob/options/PageBlobUploadPagesFromUrlOptions.class
com/azure/storage/blob/options/BlobUploadFromFileOptions.class
com/azure/storage/blob/options/BlobParallelUploadOptions.class
com/azure/storage/blob/options/BlobUploadFromUrlOptions.class
com/azure/storage/blob/options/BlockBlobSimpleUploadOptions.class
<shellId: 2 completed with exit code 0>
```

### ✅ 15. view — 5470ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\\src\\main\\java\\com\\example\\encryptedblob\\EncryptedBlobAsyncClient.java","view_range":[45,90]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java
@@ -45,46 +45,46 @@
         return blob.getProperties()
                 .map(properties -> EnvelopeMetadata.parse(properties.getMetadata()))
                 .flatMap(envelope -> keyManagement
                         .unwrapDataKey(envelope.keyId(), envelope.wrappedKey())
                         .flatMap(key -> blob.downloadContent()
                                 .map(content -> decrypt(blobName, content, envelope, key))
                                 .doFinally(ignored -> Arrays.fill(key, (byte) 0))))
                 .onErrorMap(BlobStorageException.class,
                         e -> blobFailure("download", blobName, e));
     }
 
     public Mono<Void> downloadFile(String blobName, Path destination) {
         return download(blobName)
                 .flatMap(bytes -> Mono.fromRunnable(() -> write(destination, bytes))
                         .subscribeOn(Schedulers.boundedElastic()));
     }
 
     private Mono<UploadReceipt> uploadEncrypted(
             String blobName, byte[] plaintext, WrappedDataKey dataKey) {
         byte[] plaintextKey = dataKey.plaintextKey();
         try {
             byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());
   
... (truncated)
```

### ✅ 16. view — 5089ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\\src\\main\\java\\com\\example\\encryptedblob\\EncryptedBlobClient.java","view_range":[1,50]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
@@ -1,50 +1,50 @@
 package com.example.encryptedblob;
 
 import com.azure.storage.blob.BlobClient;
 import com.azure.storage.blob.BlobContainerClient;
 import com.azure.storage.blob.models.BlobStorageException;
 import com.azure.storage.blob.options.BlobUploadFromFileOptions;
 import com.azure.storage.blob.options.BlobUploadOptions;
 import com.azure.core.util.BinaryData;
 
 import java.io.IOException;
 import java.nio.file.Files;
 import java.nio.file.Path;
 import java.util.Arrays;
 import java.util.Map;
 
 public final class EncryptedBlobClient {
     private final BlobContainerClient containerClient;
     private final KeyManagementService keyManagement;
 
     public EncryptedBlobClient(
             BlobContainerClient containerClient, KeyManagementService keyManagement) {
         this.containerClient = containerClient;
         this.keyManagement = keyManagement;
     }
 
     public UploadReceipt upload(String blobName, byte[] plaintext) {
         WrappedDataKey dataKey = keyManagement.generateAndWrapDataKey();
         byte[] plaintextKey = dataKey.plaintextKey();
         try {
             byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());
             CipherSupport.Encrypt
... (truncated)
```

### ✅ 17. powershell — 2984ms

**Input:**

```json
{"command":"$cp = \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-storage-blob\\12.35.1\\azure-storage-blob-12.35.1.jar\"; \u0026 javap -classpath $cp com.azure.storage.blob.BlobClient | Select-String 'uploadWithResponse'; \u0026 javap -classpath $cp com.azure.storage.blob.BlobAsyncClient | Select-String 'uploadWithResponse'; \u0026 javap -classpath $cp com.azure.storage.blob.options.BlobParallelUploadOptions | Select-String 'BlobParallelUploadOptions|setMetadata'","description":"Inspect blob upload method signatures"}
```

**Output:**

```
public void uploadWithResponse(java.io.InputStream, long, com.azure.storage.blob.models.ParallelTransferOptions, 
com.azure.storage.blob.models.BlobHttpHeaders, java.util.Map<java.lang.String, java.lang.String>, 
com.azure.storage.blob.models.AccessTier, com.azure.storage.blob.models.BlobRequestConditions, java.time.Duration, 
com.azure.core.util.Context);
  public com.azure.core.http.rest.Response<com.azure.storage.blob.models.BlockBlobItem> 
uploadWithResponse(com.azure.storage.blob.options.BlobParallelUploadOptions, com.azure.core.util.Context);
  public com.azure.core.http.rest.Response<com.azure.storage.blob.models.BlockBlobItem> 
uploadWithResponse(com.azure.storage.blob.options.BlobParallelUploadOptions, java.time.Duration, 
com.azure.core.util.Context);
  public reactor.core.publisher.Mono<com.azure.core.http.rest.Response<com.azure.storage.blob.models.BlockBlobItem>> 
uploadWithResponse(reactor.core.publisher.Flux<java.nio.ByteBuffer>, 
com.azure.storage.blob.models.ParallelTransferOptions, com.azure.storage.blob.models.BlobHttpHeaders, 
java.util.Map<java.lang.String, java.lang.String>, com.azure.storage.blob.models.AccessTier, 
com.azure.storage.blob.models.BlobRequestConditions);
  public reactor.core.publisher.Mono<com.azure.core.http.rest.Response<com.azure.storage.blob.models.BlockBlobItem>> 
uploadWithResponse(com.azure.storage.blob.options.BlobParallelUploadOptions);
Compiled from "BlobParallelUploadOptions.java"
public class com.azure.storage.blob.options.BlobParallelUploadOptions {
  public com.azure.storage.blob.options.BlobParallelUploadOptions(reactor.core.publisher.Flux<java.nio.ByteBuffer>);
  public com.azure.storage.blob.options.BlobParallelUploadOptions(java.io.InputStream, long);
  public com.azure.storage.blob.options.BlobParallelUploadOptions(java.io.InputStream);
  public com.azure.storage.blob.options.BlobParallelUploadOptions(com.azure.core.util.BinaryData);
  public com.azure.storage.blob.options.BlobParallelUploadOptions 
setParall
... (truncated)
```

### ✅ 18. powershell — 2605ms

**Input:**

```json
{"command":"$cp = \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-storage-blob\\12.35.1\\azure-storage-blob-12.35.1.jar\"; \u0026 javap -classpath $cp com.azure.storage.blob.BlobClient | Select-String 'downloadContent'; \u0026 javap -classpath $cp com.azure.storage.blob.BlobAsyncClient | Select-String 'downloadContent'","description":"Inspect blob download method signatures"}
```

**Output:**

```

<shellId: 4 completed with exit code 0>
```

### ✅ 19. apply_patch — 732ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/encryptedblob/EncryptedBlobClient.java\n@@\n import com.azure.storage.blob.BlobClient;\n import com.azure.storage.blob.BlobContainerClient;\n import com.azure.storage.blob.models.BlobStorageException;\n-import com.azure.storage.blob.options.BlobUploadFromFileOptions;\n-import com.azure.storage.blob.options.BlobUploadOptions;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n@@\n             Map\u003cString, String\u003e metadata = EnvelopeMetadata.create(\n                     dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());\n             blob(blobName).uploadWithResponse(\n-                    new BlobUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))\n+                    new BlobParallelUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))\n                             .setMetadata(metadata),\n                     null,\n                     null);\n*** Update File: src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java\n@@\n import com.azure.storage.blob.BlobContainerAsyncClient;\n import com.azure.storage.blob.models.BlobStorageException;\n-import com.azure.storage.blob.options.BlobUploadOptions;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n@@\n         return download(blobName)\n                 .flatMap(bytes -\u003e Mono.fromRunnable(() -\u003e write(destination, bytes))\n-                        .subscribeOn(Schedulers.boundedElastic()));\n+                        .subscribeOn(Schedulers.boundedElastic())\n+                        .then());\n@@\n             return blob(blobName)\n                     .uploadWithResponse(\n-                            new BlobUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))\n+                            new BlobParallelUploadOptions(\n+                                    BinaryData.fromBytes(encrypted.ciphertext()))\n                                     .setMetadata(metadata))\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobClient.java
@@ -3,8 +3,7 @@
 import com.azure.storage.blob.BlobClient;
 import com.azure.storage.blob.BlobContainerClient;
 import com.azure.storage.blob.models.BlobStorageException;
-import com.azure.storage.blob.options.BlobUploadFromFileOptions;
-import com.azure.storage.blob.options.BlobUploadOptions;
+import com.azure.storage.blob.options.BlobParallelUploadOptions;
 import com.azure.core.util.BinaryData;
 
 import java.io.IOException;
@@ -33,7 +32,7 @@
             Map<String, String> metadata = EnvelopeMetadata.create(
                     dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());
             blob(blobName).uploadWithResponse(
-                    new BlobUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))
+                    new BlobParallelUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))
                             .setMetadata(metadata),
                     null,
                     null);


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skill
... (truncated)
```

### ✅ 20. powershell — 15302ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Recompile project and run tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -------------< com.example:azure-encrypted-blob-uploader >--------------
[INFO] Building Azure client-side encrypted blob uploader 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ azure-encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 11 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ azure-encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.4:test (default-test) @ azure-encrypted-blob-uploader ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] --------------------
... (truncated)
```

### ✅ 21. apply_patch — 809ms

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/encryptedblob/EnvelopeMetadata.java\n@@\n-    private static final String VERSION = \"enc-version\";\n-    private static final String CONTENT_ALG = \"enc-content-alg\";\n-    private static final String WRAP_ALG = \"enc-wrap-alg\";\n-    private static final String KEY_ID = \"enc-key-id\";\n-    private static final String WRAPPED_KEY = \"enc-wrapped-key\";\n-    private static final String IV = \"enc-iv\";\n+    private static final String VERSION = \"enc_version\";\n+    private static final String CONTENT_ALG = \"enc_content_alg\";\n+    private static final String WRAP_ALG = \"enc_wrap_alg\";\n+    private static final String KEY_ID = \"enc_key_id\";\n+    private static final String WRAPPED_KEY = \"enc_wrapped_key\";\n+    private static final String IV = \"enc_iv\";\n*** Update File: src/main/java/com/example/encryptedblob/KeyManagementService.java\n@@\n import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.AzureException;\n import com.azure.core.exception.HttpResponseException;\n@@\n-        } catch (HttpResponseException e) {\n+        } catch (AzureException e) {\n             throw keyVaultFailure(\"resolve wrapping key '\" + keyName + \"'\", e);\n@@\n-        } catch (HttpResponseException e) {\n+        } catch (AzureException e) {\n             throw keyVaultFailure(\"wrap the data encryption key\", e);\n@@\n-        } catch (HttpResponseException e) {\n+        } catch (AzureException e) {\n             throw keyVaultFailure(\n@@\n     private static EnvelopeEncryptionException keyVaultFailure(\n-            String operation, HttpResponseException cause) {\n+            String operation, AzureException cause) {\n+        String status = cause instanceof HttpResponseException httpError\n+                ? \" (HTTP \" + httpError.getResponse().getStatusCode() + \")\"\n+                : \"\";\n         return new EnvelopeEncryptionException(\n-                \"Azure Key Vault could not \" + operation\n-                        + \" (HTTP \" + cause.getResponse().getStatusCode() + \")\",\n+                \"Azure Key Vault could not \" + operation + status,\n                 cause);\n*** Update File: src/main/java/com/example/encryptedblob/KeyManagementAsyncService.java\n@@\n import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.AzureException;\n import com.azure.core.exception.HttpResponseException;\n@@\n-                .onErrorMap(HttpResponseException.class,\n+                .onErrorMap(AzureException.class,\n                         e -\u003e keyVaultFailure(\"resolve wrapping key '\" + keyName + \"'\", e));\n@@\n                     .map(result -\u003e new WrappedDataKey(\n                             plaintextKey, result.getEncryptedKey(), currentKeyId))\n-                    .doOnError(ignored -\u003e Arrays.fill(plaintextKey, (byte) 0));\n-        }).onErrorMap(HttpResponseException.class,\n+                    .doFinally(ignored -\u003e Arrays.fill(plaintextKey, (byte) 0));\n+        }).onErrorMap(AzureException.class,\n                 e -\u003e keyVaultFailure(\"wrap the data encryption key\", e));\n@@\n-                .onErrorMap(HttpResponseException.class,\n+                .onErrorMap(AzureException.class,\n@@\n     private static EnvelopeEncryptionException keyVaultFailure(\n-            String operation, HttpResponseException cause) {\n+            String operation, AzureException cause) {\n+        String status = cause instanceof HttpResponseException httpError\n+                ? \" (HTTP \" + httpError.getResponse().getStatusCode() + \")\"\n+                : \"\";\n         return new EnvelopeEncryptionException(\n-                \"Azure Key Vault could not \" + operation\n-                        + \" (HTTP \" + cause.getResponse().getStatusCode() + \")\",\n+                \"Azure Key Vault could not \" + operation + status,\n                 cause);\n*** Update File: src/main/java/com/example/encryptedblob/EncryptedBlobClient.java\n@@\n import com.azure.storage.blob.BlobClient;\n import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.core.exception.AzureException;\n import com.azure.storage.blob.models.BlobStorageException;\n@@\n-        } catch (BlobStorageException e) {\n+        } catch (AzureException e) {\n             throw blobFailure(\"upload\", blobName, e);\n@@\n-        } catch (BlobStorageException e) {\n+        } catch (AzureException e) {\n             throw blobFailure(\"download\", blobName, e);\n@@\n     private static EnvelopeEncryptionException blobFailure(\n-            String operation, String blobName, BlobStorageException cause) {\n+            String operation, String blobName, AzureException cause) {\n+        String details = cause instanceof BlobStorageException storageError\n+                ? \" (HTTP \" + storageError.getStatusCode()\n+                        + \", error \" + storageError.getErrorCode() + \")\"\n+                : \"\";\n         return new EnvelopeEncryptionException(\n-                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName\n-                        + \"' (HTTP \" + cause.getStatusCode()\n-                        + \", error \" + cause.getErrorCode() + \")\",\n+                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName + \"'\"\n+                        + details,\n                 cause);\n*** Update File: src/main/java/com/example/encryptedblob/EncryptedBlobAsyncClient.java\n@@\n import com.azure.core.util.BinaryData;\n+import com.azure.core.exception.AzureException;\n@@\n-                .onErrorMap(BlobStorageException.class,\n+                .onErrorMap(AzureException.class,\n@@\n-                    .onErrorMap(BlobStorageException.class,\n+                    .onErrorMap(AzureException.class,\n@@\n     private static EnvelopeEncryptionException blobFailure(\n-            String operation, String blobName, BlobStorageException cause) {\n+            String operation, String blobName, AzureException cause) {\n+        String details = cause instanceof BlobStorageException storageError\n+                ? \" (HTTP \" + storageError.getStatusCode()\n+                        + \", error \" + storageError.getErrorCode() + \")\"\n+                : \"\";\n         return new EnvelopeEncryptionException(\n-                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName\n-                        + \"' (HTTP \" + cause.getStatusCode()\n-                        + \", error \" + cause.getErrorCode() + \")\",\n+                \"Azure Blob Storage could not \" + operation + \" blob '\" + blobName + \"'\"\n+                        + details,\n                 cause);\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EnvelopeMetadata.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EnvelopeMetadata.java
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EnvelopeMetadata.java
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/EnvelopeMetadata.java
@@ -8,12 +8,12 @@
 final class EnvelopeMetadata {
     static final String WRAP_ALGORITHM = "RSA-OAEP-256";
     private static final String FORMAT_VERSION = "1";
-    private static final String VERSION = "enc-version";
-    private static final String CONTENT_ALG = "enc-content-alg";
-    private static final String WRAP_ALG = "enc-wrap-alg";
-    private static final String KEY_ID = "enc-key-id";
-    private static final String WRAPPED_KEY = "enc-wrapped-key";
-    private static final String IV = "enc-iv";
+    private static final String VERSION = "enc_version";
+    private static final String CONTENT_ALG = "enc_content_alg";
+    private static final String WRAP_ALG = "enc_wrap_alg";
+    private static final String KEY_ID = "enc_key_id";
+    private static final String WRAPPED_KEY = "enc_wrapped_key";
+    private static final String IV = "enc_iv";
 
     private EnvelopeMetadata() {
     }


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/example/encryptedblob/KeyManagementService.java b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148/src/main/java/com/e
... (truncated)
```

### ✅ 22. powershell — 6182ms

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Run final compile and tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
WARNING: A terminally deprecated method in sun.misc.Unsafe has been called
WARNING: sun.misc.Unsafe::staticFieldBase has been called by com.google.inject.internal.aop.HiddenClassDefiner (file:/C:/apache-maven-3.9.11/lib/guice-5.1.0-classes.jar)
WARNING: Please consider reporting this to the maintainers of class com.google.inject.internal.aop.HiddenClassDefiner
WARNING: sun.misc.Unsafe::staticFieldBase will be removed in a future release
[INFO] Scanning for projects...
[INFO] 
[INFO] -------------< com.example:azure-encrypted-blob-uploader >--------------
[INFO] Building Azure client-side encrypted blob uploader 1.0.0
[INFO]   from pom.xml
[INFO] --------------------------------[ jar ]---------------------------------
[INFO] 
[INFO] --- resources:3.3.1:resources (default-resources) @ azure-encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\src\main\resources
[INFO] 
[INFO] --- compiler:3.14.1:compile (default-compile) @ azure-encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed source code.
[INFO] Compiling 11 source files with javac [debug release 17] to target\classes
[INFO] 
[INFO] --- resources:3.3.1:testResources (default-testResources) @ azure-encrypted-blob-uploader ---
[INFO] skip non existing resourceDirectory C:\Users\weidxu\AppData\Local\Temp\hyoka-storage-dp-java-encrypted-uploader-java-azure-skills-baseline-233266148\src\test\resources
[INFO] 
[INFO] --- compiler:3.14.1:testCompile (default-testCompile) @ azure-encrypted-blob-uploader ---
[INFO] Recompiling the module because of changed dependency.
[INFO] Compiling 1 source file with javac [debug release 17] to target\test-classes
[INFO] 
[INFO] --- surefire:3.5.4:test (default-test) @ azure-encrypted-blob-uploader ---
[INFO] Using auto detected provider org.apache.maven.surefire.junitplatform.JUnitPlatformProvider
[INFO] 
[INFO] --------------------
... (truncated)
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\encryptedblob\AzureConfiguration.java`
- `src\main\java\com\example\encryptedblob\CipherSupport.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobAsyncClient.java`
- `src\main\java\com\example\encryptedblob\EncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\EnvelopeEncryptionException.java`
- `src\main\java\com\example\encryptedblob\EnvelopeMetadata.java`
- `src\main\java\com\example\encryptedblob\KeyManagementAsyncService.java`
- `src\main\java\com\example\encryptedblob\KeyManagementService.java`
- `src\main\java\com\example\encryptedblob\Main.java`
- `src\main\java\com\example\encryptedblob\UploadReceipt.java`
- `src\main\java\com\example\encryptedblob\WrappedDataKey.java`
- `src\test\java\com\example\encryptedblob\CipherSupportTest.java`

## Copilot Response

I’ll inspect the workspace, then create the Maven project and compile/test the local cryptographic paths without contacting Azure.

The workspace is empty. I’m checking the current Azure SDK artifact versions and APIs before adding the source files.

I’ve settled the envelope format: AES-256-GCM locally, RSA-OAEP-256 wrapping in Key Vault, versioned blob metadata, and blob/key-bound authenticated data to detect metadata or ciphertext swaps. I’m adding the complete project now.

The project is in place. I’m compiling and running the local AES-GCM tests now; this will catch Azure SDK API mismatches without making service calls.

The compile exposed two SDK-shape issues: current Blob Storage uses `BlobParallelUploadOptions`, and the async file-write chain needs an explicit `Mono<Void>`. I’m inspecting the installed signatures, then I’ll patch both precisely.

Compilation now succeeds. A final security pass found two implementation details to tighten: Azure metadata keys should use portable identifier characters, and the async wrapper must clear its original DEK buffer after copying it into the short-lived envelope object. I’m applying those fixes and broadening service-context errors to include authentication/transport failures.

Created a complete Java 17 Maven project implementing:

- Sync and reactive async envelope encryption with AES-256-GCM and Key Vault Keys `RSA-OAEP-256` wrap/unwrap operations.
- Sync/async encrypted Blob Storage upload, download, and file APIs.
- Versioned cryptographic blob metadata containing the wrapped DEK, IV, algorithms, and versioned vault key ID.
- Shared managed-identity credential configuration with contextual Blob Storage and Key Vault errors.
- A `Main` demo running sync and async round trips and printing the key ID, wrapped DEK, and decrypted output.
- Local AES-GCM integrity tests and setup documentation in `README.md`.

Configure the documented environment variables, ensure the container and RSA Key Vault key already exist, then run:

```powershell
mvn compile exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client instances are left unclosed.

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
hyoka run --prompt-id storage-dp-java-encrypted-uploader --config java-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
