package com.example.todo;

import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final String DATABASE_NAME =
            System.getenv().getOrDefault("COSMOS_DATABASE", "todo-demo");
    private static final String CONTAINER_NAME =
            System.getenv().getOrDefault("COSMOS_CONTAINER", "items");
    private static final int PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosToDoFactory factory =
                     CosmosToDoFactory.create(DATABASE_NAME, CONTAINER_NAME)) {
            runSyncDemo(factory.syncRepository());
            runAsyncDemo(factory.asyncRepository()).block();
        }
    }

    private static void runSyncDemo(SyncToDoRepository repository) {
        String category = "sync-demo";
        ToDoItem item = newItem("Sync repository demo", category);

        System.out.println("\n=== Synchronous CRUD ===");
        print("create", repository.create(item));

        RepositoryResult<ToDoItem> read = repository.read(item.getId(), category);
        print("read", read);
        read.value().setCompleted(true);
        read.value().setTitle("Sync repository demo - completed");
        print("update", repository.update(read.value()));

        int pageNumber = 0;
        for (RepositoryPage<ToDoItem> page
                : repository.queryByCategory(category, PAGE_SIZE)) {
            printPage(++pageNumber, page);
        }

        print("delete", repository.delete(item.getId(), category));
    }

    private static Mono<Void> runAsyncDemo(AsyncToDoRepository repository) {
        String category = "async-demo";
        ToDoItem item = newItem("Async repository demo", category);

        System.out.println("\n=== Asynchronous CRUD ===");
        return repository.create(item)
                .doOnNext(result -> print("create", result))
                .then(repository.read(item.getId(), category))
                .doOnNext(result -> print("read", result))
                .map(RepositoryResult::value)
                .flatMap(readItem -> {
                    readItem.setCompleted(true);
                    readItem.setTitle("Async repository demo - completed");
                    return repository.update(readItem);
                })
                .doOnNext(result -> print("update", result))
                .thenMany(repository.queryByCategory(category, PAGE_SIZE))
                .index()
                .doOnNext(indexed -> printPage(
                        Math.toIntExact(indexed.getT1() + 1), indexed.getT2()))
                .then(repository.delete(item.getId(), category))
                .doOnNext(result -> print("delete", result))
                .then();
    }

    private static ToDoItem newItem(String title, String category) {
        return new ToDoItem(
                UUID.randomUUID().toString(),
                title,
                "This field is stored but deliberately excluded from indexing.",
                false,
                Instant.now(),
                category);
    }

    private static void print(String operation, RepositoryResult<?> result) {
        System.out.printf(
                "%-6s RU=%6.2f result=%s%n",
                operation,
                result.requestCharge(),
                result.value());
    }

    private static void printPage(int pageNumber, RepositoryPage<ToDoItem> page) {
        System.out.printf(
                "query page=%d RU=%.2f items=%d continuationToken=%s%n",
                pageNumber,
                page.requestCharge(),
                page.items().size(),
                page.continuationToken() == null ? "<end>" : "<present>");
        page.items().forEach(item -> System.out.println("  " + item));
    }
}
