package com.example.blobevents;

import com.azure.core.models.CloudEvent;
import com.azure.json.JsonProviders;
import com.azure.json.JsonReader;
import com.azure.json.JsonToken;
import com.azure.messaging.eventgrid.EventGridEvent;

import java.io.IOException;
import java.util.List;

final class EventPayloadDeserializer {
    List<BlobLifecycleEvent> deserialize(String payload) {
        EventSchema schema = detectSchema(payload);
        if (schema == EventSchema.CLOUD_EVENTS) {
            return CloudEvent.fromString(payload).stream()
                .map(event -> new BlobLifecycleEvent(
                    event.getId(),
                    event.getType(),
                    requireSubject(event.getSubject(), event.getId()),
                    event.getTime(),
                    event.getData(),
                    schema))
                .toList();
        }

        return EventGridEvent.fromString(payload).stream()
            .map(event -> new BlobLifecycleEvent(
                event.getId(),
                event.getEventType(),
                event.getSubject(),
                event.getEventTime(),
                event.getData(),
                schema))
            .toList();
    }

    private EventSchema detectSchema(String payload) {
        if (payload == null || payload.isBlank()) {
            throw new IllegalArgumentException("Event payload must not be blank");
        }

        try (JsonReader reader = JsonProviders.createReader(payload)) {
            JsonToken token = reader.nextToken();
            if (token == JsonToken.START_ARRAY) {
                token = reader.nextToken();
            }
            if (token != JsonToken.START_OBJECT) {
                throw new IllegalArgumentException("Event payload must contain a JSON object or array");
            }

            while (reader.nextToken() != JsonToken.END_OBJECT) {
                String fieldName = reader.getFieldName();
                reader.nextToken();
                if ("specversion".equals(fieldName)) {
                    return EventSchema.CLOUD_EVENTS;
                }
                if ("eventType".equals(fieldName)) {
                    return EventSchema.EVENT_GRID;
                }
                reader.skipChildren();
            }
        } catch (IOException exception) {
            throw new IllegalArgumentException("Event payload is not valid JSON", exception);
        }

        throw new IllegalArgumentException("Payload is neither Event Grid schema nor CloudEvents 1.0 schema");
    }

    private String requireSubject(String subject, String eventId) {
        if (subject == null || subject.isBlank()) {
            throw new IllegalArgumentException("CloudEvent " + eventId + " is missing a subject");
        }
        return subject;
    }
}
