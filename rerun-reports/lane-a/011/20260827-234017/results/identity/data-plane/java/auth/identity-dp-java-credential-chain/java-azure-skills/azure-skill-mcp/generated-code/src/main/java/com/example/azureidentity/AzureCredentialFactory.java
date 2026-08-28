package com.example.azureidentity;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.AzureCliCredentialBuilder;
import com.azure.identity.AzureDeveloperCliCredentialBuilder;
import com.azure.identity.AzurePipelinesCredentialBuilder;
import com.azure.identity.AzurePowerShellCredentialBuilder;
import com.azure.identity.ChainedTokenCredentialBuilder;
import com.azure.identity.EnvironmentCredentialBuilder;
import com.azure.identity.IntelliJCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.identity.WorkloadIdentityCredentialBuilder;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public final class AzureCredentialFactory {
    private static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID";
    private static final String[] PIPELINES_CREDENTIAL_VARIABLES = {
        "AZURE_CLIENT_ID",
        "AZURE_TENANT_ID",
        "AZURE_SERVICE_CONNECTION_ID",
        "SYSTEM_ACCESSTOKEN",
        "SYSTEM_OIDCREQUESTURI"
    };
    private static final String[] PIPELINES_SERVICE_CONNECTION_MARKERS = {
        "AZURE_SERVICE_CONNECTION_ID",
        "SYSTEM_ACCESSTOKEN",
        "SYSTEM_OIDCREQUESTURI"
    };
    private static final String[] WORKLOAD_IDENTITY_VARIABLES = {
        "AZURE_CLIENT_ID",
        "AZURE_TENANT_ID",
        "AZURE_FEDERATED_TOKEN_FILE"
    };

    public CredentialSelection create(DeploymentEnvironment environment, boolean enableCae) {
        return create(environment, enableCae, System.getenv());
    }

    CredentialSelection create(
        DeploymentEnvironment environment,
        boolean enableCae,
        Map<String, String> variables
    ) {
        BuiltCredential built = switch (environment) {
            case DEV -> buildDeveloperCredential();
            case CI -> buildCiCredential(variables);
            case PRODUCTION -> buildProductionCredential(variables);
        };

        TokenCredential credential = new CaeAwareTokenCredential(built.credential(), enableCae);
        return new CredentialSelection(credential, environment, built.strategy(), enableCae);
    }

    private BuiltCredential buildDeveloperCredential() {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .build();

        return new BuiltCredential(
            credential,
            "Developer tools: Azure CLI -> Azure Developer CLI -> Azure PowerShell -> IntelliJ"
        );
    }

    private BuiltCredential buildCiCredential(Map<String, String> variables) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        List<String> strategies = new ArrayList<>();

        if (hasAny(variables, PIPELINES_SERVICE_CONNECTION_MARKERS)) {
            requireAll(variables, "Azure Pipelines workload federation", PIPELINES_CREDENTIAL_VARIABLES);
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .clientId(variables.get("AZURE_CLIENT_ID"))
                .tenantId(variables.get("AZURE_TENANT_ID"))
                .serviceConnectionId(variables.get("AZURE_SERVICE_CONNECTION_ID"))
                .systemAccessToken(variables.get("SYSTEM_ACCESSTOKEN"))
                .build());
            strategies.add("Azure Pipelines service connection");
        }

        chain.addLast(new EnvironmentCredentialBuilder().build());
        strategies.add("pipeline environment variables");

        return new BuiltCredential(chain.build(), String.join(" -> ", strategies));
    }

    private BuiltCredential buildProductionCredential(Map<String, String> variables) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String managedIdentityClientId = variables.get(MANAGED_IDENTITY_CLIENT_ID);

        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();
        String managedIdentityStrategy = "system-assigned managed identity";
        if (isNonBlank(managedIdentityClientId)) {
            managedIdentity.clientId(managedIdentityClientId);
            managedIdentityStrategy = "user-assigned managed identity";
        }
        chain.addLast(managedIdentity.build());

        List<String> strategies = new ArrayList<>();
        strategies.add(managedIdentityStrategy);
        if (isNonBlank(variables.get("AZURE_FEDERATED_TOKEN_FILE"))) {
            requireAll(variables, "Kubernetes workload identity", WORKLOAD_IDENTITY_VARIABLES);
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .clientId(variables.get("AZURE_CLIENT_ID"))
                .tenantId(variables.get("AZURE_TENANT_ID"))
                .tokenFilePath(variables.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategies.add("Kubernetes workload identity");
        }

        return new BuiltCredential(chain.build(), String.join(" -> ", strategies));
    }

    private static boolean hasAny(Map<String, String> variables, String[] names) {
        for (String name : names) {
            if (isNonBlank(variables.get(name))) {
                return true;
            }
        }
        return false;
    }

    private static void requireAll(Map<String, String> variables, String label, String[] names) {
        List<String> missing = new ArrayList<>();
        for (String name : names) {
            if (!isNonBlank(variables.get(name))) {
                missing.add(name);
            }
        }
        if (!missing.isEmpty()) {
            throw new IllegalStateException(label + " is partially configured; missing: " + String.join(", ", missing));
        }
    }

    private static boolean isNonBlank(String value) {
        return value != null && !value.isBlank();
    }

    private record BuiltCredential(TokenCredential credential, String strategy) {
    }
}
