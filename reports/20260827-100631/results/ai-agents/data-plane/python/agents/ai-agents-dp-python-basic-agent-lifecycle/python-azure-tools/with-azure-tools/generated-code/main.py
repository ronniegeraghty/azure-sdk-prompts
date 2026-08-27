import os
import time
from typing import Optional

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import ListSortOrder, MessageRole, RunStatus
from azure.identity import DefaultAzureCredential


POLL_INTERVAL_SECONDS = 1
TERMINAL_RUN_STATUSES = {
    RunStatus.CANCELLED,
    RunStatus.COMPLETED,
    RunStatus.EXPIRED,
    RunStatus.FAILED,
}


def main() -> None:
    project_endpoint = os.environ["PROJECT_ENDPOINT"]
    model_deployment_name = os.environ["MODEL_DEPLOYMENT_NAME"]

    agent_id: Optional[str] = None
    thread_id: Optional[str] = None

    with DefaultAzureCredential() as credential:
        with AgentsClient(endpoint=project_endpoint, credential=credential) as client:
            try:
                agent = client.create_agent(
                    model=model_deployment_name,
                    name="hyoka-basic-agent",
                    instructions="Answer the user's question clearly and concisely.",
                )
                agent_id = agent.id

                thread = client.threads.create()
                thread_id = thread.id

                client.messages.create(
                    thread_id=thread_id,
                    role=MessageRole.USER,
                    content="What is the capital of France?",
                )

                run = client.runs.create(thread_id=thread_id, agent_id=agent_id)
                while run.status not in TERMINAL_RUN_STATUSES:
                    time.sleep(POLL_INTERVAL_SECONDS)
                    run = client.runs.get(thread_id=thread_id, run_id=run.id)

                if run.status != RunStatus.COMPLETED:
                    error_detail = f": {run.last_error}" if run.last_error else ""
                    raise RuntimeError(f"Agent run ended with status '{run.status}'{error_detail}")

                messages = client.messages.list(
                    thread_id=thread_id,
                    order=ListSortOrder.ASCENDING,
                )
                for message in messages:
                    if message.role == MessageRole.AGENT:
                        for text_message in message.text_messages:
                            print(text_message.text.value)
            finally:
                if thread_id is not None:
                    try:
                        client.threads.delete(thread_id)
                    finally:
                        if agent_id is not None:
                            client.delete_agent(agent_id)
                elif agent_id is not None:
                    client.delete_agent(agent_id)


if __name__ == "__main__":
    main()
