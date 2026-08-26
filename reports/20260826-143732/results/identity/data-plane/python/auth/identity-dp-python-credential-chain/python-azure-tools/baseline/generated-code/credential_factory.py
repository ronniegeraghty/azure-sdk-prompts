"""Azure credential chains tailored to local, CI, and production environments."""

from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Any, Mapping

from azure.identity import (
    AzureCliCredential,
    AzureDeveloperCliCredential,
    AzurePipelinesCredential,
    ChainedTokenCredential,
    EnvironmentCredential,
    ManagedIdentityCredential,
    VisualStudioCodeCredential,
    WorkloadIdentityCredential,
)
from azure.identity.aio import (
    AzureCliCredential as AsyncAzureCliCredential,
    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
    ChainedTokenCredential as AsyncChainedTokenCredential,
    EnvironmentCredential as AsyncEnvironmentCredential,
    ManagedIdentityCredential as AsyncManagedIdentityCredential,
    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
)
from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential

from environment_detector import DeploymentEnvironment


@dataclass(frozen=True)
class CredentialSelection:
    """A credential and a human-readable description of its strategy."""

    credential: TokenCredential | AsyncTokenCredential
    strategy: str
    enable_cae: bool


def _optional_environment_value(
    environment: Mapping[str, str], name: str
) -> str | None:
    value = environment.get(name)
    return value if value and value.strip() else None


def _managed_identity_options(environment: Mapping[str, str]) -> dict[str, str]:
    client_id = _optional_environment_value(
        environment, "AZURE_MANAGED_IDENTITY_CLIENT_ID"
    )
    return {"client_id": client_id} if client_id else {}


def _pipelines_options(environment: Mapping[str, str]) -> dict[str, str] | None:
    values = {
        "tenant_id": _optional_environment_value(environment, "AZURE_TENANT_ID"),
        "client_id": _optional_environment_value(environment, "AZURE_CLIENT_ID"),
        "service_connection_id": _optional_environment_value(
            environment, "AZURE_SERVICE_CONNECTION_ID"
        ),
        "system_access_token": _optional_environment_value(
            environment, "SYSTEM_ACCESSTOKEN"
        ),
    }
    if all(values.values()):
        return {name: value for name, value in values.items() if value is not None}
    return None


def _workload_identity_is_configured(environment: Mapping[str, str]) -> bool:
    return all(
        _optional_environment_value(environment, name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )


def _ci_credentials(
    environment: Mapping[str, str], *, asynchronous: bool
) -> tuple[list[Any], str]:
    environment_type = (
        AsyncEnvironmentCredential if asynchronous else EnvironmentCredential
    )
    credentials: list[Any] = [environment_type()]
    strategies = ["environment credential"]

    pipeline_options = _pipelines_options(environment)
    if pipeline_options:
        pipeline_type = (
            AsyncAzurePipelinesCredential
            if asynchronous
            else AzurePipelinesCredential
        )
        credentials.append(pipeline_type(**pipeline_options))
        strategies.append("Azure Pipelines service connection")

    return credentials, "CI chain: " + " -> ".join(strategies)


def _production_credentials(
    environment: Mapping[str, str], *, asynchronous: bool
) -> tuple[list[Any], str]:
    managed_identity_type = (
        AsyncManagedIdentityCredential if asynchronous else ManagedIdentityCredential
    )
    credentials: list[Any] = [
        managed_identity_type(**_managed_identity_options(environment))
    ]
    identity_kind = (
        "user-assigned managed identity"
        if _optional_environment_value(
            environment, "AZURE_MANAGED_IDENTITY_CLIENT_ID"
        )
        else "system-assigned managed identity"
    )
    strategies = [identity_kind]

    if _workload_identity_is_configured(environment):
        workload_type = (
            AsyncWorkloadIdentityCredential
            if asynchronous
            else WorkloadIdentityCredential
        )
        credentials.append(workload_type())
        strategies.append("workload identity")

    return credentials, "production chain: " + " -> ".join(strategies)


def _build(
    deployment_environment: DeploymentEnvironment,
    enable_cae: bool,
    environment: Mapping[str, str],
    *,
    asynchronous: bool,
) -> CredentialSelection:
    if deployment_environment is DeploymentEnvironment.DEV:
        credential_types = (
            (
                AsyncAzureCliCredential,
                AsyncAzureDeveloperCliCredential,
                AsyncVisualStudioCodeCredential,
                AsyncChainedTokenCredential,
            )
            if asynchronous
            else (
                AzureCliCredential,
                AzureDeveloperCliCredential,
                VisualStudioCodeCredential,
                ChainedTokenCredential,
            )
        )
        cli_type, developer_cli_type, vscode_type, chain_type = credential_types
        credentials = [cli_type(), developer_cli_type(), vscode_type()]
        strategy = "developer tools chain: Azure CLI -> Azure Developer CLI -> VS Code"
    elif deployment_environment is DeploymentEnvironment.CI:
        credentials, strategy = _ci_credentials(
            environment, asynchronous=asynchronous
        )
        chain_type = (
            AsyncChainedTokenCredential if asynchronous else ChainedTokenCredential
        )
    else:
        credentials, strategy = _production_credentials(
            environment, asynchronous=asynchronous
        )
        chain_type = (
            AsyncChainedTokenCredential if asynchronous else ChainedTokenCredential
        )

    return CredentialSelection(
        credential=chain_type(*credentials),
        strategy=strategy,
        enable_cae=enable_cae,
    )


def build_credential(
    deployment_environment: DeploymentEnvironment,
    *,
    enable_cae: bool = False,
    environment: Mapping[str, str] | None = None,
) -> CredentialSelection:
    """Build a synchronous credential chain for a deployment environment."""

    return _build(
        deployment_environment,
        enable_cae,
        os.environ if environment is None else environment,
        asynchronous=False,
    )


def build_async_credential(
    deployment_environment: DeploymentEnvironment,
    *,
    enable_cae: bool = False,
    environment: Mapping[str, str] | None = None,
) -> CredentialSelection:
    """Build an asynchronous credential chain for a deployment environment."""

    return _build(
        deployment_environment,
        enable_cae,
        os.environ if environment is None else environment,
        asynchronous=True,
    )
