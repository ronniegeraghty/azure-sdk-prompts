package com.example.azureauth;

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
        request.setCaeEnabled(true);
        return delegate.getToken(request);
    }
}
