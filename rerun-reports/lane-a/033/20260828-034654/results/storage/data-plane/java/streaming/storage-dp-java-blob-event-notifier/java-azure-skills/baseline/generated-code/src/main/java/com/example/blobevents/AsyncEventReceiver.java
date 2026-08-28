package com.example.blobevents;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);
    private final BlobEventHandler handler;

    public AsyncEventReceiver(BlobEventHandler handler) {
        this.handler = handler;
    }

    public Flux<EventEnvelope> receive(String jsonPayload) {
        return Flux.fromIterable(EventPayloadParser.parse(jsonPayload))
                .concatMap(event -> route(event).thenReturn(event));
    }

    private Mono<Void> route(EventEnvelope event) {
        return switch (event.type()) {
            case EventReceiver.BLOB_CREATED -> handler.handleCreatedAsync(event);
            case EventReceiver.BLOB_DELETED -> handler.handleDeletedAsync(event);
            default -> {
                LOGGER.warn("Ignoring unrecognized event type {} for event {}", event.type(), event.id());
                yield Mono.empty();
            }
        };
    }
}
