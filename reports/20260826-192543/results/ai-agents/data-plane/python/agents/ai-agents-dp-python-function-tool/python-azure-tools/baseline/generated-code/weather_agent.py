import json
import os
import time
from typing import Any

from azure.ai.agents import AgentsClient
from azure.ai.agents.models import (
    FunctionTool,
    ListSortOrder,
    MessageRole,
    RequiredFunctionToolCall,
    SubmitToolOutputsAction,
    ToolOutput,
)
from azure.identity import DefaultAzureCredential


POLL_INTERVAL_SECONDS = 1
TERMINAL_STATUSES = {"completed", "failed", "cancelled", "expired", "incomplete"}


def get_weather(location: str, unit: str) -> str:
    """Get deterministic weather data for a location.

    :param location: The city whose weather is requested.
    :param unit: The temperature unit: c for Celsius or f for Fahrenheit.
    :return: A JSON string containing the weather result.
    """
    if unit not in {"c", "f"}:
        raise ValueError("unit must be 'c' or 'f'")

    result: dict[str, Any] = {"location": location, "unit": unit}
    if location.casefold() == "seattle":
        result["temperature"] = 21 if unit == "c" else 70
    else:
        result["error"] = "Weather data is only available for Seattle."
    return json.dumps(result)


def create_weather_tool() -> FunctionTool:
    weather_tool = FunctionTool(functions={get_weather})
    parameters = weather_tool.definitions[0]["function"]["parameters"]
    parameters["properties"]["unit"]["enum"] = ["c", "f"]
    parameters["required"] = ["location", "unit"]
    return weather_tool


def execute_tool_calls(run: Any) -> list[ToolOutput]:
    if not isinstance(run.required_action, SubmitToolOutputsAction):
        raise RuntimeError(f"Unsupported required action: {run.required_action!r}")

    tool_outputs: list[ToolOutput] = []
    for tool_call in run.required_action.submit_tool_outputs.tool_calls:
        if not isinstance(tool_call, RequiredFunctionToolCall):
            raise RuntimeError(f"Unsupported tool call: {tool_call!r}")
        if tool_call.function.name != "get_weather":
            raise RuntimeError(f"Unknown function requested: {tool_call.function.name}")

        arguments = json.loads(tool_call.function.arguments)
        if not isinstance(arguments, dict):
            raise ValueError("Function arguments must decode to a JSON object.")

        output = get_weather(
            location=arguments["location"],
            unit=arguments["unit"],
        )
        tool_outputs.append(ToolOutput(tool_call_id=tool_call.id, output=output))

    if not tool_outputs:
        raise RuntimeError("The run requires action but supplied no function calls.")
    return tool_outputs


def main() -> None:
    endpoint = os.environ["PROJECT_ENDPOINT"]
    model = os.environ["MODEL_DEPLOYMENT_NAME"]
    weather_tool = create_weather_tool()

    with DefaultAzureCredential() as credential:
        with AgentsClient(endpoint=endpoint, credential=credential) as client:
            agent = None
            thread = None
            try:
                agent = client.create_agent(
                    model=model,
                    name="hyoka-weather-agent",
                    instructions=(
                        "Answer weather questions by calling get_weather. "
                        "Always use get_weather for weather information; do not guess."
                    ),
                    tools=weather_tool.definitions,
                )
                thread = client.threads.create()
                client.messages.create(
                    thread_id=thread.id,
                    role=MessageRole.USER,
                    content="What is the weather in Seattle in celsius?",
                )
                run = client.runs.create(thread_id=thread.id, agent_id=agent.id)

                while run.status not in TERMINAL_STATUSES:
                    if run.status == "requires_action":
                        tool_outputs = execute_tool_calls(run)
                        client.runs.submit_tool_outputs(
                            thread_id=thread.id,
                            run_id=run.id,
                            tool_outputs=tool_outputs,
                        )
                    elif run.status not in {"queued", "in_progress"}:
                        raise RuntimeError(f"Unexpected run status: {run.status}")

                    time.sleep(POLL_INTERVAL_SECONDS)
                    run = client.runs.get(thread_id=thread.id, run_id=run.id)

                if run.status != "completed":
                    raise RuntimeError(
                        f"Run ended with status {run.status}: {run.last_error}"
                    )

                messages = client.messages.list(
                    thread_id=thread.id,
                    order=ListSortOrder.ASCENDING,
                )
                for message in messages:
                    if message.role == MessageRole.AGENT:
                        for text_message in message.text_messages:
                            print(text_message.text.value)
            finally:
                try:
                    if thread is not None:
                        client.threads.delete(thread.id)
                finally:
                    if agent is not None:
                        client.delete_agent(agent.id)


if __name__ == "__main__":
    main()
