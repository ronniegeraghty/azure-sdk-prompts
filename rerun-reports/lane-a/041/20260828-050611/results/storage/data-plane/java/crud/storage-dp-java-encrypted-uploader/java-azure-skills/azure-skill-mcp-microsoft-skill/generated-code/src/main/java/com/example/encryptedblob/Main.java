package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();
        Path workDirectory = Files.createTempDirectory("encrypted-blob-demo-");

        runSyncDemo(configuration, workDirectory);
        runAsyncDemo(configuration, workDirectory);
    }

    private static void runSyncDemo(AzureConfiguration configuration, Path workDirectory) throws Exception {
        Path source = workDirectory.resolve("sync-source.txt");
        Path destination = workDirectory.resolve("sync-downloaded.txt");
        Files.writeString(source, "Hello from the synchronous encrypted uploader.", StandardCharsets.UTF_8);

        KeyManagement keyManagement = new KeyManagement(
            configuration.keyClient(),
            configuration.credential(),
            configuration.keyName());
        EncryptedBlobClient client = new EncryptedBlobClient(
            configuration.blobContainerClient(),
            keyManagement);

        EncryptedBlobClient.UploadResult result = client.upload(source, "demo/sync-encrypted.bin");
        client.download("demo/sync-encrypted.bin", destination);

        System.out.println("Sync vault key ID: " + result.keyId());
        System.out.println("Sync wrapped DEK (base64): " + result.wrappedKeyBase64());
        System.out.println("Sync decrypted output: "
            + Files.readString(destination, StandardCharsets.UTF_8));
    }

    private static void runAsyncDemo(AzureConfiguration configuration, Path workDirectory) throws Exception {
        Path source = workDirectory.resolve("async-source.txt");
        Path destination = workDirectory.resolve("async-downloaded.txt");
        Files.writeString(source, "Hello from the asynchronous encrypted uploader.", StandardCharsets.UTF_8);

        AsyncKeyManagement keyManagement = new AsyncKeyManagement(
            configuration.keyAsyncClient(),
            configuration.credential(),
            configuration.keyName());
        AsyncEncryptedBlobClient client = new AsyncEncryptedBlobClient(
            configuration.blobContainerAsyncClient(),
            keyManagement);

        EncryptedBlobClient.UploadResult result = client
            .upload(source, "demo/async-encrypted.bin")
            .flatMap(upload -> client.download("demo/async-encrypted.bin", destination).thenReturn(upload))
            .block();

        if (result == null) {
            throw new IllegalStateException("The asynchronous round trip completed without an upload result.");
        }

        System.out.println("Async vault key ID: " + result.keyId());
        System.out.println("Async wrapped DEK (base64): " + result.wrappedKeyBase64());
        System.out.println("Async decrypted output: "
            + Files.readString(destination, StandardCharsets.UTF_8));
    }
}
