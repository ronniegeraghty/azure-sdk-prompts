package com.example.azurecredentials;

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
    public static final String SERVICE_CONNECTION_ID = "AZURESUBSCRIPTION_SERVICE_CONNECTION_ID";
    public static final String SYSTEM_ACCESS_TOKEN = "SYSTEM_ACCESSTOKEN";

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
            case DEVELOPMENT -> localDevelopment(enableCae);
            case CI -> ciPipeline(enableCae);
            case PRODUCTION -> production(enableCae);
        };
    }

    private CredentialSelection localDevelopment(boolean enableCae) {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .build();

        return new CredentialSelection(
            credential,
            "Developer tools: Azure CLI -> Azure Developer CLI -> Azure PowerShell -> IntelliJ",
            enableCae);
    }

    private CredentialSelection ciPipeline(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder()
            .addLast(new EnvironmentCredentialBuilder().build());
        String strategy = "Pipeline environment variables";

        if (hasAzurePipelinesServiceConnection()) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .tenantId(required("AZURE_TENANT_ID"))
                .clientId(required("AZURE_CLIENT_ID"))
                .serviceConnectionId(required(SERVICE_CONNECTION_ID))
                .systemAccessToken(required(SYSTEM_ACCESS_TOKEN))
                .build());
            strategy += " -> Azure Pipelines service connection";
        } else {
            strategy += " (service-connection fallback not configured)";
        }

        return new CredentialSelection(chain.build(), strategy, enableCae);
    }

    private CredentialSelection production(boolean enableCae) {
        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = firstValue(
            MANAGED_IDENTITY_CLIENT_ID,
            "AZURE_CLIENT_ID");
        String managedIdentityDescription = "system-assigned managed identity";
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            managedIdentity.clientId(managedIdentityClientId);
            managedIdentityDescription =
                "user-assigned managed identity from "
                    + (hasValue(MANAGED_IDENTITY_CLIENT_ID)
                        ? MANAGED_IDENTITY_CLIENT_ID
                        : "AZURE_CLIENT_ID");
        }

        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(managedIdentity.build())
            .addLast(new WorkloadIdentityCredentialBuilder().build())
            .build();

        return new CredentialSelection(
            credential,
            "Production: " + managedIdentityDescription + " -> Kubernetes workload identity",
            enableCae);
    }

    private boolean hasAzurePipelinesServiceConnection() {
        return hasValue("AZURE_TENANT_ID")
            && hasValue("AZURE_CLIENT_ID")
            && hasValue(SERVICE_CONNECTION_ID)
            && hasValue(SYSTEM_ACCESS_TOKEN);
    }

    private boolean hasValue(String name) {
        String value = environment.get(name);
        return value != null && !value.isBlank();
    }

    private String required(String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is missing: " + name);
        }
        return value;
    }

    private String firstValue(String... names) {
        for (String name : names) {
            if (hasValue(name)) {
                return environment.get(name);
            }
        }
        return null;
    }
}
