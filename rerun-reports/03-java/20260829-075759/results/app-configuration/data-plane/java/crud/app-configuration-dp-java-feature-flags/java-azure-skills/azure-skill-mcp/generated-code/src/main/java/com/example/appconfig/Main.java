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
    private static final String PRODUCTION = "production";
    private static final String STAGING = "staging";
    private static final Duration POLLING_INTERVAL = Duration.ofSeconds(10);
    private static final Duration WATCH_DURATION = Duration.ofSeconds(20);
    private static final List<String> SAMPLE_USERS = List.of("alice", "bob", "carol", "dave");

    private Main() {
    }

    public static void main(String[] args) {
        String endpoint = requiredEnvironmentVariable("AZURE_APPCONFIG_ENDPOINT");
        TokenCredential credential = managedIdentityCredential();

        runSyncDemo(endpoint, credential);
        runAsyncDemo(endpoint, credential).block();
    }

    private static void runSyncDemo(String endpoint, TokenCredential credential) {
        System.out.println("=== Synchronous demo ===");
        ConfigurationClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildClient();
        ConfigurationService service = new ConfigurationService(client);
        FeatureFlagEvaluator flags = new FeatureFlagEvaluator(service);

        print("App:Title (production)", service.getSetting("App:Title", PRODUCTION).orElse("<missing>"));
        print("App:Title (staging)", service.getSetting("App:Title", STAGING).orElse("<missing>"));
        print("App:* (production)", service.listSettings("App:", PRODUCTION));
        SAMPLE_USERS.forEach(user ->
            print("BetaCheckout for " + user, flags.isEnabled("BetaCheckout", PRODUCTION, user)));

        System.out.println("Watching Demo:Sentinel for " + WATCH_DURATION.toSeconds() + " seconds...");
        try (ConfigurationWatcher watcher =
            new ConfigurationWatcher(service, List.of("Demo:Sentinel"), PRODUCTION, POLLING_INTERVAL)) {
            watcher.start();
            sleep(WATCH_DURATION);
        }
    }

    private static Mono<Void> runAsyncDemo(String endpoint, TokenCredential credential) {
        System.out.println("=== Asynchronous demo ===");
        ConfigurationAsyncClient client = new ConfigurationClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .buildAsyncClient();
        AsyncConfigurationService service = new AsyncConfigurationService(client);
        AsyncFeatureFlagEvaluator flags = new AsyncFeatureFlagEvaluator(service);

        Mono<Void> reads = Mono.when(
            service.getSetting("App:Title", PRODUCTION)
                .doOnNext(value -> print("App:Title (production)", value.orElse("<missing>"))),
            service.getSetting("App:Title", STAGING)
                .doOnNext(value -> print("App:Title (staging)", value.orElse("<missing>"))),
            service.listSettings("App:", PRODUCTION)
                .doOnNext(value -> print("App:* (production)", value)),
            Flux.fromIterable(SAMPLE_USERS)
                .concatMap(user -> flags.isEnabled("BetaCheckout", PRODUCTION, user)
                    .doOnNext(enabled -> print("BetaCheckout for " + user, enabled)))
                .then());

        return reads.then(Mono.using(
            () -> new AsyncConfigurationWatcher(
                service, List.of("Demo:Sentinel"), PRODUCTION, POLLING_INTERVAL),
            watcher -> {
                watcher.start();
                System.out.println(
                    "Watching Demo:Sentinel for " + WATCH_DURATION.toSeconds() + " seconds...");
                return Mono.delay(WATCH_DURATION).then();
            },
            AsyncConfigurationWatcher::close));
    }

    private static TokenCredential managedIdentityCredential() {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            builder.clientId(clientId);
        }
        return builder.build();
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable");
        }
        return value;
    }

    private static void sleep(Duration duration) {
        try {
            Thread.sleep(duration.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Demo interrupted", exception);
        }
    }

    private static void print(String name, Object value) {
        System.out.printf("%-35s %s%n", name + ":", value);
    }
}
