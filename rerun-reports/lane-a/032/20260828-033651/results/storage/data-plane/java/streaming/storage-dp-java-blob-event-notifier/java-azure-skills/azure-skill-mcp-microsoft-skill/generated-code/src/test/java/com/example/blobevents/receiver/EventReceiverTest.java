package com.example.blobevents.receiver;

import com.azure.core.util.BinaryData;
import com.example.blobevents.blob.BlobEventHandler;
import com.example.blobevents.blob.BlobOperations;
import com.example.blobevents.blob.BlobSummary;
import org.junit.jupiter.api.Test;
import reactor.core.publisher.Mono;

import java.util.concurrent.atomic.AtomicReference;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EventReceiverTest {
    @Test
    void parsesCloudEventAndDecodesBlobPath() {
        AtomicReference<String> requestedBlob = new AtomicReference<>();
        BlobOperations operations = new BlobOperations() {
            @Override
            public DownloadedBlob download(String container, String name) {
                requestedBlob.set(container + "/" + name);
                return new DownloadedBlob(BinaryData.fromString("x"),
                    new BlobSummary(name, 1, "text/plain", "HOT"));
            }

            @Override
            public Mono<DownloadedBlob> downloadAsync(String container, String name) {
                return Mono.fromSupplier(() -> download(container, name));
            }
        };
        String payload = """
            [{
              "specversion":"1.0",
              "type":"Microsoft.Storage.BlobCreated",
              "source":"/storage",
              "subject":"/blobServices/default/containers/docs/blobs/folder%2Ffile+one.txt",
              "id":"1",
              "time":"2026-08-28T00:00:00Z",
              "datacontenttype":"application/json",
              "data":{"url":"https://example.blob.core.windows.net/docs/folder/file+one.txt"}
            }]
            """;

        new EventReceiver(new BlobEventHandler(operations)).receive(payload);

        assertEquals("docs/folder/file+one.txt", requestedBlob.get());
    }
}
