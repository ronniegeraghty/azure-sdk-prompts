package com.example;

import com.azure.cosmos.ConsistencyLevel;
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
import java.util.Locale;

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

            client.createDatabaseIfNotExists(DATABASE_NAME);
            CosmosDatabase database = client.getDatabase(DATABASE_NAME);

            database.createContainerIfNotExists(CONTAINER_NAME, "/category");
            CosmosContainer container = database.getContainer(CONTAINER_NAME);

            Item item = new Item("item-1", CATEGORY, "Wireless headphones", 10);
            CosmosItemResponse<Item> createResponse =
                container.createItem(item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());
            System.out.printf("Created item %s (RU charge: %.2f)%n",
                createResponse.getItem().getId(), createResponse.getRequestCharge());

            PartitionKey partitionKey = new PartitionKey(CATEGORY);
            CosmosItemResponse<Item> readResponse =
                container.readItem(item.getId(), partitionKey, Item.class);
            System.out.println("Read item: " + readResponse.getItem());

            SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", CATEGORY)));
            CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions();
            queryOptions.setPartitionKey(partitionKey);

            System.out.println("Query results:");
            container.queryItems(query, queryOptions, Item.class)
                .forEach(result -> System.out.println("  " + result));

            item.setQuantity(25);
            CosmosItemResponse<Item> replaceResponse = container.replaceItem(
                item,
                item.getId(),
                partitionKey,
                new CosmosItemRequestOptions());
            System.out.println("Updated item: " + replaceResponse.getItem());

            container.deleteItem(item.getId(), partitionKey, new CosmosItemRequestOptions());
            System.out.println("Deleted item: " + item.getId());
        } catch (CosmosException exception) {
            System.err.printf(
                Locale.ROOT,
                "Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, message=%s%n",
                exception.getStatusCode(),
                exception.getSubStatusCode(),
                exception.getActivityId(),
                exception.getMessage());
            throw exception;
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
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
