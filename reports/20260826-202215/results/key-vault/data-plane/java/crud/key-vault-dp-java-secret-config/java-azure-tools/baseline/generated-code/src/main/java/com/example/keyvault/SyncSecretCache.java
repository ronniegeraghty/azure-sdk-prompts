package com.example.keyvault;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class SyncSecretCache {
    private final SyncSecretProvider provider;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, String> defaults;
    private final ConcurrentHashMap<String, SecretSnapshot> cache = new ConcurrentHashMap<>();

    public SyncSecretCache(
            SyncSecretProvider provider,
            Duration warningWindow,
            Map<String, String> defaults) {
        this(provider, warningWindow, defaults, Clock.systemUTC());
    }

    SyncSecretCache(
            SyncSecretProvider provider,
            Duration warningWindow,
            Map<String, String> defaults,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.warningWindow = requireNonNegative(warningWindow);
        this.defaults = Map.copyOf(defaults);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public void loadRequired(Collection<String> names) {
        names.forEach(this::refresh);
    }

    public String get(String name) {
        SecretSnapshot current = cache.computeIfAbsent(name, this::fetch);
        if (isNearExpiry(current)) {
            current = refresh(name);
        }
        return current.value();
    }

    public SecretSnapshot refresh(String name) {
        SecretSnapshot refreshed = fetch(name);
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<SecretSnapshot> expiringSecrets() {
        return cache.values().stream()
                .filter(this::isNearExpiry)
                .toList();
    }

    private SecretSnapshot fetch(String name) {
        return provider.get(name, defaults.getOrDefault(name, ""));
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
