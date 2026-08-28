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
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Iterator;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.atomic.AtomicInteger;

public class SyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncToDoRepository.class);

    private final CosmosContainer container;

    public SyncToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public RepositoryResponse<ToDoItem> create(ToDoItem item) {
        CosmosItemResponse<ToDoItem> response = container.createItem(
                item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());
        return response("create", item.getId(), response);
    }

    public RepositoryResponse<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response =
                container.readItem(id, new PartitionKey(category), ToDoItem.class);
        return response("read", id, response);
    }

    public RepositoryResponse<ToDoItem> update(ToDoItem item, String expectedEtag) {
        Objects.requireNonNull(expectedEtag, "expectedEtag is required for a concurrency-safe update");
        CosmosItemRequestOptions options = new CosmosItemRequestOptions().setIfMatchETag(expectedEtag);
        try {
            CosmosItemResponse<ToDoItem> response = container.replaceItem(
                    item, item.getId(), new PartitionKey(item.getCategory()), options);
            return response("update", item.getId(), response);
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                throw new OptimisticConcurrencyException(item.getId(), exception);
            }
            throw exception;
        }
    }

    public RepositoryResponse<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
                id, new PartitionKey(category), new CosmosItemRequestOptions());
        LOGGER.info("delete id={} requestCharge={} RU", id, response.getRequestCharge());
        return new RepositoryResponse<>(null, response.getETag(), response.getRequestCharge());
    }

    public Iterable<QueryPage<ToDoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category ORDER BY c.createdTimestamp",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options =
                new CosmosQueryRequestOptions().setPartitionKey(new PartitionKey(category));

        Iterable<FeedResponse<ToDoItem>> pages =
                container.queryItems(query, options, ToDoItem.class).iterableByPage(pageSize);

        return () -> {
            Iterator<FeedResponse<ToDoItem>> iterator = pages.iterator();
            AtomicInteger pageNumber = new AtomicInteger();
            return new Iterator<>() {
                @Override
                public boolean hasNext() {
                    return iterator.hasNext();
                }

                @Override
                public QueryPage<ToDoItem> next() {
                    FeedResponse<ToDoItem> page = iterator.next();
                    int number = pageNumber.incrementAndGet();
                    LOGGER.info(
                            "query category={} page={} items={} requestCharge={} RU",
                            category, number, page.getResults().size(), page.getRequestCharge());
                    return new QueryPage<>(
                            page.getResults(),
                            page.getContinuationToken(),
                            page.getRequestCharge(),
                            number);
                }
            };
        };
    }

    private RepositoryResponse<ToDoItem> response(
            String operation, String id, CosmosItemResponse<ToDoItem> response) {
        LOGGER.info("{} id={} requestCharge={} RU", operation, id, response.getRequestCharge());
        return new RepositoryResponse<>(
                response.getItem(), response.getETag(), response.getRequestCharge());
    }
}
