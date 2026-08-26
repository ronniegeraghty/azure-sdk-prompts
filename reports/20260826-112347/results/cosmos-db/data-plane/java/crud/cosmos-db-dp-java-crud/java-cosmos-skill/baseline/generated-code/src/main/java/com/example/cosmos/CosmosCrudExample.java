package com.example.cosmos;

import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosClientBuilder;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosDatabase;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.azure.cosmos.util.CosmosPagedIterable;

import java.util.List;

public final class CosmosCrudExample {
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

            runCrudOperations(client);
        } catch (CosmosException exception) {
            reportCosmosException(exception);
            System.exit(1);
        }
    }

    private static void runCrudOperations(CosmosClient client) {
        client.createDatabaseIfNotExists(DATABASE_ID);
        CosmosDatabase database = client.getDatabase(DATABASE_ID);

        database.createContainerIfNotExists(CONTAINER_ID, PARTITION_KEY_PATH);
        CosmosContainer container = database.getContainer(CONTAINER_ID);

        Item item = new Item("item-1", "electronics", "Wireless keyboard", 10);
        PartitionKey partitionKey = new PartitionKey(item.getCategory());

        CosmosItemResponse<Item> createResponse = container.createItem(item, partitionKey, null);
        System.out.printf(
            "Created item %s (request charge: %.2f RU)%n",
            createResponse.getItem().getId(),
            createResponse.getRequestCharge());

        CosmosItemResponse<Item> readResponse =
            container.readItem(item.getId(), partitionKey, Item.class);
        Item readItem = readResponse.getItem();
        System.out.printf(
            "Read item: id=%s, category=%s, name=%s, quantity=%d%n",
            readItem.getId(),
            readItem.getCategory(),
            readItem.getName(),
            readItem.getQuantity());

        String category = "electronics";
        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions queryOptions = new CosmosQueryRequestOptions();
        queryOptions.setPartitionKey(new PartitionKey(category));

        CosmosPagedIterable<Item> queryResults =
            container.queryItems(query, queryOptions, Item.class);
        queryResults.forEach(result -> System.out.printf(
            "Query result: id=%s, name=%s, quantity=%d%n",
            result.getId(),
            result.getName(),
            result.getQuantity()));

        readItem.setQuantity(25);
        CosmosItemResponse<Item> replaceResponse =
            container.replaceItem(readItem, readItem.getId(), partitionKey, null);
        System.out.printf(
            "Replaced item %s with quantity %d (request charge: %.2f RU)%n",
            replaceResponse.getItem().getId(),
            replaceResponse.getItem().getQuantity(),
            replaceResponse.getRequestCharge());

        container.deleteItem(item.getId(), partitionKey, null);
        System.out.printf("Deleted item %s%n", item.getId());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable " + name + " is not set.");
        }
        return value;
    }

    private static void reportCosmosException(CosmosException exception) {
        System.err.printf(
            "Azure Cosmos DB request failed: status=%d, substatus=%d, activityId=%s, "
                + "requestCharge=%.2f RU, message=%s%n",
            exception.getStatusCode(),
            exception.getSubStatusCode(),
            exception.getActivityId(),
            exception.getRequestCharge(),
            exception.getMessage());

        if (exception.getStatusCode() == 409) {
            System.err.println("An item with the same id and partition key already exists.");
        } else if (exception.getStatusCode() == 429) {
            System.err.printf(
                "Request rate was limited. Retry after %s.%n",
                exception.getRetryAfterDuration());
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
