import type {
  FunctionToolDefinition,
  RequiredFunctionToolCall,
  SubmitToolOutputsAction,
  ThreadRun,
  ToolOutput,
} from "@azure/ai-agents";
import { AgentsClient, ToolUtility, isOutputOfType } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

const endpoint = requiredEnvironmentVariable("PROJECT_ENDPOINT");
const model = requiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
const client = new AgentsClient(endpoint, new DefaultAzureCredential());

const weatherTool = ToolUtility.createFunctionTool({
  name: "get_weather",
  description: "Gets the current weather for a location.",
  parameters: {
    type: "object",
    properties: {
      location: { type: "string" },
      unit: { type: "string", enum: ["c", "f"] },
    },
    required: ["location", "unit"],
  },
});

let agentId: string | undefined;
let threadId: string | undefined;

try {
  const agent = await client.createAgent(model, {
    name: "hyoka-weather-agent",
    instructions: "Use get_weather for every weather question.",
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

  const onResponse = async (response: {
    parsedBody?: unknown;
  }): Promise<void> => {
    if (
      typeof response.parsedBody !== "object"
      || response.parsedBody === null
      || !("status" in response.parsedBody)
    ) {
      return;
    }

    const run = response.parsedBody as ThreadRun;
    if (
      run.status !== "requires_action"
      || !run.requiredAction
      || !isOutputOfType<SubmitToolOutputsAction>(
        run.requiredAction,
        "submit_tool_outputs",
      )
    ) {
      return;
    }

    const outputs: ToolOutput[] = [];
    for (const call of run.requiredAction.submitToolOutputs.toolCalls) {
      if (
        !isOutputOfType<RequiredFunctionToolCall>(call, "function")
        || call.function.name !== "get_weather"
      ) {
        throw new Error("Unexpected tool call.");
      }
      const args = JSON.parse(call.function.arguments) as {
        location: string;
        unit: string;
      };
      outputs.push({
        toolCallId: call.id,
        output: JSON.stringify(getWeather(args.location, args.unit)),
      });
    }
    await client.runs.submitToolOutputs(thread.id, run.id, outputs);
  };

  const run = await client.runs.createAndPoll(thread.id, agent.id, {
    pollingOptions: { intervalInMs: 500 },
    onResponse,
  });
  if (run.status !== "completed") {
    throw new Error(`Run ended with status ${run.status}.`);
  }

  const messages = client.messages.list(thread.id, { order: "asc" });
  for await (const message of messages) {
    if (message.role !== "assistant") {
      continue;
    }
    for (const content of message.content) {
      if (content.type === "text" && "text" in content) {
        console.log(content.text.value);
      }
    }
  }
} finally {
  if (threadId !== undefined) {
    await client.threads.delete(threadId);
  }
  if (agentId !== undefined) {
    await client.deleteAgent(agentId);
  }
}

function getWeather(
  location: string,
  unit: string,
): { location: string; temperature: number; unit: string } {
  if (!location.toLowerCase().includes("seattle") || !["c", "f"].includes(unit)) {
    throw new Error("Unsupported weather request.");
  }
  return {
    location: "Seattle",
    temperature: unit === "c" ? 21 : 70,
    unit,
  };
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}
