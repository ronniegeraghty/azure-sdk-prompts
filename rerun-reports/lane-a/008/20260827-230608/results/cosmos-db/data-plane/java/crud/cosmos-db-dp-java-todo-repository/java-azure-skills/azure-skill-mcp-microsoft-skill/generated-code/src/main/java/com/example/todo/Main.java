package com.example.todo;

import com.azure.cosmos.models.CosmosItemResponse;
import com.azure.cosmos.models.FeedResponse;
import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final int QUERY_PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosToDoFactory factory = CosmosToDoFactory.fromEnvironment()) {
            runSyncDemo(factory.syncRepository());
            runAsyncDemo(factory.asyncRepository()).block();
        }
    }

    private static void runSyncDemo(SyncToDoRepository repository) {
        System.out.println("\n=== Synchronous CRUD ===");
        ToDoItem item = newItem("sync");

        printItemResponse("Created", repository.create(item));

        CosmosItemResponse<ToDoItem> readResponse =
            repository.read(item.getId(), item.getCategory());
        printItemResponse("Read", readResponse);

        ToDoItem current = readResponse.getItem();
        current.setTitle("Updated synchronous ToDo");
        current.setCompleted(true);
        printItemResponse("Updated", repository.update(current));

        System.out.println("Querying synchronous pages:");
        repository.queryByCategory(
            item.getCategory(),
            QUERY_PAGE_SIZE,
            Main::printPage
        );

        CosmosItemResponse<Object> deleteResponse =
            repository.delete(item.getId(), item.getCategory());
        System.out.printf(
            "Deleted id=%s, request charge=%.2f RU%n",
            item.getId(),
            deleteResponse.getRequestCharge()
        );
    }

    private static Mono<Void> runAsyncDemo(AsyncToDoRepository repository) {
        System.out.println("\n=== Asynchronous CRUD ===");
        ToDoItem item = newItem("async");

        return repository.create(item)
            .doOnNext(response -> printItemResponse("Created", response))
            .flatMap(ignored -> repository.read(item.getId(), item.getCategory()))
            .doOnNext(response -> printItemResponse("Read", response))
            .map(CosmosItemResponse::getItem)
            .flatMap(current -> {
                current.setTitle("Updated asynchronous ToDo");
                current.setCompleted(true);
                return repository.update(current);
            })
            .doOnNext(response -> printItemResponse("Updated", response))
            .thenMany(repository.queryByCategory(item.getCategory(), QUERY_PAGE_SIZE))
            .doOnSubscribe(ignored -> System.out.println("Querying asynchronous pages:"))
            .doOnNext(Main::printPage)
            .then(repository.delete(item.getId(), item.getCategory()))
            .doOnNext(response -> System.out.printf(
                "Deleted id=%s, request charge=%.2f RU%n",
                item.getId(),
                response.getRequestCharge()
            ))
            .then();
    }

    private static ToDoItem newItem(String prefix) {
        return new ToDoItem(
            prefix + "-" + UUID.randomUUID(),
            "Demo " + prefix + " ToDo",
            "This field is stored but excluded from indexing.",
            false,
            Instant.now(),
            "demo"
        );
    }

    private static void printItemResponse(
        String operation,
        CosmosItemResponse<ToDoItem> response
    ) {
        System.out.printf(
            "%s: %s, request charge=%.2f RU%n",
            operation,
            response.getItem(),
            response.getRequestCharge()
        );
    }

    private static void printPage(FeedResponse<ToDoItem> page) {
        System.out.printf(
            "Page: %d item(s), request charge=%.2f RU%n",
            page.getResults().size(),
            page.getRequestCharge()
        );
        page.getResults().forEach(result -> System.out.println("  " + result));
    }
}
