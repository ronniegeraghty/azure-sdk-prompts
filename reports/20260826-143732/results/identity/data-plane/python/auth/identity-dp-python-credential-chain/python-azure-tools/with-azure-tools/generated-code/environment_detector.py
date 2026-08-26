"""Detect the deployment environment from well-known environment variables."""

from __future__ import annotations

import os
from dataclasses import dataclass
from enum import Enum
from typing import Mapping


class DeploymentEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


@dataclass(frozen=True)
class DetectionResult:
    environment: DeploymentEnvironment
    reason: str


_CI_MARKERS = {
    "TF_BUILD": "Azure Pipelines",
    "GITHUB_ACTIONS": "GitHub Actions",
    "GITLAB_CI": "GitLab CI",
    "JENKINS_URL": "Jenkins",
    "TEAMCITY_VERSION": "TeamCity",
    "BITBUCKET_BUILD_NUMBER": "Bitbucket Pipelines",
    "CI_PROJECT_DIR": "CI project workspace",
    "PIPELINE_WORKSPACE": "pipeline workspace",
}

_PRODUCTION_MARKERS = {
    "IDENTITY_ENDPOINT": "Azure managed identity endpoint",
    "MSI_ENDPOINT": "Azure managed identity endpoint",
    "IMDS_ENDPOINT": "Azure Instance Metadata Service endpoint",
    "WEBSITE_INSTANCE_ID": "Azure App Service host",
    "FUNCTIONS_WORKER_RUNTIME": "Azure Functions host",
    "CONTAINER_APP_NAME": "Azure Container Apps host",
}


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> DetectionResult:
    """Classify the current process as development, CI, or production."""
    values = os.environ if environ is None else environ

    override = values.get("APP_ENVIRONMENT", "").strip().lower()
    if override:
        aliases = {
            "dev": DeploymentEnvironment.DEV,
            "development": DeploymentEnvironment.DEV,
            "ci": DeploymentEnvironment.CI,
            "pipeline": DeploymentEnvironment.CI,
            "prod": DeploymentEnvironment.PRODUCTION,
            "production": DeploymentEnvironment.PRODUCTION,
        }
        if override not in aliases:
            expected = ", ".join(sorted(aliases))
            raise ValueError(
                f"Invalid APP_ENVIRONMENT value {override!r}; expected one of: {expected}"
            )
        return DetectionResult(aliases[override], "APP_ENVIRONMENT override")

    for variable, description in _CI_MARKERS.items():
        if values.get(variable):
            return DetectionResult(
                DeploymentEnvironment.CI,
                f"{description} detected from {variable}",
            )

    workload_identity_variables = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_FEDERATED_TOKEN_FILE",
    )
    if all(values.get(variable) for variable in workload_identity_variables):
        return DetectionResult(
            DeploymentEnvironment.PRODUCTION,
            "Kubernetes workload identity configuration detected",
        )

    for variable, description in _PRODUCTION_MARKERS.items():
        if values.get(variable):
            return DetectionResult(
                DeploymentEnvironment.PRODUCTION,
                f"{description} detected from {variable}",
            )

    return DetectionResult(
        DeploymentEnvironment.DEV,
        "no CI or Azure-hosted production markers were found",
    )
