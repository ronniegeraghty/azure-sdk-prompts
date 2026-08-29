package com.example.appconfig;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.atomic.AtomicReference;

public final class AsyncConfigurationWatcher implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncConfigurationWatcher.class);

    private final AsyncConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final String label;
    private final Duration pollingInterval;
    private final Map<String, Optional<String>> lastValues = new LinkedHashMap<>();
    private final AtomicReference<Disposable> subscription = new AtomicReference<>();

    public AsyncConfigurationWatcher(
        AsyncConfigurationService configurationService,
        List<String> sentinelKeys,
        Duration pollingInterval
    ) {
        this(configurationService, sentinelKeys, null, pollingInterval);
    }

    public AsyncConfigurationWatcher(
        AsyncConfigurationService configurationService,
        List<String> sentinelKeys,
        String label,
        Duration pollingInterval
    ) {
        this.configurationService = Objects.requireNonNull(configurationService, "configurationService");
        this.sentinelKeys = List.copyOf(sentinelKeys);
        if (this.sentinelKeys.isEmpty()) {
            throw new IllegalArgumentException("At least one sentinel key is required");
        }
        this.label = label;
        this.pollingInterval = requirePositive(pollingInterval);
    }

    public void start() {
        Disposable candidate = Flux.interval(Duration.ZERO, pollingInterval)
            .concatMap(tick -> pollOnce()
                .onErrorResume(exception -> {
                    LOGGER.error("Unable to poll App Configuration sentinels", exception);
                    return Mono.empty();
                }))
            .subscribe();
        if (!subscription.compareAndSet(null, candidate)) {
            candidate.dispose();
        }
    }

    @Override
    public void close() {
        Disposable current = subscription.getAndSet(null);
        if (current != null) {
            current.dispose();
        }
    }

    private Mono<Void> pollOnce() {
        return Flux.fromIterable(sentinelKeys)
            .concatMap(key -> configurationService.getSetting(key, label)
                .map(value -> Map.entry(key, value)))
            .collectMap(Map.Entry::getKey, Map.Entry::getValue, LinkedHashMap::new)
            .flatMap(currentValues -> {
                boolean initialized = !lastValues.isEmpty();
                boolean changed = initialized && !lastValues.equals(currentValues);
                lastValues.clear();
                lastValues.putAll(currentValues);
                if (!changed) {
                    return Mono.empty();
                }
                return configurationService.refreshAll()
                    .doOnSuccess(ignored ->
                        LOGGER.info("A sentinel changed; refreshed all cached configuration"));
            });
    }

    private static Duration requirePositive(Duration duration) {
        Objects.requireNonNull(duration, "pollingInterval");
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException("Polling interval must be positive");
        }
        return duration;
    }
}
