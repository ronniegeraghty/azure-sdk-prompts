package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EnvironmentDetectorTest {
    @Test
    void defaultsToDevelopment() {
        assertEquals(
            DeploymentEnvironment.DEVELOPMENT,
            new EnvironmentDetector(Map.of()).detect()
        );
    }

    @Test
    void detectsCiBeforeAzureHostingMarkers() {
        assertEquals(
            DeploymentEnvironment.CI,
            new EnvironmentDetector(Map.of(
                "PIPELINE_WORKSPACE", "/agent/work",
                "IDENTITY_ENDPOINT", "http://localhost/identity"
            )).detect()
        );
    }

    @Test
    void detectsProductionManagedIdentity() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            new EnvironmentDetector(Map.of("IDENTITY_ENDPOINT", "http://localhost/identity")).detect()
        );
    }

    @Test
    void detectsProductionWorkloadIdentity() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            new EnvironmentDetector(Map.of(
                "AZURE_FEDERATED_TOKEN_FILE",
                "/var/run/secrets/azure/tokens/azure-identity-token"
            )).detect()
        );
    }

    @Test
    void ignoresFalseCiMarker() {
        assertEquals(
            DeploymentEnvironment.DEVELOPMENT,
            new EnvironmentDetector(Map.of("CI", "false")).detect()
        );
    }
}
