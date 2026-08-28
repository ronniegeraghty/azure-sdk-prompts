"""Detect whether the process is running locally, in CI, or in production."""

from __future__ import annotations

import os
from enum import Enum
from typing import Mapping


class RuntimeEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


CI_MARKERS = (
    "CI",
    "GITHUB_ACTIONS",
    "GITLAB_CI",
    "TF_BUILD",
    "BUILD_BUILDID",
    "PIPELINE_WORKSPACE",
    "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
    "JENKINS_URL",
)

PRODUCTION_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
    "IDENTITY_HEADER",
    "MSI_SECRET",
    "WEBSITE_INSTANCE_ID",
    "CONTAINER_APP_NAME",
    "KUBERNETES_SERVICE_HOST",
    "AZURE_FEDERATED_TOKEN_FILE",
)


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> RuntimeEnvironment:
    """Classify the runtime using environment markers, with CI taking precedence."""
    values = os.environ if environ is None else environ

    if _has_marker(values, CI_MARKERS):
        return RuntimeEnvironment.CI
    if _has_marker(values, PRODUCTION_MARKERS):
        return RuntimeEnvironment.PRODUCTION
    return RuntimeEnvironment.DEV


def _has_marker(values: Mapping[str, str], names: tuple[str, ...]) -> bool:
    return any(values.get(name, "").strip() for name in names)
