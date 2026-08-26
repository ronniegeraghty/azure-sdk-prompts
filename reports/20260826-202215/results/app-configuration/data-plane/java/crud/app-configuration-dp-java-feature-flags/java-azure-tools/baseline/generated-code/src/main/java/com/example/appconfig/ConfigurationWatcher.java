package com.example.appconfig;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;

public class ConfigurationWatcher implements AutoCloseable {
    private final ConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final Duration pollingInterval;
    private final Consumer<List<String>> refreshListener;
    private final Map<String, Optional<String>> sentinelValues = new ConcurrentHashMap<>();
    private final AtomicBoolean started = new AtomicBoolean();
    private final ScheduledExecutorService executor = Executors.newSingleThreadScheduledExecutor(runnable -> {
        Thread thread = new Thread(runnable, "app-configuration-watcher");
        thread.setDaemon(true);
        return thread;
    });

    public ConfigurationWatcher(
        ConfigurationService configurationService,
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

    public void start() {
        if (started.compareAndSet(false, true)) {
            executor.scheduleWithFixedDelay(
                this::pollSafely,
                0,
                pollingInterval.toMillis(),
                TimeUnit.MILLISECONDS);
        }
    }

    private void pollSafely() {
        try {
            List<String> changedKeys = new ArrayList<>();
            for (String key : sentinelKeys) {
                Optional<String> current = configurationService.getSetting(key);
                Optional<String> previous = sentinelValues.put(key, current);
                if (previous != null && !previous.equals(current)) {
                    changedKeys.add(key);
                }
            }

            if (!changedKeys.isEmpty()) {
                configurationService.refreshAll();
                refreshListener.accept(List.copyOf(changedKeys));
            }
        } catch (RuntimeException exception) {
            System.err.println("Configuration watcher poll failed: " + exception.getMessage());
        }
    }

    @Override
    public void close() {
        executor.shutdownNow();
    }

    private static Duration requirePositive(Duration duration) {
        Objects.requireNonNull(duration, "pollingInterval");
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException("pollingInterval must be positive");
        }
        return duration;
    }
}
