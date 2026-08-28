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
class SyncCredentialSelection:
    credential: TokenCredential
    strategy: str


@dataclass(frozen=True)
class AsyncCredentialSelection:
    credential: AsyncTokenCredential
    strategy: str


def _pipeline_settings(
    environ: Mapping[str, str],
) -> tuple[str, str, str, str] | None:
    names = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_SERVICE_CONNECTION_ID",
        "SYSTEM_ACCESSTOKEN",
    )
    values = tuple(environ.get(name, "").strip() for name in names)
    if all(values):
        return values[0], values[1], values[2], values[3]
    return None


def _workload_identity_settings(
    environ: Mapping[str, str],
) -> tuple[str, str, str] | None:
    names = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_FEDERATED_TOKEN_FILE",
    )
    values = tuple(environ.get(name, "").strip() for name in names)
    if all(values):
        return values[0], values[1], values[2]
    return None


def _managed_identity_client_id(environ: Mapping[str, str]) -> str | None:
    return (
        environ.get("AZURE_MANAGED_IDENTITY_CLIENT_ID", "").strip()
        or environ.get("AZURE_CLIENT_ID", "").strip()
        or None
    )


def build_credential(
    environment: DeploymentEnvironment,
    environ: Mapping[str, str] | None = None,
) -> SyncCredentialSelection:
    """Build the synchronous credential selected for an environment."""

    values = os.environ if environ is None else environ

    if environment is DeploymentEnvironment.DEV:
        credential = ChainedTokenCredential(
            AzureCliCredential(),
            AzureDeveloperCliCredential(),
            AzurePowerShellCredential(),
            VisualStudioCodeCredential(),
        )
        return SyncCredentialSelection(
            credential,
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code",
        )

    if environment is DeploymentEnvironment.CI:
        credentials: list[TokenCredential] = [EnvironmentCredential()]
        strategy = "pipeline environment variables"
        pipeline = _pipeline_settings(values)
        if pipeline:
            tenant_id, client_id, connection_id, access_token = pipeline
            credentials.append(
                AzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=connection_id,
                    system_access_token=access_token,
                )
            )
            strategy += " -> Azure Pipelines service connection"
        else:
            strategy += (
                " (Azure Pipelines fallback not configured; requires "
                "AZURE_SERVICE_CONNECTION_ID and SYSTEM_ACCESSTOKEN)"
            )
        return SyncCredentialSelection(
            ChainedTokenCredential(*credentials),
            strategy,
        )

    client_id = _managed_identity_client_id(values)
    credentials = [ManagedIdentityCredential(client_id=client_id)]
    identity_kind = "user-assigned" if client_id else "system-assigned"
    strategy = f"{identity_kind} managed identity"
    workload = _workload_identity_settings(values)
    if workload:
        tenant_id, workload_client_id, token_file = workload
        credentials.append(
            WorkloadIdentityCredential(
                tenant_id=tenant_id,
                client_id=workload_client_id,
                token_file_path=token_file,
            )
        )
        strategy += " -> Kubernetes workload identity"
    else:
        strategy += " (workload identity fallback not configured)"
    return SyncCredentialSelection(
        ChainedTokenCredential(*credentials),
        strategy,
    )


def build_async_credential(
    environment: DeploymentEnvironment,
    environ: Mapping[str, str] | None = None,
) -> AsyncCredentialSelection:
    """Build the asynchronous credential selected for an environment."""

    values = os.environ if environ is None else environ

    if environment is DeploymentEnvironment.DEV:
        credential = AsyncChainedTokenCredential(
            AsyncAzureCliCredential(),
            AsyncAzureDeveloperCliCredential(),
            AsyncAzurePowerShellCredential(),
            AsyncVisualStudioCodeCredential(),
        )
        return AsyncCredentialSelection(
            credential,
            "developer tools: Azure CLI -> Azure Developer CLI -> "
            "Azure PowerShell -> VS Code",
        )

    if environment is DeploymentEnvironment.CI:
        credentials: list[AsyncTokenCredential] = [
            AsyncEnvironmentCredential()
        ]
        strategy = "pipeline environment variables"
        pipeline = _pipeline_settings(values)
        if pipeline:
            tenant_id, client_id, connection_id, access_token = pipeline
            credentials.append(
                AsyncAzurePipelinesCredential(
                    tenant_id=tenant_id,
                    client_id=client_id,
                    service_connection_id=connection_id,
                    system_access_token=access_token,
                )
            )
            strategy += " -> Azure Pipelines service connection"
        else:
            strategy += (
                " (Azure Pipelines fallback not configured; requires "
                "AZURE_SERVICE_CONNECTION_ID and SYSTEM_ACCESSTOKEN)"
            )
        return AsyncCredentialSelection(
            AsyncChainedTokenCredential(*credentials),
            strategy,
        )

    client_id = _managed_identity_client_id(values)
    credentials = [AsyncManagedIdentityCredential(client_id=client_id)]
    identity_kind = "user-assigned" if client_id else "system-assigned"
    strategy = f"{identity_kind} managed identity"
    workload = _workload_identity_settings(values)
    if workload:
        tenant_id, workload_client_id, token_file = workload
        credentials.append(
            AsyncWorkloadIdentityCredential(
                tenant_id=tenant_id,
                client_id=workload_client_id,
                token_file_path=token_file,
            )
        )
        strategy += " -> Kubernetes workload identity"
    else:
        strategy += " (workload identity fallback not configured)"
    return AsyncCredentialSelection(
        AsyncChainedTokenCredential(*credentials),
        strategy,
    )
