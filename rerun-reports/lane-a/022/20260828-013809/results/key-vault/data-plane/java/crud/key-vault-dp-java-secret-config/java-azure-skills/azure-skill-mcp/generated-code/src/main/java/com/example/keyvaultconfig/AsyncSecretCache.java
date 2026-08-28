package com.example.keyvaultconfig;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncSecretCache {
    private final AsyncKeyVaultSecretProvider provider;
    private final Map<String, String> requiredDefaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final ConcurrentHashMap<String, ConfigSecret> cache = new ConcurrentHashMap<>();

    public AsyncSecretCache(
            AsyncKeyVaultSecretProvider provider,
            Map<String, String> requiredDefaults,
            Duration warningWindow) {
        this(provider, requiredDefaults, warningWindow, Clock.systemUTC());
    }

    AsyncSecretCache(
            AsyncKeyVaultSecretProvider provider,
            Map<String, String> requiredDefaults,
            Duration warningWindow,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.requiredDefaults = Map.copyOf(Objects.requireNonNull(requiredDefaults, "requiredDefaults"));
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Mono<Void> loadRequired() {
        return Flux.fromIterable(requiredDefaults.entrySet())
                .flatMap(entry -> fetchAndCache(entry.getKey(), entry.getValue()))
                .then();
    }

    public Mono<ConfigSecret> get(String name) {
        String defaultValue = defaultFor(name);
        ConfigSecret current = cache.get(name);
        return current == null || isNearExpiry(current)
                ? fetchAndCache(name, defaultValue)
                : Mono.just(current);
    }

    public Mono<ConfigSecret> refresh(String name) {
        return fetchAndCache(name, defaultFor(name));
    }

    public List<ConfigSecret> expiringSecrets() {
        return cache.values().stream().filter(this::isNearExpiry).toList();
    }

    public Flux<ConfigSecret> refreshExpiringSecrets() {
        return Flux.fromIterable(expiringSecrets())
                .flatMap(secret -> refresh(secret.name()));
    }

    private Mono<ConfigSecret> fetchAndCache(String name, String defaultValue) {
        return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
    }

    private String defaultFor(String name) {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("name must not be blank");
        }
        return requiredDefaults.getOrDefault(name, "");
    }

    private boolean isNearExpiry(ConfigSecret secret) {
        return secret.expiresWithin(warningWindow, clock);
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
