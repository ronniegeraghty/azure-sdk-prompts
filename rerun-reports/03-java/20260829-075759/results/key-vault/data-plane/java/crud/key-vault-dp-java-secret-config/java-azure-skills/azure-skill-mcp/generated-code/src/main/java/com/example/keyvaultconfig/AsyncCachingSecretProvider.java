package com.example.keyvaultconfig;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

public final class AsyncCachingSecretProvider {
    private final AsyncSecretProvider provider;
    private final Duration expiryWarningWindow;
    private final Clock clock;
    private final ConcurrentMap<String, SecretSnapshot> cache = new ConcurrentHashMap<>();
    private final ConcurrentMap<String, String> defaults = new ConcurrentHashMap<>();

    public AsyncCachingSecretProvider(AsyncSecretProvider provider, Duration expiryWarningWindow) {
        this(provider, expiryWarningWindow, Clock.systemUTC());
    }

    AsyncCachingSecretProvider(
            AsyncSecretProvider provider,
            Duration expiryWarningWindow,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.expiryWarningWindow = requireNonNegative(expiryWarningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Void> loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        requiredSecrets.forEach(AsyncCachingSecretProvider::validate);
        defaults.putAll(requiredSecrets);
        return Flux.fromIterable(requiredSecrets.entrySet())
                .flatMap(entry -> fetchAndCache(entry.getKey(), entry.getValue()))
                .then();
    }

    public Mono<SecretSnapshot> get(String name) {
        return Mono.defer(() -> {
            SecretSnapshot cached = cache.get(name);
            if (cached == null) {
                return Mono.error(new IllegalArgumentException("Secret is not loaded: " + name));
            }
            return cached.expiresWithin(expiryWarningWindow, clock)
                    ? refresh(name)
                    : Mono.just(cached);
        });
    }

    public Mono<SecretSnapshot> refresh(String name) {
        return Mono.defer(() -> {
            String defaultValue = defaults.get(name);
            if (defaultValue == null) {
                return Mono.error(new IllegalArgumentException("Secret is not loaded: " + name));
            }
            return fetchAndCache(name, defaultValue);
        });
    }

    public Flux<SecretSnapshot> refreshExpiringSecrets() {
        return Flux.defer(() -> Flux.fromIterable(cache.values())
                .filter(secret -> secret.expiresWithin(expiryWarningWindow, clock))
                .flatMap(secret -> refresh(secret.name())));
    }

    public List<SecretSnapshot> expiringSecrets() {
        return cache.values().stream()
                .filter(secret -> secret.expiresWithin(expiryWarningWindow, clock))
                .toList();
    }

    private Mono<SecretSnapshot> fetchAndCache(String name, String defaultValue) {
        return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "expiryWarningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("expiryWarningWindow must not be negative");
        }
        return duration;
    }

    private static void validate(String name, String defaultValue) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("Secret name must not be blank");
        }
        Objects.requireNonNull(defaultValue, "defaultValue");
    }
}
