using System.Text.Json;
using Azure;
using Azure.AI.Agents.Persistent;
using Azure.Identity;

const string question = "What is the weather in Seattle in celsius?";

string projectEndpoint = GetRequiredEnvironmentVariable("PROJECT_ENDPOINT");
string modelDeploymentName = GetRequiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

FunctionToolDefinition weatherTool = new(
    name: "get_weather",
    description: "Gets the current weather for a location in Celsius or Fahrenheit.",
    parameters: BinaryData.FromObjectAsJson(
        new
        {
            type = "object",
            properties = new
            {
                location = new
                {
                    type = "string",
                    description = "The location whose weather is requested.",
                },
                unit = new
                {
                    type = "string",
                    description = "The temperature unit.",
                    @enum = new[] { "c", "f" },
                },
            },
            required = new[] { "location", "unit" },
            additionalProperties = false,
        }));

PersistentAgentsClient client = new(projectEndpoint, new DefaultAzureCredential());
PersistentAgent? agent = null;
PersistentAgentThread? thread = null;

try
{
    agent = await client.Administration.CreateAgentAsync(
        model: modelDeploymentName,
        name: "hyoka-weather-agent",
        instructions:
            "You are a weather assistant. For every weather question, you must call the " +
            "get_weather function and use its result to answer the user.",
        tools: [weatherTool]);

    thread = await client.Threads.CreateThreadAsync();

    await client.Messages.CreateMessageAsync(
        thread.Id,
        MessageRole.User,
        question);

    ThreadRun run = await client.Runs.CreateRunAsync(thread.Id, agent.Id);

    while (run.Status == RunStatus.Queued ||
           run.Status == RunStatus.InProgress ||
           run.Status == RunStatus.RequiresAction)
    {
        if (run.Status == RunStatus.RequiresAction)
        {
            if (run.RequiredAction is not SubmitToolOutputsAction submitAction)
            {
                throw new InvalidOperationException(
                    $"The run requested unsupported action '{run.RequiredAction?.GetType().Name}'.");
            }

            List<ToolOutput> outputs = [];
            foreach (RequiredToolCall toolCall in submitAction.ToolCalls)
            {
                if (toolCall is not RequiredFunctionToolCall functionCall)
                {
                    throw new InvalidOperationException(
                        $"Tool call '{toolCall.Id}' is not a function call.");
                }

                if (!string.Equals(functionCall.Name, weatherTool.Name, StringComparison.Ordinal))
                {
                    throw new InvalidOperationException(
                        $"Unknown function '{functionCall.Name}' in tool call '{toolCall.Id}'.");
                }

                using JsonDocument arguments = JsonDocument.Parse(functionCall.Arguments);
                string location = GetRequiredString(arguments.RootElement, "location");
                string unit = GetRequiredString(arguments.RootElement, "unit");
                string result = GetWeather(location, unit);

                outputs.Add(new ToolOutput(toolCall, result));
            }

            run = await client.Runs.SubmitToolOutputsToRunAsync(run, outputs);
            continue;
        }

        await Task.Delay(TimeSpan.FromMilliseconds(500));
        run = await client.Runs.GetRunAsync(thread.Id, run.Id);
    }

    if (run.Status != RunStatus.Completed)
    {
        throw new InvalidOperationException(
            $"Agent run ended with status '{run.Status}'. Error: {run.LastError?.Message ?? "none"}");
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
            if (content is MessageTextContent textContent)
            {
                Console.WriteLine(textContent.Text);
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

static string GetRequiredEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name) is { Length: > 0 } value
        ? value
        : throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");

static string GetRequiredString(JsonElement arguments, string propertyName)
{
    if (!arguments.TryGetProperty(propertyName, out JsonElement property) ||
        property.ValueKind != JsonValueKind.String ||
        string.IsNullOrWhiteSpace(property.GetString()))
    {
        throw new JsonException(
            $"Function argument '{propertyName}' must be a non-empty string.");
    }

    return property.GetString()!;
}

static string GetWeather(string location, string unit)
{
    if (!string.Equals(location, "Seattle", StringComparison.OrdinalIgnoreCase))
    {
        throw new ArgumentOutOfRangeException(
            nameof(location), location, "Only Seattle is supported.");
    }

    int temperature = unit.ToLowerInvariant() switch
    {
        "c" => 21,
        "f" => 70,
        _ => throw new ArgumentOutOfRangeException(
            nameof(unit), unit, "Unit must be 'c' or 'f'."),
    };

    return JsonSerializer.Serialize(new
    {
        location,
        unit = unit.ToLowerInvariant(),
        temperature,
    });
}
