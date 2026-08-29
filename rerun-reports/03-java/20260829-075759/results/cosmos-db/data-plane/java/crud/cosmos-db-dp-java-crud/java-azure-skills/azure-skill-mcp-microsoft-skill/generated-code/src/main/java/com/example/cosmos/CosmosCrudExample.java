package com.example.cosmos;

import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.azure.cosmos.util.CosmosPagedIterable;

import java.net.URI;
import java.util.Collections;
import java.util.Locale;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class CosmosCrudExample {
    private static final Logger LOGGER = Logger.getLogger(CosmosCrudExample.class.getName());

    private static final String DATABASE_ID = "TestDB";
    private static final String CONTAINER_ID = "Items";
    private static final String PARTITION_KEY_PATH = "/category";
    private static final String CATEGORY = "electronics";

    private CosmosCrudExample() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
        String key = requireEnvironmentVariable("COSMOS_KEY");
        requireLocalEmulatorEndpoint(endpoint);

        try (CosmosClient client = new CosmosClientBuilder()
            .endpoint(endpoint)
            .key(key)
            .contentResponseOnWriteEnabled(true)
            .buildClient()) {

            CosmosDatabase database = createDatabase(client);
            CosmosContainer container = createContainer(database);
            runCrudOperations(container);
        } catch (CosmosException exception) {
            logCosmosException(exception);
            System.exit(1);
        }
    }

    private static CosmosDatabase createDatabase(CosmosClient client) {
        client.createDatabaseIfNotExists(DATABASE_ID);
        LOGGER.info(() -> "Database is ready: " + DATABASE_ID);
        return client.getDatabase(DATABASE_ID);
    }

    private static CosmosContainer createContainer(CosmosDatabase database) {
        CosmosContainerProperties properties =
            new CosmosContainerProperties(CONTAINER_ID, PARTITION_KEY_PATH);
        database.createContainerIfNotExists(properties);
        LOGGER.info(() -> "Container is ready: " + CONTAINER_ID);
        return database.getContainer(CONTAINER_ID);
    }

    private static void runCrudOperations(CosmosContainer container) {
        Item item = new Item("item-001", CATEGORY, "Wireless Headphones", 10);
        PartitionKey partitionKey = new PartitionKey(item.getCategory());

        CosmosItemResponse<Item> createResponse = container.createItem(
            item,
            partitionKey,
            new CosmosItemRequestOptions());
        logOperation("Created", createResponse);

        CosmosItemResponse<Item> readResponse =
            container.readItem(item.getId(), partitionKey, Item.class);
        Item storedItem = readResponse.getItem();
        logOperation("Read", readResponse);

        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            Collections.singletonList(new SqlParameter("@category", CATEGORY)));
        CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions()
            .setPartitionKey(partitionKey);
        CosmosPagedIterable<Item> queryResults =
            container.queryItems(query, queryOptions, Item.class);

        queryResults.forEach(result ->
            LOGGER.info(() -> String.format(
                Locale.ROOT,
                "Query result: id=%s, name=%s, quantity=%d",
                result.getId(),
                result.getName(),
                result.getQuantity())));

        storedItem.setQuantity(25);
        CosmosItemResponse<Item> replaceResponse = container.replaceItem(
            storedItem,
            storedItem.getId(),
            partitionKey,
            new CosmosItemRequestOptions());
        logOperation("Replaced", replaceResponse);

        container.deleteItem(
            storedItem.getId(),
            partitionKey,
            new CosmosItemRequestOptions());
        LOGGER.info(() -> "Deleted item: " + storedItem.getId());
    }

    private static void logOperation(String operation, CosmosItemResponse<Item> response) {
        LOGGER.info(() -> String.format(
            Locale.ROOT,
            "%s item %s (status=%d, requestCharge=%.2f RU)",
            operation,
            response.getItem().getId(),
            response.getStatusCode(),
            response.getRequestCharge()));
    }

    private static void logCosmosException(CosmosException exception) {
        LOGGER.log(
            Level.SEVERE,
            String.format(
                Locale.ROOT,
                "Cosmos DB operation failed: status=%d, subStatus=%d, activityId=%s, "
                    + "requestCharge=%.2f RU, retryAfter=%s",
                exception.getStatusCode(),
                exception.getSubStatusCode(),
                exception.getActivityId(),
                exception.getRequestCharge(),
                exception.getRetryAfterDuration()),
            exception);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Environment variable " + name + " is required.");
        }
        return value;
    }

    private static void requireLocalEmulatorEndpoint(String endpoint) {
        URI endpointUri;
        try {
            endpointUri = URI.create(endpoint);
        } catch (IllegalArgumentException exception) {
            throw new IllegalStateException("COSMOS_ENDPOINT must be a valid URI.", exception);
        }

        String host = endpointUri.getHost();
        if (host == null
            || !(host.equalsIgnoreCase("localhost")
            || host.equals("127.0.0.1")
            || host.equals("::1"))) {
            throw new IllegalStateException(
                "This sample is restricted to a local Cosmos DB emulator endpoint.");
        }
    }

    public static final class Item {
        private String id;
        private String category;
        private String name;
        private int quantity;

        public Item() {
        }

        public Item(String id, String category, String name, int quantity) {
            this.id = id;
            this.category = category;
            this.name = name;
            this.quantity = quantity;
        }

        public String getId() {
            return id;
        }

        public void setId(String id) {
            this.id = id;
        }

        public String getCategory() {
            return category;
        }

        public void setCategory(String category) {
            this.category = category;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public int getQuantity() {
            return quantity;
        }

        public void setQuantity(int quantity) {
            this.quantity = quantity;
        }
    }
}
