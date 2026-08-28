package com.example.todo;

import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.azure.cosmos.util.CosmosPagedIterable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;

public class CosmosToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(CosmosToDoRepository.class);

    private final CosmosContainer container;

    public CosmosToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public CosmosOperationResult<ToDoItem> create(ToDoItem item) {
        CosmosItemResponse<ToDoItem> response = container.createItem(
                item,
                new PartitionKey(item.getCategory()),
                new CosmosItemRequestOptions());
        logCharge("create", response.getRequestCharge());
        return new CosmosOperationResult<>(response.getItem(), response.getRequestCharge());
    }

    public CosmosOperationResult<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response = container.readItem(
                id,
                new PartitionKey(category),
                ToDoItem.class);
        logCharge("read", response.getRequestCharge());
        return new CosmosOperationResult<>(response.getItem(), response.getRequestCharge());
    }

    public CosmosOperationResult<ToDoItem> update(ToDoItem item) {
        requireETag(item);
        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getETag());
        try {
            CosmosItemResponse<ToDoItem> response = container.replaceItem(
                    item,
                    item.getId(),
                    new PartitionKey(item.getCategory()),
                    options);
            logCharge("update", response.getRequestCharge());
            return new CosmosOperationResult<>(response.getItem(), response.getRequestCharge());
        } catch (com.azure.cosmos.CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                throw conflict(item, exception);
            }
            throw exception;
        }
    }

    public CosmosOperationResult<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
                id,
                new PartitionKey(category),
                new CosmosItemRequestOptions());
        logCharge("delete", response.getRequestCharge());
        return new CosmosOperationResult<>(null, response.getRequestCharge());
    }

    public void queryByCategory(
            String category,
            int pageSize,
            Consumer<ToDoPage> pageConsumer) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }
        Objects.requireNonNull(pageConsumer, "pageConsumer");

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));
        CosmosPagedIterable<ToDoItem> results =
                container.queryItems(query, options, ToDoItem.class);

        int pageNumber = 0;
        for (FeedResponse<ToDoItem> response : results.iterableByPage(pageSize)) {
            pageNumber++;
            LOGGER.info(
                    "query page={} items={} requestCharge={} RU",
                    pageNumber,
                    response.getResults().size(),
                    response.getRequestCharge());
            pageConsumer.accept(new ToDoPage(
                    response.getResults(),
                    response.getRequestCharge(),
                    response.getContinuationToken()));
        }
    }

    private static void requireETag(ToDoItem item) {
        if (item.getETag() == null || item.getETag().isBlank()) {
            throw new IllegalArgumentException(
                    "An ETag from a prior create or read is required to update item " + item.getId());
        }
    }

    private static OptimisticConcurrencyException conflict(
            ToDoItem item,
            com.azure.cosmos.CosmosException cause) {
        return new OptimisticConcurrencyException(
                "ToDo item '" + item.getId()
                        + "' was modified by another process; read it again before updating",
                cause);
    }

    private static void logCharge(String operation, double requestCharge) {
        LOGGER.info("{} requestCharge={} RU", operation, requestCharge);
    }
}
