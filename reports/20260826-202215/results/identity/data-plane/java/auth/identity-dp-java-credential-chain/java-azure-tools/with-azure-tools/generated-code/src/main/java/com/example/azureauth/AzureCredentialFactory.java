package com.example.azureauth;

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
    public static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID";
    public static final String SERVICE_CONNECTION_ID = "AZURESUBSCRIPTION_SERVICE_CONNECTION_ID";

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
            .addLast(new VisualStudioCodeCredentialBuilder().build())
            .build();

        return new CredentialSelection(
            credential,
            "Developer tools: Azure CLI -> Azure Developer CLI -> Azure PowerShell -> IntelliJ -> VS Code",
            enableCae
        );
    }

    private CredentialSelection ci(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder()
            .addLast(new EnvironmentCredentialBuilder().build());

        String strategy = "EnvironmentCredential (service principal secret or certificate)";
        if (hasAzurePipelinesServiceConnectionConfiguration()) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .serviceConnectionId(environment.get(SERVICE_CONNECTION_ID))
                .systemAccessToken(environment.get("SYSTEM_ACCESSTOKEN"))
                .build());
            strategy += " -> AzurePipelinesCredential (OIDC service connection)";
        } else if (isPresent(environment.get("TF_BUILD"))) {
            strategy += " (AzurePipelinesCredential skipped: OIDC service-connection variables are incomplete)";
        }

        return new CredentialSelection(chain.build(), strategy, enableCae);
    }

    private CredentialSelection production(boolean enableCae) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder();
        String managedIdentityClientId = environment.get(MANAGED_IDENTITY_CLIENT_ID);

        if (isPresent(managedIdentityClientId)) {
            chain.addLast(new ManagedIdentityCredentialBuilder().clientId(managedIdentityClientId).build());
        } else {
            chain.addLast(new ManagedIdentityCredentialBuilder().build());
        }

        String strategy = isPresent(managedIdentityClientId)
            ? "User-assigned managed identity"
            : "System-assigned managed identity";

        if (hasWorkloadIdentityConfiguration()) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .tenantId(environment.get("AZURE_TENANT_ID"))
                .clientId(environment.get("AZURE_CLIENT_ID"))
                .tokenFilePath(environment.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategy += " -> Workload identity fallback";
        } else {
            strategy += " (workload identity fallback inactive: federation variables are not present)";
        }

        return new CredentialSelection(chain.build(), strategy, enableCae);
    }

    private boolean hasAzurePipelinesServiceConnectionConfiguration() {
        return allPresent(
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            SERVICE_CONNECTION_ID,
            "SYSTEM_ACCESSTOKEN",
            "SYSTEM_OIDCREQUESTURI"
        );
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return allPresent("AZURE_TENANT_ID", "AZURE_CLIENT_ID", "AZURE_FEDERATED_TOKEN_FILE");
    }

    private boolean allPresent(String... names) {
        for (String name : names) {
            if (!isPresent(environment.get(name))) {
                return false;
            }
        }
        return true;
    }

    private static boolean isPresent(String value) {
        return value != null && !value.isBlank();
    }
}
