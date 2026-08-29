package com.example.todo;

import com.azure.cosmos.CosmosAsyncClient;
import com.azure.cosmos.CosmosClient;
import com.azure.cosmos.CosmosContainer;
import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final int PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        runSyncDemo();
        runAsyncDemo();
    }

    private static void runSyncDemo() {
        System.out.println("=== Synchronous repository ===");
        try (CosmosClient client = CosmosConfiguration.createSyncClient()) {
            CosmosContainer container = CosmosConfiguration.initializeSync(
                    client,
                    CosmosConfiguration.DEFAULT_DATABASE,
                    CosmosConfiguration.DEFAULT_CONTAINER);
            SyncTodoRepository repository = new SyncTodoRepository(container);
            TodoItem item = newItem("sync-demo");

            print("create", repository.create(item));
            OperationResult<TodoItem> read = repository.read(item.getId(), item.getCategory());
            print("read", read);

            TodoItem current = read.value();
            current.setCompleted(true);
            current.setTitle("Updated synchronous ToDo");
            print("update", repository.update(current));

            System.out.println("query by category:");
            repository.queryByCategory(item.getCategory(), PAGE_SIZE, page ->
                    System.out.printf("  page: %d item(s), %.2f RU, continuation=%s%n    %s%n",
                            page.items().size(),
                            page.requestCharge(),
                            page.continuationToken(),
                            page.items()));

            print("delete", repository.delete(item.getId(), item.getCategory()));
        }
    }

    private static void runAsyncDemo() {
        System.out.println("=== Asynchronous repository ===");
        CosmosAsyncClient client = CosmosConfiguration.createAsyncClient();
        CosmosConfiguration.initializeAsync(
                        client,
                        CosmosConfiguration.DEFAULT_DATABASE,
                        CosmosConfiguration.DEFAULT_CONTAINER)
                .flatMap(container -> runAsyncCrud(new AsyncTodoRepository(container)))
                .doFinally(signal -> client.close())
                .block();
    }

    private static Mono<Void> runAsyncCrud(AsyncTodoRepository repository) {
        TodoItem item = newItem("async-demo");
        return repository.create(item)
                .doOnNext(result -> print("create", result))
                .then(repository.read(item.getId(), item.getCategory()))
                .doOnNext(result -> print("read", result))
                .map(OperationResult::value)
                .flatMap(current -> {
                    current.setCompleted(true);
                    current.setTitle("Updated asynchronous ToDo");
                    return repository.update(current);
                })
                .doOnNext(result -> print("update", result))
                .thenMany(repository.queryByCategory(item.getCategory(), PAGE_SIZE))
                .doOnSubscribe(ignored -> System.out.println("query by category:"))
                .doOnNext(page -> System.out.printf(
                        "  page: %d item(s), %.2f RU, continuation=%s%n    %s%n",
                        page.items().size(),
                        page.requestCharge(),
                        page.continuationToken(),
                        page.items()))
                .then(repository.delete(item.getId(), item.getCategory()))
                .doOnNext(result -> print("delete", result))
                .then();
    }

    private static TodoItem newItem(String category) {
        return new TodoItem(
                UUID.randomUUID().toString(),
                "Demo ToDo",
                "Created by the Cosmos DB repository sample",
                false,
                Instant.now(),
                category);
    }

    private static void print(String operation, OperationResult<?> result) {
        System.out.printf("%s: %.2f RU, result=%s%n",
                operation, result.requestCharge(), result.value());
    }
}
