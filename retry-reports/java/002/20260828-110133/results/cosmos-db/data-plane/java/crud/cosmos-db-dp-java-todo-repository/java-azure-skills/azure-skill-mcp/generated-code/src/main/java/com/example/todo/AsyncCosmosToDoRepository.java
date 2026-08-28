package com.example.todo;

import com.azure.cosmos.CosmosAsyncContainer;
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

import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;

public class AsyncCosmosToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncCosmosToDoRepository.class);

    private final CosmosAsyncContainer container;

    public AsyncCosmosToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<CosmosOperationResult<ToDoItem>> create(ToDoItem item) {
        return container.createItem(
                        item,
                        new PartitionKey(item.getCategory()),
                        new CosmosItemRequestOptions())
                .map(response -> itemResult("async create", response));
    }

    public Mono<CosmosOperationResult<ToDoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
                .map(response -> itemResult("async read", response));
    }

    public Mono<CosmosOperationResult<ToDoItem>> update(ToDoItem item) {
        requireETag(item);
        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getETag());
        return container.replaceItem(
                        item,
                        item.getId(),
                        new PartitionKey(item.getCategory()),
                        options)
                .map(response -> itemResult("async update", response))
                .onErrorMap(
                        com.azure.cosmos.CosmosException.class,
                        exception -> exception.getStatusCode() == 412
                                ? conflict(item, exception)
                                : exception);
    }

    public Mono<CosmosOperationResult<Void>> delete(String id, String category) {
        return container.deleteItem(
                        id,
                        new PartitionKey(category),
                        new CosmosItemRequestOptions())
                .map(response -> {
                    logCharge("async delete", response.getRequestCharge());
                    return new CosmosOperationResult<Void>(null, response.getRequestCharge());
                });
    }

    public Flux<ToDoPage> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException(
                    "pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));

        return Flux.defer(() -> {
            AtomicInteger pageNumber = new AtomicInteger();
            return container.queryItems(query, options, ToDoItem.class)
                    .byPage(pageSize)
                    .map(response -> {
                        int currentPage = pageNumber.incrementAndGet();
                        LOGGER.info(
                                "async query page={} items={} requestCharge={} RU",
                                currentPage,
                                response.getResults().size(),
                                response.getRequestCharge());
                        return new ToDoPage(
                                response.getResults(),
                                response.getRequestCharge(),
                                response.getContinuationToken());
                    });
        });
    }

    private static CosmosOperationResult<ToDoItem> itemResult(
            String operation,
            CosmosItemResponse<ToDoItem> response) {
        logCharge(operation, response.getRequestCharge());
        return new CosmosOperationResult<>(response.getItem(), response.getRequestCharge());
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
