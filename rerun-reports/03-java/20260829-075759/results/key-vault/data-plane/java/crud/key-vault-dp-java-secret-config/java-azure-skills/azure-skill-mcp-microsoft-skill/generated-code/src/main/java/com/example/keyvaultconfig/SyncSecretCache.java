package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Collection;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class SyncSecretCache {
    private final SyncSecretProvider provider;
    private final Map<String, String> defaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final ConcurrentHashMap<String, SecretValue> cache = new ConcurrentHashMap<>();

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
        this.defaults = Map.copyOf(Objects.requireNonNull(defaults, "defaults"));
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public void loadRequired(Collection<String> names) {
        Objects.requireNonNull(names, "names").forEach(this::refresh);
    }

    public String get(String name) {
        SecretValue cached = cache.computeIfAbsent(name, this::fetch);
        if (isNearExpiry(cached)) {
            cached = refresh(name);
        }
        return cached.value();
    }

    public SecretValue refresh(String name) {
        SecretValue refreshed = fetch(name);
        cache.put(name, refreshed);
        return refreshed;
    }

    public Map<String, SecretValue> refreshExpiring() {
        cache.forEach((name, secret) -> {
            if (isNearExpiry(secret)) {
                refresh(name);
            }
        });
        return snapshot();
    }

    public Map<String, SecretValue> snapshot() {
        return Map.copyOf(cache);
    }

    public boolean isNearExpiry(SecretValue secret) {
        return secret.expiresWithin(warningWindow, OffsetDateTime.now(clock));
    }

    private SecretValue fetch(String name) {
        String defaultValue = defaults.get(name);
        if (defaultValue == null) {
            throw new IllegalArgumentException("No default configured for secret: " + name);
        }
        return provider.getSecret(name, defaultValue);
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
