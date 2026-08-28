"""Detect the deployment environment from well-known environment variables."""

from __future__ import annotations

import os
from enum import Enum
from typing import Mapping


class RuntimeEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


_CI_MARKERS = (
    "CI",
    "TF_BUILD",
    "BUILD_BUILDID",
    "BUILD_SOURCESDIRECTORY",
    "SYSTEM_TEAMPROJECT",
    "GITHUB_ACTIONS",
    "GITLAB_CI",
    "JENKINS_URL",
)

_AZURE_HOST_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
    "WEBSITE_INSTANCE_ID",
    "CONTAINER_APP_NAME",
)


def _is_truthy(value: str | None) -> bool:
    return bool(value and value.strip().lower() not in {"0", "false", "no", "off"})


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> RuntimeEnvironment:
    """Classify the runtime as dev, CI, or production.

    APP_ENV can explicitly override detection. CI markers take precedence over
    Azure-host markers because hosted pipeline agents can expose both.
    """
    values = os.environ if environ is None else environ
    override = values.get("APP_ENV", "").strip().lower()
    aliases = {
        "dev": RuntimeEnvironment.DEV,
        "development": RuntimeEnvironment.DEV,
        "local": RuntimeEnvironment.DEV,
        "ci": RuntimeEnvironment.CI,
        "pipeline": RuntimeEnvironment.CI,
        "prod": RuntimeEnvironment.PRODUCTION,
        "production": RuntimeEnvironment.PRODUCTION,
    }
    if override:
        try:
            return aliases[override]
        except KeyError as error:
            valid = ", ".join(sorted(aliases))
            raise ValueError(f"APP_ENV must be one of: {valid}") from error

    if any(_is_truthy(values.get(name)) for name in _CI_MARKERS):
        return RuntimeEnvironment.CI

    has_workload_identity = all(
        values.get(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )
    has_azure_host = any(values.get(name) for name in _AZURE_HOST_MARKERS)
    has_kubernetes_host = bool(values.get("KUBERNETES_SERVICE_HOST"))
    if has_azure_host or has_workload_identity or has_kubernetes_host:
        return RuntimeEnvironment.PRODUCTION

    return RuntimeEnvironment.DEV
