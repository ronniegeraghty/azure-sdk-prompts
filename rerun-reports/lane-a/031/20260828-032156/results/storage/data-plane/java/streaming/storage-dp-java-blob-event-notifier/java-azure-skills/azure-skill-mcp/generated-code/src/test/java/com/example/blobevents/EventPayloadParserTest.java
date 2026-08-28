package com.example.blobevents;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import com.example.blobevents.BlobLifecycleEvent.EventSchema;
import java.util.List;
import org.junit.jupiter.api.Test;

class EventPayloadParserTest {
    private final EventPayloadParser parser = new EventPayloadParser();

    @Test
    void parsesEventGridAndCloudEventsSchemas() {
        List<BlobLifecycleEvent> eventGridEvents = parser.parse("""
            [{
              "id": "eg-1",
              "eventType": "Microsoft.Storage.BlobCreated",
              "subject": "/blobServices/default/containers/docs/blobs/a.txt",
              "eventTime": "2026-08-27T19:00:00Z",
              "data": {"api": "PutBlob"}
            }]
            """);
        List<BlobLifecycleEvent> cloudEvents = parser.parse("""
            {
              "specversion": "1.0",
              "id": "ce-1",
              "type": "Microsoft.Storage.BlobDeleted",
              "source": "/storageAccounts/demo",
              "subject": "/blobServices/default/containers/docs/blobs/a.txt",
              "time": "2026-08-27T19:01:00Z",
              "data": {"api": "DeleteBlob"}
            }
            """);

        assertEquals(EventSchema.EVENT_GRID, eventGridEvents.get(0).schema());
        assertEquals(EventSchema.CLOUD_EVENTS_1_0, cloudEvents.get(0).schema());
        assertEquals(BlobEventHandler.BLOB_CREATED, eventGridEvents.get(0).eventType());
        assertEquals(BlobEventHandler.BLOB_DELETED, cloudEvents.get(0).eventType());
    }

    @Test
    void rejectsUnsupportedCloudEventsVersion() {
        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () -> parser.parse("""
            {
              "specversion": "0.3",
              "id": "ce-1",
              "type": "Microsoft.Storage.BlobCreated",
              "subject": "/blobServices/default/containers/docs/blobs/a.txt"
            }
            """));

        assertEquals("Unsupported CloudEvents specversion: 0.3", exception.getMessage());
    }
}
