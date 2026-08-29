package com.example.todo.config;

import com.azure.core.credential.TokenCredential;
import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.CosmosContainerResponse;
import com.azure.cosmos.models.CosmosDatabaseResponse;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.example.todo.repository.AsyncTodoRepository;
import com.example.todo.repository.SyncTodoRepository;

import java.util.List;
import java.util.logging.Logger;

public final class CosmosTodoFactory implements AutoCloseable {
    private static final Logger LOGGER = Logger.getLogger(CosmosTodoFactory.class.getName());
    private static final int DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60;

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final CosmosContainer syncContainer;
    private final CosmosAsyncContainer asyncContainer;

    private CosmosTodoFactory(
            CosmosClient syncClient,
            CosmosAsyncClient asyncClient,
            CosmosContainer syncContainer,
            CosmosAsyncContainer asyncContainer) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.syncContainer = syncContainer;
        this.asyncContainer = asyncContainer;
    }

    public static CosmosTodoFactory create() {
        String endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
        String databaseId = environmentVariableOrDefault("COSMOS_DATABASE", "todo-db");
        String containerId = environmentVariableOrDefault("COSMOS_CONTAINER", "todos");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String clientId = System.getenv("AZURE_CLIENT_ID");
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClient syncClient = clientBuilder(endpoint, credential).buildClient();
        CosmosAsyncClient asyncClient = clientBuilder(endpoint, credential).buildAsyncClient();
        try {
            CosmosDatabaseResponse databaseResponse =
                    syncClient.createDatabaseIfNotExists(databaseId);
            LOGGER.info(() -> "database initialization RU="
                    + String.format("%.2f", databaseResponse.getRequestCharge()));
            CosmosDatabase database = syncClient.getDatabase(databaseId);

            CosmosContainerProperties properties =
                    new CosmosContainerProperties(containerId, "/category");
            properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);
            IndexingPolicy indexingPolicy = new IndexingPolicy();
            indexingPolicy.setExcludedPaths(List.of(new ExcludedPath("/description/?")));
            properties.setIndexingPolicy(indexingPolicy);

            CosmosContainerResponse containerResponse =
                    database.createContainerIfNotExists(properties);
            LOGGER.info(() -> "container initialization RU="
                    + String.format("%.2f", containerResponse.getRequestCharge()));

            return new CosmosTodoFactory(
                    syncClient,
                    asyncClient,
                    database.getContainer(containerId),
                    asyncClient.getDatabase(databaseId).getContainer(containerId));
        } catch (RuntimeException exception) {
            asyncClient.close();
            syncClient.close();
            throw exception;
        }
    }

    public SyncTodoRepository syncRepository() {
        return new SyncTodoRepository(syncContainer);
    }

    public AsyncTodoRepository asyncRepository() {
        return new AsyncTodoRepository(asyncContainer);
    }

    @Override
    public void close() {
        asyncClient.close();
        syncClient.close();
    }

    private static CosmosClientBuilder clientBuilder(
            String endpoint,
            TokenCredential credential) {
        return new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .consistencyLevel(ConsistencyLevel.SESSION)
                .contentResponseOnWriteEnabled(true)
                .connectionSharingAcrossClientsEnabled(true)
                .userAgentSuffix("cosmos-todo-repository");
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " environment variable must be set");
        }
        return value;
    }

    private static String environmentVariableOrDefault(String name, String defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : value;
    }
}
