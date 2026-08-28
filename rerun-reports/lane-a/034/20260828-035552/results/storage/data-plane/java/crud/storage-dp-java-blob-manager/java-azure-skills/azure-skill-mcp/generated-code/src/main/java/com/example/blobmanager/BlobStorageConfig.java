package com.example.blobmanager;

import com.azure.core.credential.TokenCredential;
import com.azure.core.http.HttpClient;
import com.azure.core.http.netty.NettyAsyncHttpClientBuilder;
import com.azure.core.http.policy.ExponentialBackoffOptions;
import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.core.http.policy.RetryOptions;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.time.Duration;
import java.util.Locale;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageConfig {
    public static final String ENDPOINT_ENV = "AZURE_STORAGE_ACCOUNT_ENDPOINT";

    private final Settings settings;

    public BlobStorageConfig(Settings settings) {
        this.settings = Objects.requireNonNull(settings, "settings");
    }

    public static BlobStorageConfig fromEnvironment() {
        return new BlobStorageConfig(Settings.fromEnvironment(System.getenv()));
    }

    public Clients createClients() {
        TokenCredential credential = managedIdentityCredential(settings.managedIdentityClientId());
        HttpClient httpClient = new NettyAsyncHttpClientBuilder()
                .connectTimeout(settings.requestTimeout())
                .responseTimeout(settings.requestTimeout())
                .readTimeout(settings.requestTimeout())
                .writeTimeout(settings.requestTimeout())
                .build();

        RetryOptions retryOptions = new RetryOptions(new ExponentialBackoffOptions()
                .setMaxRetries(settings.maxRetries())
                .setBaseDelay(settings.retryDelay())
                .setMaxDelay(settings.maxRetryDelay()));

        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()
                .endpoint(settings.endpoint())
                .credential(credential)
                .httpClient(httpClient)
                .retryOptions(retryOptions)
                .httpLogOptions(new HttpLogOptions().setLogLevel(settings.logLevel()));

        return new Clients(builder.buildClient(), builder.buildAsyncClient());
    }

    private static TokenCredential managedIdentityCredential(String clientId) {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        if (clientId != null && !clientId.isBlank()) {
            builder.clientId(clientId);
        }
        return builder.build();
    }

    public record Clients(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
    }

    public record Settings(
            String endpoint,
            String managedIdentityClientId,
            int maxRetries,
            Duration retryDelay,
            Duration maxRetryDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel
    ) {
        public Settings {
            if (endpoint == null || endpoint.isBlank()) {
                throw new IllegalArgumentException(ENDPOINT_ENV + " must be set");
            }
            if (!endpoint.startsWith("https://")) {
                throw new IllegalArgumentException(ENDPOINT_ENV + " must use HTTPS");
            }
            if (maxRetries < 0) {
                throw new IllegalArgumentException("maxRetries must be non-negative");
            }
            Objects.requireNonNull(retryDelay, "retryDelay");
            Objects.requireNonNull(maxRetryDelay, "maxRetryDelay");
            Objects.requireNonNull(requestTimeout, "requestTimeout");
            Objects.requireNonNull(logLevel, "logLevel");
        }

        static Settings fromEnvironment(Map<String, String> environment) {
            return new Settings(
                    required(environment, ENDPOINT_ENV),
                    environment.get("AZURE_CLIENT_ID"),
                    integer(environment, "AZURE_STORAGE_MAX_RETRIES", 5),
                    duration(environment, "AZURE_STORAGE_RETRY_DELAY_SECONDS", 2),
                    duration(environment, "AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS", 30),
                    duration(environment, "AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120),
                    logLevel(environment.getOrDefault("AZURE_STORAGE_HTTP_LOG_LEVEL", "BASIC"))
            );
        }

        private static String required(Map<String, String> environment, String name) {
            String value = environment.get(name);
            if (value == null || value.isBlank()) {
                throw new IllegalArgumentException(name + " must be set");
            }
            return value;
        }

        private static int integer(Map<String, String> environment, String name, int defaultValue) {
            String value = environment.get(name);
            return value == null ? defaultValue : Integer.parseInt(value);
        }

        private static Duration duration(
                Map<String, String> environment,
                String name,
                int defaultSeconds
        ) {
            return Duration.ofSeconds(integer(environment, name, defaultSeconds));
        }

        private static HttpLogDetailLevel logLevel(String value) {
            return HttpLogDetailLevel.valueOf(value.toUpperCase(Locale.ROOT));
        }
    }
}
