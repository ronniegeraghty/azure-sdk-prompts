package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EnvironmentDetectorTest {
    @Test
    void defaultsToDevelopment() {
        assertEquals(
            DeploymentEnvironment.DEVELOPMENT,
            new EnvironmentDetector(Map.of(), () -> false).detect());
    }

    @Test
    void detectsCiBeforeManagedIdentity() {
        assertEquals(
            DeploymentEnvironment.CI,
            new EnvironmentDetector(
                Map.of("GITHUB_ACTIONS", "true", "IDENTITY_ENDPOINT", "http://localhost"),
                () -> true).detect());
    }

    @Test
    void detectsProductionFromWorkloadIdentity() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            new EnvironmentDetector(
                Map.of(
                    "AZURE_CLIENT_ID", "client",
                    "AZURE_TENANT_ID", "tenant",
                    "AZURE_FEDERATED_TOKEN_FILE", "token"),
                () -> false).detect());
    }

    @Test
    void detectsProductionFromImdsProbe() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            new EnvironmentDetector(Map.of(), () -> true).detect());
    }
}
