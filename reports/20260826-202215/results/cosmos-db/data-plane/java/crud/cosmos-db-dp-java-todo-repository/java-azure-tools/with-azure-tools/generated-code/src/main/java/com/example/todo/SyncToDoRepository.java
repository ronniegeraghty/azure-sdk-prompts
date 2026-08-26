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

import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;
import java.util.logging.Logger;

public final class SyncToDoRepository {
    private static final Logger LOGGER = Logger.getLogger(SyncToDoRepository.class.getName());

    private final CosmosContainer container;

    public SyncToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public RepositoryResponse<ToDoItem> create(ToDoItem item) {
        validateItem(item);
        CosmosItemResponse<ToDoItem> response = container.createItem(
            item,
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions());
        return itemResponse("create", response);
    }

    public RepositoryResponse<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response =
            container.readItem(id, new PartitionKey(category), ToDoItem.class);
        return itemResponse("read", response);
    }

    public RepositoryResponse<ToDoItem> update(ToDoItem item) {
        validateItem(item);
        if (item.getEtag() == null || item.getEtag().isBlank()) {
            throw new IllegalArgumentException(
                "An ETag is required for update; read the item before updating it");
        }

        CosmosItemRequestOptions options =
            new CosmosItemRequestOptions().setIfMatchETag(item.getEtag());
        try {
            CosmosItemResponse<ToDoItem> response = container.replaceItem(
                item,
                item.getId(),
                new PartitionKey(item.getCategory()),
                options);
            return itemResponse("update", response);
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                logCharge("update conflict", exception.getRequestCharge());
                throw new OptimisticConcurrencyException(item.getId(), exception);
            }
            throw exception;
        }
    }

    public RepositoryResponse<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
            id,
            new PartitionKey(category),
            new CosmosItemRequestOptions());
        logCharge("delete", response.getRequestCharge());
        return new RepositoryResponse<>(null, response.getRequestCharge());
    }

    public void queryByCategory(
        String category,
        int pageSize,
        Consumer<QueryPage<ToDoItem>> pageConsumer
    ) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }
        Objects.requireNonNull(pageConsumer, "pageConsumer");

        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options =
            new CosmosQueryRequestOptions().setPartitionKey(new PartitionKey(category));

        int pageNumber = 0;
        for (FeedResponse<ToDoItem> response :
            container.queryItems(query, options, ToDoItem.class).iterableByPage(pageSize)) {
            pageNumber++;
            logQueryPage(pageNumber, response);
            pageConsumer.accept(new QueryPage<>(
                pageNumber,
                response.getResults(),
                response.getRequestCharge(),
                response.getContinuationToken()));
        }
    }

    private RepositoryResponse<ToDoItem> itemResponse(
        String operation,
        CosmosItemResponse<ToDoItem> response
    ) {
        ToDoItem item = response.getItem();
        item.setEtag(response.getETag());
        logCharge(operation, response.getRequestCharge());
        return new RepositoryResponse<>(item, response.getRequestCharge());
    }

    private static void validateItem(ToDoItem item) {
        Objects.requireNonNull(item, "item");
        Objects.requireNonNull(item.getId(), "item.id");
        Objects.requireNonNull(item.getCategory(), "item.category");
    }

    private static void logCharge(String operation, double requestCharge) {
        LOGGER.info(() -> "%s consumed %.2f RUs".formatted(operation, requestCharge));
    }

    private static void logQueryPage(int pageNumber, FeedResponse<ToDoItem> response) {
        LOGGER.info(() -> "query page %d returned %d items and consumed %.2f RUs"
            .formatted(pageNumber, response.getResults().size(), response.getRequestCharge()));
    }
}
