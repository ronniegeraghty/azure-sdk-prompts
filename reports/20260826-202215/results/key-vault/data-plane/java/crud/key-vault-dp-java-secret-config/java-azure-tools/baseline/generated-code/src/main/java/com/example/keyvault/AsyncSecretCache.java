package com.example.keyvault;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncSecretCache {
    private final AsyncSecretProvider provider;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, String> defaults;
    private final ConcurrentHashMap<String, SecretSnapshot> cache = new ConcurrentHashMap<>();

    public AsyncSecretCache(
            AsyncSecretProvider provider,
            Duration warningWindow,
            Map<String, String> defaults) {
        this(provider, warningWindow, defaults, Clock.systemUTC());
    }

    AsyncSecretCache(
            AsyncSecretProvider provider,
            Duration warningWindow,
            Map<String, String> defaults,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.warningWindow = requireNonNegative(warningWindow);
        this.defaults = Map.copyOf(defaults);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Void> loadRequired(Collection<String> names) {
        return Flux.fromIterable(names).flatMap(this::refresh).then();
    }

    public Mono<String> get(String name) {
        return Mono.defer(() -> {
            SecretSnapshot current = cache.get(name);
            if (current == null || isNearExpiry(current)) {
                return refresh(name).map(SecretSnapshot::value);
            }
            return Mono.just(current.value());
        });
    }

    public Mono<SecretSnapshot> refresh(String name) {
        return provider.get(name, defaults.getOrDefault(name, ""))
                .doOnNext(secret -> cache.put(name, secret));
    }

    public List<SecretSnapshot> expiringSecrets() {
        return cache.values().stream()
                .filter(this::isNearExpiry)
                .toList();
    }

    private boolean isNearExpiry(SecretSnapshot secret) {
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
