import json
import os
import time
from typing import Literal

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import (
    FunctionDefinition,
    FunctionToolDefinition,
    ListSortOrder,
    RequiredFunctionToolCall,
    SubmitToolOutputsAction,
    ToolOutput,
)
from azure.identity import DefaultAzureCredential


POLL_INTERVAL_SECONDS = 1
ACTIVE_RUN_STATUSES = {"queued", "in_progress", "requires_action"}


def get_weather(location: str, unit: Literal["c", "f"]) -> str:
    """Get deterministic weather for a location.

    :param location: The city whose weather is requested.
    :param unit: The temperature unit, either c for Celsius or f for Fahrenheit.
    :return: A JSON string containing the location, temperature, and unit.
    """
    if unit not in ("c", "f"):
        raise ValueError("unit must be 'c' or 'f'")
    if location.casefold() != "seattle":
        raise ValueError(f"Weather is unavailable for location: {location}")

    temperature = 21 if unit == "c" else 70
    return json.dumps(
        {"location": "Seattle", "temperature": temperature, "unit": unit},
        separators=(",", ":"),
    )


def execute_tool_call(tool_call: RequiredFunctionToolCall) -> ToolOutput:
    """Execute one requested local function and correlate its result."""
    if tool_call.function.name != "get_weather":
        output = json.dumps({"error": f"Unknown function: {tool_call.function.name}"})
    else:
        try:
            arguments = json.loads(tool_call.function.arguments)
            output = get_weather(
                location=arguments["location"],
                unit=arguments["unit"],
            )
        except (json.JSONDecodeError, KeyError, TypeError, ValueError) as error:
            output = json.dumps({"error": str(error)})

    return ToolOutput(tool_call_id=tool_call.id, output=output)


def create_weather_tool() -> FunctionToolDefinition:
    return FunctionToolDefinition(
        function=FunctionDefinition(
            name="get_weather",
            description="Get deterministic weather for a location.",
            parameters={
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "The city whose weather is requested.",
                    },
                    "unit": {
                        "type": "string",
                        "enum": ["c", "f"],
                        "description": "The temperature unit.",
                    },
                },
                "required": ["location", "unit"],
                "additionalProperties": False,
            },
        )
    )


def main() -> None:
    endpoint = os.environ["PROJECT_ENDPOINT"]
    model = os.environ["MODEL_DEPLOYMENT_NAME"]
    weather_tool = create_weather_tool()

    agent_id = None
    thread_id = None

    with AgentsClient(
        endpoint=endpoint,
        credential=DefaultAzureCredential(),
    ) as agents_client:
        try:
            agent = agents_client.create_agent(
                model=model,
                name="hyoka-weather-agent",
                instructions=(
                    "For every weather question, you must call the get_weather "
                    "function. Use its result to answer the user."
                ),
                tools=[weather_tool],
            )
            agent_id = agent.id

            thread = agents_client.threads.create()
            thread_id = thread.id
            agents_client.messages.create(
                thread_id=thread_id,
                role="user",
                content="What is the weather in Seattle in celsius?",
            )
            run = agents_client.runs.create(
                thread_id=thread_id,
                agent_id=agent_id,
            )

            while run.status in ACTIVE_RUN_STATUSES:
                if (
                    run.status == "requires_action"
                    and isinstance(run.required_action, SubmitToolOutputsAction)
                ):
                    tool_outputs = [
                        execute_tool_call(tool_call)
                        for tool_call in run.required_action.submit_tool_outputs.tool_calls
                        if isinstance(tool_call, RequiredFunctionToolCall)
                    ]
                    if not tool_outputs:
                        raise RuntimeError(
                            "The run required action but provided no function tool calls."
                        )
                    run = agents_client.runs.submit_tool_outputs(
                        thread_id=thread_id,
                        run_id=run.id,
                        tool_outputs=tool_outputs,
                    )
                else:
                    time.sleep(POLL_INTERVAL_SECONDS)
                    run = agents_client.runs.get(
                        thread_id=thread_id,
                        run_id=run.id,
                    )

            if run.status != "completed":
                details = getattr(run, "last_error", None)
                raise RuntimeError(
                    f"Agent run ended with status {run.status!s}: {details}"
                )

            messages = agents_client.messages.list(
                thread_id=thread_id,
                order=ListSortOrder.ASCENDING,
            )
            for message in messages:
                if message.role == "assistant":
                    for text_message in message.text_messages:
                        print(text_message.text.value)
        finally:
            try:
                if thread_id is not None:
                    agents_client.threads.delete(thread_id)
            finally:
                if agent_id is not None:
                    agents_client.delete_agent(agent_id)


if __name__ == "__main__":
    main()
