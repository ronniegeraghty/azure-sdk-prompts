package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class SyncSecretCache {
    private final SyncSecretProvider provider;
    private final Map<String, String> defaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, SecretValue> cache = new ConcurrentHashMap<>();

    public SyncSecretCache(
            SyncSecretProvider provider,
            Map<String, String> defaults,
            Duration warningWindow) {
        this(provider, defaults, warningWindow, Clock.systemUTC());
    }

    SyncSecretCache(
            SyncSecretProvider provider,
            Map<String, String> defaults,
            Duration warningWindow,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.defaults = Map.copyOf(defaults);
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public Map<String, String> loadRequired(Collection<String> names) {
        names.forEach(this::refresh);
        return snapshot();
    }

    public String get(String name) {
        SecretValue current = cache.get(name);
        if (current == null || current.expiresWithin(warningWindow, clock)) {
            current = refresh(name);
        }
        return current.value();
    }

    public SecretValue refresh(String name) {
        SecretValue refreshed = provider.getSecret(name, defaultFor(name));
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<SecretValue> secretsNearExpiry() {
        return cache.values().stream()
                .filter(SecretValue::found)
                .filter(secret -> secret.expiresWithin(warningWindow, clock))
                .toList();
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
