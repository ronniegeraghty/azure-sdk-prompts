# Evaluation Report: storage-dp-java-encrypted-uploader

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 552.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-java-encrypted-uploader` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 31/32 |
| Duration | 552.4s |
| Timestamp | 2026-08-26T16:47:07Z |
| Files Generated | 16 |
| Event Count | 20339 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 289.0s |
| Review | 261.5s |
| **Total** | **552.4s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-tools/with-azure-tools

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-storage-blob-java, azure-security-keyvault-keys-java, azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=786193 out=16182 |
| Turn Count | 12 |

## Prompt Sent

```
Create a small Java 17 Maven project that uploads files to Azure Blob Storage with client-side encryption, where the encryption key material is managed in Azure Key Vault.

The project needs:

- A **key management class** (both sync and async versions) that interacts with Azure Key Vault's Keys service (not Secrets) to perform cryptographic operations. It should implement envelope encryption: generate a data encryption key locally, use Key Vault to protect (wrap) it, and store the protected key alongside the encrypted blob. For decryption, have Key Vault recover (unwrap) the data key, then decrypt locally. The raw data key should never be persisted anywhere, and the vault's key material should never leave Key Vault.

- A **blob uploader/downloader class** (both sync and async versions) that handles the actual encryption and storage. For upload: generate a data key, encrypt the data locally, protect the data key via Key Vault, then upload the ciphertext to Blob Storage with the protected key and any necessary cryptographic parameters stored as blob metadata. For download: read the blob and its metadata, recover the data key via Key Vault, and decrypt. Should handle errors from both services (e.g., the vault key may have been disabled, or the blob may not exist).

- A **configuration class** that builds the necessary Azure connections for both Blob Storage and Key Vault. It should read endpoints from environment variables and authenticate with managed identity. All connections should share a single credential instance.

- A **Main class** that demos both implementations: runs the full encrypt-upload-download-decrypt round-trip using the sync implementation first, then repeats with the async implementation. Print the vault key ID used, the wrapped DEK (base64), and the decrypted output.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-security-keyvault-keys-java"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-storage-blob-java"}
```

### 🔧 3. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Build a local Java 17 Maven sample using managed identity, Azure Blob Storage, and Azure Key Vault Keys for client-side envelope encryption, with synchronous and asynchronous APIs.","learn":true}
```

### 🔧 5. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate secure Java 17 code for Azure Blob Storage and Key Vault Keys client-side envelope encryption using managed identity.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 6. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Java SDK documentation for Blob Storage metadata upload/download, Key Vault CryptographyClient wrapKey/unwrapKey sync and async, and ManagedIdentityCredential reuse.","learn":true}
```

### 🔧 7. glob

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-keys-java","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java"],"pattern":"**/*"}
```

### 🔧 8. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-tools-with-azure-tools-2361299366","pattern":"**/*"}
```

### 🔧 9. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-keys-java\\references\\examples.md"}
```

### 🔧 10. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-storage-blob-java\\references\\examples.md"}
```

### 🔧 11. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find Java SDK guidance for Key Vault wrap/unwrap and Blob upload/download with metadata.","parameters":{"query":"Azure Java SDK Key Vault CryptographyClient CryptographyAsyncClient wrapKey unwrapKey BlobClient BlobAsyncClient metadata upload download managed identity"}}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Java SDK samples for Key Vault key wrapping and Blob Storage metadata.","parameters":{"language":"java","query":"Java CryptographyClient wrapKey unwrapKey CryptographyAsyncClient BlobParallelUploadOptions metadata BlobAsyncClient"}}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"$doc = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787762898805-copilot-tool-output-1974b9c7779e4475ab854f41f326bd0f.txt' | ConvertFrom-Json; $doc.results | Select-Object -First 10 title,contentUrl | Format-List","description":"Extract official documentation URLs"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$samples = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787762897646-copilot-tool-output-c8deffa79955487085b5fa80739ae2f9.txt' | ConvertFrom-Json; $samples.results | Select-Object -First 10 description,sourceUrl | Format-List","description":"Extract official sample URLs"}
```

### 🔧 16. powershell

**Input:**

```json
{"command":"mvn --version; java -version","description":"Check Java and Maven versions"}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787762898805-copilot-tool-output-1974b9c7779e4475ab854f41f326bd0f.txt","view_range":[1,6]}
```

### 🔧 18. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787762897646-copilot-tool-output-c8deffa79955487085b5fa80739ae2f9.txt","view_range":[1,6]}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read current Key Vault Keys Java cryptography client documentation.","parameters":{"url":"https://learn.microsoft.com/java/api/overview/azure/security-keyvault-keys-readme?view=azure-java-stable"}}
```

### 🔧 20. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read current Blob Storage Java client-side encryption guidance.","parameters":{"url":"https://learn.microsoft.com/azure/storage/blobs/client-side-encryption"}}
```

### 🔧 21. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eencrypted-blob-uploader\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eEncrypted Azure Blob Uploader\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cazure.keyvault.keys.version\u003e4.11.0\u003c/azure.keyvault.keys.version\u003e\n+        \u003cazure.storage.blob.version\u003e12.33.0\u003c/azure.storage.blob.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-keys\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.keyvault.keys.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-storage-blob\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.storage.blob.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.encryptedblob.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: README.md\n+# Encrypted Azure Blob Uploader\n+\n+Java 17 sample for client-side envelope encryption with Azure Blob Storage and\n+Azure Key Vault Keys. Each upload uses a new local AES-256-GCM data encryption\n+key (DEK). An RSA key in Key Vault wraps the DEK with RSA-OAEP-256, and only the\n+wrapped DEK, versioned Key Vault key ID, IV, and algorithm identifiers are\n+stored as blob metadata.\n+\n+The sample expects an existing blob container and an existing RSA or RSA-HSM\n+Key Vault key that permits `wrapKey` and `unwrapKey`. It does not create or\n+modify Azure resources.\n+\n+## Configuration\n+\n+Set these environment variables on an Azure host with managed identity:\n+\n+```text\n+AZURE_STORAGE_BLOB_ENDPOINT=https://\u003caccount\u003e.blob.core.windows.net\n+AZURE_STORAGE_CONTAINER=\u003cexisting-container\u003e\n+AZURE_KEY_VAULT_URL=https://\u003cvault\u003e.vault.azure.net\n+AZURE_KEY_NAME=\u003cexisting-rsa-key\u003e\n+AZURE_CLIENT_ID=\u003coptional-user-assigned-managed-identity-client-id\u003e\n+```\n+\n+Grant the identity only the required data-plane permissions, such as **Storage\n+Blob Data Contributor** on the target container and **Key Vault Crypto User**\n+on the target key.\n+\n+## Build and run\n+\n+```text\n+mvn clean verify\n+mvn exec:java\n+```\n+\n+`Main` performs one synchronous and one asynchronous round trip. The classes\n+also expose `Path`-based upload and download methods for small files. The sample\n+buffers each file in memory; use chunked authenticated encryption for large\n+production files.\n+\n+References:\n+\n+- https://learn.microsoft.com/java/api/overview/azure/security-keyvault-keys-readme\n+- https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme\n+- https://learn.microsoft.com/java/api/overview/azure/identity-readme\n+- https://learn.microsoft.com/azure/storage/blobs/client-side-encryption\n*** Add File: src/main/java/com/example/encryptedblob/ProtectedDataKey.java\n+package com.example.encryptedblob;\n+\n+import java.util.Objects;\n+\n+public record ProtectedDataKey(String keyId, String algorithm, byte[] wrappedKey) {\n+    public ProtectedDataKey {\n+        Objects.requireNonNull(keyId, \"keyId\");\n+        Objects.requireNonNull(algorithm, \"algorithm\");\n+        wrappedKey = Objects.requireNonNull(wrappedKey, \"wrappedKey\").clone();\n+    }\n+\n+    @Override\n+    public byte[] wrappedKey() {\n+        return wrappedKey.clone();\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EnvelopeKey.java\n+package com.example.encryptedblob;\n+\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+final class EnvelopeKey implements AutoCloseable {\n+    private final ProtectedDataKey protectedDataKey;\n+    private final byte[] dataKey;\n+    private boolean closed;\n+\n+    EnvelopeKey(ProtectedDataKey protectedDataKey, byte[] dataKey) {\n+        this.protectedDataKey = Objects.requireNonNull(protectedDataKey, \"protectedDataKey\");\n+        this.dataKey = Objects.requireNonNull(dataKey, \"dataKey\");\n+    }\n+\n+    ProtectedDataKey protectedDataKey() {\n+        ensureOpen();\n+        return protectedDataKey;\n+    }\n+\n+    byte[] dataKey() {\n+        ensureOpen();\n+        return dataKey;\n+    }\n+\n+    @Override\n+    public void close() {\n+        if (!closed) {\n+            Arrays.fill(dataKey, (byte) 0);\n+            closed = true;\n+        }\n+    }\n+\n+    private void ensureOpen() {\n+        if (closed) {\n+            throw new IllegalStateException(\"The data encryption key has already been destroyed\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/KeyManagementException.java\n+package com.example.encryptedblob;\n+\n+public final class KeyManagementException extends RuntimeException {\n+    public KeyManagementException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/SyncKeyManagementClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;\n+import com.azure.security.keyvault.keys.models.KeyVaultKey;\n+\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class SyncKeyManagementClient {\n+    static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;\n+    private static final int DATA_KEY_BYTES = 32;\n+\n+    private final KeyClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public SyncKeyManagementClient(KeyClient keyClient, TokenCredential credential, String keyName) {\n+        this(keyClient, credential, keyName, new SecureRandom());\n+    }\n+\n+    SyncKeyManagementClient(\n+        KeyClient keyClient,\n+        TokenCredential credential,\n+        String keyName,\n+        SecureRandom secureRandom\n+    ) {\n+        this.keyClient = Objects.requireNonNull(keyClient, \"keyClient\");\n+        this.credential = Objects.requireNonNull(credential, \"credential\");\n+        this.keyName = Objects.requireNonNull(keyName, \"keyName\");\n+        this.secureRandom = Objects.requireNonNull(secureRandom, \"secureRandom\");\n+    }\n+\n+    EnvelopeKey generateAndWrapDataKey() {\n+        byte[] dataKey = new byte[DATA_KEY_BYTES];\n+        secureRandom.nextBytes(dataKey);\n+\n+        try {\n+            KeyVaultKey key = keyClient.getKey(keyName);\n+            String keyId = key.getId();\n+            byte[] wrappedKey = cryptographyClient(keyId)\n+                .wrapKey(WRAP_ALGORITHM, dataKey)\n+                .getEncryptedKey();\n+            ProtectedDataKey protectedKey =\n+                new ProtectedDataKey(keyId, WRAP_ALGORITHM.toString(), wrappedKey);\n+            return new EnvelopeKey(protectedKey, dataKey);\n+        } catch (HttpResponseException e) {\n+            Arrays.fill(dataKey, (byte) 0);\n+            throw keyVaultFailure(\"wrap a new data encryption key\", e);\n+        } catch (RuntimeException e) {\n+            Arrays.fill(dataKey, (byte) 0);\n+            throw e;\n+        }\n+    }\n+\n+    byte[] unwrapDataKey(ProtectedDataKey protectedKey) {\n+        Objects.requireNonNull(protectedKey, \"protectedKey\");\n+        validateAlgorithm(protectedKey.algorithm());\n+\n+        try {\n+            return cryptographyClient(protectedKey.keyId())\n+                .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())\n+                .getKey();\n+        } catch (HttpResponseException e) {\n+            throw keyVaultFailure(\"unwrap the data encryption key with \" + protectedKey.keyId(), e);\n+        }\n+    }\n+\n+    private CryptographyClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+            .keyIdentifier(keyId)\n+            .credential(credential)\n+            .buildClient();\n+    }\n+\n+    static void validateAlgorithm(String algorithm) {\n+        if (!WRAP_ALGORITHM.toString().equals(algorithm)) {\n+            throw new IllegalArgumentException(\"Unsupported key wrap algorithm: \" + algorithm);\n+        }\n+    }\n+\n+    private static KeyManagementException keyVaultFailure(String operation, HttpResponseException e) {\n+        int status = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();\n+        return new KeyManagementException(\n+            \"Azure Key Vault could not \" + operation + \" (HTTP \" + status\n+                + \"). Verify that the key exists, is enabled, and permits wrapKey/unwrapKey.\",\n+            e\n+        );\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AsyncKeyManagementClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.exception.HttpResponseException;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;\n+import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;\n+import reactor.core.publisher.Mono;\n+\n+import java.security.SecureRandom;\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class AsyncKeyManagementClient {\n+    private static final int DATA_KEY_BYTES = 32;\n+\n+    private final KeyAsyncClient keyClient;\n+    private final TokenCredential credential;\n+    private final String keyName;\n+    private final SecureRandom secureRandom;\n+\n+    public AsyncKeyManagementClient(\n+        KeyAsyncClient keyClient,\n+        TokenCredential credential,\n+        String keyName\n+    ) {\n+        this(keyClient, credential, keyName, new SecureRandom());\n+    }\n+\n+    AsyncKeyManagementClient(\n+        KeyAsyncClient keyClient,\n+        TokenCredential credential,\n+        String keyName,\n+        SecureRandom secureRandom\n+    ) {\n+        this.keyClient = Objects.requireNonNull(keyClient, \"keyClient\");\n+        this.credential = Objects.requireNonNull(credential, \"credential\");\n+        this.keyName = Objects.requireNonNull(keyName, \"keyName\");\n+        this.secureRandom = Objects.requireNonNull(secureRandom, \"secureRandom\");\n+    }\n+\n+    Mono\u003cEnvelopeKey\u003e generateAndWrapDataKey() {\n+        return Mono.defer(() -\u003e {\n+            byte[] dataKey = new byte[DATA_KEY_BYTES];\n+            secureRandom.nextBytes(dataKey);\n+\n+            return keyClient.getKey(keyName)\n+                .flatMap(key -\u003e cryptographyClient(key.getId())\n+                    .wrapKey(SyncKeyManagementClient.WRAP_ALGORITHM, dataKey)\n+                    .map(result -\u003e new EnvelopeKey(\n+                        new ProtectedDataKey(\n+                            key.getId(),\n+                            SyncKeyManagementClient.WRAP_ALGORITHM.toString(),\n+                            result.getEncryptedKey()\n+                        ),\n+                        dataKey\n+                    )))\n+                .doOnError(ignored -\u003e Arrays.fill(dataKey, (byte) 0));\n+        }).onErrorMap(\n+            HttpResponseException.class,\n+            e -\u003e keyVaultFailure(\"wrap a new data encryption key\", e)\n+        );\n+    }\n+\n+    Mono\u003cbyte[]\u003e unwrapDataKey(ProtectedDataKey protectedKey) {\n+        return Mono.defer(() -\u003e {\n+            Objects.requireNonNull(protectedKey, \"protectedKey\");\n+            SyncKeyManagementClient.validateAlgorithm(protectedKey.algorithm());\n+            return cryptographyClient(protectedKey.keyId())\n+                .unwrapKey(SyncKeyManagementClient.WRAP_ALGORITHM, protectedKey.wrappedKey())\n+                .map(result -\u003e result.getKey());\n+        }).onErrorMap(\n+            HttpResponseException.class,\n+            e -\u003e keyVaultFailure(\n+                \"unwrap the data encryption key with \" + protectedKey.keyId(),\n+                e\n+            )\n+        );\n+    }\n+\n+    private CryptographyAsyncClient cryptographyClient(String keyId) {\n+        return new CryptographyClientBuilder()\n+            .keyIdentifier(keyId)\n+            .credential(credential)\n+            .buildAsyncClient();\n+    }\n+\n+    private static KeyManagementException keyVaultFailure(String operation, HttpResponseException e) {\n+        int status = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();\n+        return new KeyManagementException(\n+            \"Azure Key Vault could not \" + operation + \" (HTTP \" + status\n+                + \"). Verify that the key exists, is enabled, and permits wrapKey/unwrapKey.\",\n+            e\n+        );\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/BlobEncryptionException.java\n+package com.example.encryptedblob;\n+\n+import com.azure.storage.blob.models.BlobStorageException;\n+\n+public final class BlobEncryptionException extends RuntimeException {\n+    public BlobEncryptionException(String message) {\n+        super(message);\n+    }\n+\n+    public BlobEncryptionException(String message, Throwable cause) {\n+        super(message, cause);\n+    }\n+\n+    static BlobEncryptionException storageFailure(\n+        String operation,\n+        String blobName,\n+        BlobStorageException cause\n+    ) {\n+        return new BlobEncryptionException(\n+            \"Blob Storage could not \" + operation + \" '\" + blobName + \"' (HTTP \"\n+                + cause.getStatusCode() + \", error \" + cause.getErrorCode() + \")\",\n+            cause\n+        );\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/LocalAesGcm.java\n+package com.example.encryptedblob;\n+\n+import javax.crypto.AEADBadTagException;\n+import javax.crypto.Cipher;\n+import javax.crypto.spec.GCMParameterSpec;\n+import javax.crypto.spec.SecretKeySpec;\n+import java.security.GeneralSecurityException;\n+import java.security.SecureRandom;\n+import java.util.Objects;\n+\n+final class LocalAesGcm {\n+    static final String ALGORITHM = \"AES/GCM/NoPadding\";\n+    private static final int IV_BYTES = 12;\n+    private static final int TAG_BITS = 128;\n+    private static final SecureRandom SECURE_RANDOM = new SecureRandom();\n+\n+    private LocalAesGcm() {\n+    }\n+\n+    static EncryptedPayload encrypt(byte[] dataKey, byte[] plaintext, byte[] authenticatedMetadata) {\n+        Objects.requireNonNull(dataKey, \"dataKey\");\n+        Objects.requireNonNull(plaintext, \"plaintext\");\n+        Objects.requireNonNull(authenticatedMetadata, \"authenticatedMetadata\");\n+\n+        byte[] iv = new byte[IV_BYTES];\n+        SECURE_RANDOM.nextBytes(iv);\n+        try {\n+            Cipher cipher = Cipher.getInstance(ALGORITHM);\n+            cipher.init(\n+                Cipher.ENCRYPT_MODE,\n+                new SecretKeySpec(dataKey, \"AES\"),\n+                new GCMParameterSpec(TAG_BITS, iv)\n+            );\n+            cipher.updateAAD(authenticatedMetadata);\n+            return new EncryptedPayload(iv, cipher.doFinal(plaintext));\n+        } catch (GeneralSecurityException e) {\n+            throw new BlobEncryptionException(\"Local AES-GCM encryption failed\", e);\n+        }\n+    }\n+\n+    static byte[] decrypt(\n+        byte[] dataKey,\n+        byte[] ciphertext,\n+        byte[] iv,\n+        byte[] authenticatedMetadata\n+    ) {\n+        Objects.requireNonNull(dataKey, \"dataKey\");\n+        Objects.requireNonNull(ciphertext, \"ciphertext\");\n+        Objects.requireNonNull(iv, \"iv\");\n+        Objects.requireNonNull(authenticatedMetadata, \"authenticatedMetadata\");\n+        if (iv.length != IV_BYTES) {\n+            throw new BlobEncryptionException(\"Invalid AES-GCM IV length: \" + iv.length);\n+        }\n+\n+        try {\n+            Cipher cipher = Cipher.getInstance(ALGORITHM);\n+            cipher.init(\n+                Cipher.DECRYPT_MODE,\n+                new SecretKeySpec(dataKey, \"AES\"),\n+                new GCMParameterSpec(TAG_BITS, iv)\n+            );\n+            cipher.updateAAD(authenticatedMetadata);\n+            return cipher.doFinal(ciphertext);\n+        } catch (AEADBadTagException e) {\n+            throw new BlobEncryptionException(\n+                \"AES-GCM authentication failed; ciphertext or encryption metadata was modified\",\n+                e\n+            );\n+        } catch (GeneralSecurityException e) {\n+            throw new BlobEncryptionException(\"Local AES-GCM decryption failed\", e);\n+        }\n+    }\n+\n+    record EncryptedPayload(byte[] iv, byte[] ciphertext) {\n+        EncryptedPayload {\n+            iv = iv.clone();\n+            ciphertext = ciphertext.clone();\n+        }\n+\n+        @Override\n+        public byte[] iv() {\n+            return iv.clone();\n+        }\n+\n+        @Override\n+        public byte[] ciphertext() {\n+            return ciphertext.clone();\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/EncryptionMetadata.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.util.Base64;\n+import java.util.Map;\n+import java.util.Objects;\n+\n+final class EncryptionMetadata {\n+    private static final String FORMAT_VERSION = \"1\";\n+    private static final String VERSION = \"encversion\";\n+    private static final String CONTENT_ALGORITHM = \"encalgorithm\";\n+    private static final String IV = \"enciv\";\n+    private static final String KEY_ID = \"enckeyid\";\n+    private static final String WRAP_ALGORITHM = \"encwrapalgorithm\";\n+    private static final String WRAPPED_KEY = \"encwrappedkey\";\n+\n+    private final ProtectedDataKey protectedDataKey;\n+    private final byte[] iv;\n+\n+    private EncryptionMetadata(ProtectedDataKey protectedDataKey, byte[] iv) {\n+        this.protectedDataKey = Objects.requireNonNull(protectedDataKey, \"protectedDataKey\");\n+        this.iv = Objects.requireNonNull(iv, \"iv\").clone();\n+    }\n+\n+    static EncryptionMetadata create(ProtectedDataKey protectedDataKey, byte[] iv) {\n+        return new EncryptionMetadata(protectedDataKey, iv);\n+    }\n+\n+    static EncryptionMetadata parse(Map\u003cString, String\u003e metadata) {\n+        try {\n+            String version = required(metadata, VERSION);\n+            if (!FORMAT_VERSION.equals(version)) {\n+                throw new BlobEncryptionException(\n+                    \"Unsupported encryption metadata version: \" + version\n+                );\n+            }\n+\n+            String contentAlgorithm = required(metadata, CONTENT_ALGORITHM);\n+            if (!LocalAesGcm.ALGORITHM.equals(contentAlgorithm)) {\n+                throw new BlobEncryptionException(\n+                    \"Unsupported content encryption algorithm: \" + contentAlgorithm\n+                );\n+            }\n+\n+            ProtectedDataKey protectedKey = new ProtectedDataKey(\n+                required(metadata, KEY_ID),\n+                required(metadata, WRAP_ALGORITHM),\n+                Base64.getDecoder().decode(required(metadata, WRAPPED_KEY))\n+            );\n+            byte[] iv = Base64.getDecoder().decode(required(metadata, IV));\n+            return new EncryptionMetadata(protectedKey, iv);\n+        } catch (IllegalArgumentException e) {\n+            throw new BlobEncryptionException(\"Blob encryption metadata is invalid\", e);\n+        }\n+    }\n+\n+    Map\u003cString, String\u003e toMap() {\n+        return Map.of(\n+            VERSION, FORMAT_VERSION,\n+            CONTENT_ALGORITHM, LocalAesGcm.ALGORITHM,\n+            IV, Base64.getEncoder().encodeToString(iv),\n+            KEY_ID, protectedDataKey.keyId(),\n+            WRAP_ALGORITHM, protectedDataKey.algorithm(),\n+            WRAPPED_KEY, Base64.getEncoder().encodeToString(protectedDataKey.wrappedKey())\n+        );\n+    }\n+\n+    ProtectedDataKey protectedDataKey() {\n+        return protectedDataKey;\n+    }\n+\n+    byte[] iv() {\n+        return iv.clone();\n+    }\n+\n+    byte[] authenticatedData() {\n+        return authenticatedData(protectedDataKey);\n+    }\n+\n+    static byte[] authenticatedData(ProtectedDataKey protectedDataKey) {\n+        String canonicalMetadata = String.join(\n+            \"\\n\",\n+            FORMAT_VERSION,\n+            LocalAesGcm.ALGORITHM,\n+            protectedDataKey.keyId(),\n+            protectedDataKey.algorithm(),\n+            Base64.getEncoder().encodeToString(protectedDataKey.wrappedKey())\n+        );\n+        return canonicalMetadata.getBytes(StandardCharsets.UTF_8);\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e metadata, String name) {\n+        String value = metadata.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new BlobEncryptionException(\"Missing blob encryption metadata: \" + name);\n+        }\n+        return value;\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/UploadResult.java\n+package com.example.encryptedblob;\n+\n+import java.util.Base64;\n+import java.util.Objects;\n+\n+public record UploadResult(String keyId, byte[] wrappedDataKey) {\n+    public UploadResult {\n+        Objects.requireNonNull(keyId, \"keyId\");\n+        wrappedDataKey = Objects.requireNonNull(wrappedDataKey, \"wrappedDataKey\").clone();\n+    }\n+\n+    @Override\n+    public byte[] wrappedDataKey() {\n+        return wrappedDataKey.clone();\n+    }\n+\n+    public String wrappedDataKeyBase64() {\n+        return Base64.getEncoder().encodeToString(wrappedDataKey);\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/SyncEncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.core.util.Context;\n+import com.azure.storage.blob.BlobClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.models.BlobHttpHeaders;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n+\n+import java.io.IOException;\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class SyncEncryptedBlobClient {\n+    private final BlobContainerClient containerClient;\n+    private final SyncKeyManagementClient keyManagementClient;\n+\n+    public SyncEncryptedBlobClient(\n+        BlobContainerClient containerClient,\n+        SyncKeyManagementClient keyManagementClient\n+    ) {\n+        this.containerClient = Objects.requireNonNull(containerClient, \"containerClient\");\n+        this.keyManagementClient =\n+            Objects.requireNonNull(keyManagementClient, \"keyManagementClient\");\n+    }\n+\n+    public UploadResult upload(Path source, String blobName) throws IOException {\n+        return upload(Files.readAllBytes(source), blobName);\n+    }\n+\n+    public UploadResult upload(byte[] plaintext, String blobName) {\n+        Objects.requireNonNull(plaintext, \"plaintext\");\n+        Objects.requireNonNull(blobName, \"blobName\");\n+\n+        try (EnvelopeKey envelopeKey = keyManagementClient.generateAndWrapDataKey()) {\n+            ProtectedDataKey protectedKey = envelopeKey.protectedDataKey();\n+            byte[] authenticatedData = EncryptionMetadata.authenticatedData(protectedKey);\n+            LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(\n+                envelopeKey.dataKey(),\n+                plaintext,\n+                authenticatedData\n+            );\n+            EncryptionMetadata metadata =\n+                EncryptionMetadata.create(protectedKey, encrypted.iv());\n+            BlobClient blobClient = containerClient.getBlobClient(blobName);\n+            BlobParallelUploadOptions options = new BlobParallelUploadOptions(\n+                BinaryData.fromBytes(encrypted.ciphertext())\n+            )\n+                .setMetadata(metadata.toMap())\n+                .setHeaders(new BlobHttpHeaders().setContentType(\"application/octet-stream\"));\n+\n+            try {\n+                blobClient.uploadWithResponse(options, null, Context.NONE);\n+            } catch (BlobStorageException e) {\n+                throw BlobEncryptionException.storageFailure(\"upload\", blobName, e);\n+            }\n+            return new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey());\n+        }\n+    }\n+\n+    public void download(String blobName, Path destination) throws IOException {\n+        Files.write(destination, download(blobName));\n+    }\n+\n+    public byte[] download(String blobName) {\n+        Objects.requireNonNull(blobName, \"blobName\");\n+        BlobClient blobClient = containerClient.getBlobClient(blobName);\n+\n+        EncryptionMetadata metadata;\n+        byte[] ciphertext;\n+        try {\n+            metadata = EncryptionMetadata.parse(blobClient.getProperties().getMetadata());\n+            ciphertext = blobClient.downloadContent().toBytes();\n+        } catch (BlobStorageException e) {\n+            throw BlobEncryptionException.storageFailure(\"download\", blobName, e);\n+        }\n+\n+        byte[] dataKey = keyManagementClient.unwrapDataKey(metadata.protectedDataKey());\n+        try {\n+            return LocalAesGcm.decrypt(\n+                dataKey,\n+                ciphertext,\n+                metadata.iv(),\n+                metadata.authenticatedData()\n+            );\n+        } finally {\n+            Arrays.fill(dataKey, (byte) 0);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.util.BinaryData;\n+import com.azure.storage.blob.BlobAsyncClient;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.models.BlobHttpHeaders;\n+import com.azure.storage.blob.models.BlobStorageException;\n+import com.azure.storage.blob.options.BlobParallelUploadOptions;\n+import reactor.core.publisher.Mono;\n+import reactor.core.scheduler.Schedulers;\n+\n+import java.nio.file.Files;\n+import java.nio.file.Path;\n+import java.util.Arrays;\n+import java.util.Objects;\n+\n+public final class AsyncEncryptedBlobClient {\n+    private final BlobContainerAsyncClient containerClient;\n+    private final AsyncKeyManagementClient keyManagementClient;\n+\n+    public AsyncEncryptedBlobClient(\n+        BlobContainerAsyncClient containerClient,\n+        AsyncKeyManagementClient keyManagementClient\n+    ) {\n+        this.containerClient = Objects.requireNonNull(containerClient, \"containerClient\");\n+        this.keyManagementClient =\n+            Objects.requireNonNull(keyManagementClient, \"keyManagementClient\");\n+    }\n+\n+    public Mono\u003cUploadResult\u003e upload(Path source, String blobName) {\n+        return Mono.fromCallable(() -\u003e Files.readAllBytes(source))\n+            .subscribeOn(Schedulers.boundedElastic())\n+            .flatMap(bytes -\u003e upload(bytes, blobName));\n+    }\n+\n+    public Mono\u003cUploadResult\u003e upload(byte[] plaintext, String blobName) {\n+        return Mono.defer(() -\u003e {\n+            Objects.requireNonNull(plaintext, \"plaintext\");\n+            Objects.requireNonNull(blobName, \"blobName\");\n+\n+            return keyManagementClient.generateAndWrapDataKey()\n+                .flatMap(envelopeKey -\u003e {\n+                    ProtectedDataKey protectedKey = envelopeKey.protectedDataKey();\n+                    LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(\n+                        envelopeKey.dataKey(),\n+                        plaintext,\n+                        EncryptionMetadata.authenticatedData(protectedKey)\n+                    );\n+                    EncryptionMetadata metadata =\n+                        EncryptionMetadata.create(protectedKey, encrypted.iv());\n+                    BlobParallelUploadOptions options = new BlobParallelUploadOptions(\n+                        BinaryData.fromBytes(encrypted.ciphertext())\n+                    )\n+                        .setMetadata(metadata.toMap())\n+                        .setHeaders(\n+                            new BlobHttpHeaders().setContentType(\"application/octet-stream\")\n+                        );\n+\n+                    return containerClient.getBlobAsyncClient(blobName)\n+                        .uploadWithResponse(options)\n+                        .thenReturn(\n+                            new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey())\n+                        )\n+                        .doFinally(ignored -\u003e envelopeKey.close());\n+                });\n+        }).onErrorMap(\n+            BlobStorageException.class,\n+            e -\u003e BlobEncryptionException.storageFailure(\"upload\", blobName, e)\n+        );\n+    }\n+\n+    public Mono\u003cVoid\u003e download(String blobName, Path destination) {\n+        return download(blobName)\n+            .flatMap(bytes -\u003e Mono.fromCallable(() -\u003e Files.write(destination, bytes)))\n+            .subscribeOn(Schedulers.boundedElastic())\n+            .then();\n+    }\n+\n+    public Mono\u003cbyte[]\u003e download(String blobName) {\n+        return Mono.defer(() -\u003e {\n+            Objects.requireNonNull(blobName, \"blobName\");\n+            BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);\n+\n+            return blobClient.getProperties()\n+                .flatMap(properties -\u003e {\n+                    EncryptionMetadata metadata =\n+                        EncryptionMetadata.parse(properties.getMetadata());\n+                    return blobClient.downloadContent()\n+                        .flatMap(content -\u003e decrypt(content.toBytes(), metadata));\n+                });\n+        }).onErrorMap(\n+            BlobStorageException.class,\n+            e -\u003e BlobEncryptionException.storageFailure(\"download\", blobName, e)\n+        );\n+    }\n+\n+    private Mono\u003cbyte[]\u003e decrypt(byte[] ciphertext, EncryptionMetadata metadata) {\n+        return keyManagementClient.unwrapDataKey(metadata.protectedDataKey())\n+            .map(dataKey -\u003e {\n+                try {\n+                    return LocalAesGcm.decrypt(\n+                        dataKey,\n+                        ciphertext,\n+                        metadata.iv(),\n+                        metadata.authenticatedData()\n+                    );\n+                } finally {\n+                    Arrays.fill(dataKey, (byte) 0);\n+                }\n+            });\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/AzureConfiguration.java\n+package com.example.encryptedblob;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.ManagedIdentityCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.keys.KeyAsyncClient;\n+import com.azure.security.keyvault.keys.KeyClient;\n+import com.azure.security.keyvault.keys.KeyClientBuilder;\n+import com.azure.storage.blob.BlobContainerAsyncClient;\n+import com.azure.storage.blob.BlobContainerClient;\n+import com.azure.storage.blob.BlobServiceClientBuilder;\n+\n+import java.net.URI;\n+import java.net.URISyntaxException;\n+import java.util.Map;\n+\n+public final class AzureConfiguration {\n+    private final ManagedIdentityCredential credential;\n+    private final BlobContainerClient blobContainerClient;\n+    private final BlobContainerAsyncClient blobContainerAsyncClient;\n+    private final KeyClient keyClient;\n+    private final KeyAsyncClient keyAsyncClient;\n+    private final String keyName;\n+\n+    private AzureConfiguration(\n+        ManagedIdentityCredential credential,\n+        BlobContainerClient blobContainerClient,\n+        BlobContainerAsyncClient blobContainerAsyncClient,\n+        KeyClient keyClient,\n+        KeyAsyncClient keyAsyncClient,\n+        String keyName\n+    ) {\n+        this.credential = credential;\n+        this.blobContainerClient = blobContainerClient;\n+        this.blobContainerAsyncClient = blobContainerAsyncClient;\n+        this.keyClient = keyClient;\n+        this.keyAsyncClient = keyAsyncClient;\n+        this.keyName = keyName;\n+    }\n+\n+    public static AzureConfiguration fromEnvironment() {\n+        Map\u003cString, String\u003e environment = System.getenv();\n+        String storageEndpoint = required(environment, \"AZURE_STORAGE_BLOB_ENDPOINT\");\n+        String containerName = required(environment, \"AZURE_STORAGE_CONTAINER\");\n+        String vaultUrl = required(environment, \"AZURE_KEY_VAULT_URL\");\n+        String keyName = required(environment, \"AZURE_KEY_NAME\");\n+        validateHttpsEndpoint(\"AZURE_STORAGE_BLOB_ENDPOINT\", storageEndpoint);\n+        validateHttpsEndpoint(\"AZURE_KEY_VAULT_URL\", vaultUrl);\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder =\n+            new ManagedIdentityCredentialBuilder();\n+        String managedIdentityClientId = environment.get(\"AZURE_CLIENT_ID\");\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        ManagedIdentityCredential credential = credentialBuilder.build();\n+\n+        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()\n+            .endpoint(storageEndpoint)\n+            .credential(credential);\n+        BlobContainerClient syncContainer = blobBuilder\n+            .buildClient()\n+            .getBlobContainerClient(containerName);\n+        BlobContainerAsyncClient asyncContainer = blobBuilder\n+            .buildAsyncClient()\n+            .getBlobContainerAsyncClient(containerName);\n+\n+        KeyClientBuilder keyBuilder = new KeyClientBuilder()\n+            .vaultUrl(vaultUrl)\n+            .credential(credential);\n+        return new AzureConfiguration(\n+            credential,\n+            syncContainer,\n+            asyncContainer,\n+            keyBuilder.buildClient(),\n+            keyBuilder.buildAsyncClient(),\n+            keyName\n+        );\n+    }\n+\n+    public TokenCredential credential() {\n+        return credential;\n+    }\n+\n+    public BlobContainerClient blobContainerClient() {\n+        return blobContainerClient;\n+    }\n+\n+    public BlobContainerAsyncClient blobContainerAsyncClient() {\n+        return blobContainerAsyncClient;\n+    }\n+\n+    public KeyClient keyClient() {\n+        return keyClient;\n+    }\n+\n+    public KeyAsyncClient keyAsyncClient() {\n+        return keyAsyncClient;\n+    }\n+\n+    public String keyName() {\n+        return keyName;\n+    }\n+\n+    private static String required(Map\u003cString, String\u003e environment, String name) {\n+        String value = environment.get(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is missing: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static void validateHttpsEndpoint(String name, String value) {\n+        try {\n+            URI uri = new URI(value);\n+            if (!\"https\".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {\n+                throw new IllegalStateException(name + \" must be an absolute HTTPS URL\");\n+            }\n+        } catch (URISyntaxException e) {\n+            throw new IllegalStateException(name + \" is not a valid URL\", e);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/encryptedblob/Main.java\n+package com.example.encryptedblob;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.time.Duration;\n+\n+public final class Main {\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();\n+\n+        SyncKeyManagementClient syncKeys = new SyncKeyManagementClient(\n+            configuration.keyClient(),\n+            configuration.credential(),\n+            configuration.keyName()\n+        );\n+        SyncEncryptedBlobClient syncBlobs = new SyncEncryptedBlobClient(\n+            configuration.blobContainerClient(),\n+            syncKeys\n+        );\n+\n+        byte[] syncPlaintext =\n+            \"Hello from the synchronous encrypted uploader.\".getBytes(StandardCharsets.UTF_8);\n+        UploadResult syncUpload = syncBlobs.upload(syncPlaintext, \"sync-encrypted-demo.bin\");\n+        byte[] syncDecrypted = syncBlobs.download(\"sync-encrypted-demo.bin\");\n+        printResult(\"Sync\", syncUpload, syncDecrypted);\n+\n+        AsyncKeyManagementClient asyncKeys = new AsyncKeyManagementClient(\n+            configuration.keyAsyncClient(),\n+            configuration.credential(),\n+            configuration.keyName()\n+        );\n+        AsyncEncryptedBlobClient asyncBlobs = new AsyncEncryptedBlobClient(\n+            configuration.blobContainerAsyncClient(),\n+            asyncKeys\n+        );\n+\n+        byte[] asyncPlaintext =\n+            \"Hello from the asynchronous encrypted uploader.\".getBytes(StandardCharsets.UTF_8);\n+        AsyncRoundTrip asyncRoundTrip = asyncBlobs\n+            .upload(asyncPlaintext, \"async-encrypted-demo.bin\")\n+            .flatMap(upload -\u003e asyncBlobs\n+                .download(\"async-encrypted-demo.bin\")\n+                .map(decrypted -\u003e new AsyncRoundTrip(upload, decrypted)))\n+            .block(Duration.ofMinutes(2));\n+\n+        if (asyncRoundTrip == null) {\n+            throw new IllegalStateException(\"The asynchronous round trip completed without a result\");\n+        }\n+        printResult(\"Async\", asyncRoundTrip.upload(), asyncRoundTrip.decrypted());\n+    }\n+\n+    private static void printResult(String label, UploadResult upload, byte[] decrypted) {\n+        System.out.println(label + \" Key Vault key ID: \" + upload.keyId());\n+        System.out.println(label + \" wrapped DEK (base64): \" + upload.wrappedDataKeyBase64());\n+        System.out.println(\n+            label + \" decrypted output: \" + new String(decrypted, StandardCharsets.UTF_8)\n+        );\n+    }\n+\n+    private record AsyncRoundTrip(UploadResult upload, byte[] decrypted) {\n+    }\n+}\n*** Add File: src/test/java/com/example/encryptedblob/LocalAesGcmTest.java\n+package com.example.encryptedblob;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.nio.charset.StandardCharsets;\n+import java.security.SecureRandom;\n+\n+import static org.junit.jupiter.api.Assertions.assertArrayEquals;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+\n+class LocalAesGcmTest {\n+    @Test\n+    void roundTripsAuthenticatedCiphertext() {\n+        byte[] key = randomBytes(32);\n+        byte[] plaintext = \"client-side encrypted\".getBytes(StandardCharsets.UTF_8);\n+        byte[] metadata = \"authenticated metadata\".getBytes(StandardCharsets.UTF_8);\n+\n+        LocalAesGcm.EncryptedPayload encrypted =\n+            LocalAesGcm.encrypt(key, plaintext, metadata);\n+\n+        assertArrayEquals(\n+            plaintext,\n+            LocalAesGcm.decrypt(key, encrypted.ciphertext(), encrypted.iv(), metadata)\n+        );\n+    }\n+\n+    @Test\n+    void rejectsModifiedAuthenticatedMetadata() {\n+        byte[] key = randomBytes(32);\n+        byte[] plaintext = \"client-side encrypted\".getBytes(StandardCharsets.UTF_8);\n+        byte[] metadata = \"original metadata\".getBytes(StandardCharsets.UTF_8);\n+        LocalAesGcm.EncryptedPayload encrypted =\n+            LocalAesGcm.encrypt(key, plaintext, metadata);\n+\n+        assertThrows(\n+            BlobEncryptionException.class,\n+            () -\u003e LocalAesGcm.decrypt(\n+                key,\n+                encrypted.ciphertext(),\n+                encrypted.iv(),\n+                \"modified metadata\".getBytes(StandardCharsets.UTF_8)\n+            )\n+        );\n+    }\n+\n+    private static byte[] randomBytes(int size) {\n+        byte[] bytes = new byte[size];\n+        new SecureRandom().nextBytes(bytes);\n+        return bytes;\n+    }\n+}\n*** End Patch\n"
```

### 🔧 22. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress clean verify","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 23. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/encryptedblob/AsyncEncryptedBlobClient.java\n@@\n             Objects.requireNonNull(blobName, \"blobName\");\n \n             return keyManagementClient.generateAndWrapDataKey()\n-                .flatMap(envelopeKey -\u003e {\n-                    ProtectedDataKey protectedKey = envelopeKey.protectedDataKey();\n-                    LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(\n-                        envelopeKey.dataKey(),\n-                        plaintext,\n-                        EncryptionMetadata.authenticatedData(protectedKey)\n-                    );\n-                    EncryptionMetadata metadata =\n-                        EncryptionMetadata.create(protectedKey, encrypted.iv());\n-                    BlobParallelUploadOptions options = new BlobParallelUploadOptions(\n-                        BinaryData.fromBytes(encrypted.ciphertext())\n-                    )\n-                        .setMetadata(metadata.toMap())\n-                        .setHeaders(\n-                            new BlobHttpHeaders().setContentType(\"application/octet-stream\")\n-                        );\n-\n-                    return containerClient.getBlobAsyncClient(blobName)\n-                        .uploadWithResponse(options)\n-                        .thenReturn(\n-                            new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey())\n-                        )\n-                        .doFinally(ignored -\u003e envelopeKey.close());\n-                });\n+                .flatMap(envelopeKey -\u003e Mono.using(\n+                    () -\u003e envelopeKey,\n+                    key -\u003e {\n+                        ProtectedDataKey protectedKey = key.protectedDataKey();\n+                        LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(\n+                            key.dataKey(),\n+                            plaintext,\n+                            EncryptionMetadata.authenticatedData(protectedKey)\n+                        );\n+                        EncryptionMetadata metadata =\n+                            EncryptionMetadata.create(protectedKey, encrypted.iv());\n+                        BlobParallelUploadOptions options = new BlobParallelUploadOptions(\n+                            BinaryData.fromBytes(encrypted.ciphertext())\n+                        )\n+                            .setMetadata(metadata.toMap())\n+                            .setHeaders(\n+                                new BlobHttpHeaders().setContentType(\"application/octet-stream\")\n+                            );\n+\n+                        return containerClient.getBlobAsyncClient(blobName)\n+                            .uploadWithResponse(options)\n+                            .thenReturn(\n+                                new UploadResult(\n+                                    protectedKey.keyId(),\n+                                    protectedKey.wrappedKey()\n+                                )\n+                            );\n+                    },\n+                    EnvelopeKey::close\n+                ));\n         }).onErrorMap(\n*** End Patch\n"
```

### 🔧 24. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress clean verify","description":"Rebuild project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 25. rg

**Input:**

```json
{"-i":true,"-n":true,"glob":"*.{java,xml,md}","output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-tools-with-azure-tools-2361299366","pattern":"connectionString|clientSecret|accountKey|AZURE_CLIENT_SECRET|StorageSharedKeyCredential|DefaultAzureCredential"}
```

### 🔧 26. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-java-encrypted-uploader-java-azure-tools-with-azure-tools-2361299366","pattern":"src/**/*.java"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\encryptedblob\AsyncEncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\AsyncKeyManagementClient.java`
- `src\main\java\com\example\encryptedblob\AzureConfiguration.java`
- `src\main\java\com\example\encryptedblob\BlobEncryptionException.java`
- `src\main\java\com\example\encryptedblob\EncryptionMetadata.java`
- `src\main\java\com\example\encryptedblob\EnvelopeKey.java`
- `src\main\java\com\example\encryptedblob\KeyManagementException.java`
- `src\main\java\com\example\encryptedblob\LocalAesGcm.java`
- `src\main\java\com\example\encryptedblob\Main.java`
- `src\main\java\com\example\encryptedblob\ProtectedDataKey.java`
- `src\main\java\com\example\encryptedblob\SyncEncryptedBlobClient.java`
- `src\main\java\com\example\encryptedblob\SyncKeyManagementClient.java`
- `src\main\java\com\example\encryptedblob\UploadResult.java`
- `src\test\java\com\example\encryptedblob\LocalAesGcmTest.java`

## Copilot Response

I’m grounding the implementation in the current Azure SDK guidance, then I’ll create and compile the project locally without contacting Azure.

The design will use AES-256-GCM locally and RSA-OAEP-256 only to wrap the per-upload DEK in Key Vault; blob metadata will pin the exact Key Vault key version for safe rotation.

I’m implementing the sync and Reactor-based async paths with version-pinned key IDs, authenticated metadata, explicit key zeroization, and service-specific error context.

Created the complete Java 17 Maven project in the current directory.

It includes:

- Sync and async Key Vault envelope-key management using RSA-OAEP-256.
- Sync and async Blob Storage upload/download with local AES-256-GCM encryption.
- Version-pinned Key Vault key IDs and authenticated encryption metadata.
- Managed identity authentication with one shared credential instance.
- Explicit in-memory DEK zeroization.
- Service-specific error reporting for missing blobs, disabled keys, and permission failures.
- A `Main` demo for both round trips.
- Local AES-GCM tests and usage documentation in `README.md`.

Required environment variables are `AZURE_STORAGE_BLOB_ENDPOINT`, `AZURE_STORAGE_CONTAINER`, `AZURE_KEY_VAULT_URL`, and `AZURE_KEY_NAME`; `AZURE_CLIENT_ID` optionally selects a user-assigned identity.

References: [Key Vault Keys Java SDK](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-keys-readme), [Blob Storage Java SDK](https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme), and [Blob client-side encryption](https://learn.microsoft.com/azure/storage/blobs/client-side-encryption).

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Correctly limits explicit resource cleanup to the AutoCloseable EnvelopeKey, ensuring local data-encryption-key material is zeroized.

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
hyoka run --prompt-id storage-dp-java-encrypted-uploader --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
