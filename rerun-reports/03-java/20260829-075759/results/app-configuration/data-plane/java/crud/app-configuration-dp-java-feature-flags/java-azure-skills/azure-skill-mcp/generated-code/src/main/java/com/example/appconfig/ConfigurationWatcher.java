package com.example.appconfig;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;

public final class ConfigurationWatcher implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(ConfigurationWatcher.class);

    private final ConfigurationService configurationService;
    private final List<String> sentinelKeys;
    private final String label;
    private final Duration pollingInterval;
    private final Map<String, Optional<String>> lastValues = new LinkedHashMap<>();
    private final ScheduledExecutorService executor;
    private final AtomicBoolean started = new AtomicBoolean();

    public ConfigurationWatcher(
        ConfigurationService configurationService,
        List<String> sentinelKeys,
        Duration pollingInterval
    ) {
        this(configurationService, sentinelKeys, null, pollingInterval);
    }

    public ConfigurationWatcher(
        ConfigurationService configurationService,
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
        ThreadFactory threadFactory = runnable -> {
            Thread thread = new Thread(runnable, "app-configuration-watcher");
            thread.setDaemon(true);
            return thread;
        };
        this.executor = Executors.newSingleThreadScheduledExecutor(threadFactory);
    }

    public void start() {
        if (started.compareAndSet(false, true)) {
            executor.scheduleWithFixedDelay(
                this::pollSafely, 0, pollingInterval.toMillis(), TimeUnit.MILLISECONDS);
        }
    }

    @Override
    public void close() {
        executor.shutdownNow();
    }

    private void pollSafely() {
        try {
            boolean changed = false;
            for (String key : sentinelKeys) {
                Optional<String> current = configurationService.getSetting(key, label);
                Optional<String> previous = lastValues.put(key, current);
                if (previous != null && !previous.equals(current)) {
                    changed = true;
                }
            }
            if (changed) {
                configurationService.refreshAll();
                LOGGER.info("A sentinel changed; refreshed all cached configuration");
            }
        } catch (RuntimeException exception) {
            LOGGER.error("Unable to poll App Configuration sentinels", exception);
        }
    }

    private static Duration requirePositive(Duration duration) {
        Objects.requireNonNull(duration, "pollingInterval");
        if (duration.isZero() || duration.isNegative()) {
            throw new IllegalArgumentException("Polling interval must be positive");
        }
        return duration;
    }
}
