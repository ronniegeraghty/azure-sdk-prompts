package com.example.todo;

import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
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

public final class AsyncTodoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncTodoRepository.class);

    private final CosmosAsyncContainer container;

    public AsyncTodoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<OperationResult<TodoItem>> create(TodoItem item) {
        return container.createItem(
                        item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions())
                .map(response -> logged("create", response.getItem(), response.getRequestCharge()));
    }

    public Mono<OperationResult<TodoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), TodoItem.class)
                .map(response -> logged("read", response.getItem(), response.getRequestCharge()));
    }

    public Mono<OperationResult<TodoItem>> update(TodoItem item) {
        if (item.getEtag() == null || item.getEtag().isBlank()) {
            return Mono.error(
                    new IllegalArgumentException("An ETag from a prior read is required for update"));
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getEtag());
        return container.replaceItem(
                        item, item.getId(), new PartitionKey(item.getCategory()), options)
                .map(response -> logged("update", response.getItem(), response.getRequestCharge()))
                .doOnError(
                        error -> error instanceof CosmosException cosmos
                                && cosmos.getStatusCode() == 412,
                        error -> LOGGER.warn("update conflict consumed {} RU",
                                ((CosmosException) error).getRequestCharge()))
                .onErrorMap(
                        error -> error instanceof CosmosException cosmos
                                && cosmos.getStatusCode() == 412,
                        error -> new OptimisticConcurrencyException(item.getId(), error));
    }

    public Mono<OperationResult<Void>> delete(String id, String category) {
        return container.deleteItem(
                        id, new PartitionKey(category), new CosmosItemRequestOptions())
                .map(response -> {
                    LOGGER.info("delete consumed {} RU", response.getRequestCharge());
                    return new OperationResult<>(null, response.getRequestCharge());
                });
    }

    public Flux<QueryPage<TodoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));
        AtomicInteger pageNumber = new AtomicInteger();

        return container.queryItems(query, options, TodoItem.class)
                .byPage(pageSize)
                .map(response -> new QueryPage<>(
                        response.getResults(),
                        response.getRequestCharge(),
                        response.getContinuationToken()))
                .doOnNext(page -> LOGGER.info(
                        "query page {} returned {} items and consumed {} RU",
                        pageNumber.incrementAndGet(), page.items().size(), page.requestCharge()));
    }

    private OperationResult<TodoItem> logged(
            String operation, TodoItem item, double requestCharge) {
        LOGGER.info("{} consumed {} RU", operation, requestCharge);
        return new OperationResult<>(item, requestCharge);
    }
}
