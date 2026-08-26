package com.example;

import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.ai.projects.EvaluationsHelper;
import com.azure.ai.projects.models.TestingCriterionAzureAIEvaluator;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.openai.client.OpenAIClient;
import com.openai.core.JsonValue;
import com.openai.models.evals.EvalCreateParams;
import com.openai.models.evals.EvalCreateParams.DataSourceConfig.Custom;
import com.openai.models.evals.EvalCreateParams.DataSourceConfig.Custom.ItemSchema;
import com.openai.models.evals.EvalCreateResponse;
import com.openai.models.evals.EvalDeleteParams;
import com.openai.models.evals.runs.CreateEvalJsonlRunDataSource;
import com.openai.models.evals.runs.RunCreateParams;
import com.openai.models.evals.runs.RunCreateResponse;
import com.openai.models.evals.runs.RunRetrieveParams;
import com.openai.models.evals.runs.RunRetrieveResponse;
import com.openai.models.evals.runs.outputitems.OutputItemListParams;
import com.openai.models.evals.runs.outputitems.OutputItemListResponse;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public final class EvaluationRun {
    private EvaluationRun() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
        OpenAIClient openAIClient = new AIProjectClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildOpenAIClient();
        EvalCreateResponse evaluation = null;

        try {
            evaluation = openAIClient.evals().create(createEvaluation());
            System.out.println("Evaluation created: " + evaluation.id());

            RunCreateResponse run = openAIClient.evals().runs().create(
                RunCreateParams.builder()
                    .evalId(evaluation.id())
                    .name("hyoka-f1-run")
                    .dataSource(createRunDataSource())
                    .build());
            System.out.println("Evaluation run created: " + run.id());

            RunRetrieveResponse completedRun =
                waitForRun(openAIClient, evaluation.id(), run.id());
            if (!"completed".equals(completedRun.status())) {
                throw new IllegalStateException(
                    "Evaluation run ended with status " + completedRun.status());
            }
            System.out.println("Evaluation run completed: " + completedRun.id());
            printOutputItems(openAIClient, evaluation.id(), completedRun.id());
        } finally {
            if (evaluation != null) {
                openAIClient.evals().delete(EvalDeleteParams.builder()
                    .evalId(evaluation.id())
                    .build());
            }
        }
    }

    private static EvalCreateParams createEvaluation() {
        Map<String, Object> properties = new LinkedHashMap<>();
        properties.put("query", Collections.singletonMap("type", "string"));
        properties.put("response", Collections.singletonMap("type", "string"));
        properties.put(
            "ground_truth",
            Collections.singletonMap("type", "string"));

        ItemSchema itemSchema = ItemSchema.builder()
            .putAdditionalProperty("type", JsonValue.from("object"))
            .putAdditionalProperty("properties", JsonValue.from(properties))
            .putAdditionalProperty(
                "required",
                JsonValue.from(List.of("query", "response", "ground_truth")))
            .build();
        Custom dataSourceConfig = Custom.builder()
            .itemSchema(itemSchema)
            .includeSampleSchema(true)
            .build();

        TestingCriterionAzureAIEvaluator evaluator =
            new TestingCriterionAzureAIEvaluator("f1", "builtin.f1_score")
                .setDataMapping(Map.of(
                    "response", "{{item.response}}",
                    "ground_truth", "{{item.ground_truth}}"));

        return EvalCreateParams.builder()
            .name("hyoka-f1-evaluation")
            .dataSourceConfig(dataSourceConfig)
            .testingCriteria(List.of(
                EvaluationsHelper.toTestingCriterion(evaluator)))
            .build();
    }

    private static CreateEvalJsonlRunDataSource createRunDataSource() {
        CreateEvalJsonlRunDataSource.Source.FileContent.Content.Item item =
            CreateEvalJsonlRunDataSource.Source.FileContent.Content.Item.builder()
                .putAdditionalProperty(
                    "query",
                    JsonValue.from("What is the capital of France?"))
                .putAdditionalProperty("response", JsonValue.from("Paris"))
                .putAdditionalProperty("ground_truth", JsonValue.from("Paris"))
                .build();
        CreateEvalJsonlRunDataSource.Source.FileContent.Content content =
            CreateEvalJsonlRunDataSource.Source.FileContent.Content.builder()
                .item(item)
                .build();
        return CreateEvalJsonlRunDataSource.builder()
            .fileContentSource(List.of(content))
            .build();
    }

    private static RunRetrieveResponse waitForRun(
        OpenAIClient client,
        String evaluationId,
        String runId) throws InterruptedException {
        RunRetrieveResponse run = client.evals().runs().retrieve(
            RunRetrieveParams.builder()
                .evalId(evaluationId)
                .runId(runId)
                .build());
        while (!"completed".equals(run.status())
            && !"failed".equals(run.status())) {
            Thread.sleep(5000);
            run = client.evals().runs().retrieve(
                RunRetrieveParams.builder()
                    .evalId(evaluationId)
                    .runId(runId)
                    .build());
            System.out.println("Evaluation run status: " + run.status());
        }
        return run;
    }

    private static void printOutputItems(
        OpenAIClient client,
        String evaluationId,
        String runId) {
        for (OutputItemListResponse item : client.evals()
            .runs()
            .outputItems()
            .list(OutputItemListParams.builder()
                .evalId(evaluationId)
                .runId(runId)
                .build())
            .autoPager()) {
            System.out.printf(
                "Output item: id=%s status=%s%n",
                item.id(),
                item.status());
            for (OutputItemListResponse.Result result : item.results()) {
                System.out.printf(
                    "Metric: name=%s score=%s passed=%s%n",
                    result.name(),
                    result.score(),
                    result.passed());
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
