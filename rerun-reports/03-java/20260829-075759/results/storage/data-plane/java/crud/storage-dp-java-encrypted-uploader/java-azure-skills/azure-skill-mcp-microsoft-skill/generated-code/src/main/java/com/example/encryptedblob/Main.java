package com.example.encryptedblob;

import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;

import java.nio.charset.StandardCharsets;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureClientConfiguration configuration = AzureClientConfiguration.fromEnvironment();
        byte[] plaintext = (args.length == 0
                ? "Client-side envelope encryption with Azure Key Vault"
                : args[0]).getBytes(StandardCharsets.UTF_8);

        runSyncDemo(configuration, plaintext);
        runAsyncDemo(configuration, plaintext);
    }

    private static void runSyncDemo(AzureClientConfiguration configuration, byte[] plaintext) {
        BlobContainerClient container = configuration.blobServiceClient()
                .getBlobContainerClient(configuration.containerName());
        container.createIfNotExists();

        KeyManagementService keyManagement = new KeyManagementService(
                configuration.keyClient(),
                configuration.credential(),
                configuration.keyName());
        EncryptedBlobClient encryptedBlobs = new EncryptedBlobClient(container, keyManagement);

        EncryptedBlobInfo info = encryptedBlobs.upload("sync-demo.bin", plaintext);
        byte[] decrypted = encryptedBlobs.download("sync-demo.bin");

        printResult("Sync", info, decrypted);
    }

    private static void runAsyncDemo(AzureClientConfiguration configuration, byte[] plaintext) {
        BlobContainerAsyncClient container = configuration.blobServiceAsyncClient()
                .getBlobContainerAsyncClient(configuration.containerName());

        AsyncKeyManagementService keyManagement = new AsyncKeyManagementService(
                configuration.keyAsyncClient(),
                configuration.credential(),
                configuration.keyName());
        EncryptedBlobAsyncClient encryptedBlobs =
                new EncryptedBlobAsyncClient(container, keyManagement);

        container.createIfNotExists()
                .then(encryptedBlobs.upload("async-demo.bin", plaintext))
                .zipWhen(ignored -> encryptedBlobs.download("async-demo.bin"))
                .doOnNext(result -> printResult("Async", result.getT1(), result.getT2()))
                .block();
    }

    private static void printResult(String label, EncryptedBlobInfo info, byte[] decrypted) {
        System.out.println(label + " vault key ID: " + info.keyId());
        System.out.println(label + " wrapped DEK (base64): " + info.wrappedDataKeyBase64());
        System.out.println(label + " decrypted output: " + new String(decrypted, StandardCharsets.UTF_8));
    }
}
