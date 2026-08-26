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

class CredentialConnectivityTesterTest {
    private static final String SCOPE = "https://management.azure.com/.default";

    @Test
    void syncTesterAppliesCaeAndReportsExpiry() {
        AtomicReference<TokenRequestContext> capturedRequest = new AtomicReference<>();
        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);
        CredentialSelection selection = selection(capturedRequest, expiry, true);

        ConnectivityTestResult result = new CredentialConnectivityTester().test(selection, SCOPE);

        assertTrue(result.successful());
        assertTrue(result.caeRequested());
        assertEquals(expiry, result.expiresAt());
        assertEquals(SCOPE, capturedRequest.get().getScopes().get(0));
        assertTrue(capturedRequest.get().isCaeEnabled());
    }

    @Test
    void asyncTesterAppliesCaeAndReportsExpiry() {
        AtomicReference<TokenRequestContext> capturedRequest = new AtomicReference<>();
        OffsetDateTime expiry = OffsetDateTime.now().plusHours(1);
        CredentialSelection selection = selection(capturedRequest, expiry, true);

        ConnectivityTestResult result =
            new AsyncCredentialConnectivityTester().test(selection, SCOPE).block();

        assertTrue(result.successful());
        assertTrue(result.caeRequested());
        assertEquals(expiry, result.expiresAt());
        assertTrue(capturedRequest.get().isCaeEnabled());
    }

    private static CredentialSelection selection(
        AtomicReference<TokenRequestContext> capturedRequest,
        OffsetDateTime expiry,
        boolean caeEnabled
    ) {
        TokenCredential credential = request -> {
            capturedRequest.set(request);
            return Mono.just(new AccessToken("fake-token", expiry));
        };
        return new CredentialSelection(credential, "test", caeEnabled);
    }
}
