import { AgentsClient } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

const agentName = "hyoka-basic-agent";
const agentInstructions = "Answer the user's question clearly and concisely.";
const userMessage = "What is the capital of France?";

const endpoint = requiredEnvironmentVariable("PROJECT_ENDPOINT");
const modelDeployment = requiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
const client = new AgentsClient(endpoint, new DefaultAzureCredential());

let agentId: string | undefined;
let threadId: string | undefined;

try {
  const agent = await client.createAgent(modelDeployment, {
    name: agentName,
    instructions: agentInstructions,
  });
  agentId = agent.id;

  const thread = await client.threads.create();
  threadId = thread.id;

  await client.messages.create(thread.id, "user", userMessage);

  const run = await client.runs.createAndPoll(thread.id, agent.id, {
    pollingOptions: {
      intervalInMs: 500,
    },
  });

  if (run.status !== "completed") {
    throw new Error(`Agent run ended with status ${run.status}.`);
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

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}
