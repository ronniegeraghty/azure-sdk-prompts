package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class SecretCache {
    private final KeyVaultSecretProvider provider;
    private final Map<String, String> requiredDefaults;
    private final Duration warningWindow;
    private final Clock clock;
    private final ConcurrentHashMap<String, ConfigSecret> cache = new ConcurrentHashMap<>();

    public SecretCache(
            KeyVaultSecretProvider provider,
            Map<String, String> requiredDefaults,
            Duration warningWindow) {
        this(provider, requiredDefaults, warningWindow, Clock.systemUTC());
    }

    SecretCache(
            KeyVaultSecretProvider provider,
            Map<String, String> requiredDefaults,
            Duration warningWindow,
            Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.requiredDefaults = Map.copyOf(Objects.requireNonNull(requiredDefaults, "requiredDefaults"));
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public void loadRequired() {
        requiredDefaults.forEach((name, defaultValue) ->
                cache.put(name, provider.getSecret(name, defaultValue)));
    }

    public ConfigSecret get(String name) {
        String defaultValue = defaultFor(name);
        return cache.compute(name, (key, current) -> current == null || isNearExpiry(current)
                ? provider.getSecret(key, defaultValue)
                : current);
    }

    public ConfigSecret refresh(String name) {
        ConfigSecret refreshed = provider.getSecret(name, defaultFor(name));
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<ConfigSecret> expiringSecrets() {
        return cache.values().stream().filter(this::isNearExpiry).toList();
    }

    public List<ConfigSecret> refreshExpiringSecrets() {
        List<ConfigSecret> refreshed = new ArrayList<>();
        for (ConfigSecret secret : expiringSecrets()) {
            refreshed.add(refresh(secret.name()));
        }
        return List.copyOf(refreshed);
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
