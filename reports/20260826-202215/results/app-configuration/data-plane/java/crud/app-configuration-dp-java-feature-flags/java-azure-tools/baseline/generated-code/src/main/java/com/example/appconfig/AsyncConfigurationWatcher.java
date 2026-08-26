package com.example.appconfig;

import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Consumer;

public class AsyncConfigurationWatcher implements AutoCloseable {
    private final AsyncConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final Duration pollingInterval;
    private final Consumer<List<String>> refreshListener;
    private final Map<String, Optional<String>> sentinelValues = new ConcurrentHashMap<>();
    private Disposable subscription;

    public AsyncConfigurationWatcher(
        AsyncConfigurationService configurationService,
        List<String> sentinelKeys,
        Duration pollingInterval,
        Consumer<List<String>> refreshListener
    ) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
        this.sentinelKeys = List.copyOf(sentinelKeys);
        if (this.sentinelKeys.isEmpty()) {
            throw new IllegalArgumentException("At least one sentinel key is required");
        }
        this.pollingInterval = requirePositive(pollingInterval);
        this.refreshListener = Objects.requireNonNull(refreshListener, "refreshListener");
    }

    public synchronized void start() {
        if (subscription == null || subscription.isDisposed()) {
            subscription = Flux.interval(Duration.ZERO, pollingInterval)
                .concatMap(ignored -> poll()
                    .onErrorResume(exception -> {
                        System.err.println("Async configuration watcher poll failed: " + exception.getMessage());
                        return Mono.empty();
                    }))
                .subscribe();
        }
    }

    private Mono<Void> poll() {
        return Flux.fromIterable(sentinelKeys)
            .concatMap(key -> configurationService.getSetting(key)
                .map(current -> new SentinelRead(key, current)))
            .filter(read -> {
                Optional<String> previous = sentinelValues.put(read.key(), read.value());
                return previous != null && !previous.equals(read.value());
            })
            .map(SentinelRead::key)
            .collectList()
            .flatMap(changedKeys -> {
                if (changedKeys.isEmpty()) {
                    return Mono.empty();
                }
                return configurationService.refreshAll()
                    .then(Mono.fromRunnable(() -> refreshListener.accept(List.copyOf(changedKeys))));
            });
    }

    @Override
    public synchronized void close() {
        if (subscription != null) {
            subscription.dispose();
            subscription = null;
        }
    }

    private static Duration requirePositive(Duration duration) {
        Objects.requireNonNull(duration, "pollingInterval");
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException("pollingInterval must be positive");
        }
        return duration;
    }

    private record SentinelRead(String key, Optional<String> value) {
    }
}
