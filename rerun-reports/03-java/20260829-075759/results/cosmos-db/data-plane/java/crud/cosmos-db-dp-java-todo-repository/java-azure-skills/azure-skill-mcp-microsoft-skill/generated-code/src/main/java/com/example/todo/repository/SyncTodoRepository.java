package com.example.todo.repository;

import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.example.todo.model.TodoItem;

import java.util.Iterator;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.logging.Logger;

public final class SyncTodoRepository {
    private static final Logger LOGGER = Logger.getLogger(SyncTodoRepository.class.getName());

    private final CosmosContainer container;

    public SyncTodoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public OperationResult<TodoItem> create(TodoItem item) {
        validateItem(item);
        CosmosItemResponse<TodoItem> response = container.createItem(
                item,
                new PartitionKey(item.getCategory()),
                new CosmosItemRequestOptions());
        logCharge("create", item.getId(), response.getRequestCharge());
        return toResult(response);
    }

    public OperationResult<TodoItem> read(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        CosmosItemResponse<TodoItem> response = container.readItem(
                id,
                new PartitionKey(category),
                TodoItem.class);
        logCharge("read", id, response.getRequestCharge());
        return toResult(response);
    }

    public OperationResult<TodoItem> update(TodoItem item, String expectedEtag) {
        validateItem(item);
        requireText(expectedEtag, "expectedEtag");

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(expectedEtag);
        try {
            CosmosItemResponse<TodoItem> response = container.replaceItem(
                    item,
                    item.getId(),
                    new PartitionKey(item.getCategory()),
                    options);
            logCharge("update", item.getId(), response.getRequestCharge());
            return toResult(response);
        } catch (CosmosException exception) {
            logCharge("update-failed", item.getId(), exception.getRequestCharge());
            if (exception.getStatusCode() == 412) {
                throw new ConcurrentUpdateException(item.getId(), exception);
            }
            throw exception;
        }
    }

    public OperationResult<Void> delete(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        CosmosItemResponse<Object> response = container.deleteItem(
                id,
                new PartitionKey(category),
                new CosmosItemRequestOptions());
        logCharge("delete", id, response.getRequestCharge());
        return new OperationResult<>(null, response.getETag(), response.getRequestCharge());
    }

    public Iterable<QueryPage<TodoItem>> queryByCategory(String category, int pageSize) {
        requireText(category, "category");
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));
        Iterable<FeedResponse<TodoItem>> pages = container
                .queryItems(query, options, TodoItem.class)
                .iterableByPage(pageSize);

        return () -> {
            Iterator<FeedResponse<TodoItem>> iterator = pages.iterator();
            AtomicInteger pageNumber = new AtomicInteger();
            return new Iterator<>() {
                @Override
                public boolean hasNext() {
                    return iterator.hasNext();
                }

                @Override
                public QueryPage<TodoItem> next() {
                    FeedResponse<TodoItem> response = iterator.next();
                    int currentPage = pageNumber.incrementAndGet();
                    LOGGER.info(() -> "query category=" + category
                            + " page=" + currentPage
                            + " items=" + response.getResults().size()
                            + " RU=" + formatCharge(response.getRequestCharge()));
                    return new QueryPage<>(
                            response.getResults(),
                            response.getContinuationToken(),
                            response.getRequestCharge(),
                            currentPage);
                }
            };
        };
    }

    private static OperationResult<TodoItem> toResult(CosmosItemResponse<TodoItem> response) {
        return new OperationResult<>(
                response.getItem(),
                response.getETag(),
                response.getRequestCharge());
    }

    private static void validateItem(TodoItem item) {
        Objects.requireNonNull(item, "item");
        requireText(item.getId(), "item.id");
        requireText(item.getCategory(), "item.category");
    }

    private static void requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
    }

    private static void logCharge(String operation, String id, double charge) {
        LOGGER.info(() -> operation + " id=" + id + " RU=" + formatCharge(charge));
    }

    private static String formatCharge(double charge) {
        return String.format("%.2f", charge);
    }
}
