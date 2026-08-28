package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Map;

public final class Main {
    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);

    private static final String EVENT_GRID_PAYLOAD = """
            [
              {
                "id": "eg-created-001",
                "eventType": "Microsoft.Storage.BlobCreated",
                "subject": "/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1042.pdf",
                "eventTime": "2026-08-28T03:30:00Z",
                "data": {
                  "api": "PutBlob",
                  "clientRequestId": "2f46d2b0-21ad-4adb-874c-d21b8f9e2c0c",
                  "requestId": "95f80f35-901e-004f-3262-0ab54b000000",
                  "eTag": "0x8EE6D0A95A36C12",
                  "contentType": "application/pdf",
                  "contentLength": 48231,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/invoices/2026/invoice-1042.pdf",
                  "sequencer": "0000000000000000000000000002A7C0000000000012ab3c",
                  "storageDiagnostics": {"batchId": "8a4ab6c3-7084-4745-b196-81d4b4a9339e"}
                },
                "dataVersion": "",
                "metadataVersion": "1",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage"
              },
              {
                "id": "eg-deleted-001",
                "eventType": "Microsoft.Storage.BlobDeleted",
                "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
                "eventTime": "2026-08-28T03:31:00Z",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
                  "blobType": "BlockBlob"
                },
                "dataVersion": "",
                "metadataVersion": "1",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage"
              }
            ]
            """;

    private static final String CLOUD_EVENTS_PAYLOAD = """
            [
              {
                "specversion": "1.0",
                "id": "ce-created-001",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
                "type": "Microsoft.Storage.BlobCreated",
                "subject": "/blobServices/default/containers/documents/blobs/reports/quarterly-report.csv",
                "time": "2026-08-28T03:32:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "PutBlockList",
                  "contentType": "text/csv",
                  "contentLength": 16384,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/reports/quarterly-report.csv"
                }
              },
              {
                "specversion": "1.0",
                "id": "ce-deleted-001",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
                "type": "Microsoft.Storage.BlobDeleted",
                "subject": "/blobServices/default/containers/documents/blobs/tmp/import.csv",
                "time": "2026-08-28T03:33:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/tmp/import.csv",
                  "blobType": "BlockBlob"
                }
              }
            ]
            """;

    private Main() {
    }

    public static void main(String[] args) {
        BlobEventHandler handler = demoBlobHandler();
        List<CustomEvent> downstreamEvents = List.of(new CustomEvent(
                "Contoso.Documents.Processed",
                "/documents/invoices/processed",
                Map.of("documentId", "invoice-1042", "status", "processed")));

        LOGGER.info("Running synchronous Event Grid schema demo");
        new EventReceiver(handler).receive(EVENT_GRID_PAYLOAD);
        new EventPublisher(Main::logPublishedEvents).publish(downstreamEvents);

        LOGGER.info("Running asynchronous CloudEvents 1.0 demo");
        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -> {
            logPublishedEvents(events);
            return Mono.empty();
        });
        new AsyncEventReceiver(handler).receive(CLOUD_EVENTS_PAYLOAD)
                .then(Mono.defer(() -> asyncPublisher.publish(downstreamEvents)))
                .block();
    }

    private static BlobEventHandler demoBlobHandler() {
        BlobEventHandler.BlobOperations sync = address -> mockSummary(address.name());
        BlobEventHandler.AsyncBlobOperations async = address -> Mono.just(mockSummary(address.name()));
        return new BlobEventHandler(sync, async);
    }

    private static BlobEventHandler.BlobSummary mockSummary(String name) {
        String contentType = name.endsWith(".csv") ? "text/csv" : "application/pdf";
        return new BlobEventHandler.BlobSummary(name, 48_231, contentType, "Hot");
    }

    private static void logPublishedEvents(List<EventGridEvent> events) {
        events.forEach(event -> LOGGER.info(
                "Mock-published downstream event: type={}, subject={}",
                event.getEventType(),
                event.getSubject()));
    }
}
