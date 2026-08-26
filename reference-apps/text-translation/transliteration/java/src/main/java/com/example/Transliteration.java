package com.example;

import com.azure.ai.translation.text.TextTranslationClient;
import com.azure.ai.translation.text.TextTranslationClientBuilder;
import com.azure.ai.translation.text.models.TransliteratedText;
import com.azure.identity.DefaultAzureCredentialBuilder;

public final class Transliteration {
    private Transliteration() {
    }

    public static void main(String[] args) {
        if (args.length != 4) {
            throw new IllegalArgumentException(
                "Usage: transliteration "
                    + "<language> <from-script> <to-script> <text>");
        }

        String endpoint = requireEnvironmentVariable(
            "AZURE_TEXT_TRANSLATION_ENDPOINT");
        TextTranslationClient client = new TextTranslationClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
        TransliteratedText result = client.transliterate(
            args[0],
            args[1],
            args[2],
            args[3]);

        System.out.println("Transliterated text: " + result.getText());
        System.out.println("Returned script: " + result.getScript());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
