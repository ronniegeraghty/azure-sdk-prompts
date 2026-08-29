package com.example.todo;

import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosAsyncDatabase;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.ExcludedPath;
import com.azure.cosmos.models.IndexingPolicy;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import reactor.core.publisher.Mono;

import java.time.Duration;

public final class CosmosConfiguration {
    public static final String ENDPOINT_ENVIRONMENT_VARIABLE = "AZURE_COSMOS_ENDPOINT";
    public static final String DEFAULT_DATABASE = "todo-db";
    public static final String DEFAULT_CONTAINER = "items";

    private static final int DEFAULT_TTL_SECONDS = (int) Duration.ofDays(90).toSeconds();

    private CosmosConfiguration() {
    }

    public static CosmosClient createSyncClient() {
        return clientBuilder().buildClient();
    }

    public static CosmosAsyncClient createAsyncClient() {
        return clientBuilder().buildAsyncClient();
    }

    public static CosmosContainer initializeSync(
            CosmosClient client, String databaseName, String containerName) {
        client.createDatabaseIfNotExists(databaseName);
        CosmosDatabase database = client.getDatabase(databaseName);
        database.createContainerIfNotExists(containerProperties(containerName));
        return database.getContainer(containerName);
    }

    public static Mono<CosmosAsyncContainer> initializeAsync(
            CosmosAsyncClient client, String databaseName, String containerName) {
        CosmosAsyncDatabase database = client.getDatabase(databaseName);
        CosmosAsyncContainer container = database.getContainer(containerName);
        return client.createDatabaseIfNotExists(databaseName)
                .then(database.createContainerIfNotExists(containerProperties(containerName)))
                .thenReturn(container);
    }

    private static CosmosClientBuilder clientBuilder() {
        String endpoint = System.getenv(ENDPOINT_ENVIRONMENT_VARIABLE);
        if (endpoint == null || endpoint.isBlank()) {
            throw new IllegalStateException(
                    ENDPOINT_ENVIRONMENT_VARIABLE + " environment variable is required");
        }

        return new CosmosClientBuilder()
                .endpoint(endpoint)
                .credential(new ManagedIdentityCredentialBuilder().build())
                .consistencyLevel(ConsistencyLevel.SESSION);
    }

    private static CosmosContainerProperties containerProperties(String containerName) {
        IndexingPolicy indexingPolicy = new IndexingPolicy();
        indexingPolicy.getExcludedPaths().add(new ExcludedPath("/description/?"));

        CosmosContainerProperties properties =
                new CosmosContainerProperties(containerName, "/category");
        properties.setDefaultTimeToLiveInSeconds(DEFAULT_TTL_SECONDS);
        properties.setIndexingPolicy(indexingPolicy);
        return properties;
    }
}
