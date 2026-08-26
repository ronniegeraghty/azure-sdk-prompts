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

public class AsyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncToDoRepository.class);

    private final CosmosAsyncContainer container;

    public AsyncToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<RepositoryResult<ToDoItem>> create(ToDoItem item) {
        return container.createItem(
                        item,
                        new PartitionKey(item.getCategory()),
                        new CosmosItemRequestOptions())
                .map(response -> itemResult("create", response));
    }

    public Mono<RepositoryResult<ToDoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
                .map(response -> itemResult("read", response));
    }

    public Mono<RepositoryResult<ToDoItem>> update(ToDoItem item) {
        if (item.getETag() == null || item.getETag().isBlank()) {
            return Mono.error(new IllegalArgumentException(
                    "An ETag from a previous read is required for a safe update."));
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getETag());
        return container.replaceItem(
                        item,
                        item.getId(),
                        new PartitionKey(item.getCategory()),
                        options)
                .map(response -> itemResult("update", response))
                .onErrorMap(
                        error -> error instanceof CosmosException cosmosException
                                && cosmosException.getStatusCode() == 412,
                        error -> new ConcurrentUpdateException(item.getId(), error));
    }

    public Mono<RepositoryResult<Void>> delete(String id, String category) {
        return container.deleteItem(
                        id,
                        new PartitionKey(category),
                        new CosmosItemRequestOptions())
                .map(response -> {
                    LOGGER.info(
                            "delete id={} requestCharge={} RU",
                            id,
                            response.getRequestCharge());
                    return new RepositoryResult<Void>(null, response.getRequestCharge());
                });
    }

    public Flux<RepositoryPage<ToDoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException(
                    "pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));
        AtomicInteger pageNumber = new AtomicInteger();

        return container.queryItems(query, options, ToDoItem.class)
                .byPage(pageSize)
                .map(response -> toPage(response, category, pageNumber.incrementAndGet()));
    }

    private RepositoryPage<ToDoItem> toPage(
            FeedResponse<ToDoItem> response,
            String category,
            int pageNumber) {
        LOGGER.info(
                "query category={} page={} items={} requestCharge={} RU",
                category,
                pageNumber,
                response.getResults().size(),
                response.getRequestCharge());
        return new RepositoryPage<>(
                response.getResults(),
                response.getRequestCharge(),
                response.getContinuationToken());
    }

    private RepositoryResult<ToDoItem> itemResult(
            String operation,
            CosmosItemResponse<ToDoItem> response) {
        LOGGER.info(
                "{} id={} requestCharge={} RU",
                operation,
                response.getItem().getId(),
                response.getRequestCharge());
        return new RepositoryResult<>(response.getItem(), response.getRequestCharge());
    }
}
