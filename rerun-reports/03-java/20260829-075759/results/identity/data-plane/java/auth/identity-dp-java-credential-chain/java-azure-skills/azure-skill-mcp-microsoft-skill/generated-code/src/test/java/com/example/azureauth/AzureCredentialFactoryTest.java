package com.example.azureauth;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AzureCredentialFactoryTest {
    @Test
    void buildsDeveloperToolChain() {
        CredentialSelection selection =
            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.DEVELOPMENT, true);

        assertTrue(selection.strategy().contains("Azure CLI"));
        assertTrue(selection.caeEnabled());
    }

    @Test
    void buildsGenericCiChain() {
        CredentialSelection selection =
            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.CI, false);

        assertTrue(selection.strategy().contains("Environment credential"));
        assertFalse(selection.caeEnabled());
    }

    @Test
    void buildsAzurePipelinesFederatedChain() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of(
            "AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000001",
            "AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000002",
            AzureCredentialFactory.SERVICE_CONNECTION_ID,
            "00000000-0000-0000-0000-000000000003",
            "SYSTEM_ACCESSTOKEN", "fake-system-access-token",
            "SYSTEM_OIDCREQUESTURI", "https://example.invalid/oidc"
        )).create(DeploymentEnvironment.CI, true);

        assertTrue(selection.strategy().startsWith("Azure Pipelines"));
    }

    @Test
    void buildsSystemAssignedManagedIdentityChain() {
        CredentialSelection selection =
            new AzureCredentialFactory(Map.of()).create(DeploymentEnvironment.PRODUCTION, true);

        assertTrue(selection.strategy().startsWith("system-assigned"));
    }

    @Test
    void buildsUserAssignedManagedIdentityChain() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of(
            AzureCredentialFactory.MANAGED_IDENTITY_CLIENT_ID,
            "00000000-0000-0000-0000-000000000004"
        )).create(DeploymentEnvironment.PRODUCTION, true);

        assertTrue(selection.strategy().startsWith("user-assigned"));
    }

    @Test
    void addsWorkloadIdentityFallbackWhenConfigured() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of(
            "AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000001",
            "AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000002",
            "AZURE_FEDERATED_TOKEN_FILE", "fake-token-path"
        )).create(DeploymentEnvironment.PRODUCTION, true);

        assertTrue(selection.strategy().endsWith("workload identity"));
    }
}
