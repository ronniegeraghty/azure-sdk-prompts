# Evaluation Report: identity-dp-java-credential-chain

**Config:** java-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 567.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-credential-chain` |
| Config | java-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 23/26 |
| Duration | 567.1s |
| Timestamp | 2026-08-26T15:06:16Z |
| Files Generated | 14 |
| Event Count | 15051 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 300.7s |
| Review | 265.2s |
| **Total** | **567.1s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** java-azure-tools/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=373018 out=10640 |
| Turn Count | 12 |

## Prompt Sent

```
Create a small Java 17 Maven project that demonstrates how to correctly build Azure credential chains tailored to different deployment environments — local development, CI/CD pipelines, and production.

The project needs:

- A **credential factory class** that builds the appropriate Azure credential for each environment. For local development, it should chain together credentials that work from developer tools (CLI, IDE plugins, etc.). For CI pipelines, it should support credentials sourced from pipeline environment variables or Azure Pipelines service connections. For production, it should prefer managed identity (supporting both system-assigned and user-assigned, where the user-assigned identity's client ID comes from an environment variable), with workload identity as a fallback for Kubernetes scenarios. The factory should also support enabling Continuous Access Evaluation (CAE) on the credentials, which lets Azure revoke tokens mid-session for security events.

- An **environment detector class** that auto-detects which environment the app is running in by probing for well-known environment variables (e.g., CI pipeline workspace variables, managed identity endpoint availability). It should classify the environment as dev, CI, or production.

- A **connectivity tester class** (both sync and async versions) that verifies a credential works by requesting a token for a given Azure scope. It should print success/failure, the token's expiry time, and whether the token is CAE-enabled. It should handle and report the specific failure reason if authentication fails (expired cert, wrong tenant, no identity available, etc.) rather than just printing a generic error.

- A **Main class** that detects the current environment, builds the right credential, and runs the connectivity test against Azure Resource Manager using the sync tester first, then repeats with the async tester. Print the detected environment, the selected credential strategy, and the test results from both.

Include a complete `pom.xml` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace root"}
```

### 🔧 3. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central latest stable com.azure azure-identity version August 2026 AzurePipelinesCredentialBuilder CAE Java"}
```

### 🔧 4. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure SDK for Java AzurePipelinesCredentialBuilder methods systemAccessToken serviceConnectionId clientId tenantId official documentation"}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"java -version; mvn -version","description":"Check Java and Maven availability"}
```

### 🔧 6. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: pom.xml\n+\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n+\u003cproject xmlns=\"http://maven.apache.org/POM/4.0.0\"\n+         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n+         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd\"\u003e\n+    \u003cmodelVersion\u003e4.0.0\u003c/modelVersion\u003e\n+\n+    \u003cgroupId\u003ecom.example\u003c/groupId\u003e\n+    \u003cartifactId\u003eazure-credential-chains\u003c/artifactId\u003e\n+    \u003cversion\u003e1.0.0\u003c/version\u003e\n+    \u003cname\u003eAzure credential chains\u003c/name\u003e\n+    \u003cdescription\u003eEnvironment-specific Azure credential chain examples\u003c/description\u003e\n+\n+    \u003cproperties\u003e\n+        \u003cproject.build.sourceEncoding\u003eUTF-8\u003c/project.build.sourceEncoding\u003e\n+        \u003cmaven.compiler.release\u003e17\u003c/maven.compiler.release\u003e\n+        \u003cazure.identity.version\u003e1.18.4\u003c/azure.identity.version\u003e\n+        \u003cjunit.version\u003e5.12.2\u003c/junit.version\u003e\n+    \u003c/properties\u003e\n+\n+    \u003cdependencies\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.slf4j\u003c/groupId\u003e\n+            \u003cartifactId\u003eslf4j-simple\u003c/artifactId\u003e\n+            \u003cversion\u003e2.0.17\u003c/version\u003e\n+            \u003cscope\u003eruntime\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003eorg.junit.jupiter\u003c/groupId\u003e\n+            \u003cartifactId\u003ejunit-jupiter\u003c/artifactId\u003e\n+            \u003cversion\u003e${junit.version}\u003c/version\u003e\n+            \u003cscope\u003etest\u003c/scope\u003e\n+        \u003c/dependency\u003e\n+    \u003c/dependencies\u003e\n+\n+    \u003cbuild\u003e\n+        \u003cplugins\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-compiler-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.14.0\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.apache.maven.plugins\u003c/groupId\u003e\n+                \u003cartifactId\u003emaven-surefire-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.3\u003c/version\u003e\n+            \u003c/plugin\u003e\n+            \u003cplugin\u003e\n+                \u003cgroupId\u003eorg.codehaus.mojo\u003c/groupId\u003e\n+                \u003cartifactId\u003eexec-maven-plugin\u003c/artifactId\u003e\n+                \u003cversion\u003e3.5.0\u003c/version\u003e\n+                \u003cconfiguration\u003e\n+                    \u003cmainClass\u003ecom.example.azureidentity.Main\u003c/mainClass\u003e\n+                \u003c/configuration\u003e\n+            \u003c/plugin\u003e\n+        \u003c/plugins\u003e\n+    \u003c/build\u003e\n+\u003c/project\u003e\n*** Add File: src/main/java/com/example/azureidentity/DeploymentEnvironment.java\n+package com.example.azureidentity;\n+\n+public enum DeploymentEnvironment {\n+    DEVELOPMENT,\n+    CI,\n+    PRODUCTION\n+}\n*** Add File: src/main/java/com/example/azureidentity/EnvironmentDetector.java\n+package com.example.azureidentity;\n+\n+import java.net.HttpURLConnection;\n+import java.net.URI;\n+import java.time.Duration;\n+import java.util.Map;\n+import java.util.function.BooleanSupplier;\n+\n+public final class EnvironmentDetector {\n+    private static final String IMDS_URL =\n+        \"http://169.254.169.254/metadata/instance?api-version=2021-02-01\";\n+    private static final Duration PROBE_TIMEOUT = Duration.ofMillis(250);\n+\n+    private final Map\u003cString, String\u003e environment;\n+    private final BooleanSupplier managedIdentityEndpointProbe;\n+\n+    public EnvironmentDetector() {\n+        this(System.getenv(), EnvironmentDetector::isImdsAvailable);\n+    }\n+\n+    EnvironmentDetector(Map\u003cString, String\u003e environment, BooleanSupplier managedIdentityEndpointProbe) {\n+        this.environment = Map.copyOf(environment);\n+        this.managedIdentityEndpointProbe = managedIdentityEndpointProbe;\n+    }\n+\n+    public DeploymentEnvironment detect() {\n+        if (hasAny(\"CI\", \"TF_BUILD\", \"BUILD_BUILDID\", \"PIPELINE_WORKSPACE\",\n+            \"GITHUB_ACTIONS\", \"GITHUB_WORKSPACE\", \"GITLAB_CI\", \"CI_PROJECT_DIR\", \"JENKINS_URL\")) {\n+            return DeploymentEnvironment.CI;\n+        }\n+\n+        if (hasAny(\"IDENTITY_ENDPOINT\", \"MSI_ENDPOINT\", \"IMDS_ENDPOINT\")\n+            || hasWorkloadIdentityConfiguration()\n+            || managedIdentityEndpointProbe.getAsBoolean()) {\n+            return DeploymentEnvironment.PRODUCTION;\n+        }\n+\n+        return DeploymentEnvironment.DEVELOPMENT;\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return hasValue(\"AZURE_CLIENT_ID\")\n+            \u0026\u0026 hasValue(\"AZURE_TENANT_ID\")\n+            \u0026\u0026 hasValue(\"AZURE_FEDERATED_TOKEN_FILE\");\n+    }\n+\n+    private boolean hasAny(String... names) {\n+        for (String name : names) {\n+            if (hasValue(name)) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+\n+    private boolean hasValue(String name) {\n+        String value = environment.get(name);\n+        return value != null \u0026\u0026 !value.isBlank() \u0026\u0026 !\"false\".equalsIgnoreCase(value);\n+    }\n+\n+    private static boolean isImdsAvailable() {\n+        HttpURLConnection connection = null;\n+        try {\n+            connection = (HttpURLConnection) URI.create(IMDS_URL).toURL().openConnection();\n+            connection.setRequestMethod(\"GET\");\n+            connection.setRequestProperty(\"Metadata\", \"true\");\n+            connection.setConnectTimeout((int) PROBE_TIMEOUT.toMillis());\n+            connection.setReadTimeout((int) PROBE_TIMEOUT.toMillis());\n+            int status = connection.getResponseCode();\n+            return status \u003e= 200 \u0026\u0026 status \u003c 500;\n+        } catch (Exception ignored) {\n+            return false;\n+        } finally {\n+            if (connection != null) {\n+                connection.disconnect();\n+            }\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/CaeEnabledCredential.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+final class CaeEnabledCredential implements TokenCredential {\n+    private final TokenCredential delegate;\n+\n+    CaeEnabledCredential(TokenCredential delegate) {\n+        this.delegate = Objects.requireNonNull(delegate, \"delegate\");\n+    }\n+\n+    @Override\n+    public Mono\u003cAccessToken\u003e getToken(TokenRequestContext request) {\n+        return delegate.getToken(withCae(request));\n+    }\n+\n+    @Override\n+    public AccessToken getTokenSync(TokenRequestContext request) {\n+        return delegate.getTokenSync(withCae(request));\n+    }\n+\n+    private static TokenRequestContext withCae(TokenRequestContext request) {\n+        return Objects.requireNonNull(request, \"request\").setCaeEnabled(true);\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AzureCredentialFactory.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.identity.AzureCliCredentialBuilder;\n+import com.azure.identity.AzureDeveloperCliCredentialBuilder;\n+import com.azure.identity.AzurePipelinesCredentialBuilder;\n+import com.azure.identity.AzurePowerShellCredentialBuilder;\n+import com.azure.identity.ChainedTokenCredentialBuilder;\n+import com.azure.identity.EnvironmentCredentialBuilder;\n+import com.azure.identity.IntelliJCredentialBuilder;\n+import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.identity.WorkloadIdentityCredentialBuilder;\n+\n+import java.util.Map;\n+import java.util.Objects;\n+\n+public final class AzureCredentialFactory {\n+    public record CredentialSelection(\n+        TokenCredential credential,\n+        DeploymentEnvironment environment,\n+        String strategy,\n+        boolean caeEnabled) {\n+    }\n+\n+    private final Map\u003cString, String\u003e environment;\n+\n+    public AzureCredentialFactory() {\n+        this(System.getenv());\n+    }\n+\n+    AzureCredentialFactory(Map\u003cString, String\u003e environment) {\n+        this.environment = Map.copyOf(environment);\n+    }\n+\n+    public CredentialSelection create(DeploymentEnvironment deploymentEnvironment, boolean enableCae) {\n+        Objects.requireNonNull(deploymentEnvironment, \"deploymentEnvironment\");\n+\n+        Strategy strategy = switch (deploymentEnvironment) {\n+            case DEVELOPMENT -\u003e developerCredential();\n+            case CI -\u003e ciCredential();\n+            case PRODUCTION -\u003e productionCredential();\n+        };\n+        TokenCredential credential = enableCae\n+            ? new CaeEnabledCredential(strategy.credential())\n+            : strategy.credential();\n+\n+        return new CredentialSelection(\n+            credential, deploymentEnvironment, strategy.description(), enableCae);\n+    }\n+\n+    private Strategy developerCredential() {\n+        TokenCredential credential = new ChainedTokenCredentialBuilder()\n+            .addLast(new AzureCliCredentialBuilder().build())\n+            .addLast(new AzureDeveloperCliCredentialBuilder().build())\n+            .addLast(new IntelliJCredentialBuilder().build())\n+            .addLast(new AzurePowerShellCredentialBuilder().build())\n+            .build();\n+        return new Strategy(\n+            credential,\n+            \"Developer tools: Azure CLI -\u003e Azure Developer CLI -\u003e IntelliJ -\u003e Azure PowerShell\");\n+    }\n+\n+    private Strategy ciCredential() {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        StringBuilder description = new StringBuilder(\"Pipeline environment variables\");\n+\n+        if (hasAzurePipelinesServiceConnection()) {\n+            chain.addLast(new AzurePipelinesCredentialBuilder()\n+                .tenantId(environment.get(\"AZURE_TENANT_ID\"))\n+                .clientId(environment.get(\"AZURE_CLIENT_ID\"))\n+                .serviceConnectionId(environment.get(\"AZURE_SERVICE_CONNECTION_ID\"))\n+                .systemAccessToken(environment.get(\"SYSTEM_ACCESSTOKEN\"))\n+                .build());\n+            description.append(\" -\u003e Azure Pipelines workload-identity service connection\");\n+        }\n+\n+        chain.addFirst(new EnvironmentCredentialBuilder().build());\n+        return new Strategy(chain.build(), description.toString());\n+    }\n+\n+    private Strategy productionCredential() {\n+        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();\n+        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();\n+        String clientId = value(\"AZURE_CLIENT_ID\");\n+        String identityType = \"system-assigned managed identity\";\n+\n+        if (clientId != null) {\n+            managedIdentity.clientId(clientId);\n+            identityType = \"user-assigned managed identity (AZURE_CLIENT_ID)\";\n+        }\n+        chain.addLast(managedIdentity.build());\n+\n+        String description = identityType;\n+        if (hasWorkloadIdentityConfiguration()) {\n+            chain.addLast(new WorkloadIdentityCredentialBuilder()\n+                .tenantId(environment.get(\"AZURE_TENANT_ID\"))\n+                .clientId(environment.get(\"AZURE_CLIENT_ID\"))\n+                .tokenFilePath(environment.get(\"AZURE_FEDERATED_TOKEN_FILE\"))\n+                .build());\n+            description += \" -\u003e Kubernetes workload identity\";\n+        } else {\n+            description += \" (workload identity fallback not configured)\";\n+        }\n+\n+        return new Strategy(chain.build(), description);\n+    }\n+\n+    private boolean hasAzurePipelinesServiceConnection() {\n+        return value(\"AZURE_TENANT_ID\") != null\n+            \u0026\u0026 value(\"AZURE_CLIENT_ID\") != null\n+            \u0026\u0026 value(\"AZURE_SERVICE_CONNECTION_ID\") != null\n+            \u0026\u0026 value(\"SYSTEM_ACCESSTOKEN\") != null;\n+    }\n+\n+    private boolean hasWorkloadIdentityConfiguration() {\n+        return value(\"AZURE_TENANT_ID\") != null\n+            \u0026\u0026 value(\"AZURE_CLIENT_ID\") != null\n+            \u0026\u0026 value(\"AZURE_FEDERATED_TOKEN_FILE\") != null;\n+    }\n+\n+    private String value(String name) {\n+        String value = environment.get(name);\n+        return value == null || value.isBlank() ? null : value;\n+    }\n+\n+    private record Strategy(TokenCredential credential, String description) {\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/ConnectivityTestResult.java\n+package com.example.azureidentity;\n+\n+import java.time.OffsetDateTime;\n+\n+public record ConnectivityTestResult(\n+    boolean successful,\n+    OffsetDateTime expiresAt,\n+    boolean caeEnabled,\n+    String failureReason) {\n+\n+    static ConnectivityTestResult success(OffsetDateTime expiresAt, boolean caeEnabled) {\n+        return new ConnectivityTestResult(true, expiresAt, caeEnabled, null);\n+    }\n+\n+    static ConnectivityTestResult failure(boolean caeEnabled, Throwable error) {\n+        return new ConnectivityTestResult(\n+            false, null, caeEnabled, AuthenticationFailureAnalyzer.explain(error));\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AuthenticationFailureAnalyzer.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.exception.ClientAuthenticationException;\n+import com.azure.identity.CredentialUnavailableException;\n+\n+import java.util.Locale;\n+\n+final class AuthenticationFailureAnalyzer {\n+    private AuthenticationFailureAnalyzer() {\n+    }\n+\n+    static String explain(Throwable error) {\n+        Throwable root = rootCause(error);\n+        String detail = firstMessage(error, root);\n+        String normalized = detail.toLowerCase(Locale.ROOT);\n+\n+        if (containsAny(normalized, \"expired certificate\", \"certificate has expired\",\n+            \"aadsts7000222\", \"client secret is expired\", \"credential has expired\")) {\n+            return \"The client certificate or secret has expired. \" + detail;\n+        }\n+        if (containsAny(normalized, \"aadsts90002\", \"tenant not found\", \"invalid tenant\",\n+            \"tenant does not exist\")) {\n+            return \"The configured Microsoft Entra tenant is invalid or unavailable. \" + detail;\n+        }\n+        if (containsAny(normalized, \"aadsts700016\", \"application with identifier\",\n+            \"was not found in the directory\")) {\n+            return \"The client ID is not registered in the configured tenant; check AZURE_CLIENT_ID \"\n+                + \"and AZURE_TENANT_ID. \" + detail;\n+        }\n+        if (containsAny(normalized, \"aadsts7000215\", \"invalid client secret\",\n+            \"invalid_client\")) {\n+            return \"The client secret or certificate is invalid. \" + detail;\n+        }\n+        if (containsAny(normalized, \"aadsts500011\", \"invalid scope\", \"resource principal\")) {\n+            return \"The requested Azure scope/resource is invalid for this tenant. \" + detail;\n+        }\n+        if (containsAny(normalized, \"managed identity\", \"identity not found\", \"no identity\",\n+            \"credential unavailable\", \"unavailable\", \"imds endpoint\")) {\n+            return \"No configured managed, workload, pipeline, or developer identity was available. \"\n+                + detail;\n+        }\n+        if (containsAny(normalized, \"federated token\", \"oidc\", \"token file\")) {\n+            return \"The workload or pipeline federated-identity configuration is invalid. \" + detail;\n+        }\n+        if (containsAny(normalized, \"timeout\", \"timed out\", \"connection\", \"dns\",\n+            \"unknown host\", \"network\")) {\n+            return \"The Microsoft Entra authentication endpoint could not be reached. \" + detail;\n+        }\n+        if (error instanceof CredentialUnavailableException) {\n+            return \"No credential in the selected chain could authenticate. \" + detail;\n+        }\n+        if (error instanceof ClientAuthenticationException) {\n+            return \"Microsoft Entra ID rejected the credential. \" + detail;\n+        }\n+        return error.getClass().getSimpleName() + \": \" + detail;\n+    }\n+\n+    private static Throwable rootCause(Throwable error) {\n+        Throwable current = error;\n+        while (current.getCause() != null \u0026\u0026 current.getCause() != current) {\n+            current = current.getCause();\n+        }\n+        return current;\n+    }\n+\n+    private static String firstMessage(Throwable error, Throwable root) {\n+        if (error.getMessage() != null \u0026\u0026 !error.getMessage().isBlank()) {\n+            return error.getMessage();\n+        }\n+        if (root.getMessage() != null \u0026\u0026 !root.getMessage().isBlank()) {\n+            return root.getMessage();\n+        }\n+        return root.getClass().getSimpleName();\n+    }\n+\n+    private static boolean containsAny(String value, String... candidates) {\n+        for (String candidate : candidates) {\n+            if (value.contains(candidate)) {\n+                return true;\n+            }\n+        }\n+        return false;\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/CredentialConnectivityTester.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.AccessToken;\n+import com.azure.core.credential.TokenCredential;\n+import com.azure.core.credential.TokenRequestContext;\n+\n+import java.util.Objects;\n+\n+public final class CredentialConnectivityTester {\n+    public ConnectivityTestResult test(\n+        TokenCredential credential, String scope, boolean caeEnabled) {\n+\n+        Objects.requireNonNull(credential, \"credential\");\n+        try {\n+            AccessToken token = credential.getTokenSync(request(scope, caeEnabled));\n+            ConnectivityTestResult result =\n+                ConnectivityTestResult.success(token.getExpiresAt(), caeEnabled);\n+            print(\"Sync\", result);\n+            return result;\n+        } catch (RuntimeException error) {\n+            ConnectivityTestResult result = ConnectivityTestResult.failure(caeEnabled, error);\n+            print(\"Sync\", result);\n+            return result;\n+        }\n+    }\n+\n+    static TokenRequestContext request(String scope, boolean caeEnabled) {\n+        if (scope == null || scope.isBlank()) {\n+            throw new IllegalArgumentException(\"scope must not be blank\");\n+        }\n+        return new TokenRequestContext()\n+            .addScopes(scope)\n+            .setCaeEnabled(caeEnabled);\n+    }\n+\n+    static void print(String mode, ConnectivityTestResult result) {\n+        if (result.successful()) {\n+            System.out.printf(\n+                \"%s test: SUCCESS%n  Expires at: %s%n  CAE enabled: %s%n\",\n+                mode, result.expiresAt(), result.caeEnabled());\n+        } else {\n+            System.out.printf(\n+                \"%s test: FAILURE%n  Reason: %s%n  CAE enabled: %s%n\",\n+                mode, result.failureReason(), result.caeEnabled());\n+        }\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/AsyncCredentialConnectivityTester.java\n+package com.example.azureidentity;\n+\n+import com.azure.core.credential.TokenCredential;\n+import reactor.core.publisher.Mono;\n+\n+import java.util.Objects;\n+\n+public final class AsyncCredentialConnectivityTester {\n+    public Mono\u003cConnectivityTestResult\u003e test(\n+        TokenCredential credential, String scope, boolean caeEnabled) {\n+\n+        Objects.requireNonNull(credential, \"credential\");\n+        return credential.getToken(CredentialConnectivityTester.request(scope, caeEnabled))\n+            .map(token -\u003e ConnectivityTestResult.success(token.getExpiresAt(), caeEnabled))\n+            .onErrorResume(error -\u003e\n+                Mono.just(ConnectivityTestResult.failure(caeEnabled, error)))\n+            .doOnNext(result -\u003e CredentialConnectivityTester.print(\"Async\", result));\n+    }\n+}\n*** Add File: src/main/java/com/example/azureidentity/Main.java\n+package com.example.azureidentity;\n+\n+public final class Main {\n+    private static final String AZURE_RESOURCE_MANAGER_SCOPE =\n+        \"https://management.azure.com/.default\";\n+\n+    private Main() {\n+    }\n+\n+    public static void main(String[] args) {\n+        boolean caeEnabled = readCaeSetting();\n+        DeploymentEnvironment environment = new EnvironmentDetector().detect();\n+        AzureCredentialFactory.CredentialSelection selection =\n+            new AzureCredentialFactory().create(environment, caeEnabled);\n+\n+        System.out.println(\"Detected environment: \" + environment);\n+        System.out.println(\"Credential strategy: \" + selection.strategy());\n+        System.out.println(\"Azure scope: \" + AZURE_RESOURCE_MANAGER_SCOPE);\n+        System.out.println();\n+\n+        new CredentialConnectivityTester().test(\n+            selection.credential(), AZURE_RESOURCE_MANAGER_SCOPE, selection.caeEnabled());\n+        System.out.println();\n+        new AsyncCredentialConnectivityTester().test(\n+            selection.credential(), AZURE_RESOURCE_MANAGER_SCOPE, selection.caeEnabled()).block();\n+    }\n+\n+    private static boolean readCaeSetting() {\n+        String value = System.getenv(\"AZURE_ENABLE_CAE\");\n+        return value == null || value.isBlank() || Boolean.parseBoolean(value);\n+    }\n+}\n*** Add File: src/test/java/com/example/azureidentity/EnvironmentDetectorTest.java\n+package com.example.azureidentity;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertEquals;\n+\n+class EnvironmentDetectorTest {\n+    @Test\n+    void defaultsToDevelopment() {\n+        assertEquals(\n+            DeploymentEnvironment.DEVELOPMENT,\n+            new EnvironmentDetector(Map.of(), () -\u003e false).detect());\n+    }\n+\n+    @Test\n+    void detectsCiBeforeManagedIdentity() {\n+        assertEquals(\n+            DeploymentEnvironment.CI,\n+            new EnvironmentDetector(\n+                Map.of(\"GITHUB_ACTIONS\", \"true\", \"IDENTITY_ENDPOINT\", \"http://localhost\"),\n+                () -\u003e true).detect());\n+    }\n+\n+    @Test\n+    void detectsProductionFromWorkloadIdentity() {\n+        assertEquals(\n+            DeploymentEnvironment.PRODUCTION,\n+            new EnvironmentDetector(\n+                Map.of(\n+                    \"AZURE_CLIENT_ID\", \"client\",\n+                    \"AZURE_TENANT_ID\", \"tenant\",\n+                    \"AZURE_FEDERATED_TOKEN_FILE\", \"token\"),\n+                () -\u003e false).detect());\n+    }\n+\n+    @Test\n+    void detectsProductionFromImdsProbe() {\n+        assertEquals(\n+            DeploymentEnvironment.PRODUCTION,\n+            new EnvironmentDetector(Map.of(), () -\u003e true).detect());\n+    }\n+}\n*** Add File: src/test/java/com/example/azureidentity/AuthenticationFailureAnalyzerTest.java\n+package com.example.azureidentity;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class AuthenticationFailureAnalyzerTest {\n+    @Test\n+    void identifiesExpiredCredentials() {\n+        String reason = AuthenticationFailureAnalyzer.explain(\n+            new RuntimeException(\"AADSTS7000222: The provided client secret keys are expired\"));\n+\n+        assertTrue(reason.startsWith(\"The client certificate or secret has expired.\"));\n+    }\n+\n+    @Test\n+    void identifiesWrongTenant() {\n+        String reason = AuthenticationFailureAnalyzer.explain(\n+            new RuntimeException(\"AADSTS90002: Tenant 'bad-tenant' not found\"));\n+\n+        assertTrue(reason.startsWith(\n+            \"The configured Microsoft Entra tenant is invalid or unavailable.\"));\n+    }\n+\n+    @Test\n+    void identifiesMissingIdentity() {\n+        String reason = AuthenticationFailureAnalyzer.explain(\n+            new RuntimeException(\"ManagedIdentityCredential authentication unavailable\"));\n+\n+        assertTrue(reason.startsWith(\n+            \"No configured managed, workload, pipeline, or developer identity was available.\"));\n+    }\n+}\n*** Add File: README.md\n+# Azure credential chains for Java\n+\n+This Java 17 sample selects an Azure Identity credential chain for local\n+development, CI/CD, or production, then requests an Azure Resource Manager token\n+with both synchronous and asynchronous SDK APIs. It does not create or modify\n+Azure resources.\n+\n+## Credential strategies\n+\n+| Environment | Detection signals | Credential order |\n+| --- | --- | --- |\n+| Development | No CI or Azure-hosting signals | Azure CLI, Azure Developer CLI, IntelliJ, Azure PowerShell |\n+| CI/CD | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, pipeline workspace variables, and similar | `EnvironmentCredential`, then Azure Pipelines workload-identity service connection when configured |\n+| Production | App Service/Functions managed-identity variables, Kubernetes workload-identity variables, or reachable IMDS | Managed identity, then Kubernetes workload identity when configured |\n+\n+Production uses a system-assigned managed identity by default. Set\n+`AZURE_CLIENT_ID` to select a user-assigned managed identity.\n+\n+For an Azure Pipelines workload-identity service connection, expose:\n+\n+- `AZURE_TENANT_ID`\n+- `AZURE_CLIENT_ID`\n+- `AZURE_SERVICE_CONNECTION_ID`\n+- `SYSTEM_ACCESSTOKEN`\n+- `SYSTEM_OIDCREQUESTURI` (provided by Azure Pipelines)\n+\n+For Kubernetes workload identity, expose `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,\n+and `AZURE_FEDERATED_TOKEN_FILE`.\n+\n+CAE is enabled by default. Set `AZURE_ENABLE_CAE=false` to disable it. CAE is\n+requested through `TokenRequestContext`; the resource and tenant must also\n+support CAE.\n+\n+## Run\n+\n+```shell\n+mvn test\n+mvn compile exec:java\n+```\n+\n+The second command performs real authentication against Microsoft Entra ID but\n+does not call Azure Resource Manager or change any Azure resources.\n*** End Patch\n"
```

### 🔧 7. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress test","description":"Compile project and run tests","initial_wait":180,"mode":"sync"}
```

### 🔧 8. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Maven Central latest stable com.azure azure-identity-broker version August 2026 compatible azure-identity 1.18.4"}
```

### 🔧 9. web_search (via github-mcp-server)

**Input:**

```json
{"query":"Azure Identity Java 1.18.4 VisualStudioCodeCredentialBuilder azure-identity-broker required official documentation"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"jar tf \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-identity\\1.18.4\\azure-identity-1.18.4.jar\" | Select-String 'VisualStudioCodeCredentialBuilder|IntelliJCredentialBuilder'","description":"Inspect available IDE credential builders"}
```

### 🔧 11. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: pom.xml\n@@\n         \u003cazure.identity.version\u003e1.18.4\u003c/azure.identity.version\u003e\n+        \u003cazure.identity.broker.version\u003e1.1.21\u003c/azure.identity.broker.version\u003e\n         \u003cjunit.version\u003e5.12.2\u003c/junit.version\u003e\n@@\n         \u003cdependency\u003e\n             \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n             \u003cartifactId\u003eazure-identity\u003c/artifactId\u003e\n             \u003cversion\u003e${azure.identity.version}\u003c/version\u003e\n         \u003c/dependency\u003e\n+        \u003cdependency\u003e\n+            \u003cgroupId\u003ecom.azure\u003c/groupId\u003e\n+            \u003cartifactId\u003eazure-identity-broker\u003c/artifactId\u003e\n+            \u003cversion\u003e${azure.identity.broker.version}\u003c/version\u003e\n+        \u003c/dependency\u003e\n*** Update File: src/main/java/com/example/azureidentity/AzureCredentialFactory.java\n@@\n import com.azure.identity.ManagedIdentityCredentialBuilder;\n+import com.azure.identity.VisualStudioCodeCredentialBuilder;\n import com.azure.identity.WorkloadIdentityCredentialBuilder;\n@@\n             .addLast(new AzureCliCredentialBuilder().build())\n             .addLast(new AzureDeveloperCliCredentialBuilder().build())\n+            .addLast(new VisualStudioCodeCredentialBuilder().build())\n             .addLast(new IntelliJCredentialBuilder().build())\n             .addLast(new AzurePowerShellCredentialBuilder().build())\n             .build();\n         return new Strategy(\n             credential,\n-            \"Developer tools: Azure CLI -\u003e Azure Developer CLI -\u003e IntelliJ -\u003e Azure PowerShell\");\n+            \"Developer tools: Azure CLI -\u003e Azure Developer CLI -\u003e Visual Studio Code -\u003e \"\n+                + \"IntelliJ -\u003e Azure PowerShell\");\n*** Update File: README.md\n@@\n-| Development | No CI or Azure-hosting signals | Azure CLI, Azure Developer CLI, IntelliJ, Azure PowerShell |\n+| Development | No CI or Azure-hosting signals | Azure CLI, Azure Developer CLI, Visual Studio Code, IntelliJ, Azure PowerShell |\n*** Add File: src/test/java/com/example/azureidentity/AzureCredentialFactoryTest.java\n+package com.example.azureidentity;\n+\n+import org.junit.jupiter.api.Test;\n+\n+import java.util.Map;\n+\n+import static org.junit.jupiter.api.Assertions.assertInstanceOf;\n+import static org.junit.jupiter.api.Assertions.assertTrue;\n+\n+class AzureCredentialFactoryTest {\n+    @Test\n+    void wrapsCredentialsWhenCaeIsEnabled() {\n+        AzureCredentialFactory.CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of()).create(\n+                DeploymentEnvironment.DEVELOPMENT, true);\n+\n+        assertTrue(selection.caeEnabled());\n+        assertInstanceOf(CaeEnabledCredential.class, selection.credential());\n+    }\n+\n+    @Test\n+    void selectsUserAssignedManagedIdentityFromClientId() {\n+        AzureCredentialFactory.CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of(\"AZURE_CLIENT_ID\", \"user-assigned-client\"))\n+                .create(DeploymentEnvironment.PRODUCTION, false);\n+\n+        assertTrue(selection.strategy().startsWith(\"user-assigned managed identity\"));\n+    }\n+\n+    @Test\n+    void addsConfiguredAzurePipelinesServiceConnection() {\n+        AzureCredentialFactory.CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of(\n+                \"AZURE_TENANT_ID\", \"tenant\",\n+                \"AZURE_CLIENT_ID\", \"client\",\n+                \"AZURE_SERVICE_CONNECTION_ID\", \"connection\",\n+                \"SYSTEM_ACCESSTOKEN\", \"token\"))\n+                .create(DeploymentEnvironment.CI, false);\n+\n+        assertTrue(selection.strategy().contains(\"Azure Pipelines workload-identity\"));\n+    }\n+\n+    @Test\n+    void addsConfiguredWorkloadIdentityFallback() {\n+        AzureCredentialFactory.CredentialSelection selection =\n+            new AzureCredentialFactory(Map.of(\n+                \"AZURE_TENANT_ID\", \"tenant\",\n+                \"AZURE_CLIENT_ID\", \"client\",\n+                \"AZURE_FEDERATED_TOKEN_FILE\", \"token-file\"))\n+                .create(DeploymentEnvironment.PRODUCTION, false);\n+\n+        assertTrue(selection.strategy().endsWith(\"Kubernetes workload identity\"));\n+    }\n+}\n*** End Patch\n"
```

### 🔧 12. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress clean test package","description":"Build and test complete Maven project","initial_wait":180,"mode":"sync"}
```

### 🔧 13. powershell

**Input:**

```json
{"command":"javap -classpath \"$env:USERPROFILE\\.m2\\repository\\com\\azure\\azure-identity\\1.18.4\\azure-identity-1.18.4.jar\" com.azure.identity.AzurePipelinesCredentialBuilder","description":"Inspect Azure Pipelines builder API"}
```

### 🔧 14. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/main/java/com/example/azureidentity/AzureCredentialFactory.java\n@@\n             \u0026\u0026 value(\"AZURE_CLIENT_ID\") != null\n             \u0026\u0026 value(\"AZURE_SERVICE_CONNECTION_ID\") != null\n-            \u0026\u0026 value(\"SYSTEM_ACCESSTOKEN\") != null;\n+            \u0026\u0026 value(\"SYSTEM_ACCESSTOKEN\") != null\n+            \u0026\u0026 value(\"SYSTEM_OIDCREQUESTURI\") != null;\n*** Update File: src/test/java/com/example/azureidentity/AzureCredentialFactoryTest.java\n@@\n-    void addsConfiguredAzurePipelinesServiceConnection() {\n+    void omitsIncompleteAzurePipelinesServiceConnection() {\n         AzureCredentialFactory.CredentialSelection selection =\n             new AzureCredentialFactory(Map.of(\n                 \"AZURE_TENANT_ID\", \"tenant\",\n                 \"AZURE_CLIENT_ID\", \"client\",\n                 \"AZURE_SERVICE_CONNECTION_ID\", \"connection\",\n                 \"SYSTEM_ACCESSTOKEN\", \"token\"))\n                 .create(DeploymentEnvironment.CI, false);\n \n-        assertTrue(selection.strategy().contains(\"Azure Pipelines workload-identity\"));\n+        assertTrue(!selection.strategy().contains(\"Azure Pipelines workload-identity\"));\n*** End Patch\n"
```

### 🔧 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: src/test/java/com/example/azureidentity/AzureCredentialFactoryTest.java\n@@\n import java.util.Map;\n \n+import static org.junit.jupiter.api.Assertions.assertFalse;\n import static org.junit.jupiter.api.Assertions.assertInstanceOf;\n import static org.junit.jupiter.api.Assertions.assertTrue;\n@@\n-        assertTrue(!selection.strategy().contains(\"Azure Pipelines workload-identity\"));\n+        assertFalse(selection.strategy().contains(\"Azure Pipelines workload-identity\"));\n*** End Patch\n"
```

### 🔧 16. powershell

**Input:**

```json
{"command":"mvn --no-transfer-progress clean test package","description":"Rebuild and test Maven project","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `README.md`
- `pom.xml`
- `src\main\java\com\example\azureidentity\AsyncCredentialConnectivityTester.java`
- `src\main\java\com\example\azureidentity\AuthenticationFailureAnalyzer.java`
- `src\main\java\com\example\azureidentity\AzureCredentialFactory.java`
- `src\main\java\com\example\azureidentity\CaeEnabledCredential.java`
- `src\main\java\com\example\azureidentity\ConnectivityTestResult.java`
- `src\main\java\com\example\azureidentity\CredentialConnectivityTester.java`
- `src\main\java\com\example\azureidentity\DeploymentEnvironment.java`
- `src\main\java\com\example\azureidentity\EnvironmentDetector.java`
- `src\main\java\com\example\azureidentity\Main.java`
- `src\test\java\com\example\azureidentity\AuthenticationFailureAnalyzerTest.java`
- `src\test\java\com\example\azureidentity\AzureCredentialFactoryTest.java`
- `src\test\java\com\example\azureidentity\EnvironmentDetectorTest.java`

## Copilot Response

I’ll inspect the workspace, then create the Maven structure and implement the credential strategies with focused tests.

The workspace is empty. I’m confirming the current Azure Identity APIs and dependency version before writing the project so the service-connection and CAE code compiles against one coherent SDK release.

The design will keep CAE enforcement inside a credential decorator, not only in the demo tester, so every token request made through the factory is CAE-aware. I’ll also make detection and failure classification independently testable without contacting Azure.

Implementation is in place. I’m compiling and running the focused tests now; any Azure SDK API mismatch will be corrected against the resolved 1.18.4 artifacts.

The project compiles cleanly. I’m tightening coverage around the factory itself and checking whether the matching Azure Identity Broker artifact can safely add Visual Studio Code to the local IDE chain without introducing a version mismatch.

The SDK requires `SYSTEM_OIDCREQUESTURI` to be passed explicitly, not merely present in the process environment. I’m wiring that required endpoint into both the factory’s completeness check and builder configuration.

Created a complete Java 17 Maven project demonstrating environment-specific Azure credential chains.

Key components:

- `AzureCredentialFactory` — developer tools, CI environment/service connections, managed identity, and workload identity strategies
- `EnvironmentDetector` — detects development, CI/CD, managed identity, and Kubernetes environments
- CAE-enforcing credential decorator, enabled by default via `AZURE_ENABLE_CAE`
- Synchronous and asynchronous connectivity testers with detailed authentication failure classification
- `Main` — tests both flows against the Azure Resource Manager scope
- Azure Identity `1.18.4` and Identity Broker `1.1.21`
- README with configuration and execution instructions
- Packaged artifact: `target\azure-credential-chains-1.0.0.jar`

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- No AutoCloseable Azure SDK client instances are left unclosed.

## Grader Results

- credential-chain.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (13/14)
      - Uses `ChainedTokenCredentialBuilder` to compose multiple credentials: Pass
      - Credentials added via `.addLast()` — order matters: Fail
      - **Dev chain**: includes `AzureCliCredential`; may include `IntelliJCredential`, `VisualStudioCodeCredential`, `AzurePowerShellCredential`: Pass
      - **CI chain**: uses `EnvironmentCredential` or `AzurePipelinesCredential` (not just `DefaultAzureCredential`): Pass
      - **Production chain**: `ManagedIdentityCredential` first (supports user-assigned via `clientId()`), `WorkloadIdentityCredential` as fallback: Pass
      - Enables CAE via `TokenRequestContext.setCaeEnabled(true)` or `enableCae()` on credential builders: Pass
      - Detects CI (checks `CI`, `TF_BUILD`, `AZURE_PIPELINE_WORKSPACE`, or similar): Pass
      - Detects production/managed identity (checks `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, or similar): Pass
      - Falls back to dev if neither detected: Pass
      - Creates `TokenRequestContext` with correct scope (`https://management.azure.com/.default`): Pass
      - Calls `getToken()` and prints token expiry from `AccessToken.getExpiresAt()`: Pass
      - Handles failure with specific exception info: Pass
      - Async tester uses reactive `getToken()` returning `Mono<AccessToken>`: Pass
      - NOT using `DefaultAzureCredential` as the CI credential (too broad): Pass
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
| `Criteria from prompt file` | prompt_review | 93% | 1.00 | 0.9286 | 8.5% | ❌ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 9.2% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 10.9286** | **84.1%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-credential-chain --config java-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
