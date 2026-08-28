package com.example.azureauth;

import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

public final class AsyncCredentialConnectivityTester {
    public Mono<Boolean> test(CredentialSelection selection, String scope) {
        TokenRequestContext request = new TokenRequestContext()
            .addScopes(scope)
            .setCaeEnabled(selection.caeEnabled());

        return selection.credential().getToken(request)
            .map(token -> {
                System.out.println("[async] SUCCESS");
                System.out.println("  Expires: " + token.getExpiresAt());
                System.out.println("  CAE: " + TokenDetails.caeStatus(token, selection.caeEnabled()));
                return true;
            })
            .onErrorResume(failure -> {
                System.out.println("[async] FAILURE");
                System.out.println("  " + AuthenticationFailureReporter.describe(failure));
                System.out.println("  CAE requested: " + selection.caeEnabled());
                return Mono.just(false);
            });
    }
}
