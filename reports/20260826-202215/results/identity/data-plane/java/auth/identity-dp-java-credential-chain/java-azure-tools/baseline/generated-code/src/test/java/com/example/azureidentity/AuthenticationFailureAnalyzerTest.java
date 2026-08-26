package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

class AuthenticationFailureAnalyzerTest {
    @Test
    void identifiesExpiredCredentials() {
        String reason = AuthenticationFailureAnalyzer.explain(
            new RuntimeException("AADSTS7000222: The provided client secret keys are expired"));

        assertTrue(reason.startsWith("The client certificate or secret has expired."));
    }

    @Test
    void identifiesWrongTenant() {
        String reason = AuthenticationFailureAnalyzer.explain(
            new RuntimeException("AADSTS90002: Tenant 'bad-tenant' not found"));

        assertTrue(reason.startsWith(
            "The configured Microsoft Entra tenant is invalid or unavailable."));
    }

    @Test
    void identifiesMissingIdentity() {
        String reason = AuthenticationFailureAnalyzer.explain(
            new RuntimeException("ManagedIdentityCredential authentication unavailable"));

        assertTrue(reason.startsWith(
            "No configured managed, workload, pipeline, or developer identity was available."));
    }
}
