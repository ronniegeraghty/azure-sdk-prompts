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
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;

public final class AsyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncToDoRepository.class);

    private final CosmosAsyncContainer container;

    public AsyncToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<CosmosItemResponse<ToDoItem>> create(ToDoItem item) {
        validateItem(item);
        return container.createItem(
            item,
            new PartitionKey(item.getCategory()),
            new CosmosItemRequestOptions()
        ).doOnNext(response -> logCharge("create", item.getId(), response.getRequestCharge()));
    }

    public Mono<CosmosItemResponse<ToDoItem>> read(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
            .doOnNext(response -> logCharge("read", id, response.getRequestCharge()));
    }

    public Mono<CosmosItemResponse<ToDoItem>> update(ToDoItem item) {
        validateItem(item);
        if (item.getETag() == null || item.getETag().isBlank()) {
            return Mono.error(new IllegalArgumentException(
                "An ETag from a prior read is required for update"
            ));
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
            .setIfMatchETag(item.getETag());
        return container.replaceItem(
            item,
            item.getId(),
            new PartitionKey(item.getCategory()),
            options
        )
            .doOnNext(response -> logCharge("update", item.getId(), response.getRequestCharge()))
            .onErrorMap(
                exception -> exception instanceof CosmosException cosmosException
                    && cosmosException.getStatusCode() == 412,
                exception -> conflict(item, (CosmosException) exception)
            );
    }

    public Mono<CosmosItemResponse<Object>> delete(String id, String category) {
        requireText(id, "id");
        requireText(category, "category");
        return container.deleteItem(
            id,
            new PartitionKey(category),
            new CosmosItemRequestOptions()
        ).doOnNext(response -> logCharge("delete", id, response.getRequestCharge()));
    }

    public Flux<FeedResponse<ToDoItem>> queryByCategory(String category, int pageSize) {
        requireText(category, "category");
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
            "SELECT * FROM c WHERE c.category = @category",
            List.of(new SqlParameter("@category", category))
        );
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
            .setPartitionKey(new PartitionKey(category));
        return Flux.defer(() -> {
            AtomicInteger pageNumber = new AtomicInteger();
            return container.queryItems(query, options, ToDoItem.class)
                .byPage(pageSize)
                .doOnNext(page -> LOGGER.info(
                    "query category={} page={} items={} requestCharge={} RU",
                    category,
                    pageNumber.incrementAndGet(),
                    page.getResults().size(),
                    page.getRequestCharge()
                ));
        });
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
