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
                "eventTime": "2026-08-26T15:10:00Z",
                "dataVersion": "3",
                "metadataVersion": "1",
                "data": {
                  "api": "PutBlob",
                  "clientRequestId": "f4f7f56c-f5d6-4b92-86be-021789745aed",
                  "requestId": "3d2e6d94-901e-0023-37a6-a72c68000000",
                  "eTag": "0x8DC000000000001",
                  "contentType": "application/pdf",
                  "contentLength": 48291,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/invoices/2026/invoice-1042.pdf",
                  "sequencer": "000000000000000000000000000012340000000000000001"
                }
              },
              {
                "id": "eg-deleted-001",
                "eventType": "Microsoft.Storage.BlobDeleted",
                "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
                "eventTime": "2026-08-26T15:11:00Z",
                "dataVersion": "3",
                "metadataVersion": "1",
                "data": {
                  "api": "DeleteBlob",
                  "requestId": "8c4b7552-401e-0025-08a6-a77a70000000",
                  "url": "https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
                  "sequencer": "000000000000000000000000000012350000000000000001"
                }
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
                "time": "2026-08-26T15:12:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "PutBlockList",
                  "contentType": "text/csv",
                  "contentLength": 17320,
                  "blobType": "BlockBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/reports/quarterly-report.csv"
                }
              },
              {
                "specversion": "1.0",
                "id": "ce-deleted-001",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
                "type": "Microsoft.Storage.BlobDeleted",
                "subject": "/blobServices/default/containers/documents/blobs/reports/draft.csv",
                "time": "2026-08-26T15:13:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://examplestorage.blob.core.windows.net/documents/reports/draft.csv"
                }
              }
            ]
            """;

    private Main() {
    }

    public static void main(String[] args) {
        runSyncDemo();
        runAsyncDemo();
    }

    private static void runSyncDemo() {
        LOGGER.info("--- Synchronous demo ---");
        BlobEventHandler handler = new BlobEventHandler(Main::mockDownload);
        EventReceiver receiver = new EventReceiver(handler);
        receiver.receive(EVENT_GRID_PAYLOAD);
        receiver.receive(CLOUD_EVENTS_PAYLOAD);

        EventPublisher publisher = new EventPublisher(Main::logPublishedEvents);
        publisher.publish(List.of(processedEvent("invoice-1042.pdf")));
    }

    private static void runAsyncDemo() {
        LOGGER.info("--- Asynchronous demo ---");
        AsyncBlobEventHandler handler = new AsyncBlobEventHandler(
                (container, blobName) -> Mono.fromSupplier(() -> mockDownload(container, blobName)));
        AsyncEventReceiver receiver = new AsyncEventReceiver(handler);
        AsyncEventPublisher publisher = new AsyncEventPublisher(
                events -> Mono.fromRunnable(() -> logPublishedEvents(events)));

        receiver.receive(EVENT_GRID_PAYLOAD)
                .then(receiver.receive(CLOUD_EVENTS_PAYLOAD))
                .then(publisher.publish(List.of(processedEvent("quarterly-report.csv"))))
                .block();
    }

    private static BlobSummary mockDownload(String container, String blobName) {
        LOGGER.info("Offline demo download: container={}, name={}", container, blobName);
        String contentType = blobName.endsWith(".pdf") ? "application/pdf" : "text/csv";
        return new BlobSummary(blobName, blobName.endsWith(".pdf") ? 48_291 : 17_320, contentType, "Hot");
    }

    private static CustomEvent processedEvent(String name) {
        return new CustomEvent(
                "/documents/invoices/processed",
                "Contoso.Documents.Processed",
                "1.0",
                Map.of("blobName", name, "status", "processed"));
    }

    private static void logPublishedEvents(List<EventGridEvent> events) {
        events.forEach(event -> LOGGER.info(
                "Offline demo publish: type={}, subject={}, data={}",
                event.getEventType(), event.getSubject(), event.getData().toString()));
    }
}
