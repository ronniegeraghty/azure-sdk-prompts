package com.example.todo;

import com.azure.core.credential.TokenCredential;
import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.time.Duration;
import java.util.List;

public final class CosmosToDoFactory implements AutoCloseable {
    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = "COSMOS_ENDPOINT";
    public static final String DATABASE_ID = "todo-db";
    public static final String CONTAINER_ID = "items";

    private static final int DEFAULT_TTL_SECONDS = Math.toIntExact(Duration.ofDays(90).toSeconds());

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final CosmosContainer syncContainer;
    private final CosmosAsyncContainer asyncContainer;

    private CosmosToDoFactory(
        CosmosClient syncClient,
        CosmosAsyncClient asyncClient,
        CosmosContainer syncContainer,
        CosmosAsyncContainer asyncContainer
    ) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.syncContainer = syncContainer;
        this.asyncContainer = asyncContainer;
    }

    public static CosmosToDoFactory fromEnvironment() {
        String endpoint = requireEnvironmentVariable(ENDPOINT_ENVIRONMENT_VARIABLE);
        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClient syncClient = new CosmosClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .consistencyLevel(ConsistencyLevel.SESSION)
            .contentResponseOnWriteEnabled(true)
            .buildClient();

        CosmosAsyncClient asyncClient = new CosmosClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .consistencyLevel(ConsistencyLevel.SESSION)
            .contentResponseOnWriteEnabled(true)
            .buildAsyncClient();

        try {
            CosmosDatabase database = syncClient.getDatabase(
                syncClient.createDatabaseIfNotExists(DATABASE_ID).getProperties().getId()
            );

            CosmosContainerProperties properties =
                new CosmosContainerProperties(CONTAINER_ID, "/category");
            properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);
            properties.setIndexingPolicy(new IndexingPolicy().setExcludedPaths(
                List.of(new ExcludedPath("/description/?"))
            ));

            String containerId = database.createContainerIfNotExists(properties)
                .getProperties()
                .getId();
            CosmosContainer syncContainer = database.getContainer(containerId);
            CosmosAsyncContainer asyncContainer = asyncClient
                .getDatabase(database.getId())
                .getContainer(containerId);

            return new CosmosToDoFactory(
                syncClient,
                asyncClient,
                syncContainer,
                asyncContainer
            );
        } catch (RuntimeException exception) {
            asyncClient.close();
            syncClient.close();
            throw exception;
        }
    }

    public SyncToDoRepository syncRepository() {
        return new SyncToDoRepository(syncContainer);
    }

    public AsyncToDoRepository asyncRepository() {
        return new AsyncToDoRepository(asyncContainer);
    }

    @Override
    public void close() {
        asyncClient.close();
        syncClient.close();
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable " + name + " is not set"
            );
        }
        return value;
    }
}
