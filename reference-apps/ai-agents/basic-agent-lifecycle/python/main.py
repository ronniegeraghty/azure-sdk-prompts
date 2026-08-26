import os
import time

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import ListSortOrder
from azure.identity import DefaultAzureCredential

AGENT_NAME = "hyoka-basic-agent"
AGENT_INSTRUCTIONS = "Answer the user's question clearly and concisely."
USER_MESSAGE = "What is the capital of France?"


def main() -> None:
    endpoint = require_environment_variable("PROJECT_ENDPOINT")
    model_deployment = require_environment_variable("MODEL_DEPLOYMENT_NAME")

    agent = None
    thread = None
    with AgentsClient(endpoint, DefaultAzureCredential()) as client:
        try:
            agent = client.create_agent(
                model=model_deployment,
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
            while run.status in ("queued", "in_progress"):
                time.sleep(0.5)
                run = client.runs.get(thread_id=thread.id, run_id=run.id)

            if run.status != "completed":
                raise RuntimeError(f"Agent run ended with status {run.status}.")

            messages = client.messages.list(
                thread_id=thread.id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role != "assistant":
                    continue
                for text_message in message.text_messages:
                    print(text_message.text.value)
        finally:
            if thread is not None:
                client.threads.delete(thread.id)
            if agent is not None:
                client.delete_agent(agent.id)


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
