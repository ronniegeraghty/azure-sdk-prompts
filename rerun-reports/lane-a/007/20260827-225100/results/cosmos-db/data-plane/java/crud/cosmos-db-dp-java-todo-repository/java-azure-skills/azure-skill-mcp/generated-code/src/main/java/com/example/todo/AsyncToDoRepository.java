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

import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;

public class AsyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncToDoRepository.class);

    private final CosmosAsyncContainer container;

    public AsyncToDoRepository(CosmosAsyncContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public Mono<RepositoryResponse<ToDoItem>> create(ToDoItem item) {
        return container.createItem(
                        item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions())
                .map(response -> response("create", item.getId(), response));
    }

    public Mono<RepositoryResponse<ToDoItem>> read(String id, String category) {
        return container.readItem(id, new PartitionKey(category), ToDoItem.class)
                .map(response -> response("read", id, response));
    }

    public Mono<RepositoryResponse<ToDoItem>> update(ToDoItem item, String expectedEtag) {
        Objects.requireNonNull(expectedEtag, "expectedEtag is required for a concurrency-safe update");
        CosmosItemRequestOptions options = new CosmosItemRequestOptions().setIfMatchETag(expectedEtag);
        return container.replaceItem(
                        item, item.getId(), new PartitionKey(item.getCategory()), options)
                .map(response -> response("update", item.getId(), response))
                .onErrorMap(
                        exception -> exception instanceof CosmosException cosmosException
                                && cosmosException.getStatusCode() == 412,
                        exception -> new OptimisticConcurrencyException(item.getId(), exception));
    }

    public Mono<RepositoryResponse<Void>> delete(String id, String category) {
        return container.deleteItem(
                        id, new PartitionKey(category), new CosmosItemRequestOptions())
                .map(response -> {
                    LOGGER.info("delete id={} requestCharge={} RU", id, response.getRequestCharge());
                    return new RepositoryResponse<Void>(
                            null, response.getETag(), response.getRequestCharge());
                });
    }

    public Flux<QueryPage<ToDoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            return Flux.error(new IllegalArgumentException("pageSize must be greater than zero"));
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category ORDER BY c.createdTimestamp",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options =
                new CosmosQueryRequestOptions().setPartitionKey(new PartitionKey(category));

        return Flux.defer(() -> {
            AtomicInteger pageNumber = new AtomicInteger();
            return container.queryItems(query, options, ToDoItem.class)
                    .byPage(pageSize)
                    .map(page -> {
                        int number = pageNumber.incrementAndGet();
                        LOGGER.info(
                                "query category={} page={} items={} requestCharge={} RU",
                                category, number, page.getResults().size(), page.getRequestCharge());
                        return new QueryPage<>(
                                page.getResults(),
                                page.getContinuationToken(),
                                page.getRequestCharge(),
                                number);
                    });
        });
    }

    private RepositoryResponse<ToDoItem> response(
            String operation, String id, CosmosItemResponse<ToDoItem> response) {
        LOGGER.info("{} id={} requestCharge={} RU", operation, id, response.getRequestCharge());
        return new RepositoryResponse<>(
                response.getItem(), response.getETag(), response.getRequestCharge());
    }
}
