package com.example.encryptedblob;

import com.azure.security.keyvault.keys.models.KeyVaultKey;

import java.nio.charset.StandardCharsets;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();
        byte[] plaintext = "Client-side encrypted with an Azure Key Vault protected DEK."
                .getBytes(StandardCharsets.UTF_8);

        KeyVaultKey syncVaultKey = configuration.keyClient().getKey(configuration.keyName());
        String syncKeyId = syncVaultKey.getId();
        KeyManagement syncKeyManagement = new KeyManagement(
                configuration.cryptographyClient(syncKeyId),
                syncKeyId,
                configuration::cryptographyClient);
        EncryptedBlobClient syncClient = new EncryptedBlobClient(
                configuration.blobContainerClient(), syncKeyManagement);

        EncryptedBlobClient.UploadResult syncUpload =
                syncClient.upload("sync-encrypted-demo.bin", plaintext);
        byte[] syncDownload = syncClient.download("sync-encrypted-demo.bin");
        printResult("sync", syncUpload, syncDownload);

        KeyVaultKey asyncVaultKey = configuration.keyAsyncClient()
                .getKey(configuration.keyName())
                .block();
        if (asyncVaultKey == null) {
            throw new IllegalStateException("Key Vault returned no key");
        }
        String asyncKeyId = asyncVaultKey.getId();
        AsyncKeyManagement asyncKeyManagement = new AsyncKeyManagement(
                configuration.cryptographyAsyncClient(asyncKeyId),
                asyncKeyId,
                configuration::cryptographyAsyncClient);
        AsyncEncryptedBlobClient asyncClient = new AsyncEncryptedBlobClient(
                configuration.blobContainerAsyncClient(), asyncKeyManagement);

        asyncClient.upload("async-encrypted-demo.bin", plaintext)
                .flatMap(upload -> asyncClient.download("async-encrypted-demo.bin")
                        .doOnNext(download -> printResult("async", upload, download)))
                .block();
    }

    private static void printResult(
            String implementation,
            EncryptedBlobClient.UploadResult upload,
            byte[] decrypted) {
        System.out.println("Implementation: " + implementation);
        System.out.println("Vault key ID: " + upload.keyId());
        System.out.println("Wrapped DEK (base64): " + upload.wrappedDataKeyBase64());
        System.out.println("Decrypted output: " + new String(decrypted, StandardCharsets.UTF_8));
        System.out.println();
    }
}
