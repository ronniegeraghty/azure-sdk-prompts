import os
import sys
from pathlib import Path
from typing import Optional

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import FilePurpose, FileSearchTool, ListSortOrder
from azure.core.exceptions import AzureError
from azure.identity import DefaultAzureCredential


DOCUMENT_TEXT = (
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and "
    "hikers should bring a rain jacket."
)
QUESTION = (
    "According to the uploaded guide, how long is the Cascade Loop and what "
    "should hikers bring?"
)
DOCUMENT_PATH = Path(__file__).with_name("contoso_trail_guide.txt")


def cleanup_resources(
    client: AgentsClient,
    thread_id: Optional[str],
    agent_id: Optional[str],
    vector_store_id: Optional[str],
    file_id: Optional[str],
) -> None:
    cleanup_errors: list[AzureError] = []

    operations = (
        ("thread", thread_id, lambda resource_id: client.threads.delete(resource_id)),
        ("agent", agent_id, lambda resource_id: client.delete_agent(resource_id)),
        (
            "vector store",
            vector_store_id,
            lambda resource_id: client.vector_stores.delete(resource_id),
        ),
        ("uploaded file", file_id, lambda resource_id: client.files.delete(file_id=resource_id)),
    )

    for resource_name, resource_id, delete_resource in operations:
        if resource_id is None:
            continue
        try:
            delete_resource(resource_id)
        except AzureError as exc:
            cleanup_errors.append(exc)
            print(f"Failed to delete {resource_name} {resource_id}: {exc}", file=sys.stderr)

    if cleanup_errors and sys.exc_info()[0] is None:
        raise RuntimeError("One or more Azure resources could not be deleted.") from cleanup_errors[0]


def main() -> None:
    project_endpoint = os.environ["PROJECT_ENDPOINT"]
    model_deployment_name = os.environ["MODEL_DEPLOYMENT_NAME"]

    DOCUMENT_PATH.write_text(DOCUMENT_TEXT, encoding="utf-8")

    file_id: Optional[str] = None
    vector_store_id: Optional[str] = None
    agent_id: Optional[str] = None
    thread_id: Optional[str] = None

    with DefaultAzureCredential() as credential:
        with AgentsClient(endpoint=project_endpoint, credential=credential) as client:
            try:
                uploaded_file = client.files.upload_and_poll(
                    file_path=str(DOCUMENT_PATH),
                    purpose=FilePurpose.AGENTS,
                )
                file_id = uploaded_file.id

                vector_store = client.vector_stores.create_and_poll(
                    file_ids=[file_id],
                    name="hyoka-trail-guide-vector-store",
                )
                vector_store_id = vector_store.id
                if vector_store.status != "completed":
                    raise RuntimeError(
                        "Vector store indexing did not complete successfully: "
                        f"{vector_store.status}"
                    )

                file_search = FileSearchTool(vector_store_ids=[vector_store_id])
                agent = client.create_agent(
                    model=model_deployment_name,
                    name="hyoka-trail-guide-agent",
                    instructions=(
                        "Answer questions using the uploaded trail guide. Use file search "
                        "and do not invent facts that are not in the guide."
                    ),
                    tools=file_search.definitions,
                    tool_resources=file_search.resources,
                )
                agent_id = agent.id

                thread = client.threads.create()
                thread_id = thread.id
                client.messages.create(
                    thread_id=thread_id,
                    role="user",
                    content=QUESTION,
                )

                run = client.runs.create_and_process(
                    thread_id=thread_id,
                    agent_id=agent_id,
                )
                if run.status != "completed":
                    raise RuntimeError(
                        f"Agent run ended with status {run.status}: {run.last_error}"
                    )

                messages = client.messages.list(
                    thread_id=thread_id,
                    order=ListSortOrder.ASCENDING,
                )
                assistant_text_found = False
                for message in messages:
                    if message.role == "assistant":
                        for text_message in message.text_messages:
                            assistant_text_found = True
                            print(text_message.text.value)
                if not assistant_text_found:
                    raise RuntimeError("The completed run produced no assistant text.")
            finally:
                cleanup_resources(
                    client=client,
                    thread_id=thread_id,
                    agent_id=agent_id,
                    vector_store_id=vector_store_id,
                    file_id=file_id,
                )


if __name__ == "__main__":
    main()
