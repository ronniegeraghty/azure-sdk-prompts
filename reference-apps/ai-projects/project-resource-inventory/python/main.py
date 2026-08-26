import os

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import Connection, ModelDeployment
from azure.identity import DefaultAzureCredential


def main() -> None:
    endpoint = require_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
    connection_name = require_environment_variable("CONNECTION_NAME")
    deployment_name = require_environment_variable("DEPLOYMENT_NAME")

    with AIProjectClient(endpoint, DefaultAzureCredential()) as client:
        print("Connections:")
        for connection in client.connections.list():
            print_connection(connection)

        print("Selected connection:")
        print_connection(client.connections.get(connection_name))

        print("Model deployments:")
        for deployment in client.deployments.list():
            if isinstance(deployment, ModelDeployment):
                print_deployment(deployment)

        print("Selected model deployment:")
        selected_deployment = client.deployments.get(deployment_name)
        if not isinstance(selected_deployment, ModelDeployment):
            raise RuntimeError(f"{deployment_name} is not a model deployment.")
        print_deployment(selected_deployment)


def print_connection(connection: Connection) -> None:
    print(
        f"{connection.name} | {connection.type} | {connection.target} "
        f"| default={connection.is_default}"
    )


def print_deployment(deployment: ModelDeployment) -> None:
    print(
        f"{deployment.name} | {deployment.model_publisher} "
        f"| {deployment.model_name} | {deployment.model_version}"
    )


def require_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise RuntimeError(f"{name} is required.")
    return value


if __name__ == "__main__":
    main()
