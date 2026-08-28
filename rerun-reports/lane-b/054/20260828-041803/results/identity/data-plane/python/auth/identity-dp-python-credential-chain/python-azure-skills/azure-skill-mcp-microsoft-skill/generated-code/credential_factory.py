"""Build Azure credential chains tailored to the detected runtime environment."""

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
    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
    ChainedTokenCredential as AsyncChainedTokenCredential,
    EnvironmentCredential as AsyncEnvironmentCredential,
    ManagedIdentityCredential as AsyncManagedIdentityCredential,
    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,
)

from environment_detector import RuntimeEnvironment


@dataclass(frozen=True)
class SyncCredentialSelection:
    credential: TokenCredential
    strategy: str
    enable_cae: bool


@dataclass(frozen=True)
class AsyncCredentialSelection:
    credential: AsyncTokenCredential
    strategy: str
    enable_cae: bool


def build_sync_credential(
    environment: RuntimeEnvironment,
    *,
    enable_cae: bool = False,
    environ: Mapping[str, str] | None = None,
) -> SyncCredentialSelection:
    """Create a synchronous credential chain for one runtime environment."""
    values = os.environ if environ is None else environ

    if environment is RuntimeEnvironment.DEV:
        credential = ChainedTokenCredential(
            VisualStudioCodeCredential(),
            AzureCliCredential(),
            AzurePowerShellCredential(),
            AzureDeveloperCliCredential(),
        )
        strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
    elif environment is RuntimeEnvironment.CI:
        credential, strategy = _build_sync_ci_credential(values)
    else:
        credential, strategy = _build_sync_production_credential(values)

    return SyncCredentialSelection(credential, strategy, enable_cae)


def build_async_credential(
    environment: RuntimeEnvironment,
    *,
    enable_cae: bool = False,
    environ: Mapping[str, str] | None = None,
) -> AsyncCredentialSelection:
    """Create an asynchronous credential chain for one runtime environment."""
    values = os.environ if environ is None else environ

    if environment is RuntimeEnvironment.DEV:
        credential = AsyncChainedTokenCredential(
            AsyncVisualStudioCodeCredential(),
            AsyncAzureCliCredential(),
            AsyncAzurePowerShellCredential(),
            AsyncAzureDeveloperCliCredential(),
        )
        strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
    elif environment is RuntimeEnvironment.CI:
        credential, strategy = _build_async_ci_credential(values)
    else:
        credential, strategy = _build_async_production_credential(values)

    return AsyncCredentialSelection(credential, strategy, enable_cae)


def _pipeline_settings(values: Mapping[str, str]) -> dict[str, str] | None:
    names = {
        "tenant_id": ("AZURESUBSCRIPTION_TENANT_ID", "AZURE_TENANT_ID"),
        "client_id": ("AZURESUBSCRIPTION_CLIENT_ID", "AZURE_CLIENT_ID"),
        "service_connection_id": (
            "AZURESUBSCRIPTION_SERVICE_CONNECTION_ID",
            "AZURE_PIPELINES_SERVICE_CONNECTION_ID",
        ),
        "system_access_token": ("SYSTEM_ACCESSTOKEN",),
    }
    present = {
        key: next(
            (
                values.get(name, "").strip()
                for name in aliases
                if values.get(name, "").strip()
            ),
            "",
        )
        for key, aliases in names.items()
    }
    pipeline_specific = (
        present["service_connection_id"],
        present["system_access_token"],
    )
    if not any(pipeline_specific):
        return None

    missing = ["/".join(names[key]) for key, value in present.items() if not value]
    if missing:
        raise ValueError(
            "Azure Pipelines workload identity configuration is incomplete; "
            f"missing: {', '.join(missing)}"
        )
    return present


def _build_sync_ci_credential(
    values: Mapping[str, str],
) -> tuple[TokenCredential, str]:
    settings = _pipeline_settings(values)
    if settings:
        return (
            ChainedTokenCredential(
                AzurePipelinesCredential(**settings),
                EnvironmentCredential(),
            ),
            "CI: Azure Pipelines workload identity -> environment credential",
        )
    return EnvironmentCredential(), "CI: environment credential (secret or certificate)"


def _build_async_ci_credential(
    values: Mapping[str, str],
) -> tuple[AsyncTokenCredential, str]:
    settings = _pipeline_settings(values)
    if settings:
        return (
            AsyncChainedTokenCredential(
                AsyncAzurePipelinesCredential(**settings),
                AsyncEnvironmentCredential(),
            ),
            "CI: Azure Pipelines workload identity -> environment credential",
        )
    return (
        AsyncEnvironmentCredential(),
        "CI: environment credential (secret or certificate)",
    )


def _workload_identity_is_configured(values: Mapping[str, str]) -> bool:
    required = (
        "AZURE_TENANT_ID",
        "AZURE_CLIENT_ID",
        "AZURE_FEDERATED_TOKEN_FILE",
    )
    return all(values.get(name, "").strip() for name in required)


def _managed_identity_client_id(values: Mapping[str, str]) -> str | None:
    return (
        values.get("AZURE_MANAGED_IDENTITY_CLIENT_ID", "").strip()
        or values.get("AZURE_CLIENT_ID", "").strip()
        or None
    )


def _build_sync_production_credential(
    values: Mapping[str, str],
) -> tuple[TokenCredential, str]:
    client_id = _managed_identity_client_id(values)
    credentials: list[TokenCredential] = [
        ManagedIdentityCredential(client_id=client_id)
    ]
    strategy = (
        "production: user-assigned managed identity"
        if client_id
        else "production: system-assigned managed identity"
    )
    if _workload_identity_is_configured(values):
        credentials.append(
            WorkloadIdentityCredential(
                tenant_id=values["AZURE_TENANT_ID"],
                client_id=values["AZURE_CLIENT_ID"],
                token_file_path=values["AZURE_FEDERATED_TOKEN_FILE"],
            )
        )
        strategy += " -> Kubernetes workload identity"
    return ChainedTokenCredential(*credentials), strategy


def _build_async_production_credential(
    values: Mapping[str, str],
) -> tuple[AsyncTokenCredential, str]:
    client_id = _managed_identity_client_id(values)
    credentials: list[AsyncTokenCredential] = [
        AsyncManagedIdentityCredential(client_id=client_id)
    ]
    strategy = (
        "production: user-assigned managed identity"
        if client_id
        else "production: system-assigned managed identity"
    )
    if _workload_identity_is_configured(values):
        credentials.append(
            AsyncWorkloadIdentityCredential(
                tenant_id=values["AZURE_TENANT_ID"],
                client_id=values["AZURE_CLIENT_ID"],
                token_file_path=values["AZURE_FEDERATED_TOKEN_FILE"],
            )
        )
        strategy += " -> Kubernetes workload identity"
    return AsyncChainedTokenCredential(*credentials), strategy
