package com.example.azure.identity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

import java.util.Objects;

final class CaeTokenCredential implements TokenCredential {
    private final TokenCredential delegate;

    CaeTokenCredential(TokenCredential delegate) {
        this.delegate = Objects.requireNonNull(delegate, "delegate");
    }

    @Override
    public Mono<AccessToken> getToken(TokenRequestContext request) {
        TokenRequestContext caeRequest = new TokenRequestContext()
            .setScopes(request.getScopes())
            .setClaims(request.getClaims())
            .setTenantId(request.getTenantId())
            .setProofOfPossessionOptions(request.getProofOfPossessionOptions())
            .setCaeEnabled(true);
        return delegate.getToken(caeRequest);
    }
}
