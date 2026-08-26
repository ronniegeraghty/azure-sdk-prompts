using System.Text.Json;
using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string UserQuestion = "What is the weather in Seattle in celsius?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

FunctionToolDefinition weatherTool = new(
    name: "get_weather",
    description: "Gets the weather for a location in Celsius or Fahrenheit.",
    parameters: BinaryData.FromObjectAsJson(new
    {
        type = "object",
        properties = new
        {
            location = new
            {
                type = "string",
                description = "The city whose weather is requested."
            },
            unit = new
            {
                type = "string",
                @enum = new[] { "c", "f" },
                description = "Temperature unit: c for Celsius or f for Fahrenheit."
            }
        },
        required = new[] { "location", "unit" },
        additionalProperties = false
    }));

PersistentAgentsClient client = new(
    projectEndpoint,
    new DefaultAzureCredential());

PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: "hyoka-weather-agent",
        instructions:
            "Answer weather questions by calling get_weather. " +
            "Always use get_weather before giving a weather answer.",
        tools: [weatherTool]);

    thread = await client.Threads.CreateThreadAsync();

    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        UserQuestion);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);

    while (!IsTerminal(run.Status))
    {
        if (run.Status == RunStatus.RequiresAction)
        {
            if (run.RequiredAction is not SubmitToolOutputsAction submitAction)
            {
                throw new InvalidOperationException(
                    $"Run requires unsupported action '{run.RequiredAction?.GetType().Name ?? "unknown"}'.");
            }

            List<ToolOutput> outputs = [];

            foreach (RequiredToolCall toolCall in submitAction.ToolCalls)
            {
                if (toolCall is not RequiredFunctionToolCall functionCall)
                {
                    throw new InvalidOperationException(
                        $"Unsupported tool call type '{toolCall.GetType().Name}'.");
                }

                if (!string.Equals(functionCall.Name, "get_weather", StringComparison.Ordinal))
                {
                    throw new InvalidOperationException(
                        $"Unsupported function '{functionCall.Name}'.");
                }

                string result = await GetWeatherAsync(functionCall.Arguments);
                outputs.Add(new ToolOutput(toolCall.Id, result));
            }

            run = await client.Runs.SubmitToolOutputsToRunAsync(run, outputs);
        }
        else
        {
            await Task.Delay(TimeSpan.FromMilliseconds(500));
            run = await client.Runs.GetRunAsync(thread.Id, run.Id);
        }
    }

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run ended with status '{run.Status}'.");
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
            $"Required environment variable '{name}' is not set.");
    }

    return value;
}

static ValueTask<string> GetWeatherAsync(string arguments)
{
    WeatherArguments? request = JsonSerializer.Deserialize<WeatherArguments>(
        arguments,
        new JsonSerializerOptions
        {
            PropertyNameCaseInsensitive = true
        });

    if (request is null ||
        string.IsNullOrWhiteSpace(request.Location) ||
        request.Unit is not ("c" or "f"))
    {
        throw new JsonException(
            "get_weather requires a location and a unit of either 'c' or 'f'.");
    }

    if (!string.Equals(request.Location, "Seattle", StringComparison.OrdinalIgnoreCase))
    {
        throw new ArgumentException(
            $"No deterministic weather data is available for '{request.Location}'.",
            nameof(arguments));
    }

    int temperature = request.Unit == "c" ? 21 : 70;
    string result = JsonSerializer.Serialize(new
    {
        location = "Seattle",
        temperature,
        unit = request.Unit
    });

    return ValueTask.FromResult(result);
}

static bool IsTerminal(RunStatus status) =>
    status == RunStatus.Completed ||
    status == RunStatus.Failed ||
    status == RunStatus.Cancelled ||
    status == RunStatus.Expired;

internal sealed record WeatherArguments(string Location, string Unit);
