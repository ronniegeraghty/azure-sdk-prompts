package com.example.blobevents;

import com.example.blobevents.BlobLifecycleEvent.EventSchema;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

public final class EventPayloadParser {
    private final ObjectMapper objectMapper;

    public EventPayloadParser() {
        this(new ObjectMapper());
    }

    EventPayloadParser(ObjectMapper objectMapper) {
        this.objectMapper = objectMapper;
    }

    public List<BlobLifecycleEvent> parse(String payload) {
        if (payload == null || payload.isBlank()) {
            throw new IllegalArgumentException("Event payload must not be blank");
        }

        try {
            JsonNode root = objectMapper.readTree(payload);
            List<BlobLifecycleEvent> events = new ArrayList<>();
            if (root.isArray()) {
                for (JsonNode node : (ArrayNode) root) {
                    events.add(parseEvent(node));
                }
            } else if (root.isObject()) {
                events.add(parseEvent(root));
            } else {
                throw new IllegalArgumentException("Event payload must be a JSON object or array");
            }
            return List.copyOf(events);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Event payload is not valid JSON", exception);
        }
    }

    private BlobLifecycleEvent parseEvent(JsonNode node) {
        if (node.hasNonNull("specversion")) {
            String specVersion = requiredText(node, "specversion");
            if (!"1.0".equals(specVersion)) {
                throw new IllegalArgumentException("Unsupported CloudEvents specversion: " + specVersion);
            }
            return new BlobLifecycleEvent(
                requiredText(node, "id"),
                requiredText(node, "type"),
                requiredText(node, "subject"),
                optionalTime(node, "time"),
                node.path("data"),
                EventSchema.CLOUD_EVENTS_1_0
            );
        }

        if (node.hasNonNull("eventType")) {
            return new BlobLifecycleEvent(
                requiredText(node, "id"),
                requiredText(node, "eventType"),
                requiredText(node, "subject"),
                optionalTime(node, "eventTime"),
                node.path("data"),
                EventSchema.EVENT_GRID
            );
        }

        throw new IllegalArgumentException("Event is neither Event Grid schema nor CloudEvents 1.0 schema");
    }

    private static String requiredText(JsonNode node, String field) {
        JsonNode value = node.get(field);
        if (value == null || !value.isTextual() || value.textValue().isBlank()) {
            throw new IllegalArgumentException("Event field '" + field + "' must be a non-empty string");
        }
        return value.textValue();
    }

    private static OffsetDateTime optionalTime(JsonNode node, String field) {
        JsonNode value = node.get(field);
        return value == null || value.isNull() ? null : OffsetDateTime.parse(requiredText(node, field));
    }
}
