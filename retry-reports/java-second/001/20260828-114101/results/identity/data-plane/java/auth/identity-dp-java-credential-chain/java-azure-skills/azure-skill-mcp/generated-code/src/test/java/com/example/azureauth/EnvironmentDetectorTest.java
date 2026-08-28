package com.example.azureauth;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EnvironmentDetectorTest {
    private final EnvironmentDetector detector = new EnvironmentDetector();

    @Test
    void defaultsToDevelopment() {
        assertEquals(DeploymentEnvironment.DEVELOPMENT, detector.detect(Map.of()));
    }

    @Test
    void detectsCiBeforeProductionMarkers() {
        assertEquals(
            DeploymentEnvironment.CI,
            detector.detect(Map.of("TF_BUILD", "True", "IDENTITY_ENDPOINT", "http://localhost"))
        );
    }

    @Test
    void detectsManagedIdentityAsProduction() {
        assertEquals(
            DeploymentEnvironment.PRODUCTION,
            detector.detect(Map.of("IDENTITY_ENDPOINT", "http://localhost"))
        );
    }

    @Test
    void explicitOverrideWins() {
        assertEquals(
            DeploymentEnvironment.DEVELOPMENT,
            detector.detect(Map.of("APP_DEPLOYMENT_ENVIRONMENT", "local", "TF_BUILD", "True"))
        );
    }

    @Test
    void rejectsUnknownOverride() {
        assertThrows(
            IllegalArgumentException.class,
            () -> detector.detect(Map.of("APP_DEPLOYMENT_ENVIRONMENT", "staging"))
        );
    }
}
