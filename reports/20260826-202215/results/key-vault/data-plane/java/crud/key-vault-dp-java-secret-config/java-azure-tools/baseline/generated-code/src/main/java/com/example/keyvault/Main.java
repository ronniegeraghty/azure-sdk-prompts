package com.example.keyvault;

import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final List<String> REQUIRED_KEYS =
            List.of("database-url", "api-key", "feature-flags");
    private static final Map<String, String> DEFAULTS = Map.of(
            "database-url", "jdbc:postgresql://localhost/example",
            "api-key", "not-configured",
            "feature-flags", "{}");
    private static final Duration WARNING_WINDOW = Duration.ofDays(7);
    private static final String ROTATION_SECRET = "demo-rotating-secret";

    private Main() {
    }

    public static void main(String[] args) {
        runSyncDemo();
        runAsyncDemo();
    }

    private static void runSyncDemo() {
        System.out.println("=== Synchronous implementation ===");
        SecretClient client = KeyVaultClients.syncClient();
        SyncSecretCache cache = new SyncSecretCache(
                new KeyVaultSyncSecretProvider(client), WARNING_WINDOW, DEFAULTS);

        cache.loadRequired(REQUIRED_KEYS);
        REQUIRED_KEYS.forEach(name -> {
            cache.get(name);
            System.out.printf("%s read from cache%n", name);
        });
        cache.refresh("api-key");
        printExpiryWarnings(cache.expiringSecrets());

        new SyncSecretRotator(client, Duration.ofSeconds(2), Duration.ofMinutes(2))
                .rotate(ROTATION_SECRET, "sync-rotated-value", OffsetDateTime.now().plusDays(90));
        System.out.println("Synchronous rotation complete.");
    }

    private static void runAsyncDemo() {
        System.out.println("=== Asynchronous implementation ===");
        SecretAsyncClient client = KeyVaultClients.asyncClient();
        AsyncSecretCache cache = new AsyncSecretCache(
                new KeyVaultAsyncSecretProvider(client), WARNING_WINDOW, DEFAULTS);

        cache.loadRequired(REQUIRED_KEYS)
                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_KEYS))
                .flatMap(name -> cache.get(name)
                        .doOnNext(value -> System.out.printf("%s read from cache%n", name)))
                .then(cache.refresh("api-key"))
                .doOnSuccess(ignored -> printExpiryWarnings(cache.expiringSecrets()))
                .then(new AsyncSecretRotator(client, Duration.ofSeconds(2), Duration.ofMinutes(2))
                        .rotate(ROTATION_SECRET, "async-rotated-value", OffsetDateTime.now().plusDays(90)))
                .doOnSuccess(ignored -> System.out.println("Asynchronous rotation complete."))
                .block();
    }

    private static void printExpiryWarnings(List<SecretSnapshot> secrets) {
        secrets.forEach(secret -> System.out.printf(
                "WARNING: %s expires at %s%n", secret.name(), secret.expiresOn()));
    }
}
