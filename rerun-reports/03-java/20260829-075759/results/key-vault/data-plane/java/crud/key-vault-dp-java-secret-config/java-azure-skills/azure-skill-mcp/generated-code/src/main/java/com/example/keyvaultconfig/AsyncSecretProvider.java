package com.example.keyvaultconfig;

import reactor.core.publisher.Mono;

public interface AsyncSecretProvider {
    Mono<SecretSnapshot> getSecret(String name, String defaultValue);

    Mono<SecretSnapshot> getSecret(String name, String version, String defaultValue);
}
