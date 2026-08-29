package com.example.blobevents;

import static org.junit.jupiter.api.Assertions.assertEquals;

import com.example.blobevents.model.BlobLocation;
import com.example.blobevents.model.IncomingEvent.Schema;
import org.junit.jupiter.api.Test;

class EventReceiverTest {
    @Test
    void deserializesEventGridSchema() {
        String payload = """
            [{"id":"1","eventType":"Microsoft.Storage.BlobDeleted",
              "subject":"/blobServices/default/containers/docs/blobs/a.txt",
              "eventTime":"2026-08-29T03:00:00Z","data":{},"dataVersion":"","metadataVersion":"1"}]
            """;

        var events = EventReceiver.deserialize(payload);

        assertEquals(1, events.size());
        assertEquals(Schema.EVENT_GRID, events.get(0).schema());
        assertEquals(BlobEventHandler.BLOB_DELETED, events.get(0).type());
    }

    @Test
    void deserializesCloudEventsSchema() {
        String payload = """
            [{"specversion":"1.0","type":"Microsoft.Storage.BlobCreated","source":"/storage",
              "id":"2","time":"2026-08-29T03:00:00Z",
              "subject":"/blobServices/default/containers/docs/blobs/a.txt",
              "datacontenttype":"application/json","data":{}}]
            """;

        var events = EventReceiver.deserialize(payload);

        assertEquals(1, events.size());
        assertEquals(Schema.CLOUD_EVENT, events.get(0).schema());
        assertEquals(BlobEventHandler.BLOB_CREATED, events.get(0).type());
    }

    @Test
    void parsesEncodedBlobSubject() {
        BlobLocation location = BlobLocation.fromSubject(
            "/blobServices/default/containers/my-docs/blobs/2026/paid%20invoice+copy.pdf");

        assertEquals("my-docs", location.container());
        assertEquals("2026/paid invoice+copy.pdf", location.blobName());
    }
}
