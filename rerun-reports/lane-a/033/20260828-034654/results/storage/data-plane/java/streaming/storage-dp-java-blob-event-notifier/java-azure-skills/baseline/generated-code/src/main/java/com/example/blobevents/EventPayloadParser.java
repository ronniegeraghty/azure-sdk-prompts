package com.example.blobevents;

import com.azure.core.models.CloudEvent;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public final class EventPayloadParser {
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    private EventPayloadParser() {
    }

    public static List<EventEnvelope> parse(String payload) {
        try {
            JsonNode root = OBJECT_MAPPER.readTree(payload);
            List<JsonNode> nodes = new ArrayList<>();
            if (root.isArray()) {
                root.forEach(nodes::add);
            } else if (root.isObject()) {
                nodes.add(root);
            } else {
                throw new IllegalArgumentException("Event payload must be a JSON object or array");
            }

            List<EventEnvelope> events = new ArrayList<>(nodes.size());
            for (JsonNode node : nodes) {
                events.add(node.hasNonNull("specversion")
                        ? fromCloudEvent(node)
                        : fromEventGridEvent(node));
            }
            return List.copyOf(events);
        } catch (IOException | RuntimeException exception) {
            throw new IllegalArgumentException("Invalid Event Grid webhook payload", exception);
        }
    }

    private static EventEnvelope fromEventGridEvent(JsonNode node) {
        EventGridEvent event = EventGridEvent.fromString(node.toString()).get(0);
        return new EventEnvelope(
                event.getId(),
                event.getEventType(),
                event.getSubject(),
                event.getEventTime(),
                event.getData(),
                EventEnvelope.Schema.EVENT_GRID);
    }

    private static EventEnvelope fromCloudEvent(JsonNode node) {
        CloudEvent event = CloudEvent.fromString(node.toString()).get(0);
        return new EventEnvelope(
                event.getId(),
                event.getType(),
                event.getSubject(),
                event.getTime(),
                event.getData(),
                EventEnvelope.Schema.CLOUD_EVENTS);
    }
}
