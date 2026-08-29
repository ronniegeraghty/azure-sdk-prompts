package com.example.blobmanager;

import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import reactor.core.publisher.Mono;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.Map;

public final class Main {
    private static final String CONTAINER_ENV = "AZURE_STORAGE_CONTAINER";
    private static final String DEFAULT_CONTAINER = "blob-manager-demo";

    private Main() {
    }

    public static void main(String[] args) throws IOException {
        BlobStorageConfiguration.BlobStorageClients clients =
            BlobStorageConfiguration.fromEnvironment();
        String containerName = System.getenv().getOrDefault(CONTAINER_ENV, DEFAULT_CONTAINER);
        Path workDirectory = Files.createTempDirectory("azure-blob-manager-");

        try {
            Path sample = writeSample(workDirectory.resolve("sample.txt"), "sync sample");
            runSyncDemo(clients, containerName, sample, workDirectory);

            Path asyncSample = writeSample(workDirectory.resolve("async-sample.txt"), "async sample");
            runAsyncDemo(clients, containerName, asyncSample, workDirectory).block();
        } finally {
            deleteLocalFiles(workDirectory);
        }
    }

    private static void runSyncDemo(
        BlobStorageConfiguration.BlobStorageClients clients,
        String containerName,
        Path sample,
        Path workDirectory
    ) {
        String blobName = "sync/sample.txt";
        BlobContainerClient container = clients.syncClient().getBlobContainerClient(containerName);
        container.createIfNotExists();
        BlobStorageService service = new BlobStorageService(container, clients.requestTimeout());

        System.out.println("[sync] Uploading " + blobName);
        service.upload(sample, blobName, Map.of("source", "sync-demo"), Map.of("demo", "sync"));

        System.out.println("[sync] Listing blobs");
        service.listBlobs().forEach(item -> System.out.println("[sync] - " + item.getName()));

        Path download = workDirectory.resolve("sync-download.txt");
        System.out.println("[sync] Downloading to " + download);
        service.download(blobName, download);

        System.out.println("[sync] Acquiring lease and overwriting");
        String leaseId = service.acquireLease(blobName, 60);
        try {
            writeSample(sample, "sync lease-protected update");
            service.upload(
                sample, blobName, Map.of("source", "sync-demo"), Map.of("demo", "sync"), leaseId);
        } catch (IOException exception) {
            throw new IllegalStateException("Could not update the local sample", exception);
        } finally {
            service.releaseLease(blobName, leaseId);
        }

        System.out.println("[sync] Deleting " + blobName);
        service.delete(blobName);
        System.out.println("[sync] Complete");
    }

    private static Mono<Void> runAsyncDemo(
        BlobStorageConfiguration.BlobStorageClients clients,
        String containerName,
        Path sample,
        Path workDirectory
    ) {
        String blobName = "async/sample.txt";
        BlobContainerAsyncClient container =
            clients.asyncClient().getBlobContainerAsyncClient(containerName);
        AsyncBlobStorageService service =
            new AsyncBlobStorageService(container);
        Path download = workDirectory.resolve("async-download.txt");

        return container.createIfNotExists()
            .then(Mono.defer(() -> {
                System.out.println("[async] Uploading " + blobName);
                return service.upload(
                    sample, blobName, Map.of("source", "async-demo"), Map.of("demo", "async"));
            }))
            .then(Mono.defer(() -> {
                System.out.println("[async] Listing blobs");
                return service.listBlobs()
                    .doOnNext(item -> System.out.println("[async] - " + item.getName()))
                    .then();
            }))
            .then(Mono.defer(() -> {
                System.out.println("[async] Downloading to " + download);
                return service.download(blobName, download);
            }))
            .then(Mono.defer(() -> {
                System.out.println("[async] Acquiring lease and overwriting");
                return service.acquireLease(blobName, 60)
                    .flatMap(leaseId -> Mono.usingWhen(
                        Mono.just(leaseId),
                        id -> overwriteWithLease(service, sample, blobName, id),
                        id -> service.releaseLease(blobName, id),
                        (id, exception) -> service.releaseLease(blobName, id),
                        id -> service.releaseLease(blobName, id)));
            }))
            .then(Mono.defer(() -> {
                System.out.println("[async] Deleting " + blobName);
                return service.delete(blobName);
            }))
            .doOnSuccess(ignored -> System.out.println("[async] Complete"))
            .then();
    }

    private static Mono<Void> overwriteWithLease(
        AsyncBlobStorageService service,
        Path sample,
        String blobName,
        String leaseId
    ) {
        try {
            writeSample(sample, "async lease-protected update");
        } catch (IOException exception) {
            return Mono.error(new IllegalStateException(
                "Could not update the local sample", exception));
        }

        return service.upload(
                sample,
                blobName,
                Map.of("source", "async-demo"),
                Map.of("demo", "async"),
                leaseId)
            .then();
    }

    private static Path writeSample(Path path, String text) throws IOException {
        return Files.writeString(
            path,
            text + System.lineSeparator(),
            StandardCharsets.UTF_8,
            StandardOpenOption.CREATE,
            StandardOpenOption.TRUNCATE_EXISTING);
    }

    private static void deleteLocalFiles(Path directory) throws IOException {
        try (var paths = Files.walk(directory)) {
            for (Path path : paths.sorted((left, right) -> right.compareTo(left)).toList()) {
                Files.deleteIfExists(path);
            }
        }
    }
}
