package com.example.blobmanager;

import com.azure.core.credential.TokenCredential;
import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.common.policy.RequestRetryOptions;
import com.azure.storage.common.policy.RetryPolicyType;

import java.net.URI;
import java.time.Duration;
import java.util.Locale;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageConfiguration {
    public static final String ENDPOINT_ENV = "AZURE_STORAGE_ACCOUNT_URL";

    private BlobStorageConfiguration() {
    }

    public static BlobStorageClients fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static BlobStorageClients fromEnvironment(Map<String, String> environment) {
        Objects.requireNonNull(environment, "environment");

        String endpoint = required(environment, ENDPOINT_ENV);
        validateEndpoint(endpoint);

        int maxRetries = nonNegativeInt(environment, "AZURE_STORAGE_MAX_RETRIES", 5);
        int retryDelayMs = positiveInt(environment, "AZURE_STORAGE_RETRY_DELAY_MS", 800);
        int maxRetryDelayMs = positiveInt(environment, "AZURE_STORAGE_MAX_RETRY_DELAY_MS", 10_000);
        int requestTimeoutSeconds = positiveInt(environment, "AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120);
        HttpLogDetailLevel logLevel = logLevel(environment.getOrDefault(
            "AZURE_STORAGE_HTTP_LOG_LEVEL", "BASIC"));

        TokenCredential credential = managedIdentityCredential(environment);
        BlobServiceClient syncClient = builder(
            endpoint, credential, maxRetries, retryDelayMs, maxRetryDelayMs,
            requestTimeoutSeconds, logLevel).buildClient();
        BlobServiceAsyncClient asyncClient = builder(
            endpoint, credential, maxRetries, retryDelayMs, maxRetryDelayMs,
            requestTimeoutSeconds, logLevel).buildAsyncClient();

        return new BlobStorageClients(
            syncClient,
            asyncClient,
            Duration.ofSeconds(requestTimeoutSeconds));
    }

    private static BlobServiceClientBuilder builder(
        String endpoint,
        TokenCredential credential,
        int maxRetries,
        int retryDelayMs,
        int maxRetryDelayMs,
        int requestTimeoutSeconds,
        HttpLogDetailLevel logLevel
    ) {
        RequestRetryOptions retryOptions = new RequestRetryOptions(
            RetryPolicyType.EXPONENTIAL,
            maxRetries + 1,
            requestTimeoutSeconds,
            (long) retryDelayMs,
            (long) maxRetryDelayMs,
            null);

        return new BlobServiceClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .retryOptions(retryOptions)
            .httpLogOptions(new HttpLogOptions().setLogLevel(logLevel));
    }

    private static TokenCredential managedIdentityCredential(Map<String, String> environment) {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        String clientId = environment.get("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            builder.clientId(clientId);
        }
        return builder.build();
    }

    private static void validateEndpoint(String endpoint) {
        URI uri;
        try {
            uri = URI.create(endpoint);
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException(ENDPOINT_ENV + " must be a valid URI", exception);
        }
        if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null
            || uri.getRawQuery() != null || uri.getUserInfo() != null) {
            throw new IllegalArgumentException(
                ENDPOINT_ENV + " must be an HTTPS storage endpoint without credentials or query parameters");
        }
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static int positiveInt(Map<String, String> environment, String name, int defaultValue) {
        int value = integer(environment, name, defaultValue);
        if (value <= 0) {
            throw new IllegalArgumentException(name + " must be greater than zero");
        }
        return value;
    }

    private static int nonNegativeInt(Map<String, String> environment, String name, int defaultValue) {
        int value = integer(environment, name, defaultValue);
        if (value < 0) {
            throw new IllegalArgumentException(name + " must be zero or greater");
        }
        return value;
    }

    private static int integer(Map<String, String> environment, String name, int defaultValue) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            return defaultValue;
        }
        try {
            return Integer.parseInt(value);
        } catch (NumberFormatException exception) {
            throw new IllegalArgumentException(name + " must be an integer", exception);
        }
    }

    private static HttpLogDetailLevel logLevel(String value) {
        try {
            return HttpLogDetailLevel.valueOf(value.toUpperCase(Locale.ROOT));
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException(
                "AZURE_STORAGE_HTTP_LOG_LEVEL must be one of NONE, BASIC, HEADERS, or BODY_AND_HEADERS",
                exception);
        }
    }

    public record BlobStorageClients(
        BlobServiceClient syncClient,
        BlobServiceAsyncClient asyncClient,
        Duration requestTimeout
    ) {
        public BlobStorageClients {
            Objects.requireNonNull(syncClient, "syncClient");
            Objects.requireNonNull(asyncClient, "asyncClient");
            Objects.requireNonNull(requestTimeout, "requestTimeout");
        }
    }
}
