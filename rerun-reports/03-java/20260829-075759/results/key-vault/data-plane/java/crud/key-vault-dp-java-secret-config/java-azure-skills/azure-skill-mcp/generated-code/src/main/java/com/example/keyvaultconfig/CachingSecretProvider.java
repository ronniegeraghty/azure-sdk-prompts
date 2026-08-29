package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

public final class CachingSecretProvider {
    private final SecretProvider provider;
    private final Duration expiryWarningWindow;
    private final Clock clock;
    private final ConcurrentMap<String, SecretSnapshot> cache = new ConcurrentHashMap<>();
    private final ConcurrentMap<String, String> defaults = new ConcurrentHashMap<>();

    public CachingSecretProvider(SecretProvider provider, Duration expiryWarningWindow) {
        this(provider, expiryWarningWindow, Clock.systemUTC());
    }

    CachingSecretProvider(SecretProvider provider, Duration expiryWarningWindow, Clock clock) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.expiryWarningWindow = requireNonNegative(expiryWarningWindow);
        this.clock = Objects.requireNonNull(clock, "clock");
    }

    public void loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        requiredSecrets.forEach((name, defaultValue) -> {
            validate(name, defaultValue);
            defaults.put(name, defaultValue);
            cache.put(name, provider.getSecret(name, defaultValue));
        });
    }

    public SecretSnapshot get(String name) {
        SecretSnapshot cached = requireCached(name);
        return cached.expiresWithin(expiryWarningWindow, clock) ? refresh(name) : cached;
    }

    public SecretSnapshot refresh(String name) {
        String defaultValue = requireDefault(name);
        SecretSnapshot refreshed = provider.getSecret(name, defaultValue);
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<SecretSnapshot> refreshExpiringSecrets() {
        List<SecretSnapshot> refreshed = new ArrayList<>();
        cache.forEach((name, secret) -> {
            if (secret.expiresWithin(expiryWarningWindow, clock)) {
                refreshed.add(refresh(name));
            }
        });
        return List.copyOf(refreshed);
    }

    public List<SecretSnapshot> expiringSecrets() {
        return cache.values().stream()
                .filter(secret -> secret.expiresWithin(expiryWarningWindow, clock))
                .toList();
    }

    private SecretSnapshot requireCached(String name) {
        SecretSnapshot secret = cache.get(name);
        if (secret == null) {
            throw new IllegalArgumentException("Secret is not loaded: " + name);
        }
        return secret;
    }

    private String requireDefault(String name) {
        String defaultValue = defaults.get(name);
        if (defaultValue == null) {
            throw new IllegalArgumentException("Secret is not loaded: " + name);
        }
        return defaultValue;
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
