from __future__ import annotations

import os
import time
from contextlib import ExitStack

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


def _required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    project_endpoint = _required_environment_variable("PROJECT_ENDPOINT")
    model_deployment_name = _required_environment_variable("MODEL_DEPLOYMENT_NAME")

    credential = DefaultAzureCredential()
    with credential, AgentsClient(
        endpoint=project_endpoint,
        credential=credential,
    ) as agents_client:
        with ExitStack() as resources:
            agent = agents_client.create_agent(
                model=model_deployment_name,
                name=AGENT_NAME,
                instructions=AGENT_INSTRUCTIONS,
            )
            resources.callback(agents_client.delete_agent, agent_id=agent.id)

            thread = agents_client.threads.create()
            resources.callback(agents_client.threads.delete, thread_id=thread.id)

            agents_client.messages.create(
                thread_id=thread.id,
                role=MessageRole.USER,
                content=USER_MESSAGE,
            )

            run = agents_client.runs.create(
                thread_id=thread.id,
                agent_id=agent.id,
            )
            while run.status not in TERMINAL_STATUSES:
                time.sleep(POLL_INTERVAL_SECONDS)
                run = agents_client.runs.get(
                    thread_id=thread.id,
                    run_id=run.id,
                )

            if run.status != RunStatus.COMPLETED:
                raise RuntimeError(
                    f"Agent run ended with status {run.status}: {run.last_error}"
                )

            messages = agents_client.messages.list(
                thread_id=thread.id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role == MessageRole.AGENT:
                    for text_message in message.text_messages:
                        print(text_message.text.value)
