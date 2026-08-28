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

import java.util.Map;
import java.util.Objects;

public final class AzureCredentialFactory {
    public static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID";
    public static final String PIPELINES_SERVICE_CONNECTION_ID = "AZURE_PIPELINES_SERVICE_CONNECTION_ID";

    private final Map<String, String> environment;

    public AzureCredentialFactory() {
        this(System.getenv());
    }

    AzureCredentialFactory(Map<String, String> environment) {
        this.environment = Map.copyOf(Objects.requireNonNull(environment, "environment"));
    }

    public CredentialSelection create(DeploymentEnvironment deploymentEnvironment, boolean enableCae) {
        Objects.requireNonNull(deploymentEnvironment, "deploymentEnvironment");
        return switch (deploymentEnvironment) {
            case DEVELOPMENT -> development(enableCae);
            case CI -> ci(enableCae);
            case PRODUCTION -> production(enableCae);
        };
    }

    private CredentialSelection development(boolean enableCae) {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .build();
        String strategy = "Azure CLI -> Azure Developer CLI -> Azure PowerShell -> IntelliJ";
        return new CredentialSelection(credential, strategy, enableCae);
    }

    private CredentialSelection ci(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String strategy;

        if (hasAll(
            "AZURE_CLIENT_ID",
            "AZURE_TENANT_ID",
            PIPELINES_SERVICE_CONNECTION_ID,
            "SYSTEM_ACCESSTOKEN",
            "SYSTEM_OIDCREQUESTURI"
        )) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .serviceConnectionId(environment.get(PIPELINES_SERVICE_CONNECTION_ID))
                .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
                .build());
            strategy = "Azure Pipelines workload identity service connection -> environment credential";
        } else {
            strategy = "Environment credential";
        }

        chain.addLast(new EnvironmentCredentialBuilder().build());
        return new CredentialSelection(chain.build(), strategy, enableCae);
    }

    private CredentialSelection production(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String managedIdentityClientId = environment.get(MANAGED_IDENTITY_CLIENT_ID);
        String managedIdentityStrategy;

        if (managedIdentityClientId == null || managedIdentityClientId.isBlank()) {
            chain.addLast(new ManagedIdentityCredentialBuilder().build());
            managedIdentityStrategy = "system-assigned managed identity";
        } else {
            chain.addLast(new ManagedIdentityCredentialBuilder()
                .clientId(managedIdentityClientId)
                .build());
            managedIdentityStrategy = "user-assigned managed identity (" + MANAGED_IDENTITY_CLIENT_ID + ")";
        }

        String strategy = managedIdentityStrategy;
        if (hasAll("AZURE_TENANT_ID", "AZURE_CLIENT_ID", "AZURE_FEDERATED_TOKEN_FILE")) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .tokenFilePath(environment.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategy += " -> workload identity";
        }
        return new CredentialSelection(
            chain.build(),
            strategy,
            enableCae
        );
    }

    private boolean hasAll(String... names) {
        for (String name : names) {
            String value = environment.get(name);
            if (value == null || value.isBlank()) {
                return false;
            }
        }
        return true;
    }
}
