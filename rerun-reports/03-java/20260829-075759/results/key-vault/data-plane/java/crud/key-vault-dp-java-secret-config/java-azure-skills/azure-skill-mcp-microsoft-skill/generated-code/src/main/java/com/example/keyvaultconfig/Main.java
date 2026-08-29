package com.example.keyvaultconfig;

import com.azure.security.keyvault.secrets.SecretAsyncClient;
import com.azure.security.keyvault.secrets.SecretClient;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final List<String> REQUIRED_KEYS =
            List.of("database-url", "service-api-key", "feature-mode");
    private static final Map<String, String> DEFAULTS = Map.of(
            "database-url", "jdbc:h2:mem:local",
            "service-api-key", "not-configured",
            "feature-mode", "safe");
    private static final Duration WARNING_WINDOW = Duration.ofDays(7);
    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);
    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(2);
    private static final String ROTATION_SECRET = "rotating-demo-secret";

    private Main() {
    }

    public static void main(String[] args) {
        String rotatedValue = requireEnvironmentVariable("DEMO_ROTATED_SECRET_VALUE");

        runSyncDemo(rotatedValue);
        runAsyncDemo(rotatedValue + "-async");
    }

    private static void runSyncDemo(String rotatedValue) {
        System.out.println("=== Synchronous implementation ===");
        SecretClient client = KeyVaultClientFactory.createSyncClient();
        SyncSecretCache cache = new SyncSecretCache(
                new SyncSecretProvider(client), DEFAULTS, WARNING_WINDOW);

        cache.loadRequired(REQUIRED_KEYS);
        REQUIRED_KEYS.forEach(name ->
                System.out.printf("%s loaded (%d characters)%n", name, cache.get(name).length()));

        cache.refresh("service-api-key");
        cache.refreshExpiring();
        printExpiryWarnings(cache.snapshot(), cache::isNearExpiry);

        SyncSecretRotator rotator =
                new SyncSecretRotator(client, ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT);
        rotator.rotate(ROTATION_SECRET, rotatedValue, OffsetDateTime.now().plusDays(90));
        System.out.println("Synchronous rotation completed.");
    }

    private static void runAsyncDemo(String rotatedValue) {
        System.out.println("=== Asynchronous implementation ===");
        SecretAsyncClient client = KeyVaultClientFactory.createAsyncClient();
        AsyncSecretCache cache = new AsyncSecretCache(
                new AsyncSecretProvider(client), DEFAULTS, WARNING_WINDOW);

        cache.loadRequired(REQUIRED_KEYS)
                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_KEYS)
                        .concatMap(name -> cache.get(name)
                                .doOnNext(value ->
                                        System.out.printf(
                                                "%s loaded (%d characters)%n",
                                                name,
                                                value.length()))))
                .then(cache.refresh("service-api-key"))
                .then(cache.refreshExpiring())
                .doOnNext(snapshot ->
                        printExpiryWarnings(snapshot, cache::isNearExpiry))
                .then(new AsyncSecretRotator(
                        client, ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)
                        .rotate(
                                ROTATION_SECRET,
                                rotatedValue,
                                OffsetDateTime.now().plusDays(90)))
                .doOnSuccess(secret ->
                        System.out.println("Asynchronous rotation completed."))
                .block();
    }

    private static void printExpiryWarnings(
            Map<String, SecretValue> secrets,
            java.util.function.Predicate<SecretValue> nearExpiry) {
        secrets.values().stream()
                .filter(nearExpiry)
                .forEach(secret -> System.out.printf(
                        "WARNING: %s expires at %s%n",
                        secret.name(),
                        secret.expiresOn()));
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Required environment variable " + name + " is not set");
        }
        return value;
    }
}
