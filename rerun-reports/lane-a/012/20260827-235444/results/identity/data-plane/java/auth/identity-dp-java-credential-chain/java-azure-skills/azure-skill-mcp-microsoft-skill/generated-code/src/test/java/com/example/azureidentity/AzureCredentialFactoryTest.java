package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AzureCredentialFactoryTest {
    @Test
    void selectsUserAssignedManagedIdentityWhenConfigured() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of(
            AzureCredentialFactory.MANAGED_IDENTITY_CLIENT_ID,
            "00000000-0000-0000-0000-000000000000"
        )).create(DeploymentEnvironment.PRODUCTION, true);

        assertTrue(selection.strategy().startsWith("user-assigned managed identity"));
        assertTrue(selection.caeEnabled());
    }

    @Test
    void selectsSystemAssignedManagedIdentityByDefault() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of())
            .create(DeploymentEnvironment.PRODUCTION, false);

        assertTrue(selection.strategy().startsWith("system-assigned managed identity"));
        assertEquals(false, selection.caeEnabled());
    }

    @Test
    void addsWorkloadIdentityFallbackWhenFederationIsConfigured() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of(
            "AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000001",
            "AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000002",
            "AZURE_FEDERATED_TOKEN_FILE", "target/federated-token"
        )).create(DeploymentEnvironment.PRODUCTION, true);

        assertTrue(selection.strategy().endsWith("-> workload identity"));
    }

    @Test
    void buildsEnvironmentCredentialForGenericCi() {
        CredentialSelection selection = new AzureCredentialFactory(Map.of())
            .create(DeploymentEnvironment.CI, true);

        assertEquals("Environment credential", selection.strategy());
    }
}
