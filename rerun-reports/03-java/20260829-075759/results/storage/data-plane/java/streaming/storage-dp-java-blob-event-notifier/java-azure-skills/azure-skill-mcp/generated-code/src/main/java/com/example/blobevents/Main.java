package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Map;
import java.util.logging.Logger;
import java.util.stream.StreamSupport;

public final class Main {
    private static final Logger LOGGER = Logger.getLogger(Main.class.getName());

    private static final String EVENT_GRID_PAYLOAD = """
        [
          {
            "id": "f5f0a761-8a0f-4f9d-bf19-8014f5997a4f",
            "eventType": "Microsoft.Storage.BlobCreated",
            "subject": "/blobServices/default/containers/documents/blobs/invoices/2026-08/invoice-1001.pdf",
            "eventTime": "2026-08-29T03:40:00Z",
            "data": {
              "api": "PutBlob",
              "clientRequestId": "4d9329b5-efec-4f3a-b68d-9f3eb31f7857",
              "requestId": "8caa1a6e-501e-0030-1d80-9d3a23000000",
              "eTag": "0x8DE000000000001",
              "contentType": "application/pdf",
              "contentLength": 24576,
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/invoices/2026-08/invoice-1001.pdf",
              "sequencer": "000000000000000000000000000000A10000000000000123"
            },
            "dataVersion": "",
            "metadataVersion": "1"
          },
          {
            "id": "6f39f7e3-6368-45cd-8936-67bc4ea0376c",
            "eventType": "Microsoft.Storage.BlobDeleted",
            "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
            "eventTime": "2026-08-29T03:41:00Z",
            "data": {
              "api": "DeleteBlob",
              "requestId": "c48e8542-301e-0059-2f80-9dd848000000",
              "contentType": "application/pdf",
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/archive/old-invoice.pdf",
              "sequencer": "000000000000000000000000000000A20000000000000124"
            },
            "dataVersion": "",
            "metadataVersion": "1"
          }
        ]
        """;

    private static final String CLOUD_EVENTS_PAYLOAD = """
        [
          {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobCreated",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/reports/quarterly-summary.csv",
            "id": "cc32dad8-5ed4-4e3f-b829-bd55d72fdf0a",
            "time": "2026-08-29T03:42:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "PutBlockList",
              "contentType": "text/csv",
              "contentLength": 4096,
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/reports/quarterly-summary.csv",
              "sequencer": "000000000000000000000000000000A30000000000000125"
            }
          },
          {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobDeleted",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/temp/upload.tmp",
            "id": "c18e061b-b927-48cb-9bda-4836d3424e0c",
            "time": "2026-08-29T03:43:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "DeleteBlob",
              "contentType": "application/octet-stream",
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/temp/upload.tmp",
              "sequencer": "000000000000000000000000000000A40000000000000126"
            }
          }
        ]
        """;

    private Main() {
    }

    public static void main(String[] args) {
        BlobEventHandler handler = offlineBlobHandler();
        List<CustomEvent> downstreamEvents = List.of(new CustomEvent(
            "Contoso.Documents.DocumentProcessed",
            "/documents/invoices/processed",
            Map.of("documentId", "invoice-1001", "status", "processed"),
            "1.0"));

        LOGGER.info("Running synchronous demo");
        EventReceiver receiver = new EventReceiver();
        receiver.receive(EVENT_GRID_PAYLOAD, handler);
        receiver.receive(CLOUD_EVENTS_PAYLOAD, handler);
        new EventPublisher(events -> logPublished("sync", events)).publish(downstreamEvents);

        LOGGER.info("Running asynchronous demo");
        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver();
        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -> {
            logPublished("async", events);
            return Mono.empty();
        });
        asyncReceiver.receive(EVENT_GRID_PAYLOAD, handler)
            .then(asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD, handler))
            .then(asyncPublisher.publish(downstreamEvents))
            .block();
    }

    private static BlobEventHandler offlineBlobHandler() {
        BlobEventHandler.SyncBlobReader syncReader = location -> sampleSummary(location.name());
        BlobEventHandler.AsyncBlobReader asyncReader = location -> Mono.just(sampleSummary(location.name()));
        return new BlobEventHandler(syncReader, asyncReader);
    }

    private static BlobEventHandler.BlobSummary sampleSummary(String name) {
        String contentType = name.endsWith(".pdf") ? "application/pdf" : "text/csv";
        long size = name.endsWith(".pdf") ? 24_576 : 4_096;
        return new BlobEventHandler.BlobSummary(name, size, contentType, "HOT");
    }

    private static void logPublished(String mode, Iterable<EventGridEvent> events) {
        StreamSupport.stream(events.spliterator(), false)
            .forEach(event -> LOGGER.info(() -> mode + " published custom event: type="
                + event.getEventType() + ", subject=" + event.getSubject()));
    }
}
