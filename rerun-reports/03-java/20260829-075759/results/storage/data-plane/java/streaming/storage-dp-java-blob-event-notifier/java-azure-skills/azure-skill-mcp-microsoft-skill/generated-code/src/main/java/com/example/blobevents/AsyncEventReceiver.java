package com.example.blobevents;

import com.example.blobevents.model.IncomingEvent;
import java.util.List;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);
    private final BlobEventHandler handler;

    public AsyncEventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public Flux<IncomingEvent> receive(String jsonPayload) {
        return Flux.defer(() -> {
            List<IncomingEvent> events = EventReceiver.deserialize(jsonPayload);
            return Flux.fromIterable(events);
        }).concatMap(event -> route(event).thenReturn(event));
    }

    private reactor.core.publisher.Mono<Void> route(IncomingEvent event) {
        return switch (event.type()) {
            case BlobEventHandler.BLOB_CREATED -> handler.handleCreatedAsync(event);
            case BlobEventHandler.BLOB_DELETED -> handler.handleDeletedAsync(event);
            default -> reactor.core.publisher.Mono.fromRunnable(
                () -> LOGGER.warn("Ignoring unrecognized Event Grid event type: {}", event.type()));
        };
    }
}
