package com.example.keyvaultconfig;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Map;
import java.util.UUID;

public final class Main {
    private static final Duration EXPIRY_WARNING_WINDOW = Duration.ofDays(7);
    private static final Map<String, String> REQUIRED_CONFIG = Map.of(
            "database-url", "jdbc:postgresql://localhost/app",
            "external-api-key", "development-only-default",
            "feature-flag", "false");

    private Main() {
    }

    public static void main(String[] args) {
        KeyVaultClientFactory.Clients clients = KeyVaultClientFactory.fromEnvironment();
        SecretRotationHelper rotationHelper = new SecretRotationHelper(
                clients.syncClient(),
                clients.asyncClient(),
                Duration.ofMinutes(5),
                Duration.ofSeconds(2));

        runSyncDemo(clients, rotationHelper);
        runAsyncDemo(clients, rotationHelper);
    }

    private static void runSyncDemo(
            KeyVaultClientFactory.Clients clients,
            SecretRotationHelper rotationHelper) {
        System.out.println("=== Synchronous provider ===");
        KeyVaultSecretProvider provider = new KeyVaultSecretProvider(clients.syncClient());
        CachingSecretProvider cache =
                new CachingSecretProvider(provider, EXPIRY_WARNING_WINDOW);

        cache.loadRequired(REQUIRED_CONFIG);
        REQUIRED_CONFIG.keySet().forEach(name -> print(cache.get(name)));

        System.out.println("Refreshing feature-flag");
        print(cache.refresh("feature-flag"));

        cache.refreshExpiringSecrets();
        warnNearExpiry(cache.expiringSecrets());

        rotationHelper.rotate(
                rotationSecretName(),
                "sync-" + UUID.randomUUID(),
                OffsetDateTime.now().plusDays(30));
        System.out.println("Synchronous rotation completed");
    }

    private static void runAsyncDemo(
            KeyVaultClientFactory.Clients clients,
            SecretRotationHelper rotationHelper) {
        System.out.println("=== Asynchronous provider ===");
        AsyncKeyVaultSecretProvider provider =
                new AsyncKeyVaultSecretProvider(clients.asyncClient());
        AsyncCachingSecretProvider cache =
                new AsyncCachingSecretProvider(provider, EXPIRY_WARNING_WINDOW);

        cache.loadRequired(REQUIRED_CONFIG)
                .thenMany(reactor.core.publisher.Flux.fromIterable(REQUIRED_CONFIG.keySet()))
                .concatMap(cache::get)
                .doOnNext(Main::print)
                .then(cache.refresh("feature-flag"))
                .doOnNext(secret -> System.out.println("Refreshed feature-flag"))
                .doOnNext(Main::print)
                .thenMany(cache.refreshExpiringSecrets())
                .then(MonoSupport.fromRunnable(() -> warnNearExpiry(cache.expiringSecrets())))
                .then(rotationHelper.rotateAsync(
                        rotationSecretName(),
                        "async-" + UUID.randomUUID(),
                        OffsetDateTime.now().plusDays(30)))
                .doOnSuccess(secret -> System.out.println("Asynchronous rotation completed"))
                .block();
    }

    private static String rotationSecretName() {
        return System.getenv().getOrDefault("ROTATION_SECRET_NAME", "rotating-demo-secret");
    }

    private static void print(SecretSnapshot secret) {
        System.out.printf(
                "%s loaded (valueLength=%d, version=%s, expires=%s, source=%s)%n",
                secret.name(),
                secret.value().length(),
                secret.version(),
                secret.expiresOn(),
                secret.found() ? "Key Vault" : "default");
    }

    private static void warnNearExpiry(Iterable<SecretSnapshot> secrets) {
        secrets.forEach(secret -> System.out.printf(
                "WARNING: secret '%s' expires on %s%n",
                secret.name(),
                secret.expiresOn()));
    }

    private static final class MonoSupport {
        private MonoSupport() {
        }

        private static reactor.core.publisher.Mono<Void> fromRunnable(Runnable runnable) {
            return reactor.core.publisher.Mono.fromRunnable(runnable);
        }
    }
}
