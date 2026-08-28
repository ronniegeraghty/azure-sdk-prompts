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
    private static final String ENVIRONMENT_LABEL = "production";
    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(5);
    private static final List<String> SENTINELS = List.of("demo:sentinel");
    private static final List<String> SAMPLE_USERS = List.of("alice", "bob", "carol", "dave");

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable("AZURE_APPCONFIG_ENDPOINT");
        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

        runSyncDemo(new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient());

        runAsyncDemo(new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient());
    }

    private static void runSyncDemo(ConfigurationClient client) throws InterruptedException {
        System.out.println("=== Synchronous demo ===");
        ConfigurationService configuration = new ConfigurationService(client);
        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(configuration);

        printConfiguration(
            configuration.getSetting("demo:message"),
            configuration.getSetting("demo:message", ENVIRONMENT_LABEL),
            configuration.listSettings("demo:", ENVIRONMENT_LABEL));
        SAMPLE_USERS.forEach(user -> System.out.printf(
            "beta-dashboard for %-5s: %s%n",
            user,
            flags.isEnabled("beta-dashboard", user, ENVIRONMENT_LABEL)));

        try (ConfigurationWatcher watcher = ConfigurationWatcher.forSync(
            configuration,
            SENTINELS,
            ENVIRONMENT_LABEL,
            POLLING_INTERVAL,
            () -> System.out.println("Sync cache refreshed after sentinel change."),
            error -> System.err.println("Sync watcher failed: " + error.getMessage()))) {
            watcher.start();
            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());
        }
    }

    private static void runAsyncDemo(ConfigurationAsyncClient client) throws InterruptedException {
        System.out.println("\n=== Asynchronous demo ===");
        AsyncConfigurationService configuration = new AsyncConfigurationService(client);
        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(configuration);

        String unlabeled = configuration.getSetting("demo:message").block();
        String labeled = configuration.getSetting("demo:message", ENVIRONMENT_LABEL).block();
        Map<String, String> prefixed = configuration.listSettings("demo:", ENVIRONMENT_LABEL).block();
        printConfiguration(unlabeled, labeled, prefixed);

        flagsForUsers(flags).block();

        try (ConfigurationWatcher watcher = ConfigurationWatcher.forAsync(
            configuration,
            SENTINELS,
            ENVIRONMENT_LABEL,
            POLLING_INTERVAL,
            () -> System.out.println("Async cache refreshed after sentinel change."),
            error -> System.err.println("Async watcher failed: " + error.getMessage()))) {
            watcher.start();
            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());
        }
    }

    private static reactor.core.publisher.Mono<Void> flagsForUsers(AsyncFeatureFlagEvaluator flags) {
        return reactor.core.publisher.Flux.fromIterable(SAMPLE_USERS)
            .concatMap(user -> flags.isEnabled("beta-dashboard", user, ENVIRONMENT_LABEL)
                .doOnNext(enabled -> System.out.printf(
                    "beta-dashboard for %-5s: %s%n", user, enabled)))
            .then();
    }

    private static void printConfiguration(
        String unlabeled,
        String labeled,
        Map<String, String> prefixed
    ) {
        System.out.println("Unlabeled message: " + unlabeled);
        System.out.println("Production message: " + labeled);
        System.out.println("Production demo settings: " + prefixed);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must contain the App Configuration endpoint");
        }
        return value;
    }
}
