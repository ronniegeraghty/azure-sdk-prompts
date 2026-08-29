package com.example.todo;

import com.example.todo.config.CosmosTodoFactory;
import com.example.todo.model.TodoItem;
import com.example.todo.repository.AsyncTodoRepository;
import com.example.todo.repository.OperationResult;
import com.example.todo.repository.QueryPage;
import com.example.todo.repository.SyncTodoRepository;
import reactor.core.publisher.Mono;

import java.time.Instant;
import java.util.UUID;

public final class Main {
    private static final int DEMO_PAGE_SIZE = 2;

    private Main() {
    }

    public static void main(String[] args) {
        try (CosmosTodoFactory factory = CosmosTodoFactory.create()) {
            runSyncDemo(factory.syncRepository());
            runAsyncDemo(factory.asyncRepository()).block();
        }
    }

    private static void runSyncDemo(SyncTodoRepository repository) {
        System.out.println("\n=== Synchronous repository ===");
        TodoItem item = new TodoItem(
                UUID.randomUUID().toString(),
                "Try the synchronous repository",
                "Create, read, update, query, and delete a ToDo item.",
                false,
                Instant.now(),
                "sync-demo");

        OperationResult<TodoItem> created = repository.create(item);
        printOperation("Created", created);

        OperationResult<TodoItem> read =
                repository.read(item.getId(), item.getCategory());
        printOperation("Read", read);

        read.item().setCompleted(true);
        read.item().setTitle("Synchronous repository complete");
        OperationResult<TodoItem> updated =
                repository.update(read.item(), read.etag());
        printOperation("Updated", updated);

        for (QueryPage<TodoItem> page
                : repository.queryByCategory(item.getCategory(), DEMO_PAGE_SIZE)) {
            printPage(page);
        }

        OperationResult<Void> deleted =
                repository.delete(item.getId(), item.getCategory());
        System.out.printf("Deleted id=%s, RU=%.2f%n",
                item.getId(), deleted.requestCharge());
    }

    private static Mono<Void> runAsyncDemo(AsyncTodoRepository repository) {
        System.out.println("\n=== Asynchronous repository ===");
        TodoItem item = new TodoItem(
                UUID.randomUUID().toString(),
                "Try the asynchronous repository",
                "Process query pages as each page arrives.",
                false,
                Instant.now(),
                "async-demo");

        return repository.create(item)
                .doOnNext(result -> printOperation("Created", result))
                .flatMap(ignored -> repository.read(item.getId(), item.getCategory()))
                .doOnNext(result -> printOperation("Read", result))
                .flatMap(read -> {
                    read.item().setCompleted(true);
                    read.item().setTitle("Asynchronous repository complete");
                    return repository.update(read.item(), read.etag());
                })
                .doOnNext(result -> printOperation("Updated", result))
                .thenMany(repository.queryByCategory(item.getCategory(), DEMO_PAGE_SIZE))
                .doOnNext(Main::printPage)
                .then(repository.delete(item.getId(), item.getCategory()))
                .doOnNext(result -> System.out.printf("Deleted id=%s, RU=%.2f%n",
                        item.getId(), result.requestCharge()))
                .then();
    }

    private static void printOperation(
            String operation,
            OperationResult<TodoItem> result) {
        System.out.printf("%s %s, ETag=%s, RU=%.2f%n",
                operation,
                result.item(),
                result.etag(),
                result.requestCharge());
    }

    private static void printPage(QueryPage<TodoItem> page) {
        System.out.printf("Page %d: %d item(s), RU=%.2f, hasMore=%s%n",
                page.pageNumber(),
                page.items().size(),
                page.requestCharge(),
                page.continuationToken() != null);
        page.items().forEach(result -> System.out.println("  " + result));
    }
}
