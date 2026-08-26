import {
  AgentsClient,
  type MessageContentUnion,
  type MessageTextContent,
  type RunStatus,
} from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

const AGENT_NAME = "hyoka-basic-agent";
const AGENT_INSTRUCTIONS = "Answer the user's question clearly and concisely.";
const USER_MESSAGE = "What is the capital of France?";
const POLL_INTERVAL_MS = 1_000;
const TERMINAL_STATUSES = new Set<RunStatus>(["completed", "failed", "cancelled", "expired"]);

function requireEnvironmentVariable(name: "PROJECT_ENDPOINT" | "MODEL_DEPLOYMENT_NAME"): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Environment variable ${name} is required.`);
  }

  return value;
}

async function delay(milliseconds: number): Promise<void> {
  await new Promise<void>((resolve) => setTimeout(resolve, milliseconds));
}

function isTextContent(content: MessageContentUnion): content is MessageTextContent {
  return content.type === "text";
}

async function main(): Promise<void> {
  const projectEndpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
  const modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  let agentId: string | undefined;
  let threadId: string | undefined;

  try {
    const agent = await client.createAgent(modelDeploymentName, {
      name: AGENT_NAME,
      instructions: AGENT_INSTRUCTIONS,
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;

    await client.messages.create(thread.id, "user", USER_MESSAGE);

    let run = await client.runs.create(thread.id, agent.id);
    while (!TERMINAL_STATUSES.has(run.status)) {
      await delay(POLL_INTERVAL_MS);
      run = await client.runs.get(thread.id, run.id);
    }

    if (run.status !== "completed") {
      const details = run.lastError?.message ? `: ${run.lastError.message}` : "";
      throw new Error(`Agent run ended with status "${run.status}"${details}`);
    }

    const messages = client.messages.list(thread.id, { order: "asc" });
    for await (const message of messages) {
      if (message.role !== "assistant") {
        continue;
      }

      for (const content of message.content) {
        if (isTextContent(content)) {
          console.log(content.text.value);
        }
      }
    }
  } finally {
    const cleanupOperations: Promise<unknown>[] = [];
    if (threadId) {
      cleanupOperations.push(client.threads.delete(threadId));
    }
    if (agentId) {
      cleanupOperations.push(client.deleteAgent(agentId));
    }
    await Promise.all(cleanupOperations);
  }
}

main().catch((error: unknown) => {
  console.error(error instanceof Error ? error.message : error);
  process.exitCode = 1;
});
