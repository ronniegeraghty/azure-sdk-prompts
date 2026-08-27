import os
import time

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import ListSortOrder
from azure.identity import DefaultAzureCredential

AGENT_NAME = "hyoka-basic-agent"
AGENT_INSTRUCTIONS = "Answer the user's question clearly and concisely."
USER_MESSAGE = "What is the capital of France?"
POLL_INTERVAL_SECONDS = 1
TERMINAL_STATUSES = {"cancelled", "completed", "expired", "failed", "incomplete"}


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    project_endpoint = required_environment_variable("PROJECT_ENDPOINT")
    model_deployment_name = required_environment_variable("MODEL_DEPLOYMENT_NAME")

    credential = DefaultAzureCredential()
    agent = None
    thread = None

    try:
        with AgentsClient(endpoint=project_endpoint, credential=credential) as client:
            try:
                agent = client.create_agent(
                    model=model_deployment_name,
                    name=AGENT_NAME,
                    instructions=AGENT_INSTRUCTIONS,
                )
                thread = client.threads.create()
                client.messages.create(
                    thread_id=thread.id,
                    role="user",
                    content=USER_MESSAGE,
                )

                run = client.runs.create(thread_id=thread.id, agent_id=agent.id)
                while run.status not in TERMINAL_STATUSES:
                    time.sleep(POLL_INTERVAL_SECONDS)
                    run = client.runs.get(thread_id=thread.id, run_id=run.id)

                if run.status != "completed":
                    raise RuntimeError(
                        f"Agent run ended with status '{run.status}': {run.last_error}"
                    )

                messages = client.messages.list(
                    thread_id=thread.id,
                    order=ListSortOrder.ASCENDING,
                )
                for message in messages:
                    if message.role == "assistant":
                        for text_message in message.text_messages:
                            print(text_message.text.value)
            finally:
                try:
                    if thread is not None:
                        client.threads.delete(thread.id)
                finally:
                    if agent is not None:
                        client.delete_agent(agent.id)
    finally:
        credential.close()


if __name__ == "__main__":
    main()
