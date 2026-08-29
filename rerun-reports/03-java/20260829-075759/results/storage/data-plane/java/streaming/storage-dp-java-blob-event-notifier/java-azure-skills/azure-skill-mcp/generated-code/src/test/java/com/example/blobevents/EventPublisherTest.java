package com.example.blobevents;

import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicBoolean;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EventPublisherTest {
    @Test
    void preservesSubjectHierarchy() {
        CustomEvent customEvent = new CustomEvent(
            "Contoso.Documents.Processed",
            "/documents/invoices/processed",
            Map.of("id", "42"),
            "1.0");

        var event = EventPublisher.toEventGridEvents(List.of(customEvent)).get(0);

        assertEquals("/documents/invoices/processed", event.getSubject());
        assertEquals("Contoso.Documents.Processed", event.getEventType());
    }

    @Test
    void rejectsRelativeSubject() {
        assertThrows(IllegalArgumentException.class, () -> new CustomEvent(
            "Contoso.Documents.Processed",
            "documents/processed",
            Map.of(),
            "1.0"));
    }

    @Test
    void asyncPublishIsLazy() {
        AtomicBoolean called = new AtomicBoolean();
        AsyncEventPublisher publisher = new AsyncEventPublisher(events -> {
            called.set(true);
            return reactor.core.publisher.Mono.empty();
        });
        CustomEvent event = new CustomEvent(
            "Contoso.Documents.Processed",
            "/documents/processed",
            Map.of(),
            "1.0");

        var publication = publisher.publish(List.of(event));
        assertEquals(false, called.get());

        publication.block();
        assertEquals(true, called.get());
    }
}
