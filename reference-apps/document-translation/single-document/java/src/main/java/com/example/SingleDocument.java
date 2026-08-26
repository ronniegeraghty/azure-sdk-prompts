package com.example;

import com.azure.ai.translation.document.SingleDocumentTranslationClient;
import com.azure.ai.translation.document.SingleDocumentTranslationClientBuilder;
import com.azure.ai.translation.document.models.DocumentFileDetails;
import com.azure.ai.translation.document.models.DocumentTranslateContent;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public final class SingleDocument {
    private SingleDocument() {
    }

    public static void main(String[] args) throws IOException {
        if (args.length != 3) {
            throw new IllegalArgumentException(
                "Usage: single-document "
                    + "<input-path> <target-language> <output-path>");
        }

        String endpoint = requireEnvironmentVariable(
            "AZURE_DOCUMENT_TRANSLATION_ENDPOINT");
        Path inputPath = Path.of(args[0]);
        SingleDocumentTranslationClient client =
            new SingleDocumentTranslationClientBuilder()
                .endpoint(endpoint)
                .credential(new DefaultAzureCredentialBuilder().build())
                .buildClient();
        DocumentFileDetails document = new DocumentFileDetails(
            BinaryData.fromBytes(Files.readAllBytes(inputPath)))
            .setFilename(inputPath.getFileName().toString())
            .setContentType(getMediaType(inputPath));
        BinaryData translated = client.translate(
            args[1],
            new DocumentTranslateContent(document));

        Files.write(Path.of(args[2]), translated.toBytes());
    }

    private static String getMediaType(Path path) {
        String filename = path.getFileName().toString().toLowerCase();
        if (filename.endsWith(".txt")) {
            return "text/plain";
        }
        if (filename.endsWith(".html") || filename.endsWith(".htm")) {
            return "text/html";
        }
        if (filename.endsWith(".docx")) {
            return "application/vnd.openxmlformats-officedocument"
                + ".wordprocessingml.document";
        }
        if (filename.endsWith(".pdf")) {
            return "application/pdf";
        }
        throw new IllegalArgumentException(
            "Unsupported document extension: " + filename);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
