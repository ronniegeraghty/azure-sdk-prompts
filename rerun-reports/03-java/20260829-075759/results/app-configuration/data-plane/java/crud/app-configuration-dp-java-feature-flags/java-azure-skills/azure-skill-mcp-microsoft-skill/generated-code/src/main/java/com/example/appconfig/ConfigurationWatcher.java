package com.example.appconfig;

import java.time.Duration;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;

public final class ConfigurationWatcher implements AutoCloseable {
    private static final System.Logger LOGGER =
        System.getLogger(ConfigurationWatcher.class.getName());

    private final ConfigurationService configuration;
    private final List<Sentinel> sentinels;
    private final Duration pollingInterval;
    private final Runnable onRefresh;
    private final ScheduledExecutorService scheduler =
        Executors.newSingleThreadScheduledExecutor(runnable -> {
            Thread thread = new Thread(runnable, "app-configuration-watcher");
            thread.setDaemon(true);
            return thread;
        });
    private final AtomicBoolean started = new AtomicBoolean();

    public ConfigurationWatcher(
        ConfigurationService configuration,
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
        if (!started.compareAndSet(false, true)) {
            throw new IllegalStateException("Watcher has already been started");
        }
        scheduler.scheduleWithFixedDelay(
            this::pollSafely,
            0,
            pollingInterval.toMillis(),
            TimeUnit.MILLISECONDS
        );
    }

    private void pollSafely() {
        try {
            boolean changed = sentinels.stream()
                .map(configuration::checkForUpdate)
                .reduce(false, Boolean::logicalOr);
            if (changed) {
                configuration.refreshAll();
                onRefresh.run();
            }
        } catch (RuntimeException exception) {
            LOGGER.log(System.Logger.Level.ERROR, "Configuration polling failed", exception);
        }
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
        scheduler.shutdownNow();
    }
}
