package com.example.blobmanager;

import com.azure.storage.blob.models.BlobItem;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        String containerName = System.getenv().getOrDefault("AZURE_STORAGE_CONTAINER", "blob-manager-demo");
        AzureBlobStorageConfiguration configuration =
                AzureBlobStorageConfiguration.fromEnvironment();
        AzureBlobStorageConfiguration.Clients clients = configuration.createClients();

        AzureBlobManager sync = new AzureBlobManager(clients.syncClient(), containerName);
        AzureBlobManagerAsync async = new AzureBlobManagerAsync(clients.asyncClient(), containerName);

        Path demoDirectory = Files.createTempDirectory("azure-blob-manager-");
        try {
            runSyncDemo(sync, demoDirectory);
            runAsyncDemo(async, demoDirectory);
        } finally {
            deleteDirectory(demoDirectory);
        }
    }

    private static void runSyncDemo(AzureBlobManager manager, Path directory) throws Exception {
        String blobName = "sync-demo-" + UUID.randomUUID() + ".txt";
        Path source = Files.writeString(directory.resolve("sync-source.txt"), "sync upload\n");
        Path download = directory.resolve("sync-download.txt");
        Map<String, String> metadata = Map.of("demo", "sync");
        Map<String, String> tags = Map.of("environment", "demo", "implementation", "sync");

        System.out.println("[sync] Uploading " + blobName);
        String etag = manager.upload(
                blobName, source, metadata, tags, BlobWriteCondition.createOnly());
        System.out.println("[sync] Uploaded with ETag " + etag);

        System.out.println("[sync] Listing blobs");
        manager.list().stream().map(BlobItem::getName).forEach(name -> System.out.println("  " + name));

        System.out.println("[sync] Downloading to " + download);
        manager.download(blobName, download, true);

        System.out.println("[sync] Acquiring lease and overwriting");
        String leaseId = manager.acquireLease(blobName, 60);
        try {
            Files.writeString(source, "sync lease-protected overwrite\n");
            String updatedEtag = manager.upload(
                    blobName, source, metadata, tags, BlobWriteCondition.withLease(leaseId));
            System.out.println("[sync] Overwritten with ETag " + updatedEtag);
        } finally {
            manager.releaseLease(blobName, leaseId);
        }

        System.out.println("[sync] Deleting " + blobName);
        manager.delete(blobName);
        System.out.println("[sync] Complete");
    }

    private static void runAsyncDemo(AzureBlobManagerAsync manager, Path directory) throws Exception {
        String blobName = "async-demo-" + UUID.randomUUID() + ".txt";
        Path source = Files.writeString(directory.resolve("async-source.txt"), "async upload\n");
        Path download = directory.resolve("async-download.txt");
        Map<String, String> metadata = Map.of("demo", "async");
        Map<String, String> tags = Map.of("environment", "demo", "implementation", "async");

        System.out.println("[async] Uploading " + blobName);
        String etag = manager.upload(
                blobName, source, metadata, tags, BlobWriteCondition.createOnly()).block();
        System.out.println("[async] Uploaded with ETag " + etag);

        System.out.println("[async] Listing blobs");
        List<BlobItem> blobs = manager.list().collectList().block();
        if (blobs != null) {
            blobs.stream().map(BlobItem::getName).forEach(name -> System.out.println("  " + name));
        }

        System.out.println("[async] Downloading to " + download);
        manager.download(blobName, download, true).block();

        System.out.println("[async] Acquiring lease and overwriting");
        String leaseId = manager.acquireLease(blobName, 60).block();
        if (leaseId == null) {
            throw new IllegalStateException("Azure returned no lease ID");
        }
        try {
            Files.writeString(source, "async lease-protected overwrite\n");
            String updatedEtag = manager.upload(
                    blobName, source, metadata, tags, BlobWriteCondition.withLease(leaseId)).block();
            System.out.println("[async] Overwritten with ETag " + updatedEtag);
        } finally {
            manager.releaseLease(blobName, leaseId).block();
        }

        System.out.println("[async] Deleting " + blobName);
        manager.delete(blobName).block();
        System.out.println("[async] Complete");
    }

    private static void deleteDirectory(Path directory) throws Exception {
        try (var paths = Files.walk(directory)) {
            for (Path path : paths.sorted((left, right) -> right.compareTo(left)).toList()) {
                Files.deleteIfExists(path);
            }
        }
    }
}
