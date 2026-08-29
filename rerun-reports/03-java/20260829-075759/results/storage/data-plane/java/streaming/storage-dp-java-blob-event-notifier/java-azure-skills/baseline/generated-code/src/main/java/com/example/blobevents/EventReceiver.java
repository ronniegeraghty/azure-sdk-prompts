package com.example.blobevents;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

public final class EventReceiver {
    public static final String BLOB_CREATED = "Microsoft.Storage.BlobCreated";
    public static final String BLOB_DELETED = "Microsoft.Storage.BlobDeleted";

    private static final Logger LOGGER = LoggerFactory.getLogger(EventReceiver.class);
    private final ObjectMapper mapper;

    public EventReceiver(ObjectMapper mapper) {
        this.mapper = mapper;
    }

    public List<LifecycleEvent> receive(String payload, BlobEventHandler handler) {
        List<LifecycleEvent> events = deserialize(payload);
        events.forEach(event -> route(event, handler));
        return events;
    }

    public Mono<List<LifecycleEvent>> receiveAsync(String payload, AsyncBlobEventHandler handler) {
        return Mono.fromCallable(() -> deserialize(payload))
                .flatMap(events -> Flux.fromIterable(events)
                        .concatMap(event -> routeAsync(event, handler))
                        .then(Mono.just(events)));
    }

    public List<LifecycleEvent> deserialize(String payload) {
        try {
            JsonNode root = mapper.readTree(payload);
            ArrayNode array = root.isArray()
                    ? (ArrayNode) root
                    : mapper.createArrayNode().add(root);
            List<LifecycleEvent> events = new ArrayList<>(array.size());
            array.forEach(node -> events.add(normalize(node)));
            return List.copyOf(events);
        } catch (JsonProcessingException | IllegalArgumentException exception) {
            throw new IllegalArgumentException("Invalid Event Grid webhook payload", exception);
        }
    }

    private LifecycleEvent normalize(JsonNode node) {
        boolean cloudEvent = node.hasNonNull("specversion");
        String type = requiredText(node, cloudEvent ? "type" : "eventType");
        String subject = requiredText(node, "subject");
        String timeField = cloudEvent ? "time" : "eventTime";
        String dataVersion = cloudEvent
                ? node.path("dataversion").asText("1.0")
                : node.path("dataVersion").asText("");

        return new LifecycleEvent(
                requiredText(node, "id"),
                type,
                subject,
                OffsetDateTime.parse(requiredText(node, timeField)),
                dataVersion,
                node.path("data"),
                cloudEvent ? LifecycleEvent.Schema.CLOUD_EVENTS : LifecycleEvent.Schema.EVENT_GRID);
    }

    private static String requiredText(JsonNode node, String field) {
        String value = node.path(field).asText(null);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Event is missing required field: " + field);
        }
        return value;
    }

    private static void route(LifecycleEvent event, BlobEventHandler handler) {
        switch (event.type()) {
            case BLOB_CREATED -> handler.onBlobCreated(event);
            case BLOB_DELETED -> handler.onBlobDeleted(event);
            default -> LOGGER.warn("Ignoring unrecognized event type '{}' (event id {})", event.type(), event.id());
        }
    }

    private static Mono<Void> routeAsync(LifecycleEvent event, AsyncBlobEventHandler handler) {
        return switch (event.type()) {
            case BLOB_CREATED -> handler.onBlobCreated(event);
            case BLOB_DELETED -> handler.onBlobDeleted(event);
            default -> {
                LOGGER.warn("Ignoring unrecognized event type '{}' (event id {})", event.type(), event.id());
                yield Mono.empty();
            }
        };
    }
}
