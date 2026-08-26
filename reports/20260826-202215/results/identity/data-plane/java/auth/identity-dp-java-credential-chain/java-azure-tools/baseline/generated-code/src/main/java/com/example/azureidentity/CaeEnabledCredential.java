package com.example.azureidentity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

import java.util.Objects;

final class CaeEnabledCredential implements TokenCredential {
    private final TokenCredential delegate;

    CaeEnabledCredential(TokenCredential delegate) {
        this.delegate = Objects.requireNonNull(delegate, "delegate");
    }

    @Override
    public Mono<AccessToken> getToken(TokenRequestContext request) {
        return delegate.getToken(withCae(request));
    }

    @Override
    public AccessToken getTokenSync(TokenRequestContext request) {
        return delegate.getTokenSync(withCae(request));
    }

    private static TokenRequestContext withCae(TokenRequestContext request) {
        return Objects.requireNonNull(request, "request").setCaeEnabled(true);
    }
}
