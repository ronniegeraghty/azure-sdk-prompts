import os
import sys
from pathlib import Path

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import (
    FilePurpose,
    FileSearchTool,
    ListSortOrder,
    MessageRole,
    RunStatus,
    VectorStoreStatus,
)
from azure.core.exceptions import AzureError
from azure.identity import DefaultAzureCredential


GUIDE_FACT = "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket."
QUESTION = "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?"
GUIDE_PATH = Path(__file__).with_name("contoso_trail_guide.txt")


def main() -> None:
    endpoint = os.environ["PROJECT_ENDPOINT"]
    model_deployment = os.environ["MODEL_DEPLOYMENT_NAME"]
    GUIDE_PATH.write_text(GUIDE_FACT, encoding="utf-8")

    uploaded_file_id: str | None = None
    vector_store_id: str | None = None
    agent_id: str | None = None
    thread_id: str | None = None

    with DefaultAzureCredential() as credential, AgentsClient(
        endpoint=endpoint,
        credential=credential,
    ) as client:
        try:
            uploaded_file = client.files.upload_and_poll(
                file_path=str(GUIDE_PATH),
                purpose=FilePurpose.AGENTS,
            )
            uploaded_file_id = uploaded_file.id

            vector_store = client.vector_stores.create_and_poll(
                file_ids=[uploaded_file_id],
                name="hyoka-trail-guide-vector-store",
            )
            vector_store_id = vector_store.id
            if vector_store.status != VectorStoreStatus.COMPLETED:
                raise RuntimeError(
                    "Vector store indexing did not complete successfully "
                    f"(status: {vector_store.status})."
                )

            file_search = FileSearchTool(vector_store_ids=[vector_store_id])
            agent = client.create_agent(
                model=model_deployment,
                name="hyoka-trail-guide-agent",
                instructions=(
                    "Answer questions using the uploaded trail guide. Use file "
                    "search and do not invent details that are absent from the guide."
                ),
                tools=file_search.definitions,
                tool_resources=file_search.resources,
            )
            agent_id = agent.id

            thread = client.threads.create()
            thread_id = thread.id
            client.messages.create(
                thread_id=thread_id,
                role=MessageRole.USER,
                content=QUESTION,
            )

            run = client.runs.create_and_process(
                thread_id=thread_id,
                agent_id=agent_id,
            )
            if run.status != RunStatus.COMPLETED:
                details = f": {run.last_error}" if run.last_error else ""
                raise RuntimeError(
                    f"Agent run ended with status {run.status}{details}"
                )

            messages = client.messages.list(
                thread_id=thread_id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role == MessageRole.AGENT:
                    for text_message in message.text_messages:
                        print(text_message.text.value)
        finally:
            cleanup_errors: list[str] = []
            cleanup_actions = (
                (
                    "thread",
                    thread_id,
                    lambda resource_id: client.threads.delete(resource_id),
                ),
                (
                    "agent",
                    agent_id,
                    lambda resource_id: client.delete_agent(resource_id),
                ),
                (
                    "vector store",
                    vector_store_id,
                    lambda resource_id: client.vector_stores.delete(resource_id),
                ),
                (
                    "uploaded file",
                    uploaded_file_id,
                    lambda resource_id: client.files.delete(file_id=resource_id),
                ),
            )
            for resource_name, resource_id, delete_resource in cleanup_actions:
                if resource_id is None:
                    continue
                try:
                    delete_resource(resource_id)
                except AzureError as error:
                    cleanup_errors.append(f"{resource_name}: {error}")

            if cleanup_errors:
                message = "Cleanup failed for " + "; ".join(cleanup_errors)
                if sys.exc_info()[0] is None:
                    raise RuntimeError(message)
                print(message, file=sys.stderr)


if __name__ == "__main__":
    main()
