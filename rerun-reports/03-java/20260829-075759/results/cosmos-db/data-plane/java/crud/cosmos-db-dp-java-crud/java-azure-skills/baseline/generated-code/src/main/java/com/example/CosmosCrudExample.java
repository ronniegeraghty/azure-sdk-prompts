package com.example;

import com.azure.cosmos.ConsistencyLevel;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.models.CosmosContainerProperties;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;

import java.util.List;
import java.util.UUID;

public final class CosmosCrudExample {
    private static final String DATABASE_NAME = "TestDB";
    private static final String CONTAINER_NAME = "Items";

    private CosmosCrudExample() {
    }

    public static void main(String[] args) {
        CosmosClient client = null;

        try {
            String endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
            String key = requiredEnvironmentVariable("COSMOS_KEY");

            client = new CosmosClientBuilder()
                    .endpoint(endpoint)
                    .key(key)
                    .consistencyLevel(ConsistencyLevel.SESSION)
                    .buildClient();

            client.createDatabaseIfNotExists(DATABASE_NAME);
            CosmosDatabase database = client.getDatabase(DATABASE_NAME);

            CosmosContainerProperties properties =
                    new CosmosContainerProperties(CONTAINER_NAME, "/category");
            database.createContainerIfNotExists(properties);
            CosmosContainer container = database.getContainer(CONTAINER_NAME);

            Item item = new Item(
                    UUID.randomUUID().toString(),
                    "electronics",
                    "Wireless keyboard",
                    10);
            PartitionKey partitionKey = new PartitionKey(item.getCategory());

            container.createItem(item, partitionKey, new CosmosItemRequestOptions());
            System.out.printf("Created item %s%n", item.getId());

            Item readItem = container.readItem(
                    item.getId(),
                    partitionKey,
                    Item.class).getItem();
            System.out.printf(
                    "Read item: %s, quantity=%d%n",
                    readItem.getName(),
                    readItem.getQuantity());

            SqlQuerySpec query = new SqlQuerySpec(
                    "SELECT * FROM items i WHERE i.category = @category",
                    List.of(new SqlParameter("@category", "electronics")));

            System.out.println("Electronics:");
            container.queryItems(query, new CosmosQueryRequestOptions(), Item.class)
                    .forEach(result -> System.out.printf(
                            "  %s: %s (quantity=%d)%n",
                            result.getId(),
                            result.getName(),
                            result.getQuantity()));

            readItem.setQuantity(20);
            Item replacedItem = container.replaceItem(
                    readItem,
                    readItem.getId(),
                    partitionKey,
                    new CosmosItemRequestOptions()).getItem();
            System.out.printf(
                    "Updated item %s to quantity %d%n",
                    replacedItem.getId(),
                    replacedItem.getQuantity());

            container.deleteItem(
                    replacedItem.getId(),
                    partitionKey,
                    new CosmosItemRequestOptions());
            System.out.printf("Deleted item %s%n", replacedItem.getId());
        } catch (com.azure.cosmos.CosmosException exception) {
            System.err.printf(
                    "Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, message=%s%n",
                    exception.getStatusCode(),
                    exception.getSubStatusCode(),
                    exception.getActivityId(),
                    exception.getMessage());
            System.exit(1);
        } catch (IllegalStateException exception) {
            System.err.println(exception.getMessage());
            System.exit(1);
        } finally {
            if (client != null) {
                client.close();
            }
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Set the " + name + " environment variable before running the program.");
        }
        return value;
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
