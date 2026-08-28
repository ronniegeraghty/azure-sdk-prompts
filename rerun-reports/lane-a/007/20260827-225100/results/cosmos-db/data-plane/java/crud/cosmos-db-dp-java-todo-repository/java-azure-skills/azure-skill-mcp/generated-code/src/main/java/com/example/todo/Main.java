package com.example.todo;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final int QUERY_PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosToDoFactory factory = CosmosToDoFactory.create()) {
            runSyncDemo(factory.syncRepository());
            runAsyncDemo(factory.asyncRepository());
        }
    }

    private static void runSyncDemo(SyncToDoRepository repository) {
        System.out.println("\n=== Synchronous CRUD ===");
        ToDoItem item = newItem("sync", "Run synchronous demo");

        RepositoryResponse<ToDoItem> created = repository.create(item);
        print("Created", created);

        RepositoryResponse<ToDoItem> read =
                repository.read(item.getId(), item.getCategory());
        print("Read", read);

        read.value().setCompleted(true);
        RepositoryResponse<ToDoItem> updated =
                repository.update(read.value(), read.etag());
        print("Updated", updated);

        System.out.println("Query results:");
        for (QueryPage<ToDoItem> page
                : repository.queryByCategory(item.getCategory(), QUERY_PAGE_SIZE)) {
            printPage(page);
        }

        RepositoryResponse<Void> deleted =
                repository.delete(item.getId(), item.getCategory());
        print("Deleted", deleted);
    }

    private static void runAsyncDemo(AsyncToDoRepository repository) {
        System.out.println("\n=== Asynchronous CRUD ===");
        ToDoItem item = newItem("async", "Run asynchronous demo");

        repository.create(item)
                .doOnNext(response -> print("Created", response))
                .flatMap(created -> repository.read(item.getId(), item.getCategory()))
                .doOnNext(response -> print("Read", response))
                .flatMap(read -> {
                    read.value().setCompleted(true);
                    return repository.update(read.value(), read.etag());
                })
                .doOnNext(response -> print("Updated", response))
                .thenMany(repository.queryByCategory(item.getCategory(), QUERY_PAGE_SIZE))
                .doOnSubscribe(ignored -> System.out.println("Query results:"))
                .doOnNext(Main::printPage)
                .then(repository.delete(item.getId(), item.getCategory()))
                .doOnNext(response -> print("Deleted", response))
                .block();
    }

    private static ToDoItem newItem(String category, String title) {
        return new ToDoItem(
                UUID.randomUUID().toString(),
                title,
                "Created by the Azure Cosmos DB Java repository sample",
                false,
                Instant.now(),
                category);
    }

    private static void print(String operation, RepositoryResponse<?> response) {
        System.out.printf(
                "%s: value=%s, etag=%s, requestCharge=%.2f RU%n",
                operation, response.value(), response.etag(), response.requestCharge());
    }

    private static void printPage(QueryPage<ToDoItem> page) {
        System.out.printf(
                "  page=%d, items=%d, requestCharge=%.2f RU, continuationToken=%s%n",
                page.pageNumber(),
                page.results().size(),
                page.requestCharge(),
                page.continuationToken());
        page.results().forEach(item -> System.out.println("    " + item));
    }
}
