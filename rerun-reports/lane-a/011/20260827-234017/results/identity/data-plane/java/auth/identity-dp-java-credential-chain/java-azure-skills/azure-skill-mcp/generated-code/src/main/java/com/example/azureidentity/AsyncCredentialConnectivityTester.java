package com.example.azureidentity;

import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

public final class AsyncCredentialConnectivityTester {
    public Mono<ConnectivityResult> test(
        TokenCredential credential,
        String scope,
        boolean caeEnabled
    ) {
        return credential.getToken(new TokenRequestContext().addScopes(scope))
            .map(token -> ConnectivityResult.success(scope, token.getExpiresAt(), caeEnabled))
            .onErrorResume(failure -> Mono.just(ConnectivityResult.failure(
                scope,
                caeEnabled,
                AuthenticationFailureAnalyzer.describe(failure)
            )))
            .doOnNext(result -> result.print("Asynchronous"));
    }
}
