package com.example.blobevents;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EventReceiverTest {
    private final EventReceiver receiver = new EventReceiver();

    @Test
    void deserializesEventGridSchema() {
        String payload = """
            [{
              "id":"1",
              "eventType":"Microsoft.Storage.BlobCreated",
              "subject":"/blobServices/default/containers/docs/blobs/a.pdf",
              "eventTime":"2026-08-29T03:40:00Z",
              "data":{"url":"https://example.blob.core.windows.net/docs/a.pdf"},
              "dataVersion":"",
              "metadataVersion":"1"
            }]
            """;

        BlobLifecycleEvent event = receiver.deserialize(payload).get(0);

        assertEquals(EventSchema.EVENT_GRID, event.schema());
        assertEquals("Microsoft.Storage.BlobCreated", event.type());
        assertEquals("1", event.id());
    }

    @Test
    void deserializesCloudEventsSchema() {
        String payload = """
            [{
              "specversion":"1.0",
              "type":"Microsoft.Storage.BlobDeleted",
              "source":"/storageAccounts/example",
              "subject":"/blobServices/default/containers/docs/blobs/a.pdf",
              "id":"2",
              "time":"2026-08-29T03:40:00Z",
              "datacontenttype":"application/json",
              "data":{"url":"https://example.blob.core.windows.net/docs/a.pdf"}
            }]
            """;

        BlobLifecycleEvent event = receiver.deserialize(payload).get(0);

        assertEquals(EventSchema.CLOUD_EVENTS, event.schema());
        assertEquals("Microsoft.Storage.BlobDeleted", event.type());
        assertEquals("2", event.id());
    }
}
