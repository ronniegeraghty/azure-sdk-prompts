import {
  AgentsClient,
  isOutputOfType,
  ToolUtility,
  type MessageContentUnion,
  type MessageTextContent,
  type RequiredFunctionToolCall,
  type SubmitToolOutputsAction,
  type ThreadRun,
  type ToolOutput,
} from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

const userMessage = "What is the weather in Seattle in celsius?";
const activeStatuses = new Set(["queued", "in_progress", "requires_action"]);

type TemperatureUnit = "c" | "f";

interface WeatherArguments {
  location: string;
  unit: TemperatureUnit;
}

interface WeatherResult extends WeatherArguments {
  temperature: number;
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

function decodeWeatherArguments(serializedArguments: string): WeatherArguments {
  const value: unknown = JSON.parse(serializedArguments);
  if (
    typeof value !== "object" ||
    value === null ||
    !("location" in value) ||
    typeof value.location !== "string" ||
    !("unit" in value) ||
    (value.unit !== "c" && value.unit !== "f")
  ) {
    throw new Error(`Invalid get_weather arguments: ${serializedArguments}`);
  }

  return { location: value.location, unit: value.unit };
}

function isMessageTextContent(
  content: MessageContentUnion,
): content is MessageTextContent {
  return content.type === "text" && "text" in content;
}

async function getWeather(
  location: string,
  unit: TemperatureUnit,
): Promise<WeatherResult> {
  if (location.trim().toLowerCase() !== "seattle") {
    throw new Error(`Unsupported location: ${location}`);
  }

  return {
    location,
    unit,
    temperature: unit === "c" ? 21 : 70,
  };
}

async function wait(milliseconds: number): Promise<void> {
  await new Promise<void>((resolve) => setTimeout(resolve, milliseconds));
}

async function main(): Promise<void> {
  const projectEndpoint = requiredEnvironmentVariable("PROJECT_ENDPOINT");
  const modelDeploymentName = requiredEnvironmentVariable(
    "MODEL_DEPLOYMENT_NAME",
  );

  const client = new AgentsClient(
    projectEndpoint,
    new DefaultAzureCredential(),
  );

  const weatherTool = ToolUtility.createFunctionTool({
    name: "get_weather",
    description: "Get the deterministic weather for a location.",
    parameters: {
      type: "object",
      properties: {
        location: {
          type: "string",
          description: "The city whose weather is requested.",
        },
        unit: {
          type: "string",
          description: "The temperature unit.",
          enum: ["c", "f"],
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
        "You answer weather questions. You must use the get_weather function for every weather question and base your answer on its result.",
      tools: [weatherTool.definition],
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;

    await client.messages.create(thread.id, "user", userMessage);

    let run: ThreadRun = await client.runs.create(thread.id, agent.id);

    while (activeStatuses.has(run.status)) {
      if (run.status === "requires_action") {
        if (
          !run.requiredAction ||
          !isOutputOfType<SubmitToolOutputsAction>(
            run.requiredAction,
            "submit_tool_outputs",
          )
        ) {
          throw new Error("Run requested an unsupported action.");
        }

        const outputs: ToolOutput[] = [];

        for (const toolCall of run.requiredAction.submitToolOutputs.toolCalls) {
          if (!isOutputOfType<RequiredFunctionToolCall>(toolCall, "function")) {
            throw new Error(`Unsupported tool call type: ${toolCall.type}`);
          }
          if (toolCall.function.name !== "get_weather") {
            throw new Error(
              `Unsupported function call: ${toolCall.function.name}`,
            );
          }

          const args = decodeWeatherArguments(toolCall.function.arguments);
          const result = await getWeather(args.location, args.unit);
          outputs.push({
            toolCallId: toolCall.id,
            output: JSON.stringify(result),
          });
        }

        if (outputs.length === 0) {
          throw new Error("Run requested action without function tool calls.");
        }

        run = await client.runs.submitToolOutputs(
          thread.id,
          run.id,
          outputs,
        );
      } else {
        await wait(1_000);
        run = await client.runs.get(thread.id, run.id);
      }
    }

    if (run.status !== "completed") {
      throw new Error(
        `Run ended with status ${run.status}: ${JSON.stringify(run.lastError)}`,
      );
    }

    const messages = [];
    for await (const message of client.messages.list(thread.id)) {
      messages.push(message);
    }
    messages.reverse();

    for (const message of messages) {
      if (message.role !== "assistant") {
        continue;
      }
      for (const content of message.content) {
        if (isMessageTextContent(content)) {
          console.log(content.text.value);
        }
      }
    }
  } finally {
    try {
      if (threadId) {
        await client.threads.delete(threadId);
      }
    } finally {
      if (agentId) {
        await client.deleteAgent(agentId);
      }
    }
  }
}

await main();
