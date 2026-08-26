package com.example.todo;

import com.azure.core.credential.TokenCredential;
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
import java.util.Objects;

public final class CosmosConfiguration implements AutoCloseable {
    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = "COSMOS_ENDPOINT";
    public static final String MANAGED_IDENTITY_CLIENT_ID_VARIABLE = "AZURE_CLIENT_ID";
    public static final String DATABASE_ID = "todo-db";
    public static final String CONTAINER_ID = "items";

    private static final int DEFAULT_TTL_SECONDS = Math.toIntExact(Duration.ofDays(90).toSeconds());

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final CosmosContainer syncContainer;
    private final CosmosAsyncContainer asyncContainer;

    private CosmosConfiguration(
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

    public static CosmosConfiguration createFromEnvironment() {
        String endpoint = requireEnvironmentVariable(ENDPOINT_ENVIRONMENT_VARIABLE);
        String clientId = System.getenv(MANAGED_IDENTITY_CLIENT_ID_VARIABLE);

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (clientId != null && !clientId.isBlank()) {
            credentialBuilder.clientId(clientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClient syncClient = new CosmosClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .contentResponseOnWriteEnabled(true)
            .buildClient();

        CosmosAsyncClient asyncClient = new CosmosClientBuilder()
            .endpoint(endpoint)
            .credential(credential)
            .contentResponseOnWriteEnabled(true)
            .buildAsyncClient();

        try {
            CosmosDatabase database = syncClient.getDatabase(
                syncClient.createDatabaseIfNotExists(DATABASE_ID).getProperties().getId());

            CosmosContainerProperties properties =
                new CosmosContainerProperties(CONTAINER_ID, "/category");
            properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);
            properties.setIndexingPolicy(new IndexingPolicy()
                .setExcludedPaths(List.of(new ExcludedPath("/description/?"))));

            String containerId = database.createContainerIfNotExists(properties)
                .getProperties()
                .getId();

            return new CosmosConfiguration(
                syncClient,
                asyncClient,
                database.getContainer(containerId),
                asyncClient.getDatabase(DATABASE_ID).getContainer(containerId));
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
            throw new IllegalStateException("Required environment variable " + name + " is not set");
        }
        return Objects.requireNonNull(value);
    }
}
