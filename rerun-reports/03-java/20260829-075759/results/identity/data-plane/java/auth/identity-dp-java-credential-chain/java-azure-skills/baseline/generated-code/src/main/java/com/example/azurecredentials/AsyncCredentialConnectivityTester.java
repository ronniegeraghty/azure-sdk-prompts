package com.example.azurecredentials;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

public final class AsyncCredentialConnectivityTester {
    public Mono<Boolean> test(CredentialSelection selection, String scope) {
        TokenRequestContext request =
            CredentialConnectivityTester.requestFor(scope, selection.caeEnabled());

        return selection.credential().getToken(request)
            .doOnNext(token -> printSuccess(token, selection.caeEnabled()))
            .map(token -> true)
            .onErrorResume(exception -> {
                System.out.println("[async] Authentication failed: "
                    + AuthenticationFailureReporter.describe(exception));
                System.out.println("[async] CAE requested: " + selection.caeEnabled());
                return Mono.just(false);
            });
    }

    private static void printSuccess(AccessToken token, boolean caeRequested) {
        System.out.println("[async] Authentication succeeded");
        System.out.println("[async] Token expires at: " + token.getExpiresAt());
        System.out.println("[async] CAE: " + CaeTokenInspector.status(token, caeRequested));
    }
}
