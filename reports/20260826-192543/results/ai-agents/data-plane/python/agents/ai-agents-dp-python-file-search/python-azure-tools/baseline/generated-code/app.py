import os
import sys
from pathlib import Path
from typing import Optional

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import FilePurpose, FileSearchTool, ListSortOrder, RunStatus
from azure.core.exceptions import AzureError
from azure.identity import DefaultAzureCredential


GUIDE_FACT = (
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and "
    "hikers should bring a rain jacket."
)
QUESTION = (
    "According to the uploaded guide, how long is the Cascade Loop and what "
    "should hikers bring?"
)
GUIDE_PATH = Path(__file__).resolve().with_name("contoso_trail_guide.txt")


def delete_remote_resources(
    client: AgentsClient,
    thread_id: Optional[str],
    agent_id: Optional[str],
    vector_store_id: Optional[str],
    file_id: Optional[str],
) -> list[str]:
    errors: list[str] = []

    if thread_id is not None:
        try:
            client.threads.delete(thread_id)
        except AzureError as exc:
            errors.append(f"thread {thread_id}: {exc}")

    if agent_id is not None:
        try:
            client.delete_agent(agent_id)
        except AzureError as exc:
            errors.append(f"agent {agent_id}: {exc}")

    if vector_store_id is not None:
        try:
            client.vector_stores.delete(vector_store_id)
        except AzureError as exc:
            errors.append(f"vector store {vector_store_id}: {exc}")

    if file_id is not None:
        try:
            client.files.delete(file_id=file_id)
        except AzureError as exc:
            errors.append(f"file {file_id}: {exc}")

    return errors


def main() -> None:
    endpoint = os.environ["PROJECT_ENDPOINT"]
    model = os.environ["MODEL_DEPLOYMENT_NAME"]
    GUIDE_PATH.write_text(GUIDE_FACT, encoding="utf-8")

    file_id: Optional[str] = None
    vector_store_id: Optional[str] = None
    agent_id: Optional[str] = None
    thread_id: Optional[str] = None

    with AgentsClient(
        endpoint=endpoint,
        credential=DefaultAzureCredential(),
    ) as client:
        try:
            uploaded_file = client.files.upload_and_poll(
                file_path=str(GUIDE_PATH),
                purpose=FilePurpose.AGENTS,
            )
            file_id = uploaded_file.id

            vector_store = client.vector_stores.create_and_poll(
                file_ids=[file_id],
                name="hyoka-trail-guide-vector-store",
            )
            vector_store_id = vector_store.id

            if (
                vector_store.status != "completed"
                or vector_store.file_counts.completed != 1
                or vector_store.file_counts.failed != 0
            ):
                raise RuntimeError(
                    "Document indexing did not complete successfully "
                    f"(status={vector_store.status}, "
                    f"completed={vector_store.file_counts.completed}, "
                    f"failed={vector_store.file_counts.failed})."
                )

            file_search = FileSearchTool(vector_store_ids=[vector_store_id])
            agent = client.create_agent(
                model=model,
                name="hyoka-trail-guide-agent",
                instructions=(
                    "Answer questions using the uploaded trail guide. Use file "
                    "search and do not invent facts that are absent from the guide."
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
            if run.status != RunStatus.COMPLETED:
                raise RuntimeError(
                    f"Agent run ended with status {run.status}: {run.last_error}"
                )

            messages = client.messages.list(
                thread_id=thread_id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role == "assistant" and message.text_messages:
                    for text_message in message.text_messages:
                        print(text_message.text.value)
        finally:
            operation_failed = sys.exc_info()[0] is not None
            cleanup_errors = delete_remote_resources(
                client,
                thread_id,
                agent_id,
                vector_store_id,
                file_id,
            )
            if cleanup_errors:
                details = "; ".join(cleanup_errors)
                if operation_failed:
                    print(f"Cleanup failed: {details}", file=sys.stderr)
                else:
                    raise RuntimeError(f"Cleanup failed: {details}")


if __name__ == "__main__":
    main()
