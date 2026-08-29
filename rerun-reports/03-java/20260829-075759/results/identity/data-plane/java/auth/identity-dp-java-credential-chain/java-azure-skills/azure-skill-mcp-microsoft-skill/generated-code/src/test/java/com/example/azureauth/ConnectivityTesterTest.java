package com.example.azureauth;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import org.junit.jupiter.api.Test;
import reactor.core.publisher.Mono;

import java.time.OffsetDateTime;
import java.util.concurrent.atomic.AtomicReference;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class ConnectivityTesterTest {
    @Test
    void syncTesterRequestsScopeAndCae() {
        AtomicReference<TokenRequestContext> captured = new AtomicReference<>();
        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);
        CredentialSelection selection =
            new CredentialSelection(successfulCredential(captured, expiry), "test", true);

        ConnectivityTestResult result =
            new SyncCredentialConnectivityTester().test(selection, "https://management.azure.com/.default");

        assertTrue(result.successful());
        assertEquals(expiry, result.expiresAt());
        assertTrue(result.caeRequested());
        assertTrue(captured.get().isCaeEnabled());
        assertEquals(
            "https://management.azure.com/.default",
            captured.get().getScopes().get(0)
        );
    }

    @Test
    void asyncTesterRequestsScopeAndCae() {
        AtomicReference<TokenRequestContext> captured = new AtomicReference<>();
        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);
        CredentialSelection selection =
            new CredentialSelection(successfulCredential(captured, expiry), "test", true);

        ConnectivityTestResult result = new AsyncCredentialConnectivityTester()
            .test(selection, "https://management.azure.com/.default")
            .block();

        assertTrue(result.successful());
        assertEquals(expiry, result.expiresAt());
        assertTrue(captured.get().isCaeEnabled());
    }

    private static TokenCredential successfulCredential(
        AtomicReference<TokenRequestContext> captured,
        OffsetDateTime expiry
    ) {
        return request -> {
            captured.set(request);
            return Mono.just(new AccessToken("fake-token", expiry));
        };
    }
}
