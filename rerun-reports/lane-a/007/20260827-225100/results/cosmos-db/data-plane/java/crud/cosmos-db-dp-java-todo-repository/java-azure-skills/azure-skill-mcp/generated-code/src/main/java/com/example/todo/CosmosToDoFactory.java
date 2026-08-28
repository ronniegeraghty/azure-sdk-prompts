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
import com.azure.cosmos.models.IncludedPath;
import com.azure.cosmos.models.IndexingMode;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;

import java.util.List;

public final class CosmosToDoFactory implements AutoCloseable {
    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = "COSMOS_ENDPOINT";
    public static final String DATABASE_NAME = "todo-db";
    public static final String CONTAINER_NAME = "todos";
    public static final int DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60;

    private final CosmosClient syncClient;
    private final CosmosAsyncClient asyncClient;
    private final CosmosContainer syncContainer;
    private final CosmosAsyncContainer asyncContainer;

    private CosmosToDoFactory(
            CosmosClient syncClient,
            CosmosAsyncClient asyncClient,
            CosmosContainer syncContainer,
            CosmosAsyncContainer asyncContainer) {
        this.syncClient = syncClient;
        this.asyncClient = asyncClient;
        this.syncContainer = syncContainer;
        this.asyncContainer = asyncContainer;
    }

    public static CosmosToDoFactory create() {
        String endpoint = requireEnvironmentVariable(ENDPOINT_ENVIRONMENT_VARIABLE);
        String managedIdentityClientId = System.getenv("AZURE_CLIENT_ID");

        ManagedIdentityCredentialBuilder credentialBuilder = new ManagedIdentityCredentialBuilder();
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        TokenCredential credential = credentialBuilder.build();

        CosmosClientBuilder clientBuilder = new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(credential)
                .consistencyLevel(ConsistencyLevel.SESSION);

        CosmosClient syncClient = clientBuilder.buildClient();
        CosmosAsyncClient asyncClient = clientBuilder.buildAsyncClient();

        try {
            syncClient.createDatabaseIfNotExists(DATABASE_NAME);
            CosmosDatabase database = syncClient.getDatabase(DATABASE_NAME);
            database.createContainerIfNotExists(containerProperties());

            return new CosmosToDoFactory(
                    syncClient,
                    asyncClient,
                    database.getContainer(CONTAINER_NAME),
                    asyncClient.getDatabase(DATABASE_NAME).getContainer(CONTAINER_NAME));
        } catch (RuntimeException exception) {
            asyncClient.close();
            syncClient.close();
            throw exception;
        }
    }

    private static CosmosContainerProperties containerProperties() {
        CosmosContainerProperties properties =
                new CosmosContainerProperties(CONTAINER_NAME, "/category");
        properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);

        IndexingPolicy indexingPolicy = new IndexingPolicy();
        indexingPolicy.setAutomatic(true);
        indexingPolicy.setIndexingMode(IndexingMode.CONSISTENT);
        indexingPolicy.setIncludedPaths(List.of(new IncludedPath("/*")));
        indexingPolicy.setExcludedPaths(List.of(new ExcludedPath("/description/?")));
        properties.setIndexingPolicy(indexingPolicy);
        return properties;
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
        return value;
    }
}
