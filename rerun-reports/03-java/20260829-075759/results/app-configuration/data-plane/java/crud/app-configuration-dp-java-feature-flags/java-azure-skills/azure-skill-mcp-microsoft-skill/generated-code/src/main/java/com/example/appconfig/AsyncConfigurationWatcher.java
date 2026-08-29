package com.example.appconfig;

import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicReference;

public final class AsyncConfigurationWatcher implements AutoCloseable {
    private static final System.Logger LOGGER =
        System.getLogger(AsyncConfigurationWatcher.class.getName());

    private final AsyncConfigurationService configuration;
    private final List<Sentinel> sentinels;
    private final Duration pollingInterval;
    private final Runnable onRefresh;
    private final AtomicReference<Disposable> subscription = new AtomicReference<>();

    public AsyncConfigurationWatcher(
        AsyncConfigurationService configuration,
        List<Sentinel> sentinels,
        Duration pollingInterval,
        Runnable onRefresh
    ) {
        this.configuration = Objects.requireNonNull(configuration, "configuration");
        this.sentinels = List.copyOf(sentinels);
        if (this.sentinels.isEmpty()) {
            throw new IllegalArgumentException("At least one sentinel is required");
        }
        this.pollingInterval = requirePositive(pollingInterval);
        this.onRefresh = Objects.requireNonNull(onRefresh, "onRefresh");
    }

    public void start() {
        Disposable watcher = Flux.interval(Duration.ZERO, pollingInterval)
            .concatMap(ignored -> pollOnce()
                .onErrorResume(error -> {
                    LOGGER.log(System.Logger.Level.ERROR, "Async configuration polling failed", error);
                    return Mono.empty();
                }))
            .subscribe();
        if (!subscription.compareAndSet(null, watcher)) {
            watcher.dispose();
            throw new IllegalStateException("Watcher has already been started");
        }
    }

    private Mono<Void> pollOnce() {
        return Flux.fromIterable(sentinels)
            .concatMap(configuration::checkForUpdate)
            .any(Boolean::booleanValue)
            .flatMap(changed -> {
                if (!changed) {
                    return Mono.empty();
                }
                return configuration.refreshAll().doOnSuccess(ignored -> onRefresh.run());
            });
    }

    private static Duration requirePositive(Duration duration) {
        Objects.requireNonNull(duration, "pollingInterval");
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException("Polling interval must be positive");
        }
        return duration;
    }

    @Override
    public void close() {
        Disposable watcher = subscription.getAndSet(null);
        if (watcher != null) {
            watcher.dispose();
        }
    }
}
