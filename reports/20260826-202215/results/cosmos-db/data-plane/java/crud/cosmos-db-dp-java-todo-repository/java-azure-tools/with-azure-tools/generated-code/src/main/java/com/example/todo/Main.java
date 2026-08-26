package com.example.todo;

import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final int PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosConfiguration configuration = CosmosConfiguration.createFromEnvironment()) {
            runSyncDemo(configuration.syncRepository());
            runAsyncDemo(configuration.asyncRepository()).block();
        }
    }

    private static void runSyncDemo(SyncToDoRepository repository) {
        System.out.println("=== Synchronous repository ===");
        String category = "sync-demo-" + UUID.randomUUID();
        ToDoItem item = newItem("Sync Cosmos DB demo", category);
        ToDoItem secondItem = newItem("Sync query item 2", category);
        ToDoItem thirdItem = newItem("Sync query item 3", category);

        RepositoryResponse<ToDoItem> created = repository.create(item);
        printOperation("create", created);

        RepositoryResponse<ToDoItem> read =
            repository.read(item.getId(), item.getCategory());
        printOperation("read", read);

        ToDoItem current = read.value();
        current.setCompleted(true);
        RepositoryResponse<ToDoItem> updated = repository.update(current);
        printOperation("update", updated);

        printOperation("create query item", repository.create(secondItem));
        printOperation("create query item", repository.create(thirdItem));
        repository.queryByCategory(category, PAGE_SIZE, Main::printPage);

        RepositoryResponse<Void> deleted =
            repository.delete(item.getId(), category);
        printOperation("delete", deleted);
        printOperation("delete query item", repository.delete(secondItem.getId(), category));
        printOperation("delete query item", repository.delete(thirdItem.getId(), category));
    }

    private static Mono<Void> runAsyncDemo(AsyncToDoRepository repository) {
        System.out.println("=== Asynchronous repository ===");
        String category = "async-demo-" + UUID.randomUUID();
        ToDoItem item = newItem("Async Cosmos DB demo", category);
        ToDoItem secondItem = newItem("Async query item 2", category);
        ToDoItem thirdItem = newItem("Async query item 3", category);

        return repository.create(item)
            .doOnNext(response -> printOperation("async create", response))
            .then(repository.read(item.getId(), item.getCategory()))
            .doOnNext(response -> printOperation("async read", response))
            .map(RepositoryResponse::value)
            .flatMap(current -> {
                current.setCompleted(true);
                return repository.update(current);
            })
            .doOnNext(response -> printOperation("async update", response))
            .then(Mono.when(
                repository.create(secondItem)
                    .doOnNext(response -> printOperation("async create query item", response)),
                repository.create(thirdItem)
                    .doOnNext(response -> printOperation("async create query item", response))))
            .thenMany(repository.queryByCategory(category, PAGE_SIZE))
            .doOnNext(Main::printPage)
            .then(repository.delete(item.getId(), category))
            .doOnNext(response -> printOperation("async delete", response))
            .then(repository.delete(secondItem.getId(), category))
            .doOnNext(response -> printOperation("async delete query item", response))
            .then(repository.delete(thirdItem.getId(), category))
            .doOnNext(response -> printOperation("async delete query item", response))
            .then();
    }

    private static ToDoItem newItem(String title, String category) {
        return new ToDoItem(
            UUID.randomUUID().toString(),
            title,
            "This field is deliberately excluded from the Cosmos DB index.",
            false,
            Instant.now(),
            category);
    }

    private static void printOperation(String operation, RepositoryResponse<?> response) {
        System.out.printf(
            "%s: result=%s, requestCharge=%.2f RUs%n",
            operation,
            response.value(),
            response.requestCharge());
    }

    private static void printPage(QueryPage<ToDoItem> page) {
        System.out.printf(
            "page %d: %d item(s), requestCharge=%.2f RUs, continuationToken=%s%n",
            page.pageNumber(),
            page.items().size(),
            page.requestCharge(),
            page.continuationToken() == null ? "<end>" : "<available>");
        page.items().forEach(result -> System.out.println("  " + result));
    }
}
