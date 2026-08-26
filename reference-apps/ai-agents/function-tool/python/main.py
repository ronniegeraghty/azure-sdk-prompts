import json
import os
import time
from typing import Literal

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import (
    FunctionTool,
    ListSortOrder,
    RequiredFunctionToolCall,
    SubmitToolOutputsAction,
    ToolOutput,
)
from azure.identity import DefaultAzureCredential


def get_weather(location: str, unit: Literal["c", "f"]) -> str:
    """Get the current weather for a location.

    :param location: The city for the weather request.
    :param unit: The temperature unit, either c or f.
    """
    if "seattle" not in location.lower() or unit not in ("c", "f"):
        raise ValueError("Unsupported weather request.")
    return json.dumps(
        {
            "location": "Seattle",
            "temperature": 21 if unit == "c" else 70,
            "unit": unit,
        }
    )


def main() -> None:
    endpoint = require_environment_variable("PROJECT_ENDPOINT")
    model = require_environment_variable("MODEL_DEPLOYMENT_NAME")
    functions = FunctionTool(functions={get_weather})

    agent = None
    thread = None
    with AgentsClient(endpoint, DefaultAzureCredential()) as client:
        try:
            agent = client.create_agent(
                model=model,
                name="hyoka-weather-agent",
                instructions="Use get_weather for every weather question.",
                tools=functions.definitions,
            )
            thread = client.threads.create()
            client.messages.create(
                thread_id=thread.id,
                role="user",
                content="What is the weather in Seattle in celsius?",
            )
            run = client.runs.create(thread_id=thread.id, agent_id=agent.id)

            while run.status in ("queued", "in_progress", "requires_action"):
                time.sleep(0.5)
                run = client.runs.get(thread_id=thread.id, run_id=run.id)
                if (
                    run.status == "requires_action"
                    and isinstance(run.required_action, SubmitToolOutputsAction)
                ):
                    outputs = []
                    for call in run.required_action.submit_tool_outputs.tool_calls:
                        if not isinstance(call, RequiredFunctionToolCall):
                            raise RuntimeError("Unexpected tool call type.")
                        if call.function.name != "get_weather":
                            raise RuntimeError("Unexpected function name.")
                        outputs.append(
                            ToolOutput(
                                tool_call_id=call.id,
                                output=functions.execute(call),
                            )
                        )
                    run = client.runs.submit_tool_outputs(
                        thread_id=thread.id,
                        run_id=run.id,
                        tool_outputs=outputs,
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


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
