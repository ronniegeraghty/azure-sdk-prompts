using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string agentName = "hyoka-basic-agent";
const string agentInstructions = "Answer the user's question clearly and concisely.";
const string userMessage = "What is the capital of France?";

string endpoint = Environment.GetEnvironmentVariable("PROJECT_ENDPOINT")
    ?? throw new InvalidOperationException("PROJECT_ENDPOINT is required.");
string modelDeployment = Environment.GetEnvironmentVariable("MODEL_DEPLOYMENT_NAME")
    ?? throw new InvalidOperationException("MODEL_DEPLOYMENT_NAME is required.");

PersistentAgentsClient client = new(endpoint, new DefaultAzureCredential());
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeployment,
        name: agentName,
        instructions: agentInstructions);

    thread = await client.Threads.CreateThreadAsync();
    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        userMessage);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);
    while (run.Status == RunStatus.Queued || run.Status == RunStatus.InProgress)
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);
    }

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run ended with status {run.Status}: {run.LastError?.Message}");
    }

    AsyncPageable<PersistentThreadMessage> messages = client.Messages.GetMessagesAsync(
        threadId: thread.Id,
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
}
