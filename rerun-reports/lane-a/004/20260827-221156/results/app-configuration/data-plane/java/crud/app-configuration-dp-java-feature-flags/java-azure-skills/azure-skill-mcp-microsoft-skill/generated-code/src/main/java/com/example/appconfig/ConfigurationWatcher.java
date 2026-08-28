package com.example.appconfig;

import reactor.core.Disposable;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class ConfigurationWatcher implements AutoCloseable {
    private final Runnable startAction;
    private final Runnable closeAction;
    private boolean started;

    private ConfigurationWatcher(Runnable startAction, Runnable closeAction) {
        this.startAction = startAction;
        this.closeAction = closeAction;
    }

    public static ConfigurationWatcher forSync(
        ConfigurationService service,
        List<String> sentinelKeys,
        String label,
        Duration pollingInterval,
        Runnable onRefresh,
        Consumer<Throwable> errorHandler
    ) {
        validate(sentinelKeys, pollingInterval);
        ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(runnable -> {
            Thread thread = new Thread(runnable, "app-configuration-watcher");
            thread.setDaemon(true);
            return thread;
        });
        Map<String, String> previousValues = new LinkedHashMap<>();

        Runnable poll = () -> {
            try {
                boolean changed = false;
                for (String key : sentinelKeys) {
                    String current = service.getSetting(key, label);
                    String previous = previousValues.put(key, current);
                    changed |= previous != null && !Objects.equals(previous, current);
                }
                if (changed) {
                    service.refreshAll();
                    onRefresh.run();
                }
            } catch (RuntimeException error) {
                errorHandler.accept(error);
            }
        };

        return new ConfigurationWatcher(
            () -> scheduler.scheduleWithFixedDelay(
                poll, 0, pollingInterval.toMillis(), TimeUnit.MILLISECONDS),
            scheduler::shutdownNow);
    }

    public static ConfigurationWatcher forAsync(
        AsyncConfigurationService service,
        List<String> sentinelKeys,
        String label,
        Duration pollingInterval,
        Runnable onRefresh,
        Consumer<Throwable> errorHandler
    ) {
        validate(sentinelKeys, pollingInterval);
        Map<String, String> previousValues = new LinkedHashMap<>();
        Disposable[] subscription = new Disposable[1];

        Mono<Void> poll = Flux.fromIterable(sentinelKeys)
            .concatMap(key -> service.getSetting(key, label).map(value -> Map.entry(key, value)))
            .collectMap(Map.Entry::getKey, Map.Entry::getValue, LinkedHashMap::new)
            .flatMap(currentValues -> {
                boolean changed = !previousValues.isEmpty() && sentinelKeys.stream()
                    .anyMatch(key -> !Objects.equals(previousValues.get(key), currentValues.get(key)));
                previousValues.clear();
                previousValues.putAll(currentValues);
                return changed ? service.refreshAll().doOnSuccess(ignored -> onRefresh.run()) : Mono.empty();
            });

        return new ConfigurationWatcher(
            () -> subscription[0] = Flux.interval(Duration.ZERO, pollingInterval)
                .concatMap(ignored -> poll)
                .subscribe(ignored -> {
                }, errorHandler),
            () -> {
                if (subscription[0] != null) {
                    subscription[0].dispose();
                }
            });
    }

    public synchronized void start() {
        if (started) {
            throw new IllegalStateException("Watcher has already been started");
        }
        started = true;
        startAction.run();
    }

    @Override
    public synchronized void close() {
        if (started) {
            closeAction.run();
            started = false;
        }
    }

    private static void validate(List<String> sentinelKeys, Duration pollingInterval) {
        if (sentinelKeys.isEmpty()) {
            throw new IllegalArgumentException("At least one sentinel key is required");
        }
        if (pollingInterval.isZero() || pollingInterval.isNegative()) {
            throw new IllegalArgumentException("Polling interval must be positive");
        }
    }
}
