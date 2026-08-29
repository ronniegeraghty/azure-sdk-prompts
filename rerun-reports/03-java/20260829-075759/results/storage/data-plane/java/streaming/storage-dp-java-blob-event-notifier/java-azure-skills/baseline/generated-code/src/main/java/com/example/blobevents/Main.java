package com.example.blobevents;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.json.JsonMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Map;

public final class Main {
    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);

    private Main() {
    }

    public static void main(String[] args) {
        ObjectMapper mapper = JsonMapper.builder().findAndAddModules().build();
        EventReceiver receiver = new EventReceiver(mapper);

        LOGGER.info("=== Synchronous demo ===");
        receiver.receive(EVENT_GRID_PAYLOAD, new DemoBlobEventHandler());
        publishSyncDemo();

        LOGGER.info("=== Asynchronous demo ===");
        receiver.receiveAsync(CLOUD_EVENTS_PAYLOAD, new DemoAsyncBlobEventHandler())
                .then(Mono.defer(Main::publishAsyncDemo))
                .block();
    }

    private static void publishSyncDemo() {
        EventPublisher.CustomEvent event = downstreamEvent();
        String endpoint = System.getenv("EVENT_GRID_TOPIC_ENDPOINT");
        if (endpoint == null || endpoint.isBlank()) {
            LOGGER.info("Mock publish: type='{}', subject='{}', data={}",
                    event.type(), event.subject(), event.data());
            return;
        }
        AzureConfiguration configuration = new AzureConfiguration();
        new EventPublisher(configuration.eventGridPublisherClient(endpoint)).publish(List.of(event));
    }

    private static Mono<Void> publishAsyncDemo() {
        EventPublisher.CustomEvent event = downstreamEvent();
        String endpoint = System.getenv("EVENT_GRID_TOPIC_ENDPOINT");
        if (endpoint == null || endpoint.isBlank()) {
            LOGGER.info("Mock async publish: type='{}', subject='{}', data={}",
                    event.type(), event.subject(), event.data());
            return Mono.empty();
        }
        AzureConfiguration configuration = new AzureConfiguration();
        return new EventPublisher(configuration.eventGridPublisherAsyncClient(endpoint))
                .publishAsync(List.of(event));
    }

    private static EventPublisher.CustomEvent downstreamEvent() {
        return new EventPublisher.CustomEvent(
                "/documents/invoices/processed",
                "Contoso.Documents.Processed",
                "1.0",
                Map.of("documentId", "invoice-2026-00842", "status", "processed"));
    }

    private static final class DemoBlobEventHandler extends BlobEventHandler {
        private DemoBlobEventHandler() {
            super(null);
        }

        @Override
        public void onBlobCreated(LifecycleEvent event) {
            BlobAddress address = BlobAddress.fromSubject(event.subject());
            LOGGER.info("Mock blob created: container='{}', name='{}', size={}, contentType='{}', accessTier='{}'",
                    address.container(), address.name(), event.data().path("contentLength").asLong(),
                    event.data().path("contentType").asText(), "Hot");
        }
    }

    private static final class DemoAsyncBlobEventHandler extends AsyncBlobEventHandler {
        private DemoAsyncBlobEventHandler() {
            super(null);
        }

        @Override
        public Mono<Void> onBlobCreated(LifecycleEvent event) {
            BlobEventHandler.BlobAddress address = BlobEventHandler.BlobAddress.fromSubject(event.subject());
            LOGGER.info("Mock async blob created: container='{}', name='{}', size={}, contentType='{}', accessTier='{}'",
                    address.container(), address.name(), event.data().path("contentLength").asLong(),
                    event.data().path("contentType").asText(), "Hot");
            return Mono.empty();
        }
    }

    private static final String EVENT_GRID_PAYLOAD = """
            [
              {
                "id": "7b1d78a2-a13b-4d67-8f74-2dba8495b22f",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": "/blobServices/default/containers/documents/blobs/invoices/invoice-2026-00842.pdf",
                "eventType": "Microsoft.Storage.BlobCreated",
                "eventTime": "2026-08-29T03:30:00Z",
                "data": {
                  "api": "PutBlob",
                  "clientRequestId": "c851a9a0-f385-4c48-bf74-aabfc8d06288",
                  "requestId": "37f64f8b-701e-0015-65c4-1e8bb7000000",
                  "eTag": "0x8DDAF04711A0B92",
                  "contentType": "application/pdf",
                  "contentLength": 184532,
                  "blobType": "BlockBlob",
                  "url": "https://demostorage.blob.core.windows.net/documents/invoices/invoice-2026-00842.pdf",
                  "sequencer": "00000000000000000000000000010A8A0000000000008c3d"
                },
                "dataVersion": "",
                "metadataVersion": "1"
              },
              {
                "id": "5a4177a2-13ea-497f-91c4-a4e2a6a22822",
                "topic": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": "/blobServices/default/containers/documents/blobs/archive/old-invoice.pdf",
                "eventType": "Microsoft.Storage.BlobDeleted",
                "eventTime": "2026-08-29T03:31:00Z",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://demostorage.blob.core.windows.net/documents/archive/old-invoice.pdf",
                  "sequencer": "00000000000000000000000000010A8B0000000000008c3e"
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
                "id": "a7c19351-5781-4eaa-bf56-12a417d4794d",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": "/blobServices/default/containers/documents/blobs/reports/quarterly-report.csv",
                "type": "Microsoft.Storage.BlobCreated",
                "time": "2026-08-29T03:35:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "PutBlockList",
                  "contentType": "text/csv",
                  "contentLength": 43218,
                  "blobType": "BlockBlob",
                  "url": "https://demostorage.blob.core.windows.net/documents/reports/quarterly-report.csv",
                  "sequencer": "00000000000000000000000000010A8C0000000000008c3f"
                }
              },
              {
                "specversion": "1.0",
                "id": "b690cf55-71b8-4dce-9655-f7ee93c26702",
                "source": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo/providers/Microsoft.Storage/storageAccounts/demostorage",
                "subject": "/blobServices/default/containers/documents/blobs/temp/upload.tmp",
                "type": "Microsoft.Storage.BlobDeleted",
                "time": "2026-08-29T03:36:00Z",
                "datacontenttype": "application/json",
                "data": {
                  "api": "DeleteBlob",
                  "url": "https://demostorage.blob.core.windows.net/documents/temp/upload.tmp",
                  "sequencer": "00000000000000000000000000010A8D0000000000008c40"
                }
              }
            ]
            """;
}
