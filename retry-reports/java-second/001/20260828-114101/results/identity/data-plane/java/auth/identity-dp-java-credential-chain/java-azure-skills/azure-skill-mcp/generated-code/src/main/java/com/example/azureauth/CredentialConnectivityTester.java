package com.example.azureauth;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;

public final class CredentialConnectivityTester {
    public boolean test(CredentialSelection selection, String scope) {
        TokenRequestContext request = new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(selection.caeEnabled());

        try {
            AccessToken token = selection.credential().getTokenSync(request);
            System.out.println("[sync] SUCCESS");
            System.out.println("  Expires: " + token.getExpiresAt());
            System.out.println("  CAE: " + TokenDetails.caeStatus(token, selection.caeEnabled()));
            return true;
        } catch (RuntimeException failure) {
            System.out.println("[sync] FAILURE");
            System.out.println("  " + AuthenticationFailureReporter.describe(failure));
            System.out.println("  CAE requested: " + selection.caeEnabled());
            return false;
        }
    }
}
