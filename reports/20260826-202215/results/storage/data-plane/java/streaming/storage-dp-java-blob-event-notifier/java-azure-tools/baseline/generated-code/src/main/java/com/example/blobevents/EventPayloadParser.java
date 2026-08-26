package com.example.blobevents;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

final class EventPayloadParser {
    private final ObjectMapper objectMapper;

    EventPayloadParser(ObjectMapper objectMapper) {
        this.objectMapper = objectMapper;
    }

    List<BlobLifecycleEvent> parse(String payload) {
        try {
            JsonNode root = objectMapper.readTree(payload);
            if (!root.isArray()) {
                throw new IllegalArgumentException("Event Grid webhook payload must be a JSON array");
            }

            List<BlobLifecycleEvent> events = new ArrayList<>();
            for (JsonNode node : root) {
                boolean cloudEvent = node.hasNonNull("specversion");
                EventSchema schema = cloudEvent ? EventSchema.CLOUD_EVENTS : EventSchema.EVENT_GRID;
                String typeField = cloudEvent ? "type" : "eventType";
                String timeField = cloudEvent ? "time" : "eventTime";
                events.add(new BlobLifecycleEvent(
                        schema,
                        requiredText(node, "id"),
                        requiredText(node, typeField),
                        requiredText(node, "subject"),
                        OffsetDateTime.parse(requiredText(node, timeField)),
                        required(node, "data")));
            }
            return List.copyOf(events);
        } catch (JsonProcessingException e) {
            throw new IllegalArgumentException("Invalid event JSON payload", e);
        }
    }

    private static String requiredText(JsonNode node, String field) {
        JsonNode value = required(node, field);
        if (!value.isTextual() || value.textValue().isBlank()) {
            throw new IllegalArgumentException("Event field '" + field + "' must be a non-empty string");
        }
        return value.textValue();
    }

    private static JsonNode required(JsonNode node, String field) {
        JsonNode value = node.get(field);
        if (value == null || value.isNull()) {
            throw new IllegalArgumentException("Event is missing required field '" + field + "'");
        }
        return value;
    }
}
