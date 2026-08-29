package com.example.azureauth;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.ConfigurationBuilder;
import com.azure.identity.AzureCliCredentialBuilder;
import com.azure.identity.AzureDeveloperCliCredentialBuilder;
import com.azure.identity.AzurePipelinesCredentialBuilder;
import com.azure.identity.ChainedTokenCredentialBuilder;
import com.azure.identity.EnvironmentCredentialBuilder;
import com.azure.identity.IntelliJCredentialBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.identity.WorkloadIdentityCredentialBuilder;

import java.util.Map;
import java.util.Objects;

public final class AzureCredentialFactory {
    public static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID";
    public static final String SERVICE_CONNECTION_ID = "AZURE_SERVICE_CONNECTION_ID";

    private final Map<String, String> environment;

    public AzureCredentialFactory() {
        this(System.getenv());
    }

    AzureCredentialFactory(Map<String, String> environment) {
        this.environment = Map.copyOf(environment);
    }

    public CredentialSelection create(DeploymentEnvironment deploymentEnvironment, boolean enableCae) {
        Objects.requireNonNull(deploymentEnvironment, "deploymentEnvironment");
        return switch (deploymentEnvironment) {
            case DEVELOPMENT -> developmentCredential(enableCae);
            case CI -> ciCredential(enableCae);
            case PRODUCTION -> productionCredential(enableCae);
        };
    }

    private CredentialSelection developmentCredential(boolean enableCae) {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .build();

        return new CredentialSelection(
            credential,
            "Azure CLI -> Azure Developer CLI -> IntelliJ",
            enableCae
        );
    }

    private CredentialSelection ciCredential(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String strategy;

        if (hasAzurePipelinesFederationConfiguration()) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))
                .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
                .configuration(new ConfigurationBuilder()
                    .putProperty("SYSTEM_OIDCREQUESTURI", environment.get("SYSTEM_OIDCREQUESTURI"))
                    .build())
                .build());
            strategy = "Azure Pipelines workload federation -> environment credential";
        } else {
            strategy = "Environment credential";
        }

        chain.addLast(new EnvironmentCredentialBuilder().build());
        return new CredentialSelection(chain.build(), strategy, enableCae);
    }

    private CredentialSelection productionCredential(boolean enableCae) {
        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = environment.get(MANAGED_IDENTITY_CLIENT_ID);
        String managedIdentityDescription = "system-assigned managed identity";

        if (hasText(managedIdentityClientId)) {
            managedIdentity.clientId(managedIdentityClientId);
            managedIdentityDescription = "user-assigned managed identity";
        }

        ChainedTokenCredentialBuilder chain =
            new ChainedTokenCredentialBuilder().addLast(managedIdentity.build());
        String strategy = managedIdentityDescription;

        if (hasWorkloadIdentityConfiguration()) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .tokenFilePath(environment.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategy += " -> workload identity";
        } else {
            strategy += " (workload identity fallback not configured)";
        }

        return new CredentialSelection(
            chain.build(),
            strategy,
            enableCae
        );
    }

    private boolean hasAzurePipelinesFederationConfiguration() {
        return hasText(environment.get("AZURE_TENANT_ID"))
            && hasText(environment.get("AZURE_CLIENT_ID"))
            && hasText(environment.get(SERVICE_CONNECTION_ID))
            && hasText(environment.get("SYSTEM_ACCESSTOKEN"))
            && hasText(environment.get("SYSTEM_OIDCREQUESTURI"));
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return hasText(environment.get("AZURE_TENANT_ID"))
            && hasText(environment.get("AZURE_CLIENT_ID"))
            && hasText(environment.get("AZURE_FEDERATED_TOKEN_FILE"));
    }

    private static boolean hasText(String value) {
        return value != null && !value.isBlank();
    }
}
