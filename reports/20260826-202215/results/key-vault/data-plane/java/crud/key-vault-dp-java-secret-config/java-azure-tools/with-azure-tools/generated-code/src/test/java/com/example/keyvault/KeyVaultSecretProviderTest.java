package com.example.keyvault;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import java.time.OffsetDateTime;
import org.junit.jupiter.api.Test;
import reactor.core.publisher.Mono;

class KeyVaultSecretProviderTest {
    @Test
    void syncProviderGetsSpecificVersionAndExpiry() {
        SecretClient client = mock(SecretClient.class);
        KeyVaultSecret secret = new KeyVaultSecret("api-key", "value");
        OffsetDateTime expiry = OffsetDateTime.parse("2026-09-01T00:00:00Z");
        secret.getProperties().setExpiresOn(expiry);
        when(client.getSecret("api-key", "v1")).thenReturn(secret);

        SecretSnapshot result =
            new SyncKeyVaultSecretProvider(client).getSecret("api-key", "v1", "fallback");

        assertEquals("value", result.value());
        assertEquals(expiry, result.expiresOn());
        verify(client).getSecret("api-key", "v1");
    }

    @Test
    void syncProviderReturnsDefaultWhenMissing() {
        SecretClient client = mock(SecretClient.class);
        when(client.getSecret("missing")).thenThrow(mock(ResourceNotFoundException.class));

        SecretSnapshot result =
            new SyncKeyVaultSecretProvider(client).getSecret("missing", "fallback");

        assertEquals("fallback", result.value());
        assertTrue(result.defaultValue());
    }

    @Test
    void asyncProviderReturnsDefaultWhenMissing() {
        SecretAsyncClient client = mock(SecretAsyncClient.class);
        when(client.getSecret("missing"))
            .thenReturn(Mono.error(mock(ResourceNotFoundException.class)));

        SecretSnapshot result =
            new AsyncKeyVaultSecretProvider(client).getSecret("missing", "fallback").block();

        assertEquals("fallback", result.value());
        assertTrue(result.defaultValue());
    }
}
