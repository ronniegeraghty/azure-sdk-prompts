package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncCredentialConnectivityTester {
    public Mono<ConnectivityTestResult> test(CredentialSelection selection, String scope) {
        Objects.requireNonNull(selection, "selection");

        return selection.credential()
            .getToken(SyncCredentialConnectivityTester.request(scope, selection.caeEnabled()))
            .map(token -> ConnectivityTestResult.success(token.getExpiresAt(), selection.caeEnabled()))
            .onErrorResume(
                failure -> failure instanceof CredentialUnavailableException
                    || failure instanceof ClientAuthenticationException,
                failure -> Mono.just(ConnectivityTestResult.failure(
                    selection.caeEnabled(),
                    AuthenticationFailureAnalyzer.explain(failure)
                ))
            )
            .doOnNext(SyncCredentialConnectivityTester::print);
    }
}
