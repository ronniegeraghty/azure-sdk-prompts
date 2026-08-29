package com.example.keyvaultconfig;

import reactor.core.publisher.Mono;

public interface AsyncSecretProvider {
    Mono<SecretValue> getSecret(String name, String defaultValue);

    Mono<SecretValue> getSecret(String name, String version, String defaultValue);
}
