package com.example.blobevents.receiver;

import com.example.blobevents.blob.AsyncBlobEventHandler;
import com.example.blobevents.model.BlobLifecycleEvent;
import reactor.core.publisher.Flux;

import java.util.logging.Logger;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = Logger.getLogger(AsyncEventReceiver.class.getName());
    private static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    private static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private final AsyncBlobEventHandler handler;

    public AsyncEventReceiver(AsyncBlobEventHandler handler) {
        this.handler = handler;
    }

    public Flux<BlobLifecycleEvent> receive(String jsonPayload) {
        return Flux.fromIterable(EventPayloadParser.parse(jsonPayload))
            .concatMap(event -> route(event).thenReturn(event));
    }

    private reactor.core.publisher.Mono<Void> route(BlobLifecycleEvent event) {
        return switch (event.type()) {
            case BLOB_CREATED -> handler.handleCreated(event);
            case BLOB_DELETED -> handler.handleDeleted(event);
            default -> reactor.core.publisher.Mono.fromRunnable(
                () -> LOGGER.warning("Ignoring unrecognized event type: " + event.type()));
        };
    }
}
