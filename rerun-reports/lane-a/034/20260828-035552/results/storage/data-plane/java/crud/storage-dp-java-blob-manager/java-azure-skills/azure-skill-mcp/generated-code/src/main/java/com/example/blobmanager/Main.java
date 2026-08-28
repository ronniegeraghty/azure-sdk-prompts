package com.example.blobmanager;

import com.azure.storage.blob.models.BlobProperties;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Map;

public final class Main {
    private static final String CONTAINER_ENV = "AZURE_STORAGE_CONTAINER";
    private static final int LEASE_SECONDS = 30;

    private Main() {
    }

    public static void main(String[] args) throws Exception {
        String containerName = System.getenv().getOrDefault(CONTAINER_ENV, "blob-manager-demo");
        BlobStorageConfig.Clients clients = BlobStorageConfig.fromEnvironment().createClients();

        Path workDirectory = Files.createTempDirectory("azure-blob-manager-");
        try {
            runSyncDemo(clients, containerName, workDirectory);
            runAsyncDemo(clients, containerName, workDirectory);
        } finally {
            deleteLocalFiles(workDirectory);
        }
    }

    private static void runSyncDemo(
            BlobStorageConfig.Clients clients,
            String containerName,
            Path workDirectory
    ) throws Exception {
        System.out.println("\n=== Synchronous demo ===");
        BlobStorageService service = new BlobStorageService(clients.syncClient(), containerName);
        service.ensureContainerExists();

        String blobName = "sync-sample.txt";
        Path source = workDirectory.resolve("sync-source.txt");
        Path download = workDirectory.resolve("sync-download.txt");
        Files.writeString(source, "Initial synchronous content\n");

        System.out.println("Uploading " + blobName);
        BlobProperties properties = service.upload(
                blobName,
                source,
                Map.of("demo", "sync"),
                Map.of("project", "blob-manager", "implementation", "sync"));

        System.out.println("Listing blobs:");
        service.listBlobs().forEach(item -> System.out.println(" - " + item.getName()));

        System.out.println("Downloading to " + download);
        service.download(blobName, download, true);

        System.out.println("Acquiring lease and conditionally overwriting " + blobName);
        Files.writeString(source, "Updated synchronous content under a lease\n");
        try (BlobStorageService.Lease lease = service.acquireLease(blobName, LEASE_SECONDS)) {
            properties = service.upload(
                    blobName,
                    source,
                    Map.of("demo", "sync", "version", "2"),
                    Map.of("project", "blob-manager", "implementation", "sync"),
                    properties.getETag(),
                    lease.leaseId());
            System.out.println("Overwrite complete; new ETag: " + properties.getETag());
        }

        System.out.println("Deleting " + blobName + ": " + service.delete(blobName));
    }

    private static void runAsyncDemo(
            BlobStorageConfig.Clients clients,
            String containerName,
            Path workDirectory
    ) throws Exception {
        System.out.println("\n=== Asynchronous demo ===");
        BlobStorageAsyncService service =
                new BlobStorageAsyncService(clients.asyncClient(), containerName);
        service.ensureContainerExists().block();

        String blobName = "async-sample.txt";
        Path source = workDirectory.resolve("async-source.txt");
        Path download = workDirectory.resolve("async-download.txt");
        Files.writeString(source, "Initial asynchronous content\n");

        System.out.println("Uploading " + blobName);
        BlobProperties properties = service.upload(
                blobName,
                source,
                Map.of("demo", "async"),
                Map.of("project", "blob-manager", "implementation", "async")).block();

        System.out.println("Listing blobs:");
        service.listBlobs()
                .doOnNext(item -> System.out.println(" - " + item.getName()))
                .then()
                .block();

        System.out.println("Downloading to " + download);
        service.download(blobName, download, true).block();

        System.out.println("Acquiring lease and conditionally overwriting " + blobName);
        Files.writeString(source, "Updated asynchronous content under a lease\n");
        String leaseId = service.acquireLease(blobName, LEASE_SECONDS).block();
        try {
            BlobProperties updated = service.upload(
                    blobName,
                    source,
                    Map.of("demo", "async", "version", "2"),
                    Map.of("project", "blob-manager", "implementation", "async"),
                    properties.getETag(),
                    leaseId).block();
            System.out.println("Overwrite complete; new ETag: " + updated.getETag());
        } finally {
            if (leaseId != null) {
                service.releaseLease(blobName, leaseId).block();
            }
        }

        System.out.println("Deleting " + blobName + ": " + service.delete(blobName).block());
    }

    private static void deleteLocalFiles(Path directory) throws Exception {
        try (var paths = Files.walk(directory)) {
            for (Path path : paths.sorted((left, right) -> right.compareTo(left)).toList()) {
                Files.deleteIfExists(path);
            }
        }
    }
}
