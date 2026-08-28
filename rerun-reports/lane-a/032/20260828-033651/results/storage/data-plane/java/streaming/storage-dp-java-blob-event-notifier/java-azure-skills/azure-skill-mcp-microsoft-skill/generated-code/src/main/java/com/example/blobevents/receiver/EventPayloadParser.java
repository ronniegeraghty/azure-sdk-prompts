package com.example.blobevents.receiver;

import com.azure.core.models.CloudEvent;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.example.blobevents.model.BlobLifecycleEvent;
import com.example.blobevents.model.BlobLifecycleEvent.EventSchema;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.io.UncheckedIOException;
import java.util.List;

final class EventPayloadParser {
    private static final ObjectMapper MAPPER = new ObjectMapper();

    private EventPayloadParser() {
    }

    static List<BlobLifecycleEvent> parse(String payload) {
        try {
            JsonNode root = MAPPER.readTree(payload);
            if (!root.isArray() || root.isEmpty()) {
                throw new IllegalArgumentException("Event Grid webhook payload must be a non-empty JSON array");
            }
            if (root.get(0).has("specversion")) {
                return CloudEvent.fromString(payload).stream()
                    .map(EventPayloadParser::fromCloudEvent)
                    .toList();
            }
            return EventGridEvent.fromString(payload).stream()
                .map(EventPayloadParser::fromEventGridEvent)
                .toList();
        } catch (IOException exception) {
            throw new UncheckedIOException("Invalid event JSON payload", exception);
        }
    }

    private static BlobLifecycleEvent fromCloudEvent(CloudEvent event) {
        return new BlobLifecycleEvent(
            event.getId(), event.getType(), event.getSubject(), event.getTime(),
            event.getData(), EventSchema.CLOUD_EVENTS);
    }

    private static BlobLifecycleEvent fromEventGridEvent(EventGridEvent event) {
        return new BlobLifecycleEvent(
            event.getId(), event.getEventType(), event.getSubject(), event.getEventTime(),
            event.getData(), EventSchema.EVENT_GRID);
    }
}
