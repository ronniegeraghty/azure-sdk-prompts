# Evaluation Report: key-vault-dp-java-secret-config

**Config:** java-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 597.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-java-secret-config` |
| Config | java-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 20/22 |
| Duration | 597.4s |
| Timestamp | 2026-08-26T13:02:53Z |
| Files Generated | 14 |
| Event Count | 17218 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 282.4s |
| Review | 313.7s |
| **Total** | **597.4s** |

## Configuration

- **name:** java-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-java, azure-ai-anomalydetector-java, azure-ai-contentsafety-java, azure-ai-formrecognizer-java, azure-ai-projects-java, azure-ai-vision-imageanalysis-java, azure-ai-voicelive-java, azure-appconfiguration-java, azure-communication-callautomation-java, azure-communication-callingserver-java, azure-communication-chat-java, azure-communication-common-java, azure-communication-sms-java, azure-compute-batch-java, azure-cosmos-java, azure-data-tables-java, azure-eventgrid-java, azure-eventhub-java, azure-identity-java, azure-messaging-webpubsub-java, azure-monitor-ingestion-java, azure-monitor-opentelemetry-exporter-java, azure-monitor-query-java, azure-security-keyvault-keys-java, azure-security-keyvault-secrets-java, azure-storage-blob-java, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-security-keyvault-secrets-java, azure-identity-java |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=791001 out=13433 |
| Turn Count | 14 |

## Prompt Sent

```
Create a small Java 17 Maven project that implements an application configuration provider backed by Azure Key Vault.

The project needs:

- A **secret provider class** (both sync and async versions) that retrieves secrets from Key Vault by name, with graceful handling when a secret doesn't exist (return a default value instead of crashing). It should also be able to retrieve a specific version of a secret (not just the latest), and inspect a secret's expiry date so the caller can tell if a secret is about to expire.

- A **caching layer** on top of the provider that stores secret values in memory after first retrieval. It should support bulk-loading a predefined set of required config keys at startup, on-demand refresh of individual keys, and automatic re-fetch of any secret whose expiry date is within a configurable warning window (e.g., 7 days out).

- A **configuration/factory class** that connects securely to the Key Vault using the vault URL from an environment variable. The application runs in Azure and should authenticate using managed identity — no client secrets or certificates in code.

- A **secret rotation helper** that safely rotates a secret: delete the old secret, ensure the deletion is fully complete, then create the new secret with an updated value and expiry date. The rotation must be safe — don't assume deletion is instantaneous, since Key Vault's soft-delete feature means the secret may not be immediately gone.

- A **Main class** that demos both implementations: loading several config keys at startup, reading them from cache, refreshing one, printing a warning if any secret is near expiry, and performing a secret rotation (delete old, wait for completion, create new). Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-security-keyvault-secrets-java"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-identity-java"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local Java 17 Maven sample using Azure Key Vault Secrets with sync/async clients, managed identity authentication, caching, expiry checks, and safe delete-and-recreate rotation. No Azure deployment or live resource modification.","learn":true}
```

### 🔧 4. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Generate a Java 17 Azure Key Vault Secrets application using managed identity, sync/async APIs, caching, expiry handling, and safe rotation. Local code only; do not deploy or modify resources.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java\\references\\examples.md"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java\\references\\examples.md"}
```

### 🔧 7. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-java-secret-config-java-azure-tools-with-azure-tools-358467483","pattern":"**/*"}
```

### 🔧 8. glob

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-security-keyvault-secrets-java","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-java\\skills\\azure-identity-java"],"pattern":"**/acceptance-criteria.md"}
```

### 🔧 9. web_fetch

**Input:**

```json
{"max_length":4000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/maven-metadata.xml"}
```

### 🔧 10. web_fetch

**Input:**

```json
{"max_length":3000,"url":"https://repo1.maven.org/maven2/org/junit/junit-bom/maven-metadata.xml"}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Java Azure Key Vault Secrets guidance for managed identity client creation, secret retrieval/version/expiry, and delete/purge polling behavior.","learn":true}
```

### 🔧 12. powershell

**Input:**

```json
{"command":"$java = Get-Command java -ErrorAction SilentlyContinue; $mvn = Get-Command mvn -ErrorAction SilentlyContinue; [pscustomobject]@{Java=if($java){\u0026 java -version 2\u003e\u00261 | Select-Object -First 1}else{'missing'}; Maven=if($mvn){\u0026 mvn -version 2\u003e\u00261 | Select-Object -First 1}else{'missing'}} | Format-List","description":"Check Java and Maven availability"}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Retrieve latest official Java samples for Azure Key Vault SecretClient/SecretAsyncClient with ManagedIdentityCredential, getSecret by version, expiry properties, and beginDeleteSecret/purgeDeletedSecret.","parameters":{"language":"java","query":"Azure Key Vault Secrets Java SecretClient SecretAsyncClient ManagedIdentityCredential getSecret version expiresOn beginDeleteSecret purgeDeletedSecret"}}
```

### 🔧 14. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official documentation for Java Key Vault secret deletion polling, soft-delete purge behavior, and managed identity authentication.","parameters":{"query":"Azure Key Vault Secrets Java beginDeleteSecret waitForCompletion purge deleted secret soft delete managed identity"}}
```

### 🔧 15. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official Java Key Vault Secrets quickstart identified by documentation search.","parameters":{"url":"https://learn.microsoft.com/en-us/azure/key-vault/secrets/quick-create-java"}}
```

### 🔧 16. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003ekey-vault-config-provider\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0-SNAPSHOT\u003c/version\u003e\n+    \u003cname\u003eKey Vault configuration provider\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.sdk.bom.version\u003e1.3.8\u003c/azure.sdk.bom.version\u003e\n+        \u003cjunit.version\u003e5.14.3\u003c/junit.version\u003e\n+        \u003cmockito.version\u003e5.21.0\u003c/mockito.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencyManagement\u003e\n+        \u003cdependencies\u003e\n+            \u003cdependency\u003e\n+                \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+                \u003cartifactId\u003eazure-sdk-bom\u003c/artifactId\u003e\n+                \u003cversion\u003e${azure.sdk.bom.version}\u003c/version\u003e\n+                \u003ctype\u003epom\u003c/type\u003e\n+                \u003cscope\u003eimport\u003c/scope\u003e\n+            \u003c/dependency\u003e\n+        \u003c/dependencies\u003e\n+    \u003c/dependencyManagement\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-security-keyvault-secrets\u003c/artifactId\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.mockito\u003c/groupId\u003e\n+            \u003cartifactId\u003emockito-core\u003c/artifactId\u003e\n+            \u003cversion\u003e${mockito.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eio.projectreactor\u003c/groupId\u003e\n+            \u003cartifactId\u003ereactor-test\u003c/artifactId\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.4\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.6.2\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.keyvault.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/keyvault/SecretSnapshot.java\n+package com.example.keyvault;\n+\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public record SecretSnapshot(\n+    String name,\n+    String value,\n+    String version,\n+    OffsetDateTime expiresOn,\n+    boolean defaultValue\n+) {\n+    public SecretSnapshot {\n+        Objects.requireNonNull(name, \"name\");\n+        Objects.requireNonNull(value, \"value\");\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/SecretExpiry.java\n+package com.example.keyvault;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+final class SecretExpiry {\n+    private SecretExpiry() {\n+    }\n+\n+    static boolean isWithin(SecretSnapshot secret, Duration warningWindow, Clock clock) {\n+        Objects.requireNonNull(secret, \"secret\");\n+        Objects.requireNonNull(warningWindow, \"warningWindow\");\n+        Objects.requireNonNull(clock, \"clock\");\n+        if (warningWindow.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+\n+        OffsetDateTime expiresOn = secret.expiresOn();\n+        return expiresOn != null\n+            \u0026\u0026 !expiresOn.isAfter(OffsetDateTime.now(clock).plus(warningWindow));\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/SyncKeyVaultSecretProvider.java\n+package com.example.keyvault;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import java.util.Objects;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+\n+public class SyncKeyVaultSecretProvider {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(SyncKeyVaultSecretProvider.class);\n+\n+    private final SecretClient client;\n+\n+    public SyncKeyVaultSecretProvider(SecretClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public SecretSnapshot getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    public SecretSnapshot getSecret(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+        try {\n+            KeyVaultSecret secret = version == null\n+                ? client.getSecret(name)\n+                : client.getSecret(name, version);\n+            return toSnapshot(secret);\n+        } catch (ResourceNotFoundException exception) {\n+            LOGGER.warn(\"Key Vault secret '{}' was not found; using its configured default\", name);\n+            return new SecretSnapshot(name, defaultValue, version, null, true);\n+        }\n+    }\n+\n+    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {\n+        return new SecretSnapshot(\n+            secret.getName(),\n+            secret.getValue(),\n+            secret.getProperties().getVersion(),\n+            secret.getProperties().getExpiresOn(),\n+            false\n+        );\n+    }\n+\n+    private static void requireText(String value, String label) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(label + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/AsyncKeyVaultSecretProvider.java\n+package com.example.keyvault;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import java.util.Objects;\n+import org.slf4j.Logger;\n+import org.slf4j.LoggerFactory;\n+import reactor.core.publisher.Mono;\n+\n+public class AsyncKeyVaultSecretProvider {\n+    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncKeyVaultSecretProvider.class);\n+\n+    private final SecretAsyncClient client;\n+\n+    public AsyncKeyVaultSecretProvider(SecretAsyncClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e getSecret(String name, String defaultValue) {\n+        return getSecret(name, null, defaultValue);\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e getSecret(String name, String version, String defaultValue) {\n+        requireText(name, \"name\");\n+        Objects.requireNonNull(defaultValue, \"defaultValue\");\n+\n+        Mono\u003cKeyVaultSecret\u003e request = version == null\n+            ? client.getSecret(name)\n+            : client.getSecret(name, version);\n+        return request\n+            .map(AsyncKeyVaultSecretProvider::toSnapshot)\n+            .onErrorResume(ResourceNotFoundException.class, exception -\u003e {\n+                LOGGER.warn(\"Key Vault secret '{}' was not found; using its configured default\", name);\n+                return Mono.just(new SecretSnapshot(name, defaultValue, version, null, true));\n+            });\n+    }\n+\n+    private static SecretSnapshot toSnapshot(KeyVaultSecret secret) {\n+        return new SecretSnapshot(\n+            secret.getName(),\n+            secret.getValue(),\n+            secret.getProperties().getVersion(),\n+            secret.getProperties().getExpiresOn(),\n+            false\n+        );\n+    }\n+\n+    private static void requireText(String value, String label) {\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalArgumentException(label + \" must not be blank\");\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/SyncSecretCache.java\n+package com.example.keyvault;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public final class SyncSecretCache {\n+    private final SyncKeyVaultSecretProvider provider;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final Map\u003cString, String\u003e defaults = new ConcurrentHashMap\u003c\u003e();\n+    private final Map\u003cString, SecretSnapshot\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public SyncSecretCache(SyncKeyVaultSecretProvider provider, Duration warningWindow) {\n+        this(provider, warningWindow, Clock.systemUTC());\n+    }\n+\n+    SyncSecretCache(SyncKeyVaultSecretProvider provider, Duration warningWindow, Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public void loadRequired(Map\u003cString, String\u003e requiredSecrets) {\n+        Objects.requireNonNull(requiredSecrets, \"requiredSecrets\");\n+        requiredSecrets.forEach((name, defaultValue) -\u003e {\n+            defaults.put(name, Objects.requireNonNull(defaultValue, \"defaultValue\"));\n+            refresh(name);\n+        });\n+    }\n+\n+    public SecretSnapshot get(String name) {\n+        SecretSnapshot secret = requireCached(name);\n+        return SecretExpiry.isWithin(secret, warningWindow, clock) ? refresh(name) : secret;\n+    }\n+\n+    public SecretSnapshot refresh(String name) {\n+        String defaultValue = defaults.get(name);\n+        if (defaultValue == null) {\n+            throw new IllegalArgumentException(\"No default is registered for secret: \" + name);\n+        }\n+        SecretSnapshot refreshed = provider.getSecret(name, defaultValue);\n+        cache.put(name, refreshed);\n+        return refreshed;\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e refreshExpiring() {\n+        List\u003cSecretSnapshot\u003e refreshed = new ArrayList\u003c\u003e();\n+        List.copyOf(cache.values()).stream()\n+            .filter(secret -\u003e SecretExpiry.isWithin(secret, warningWindow, clock))\n+            .map(SecretSnapshot::name)\n+            .map(this::refresh)\n+            .forEach(refreshed::add);\n+        return List.copyOf(refreshed);\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e secretsNearExpiry() {\n+        return cache.values().stream()\n+            .filter(secret -\u003e SecretExpiry.isWithin(secret, warningWindow, clock))\n+            .toList();\n+    }\n+\n+    private SecretSnapshot requireCached(String name) {\n+        SecretSnapshot secret = cache.get(name);\n+        if (secret == null) {\n+            throw new IllegalArgumentException(\"Secret has not been loaded: \" + name);\n+        }\n+        return secret;\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/AsyncSecretCache.java\n+package com.example.keyvault;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.concurrent.ConcurrentHashMap;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncSecretCache {\n+    private final AsyncKeyVaultSecretProvider provider;\n+    private final Duration warningWindow;\n+    private final Clock clock;\n+    private final Map\u003cString, String\u003e defaults = new ConcurrentHashMap\u003c\u003e();\n+    private final Map\u003cString, SecretSnapshot\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncSecretCache(AsyncKeyVaultSecretProvider provider, Duration warningWindow) {\n+        this(provider, warningWindow, Clock.systemUTC());\n+    }\n+\n+    AsyncSecretCache(AsyncKeyVaultSecretProvider provider, Duration warningWindow, Clock clock) {\n+        this.provider = Objects.requireNonNull(provider, \"provider\");\n+        this.warningWindow = requireNonNegative(warningWindow);\n+        this.clock = Objects.requireNonNull(clock, \"clock\");\n+    }\n+\n+    public Mono\u003cVoid\u003e loadRequired(Map\u003cString, String\u003e requiredSecrets) {\n+        Objects.requireNonNull(requiredSecrets, \"requiredSecrets\");\n+        requiredSecrets.forEach((name, defaultValue) -\u003e\n+            defaults.put(name, Objects.requireNonNull(defaultValue, \"defaultValue\")));\n+        return Flux.fromIterable(requiredSecrets.keySet())\n+            .flatMap(this::refresh)\n+            .then();\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e get(String name) {\n+        return Mono.defer(() -\u003e {\n+            SecretSnapshot secret = cache.get(name);\n+            if (secret == null) {\n+                return Mono.error(new IllegalArgumentException(\"Secret has not been loaded: \" + name));\n+            }\n+            return SecretExpiry.isWithin(secret, warningWindow, clock)\n+                ? refresh(name)\n+                : Mono.just(secret);\n+        });\n+    }\n+\n+    public Mono\u003cSecretSnapshot\u003e refresh(String name) {\n+        return Mono.defer(() -\u003e {\n+            String defaultValue = defaults.get(name);\n+            if (defaultValue == null) {\n+                return Mono.error(new IllegalArgumentException(\n+                    \"No default is registered for secret: \" + name));\n+            }\n+            return provider.getSecret(name, defaultValue)\n+                .doOnNext(secret -\u003e cache.put(name, secret));\n+        });\n+    }\n+\n+    public Flux\u003cSecretSnapshot\u003e refreshExpiring() {\n+        return Flux.defer(() -\u003e Flux.fromIterable(List.copyOf(cache.values())))\n+            .filter(secret -\u003e SecretExpiry.isWithin(secret, warningWindow, clock))\n+            .map(SecretSnapshot::name)\n+            .flatMap(this::refresh);\n+    }\n+\n+    public List\u003cSecretSnapshot\u003e secretsNearExpiry() {\n+        return cache.values().stream()\n+            .filter(secret -\u003e SecretExpiry.isWithin(secret, warningWindow, clock))\n+            .toList();\n+    }\n+\n+    private static Duration requireNonNegative(Duration duration) {\n+        Objects.requireNonNull(duration, \"warningWindow\");\n+        if (duration.isNegative()) {\n+            throw new IllegalArgumentException(\"warningWindow must not be negative\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/KeyVaultClientFactory.java\n+package com.example.keyvault;\n+\n+import com.azure.identity.ManagedIdentityCredential;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.SecretClientBuilder;\n+import java.net.URI;\n+import java.net.URISyntaxException;\n+\n+public final class KeyVaultClientFactory {\n+    public static final String VAULT_URL_ENV = \"AZURE_KEY_VAULT_URL\";\n+    public static final String MANAGED_IDENTITY_CLIENT_ID_ENV = \"AZURE_MANAGED_IDENTITY_CLIENT_ID\";\n+\n+    private KeyVaultClientFactory() {\n+    }\n+\n+    public static KeyVaultClients fromEnvironment() {\n+        String vaultUrl = requireEnvironment(VAULT_URL_ENV);\n+        validateHttpsUrl(vaultUrl);\n+\n+        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();\n+        String managedIdentityClientId = System.getenv(MANAGED_IDENTITY_CLIENT_ID_ENV);\n+        if (managedIdentityClientId != null \u0026\u0026 !managedIdentityClientId.isBlank()) {\n+            credentialBuilder.clientId(managedIdentityClientId);\n+        }\n+        ManagedIdentityCredential credential = credentialBuilder.build();\n+        SecretClientBuilder clientBuilder = new SecretClientBuilder()\n+            .vaultUrl(vaultUrl)\n+            .credential(credential);\n+\n+        return new KeyVaultClients(clientBuilder.buildClient(), clientBuilder.buildAsyncClient());\n+    }\n+\n+    private static String requireEnvironment(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static void validateHttpsUrl(String value) {\n+        try {\n+            URI uri = new URI(value);\n+            if (!\"https\".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {\n+                throw new IllegalStateException(VAULT_URL_ENV + \" must be an absolute HTTPS URL\");\n+            }\n+        } catch (URISyntaxException exception) {\n+            throw new IllegalStateException(VAULT_URL_ENV + \" is not a valid URL\", exception);\n+        }\n+    }\n+\n+    public record KeyVaultClients(SecretClient syncClient, SecretAsyncClient asyncClient) {\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/SyncSecretRotator.java\n+package com.example.keyvault;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import com.azure.security.keyvault.secrets.models.SecretProperties;\n+import java.time.Duration;\n+import java.time.Instant;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+\n+public final class SyncSecretRotator {\n+    private final SecretClient client;\n+    private final Duration pollInterval;\n+    private final Duration timeout;\n+\n+    public SyncSecretRotator(SecretClient client, Duration pollInterval, Duration timeout) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.pollInterval = requirePositive(pollInterval, \"pollInterval\");\n+        this.timeout = requirePositive(timeout, \"timeout\");\n+    }\n+\n+    public KeyVaultSecret rotate(String name, String newValue, OffsetDateTime expiresOn) {\n+        requireRotationArguments(name, newValue, expiresOn);\n+        client.beginDeleteSecret(name).waitForCompletion();\n+        client.purgeDeletedSecret(name);\n+        waitUntilPurged(name);\n+\n+        return client.setSecret(new KeyVaultSecret(name, newValue)\n+            .setProperties(new SecretProperties().setExpiresOn(expiresOn)));\n+    }\n+\n+    private void waitUntilPurged(String name) {\n+        Instant deadline = Instant.now().plus(timeout);\n+        while (Instant.now().isBefore(deadline)) {\n+            try {\n+                client.getDeletedSecret(name);\n+                sleep();\n+            } catch (ResourceNotFoundException exception) {\n+                return;\n+            }\n+        }\n+        throw new IllegalStateException(\"Timed out waiting for purged secret name to become available: \" + name);\n+    }\n+\n+    private void sleep() {\n+        try {\n+            Thread.sleep(pollInterval.toMillis());\n+        } catch (InterruptedException exception) {\n+            Thread.currentThread().interrupt();\n+            throw new IllegalStateException(\"Interrupted while waiting for secret purge\", exception);\n+        }\n+    }\n+\n+    static void requireRotationArguments(String name, String newValue, OffsetDateTime expiresOn) {\n+        if (name == null || name.isBlank()) {\n+            throw new IllegalArgumentException(\"name must not be blank\");\n+        }\n+        if (newValue == null || newValue.isEmpty()) {\n+            throw new IllegalArgumentException(\"newValue must not be empty\");\n+        }\n+        Objects.requireNonNull(expiresOn, \"expiresOn\");\n+    }\n+\n+    static Duration requirePositive(Duration duration, String name) {\n+        Objects.requireNonNull(duration, name);\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(name + \" must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/AsyncSecretRotator.java\n+package com.example.keyvault;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import com.azure.security.keyvault.secrets.models.SecretProperties;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.util.Objects;\n+import reactor.core.publisher.Mono;\n+\n+public final class AsyncSecretRotator {\n+    private final SecretAsyncClient client;\n+    private final Duration pollInterval;\n+    private final Duration timeout;\n+\n+    public AsyncSecretRotator(SecretAsyncClient client, Duration pollInterval, Duration timeout) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+        this.pollInterval = SyncSecretRotator.requirePositive(pollInterval, \"pollInterval\");\n+        this.timeout = SyncSecretRotator.requirePositive(timeout, \"timeout\");\n+    }\n+\n+    public Mono\u003cKeyVaultSecret\u003e rotate(String name, String newValue, OffsetDateTime expiresOn) {\n+        SyncSecretRotator.requireRotationArguments(name, newValue, expiresOn);\n+        KeyVaultSecret replacement = new KeyVaultSecret(name, newValue)\n+            .setProperties(new SecretProperties().setExpiresOn(expiresOn));\n+\n+        return client.beginDeleteSecret(name)\n+            .last()\n+            .then(client.purgeDeletedSecret(name))\n+            .then(waitUntilPurged(name))\n+            .then(client.setSecret(replacement));\n+    }\n+\n+    private Mono\u003cVoid\u003e waitUntilPurged(String name) {\n+        return Mono.defer(() -\u003e client.getDeletedSecret(name)\n+                .then(Mono.delay(pollInterval))\n+                .then(waitUntilPurged(name)))\n+            .onErrorResume(ResourceNotFoundException.class, exception -\u003e Mono.empty())\n+            .timeout(timeout, Mono.error(new IllegalStateException(\n+                \"Timed out waiting for purged secret name to become available: \" + name)));\n+    }\n+}\n*** Add File: src/main/java/com/example/keyvault/Main.java\n+package com.example.keyvault;\n+\n+import com.example.keyvault.KeyVaultClientFactory.KeyVaultClients;\n+import java.time.Duration;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+import java.util.LinkedHashMap;\n+import java.util.List;\n+import java.util.Map;\n+\n+public final class Main {\n+    private static final Duration WARNING_WINDOW = Duration.ofDays(7);\n+    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);\n+    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(2);\n+    private static final Map\u003cString, String\u003e REQUIRED_CONFIG = requiredConfig();\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        KeyVaultClients clients = KeyVaultClientFactory.fromEnvironment();\n+        runSyncDemo(clients);\n+        runAsyncDemo(clients).block();\n+    }\n+\n+    private static void runSyncDemo(KeyVaultClients clients) {\n+        System.out.println(\"=== Synchronous implementation ===\");\n+        SyncKeyVaultSecretProvider provider = new SyncKeyVaultSecretProvider(clients.syncClient());\n+        SyncSecretCache cache = new SyncSecretCache(provider, WARNING_WINDOW);\n+\n+        cache.loadRequired(REQUIRED_CONFIG);\n+        REQUIRED_CONFIG.keySet().forEach(name -\u003e printRead(\"sync\", cache.get(name)));\n+        cache.refresh(\"api-base-url\");\n+        cache.refreshExpiring();\n+        printExpiryWarnings(cache.secretsNearExpiry());\n+\n+        String rotationName = requireEnvironment(\"DEMO_SYNC_ROTATION_SECRET_NAME\");\n+        String rotationValue = requireEnvironment(\"DEMO_SYNC_ROTATION_NEW_VALUE\");\n+        new SyncSecretRotator(clients.syncClient(), ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)\n+            .rotate(rotationName, rotationValue, OffsetDateTime.now(ZoneOffset.UTC).plusDays(90));\n+        System.out.printf(\"sync rotated '%s'%n\", rotationName);\n+    }\n+\n+    private static reactor.core.publisher.Mono\u003cVoid\u003e runAsyncDemo(KeyVaultClients clients) {\n+        System.out.println(\"=== Asynchronous implementation ===\");\n+        AsyncKeyVaultSecretProvider provider = new AsyncKeyVaultSecretProvider(clients.asyncClient());\n+        AsyncSecretCache cache = new AsyncSecretCache(provider, WARNING_WINDOW);\n+        String rotationName = requireEnvironment(\"DEMO_ASYNC_ROTATION_SECRET_NAME\");\n+        String rotationValue = requireEnvironment(\"DEMO_ASYNC_ROTATION_NEW_VALUE\");\n+\n+        return cache.loadRequired(REQUIRED_CONFIG)\n+            .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_CONFIG.keySet()))\n+            .concatMap(cache::get)\n+            .doOnNext(secret -\u003e printRead(\"async\", secret))\n+            .then(cache.refresh(\"api-base-url\"))\n+            .thenMany(cache.refreshExpiring())\n+            .then()\n+            .doOnSuccess(ignored -\u003e printExpiryWarnings(cache.secretsNearExpiry()))\n+            .then(new AsyncSecretRotator(clients.asyncClient(), ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)\n+                .rotate(rotationName, rotationValue, OffsetDateTime.now(ZoneOffset.UTC).plusDays(90)))\n+            .doOnNext(secret -\u003e System.out.printf(\"async rotated '%s'%n\", secret.getName()))\n+            .then();\n+    }\n+\n+    private static void printRead(String implementation, SecretSnapshot secret) {\n+        System.out.printf(\n+            \"%s cache read: name=%s, version=%s, default=%s%n\",\n+            implementation,\n+            secret.name(),\n+            secret.version() == null ? \"\u003cnone\u003e\" : secret.version(),\n+            secret.defaultValue()\n+        );\n+    }\n+\n+    private static void printExpiryWarnings(List\u003cSecretSnapshot\u003e expiringSecrets) {\n+        expiringSecrets.forEach(secret -\u003e System.out.printf(\n+            \"WARNING: secret '%s' expires on %s%n\", secret.name(), secret.expiresOn()));\n+    }\n+\n+    private static String requireEnvironment(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Required environment variable is not set: \" + name);\n+        }\n+        return value;\n+    }\n+\n+    private static Map\u003cString, String\u003e requiredConfig() {\n+        Map\u003cString, String\u003e config = new LinkedHashMap\u003c\u003e();\n+        config.put(\"database-connection-string\", \"jdbc:postgresql://localhost/app\");\n+        config.put(\"api-base-url\", \"https://localhost:8443\");\n+        config.put(\"feature-flags\", \"{}\");\n+        return Map.copyOf(config);\n+    }\n+}\n*** Add File: src/test/java/com/example/keyvault/SyncSecretCacheTest.java\n+package com.example.keyvault;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.mockito.Mockito.mock;\n+import static org.mockito.Mockito.verify;\n+import static org.mockito.Mockito.when;\n+\n+import java.time.Clock;\n+import java.time.Duration;\n+import java.time.Instant;\n+import java.time.OffsetDateTime;\n+import java.time.ZoneOffset;\n+import java.util.Map;\n+import org.junit.jupiter.api.Test;\n+\n+class SyncSecretCacheTest {\n+    private static final Clock CLOCK =\n+        Clock.fixed(Instant.parse(\"2026-08-26T00:00:00Z\"), ZoneOffset.UTC);\n+\n+    @Test\n+    void refreshesNearExpirySecretWhenRead() {\n+        SyncKeyVaultSecretProvider provider = mock(SyncKeyVaultSecretProvider.class);\n+        SecretSnapshot nearExpiry = new SecretSnapshot(\n+            \"api-key\", \"old\", \"v1\", OffsetDateTime.now(CLOCK).plusDays(2), false);\n+        SecretSnapshot refreshed = new SecretSnapshot(\n+            \"api-key\", \"new\", \"v2\", OffsetDateTime.now(CLOCK).plusDays(30), false);\n+        when(provider.getSecret(\"api-key\", \"fallback\")).thenReturn(nearExpiry, refreshed);\n+\n+        SyncSecretCache cache = new SyncSecretCache(provider, Duration.ofDays(7), CLOCK);\n+        cache.loadRequired(Map.of(\"api-key\", \"fallback\"));\n+\n+        assertEquals(\"new\", cache.get(\"api-key\").value());\n+        verify(provider, org.mockito.Mockito.times(2)).getSecret(\"api-key\", \"fallback\");\n+    }\n+}\n*** Add File: src/test/java/com/example/keyvault/KeyVaultSecretProviderTest.java\n+package com.example.keyvault;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+import static org.mockito.Mockito.mock;\n+import static org.mockito.Mockito.verify;\n+import static org.mockito.Mockito.when;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.security.keyvault.secrets.SecretAsyncClient;\n+import com.azure.security.keyvault.secrets.SecretClient;\n+import com.azure.security.keyvault.secrets.models.KeyVaultSecret;\n+import java.time.OffsetDateTime;\n+import org.junit.jupiter.api.Test;\n+import reactor.core.publisher.Mono;\n+import reactor.test.StepVerifier;\n+\n+class KeyVaultSecretProviderTest {\n+    @Test\n+    void syncProviderGetsSpecificVersionAndExpiry() {\n+        SecretClient client = mock(SecretClient.class);\n+        KeyVaultSecret secret = new KeyVaultSecret(\"api-key\", \"value\");\n+        OffsetDateTime expiry = OffsetDateTime.parse(\"2026-09-01T00:00:00Z\");\n+        secret.getProperties().setExpiresOn(expiry);\n+        when(client.getSecret(\"api-key\", \"v1\")).thenReturn(secret);\n+\n+        SecretSnapshot result =\n+            new SyncKeyVaultSecretProvider(client).getSecret(\"api-key\", \"v1\", \"fallback\");\n+\n+        assertEquals(\"value\", result.value());\n+        assertEquals(expiry, result.expiresOn());\n+        verify(client).getSecret(\"api-key\", \"v1\");\n+    }\n+\n+    @Test\n+    void syncProviderReturnsDefaultWhenMissing() {\n+        SecretClient client = mock(SecretClient.class);\n+        when(client.getSecret(\"missing\")).thenThrow(mock(ResourceNotFoundException.class));\n+\n+        SecretSnapshot result =\n+            new SyncKeyVaultSecretProvider(client).getSecret(\"missing\", \"fallback\");\n+\n+        assertEquals(\"fallback\", result.value());\n+        assertTrue(result.defaultValue());\n+    }\n+\n+    @Test\n+    void asyncProviderReturnsDefaultWhenMissing() {\n+        SecretAsyncClient client = mock(SecretAsyncClient.class);\n+        when(client.getSecret(\"missing\"))\n+            .thenReturn(Mono.error(mock(ResourceNotFoundException.class)));\n+\n+        Mono\u003cSecretSnapshot\u003e result =\n+            new AsyncKeyVaultSecretProvider(client).getSecret(\"missing\", \"fallback\");\n+\n+        StepVerifier.create(result)\n+            .assertNext(secret -\u003e {\n+                assertEquals(\"fallback\", secret.value());\n+                assertTrue(secret.defaultValue());\n+            })\n+            .verifyComplete();\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault configuration provider\n+\n+Small Java 17 sample with synchronous and asynchronous Key Vault secret providers, in-memory\n+caches, expiry-aware refresh, managed identity authentication, and delete-and-recreate rotation.\n+Secret values are never printed.\n+\n+## Configure and run\n+\n+The managed identity needs Key Vault data-plane permissions to read, delete, purge, and create\n+secrets. Delete-and-recreate rotation requires purge protection to be disabled; if organizational\n+policy enables purge protection, use normal Key Vault version rotation (`setSecret`) instead.\n+\n+Set these environment variables:\n+\n+```text\n+AZURE_KEY_VAULT_URL=https://\u003cvault-name\u003e.vault.azure.net\n+AZURE_MANAGED_IDENTITY_CLIENT_ID=\u003coptional-user-assigned-managed-identity-client-id\u003e\n+DEMO_SYNC_ROTATION_SECRET_NAME=\u003cexisting-secret-name\u003e\n+DEMO_SYNC_ROTATION_NEW_VALUE=\u003cnew-secret-value\u003e\n+DEMO_ASYNC_ROTATION_SECRET_NAME=\u003cdifferent-existing-secret-name\u003e\n+DEMO_ASYNC_ROTATION_NEW_VALUE=\u003cnew-secret-value\u003e\n+```\n+\n+Run:\n+\n+```text\n+mvn test\n+mvn exec:java\n+```\n+\n+The demo bulk-loads `database-connection-string`, `api-base-url`, and `feature-flags`, reads each\n+from cache, refreshes `api-base-url`, refreshes secrets within the seven-day warning window, prints\n+expiry warnings without exposing values, and rotates one sync and one async demo secret.\n+\n+## Design notes\n+\n+- Only `ResourceNotFoundException` is converted to a caller-provided default. Authentication,\n+  authorization, throttling, and service failures remain visible.\n+- A cached secret inside the warning window is automatically fetched again when read. The\n+  `refreshExpiring` methods support proactive sweeps as well.\n+- Rotation waits for the delete long-running operation, purges the soft-deleted secret, polls until\n+  the deleted name is no longer visible, and only then creates the replacement. Polling has a\n+  two-minute timeout.\n+- Azure SDK versions are managed by the current stable Azure SDK BOM (`1.3.8`).\n+\n+## References\n+\n+- [Azure Key Vault Secret client library for Java quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-java)\n+- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)\n+- [Azure SDK for Java Key Vault Secrets samples](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/keyvault/azure-security-keyvault-secrets/src/samples/java/com/azure/security/keyvault/secrets)\n*** End Patch\n"
```

### 🔧 17. powershell

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 18. web_fetch

**Input:**

```json
{"max_length":2000,"url":"https://repo1.maven.org/maven2/io/projectreactor/reactor-test/maven-metadata.xml"}
```

### 🔧 19. web_fetch

**Input:**

```json
{"max_length":20000,"url":"https://repo1.maven.org/maven2/com/azure/azure-sdk-bom/1.3.8/azure-sdk-bom-1.3.8.pom"}
```

### 🔧 20. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n-        \u003cdependency\u003e\n-            \u003cgroupId\u003eio.projectreactor\u003c/groupId\u003e\n-            \u003cartifactId\u003ereactor-test\u003c/artifactId\u003e\n-            \u003cscope\u003etest\u003c/scope\u003e\n-        \u003c/dependency\u003e\n*** Update File: src/test/java/com/example/keyvault/KeyVaultSecretProviderTest.java\n@@\n import java.time.OffsetDateTime;\n import org.junit.jupiter.api.Test;\n import reactor.core.publisher.Mono;\n-import reactor.test.StepVerifier;\n@@\n-        Mono\u003cSecretSnapshot\u003e result =\n-            new AsyncKeyVaultSecretProvider(client).getSecret(\"missing\", \"fallback\");\n-\n-        StepVerifier.create(result)\n-            .assertNext(secret -\u003e {\n-                assertEquals(\"fallback\", secret.value());\n-                assertTrue(secret.defaultValue());\n-            })\n-            .verifyComplete();\n+        SecretSnapshot result =\n+            new AsyncKeyVaultSecretProvider(client).getSecret(\"missing\", \"fallback\").block();\n+\n+        assertEquals(\"fallback\", result.value());\n+        assertTrue(result.defaultValue());\n     }\n }\n*** End Patch\n"
```

### 🔧 21. powershell

**Input:**

```json
{"command":"mvn test --no-transfer-progress","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\keyvault\AsyncKeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvault\AsyncSecretCache.java`
- `src\main\java\com\example\keyvault\AsyncSecretRotator.java`
- `src\main\java\com\example\keyvault\KeyVaultClientFactory.java`
- `src\main\java\com\example\keyvault\Main.java`
- `src\main\java\com\example\keyvault\SecretExpiry.java`
- `src\main\java\com\example\keyvault\SecretSnapshot.java`
- `src\main\java\com\example\keyvault\SyncKeyVaultSecretProvider.java`
- `src\main\java\com\example\keyvault\SyncSecretCache.java`
- `src\main\java\com\example\keyvault\SyncSecretRotator.java`
- `src\test\java\com\example\keyvault\KeyVaultSecretProviderTest.java`
- `src\test\java\com\example\keyvault\SyncSecretCacheTest.java`

## Copilot Response

I’m grounding the implementation in the SDK guidance, then I’ll build the project around testable client abstractions so the cache and rotation behavior can be verified without contacting Azure.

The rotation path needs one extra safeguard beyond waiting for deletion: a soft-deleted name cannot be recreated until it is purged. I’m implementing delete-completion polling, purge, and purge-propagation polling, with an explicit timeout and a clear failure when purge protection prevents delete-and-recreate rotation.

Created the complete Java 17 Maven project in the current directory. It includes:

- Sync/async Key Vault secret providers with defaults, version retrieval, and expiry metadata
- Expiry-aware sync/async in-memory caches with bulk loading and refresh
- Managed identity client factory using `AZURE_KEY_VAULT_URL`
- Delete, completion wait, purge, propagation wait, and recreate rotation helpers
- A sync-first, async-second demo without printing secret values
- Unit tests and documented environment variables in `README.md`
- Azure SDK BOM `1.3.8`, Key Vault Secrets `4.11.1`, and Identity `1.18.4`

Run with `mvn exec:java` after configuring the documented variables. Delete-and-recreate rotation requires purge permission and a vault without purge protection; otherwise Key Vault intentionally prevents reuse of the deleted name. References follow the official [Key Vault Java quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-java) and [Azure Identity guidance](https://learn.microsoft.com/java/api/overview/azure/identity-readme).

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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.1% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 11.0000** | **84.6%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-java-secret-config --config java-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
