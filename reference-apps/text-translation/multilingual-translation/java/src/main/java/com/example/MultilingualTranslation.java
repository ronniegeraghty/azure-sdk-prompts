package com.example;

import com.azure.ai.translation.text.TextTranslationClient;
import com.azure.ai.translation.text.TextTranslationClientBuilder;
import com.azure.ai.translation.text.models.TranslateInputItem;
import com.azure.ai.translation.text.models.TranslatedTextItem;
import com.azure.ai.translation.text.models.TranslationTarget;
import com.azure.ai.translation.text.models.TranslationText;
import com.azure.identity.DefaultAzureCredentialBuilder;
import java.util.List;

public final class MultilingualTranslation {
    private MultilingualTranslation() {
    }

    public static void main(String[] args) {
        if (args.length != 1) {
            throw new IllegalArgumentException(
                "Usage: multilingual-translation <text>");
        }

        String endpoint = requireEnvironmentVariable(
            "AZURE_TEXT_TRANSLATION_ENDPOINT");
        TextTranslationClient client = new TextTranslationClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
        TranslateInputItem input = new TranslateInputItem(
            args[0],
            List.of(
                new TranslationTarget("fr"),
                new TranslationTarget("ja")));

        TranslatedTextItem result = client.translate(List.of(input)).get(0);
        System.out.println(
            "Detected source language: "
                + result.getDetectedLanguage().getLanguage());
        for (TranslationText translation : result.getTranslations()) {
            System.out.printf(
                "Translation (%s): %s%n",
                translation.getLanguage(),
                translation.getText());
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
