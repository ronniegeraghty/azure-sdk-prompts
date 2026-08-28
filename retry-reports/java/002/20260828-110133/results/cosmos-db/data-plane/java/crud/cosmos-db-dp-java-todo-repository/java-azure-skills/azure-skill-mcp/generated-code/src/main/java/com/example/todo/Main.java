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
        try (CosmosClientFactory factory = CosmosClientFactory.createFromEnvironment()) {
            runSynchronousDemo(new CosmosToDoRepository(factory.syncContainer()));
            runAsynchronousDemo(new AsyncCosmosToDoRepository(factory.asyncContainer()));
        }
    }

    private static void runSynchronousDemo(CosmosToDoRepository repository) {
        System.out.println("=== Synchronous repository ===");
        ToDoItem newItem = newItem("Sync Cosmos DB demo");

        CosmosOperationResult<ToDoItem> created = repository.create(newItem);
        printOperation("created", created);

        CosmosOperationResult<ToDoItem> read =
                repository.read(created.value().getId(), created.value().getCategory());
        printOperation("read", read);

        ToDoItem itemToUpdate = read.value();
        itemToUpdate.setCompleted(true);
        CosmosOperationResult<ToDoItem> updated = repository.update(itemToUpdate);
        printOperation("updated", updated);

        System.out.println("query pages:");
        repository.queryByCategory(CATEGORY, PAGE_SIZE, Main::printPage);

        CosmosOperationResult<Void> deleted =
                repository.delete(updated.value().getId(), updated.value().getCategory());
        System.out.printf("deleted requestCharge=%.2f RU%n%n", deleted.requestCharge());
    }

    private static void runAsynchronousDemo(AsyncCosmosToDoRepository repository) {
        System.out.println("=== Asynchronous repository ===");
        ToDoItem newItem = newItem("Async Cosmos DB demo");

        repository.create(newItem)
                .doOnNext(result -> printOperation("created", result))
                .flatMap(created -> repository.read(
                        created.value().getId(),
                        created.value().getCategory()))
                .doOnNext(result -> printOperation("read", result))
                .flatMap(read -> {
                    read.value().setCompleted(true);
                    return repository.update(read.value());
                })
                .doOnNext(result -> printOperation("updated", result))
                .flatMap(updated -> repository.queryByCategory(CATEGORY, PAGE_SIZE)
                        .doOnSubscribe(ignored -> System.out.println("query pages:"))
                        .doOnNext(Main::printPage)
                        .then(Mono.just(updated)))
                .flatMap(updated -> repository.delete(
                        updated.value().getId(),
                        updated.value().getCategory()))
                .doOnNext(result -> System.out.printf(
                        "deleted requestCharge=%.2f RU%n",
                        result.requestCharge()))
                .block();
    }

    private static ToDoItem newItem(String title) {
        return new ToDoItem(
                UUID.randomUUID().toString(),
                title,
                "This field is intentionally excluded from the indexing policy.",
                false,
                Instant.now(),
                CATEGORY);
    }

    private static void printOperation(
            String operation,
            CosmosOperationResult<ToDoItem> result) {
        System.out.printf(
                "%s requestCharge=%.2f RU result=%s%n",
                operation,
                result.requestCharge(),
                result.value());
    }

    private static void printPage(ToDoPage page) {
        System.out.printf(
                "page requestCharge=%.2f RU itemCount=%d continuationToken=%s%n",
                page.requestCharge(),
                page.items().size(),
                page.continuationToken());
        page.items().forEach(item -> System.out.println("  " + item));
    }
}
