package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class CachingSecretProvider implements AutoCloseable {
    private final SecretProvider provider;
    private final Duration expiryWarningWindow;
    private final Clock clock;
    private final Consumer<Throwable> refreshErrorHandler;
    private final Map<String, String> defaultValues = new ConcurrentHashMap<>();
    private final Map<String, SecretValue> cache = new ConcurrentHashMap<>();
    private final ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(runnable -> {
        Thread thread = new Thread(runnable, "key-vault-cache-refresh");
        thread.setDaemon(true);
        return thread;
    });

    public CachingSecretProvider(
            SecretProvider provider,
            Duration expiryWarningWindow,
            Consumer<Throwable> refreshErrorHandler) {
        this(provider, expiryWarningWindow, Clock.systemUTC(), refreshErrorHandler);
    }

    CachingSecretProvider(
            SecretProvider provider,
            Duration expiryWarningWindow,
            Clock clock,
            Consumer<Throwable> refreshErrorHandler) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.expiryWarningWindow = requirePositive(expiryWarningWindow, "expiryWarningWindow");
        this.clock = Objects.requireNonNull(clock, "clock");
        this.refreshErrorHandler = Objects.requireNonNull(refreshErrorHandler, "refreshErrorHandler");
    }

    public void loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        requiredSecrets.forEach(this::get);
    }

    public SecretValue get(String name, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        defaultValues.putIfAbsent(name, defaultValue);
        return cache.computeIfAbsent(name, key -> provider.getSecret(key, defaultValue));
    }

    public Optional<SecretValue> getCached(String name) {
        return Optional.ofNullable(cache.get(Objects.requireNonNull(name, "name")));
    }

    public SecretValue refresh(String name) {
        Objects.requireNonNull(name, "name");
        String defaultValue = defaultValues.get(name);
        if (defaultValue == null) {
            throw new IllegalArgumentException("No default value registered for secret: " + name);
        }
        SecretValue refreshed = provider.getSecret(name, defaultValue);
        cache.put(name, refreshed);
        return refreshed;
    }

    public List<SecretValue> secretsNearExpiry() {
        OffsetDateTime threshold = OffsetDateTime.now(clock).plus(expiryWarningWindow);
        return cache.values().stream()
                .filter(secret -> secret.expiry().map(expiry -> !expiry.isAfter(threshold)).orElse(false))
                .toList();
    }

    public void refreshExpiringSecrets() {
        new ArrayList<>(secretsNearExpiry()).forEach(secret -> refresh(secret.name()));
    }

    public void startAutomaticRefresh(Duration checkInterval) {
        Duration interval = requirePositive(checkInterval, "checkInterval");
        scheduler.scheduleWithFixedDelay(
                this::refreshExpiringSecretsSafely,
                interval.toMillis(),
                interval.toMillis(),
                TimeUnit.MILLISECONDS);
    }

    private void refreshExpiringSecretsSafely() {
        try {
            refreshExpiringSecrets();
        } catch (RuntimeException exception) {
            refreshErrorHandler.accept(exception);
        }
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }

    @Override
    public void close() {
        scheduler.shutdownNow();
    }
}
