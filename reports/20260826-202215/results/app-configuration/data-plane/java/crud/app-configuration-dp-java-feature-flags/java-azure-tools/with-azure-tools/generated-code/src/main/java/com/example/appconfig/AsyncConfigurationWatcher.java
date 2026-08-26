package com.example.appconfig;

import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class AsyncConfigurationWatcher implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(AsyncConfigurationWatcher.class.getName());

    private final AsyncConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final String label;
    private final Duration pollingInterval;
    private final Map<String, String> lastValues = new HashMap<>();
    private Disposable subscription;

    public AsyncConfigurationWatcher(
        AsyncConfigurationService configurationService,
        List<String> sentinelKeys,
        String label,
        Duration pollingInterval
    ) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
        this.sentinelKeys = List.copyOf(sentinelKeys);
        if (this.sentinelKeys.isEmpty()) {
            throw new IllegalArgumentException("sentinelKeys must not be empty");
        }
        this.label = label;
        this.pollingInterval = requirePositive(pollingInterval);
    }

    public void start() {
        if (subscription != null && !subscription.isDisposed()) {
            throw new IllegalStateException("Watcher is already running");
        }
        subscription = Flux.interval(Duration.ZERO, pollingInterval)
            .concatMap(ignored -> poll())
            .subscribe(
                ignored -> {
                },
                error -> LOGGER.log(Level.SEVERE, "Configuration watcher stopped", error)
            );
    }

    private Mono<Void> poll() {
        return Flux.fromIterable(sentinelKeys)
            .concatMap(key -> configurationService.getSetting(key, label).map(value -> Map.entry(key, value)))
            .collectMap(Map.Entry::getKey, Map.Entry::getValue)
            .flatMap(currentValues -> {
                boolean initialized = !lastValues.isEmpty();
                boolean changed = initialized && currentValues.entrySet().stream()
                    .anyMatch(entry -> !Objects.equals(lastValues.get(entry.getKey()), entry.getValue()));
                lastValues.clear();
                lastValues.putAll(currentValues);
                if (!changed) {
                    return Mono.empty();
                }
                LOGGER.info("Sentinel changed; refreshing all cached configuration");
                return configurationService.refreshAll();
            })
            .onErrorResume(error -> {
                LOGGER.log(Level.WARNING, "Configuration polling failed", error);
                return Mono.empty();
            });
    }

    private static Duration requirePositive(Duration interval) {
        Objects.requireNonNull(interval, "pollingInterval");
        if (interval.isZero() || interval.isNegative()) {
            throw new IllegalArgumentException("pollingInterval must be positive");
        }
        return interval;
    }

    @Override
    public void close() {
        if (subscription != null) {
            subscription.dispose();
        }
    }
}
