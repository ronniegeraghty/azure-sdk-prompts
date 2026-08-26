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
import com.azure.storage.blob.models.ParallelTransferOptions;

import java.time.Duration;
import java.util.Locale;
import java.util.Map;
import java.util.Objects;

public final class BlobStorageConfiguration {
    private static final int MEBIBYTE = 1024 * 1024;

    private final String endpoint;
    private final int maxRetries;
    private final Duration retryBaseDelay;
    private final Duration retryMaxDelay;
    private final Duration requestTimeout;
    private final HttpLogDetailLevel logLevel;
    private final long blockSize;
    private final int maxConcurrency;
    private final String managedIdentityClientId;

    private BlobStorageConfiguration(
            String endpoint,
            int maxRetries,
            Duration retryBaseDelay,
            Duration retryMaxDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel,
            long blockSize,
            int maxConcurrency,
            String managedIdentityClientId) {
        this.endpoint = endpoint;
        this.maxRetries = maxRetries;
        this.retryBaseDelay = retryBaseDelay;
        this.retryMaxDelay = retryMaxDelay;
        this.requestTimeout = requestTimeout;
        this.logLevel = logLevel;
        this.blockSize = blockSize;
        this.maxConcurrency = maxConcurrency;
        this.managedIdentityClientId = managedIdentityClientId;
    }

    public static BlobStorageConfiguration fromEnvironment() {
        return fromEnvironment(System.getenv());
    }

    static BlobStorageConfiguration fromEnvironment(Map<String, String> environment) {
        Objects.requireNonNull(environment, "environment");

        String endpoint = required(environment, "AZURE_STORAGE_ACCOUNT_ENDPOINT");
        if (!endpoint.startsWith("https://")) {
            throw new IllegalArgumentException("AZURE_STORAGE_ACCOUNT_ENDPOINT must use HTTPS");
        }

        int maxRetries = positiveInt(environment, "BLOB_MAX_RETRIES", 5);
        int baseDelaySeconds = positiveInt(environment, "BLOB_RETRY_BASE_DELAY_SECONDS", 1);
        int maxDelaySeconds = positiveInt(environment, "BLOB_RETRY_MAX_DELAY_SECONDS", 30);
        int requestTimeoutSeconds = positiveInt(environment, "BLOB_REQUEST_TIMEOUT_SECONDS", 60);
        int blockSizeMiB = positiveInt(environment, "BLOB_BLOCK_SIZE_MIB", 8);
        int maxConcurrency = positiveInt(environment, "BLOB_MAX_CONCURRENCY", 4);

        if (maxDelaySeconds < baseDelaySeconds) {
            throw new IllegalArgumentException(
                    "BLOB_RETRY_MAX_DELAY_SECONDS must be at least BLOB_RETRY_BASE_DELAY_SECONDS");
        }

        String configuredLogLevel = environment.getOrDefault("BLOB_HTTP_LOG_LEVEL", "BASIC")
                .trim()
                .toUpperCase(Locale.ROOT);
        HttpLogDetailLevel logLevel;
        try {
            logLevel = HttpLogDetailLevel.valueOf(configuredLogLevel);
        } catch (IllegalArgumentException exception) {
            throw new IllegalArgumentException(
                    "BLOB_HTTP_LOG_LEVEL must be NONE, BASIC, HEADERS, or BODY_AND_HEADERS", exception);
        }

        return new BlobStorageConfiguration(
                endpoint,
                maxRetries,
                Duration.ofSeconds(baseDelaySeconds),
                Duration.ofSeconds(maxDelaySeconds),
                Duration.ofSeconds(requestTimeoutSeconds),
                logLevel,
                Math.multiplyExact((long) blockSizeMiB, MEBIBYTE),
                maxConcurrency,
                blankToNull(environment.get("AZURE_CLIENT_ID")));
    }

    public BlobServiceClient createSyncClient() {
        return clientBuilder().buildClient();
    }

    public BlobServiceAsyncClient createAsyncClient() {
        return clientBuilder().buildAsyncClient();
    }

    public ParallelTransferOptions createTransferOptions() {
        return new ParallelTransferOptions()
                .setBlockSizeLong(blockSize)
                .setMaxConcurrency(maxConcurrency);
    }

    private BlobServiceClientBuilder clientBuilder() {
        ExponentialBackoffOptions backoff = new ExponentialBackoffOptions()
                .setMaxRetries(maxRetries)
                .setBaseDelay(retryBaseDelay)
                .setMaxDelay(retryMaxDelay);

        HttpClient httpClient = new NettyAsyncHttpClientBuilder()
                .connectTimeout(requestTimeout)
                .responseTimeout(requestTimeout)
                .readTimeout(requestTimeout)
                .writeTimeout(requestTimeout)
                .build();

        return new BlobServiceClientBuilder()
                .endpoint(endpoint)
                .credential(createCredential())
                .retryOptions(new RetryOptions(backoff))
                .httpClient(httpClient)
                .httpLogOptions(new HttpLogOptions().setLogLevel(logLevel));
    }

    private TokenCredential createCredential() {
        ManagedIdentityCredentialBuilder builder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null) {
            builder.clientId(managedIdentityClientId);
        }
        return builder.build();
    }

    private static String required(Map<String, String> environment, String name) {
        String value = blankToNull(environment.get(name));
        if (value == null) {
            throw new IllegalArgumentException("Missing required environment variable: " + name);
        }
        return value;
    }

    private static int positiveInt(Map<String, String> environment, String name, int defaultValue) {
        String value = blankToNull(environment.get(name));
        if (value == null) {
            return defaultValue;
        }

        try {
            int parsed = Integer.parseInt(value);
            if (parsed <= 0) {
                throw new IllegalArgumentException(name + " must be greater than zero");
            }
            return parsed;
        } catch (NumberFormatException exception) {
            throw new IllegalArgumentException(name + " must be an integer", exception);
        }
    }

    private static String blankToNull(String value) {
        return value == null || value.isBlank() ? null : value.trim();
    }
}
