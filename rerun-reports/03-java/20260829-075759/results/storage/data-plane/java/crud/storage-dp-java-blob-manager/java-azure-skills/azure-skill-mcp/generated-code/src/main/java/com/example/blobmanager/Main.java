package com.example.blobmanager;

import com.azure.storage.blob.models.BlobItem;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Map;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        if (args.length != 1) {
            System.err.println("Usage: mvn exec:java -Dexec.args=\"<container-name>\"");
            System.exit(2);
        }

        String containerName = args[0];
        AzureBlobConfiguration.StorageClients clients =
                AzureBlobConfiguration.fromEnvironment().createClients();
        BlobStorageService sync = new BlobStorageService(clients.syncClient());
        AsyncBlobStorageService async =
                new AsyncBlobStorageService(clients.asyncClient(), clients.requestTimeout());

        Path workDirectory = Files.createTempDirectory("azure-blob-manager-demo-");
        Path source = workDirectory.resolve("sample.txt");
        Files.writeString(source, "Initial content\n");
        Map<String, String> metadata = Map.of("demo", "azure-blob-manager");
        Map<String, String> tags = Map.of("project", "blob-manager", "environment", "demo");

        try {
            runSyncDemo(sync, containerName, source, workDirectory, metadata, tags);
            runAsyncDemo(async, containerName, source, workDirectory, metadata, tags);
        } finally {
            Files.deleteIfExists(workDirectory.resolve("sync-download.txt"));
            Files.deleteIfExists(workDirectory.resolve("async-download.txt"));
            Files.deleteIfExists(source);
            Files.deleteIfExists(workDirectory);
        }
    }

    private static void runSyncDemo(
            BlobStorageService service,
            String container,
            Path source,
            Path workDirectory,
            Map<String, String> metadata,
            Map<String, String> tags) throws Exception {
        String blobName = "sync-sample.txt";
        System.out.println("[sync] Uploading " + blobName);
        service.upload(container, blobName, source, metadata, tags);

        System.out.println("[sync] Listing blobs");
        service.list(container).stream().map(BlobItem::getName).forEach(name -> System.out.println("  " + name));

        Path download = workDirectory.resolve("sync-download.txt");
        System.out.println("[sync] Downloading to " + download);
        service.download(container, blobName, download);

        Files.writeString(source, "Content written while holding an Azure Blob lease\n");
        System.out.println("[sync] Acquiring lease and overwriting " + blobName);
        service.overwriteWithLease(container, blobName, source, metadata, tags);

        System.out.println("[sync] Deleting " + blobName);
        service.delete(container, blobName);
    }

    private static void runAsyncDemo(
            AsyncBlobStorageService service,
            String container,
            Path source,
            Path workDirectory,
            Map<String, String> metadata,
            Map<String, String> tags) throws Exception {
        String blobName = "async-sample.txt";
        Files.writeString(source, "Initial async content\n");
        System.out.println("[async] Uploading " + blobName);
        service.upload(container, blobName, source, metadata, tags).block();

        System.out.println("[async] Listing blobs");
        service.list(container).map(BlobItem::getName).doOnNext(name -> System.out.println("  " + name)).then().block();

        Path download = workDirectory.resolve("async-download.txt");
        System.out.println("[async] Downloading to " + download);
        service.download(container, blobName, download).block();

        Files.writeString(source, "Async content written while holding an Azure Blob lease\n");
        System.out.println("[async] Acquiring lease and overwriting " + blobName);
        service.overwriteWithLease(container, blobName, source, metadata, tags).block();

        System.out.println("[async] Deleting " + blobName);
        service.delete(container, blobName).block();
    }
}
