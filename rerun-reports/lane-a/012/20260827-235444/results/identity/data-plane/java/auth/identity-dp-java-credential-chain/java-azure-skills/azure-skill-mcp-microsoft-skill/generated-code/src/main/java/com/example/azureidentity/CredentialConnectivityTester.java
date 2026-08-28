package com.example.azureidentity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import com.azure.core.exception.ClientAuthenticationException;

import java.util.Objects;

public final class CredentialConnectivityTester {
    public boolean test(TokenCredential credential, String scope, boolean enableCae) {
        Objects.requireNonNull(credential, "credential");
        TokenRequestContext request = request(scope, enableCae);

        try {
            AccessToken token = credential.getTokenSync(request);
            printSuccess(token, request.isCaeEnabled());
            return true;
        } catch (ClientAuthenticationException exception) {
            printFailure(exception, request.isCaeEnabled());
            return false;
        } catch (RuntimeException exception) {
            printFailure(exception, request.isCaeEnabled());
            return false;
        }
    }

    static TokenRequestContext request(String scope, boolean enableCae) {
        if (scope == null || scope.isBlank()) {
            throw new IllegalArgumentException("scope must not be blank");
        }
        return new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(enableCae);
    }

    static void printSuccess(AccessToken token, boolean caeRequested) {
        System.out.println("  Result: SUCCESS");
        System.out.println("  Token expires: " + token.getExpiresAt());
        printCae(caeRequested);
    }

    static void printFailure(Throwable failure, boolean caeRequested) {
        System.out.println("  Result: FAILURE");
        System.out.println("  Reason: " + AuthenticationFailureReporter.describe(failure));
        printCae(caeRequested);
    }

    private static void printCae(boolean caeRequested) {
        System.out.println("  CAE requested: " + caeRequested
            + " (the opaque access token does not expose whether the resource granted a CAE token)");
    }
}
