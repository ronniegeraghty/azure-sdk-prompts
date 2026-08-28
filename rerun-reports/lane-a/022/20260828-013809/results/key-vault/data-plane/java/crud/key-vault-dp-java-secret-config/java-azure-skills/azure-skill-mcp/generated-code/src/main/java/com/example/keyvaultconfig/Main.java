package com.example.keyvaultconfig;

import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public final class Main {
    private static final Duration WARNING_WINDOW = Duration.ofDays(7);
    private static final Duration ROTATION_TIMEOUT = Duration.ofMinutes(5);
    private static final Duration ROTATION_POLL_INTERVAL = Duration.ofSeconds(2);
    private static final String ROTATION_SECRET = "demo-rotating-secret";
    private static final Map<String, String> REQUIRED_CONFIG = Map.of(
            "database-connection-string", "not-configured",
            "external-api-key", "not-configured",
            "feature-flags", "{}");

    private Main() {
    }

    public static void main(String[] args) {
        KeyVaultClientFactory.Clients clients = KeyVaultClientFactory.fromEnvironment();
        runSyncDemo(clients);
        runAsyncDemo(clients).block();
    }

    private static void runSyncDemo(KeyVaultClientFactory.Clients clients) {
        System.out.println("=== Synchronous provider ===");
        KeyVaultSecretProvider provider = new KeyVaultSecretProvider(clients.syncClient());
        SecretCache cache = new SecretCache(provider, REQUIRED_CONFIG, WARNING_WINDOW);

        cache.loadRequired();
        REQUIRED_CONFIG.keySet().forEach(name -> printRead(cache.get(name)));

        printRead(cache.refresh("external-api-key"));
        warnAbout(cache.expiringSecrets());
        cache.refreshExpiringSecrets();

        ConfigSecret versioned = provider.getSecret(
                "external-api-key",
                System.getenv("DEMO_SECRET_VERSION"),
                "not-configured");
        System.out.printf("Versioned read: %s (version=%s, default=%s)%n",
                versioned.name(), versioned.version(), versioned.defaultValue());

        SecretRotationHelper rotation = new SecretRotationHelper(
                clients.syncClient(), ROTATION_TIMEOUT, ROTATION_POLL_INTERVAL);
        KeyVaultSecret rotated = rotation.rotate(
                ROTATION_SECRET,
                "sync-" + UUID.randomUUID(),
                OffsetDateTime.now(ZoneOffset.UTC).plusDays(30));
        System.out.printf("Rotated %s to version %s%n",
                rotated.getName(), rotated.getProperties().getVersion());
    }

    private static Mono<Void> runAsyncDemo(KeyVaultClientFactory.Clients clients) {
        System.out.println("=== Asynchronous provider ===");
        AsyncKeyVaultSecretProvider provider = new AsyncKeyVaultSecretProvider(clients.asyncClient());
        AsyncSecretCache cache = new AsyncSecretCache(provider, REQUIRED_CONFIG, WARNING_WINDOW);
        AsyncSecretRotationHelper rotation = new AsyncSecretRotationHelper(
                clients.asyncClient(), ROTATION_TIMEOUT, ROTATION_POLL_INTERVAL);

        return cache.loadRequired()
                .thenMany(Flux.fromIterable(REQUIRED_CONFIG.keySet())
                        .concatMap(cache::get)
                        .doOnNext(Main::printRead))
                .then(cache.refresh("external-api-key").doOnNext(Main::printRead))
                .then(Mono.fromRunnable(() -> warnAbout(cache.expiringSecrets())))
                .thenMany(cache.refreshExpiringSecrets())
                .then(provider.getSecret(
                        "external-api-key",
                        System.getenv("DEMO_SECRET_VERSION"),
                        "not-configured"))
                .doOnNext(secret -> System.out.printf(
                        "Versioned read: %s (version=%s, default=%s)%n",
                        secret.name(), secret.version(), secret.defaultValue()))
                .then(rotation.rotate(
                        ROTATION_SECRET,
                        "async-" + UUID.randomUUID(),
                        OffsetDateTime.now(ZoneOffset.UTC).plusDays(30)))
                .doOnNext(secret -> System.out.printf(
                        "Rotated %s to version %s%n",
                        secret.getName(), secret.getProperties().getVersion()))
                .then();
    }

    private static void printRead(ConfigSecret secret) {
        System.out.printf("Cache read: %s (version=%s, default=%s)%n",
                secret.name(), secret.version(), secret.defaultValue());
    }

    private static void warnAbout(List<ConfigSecret> secrets) {
        secrets.forEach(secret -> System.out.printf(
                "WARNING: %s expires at %s%n", secret.name(), secret.expiresOn()));
    }
}
