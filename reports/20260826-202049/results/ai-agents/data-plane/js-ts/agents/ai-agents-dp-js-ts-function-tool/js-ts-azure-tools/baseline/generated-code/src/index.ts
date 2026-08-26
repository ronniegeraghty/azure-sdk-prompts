import {
  AgentsClient,
  isOutputOfType,
  ToolUtility,
  type FunctionToolDefinition,
  type MessageTextContent,
  type SubmitToolOutputsAction,
  type ThreadRun,
  type ToolOutput,
} from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

type TemperatureUnit = "c" | "f";

interface WeatherArguments {
  location: string;
  unit: TemperatureUnit;
}

interface WeatherResult extends WeatherArguments {
  temperature: number;
}

const terminalStatuses = new Set([
  "completed",
  "failed",
  "cancelled",
  "expired",
  "incomplete",
]);

function requireEnvironmentVariable(name: "PROJECT_ENDPOINT" | "MODEL_DEPLOYMENT_NAME"): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function decodeWeatherArguments(encodedArguments: string): WeatherArguments {
  const parsed: unknown = JSON.parse(encodedArguments);
  if (
    typeof parsed !== "object" ||
    parsed === null ||
    !("location" in parsed) ||
    typeof parsed.location !== "string" ||
    !("unit" in parsed) ||
    (parsed.unit !== "c" && parsed.unit !== "f")
  ) {
    throw new Error("get_weather requires string location and unit ('c' or 'f') arguments");
  }

  return { location: parsed.location, unit: parsed.unit };
}

async function getWeather(location: string, unit: TemperatureUnit): Promise<WeatherResult> {
  if (location.trim().toLowerCase() !== "seattle") {
    throw new Error(`Weather data is unavailable for location: ${location}`);
  }

  return {
    location,
    unit,
    temperature: unit === "c" ? 21 : 70,
  };
}

async function main(): Promise<void> {
  const projectEndpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
  const modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  const weatherTool = ToolUtility.createFunctionTool({
    name: "get_weather",
    description: "Gets deterministic weather data for a location.",
    parameters: {
      type: "object",
      properties: {
        location: {
          type: "string",
          description: "The city whose weather is requested.",
        },
        unit: {
          type: "string",
          enum: ["c", "f"],
          description: "The temperature unit: c for Celsius or f for Fahrenheit.",
        },
      },
      required: ["location", "unit"],
      additionalProperties: false,
    },
  });

  let agentId: string | undefined;
  let threadId: string | undefined;

  try {
    const agent = await client.createAgent(modelDeploymentName, {
      name: "hyoka-weather-agent",
      instructions:
        "Answer weather questions by calling get_weather. You must use get_weather for every weather question and base your answer on its result.",
      tools: [weatherTool.definition],
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;
    await client.messages.create(
      thread.id,
      "user",
      "What is the weather in Seattle in celsius?",
    );

    const handleResponse = async (response: { parsedBody?: ThreadRun }): Promise<void> => {
      const run = response.parsedBody;
      if (run?.status !== "requires_action" || !run.requiredAction) {
        return;
      }
      if (!isOutputOfType<SubmitToolOutputsAction>(run.requiredAction, "submit_tool_outputs")) {
        throw new Error(`Unsupported required action for run ${run.id}`);
      }

      const outputs: ToolOutput[] = [];
      for (const toolCall of run.requiredAction.submitToolOutputs.toolCalls) {
        if (!isOutputOfType<FunctionToolDefinition>(toolCall, "function")) {
          throw new Error(`Unsupported tool call type for tool call ${toolCall.id}`);
        }
        if (toolCall.function.name !== "get_weather") {
          throw new Error(`Unknown function tool: ${toolCall.function.name}`);
        }

        const args = decodeWeatherArguments(toolCall.function.parameters);
        const result = await getWeather(args.location, args.unit);
        outputs.push({
          toolCallId: toolCall.id,
          output: JSON.stringify(result),
        });
      }

      await client.runs.submitToolOutputs(thread.id, run.id, outputs);
    };

    const run = await client.runs.createAndPoll(thread.id, agent.id, {
      pollingOptions: {
        intervalInMs: 1000,
      },
      onResponse: handleResponse,
    });

    if (!terminalStatuses.has(run.status)) {
      throw new Error(`Polling stopped before the run reached a terminal status: ${run.status}`);
    }
    if (run.status !== "completed") {
      const details = run.lastError ? `: ${JSON.stringify(run.lastError)}` : "";
      throw new Error(`Run ended with status ${run.status}${details}`);
    }

    const messages = [];
    for await (const message of client.messages.list(thread.id)) {
      messages.push(message);
    }
    messages.sort((left, right) => left.createdAt.getTime() - right.createdAt.getTime());

    for (const message of messages) {
      if (message.role !== "assistant") {
        continue;
      }
      for (const content of message.content) {
        if (isOutputOfType<MessageTextContent>(content, "text")) {
          console.log(content.text.value);
        }
      }
    }
  } finally {
    if (threadId) {
      await client.threads.delete(threadId);
    }
    if (agentId) {
      await client.deleteAgent(agentId);
    }
  }
}

await main().catch((error: unknown) => {
  console.error(error instanceof Error ? error.message : error);
  process.exitCode = 1;
});
