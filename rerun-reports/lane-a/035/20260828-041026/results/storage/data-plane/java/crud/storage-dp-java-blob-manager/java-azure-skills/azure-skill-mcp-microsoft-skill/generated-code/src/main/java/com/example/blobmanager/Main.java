package com.example.blobmanager;

import reactor.core.publisher.Mono;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.time.Duration;
import java.util.Map;

public final class Main {
    private static final Duration LEASE_DURATION = Duration.ofSeconds(30);

    private Main() {
    }

    public static void main(String[] args) throws IOException {
        String containerName = requireEnvironment("AZURE_STORAGE_CONTAINER");
        BlobStorageConfiguration configuration = BlobStorageConfiguration.fromEnvironment();
        Path workDirectory = Files.createTempDirectory("azure-blob-manager-");
        Path sampleFile = writeFile(workDirectory.resolve("sample.txt"), "Initial sample content.\n");
        Path updatedFile = writeFile(workDirectory.resolve("updated.txt"), "Updated while holding a lease.\n");
        Map<String, String> metadata = Map.of("source", "java-demo");
        Map<String, String> tags = Map.of("Project", "BlobManager", "Environment", "Demo");

        System.out.println("Using existing container: " + containerName);
        runSyncDemo(configuration, containerName, sampleFile, updatedFile, workDirectory, metadata, tags);
        runAsyncDemo(configuration, containerName, sampleFile, updatedFile, workDirectory, metadata, tags).block();
        System.out.println("Demo complete. Local files are in " + workDirectory);
    }

    private static void runSyncDemo(
            BlobStorageConfiguration configuration,
            String containerName,
            Path sampleFile,
            Path updatedFile,
            Path workDirectory,
            Map<String, String> metadata,
            Map<String, String> tags) {
        String blobName = "sync-sample.txt";
        BlobStorageService service = new BlobStorageService(configuration.createSyncClient(), containerName);

        System.out.println("\n--- Sync demo ---");
        BlobUploadResult uploaded = service.upload(sampleFile, blobName, metadata, tags);
        System.out.println("Uploaded " + blobName + " with ETag " + uploaded.eTag());

        service.listBlobs().forEach(item -> System.out.println("Listed blob: " + item.getName()));

        Path download = workDirectory.resolve("sync-download.txt");
        service.download(blobName, download, true);
        System.out.println("Downloaded " + blobName + " to " + download);

        String leaseId = service.acquireLease(blobName, LEASE_DURATION);
        System.out.println("Acquired lease " + leaseId);
        try {
            String expectedETag = service.getETag(blobName);
            BlobUploadResult updated = service.upload(updatedFile, blobName, metadata, tags, expectedETag, leaseId);
            System.out.println("Overwrote leased blob; new ETag " + updated.eTag());
        } finally {
            service.releaseLease(blobName, leaseId);
            System.out.println("Released lease");
        }

        System.out.println("Deleted " + blobName + ": " + service.delete(blobName));
    }

    private static Mono<Void> runAsyncDemo(
            BlobStorageConfiguration configuration,
            String containerName,
            Path sampleFile,
            Path updatedFile,
            Path workDirectory,
            Map<String, String> metadata,
            Map<String, String> tags) {
        String blobName = "async-sample.txt";
        Path download = workDirectory.resolve("async-download.txt");
        BlobStorageAsyncService service =
                new BlobStorageAsyncService(configuration.createAsyncClient(), containerName);

        System.out.println("\n--- Async demo ---");
        return service.upload(sampleFile, blobName, metadata, tags)
                .doOnNext(result -> System.out.println(
                        "Uploaded " + blobName + " with ETag " + result.eTag()))
                .thenMany(service.listBlobs()
                        .doOnNext(item -> System.out.println("Listed blob: " + item.getName())))
                .then(service.download(blobName, download, true))
                .doOnSuccess(ignored -> System.out.println("Downloaded " + blobName + " to " + download))
                .then(Mono.usingWhen(
                        service.acquireLease(blobName, LEASE_DURATION)
                                .doOnNext(leaseId -> System.out.println("Acquired lease " + leaseId)),
                        leaseId -> service.getETag(blobName)
                                .flatMap(eTag -> service.upload(
                                        updatedFile, blobName, metadata, tags, eTag, leaseId))
                                .doOnNext(result -> System.out.println(
                                        "Overwrote leased blob; new ETag " + result.eTag())),
                        leaseId -> service.releaseLease(blobName, leaseId)
                                .doOnSuccess(ignored -> System.out.println("Released lease"))))
                .then(service.delete(blobName))
                .doOnNext(deleted -> System.out.println("Deleted " + blobName + ": " + deleted))
                .then();
    }

    private static Path writeFile(Path path, String content) throws IOException {
        return Files.writeString(
                path,
                content,
                StandardOpenOption.CREATE,
                StandardOpenOption.TRUNCATE_EXISTING,
                StandardOpenOption.WRITE);
    }

    private static String requireEnvironment(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
