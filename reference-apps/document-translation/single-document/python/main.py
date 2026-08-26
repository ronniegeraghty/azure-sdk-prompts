import os
from pathlib import Path
import sys

from azure.ai.translation.document import SingleDocumentTranslationClient
from azure.ai.translation.document.models import DocumentTranslateContent
from azure.identity import DefaultAzureCredential


def main() -> None:
    if len(sys.argv) != 4:
        raise RuntimeError(
            "Usage: single-document "
            "<input-path> <target-language> <output-path>"
        )

    input_path = Path(sys.argv[1])
    target_language = sys.argv[2]
    output_path = Path(sys.argv[3])
    endpoint = require_environment_variable(
        "AZURE_DOCUMENT_TRANSLATION_ENDPOINT"
    )

    with (
        input_path.open("rb") as source,
        DefaultAzureCredential() as credential,
        SingleDocumentTranslationClient(
            endpoint,
            credential,
        ) as client,
    ):
        content = DocumentTranslateContent(
            document=(
                input_path.name,
                source,
                get_media_type(input_path),
            )
        )
        translated = client.translate(
            body=content,
            target_language=target_language,
        )

        with output_path.open("wb") as output:
            output.write(translated)


def get_media_type(path: Path) -> str:
    media_types = {
        ".txt": "text/plain",
        ".html": "text/html",
        ".htm": "text/html",
        ".docx": (
            "application/vnd.openxmlformats-officedocument"
            ".wordprocessingml.document"
        ),
        ".pdf": "application/pdf",
    }
    try:
        return media_types[path.suffix.lower()]
    except KeyError as error:
        raise RuntimeError(
            f"Unsupported document extension: {path.suffix}"
        ) from error


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
