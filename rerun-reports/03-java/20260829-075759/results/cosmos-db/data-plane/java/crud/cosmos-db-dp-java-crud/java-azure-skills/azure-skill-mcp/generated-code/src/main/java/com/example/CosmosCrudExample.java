package com.example;

import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;

import java.util.List;
import java.util.UUID;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class CosmosCrudExample {
    private static final Logger LOGGER = Logger.getLogger(CosmosCrudExample.class.getName());
    private static final String DATABASE_ID = "TestDB";
    private static final String CONTAINER_ID = "Items";
    private static final String PARTITION_KEY_PATH = "/category";

    private CosmosCrudExample() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
        String key = requireEnvironmentVariable("COSMOS_KEY");

        try (CosmosClient client = new CosmosClientBuilder()
                .endpoint(endpoint)
                .key(key)
                .buildClient()) {
            CosmosContainer container = createContainer(client);
            runCrudOperations(container);
        } catch (CosmosException exception) {
            logCosmosException(exception);
            System.exit(1);
        }
    }

    private static CosmosContainer createContainer(CosmosClient client) {
        client.createDatabaseIfNotExists(DATABASE_ID);
        CosmosDatabase database = client.getDatabase(DATABASE_ID);

        database.createContainerIfNotExists(CONTAINER_ID, PARTITION_KEY_PATH);
        LOGGER.info(() -> "Database '" + DATABASE_ID + "' and container '"
                + CONTAINER_ID + "' are ready.");
        return database.getContainer(CONTAINER_ID);
    }

    private static void runCrudOperations(CosmosContainer container) {
        Item item = new Item(
                UUID.randomUUID().toString(),
                "electronics",
                "Wireless keyboard",
                10);
        PartitionKey partitionKey = new PartitionKey(item.getCategory());

        CosmosItemResponse<Item> createResponse =
                container.createItem(item, partitionKey, new CosmosItemRequestOptions());
        LOGGER.info(() -> "Created item " + createResponse.getItem().getId());

        CosmosItemResponse<Item> readResponse =
                container.readItem(item.getId(), partitionKey, Item.class);
        Item readItem = readResponse.getItem();
        LOGGER.info(() -> "Read item: " + readItem);

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM items i WHERE i.category = @category",
                List.of(new SqlParameter("@category", "electronics")));
        CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey("electronics"));

        LOGGER.info("Electronics items:");
        container.queryItems(query, queryOptions, Item.class)
                .forEach(result -> LOGGER.info(result::toString));

        readItem.setQuantity(25);
        CosmosItemResponse<Item> replaceResponse = container.replaceItem(
                readItem,
                readItem.getId(),
                partitionKey,
                new CosmosItemRequestOptions());
        LOGGER.info(() -> "Updated quantity to " + replaceResponse.getItem().getQuantity());

        container.deleteItem(item.getId(), partitionKey, new CosmosItemRequestOptions());
        LOGGER.info(() -> "Deleted item " + item.getId());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static void logCosmosException(CosmosException exception) {
        LOGGER.log(
                Level.SEVERE,
                "Cosmos DB request failed. Status code: {0}, substatus code: {1}, "
                        + "activity ID: {2}, message: {3}, diagnostics: {4}",
                new Object[]{
                        exception.getStatusCode(),
                        exception.getSubStatusCode(),
                        exception.getActivityId(),
                        exception.getMessage(),
                        exception.getDiagnostics()
                });
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

        @Override
        public String toString() {
            return "Item{"
                    + "id='" + id + '\''
                    + ", category='" + category + '\''
                    + ", name='" + name + '\''
                    + ", quantity=" + quantity
                    + '}';
        }
    }
}
