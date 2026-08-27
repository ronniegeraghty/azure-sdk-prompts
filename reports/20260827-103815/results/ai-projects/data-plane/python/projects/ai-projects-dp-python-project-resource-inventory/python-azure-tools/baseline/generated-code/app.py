"""Inspect connections and model deployments in a Microsoft Foundry project."""

from __future__ import annotations

import os
from typing import NoReturn

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import Connection, ModelDeployment
from azure.identity import DefaultAzureCredential


def required_environment_variable(name: str) -> str:
    """Return a required, non-empty environment variable."""
    value = os.environ.get(name, "").strip()
    if not value:
        raise ValueError(f"Environment variable {name} must be set to a non-empty value.")
    return value


def display_value(value: object) -> object:
    """Use an enum's wire value when the SDK returns an enum instance."""
    return getattr(value, "value", value)


def print_connection(connection: Connection) -> None:
    """Print the requested connection metadata."""
    print(f"  Name: {connection.name}")
    print(f"  Type: {display_value(connection.type)}")
    print(f"  Target: {connection.target}")
    print(f"  Default: {connection.is_default}")


def print_model_deployment(deployment: ModelDeployment) -> None:
    """Print the requested model deployment metadata."""
    print(f"  Name: {deployment.name}")
    print(f"  Model publisher: {deployment.model_publisher}")
    print(f"  Model name: {deployment.model_name}")
    print(f"  Model version: {deployment.model_version}")


def require_model_deployment(deployment: object, name: str) -> NoReturn:
    """Raise a useful error when a named deployment is not a model deployment."""
    deployment_type = display_value(getattr(deployment, "type", type(deployment).__name__))
    raise TypeError(
        f"Deployment {name!r} has type {deployment_type!r}; "
        "a model deployment is required."
    )


def main() -> None:
    endpoint = required_environment_variable("FOUNDRY_PROJECT_ENDPOINT")
    connection_name = required_environment_variable("CONNECTION_NAME")
    deployment_name = required_environment_variable("DEPLOYMENT_NAME")

    with (
        DefaultAzureCredential() as credential,
        AIProjectClient(endpoint=endpoint, credential=credential) as project_client,
    ):
        print("Project connections")
        for connection in project_client.connections.list():
            print_connection(connection)

        print(f"\nConnection {connection_name!r}")
        connection = project_client.connections.get(
            connection_name,
            include_credentials=False,
        )
        print_connection(connection)

        print("\nModel deployments")
        for deployment in project_client.deployments.list():
            if isinstance(deployment, ModelDeployment):
                print_model_deployment(deployment)

        print(f"\nModel deployment {deployment_name!r}")
        deployment = project_client.deployments.get(deployment_name)
        if not isinstance(deployment, ModelDeployment):
            require_model_deployment(deployment, deployment_name)
        print_model_deployment(deployment)


if __name__ == "__main__":
    main()
