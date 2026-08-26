import os
import tempfile
from pathlib import Path

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import FilePurpose, FileSearchTool, ListSortOrder
from azure.identity import DefaultAzureCredential

GUIDE = (
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long "
    "and hikers should bring a rain jacket."
)
QUESTION = (
    "According to the uploaded guide, how long is the Cascade Loop "
    "and what should hikers bring?"
)


def main() -> None:
    endpoint = require_environment_variable("PROJECT_ENDPOINT")
    model = require_environment_variable("MODEL_DEPLOYMENT_NAME")
    with tempfile.NamedTemporaryFile(
        mode="w",
        encoding="utf-8",
        suffix=".txt",
        delete=False,
    ) as document:
        document.write(GUIDE)
        document_path = Path(document.name)

    uploaded_file = None
    vector_store = None
    agent = None
    thread = None
    with AgentsClient(endpoint, DefaultAzureCredential()) as client:
        try:
            uploaded_file = client.files.upload_and_poll(
                file_path=str(document_path),
                purpose=FilePurpose.AGENTS,
            )
            vector_store = client.vector_stores.create_and_poll(
                file_ids=[uploaded_file.id],
                name="hyoka-trail-guide",
            )
            if vector_store.status != "completed":
                raise RuntimeError(
                    f"Vector store ended with status {vector_store.status}."
                )

            search_tool = FileSearchTool(vector_store_ids=[vector_store.id])
            agent = client.create_agent(
                model=model,
                name="hyoka-trail-guide-agent",
                instructions=(
                    "Use file search to answer questions about the uploaded guide."
                ),
                tools=search_tool.definitions,
                tool_resources=search_tool.resources,
            )
            thread = client.threads.create()
            client.messages.create(
                thread_id=thread.id,
                role="user",
                content=QUESTION,
            )
            run = client.runs.create_and_process(
                thread_id=thread.id,
                agent_id=agent.id,
            )
            if run.status != "completed":
                raise RuntimeError(f"Run ended with status {run.status}.")

            messages = client.messages.list(
                thread_id=thread.id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role == "assistant":
                    for text_message in message.text_messages:
                        print(text_message.text.value)
        finally:
            if thread is not None:
                client.threads.delete(thread.id)
            if agent is not None:
                client.delete_agent(agent.id)
            if vector_store is not None:
                client.vector_stores.delete(vector_store.id)
            if uploaded_file is not None:
                client.files.delete(file_id=uploaded_file.id)
            document_path.unlink(missing_ok=True)


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
