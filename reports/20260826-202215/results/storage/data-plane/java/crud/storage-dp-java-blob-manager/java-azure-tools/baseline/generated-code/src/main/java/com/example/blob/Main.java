package com.example.blob;

import com.azure.storage.blob.BlobServiceAsyncClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobItem;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.Map;

public final class Main {
    private static final Map<String, String> METADATA = Map.of("demo", "azure-blob-manager");
    private static final Map<String, String> TAGS = Map.of("project", "blob-manager", "stage", "demo");

    private Main() {
    }

    public static void main(String[] args) throws IOException {
        AzureBlobConfiguration configuration = AzureBlobConfiguration.fromEnvironment();
        AzureBlobConfiguration.Settings settings = configuration.settings();
        Path workDirectory = Files.createDirectories(Path.of("target", "demo"));

        BlobServiceClient syncClient = configuration.createSyncClient();
        BlobStorageService sync = new BlobStorageService(
                syncClient.getBlobContainerClient(settings.containerName()));
        runSyncDemo(sync, workDirectory);

        BlobServiceAsyncClient asyncClient = configuration.createAsyncClient();
        BlobStorageAsyncService async = new BlobStorageAsyncService(
                asyncClient.getBlobContainerAsyncClient(settings.containerName()));
        runAsyncDemo(async, workDirectory).block();
    }

    private static void runSyncDemo(BlobStorageService service, Path directory) throws IOException {
        String blobName = "sync-demo.txt";
        Path source = Files.writeString(directory.resolve("sync-source.txt"), "Initial sync content\n");
        Path download = directory.resolve("sync-download.txt");

        System.out.println("[sync] Uploading " + blobName);
        service.upload(blobName, source, METADATA, TAGS);

        System.out.println("[sync] Listing blobs");
        service.list().forEach(item -> printBlob("[sync]", item));

        System.out.println("[sync] Downloading to " + download);
        service.download(blobName, download);

        System.out.println("[sync] Acquiring lease and overwriting " + blobName);
        String leaseId = service.acquireLease(blobName);
        try {
            Files.writeString(source, "Updated sync content under lease\n", StandardOpenOption.TRUNCATE_EXISTING);
            service.upload(blobName, source, METADATA, TAGS, leaseId);
        } finally {
            service.releaseLease(blobName, leaseId);
        }

        System.out.println("[sync] Deleting " + blobName);
        service.delete(blobName);
        System.out.println("[sync] Complete");
    }

    private static Mono<Void> runAsyncDemo(BlobStorageAsyncService service, Path directory) {
        String blobName = "async-demo.txt";
        Path source = directory.resolve("async-source.txt");
        Path download = directory.resolve("async-download.txt");

        return writeFile(source, "Initial async content\n")
                .then(Mono.defer(() -> {
                    System.out.println("[async] Uploading " + blobName);
                    return service.upload(blobName, source, METADATA, TAGS);
                }))
                .then(Mono.defer(() -> {
                    System.out.println("[async] Listing blobs");
                    return service.list();
                }))
                .doOnNext(items -> items.forEach(item -> printBlob("[async]", item)))
                .then(Mono.defer(() -> {
                    System.out.println("[async] Downloading to " + download);
                    return service.download(blobName, download);
                }))
                .then(Mono.defer(() -> {
                    System.out.println("[async] Acquiring lease and overwriting " + blobName);
                    return Mono.usingWhen(
                            service.acquireLease(blobName),
                            leaseId -> writeFile(source, "Updated async content under lease\n")
                                    .then(service.upload(blobName, source, METADATA, TAGS, leaseId)),
                            leaseId -> service.releaseLease(blobName, leaseId));
                }))
                .then(Mono.defer(() -> {
                    System.out.println("[async] Deleting " + blobName);
                    return service.delete(blobName);
                }))
                .doOnSuccess(deleted -> System.out.println("[async] Complete"))
                .then();
    }

    private static Mono<Path> writeFile(Path path, String content) {
        return Mono.fromCallable(() -> Files.writeString(
                        path, content, StandardOpenOption.CREATE, StandardOpenOption.TRUNCATE_EXISTING))
                .subscribeOn(Schedulers.boundedElastic());
    }

    private static void printBlob(String prefix, BlobItem item) {
        System.out.printf("%s - %s tags=%s%n", prefix, item.getName(), item.getTags());
    }
}
