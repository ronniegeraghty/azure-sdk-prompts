package com.example.keyvault;

import com.example.keyvault.KeyVaultClientFactory.KeyVaultClients;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final Duration WARNING_WINDOW = Duration.ofDays(7);
    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);
    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(2);
    private static final Map<String, String> REQUIRED_CONFIG = requiredConfig();

    private Main() {
    }

    public static void main(String[] args) {
        KeyVaultClients clients = KeyVaultClientFactory.fromEnvironment();
        runSyncDemo(clients);
        runAsyncDemo(clients).block();
    }

    private static void runSyncDemo(KeyVaultClients clients) {
        System.out.println("=== Synchronous implementation ===");
        SyncKeyVaultSecretProvider provider = new SyncKeyVaultSecretProvider(clients.syncClient());
        SyncSecretCache cache = new SyncSecretCache(provider, WARNING_WINDOW);

        cache.loadRequired(REQUIRED_CONFIG);
        REQUIRED_CONFIG.keySet().forEach(name -> printRead("sync", cache.get(name)));
        cache.refresh("api-base-url");
        cache.refreshExpiring();
        printExpiryWarnings(cache.secretsNearExpiry());

        String rotationName = requireEnvironment("DEMO_SYNC_ROTATION_SECRET_NAME");
        String rotationValue = requireEnvironment("DEMO_SYNC_ROTATION_NEW_VALUE");
        new SyncSecretRotator(clients.syncClient(), ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)
            .rotate(rotationName, rotationValue, OffsetDateTime.now(ZoneOffset.UTC).plusDays(90));
        System.out.printf("sync rotated '%s'%n", rotationName);
    }

    private static reactor.core.publisher.Mono<Void> runAsyncDemo(KeyVaultClients clients) {
        System.out.println("=== Asynchronous implementation ===");
        AsyncKeyVaultSecretProvider provider = new AsyncKeyVaultSecretProvider(clients.asyncClient());
        AsyncSecretCache cache = new AsyncSecretCache(provider, WARNING_WINDOW);
        String rotationName = requireEnvironment("DEMO_ASYNC_ROTATION_SECRET_NAME");
        String rotationValue = requireEnvironment("DEMO_ASYNC_ROTATION_NEW_VALUE");

        return cache.loadRequired(REQUIRED_CONFIG)
            .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_CONFIG.keySet()))
            .concatMap(cache::get)
            .doOnNext(secret -> printRead("async", secret))
            .then(cache.refresh("api-base-url"))
            .thenMany(cache.refreshExpiring())
            .then()
            .doOnSuccess(ignored -> printExpiryWarnings(cache.secretsNearExpiry()))
            .then(new AsyncSecretRotator(clients.asyncClient(), ROTATION_POLL_INTERVAL, ROTATION_TIMEOUT)
                .rotate(rotationName, rotationValue, OffsetDateTime.now(ZoneOffset.UTC).plusDays(90)))
            .doOnNext(secret -> System.out.printf("async rotated '%s'%n", secret.getName()))
            .then();
    }

    private static void printRead(String implementation, SecretSnapshot secret) {
        System.out.printf(
            "%s cache read: name=%s, version=%s, default=%s%n",
            implementation,
            secret.name(),
            secret.version() == null ? "<none>" : secret.version(),
            secret.defaultValue()
        );
    }

    private static void printExpiryWarnings(List<SecretSnapshot> expiringSecrets) {
        expiringSecrets.forEach(secret -> System.out.printf(
            "WARNING: secret '%s' expires on %s%n", secret.name(), secret.expiresOn()));
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static Map<String, String> requiredConfig() {
        Map<String, String> config = new LinkedHashMap<>();
        config.put("database-connection-string", "jdbc:postgresql://localhost/app");
        config.put("api-base-url", "https://localhost:8443");
        config.put("feature-flags", "{}");
        return Map.copyOf(config);
    }
}
