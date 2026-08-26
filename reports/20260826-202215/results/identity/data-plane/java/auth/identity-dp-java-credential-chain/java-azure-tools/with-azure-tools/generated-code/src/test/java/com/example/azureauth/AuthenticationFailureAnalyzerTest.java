package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

class AuthenticationFailureAnalyzerTest {
    @Test
    void identifiesExpiredSecret() {
        String reason = AuthenticationFailureAnalyzer.describe(
            new ClientAuthenticationException("AADSTS7000222: client secret keys for app have expired", null)
        );

        assertTrue(reason.startsWith("The client secret has expired."));
    }

    @Test
    void identifiesWrongTenant() {
        String reason = AuthenticationFailureAnalyzer.describe(
            new ClientAuthenticationException("AADSTS90002: Tenant 'bad-id' not found", null)
        );

        assertTrue(reason.startsWith("The tenant ID is wrong"));
    }

    @Test
    void identifiesExpiredCertificate() {
        String reason = AuthenticationFailureAnalyzer.describe(
            new ClientAuthenticationException("Client certificate has expired", null)
        );

        assertTrue(reason.startsWith("The client certificate is expired"));
    }

    @Test
    void identifiesUnavailableChain() {
        String reason = AuthenticationFailureAnalyzer.describe(
            new CredentialUnavailableException("Azure CLI executable not found")
        );

        assertTrue(reason.startsWith("No credential in the selected chain"));
    }
}
