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
import com.azure.identity.VisualStudioCodeCredentialBuilder;
import com.azure.identity.WorkloadIdentityCredentialBuilder;

import java.util.Map;
import java.util.Objects;

public final class AzureCredentialFactory {
    public record CredentialSelection(
        TokenCredential credential,
        DeploymentEnvironment environment,
        String strategy,
        boolean caeEnabled) {
    }

    private final Map<String, String> environment;

    public AzureCredentialFactory() {
        this(System.getenv());
    }

    AzureCredentialFactory(Map<String, String> environment) {
        this.environment = Map.copyOf(environment);
    }

    public CredentialSelection create(DeploymentEnvironment deploymentEnvironment, boolean enableCae) {
        Objects.requireNonNull(deploymentEnvironment, "deploymentEnvironment");

        Strategy strategy = switch (deploymentEnvironment) {
            case DEVELOPMENT -> developerCredential();
            case CI -> ciCredential();
            case PRODUCTION -> productionCredential();
        };
        TokenCredential credential = enableCae
            ? new CaeEnabledCredential(strategy.credential())
            : strategy.credential();

        return new CredentialSelection(
            credential, deploymentEnvironment, strategy.description(), enableCae);
    }

    private Strategy developerCredential() {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new VisualStudioCodeCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .build();
        return new Strategy(
            credential,
            "Developer tools: Azure CLI -> Azure Developer CLI -> Visual Studio Code -> "
                + "IntelliJ -> Azure PowerShell");
    }

    private Strategy ciCredential() {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        StringBuilder description = new StringBuilder("Pipeline environment variables");

        if (hasAzurePipelinesServiceConnection()) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .serviceConnectionId(environment.get("AZURE_SERVICE_CONNECTION_ID"))
                .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
                .build());
            description.append(" -> Azure Pipelines workload-identity service connection");
        }

        chain.addFirst(new EnvironmentCredentialBuilder().build());
        return new Strategy(chain.build(), description.toString());
    }

    private Strategy productionCredential() {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();
        String clientId = value("AZURE_CLIENT_ID");
        String identityType = "system-assigned managed identity";

        if (clientId != null) {
            managedIdentity.clientId(clientId);
            identityType = "user-assigned managed identity (AZURE_CLIENT_ID)";
        }
        chain.addLast(managedIdentity.build());

        String description = identityType;
        if (hasWorkloadIdentityConfiguration()) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .tokenFilePath(environment.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            description += " -> Kubernetes workload identity";
        } else {
            description += " (workload identity fallback not configured)";
        }

        return new Strategy(chain.build(), description);
    }

    private boolean hasAzurePipelinesServiceConnection() {
        return value("AZURE_TENANT_ID") != null
            && value("AZURE_CLIENT_ID") != null
            && value("AZURE_SERVICE_CONNECTION_ID") != null
            && value("SYSTEM_ACCESSTOKEN") != null
            && value("SYSTEM_OIDCREQUESTURI") != null;
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return value("AZURE_TENANT_ID") != null
            && value("AZURE_CLIENT_ID") != null
            && value("AZURE_FEDERATED_TOKEN_FILE") != null;
    }

    private String value(String name) {
        String value = environment.get(name);
        return value == null || value.isBlank() ? null : value;
    }

    private record Strategy(TokenCredential credential, String description) {
    }
}
