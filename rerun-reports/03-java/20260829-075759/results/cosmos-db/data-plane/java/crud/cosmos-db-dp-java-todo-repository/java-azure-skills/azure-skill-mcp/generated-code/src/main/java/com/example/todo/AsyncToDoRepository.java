package com.example.todo;

import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.ConcurrentModificationException;
import java.util.List;
import java.util.Objects;

public final class AsyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncToDoRepository.class);
    private static final String QUERY_BY_CATEGORY =
            "SELECT * FROM todo t WHERE t.category = @category ORDER BY t.createdAt";

    private final CosmosAsyncContainer container;

    public AsyncToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<RepositoryResult<ToDoItem>> create(ToDoItem item) {
        validateItem(item);
        return container.createItem(
                        item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions())
                .map(response -> result("create", response));
    }

    public Mono<RepositoryResult<ToDoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
                .map(response -> result("read", response));
    }

    public Mono<RepositoryResult<ToDoItem>> update(ToDoItem item) {
        validateItem(item);
        requireEtag(item);

        CosmosItemRequestOptions options = new CosmosItemRequestOptions().setIfMatchETag(item.getEtag());
        return container.replaceItem(
                        item, item.getId(), new PartitionKey(item.getCategory()), options)
                .map(response -> result("update", response))
                .onErrorMap(
                        error -> error instanceof CosmosException cosmos && cosmos.getStatusCode() == 412,
                        error -> conflict(item, error));
    }

    public Mono<RepositoryResult<Void>> delete(String id, String category) {
        return container.deleteItem(
                        id, new PartitionKey(category), new CosmosItemRequestOptions())
                .map(response -> {
                    logCharge("delete", response.getRequestCharge());
                    return new RepositoryResult<Void>(null, response.getRequestCharge());
                });
    }

    public Flux<QueryPage> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                QUERY_BY_CATEGORY,
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));

        return container.queryItems(query, options, ToDoItem.class)
                .byPage(pageSize)
                .index()
                .map(indexedResponse -> {
                    long pageNumber = indexedResponse.getT1() + 1;
                    var response = indexedResponse.getT2();
                    LOGGER.info("query page={} items={} requestCharge={} RU",
                            pageNumber, response.getResults().size(), response.getRequestCharge());
                    return new QueryPage(response.getResults(), response.getRequestCharge());
                });
    }

    private static RepositoryResult<ToDoItem> result(
            String operation, CosmosItemResponse<ToDoItem> response) {
        logCharge(operation, response.getRequestCharge());
        return new RepositoryResult<>(response.getItem(), response.getRequestCharge());
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

    private static ConcurrentModificationException conflict(ToDoItem item, Throwable cause) {
        return new ConcurrentModificationException(
                "Update conflict for item '%s' in category '%s': the item changed after it was read"
                        .formatted(item.getId(), item.getCategory()),
                cause);
    }

    private static void logCharge(String operation, double requestCharge) {
        LOGGER.info("{} requestCharge={} RU", operation, requestCharge);
    }
}
