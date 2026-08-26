package com.example.azureidentity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;

import java.util.Objects;

public final class CredentialConnectivityTester {
    public ConnectivityTestResult test(
        TokenCredential credential, String scope, boolean caeEnabled) {

        Objects.requireNonNull(credential, "credential");
        try {
            AccessToken token = credential.getTokenSync(request(scope, caeEnabled));
            ConnectivityTestResult result =
                ConnectivityTestResult.success(token.getExpiresAt(), caeEnabled);
            print("Sync", result);
            return result;
        } catch (RuntimeException error) {
            ConnectivityTestResult result = ConnectivityTestResult.failure(caeEnabled, error);
            print("Sync", result);
            return result;
        }
    }

    static TokenRequestContext request(String scope, boolean caeEnabled) {
        if (scope == null || scope.isBlank()) {
            throw new IllegalArgumentException("scope must not be blank");
        }
        return new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(caeEnabled);
    }

    static void print(String mode, ConnectivityTestResult result) {
        if (result.successful()) {
            System.out.printf(
                "%s test: SUCCESS%n  Expires at: %s%n  CAE enabled: %s%n",
                mode, result.expiresAt(), result.caeEnabled());
        } else {
            System.out.printf(
                "%s test: FAILURE%n  Reason: %s%n  CAE enabled: %s%n",
                mode, result.failureReason(), result.caeEnabled());
        }
    }
}
