package com.example.blobevents;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EventPayloadParserTest {
    @Test
    void parsesEventGridSchema() {
        String payload = """
            [{
              "id":"1",
              "eventType":"Microsoft.Storage.BlobCreated",
              "subject":"/blobServices/default/containers/docs/blobs/a.pdf",
              "eventTime":"2026-08-26T15:58:12Z",
              "data":{"url":"https://example.blob.core.windows.net/docs/a.pdf"},
              "dataVersion":"",
              "metadataVersion":"1"
            }]
            """;

        BlobLifecycleEvent event = EventPayloadParser.parse(payload).get(0);

        assertEquals("Microsoft.Storage.BlobCreated", event.type());
        assertEquals("/blobServices/default/containers/docs/blobs/a.pdf", event.subject());
    }

    @Test
    void parsesCloudEventsSchema() {
        String payload = """
            [{
              "specversion":"1.0",
              "type":"Microsoft.Storage.BlobDeleted",
              "source":"/demo",
              "id":"2",
              "time":"2026-08-26T15:59:12Z",
              "subject":"/blobServices/default/containers/docs/blobs/a.pdf",
              "datacontenttype":"application/json",
              "data":{"url":"https://example.blob.core.windows.net/docs/a.pdf"}
            }]
            """;

        BlobLifecycleEvent event = EventPayloadParser.parse(payload).get(0);

        assertEquals("Microsoft.Storage.BlobDeleted", event.type());
        assertEquals("2", event.id());
    }

    @Test
    void rejectsBlankPayload() {
        assertThrows(IllegalArgumentException.class, () -> EventPayloadParser.parse(" "));
    }
}
