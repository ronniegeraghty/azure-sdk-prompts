using System.Text.Json;
using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string userMessage = "What is the weather in Seattle in celsius?";
string endpoint = RequireEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeployment = RequireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

FunctionToolDefinition weatherTool = new(
    name: "get_weather",
    description: "Gets the current weather for a location.",
    parameters: BinaryData.FromObjectAsJson(new
    {
        type = "object",
        properties = new
        {
            location = new { type = "string" },
            unit = new { type = "string", @enum = new[] { "c", "f" } },
        },
        required = new[] { "location", "unit" },
    }));

PersistentAgentsClient client = new(endpoint, new DefaultAzureCredential());
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeployment,
        name: "hyoka-weather-agent",
        instructions: "Use get_weather for every weather question.",
        tools: [weatherTool]);
    thread = await client.Threads.CreateThreadAsync();
    await client.Messages.CreateMessageAsync(thread.Id, MessageRole.User, userMessage);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);
    do
    {
        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);

        if (run.Status == RunStatus.RequiresAction
            && run.RequiredAction is SubmitToolOutputsAction action)
        {
            List<ToolOutput> outputs = [];
            foreach (RequiredToolCall call in action.ToolCalls)
            {
                if (call is not RequiredFunctionToolCall functionCall
                    || functionCall.Name != weatherTool.Name)
                {
                    throw new InvalidOperationException("Unexpected tool call.");
                }

                using JsonDocument arguments = JsonDocument.Parse(functionCall.Arguments);
                string location = arguments.RootElement.GetProperty("location").GetString()
                    ?? throw new InvalidOperationException("location is required.");
                string unit = arguments.RootElement.GetProperty("unit").GetString()
                    ?? throw new InvalidOperationException("unit is required.");
                outputs.Add(new ToolOutput(call, GetWeather(location, unit)));
            }

            run = await client.Runs.SubmitToolOutputsToRunAsync(
                run,
                outputs,
                toolApprovals: null);
        }
    }
    while (run.Status == RunStatus.Queued
        || run.Status == RunStatus.InProgress
        || run.Status == RunStatus.RequiresAction);

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
}

static string GetWeather(string location, string unit)
{
    if (!location.Contains("Seattle", StringComparison.OrdinalIgnoreCase)
        || unit is not ("c" or "f"))
    {
        throw new ArgumentException("Unsupported weather request.");
    }

    return JsonSerializer.Serialize(new
    {
        location = "Seattle",
        temperature = unit == "c" ? 21 : 70,
        unit,
    });
}

static string RequireEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException($"{name} is required.");
