package com.example.blobevents;

import java.util.List;
import java.util.Objects;

public final class EventReceiver {
    private final EventPayloadParser parser;
    private final BlobEventHandler handler;

    public EventReceiver(EventPayloadParser parser, BlobEventHandler handler) {
        this.parser = Objects.requireNonNull(parser, "parser");
        this.handler = Objects.requireNonNull(handler, "handler");
    }

    public List<BlobLifecycleEvent> receive(String jsonPayload) {
        List<BlobLifecycleEvent> events = parser.parse(jsonPayload);
        events.forEach(handler::handle);
        return events;
    }
}
