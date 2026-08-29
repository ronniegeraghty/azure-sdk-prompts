package com.example.azure.identity;

import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

public final class AsyncCredentialConnectivityTester {
    public Mono<Boolean> test(CredentialFactory.BuiltCredential builtCredential, String scope) {
        TokenRequestContext request = new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(builtCredential.caeEnabled());

        System.out.println("[async] Requesting a token...");
        return builtCredential.credential()
            .getToken(request)
            .doOnNext(token -> CredentialConnectivityTester.printSuccess(
                "async", token, builtCredential.caeEnabled()))
            .map(token -> true)
            .onErrorResume(failure -> {
                CredentialConnectivityTester.printFailure("async", failure, builtCredential.caeEnabled());
                return Mono.just(false);
            });
    }
}
