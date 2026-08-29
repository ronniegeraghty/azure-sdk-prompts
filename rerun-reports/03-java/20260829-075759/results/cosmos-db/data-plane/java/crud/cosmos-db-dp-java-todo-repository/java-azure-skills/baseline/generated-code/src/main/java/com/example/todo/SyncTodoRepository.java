package com.example.todo;

import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.azure.cosmos.CosmosException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;

public final class SyncTodoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncTodoRepository.class);

    private final CosmosContainer container;

    public SyncTodoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public OperationResult<TodoItem> create(TodoItem item) {
        CosmosItemResponse<TodoItem> response = container.createItem(
                item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());
        return logged("create", response);
    }

    public OperationResult<TodoItem> read(String id, String category) {
        CosmosItemResponse<TodoItem> response = container.readItem(
                id, new PartitionKey(category), TodoItem.class);
        return logged("read", response);
    }

    public OperationResult<TodoItem> update(TodoItem item) {
        if (item.getEtag() == null || item.getEtag().isBlank()) {
            throw new IllegalArgumentException("An ETag from a prior read is required for update");
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getEtag());
        try {
            CosmosItemResponse<TodoItem> response = container.replaceItem(
                    item, item.getId(), new PartitionKey(item.getCategory()), options);
            return logged("update", response);
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                LOGGER.warn("update conflict consumed {} RU", exception.getRequestCharge());
                throw new OptimisticConcurrencyException(item.getId(), exception);
            }
            throw exception;
        }
    }

    public OperationResult<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
                id, new PartitionKey(category), new CosmosItemRequestOptions());
        LOGGER.info("delete consumed {} RU", response.getRequestCharge());
        return new OperationResult<>(null, response.getRequestCharge());
    }

    public void queryByCategory(String category, int pageSize,
                                Consumer<QueryPage<TodoItem>> pageConsumer) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));

        int pageNumber = 0;
        for (FeedResponse<TodoItem> response
                : container.queryItems(query, options, TodoItem.class).iterableByPage(pageSize)) {
            pageNumber++;
            QueryPage<TodoItem> page = new QueryPage<>(
                    response.getResults(),
                    response.getRequestCharge(),
                    response.getContinuationToken());
            LOGGER.info("query page {} returned {} items and consumed {} RU",
                    pageNumber, page.items().size(), page.requestCharge());
            pageConsumer.accept(page);
        }
    }

    private OperationResult<TodoItem> logged(
            String operation, CosmosItemResponse<TodoItem> response) {
        LOGGER.info("{} consumed {} RU", operation, response.getRequestCharge());
        return new OperationResult<>(response.getItem(), response.getRequestCharge());
    }
}
