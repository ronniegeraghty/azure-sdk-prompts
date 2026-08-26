import os
import sys

from azure.ai.translation.document import DocumentTranslationClient
from azure.identity import DefaultAzureCredential


def main() -> None:
    if len(sys.argv) != 4:
        raise RuntimeError(
            "Usage: batch-container "
            "<source-container-uri> <target-container-uri> "
            "<target-language>"
        )

    endpoint = require_environment_variable(
        "AZURE_DOCUMENT_TRANSLATION_ENDPOINT"
    )
    source_url, target_url, language = sys.argv[1:]
    with (
        DefaultAzureCredential() as credential,
        DocumentTranslationClient(endpoint, credential) as client,
    ):
        poller = client.begin_translation(
            source_url,
            target_url,
            language,
        )
        poller.result()
        print(f"Overall status: {poller.status()}")

        for document in client.list_document_statuses(poller.id):
            print(
                f"Document: id={document.id} "
                f"status={document.status}"
            )
            if document.error:
                print(
                    f"Failure: code={document.error.code} "
                    f"message={document.error.message}"
                )


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
