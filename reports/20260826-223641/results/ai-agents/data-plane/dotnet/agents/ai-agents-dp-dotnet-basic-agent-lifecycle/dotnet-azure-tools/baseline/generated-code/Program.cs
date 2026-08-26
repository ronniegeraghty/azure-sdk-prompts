using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string agentName = "hyoka-basic-agent";
const string agentInstructions = "Answer the user's question clearly and concisely.";
const string userMessage = "What is the capital of France?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

PersistentAgentsClient client = new(projectEndpoint, new DefaultAzureCredential());
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: agentName,
        instructions: agentInstructions);

    thread = await client.Threads.CreateThreadAsync();

    await client.Messages.CreateMessageAsync(
        threadId: thread.Id,
        role: MessageRole.User,
        content: userMessage);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);

    while (!IsTerminal(run.Status))
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);
    }

    if (run.Status != RunStatus.Completed)
    {
        string detail = run.LastError?.Message ?? "No error details were provided.";
        throw new InvalidOperationException(
            $"Agent run ended with status '{run.Status}': {detail}");
    }

    AsyncPageable<PersistentThreadMessage> messages =
        client.Messages.GetMessagesAsync(
            threadId: thread.Id,
            order: ListSortOrder.Ascending);

    await foreach (PersistentThreadMessage message in messages)
    {
        if (message.Role != MessageRole.Agent)
        {
            continue;
        }

        foreach (MessageContent contentItem in message.ContentItems)
        {
            if (contentItem is MessageTextContent textContent)
            {
                Console.WriteLine(textContent.Text);
            }
        }
    }
}
finally
{
    try
    {
        if (thread is not null)
        {
            await client.Threads.DeleteThreadAsync(thread.Id);
        }
    }
    finally
    {
        if (agent is not null)
        {
            await client.Administration.DeleteAgentAsync(agent.Id);
        }
    }
}

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Environment variable '{name}' must be set.");
    }

    return value;
}

static bool IsTerminal(RunStatus status) =>
    status == RunStatus.Completed
    || status == RunStatus.Failed
    || status == RunStatus.Cancelled
    || status == RunStatus.Expired
    || status == RunStatus.Incomplete;
