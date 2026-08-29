package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();

        byte[] syncInput = "Hello from the synchronous encrypted uploader.".getBytes(StandardCharsets.UTF_8);
        EncryptedBlobStore syncStore = configuration.syncBlobStore("sync-encrypted-demo.bin");
        BlobEncryptionMetadata syncMetadata = syncStore.upload(syncInput, true);
        byte[] syncOutput = syncStore.download();
        printResult("Sync", syncMetadata, syncOutput);

        byte[] asyncInput = "Hello from the asynchronous encrypted uploader.".getBytes(StandardCharsets.UTF_8);
        configuration.asyncBlobStore("async-encrypted-demo.bin")
            .flatMap(store -> store.upload(asyncInput, true)
                .flatMap(metadata -> store.download()
                    .doOnNext(output -> printResult("Async", metadata, output))))
            .block();
    }

    private static void printResult(
        String implementation,
        BlobEncryptionMetadata metadata,
        byte[] plaintext
    ) {
        System.out.println(implementation + " vault key ID: " + metadata.keyId());
        System.out.println(implementation + " wrapped DEK (base64): " + metadata.wrappedDataKeyBase64());
        System.out.println(implementation + " decrypted output: "
            + new String(plaintext, StandardCharsets.UTF_8));
    }
}
