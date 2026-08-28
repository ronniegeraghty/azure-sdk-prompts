package com.example.azureidentity;

import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncCredentialConnectivityTester {
    public Mono<Boolean> test(TokenCredential credential, String scope, boolean enableCae) {
        Objects.requireNonNull(credential, "credential");
        TokenRequestContext request = CredentialConnectivityTester.request(scope, enableCae);

        return credential.getToken(request)
            .doOnNext(token -> CredentialConnectivityTester.printSuccess(token, request.isCaeEnabled()))
            .map(token -> true)
            .onErrorResume(failure -> {
                CredentialConnectivityTester.printFailure(failure, request.isCaeEnabled());
                return Mono.just(false);
            });
    }
}
