"""Detect the deployment environment from well-known host variables."""

from __future__ import annotations

import os
from enum import Enum
from typing import Mapping


class RuntimeEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


CI_MARKERS = (
    "TF_BUILD",
    "BUILD_BUILDID",
    "PIPELINE_WORKSPACE",
    "GITHUB_ACTIONS",
    "GITHUB_WORKSPACE",
    "GITLAB_CI",
    "CI_PROJECT_DIR",
    "JENKINS_URL",
    "CI",
)

PRODUCTION_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
    "AZURE_FEDERATED_TOKEN_FILE",
    "WEBSITE_INSTANCE_ID",
    "KUBERNETES_SERVICE_HOST",
    "CONTAINER_APP_NAME",
)


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> RuntimeEnvironment:
    """Classify the current host as CI, production, or local development."""

    variables = os.environ if environ is None else environ
    if any(variables.get(name) for name in CI_MARKERS):
        return RuntimeEnvironment.CI
    if any(variables.get(name) for name in PRODUCTION_MARKERS):
        return RuntimeEnvironment.PRODUCTION
    return RuntimeEnvironment.DEV
