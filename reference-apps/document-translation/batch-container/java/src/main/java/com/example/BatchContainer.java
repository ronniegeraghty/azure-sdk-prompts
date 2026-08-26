package com.example;

import com.azure.ai.translation.document.DocumentTranslationClient;
import com.azure.ai.translation.document.DocumentTranslationClientBuilder;
import com.azure.ai.translation.document.models.DocumentStatusResult;
import com.azure.ai.translation.document.models.DocumentTranslationInput;
import com.azure.ai.translation.document.models.TranslationSource;
import com.azure.ai.translation.document.models.TranslationStatusResult;
import com.azure.ai.translation.document.models.TranslationTarget;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.DefaultAzureCredentialBuilder;
import java.util.List;

public final class BatchContainer {
    private BatchContainer() {
    }

    public static void main(String[] args) {
        if (args.length != 3) {
            throw new IllegalArgumentException(
                "Usage: batch-container "
                    + "<source-container-uri> <target-container-uri> "
                    + "<target-language>");
        }

        String endpoint = requireEnvironmentVariable(
            "AZURE_DOCUMENT_TRANSLATION_ENDPOINT");
        DocumentTranslationClient client =
            new DocumentTranslationClientBuilder()
                .endpoint(endpoint)
                .credential(new DefaultAzureCredentialBuilder().build())
                .buildClient();
        DocumentTranslationInput input = new DocumentTranslationInput(
            new TranslationSource(args[0]),
            List.of(new TranslationTarget(args[1], args[2])));
        SyncPoller<TranslationStatusResult, TranslationStatusResult> poller =
            client.beginTranslation(List.of(input), false);
        TranslationStatusResult result =
            poller.waitForCompletion().getValue();

        System.out.println("Overall status: " + result.getStatus());
        for (
            DocumentStatusResult document
                : client.listDocumentStatuses(result.getId())
        ) {
            System.out.printf(
                "Document: id=%s status=%s%n",
                document.getId(),
                document.getStatus());
            if (document.getError() != null) {
                System.out.printf(
                    "Failure: code=%s message=%s%n",
                    document.getError().getCode(),
                    document.getError().getMessage());
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
