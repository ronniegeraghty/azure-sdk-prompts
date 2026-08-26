import os
import sys

from azure.ai.translation.text import TextTranslationClient
from azure.identity import DefaultAzureCredential


def main() -> None:
    if len(sys.argv) != 5:
        raise RuntimeError(
            "Usage: transliteration "
            "<language> <from-script> <to-script> <text>"
        )

    endpoint = require_environment_variable(
        "AZURE_TEXT_TRANSLATION_ENDPOINT"
    )
    language, from_script, to_script, text = sys.argv[1:]
    with (
        DefaultAzureCredential() as credential,
        TextTranslationClient(
            endpoint=endpoint,
            credential=credential,
        ) as client,
    ):
        response = client.transliterate(
            body=[text],
            language=language,
            from_script=from_script,
            to_script=to_script,
        )

    if not response:
        raise RuntimeError("The service returned no transliteration.")

    print(f"Transliterated text: {response[0].text}")
    print(f"Returned script: {response[0].script}")


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
