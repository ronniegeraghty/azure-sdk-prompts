import {
  AgentsClient,
  isOutputOfType,
  type MessageTextContent,
  type RunStatus,
} from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";

const USER_MESSAGE = "What is the capital of France?";
const TERMINAL_STATUSES = new Set<RunStatus>([
  "cancelled",
  "completed",
  "expired",
  "failed",
]);

function getRequiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();

  if (!value) {
    throw new Error(`The ${name} environment variable is required.`);
  }

  return value;
}

function delay(milliseconds: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

async function runConversation(): Promise<void> {
  const projectEndpoint = getRequiredEnvironmentVariable("PROJECT_ENDPOINT");
  const modelDeploymentName = getRequiredEnvironmentVariable(
    "MODEL_DEPLOYMENT_NAME",
  );

  const client = new AgentsClient(
    projectEndpoint,
    new DefaultAzureCredential(),
  );

  let agentId: string | undefined;
  let threadId: string | undefined;
  let workflowError: unknown;

  try {
    const agent = await client.createAgent(modelDeploymentName, {
      name: "hyoka-basic-agent",
      instructions: "Answer the user's question clearly and concisely.",
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;

    await client.messages.create(thread.id, "user", USER_MESSAGE);

    let run = await client.runs.create(thread.id, agent.id);

    while (!TERMINAL_STATUSES.has(run.status)) {
      await delay(1_000);
      run = await client.runs.get(thread.id, run.id);
    }

    if (run.status !== "completed") {
      const details = run.lastError
        ? `: ${run.lastError.code} - ${run.lastError.message}`
        : "";
      throw new Error(`Agent run ended with status "${run.status}"${details}`);
    }

    const messages = client.messages.list(thread.id, { order: "asc" });

    for await (const message of messages) {
      if (message.role !== "assistant") {
        continue;
      }

      for (const content of message.content) {
        if (isOutputOfType<MessageTextContent>(content, "text")) {
          console.log(content.text.value);
        }
      }
    }
  } catch (error) {
    workflowError = error;
  }

  const cleanupResults = await Promise.allSettled([
    ...(threadId ? [client.threads.delete(threadId)] : []),
    ...(agentId ? [client.deleteAgent(agentId)] : []),
  ]);
  const cleanupErrors = cleanupResults
    .filter((result) => result.status === "rejected")
    .map((result) => result.reason);
  const errors = workflowError
    ? [workflowError, ...cleanupErrors]
    : cleanupErrors;

  if (errors.length === 1) {
    throw errors[0];
  }

  if (errors.length > 1) {
    throw new AggregateError(errors, "The conversation or cleanup failed.");
  }
}

runConversation().catch((error: unknown) => {
  console.error(error instanceof Error ? error.message : error);
  process.exitCode = 1;
});
