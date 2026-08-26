import os
import time

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import ListSortOrder, MessageRole, RunStatus
from azure.identity import DefaultAzureCredential

AGENT_NAME = "hyoka-basic-agent"
AGENT_INSTRUCTIONS = "Answer the user's question clearly and concisely."
USER_MESSAGE = "What is the capital of France?"
POLL_INTERVAL_SECONDS = 1
TERMINAL_STATUSES = {
    RunStatus.COMPLETED,
    RunStatus.FAILED,
    RunStatus.CANCELLED,
    RunStatus.EXPIRED,
}


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    project_endpoint = required_environment_variable("PROJECT_ENDPOINT")
    model_deployment_name = required_environment_variable("MODEL_DEPLOYMENT_NAME")

    credential = DefaultAzureCredential()
    agent_id = None
    thread_id = None

    with credential, AgentsClient(
        endpoint=project_endpoint,
        credential=credential,
    ) as client:
        try:
            agent = client.create_agent(
                model=model_deployment_name,
                name=AGENT_NAME,
                instructions=AGENT_INSTRUCTIONS,
            )
            agent_id = agent.id

            thread = client.threads.create()
            thread_id = thread.id
            client.messages.create(
                thread_id=thread_id,
                role=MessageRole.USER,
                content=USER_MESSAGE,
            )

            run = client.runs.create(thread_id=thread_id, agent_id=agent_id)
            while run.status not in TERMINAL_STATUSES:
                time.sleep(POLL_INTERVAL_SECONDS)
                run = client.runs.get(thread_id=thread_id, run_id=run.id)

            if run.status != RunStatus.COMPLETED:
                details = run.last_error or run.incomplete_details or "No details available."
                raise RuntimeError(f"Agent run ended with status {run.status}: {details}")

            messages = client.messages.list(
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
                    client.threads.delete(thread_id=thread_id)
            finally:
                if agent_id is not None:
                    client.delete_agent(agent_id=agent_id)


if __name__ == "__main__":
    main()
