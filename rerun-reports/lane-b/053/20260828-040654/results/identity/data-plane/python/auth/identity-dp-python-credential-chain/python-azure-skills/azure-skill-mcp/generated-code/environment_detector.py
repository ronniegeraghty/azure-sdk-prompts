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
    evidence: tuple[str, ...]


_CI_MARKERS = (
    "TF_BUILD",
    "GITHUB_ACTIONS",
    "GITLAB_CI",
    "JENKINS_URL",
    "BUILD_BUILDID",
    "SYSTEM_TEAMPROJECTID",
    "PIPELINE_WORKSPACE",
    "RUNNER_WORKSPACE",
)

_PRODUCTION_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
    "WEBSITE_INSTANCE_ID",
    "CONTAINER_APP_NAME",
    "AZURE_FEDERATED_TOKEN_FILE",
    "KUBERNETES_SERVICE_HOST",
)


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> DetectionResult:
    """Classify the process as dev, CI, or production.

    AZURE_CREDENTIAL_ENVIRONMENT can explicitly select dev, ci, or production.
    CI markers take precedence over hosting markers because build agents can run
    on Azure infrastructure that also exposes managed identity endpoints.
    """

    values = os.environ if environ is None else environ
    override = values.get("AZURE_CREDENTIAL_ENVIRONMENT", "").strip().lower()
    if override:
        try:
            selected = DeploymentEnvironment(override)
        except ValueError as error:
            allowed = ", ".join(item.value for item in DeploymentEnvironment)
            raise ValueError(
                "AZURE_CREDENTIAL_ENVIRONMENT must be one of: "
                f"{allowed}; received {override!r}"
            ) from error
        return DetectionResult(selected, ("AZURE_CREDENTIAL_ENVIRONMENT",))

    ci_evidence = tuple(name for name in _CI_MARKERS if values.get(name))
    if values.get("CI", "").strip().lower() in {"1", "true", "yes"}:
        ci_evidence = ("CI",) + ci_evidence
    if ci_evidence:
        return DetectionResult(DeploymentEnvironment.CI, ci_evidence)

    production_evidence = tuple(
        name for name in _PRODUCTION_MARKERS if values.get(name)
    )
    if production_evidence:
        return DetectionResult(
            DeploymentEnvironment.PRODUCTION,
            production_evidence,
        )

    return DetectionResult(
        DeploymentEnvironment.DEV,
        ("no CI or Azure hosting markers found",),
    )
