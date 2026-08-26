import os
import time

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import TestingCriterionAzureAIEvaluator
from azure.identity import DefaultAzureCredential
from openai.types.eval_create_params import DataSourceConfigCustom
from openai.types.evals.create_eval_jsonl_run_data_source_param import (
    CreateEvalJSONLRunDataSourceParam,
    SourceFileContent,
    SourceFileContentContent,
)


def main() -> None:
    endpoint = require_environment_variable("FOUNDRY_PROJECT_ENDPOINT")

    with (
        DefaultAzureCredential() as credential,
        AIProjectClient(endpoint, credential) as project_client,
        project_client.get_openai_client() as client,
    ):
        evaluation_id = None
        try:
            evaluation = client.evals.create(
                name="hyoka-f1-evaluation",
                data_source_config=DataSourceConfigCustom(
                    type="custom",
                    item_schema={
                        "type": "object",
                        "properties": {
                            "query": {"type": "string"},
                            "response": {"type": "string"},
                            "ground_truth": {"type": "string"},
                        },
                        "required": [
                            "query",
                            "response",
                            "ground_truth",
                        ],
                    },
                    include_sample_schema=True,
                ),
                testing_criteria=[
                    TestingCriterionAzureAIEvaluator(
                        type="azure_ai_evaluator",
                        name="f1",
                        evaluator_name="builtin.f1_score",
                        data_mapping={
                            "response": "{{item.response}}",
                            "ground_truth": "{{item.ground_truth}}",
                        },
                    )
                ],
            )
            evaluation_id = evaluation.id
            print(f"Evaluation created: {evaluation.id}")

            run = client.evals.runs.create(
                eval_id=evaluation.id,
                name="hyoka-f1-run",
                data_source=CreateEvalJSONLRunDataSourceParam(
                    type="jsonl",
                    source=SourceFileContent(
                        type="file_content",
                        content=[
                            SourceFileContentContent(
                                item={
                                    "query": (
                                        "What is the capital of France?"
                                    ),
                                    "response": "Paris",
                                    "ground_truth": "Paris",
                                }
                            )
                        ],
                    ),
                ),
            )
            print(f"Evaluation run created: {run.id}")

            while run.status not in ("completed", "failed"):
                time.sleep(5)
                run = client.evals.runs.retrieve(
                    run_id=run.id,
                    eval_id=evaluation.id,
                )
                print(f"Evaluation run status: {run.status}")
            if run.status != "completed":
                raise RuntimeError(
                    f"Evaluation run ended with status {run.status}."
                )

            print(f"Evaluation run completed: {run.id}")
            output_items = client.evals.runs.output_items.list(
                run_id=run.id,
                eval_id=evaluation.id,
            )
            for item in output_items:
                print(f"Output item: id={item.id} status={item.status}")
                for result in item.results:
                    print(
                        f"Metric: name={result.name} score={result.score} "
                        f"passed={result.passed}"
                    )
        finally:
            if evaluation_id is not None:
                client.evals.delete(eval_id=evaluation_id)


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
