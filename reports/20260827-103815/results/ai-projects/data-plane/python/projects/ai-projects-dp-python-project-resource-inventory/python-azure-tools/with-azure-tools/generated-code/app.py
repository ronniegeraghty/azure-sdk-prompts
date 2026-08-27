import os
import sys

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import Connection, ModelDeployment
from azure.core.exceptions import AzureError
from azure.identity import DefaultAzureCredential


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name, "").strip()
    if not value:
        raise ValueError(f"Environment variable {name} must be set.")
    return value


def display_value(value: object) -> str:
    return str(getattr(value, "value", value))


def print_connection(connection: Connection) -> None:
    print(f"  Name: {connection.name}")
    print(f"  Type: {display_value(connection.type)}")
    print(f"  Target: {connection.target}")
    print(f"  Default: {connection.is_default}")


def print_model_deployment(deployment: ModelDeployment) -> None:
    print(f"  Name: {deployment.name}")
    print(f"  Model publisher: {deployment.model_publisher}")
    print(f"  Model name: {deployment.model_name}")
    print(f"  Model version: {deployment.model_version}")


def inspect_project(
    project_client: AIProjectClient,
    connection_name: str,
    deployment_name: str,
) -> None:
    print("Project connections")
    connection_count = 0
    for connection in project_client.connections.list():
        connection_count += 1
        print_connection(connection)
        print()
    if connection_count == 0:
        print("  No connections found.")
        print()

    print(f"Requested connection: {connection_name}")
    connection = project_client.connections.get(
        connection_name,
        include_credentials=False,
    )
    print_connection(connection)
    print()

    print("Project model deployments")
    model_deployment_count = 0
    for deployment in project_client.deployments.list():
        if isinstance(deployment, ModelDeployment):
            model_deployment_count += 1
            print_model_deployment(deployment)
            print()
    if model_deployment_count == 0:
        print("  No model deployments found.")
        print()

    print(f"Requested deployment: {deployment_name}")
    deployment = project_client.deployments.get(deployment_name)
    if not isinstance(deployment, ModelDeployment):
        raise TypeError(
            f"Deployment {deployment_name!r} is not a model deployment "
            f"(received {type(deployment).__name__})."
        )
    print_model_deployment(deployment)


def main() -> int:
    try:
        endpoint = required_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
        connection_name = required_environment_variable("CONNECTION_NAME")
        deployment_name = required_environment_variable("DEPLOYMENT_NAME")

        with (
            DefaultAzureCredential() as credential,
            AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
        ):
            inspect_project(project_client, connection_name, deployment_name)
    except (AzureError, TypeError, ValueError) as error:
        print(f"Error: {error}", file=sys.stderr)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
