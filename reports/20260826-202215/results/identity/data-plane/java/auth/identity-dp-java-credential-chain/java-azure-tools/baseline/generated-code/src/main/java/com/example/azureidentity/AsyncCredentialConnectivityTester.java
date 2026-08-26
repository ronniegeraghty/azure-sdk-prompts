package com.example.azureidentity;

import com.azure.core.credential.TokenCredential;
import reactor.core.publisher.Mono;

import java.util.Objects;

public final class AsyncCredentialConnectivityTester {
    public Mono<ConnectivityTestResult> test(
        TokenCredential credential, String scope, boolean caeEnabled) {

        Objects.requireNonNull(credential, "credential");
        return credential.getToken(CredentialConnectivityTester.request(scope, caeEnabled))
            .map(token -> ConnectivityTestResult.success(token.getExpiresAt(), caeEnabled))
            .onErrorResume(error ->
                Mono.just(ConnectivityTestResult.failure(caeEnabled, error)))
            .doOnNext(result -> CredentialConnectivityTester.print("Async", result));
    }
}
