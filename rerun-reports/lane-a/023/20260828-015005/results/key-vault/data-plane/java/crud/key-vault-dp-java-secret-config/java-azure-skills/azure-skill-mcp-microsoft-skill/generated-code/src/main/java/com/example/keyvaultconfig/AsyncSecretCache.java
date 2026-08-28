package com.example.keyvaultconfig;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncSecretCache {
    private final AsyncSecretProvider provider;
    private final Map<String, String> defaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, SecretValue> cache = new ConcurrentHashMap<>();

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
        this.defaults = Map.copyOf(defaults);
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Map<String, String>> loadRequired(Collection<String> names) {
        return Flux.fromIterable(names)
                .flatMap(this::refresh)
                .then(Mono.fromSupplier(this::snapshot));
    }

    public Mono<String> get(String name) {
        return Mono.defer(() -> {
            SecretValue current = cache.get(name);
            if (current == null || current.expiresWithin(warningWindow, clock)) {
                return refresh(name).map(SecretValue::value);
            }
            return Mono.just(current.value());
        });
    }

    public Mono<SecretValue> refresh(String name) {
        return provider.getSecret(name, defaultFor(name))
                .doOnNext(secret -> cache.put(name, secret));
    }

    public Mono<List<SecretValue>> secretsNearExpiry() {
        return Mono.fromSupplier(() -> cache.values().stream()
                .filter(SecretValue::found)
                .filter(secret -> secret.expiresWithin(warningWindow, clock))
                .toList());
    }

    public Map<String, String> snapshot() {
        Map<String, String> values = new ConcurrentHashMap<>();
        cache.forEach((name, secret) -> values.put(name, secret.value()));
        return Map.copyOf(values);
    }

    private String defaultFor(String name) {
        Objects.requireNonNull(name, "name");
        return defaults.getOrDefault(name, "");
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
