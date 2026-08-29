package com.example.encryptedblob;

import java.nio.charset.StandardCharsets;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) {
        AzureConfiguration configuration = AzureConfiguration.fromEnvironment();
        byte[] plaintext = "Client-side encryption with an Azure Key Vault KEK."
                .getBytes(StandardCharsets.UTF_8);

        EncryptedBlobClient syncClient = configuration.encryptedBlobClient();
        UploadReceipt syncReceipt = syncClient.upload("sync-encrypted-demo.bin", plaintext);
        byte[] syncDecrypted = syncClient.download("sync-encrypted-demo.bin");
        printResult("sync", syncReceipt, syncDecrypted);

        configuration.encryptedBlobAsyncClient()
                .flatMap(asyncClient ->
                        asyncClient.upload("async-encrypted-demo.bin", plaintext)
                                .flatMap(receipt -> asyncClient
                                        .download("async-encrypted-demo.bin")
                                        .doOnNext(decrypted ->
                                                printResult("async", receipt, decrypted))))
                .block();
    }

    private static void printResult(
            String implementation, UploadReceipt receipt, byte[] decrypted) {
        System.out.println("[" + implementation + "] Vault key ID: " + receipt.keyId());
        System.out.println("[" + implementation + "] Wrapped DEK (base64): "
                + receipt.wrappedDataKeyBase64());
        System.out.println("[" + implementation + "] Decrypted output: "
                + new String(decrypted, StandardCharsets.UTF_8));
    }
}
