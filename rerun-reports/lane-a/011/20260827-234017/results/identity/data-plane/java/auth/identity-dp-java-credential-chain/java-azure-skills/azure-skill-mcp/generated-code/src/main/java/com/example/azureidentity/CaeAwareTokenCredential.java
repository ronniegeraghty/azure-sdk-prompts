package com.example.azureidentity;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenCredential;
import com.azure.core.credential.TokenRequestContext;
import reactor.core.publisher.Mono;

import java.util.Objects;

final class CaeAwareTokenCredential implements TokenCredential {
    private final TokenCredential delegate;
    private final boolean caeEnabled;

    CaeAwareTokenCredential(TokenCredential delegate, boolean caeEnabled) {
        this.delegate = Objects.requireNonNull(delegate, "delegate");
        this.caeEnabled = caeEnabled;
    }

    @Override
    public Mono<AccessToken> getToken(TokenRequestContext request) {
        return delegate.getToken(applyCaeSetting(request));
    }

    @Override
    public AccessToken getTokenSync(TokenRequestContext request) {
        return delegate.getTokenSync(applyCaeSetting(request));
    }

    private TokenRequestContext applyCaeSetting(TokenRequestContext request) {
        return Objects.requireNonNull(request, "request").setCaeEnabled(caeEnabled);
    }
}
