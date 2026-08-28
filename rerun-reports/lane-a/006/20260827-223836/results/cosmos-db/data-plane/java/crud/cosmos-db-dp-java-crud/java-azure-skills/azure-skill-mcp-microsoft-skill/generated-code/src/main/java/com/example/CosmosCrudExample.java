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
import com.azure.cosmos.util.CosmosPagedIterable;

import java.util.List;
import java.util.UUID;

public final class CosmosCrudExample {
    private static final String DATABASE_ID = "TestDB";
    private static final String CONTAINER_ID = "Items";
    private static final String CATEGORY = "electronics";

    private CosmosCrudExample() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
        String key = requireEnvironmentVariable("COSMOS_KEY");

        try (CosmosClient client = new CosmosClientBuilder()
            .endpoint(endpoint)
            .key(key)
            .buildClient()) {

            CosmosDatabase database = createDatabase(client);
            CosmosContainer container = createContainer(database);

            Item item = new Item(
                UUID.randomUUID().toString(),
                CATEGORY,
                "Wireless headphones",
                10
            );

            createItem(container, item);
            Item readItem = readItem(container, item.getId(), item.getCategory());
            queryItemsByCategory(container, CATEGORY);
            replaceItem(container, readItem, 25);
            deleteItem(container, item.getId(), item.getCategory());
        } catch (CosmosException exception) {
            reportCosmosException(exception);
            System.exit(1);
        }
    }

    private static CosmosDatabase createDatabase(CosmosClient client) {
        client.createDatabaseIfNotExists(DATABASE_ID);
        System.out.printf("Database ready: %s%n", DATABASE_ID);
        return client.getDatabase(DATABASE_ID);
    }

    private static CosmosContainer createContainer(CosmosDatabase database) {
        database.createContainerIfNotExists(CONTAINER_ID, "/category");
        System.out.printf("Container ready: %s (partition key: /category)%n", CONTAINER_ID);
        return database.getContainer(CONTAINER_ID);
    }

    private static void createItem(CosmosContainer container, Item item) {
        CosmosItemResponse<Item> response = container.createItem(
            item,
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        );
        System.out.printf(
            "Created item %s (request charge: %.2f RU)%n",
            item.getId(),
            response.getRequestCharge()
        );
    }

    private static Item readItem(CosmosContainer container, String id, String category) {
        CosmosItemResponse<Item> response = container.readItem(
            id,
            new PartitionKey(category),
            Item.class
        );
        Item item = response.getItem();
        System.out.printf(
            "Read item: id=%s, name=%s, quantity=%d%n",
            item.getId(),
            item.getName(),
            item.getQuantity()
        );
        return item;
    }

    private static void queryItemsByCategory(CosmosContainer container, String category) {
        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category))
        );

        CosmosPagedIterable<Item> items = container.queryItems(
            query,
            new CosmosQueryRequestOptions(),
            Item.class
        );

        System.out.printf("Items in category '%s':%n", category);
        items.forEach(item -> System.out.printf(
            "  id=%s, name=%s, quantity=%d%n",
            item.getId(),
            item.getName(),
            item.getQuantity()
        ));
    }

    private static void replaceItem(CosmosContainer container, Item item, int newQuantity) {
        item.setQuantity(newQuantity);
        CosmosItemResponse<Item> response = container.replaceItem(
            item,
            item.getId(),
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        );
        System.out.printf(
            "Replaced item %s with quantity %d (request charge: %.2f RU)%n",
            item.getId(),
            response.getItem().getQuantity(),
            response.getRequestCharge()
        );
    }

    private static void deleteItem(CosmosContainer container, String id, String category) {
        container.deleteItem(
            id,
            new PartitionKey(category),
            new CosmosItemRequestOptions()
        );
        System.out.printf("Deleted item %s%n", id);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    private static void reportCosmosException(CosmosException exception) {
        System.err.printf(
            "Cosmos DB operation failed: status=%d, substatus=%d, activityId=%s, requestCharge=%.2f, message=%s%n",
            exception.getStatusCode(),
            exception.getSubStatusCode(),
            exception.getActivityId(),
            exception.getRequestCharge(),
            exception.getMessage()
        );

        switch (exception.getStatusCode()) {
            case 404 -> System.err.println("The requested database, container, or item was not found.");
            case 409 -> System.err.println("A resource with the same id and partition key already exists.");
            case 429 -> System.err.printf(
                "Request rate was limited; retry after %s.%n",
                exception.getRetryAfterDuration()
            );
            default -> {
                // The diagnostics above contain the request details needed for investigation.
            }
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
