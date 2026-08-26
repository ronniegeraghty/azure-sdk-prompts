using System.ClientModel;
using System.Text.Json;
using Azure.AI.Projects;
using Azure.Identity;
using OpenAI.Evals;

string endpoint = RequireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
AIProjectClient projectClient = new(new Uri(endpoint), new DefaultAzureCredential());
EvaluationClient evaluationClient =
    projectClient.ProjectOpenAIClient.GetEvaluationClient();
string? evaluationId = null;

try
{
    object dataSourceConfig = new
    {
        type = "custom",
        item_schema = new
        {
            type = "object",
            properties = new
            {
                query = new { type = "string" },
                response = new { type = "string" },
                ground_truth = new { type = "string" },
            },
            required = new[] { "query", "response", "ground_truth" },
        },
        include_sample_schema = true,
    };
    object[] testingCriteria =
    [
        new
        {
            type = "azure_ai_evaluator",
            name = "f1",
            evaluator_name = "builtin.f1_score",
            data_mapping = new
            {
                response = "{{item.response}}",
                ground_truth = "{{item.ground_truth}}",
            },
        },
    ];
    BinaryData evaluationData = BinaryData.FromObjectAsJson(new
    {
        name = "hyoka-f1-evaluation",
        data_source_config = dataSourceConfig,
        testing_criteria = testingCriteria,
    });

    using (BinaryContent content = BinaryContent.Create(evaluationData))
    {
        ClientResult evaluation =
            await evaluationClient.CreateEvaluationAsync(content);
        evaluationId = GetString(evaluation, "id");
    }
    Console.WriteLine($"Evaluation created: {evaluationId}");

    BinaryData runData = BinaryData.FromObjectAsJson(new
    {
        name = "hyoka-f1-run",
        data_source = new
        {
            type = "jsonl",
            source = new
            {
                type = "file_content",
                content = new[]
                {
                    new
                    {
                        item = new
                        {
                            query = "What is the capital of France?",
                            response = "Paris",
                            ground_truth = "Paris",
                        },
                    },
                },
            },
        },
    });

    ClientResult run;
    using (BinaryContent content = BinaryContent.Create(runData))
    {
        run = await evaluationClient.CreateEvaluationRunAsync(
            evaluationId,
            content);
    }

    string runId = GetString(run, "id");
    string status = GetString(run, "status");
    while (status is not ("completed" or "failed"))
    {
        await Task.Delay(TimeSpan.FromSeconds(5));
        run = await evaluationClient.GetEvaluationRunAsync(
            evaluationId,
            runId,
            options: new());
        status = GetString(run, "status");
        Console.WriteLine($"Evaluation run status: {status}");
    }
    if (status != "completed")
    {
        throw new InvalidOperationException(
            $"Evaluation run ended with status {status}.");
    }

    Console.WriteLine($"Evaluation run completed: {runId}");
    await PrintOutputItemsAsync(evaluationClient, evaluationId, runId);
}
finally
{
    if (evaluationId is not null)
    {
        await evaluationClient.DeleteEvaluationAsync(
            evaluationId,
            new System.ClientModel.Primitives.RequestOptions());
    }
}

static async Task PrintOutputItemsAsync(
    EvaluationClient client,
    string evaluationId,
    string runId)
{
    bool hasMore;
    string? after = null;
    do
    {
        ClientResult page = await client.GetEvaluationRunOutputItemsAsync(
            evaluationId: evaluationId,
            evaluationRunId: runId,
            limit: null,
            order: "asc",
            after: after,
            outputItemStatus: default,
            options: new());
        using JsonDocument document = Parse(page);
        JsonElement root = document.RootElement;

        foreach (JsonElement item in root.GetProperty("data").EnumerateArray())
        {
            Console.WriteLine(
                $"Output item: id={item.GetProperty("id").GetString()} "
                + $"status={item.GetProperty("status").GetString()}");
            foreach (JsonElement result in item.GetProperty("results").EnumerateArray())
            {
                Console.WriteLine(
                    $"Metric: name={result.GetProperty("name").GetString()} "
                    + $"score={result.GetProperty("score").GetDouble()} "
                    + $"passed={result.GetProperty("passed").GetBoolean()}");
            }
        }

        hasMore = root.GetProperty("has_more").GetBoolean();
        after = hasMore ? root.GetProperty("last_id").GetString() : null;
    }
    while (hasMore);
}

static string GetString(ClientResult result, string property)
{
    using JsonDocument document = Parse(result);
    return document.RootElement.GetProperty(property).GetString()
        ?? throw new InvalidOperationException($"{property} is missing.");
}

static JsonDocument Parse(ClientResult result) =>
    JsonDocument.Parse(result.GetRawResponse().Content.ToMemory());

static string RequireEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException($"{name} is required.");
