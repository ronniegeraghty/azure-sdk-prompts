package com.example.blobevents;

import com.example.blobevents.BlobEventHandler.BlobAddress;
import com.example.blobevents.BlobEventHandler.BlobSummary;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final String EVENT_GRID_PAYLOAD = """
        [
          {
            "id": "eg-created-001",
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "subject": "/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf",
            "eventType": "Microsoft.Storage.BlobCreated",
            "eventTime": "2026-08-27T19:00:00Z",
            "data": {
              "api": "PutBlob",
              "contentType": "application/pdf",
              "contentLength": 24576,
              "url": "https://demostore.blob.core.windows.net/documents/invoices/invoice-1001.pdf"
            },
            "dataVersion": "3",
            "metadataVersion": "1"
          },
          {
            "id": "eg-deleted-001",
            "subject": "/blobServices/default/containers/documents/blobs/archive%2Finvoice-0999.pdf",
            "eventType": "Microsoft.Storage.BlobDeleted",
            "eventTime": "2026-08-27T19:01:00Z",
            "data": {
              "api": "DeleteBlob",
              "url": "https://demostore.blob.core.windows.net/documents/archive/invoice-0999.pdf"
            },
            "dataVersion": "3",
            "metadataVersion": "1"
          },
          {
            "id": "eg-unknown-001",
            "subject": "/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf",
            "eventType": "Contoso.Storage.BlobReviewed",
            "eventTime": "2026-08-27T19:01:30Z",
            "data": {
              "review": "approved"
            },
            "dataVersion": "1",
            "metadataVersion": "1"
          }
        ]
        """;

    private static final String CLOUD_EVENTS_PAYLOAD = """
        [
          {
            "specversion": "1.0",
            "id": "ce-created-001",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "type": "Microsoft.Storage.BlobCreated",
            "subject": "/blobServices/default/containers/documents/blobs/reports%2Fquarterly.txt",
            "time": "2026-08-27T19:02:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "PutBlob",
              "contentType": "text/plain",
              "contentLength": 1024,
              "url": "https://demostore.blob.core.windows.net/documents/reports/quarterly.txt"
            }
          },
          {
            "specversion": "1.0",
            "id": "ce-deleted-001",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostore",
            "type": "Microsoft.Storage.BlobDeleted",
            "subject": "/blobServices/default/containers/documents/blobs/reports%2Fold-quarterly.txt",
            "time": "2026-08-27T19:03:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "DeleteBlob",
              "url": "https://demostore.blob.core.windows.net/documents/reports/old-quarterly.txt"
            }
          }
        ]
        """;

    private Main() {
    }

    public static void main(String[] args) {
        BlobEventHandler handler = BlobEventHandler.forDemo(Map.of(
            new BlobAddress("documents", "invoices/invoice-1001.pdf"),
            new BlobSummary(24_576, "application/pdf", "HOT"),
            new BlobAddress("documents", "reports/quarterly.txt"),
            new BlobSummary(1_024, "text/plain", "COOL")
        ));
        EventPayloadParser parser = new EventPayloadParser();
        List<CustomEvent> downstreamEvents = List.of(new CustomEvent(
            "/documents/invoices/processed",
            "Contoso.Documents.Processed",
            "1.0",
            Map.of("documentId", "invoice-1001", "status", "processed")
        ));

        System.out.println("=== Synchronous implementation ===");
        EventReceiver receiver = new EventReceiver(parser, handler);
        receiver.receive(EVENT_GRID_PAYLOAD);
        receiver.receive(CLOUD_EVENTS_PAYLOAD);
        EventPublisher.dryRun().publish(downstreamEvents);

        System.out.println("=== Asynchronous implementation ===");
        EventReceiverAsync asyncReceiver = new EventReceiverAsync(parser, handler);
        asyncReceiver.receive(EVENT_GRID_PAYLOAD).then().block();
        asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD).then().block();
        EventPublisherAsync.dryRun().publish(downstreamEvents).block();
    }
}
