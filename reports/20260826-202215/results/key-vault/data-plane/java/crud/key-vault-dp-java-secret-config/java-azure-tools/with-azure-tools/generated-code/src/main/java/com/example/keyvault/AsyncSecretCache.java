package com.example.keyvault;

import java.time.Clock;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

public final class AsyncSecretCache {
    private final AsyncKeyVaultSecretProvider provider;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, String> defaults = new ConcurrentHashMap<>();
    private final Map<String, SecretSnapshot> cache = new ConcurrentHashMap<>();

    public AsyncSecretCache(AsyncKeyVaultSecretProvider provider, Duration warningWindow) {
        this(provider, warningWindow, Clock.systemUTC());
    }

    AsyncSecretCache(AsyncKeyVaultSecretProvider provider, Duration warningWindow, Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Void> loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        requiredSecrets.forEach((name, defaultValue) ->
            defaults.put(name, Objects.requireNonNull(defaultValue, "defaultValue")));
        return Flux.fromIterable(requiredSecrets.keySet())
            .flatMap(this::refresh)
            .then();
    }

    public Mono<SecretSnapshot> get(String name) {
        return Mono.defer(() -> {
            SecretSnapshot secret = cache.get(name);
            if (secret == null) {
                return Mono.error(new IllegalArgumentException("Secret has not been loaded: " + name));
            }
            return SecretExpiry.isWithin(secret, warningWindow, clock)
                ? refresh(name)
                : Mono.just(secret);
        });
    }

    public Mono<SecretSnapshot> refresh(String name) {
        return Mono.defer(() -> {
            String defaultValue = defaults.get(name);
            if (defaultValue == null) {
                return Mono.error(new IllegalArgumentException(
                    "No default is registered for secret: " + name));
            }
            return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
        });
    }

    public Flux<SecretSnapshot> refreshExpiring() {
        return Flux.defer(() -> Flux.fromIterable(List.copyOf(cache.values())))
            .filter(secret -> SecretExpiry.isWithin(secret, warningWindow, clock))
            .map(SecretSnapshot::name)
            .flatMap(this::refresh);
    }

    public List<SecretSnapshot> secretsNearExpiry() {
        return cache.values().stream()
            .filter(secret -> SecretExpiry.isWithin(secret, warningWindow, clock))
            .toList();
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
