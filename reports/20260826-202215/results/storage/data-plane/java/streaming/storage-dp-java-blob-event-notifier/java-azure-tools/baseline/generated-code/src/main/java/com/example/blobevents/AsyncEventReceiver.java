package com.example.blobevents;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

public final class AsyncEventReceiver {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncEventReceiver.class);

    private final EventPayloadParser parser;
    private final AsyncBlobEventProcessor processor;

    public AsyncEventReceiver(AsyncBlobEventProcessor processor) {
        this(processor, new ObjectMapper());
    }

    AsyncEventReceiver(AsyncBlobEventProcessor processor, ObjectMapper objectMapper) {
        this.processor = processor;
        this.parser = new EventPayloadParser(objectMapper);
    }

    public Mono<Void> receive(String jsonPayload) {
        return Mono.fromCallable(() -> parser.parse(jsonPayload))
                .flatMapMany(Flux::fromIterable)
                .concatMap(this::route)
                .then();
    }

    private Mono<Void> route(BlobLifecycleEvent event) {
        return switch (event.type()) {
            case EventReceiver.BLOB_CREATED -> processor.onBlobCreated(event);
            case EventReceiver.BLOB_DELETED -> processor.onBlobDeleted(event);
            default -> Mono.fromRunnable(() ->
                    LOGGER.warn("Ignoring unsupported event type '{}' (event id {})", event.type(), event.id()));
        };
    }
}
