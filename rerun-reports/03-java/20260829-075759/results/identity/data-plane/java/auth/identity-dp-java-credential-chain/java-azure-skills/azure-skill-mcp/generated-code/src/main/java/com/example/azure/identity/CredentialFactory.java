package com.example.azure.identity;

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

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public final class CredentialFactory {
    public static final String USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID
        = "AZURE_MANAGED_IDENTITY_CLIENT_ID";

    private final Map<String, String> environment;

    public CredentialFactory() {
        this(System.getenv());
    }

    CredentialFactory(Map<String, String> environment) {
        this.environment = Map.copyOf(environment);
    }

    public BuiltCredential create(DeploymentEnvironment deploymentEnvironment, boolean caeEnabled) {
        BuiltCredential built = switch (deploymentEnvironment) {
            case DEV -> createDevelopmentCredential();
            case CI -> createCiCredential();
            case PRODUCTION -> createProductionCredential();
        };

        TokenCredential credential = caeEnabled ? new CaeTokenCredential(built.credential()) : built.credential();
        return new BuiltCredential(credential, built.strategy(), caeEnabled);
    }

    private BuiltCredential createDevelopmentCredential() {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .addLast(new VisualStudioCodeCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .build();

        return new BuiltCredential(
            credential,
            "Azure CLI -> IntelliJ -> Visual Studio Code -> Azure Developer CLI -> Azure PowerShell",
            false
        );
    }

    private BuiltCredential createCiCredential() {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        List<String> strategies = new ArrayList<>();

        if (hasAzurePipelinesServiceConnectionConfiguration()) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .tenantId(required("AZURE_TENANT_ID"))
                .clientId(required("AZURE_CLIENT_ID"))
                .serviceConnectionId(required("AZURE_SERVICE_CONNECTION_ID"))
                .systemAccessToken(required("SYSTEM_ACCESSTOKEN"))
                .build());
            strategies.add("Azure Pipelines workload identity service connection");
        }

        chain.addLast(new EnvironmentCredentialBuilder().build());
        strategies.add("environment-configured service principal (secret or certificate)");

        return new BuiltCredential(chain.build(), String.join(" -> ", strategies), false);
    }

    private BuiltCredential createProductionCredential() {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String managedIdentityClientId = environment.get(USER_ASSIGNED_MANAGED_IDENTITY_CLIENT_ID);
        String managedIdentityStrategy;

        if (isPresent(managedIdentityClientId)) {
            chain.addLast(new ManagedIdentityCredentialBuilder().clientId(managedIdentityClientId).build());
            managedIdentityStrategy = "user-assigned managed identity";
        } else {
            chain.addLast(new ManagedIdentityCredentialBuilder().build());
            managedIdentityStrategy = "system-assigned managed identity";
        }

        String strategy = managedIdentityStrategy;
        if (hasWorkloadIdentityConfiguration()) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .tenantId(required("AZURE_TENANT_ID"))
                .clientId(required("AZURE_CLIENT_ID"))
                .tokenFilePath(required("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategy += " -> Kubernetes workload identity";
        }

        return new BuiltCredential(chain.build(), strategy, false);
    }

    private boolean hasAzurePipelinesServiceConnectionConfiguration() {
        return has("TF_BUILD")
            && has("SYSTEM_OIDCREQUESTURI")
            && has("SYSTEM_ACCESSTOKEN")
            && has("AZURE_SERVICE_CONNECTION_ID")
            && has("AZURE_TENANT_ID")
            && has("AZURE_CLIENT_ID");
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return has("AZURE_FEDERATED_TOKEN_FILE")
            && has("AZURE_TENANT_ID")
            && has("AZURE_CLIENT_ID");
    }

    private boolean has(String name) {
        return isPresent(environment.get(name));
    }

    private String required(String name) {
        String value = environment.get(name);
        if (!isPresent(value)) {
            throw new IllegalStateException("Required environment variable is missing: " + name);
        }
        return value;
    }

    private static boolean isPresent(String value) {
        return value != null && !value.isBlank();
    }

    public record BuiltCredential(TokenCredential credential, String strategy, boolean caeEnabled) {
    }
}
