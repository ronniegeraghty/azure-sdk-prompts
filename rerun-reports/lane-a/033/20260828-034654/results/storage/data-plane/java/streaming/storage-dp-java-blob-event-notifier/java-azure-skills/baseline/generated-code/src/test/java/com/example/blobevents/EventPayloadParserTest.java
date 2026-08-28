package com.example.blobevents;

import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EventPayloadParserTest {
    @Test
    void parsesBothSupportedSchemasInOneBatch() {
        String payload = """
                [
                  {
                    "id": "eg-1",
                    "eventType": "Microsoft.Storage.BlobCreated",
                    "subject": "/blobServices/default/containers/docs/blobs/a.pdf",
                    "eventTime": "2026-08-28T00:00:00Z",
                    "data": {"api": "PutBlob"},
                    "dataVersion": "",
                    "metadataVersion": "1"
                  },
                  {
                    "specversion": "1.0",
                    "id": "ce-1",
                    "source": "/storage/demo",
                    "type": "Microsoft.Storage.BlobDeleted",
                    "subject": "/blobServices/default/containers/docs/blobs/b.pdf",
                    "time": "2026-08-28T00:01:00Z",
                    "datacontenttype": "application/json",
                    "data": {"api": "DeleteBlob"}
                  }
                ]
                """;

        List<EventEnvelope> events = EventPayloadParser.parse(payload);

        assertEquals(2, events.size());
        assertEquals(EventEnvelope.Schema.EVENT_GRID, events.get(0).schema());
        assertEquals(EventEnvelope.Schema.CLOUD_EVENTS, events.get(1).schema());
        assertEquals(EventReceiver.BLOB_CREATED, events.get(0).type());
        assertEquals(EventReceiver.BLOB_DELETED, events.get(1).type());
    }

    @Test
    void parsesEncodedBlobSubjectWithoutTreatingPlusAsSpace() {
        BlobEventHandler.BlobAddress address = BlobEventHandler.parseSubject(
                "/blobServices/default/containers/my%2Ddocs/blobs/2026/a+b%20c.pdf");

        assertEquals("my-docs", address.container());
        assertEquals("2026/a+b c.pdf", address.name());
    }
}
