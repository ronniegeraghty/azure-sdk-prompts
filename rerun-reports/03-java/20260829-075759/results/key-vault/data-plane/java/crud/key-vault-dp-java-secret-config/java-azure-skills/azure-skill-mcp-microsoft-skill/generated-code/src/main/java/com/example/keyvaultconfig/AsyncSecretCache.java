package com.example.keyvaultconfig;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Collection;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncSecretCache {
    private final AsyncSecretProvider provider;
    private final Map<String, String> defaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final ConcurrentHashMap<String, SecretValue> cache = new ConcurrentHashMap<>();

    public AsyncSecretCache(
            AsyncSecretProvider provider,
            Map<String, String> defaults,
            Duration warningWindow) {
        this(provider, defaults, warningWindow, Clock.systemUTC());
    }

    AsyncSecretCache(
            AsyncSecretProvider provider,
            Map<String, String> defaults,
            Duration warningWindow,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.defaults = Map.copyOf(Objects.requireNonNull(defaults, "defaults"));
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Void> loadRequired(Collection<String> names) {
        return Flux.fromIterable(Objects.requireNonNull(names, "names"))
                .flatMap(this::refresh)
                .then();
    }

    public Mono<String> get(String name) {
        return Mono.defer(() -> {
            SecretValue cached = cache.get(name);
            if (cached == null || isNearExpiry(cached)) {
                return refresh(name).map(SecretValue::value);
            }
            return Mono.just(cached.value());
        });
    }

    public Mono<SecretValue> refresh(String name) {
        String defaultValue = defaults.get(name);
        if (defaultValue == null) {
            return Mono.error(new IllegalArgumentException(
                    "No default configured for secret: " + name));
        }
        return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
    }

    public Mono<Map<String, SecretValue>> refreshExpiring() {
        return Flux.fromIterable(cache.entrySet())
                .filter(entry -> isNearExpiry(entry.getValue()))
                .flatMap(entry -> refresh(entry.getKey()))
                .then(Mono.fromSupplier(this::snapshot));
    }

    public Map<String, SecretValue> snapshot() {
        return Map.copyOf(cache);
    }

    public boolean isNearExpiry(SecretValue secret) {
        return secret.expiresWithin(warningWindow, OffsetDateTime.now(clock));
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
