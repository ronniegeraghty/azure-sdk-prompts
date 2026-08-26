import os

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import Connection, ModelDeployment
from azure.identity import DefaultAzureCredential


def print_connection(connection: Connection) -> None:
    print(f"  Name: {connection.name}")
    print(f"  Type: {connection.type}")
    print(f"  Target: {connection.target}")
    print(f"  Default: {connection.is_default}")


def print_model_deployment(deployment: ModelDeployment) -> None:
    print(f"  Name: {deployment.name}")
    print(f"  Model publisher: {deployment.model_publisher}")
    print(f"  Model name: {deployment.model_name}")
    print(f"  Model version: {deployment.model_version}")


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    endpoint = required_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
    connection_name = required_environment_variable("CONNECTION_NAME")
    deployment_name = required_environment_variable("DEPLOYMENT_NAME")

    with (
        DefaultAzureCredential() as credential,
        AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
    ):
        print("Project connections")
        print("===================")
        for connection in project_client.connections.list():
            print_connection(connection)
            print()

        print(f"Connection: {connection_name}")
        print("=" * (12 + len(connection_name)))
        connection = project_client.connections.get(
            connection_name,
            include_credentials=False,
        )
        print_connection(connection)
        print()

        print("Model deployments")
        print("=================")
        for deployment in project_client.deployments.list():
            if isinstance(deployment, ModelDeployment):
                print_model_deployment(deployment)
                print()

        print(f"Deployment: {deployment_name}")
        print("=" * (12 + len(deployment_name)))
        deployment = project_client.deployments.get(deployment_name)
        if not isinstance(deployment, ModelDeployment):
            raise TypeError(
                f"Deployment {deployment_name!r} is not a model deployment "
                f"(received {type(deployment).__name__})."
            )
        print_model_deployment(deployment)


if __name__ == "__main__":
    main()
