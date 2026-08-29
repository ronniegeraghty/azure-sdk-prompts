package com.example.appconfig;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;
import java.util.function.Supplier;
import java.util.stream.Collectors;

public final class ConfigurationWatcher implements AutoCloseable {
    public record Sentinel(String key, String label) {
        public Sentinel {
            if (key == null || key.isBlank()) {
                throw new IllegalArgumentException("Sentinel key must not be blank");
            }
        }

        public Sentinel(String key) {
            this(key, null);
        }
    }

    private final Duration pollingInterval;
    private final Supplier<CompletionStage<Set<Sentinel>>> poll;
    private final Consumer<Set<Sentinel>> changeListener;
    private final Consumer<Throwable> errorListener;
    private final ScheduledExecutorService scheduler;
    private final AtomicBoolean started = new AtomicBoolean();
    private final AtomicBoolean polling = new AtomicBoolean();

    private ConfigurationWatcher(
        Duration pollingInterval,
        Supplier<CompletionStage<Set<Sentinel>>> poll,
        Consumer<Set<Sentinel>> changeListener,
        Consumer<Throwable> errorListener
    ) {
        if (pollingInterval == null || pollingInterval.isZero() || pollingInterval.isNegative()) {
            throw new IllegalArgumentException("Polling interval must be positive");
        }
        this.pollingInterval = pollingInterval;
        this.poll = poll;
        this.changeListener = changeListener;
        this.errorListener = errorListener;
        this.scheduler = Executors.newSingleThreadScheduledExecutor(runnable -> {
            Thread thread = new Thread(runnable, "app-configuration-watcher");
            thread.setDaemon(true);
            return thread;
        });
    }

    public static ConfigurationWatcher forSync(
        SyncConfigurationService service,
        List<Sentinel> sentinels,
        Duration pollingInterval,
        Consumer<Set<Sentinel>> changeListener
    ) {
        Objects.requireNonNull(service, "service");
        List<Sentinel> watched = validateSentinels(sentinels);
        Supplier<CompletionStage<Set<Sentinel>>> poll = () -> {
            Set<Sentinel> changed = watched.stream()
                .filter(sentinel -> service.hasSettingChanged(sentinel.key(), sentinel.label()))
                .collect(Collectors.toUnmodifiableSet());
            if (!changed.isEmpty()) {
                service.refreshAll();
            }
            return CompletableFuture.completedFuture(changed);
        };
        return new ConfigurationWatcher(
            pollingInterval, poll, changeListener, ConfigurationWatcher::reportError);
    }

    public static ConfigurationWatcher forAsync(
        AsyncConfigurationService service,
        List<Sentinel> sentinels,
        Duration pollingInterval,
        Consumer<Set<Sentinel>> changeListener
    ) {
        Objects.requireNonNull(service, "service");
        List<Sentinel> watched = validateSentinels(sentinels);
        Supplier<CompletionStage<Set<Sentinel>>> poll = () -> Flux.fromIterable(watched)
            .concatMap(sentinel -> service.hasSettingChanged(sentinel.key(), sentinel.label())
                .filter(Boolean::booleanValue)
                .map(ignored -> sentinel))
            .collect(Collectors.toUnmodifiableSet())
            .flatMap(changed -> changed.isEmpty()
                ? Mono.just(changed)
                : service.refreshAll().thenReturn(changed))
            .toFuture();
        return new ConfigurationWatcher(
            pollingInterval, poll, changeListener, ConfigurationWatcher::reportError);
    }

    public void start() {
        if (!started.compareAndSet(false, true)) {
            throw new IllegalStateException("Configuration watcher has already been started");
        }
        scheduler.scheduleWithFixedDelay(
            this::pollOnce,
            0,
            pollingInterval.toMillis(),
            TimeUnit.MILLISECONDS);
    }

    private void pollOnce() {
        if (!polling.compareAndSet(false, true)) {
            return;
        }
        try {
            poll.get().whenComplete((changed, error) -> {
                try {
                    if (error != null) {
                        errorListener.accept(unwrap(error));
                    } else if (!changed.isEmpty()) {
                        changeListener.accept(changed);
                    }
                } finally {
                    polling.set(false);
                }
            });
        } catch (RuntimeException exception) {
            polling.set(false);
            errorListener.accept(exception);
        }
    }

    @Override
    public void close() {
        scheduler.shutdownNow();
    }

    private static List<Sentinel> validateSentinels(List<Sentinel> sentinels) {
        Objects.requireNonNull(sentinels, "sentinels");
        if (sentinels.isEmpty()) {
            throw new IllegalArgumentException("At least one sentinel is required");
        }
        return List.copyOf(sentinels);
    }

    private static Throwable unwrap(Throwable error) {
        return error.getCause() == null ? error : error.getCause();
    }

    private static void reportError(Throwable error) {
        System.err.println("Configuration watcher poll failed: " + error.getMessage());
    }
}
