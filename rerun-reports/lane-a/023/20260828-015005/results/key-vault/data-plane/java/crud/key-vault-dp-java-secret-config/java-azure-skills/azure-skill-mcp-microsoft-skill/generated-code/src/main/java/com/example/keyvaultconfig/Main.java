package com.example.keyvaultconfig;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public final class Main {
    private static final List<String> REQUIRED_KEYS =
            List.of("database-url", "api-key", "feature-flag");
    private static final Map<String, String> DEFAULTS = Map.of(
            "database-url", "jdbc:h2:mem:fallback",
            "api-key", "not-configured",
            "feature-flag", "false");
    private static final Duration EXPIRY_WARNING_WINDOW = Duration.ofDays(7);
    private static final String ROTATION_SECRET = "rotating-demo-secret";

    private Main() {
    }

    public static void main(String[] args) {
        KeyVaultClientFactory.Clients clients = KeyVaultClientFactory.fromEnvironment();
        runSyncDemo(clients);
        runAsyncDemo(clients);
    }

    private static void runSyncDemo(KeyVaultClientFactory.Clients clients) {
        System.out.println("=== Synchronous provider ===");
        SyncSecretProvider provider = new SyncSecretProvider(clients.sync());
        SyncSecretCache cache =
                new SyncSecretCache(provider, DEFAULTS, EXPIRY_WARNING_WINDOW);

        cache.loadRequired(REQUIRED_KEYS);
        REQUIRED_KEYS.forEach(name ->
                System.out.printf("%s loaded from cache (%d characters)%n",
                        name, cache.get(name).length()));
        cache.refresh("api-key");
        printExpiryWarnings(cache.secretsNearExpiry());

        SyncSecretRotationHelper rotation = new SyncSecretRotationHelper(clients.sync());
        rotation.rotate(
                ROTATION_SECRET,
                "sync-" + UUID.randomUUID(),
                OffsetDateTime.now(ZoneOffset.UTC).plusDays(90));
        cache.refresh(ROTATION_SECRET);
        System.out.println("Synchronous rotation completed.");
    }

    private static void runAsyncDemo(KeyVaultClientFactory.Clients clients) {
        System.out.println("=== Asynchronous provider ===");
        AsyncSecretProvider provider = new AsyncSecretProvider(clients.async());
        AsyncSecretCache cache =
                new AsyncSecretCache(provider, DEFAULTS, EXPIRY_WARNING_WINDOW);

        cache.loadRequired(REQUIRED_KEYS)
                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_KEYS))
                .concatMap(name -> cache.get(name)
                        .doOnNext(value -> System.out.printf(
                                "%s loaded from cache (%d characters)%n",
                                name,
                                value.length())))
                .then(cache.refresh("api-key"))
                .then(cache.secretsNearExpiry())
                .doOnNext(Main::printExpiryWarnings)
                .then(new AsyncSecretRotationHelper(clients.async()).rotate(
                        ROTATION_SECRET,
                        "async-" + UUID.randomUUID(),
                        OffsetDateTime.now(ZoneOffset.UTC).plusDays(90)))
                .then(cache.refresh(ROTATION_SECRET))
                .doOnSuccess(ignored -> System.out.println("Asynchronous rotation completed."))
                .block();
    }

    private static void printExpiryWarnings(List<SecretValue> secrets) {
        secrets.forEach(secret -> System.out.printf(
                "WARNING: %s version %s expires at %s%n",
                secret.name(),
                secret.version(),
                secret.expiresOn()));
    }
}
