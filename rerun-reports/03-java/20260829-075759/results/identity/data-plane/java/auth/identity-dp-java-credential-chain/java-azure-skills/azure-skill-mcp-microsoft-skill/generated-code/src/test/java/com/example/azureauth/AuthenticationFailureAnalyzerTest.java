package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

class AuthenticationFailureAnalyzerTest {
    @Test
    void identifiesWrongTenantOrApplication() {
        ClientAuthenticationException failure =
            new ClientAuthenticationException("AADSTS700016: Application with identifier was not found", null);

        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains("client application"));
    }

    @Test
    void identifiesExpiredCredential() {
        ClientAuthenticationException failure =
            new ClientAuthenticationException("AADSTS7000222: The provided client secret has expired", null);

        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains("expired"));
    }

    @Test
    void identifiesUnavailableChain() {
        CredentialUnavailableException failure =
            new CredentialUnavailableException("No credential was able to authenticate");

        assertTrue(AuthenticationFailureAnalyzer.explain(failure).contains("No credential"));
    }
}
