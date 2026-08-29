package com.example.appconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import reactor.core.publisher.Flux;

import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.Optional;

public final class Main {
    private static final String ENVIRONMENT_LABEL = "staging";
    private static final List<String> SAMPLE_USERS = List.of("alice", "bob", "carol", "dave");
    private static final List<ConfigurationWatcher.Sentinel> SENTINELS =
        List.of(new ConfigurationWatcher.Sentinel("Demo:Sentinel", ENVIRONMENT_LABEL));

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = System.getenv("AZURE_APP_CONFIGURATION_ENDPOINT");
        if (endpoint == null || endpoint.isBlank()) {
            System.out.println("Set AZURE_APP_CONFIGURATION_ENDPOINT to run the Azure App Configuration demo.");
            return;
        }

        Duration pollInterval = Duration.ofSeconds(readPositiveLong("CONFIG_POLL_INTERVAL_SECONDS", 2));
        Duration watchDuration = Duration.ofSeconds(readPositiveLong("DEMO_WATCH_SECONDS", 5));
        TokenCredential credential = managedIdentityCredential();

        runSyncDemo(endpoint, credential, pollInterval, watchDuration);
        runAsyncDemo(endpoint, credential, pollInterval, watchDuration);
    }

    private static void runSyncDemo(
        String endpoint,
        TokenCredential credential,
        Duration pollInterval,
        Duration watchDuration
    ) throws InterruptedException {
        System.out.println("\n=== Synchronous implementation ===");
        ConfigurationClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient();
        SyncConfigurationService service = new SyncConfigurationService(client);

        printValue("Demo:Message (no label)", service.getSetting("Demo:Message"));
        printValue("Demo:Message (staging)", service.getSetting("Demo:Message", ENVIRONMENT_LABEL));
        printSettings(service.listSettings("Demo:", ENVIRONMENT_LABEL));

        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(service, ENVIRONMENT_LABEL);
        for (String user : SAMPLE_USERS) {
            System.out.printf("BetaCheckout for %-5s: %s%n", user, flags.isEnabled("BetaCheckout", user));
        }

        try (ConfigurationWatcher watcher = ConfigurationWatcher.forSync(
            service,
            SENTINELS,
            pollInterval,
            changed -> System.out.println("Sync sentinel changed; cache refreshed: " + changed))) {
            watcher.start();
            Thread.sleep(watchDuration.toMillis());
        }
    }

    private static void runAsyncDemo(
        String endpoint,
        TokenCredential credential,
        Duration pollInterval,
        Duration watchDuration
    ) throws InterruptedException {
        System.out.println("\n=== Asynchronous implementation ===");
        ConfigurationAsyncClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient();
        AsyncConfigurationService service = new AsyncConfigurationService(client);
        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(service, ENVIRONMENT_LABEL);

        service.getSetting("Demo:Message")
            .doOnNext(value -> printValue("Demo:Message (no label)", value))
            .then(service.getSetting("Demo:Message", ENVIRONMENT_LABEL))
            .doOnNext(value -> printValue("Demo:Message (staging)", value))
            .then(service.listSettings("Demo:", ENVIRONMENT_LABEL))
            .doOnNext(Main::printSettings)
            .thenMany(Flux.fromIterable(SAMPLE_USERS)
                .concatMap(user -> flags.isEnabled("BetaCheckout", user)
                    .doOnNext(enabled ->
                        System.out.printf("BetaCheckout for %-5s: %s%n", user, enabled))))
            .then()
            .block();

        try (ConfigurationWatcher watcher = ConfigurationWatcher.forAsync(
            service,
            SENTINELS,
            pollInterval,
            changed -> System.out.println("Async sentinel changed; cache refreshed: " + changed))) {
            watcher.start();
            Thread.sleep(watchDuration.toMillis());
        }
    }

    private static TokenCredential managedIdentityCredential() {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            builder.clientId(clientId);
        }
        return builder.build();
    }

    private static long readPositiveLong(String name, long defaultValue) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            return defaultValue;
        }
        long parsed = Long.parseLong(value);
        if (parsed <= 0) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return parsed;
    }

    private static void printValue(String name, Optional<String> value) {
        System.out.printf("%s: %s%n", name, value.orElse("<missing>"));
    }

    private static void printSettings(Map<String, String> settings) {
        System.out.println("Settings with Demo: prefix:");
        settings.forEach((key, value) -> System.out.printf("  %s=%s%n", key, value));
    }
}
