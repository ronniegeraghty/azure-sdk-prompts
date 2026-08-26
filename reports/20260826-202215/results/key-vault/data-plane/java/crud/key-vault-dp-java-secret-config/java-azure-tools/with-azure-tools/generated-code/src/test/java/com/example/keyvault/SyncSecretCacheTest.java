package com.example.keyvault;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import java.time.Clock;
import java.time.Duration;
import java.time.Instant;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.Map;
import org.junit.jupiter.api.Test;

class SyncSecretCacheTest {
    private static final Clock CLOCK =
        Clock.fixed(Instant.parse("2026-08-26T00:00:00Z"), ZoneOffset.UTC);

    @Test
    void refreshesNearExpirySecretWhenRead() {
        SyncKeyVaultSecretProvider provider = mock(SyncKeyVaultSecretProvider.class);
        SecretSnapshot nearExpiry = new SecretSnapshot(
            "api-key", "old", "v1", OffsetDateTime.now(CLOCK).plusDays(2), false);
        SecretSnapshot refreshed = new SecretSnapshot(
            "api-key", "new", "v2", OffsetDateTime.now(CLOCK).plusDays(30), false);
        when(provider.getSecret("api-key", "fallback")).thenReturn(nearExpiry, refreshed);

        SyncSecretCache cache = new SyncSecretCache(provider, Duration.ofDays(7), CLOCK);
        cache.loadRequired(Map.of("api-key", "fallback"));

        assertEquals("new", cache.get("api-key").value());
        verify(provider, org.mockito.Mockito.times(2)).getSecret("api-key", "fallback");
    }
}
