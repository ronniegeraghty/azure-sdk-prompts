package com.example.blobmanager;

import com.azure.core.http.HttpClient;
import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.core.util.ClientOptions;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.common.policy.RequestRetryOptions;
import com.azure.storage.common.policy.RetryPolicyType;
import com.azure.core.http.netty.NettyAsyncHttpClientBuilder;

import java.time.Duration;
import java.util.Locale;
import java.util.Objects;

public final class AzureBlobConfiguration {
    public static final String ENDPOINT_ENV = "AZURE_STORAGE_ENDPOINT";

    private final String endpoint;
    private final int maxRetries;
    private final Duration retryDelay;
    private final Duration maxRetryDelay;
    private final Duration requestTimeout;
    private final HttpLogDetailLevel logLevel;

    private AzureBlobConfiguration(
            String endpoint,
            int maxRetries,
            Duration retryDelay,
            Duration maxRetryDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel) {
        this.endpoint = Objects.requireNonNull(endpoint, "endpoint");
        this.maxRetries = maxRetries;
        this.retryDelay = retryDelay;
        this.maxRetryDelay = maxRetryDelay;
        this.requestTimeout = requestTimeout;
        this.logLevel = logLevel;
    }

    public static AzureBlobConfiguration fromEnvironment() {
        String endpoint = requiredEnvironmentVariable(ENDPOINT_ENV);
        if (!endpoint.startsWith("https://")) {
            throw new IllegalArgumentException(ENDPOINT_ENV + " must use HTTPS");
        }

        int maxRetries = integerEnvironmentVariable("AZURE_STORAGE_MAX_RETRIES", 5, 0);
        Duration retryDelay = durationEnvironmentVariable("AZURE_STORAGE_RETRY_DELAY_SECONDS", 2);
        Duration maxRetryDelay = durationEnvironmentVariable("AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS", 30);
        Duration requestTimeout = durationEnvironmentVariable("AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120);
        HttpLogDetailLevel logLevel = enumEnvironmentVariable(
                "AZURE_STORAGE_HTTP_LOG_LEVEL", HttpLogDetailLevel.class, HttpLogDetailLevel.BASIC);

        return new AzureBlobConfiguration(
                endpoint, maxRetries, retryDelay, maxRetryDelay, requestTimeout, logLevel);
    }

    public StorageClients createClients() {
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId.trim());
        }

        var credential = credentialBuilder.build();
        HttpClient httpClient = new NettyAsyncHttpClientBuilder()
                .connectTimeout(requestTimeout)
                .responseTimeout(requestTimeout)
                .readTimeout(requestTimeout)
                .writeTimeout(requestTimeout)
                .build();

        RequestRetryOptions retryOptions = new RequestRetryOptions(
                RetryPolicyType.EXPONENTIAL,
                maxRetries + 1,
                requestTimeout,
                retryDelay,
                maxRetryDelay,
                null);

        BlobServiceClientBuilder builder = new BlobServiceClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .httpClient(httpClient)
                .retryOptions(retryOptions)
                .clientOptions(new ClientOptions().setApplicationId("azure-blob-manager"))
                .httpLogOptions(new HttpLogOptions().setLogLevel(logLevel));

        return new StorageClients(builder.buildClient(), builder.buildAsyncClient(), requestTimeout);
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable " + name + " is not set");
        }
        return value.trim();
    }

    private static int integerEnvironmentVariable(String name, int defaultValue, int minimum) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            return defaultValue;
        }
        int parsed = Integer.parseInt(value);
        if (parsed < minimum) {
            throw new IllegalArgumentException(name + " must be at least " + minimum);
        }
        return parsed;
    }

    private static Duration durationEnvironmentVariable(String name, long defaultSeconds) {
        String value = System.getenv(name);
        long seconds = value == null || value.isBlank() ? defaultSeconds : Long.parseLong(value);
        if (seconds <= 0) {
            throw new IllegalArgumentException(name + " must be greater than zero");
        }
        return Duration.ofSeconds(seconds);
    }

    private static <E extends Enum<E>> E enumEnvironmentVariable(
            String name, Class<E> type, E defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank()
                ? defaultValue
                : Enum.valueOf(type, value.trim().toUpperCase(Locale.ROOT));
    }

    public record StorageClients(
            BlobServiceClient syncClient,
            BlobServiceAsyncClient asyncClient,
            Duration requestTimeout) {
    }
}
