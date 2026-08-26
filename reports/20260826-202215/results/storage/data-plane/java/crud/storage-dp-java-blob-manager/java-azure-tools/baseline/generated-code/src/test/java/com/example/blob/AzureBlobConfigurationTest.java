package com.example.blob;

import com.azure.core.http.policy.HttpLogDetailLevel;
import org.junit.jupiter.api.Test;

import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class AzureBlobConfigurationTest {
    @Test
    void readsSettingsFromEnvironment() {
        Map<String, String> environment = new HashMap<>();
        environment.put("AZURE_STORAGE_ENDPOINT", "https://example.blob.core.windows.net");
        environment.put("AZURE_STORAGE_CONTAINER", "documents");
        environment.put("AZURE_STORAGE_MAX_RETRIES", "7");
        environment.put("AZURE_STORAGE_RETRY_DELAY_SECONDS", "3");
        environment.put("AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS", "45");
        environment.put("AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", "180");
        environment.put("AZURE_STORAGE_LOG_LEVEL", "headers");

        AzureBlobConfiguration.Settings settings = AzureBlobConfiguration.Settings.from(environment);

        assertEquals("https://example.blob.core.windows.net", settings.endpoint());
        assertEquals("documents", settings.containerName());
        assertEquals(7, settings.maxRetries());
        assertEquals(Duration.ofSeconds(3), settings.retryDelay());
        assertEquals(Duration.ofSeconds(45), settings.maxRetryDelay());
        assertEquals(Duration.ofSeconds(180), settings.requestTimeout());
        assertEquals(HttpLogDetailLevel.HEADERS, settings.logLevel());
    }

    @Test
    void rejectsInsecureEndpoint() {
        Map<String, String> environment = Map.of(
                "AZURE_STORAGE_ENDPOINT", "http://example.blob.core.windows.net",
                "AZURE_STORAGE_CONTAINER", "documents");

        assertThrows(IllegalArgumentException.class,
                () -> AzureBlobConfiguration.Settings.from(environment));
    }
}
