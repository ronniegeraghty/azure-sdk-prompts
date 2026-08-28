package com.example.blobevents;

import java.util.List;
import java.util.Objects;
import reactor.core.publisher.Flux;

public final class EventReceiverAsync {
    private final EventPayloadParser parser;
    private final BlobEventHandler handler;

    public EventReceiverAsync(EventPayloadParser parser, BlobEventHandler handler) {
        this.parser = Objects.requireNonNull(parser, "parser");
        this.handler = Objects.requireNonNull(handler, "handler");
    }

    public Flux<BlobLifecycleEvent> receive(String jsonPayload) {
        List<BlobLifecycleEvent> events = parser.parse(jsonPayload);
        return Flux.fromIterable(events)
            .concatMap(event -> handler.handleAsync(event).thenReturn(event));
    }
}
