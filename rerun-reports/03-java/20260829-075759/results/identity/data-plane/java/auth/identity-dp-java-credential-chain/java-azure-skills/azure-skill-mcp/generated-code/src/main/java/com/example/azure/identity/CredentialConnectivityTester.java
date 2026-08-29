package com.example.azure.identity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;

import java.time.Duration;

public final class CredentialConnectivityTester {
    private static final Duration TOKEN_REQUEST_TIMEOUT = Duration.ofSeconds(30);

    public boolean test(CredentialFactory.BuiltCredential builtCredential, String scope) {
        TokenRequestContext request = new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(builtCredential.caeEnabled());

        System.out.println("[sync] Requesting a token...");
        try {
            AccessToken token = builtCredential.credential().getToken(request).block(TOKEN_REQUEST_TIMEOUT);
            if (token == null) {
                throw new IllegalStateException("The credential completed without returning a token.");
            }
            printSuccess("sync", token, builtCredential.caeEnabled());
            return true;
        } catch (RuntimeException failure) {
            printFailure("sync", failure, builtCredential.caeEnabled());
            return false;
        }
    }

    static void printSuccess(String mode, AccessToken token, boolean caeEnabled) {
        System.out.printf("[%s] SUCCESS%n", mode);
        System.out.printf("[%s] Token expires: %s%n", mode, token.getExpiresAt());
        printCaeStatus(mode, caeEnabled);
    }

    static void printFailure(String mode, Throwable failure, boolean caeEnabled) {
        System.out.printf("[%s] FAILURE%n", mode);
        System.out.printf("[%s] Reason: %s%n", mode, AuthenticationFailureAnalyzer.describe(failure));
        printCaeStatus(mode, caeEnabled);
    }

    private static void printCaeStatus(String mode, boolean caeEnabled) {
        System.out.printf(
            "[%s] CAE-enabled token request: %s"
                + " (CAE issuance is controlled by Microsoft Entra ID and the target resource)%n",
            mode,
            caeEnabled ? "yes" : "no"
        );
    }
}
