package com.example.azureauth;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EnvironmentDetectorTest {
    @Test
    void defaultsToDevelopment() {
        assertEquals(DeploymentEnvironment.DEVELOPMENT, new EnvironmentDetector(Map.of()).detect());
    }

    @Test
    void detectsCiBeforeAzureHostingSignals() {
        EnvironmentDetector detector = new EnvironmentDetector(Map.of(
            "TF_BUILD", "True",
            "IDENTITY_ENDPOINT", "http://localhost/identity"
        ));

        assertEquals(DeploymentEnvironment.CI, detector.detect());
    }

    @Test
    void detectsProductionFromManagedIdentityEndpoint() {
        EnvironmentDetector detector =
            new EnvironmentDetector(Map.of("IDENTITY_ENDPOINT", "http://localhost/identity"));

        assertEquals(DeploymentEnvironment.PRODUCTION, detector.detect());
    }

    @Test
    void detectsProductionFromWorkloadIdentity() {
        EnvironmentDetector detector = new EnvironmentDetector(Map.of(
            "AZURE_TENANT_ID", "tenant",
            "AZURE_CLIENT_ID", "client",
            "AZURE_FEDERATED_TOKEN_FILE", "token"
        ));

        assertEquals(DeploymentEnvironment.PRODUCTION, detector.detect());
    }

    @Test
    void explicitOverrideWins() {
        EnvironmentDetector detector = new EnvironmentDetector(Map.of(
            "APP_ENVIRONMENT", "dev",
            "TF_BUILD", "True"
        ));

        assertEquals(DeploymentEnvironment.DEVELOPMENT, detector.detect());
    }

    @Test
    void rejectsUnknownOverride() {
        assertThrows(
            IllegalArgumentException.class,
            () -> new EnvironmentDetector(Map.of("APP_ENVIRONMENT", "staging")).detect()
        );
    }
}
