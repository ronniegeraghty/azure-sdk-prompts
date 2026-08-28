package com.example.azureencrypted;

import java.nio.charset.StandardCharsets;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();

        KeyManagement keyManagement =
                new KeyManagement(
                        configuration.cryptographyClient(),
                        configuration.credential(),
                        configuration.keyId());
        EncryptedBlobClient syncClient =
                new EncryptedBlobClient(configuration.blobContainerClient(), keyManagement);

        byte[] syncPlaintext = "Hello from synchronous envelope encryption."
                .getBytes(StandardCharsets.UTF_8);
        EncryptedBlobClient.UploadResult syncUpload =
                syncClient.upload("sync-encrypted-demo.bin", syncPlaintext);
        byte[] syncDownloaded = syncClient.download("sync-encrypted-demo.bin");
        printResult("Sync", syncUpload, syncDownloaded);

        AsyncKeyManagement asyncKeyManagement =
                new AsyncKeyManagement(
                        configuration.cryptographyAsyncClient(),
                        configuration.credential(),
                        configuration.keyId());
        AsyncEncryptedBlobClient asyncClient =
                new AsyncEncryptedBlobClient(
                        configuration.blobContainerAsyncClient(), asyncKeyManagement);

        byte[] asyncPlaintext = "Hello from asynchronous envelope encryption."
                .getBytes(StandardCharsets.UTF_8);
        EncryptedBlobClient.UploadResult asyncUpload =
                asyncClient.upload("async-encrypted-demo.bin", asyncPlaintext).block();
        byte[] asyncDownloaded = asyncClient.download("async-encrypted-demo.bin").block();

        if (asyncUpload == null || asyncDownloaded == null) {
            throw new IllegalStateException("The asynchronous round-trip completed without a result");
        }
        printResult("Async", asyncUpload, asyncDownloaded);
    }

    private static void printResult(
            String implementation,
            EncryptedBlobClient.UploadResult upload,
            byte[] decrypted) {
        System.out.println(implementation + " vault key ID: " + upload.keyId());
        System.out.println(implementation + " wrapped DEK (base64): "
                + upload.wrappedDataKeyBase64());
        System.out.println(implementation + " decrypted output: "
                + new String(decrypted, StandardCharsets.UTF_8));
    }
}
