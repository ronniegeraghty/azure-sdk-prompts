package com.example.appconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import reactor.core.publisher.Flux;

import java.time.Duration;
import java.util.List;

public final class Main {
    private static final String ENVIRONMENT_LABEL = "production";
    private static final List<String> SAMPLE_USERS = List.of("alice", "bob", "carol", "dave");

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable("AZURE_APPCONFIG_ENDPOINT");
        int pollSeconds = Integer.parseInt(System.getenv().getOrDefault("CONFIG_POLL_SECONDS", "10"));
        Duration pollingInterval = Duration.ofSeconds(pollSeconds);

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        ConfigurationClient syncClient = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient();
        ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient();

        runSyncDemo(syncClient, pollingInterval);
        runAsyncDemo(asyncClient, pollingInterval);
    }

    private static void runSyncDemo(ConfigurationClient client, Duration pollingInterval)
        throws InterruptedException {
        System.out.println("=== Synchronous implementation ===");
        ConfigurationService service = new ConfigurationService(client);
        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(service);

        System.out.println("application:name = " + service.getSetting("application:name"));
        System.out.println("application:message [production] = "
            + service.getSetting("application:message", ENVIRONMENT_LABEL));
        System.out.println("application:* [production] = "
            + service.listSettings("application:", ENVIRONMENT_LABEL));
        SAMPLE_USERS.forEach(user -> System.out.printf(
            "BetaDashboard for %-5s = %s%n",
            user,
            flags.isEnabled("BetaDashboard", user, ENVIRONMENT_LABEL)
        ));

        try (ConfigurationWatcher watcher = new ConfigurationWatcher(
            service,
            List.of("application:sentinel"),
            ENVIRONMENT_LABEL,
            pollingInterval
        )) {
            watcher.start();
            System.out.println("Watching the sync sentinel for " + pollingInterval.multipliedBy(2) + "...");
            Thread.sleep(pollingInterval.multipliedBy(2).toMillis());
        }
    }

    private static void runAsyncDemo(ConfigurationAsyncClient client, Duration pollingInterval)
        throws InterruptedException {
        System.out.println("\n=== Asynchronous implementation ===");
        AsyncConfigurationService service = new AsyncConfigurationService(client);
        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(service);

        service.getSetting("application:name")
            .doOnNext(value -> System.out.println("application:name = " + value))
            .then(service.getSetting("application:message", ENVIRONMENT_LABEL)
                .doOnNext(value -> System.out.println("application:message [production] = " + value)))
            .then(service.listSettings("application:", ENVIRONMENT_LABEL)
                .doOnNext(value -> System.out.println("application:* [production] = " + value)))
            .thenMany(Flux.fromIterable(SAMPLE_USERS)
                .concatMap(user -> flags.isEnabled("BetaDashboard", user, ENVIRONMENT_LABEL)
                    .doOnNext(enabled -> System.out.printf(
                        "BetaDashboard for %-5s = %s%n",
                        user,
                        enabled
                    ))))
            .then()
            .block();

        try (AsyncConfigurationWatcher watcher = new AsyncConfigurationWatcher(
            service,
            List.of("application:sentinel"),
            ENVIRONMENT_LABEL,
            pollingInterval
        )) {
            watcher.start();
            System.out.println("Watching the async sentinel for " + pollingInterval.multipliedBy(2) + "...");
            Thread.sleep(pollingInterval.multipliedBy(2).toMillis());
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must contain the App Configuration endpoint");
        }
        return value;
    }
}
