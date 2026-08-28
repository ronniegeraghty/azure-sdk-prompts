package com.example.todo;

import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;

public final class SyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncToDoRepository.class);

    private final CosmosContainer container;

    public SyncToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public CosmosItemResponse<ToDoItem> create(ToDoItem item) {
        validateItem(item);
        CosmosItemResponse<ToDoItem> response = container.createItem(
            item,
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        );
        logCharge("create", item.getId(), response.getRequestCharge());
        return response;
    }

    public CosmosItemResponse<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response = container.readItem(
            requireText(id, "id"),
            new PartitionKey(requireText(category, "category")),
            ToDoItem.class
        );
        logCharge("read", id, response.getRequestCharge());
        return response;
    }

    public CosmosItemResponse<ToDoItem> update(ToDoItem item) {
        validateItem(item);
        if (item.getETag() == null || item.getETag().isBlank()) {
            throw new IllegalArgumentException("An ETag from a prior read is required for update");
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
            .setIfMatchETag(item.getETag());
        try {
            CosmosItemResponse<ToDoItem> response = container.replaceItem(
                item,
                item.getId(),
                new PartitionKey(item.getCategory()),
                options
            );
            logCharge("update", item.getId(), response.getRequestCharge());
            return response;
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                throw conflict(item, exception);
            }
            throw exception;
        }
    }

    public CosmosItemResponse<Object> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
            requireText(id, "id"),
            new PartitionKey(requireText(category, "category")),
            new CosmosItemRequestOptions()
        );
        logCharge("delete", id, response.getRequestCharge());
        return response;
    }

    public void queryByCategory(
        String category,
        int pageSize,
        Consumer<FeedResponse<ToDoItem>> pageConsumer
    ) {
        requireText(category, "category");
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }
        Objects.requireNonNull(pageConsumer, "pageConsumer");

        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category))
        );
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
            .setPartitionKey(new PartitionKey(category));

        int pageNumber = 0;
        for (FeedResponse<ToDoItem> page
            : container.queryItems(query, options, ToDoItem.class).iterableByPage(pageSize)) {
            pageNumber++;
            LOGGER.info(
                "query category={} page={} items={} requestCharge={} RU",
                category,
                pageNumber,
                page.getResults().size(),
                page.getRequestCharge()
            );
            pageConsumer.accept(page);
        }
    }

    private static OptimisticConcurrencyException conflict(ToDoItem item, CosmosException cause) {
        return new OptimisticConcurrencyException(
            "Update conflict for ToDo item '" + item.getId()
                + "': it was modified after it was read; read the latest item and retry",
            cause
        );
    }

    private static void logCharge(String operation, String id, double requestCharge) {
        LOGGER.info("{} id={} requestCharge={} RU", operation, id, requestCharge);
    }

    private static void validateItem(ToDoItem item) {
        Objects.requireNonNull(item, "item");
        requireText(item.getId(), "item.id");
        requireText(item.getCategory(), "item.category");
    }

    private static String requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}
