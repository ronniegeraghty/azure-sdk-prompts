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
import com.azure.identity.WorkloadIdentityCredentialBuilder;

import java.util.Map;

public final class CredentialFactory {
    private static final String MANAGED_IDENTITY_CLIENT_ID = "AZURE_MANAGED_IDENTITY_CLIENT_ID";

    public CredentialSelection create(DeploymentEnvironment environment, boolean enableCae) {
        return create(environment, enableCae, System.getenv());
    }

    CredentialSelection create(
        DeploymentEnvironment environment,
        boolean enableCae,
        Map<String, String> variables
    ) {
        CredentialSelection selection = switch (environment) {
            case DEVELOPMENT -> developerCredential();
            case CI -> ciCredential(variables);
            case PRODUCTION -> productionCredential(variables);
        };

        TokenCredential credential = enableCae
            ? new CaeEnabledCredential(selection.credential())
            : selection.credential();
        return new CredentialSelection(credential, selection.strategy(), enableCae);
    }

    private CredentialSelection developerCredential() {
        TokenCredential credential = new ChainedTokenCredentialBuilder()
            .addLast(new AzureCliCredentialBuilder().build())
            .addLast(new AzureDeveloperCliCredentialBuilder().build())
            .addLast(new AzurePowerShellCredentialBuilder().build())
            .addLast(new IntelliJCredentialBuilder().build())
            .build();

        return new CredentialSelection(
            credential,
            "Azure CLI -> Azure Developer CLI -> Azure PowerShell -> IntelliJ",
            false
        );
    }

    private CredentialSelection ciCredential(Map<String, String> variables) {
        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder()
            .addLast(new EnvironmentCredentialBuilder().build());
        String strategy = "EnvironmentCredential (secret, certificate, or federated environment variables)";

        if (hasAzurePipelinesServiceConnection(variables)) {
            chain.addLast(new AzurePipelinesCredentialBuilder()
                .clientId(variables.get("AZURESUBSCRIPTION_CLIENT_ID"))
                .tenantId(variables.get("AZURESUBSCRIPTION_TENANT_ID"))
                .serviceConnectionId(variables.get("AZURESUBSCRIPTION_SERVICE_CONNECTION_ID"))
                .systemAccessToken(variables.get("SYSTEM_ACCESSTOKEN"))
                .build());
            strategy += " -> Azure Pipelines workload identity service connection";
        } else if (hasText(variables.get("TF_BUILD"))) {
            strategy += " (Azure Pipelines service connection variables incomplete)";
        }

        return new CredentialSelection(chain.build(), strategy, false);
    }

    private CredentialSelection productionCredential(Map<String, String> variables) {
        ManagedIdentityCredentialBuilder managedIdentity = new ManagedIdentityCredentialBuilder();
        String userAssignedClientId = variables.get(MANAGED_IDENTITY_CLIENT_ID);
        String strategy;

        if (hasText(userAssignedClientId)) {
            managedIdentity.clientId(userAssignedClientId);
            strategy = "User-assigned managed identity (" + MANAGED_IDENTITY_CLIENT_ID + ")";
        } else {
            strategy = "System-assigned managed identity";
        }

        ChainedTokenCredentialBuilder chain = new ChainedTokenCredentialBuilder()
            .addLast(managedIdentity.build());

        if (hasWorkloadIdentityConfiguration(variables)) {
            chain.addLast(new WorkloadIdentityCredentialBuilder()
                .clientId(variables.get("AZURE_CLIENT_ID"))
                .tenantId(variables.get("AZURE_TENANT_ID"))
                .tokenFilePath(variables.get("AZURE_FEDERATED_TOKEN_FILE"))
                .build());
            strategy += " -> Kubernetes workload identity fallback";
        } else {
            strategy += " (workload identity fallback not configured)";
        }

        return new CredentialSelection(chain.build(), strategy, false);
    }

    private boolean hasAzurePipelinesServiceConnection(Map<String, String> variables) {
        return hasText(variables.get("AZURESUBSCRIPTION_CLIENT_ID"))
            && hasText(variables.get("AZURESUBSCRIPTION_TENANT_ID"))
            && hasText(variables.get("AZURESUBSCRIPTION_SERVICE_CONNECTION_ID"))
            && hasText(variables.get("SYSTEM_ACCESSTOKEN"))
            && hasText(variables.get("SYSTEM_OIDCREQUESTURI"));
    }

    private boolean hasWorkloadIdentityConfiguration(Map<String, String> variables) {
        return hasText(variables.get("AZURE_CLIENT_ID"))
            && hasText(variables.get("AZURE_TENANT_ID"))
            && hasText(variables.get("AZURE_FEDERATED_TOKEN_FILE"));
    }

    private boolean hasText(String value) {
        return value != null && !value.isBlank();
    }
}
