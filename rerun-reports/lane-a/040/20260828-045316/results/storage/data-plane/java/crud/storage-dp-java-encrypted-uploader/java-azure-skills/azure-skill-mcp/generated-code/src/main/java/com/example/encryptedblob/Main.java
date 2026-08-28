package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();
        Path source = args.length > 0
                ? Path.of(args[0])
                : Files.createTempFile("encrypted-blob-demo-", ".txt");
        if (args.length == 0) {
            Files.writeString(
                    source,
                    "Client-side envelope encryption with Azure Key Vault.",
                    StandardCharsets.UTF_8);
        }

        KeyManagementClient keyManager = new KeyManagementClient(
                configuration.keyClient(),
                configuration.credential(),
                configuration.keyName());
        EncryptedBlobClient syncClient = new EncryptedBlobClient(
                configuration.blobContainerClient(),
                keyManager);

        UploadResult syncResult = syncClient.upload(source, "sync-encrypted-demo.bin");
        byte[] syncPlaintext = syncClient.download("sync-encrypted-demo.bin");
        printResult("sync", syncResult, syncPlaintext);

        AsyncKeyManagementClient asyncKeyManager = new AsyncKeyManagementClient(
                configuration.keyAsyncClient(),
                configuration.credential(),
                configuration.keyName());
        EncryptedBlobAsyncClient asyncClient = new EncryptedBlobAsyncClient(
                configuration.blobContainerAsyncClient(),
                asyncKeyManager);

        asyncClient.upload(source, "async-encrypted-demo.bin")
                .flatMap(result -> asyncClient.download("async-encrypted-demo.bin")
                        .doOnNext(plaintext -> printResult("async", result, plaintext)))
                .block();
    }

    private static void printResult(String implementation, UploadResult result, byte[] plaintext) {
        System.out.println("[" + implementation + "] vault key ID: " + result.keyId());
        System.out.println("[" + implementation + "] wrapped DEK (base64): "
                + result.wrappedDataKeyBase64());
        System.out.println("[" + implementation + "] decrypted output: "
                + new String(plaintext, StandardCharsets.UTF_8));
    }
}
