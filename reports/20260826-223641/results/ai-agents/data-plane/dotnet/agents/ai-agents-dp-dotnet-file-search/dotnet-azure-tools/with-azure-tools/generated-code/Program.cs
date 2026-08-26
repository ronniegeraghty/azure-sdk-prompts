using System.Runtime.ExceptionServices;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string GuideFact =
    "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
const string Question =
    "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
string documentPath = Path.Combine(Path.GetTempPath(), $"contoso-trail-guide-{Guid.NewGuid():N}.txt");

PersistentAgentsClient client = new(
    projectEndpoint,
    new DefaultAzureCredential());

string? uploadedFileId = null;
string? vectorStoreId = null;
string? agentId = null;
string? threadId = null;
Exception? operationException = null;
List<Exception> cleanupExceptions = [];

try
{
    await File.WriteAllTextAsync(documentPath, GuideFact);

    PersistentAgentFileInfo uploadedFile = await client.Files.UploadFileAsync(
        documentPath,
        PersistentAgentFilePurpose.Agents);
    uploadedFileId = uploadedFile.Id;

    PersistentAgentsVectorStore vectorStore = await client.VectorStores.CreateVectorStoreAsync(
        fileIds: [uploadedFile.Id],
        name: "hyoka-trail-guide-vector-store");
    vectorStoreId = vectorStore.Id;

    vectorStore = await WaitForVectorStoreAsync(client, vectorStore.Id);

    FileSearchToolResource fileSearchResource = new();
    fileSearchResource.VectorStoreIds.Add(vectorStore.Id);

    PersistentAgent agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: "hyoka-trail-guide-agent",
        instructions: "Answer questions using the uploaded guide. Use file search and do not invent facts.",
        tools: [new FileSearchToolDefinition()],
        toolResources: new ToolResources { FileSearch = fileSearchResource });
    agentId = agent.Id;

    PersistentAgentThread thread = await client.Threads.CreateThreadAsync();
    threadId = thread.Id;

    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        Question);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);
    run = await WaitForRunAsync(client, thread.Id, run);

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run ended with status '{run.Status}': {run.LastError?.Message ?? "No error details were provided."}");
    }

    await foreach (PersistentThreadMessage message in client.Messages.GetMessagesAsync(
        threadId: thread.Id,
        order: ListSortOrder.Ascending))
    {
        if (message.Role != MessageRole.Agent)
        {
            continue;
        }

        foreach (MessageContent content in message.ContentItems)
        {
            if (content is MessageTextContent textContent)
            {
                Console.WriteLine(textContent.Text);
            }
        }
    }
}
catch (Exception ex)
{
    operationException = ex;
}
finally
{
    if (threadId is not null)
    {
        await TryCleanupAsync(
            () => client.Threads.DeleteThreadAsync(threadId),
            $"thread '{threadId}'",
            cleanupExceptions);
    }

    if (agentId is not null)
    {
        await TryCleanupAsync(
            () => client.Administration.DeleteAgentAsync(agentId),
            $"agent '{agentId}'",
            cleanupExceptions);
    }

    if (vectorStoreId is not null)
    {
        await TryCleanupAsync(
            () => client.VectorStores.DeleteVectorStoreAsync(vectorStoreId),
            $"vector store '{vectorStoreId}'",
            cleanupExceptions);
    }

    if (uploadedFileId is not null)
    {
        await TryCleanupAsync(
            () => client.Files.DeleteFileAsync(uploadedFileId),
            $"uploaded file '{uploadedFileId}'",
            cleanupExceptions);
    }

    File.Delete(documentPath);
}

if (operationException is not null)
{
    if (cleanupExceptions.Count > 0)
    {
        Console.Error.WriteLine(
            $"Cleanup also failed: {string.Join(" | ", cleanupExceptions.Select(ex => ex.Message))}");
    }

    ExceptionDispatchInfo.Capture(operationException).Throw();
}

if (cleanupExceptions.Count > 0)
{
    throw new AggregateException("One or more resources could not be deleted.", cleanupExceptions);
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);
    return string.IsNullOrWhiteSpace(value)
        ? throw new InvalidOperationException($"Environment variable '{name}' is required.")
        : value;
}

static async Task<PersistentAgentsVectorStore> WaitForVectorStoreAsync(
    PersistentAgentsClient client,
    string vectorStoreId)
{
    using CancellationTokenSource timeout = new(TimeSpan.FromMinutes(5));

    while (true)
    {
        PersistentAgentsVectorStore vectorStore =
            await client.VectorStores.GetVectorStoreAsync(vectorStoreId, timeout.Token);

        if (vectorStore.FileCounts.Failed > 0 || vectorStore.FileCounts.Cancelled > 0)
        {
            throw new InvalidOperationException(
                $"Vector-store indexing failed: {vectorStore.FileCounts.Failed} failed and " +
                $"{vectorStore.FileCounts.Cancelled} cancelled file(s).");
        }

        if (vectorStore.Status == VectorStoreStatus.Completed)
        {
            if (vectorStore.FileCounts.Total == 0 ||
                vectorStore.FileCounts.Completed != vectorStore.FileCounts.Total)
            {
                throw new InvalidOperationException(
                    "Vector store completed without successfully indexing every file.");
            }

            return vectorStore;
        }

        if (vectorStore.Status != VectorStoreStatus.InProgress)
        {
            throw new InvalidOperationException(
                $"Vector store ended with unexpected status '{vectorStore.Status}'.");
        }

        await Task.Delay(TimeSpan.FromMilliseconds(500), timeout.Token);
    }
}

static async Task<ThreadRun> WaitForRunAsync(
    PersistentAgentsClient client,
    string threadId,
    ThreadRun run)
{
    using CancellationTokenSource timeout = new(TimeSpan.FromMinutes(5));

    while (run.Status == RunStatus.Queued ||
           run.Status == RunStatus.InProgress ||
           run.Status == RunStatus.Cancelling)
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500), timeout.Token);
        run = await client.Runs.GetRunAsync(threadId, run.Id, timeout.Token);
    }

    return run;
}

static async Task TryCleanupAsync(
    Func<Task> cleanup,
    string resourceDescription,
    ICollection<Exception> cleanupExceptions)
{
    try
    {
        await cleanup();
    }
    catch (Exception ex)
    {
        cleanupExceptions.Add(
            new InvalidOperationException($"Failed to delete {resourceDescription}.", ex));
    }
}
