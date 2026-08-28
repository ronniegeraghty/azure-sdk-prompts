package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class AzureCredentialFactoryTest {
    private final AzureCredentialFactory factory = new AzureCredentialFactory();

    @Test
    void supportsCiServicePrincipalEnvironmentVariablesWithoutPipelineFederation() {
        CredentialSelection selection = factory.create(
            DeploymentEnvironment.CI,
            false,
            Map.of(
                "AZURE_TENANT_ID", "tenant",
                "AZURE_CLIENT_ID", "client",
                "AZURE_CLIENT_SECRET", "secret"
            )
        );

        assertEquals("pipeline environment variables", selection.strategy());
    }

    @Test
    void rejectsPartialAzurePipelinesServiceConnectionConfiguration() {
        IllegalStateException failure = assertThrows(
            IllegalStateException.class,
            () -> factory.create(
                DeploymentEnvironment.CI,
                false,
                Map.of("AZURE_SERVICE_CONNECTION_ID", "connection")
            )
        );

        assertTrue(failure.getMessage().contains("SYSTEM_ACCESSTOKEN"));
        assertTrue(failure.getMessage().contains("SYSTEM_OIDCREQUESTURI"));
    }

    @Test
    void selectsUserAssignedManagedIdentityFromDedicatedVariable() {
        CredentialSelection selection = factory.create(
            DeploymentEnvironment.PRODUCTION,
            true,
            Map.of("AZURE_MANAGED_IDENTITY_CLIENT_ID", "managed-identity-client")
        );

        assertEquals("user-assigned managed identity", selection.strategy());
        assertTrue(selection.caeEnabled());
    }

    @Test
    void rejectsPartialWorkloadIdentityConfiguration() {
        IllegalStateException failure = assertThrows(
            IllegalStateException.class,
            () -> factory.create(
                DeploymentEnvironment.PRODUCTION,
                false,
                Map.of("AZURE_FEDERATED_TOKEN_FILE", "token-file")
            )
        );

        assertTrue(failure.getMessage().contains("AZURE_CLIENT_ID"));
        assertTrue(failure.getMessage().contains("AZURE_TENANT_ID"));
    }
}
