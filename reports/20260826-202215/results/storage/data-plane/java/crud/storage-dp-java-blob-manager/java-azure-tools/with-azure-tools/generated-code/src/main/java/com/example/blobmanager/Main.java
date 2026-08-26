package com.example.blobmanager;

import com.azure.storage.blob.models.BlobItem;
import reactor.core.publisher.Mono;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.List;
import java.util.Map;

public final class Main {
    private static final String DEFAULT_CONTAINER = "blob-manager-demo";
    private static final String SYNC_BLOB = "sync/sample.txt";
    private static final String ASYNC_BLOB = "async/sample.txt";
    private static final int LEASE_SECONDS = 60;

    private Main() {
    }

    public static void main(String[] args) throws IOException {
        BlobStorageConfiguration configuration = BlobStorageConfiguration.fromEnvironment();
        String containerName = environmentOrDefault("AZURE_STORAGE_CONTAINER", DEFAULT_CONTAINER);
        Path workDirectory = Files.createTempDirectory("azure-blob-manager-demo-");

        Path sample = workDirectory.resolve("sample.txt");
        Path replacement = workDirectory.resolve("replacement.txt");
        Files.writeString(sample, "Initial sample content\n", StandardOpenOption.CREATE_NEW);
        Files.writeString(replacement, "Content written while holding a lease\n", StandardOpenOption.CREATE_NEW);

        Map<String, String> metadata = Map.of("source", "blob-manager-demo");
        Map<String, String> tags = Map.of("project", "blob-manager", "stage", "demo");

        System.out.println("Demo files: " + workDirectory);
        runSyncDemo(configuration, containerName, sample, replacement, metadata, tags, workDirectory);
        runAsyncDemo(configuration, containerName, sample, replacement, metadata, tags, workDirectory);
        System.out.println("All operations completed.");
    }

    private static void runSyncDemo(
            BlobStorageConfiguration configuration,
            String containerName,
            Path sample,
            Path replacement,
            Map<String, String> metadata,
            Map<String, String> tags,
            Path workDirectory) {
        System.out.println("\n=== Synchronous demo ===");
        BlobStorageService service = new BlobStorageService(
                configuration.createSyncClient(),
                containerName,
                configuration.createTransferOptions());

        service.ensureContainerExists();
        System.out.println("Container ready: " + containerName);

        BlobUploadResult upload = service.upload(SYNC_BLOB, sample, metadata, tags);
        System.out.println("Uploaded " + upload.blobName() + " (ETag " + upload.eTag() + ")");

        printBlobs(service.listBlobs());

        Path download = workDirectory.resolve("sync-download.txt");
        service.download(SYNC_BLOB, download);
        System.out.println("Downloaded to " + download);

        String leaseId = service.acquireLease(SYNC_BLOB, LEASE_SECONDS);
        System.out.println("Lease acquired: " + leaseId);
        try {
            BlobUploadResult leasedUpload =
                    service.uploadWithLease(SYNC_BLOB, replacement, metadata, tags, leaseId);
            System.out.println("Overwrote under lease (ETag " + leasedUpload.eTag() + ")");
        } finally {
            service.releaseLease(SYNC_BLOB, leaseId);
            System.out.println("Lease released.");
        }

        System.out.println("Deleted: " + service.delete(SYNC_BLOB));
    }

    private static void runAsyncDemo(
            BlobStorageConfiguration configuration,
            String containerName,
            Path sample,
            Path replacement,
            Map<String, String> metadata,
            Map<String, String> tags,
            Path workDirectory) {
        System.out.println("\n=== Asynchronous demo ===");
        AsyncBlobStorageService service = new AsyncBlobStorageService(
                configuration.createAsyncClient(),
                containerName,
                configuration.createTransferOptions());

        service.ensureContainerExists()
                .doOnSuccess(ignored -> System.out.println("Container ready: " + containerName))
                .then(service.upload(ASYNC_BLOB, sample, metadata, tags))
                .doOnNext(result ->
                        System.out.println("Uploaded " + result.blobName() + " (ETag " + result.eTag() + ")"))
                .then(service.listBlobs())
                .doOnNext(Main::printBlobs)
                .then(service.download(ASYNC_BLOB, workDirectory.resolve("async-download.txt")))
                .doOnSuccess(ignored ->
                        System.out.println("Downloaded to " + workDirectory.resolve("async-download.txt")))
                .then(Mono.usingWhen(
                        service.acquireLease(ASYNC_BLOB, LEASE_SECONDS)
                                .doOnNext(leaseId -> System.out.println("Lease acquired: " + leaseId)),
                        leaseId -> service.uploadWithLease(
                                        ASYNC_BLOB, replacement, metadata, tags, leaseId)
                                .doOnNext(result -> System.out.println(
                                        "Overwrote under lease (ETag " + result.eTag() + ")"))
                                .then(),
                        leaseId -> service.releaseLease(ASYNC_BLOB, leaseId)
                                .doOnSuccess(ignored -> System.out.println("Lease released."))))
                .then(service.delete(ASYNC_BLOB))
                .doOnNext(deleted -> System.out.println("Deleted: " + deleted))
                .then()
                .block();
    }

    private static void printBlobs(List<BlobItem> blobs) {
        System.out.println("Blobs in container:");
        blobs.forEach(blob -> System.out.println("  - " + blob.getName()));
    }

    private static String environmentOrDefault(String name, String defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : value.trim();
    }
}
