"""Detect the deployment environment from well-known environment variables."""

from __future__ import annotations

import os
from enum import Enum
from typing import Mapping


class DeploymentEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


_CI_MARKERS = (
    "CI",
    "TF_BUILD",
    "GITHUB_ACTIONS",
    "GITHUB_WORKSPACE",
    "BUILD_BUILDID",
    "BUILD_SOURCESDIRECTORY",
    "SYSTEM_TEAMPROJECT",
    "JENKINS_URL",
    "GITLAB_CI",
)

_MANAGED_IDENTITY_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
)


def _is_truthy(value: str | None) -> bool:
    return bool(value and value.strip().lower() not in {"0", "false", "no", "off"})


def detect_environment(
    environment: Mapping[str, str] | None = None,
) -> DeploymentEnvironment:
    """Classify the current process as dev, CI, or production.

    APP_ENV and AZURE_DEPLOYMENT_ENVIRONMENT can explicitly override detection.
    CI markers take precedence over identity markers because hosted pipeline agents
    can themselves run on Azure resources with managed identity available.
    """

    values = os.environ if environment is None else environment

    override = (
        values.get("AZURE_DEPLOYMENT_ENVIRONMENT") or values.get("APP_ENV") or ""
    ).strip().lower()
    aliases = {
        "dev": DeploymentEnvironment.DEV,
        "development": DeploymentEnvironment.DEV,
        "local": DeploymentEnvironment.DEV,
        "ci": DeploymentEnvironment.CI,
        "pipeline": DeploymentEnvironment.CI,
        "prod": DeploymentEnvironment.PRODUCTION,
        "production": DeploymentEnvironment.PRODUCTION,
    }
    if override:
        try:
            return aliases[override]
        except KeyError as error:
            allowed = ", ".join(sorted(aliases))
            raise ValueError(
                f"Unsupported environment override {override!r}; use one of: {allowed}"
            ) from error

    if any(_is_truthy(values.get(name)) for name in _CI_MARKERS):
        return DeploymentEnvironment.CI

    has_workload_identity = all(
        values.get(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )
    if has_workload_identity or any(
        values.get(name) for name in _MANAGED_IDENTITY_MARKERS
    ):
        return DeploymentEnvironment.PRODUCTION

    return DeploymentEnvironment.DEV
