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

import java.util.List;
import java.util.UUID;

public final class CosmosCrudApp {
    private static final String DATABASE_NAME = "TestDB";
    private static final String CONTAINER_NAME = "Items";
    private static final String CATEGORY = "electronics";

    private CosmosCrudApp() {
    }

    public static void main(String[] args) {
        String endpoint = requiredEnvironmentVariable("COSMOS_ENDPOINT");
        String key = requiredEnvironmentVariable("COSMOS_KEY");

        try (CosmosClient client = new CosmosClientBuilder()
                .endpoint(endpoint)
                .key(key)
                .consistencyLevel(ConsistencyLevel.SESSION)
                .buildClient()) {

            client.createDatabaseIfNotExists(DATABASE_NAME);
            CosmosDatabase database = client.getDatabase(DATABASE_NAME);

            CosmosContainerProperties properties =
                    new CosmosContainerProperties(CONTAINER_NAME, "/category");
            database.createContainerIfNotExists(properties);
            CosmosContainer container = database.getContainer(CONTAINER_NAME);

            Item item = new Item(
                    UUID.randomUUID().toString(),
                    CATEGORY,
                    "Wireless headphones",
                    10);

            CosmosItemResponse<Item> createResponse =
                    container.createItem(item, new PartitionKey(item.getCategory()),
                            new CosmosItemRequestOptions());
            System.out.printf("Created item %s (request charge: %.2f RU)%n",
                    item.getId(), createResponse.getRequestCharge());

            CosmosItemResponse<Item> readResponse =
                    container.readItem(item.getId(), new PartitionKey(item.getCategory()), Item.class);
            Item readItem = readResponse.getItem();
            System.out.printf("Read item: %s, quantity: %d%n",
                    readItem.getName(), readItem.getQuantity());

            SqlQuerySpec query = new SqlQuerySpec(
                    "SELECT * FROM c WHERE c.category = @category",
                    List.of(new SqlParameter("@category", CATEGORY)));

            System.out.println("Items in category '" + CATEGORY + "':");
            container.queryItems(query, new CosmosQueryRequestOptions(), Item.class)
                    .forEach(result -> System.out.printf("- %s: %s (quantity: %d)%n",
                            result.getId(), result.getName(), result.getQuantity()));

            readItem.setQuantity(25);
            CosmosItemResponse<Item> replaceResponse =
                    container.replaceItem(
                            readItem,
                            readItem.getId(),
                            new PartitionKey(readItem.getCategory()),
                            new CosmosItemRequestOptions());
            System.out.printf("Updated quantity to %d (request charge: %.2f RU)%n",
                    replaceResponse.getItem().getQuantity(), replaceResponse.getRequestCharge());

            CosmosItemResponse<Object> deleteResponse =
                    container.deleteItem(
                            readItem.getId(),
                            new PartitionKey(readItem.getCategory()),
                            new CosmosItemRequestOptions());
            System.out.printf("Deleted item %s (request charge: %.2f RU)%n",
                    readItem.getId(), deleteResponse.getRequestCharge());
        } catch (CosmosException exception) {
            System.err.printf(
                    "Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, message=%s%n",
                    exception.getStatusCode(),
                    exception.getSubStatusCode(),
                    exception.getActivityId(),
                    exception.getMessage());
            System.exit(1);
        }
    }

    private static String requiredEnvironmentVariable(String name) {
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
    }
}
