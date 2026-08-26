package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;
import java.time.Duration;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();

        SyncKeyManagementClient syncKeys = new SyncKeyManagementClient(
            configuration.keyClient(),
            configuration.credential(),
            configuration.keyName()
        );
        SyncEncryptedBlobClient syncBlobs = new SyncEncryptedBlobClient(
            configuration.blobContainerClient(),
            syncKeys
        );

        byte[] syncPlaintext =
            "Hello from the synchronous encrypted uploader.".getBytes(StandardCharsets.UTF_8);
        UploadResult syncUpload = syncBlobs.upload(syncPlaintext, "sync-encrypted-demo.bin");
        byte[] syncDecrypted = syncBlobs.download("sync-encrypted-demo.bin");
        printResult("Sync", syncUpload, syncDecrypted);

        AsyncKeyManagementClient asyncKeys = new AsyncKeyManagementClient(
            configuration.keyAsyncClient(),
            configuration.credential(),
            configuration.keyName()
        );
        AsyncEncryptedBlobClient asyncBlobs = new AsyncEncryptedBlobClient(
            configuration.blobContainerAsyncClient(),
            asyncKeys
        );

        byte[] asyncPlaintext =
            "Hello from the asynchronous encrypted uploader.".getBytes(StandardCharsets.UTF_8);
        AsyncRoundTrip asyncRoundTrip = asyncBlobs
            .upload(asyncPlaintext, "async-encrypted-demo.bin")
            .flatMap(upload -> asyncBlobs
                .download("async-encrypted-demo.bin")
                .map(decrypted -> new AsyncRoundTrip(upload, decrypted)))
            .block(Duration.ofMinutes(2));

        if (asyncRoundTrip == null) {
            throw new IllegalStateException("The asynchronous round trip completed without a result");
        }
        printResult("Async", asyncRoundTrip.upload(), asyncRoundTrip.decrypted());
    }

    private static void printResult(String label, UploadResult upload, byte[] decrypted) {
        System.out.println(label + " Key Vault key ID: " + upload.keyId());
        System.out.println(label + " wrapped DEK (base64): " + upload.wrappedDataKeyBase64());
        System.out.println(
            label + " decrypted output: " + new String(decrypted, StandardCharsets.UTF_8)
        );
    }

    private record AsyncRoundTrip(UploadResult upload, byte[] decrypted) {
    }
}
