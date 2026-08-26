package com.example.blobevents;

import com.azure.core.models.CloudEvent;
import com.azure.messaging.eventgrid.EventGridEvent;

import java.util.List;
import java.util.regex.Pattern;

final class EventPayloadParser {
    private static final Pattern CLOUD_EVENT_MARKER =
        Pattern.compile("\"specversion\"\\s*:", Pattern.CASE_INSENSITIVE);

    private EventPayloadParser() {
    }

    static List<BlobLifecycleEvent> parse(String jsonPayload) {
        if (jsonPayload == null || jsonPayload.isBlank()) {
            throw new IllegalArgumentException("Event payload must not be blank");
        }

        if (CLOUD_EVENT_MARKER.matcher(jsonPayload).find()) {
            return CloudEvent.fromString(jsonPayload).stream()
                .map(event -> new BlobLifecycleEvent(
                    event.getId(),
                    event.getType(),
                    event.getSubject(),
                    event.getTime(),
                    event.getData()))
                .toList();
        }

        return EventGridEvent.fromString(jsonPayload).stream()
            .map(event -> new BlobLifecycleEvent(
                event.getId(),
                event.getEventType(),
                event.getSubject(),
                event.getEventTime(),
                event.getData()))
            .toList();
    }
}
