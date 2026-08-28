package com.example.blobmanager;

import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.core.http.policy.HttpLoggingPolicy;
import com.azure.core.http.policy.TimeoutPolicy;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.common.policy.RequestRetryOptions;
import com.azure.storage.common.policy.RetryPolicyType;

import java.time.Duration;
import java.util.Locale;
import java.util.Map;

public final class AzureBlobStorageConfiguration {
    public static final String ENDPOINT_ENV = "AZURE_STORAGE_BLOB_ENDPOINT";

    private final String endpoint;
    private final int maxRetries;
    private final Duration retryDelay;
    private final Duration maxRetryDelay;
    private final Duration requestTimeout;
    private final HttpLogDetailLevel logLevel;

    public AzureBlobStorageConfiguration(
            String endpoint,
            int maxRetries,
            Duration retryDelay,
            Duration maxRetryDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel) {
        if (endpoint == null || endpoint.isBlank()) {
            throw new IllegalArgumentException("Storage account endpoint must not be blank");
        }
        if (maxRetries < 0) {
            throw new IllegalArgumentException("maxRetries must be non-negative");
        }
        this.endpoint = endpoint;
        this.maxRetries = maxRetries;
        this.retryDelay = requirePositive(retryDelay, "retryDelay");
        this.maxRetryDelay = requirePositive(maxRetryDelay, "maxRetryDelay");
        this.requestTimeout = requirePositive(requestTimeout, "requestTimeout");
        this.logLevel = logLevel;
    }

    public static AzureBlobStorageConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static AzureBlobStorageConfiguration fromEnvironment(Map<String, String> environment) {
        String endpoint = required(environment, ENDPOINT_ENV);
        int maxRetries = parseNonNegativeInt(environment, "AZURE_STORAGE_MAX_RETRIES", 5);
        Duration retryDelay = Duration.ofMillis(
                parsePositiveLong(environment, "AZURE_STORAGE_RETRY_DELAY_MS", 800));
        Duration maxRetryDelay = Duration.ofMillis(
                parsePositiveLong(environment, "AZURE_STORAGE_MAX_RETRY_DELAY_MS", 10_000));
        Duration requestTimeout = Duration.ofSeconds(
                parsePositiveLong(environment, "AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120));
        HttpLogDetailLevel logLevel = parseLogLevel(
                environment.getOrDefault("AZURE_STORAGE_HTTP_LOG_LEVEL", "BASIC"));

        return new AzureBlobStorageConfiguration(
                endpoint, maxRetries, retryDelay, maxRetryDelay, requestTimeout, logLevel);
    }

    public Clients createClients() {
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        RequestRetryOptions retryOptions = new RequestRetryOptions(
                RetryPolicyType.EXPONENTIAL,
                maxRetries + 1,
                null,
                retryDelay.toMillis(),
                maxRetryDelay.toMillis(),
                null);
        HttpLogOptions logOptions = new HttpLogOptions().setLogLevel(logLevel);

        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .retryOptions(retryOptions)
                .addPolicy(new TimeoutPolicy(requestTimeout))
                .addPolicy(new HttpLoggingPolicy(logOptions));

        return new Clients(builder.buildClient(), builder.buildAsyncClient());
    }

    public Duration requestTimeout() {
        return requestTimeout;
    }

    public record Clients(BlobServiceClient syncClient, BlobServiceAsyncClient asyncClient) {
    }

    private static Duration requirePositive(Duration value, String name) {
        if (value == null || value.isZero() || value.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return value;
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable " + name + " is not set");
        }
        return value;
    }

    private static int parseNonNegativeInt(Map<String, String> environment, String name, int defaultValue) {
        int value = Integer.parseInt(environment.getOrDefault(name, Integer.toString(defaultValue)));
        if (value < 0) {
            throw new IllegalArgumentException(name + " must be non-negative");
        }
        return value;
    }

    private static long parsePositiveLong(Map<String, String> environment, String name, long defaultValue) {
        long value = Long.parseLong(environment.getOrDefault(name, Long.toString(defaultValue)));
        if (value <= 0) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return value;
    }

    private static HttpLogDetailLevel parseLogLevel(String value) {
        try {
            return HttpLogDetailLevel.valueOf(value.toUpperCase(Locale.ROOT));
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException(
                    "AZURE_STORAGE_HTTP_LOG_LEVEL must be one of NONE, BASIC, HEADERS, BODY, BODY_AND_HEADERS",
                    exception);
        }
    }
}
