import os
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
from azure.identity import DefaultAzureCredential


DOCUMENT_FACT = (
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and "
    "hikers should bring a rain jacket."
)
QUESTION = (
    "According to the uploaded guide, how long is the Cascade Loop and what "
    "should hikers bring?"
)
DOCUMENT_PATH = Path(__file__).with_name("contoso_trail_guide.txt")


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    endpoint = required_environment_variable("PROJECT_ENDPOINT")
    model_deployment_name = required_environment_variable("MODEL_DEPLOYMENT_NAME")
    DOCUMENT_PATH.write_text(DOCUMENT_FACT, encoding="utf-8")

    uploaded_file_id: str | None = None
    vector_store_id: str | None = None
    agent_id: str | None = None
    thread_id: str | None = None

    credential = DefaultAzureCredential()
    try:
        with AgentsClient(endpoint=endpoint, credential=credential) as agents_client:
            try:
                uploaded_file = agents_client.files.upload_and_poll(
                    file_path=DOCUMENT_PATH,
                    purpose=FilePurpose.AGENTS,
                )
                uploaded_file_id = uploaded_file.id

                vector_store = agents_client.vector_stores.create_and_poll(
                    file_ids=[uploaded_file_id],
                    name="hyoka-trail-guide-vector-store",
                )
                vector_store_id = vector_store.id
                if vector_store.status != VectorStoreStatus.COMPLETED:
                    raise RuntimeError(
                        "Vector store indexing did not complete successfully "
                        f"(status: {vector_store.status})."
                    )

                file_search = FileSearchTool(
                    vector_store_ids=[vector_store_id]
                )
                agent = agents_client.create_agent(
                    model=model_deployment_name,
                    name="hyoka-trail-guide-agent",
                    instructions=(
                        "Answer questions using the uploaded trail guide. "
                        "Use file search and do not invent facts."
                    ),
                    tools=file_search.definitions,
                    tool_resources=file_search.resources,
                )
                agent_id = agent.id

                thread = agents_client.threads.create()
                thread_id = thread.id
                agents_client.messages.create(
                    thread_id=thread_id,
                    role=MessageRole.USER,
                    content=QUESTION,
                )

                run = agents_client.runs.create_and_process(
                    thread_id=thread_id,
                    agent_id=agent_id,
                )
                if run.status != RunStatus.COMPLETED:
                    detail = f": {run.last_error}" if run.last_error else ""
                    raise RuntimeError(
                        f"Agent run ended with status {run.status}{detail}"
                    )

                messages = agents_client.messages.list(
                    thread_id=thread_id,
                    order=ListSortOrder.ASCENDING,
                )
                for message in messages:
                    if message.role == MessageRole.AGENT:
                        for text_message in message.text_messages:
                            print(text_message.text.value)
            finally:
                try:
                    if thread_id is not None:
                        agents_client.threads.delete(thread_id=thread_id)
                finally:
                    try:
                        if agent_id is not None:
                            agents_client.delete_agent(agent_id=agent_id)
                    finally:
                        try:
                            if vector_store_id is not None:
                                agents_client.vector_stores.delete(
                                    vector_store_id=vector_store_id
                                )
                        finally:
                            if uploaded_file_id is not None:
                                agents_client.files.delete(
                                    file_id=uploaded_file_id
                                )
    finally:
        credential.close()


if __name__ == "__main__":
    main()
