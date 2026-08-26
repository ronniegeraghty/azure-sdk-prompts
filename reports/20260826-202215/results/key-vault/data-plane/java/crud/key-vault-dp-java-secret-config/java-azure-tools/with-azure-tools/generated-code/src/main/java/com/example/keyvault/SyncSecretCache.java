package com.example.keyvault;

import java.time.Clock;
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

public final class SyncSecretCache {
    private final SyncKeyVaultSecretProvider provider;
    private final Duration warningWindow;
    private final Clock clock;
    private final Map<String, String> defaults = new ConcurrentHashMap<>();
    private final Map<String, SecretSnapshot> cache = new ConcurrentHashMap<>();

    public SyncSecretCache(SyncKeyVaultSecretProvider provider, Duration warningWindow) {
        this(provider, warningWindow, Clock.systemUTC());
    }

    SyncSecretCache(SyncKeyVaultSecretProvider provider, Duration warningWindow, Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.warningWindow = requireNonNegative(warningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public void loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        requiredSecrets.forEach((name, defaultValue) -> {
            defaults.put(name, Objects.requireNonNull(defaultValue, "defaultValue"));
            refresh(name);
        });
    }

    public SecretSnapshot get(String name) {
        SecretSnapshot secret = requireCached(name);
        return SecretExpiry.isWithin(secret, warningWindow, clock) ? refresh(name) : secret;
    }

    public SecretSnapshot refresh(String name) {
        String defaultValue = defaults.get(name);
        if (defaultValue == null) {
            throw new IllegalArgumentException("No default is registered for secret: " + name);
        }
        SecretSnapshot refreshed = provider.getSecret(name, defaultValue);
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<SecretSnapshot> refreshExpiring() {
        List<SecretSnapshot> refreshed = new ArrayList<>();
        List.copyOf(cache.values()).stream()
            .filter(secret -> SecretExpiry.isWithin(secret, warningWindow, clock))
            .map(SecretSnapshot::name)
            .map(this::refresh)
            .forEach(refreshed::add);
        return List.copyOf(refreshed);
    }

    public List<SecretSnapshot> secretsNearExpiry() {
        return cache.values().stream()
            .filter(secret -> SecretExpiry.isWithin(secret, warningWindow, clock))
            .toList();
    }

    private SecretSnapshot requireCached(String name) {
        SecretSnapshot secret = cache.get(name);
        if (secret == null) {
            throw new IllegalArgumentException("Secret has not been loaded: " + name);
        }
        return secret;
    }

    private static Duration requireNonNegative(Duration duration) {
        Objects.requireNonNull(duration, "warningWindow");
        if (duration.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return duration;
    }
}
