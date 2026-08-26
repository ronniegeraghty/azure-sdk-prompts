package com.example.blob;

import com.azure.core.http.HttpClient;
import com.azure.core.http.netty.NettyAsyncHttpClientBuilder;
import com.azure.core.http.policy.HttpLogDetailLevel;
import com.azure.core.http.policy.HttpLogOptions;
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
import java.util.Objects;

public final class AzureBlobConfiguration {
    private final Settings settings;

    public AzureBlobConfiguration(Settings settings) {
        this.settings = Objects.requireNonNull(settings, "settings");
    }

    public static AzureBlobConfiguration fromEnvironment() {
        return new AzureBlobConfiguration(Settings.from(System.getenv()));
    }

    public BlobServiceClient createSyncClient() {
        return baseBuilder().buildClient();
    }

    public BlobServiceAsyncClient createAsyncClient() {
        return baseBuilder().buildAsyncClient();
    }

    public Settings settings() {
        return settings;
    }

    private BlobServiceClientBuilder baseBuilder() {
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        HttpClient httpClient = new NettyAsyncHttpClientBuilder()
                .connectTimeout(settings.requestTimeout())
                .responseTimeout(settings.requestTimeout())
                .readTimeout(settings.requestTimeout())
                .writeTimeout(settings.requestTimeout())
                .build();

        RequestRetryOptions retryOptions = new RequestRetryOptions(
                RetryPolicyType.EXPONENTIAL,
                Math.addExact(settings.maxRetries(), 1),
                settings.requestTimeout(),
                settings.retryDelay(),
                settings.maxRetryDelay(),
                null);

        return new BlobServiceClientBuilder()
                .endpoint(settings.endpoint())
                .credential(credential)
                .httpClient(httpClient)
                .retryOptions(retryOptions)
                .httpLogOptions(new HttpLogOptions().setLogLevel(settings.logLevel()));
    }

    public record Settings(
            String endpoint,
            String containerName,
            int maxRetries,
            Duration retryDelay,
            Duration maxRetryDelay,
            Duration requestTimeout,
            HttpLogDetailLevel logLevel) {

        public Settings {
            if (endpoint == null || endpoint.isBlank()) {
                throw new IllegalArgumentException("AZURE_STORAGE_ENDPOINT is required");
            }
            if (!endpoint.startsWith("https://")) {
                throw new IllegalArgumentException("AZURE_STORAGE_ENDPOINT must use HTTPS");
            }
            if (containerName == null || containerName.isBlank()) {
                throw new IllegalArgumentException("AZURE_STORAGE_CONTAINER is required");
            }
            if (maxRetries < 0) {
                throw new IllegalArgumentException("AZURE_STORAGE_MAX_RETRIES must not be negative");
            }
            requirePositive(retryDelay, "AZURE_STORAGE_RETRY_DELAY_SECONDS");
            requirePositive(maxRetryDelay, "AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS");
            requirePositive(requestTimeout, "AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS");
            Objects.requireNonNull(logLevel, "logLevel");
        }

        public static Settings from(Map<String, String> environment) {
            return new Settings(
                    required(environment, "AZURE_STORAGE_ENDPOINT"),
                    required(environment, "AZURE_STORAGE_CONTAINER"),
                    integer(environment, "AZURE_STORAGE_MAX_RETRIES", 5),
                    seconds(environment, "AZURE_STORAGE_RETRY_DELAY_SECONDS", 2),
                    seconds(environment, "AZURE_STORAGE_MAX_RETRY_DELAY_SECONDS", 30),
                    seconds(environment, "AZURE_STORAGE_REQUEST_TIMEOUT_SECONDS", 120),
                    logLevel(environment.getOrDefault("AZURE_STORAGE_LOG_LEVEL", "BASIC")));
        }

        private static String required(Map<String, String> environment, String name) {
            String value = environment.get(name);
            if (value == null || value.isBlank()) {
                throw new IllegalArgumentException(name + " is required");
            }
            return value;
        }

        private static int integer(Map<String, String> environment, String name, int defaultValue) {
            String value = environment.get(name);
            try {
                return value == null ? defaultValue : Integer.parseInt(value);
            } catch (NumberFormatException e) {
                throw new IllegalArgumentException(name + " must be an integer", e);
            }
        }

        private static Duration seconds(Map<String, String> environment, String name, long defaultValue) {
            String value = environment.get(name);
            try {
                return Duration.ofSeconds(value == null ? defaultValue : Long.parseLong(value));
            } catch (NumberFormatException e) {
                throw new IllegalArgumentException(name + " must be an integer number of seconds", e);
            }
        }

        private static HttpLogDetailLevel logLevel(String value) {
            try {
                return HttpLogDetailLevel.valueOf(value.toUpperCase(Locale.ROOT));
            } catch (IllegalArgumentException e) {
                throw new IllegalArgumentException(
                        "AZURE_STORAGE_LOG_LEVEL must be NONE, BASIC, HEADERS, BODY, or BODY_AND_HEADERS", e);
            }
        }

        private static void requirePositive(Duration duration, String name) {
            if (duration == null || duration.isZero() || duration.isNegative()) {
                throw new IllegalArgumentException(name + " must be positive");
            }
        }
    }
}
