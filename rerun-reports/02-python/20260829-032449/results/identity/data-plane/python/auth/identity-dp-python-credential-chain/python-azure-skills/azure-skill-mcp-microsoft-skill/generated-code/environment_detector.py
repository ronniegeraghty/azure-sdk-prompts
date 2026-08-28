"""Detect whether the process is running locally, in CI, or in production."""

from __future__ import annotations

import os
from dataclasses import dataclass
from enum import Enum
from typing import Mapping


class RuntimeEnvironment(str, Enum):
    DEV = "dev"
    CI = "ci"
    PRODUCTION = "production"


@dataclass(frozen=True)
class DetectionResult:
    environment: RuntimeEnvironment
    reason: str


_CI_MARKERS = (
    "CI",
    "TF_BUILD",
    "GITHUB_ACTIONS",
    "GITHUB_WORKSPACE",
    "BUILD_BUILDID",
    "BUILD_SOURCESDIRECTORY",
    "JENKINS_URL",
    "GITLAB_CI",
)

_MANAGED_IDENTITY_ENDPOINT_MARKERS = (
    "IDENTITY_ENDPOINT",
    "MSI_ENDPOINT",
    "IMDS_ENDPOINT",
)

_AZURE_HOST_MARKERS = (
    "WEBSITE_INSTANCE_ID",
    "FUNCTIONS_WORKER_RUNTIME",
    "CONTAINER_APP_NAME",
)


def detect_environment(
    environ: Mapping[str, str] | None = None,
) -> DetectionResult:
    """Classify the runtime using an override and well-known host markers."""
    values = os.environ if environ is None else environ

    override = values.get("AZURE_AUTH_ENVIRONMENT", "").strip().lower()
    if override:
        try:
            environment = RuntimeEnvironment(override)
        except ValueError as error:
            allowed = ", ".join(item.value for item in RuntimeEnvironment)
            raise ValueError(
                f"AZURE_AUTH_ENVIRONMENT must be one of: {allowed}"
            ) from error
        return DetectionResult(environment, "AZURE_AUTH_ENVIRONMENT override")

    marker = _first_present(values, _CI_MARKERS)
    if marker:
        return DetectionResult(RuntimeEnvironment.CI, f"CI marker {marker}")

    marker = _first_present(values, _MANAGED_IDENTITY_ENDPOINT_MARKERS)
    if marker:
        return DetectionResult(
            RuntimeEnvironment.PRODUCTION,
            f"managed identity endpoint marker {marker}",
        )

    if _workload_identity_is_configured(values):
        return DetectionResult(
            RuntimeEnvironment.PRODUCTION,
            "Kubernetes workload identity variables",
        )

    marker = _first_present(values, _AZURE_HOST_MARKERS)
    if marker:
        return DetectionResult(
            RuntimeEnvironment.PRODUCTION, f"Azure host marker {marker}"
        )

    return DetectionResult(RuntimeEnvironment.DEV, "no CI or Azure host markers")


def _first_present(values: Mapping[str, str], names: tuple[str, ...]) -> str | None:
    return next((name for name in names if values.get(name)), None)


def _workload_identity_is_configured(values: Mapping[str, str]) -> bool:
    return all(
        values.get(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )
