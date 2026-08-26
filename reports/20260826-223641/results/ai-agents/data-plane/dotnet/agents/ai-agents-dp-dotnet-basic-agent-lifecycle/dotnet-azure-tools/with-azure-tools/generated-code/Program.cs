using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string AgentName = "hyoka-basic-agent";
const string AgentInstructions = "Answer the user's question clearly and concisely.";
const string UserMessage = "What is the capital of France?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

PersistentAgentsClient client = new(
    projectEndpoint,
    new DefaultAzureCredential());

PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: AgentName,
        instructions: AgentInstructions);

    thread = await client.Threads.CreateThreadAsync();

    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        UserMessage);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);

    while (run.Status == RunStatus.Queued || run.Status == RunStatus.InProgress)
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);
    }

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run '{run.Id}' ended with status '{run.Status}'. " +
            $"Error: {run.LastError?.Message ?? "No error details were provided."}");
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
            $"Environment variable '{name}' is required.");
    }

    return value;
}
