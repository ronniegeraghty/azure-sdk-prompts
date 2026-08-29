package com.example.keyvaultconfig;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.LinkedHashMap;
import java.util.Map;

public final class Main {
    private static final Duration EXPIRY_WARNING_WINDOW = Duration.ofDays(7);
    private static final Duration EXPIRY_CHECK_INTERVAL = Duration.ofHours(1);
    private static final Duration PURGE_TIMEOUT = Duration.ofMinutes(5);
    private static final Duration PURGE_POLL_INTERVAL = Duration.ofSeconds(2);
    private static final String ROTATING_SECRET = "demo-rotating-secret";

    private Main() {
    }

    public static void main(String[] args) {
        KeyVaultConfiguration configuration = KeyVaultConfiguration.fromEnvironment();
        Map<String, String> requiredSecrets = requiredSecrets();
        String rotatedValue = requiredEnvironmentVariable("ROTATED_SECRET_VALUE");
        OffsetDateTime rotatedExpiry = OffsetDateTime.now().plusDays(90);

        runSyncDemo(configuration, requiredSecrets, rotatedValue, rotatedExpiry);
        runAsyncDemo(configuration, requiredSecrets, rotatedValue, rotatedExpiry);
    }

    private static void runSyncDemo(
            KeyVaultConfiguration configuration,
            Map<String, String> requiredSecrets,
            String rotatedValue,
            OffsetDateTime rotatedExpiry) {
        System.out.println("Running synchronous demo");
        try (CachingSecretProvider cache = new CachingSecretProvider(
                configuration.secretProvider(),
                EXPIRY_WARNING_WINDOW,
                error -> System.err.println("Synchronous cache refresh failed: " + error.getMessage()))) {
            cache.loadRequired(requiredSecrets);
            cache.startAutomaticRefresh(EXPIRY_CHECK_INTERVAL);
            printCachedNames(cache, requiredSecrets);
            cache.refresh("database-password");
            printExpiryWarnings(cache.secretsNearExpiry());
            cache.refreshExpiringSecrets();

            new SecretRotationHelper(
                    configuration.secretClient(),
                    PURGE_TIMEOUT,
                    PURGE_POLL_INTERVAL)
                    .rotate(ROTATING_SECRET, rotatedValue, rotatedExpiry);
        }
    }

    private static void runAsyncDemo(
            KeyVaultConfiguration configuration,
            Map<String, String> requiredSecrets,
            String rotatedValue,
            OffsetDateTime rotatedExpiry) {
        System.out.println("Running asynchronous demo");
        try (AsyncCachingSecretProvider cache = new AsyncCachingSecretProvider(
                configuration.asyncSecretProvider(),
                EXPIRY_WARNING_WINDOW,
                error -> System.err.println("Asynchronous cache refresh failed: " + error.getMessage()))) {
            cache.loadRequired(requiredSecrets).block();
            cache.startAutomaticRefresh(EXPIRY_CHECK_INTERVAL);
            printCachedNames(cache, requiredSecrets);
            cache.refresh("database-password").block();
            printExpiryWarnings(cache.secretsNearExpiry());
            cache.refreshExpiringSecrets().block();

            new AsyncSecretRotationHelper(
                    configuration.secretAsyncClient(),
                    PURGE_TIMEOUT,
                    PURGE_POLL_INTERVAL)
                    .rotate(ROTATING_SECRET, rotatedValue, rotatedExpiry)
                    .block();
        }
    }

    private static Map<String, String> requiredSecrets() {
        Map<String, String> required = new LinkedHashMap<>();
        required.put("database-host", "localhost");
        required.put("database-password", "development-only");
        required.put("service-api-key", "not-configured");
        required.put(ROTATING_SECRET, "not-configured");
        return required;
    }

    private static void printCachedNames(
            CachingSecretProvider cache,
            Map<String, String> requiredSecrets) {
        requiredSecrets.keySet().forEach(name -> cache.getCached(name)
                .ifPresent(secret -> System.out.println("Cached secret: " + secret.name())));
    }

    private static void printCachedNames(
            AsyncCachingSecretProvider cache,
            Map<String, String> requiredSecrets) {
        requiredSecrets.keySet().forEach(name -> cache.getCached(name)
                .ifPresent(secret -> System.out.println("Cached secret: " + secret.name())));
    }

    private static void printExpiryWarnings(Iterable<SecretValue> secrets) {
        secrets.forEach(secret -> System.out.printf(
                "WARNING: secret '%s' expires at %s%n",
                secret.name(),
                secret.expiresOn()));
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}
