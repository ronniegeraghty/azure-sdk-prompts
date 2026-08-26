"""Inspect connections and model deployments in a Microsoft Foundry project."""

from __future__ import annotations

import os
import sys
from dataclasses import dataclass

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import Connection, ModelDeployment
from azure.core.exceptions import AzureError
from azure.identity import DefaultAzureCredential


@dataclass(frozen=True)
class Settings:
    project_endpoint: str
    connection_name: str
    deployment_name: str

    @classmethod
    def from_environment(cls) -> Settings:
        variable_names = (
            "FOUNDRY_PROJECT_ENDPOINT",
            "CONNECTION_NAME",
            "DEPLOYMENT_NAME",
        )
        values = {name: os.environ.get(name, "").strip() for name in variable_names}
        missing = [name for name, value in values.items() if not value]
        if missing:
            raise ValueError(
                "Missing required environment variable(s): " + ", ".join(missing)
            )

        return cls(
            project_endpoint=values["FOUNDRY_PROJECT_ENDPOINT"],
            connection_name=values["CONNECTION_NAME"],
            deployment_name=values["DEPLOYMENT_NAME"],
        )


def print_connection(connection: Connection) -> None:
    """Print non-secret connection metadata."""
    print(f"  Name: {connection.name}")
    print(f"  Type: {connection.type}")
    print(f"  Target: {connection.target}")
    print(f"  Default: {connection.is_default}")


def print_model_deployment(deployment: ModelDeployment) -> None:
    """Print typed model deployment metadata."""
    print(f"  Name: {deployment.name}")
    print(f"  Model publisher: {deployment.model_publisher}")
    print(f"  Model name: {deployment.model_name}")
    print(f"  Model version: {deployment.model_version}")


def inspect_project(client: AIProjectClient, settings: Settings) -> None:
    print("Project connections")
    connection_count = 0
    for connection in client.connections.list():
        connection_count += 1
        print_connection(connection)
        print()
    if connection_count == 0:
        print("  No connections found.\n")

    print(f"Requested connection: {settings.connection_name}")
    connection = client.connections.get(
        settings.connection_name,
        include_credentials=False,
    )
    print_connection(connection)

    print("\nProject model deployments")
    model_deployment_count = 0
    for deployment in client.deployments.list():
        if isinstance(deployment, ModelDeployment):
            model_deployment_count += 1
            print_model_deployment(deployment)
            print()
    if model_deployment_count == 0:
        print("  No model deployments found.\n")

    print(f"Requested model deployment: {settings.deployment_name}")
    deployment = client.deployments.get(settings.deployment_name)
    if not isinstance(deployment, ModelDeployment):
        raise TypeError(
            f"Deployment '{settings.deployment_name}' is not a model deployment "
            f"(type: {deployment.type})."
        )
    print_model_deployment(deployment)


def main() -> int:
    try:
        settings = Settings.from_environment()
        with DefaultAzureCredential() as credential:
            with AIProjectClient(
                endpoint=settings.project_endpoint,
                credential=credential,
            ) as client:
                inspect_project(client, settings)
    except (AzureError, TypeError, ValueError) as error:
        print(f"Error: {error}", file=sys.stderr)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
