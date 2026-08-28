package com.example.todo;

import com.azure.core.credential.TokenCredential;
import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.time.Duration;
import java.util.List;

public final class CosmosClientFactory implements AutoCloseable {
    private static final int DEFAULT_TTL_SECONDS = Math.toIntExact(Duration.ofDays(90).toSeconds());

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final String databaseName;
    private final String containerName;

    private CosmosClientFactory(
            CosmosClient syncClient,
            CosmosAsyncClient asyncClient,
            String databaseName,
            String containerName) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.databaseName = databaseName;
        this.containerName = containerName;
    }

    public static CosmosClientFactory createFromEnvironment() {
        String endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
        String databaseName = environmentVariableOrDefault("COSMOS_DATABASE", "todo-db");
        String containerName = environmentVariableOrDefault("COSMOS_CONTAINER", "todos");
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClientBuilder clientBuilder = new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .consistencyLevel(ConsistencyLevel.SESSION)
                .contentResponseOnWriteEnabled(true);

        CosmosClient syncClient = clientBuilder.buildClient();
        try {
            initialize(syncClient, databaseName, containerName);
            CosmosAsyncClient asyncClient = clientBuilder.buildAsyncClient();
            return new CosmosClientFactory(
                    syncClient,
                    asyncClient,
                    databaseName,
                    containerName);
        } catch (RuntimeException exception) {
            syncClient.close();
            throw exception;
        }
    }

    public CosmosContainer syncContainer() {
        return syncClient.getDatabase(databaseName).getContainer(containerName);
    }

    public CosmosAsyncContainer asyncContainer() {
        return asyncClient.getDatabase(databaseName).getContainer(containerName);
    }

    private static void initialize(
            CosmosClient client,
            String databaseName,
            String containerName) {
        client.createDatabaseIfNotExists(databaseName);

        IndexingPolicy indexingPolicy = new IndexingPolicy()
                .setExcludedPaths(List.of(new ExcludedPath("/description/?")));
        CosmosContainerProperties properties =
                new CosmosContainerProperties(containerName, "/category")
                        .setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS)
                        .setIndexingPolicy(indexingPolicy);

        client.getDatabase(databaseName).createContainerIfNotExists(properties);
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Required environment variable " + name + " is not set");
        }
        return value;
    }

    private static String environmentVariableOrDefault(String name, String defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : value;
    }

    @Override
    public void close() {
        asyncClient.close();
        syncClient.close();
    }
}
