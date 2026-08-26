package com.example.blobevents;

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
            "id": "7f4c87d8-904f-4a86-a8b7-94f8e342f256",
            "eventType": "Microsoft.Storage.BlobCreated",
            "subject": "/blobServices/default/containers/documents/blobs/invoices/2026/invoice-1042.pdf",
            "eventTime": "2026-08-26T15:58:12.123Z",
            "data": {
              "api": "PutBlob",
              "clientRequestId": "a8e90b14-2b28-4da7-8ab7-b0f14fb69521",
              "requestId": "f516a3cf-901e-0012-11f0-b69442000000",
              "eTag": "0x8DE44F3B6A2CC52",
              "contentType": "application/pdf",
              "contentLength": 184320,
              "blobType": "BlockBlob",
              "accessTier": "Hot",
              "url": "https://examplestorage.blob.core.windows.net/documents/invoices/2026/invoice-1042.pdf"
            },
            "dataVersion": "",
            "metadataVersion": "1"
          },
          {
            "id": "93a3ace3-6e73-47b9-aa02-0bce35757a41",
            "eventType": "Microsoft.Storage.BlobDeleted",
            "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
            "eventTime": "2026-08-26T15:59:03.407Z",
            "data": {
              "api": "DeleteBlob",
              "url": "https://examplestorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
              "blobType": "BlockBlob"
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
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
            "id": "f2d58ea0-c76b-45a5-a860-5c611afc50c8",
            "time": "2026-08-26T16:02:17.811Z",
            "subject": "/blobServices/default/containers/documents/blobs/contracts/contract-88.docx",
            "datacontenttype": "application/json",
            "data": {
              "api": "PutBlockList",
              "contentType": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
              "contentLength": 97280,
              "blobType": "BlockBlob",
              "accessTier": "Cool",
              "url": "https://examplestorage.blob.core.windows.net/documents/contracts/contract-88.docx"
            }
          },
          {
            "specversion": "1.0",
            "type": "Microsoft.Storage.BlobDeleted",
            "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/examplestorage",
            "id": "18e99ea5-6274-42eb-9920-9243565ad6ed",
            "time": "2026-08-26T16:03:45.113Z",
            "subject": "/blobServices/default/containers/documents/blobs/contracts/draft.docx",
            "datacontenttype": "application/json",
            "data": {
              "api": "DeleteBlob",
              "url": "https://examplestorage.blob.core.windows.net/documents/contracts/draft.docx",
              "blobType": "BlockBlob"
            }
          },
          {
            "specversion": "1.0",
            "type": "Contoso.Storage.BlobReviewed",
            "source": "/demo",
            "id": "7ee54594-bc84-4c03-ac8e-a2068204fcef",
            "time": "2026-08-26T16:04:00Z",
            "subject": "/blobServices/default/containers/documents/blobs/contracts/contract-88.docx",
            "datacontenttype": "application/json",
            "data": { "reviewed": true }
          }
        ]
        """;

    private Main() {
    }

    public static void main(String[] args) {
        BlobEventHandler.BlobDownloader syncDownloader = (container, blobName) ->
            new BlobSummary(blobName, 184_320, contentType(blobName), blobName.endsWith(".docx") ? "Cool" : "Hot");
        AsyncBlobEventHandler.AsyncBlobDownloader asyncDownloader = (container, blobName) ->
            Mono.just(new BlobSummary(
                blobName, 184_320, contentType(blobName), blobName.endsWith(".docx") ? "Cool" : "Hot"));

        EventPublisher syncPublisher = new EventPublisher(events ->
            events.forEach(event -> LOGGER.info(
                "Demo publish: type={}, subject={}", event.getEventType(), event.getSubject())));
        AsyncEventPublisher asyncPublisher = new AsyncEventPublisher(events -> Mono.fromRunnable(() ->
            events.forEach(event -> LOGGER.info(
                "Demo async publish: type={}, subject={}", event.getEventType(), event.getSubject()))));

        List<CustomEvent> downstreamEvents = List.of(new CustomEvent(
            "Contoso.Documents.Processed",
            "/documents/invoices/processed",
            Map.of("documentId", "invoice-1042", "status", "processed"),
            "1.0"));

        LOGGER.info("Starting synchronous demo");
        EventReceiver syncReceiver = new EventReceiver(new BlobEventHandler(syncDownloader));
        syncReceiver.receive(EVENT_GRID_PAYLOAD);
        syncReceiver.receive(CLOUD_EVENTS_PAYLOAD);
        syncPublisher.publish(downstreamEvents);

        LOGGER.info("Starting asynchronous demo");
        AsyncEventReceiver asyncReceiver = new AsyncEventReceiver(new AsyncBlobEventHandler(asyncDownloader));
        asyncReceiver.receive(EVENT_GRID_PAYLOAD)
            .then(asyncReceiver.receive(CLOUD_EVENTS_PAYLOAD))
            .then(asyncPublisher.publish(downstreamEvents))
            .block();
    }

    private static String contentType(String blobName) {
        return blobName.endsWith(".pdf")
            ? "application/pdf"
            : "application/vnd.openxmlformats-officedocument.wordprocessingml.document";
    }
}
