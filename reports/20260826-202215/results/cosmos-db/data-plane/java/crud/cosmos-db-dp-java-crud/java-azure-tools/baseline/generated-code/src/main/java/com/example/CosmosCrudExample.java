package com.example;

import com.azure.cosmos.ConsistencyLevel;
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

import java.util.Collections;
import java.util.UUID;

public final class CosmosCrudExample {
    private static final String DATABASE_NAME = "TestDB";
    private static final String CONTAINER_NAME = "Items";
    private static final String CATEGORY = "electronics";

    private CosmosCrudExample() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("COSMOS_ENDPOINT");
        String key = requireEnvironmentVariable("COSMOS_KEY");

        try (CosmosClient client = new CosmosClientBuilder()
            .endpoint(endpoint)
            .key(key)
            .consistencyLevel(ConsistencyLevel.SESSION)
            .buildClient()) {

            CosmosDatabase database = createDatabase(client);
            CosmosContainer container = createContainer(database);

            Item item = new Item(
                UUID.randomUUID().toString(),
                CATEGORY,
                "Laptop",
                1
            );

            createItem(container, item);
            readItem(container, item.getId(), item.getCategory());
            queryItems(container, CATEGORY);
            replaceItem(container, item);
            deleteItem(container, item.getId(), item.getCategory());
        } catch (CosmosException exception) {
            reportCosmosException(exception);
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.println("Configuration error: " + exception.getMessage());
            System.exit(1);
        }
    }

    private static CosmosDatabase createDatabase(CosmosClient client) {
        client.createDatabaseIfNotExists(DATABASE_NAME);
        System.out.printf("Database ready: %s%n", DATABASE_NAME);
        return client.getDatabase(DATABASE_NAME);
    }

    private static CosmosContainer createContainer(CosmosDatabase database) {
        CosmosContainerProperties properties =
            new CosmosContainerProperties(CONTAINER_NAME, "/category");
        database.createContainerIfNotExists(properties);
        System.out.printf("Container ready: %s (partition key: /category)%n", CONTAINER_NAME);
        return database.getContainer(CONTAINER_NAME);
    }

    private static void createItem(CosmosContainer container, Item item) {
        CosmosItemResponse<Item> response = container.createItem(
            item,
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        );
        System.out.printf(
            "Created item %s (request charge: %.2f RU)%n",
            response.getItem().getId(),
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

    private static void queryItems(CosmosContainer container, String category) {
        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM items i WHERE i.category = @category",
            Collections.singletonList(new SqlParameter("@category", category))
        );
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
            .setPartitionKey(new PartitionKey(category));

        System.out.printf("Items in category '%s':%n", category);
        container.queryItems(query, options, Item.class)
            .iterableByPage()
            .forEach(page -> page.getResults().forEach(item ->
                System.out.printf(
                    "  id=%s, name=%s, quantity=%d%n",
                    item.getId(),
                    item.getName(),
                    item.getQuantity()
                )
            ));
    }

    private static void replaceItem(CosmosContainer container, Item item) {
        item.setQuantity(2);
        CosmosItemResponse<Item> response = container.replaceItem(
            item,
            item.getId(),
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        );
        System.out.printf(
            "Replaced item %s with quantity %d%n",
            response.getItem().getId(),
            response.getItem().getQuantity()
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
            throw new IllegalArgumentException(
                "Environment variable " + name + " must be set."
            );
        }
        return value;
    }

    private static void reportCosmosException(CosmosException exception) {
        System.err.printf(
            "Cosmos DB request failed: status=%d, subStatus=%d, activityId=%s, message=%s%n",
            exception.getStatusCode(),
            exception.getSubStatusCode(),
            exception.getActivityId(),
            exception.getMessage()
        );
        if (exception.getRetryAfterDuration() != null
            && !exception.getRetryAfterDuration().isZero()) {
            System.err.printf(
                "Service requested a retry after %d ms.%n",
                exception.getRetryAfterDuration().toMillis()
            );
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
