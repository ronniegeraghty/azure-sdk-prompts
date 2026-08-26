package com.example.keyvault;

import reactor.core.publisher.Mono;

public interface AsyncSecretProvider {
    Mono<SecretSnapshot> get(String name, String defaultValue);

    Mono<SecretSnapshot> getVersion(String name, String version, String defaultValue);
}
