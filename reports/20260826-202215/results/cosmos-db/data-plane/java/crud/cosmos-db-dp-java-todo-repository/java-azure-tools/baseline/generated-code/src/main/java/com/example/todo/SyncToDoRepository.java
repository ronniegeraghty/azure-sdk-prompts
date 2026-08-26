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
import java.util.NoSuchElementException;
import java.util.Objects;

public class SyncToDoRepository {
    private static final Logger LOGGER = LoggerFactory.getLogger(SyncToDoRepository.class);

    private final CosmosContainer container;

    public SyncToDoRepository(CosmosContainer container) {
        this.container = Objects.requireNonNull(container, "container");
    }

    public RepositoryResult<ToDoItem> create(ToDoItem item) {
        CosmosItemResponse<ToDoItem> response = container.createItem(
                item, new PartitionKey(item.getCategory()), new CosmosItemRequestOptions());
        return itemResult("create", response);
    }

    public RepositoryResult<ToDoItem> read(String id, String category) {
        CosmosItemResponse<ToDoItem> response = container.readItem(
                id, new PartitionKey(category), ToDoItem.class);
        return itemResult("read", response);
    }

    public RepositoryResult<ToDoItem> update(ToDoItem item) {
        if (item.getETag() == null || item.getETag().isBlank()) {
            throw new IllegalArgumentException(
                    "An ETag from a previous read is required for a safe update.");
        }

        CosmosItemRequestOptions options = new CosmosItemRequestOptions()
                .setIfMatchETag(item.getETag());
        try {
            CosmosItemResponse<ToDoItem> response = container.replaceItem(
                    item, item.getId(), new PartitionKey(item.getCategory()), options);
            return itemResult("update", response);
        } catch (CosmosException exception) {
            if (exception.getStatusCode() == 412) {
                throw new ConcurrentUpdateException(item.getId(), exception);
            }
            throw exception;
        }
    }

    public RepositoryResult<Void> delete(String id, String category) {
        CosmosItemResponse<Object> response = container.deleteItem(
                id, new PartitionKey(category), new CosmosItemRequestOptions());
        LOGGER.info("delete id={} requestCharge={} RU", id, response.getRequestCharge());
        return new RepositoryResult<>(null, response.getRequestCharge());
    }

    public Iterable<RepositoryPage<ToDoItem>> queryByCategory(String category, int pageSize) {
        if (pageSize <= 0) {
            throw new IllegalArgumentException("pageSize must be greater than zero");
        }

        SqlQuerySpec query = new SqlQuerySpec(
                "SELECT * FROM c WHERE c.category = @category",
                List.of(new SqlParameter("@category", category)));
        CosmosQueryRequestOptions options = new CosmosQueryRequestOptions()
                .setPartitionKey(new PartitionKey(category));
        Iterable<FeedResponse<ToDoItem>> pages = container
                .queryItems(query, options, ToDoItem.class)
                .iterableByPage(pageSize);

        return () -> mapPages(pages.iterator(), category);
    }

    private Iterator<RepositoryPage<ToDoItem>> mapPages(
            Iterator<FeedResponse<ToDoItem>> source,
            String category) {
        return new Iterator<>() {
            private int pageNumber;

            @Override
            public boolean hasNext() {
                return source.hasNext();
            }

            @Override
            public RepositoryPage<ToDoItem> next() {
                if (!hasNext()) {
                    throw new NoSuchElementException();
                }
                FeedResponse<ToDoItem> response = source.next();
                pageNumber++;
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
        };
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
