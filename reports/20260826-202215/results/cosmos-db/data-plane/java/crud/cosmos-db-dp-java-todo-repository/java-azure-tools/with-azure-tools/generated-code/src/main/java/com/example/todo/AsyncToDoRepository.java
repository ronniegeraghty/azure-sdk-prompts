package com.example.todo;

import com.azure.cosmos.CosmosAsyncContainer;
import com.azure.cosmos.CosmosException;
import com.azure.cosmos.models.CosmosItemRequestOptions;
import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.CosmosQueryRequestOptions;
import com.azure.cosmos.models.FeedResponse;
import com.azure.cosmos.models.PartitionKey;
import com.azure.cosmos.models.SqlParameter;
import com.azure.cosmos.models.SqlQuerySpec;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.logging.Logger;

public final class AsyncToDoRepository {
    private static final Logger LOGGER = Logger.getLogger(AsyncToDoRepository.class.getName());

    private final CosmosAsyncContainer container;

    public AsyncToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<RepositoryResponse<ToDoItem>> create(ToDoItem item) {
        validateItem(item);
        return container.createItem(
                item,
                new PartitionKey(item.getCategory()),
                new CosmosItemRequestOptions())
            .map(response -> itemResponse("async create", response));
    }

    public Mono<RepositoryResponse<ToDoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
            .map(response -> itemResponse("async read", response));
    }

    public Mono<RepositoryResponse<ToDoItem>> update(ToDoItem item) {
        validateItem(item);
        if (item.getEtag() == null || item.getEtag().isBlank()) {
            return Mono.error(new IllegalArgumentException(
                "An ETag is required for update; read the item before updating it"));
        }

        CosmosItemRequestOptions options =
            new CosmosItemRequestOptions().setIfMatchETag(item.getEtag());
        return container.replaceItem(
                item,
                item.getId(),
                new PartitionKey(item.getCategory()),
                options)
            .map(response -> itemResponse("async update", response))
            .doOnError(exception -> {
                if (exception instanceof CosmosException cosmosException
                    && cosmosException.getStatusCode() == 412) {
                    logCharge("async update conflict", cosmosException.getRequestCharge());
                }
            })
            .onErrorMap(
                exception -> exception instanceof CosmosException cosmosException
                    && cosmosException.getStatusCode() == 412,
                exception -> new OptimisticConcurrencyException(item.getId(), exception));
    }

    public Mono<RepositoryResponse<Void>> delete(String id, String category) {
        return container.deleteItem(
                id,
                new PartitionKey(category),
                new CosmosItemRequestOptions())
            .map(response -> {
                logCharge("async delete", response.getRequestCharge());
                return new RepositoryResponse<Void>(null, response.getRequestCharge());
            });
    }

    public Flux<QueryPage<ToDoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options =
            new CosmosQueryRequestOptions().setPartitionKey(new PartitionKey(category));
        AtomicInteger pageNumber = new AtomicInteger();

        return container.queryItems(query, options, ToDoItem.class)
            .byPage(pageSize)
            .map(response -> {
                int currentPage = pageNumber.incrementAndGet();
                logQueryPage(currentPage, response);
                return new QueryPage<>(
                    currentPage,
                    response.getResults(),
                    response.getRequestCharge(),
                    response.getContinuationToken());
            });
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
        LOGGER.info(() -> "async query page %d returned %d items and consumed %.2f RUs"
            .formatted(pageNumber, response.getResults().size(), response.getRequestCharge()));
    }
}
