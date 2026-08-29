package com.example.azurecredentials;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;

import java.util.List;

public final class CredentialConnectivityTester {
    public boolean test(CredentialSelection selection, String scope) {
        TokenRequestContext request = requestFor(scope, selection.caeEnabled());
        try {
            AccessToken token = selection.credential().getTokenSync(request);
            System.out.println("[sync] Authentication succeeded");
            System.out.println("[sync] Token expires at: " + token.getExpiresAt());
            System.out.println("[sync] CAE: "
                + CaeTokenInspector.status(token, selection.caeEnabled()));
            return true;
        } catch (RuntimeException exception) {
            System.out.println("[sync] Authentication failed: "
                + AuthenticationFailureReporter.describe(exception));
            System.out.println("[sync] CAE requested: " + selection.caeEnabled());
            return false;
        }
    }

    static TokenRequestContext requestFor(String scope, boolean caeEnabled) {
        if (scope == null || scope.isBlank()) {
            throw new IllegalArgumentException("scope must not be blank");
        }
        return new TokenRequestContext()
            .setScopes(List.of(scope))
            .setCaeEnabled(caeEnabled);
    }
}
