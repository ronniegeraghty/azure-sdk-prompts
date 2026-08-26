package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertInstanceOf;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AzureCredentialFactoryTest {
    @Test
    void wrapsCredentialsWhenCaeIsEnabled() {
        AzureCredentialFactory.CredentialSelection selection =
            new AzureCredentialFactory(Map.of()).create(
                DeploymentEnvironment.DEVELOPMENT, true);

        assertTrue(selection.caeEnabled());
        assertInstanceOf(CaeEnabledCredential.class, selection.credential());
    }

    @Test
    void selectsUserAssignedManagedIdentityFromClientId() {
        AzureCredentialFactory.CredentialSelection selection =
            new AzureCredentialFactory(Map.of("AZURE_CLIENT_ID", "user-assigned-client"))
                .create(DeploymentEnvironment.PRODUCTION, false);

        assertTrue(selection.strategy().startsWith("user-assigned managed identity"));
    }

    @Test
    void omitsIncompleteAzurePipelinesServiceConnection() {
        AzureCredentialFactory.CredentialSelection selection =
            new AzureCredentialFactory(Map.of(
                "AZURE_TENANT_ID", "tenant",
                "AZURE_CLIENT_ID", "client",
                "AZURE_SERVICE_CONNECTION_ID", "connection",
                "SYSTEM_ACCESSTOKEN", "token"))
                .create(DeploymentEnvironment.CI, false);

        assertFalse(selection.strategy().contains("Azure Pipelines workload-identity"));
    }

    @Test
    void addsConfiguredWorkloadIdentityFallback() {
        AzureCredentialFactory.CredentialSelection selection =
            new AzureCredentialFactory(Map.of(
                "AZURE_TENANT_ID", "tenant",
                "AZURE_CLIENT_ID", "client",
                "AZURE_FEDERATED_TOKEN_FILE", "token-file"))
                .create(DeploymentEnvironment.PRODUCTION, false);

        assertTrue(selection.strategy().endsWith("Kubernetes workload identity"));
    }
}
