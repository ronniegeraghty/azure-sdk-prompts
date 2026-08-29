package com.example.todo.repository;

import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import com.example.todo.model.TodoItem;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.logging.Logger;

public final class AsyncTodoRepository {
    private static final Logger LOGGER = Logger.getLogger(AsyncTodoRepository.class.getName());

    private final CosmosAsyncContainer container;

    public AsyncTodoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<OperationResult<TodoItem>> create(TodoItem item) {
        validateItem(item);
        return container.createItem(
                        item,
                        new PartitionKey(item.getCategory()),
                        new CosmosItemRequestOptions())
                .doOnNext(response -> logCharge("async-create", item.getId(), response.getRequestCharge()))
                .map(AsyncTodoRepository::toResult);
    }

    public Mono<OperationResult<TodoItem>> read(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        return container.readItem(id, new PartitionKey(category), TodoItem.class)
                .doOnNext(response -> logCharge("async-read", id, response.getRequestCharge()))
                .map(AsyncTodoRepository::toResult);
    }

    public Mono<OperationResult<TodoItem>> update(TodoItem item, String expectedEtag) {
        validateItem(item);
        requireText(expectedEtag, "expectedEtag");

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(expectedEtag);
        return container.replaceItem(
                        item,
                        item.getId(),
                        new PartitionKey(item.getCategory()),
                        options)
                .doOnNext(response -> logCharge("async-update", item.getId(), response.getRequestCharge()))
                .doOnError(CosmosException.class, exception ->
                        logCharge("async-update-failed", item.getId(), exception.getRequestCharge()))
                .onErrorMap(
                        error -> error instanceof CosmosException exception
                                && exception.getStatusCode() == 412,
                        error -> new ConcurrentUpdateException(item.getId(), error))
                .map(AsyncTodoRepository::toResult);
    }

    public Mono<OperationResult<Void>> delete(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        return container.deleteItem(
                        id,
                        new PartitionKey(category),
                        new CosmosItemRequestOptions())
                .doOnNext(response -> logCharge("async-delete", id, response.getRequestCharge()))
                .map(response -> new OperationResult<>(
                        null,
                        response.getETag(),
                        response.getRequestCharge()));
    }

    public Flux<QueryPage<TodoItem>> queryByCategory(String category, int pageSize) {
        requireText(category, "category");
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));

        return Flux.defer(() -> {
            AtomicInteger pageNumber = new AtomicInteger();
            return container.queryItems(query, options, TodoItem.class)
                    .byPage(pageSize)
                    .map(response -> {
                        int currentPage = pageNumber.incrementAndGet();
                        LOGGER.info(() -> "async-query category=" + category
                                + " page=" + currentPage
                                + " items=" + response.getResults().size()
                                + " RU=" + formatCharge(response.getRequestCharge()));
                        return new QueryPage<>(
                                response.getResults(),
                                response.getContinuationToken(),
                                response.getRequestCharge(),
                                currentPage);
                    });
        });
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
