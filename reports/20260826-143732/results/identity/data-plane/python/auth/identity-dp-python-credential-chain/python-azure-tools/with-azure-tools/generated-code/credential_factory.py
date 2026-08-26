"""Build explicit Azure credential chains for each deployment environment."""

from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Mapping

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.identity import (
    AzureCliCredential,
    AzureDeveloperCliCredential,
    AzurePipelinesCredential,
    AzurePowerShellCredential,
    ChainedTokenCredential,
    EnvironmentCredential,
    ManagedIdentityCredential,
    VisualStudioCodeCredential,
    WorkloadIdentityCredential,
)
from azure.identity.aio import (
    AzureCliCredential as AsyncAzureCliCredential,
)
from azure.identity.aio import (
    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
)
from azure.identity.aio import (
    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
)
from azure.identity.aio import (
    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
)
from azure.identity.aio import (
    ChainedTokenCredential as AsyncChainedTokenCredential,
)
from azure.identity.aio import (
    EnvironmentCredential as AsyncEnvironmentCredential,
)
from azure.identity.aio import (
    ManagedIdentityCredential as AsyncManagedIdentityCredential,
)
from azure.identity.aio import (
    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
)
from azure.identity.aio import (
    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
)

from environment_detector import DeploymentEnvironment


@dataclass(frozen=True)
class CredentialBundle:
    credential: TokenCredential
    strategy: str


@dataclass(frozen=True)
class AsyncCredentialBundle:
    credential: AsyncTokenCredential
    strategy: str


def _azure_pipelines_configuration(
    environ: Mapping[str, str],
) -> tuple[str, str, str, str] | None:
    names = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_SERVICE_CONNECTION_ID",
        "SYSTEM_ACCESSTOKEN",
    )
    if not all(environ.get(name) for name in names):
        return None
    return (
        environ["AZURE_TENANT_ID"],
        environ["AZURE_CLIENT_ID"],
        environ["AZURE_SERVICE_CONNECTION_ID"],
        environ["SYSTEM_ACCESSTOKEN"],
    )


def _has_workload_identity_configuration(environ: Mapping[str, str]) -> bool:
    return all(
        environ.get(name)
        for name in (
            "AZURE_TENANT_ID",
            "AZURE_CLIENT_ID",
            "AZURE_FEDERATED_TOKEN_FILE",
        )
    )


def _managed_identity_client_id(environ: Mapping[str, str]) -> str | None:
    return environ.get("AZURE_MANAGED_IDENTITY_CLIENT_ID") or None


def build_credential(
    environment: DeploymentEnvironment,
    environ: Mapping[str, str] | None = None,
) -> CredentialBundle:
    """Create a synchronous credential and describe its ordered strategy."""
    values = os.environ if environ is None else environ

    if environment is DeploymentEnvironment.DEV:
        credential = ChainedTokenCredential(
            AzureCliCredential(),
            AzureDeveloperCliCredential(),
            AzurePowerShellCredential(),
            VisualStudioCodeCredential(),
        )
        return CredentialBundle(
            credential,
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code",
        )

    if environment is DeploymentEnvironment.CI:
        credentials: list[TokenCredential] = []
        strategy: list[str] = []
        pipeline_config = _azure_pipelines_configuration(values)
        if pipeline_config:
            tenant_id, client_id, service_connection_id, system_access_token = (
                pipeline_config
            )
            credentials.append(
                AzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=service_connection_id,
                    system_access_token=system_access_token,
                )
            )
            strategy.append("Azure Pipelines workload identity service connection")

        credentials.append(EnvironmentCredential())
        strategy.append("pipeline environment credential (secret or certificate)")
        return CredentialBundle(
            ChainedTokenCredential(*credentials),
            " -> ".join(strategy),
        )

    managed_identity_client_id = _managed_identity_client_id(values)
    credentials = [
        ManagedIdentityCredential(client_id=managed_identity_client_id)
    ]
    strategy = [
        "user-assigned managed identity"
        if managed_identity_client_id
        else "system-assigned managed identity"
    ]
    if _has_workload_identity_configuration(values):
        credentials.append(
            WorkloadIdentityCredential(
                tenant_id=values["AZURE_TENANT_ID"],
                client_id=values["AZURE_CLIENT_ID"],
                token_file_path=values["AZURE_FEDERATED_TOKEN_FILE"],
            )
        )
        strategy.append("Kubernetes workload identity fallback")

    return CredentialBundle(
        ChainedTokenCredential(*credentials),
        " -> ".join(strategy),
    )


def build_async_credential(
    environment: DeploymentEnvironment,
    environ: Mapping[str, str] | None = None,
) -> AsyncCredentialBundle:
    """Create an asynchronous credential with the same environment strategy."""
    values = os.environ if environ is None else environ

    if environment is DeploymentEnvironment.DEV:
        credential = AsyncChainedTokenCredential(
            AsyncAzureCliCredential(),
            AsyncAzureDeveloperCliCredential(),
            AsyncAzurePowerShellCredential(),
            AsyncVisualStudioCodeCredential(),
        )
        return AsyncCredentialBundle(
            credential,
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code",
        )

    if environment is DeploymentEnvironment.CI:
        credentials: list[AsyncTokenCredential] = []
        strategy: list[str] = []
        pipeline_config = _azure_pipelines_configuration(values)
        if pipeline_config:
            tenant_id, client_id, service_connection_id, system_access_token = (
                pipeline_config
            )
            credentials.append(
                AsyncAzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=service_connection_id,
                    system_access_token=system_access_token,
                )
            )
            strategy.append("Azure Pipelines workload identity service connection")

        credentials.append(AsyncEnvironmentCredential())
        strategy.append("pipeline environment credential (secret or certificate)")
        return AsyncCredentialBundle(
            AsyncChainedTokenCredential(*credentials),
            " -> ".join(strategy),
        )

    managed_identity_client_id = _managed_identity_client_id(values)
    credentials = [
        AsyncManagedIdentityCredential(client_id=managed_identity_client_id)
    ]
    strategy = [
        "user-assigned managed identity"
        if managed_identity_client_id
        else "system-assigned managed identity"
    ]
    if _has_workload_identity_configuration(values):
        credentials.append(
            AsyncWorkloadIdentityCredential(
                tenant_id=values["AZURE_TENANT_ID"],
                client_id=values["AZURE_CLIENT_ID"],
                token_file_path=values["AZURE_FEDERATED_TOKEN_FILE"],
            )
        )
        strategy.append("Kubernetes workload identity fallback")

    return AsyncCredentialBundle(
        AsyncChainedTokenCredential(*credentials),
        " -> ".join(strategy),
    )
