package com.example.blobmanager;

import com.azure.core.credential.TokenCredential;
import com.azure.core.http.HttpClient;
import com.azure.core.http.netty.NettyAsyncHttpClientBuilder;
import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.common.policy.RequestRetryOptions;
import com.azure.storage.common.policy.RetryPolicyType;

import java.time.Duration;
import java.util.Locale;
import java.util.Objects;

public final class BlobStorageConfiguration {
    private static final String ENDPOINT_ENV = "AZURE_STORAGE_ACCOUNT_ENDPOINT";

    private final String endpoint;
    private final String managedIdentityClientId;
    private final int maxRetries;
    private final Duration retryDelay;
    private final Duration maxRetryDelay;
    private final Duration requestTimeout;
    private final HttpLogDetailLevel logLevel;
    private final HttpClient httpClient;
    private final TokenCredential credential;

    public BlobStorageConfiguration(
            String endpoint,
            String managedIdentityClientId,
            int maxRetries,
            Duration retryDelay,
            Duration maxRetryDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel) {
        this.endpoint = requireHttpsEndpoint(endpoint);
        this.managedIdentityClientId = managedIdentityClientId;
        this.maxRetries = requireNonNegative(maxRetries, "maxRetries");
        this.retryDelay = requirePositive(retryDelay, "retryDelay");
        this.maxRetryDelay = requirePositive(maxRetryDelay, "maxRetryDelay");
        this.requestTimeout = requirePositive(requestTimeout, "requestTimeout");
        this.logLevel = Objects.requireNonNull(logLevel, "logLevel");
        if (retryDelay.compareTo(maxRetryDelay) > 0) {
            throw new IllegalArgumentException("retryDelay must not exceed maxRetryDelay");
        }

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        this.credential = credentialBuilder.build();
        this.httpClient = new NettyAsyncHttpClientBuilder()
                .connectTimeout(requestTimeout)
                .readTimeout(requestTimeout)
                .responseTimeout(requestTimeout)
                .writeTimeout(requestTimeout)
                .build();
    }

    public static BlobStorageConfiguration fromEnvironment() {
        return new BlobStorageConfiguration(
                requireEnvironment(ENDPOINT_ENV),
                System.getenv("AZURE_CLIENT_ID"),
                integerEnvironment("AZURE_STORAGE_MAX_RETRIES", 5),
                Duration.ofMillis(longEnvironment("AZURE_STORAGE_RETRY_DELAY_MS", 800)),
                Duration.ofMillis(longEnvironment("AZURE_STORAGE_MAX_RETRY_DELAY_MS", 10_000)),
                Duration.ofSeconds(longEnvironment("AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120)),
                logLevelEnvironment("AZURE_STORAGE_HTTP_LOG_LEVEL", HttpLogDetailLevel.BASIC));
    }

    public BlobServiceClient createSyncClient() {
        return clientBuilder().buildClient();
    }

    public BlobServiceAsyncClient createAsyncClient() {
        return clientBuilder().buildAsyncClient();
    }

    private BlobServiceClientBuilder clientBuilder() {
        RequestRetryOptions retryOptions = new RequestRetryOptions(
                RetryPolicyType.EXPONENTIAL,
                maxRetries + 1,
                Math.toIntExact(requestTimeout.toSeconds()),
                retryDelay.toMillis(),
                maxRetryDelay.toMillis(),
                null);

        HttpLogOptions logOptions = new HttpLogOptions().setLogLevel(logLevel);

        return new BlobServiceClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .httpClient(httpClient)
                .retryOptions(retryOptions)
                .httpLogOptions(logOptions);
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static int integerEnvironment(String name, int defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : Integer.parseInt(value);
    }

    private static long longEnvironment(String name, long defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : Long.parseLong(value);
    }

    private static HttpLogDetailLevel logLevelEnvironment(String name, HttpLogDetailLevel defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank()
                ? defaultValue
                : HttpLogDetailLevel.valueOf(value.trim().toUpperCase(Locale.ROOT));
    }

    private static String requireHttpsEndpoint(String endpoint) {
        Objects.requireNonNull(endpoint, "endpoint");
        if (!endpoint.startsWith("https://")) {
            throw new IllegalArgumentException(ENDPOINT_ENV + " must use HTTPS");
        }
        return endpoint;
    }

    private static int requireNonNegative(int value, String name) {
        if (value < 0) {
            throw new IllegalArgumentException(name + " must not be negative");
        }
        return value;
    }

    private static Duration requirePositive(Duration value, String name) {
        Objects.requireNonNull(value, name);
        if (value.isZero() || value.isNegative()) {
            throw new IllegalArgumentException(name + " must be positive");
        }
        return value;
    }
}
