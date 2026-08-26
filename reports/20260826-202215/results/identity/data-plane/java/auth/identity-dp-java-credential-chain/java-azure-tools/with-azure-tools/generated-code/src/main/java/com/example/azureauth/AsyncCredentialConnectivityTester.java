package com.example.azureauth;

import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncCredentialConnectivityTester {
    public Mono<ConnectivityTestResult> test(CredentialSelection selection, String scope) {
        Objects.requireNonNull(selection, "selection");
        TokenRequestContext request = CredentialConnectivityTester.request(scope, selection.caeEnabled());

        return selection.credential().getToken(request)
            .map(token -> ConnectivityTestResult.success(token.getExpiresAt(), request.isCaeEnabled()))
            .onErrorResume(failure -> Mono.just(ConnectivityTestResult.failure(
                request.isCaeEnabled(),
                AuthenticationFailureAnalyzer.describe(failure)
            )))
            .doOnNext(CredentialConnectivityTester::print);
    }
}
