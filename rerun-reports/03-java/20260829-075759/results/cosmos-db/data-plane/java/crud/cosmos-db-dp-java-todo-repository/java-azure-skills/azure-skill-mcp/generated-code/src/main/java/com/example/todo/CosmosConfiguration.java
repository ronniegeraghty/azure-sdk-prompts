package com.example.todo;

import com.azure.core.credential.TokenCredential;
import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.util.List;

public final class CosmosConfiguration implements AutoCloseable {
    private static final int DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60;

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final String databaseId;
    private final String containerId;

    private CosmosConfiguration(
            CosmosClient syncClient,
            CosmosAsyncClient asyncClient,
            String databaseId,
            String containerId) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.databaseId = databaseId;
        this.containerId = containerId;
    }

    public static CosmosConfiguration fromEnvironment() {
        String endpoint = requireEnvironment("AZURE_COSMOS_ENDPOINT");
        String databaseId = environmentOrDefault("AZURE_COSMOS_DATABASE", "todo-db");
        String containerId = environmentOrDefault("AZURE_COSMOS_CONTAINER", "items");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClientBuilder clientBuilder = new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .consistencyLevel(ConsistencyLevel.SESSION)
                .contentResponseOnWriteEnabled(true);

        CosmosClient syncClient = clientBuilder.buildClient();
        CosmosAsyncClient asyncClient = clientBuilder.buildAsyncClient();

        try {
            initialize(syncClient, databaseId, containerId);
            return new CosmosConfiguration(syncClient, asyncClient, databaseId, containerId);
        } catch (RuntimeException exception) {
            syncClient.close();
            asyncClient.close();
            throw exception;
        }
    }

    public SyncToDoRepository syncRepository() {
        CosmosContainer container = syncClient.getDatabase(databaseId).getContainer(containerId);
        return new SyncToDoRepository(container);
    }

    public AsyncToDoRepository asyncRepository() {
        return new AsyncToDoRepository(
                asyncClient.getDatabase(databaseId).getContainer(containerId));
    }

    @Override
    public void close() {
        syncClient.close();
        asyncClient.close();
    }

    private static void initialize(CosmosClient client, String databaseId, String containerId) {
        client.createDatabaseIfNotExists(databaseId);

        CosmosContainerProperties properties = new CosmosContainerProperties(containerId, "/category");
        properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);

        IndexingPolicy indexingPolicy = new IndexingPolicy();
        indexingPolicy.setExcludedPaths(List.of(new ExcludedPath("/description/?")));
        properties.setIndexingPolicy(indexingPolicy);

        client.getDatabase(databaseId).createContainerIfNotExists(properties);
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static String environmentOrDefault(String name, String defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : value;
    }
}
