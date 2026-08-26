import { AgentsClient, isOutputOfType, ToolUtility } from "@azure/ai-agents";
import type { MessageTextContent } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";
import { createReadStream } from "node:fs";
import { writeFile } from "node:fs/promises";
import { resolve } from "node:path";

const documentText =
  "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const question =
  "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";
const documentPath = resolve("trail-guide.txt");

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing required environment variable: ${name}`);
  }
  return value;
}

async function main(): Promise<void> {
  const projectEndpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
  const modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
  const client = new AgentsClient(projectEndpoint, new DefaultAzureCredential());

  let uploadedFileId: string | undefined;
  let vectorStoreId: string | undefined;
  let agentId: string | undefined;
  let threadId: string | undefined;

  await writeFile(documentPath, documentText, "utf8");

  try {
    const uploadedFile = await client.files.upload(
      createReadStream(documentPath),
      "assistants",
      { fileName: "trail-guide.txt" },
    );
    uploadedFileId = uploadedFile.id;

    const vectorStore = await client.vectorStores.create({
      name: "hyoka-trail-guide-vector-store",
    });
    vectorStoreId = vectorStore.id;

    const indexedFile = await client.vectorStoreFiles
      .createAndPoll(vectorStore.id, {
        fileId: uploadedFile.id,
        pollingOptions: { intervalInMs: 2_000 },
      })
      .pollUntilDone();

    if (indexedFile.status !== "completed") {
      throw new Error(
        `Document indexing failed with status "${indexedFile.status}"` +
          (indexedFile.lastError ? `: ${indexedFile.lastError.message}` : ""),
      );
    }

    const fileSearchTool = ToolUtility.createFileSearchTool([vectorStore.id]);
    const agent = await client.createAgent(modelDeploymentName, {
      name: "hyoka-trail-guide-agent",
      instructions:
        "Answer questions using the uploaded guide. Base factual claims on file search results.",
      tools: [fileSearchTool.definition],
      toolResources: fileSearchTool.resources,
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;
    await client.messages.create(thread.id, "user", question);

    const run = await client.runs.createAndPoll(thread.id, agent.id, {
      pollingOptions: { intervalInMs: 2_000 },
    });

    if (run.status !== "completed") {
      throw new Error(
        `Agent run ended with status "${run.status}"` +
          (run.lastError ? `: ${run.lastError.message}` : ""),
      );
    }

    const messages = [];
    for await (const message of client.messages.list(thread.id, {
      order: "asc",
    })) {
      messages.push(message);
    }

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
    const cleanupErrors: unknown[] = [];

    const cleanUp = async (operation: () => Promise<unknown>): Promise<void> => {
      try {
        await operation();
      } catch (error: unknown) {
        cleanupErrors.push(error);
      }
    };

    if (threadId) {
      const id = threadId;
      await cleanUp(() => client.threads.delete(id));
    }
    if (agentId) {
      const id = agentId;
      await cleanUp(() => client.deleteAgent(id));
    }
    if (vectorStoreId) {
      const id = vectorStoreId;
      await cleanUp(() => client.vectorStores.delete(id));
    }
    if (uploadedFileId) {
      const id = uploadedFileId;
      await cleanUp(() => client.files.delete(id));
    }

    if (cleanupErrors.length > 0) {
      throw new AggregateError(cleanupErrors, "One or more Azure resources could not be deleted");
    }
  }
}

main().catch((error: unknown) => {
  console.error(error);
  process.exitCode = 1;
});
