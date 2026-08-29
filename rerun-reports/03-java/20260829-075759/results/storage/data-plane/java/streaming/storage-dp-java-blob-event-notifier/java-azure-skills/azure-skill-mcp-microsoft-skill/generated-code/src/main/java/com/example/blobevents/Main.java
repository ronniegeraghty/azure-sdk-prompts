package com.example.blobevents;

import com.example.blobevents.model.BlobSummary;
import com.example.blobevents.model.CustomEvent;
import com.example.blobevents.storage.InMemoryBlobStore;
import java.util.List;
import java.util.Map;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

public final class Main {
    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);

    private Main() {
    }

    public static void main(String[] args) {
        InMemoryBlobStore demoStore = new InMemoryBlobStore(Map.of(
            "documents/invoices/august-2026.pdf",
            new BlobSummary("invoices/august-2026.pdf", 184_320, "application/pdf", "COOL")));
        BlobEventHandler handler = new BlobEventHandler(demoStore);
        List<CustomEvent> downstreamEvents = List.of(new CustomEvent(
            "/documents/invoices/processed",
            "Contoso.Documents.Processed",
            Map.of("documentId", "august-2026", "status", "processed"),
            "1.0"));

        LOGGER.info("----- synchronous demo -----");
        EventReceiver receiver = new EventReceiver(handler);
        receiver.receive(eventGridPayload());
        receiver.receive(cloudEventsPayload());
        new EventPublisher(events -> events.forEach(event ->
            LOGGER.info("Mock publish: type={}, subject={}", event.getEventType(), event.getSubject())))
            .publish(downstreamEvents);

        LOGGER.info("----- asynchronous demo -----");
        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(handler);
        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -> Mono.fromRunnable(() ->
            events.forEach(event ->
                LOGGER.info("Mock async publish: type={}, subject={}", event.getEventType(), event.getSubject()))));

        asyncReceiver.receive(eventGridPayload())
            .thenMany(asyncReceiver.receive(cloudEventsPayload()))
            .then(asyncPublisher.publish(downstreamEvents))
            .block();
    }

    private static String eventGridPayload() {
        return """
            [
              {
                "id": "2f01f1f8-4f52-4b6d-a9e1-0cf1a833f000",
                "eventType": "Microsoft.Storage.BlobCreated",
                "subject": "/blobServices/default/containers/documents/blobs/invoices/august-2026.pdf",
                "eventTime": "2026-08-29T03:50:00Z",
                "data": {
                  "api": "PutBlob",
                  "clientRequestId": "9f621d84-76c8-4a9d-81f3-fdf58d7b1077",
                  "requestId": "e4b171fe-501e-0013-137c-f0cacc000000",
                  "eTag": "0x8DEDEADBEEF0000",
                  "contentType": "application/pdf",
                  "contentLength": 184320,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/invoices/august-2026.pdf",
                  "sequencer": "0000000000000000000000000000001"
                },
                "dataVersion": "",
                "metadataVersion": "1"
              },
              {
                "id": "576499b2-8f4d-48ab-a98d-a6cd53eaf000",
                "eventType": "Microsoft.Storage.BlobDeleted",
                "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
                "eventTime": "2026-08-29T03:51:00Z",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
                  "sequencer": "0000000000000000000000000000002"
                },
                "dataVersion": "",
                "metadataVersion": "1"
              }
            ]
            """;
    }

    private static String cloudEventsPayload() {
        return """
            [
              {
                "specversion": "1.0",
                "type": "Microsoft.Storage.BlobCreated",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
                "id": "825af45f-1d7a-4f2f-a93b-13568c24f000",
                "time": "2026-08-29T03:52:00Z",
                "subject": "/blobServices/default/containers/documents/blobs/invoices/august-2026.pdf",
                "datacontenttype": "application/json",
                "data": {
                  "api": "PutBlob",
                  "contentType": "application/pdf",
                  "contentLength": 184320,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/invoices/august-2026.pdf",
                  "sequencer": "0000000000000000000000000000003"
                }
              },
              {
                "specversion": "1.0",
                "type": "Microsoft.Storage.BlobDeleted",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
                "id": "dfab8ab6-abdb-4ae5-acf5-bd605c56f000",
                "time": "2026-08-29T03:53:00Z",
                "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
                "datacontenttype": "application/json",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
                  "sequencer": "0000000000000000000000000000004"
                }
              }
            ]
            """;
    }
}
