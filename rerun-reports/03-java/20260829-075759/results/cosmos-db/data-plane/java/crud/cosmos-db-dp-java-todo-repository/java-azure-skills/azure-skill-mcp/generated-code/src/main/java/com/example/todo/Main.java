package com.example.todo;

import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final String CATEGORY = "demo";
    private static final int PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosConfiguration configuration = CosmosConfiguration.fromEnvironment()) {
            runSyncDemo(configuration.syncRepository());
            runAsyncDemo(configuration.asyncRepository()).block();
        }
    }

    private static void runSyncDemo(SyncToDoRepository repository) {
        System.out.println("\n=== Synchronous repository ===");
        String id = "sync-" + UUID.randomUUID();
        ToDoItem newItem = new ToDoItem(
                id, "Try synchronous Cosmos SDK", "CRUD repository demo",
                false, Instant.now(), CATEGORY);

        print("create", repository.create(newItem));
        RepositoryResult<ToDoItem> read = repository.read(id, CATEGORY);
        print("read", read);

        ToDoItem item = read.value();
        item.setCompleted(true);
        item.setTitle("Synchronous Cosmos SDK complete");
        print("update", repository.update(item));

        System.out.println("query by category (streamed page by page):");
        repository.queryByCategory(CATEGORY, PAGE_SIZE, Main::printPage);

        print("delete", repository.delete(id, CATEGORY));
    }

    private static Mono<Void> runAsyncDemo(AsyncToDoRepository repository) {
        System.out.println("\n=== Asynchronous repository ===");
        String id = "async-" + UUID.randomUUID();
        ToDoItem newItem = new ToDoItem(
                id, "Try asynchronous Cosmos SDK", "Reactive CRUD repository demo",
                false, Instant.now(), CATEGORY);

        return repository.create(newItem)
                .doOnNext(result -> print("create", result))
                .then(repository.read(id, CATEGORY))
                .doOnNext(result -> print("read", result))
                .map(RepositoryResult::value)
                .flatMap(item -> {
                    item.setCompleted(true);
                    item.setTitle("Asynchronous Cosmos SDK complete");
                    return repository.update(item);
                })
                .doOnNext(result -> print("update", result))
                .thenMany(repository.queryByCategory(CATEGORY, PAGE_SIZE))
                .doOnSubscribe(ignored ->
                        System.out.println("query by category (pages arrive asynchronously):"))
                .doOnNext(Main::printPage)
                .then(repository.delete(id, CATEGORY))
                .doOnNext(result -> print("delete", result))
                .then();
    }

    private static void print(String operation, RepositoryResult<?> result) {
        System.out.printf("%s: RU=%.2f result=%s%n",
                operation, result.requestCharge(), result.value());
    }

    private static void printPage(QueryPage page) {
        System.out.printf("page: RU=%.2f itemCount=%d%n",
                page.requestCharge(), page.items().size());
        page.items().forEach(item -> System.out.println("  " + item));
    }
}
