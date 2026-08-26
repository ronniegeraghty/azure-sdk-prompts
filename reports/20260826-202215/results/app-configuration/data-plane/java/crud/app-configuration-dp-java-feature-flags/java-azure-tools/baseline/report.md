# Evaluation Report: app-configuration-dp-java-feature-flags

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 646.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-java-feature-flags` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 18/21 |
| Duration | 646.0s |
| Timestamp | 2026-08-26T14:14:23Z |
| Files Generated | 11 |
| Event Count | 15755 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 264.2s |
| Review | 380.7s |
| **Total** | **646.0s** |

## Configuration

- **name:** java-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=394592 out=11839 |
| Turn Count | 12 |

## Prompt Sent

```
Create a small Java 17 Maven project that implements a feature flag and configuration management system backed by Azure App Configuration.

The project needs:

- A **configuration service class** (both sync and async versions) that retrieves settings from App Configuration. It should support fetching a single setting by key, fetching a setting with a specific label (to distinguish between environments like "production" vs "staging"), and listing all settings that match a key prefix (returned as a map). It should also avoid re-downloading values that haven't changed since the last read — minimize unnecessary network traffic when polling for config changes.

- A **feature flag evaluator class** that reads feature flags from App Configuration. Feature flags in App Configuration use a special key prefix and store their state as a JSON payload. The evaluator should be able to check if a flag is enabled, and also support percentage-based rollout — if a flag is configured for a percentage rollout (e.g., 30% of users), the evaluator should deterministically decide whether a given user ID falls within the rollout percentage using a consistent hash, so the same user always gets the same result.

- A **configuration watcher class** that periodically polls for configuration changes. It should accept a list of "sentinel" keys to watch and a polling interval. When a sentinel key's value changes, the watcher should trigger a full refresh of all cached configuration. This is the recommended pattern for coordinating config updates in App Configuration.

- A **Main class** that demos both implementations: connecting to App Configuration (endpoint from environment variable, authenticated with managed identity), reading some config values with labels, evaluating feature flags for a few sample user IDs with percentage rollout, and starting the config watcher to detect a change. Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure Java SDK azure-data-appconfiguration latest ConfigurationClient conditional getConfigurationSetting ETag MatchConditions API Maven 2026"}
```

### 🔧 3. powershell

**Input:**

```json
{"command":"java -version; mvn -version","description":"Check Java and Maven versions","initial_wait":30,"mode":"sync"}
```

### 🔧 4. powershell

**Input:**

```json
{"command":"mvn --quiet dependency:get -Dartifact=com.azure:azure-data-appconfiguration:1.10.0 \u0026\u0026 javap -classpath \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-data-appconfiguration\\1.10.0\\azure-data-appconfiguration-1.10.0.jar\" com.azure.data.appconfiguration.ConfigurationClient com.azure.data.appconfiguration.ConfigurationAsyncClient com.azure.data.appconfiguration.models.SettingSelector com.azure.data.appconfiguration.models.ConfigurationSetting","description":"Inspect Azure App Configuration APIs","initial_wait":120,"mode":"sync"}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"mvn --quiet dependency:get '-Dartifact=com.azure:azure-data-appconfiguration:1.10.0'; if ($LASTEXITCODE -eq 0) { javap -classpath \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-data-appconfiguration\\1.10.0\\azure-data-appconfiguration-1.10.0.jar\" com.azure.data.appconfiguration.ConfigurationClient com.azure.data.appconfiguration.ConfigurationAsyncClient com.azure.data.appconfiguration.models.SettingSelector com.azure.data.appconfiguration.models.ConfigurationSetting }","description":"Download and inspect Azure SDK APIs","initial_wait":120,"mode":"sync"}
```

### 🔧 6. powershell

**Input:**

```json
{"command":"mvn --quiet dependency:get '-Dartifact=com.azure:azure-data-appconfiguration:1.10.0:jar:sources'; $src = \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-data-appconfiguration\\1.10.0\\azure-data-appconfiguration-1.10.0-sources.jar\"; \u0026 jar xf $src com/azure/data/appconfiguration/ConfigurationClient.java com/azure/data/appconfiguration/ConfigurationAsyncClient.java; Select-String -Path 'com\\azure\\data\\appconfiguration\\ConfigurationClient.java','com\\azure\\data\\appconfiguration\\ConfigurationAsyncClient.java' -Pattern 'getConfigurationSettingWithResponse' -Context 0,45; Remove-Item -Recurse -Force 'com'","description":"Inspect conditional read implementation","initial_wait":120,"mode":"sync"}
```

### 🔧 7. rg

**Input:**

```json
{"-C":4,"head_limit":120,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787753782001-copilot-tool-output-53140-94ee2dbb-255a-4fc5-ae52-81421a02b610.txt","pattern":"onlyIfChanged|statusCode|NOT_MODIFIED|getConfigurationSettingWithResponse"}
```

### 🔧 8. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787753782001-copilot-tool-output-53140-94ee2dbb-255a-4fc5-ae52-81421a02b610.txt","view_range":[80,145]}
```

### 🔧 9. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787753782001-copilot-tool-output-53140-94ee2dbb-255a-4fc5-ae52-81421a02b610.txt","view_range":[137,170]}
```

### 🔧 10. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-app-configuration-demo\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure App Configuration Demo\u003c/name\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cazure.appconfiguration.version\u003e1.10.0\u003c/azure.appconfiguration.version\u003e\n+        \u003cazure.identity.version\u003e1.18.2\u003c/azure.identity.version\u003e\n+        \u003cjackson.version\u003e2.20.0\u003c/jackson.version\u003e\n+        \u003cjunit.version\u003e5.13.4\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-data-appconfiguration\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.appconfiguration.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.fasterxml.jackson.core\u003c/groupId\u003e\n+            \u003cartifactId\u003ejackson-databind\u003c/artifactId\u003e\n+            \u003cversion\u003e${jackson.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.1\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.1\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.appconfig.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/appconfig/ConfigurationService.java\n+package com.example.appconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.core.http.rest.Response;\n+import com.azure.core.util.Context;\n+import com.azure.data.appconfiguration.ConfigurationClient;\n+import com.azure.data.appconfiguration.models.ConfigurationSetting;\n+import com.azure.data.appconfiguration.models.SettingSelector;\n+\n+import java.util.LinkedHashMap;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.Optional;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public class ConfigurationService {\n+    private final ConfigurationClient client;\n+    private final Map\u003cCacheKey, ConfigurationSetting\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public ConfigurationService(ConfigurationClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public Optional\u003cString\u003e getSetting(String key) {\n+        return getSetting(key, null);\n+    }\n+\n+    public Optional\u003cString\u003e getSetting(String key, String label) {\n+        Objects.requireNonNull(key, \"key\");\n+        CacheKey cacheKey = new CacheKey(key, label);\n+        ConfigurationSetting cached = cache.get(cacheKey);\n+        ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);\n+\n+        if (cached != null) {\n+            request.setETag(cached.getETag());\n+        }\n+\n+        try {\n+            Response\u003cConfigurationSetting\u003e response =\n+                client.getConfigurationSettingWithResponse(request, null, cached != null, Context.NONE);\n+            if (response.getStatusCode() == 304) {\n+                return Optional.ofNullable(cached.getValue());\n+            }\n+\n+            ConfigurationSetting current = response.getValue();\n+            cache.put(cacheKey, current);\n+            return Optional.ofNullable(current.getValue());\n+        } catch (ResourceNotFoundException exception) {\n+            cache.remove(cacheKey);\n+            return Optional.empty();\n+        }\n+    }\n+\n+    public Map\u003cString, String\u003e listSettings(String keyPrefix) {\n+        return listSettings(keyPrefix, null);\n+    }\n+\n+    public Map\u003cString, String\u003e listSettings(String keyPrefix, String label) {\n+        Objects.requireNonNull(keyPrefix, \"keyPrefix\");\n+        SettingSelector selector = new SettingSelector()\n+            .setKeyFilter(keyPrefix + \"*\")\n+            .setLabelFilter(label == null ? ConfigurationSetting.NO_LABEL : label);\n+\n+        Map\u003cString, String\u003e result = new LinkedHashMap\u003c\u003e();\n+        for (ConfigurationSetting setting : client.listConfigurationSettings(selector)) {\n+            result.put(setting.getKey(), setting.getValue());\n+            cache.put(new CacheKey(setting.getKey(), setting.getLabel()), setting);\n+        }\n+        return Map.copyOf(result);\n+    }\n+\n+    public void refreshAll() {\n+        cache.clear();\n+    }\n+\n+    private record CacheKey(String key, String label) {\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncConfigurationService.java\n+package com.example.appconfig;\n+\n+import com.azure.core.exception.ResourceNotFoundException;\n+import com.azure.data.appconfiguration.ConfigurationAsyncClient;\n+import com.azure.data.appconfiguration.models.ConfigurationSetting;\n+import com.azure.data.appconfiguration.models.SettingSelector;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.LinkedHashMap;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.Optional;\n+import java.util.concurrent.ConcurrentHashMap;\n+\n+public class AsyncConfigurationService {\n+    private final ConfigurationAsyncClient client;\n+    private final Map\u003cCacheKey, ConfigurationSetting\u003e cache = new ConcurrentHashMap\u003c\u003e();\n+\n+    public AsyncConfigurationService(ConfigurationAsyncClient client) {\n+        this.client = Objects.requireNonNull(client, \"client\");\n+    }\n+\n+    public Mono\u003cOptional\u003cString\u003e\u003e getSetting(String key) {\n+        return getSetting(key, null);\n+    }\n+\n+    public Mono\u003cOptional\u003cString\u003e\u003e getSetting(String key, String label) {\n+        Objects.requireNonNull(key, \"key\");\n+        return Mono.defer(() -\u003e {\n+            CacheKey cacheKey = new CacheKey(key, label);\n+            ConfigurationSetting cached = cache.get(cacheKey);\n+            ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);\n+            if (cached != null) {\n+                request.setETag(cached.getETag());\n+            }\n+\n+            return client.getConfigurationSettingWithResponse(request, null, cached != null)\n+                .map(response -\u003e {\n+                    if (response.getStatusCode() == 304) {\n+                        return Optional.ofNullable(cached.getValue());\n+                    }\n+\n+                    ConfigurationSetting current = response.getValue();\n+                    cache.put(cacheKey, current);\n+                    return Optional.ofNullable(current.getValue());\n+                })\n+                .onErrorResume(ResourceNotFoundException.class, exception -\u003e {\n+                    cache.remove(cacheKey);\n+                    return Mono.just(Optional.empty());\n+                });\n+        });\n+    }\n+\n+    public Mono\u003cMap\u003cString, String\u003e\u003e listSettings(String keyPrefix) {\n+        return listSettings(keyPrefix, null);\n+    }\n+\n+    public Mono\u003cMap\u003cString, String\u003e\u003e listSettings(String keyPrefix, String label) {\n+        Objects.requireNonNull(keyPrefix, \"keyPrefix\");\n+        SettingSelector selector = new SettingSelector()\n+            .setKeyFilter(keyPrefix + \"*\")\n+            .setLabelFilter(label == null ? ConfigurationSetting.NO_LABEL : label);\n+\n+        return client.listConfigurationSettings(selector)\n+            .collect(LinkedHashMap\u003cString, String\u003e::new, (result, setting) -\u003e {\n+                result.put(setting.getKey(), setting.getValue());\n+                cache.put(new CacheKey(setting.getKey(), setting.getLabel()), setting);\n+            })\n+            .map(Map::copyOf);\n+    }\n+\n+    public Mono\u003cVoid\u003e refreshAll() {\n+        return Mono.fromRunnable(cache::clear);\n+    }\n+\n+    private record CacheKey(String key, String label) {\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/FeatureFlagLogic.java\n+package com.example.appconfig;\n+\n+import com.fasterxml.jackson.core.JsonProcessingException;\n+import com.fasterxml.jackson.databind.JsonNode;\n+import com.fasterxml.jackson.databind.ObjectMapper;\n+\n+import java.math.BigInteger;\n+import java.nio.charset.StandardCharsets;\n+import java.security.MessageDigest;\n+import java.security.NoSuchAlgorithmException;\n+import java.util.Locale;\n+import java.util.OptionalDouble;\n+\n+final class FeatureFlagLogic {\n+    private static final ObjectMapper JSON = new ObjectMapper();\n+    private static final BigInteger TEN_THOUSAND = BigInteger.valueOf(10_000);\n+\n+    private FeatureFlagLogic() {\n+    }\n+\n+    static boolean evaluate(String flagName, String payload, String userId) {\n+        JsonNode flag;\n+        try {\n+            flag = JSON.readTree(payload);\n+        } catch (JsonProcessingException exception) {\n+            throw new IllegalArgumentException(\"Invalid JSON for feature flag \" + flagName, exception);\n+        }\n+\n+        if (!flag.path(\"enabled\").asBoolean(false)) {\n+            return false;\n+        }\n+\n+        OptionalDouble percentage = findPercentage(flag);\n+        if (percentage.isEmpty()) {\n+            return true;\n+        }\n+        if (userId == null || userId.isBlank()) {\n+            return false;\n+        }\n+\n+        double value = percentage.getAsDouble();\n+        if (value \u003c 0 || value \u003e 100) {\n+            throw new IllegalArgumentException(\n+                \"Percentage for feature flag \" + flagName + \" must be between 0 and 100\");\n+        }\n+        return rolloutBucket(flagName, userId) \u003c value * 100;\n+    }\n+\n+    private static OptionalDouble findPercentage(JsonNode flag) {\n+        JsonNode filters = flag.path(\"conditions\").path(\"client_filters\");\n+        if (!filters.isArray()) {\n+            return OptionalDouble.empty();\n+        }\n+\n+        for (JsonNode filter : filters) {\n+            String name = filter.path(\"name\").asText(\"\").toLowerCase(Locale.ROOT);\n+            if (name.endsWith(\"percentage\")) {\n+                JsonNode parameters = filter.path(\"parameters\");\n+                JsonNode value = parameters.has(\"Value\")\n+                    ? parameters.get(\"Value\")\n+                    : parameters.get(\"value\");\n+                if (value == null || (!value.isNumber() \u0026\u0026 !value.isTextual())) {\n+                    throw new IllegalArgumentException(\"Percentage filter is missing its Value parameter\");\n+                }\n+                try {\n+                    return OptionalDouble.of(value.isNumber()\n+                        ? value.asDouble()\n+                        : Double.parseDouble(value.asText()));\n+                } catch (NumberFormatException exception) {\n+                    throw new IllegalArgumentException(\"Percentage filter Value must be numeric\", exception);\n+                }\n+            }\n+        }\n+        return OptionalDouble.empty();\n+    }\n+\n+    private static int rolloutBucket(String flagName, String userId) {\n+        try {\n+            MessageDigest digest = MessageDigest.getInstance(\"SHA-256\");\n+            byte[] hash = digest.digest((flagName + \":\" + userId).getBytes(StandardCharsets.UTF_8));\n+            byte[] firstEightBytes = new byte[8];\n+            System.arraycopy(hash, 0, firstEightBytes, 0, firstEightBytes.length);\n+            return new BigInteger(1, firstEightBytes).mod(TEN_THOUSAND).intValue();\n+        } catch (NoSuchAlgorithmException exception) {\n+            throw new IllegalStateException(\"SHA-256 is unavailable\", exception);\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/FeatureFlagEvaluator.java\n+package com.example.appconfig;\n+\n+import java.util.Objects;\n+\n+public class FeatureFlagEvaluator {\n+    public static final String FEATURE_FLAG_PREFIX = \".appconfig.featureflag/\";\n+\n+    private final ConfigurationService configurationService;\n+    private final String label;\n+\n+    public FeatureFlagEvaluator(ConfigurationService configurationService) {\n+        this(configurationService, null);\n+    }\n+\n+    public FeatureFlagEvaluator(ConfigurationService configurationService, String label) {\n+        this.configurationService = Objects.requireNonNull(configurationService, \"configurationService\");\n+        this.label = label;\n+    }\n+\n+    public boolean isEnabled(String flagName) {\n+        return isEnabled(flagName, null);\n+    }\n+\n+    public boolean isEnabled(String flagName, String userId) {\n+        String key = flagName.startsWith(FEATURE_FLAG_PREFIX)\n+            ? flagName\n+            : FEATURE_FLAG_PREFIX + flagName;\n+        return configurationService.getSetting(key, label)\n+            .map(payload -\u003e FeatureFlagLogic.evaluate(flagName, payload, userId))\n+            .orElse(false);\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncFeatureFlagEvaluator.java\n+package com.example.appconfig;\n+\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public class AsyncFeatureFlagEvaluator {\n+    private final AsyncConfigurationService configurationService;\n+    private final String label;\n+\n+    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService) {\n+        this(configurationService, null);\n+    }\n+\n+    public AsyncFeatureFlagEvaluator(AsyncConfigurationService configurationService, String label) {\n+        this.configurationService = Objects.requireNonNull(configurationService, \"configurationService\");\n+        this.label = label;\n+    }\n+\n+    public Mono\u003cBoolean\u003e isEnabled(String flagName) {\n+        return isEnabled(flagName, null);\n+    }\n+\n+    public Mono\u003cBoolean\u003e isEnabled(String flagName, String userId) {\n+        String key = flagName.startsWith(FeatureFlagEvaluator.FEATURE_FLAG_PREFIX)\n+            ? flagName\n+            : FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + flagName;\n+        return configurationService.getSetting(key, label)\n+            .map(payload -\u003e payload\n+                .map(value -\u003e FeatureFlagLogic.evaluate(flagName, value, userId))\n+                .orElse(false));\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/ConfigurationWatcher.java\n+package com.example.appconfig;\n+\n+import java.time.Duration;\n+import java.util.ArrayList;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.Optional;\n+import java.util.concurrent.ConcurrentHashMap;\n+import java.util.concurrent.Executors;\n+import java.util.concurrent.ScheduledExecutorService;\n+import java.util.concurrent.TimeUnit;\n+import java.util.concurrent.atomic.AtomicBoolean;\n+import java.util.function.Consumer;\n+\n+public class ConfigurationWatcher implements AutoCloseable {\n+    private final ConfigurationService configurationService;\n+    private final List\u003cString\u003e sentinelKeys;\n+    private final Duration pollingInterval;\n+    private final Consumer\u003cList\u003cString\u003e\u003e refreshListener;\n+    private final Map\u003cString, Optional\u003cString\u003e\u003e sentinelValues = new ConcurrentHashMap\u003c\u003e();\n+    private final AtomicBoolean started = new AtomicBoolean();\n+    private final ScheduledExecutorService executor = Executors.newSingleThreadScheduledExecutor(runnable -\u003e {\n+        Thread thread = new Thread(runnable, \"app-configuration-watcher\");\n+        thread.setDaemon(true);\n+        return thread;\n+    });\n+\n+    public ConfigurationWatcher(\n+        ConfigurationService configurationService,\n+        List\u003cString\u003e sentinelKeys,\n+        Duration pollingInterval,\n+        Consumer\u003cList\u003cString\u003e\u003e refreshListener\n+    ) {\n+        this.configurationService = Objects.requireNonNull(configurationService, \"configurationService\");\n+        this.sentinelKeys = List.copyOf(sentinelKeys);\n+        if (this.sentinelKeys.isEmpty()) {\n+            throw new IllegalArgumentException(\"At least one sentinel key is required\");\n+        }\n+        this.pollingInterval = requirePositive(pollingInterval);\n+        this.refreshListener = Objects.requireNonNull(refreshListener, \"refreshListener\");\n+    }\n+\n+    public void start() {\n+        if (started.compareAndSet(false, true)) {\n+            executor.scheduleWithFixedDelay(\n+                this::pollSafely,\n+                0,\n+                pollingInterval.toMillis(),\n+                TimeUnit.MILLISECONDS);\n+        }\n+    }\n+\n+    private void pollSafely() {\n+        try {\n+            List\u003cString\u003e changedKeys = new ArrayList\u003c\u003e();\n+            for (String key : sentinelKeys) {\n+                Optional\u003cString\u003e current = configurationService.getSetting(key);\n+                Optional\u003cString\u003e previous = sentinelValues.put(key, current);\n+                if (previous != null \u0026\u0026 !previous.equals(current)) {\n+                    changedKeys.add(key);\n+                }\n+            }\n+\n+            if (!changedKeys.isEmpty()) {\n+                configurationService.refreshAll();\n+                refreshListener.accept(List.copyOf(changedKeys));\n+            }\n+        } catch (RuntimeException exception) {\n+            System.err.println(\"Configuration watcher poll failed: \" + exception.getMessage());\n+        }\n+    }\n+\n+    @Override\n+    public void close() {\n+        executor.shutdownNow();\n+    }\n+\n+    private static Duration requirePositive(Duration duration) {\n+        Objects.requireNonNull(duration, \"pollingInterval\");\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(\"pollingInterval must be positive\");\n+        }\n+        return duration;\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/AsyncConfigurationWatcher.java\n+package com.example.appconfig;\n+\n+import reactor.core.Disposable;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.List;\n+import java.util.Map;\n+import java.util.Objects;\n+import java.util.Optional;\n+import java.util.concurrent.ConcurrentHashMap;\n+import java.util.function.Consumer;\n+\n+public class AsyncConfigurationWatcher implements AutoCloseable {\n+    private final AsyncConfigurationService configurationService;\n+    private final List\u003cString\u003e sentinelKeys;\n+    private final Duration pollingInterval;\n+    private final Consumer\u003cList\u003cString\u003e\u003e refreshListener;\n+    private final Map\u003cString, Optional\u003cString\u003e\u003e sentinelValues = new ConcurrentHashMap\u003c\u003e();\n+    private Disposable subscription;\n+\n+    public AsyncConfigurationWatcher(\n+        AsyncConfigurationService configurationService,\n+        List\u003cString\u003e sentinelKeys,\n+        Duration pollingInterval,\n+        Consumer\u003cList\u003cString\u003e\u003e refreshListener\n+    ) {\n+        this.configurationService = Objects.requireNonNull(configurationService, \"configurationService\");\n+        this.sentinelKeys = List.copyOf(sentinelKeys);\n+        if (this.sentinelKeys.isEmpty()) {\n+            throw new IllegalArgumentException(\"At least one sentinel key is required\");\n+        }\n+        this.pollingInterval = requirePositive(pollingInterval);\n+        this.refreshListener = Objects.requireNonNull(refreshListener, \"refreshListener\");\n+    }\n+\n+    public synchronized void start() {\n+        if (subscription == null || subscription.isDisposed()) {\n+            subscription = Flux.interval(Duration.ZERO, pollingInterval)\n+                .concatMap(ignored -\u003e poll()\n+                    .onErrorResume(exception -\u003e {\n+                        System.err.println(\"Async configuration watcher poll failed: \" + exception.getMessage());\n+                        return Mono.empty();\n+                    }))\n+                .subscribe();\n+        }\n+    }\n+\n+    private Mono\u003cVoid\u003e poll() {\n+        return Flux.fromIterable(sentinelKeys)\n+            .concatMap(key -\u003e configurationService.getSetting(key)\n+                .map(current -\u003e new SentinelRead(key, current)))\n+            .filter(read -\u003e {\n+                Optional\u003cString\u003e previous = sentinelValues.put(read.key(), read.value());\n+                return previous != null \u0026\u0026 !previous.equals(read.value());\n+            })\n+            .map(SentinelRead::key)\n+            .collectList()\n+            .flatMap(changedKeys -\u003e {\n+                if (changedKeys.isEmpty()) {\n+                    return Mono.empty();\n+                }\n+                return configurationService.refreshAll()\n+                    .then(Mono.fromRunnable(() -\u003e refreshListener.accept(List.copyOf(changedKeys))));\n+            });\n+    }\n+\n+    @Override\n+    public synchronized void close() {\n+        if (subscription != null) {\n+            subscription.dispose();\n+            subscription = null;\n+        }\n+    }\n+\n+    private static Duration requirePositive(Duration duration) {\n+        Objects.requireNonNull(duration, \"pollingInterval\");\n+        if (duration.isZero() || duration.isNegative()) {\n+            throw new IllegalArgumentException(\"pollingInterval must be positive\");\n+        }\n+        return duration;\n+    }\n+\n+    private record SentinelRead(String key, Optional\u003cString\u003e value) {\n+    }\n+}\n*** Add File: src/main/java/com/example/appconfig/Main.java\n+package com.example.appconfig;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.data.appconfiguration.ConfigurationAsyncClient;\n+import com.azure.data.appconfiguration.ConfigurationClient;\n+import com.azure.data.appconfiguration.ConfigurationClientBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import reactor.core.publisher.Flux;\n+import reactor.core.publisher.Mono;\n+\n+import java.time.Duration;\n+import java.util.List;\n+\n+public final class Main {\n+    private static final String ENDPOINT_ENV = \"AZURE_APPCONFIG_ENDPOINT\";\n+    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(5);\n+    private static final List\u003cString\u003e SAMPLE_USERS = List.of(\"alice\", \"bob\", \"charlie\", \"diana\");\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) throws InterruptedException {\n+        String endpoint = requireEnvironmentVariable(ENDPOINT_ENV);\n+        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();\n+\n+        ConfigurationClient syncClient = new ConfigurationClientBuilder()\n+            .endpoint(endpoint)\n+            .credential(credential)\n+            .buildClient();\n+        runSyncDemo(new ConfigurationService(syncClient));\n+\n+        ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()\n+            .endpoint(endpoint)\n+            .credential(credential)\n+            .buildAsyncClient();\n+        runAsyncDemo(new AsyncConfigurationService(asyncClient));\n+    }\n+\n+    private static void runSyncDemo(ConfigurationService service) throws InterruptedException {\n+        System.out.println(\"=== Synchronous demo ===\");\n+        print(\"Production message\", service.getSetting(\"app:message\", \"production\").orElse(\"\u003cmissing\u003e\"));\n+        print(\"Staging message\", service.getSetting(\"app:message\", \"staging\").orElse(\"\u003cmissing\u003e\"));\n+        print(\"Production app settings\", service.listSettings(\"app:\", \"production\"));\n+\n+        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(service, \"production\");\n+        for (String userId : SAMPLE_USERS) {\n+            print(\"BetaFeature for \" + userId, flags.isEnabled(\"BetaFeature\", userId));\n+        }\n+\n+        try (ConfigurationWatcher watcher = new ConfigurationWatcher(\n+            service,\n+            List.of(\"app:sentinel\"),\n+            POLLING_INTERVAL,\n+            changed -\u003e System.out.println(\"Sync refresh triggered by \" + changed)\n+        )) {\n+            watcher.start();\n+            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());\n+        }\n+    }\n+\n+    private static void runAsyncDemo(AsyncConfigurationService service) {\n+        System.out.println(\"=== Asynchronous demo ===\");\n+        Mono.zip(\n+                service.getSetting(\"app:message\", \"production\"),\n+                service.getSetting(\"app:message\", \"staging\"),\n+                service.listSettings(\"app:\", \"production\"))\n+            .doOnNext(values -\u003e {\n+                print(\"Production message\", values.getT1().orElse(\"\u003cmissing\u003e\"));\n+                print(\"Staging message\", values.getT2().orElse(\"\u003cmissing\u003e\"));\n+                print(\"Production app settings\", values.getT3());\n+            })\n+            .block();\n+\n+        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(service, \"production\");\n+        Flux.fromIterable(SAMPLE_USERS)\n+            .concatMap(userId -\u003e flags.isEnabled(\"BetaFeature\", userId)\n+                .doOnNext(enabled -\u003e print(\"BetaFeature for \" + userId, enabled)))\n+            .then()\n+            .block();\n+\n+        try (AsyncConfigurationWatcher watcher = new AsyncConfigurationWatcher(\n+            service,\n+            List.of(\"app:sentinel\"),\n+            POLLING_INTERVAL,\n+            changed -\u003e System.out.println(\"Async refresh triggered by \" + changed)\n+        )) {\n+            watcher.start();\n+            Mono.delay(POLLING_INTERVAL.multipliedBy(2)).block();\n+        }\n+    }\n+\n+    private static String requireEnvironmentVariable(String name) {\n+        String value = System.getenv(name);\n+        if (value == null || value.isBlank()) {\n+            throw new IllegalStateException(\"Set \" + name + \" to your Azure App Configuration endpoint\");\n+        }\n+        return value;\n+    }\n+\n+    private static void print(String label, Object value) {\n+        System.out.printf(\"%-28s %s%n\", label + \":\", value);\n+    }\n+}\n*** Add File: src/test/java/com/example/appconfig/FeatureFlagLogicTest.java\n+package com.example.appconfig;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+import static org.junit.jupiter.api.Assertions.assertFalse;\n+import static org.junit.jupiter.api.Assertions.assertThrows;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class FeatureFlagLogicTest {\n+    @Test\n+    void evaluatesSimpleEnabledAndDisabledFlags() {\n+        assertTrue(FeatureFlagLogic.evaluate(\"simple\", \"{\\\"enabled\\\":true}\", null));\n+        assertFalse(FeatureFlagLogic.evaluate(\"simple\", \"{\\\"enabled\\\":false}\", \"alice\"));\n+    }\n+\n+    @Test\n+    void percentageRolloutIsDeterministic() {\n+        String payload = percentageFlag(30);\n+        boolean first = FeatureFlagLogic.evaluate(\"beta\", payload, \"alice\");\n+\n+        for (int attempt = 0; attempt \u003c 20; attempt++) {\n+            assertEquals(first, FeatureFlagLogic.evaluate(\"beta\", payload, \"alice\"));\n+        }\n+    }\n+\n+    @Test\n+    void percentageBoundaryValuesAreRespected() {\n+        assertFalse(FeatureFlagLogic.evaluate(\"beta\", percentageFlag(0), \"alice\"));\n+        assertTrue(FeatureFlagLogic.evaluate(\"beta\", percentageFlag(100), \"alice\"));\n+        assertFalse(FeatureFlagLogic.evaluate(\"beta\", percentageFlag(30), null));\n+    }\n+\n+    @Test\n+    void percentageDistributionTracksConfiguredValue() {\n+        int enabled = 0;\n+        for (int user = 0; user \u003c 10_000; user++) {\n+            if (FeatureFlagLogic.evaluate(\"beta\", percentageFlag(30), \"user-\" + user)) {\n+                enabled++;\n+            }\n+        }\n+\n+        assertTrue(enabled \u003e= 2_800 \u0026\u0026 enabled \u003c= 3_200, \"enabled users: \" + enabled);\n+    }\n+\n+    @Test\n+    void rejectsMalformedPayloadsAndPercentages() {\n+        assertThrows(IllegalArgumentException.class,\n+            () -\u003e FeatureFlagLogic.evaluate(\"beta\", \"not-json\", \"alice\"));\n+        assertThrows(IllegalArgumentException.class,\n+            () -\u003e FeatureFlagLogic.evaluate(\"beta\", percentageFlag(101), \"alice\"));\n+    }\n+\n+    private static String percentageFlag(int percentage) {\n+        return \"\"\"\n+            {\n+              \"enabled\": true,\n+              \"conditions\": {\n+                \"client_filters\": [\n+                  {\n+                    \"name\": \"Microsoft.Percentage\",\n+                    \"parameters\": {\"Value\": \"%d\"}\n+                  }\n+                ]\n+              }\n+            }\n+            \"\"\".formatted(percentage);\n+    }\n+}\n*** Add File: README.md\n+# Azure App Configuration demo\n+\n+Java 17 example with synchronous and Reactor-based asynchronous configuration access, ETag-aware\n+conditional reads, percentage feature flags, and sentinel-driven cache refresh.\n+\n+Expected sample data:\n+\n+| Key | Label | Example value |\n+| --- | --- | --- |\n+| `app:message` | `production` | `Hello from production` |\n+| `app:message` | `staging` | `Hello from staging` |\n+| `app:sentinel` | no label | `1` |\n+| `.appconfig.featureflag/BetaFeature` | `production` | JSON below |\n+\n+```json\n+{\n+  \"id\": \"BetaFeature\",\n+  \"enabled\": true,\n+  \"conditions\": {\n+    \"client_filters\": [\n+      {\n+        \"name\": \"Microsoft.Percentage\",\n+        \"parameters\": {\n+          \"Value\": \"30\"\n+        }\n+      }\n+    ]\n+  }\n+}\n+```\n+\n+Build and run:\n+\n+```powershell\n+mvn verify\n+$env:AZURE_APPCONFIG_ENDPOINT = \"https://\u003cstore-name\u003e.azconfig.io\"\n+mvn exec:java\n+```\n+\n+The runtime environment must provide a managed identity with App Configuration Data Reader access.\n+The demo does not create or modify Azure resources.\n*** End Patch\n"
```

### 🔧 11. powershell

**Input:**

```json
{"command":"mvn --quiet verify","description":"Build project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 12. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/appconfig/ConfigurationService.java\n@@\n import java.util.LinkedHashMap;\n+import java.util.List;\n import java.util.Map;\n@@\n     public void refreshAll() {\n-        cache.clear();\n+        List\u003cCacheKey\u003e cachedKeys = List.copyOf(cache.keySet());\n+        for (CacheKey cachedKey : cachedKeys) {\n+            getSetting(cachedKey.key(), cachedKey.label());\n+        }\n     }\n*** Update File: src/main/java/com/example/appconfig/AsyncConfigurationService.java\n@@\n import com.azure.data.appconfiguration.models.ConfigurationSetting;\n import com.azure.data.appconfiguration.models.SettingSelector;\n+import reactor.core.publisher.Flux;\n import reactor.core.publisher.Mono;\n \n import java.util.LinkedHashMap;\n+import java.util.List;\n@@\n     public Mono\u003cVoid\u003e refreshAll() {\n-        return Mono.fromRunnable(cache::clear);\n+        return Mono.defer(() -\u003e Flux.fromIterable(List.copyOf(cache.keySet()))\n+            .concatMap(cachedKey -\u003e getSetting(cachedKey.key(), cachedKey.label()))\n+            .then());\n     }\n*** End Patch\n"
```

### 🔧 13. powershell

**Input:**

```json
{"command":"mvn --quiet verify","description":"Run final Maven verification","initial_wait":180,"mode":"sync"}
```

### 🔧 14. rg

**Input:**

```json
{"-n":true,"output_mode":"content","paths":["src","pom.xml","README.md"],"pattern":"TODO|FIXME|connectionString|DefaultAzureCredential"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\appconfig\AsyncConfigurationService.java`
- `src\main\java\com\example\appconfig\AsyncConfigurationWatcher.java`
- `src\main\java\com\example\appconfig\AsyncFeatureFlagEvaluator.java`
- `src\main\java\com\example\appconfig\ConfigurationService.java`
- `src\main\java\com\example\appconfig\ConfigurationWatcher.java`
- `src\main\java\com\example\appconfig\FeatureFlagEvaluator.java`
- `src\main\java\com\example\appconfig\FeatureFlagLogic.java`
- `src\main\java\com\example\appconfig\Main.java`
- `src\test\java\com\example\appconfig\FeatureFlagLogicTest.java`

## Copilot Response

I’ll scaffold the project around the Azure SDK’s conditional-read APIs, then compile and test the cache, rollout, and watcher behavior without contacting Azure.

The SDK exposes ETag-aware `onlyIfChanged` reads for both clients, so polling can use conditional requests while sentinel changes remain the only trigger for a full cache refresh. I’m implementing that path directly and keeping rollout logic independently testable.

The initial build is clean. One semantic improvement remains: a sentinel change should eagerly revalidate every cached key, not merely invalidate the cache; I’m tightening both refresh implementations to do conditional ETag refreshes so unchanged payloads still are not downloaded.

Created a complete Java 17 Maven project with:

- Sync and Reactor-based async configuration services
- ETag conditional reads to avoid re-downloading unchanged settings
- Key, label, and prefix-based configuration retrieval
- JSON feature flag evaluation with deterministic SHA-256 percentage rollout
- Sync and async sentinel watchers that refresh all cached settings
- Managed identity authentication through `AZURE_APPCONFIG_ENDPOINT`
- A `Main` demo running sync first, then async
- Maven dependencies, unit tests, and setup documentation in `README.md`

Run with:

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
mvn exec:java
```

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated code fully satisfies the client resource-management criterion.

### Strengths

- AutoCloseable watcher resources are deterministically closed with try-with-resources.
- Azure SDK client instances used by the project do not expose a required close operation.

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
hyoka run --prompt-id app-configuration-dp-java-feature-flags --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
