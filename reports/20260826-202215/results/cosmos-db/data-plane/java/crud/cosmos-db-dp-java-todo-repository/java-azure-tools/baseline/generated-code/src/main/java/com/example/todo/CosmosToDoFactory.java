package com.example.todo;

import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingMode;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;

import java.time.Duration;
import java.util.List;
import java.util.Objects;

public final class CosmosToDoFactory implements AutoCloseable {
    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = "COSMOS_ENDPOINT";
    public static final int DEFAULT_TTL_SECONDS = (int) Duration.ofDays(90).toSeconds();

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final String databaseName;
    private final String containerName;

    private CosmosToDoFactory(
            CosmosClient syncClient,
            CosmosAsyncClient asyncClient,
            String databaseName,
            String containerName) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.databaseName = databaseName;
        this.containerName = containerName;
    }

    public static CosmosToDoFactory create(String databaseName, String containerName) {
        String endpoint = requireEnvironmentVariable(ENDPOINT_ENVIRONMENT_VARIABLE);
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();

        CosmosClient syncClient = clientBuilder(endpoint, credential).buildClient();
        CosmosAsyncClient asyncClient = null;
        try {
            initializeSchema(syncClient, databaseName, containerName);
            asyncClient = clientBuilder(endpoint, credential).buildAsyncClient();
            return new CosmosToDoFactory(
                    syncClient, asyncClient, databaseName, containerName);
        } catch (RuntimeException exception) {
            if (asyncClient != null) {
                asyncClient.close();
            }
            syncClient.close();
            throw exception;
        }
    }

    public SyncToDoRepository syncRepository() {
        return new SyncToDoRepository(
                syncClient.getDatabase(databaseName).getContainer(containerName));
    }

    public AsyncToDoRepository asyncRepository() {
        return new AsyncToDoRepository(
                asyncClient.getDatabase(databaseName).getContainer(containerName));
    }

    @Override
    public void close() {
        asyncClient.close();
        syncClient.close();
    }

    private static CosmosClientBuilder clientBuilder(
            String endpoint,
            DefaultAzureCredential credential) {
        return new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .consistencyLevel(ConsistencyLevel.SESSION);
    }

    private static void initializeSchema(
            CosmosClient client,
            String databaseName,
            String containerName) {
        Objects.requireNonNull(databaseName, "databaseName");
        Objects.requireNonNull(containerName, "containerName");

        client.createDatabaseIfNotExists(databaseName);

        IndexingPolicy indexingPolicy = new IndexingPolicy()
                .setAutomatic(true)
                .setIndexingMode(IndexingMode.CONSISTENT)
                .setExcludedPaths(List.of(new ExcludedPath("/description/?")));
        CosmosContainerProperties properties = new CosmosContainerProperties(
                containerName, "/category");
        properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);
        properties.setIndexingPolicy(indexingPolicy);

        client.getDatabase(databaseName).createContainerIfNotExists(properties);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Environment variable " + name + " must contain the Cosmos DB endpoint.");
        }
        return value;
    }
}
