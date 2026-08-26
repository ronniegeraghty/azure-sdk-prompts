import type { MessageTextContent } from "@azure/ai-agents";
import { AgentsClient, ToolUtility, isOutputOfType } from "@azure/ai-agents";
import { DefaultAzureCredential } from "@azure/identity";
import { Readable } from "node:stream";

const guide =
  "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const question =
  "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";
const endpoint = requiredEnvironmentVariable("PROJECT_ENDPOINT");
const model = requiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
const client = new AgentsClient(endpoint, new DefaultAzureCredential());

let fileId: string | undefined;
let vectorStoreId: string | undefined;
let agentId: string | undefined;
let threadId: string | undefined;

try {
  const file = await client.files.upload(Readable.from([guide]), "assistants", {
    fileName: "trail-guide.txt",
  });
  fileId = file.id;

  const vectorStore = await client.vectorStores.create({
    name: "hyoka-trail-guide",
  });
  vectorStoreId = vectorStore.id;
  const indexedFile = await client.vectorStoreFiles.create(vectorStore.id, {
    fileId: file.id,
    pollingOptions: { intervalInMs: 500 },
  });
  if (indexedFile.status !== "completed") {
    throw new Error(`Vector-store file ended with status ${indexedFile.status}.`);
  }

  const searchTool = ToolUtility.createFileSearchTool([vectorStore.id]);
  const agent = await client.createAgent(model, {
    name: "hyoka-trail-guide-agent",
    instructions: "Use file search to answer questions about the uploaded guide.",
    tools: [searchTool.definition],
    toolResources: searchTool.resources,
  });
  agentId = agent.id;

  const thread = await client.threads.create();
  threadId = thread.id;
  await client.messages.create(thread.id, "user", question);
  const run = await client.runs.createAndPoll(thread.id, agent.id, {
    pollingOptions: { intervalInMs: 500 },
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
      if (isOutputOfType<MessageTextContent>(content, "text")) {
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
  if (vectorStoreId !== undefined) {
    await client.vectorStores.delete(vectorStoreId);
  }
  if (fileId !== undefined) {
    await client.files.delete(fileId);
  }
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}
