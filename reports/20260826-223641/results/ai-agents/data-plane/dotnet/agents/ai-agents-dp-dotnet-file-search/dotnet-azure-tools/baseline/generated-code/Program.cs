using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string GuideFact =
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const string Question =
    "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
string documentPath = Path.Combine(Path.GetTempPath(), $"contoso-trail-guide-{Guid.NewGuid():N}.txt");

PersistentAgentsClient client = new(projectEndpoint, new DefaultAzureCredential());
PersistentAgentFileInfo? uploadedFile = null;
PersistentAgentsVectorStore? vectorStore = null;
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

using CancellationTokenSource timeout = new(TimeSpan.FromMinutes(10));
CancellationToken cancellationToken = timeout.Token;

try
{
    await File.WriteAllTextAsync(documentPath, GuideFact, cancellationToken);

    uploadedFile = await client.Files.UploadFileAsync(
        documentPath,
        PersistentAgentFilePurpose.Agents,
        cancellationToken);

    vectorStore = await client.VectorStores.CreateVectorStoreAsync(
        fileIds: [uploadedFile.Id],
        name: "hyoka-trail-guide-vector-store",
        cancellationToken: cancellationToken);

    while (vectorStore.Status == VectorStoreStatus.InProgress)
    {
        await Task.Delay(TimeSpan.FromSeconds(1), cancellationToken);
        vectorStore = await client.VectorStores.GetVectorStoreAsync(
            vectorStore.Id,
            cancellationToken);
    }

    if (vectorStore.Status != VectorStoreStatus.Completed ||
        vectorStore.FileCounts.Completed != 1 ||
        vectorStore.FileCounts.Failed != 0 ||
        vectorStore.FileCounts.Cancelled != 0)
    {
        throw new InvalidOperationException(
            $"Vector store indexing did not complete successfully. " +
            $"Status: {vectorStore.Status}; completed: {vectorStore.FileCounts.Completed}; " +
            $"failed: {vectorStore.FileCounts.Failed}; cancelled: {vectorStore.FileCounts.Cancelled}.");
    }

    FileSearchToolResource fileSearchResource = new();
    fileSearchResource.VectorStoreIds.Add(vectorStore.Id);

    agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: "hyoka-trail-guide-agent",
        instructions: "Answer questions using the uploaded trail guide. Do not invent facts.",
        tools: [new FileSearchToolDefinition()],
        toolResources: new ToolResources { FileSearch = fileSearchResource },
        cancellationToken: cancellationToken);

    thread = await client.Threads.CreateThreadAsync(cancellationToken: cancellationToken);

    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        Question,
        cancellationToken: cancellationToken);

    ThreadRun run = await client.Runs.CreateRunAsync(
        thread,
        agent,
        cancellationToken: cancellationToken);

    while (run.Status == RunStatus.Queued ||
           run.Status == RunStatus.InProgress ||
           run.Status == RunStatus.Cancelling)
    {
        await Task.Delay(TimeSpan.FromSeconds(1), cancellationToken);
        run = await client.Runs.GetRunAsync(thread.Id, run.Id, cancellationToken);
    }

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run ended with status '{run.Status}': {run.LastError?.Message ?? "No error details were returned."}");
    }

    await foreach (PersistentThreadMessage message in client.Messages.GetMessagesAsync(
        threadId: thread.Id,
        order: ListSortOrder.Ascending,
        cancellationToken: cancellationToken))
    {
        if (message.Role == MessageRole.Agent)
        {
            foreach (MessageTextContent text in message.ContentItems.OfType<MessageTextContent>())
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

    if (File.Exists(documentPath))
    {
        File.Delete(documentPath);
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException($"Environment variable '{name}' is required.")
        : value;
}
