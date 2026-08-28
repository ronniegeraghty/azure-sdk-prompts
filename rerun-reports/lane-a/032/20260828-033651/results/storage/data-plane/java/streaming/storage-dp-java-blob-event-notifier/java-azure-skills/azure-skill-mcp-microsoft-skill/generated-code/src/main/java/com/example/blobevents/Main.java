package com.example.blobevents;

import com.azure.core.util.BinaryData;
import com.example.blobevents.blob.AsyncBlobEventHandler;
import com.example.blobevents.blob.BlobEventHandler;
import com.example.blobevents.blob.BlobOperations;
import com.example.blobevents.blob.BlobSummary;
import com.example.blobevents.model.CustomEvent;
import com.example.blobevents.publisher.AsyncEventPublisher;
import com.example.blobevents.publisher.EventPublisher;
import com.example.blobevents.receiver.AsyncEventReceiver;
import com.example.blobevents.receiver.EventReceiver;
import reactor.core.publisher.Mono;

import java.nio.charset.StandardCharsets;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;
import java.util.logging.Logger;

public final class Main {
    private static final Logger LOGGER = Logger.getLogger(Main.class.getName());

    private static final String EVENT_GRID_PAYLOAD = """
        [
          {
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/invoices%2Finvoice-1001.pdf",
            "eventType": "Microsoft.Storage.BlobCreated",
            "id": "11111111-1111-1111-1111-111111111111",
            "data": {
              "api": "PutBlob",
              "clientRequestId": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
              "requestId": "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
              "eTag": "0x8DB123456789ABC",
              "contentType": "application/pdf",
              "contentLength": 2048,
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/invoices/invoice-1001.pdf",
              "sequencer": "000000000000000000000000000000010000000000000001"
            },
            "dataVersion": "",
            "metadataVersion": "1",
            "eventTime": "2026-08-28T01:30:00Z"
          },
          {
            "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/archive%2Fold-invoice.pdf",
            "eventType": "Microsoft.Storage.BlobDeleted",
            "id": "22222222-2222-2222-2222-222222222222",
            "data": {
              "api": "DeleteBlob",
              "url": "https://example.blob.core.windows.net/documents/archive/old-invoice.pdf",
              "sequencer": "000000000000000000000000000000020000000000000001"
            },
            "dataVersion": "",
            "metadataVersion": "1",
            "eventTime": "2026-08-28T01:31:00Z"
          }
        ]
        """;

    private static final String CLOUD_EVENTS_PAYLOAD = """
        [
          {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobCreated",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/reports%2Fquarterly.csv",
            "id": "33333333-3333-3333-3333-333333333333",
            "time": "2026-08-28T01:32:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "PutBlob",
              "contentType": "text/csv",
              "contentLength": 512,
              "blobType": "BlockBlob",
              "url": "https://example.blob.core.windows.net/documents/reports/quarterly.csv"
            }
          },
          {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobDeleted",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/example",
            "subject": "/blobServices/default/containers/documents/blobs/reports%2Fold.csv",
            "id": "44444444-4444-4444-4444-444444444444",
            "time": "2026-08-28T01:33:00Z",
            "datacontenttype": "application/json",
            "data": {
              "api": "DeleteBlob",
              "url": "https://example.blob.core.windows.net/documents/reports/old.csv"
            }
          }
        ]
        """;

    private Main() {
    }

    public static void main(String[] args) {
        BlobOperations demoBlobs = new DemoBlobOperations();
        CustomEvent notification = new CustomEvent(
            "Contoso.Documents.Processed",
            "/documents/invoices/processed",
            "1.0",
            Map.of("documentId", "invoice-1001", "processedAt", OffsetDateTime.now().toString()));

        LOGGER.info("=== Synchronous demo ===");
        EventReceiver receiver = new EventReceiver(new BlobEventHandler(demoBlobs));
        receiver.receive(EVENT_GRID_PAYLOAD);
        receiver.receive(CLOUD_EVENTS_PAYLOAD);
        new EventPublisher(events -> events.forEach(event ->
            LOGGER.info(() -> "Would publish sync event: type=" + event.getEventType()
                + ", subject=" + event.getSubject())))
            .publish(List.of(notification));

        LOGGER.info("=== Asynchronous demo ===");
        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(new AsyncBlobEventHandler(demoBlobs));
        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -> Mono.fromRunnable(() ->
            events.forEach(event -> LOGGER.info(() -> "Would publish async event: type="
                + event.getEventType() + ", subject=" + event.getSubject()))));

        asyncReceiver.receive(EVENT_GRID_PAYLOAD)
            .thenMany(asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD))
            .then(asyncPublisher.publish(List.of(notification)))
            .block();
    }

    private static final class DemoBlobOperations implements BlobOperations {
        @Override
        public DownloadedBlob download(String container, String name) {
            byte[] content = ("mock content for " + container + "/" + name)
                .getBytes(StandardCharsets.UTF_8);
            String contentType = name.endsWith(".pdf") ? "application/pdf" : "text/csv";
            return new DownloadedBlob(
                BinaryData.fromBytes(content),
                new BlobSummary(name, content.length, contentType, "HOT"));
        }

        @Override
        public Mono<DownloadedBlob> downloadAsync(String container, String name) {
            return Mono.fromSupplier(() -> download(container, name));
        }
    }
}
