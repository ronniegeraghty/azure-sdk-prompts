package com.example.appconfig;

import java.time.Duration;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class ConfigurationWatcher implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(ConfigurationWatcher.class.getName());

    private final ConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final String label;
    private final Duration pollingInterval;
    private final Map<String, String> lastValues = new HashMap<>();
    private final ScheduledExecutorService executor = Executors.newSingleThreadScheduledExecutor(task -> {
        Thread thread = new Thread(task, "app-configuration-watcher");
        thread.setDaemon(true);
        return thread;
    });

    public ConfigurationWatcher(
        ConfigurationService configurationService,
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
        executor.scheduleAtFixedRate(
            this::pollSafely,
            0,
            pollingInterval.toMillis(),
            TimeUnit.MILLISECONDS
        );
    }

    private void pollSafely() {
        try {
            boolean initialized = !lastValues.isEmpty();
            boolean changed = false;
            for (String key : sentinelKeys) {
                String current = configurationService.getSetting(key, label);
                String previous = lastValues.put(key, current);
                changed |= initialized && !Objects.equals(previous, current);
            }
            if (changed) {
                LOGGER.info("Sentinel changed; refreshing all cached configuration");
                configurationService.refreshAll();
            }
        } catch (RuntimeException exception) {
            LOGGER.log(Level.WARNING, "Configuration polling failed", exception);
        }
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
        executor.shutdownNow();
    }
}
