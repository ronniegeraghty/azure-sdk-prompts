import os
import sys

from azure.ai.translation.text import TextTranslationClient
from azure.identity import DefaultAzureCredential


def main() -> None:
    if len(sys.argv) != 2:
        raise RuntimeError("Usage: multilingual-translation <text>")

    endpoint = require_environment_variable(
        "AZURE_TEXT_TRANSLATION_ENDPOINT"
    )
    with (
        DefaultAzureCredential() as credential,
        TextTranslationClient(
            endpoint=endpoint,
            credential=credential,
        ) as client,
    ):
        response = client.translate(
            body=[sys.argv[1]],
            to_language=["fr", "ja"],
        )

    if not response or response[0].detected_language is None:
        raise RuntimeError(
            "The service did not return a detected source language."
        )

    result = response[0]
    print(
        f"Detected source language: "
        f"{result.detected_language.language}"
    )
    for translation in result.translations:
        print(f"Translation ({translation.language}): {translation.text}")


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
