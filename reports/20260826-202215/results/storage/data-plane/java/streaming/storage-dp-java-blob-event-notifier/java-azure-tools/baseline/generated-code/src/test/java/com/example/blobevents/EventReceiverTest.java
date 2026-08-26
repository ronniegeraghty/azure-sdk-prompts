package com.example.blobevents;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class EventReceiverTest {
    @Test
    void parsesAndRoutesBothSchemas() {
        CapturingProcessor processor = new CapturingProcessor();
        EventReceiver receiver = new EventReceiver(processor);

        List<BlobLifecycleEvent> gridEvents = receiver.receive(payload("eventType", "eventTime", null));
        List<BlobLifecycleEvent> cloudEvents = receiver.receive(payload("type", "time", "\"specversion\":\"1.0\","));

        assertEquals(EventSchema.EVENT_GRID, gridEvents.get(0).schema());
        assertEquals(EventSchema.CLOUD_EVENTS, cloudEvents.get(0).schema());
        assertEquals(List.of("created", "deleted", "created", "deleted"), processor.calls);
    }

    @Test
    void rejectsNonArrayPayload() {
        EventReceiver receiver = new EventReceiver(new CapturingProcessor());
        assertThrows(IllegalArgumentException.class, () -> receiver.receive("{}"));
    }

    private static String payload(String typeField, String timeField, String extra) {
        String prefix = extra == null ? "" : extra;
        return """
                [
                  {%s"id":"1","%s":"Microsoft.Storage.BlobCreated","subject":"/blobServices/default/containers/c/blobs/a.txt","%s":"2026-08-26T15:00:00Z","data":{}},
                  {%s"id":"2","%s":"Microsoft.Storage.BlobDeleted","subject":"/blobServices/default/containers/c/blobs/b.txt","%s":"2026-08-26T15:01:00Z","data":{}}
                ]
                """.formatted(prefix, typeField, timeField, prefix, typeField, timeField);
    }

    private static final class CapturingProcessor implements BlobEventProcessor {
        private final List<String> calls = new ArrayList<>();

        @Override
        public void onBlobCreated(BlobLifecycleEvent event) {
            calls.add("created");
        }

        @Override
        public void onBlobDeleted(BlobLifecycleEvent event) {
            calls.add("deleted");
        }
    }
}
