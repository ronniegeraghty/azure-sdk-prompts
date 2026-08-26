package com.example.appconfig;

import com.azure.core.credential.TokenCredential;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.ConfigurationClientBuilder;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.util.List;

public final class Main {
    private static final String ENDPOINT_ENV = "AZURE_APPCONFIG_ENDPOINT";
    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(5);
    private static final List<String> SAMPLE_USERS = List.of("alice", "bob", "charlie", "diana");

    private Main() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable(ENDPOINT_ENV);
        TokenCredential credential = new ManagedIdentityCredentialBuilder().build();

        ConfigurationClient syncClient = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient();
        runSyncDemo(new ConfigurationService(syncClient));

        ConfigurationAsyncClient asyncClient = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient();
        runAsyncDemo(new AsyncConfigurationService(asyncClient));
    }

    private static void runSyncDemo(ConfigurationService service) throws InterruptedException {
        System.out.println("=== Synchronous demo ===");
        print("Production message", service.getSetting("app:message", "production").orElse("<missing>"));
        print("Staging message", service.getSetting("app:message", "staging").orElse("<missing>"));
        print("Production app settings", service.listSettings("app:", "production"));

        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(service, "production");
        for (String userId : SAMPLE_USERS) {
            print("BetaFeature for " + userId, flags.isEnabled("BetaFeature", userId));
        }

        try (ConfigurationWatcher watcher = new ConfigurationWatcher(
            service,
            List.of("app:sentinel"),
            POLLING_INTERVAL,
            changed -> System.out.println("Sync refresh triggered by " + changed)
        )) {
            watcher.start();
            Thread.sleep(POLLING_INTERVAL.multipliedBy(2).toMillis());
        }
    }

    private static void runAsyncDemo(AsyncConfigurationService service) {
        System.out.println("=== Asynchronous demo ===");
        Mono.zip(
                service.getSetting("app:message", "production"),
                service.getSetting("app:message", "staging"),
                service.listSettings("app:", "production"))
            .doOnNext(values -> {
                print("Production message", values.getT1().orElse("<missing>"));
                print("Staging message", values.getT2().orElse("<missing>"));
                print("Production app settings", values.getT3());
            })
            .block();

        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(service, "production");
        Flux.fromIterable(SAMPLE_USERS)
            .concatMap(userId -> flags.isEnabled("BetaFeature", userId)
                .doOnNext(enabled -> print("BetaFeature for " + userId, enabled)))
            .then()
            .block();

        try (AsyncConfigurationWatcher watcher = new AsyncConfigurationWatcher(
            service,
            List.of("app:sentinel"),
            POLLING_INTERVAL,
            changed -> System.out.println("Async refresh triggered by " + changed)
        )) {
            watcher.start();
            Mono.delay(POLLING_INTERVAL.multipliedBy(2)).block();
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set " + name + " to your Azure App Configuration endpoint");
        }
        return value;
    }

    private static void print(String label, Object value) {
        System.out.printf("%-28s %s%n", label + ":", value);
    }
}
