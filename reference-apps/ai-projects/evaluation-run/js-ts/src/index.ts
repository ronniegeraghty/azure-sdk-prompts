import { AIProjectClient } from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";

const endpoint = requiredEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
const project = new AIProjectClient(endpoint, new DefaultAzureCredential());
const openAIClient = project.getOpenAIClient();
let evaluationId: string | undefined;

try {
  const testingCriteria = [
    {
      type: "azure_ai_evaluator" as const,
      name: "f1",
      evaluator_name: "builtin.f1_score",
      data_mapping: {
        response: "{{item.response}}",
        ground_truth: "{{item.ground_truth}}",
      },
    },
  ];

  const evaluation = await openAIClient.evals.create({
    name: "hyoka-f1-evaluation",
    data_source_config: {
      type: "custom",
      item_schema: {
        type: "object",
        properties: {
          query: { type: "string" },
          response: { type: "string" },
          ground_truth: { type: "string" },
        },
        required: ["query", "response", "ground_truth"],
      },
      include_sample_schema: true,
    },
    testing_criteria: testingCriteria as any,
  });
  evaluationId = evaluation.id;
  console.log(`Evaluation created: ${evaluation.id}`);

  let run = await openAIClient.evals.runs.create(evaluation.id, {
    name: "hyoka-f1-run",
    data_source: {
      type: "jsonl",
      source: {
        type: "file_content",
        content: [
          {
            item: {
              query: "What is the capital of France?",
              response: "Paris",
              ground_truth: "Paris",
            },
          },
        ],
      },
    },
  });
  console.log(`Evaluation run created: ${run.id}`);

  while (!["completed", "failed"].includes(run.status)) {
    await delay(5000);
    run = await openAIClient.evals.runs.retrieve(run.id, {
      eval_id: evaluation.id,
    });
    console.log(`Evaluation run status: ${run.status}`);
  }
  if (run.status !== "completed") {
    throw new Error(`Evaluation run ended with status ${run.status}.`);
  }

  console.log(`Evaluation run completed: ${run.id}`);
  for await (const item of openAIClient.evals.runs.outputItems.list(run.id, {
    eval_id: evaluation.id,
  })) {
    console.log(`Output item: id=${item.id} status=${item.status}`);
    for (const result of item.results) {
      console.log(
        `Metric: name=${result.name} score=${result.score} passed=${result.passed}`,
      );
    }
  }
} finally {
  if (evaluationId !== undefined) {
    await openAIClient.evals.delete(evaluationId);
  }
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}

function delay(milliseconds: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}
