package com.example.appconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.time.Duration;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final String PRODUCTION = "production";
    private static final String STAGING = "staging";
    private static final String SENTINEL_KEY = "app:sentinel";
    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(5);

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requiredEnvironmentVariable("AZURE_APPCONFIG_ENDPOINT");
        long watchSeconds = Long.parseLong(
            System.getenv().getOrDefault("DEMO_WATCH_SECONDS", "15")
        );
        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

        runSyncDemo(endpoint, credential, Duration.ofSeconds(watchSeconds));
        runAsyncDemo(endpoint, credential, Duration.ofSeconds(watchSeconds));
    }

    private static void runSyncDemo(
        String endpoint,
        TokenCredential credential,
        Duration watchDuration
    ) throws InterruptedException {
        System.out.println("=== Synchronous implementation ===");
        ConfigurationClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient();
        ConfigurationService configuration = new ConfigurationService(client);
        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(configuration);

        printSetting("Production greeting", configuration.getSetting("app:greeting", PRODUCTION).orElse("<missing>"));
        printSetting("Staging greeting", configuration.getSetting("app:greeting", STAGING).orElse("<missing>"));
        printSettings(configuration.listSettings("app:", PRODUCTION));
        printUsers(user -> flags.isEnabledForUser("beta-dashboard", PRODUCTION, user));

        System.out.printf("Watching '%s' for %d seconds...%n", SENTINEL_KEY, watchDuration.toSeconds());
        try (ConfigurationWatcher watcher = new ConfigurationWatcher(
            configuration,
            List.of(new Sentinel(SENTINEL_KEY, PRODUCTION)),
            POLLING_INTERVAL,
            () -> System.out.println("Sync cache refreshed after sentinel change")
        )) {
            watcher.start();
            Thread.sleep(watchDuration.toMillis());
        }
    }

    private static void runAsyncDemo(
        String endpoint,
        TokenCredential credential,
        Duration watchDuration
    ) throws InterruptedException {
        System.out.println("\n=== Asynchronous implementation ===");
        ConfigurationAsyncClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient();
        AsyncConfigurationService configuration = new AsyncConfigurationService(client);
        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(configuration);

        printSetting("Production greeting",
            configuration.getSetting("app:greeting", PRODUCTION).defaultIfEmpty("<missing>").block());
        printSetting("Staging greeting",
            configuration.getSetting("app:greeting", STAGING).defaultIfEmpty("<missing>").block());
        printSettings(configuration.listSettings("app:", PRODUCTION).block());
        for (String user : List.of("alice", "bob", "carol", "dave")) {
            Boolean enabled = flags.isEnabledForUser("beta-dashboard", PRODUCTION, user).block();
            System.out.printf("beta-dashboard for %-5s: %s%n", user, enabled);
        }

        System.out.printf("Watching '%s' for %d seconds...%n", SENTINEL_KEY, watchDuration.toSeconds());
        try (AsyncConfigurationWatcher watcher = new AsyncConfigurationWatcher(
            configuration,
            List.of(new Sentinel(SENTINEL_KEY, PRODUCTION)),
            POLLING_INTERVAL,
            () -> System.out.println("Async cache refreshed after sentinel change")
        )) {
            watcher.start();
            Thread.sleep(watchDuration.toMillis());
        }
    }

    private static void printSetting(String name, String value) {
        System.out.printf("%s: %s%n", name, value);
    }

    private static void printSettings(Map<String, String> settings) {
        System.out.println("Production settings:");
        settings.forEach((key, value) -> System.out.printf("  %s = %s%n", key, value));
    }

    private static void printUsers(UserFlagCheck check) {
        for (String user : List.of("alice", "bob", "carol", "dave")) {
            System.out.printf("beta-dashboard for %-5s: %s%n", user, check.isEnabled(user));
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must contain the App Configuration endpoint");
        }
        return value;
    }

    @FunctionalInterface
    private interface UserFlagCheck {
        boolean isEnabled(String user);
    }
}
