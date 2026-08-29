package com.example.todo;

import com.azure.cosmos.CosmosException;
import com.azure.cosmos.CosmosContainer;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ConcurrentModificationException;
import java.util.List;
import java.util.Objects;
import java.util.function.Consumer;

public final class SyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncToDoRepository.class);
    private static final String QUERY_BY_CATEGORY =
            "SELECT * FROM todo t WHERE t.category = @category ORDER BY t.createdAt";

    private final CosmosContainer container;

    public SyncToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public RepositoryResult<ToDoItem> create(ToDoItem item) {
        validateItem(item);
        CosmosItemResponse<ToDoItem> response =
                container.createItem(item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());
        logCharge("create", response.getRequestCharge());
        return new RepositoryResult<>(response.getItem(), response.getRequestCharge());
    }

    public RepositoryResult<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response =
                container.readItem(id, new PartitionKey(category), ToDoItem.class);
        logCharge("read", response.getRequestCharge());
        return new RepositoryResult<>(response.getItem(), response.getRequestCharge());
    }

    public RepositoryResult<ToDoItem> update(ToDoItem item) {
        validateItem(item);
        requireEtag(item);

        CosmosItemRequestOptions options = new CosmosItemRequestOptions().setIfMatchETag(item.getEtag());
        try {
            CosmosItemResponse<ToDoItem> response =
                    container.replaceItem(item, item.getId(), new PartitionKey(item.getCategory()), options);
            logCharge("update", response.getRequestCharge());
            return new RepositoryResult<>(response.getItem(), response.getRequestCharge());
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                throw conflict(item, exception);
            }
            throw exception;
        }
    }

    public RepositoryResult<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response =
                container.deleteItem(id, new PartitionKey(category), new CosmosItemRequestOptions());
        logCharge("delete", response.getRequestCharge());
        return new RepositoryResult<>(null, response.getRequestCharge());
    }

    public void queryByCategory(String category, int pageSize, Consumer<QueryPage> pageConsumer) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }

        SqlQuerySpec query = new SqlQuerySpec(
                QUERY_BY_CATEGORY,
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));

        int pageNumber = 0;
        for (FeedResponse<ToDoItem> response
                : container.queryItems(query, options, ToDoItem.class).iterableByPage(pageSize)) {
            pageNumber++;
            double charge = response.getRequestCharge();
            LOGGER.info("query page={} items={} requestCharge={} RU",
                    pageNumber, response.getResults().size(), charge);
            pageConsumer.accept(new QueryPage(response.getResults(), charge));
        }
    }

    private static void validateItem(ToDoItem item) {
        Objects.requireNonNull(item, "item");
        Objects.requireNonNull(item.getId(), "item.id");
        Objects.requireNonNull(item.getCategory(), "item.category");
    }

    private static void requireEtag(ToDoItem item) {
        if (item.getEtag() == null || item.getEtag().isBlank()) {
            throw new IllegalArgumentException(
                    "An ETag from a prior read is required to update item " + item.getId());
        }
    }

    private static ConcurrentModificationException conflict(ToDoItem item, CosmosException cause) {
        return new ConcurrentModificationException(
                "Update conflict for item '%s' in category '%s': the item changed after it was read"
                        .formatted(item.getId(), item.getCategory()),
                cause);
    }

    private static void logCharge(String operation, double requestCharge) {
        LOGGER.info("{} requestCharge={} RU", operation, requestCharge);
    }
}
