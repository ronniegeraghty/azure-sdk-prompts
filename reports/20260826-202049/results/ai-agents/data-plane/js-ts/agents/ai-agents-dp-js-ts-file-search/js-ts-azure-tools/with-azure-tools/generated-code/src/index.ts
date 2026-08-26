import {
  AgentsClient,
  isOutputOfType,
  ToolUtility,
  type MessageTextContent,
} from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";
import { createReadStream } from "node:fs";
import { mkdir, writeFile } from "node:fs/promises";
import { join } from "node:path";

const DOCUMENT_FACT =
  "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const USER_QUESTION =
  "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";
const AGENT_NAME = "hyoka-trail-guide-agent";
const DATA_DIRECTORY = "data";
const DOCUMENT_NAME = "contoso-trail-guide.txt";

function requireEnvironmentVariable(name: "PROJECT_ENDPOINT" | "MODEL_DEPLOYMENT_NAME"): string {
  const value = process.env[name]?.trim();
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
  const errors: unknown[] = [];

  try {
    await mkdir(DATA_DIRECTORY, { recursive: true });
    const documentPath = join(DATA_DIRECTORY, DOCUMENT_NAME);
    await writeFile(documentPath, DOCUMENT_FACT, "utf8");

    const uploadedFile = await client.files.upload(
      createReadStream(documentPath),
      "assistants",
      { fileName: DOCUMENT_NAME },
    );
    uploadedFileId = uploadedFile.id;

    const vectorStore = await client.vectorStores.create({
      name: "hyoka-trail-guide-vector-store",
    });
    vectorStoreId = vectorStore.id;

    const vectorStoreFilePoller = client.vectorStoreFiles.createAndPoll(vectorStore.id, {
      fileId: uploadedFile.id,
      pollingOptions: { intervalInMs: 2_000 },
    });
    const vectorStoreFile = await vectorStoreFilePoller.pollUntilDone();
    if (vectorStoreFile.status !== "completed") {
      throw new Error(
        `Vector store file indexing did not complete successfully (status: ${vectorStoreFile.status}).`,
      );
    }

    const fileSearchTool = ToolUtility.createFileSearchTool([vectorStore.id]);
    const agent = await client.createAgent(modelDeploymentName, {
      name: AGENT_NAME,
      instructions: "Answer questions using facts from the uploaded trail guide.",
      tools: [fileSearchTool.definition],
      toolResources: fileSearchTool.resources,
    });
    agentId = agent.id;

    const thread = await client.threads.create();
    threadId = thread.id;
    await client.messages.create(thread.id, "user", USER_QUESTION);

    const run = await client.runs.createAndPoll(thread.id, agent.id, {
      pollingOptions: { intervalInMs: 2_000 },
    });
    if (run.status !== "completed") {
      const details = run.lastError
        ? ` ${run.lastError.code}: ${run.lastError.message}`
        : "";
      throw new Error(`Agent run did not complete successfully (status: ${run.status}).${details}`);
    }

    const messages = await client.messages.list(thread.id, { order: "asc" });
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
  } catch (error: unknown) {
    errors.push(error);
  } finally {
    const cleanupSteps: Array<[string, string | undefined, (id: string) => Promise<unknown>]> = [
      ["thread", threadId, (id) => client.threads.delete(id)],
      ["agent", agentId, (id) => client.deleteAgent(id)],
      ["vector store", vectorStoreId, (id) => client.vectorStores.delete(id)],
      ["uploaded file", uploadedFileId, (id) => client.files.delete(id)],
    ];

    for (const [resourceName, resourceId, remove] of cleanupSteps) {
      if (!resourceId) {
        continue;
      }
      try {
        await remove(resourceId);
      } catch (error: unknown) {
        errors.push(new Error(`Failed to delete ${resourceName} ${resourceId}.`, { cause: error }));
      }
    }
  }

  if (errors.length > 0) {
    throw new AggregateError(errors, "The application failed.");
  }
}

await main();
