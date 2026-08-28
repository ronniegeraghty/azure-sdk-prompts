package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EnvironmentDetectorTest {
    private final EnvironmentDetector detector = new EnvironmentDetector();

    @Test
    void defaultsToDeveloperEnvironment() {
        assertEquals(DeploymentEnvironment.DEV, detector.detect(Map.of()));
    }

    @Test
    void detectsCiBeforeProductionMarkers() {
        assertEquals(
            DeploymentEnvironment.CI,
            detector.detect(Map.of("TF_BUILD", "True", "IDENTITY_ENDPOINT", "http://localhost/identity"))
        );
    }

    @Test
    void detectsManagedIdentityProductionEnvironment() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            detector.detect(Map.of("IDENTITY_ENDPOINT", "http://localhost/identity"))
        );
    }

    @Test
    void detectsKubernetesWorkloadIdentityEnvironment() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            detector.detect(Map.of("AZURE_FEDERATED_TOKEN_FILE", "/var/run/secrets/azure/tokens/azure-identity-token"))
        );
    }
}
