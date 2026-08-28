package com.example.azureidentity;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

class AuthenticationFailureReporterTest {
    @Test
    void reportsExpiredSecret() {
        String result = AuthenticationFailureReporter.describe(
            new ClientAuthenticationException("AADSTS7000222: client secret has expired", null)
        );
        assertTrue(result.startsWith("The service principal client secret has expired."));
    }

    @Test
    void reportsWrongTenant() {
        String result = AuthenticationFailureReporter.describe(
            new ClientAuthenticationException("AADSTS90002: Tenant not found", null)
        );
        assertTrue(result.startsWith("The configured tenant does not exist"));
    }

    @Test
    void reportsUnavailableIdentity() {
        String result = AuthenticationFailureReporter.describe(
            new CredentialUnavailableException("Managed Identity endpoint is unavailable")
        );
        assertTrue(result.startsWith("No identity is available"));
    }
}
