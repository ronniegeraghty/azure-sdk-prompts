package com.example.azureidentity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;

public final class SyncCredentialConnectivityTester {
    public ConnectivityResult test(
        TokenCredential credential,
        String scope,
        boolean caeEnabled
    ) {
        try {
            AccessToken token = credential.getTokenSync(new TokenRequestContext().addScopes(scope));
            ConnectivityResult result = ConnectivityResult.success(scope, token.getExpiresAt(), caeEnabled);
            result.print("Synchronous");
            return result;
        } catch (RuntimeException failure) {
            ConnectivityResult result = ConnectivityResult.failure(
                scope,
                caeEnabled,
                AuthenticationFailureAnalyzer.describe(failure)
            );
            result.print("Synchronous");
            return result;
        }
    }
}
