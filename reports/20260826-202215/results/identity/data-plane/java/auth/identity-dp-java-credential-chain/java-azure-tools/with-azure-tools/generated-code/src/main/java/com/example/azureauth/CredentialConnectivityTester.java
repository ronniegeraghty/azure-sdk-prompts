package com.example.azureauth;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;

import java.util.Objects;

public final class CredentialConnectivityTester {
    public ConnectivityTestResult test(CredentialSelection selection, String scope) {
        Objects.requireNonNull(selection, "selection");
        TokenRequestContext request = request(scope, selection.caeEnabled());

        try {
            AccessToken token = selection.credential().getTokenSync(request);
            ConnectivityTestResult result =
                ConnectivityTestResult.success(token.getExpiresAt(), request.isCaeEnabled());
            print(result);
            return result;
        } catch (RuntimeException failure) {
            ConnectivityTestResult result = ConnectivityTestResult.failure(
                request.isCaeEnabled(),
                AuthenticationFailureAnalyzer.describe(failure)
            );
            print(result);
            return result;
        }
    }

    static TokenRequestContext request(String scope, boolean enableCae) {
        if (scope == null || scope.isBlank()) {
            throw new IllegalArgumentException("scope must not be blank");
        }
        return new TokenRequestContext().addScopes(scope).setCaeEnabled(enableCae);
    }

    static void print(ConnectivityTestResult result) {
        if (result.successful()) {
            System.out.printf(
                "SUCCESS - token expires at %s; CAE requested: %s%n",
                result.expiresAt(),
                result.caeRequested()
            );
        } else {
            System.out.printf(
                "FAILURE - CAE requested: %s; reason: %s%n",
                result.caeRequested(),
                result.failureReason()
            );
        }
    }
}
