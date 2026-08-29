package com.example.keyvaultconfig;

import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Consumer;

public final class AsyncCachingSecretProvider implements AutoCloseable {
    private final AsyncSecretProvider provider;
    private final Duration expiryWarningWindow;
    private final Clock clock;
    private final Consumer<Throwable> refreshErrorHandler;
    private final Map<String, String> defaultValues = new ConcurrentHashMap<>();
    private final Map<String, SecretValue> cache = new ConcurrentHashMap<>();
    private volatile Disposable automaticRefresh;

    public AsyncCachingSecretProvider(
            AsyncSecretProvider provider,
            Duration expiryWarningWindow,
            Consumer<Throwable> refreshErrorHandler) {
        this(provider, expiryWarningWindow, Clock.systemUTC(), refreshErrorHandler);
    }

    AsyncCachingSecretProvider(
            AsyncSecretProvider provider,
            Duration expiryWarningWindow,
            Clock clock,
            Consumer<Throwable> refreshErrorHandler) {
        this.provider = Objects.requireNonNull(provider, "provider");
        this.expiryWarningWindow = requirePositive(expiryWarningWindow, "expiryWarningWindow");
        this.clock = Objects.requireNonNull(clock, "clock");
        this.refreshErrorHandler = Objects.requireNonNull(refreshErrorHandler, "refreshErrorHandler");
    }

    public Mono<Void> loadRequired(Map<String, String> requiredSecrets) {
        Objects.requireNonNull(requiredSecrets, "requiredSecrets");
        return Flux.fromIterable(requiredSecrets.entrySet())
                .flatMap(entry -> get(entry.getKey(), entry.getValue()))
                .then();
    }

    public Mono<SecretValue> get(String name, String defaultValue) {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(defaultValue, "defaultValue");
        defaultValues.putIfAbsent(name, defaultValue);
        SecretValue cached = cache.get(name);
        if (cached != null) {
            return Mono.just(cached);
        }
        return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
    }

    public Optional<SecretValue> getCached(String name) {
        return Optional.ofNullable(cache.get(Objects.requireNonNull(name, "name")));
    }

    public Mono<SecretValue> refresh(String name) {
        Objects.requireNonNull(name, "name");
        String defaultValue = defaultValues.get(name);
        if (defaultValue == null) {
            return Mono.error(new IllegalArgumentException(
                    "No default value registered for secret: " + name));
        }
        return provider.getSecret(name, defaultValue)
                .doOnNext(secret -> cache.put(name, secret));
    }

    public List<SecretValue> secretsNearExpiry() {
        OffsetDateTime threshold = OffsetDateTime.now(clock).plus(expiryWarningWindow);
        return cache.values().stream()
                .filter(secret -> secret.expiry().map(expiry -> !expiry.isAfter(threshold)).orElse(false))
                .toList();
    }

    public Mono<Void> refreshExpiringSecrets() {
        return Flux.fromIterable(secretsNearExpiry())
                .flatMap(secret -> refresh(secret.name()))
                .then();
    }

    public synchronized void startAutomaticRefresh(Duration checkInterval) {
        Duration interval = requirePositive(checkInterval, "checkInterval");
        if (automaticRefresh != null && !automaticRefresh.isDisposed()) {
            throw new IllegalStateException("Automatic refresh is already running");
        }
        automaticRefresh = Flux.interval(interval)
                .concatMap(tick -> refreshExpiringSecrets()
                        .doOnError(refreshErrorHandler)
                        .onErrorComplete())
                .subscribe();
    }

    private static Duration requirePositive(Duration duration, String name) {
        Objects.requireNonNull(duration, name);
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return duration;
    }

    @Override
    public synchronized void close() {
        if (automaticRefresh != null) {
            automaticRefresh.dispose();
        }
    }
}
