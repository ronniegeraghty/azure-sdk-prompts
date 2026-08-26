using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string guide =
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const string question =
    "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";

string endpoint = RequireEnvironmentVariable("PROJECT_ENDPOINT");
string model = RequireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
string documentPath = Path.Combine(Path.GetTempPath(), $"hyoka-trail-guide-{Guid.NewGuid():N}.txt");
await File.WriteAllTextAsync(documentPath, guide);

PersistentAgentsClient client = new(endpoint, new DefaultAzureCredential());
PersistentAgentFileInfo? uploadedFile = null;
PersistentAgentsVectorStore? vectorStore = null;
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    uploadedFile = await client.Files.UploadFileAsync(
        documentPath,
        PersistentAgentFilePurpose.Agents);
    vectorStore = await client.VectorStores.CreateVectorStoreAsync(
        fileIds: [uploadedFile.Id],
        name: "hyoka-trail-guide");

    while (vectorStore.Status == VectorStoreStatus.InProgress)
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        vectorStore = await client.VectorStores.GetVectorStoreAsync(vectorStore.Id);
    }
    if (vectorStore.Status != VectorStoreStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Vector store ended with status {vectorStore.Status}.");
    }

    FileSearchToolResource searchResource = new();
    searchResource.VectorStoreIds.Add(vectorStore.Id);
    agent = await client.Administration.CreateAgentAsync(
        model: model,
        name: "hyoka-trail-guide-agent",
        instructions: "Use file search to answer questions about the uploaded guide.",
        tools: [new FileSearchToolDefinition()],
        toolResources: new ToolResources { FileSearch = searchResource });

    thread = await client.Threads.CreateThreadAsync();
    await client.Messages.CreateMessageAsync(thread.Id, MessageRole.User, question);
    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);
    do
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);
    }
    while (run.Status == RunStatus.Queued || run.Status == RunStatus.InProgress);

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException($"Run ended with status {run.Status}.");
    }

    AsyncPageable<PersistentThreadMessage> messages = client.Messages.GetMessagesAsync(
        thread.Id,
        order: ListSortOrder.Ascending);
    await foreach (PersistentThreadMessage message in messages)
    {
        if (message.Role != MessageRole.Agent)
        {
            continue;
        }
        foreach (MessageContent content in message.ContentItems)
        {
            if (content is MessageTextContent text)
            {
                Console.WriteLine(text.Text);
            }
        }
    }
}
finally
{
    if (thread is not null)
    {
        await client.Threads.DeleteThreadAsync(thread.Id);
    }
    if (agent is not null)
    {
        await client.Administration.DeleteAgentAsync(agent.Id);
    }
    if (vectorStore is not null)
    {
        await client.VectorStores.DeleteVectorStoreAsync(vectorStore.Id);
    }
    if (uploadedFile is not null)
    {
        await client.Files.DeleteFileAsync(uploadedFile.Id);
    }
    File.Delete(documentPath);
}

static string RequireEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException($"{name} is required.");
