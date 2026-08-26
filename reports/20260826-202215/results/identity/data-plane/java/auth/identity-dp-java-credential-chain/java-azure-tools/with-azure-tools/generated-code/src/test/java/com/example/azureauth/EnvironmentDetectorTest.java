package com.example.azureauth;

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
    void detectsCiBeforeProductionSignals() {
        assertEquals(
            DeploymentEnvironment.CI,
            new EnvironmentDetector(Map.of(
                "TF_BUILD", "True",
                "IDENTITY_ENDPOINT", "http://localhost/identity"
            )).detect()
        );
    }

    @Test
    void detectsProductionFromWorkloadIdentity() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            new EnvironmentDetector(Map.of(
                "AZURE_FEDERATED_TOKEN_FILE", "/var/run/secrets/azure/tokens/identity-token"
            )).detect()
        );
    }

    @Test
    void ignoresFalseCiFlag() {
        assertEquals(
            DeploymentEnvironment.DEVELOPMENT,
            new EnvironmentDetector(Map.of("CI", "false")).detect()
        );
    }
}
